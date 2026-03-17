package parser

import (
	"os"
	"path/filepath"
	"testing"
)

const sampleBUILD = `{
  "targets": [
    {
      "label": "//src/lib:util",
      "srcs":  ["util.go"],
      "deps":  [],
      "cmd":   "go build -o $OUT $SRCS",
      "outs":  ["util.a"]
    },
    {
      "label": "//src/hello:hello",
      "srcs":  ["main.go"],
      "deps":  ["//src/lib:util"],
      "cmd":   "go build -o $OUT $SRCS",
      "outs":  ["hello"]
    }
  ]
}`

func TestParseValid(t *testing.T) {
	targets, err := parse([]byte(sampleBUILD), "test")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(targets) != 2 {
		t.Fatalf("expected 2 targets, got %d", len(targets))
	}
	if targets[0].Label != "//src/lib:util" {
		t.Errorf("unexpected label: %s", targets[0].Label)
	}
	if len(targets[1].Deps) != 1 || targets[1].Deps[0] != "//src/lib:util" {
		t.Errorf("unexpected deps: %v", targets[1].Deps)
	}
}

func TestParseMissingLabel(t *testing.T) {
	bad := `{"targets":[{"srcs":["a.go"],"cmd":"go build -o $OUT $SRCS","outs":["a"]}]}`
	_, err := parse([]byte(bad), "test")
	if err == nil {
		t.Fatal("expected error for missing label")
	}
}

func TestParseBadLabel(t *testing.T) {
	bad := `{"targets":[{"label":"src/hello","srcs":["a.go"],"cmd":"go build -o $OUT $SRCS","outs":["a"]}]}`
	_, err := parse([]byte(bad), "test")
	if err == nil {
		t.Fatal("expected error for bad label")
	}
}

func TestParseMissingCmd(t *testing.T) {
	bad := `{"targets":[{"label":"//src:a","srcs":["a.go"],"outs":["a"]}]}`
	_, err := parse([]byte(bad), "test")
	if err == nil {
		t.Fatal("expected error for missing cmd")
	}
}

func TestParseDir(t *testing.T) {
	dir := t.TempDir()

	// Write a BUILD file in a subdirectory.
	sub := filepath.Join(dir, "src", "lib")
	if err := os.MkdirAll(sub, 0755); err != nil {
		t.Fatal(err)
	}
	content := `{"targets":[{"label":"//src/lib:util","srcs":["util.go"],"cmd":"go build -o $OUT $SRCS","outs":["util.a"]}]}`
	if err := os.WriteFile(filepath.Join(sub, "BUILD"), []byte(content), 0644); err != nil {
		t.Fatal(err)
	}

	targets, err := ParseDir(dir)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(targets) != 1 {
		t.Fatalf("expected 1 target, got %d", len(targets))
	}
}
