// Package cache implements a local content-addressable store (CAS) and action cache.
//
// Layout on disk (~/.cache/build-system/):
//
//	cas/
//	  <hex-digest>        ← raw file contents keyed by SHA-256
//	actions/
//	  <action-hex>        ← JSON-encoded ActionResult keyed by action digest
package cache

import (
	"crypto/sha256"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

// ActionResult records the outputs produced by a cached action.
type ActionResult struct {
	OutputDigests map[string][32]byte // output path → CAS digest
	ExitCode      int
}

// Cache is a local disk-backed CAS + action cache.
type Cache struct {
	casDir    string
	actionDir string
}

// New creates (or opens) a cache rooted at dir.
func New(dir string) (*Cache, error) {
	casDir := filepath.Join(dir, "cas")
	actionDir := filepath.Join(dir, "actions")
	for _, d := range []string{casDir, actionDir} {
		if err := os.MkdirAll(d, 0755); err != nil {
			return nil, fmt.Errorf("cache init: %w", err)
		}
	}
	return &Cache{casDir: casDir, actionDir: actionDir}, nil
}

// DefaultDir returns the default cache directory.
func DefaultDir() string {
	home, _ := os.UserHomeDir()
	return filepath.Join(home, ".cache", "build-system")
}

// --- CAS ---

// StoreFile copies a file into the CAS and returns its digest.
func (c *Cache) StoreFile(path string) ([32]byte, error) {
	f, err := os.Open(path)
	if err != nil {
		return [32]byte{}, fmt.Errorf("cas store: %w", err)
	}
	defer f.Close()

	h := sha256.New()
	tmp, err := os.CreateTemp(c.casDir, "tmp-*")
	if err != nil {
		return [32]byte{}, fmt.Errorf("cas store tmp: %w", err)
	}
	tmpPath := tmp.Name()

	if _, err := io.Copy(io.MultiWriter(h, tmp), f); err != nil {
		tmp.Close()
		os.Remove(tmpPath)
		return [32]byte{}, fmt.Errorf("cas store copy: %w", err)
	}
	tmp.Close()

	var digest [32]byte
	copy(digest[:], h.Sum(nil))
	dest := c.casPath(digest)
	if err := os.Rename(tmpPath, dest); err != nil {
		os.Remove(tmpPath)
		return [32]byte{}, fmt.Errorf("cas store rename: %w", err)
	}
	return digest, nil
}

// RetrieveFile copies a CAS object to dest path.
func (c *Cache) RetrieveFile(digest [32]byte, dest string) error {
	src := c.casPath(digest)
	if err := os.MkdirAll(filepath.Dir(dest), 0755); err != nil {
		return err
	}
	in, err := os.Open(src)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return fmt.Errorf("cas miss: %x", digest)
		}
		return err
	}
	defer in.Close()

	out, err := os.Create(dest)
	if err != nil {
		return err
	}
	defer out.Close()
	_, err = io.Copy(out, in)
	return err
}

// HasCAS reports whether the CAS contains the given digest.
func (c *Cache) HasCAS(digest [32]byte) bool {
	_, err := os.Stat(c.casPath(digest))
	return err == nil
}

func (c *Cache) casPath(digest [32]byte) string {
	return filepath.Join(c.casDir, fmt.Sprintf("%x", digest))
}

// --- Action cache ---

// StoreAction saves an ActionResult keyed by the action digest.
func (c *Cache) StoreAction(actionDigest [32]byte, result ActionResult) error {
	data, err := json.Marshal(result)
	if err != nil {
		return err
	}
	return os.WriteFile(c.actionPath(actionDigest), data, 0644)
}

// LookupAction retrieves an ActionResult. Returns (result, true) on hit.
func (c *Cache) LookupAction(actionDigest [32]byte) (ActionResult, bool) {
	data, err := os.ReadFile(c.actionPath(actionDigest))
	if err != nil {
		return ActionResult{}, false
	}
	var result ActionResult
	if err := json.Unmarshal(data, &result); err != nil {
		return ActionResult{}, false
	}
	return result, true
}

func (c *Cache) actionPath(digest [32]byte) string {
	return filepath.Join(c.actionDir, fmt.Sprintf("%x", digest))
}

// Clean removes all cached data.
func (c *Cache) Clean() error {
	dir := filepath.Dir(c.casDir)
	if err := os.RemoveAll(dir); err != nil {
		return err
	}
	return os.MkdirAll(dir, 0755)
}
