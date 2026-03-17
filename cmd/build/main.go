// build is the CLI entry point for the toy build system.
//
// Usage:
//
//	build <target>              — build a target and its dependencies
//	build //...                 — build all targets in the workspace
//	build query deps <target>   — list transitive dependencies
//	build query rdeps <target>  — list targets that depend on <target>
//	build clean                 — clear the local cache
//	build graph <target>        — print dependency graph as ASCII
package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/xitingxie/build-system/internal/cache"
	"github.com/xitingxie/build-system/internal/executor"
	"github.com/xitingxie/build-system/internal/graph"
	"github.com/xitingxie/build-system/internal/parser"
	"github.com/xitingxie/build-system/internal/scheduler"
)

func main() {
	if err := run(os.Args[1:]); err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}

func run(args []string) error {
	if len(args) == 0 {
		return fmt.Errorf("usage: build <target> | build //... | build query deps|rdeps <target> | build clean | build graph <target>")
	}

	workDir, err := os.Getwd()
	if err != nil {
		return err
	}

	switch args[0] {
	case "clean":
		return cmdClean()
	case "query":
		return cmdQuery(workDir, args[1:])
	case "graph":
		return cmdGraph(workDir, args[1:])
	default:
		return cmdBuild(workDir, args[0])
	}
}

// cmdBuild builds a target (or all targets with //...).
func cmdBuild(workDir, label string) error {
	targets, err := parser.ParseDir(workDir)
	if err != nil {
		return err
	}
	g, err := graph.New(targets)
	if err != nil {
		return err
	}
	if err := g.DetectCycles(); err != nil {
		return err
	}

	// Determine the subgraph to build.
	var buildGraph *graph.Graph
	if label == "//..." {
		buildGraph = g
	} else {
		nodes, err := g.Subgraph(label)
		if err != nil {
			return err
		}
		// Rebuild a graph from just those nodes.
		sub := make([]parser.Target, 0, len(nodes))
		for _, n := range nodes {
			sub = append(sub, n.Target)
		}
		buildGraph, err = graph.New(sub)
		if err != nil {
			return err
		}
	}

	c, err := cache.New(cache.DefaultDir())
	if err != nil {
		return err
	}
	exec := executor.New(c, workDir)
	sched := scheduler.New(exec, 0) // 0 = use GOMAXPROCS

	fmt.Printf("Building %s\n", label)
	return sched.Run(buildGraph)
}

// cmdClean removes the local cache.
func cmdClean() error {
	c, err := cache.New(cache.DefaultDir())
	if err != nil {
		return err
	}
	if err := c.Clean(); err != nil {
		return err
	}
	fmt.Println("Cache cleared.")
	return nil
}

// cmdQuery handles "query deps <target>" and "query rdeps <target>".
func cmdQuery(workDir string, args []string) error {
	if len(args) < 2 {
		return fmt.Errorf("usage: build query deps|rdeps <target>")
	}
	subCmd, label := args[0], args[1]

	targets, err := parser.ParseDir(workDir)
	if err != nil {
		return err
	}
	g, err := graph.New(targets)
	if err != nil {
		return err
	}

	switch subCmd {
	case "deps":
		nodes, err := g.Subgraph(label)
		if err != nil {
			return err
		}
		for _, n := range nodes {
			if n.Target.Label != label {
				fmt.Println(n.Target.Label)
			}
		}
	case "rdeps":
		root, ok := g.Nodes[label]
		if !ok {
			return fmt.Errorf("unknown target: %s", label)
		}
		visited := map[string]bool{}
		var walk func(n *graph.Node)
		walk = func(n *graph.Node) {
			for _, r := range n.Rdeps {
				if !visited[r.Target.Label] {
					visited[r.Target.Label] = true
					fmt.Println(r.Target.Label)
					walk(r)
				}
			}
		}
		walk(root)
	default:
		return fmt.Errorf("unknown query subcommand: %s (use deps or rdeps)", subCmd)
	}
	return nil
}

// cmdGraph prints an ASCII representation of a target's dependency tree.
func cmdGraph(workDir string, args []string) error {
	if len(args) < 1 {
		return fmt.Errorf("usage: build graph <target>")
	}
	label := args[0]

	targets, err := parser.ParseDir(workDir)
	if err != nil {
		return err
	}
	g, err := graph.New(targets)
	if err != nil {
		return err
	}

	root, ok := g.Nodes[label]
	if !ok {
		return fmt.Errorf("unknown target: %s", label)
	}

	printTree(root, "", true)
	return nil
}

func printTree(n *graph.Node, prefix string, isLast bool) {
	connector := "├── "
	if isLast {
		connector = "└── "
	}
	if prefix == "" {
		fmt.Println(n.Target.Label)
	} else {
		fmt.Println(prefix + connector + n.Target.Label)
	}
	childPrefix := prefix
	if prefix == "" {
		childPrefix = ""
	} else if isLast {
		childPrefix += "    "
	} else {
		childPrefix += "│   "
	}
	for i, dep := range n.Deps {
		printTree(dep, childPrefix, i == len(n.Deps)-1)
	}
}

// labelToDir converts //src/hello:hello → src/hello (used by executor, duplicated here for clarity).
func labelToDir(label string) string {
	s := strings.TrimPrefix(label, "//")
	if idx := strings.LastIndex(s, ":"); idx >= 0 {
		s = s[:idx]
	}
	return s
}
