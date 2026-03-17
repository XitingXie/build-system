// Package parser reads BUILD files and returns a list of targets.
//
// BUILD file format (JSON):
//
//	{
//	  "targets": [
//	    {
//	      "label": "//src/hello:hello",
//	      "srcs":  ["main.go"],
//	      "deps":  ["//src/lib:util"],
//	      "cmd":   "go build -o $OUT $SRCS",
//	      "outs":  ["hello"]
//	    }
//	  ]
//	}
package parser

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// Target represents a single build target declared in a BUILD file.
type Target struct {
	Label string   `json:"label"` // e.g. //src/hello:hello
	Srcs  []string `json:"srcs"`  // source files relative to the BUILD file
	Deps  []string `json:"deps"`  // labels of dependencies
	Cmd   string   `json:"cmd"`   // shell command; $SRCS and $OUT are expanded
	Outs  []string `json:"outs"`  // expected output file names
}

// buildFile is the top-level structure of a BUILD file.
type buildFile struct {
	Targets []Target `json:"targets"`
}

// ParseFile reads and parses a single BUILD file.
func ParseFile(path string) ([]Target, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("read %s: %w", path, err)
	}
	return parse(data, path)
}

// ParseDir walks a directory tree and parses every file named "BUILD".
// All targets from all BUILD files are returned in the order they are found.
func ParseDir(root string) ([]Target, error) {
	var all []Target
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() && info.Name() == "BUILD" {
			targets, err := ParseFile(path)
			if err != nil {
				return err
			}
			all = append(all, targets...)
		}
		return nil
	})
	return all, err
}

// parse unmarshals raw JSON and validates each target.
func parse(data []byte, source string) ([]Target, error) {
	var bf buildFile
	if err := json.Unmarshal(data, &bf); err != nil {
		return nil, fmt.Errorf("parse %s: %w", source, err)
	}
	for i, t := range bf.Targets {
		if err := validate(t); err != nil {
			return nil, fmt.Errorf("%s target[%d]: %w", source, i, err)
		}
	}
	return bf.Targets, nil
}

// validate checks that a target is well-formed.
func validate(t Target) error {
	if t.Label == "" {
		return fmt.Errorf("missing label")
	}
	if !strings.HasPrefix(t.Label, "//") {
		return fmt.Errorf("label %q must start with //", t.Label)
	}
	if !strings.Contains(t.Label, ":") {
		return fmt.Errorf("label %q must contain ':'", t.Label)
	}
	if t.Cmd == "" {
		return fmt.Errorf("label %q: missing cmd", t.Label)
	}
	if len(t.Outs) == 0 {
		return fmt.Errorf("label %q: outs must have at least one entry", t.Label)
	}
	return nil
}
