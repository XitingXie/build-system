// Package executor runs a build action and stores its outputs in the cache.
package executor

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/xitingxie/build-system/internal/cache"
	"github.com/xitingxie/build-system/internal/graph"
	"github.com/xitingxie/build-system/internal/hasher"
	"github.com/xitingxie/build-system/internal/parser"
)

// Executor runs actions and manages cache interactions.
type Executor struct {
	cache   *cache.Cache
	workDir string // root of the source tree
}

// New creates an Executor.
func New(c *cache.Cache, workDir string) *Executor {
	return &Executor{cache: c, workDir: workDir}
}

// Run executes the action for node n.
// It checks the cache first; on a miss it runs the command and stores outputs.
// depDigests are the action-key digests of transitive dependencies (for cache key).
func (e *Executor) Run(n *graph.Node, depDigests [][32]byte) error {
	t := n.Target

	// Resolve absolute source paths.
	srcPaths := make([]string, 0, len(t.Srcs))
	for _, s := range t.Srcs {
		srcPaths = append(srcPaths, filepath.Join(e.workDir, labelToDir(t.Label), s))
	}

	// Hash input files.
	inputs, err := hasher.HashFiles(srcPaths)
	if err != nil {
		return fmt.Errorf("hash inputs for %s: %w", t.Label, err)
	}

	// Compute action key.
	actionKey := hasher.ActionKey(inputs, t.Cmd, nil, depDigests)

	// Cache hit?
	if result, ok := e.cache.LookupAction(actionKey); ok {
		fmt.Printf("  [cached] %s\n", t.Label)
		return e.restoreOutputs(t.Label, result)
	}

	// Cache miss — execute the action.
	fmt.Printf("  [build]  %s\n", t.Label)
	outDir, err := e.runAction(t, srcPaths)
	if err != nil {
		return err
	}

	// Store outputs in CAS and record the action result.
	result := cache.ActionResult{
		OutputDigests: make(map[string][32]byte),
		ExitCode:      0,
	}
	for _, out := range t.Outs {
		outPath := filepath.Join(outDir, out)
		digest, err := e.cache.StoreFile(outPath)
		if err != nil {
			return fmt.Errorf("store output %s: %w", out, err)
		}
		result.OutputDigests[out] = digest
	}
	if err := e.cache.StoreAction(actionKey, result); err != nil {
		return fmt.Errorf("store action result for %s: %w", t.Label, err)
	}

	// Copy outputs to workspace.
	return e.restoreOutputs(t.Label, result)
}

// runAction executes the shell command for a target and returns the output directory.
func (e *Executor) runAction(t parser.Target, srcPaths []string) (string, error) {
	outDir := filepath.Join(e.workDir, "out", labelToDir(t.Label))
	if err := os.MkdirAll(outDir, 0755); err != nil {
		return "", err
	}

	// Expand $SRCS and $OUT in the command.
	outPath := filepath.Join(outDir, t.Outs[0])
	cmd := t.Cmd
	cmd = strings.ReplaceAll(cmd, "$SRCS", strings.Join(srcPaths, " "))
	cmd = strings.ReplaceAll(cmd, "$OUT", outPath)

	c := exec.Command("sh", "-c", cmd)
	c.Dir = filepath.Join(e.workDir, labelToDir(t.Label))
	c.Stdout = os.Stdout
	c.Stderr = os.Stderr
	if err := c.Run(); err != nil {
		return "", fmt.Errorf("action failed for %s: %w", t.Label, err)
	}
	return outDir, nil
}

// restoreOutputs copies CAS objects to the workspace output directory.
func (e *Executor) restoreOutputs(label string, result cache.ActionResult) error {
	outDir := filepath.Join(e.workDir, "out", labelToDir(label))
	if err := os.MkdirAll(outDir, 0755); err != nil {
		return err
	}
	for name, digest := range result.OutputDigests {
		dest := filepath.Join(outDir, name)
		if err := e.cache.RetrieveFile(digest, dest); err != nil {
			return fmt.Errorf("restore output %s: %w", name, err)
		}
	}
	return nil
}

// labelToDir converts //src/hello:hello → src/hello.
func labelToDir(label string) string {
	// Strip leading //
	s := strings.TrimPrefix(label, "//")
	// Drop the :name suffix
	if idx := strings.LastIndex(s, ":"); idx >= 0 {
		s = s[:idx]
	}
	return s
}
