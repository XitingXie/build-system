// Package scheduler executes build actions in parallel, respecting dependency order.
// Actions are prioritised by critical-path length so the longest chain starts first.
package scheduler

import (
	"fmt"
	"runtime"
	"sync"

	"github.com/xitingxie/build-system/internal/executor"
	"github.com/xitingxie/build-system/internal/graph"
)

// Scheduler runs the build graph in parallel up to maxWorkers goroutines.
type Scheduler struct {
	exec       *executor.Executor
	maxWorkers int
}

// New creates a Scheduler. If maxWorkers <= 0 it defaults to GOMAXPROCS.
func New(exec *executor.Executor, maxWorkers int) *Scheduler {
	if maxWorkers <= 0 {
		maxWorkers = runtime.GOMAXPROCS(0)
	}
	return &Scheduler{exec: exec, maxWorkers: maxWorkers}
}

// Run executes all nodes in the graph, respecting dependencies.
// It returns the first error encountered (remaining in-flight actions are waited on).
func (s *Scheduler) Run(g *graph.Graph) error {
	// Compute critical paths so we can prioritise.
	order, err := g.CriticalPath()
	if err != nil {
		return err
	}

	type state struct {
		done       bool
		err        error
		actionKey  [32]byte // digest produced for this node
		mu         sync.Mutex
		waiters    []chan struct{} // closed when this node finishes
	}

	states := make(map[string]*state, len(order))
	for _, n := range order {
		states[n.Target.Label] = &state{}
	}

	// notify closes all waiters for a node.
	notify := func(st *state) {
		st.mu.Lock()
		for _, ch := range st.waiters {
			close(ch)
		}
		st.waiters = nil
		st.mu.Unlock()
	}

	// subscribe returns a channel that will be closed when st finishes.
	subscribe := func(st *state) chan struct{} {
		st.mu.Lock()
		defer st.mu.Unlock()
		if st.done {
			ch := make(chan struct{})
			close(ch)
			return ch
		}
		ch := make(chan struct{})
		st.waiters = append(st.waiters, ch)
		return ch
	}

	sem := make(chan struct{}, s.maxWorkers)
	var wg sync.WaitGroup
	var firstErr error
	var errMu sync.Mutex

	setErr := func(err error) {
		errMu.Lock()
		if firstErr == nil {
			firstErr = err
		}
		errMu.Unlock()
	}

	for _, n := range order {
		n := n // capture
		st := states[n.Target.Label]

		wg.Add(1)
		go func() {
			defer wg.Done()

			// Wait for all dependencies to finish.
			depDigests := make([][32]byte, 0, len(n.Deps))
			for _, dep := range n.Deps {
				depSt := states[dep.Target.Label]
				ch := subscribe(depSt)
				<-ch
				depSt.mu.Lock()
				if depSt.err != nil {
					setErr(fmt.Errorf("dependency %s failed", dep.Target.Label))
					depSt.mu.Unlock()
					st.mu.Lock()
					st.done = true
					st.err = firstErr
					st.mu.Unlock()
					notify(st)
					return
				}
				depDigests = append(depDigests, depSt.actionKey)
				depSt.mu.Unlock()
			}

			// Acquire a worker slot (respects maxWorkers).
			sem <- struct{}{}
			defer func() { <-sem }()

			err := s.exec.Run(n, depDigests)

			st.mu.Lock()
			st.done = true
			st.err = err
			st.mu.Unlock()
			if err != nil {
				setErr(err)
			}
			notify(st)
		}()
	}

	wg.Wait()
	return firstErr
}
