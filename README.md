# Build System

A toy build system written in Go, inspired by Bazel. It supports declarative targets, content-addressed caching, parallel execution, and critical-path scheduling. Every build is recorded to a metrics database for analysis with cloud BI tools.

## Features

- **Declarative BUILD files** — JSON-based target definitions
- **Dependency graph** — cycle detection, topological sort, subgraph queries
- **Content-addressed cache** — skip actions whose inputs haven't changed
- **Parallel execution** — runs independent actions concurrently, prioritised by critical path
- **Metrics DB** — every build and action is persisted to a two-table JSONL database

## Installation

```bash
git clone https://github.com/xitingxie/build-system
cd build-system
go install ./cmd/build
```

## Usage

```bash
build <target>                  # build a target and its dependencies
build //...                     # build every target in the workspace
build query deps <target>       # list transitive dependencies
build query rdeps <target>      # list targets that depend on <target>
build graph <target>            # print dependency tree as ASCII
build clean                     # clear the local cache
```

## BUILD File Format

BUILD files are JSON. Place one in any directory of your workspace.

```json
{
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
}
```

| Field   | Description |
|---------|-------------|
| `label` | Fully-qualified name: `//path/to/dir:name` |
| `srcs`  | Source files relative to the BUILD file's directory |
| `deps`  | Labels of targets this target depends on |
| `cmd`   | Shell command to run. `$SRCS` expands to source paths, `$OUT` to the first output path |
| `outs`  | Output file names produced by the command |

## Caching

The cache lives at `~/.cache/build-system/` and has two layers:

| Layer | Key | Value |
|-------|-----|-------|
| Action cache | SHA-256 of (inputs + command + dep digests) | Map of output name → CAS digest |
| CAS | SHA-256 of file contents | Raw file bytes |

On a cache hit the action is skipped entirely and outputs are restored from the CAS. Running `build clean` removes both layers.

## Metrics Database

Every build appends records to `~/.cache/build-system/db/`:

```
builds.jsonl   — one record per build invocation
actions.jsonl  — one record per action, linked by build_id
```

**builds.jsonl schema**

| Field | Type | Description |
|-------|------|-------------|
| `id` | int | Unix-nanosecond timestamp (primary key) |
| `started_at` | timestamp | When the build started |
| `duration_ms` | int | Total wall time in milliseconds |
| `target` | string | Target label requested (e.g. `//...`) |
| `success` | bool | Whether the build succeeded |

**actions.jsonl schema**

| Field | Type | Description |
|-------|------|-------------|
| `build_id` | int | FK → `builds.id` |
| `label` | string | Target label |
| `cache_hit` | bool | Whether the action was served from cache |
| `duration_ms` | int | Action wall time in milliseconds |
| `exit_code` | int | Exit code of the shell command |

### Querying locally

```bash
# Pretty-print all builds
cat ~/.cache/build-system/db/builds.jsonl | python3 -m json.tool

# Cache hit rate across all actions
jq -s '(map(select(.cache_hit)) | length) / length * 100' \
  ~/.cache/build-system/db/actions.jsonl

# Slowest targets by average duration
jq -s 'group_by(.label)
  | map({label: .[0].label, avg_ms: (map(.duration_ms) | add/length)})
  | sort_by(-.avg_ms)' \
  ~/.cache/build-system/db/actions.jsonl
```

### Connecting to a cloud dashboard

The JSONL files load directly into cloud BI tools without any transformation:

- **Google Cloud** — upload to GCS, create BigQuery external tables, connect Looker Studio
- **AWS** — sync to S3, define an Athena table, connect QuickSight or Grafana
- **Self-hosted** — serve the files over HTTP and use Grafana's JSON datasource plugin

## Project Structure

```
cmd/build/          CLI entry point
internal/
  parser/           BUILD file parser
  graph/            DAG construction, cycle detection, topological sort
  hasher/           SHA-256 content hashing and action key computation
  cache/            Content-addressable store and action cache
  executor/         Action execution, variable expansion, output management
  scheduler/        Parallel scheduler with critical-path prioritisation
  metrics/          Two-table JSONL metrics database
examples/
  greeting/         Sample three-target package
  upper/            Sample cross-package target
verify.sh           End-to-end verification script (10 checks)
```

## Verification

```bash
bash verify.sh
```

Runs 10 checks: fresh build, cached build, incremental rebuild after source change, cross-package dependencies, `query deps`, `query rdeps`, `graph`, `build //...`, and `build clean`.

## Architecture

```
ParseDir → Graph → DetectCycles → Subgraph → Scheduler ──► Executor ──► Cache
                                                  │                        │
                                             critical-path             CAS + action
                                             ordering                  cache (SHA-256)
                                                  │
                                             metrics.DB
                                         (builds + actions)
```
