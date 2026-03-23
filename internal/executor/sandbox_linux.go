//go:build linux

// Package executor runs a build action and stores its outputs in the cache.
package executor

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"syscall"
)

// sandboxRun executes cmd inside a hermetic Linux sandbox.
//
// A new user+mount namespace is created so the process appears as root inside
// but is still constrained to the real user's file permissions on the host.
// Inside the namespace:
//   - A tmpfs is mounted as the sandbox root.
//   - Standard system directories (/bin, /usr, /lib, …) are bind-mounted read-only.
//   - Only the explicitly declared inputPaths are bind-mounted read-only from
//     the workspace; all other workspace files are invisible.
//   - outDir is bind-mounted read-write so the action can write its outputs.
//   - The command runs inside a chroot(2) into the sandbox root.
//
// On a cache miss this is what enforces hermeticity: if the action tries to
// read a file that was not declared as an input the open(2) call will return
// ENOENT, failing the build immediately rather than silently succeeding.
func sandboxRun(cmd, dir string, inputPaths []string, outDir string) error {
	sandboxRoot, err := os.MkdirTemp("", "build-sandbox-*")
	if err != nil {
		return fmt.Errorf("sandbox: create root dir: %w", err)
	}
	defer os.RemoveAll(sandboxRoot)

	// Build a shell script that runs inside the new namespace.
	// All mount operations happen after CLONE_NEWNS so they are invisible to
	// the parent process and are cleaned up automatically when the namespace exits.
	var s strings.Builder
	s.WriteString("set -e\n")

	// A fresh tmpfs provides the skeleton of the new root filesystem.
	s.WriteString(fmt.Sprintf("mount -t tmpfs tmpfs %s\n", q(sandboxRoot)))

	// Read-only system directories (compiler, linker, libc, shell, …).
	for _, d := range []string{"/bin", "/usr", "/sbin", "/lib", "/lib64", "/etc", "/dev", "/proc", "/sys"} {
		if _, err := os.Stat(d); err != nil {
			continue // silently skip dirs that don't exist on this host
		}
		dst := filepath.Join(sandboxRoot, d)
		s.WriteString(fmt.Sprintf(
			"mkdir -p %s && mount --bind %s %s && mount -o remount,bind,ro %s\n",
			q(dst), q(d), q(dst), q(dst),
		))
	}

	// Declared input files — bind-mounted read-only.
	// Each file is mounted at the same absolute path it has on the host so
	// the already-expanded $SRCS paths in cmd remain valid inside the chroot.
	for _, inp := range inputPaths {
		dst := filepath.Join(sandboxRoot, inp)
		s.WriteString(fmt.Sprintf(
			"mkdir -p %s && touch %s && mount --bind %s %s && mount -o remount,bind,ro %s\n",
			q(filepath.Dir(dst)), q(dst), q(inp), q(dst), q(dst),
		))
	}

	// Output directory — bind-mounted read-write.
	dstOut := filepath.Join(sandboxRoot, outDir)
	s.WriteString(fmt.Sprintf(
		"mkdir -p %s && mount --bind %s %s\n",
		q(dstOut), q(outDir), q(dstOut),
	))

	// chroot into the sandbox and run the action.
	// We cd to the original working directory first (which is visible because
	// we bind-mounted the input files, implicitly creating its path).
	innerCmd := fmt.Sprintf("cd %s && %s", q(dir), cmd)
	s.WriteString(fmt.Sprintf("exec chroot %s /bin/sh -c %s\n", q(sandboxRoot), q(innerCmd)))

	c := exec.Command("/bin/sh", "-c", s.String())
	c.Stdout = os.Stdout
	c.Stderr = os.Stderr
	c.SysProcAttr = &syscall.SysProcAttr{
		// CLONE_NEWUSER: unprivileged sandbox — the process maps to uid 0 inside
		// the namespace, allowing chroot and bind-mount without host root.
		// CLONE_NEWNS:   private mount namespace — our mounts are invisible to
		// the rest of the system and are torn down when the process exits.
		Cloneflags: syscall.CLONE_NEWUSER | syscall.CLONE_NEWNS,
		UidMappings: []syscall.SysProcIDMap{
			{ContainerID: 0, HostID: os.Getuid(), Size: 1},
		},
		GidMappings: []syscall.SysProcIDMap{
			{ContainerID: 0, HostID: os.Getgid(), Size: 1},
		},
	}
	return c.Run()
}

// q returns a single-quoted, shell-safe representation of s.
func q(s string) string {
	return "'" + strings.ReplaceAll(s, "'", `'\''`) + "'"
}
