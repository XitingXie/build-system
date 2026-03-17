package graph

import (
	"testing"

	"github.com/xitingxie/build-system/internal/parser"
)

func makeTargets() []parser.Target {
	return []parser.Target{
		{Label: "//a:a", Srcs: []string{"a.go"}, Cmd: "build", Outs: []string{"a"}},
		{Label: "//b:b", Srcs: []string{"b.go"}, Deps: []string{"//a:a"}, Cmd: "build", Outs: []string{"b"}},
		{Label: "//c:c", Srcs: []string{"c.go"}, Deps: []string{"//b:b"}, Cmd: "build", Outs: []string{"c"}},
	}
}

func TestNewGraph(t *testing.T) {
	g, err := New(makeTargets())
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(g.Nodes) != 3 {
		t.Fatalf("expected 3 nodes, got %d", len(g.Nodes))
	}
}

func TestDuplicateLabel(t *testing.T) {
	targets := makeTargets()
	targets = append(targets, targets[0]) // duplicate //a:a
	_, err := New(targets)
	if err == nil {
		t.Fatal("expected error for duplicate label")
	}
}

func TestUnknownDep(t *testing.T) {
	targets := []parser.Target{
		{Label: "//a:a", Deps: []string{"//missing:x"}, Cmd: "build", Outs: []string{"a"}},
	}
	_, err := New(targets)
	if err == nil {
		t.Fatal("expected error for unknown dep")
	}
}

func TestNoCycle(t *testing.T) {
	g, _ := New(makeTargets())
	if err := g.DetectCycles(); err != nil {
		t.Fatalf("unexpected cycle error: %v", err)
	}
}

func TestCycleDetected(t *testing.T) {
	targets := []parser.Target{
		{Label: "//a:a", Deps: []string{"//b:b"}, Cmd: "build", Outs: []string{"a"}},
		{Label: "//b:b", Deps: []string{"//a:a"}, Cmd: "build", Outs: []string{"b"}},
	}
	g, err := New(targets)
	if err != nil {
		t.Fatalf("unexpected error building graph: %v", err)
	}
	if err := g.DetectCycles(); err == nil {
		t.Fatal("expected cycle error")
	}
}

func TestTopoSort(t *testing.T) {
	g, _ := New(makeTargets())
	order, err := g.TopoSort()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(order) != 3 {
		t.Fatalf("expected 3 nodes, got %d", len(order))
	}
	// //a:a must come before //b:b, //b:b before //c:c
	pos := make(map[string]int)
	for i, n := range order {
		pos[n.Target.Label] = i
	}
	if pos["//a:a"] > pos["//b:b"] {
		t.Error("//a:a should come before //b:b")
	}
	if pos["//b:b"] > pos["//c:c"] {
		t.Error("//b:b should come before //c:c")
	}
}

func TestCriticalPath(t *testing.T) {
	g, _ := New(makeTargets())
	order, err := g.CriticalPath()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	cp := make(map[string]int)
	for _, n := range order {
		cp[n.Target.Label] = n.CritPath
	}
	// //a:a has no deps → CritPath = 1
	// //b:b depends on //a:a → CritPath = 2
	// //c:c depends on //b:b → CritPath = 3
	if cp["//a:a"] != 1 {
		t.Errorf("expected CritPath 1 for //a:a, got %d", cp["//a:a"])
	}
	if cp["//b:b"] != 2 {
		t.Errorf("expected CritPath 2 for //b:b, got %d", cp["//b:b"])
	}
	if cp["//c:c"] != 3 {
		t.Errorf("expected CritPath 3 for //c:c, got %d", cp["//c:c"])
	}
}

func TestSubgraph(t *testing.T) {
	g, _ := New(makeTargets())
	nodes, err := g.Subgraph("//c:c")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(nodes) != 3 {
		t.Errorf("expected 3 nodes in subgraph, got %d", len(nodes))
	}
}
