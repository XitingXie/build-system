// Package graph builds a DAG from parsed targets and provides
// cycle detection, topological sort, and critical-path calculation.
package graph

import (
	"fmt"
	"strings"

	"github.com/xitingxie/build-system/internal/parser"
)

// Graph is a directed acyclic graph of build targets.
type Graph struct {
	Nodes map[string]*Node // label → node
}

// Node wraps a target and its resolved edges.
type Node struct {
	Target   parser.Target
	Deps     []*Node // outgoing edges (dependencies)
	Rdeps    []*Node // reverse edges (dependents)
	CritPath int     // longest path to a leaf (set by CriticalPath)
}

// New builds a Graph from a slice of targets.
// Returns an error if a label is duplicated or a dependency is undeclared.
func New(targets []parser.Target) (*Graph, error) {
	g := &Graph{Nodes: make(map[string]*Node, len(targets))}

	// First pass: register all nodes.
	for _, t := range targets {
		if _, exists := g.Nodes[t.Label]; exists {
			return nil, fmt.Errorf("duplicate label: %s", t.Label)
		}
		g.Nodes[t.Label] = &Node{Target: t}
	}

	// Second pass: wire edges.
	for _, t := range targets {
		n := g.Nodes[t.Label]
		for _, dep := range t.Deps {
			d, ok := g.Nodes[dep]
			if !ok {
				return nil, fmt.Errorf("target %s depends on unknown label %s", t.Label, dep)
			}
			n.Deps = append(n.Deps, d)
			d.Rdeps = append(d.Rdeps, n)
		}
	}

	return g, nil
}

// DetectCycles returns an error describing the first cycle found, or nil.
func (g *Graph) DetectCycles() error {
	// DFS with three colours: 0=white, 1=grey (on stack), 2=black (done).
	color := make(map[string]int, len(g.Nodes))
	path := make([]string, 0)

	var dfs func(n *Node) error
	dfs = func(n *Node) error {
		color[n.Target.Label] = 1
		path = append(path, n.Target.Label)
		for _, dep := range n.Deps {
			switch color[dep.Target.Label] {
			case 1:
				// Found a back-edge: report the cycle.
				start := dep.Target.Label
				cycle := []string{}
				for i := len(path) - 1; i >= 0; i-- {
					cycle = append([]string{path[i]}, cycle...)
					if path[i] == start {
						break
					}
				}
				cycle = append(cycle, start) // close the loop
				return fmt.Errorf("cycle detected: %s", strings.Join(cycle, " → "))
			case 0:
				if err := dfs(dep); err != nil {
					return err
				}
			}
		}
		path = path[:len(path)-1]
		color[n.Target.Label] = 2
		return nil
	}

	for _, n := range g.Nodes {
		if color[n.Target.Label] == 0 {
			if err := dfs(n); err != nil {
				return err
			}
		}
	}
	return nil
}

// TopoSort returns all nodes in topological order (dependencies before dependents).
// Returns an error if a cycle exists.
func (g *Graph) TopoSort() ([]*Node, error) {
	// Kahn's algorithm.
	inDegree := make(map[string]int, len(g.Nodes))
	for _, n := range g.Nodes {
		if _, ok := inDegree[n.Target.Label]; !ok {
			inDegree[n.Target.Label] = 0
		}
		for _, dep := range n.Deps {
			inDegree[dep.Target.Label] = inDegree[dep.Target.Label] // ensure key exists
			_ = dep
		}
	}
	// Count how many nodes depend on each node (in-degree in reversed sense).
	// We want deps-first order, so "in-degree" here = number of declared deps.
	depCount := make(map[string]int, len(g.Nodes))
	for _, n := range g.Nodes {
		depCount[n.Target.Label] = len(n.Deps)
	}

	queue := make([]*Node, 0)
	for _, n := range g.Nodes {
		if depCount[n.Target.Label] == 0 {
			queue = append(queue, n)
		}
	}

	var order []*Node
	for len(queue) > 0 {
		// Pop front.
		cur := queue[0]
		queue = queue[1:]
		order = append(order, cur)

		// Reduce dep-count for all nodes that depend on cur.
		for _, rdep := range cur.Rdeps {
			depCount[rdep.Target.Label]--
			if depCount[rdep.Target.Label] == 0 {
				queue = append(queue, rdep)
			}
		}
	}

	if len(order) != len(g.Nodes) {
		return nil, fmt.Errorf("cycle detected: topological sort incomplete")
	}
	return order, nil
}

// CriticalPath computes the critical-path length for every node.
// CritPath is defined as the number of actions on the longest dependency chain
// from the node to any leaf it transitively depends on (inclusive).
// Must be called after TopoSort succeeds.
func (g *Graph) CriticalPath() ([]*Node, error) {
	order, err := g.TopoSort()
	if err != nil {
		return nil, err
	}
	// Process in topo order (deps first).
	for _, n := range order {
		max := 0
		for _, dep := range n.Deps {
			if dep.CritPath > max {
				max = dep.CritPath
			}
		}
		n.CritPath = max + 1
	}
	return order, nil
}

// Subgraph returns the transitive closure of nodes reachable from the given label.
func (g *Graph) Subgraph(label string) ([]*Node, error) {
	root, ok := g.Nodes[label]
	if !ok {
		return nil, fmt.Errorf("unknown target: %s", label)
	}
	visited := make(map[string]bool)
	var result []*Node
	var dfs func(n *Node)
	dfs = func(n *Node) {
		if visited[n.Target.Label] {
			return
		}
		visited[n.Target.Label] = true
		for _, dep := range n.Deps {
			dfs(dep)
		}
		result = append(result, n)
	}
	dfs(root)
	return result, nil
}
