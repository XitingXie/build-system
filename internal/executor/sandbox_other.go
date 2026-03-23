//go:build !linux

// Package executor runs a build action and stores its outputs in the cache.
package executor

import (
	"os"
	"os/exec"
)

// sandboxRun on non-Linux platforms runs cmd without namespace isolation.
// The declared inputPaths and outDir parameters are accepted for API
// compatibility but no filesystem restriction is applied.
func sandboxRun(cmd, dir string, _ []string, _ string) error {
	c := exec.Command("/bin/sh", "-c", cmd)
	c.Dir = dir
	c.Stdout = os.Stdout
	c.Stderr = os.Stderr
	return c.Run()
}
