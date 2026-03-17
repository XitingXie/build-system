# Build System Implementation Plan

## Overview

A minimal but complete artifact-based build system with content-addressed caching, hermetic execution, and parallel scheduling. Built in **Go**. Inspired by Bazel's design principles — independent implementation, no Bazel dependency.

---

## Architecture

```
                    ┌─────────────────────────────────┐
                    │         CLI (build tool)         │
                    └────────────────┬────────────────┘
                                     │
          ┌──────────────────────────▼──────────────────────────┐
          │                    Core Engine                       │
          │   Load Phase → Analysis Phase → Execution Phase      │
          └───┬──────────────┬──────────────────────┬──────────┘
              │              │                      │
       ┌──────▼─────┐ ┌──────▼──────┐      ┌───────▼──────┐
       │  BUILD File│ │  Dependency │      │   Scheduler  │
       │   Parser   │ │    Graph    │      │  (topo sort) │
       └────────────┘ └─────────────┘      └───────┬──────┘
                                                   │
                            ┌──────────────────────▼──────────┐
                            │          Executor Pool           │
                            │  (hermetic sandbox per action)   │
                            └───────────────┬─────────────────┘
                                            │
                               ┌────────────▼────────────┐
                               │       Cache Layer        │
                               │  Local CAS + Remote CAS  │
                               └─────────────────────────┘
```

---

## Phases & Components

### Phase 1 — Core Engine (Foundations)

| Component | Description |
|-----------|-------------|
| **BUILD file parser** | Parse declarative target definitions (simple JSON/TOML DSL) |
| **Target graph** | DAG of targets with dependency edges |
| **Cycle detection** | DFS pre-execution; hard error on cycle |
| **Topological sort** | Kahn's algorithm for valid build order |
| **Critical path calculator** | Dynamic programming; exposes parallelism limit |

### Phase 2 — Change Detection & Caching

| Component | Description |
|-----------|-------------|
| **Content hasher** | SHA-256 hash of inputs: source files, compiler binary, flags, transitive dep hashes |
| **Cache key builder** | Combines all input hashes into single action digest |
| **Local CAS** | Content-addressable store on disk (`~/.cache/build-system/`) |
| **Action cache** | `action_digest → action_result` mapping |
| **Merkle tree** | Efficient verification of large input sets |

### Phase 3 — Execution & Hermeticity

| Component | Description |
|-----------|-------------|
| **Sandbox** | Linux: mount namespaces; macOS: sandbox-exec; only declared inputs visible |
| **Action executor** | Execute `(inputs, command, expected_outputs)` in hermetic environment |
| **Parallel scheduler** | Ready queue ordered by critical path length; work stealing for idle workers |
| **Output uploader** | Upload outputs to local CAS after execution |

### Phase 4 — Remote Execution (Optional / Advanced)

| Component | Description |
|-----------|-------------|
| **RBE protocol** | gRPC-based Remote Build Execution API |
| **Remote CAS** | S3/GCS-backed content-addressable store |
| **Stateless workers** | Download inputs from CAS, execute, upload outputs |
| **Remote action cache** | Shared across all machines; warm cache grows as team builds |

### Phase 5 — Graph Health & Tooling

| Component | Description |
|-----------|-------------|
| **Dependency query** | `query` command: who depends on X? what does X depend on? |
| **Metrics/reporting** | Critical path length, cache hit rate, module fanout, build time percentiles |
| **Unused dep detection** | Automated detection of declared but unused dependencies |

---

## Data Structures

```go
type Target struct {
    Label string            // "//src/foo:bar"
    Srcs  []string
    Deps  []string
    Cmd   string
    Outs  []string
}

type Action struct {
    Inputs       []FileDigest        // content-addressed
    Command      []string
    Env          map[string]string
    Outputs      []string
    ActionDigest [32]byte            // SHA-256 cache key
}

type DAG struct {
    Nodes map[string]*Target
    Edges map[string][]string        // adjacency list
}

type CacheEntry struct {
    ActionDigest  [32]byte
    OutputDigests map[string][32]byte // path → digest
    ExitCode      int
}
```

---

## Key Algorithms

| Algorithm | Purpose | Complexity |
|-----------|---------|------------|
| DFS with recursion stack | Cycle detection | O(V+E) |
| Kahn's algorithm | Topological sort | O(V+E) |
| DP on DAG (longest path) | Critical path | O(V+E) |
| BFS/DFS from changed nodes | Change propagation | O(V+E) |
| Priority queue by critical path | List scheduling | O(V log V) |
| Work stealing | Idle worker utilization | — |

---

## Implementation Steps

```
Step 1:  BUILD file format + parser
Step 2:  Target graph + cycle detection + topological sort
Step 3:  Content hashing + local cache (CAS + action cache)
Step 4:  Basic sequential executor (no sandbox yet)
Step 5:  Parallel scheduler (critical path ordering)
Step 6:  Hermetic sandbox (Linux namespaces / macOS sandbox-exec)
Step 7:  CLI (build, query, clean commands)
Step 8:  Remote cache (S3/GCS backend)
Step 9:  Graph health metrics + query tool
Step 10: Remote execution (RBE protocol)
```

---

## Technology Choices

| Layer | Choice | Reason |
|-------|--------|--------|
| Language | **Go** | Native goroutines, single binary, simple concurrency model |
| BUILD file DSL | JSON or TOML | Simple to parse, no external dependencies |
| Hashing | SHA-256 (`crypto/sha256`) | Collision resistance, standard library |
| Sandboxing | Linux mount namespaces + seccomp | Kernel-level hermeticity |
| Remote cache backend | S3-compatible API | Wide compatibility |
| RPC | gRPC | Standard for RBE protocol |

---

## Project Structure

```
build-system/
├── docs/
│   ├── build_system_book.md
│   └── plan.md
├── cmd/
│   └── build/
│       └── main.go          # CLI entrypoint
├── internal/
│   ├── parser/              # BUILD file parser
│   ├── graph/               # DAG, cycle detection, topo sort, critical path
│   ├── hasher/              # SHA-256 content hashing, cache key builder
│   ├── cache/               # Local CAS + action cache
│   ├── executor/            # Action executor + sandbox
│   ├── scheduler/           # Parallel scheduler + work stealing
│   └── query/               # Dependency query tool
└── go.mod
```
