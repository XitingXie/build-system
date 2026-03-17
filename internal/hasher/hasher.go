// Package hasher computes SHA-256 content hashes for files and actions.
package hasher

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io"
	"os"
	"sort"
)

// FileDigest holds the path and SHA-256 digest of a file.
type FileDigest struct {
	Path   string
	Digest [32]byte
}

// HashFile computes the SHA-256 digest of a single file.
func HashFile(path string) ([32]byte, error) {
	f, err := os.Open(path)
	if err != nil {
		return [32]byte{}, fmt.Errorf("hash file %s: %w", path, err)
	}
	defer f.Close()

	h := sha256.New()
	if _, err := io.Copy(h, f); err != nil {
		return [32]byte{}, fmt.Errorf("hash file %s: %w", path, err)
	}
	var digest [32]byte
	copy(digest[:], h.Sum(nil))
	return digest, nil
}

// HashFiles computes digests for a list of file paths.
func HashFiles(paths []string) ([]FileDigest, error) {
	result := make([]FileDigest, 0, len(paths))
	for _, p := range paths {
		d, err := HashFile(p)
		if err != nil {
			return nil, err
		}
		result = append(result, FileDigest{Path: p, Digest: d})
	}
	return result, nil
}

// ActionKey computes a stable cache key for an action from:
//   - input file digests
//   - the command string
//   - sorted environment variables
//   - transitive dependency digests
func ActionKey(inputs []FileDigest, cmd string, env map[string]string, depDigests [][32]byte) [32]byte {
	h := sha256.New()

	// Hash each input file digest (sorted by path for stability).
	sorted := make([]FileDigest, len(inputs))
	copy(sorted, inputs)
	sort.Slice(sorted, func(i, j int) bool { return sorted[i].Path < sorted[j].Path })
	for _, fd := range sorted {
		h.Write([]byte(fd.Path))
		h.Write(fd.Digest[:])
	}

	// Hash the command.
	h.Write([]byte(cmd))

	// Hash env vars (sorted by key for stability).
	keys := make([]string, 0, len(env))
	for k := range env {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, k := range keys {
		h.Write([]byte(k + "=" + env[k]))
	}

	// Hash transitive dependency action digests.
	for _, d := range depDigests {
		h.Write(d[:])
	}

	var key [32]byte
	copy(key[:], h.Sum(nil))
	return key
}

// Hex returns the hex string representation of a digest.
func Hex(d [32]byte) string {
	return hex.EncodeToString(d[:])
}
