# Build Systems: A Deep Dive
**Status:** Draft — outline + content in progress  
**Last updated:** 2026-03-16  
**Progress:** 6 / 18 chapters fully drafted · 7 chapters partially discussed

---

## Progress legend

| Symbol | Meaning |
|--------|---------|
| ✅ | Drafted — content exists, needs polish |
| 🔄 | Discussed in depth — content written below |
| ⬜ | Not started |

---

# OUTLINE

*Each section links to written content below. Sections without links are not yet written.*

---

### [Chapter 0 — The complete map: build systems at a glance](#ch0) 🔄
*A full-picture orientation before the deep dive. Covers the two generations, the major systems, interview-ready architecture, the five critical design decisions, and the comparison table. Read this first for the lay of the land; read the rest of the book to understand why.*
- [0.1 The core theory: two generations of build systems](#ch0-1)
- [0.2 The major systems, one by one](#ch0-2)
- [0.3 System architecture: what you would build in an interview](#ch0-3)
- [0.4 The critical design decisions — five interview questions](#ch0-4)
- [0.5 Comparison table at a glance](#ch0-5)

---

## Part I — Foundations

### [Chapter 1 — What is a build system?](#ch1) ✅
- [1.1 The gap between source code and a running program](#ch1-1)
- [1.2 The progression of pain — why build systems exist](#ch1-2)
- [1.3 The four jobs of a build system](#ch1-3)
- [1.4 What happens without a build system](#ch1-4)
- [1.5 What a build system gives you — the concrete transformation](#ch1-5)
- [1.6 Why multiple files? Why the dependency graph reflects your organization](#ch1-6)

### [Chapter 2 — Why the dependency graph is everything](#ch2) ✅
- [2.1 The critical path — the hard floor](#ch2-1)
- [2.2 Wide flat graph vs. deep chain](#ch2-2)
- [2.3 Cycles: the dependency pattern that makes the build impossible](#ch2-3)
- [2.4 Change propagation](#ch2-4)
- [2.5 Real numbers: what graph shape means at scale](#ch2-5)
- [2.6 Why Google tracks critical path length as a metric](#ch2-6)
- [2.7 Concurrent changes: independent files vs. a chain](#ch2-7)
- [2.8 The combinatorial explosion of cache misses](#ch2-8)

---

## Part II — How a Build System Works Internally

### [Chapter 3 — The two generations of build systems](#ch3) 🔄
- [3.1 Generation 1: task-based systems](#ch3-1)
- [3.2 Make in depth — the canonical example](#ch3-2)
- [3.3 The Build Engineer — a named, essential organizational role](#ch3-3)
- [3.4 Generation 2: artifact-based systems](#ch3-4)
- [3.5 The three properties of a modern build system](#ch3-5)
- [3.6 Why these properties are hard to achieve simultaneously](#ch3-6)
- [3.7 What Bazel changed about the build engineering role](#ch3-7)

### Chapter 4 — Dependency resolution and graph traversal ⬜
- 4.1 How BUILD files declare targets and dependencies
- 4.2 Loading phase: parsing Starlark / build files
- 4.3 Building the target graph in memory
- 4.4 Graph traversal algorithms — topological sort, BFS, DFS
- 4.5 Handling large graphs efficiently
- 4.6 Target graph vs. action graph

### [Chapter 5 — Change detection and content hashing](#ch5) 🔄
- [5.1 Why timestamps fail](#ch5-1)
- [5.2 Content-addressed storage — the core idea](#ch5-2)
- [5.3 Computing a fingerprint for a build action](#ch5-3)
- [5.4 The cache key: anatomy and rationale](#ch5-4)
- [5.5 Cache invalidation — the hardest problem](#ch5-5)
- [5.6 Merkle trees](#ch5-6)
- [5.7 When is a stale cache hit truly impossible — unpacking the claim](#ch5-7)

### Chapter 6 — Execution: parallelism and sandboxing ⬜
- 6.1 Scheduling: topological sort + priority queue
- 6.2 Critical path scheduling
- 6.3 Local parallelism: worker pools
- 6.4 Sandboxing: why isolation matters
- 6.5 Handling non-deterministic actions
- 6.6 Execution failures and retry semantics

### [Chapter 7 — Caching: local, remote, and distributed](#ch7) 🔄
- [7.1 Local disk cache](#ch7-1)
- [7.2 How content addressing keeps the cache correct automatically](#ch7-2)
- [7.3 The two failure modes that look identical from the outside](#ch7-3)
- [7.4 Environmental leaks in depth](#ch7-4)
- [7.5 Remote cache — shared across engineers and CI](#ch7-5)
- [7.6 The remote build execution (RBE) protocol](#ch7-6)
- [7.7 Cache poisoning](#ch7-7)
- [7.8 When local-vs-CI divergence doesn't matter](#ch7-8)
- [7.9 Garbage collection and cache eviction](#ch7-9)

### [Chapter 8 — Build systems at team scale](#ch8) 🔄
- [8.1 The fundamental tension: concurrent work on a shared graph](#ch8-1)
- [8.2 The stable interface pattern](#ch8-2)
- [8.3 Module ownership as a build performance tool](#ch8-3)
- [8.4 Keeping modules small and focused](#ch8-4)
- [8.5 The submit queue — serializing merges](#ch8-5)
- [8.6 Trunk-based development](#ch8-6)
- [8.7 The remote cache as a coordination mechanism — and its limits](#ch8-7)

### [Chapter 9 — Version control philosophy and its build consequences](#ch9) 🔄
- [9.1 The foundational question: where does your dependency live?](#ch9-1)
- [9.2 Google's model: Piper, G4, and the monorepo](#ch9-2)
- [9.3 What the monorepo makes possible](#ch9-3)
- [9.4 What the monorepo requires](#ch9-4)
- [9.5 Amazon's model: polyrepo, service ownership, artifact versioning](#ch9-5)
- [9.6 What polyrepo makes possible — and what it breaks](#ch9-6)
- [9.7 Git's philosophy and its build consequences](#ch9-7)
- [9.8 The mismatch problem](#ch9-8)
- [9.9 The philosophical divide: a comparison](#ch9-9)

### [Chapter 10 — Theory meets reality: why even Google has slow builds](#ch10) 🔄
- [10.1 What the theory actually promises](#ch10-1)
- [10.2 The critical path minimum is still real, and can be very large](#ch10-2)
- [10.3 Interface changes: the worst case, and they happen](#ch10-3)
- [10.4 Tests are the real bottleneck, not compilation](#ch10-4)
- [10.5 The submit queue creates its own serialization latency](#ch10-5)
- [10.6 Distributed execution has overhead of its own](#ch10-6)
- [10.7 Some work genuinely cannot be cached](#ch10-7)
- [10.8 Clean builds: when all the theory hits zero](#ch10-8)
- [10.9 Mitigations for clean build latency](#ch10-9)
- [10.10 The honest picture](#ch10-10)

### [Chapter 11 — Optimizing the build: keeping the graph healthy over time](#ch11) 🔄
*Chapter 10 explains why builds are slow. This chapter explains what you do about it — the concrete interventions an engineer or architect applies to the code, the dependency graph, and the build configuration to manage both clean and incremental build time. The central insight: build optimization is not a one-time design exercise. The graph degrades continuously through individually rational decisions, and maintaining build health requires ongoing, increasingly automated intervention.*
- [11.1 The build graph degrades — and it's nobody's fault](#ch11-1)
- [11.2 The three stages of growth: when different interventions matter](#ch11-2)
- [11.3 Modularization: the primary technique and its real costs](#ch11-3)
- [11.4 Beyond modularization: the full optimization toolkit](#ch11-4)
- [11.5 The restructuring problem: why optimization doesn't happen](#ch11-5)
- [11.6 Automated detection: monitoring graph health at scale](#ch11-6)
- [11.7 Where AI changes the economics](#ch11-7)

---

## Part III — Industry Systems

### Chapter 12 — Google: Blaze and Bazel ⬜
- 12.1 The problem: one monorepo, tens of thousands of engineers
- 12.2 Core architecture: client-server model
- 12.3 Starlark: why a restricted Python dialect
- 12.4 The sandbox: symlink forests and hermetic compilation
- 12.5 Remote caching and RBE at Google scale
- 12.6 The submit queue (TAP)
- 12.7 Bazel as open source: what was kept, what was left out
- 12.8 Strengths and weaknesses

### Chapter 13 — Meta: Buck and Buck2 ⬜
- 13.1 Why Buck1 existed and why it needed replacing
- 13.2 Buck2's architectural innovations
- 13.3 Starlark in Buck2: the prelude and user rules
- 13.4 Dynamic dependencies
- 13.5 Remote execution and hermeticity in Buck2
- 13.6 Comparison with Bazel

### Chapter 14 — Beyond the monorepo giants: Gradle, Android, and the JavaScript ecosystem 🔄
- [14.1 Maven, Gradle, and Android — three things that look the same and aren't](#ch13-1)
- [14.2 Gradle and the Android build: why it is slow and why it is hard to fix](#ch13-2)
- [14.3 What Google uses internally for Android — and why it's different](#ch13-3)
- 14.4 The JavaScript monorepo problem
- 14.5 Turborepo: minimalism and speed
- 14.6 Nx: the full platform
- 14.7 When to use which
- 14.8 The Rust migration trend

---

## Part IV — Building Your Own

### Chapter 15 — Systems design interview: designing a build system ⬜
- 15.1 Clarifying questions
- 15.2 Core components and interfaces
- 15.3 The dependency graph
- 15.4 The cache
- 15.5 The execution engine
- 15.6 Handling non-determinism and failures
- 15.7 Scaling to distributed execution
- 15.8 Handling team scale

### Chapter 16 — Implementing a minimal build system ⬜
- 16.1 Design decisions
- 16.2 Representing the dependency graph
- 16.3 Computing content hashes
- 16.4 Topological sort and scheduling
- 16.5 A simple local cache
- 16.6 Running actions in isolated subprocesses
- 16.7 End-to-end walkthrough
- 16.8 What to add next

### [Chapter 17 — The permanent foundations: what will not change in the AI era](#ch16) 🔄
- [17.1 Build systems are applied computer science, not technology fashion](#ch16-1)
- [17.2 Graph theory: the dependency graph is a DAG, forever](#ch16-2)
- [17.3 Hashing and content addressing: from cryptography and distributed systems](#ch16-3)
- [17.4 Parallelism and scheduling theory: Amdahl's Law and critical path](#ch16-4)
- [17.5 Caching theory: the fundamental space-time tradeoff](#ch16-5)
- [17.6 Formal language theory: why the BUILD file must be restricted](#ch16-6)
- [17.7 Distributed systems: the remote cache is a distributed store](#ch16-7)
- [17.8 Operating systems: sandboxing is kernel-level least privilege](#ch16-8)
- [17.9 What AI changes — and what it doesn't](#ch16-9)
- [17.10 Why this knowledge does not depreciate](#ch16-10)

---

## Appendices

### Appendix A — Glossary ⬜
### Appendix B — Comparison table ⬜
### Appendix C — Further reading ⬜
### [Appendix D — A history of build systems: milestones, ideas, and turning points](#appendix-d) 🔄

---
---

# CONTENT

---

<a id="ch0"></a>
## Chapter 0 — The complete map: build systems at a glance

*This chapter is a deliberate shortcut. It presents the entire landscape — theory, major systems, architecture, design decisions, and comparison table — as a single coherent overview. It is written for the engineer who wants to orient themselves quickly, or who is preparing for a systems design interview and needs the full picture before drilling into any one area. Every claim made here is explained in depth in the chapters that follow.*

---

<a id="ch0-1"></a>
### 0.1 The core theory: two generations of build systems

The entire build systems field splits on one fundamental design decision: do you tell the system *what to run* (tasks), or *what to produce* (artifacts)?

**Generation 1 — Task-based** (Make, Ant, Maven, Gradle, npm scripts)

You write imperative scripts. "Run this command, then that one." The system has no insight into what the task actually does, so it cannot know when to skip it, parallelize it safely, or cache it. As teams grow, builds become slow, fragile, and non-reproducible. The build engineer role existed precisely because these systems required constant human intervention to function correctly at scale.

**Generation 2 — Artifact-based** (Blaze/Bazel, Buck2, Pants)

You declare *what you want to produce* and *what it depends on*. The system owns the execution graph. It can cache, parallelize, and distribute builds because it understands the semantics of every step. The key insight is taking power out of engineers' hands and giving it to the system — not running arbitrary tasks but producing declared artifacts from declared inputs.

Three properties define a modern artifact-based build system:

**Hermeticity** — a build action can only see its explicitly declared inputs. If you forget to declare a header file, the build fails rather than silently succeeding on your machine and failing on CI. Bazel achieves this with a sandbox: when it performs a compilation, it creates a new directory filled only with symlinks to the declared input dependencies. Undeclared inputs are structurally inaccessible, not merely discouraged.

**Incrementality** — the system tracks a content-addressed dependency graph. If inputs haven't changed (same bytes, same hash), outputs are served from cache. This applies locally, across a team via a shared remote cache, and across CI runs. The cache key is a hash of all declared inputs: source content, compiler binary, compiler flags, transitive dependency hashes, environment variables, target platform.

**Reproducibility** — given the same inputs, the build always produces the same outputs, on any machine, at any time. This is a direct consequence of hermeticity (the build can't see undeclared inputs) and content addressing (same inputs always produce the same key). It is what makes remote caching safe: if engineer A and engineer B have the same declared inputs, they can share the result.

---

<a id="ch0-2"></a>
### 0.2 The major systems, one by one

**Google — Blaze (internal) / Bazel (open source, 2015)**

The problem: one of the world's largest monorepos — billions of lines of code, tens of thousands of engineers. The motivation was a build system that provides both speed and correctness at that scale.

Bazel is the canonical artifact-based system. `BUILD` files written in Starlark (a restricted Python dialect) declare targets and their dependencies. Bazel constructs an action graph, computes a content hash for every node, and only re-executes a node when its inputs change. The client-server architecture keeps one long-running server per workspace that maintains the build graph in memory across builds, avoiding cold-start overhead.

Remote execution is the real multiplier: build results are stored in a content-addressable shared cache. If any engineer built `//auth:server` this morning, every other engineer gets it from cache this afternoon. The team's collective build history is a shared asset.

**Meta — Buck (2013) / Buck2 (2022)**

The problem: Meta's monorepo spans C++, Python, Rust, Kotlin, Swift, Haskell, OCaml and more. Buck1 hit fundamental scaling limits after a decade.

Buck2's key architectural advances over both Buck1 and Bazel: it is written in Rust (for performance and correctness), it has complete core/rules separation (the build system core has no knowledge of any language — language rules are written in Starlark separately from the core, which means rules can evolve without rebuilding the core), and it runs on a single incremental dependency graph rather than separate loading/analysis/execution phases. Eliminating phases removes whole categories of bugs and increases parallelism. Meta's internal benchmarks show Buck2 completing builds 2× faster than Buck1.

**JS Ecosystem — Turborepo (Vercel) / Nx (Nrwl)**

These apply artifact-based thinking to the JavaScript/TypeScript monorepo, where the unit of work is an npm package rather than a compilation target.

Turborepo's philosophy: radical minimalism. Define a `turbo.json` pipeline declaring task dependencies (`build` depends on `^build` — the same task in all upstream packages). Turborepo hashes inputs, caches outputs, runs tasks in parallel where the graph permits, and integrates with Vercel's remote cache. It does one thing fast and stays out of the way.

Nx's philosophy: the full platform. Beyond caching and parallelism, Nx provides project graph visualization, enforced module boundaries, code generators, distributed task execution across CI machines, and a plugin ecosystem. The core is written in Rust for speed. Nx is what you reach for when you want Bazel-like architectural enforcement in a JavaScript codebase without adopting Bazel.

The practical distinction: Turborepo gives you an incredibly fast car. Nx gives you the car, the map, the guardrails, and the mechanic.

---

<a id="ch0-3"></a>
### 0.3 System architecture: what you would build in an interview

A modern build system has four phases that execute in strict order, and a cache layer that intercepts between analysis and execution:

```
Source files + BUILD files
        │
        ▼
┌─────────────────┐
│   LOAD PHASE    │  Parse BUILD files, evaluate Starlark macros,
│                 │  construct raw target graph
└────────┬────────┘
         │
         ▼
┌─────────────────┐
│ ANALYSIS PHASE  │  Traverse target graph → produce action graph.
│                 │  Each action = (inputs, command, expected outputs).
│                 │  Content hasher computes fingerprint per action.
│                 │  Fingerprint = cache key.
└────────┬────────┘
         │
         ▼  ◄── Cache lookup here: fingerprint → hit or miss
┌─────────────────┐
│ EXECUTION PHASE │  Scheduler topologically sorts action graph,
│                 │  identifies critical path, dispatches ready
│                 │  actions to executors.
│                 │
│  Local executor │  hermetic sandbox (Linux mount namespaces /
│                 │  macOS sandbox-exec)
│  Remote executor│  RBE protocol over gRPC — same protocol that
│                 │  Bazel, Buck2, Pants all speak
└────────┬────────┘
         │
         ▼
┌─────────────────┐
│   CACHE LAYER   │  Content-addressable store (CAS).
│                 │  Local disk → remote (S3/GCS/custom).
│                 │  Hit: download artifact, skip execution.
│                 │  Miss: execute, upload result to both caches.
└─────────────────┘
```

**The load phase** parses BUILD files and evaluates Starlark macros. The Starlark evaluator is sandboxed — no I/O, no randomness — so this phase is deterministic and cacheable. Lazy evaluation: only load BUILD files for the transitive closure of what was requested, not the entire repository.

**The analysis phase** traverses the target graph and produces an action graph. Each action is a tuple of `(inputs, command, expected_outputs)`. The content hasher computes a fingerprint for each action node. This fingerprint is the cache key. Cycle detection runs during this phase using DFS on the completed graph — if a cycle exists, the build fails immediately with a precise error before any compilation begins.

**The execution phase** uses a scheduler that topologically sorts the action graph and maintains a priority queue ordered by estimated remaining critical path length — dispatching the work that would block the most downstream actions first. Independent actions execute in parallel up to the worker pool limit. Each local action runs in a hermetic sandbox. Each remote action speaks the RBE (Remote Build Execution) protocol: the client sends `(action_digest, input_root_digest)`, the worker fetches inputs from CAS, executes, and uploads outputs back.

**The cache layer** is a content-addressable store. The cache key is the action fingerprint. Before executing any action, the scheduler queries local cache, then remote cache. A hit returns the stored output immediately. A miss executes and writes the result back to both caches. The cache is logically append-only; eviction is handled by LRU or TTL separately from writes.

---

<a id="ch0-4"></a>
### 0.4 The critical design decisions — five interview questions

These are the questions a good interviewer asks after you sketch the basic architecture. Having precise answers to each is the difference between a passing and a strong performance.

**1. How do you handle non-deterministic actions?**

Actions that embed timestamps, generate random seeds, or depend on the current time cannot be cached safely — running them twice with identical declared inputs produces different outputs. Three strategies: (a) force hermeticity by sandboxing the process, overriding `SOURCE_DATE_EPOCH` to a fixed value, stripping metadata from archives; (b) for truly non-deterministic actions, mark them as non-cacheable — they always execute, never write to cache; (c) use double-build verification (run the action twice, compare outputs — if they differ, flag it as flaky and don't cache).

**2. What is your cache invalidation strategy?**

Content-addressed: the key is the hash of all declared inputs plus the command. No time-based expiry — a cache entry is valid forever if the key matches. Staleness is structurally impossible given correct input declarations (it would require a SHA-256 collision). Old entries are garbage collected by LRU or TTL independently of correctness. The only way a cache entry becomes wrong is if an input was not declared — which is a declaration bug, not a cache bug.

**3. How do you handle external dependencies?**

Lock files with cryptographic hashes. Bazel uses `MODULE.lock`; npm uses `package-lock.json`; Cargo uses `Cargo.lock`. The build system validates the hash of every downloaded artifact before using it. If the hash doesn't match the lock file, the build fails. This prevents supply chain attacks (a compromised registry cannot silently substitute a malicious artifact) and ensures reproducibility (the same lock file always produces the same dependency set).

**4. How do you scale the dependency graph for a billion-line codebase?**

Three mechanisms: lazy evaluation (only load BUILD files for the transitive closure of the requested target — not the entire repository), server process (keep the analyzed graph in memory across builds so incremental builds don't re-parse unchanged BUILD files), and incremental graph updates (when a BUILD file changes, only re-analyze the affected portion of the graph, not the whole thing). Buck2's single incremental graph across all phases is the current state of the art here.

**5. How do you distribute work across thousands of machines?**

The RBE protocol over gRPC. The client submits an `Action` — a content-addressed tuple of `(command, input_root_digest, output_paths, platform)`. Workers download inputs from the content-addressable store (CAS) by digest, execute the command in a clean environment, and upload outputs back to CAS. The client retrieves outputs by digest. The action cache maps `action_digest → action_result_digest`, so if any worker has already executed this exact action, the client gets the result without re-executing. All state is in the CAS and action cache, which means workers are stateless and horizontally scalable.

---

<a id="ch0-5"></a>
### 0.5 Comparison table at a glance

| | **Make** | **Gradle** | **Bazel** | **Buck2** | **Turborepo** | **Nx** |
|---|---|---|---|---|---|---|
| Origin | Bell Labs | Open source | Google | Meta | Vercel | Nrwl |
| Generation | 1 — task | 1 — task | 2 — artifact | 2 — artifact | 1.5 — hash-based | 1.5 — hash-based |
| Core language | C | JVM | Java + C++ | Rust | Go → Rust | TypeScript → Rust |
| Rule language | Makefile syntax | Groovy / Kotlin DSL | Starlark | Starlark | turbo.json | JSON + plugins |
| Scope | Any language | JVM / Android | Polyglot monorepo | Polyglot monorepo | JS/TS monorepo | JS/TS + polyglot |
| Hermeticity | None | None | Full sandbox | Remote-exec hermetic | Hash-based only | Hash-based only |
| Remote cache | No | Optional (build cache) | Yes (RBE) | Yes (RBE) | Vercel / self-hosted | Nx Cloud / self-hosted |
| Cycle detection | Partial / broken | Warns | Hard error, pre-execution | Hard error, pre-execution | Warns | Warns |
| Learning curve | Medium | Medium | High | High | Low | Medium |
| Best fit | Small C/C++ projects | JVM / Android apps | Large org, many languages | Meta-scale monorepo | JS teams, fast setup | Enterprise JS platform |

**The two-line summary of where each system sits:**

Make and Gradle are Generation 1: flexible, familiar, and unable to provide safe caching or guaranteed hermeticity because their tasks can do anything.

Bazel and Buck2 are Generation 2: correct by design, fast at scale, high setup cost, and designed for organizations that have outgrown what task-based systems can provide.

Turborepo and Nx occupy the practical middle for JavaScript: they use content hashing for caching (better than task-based) but do not enforce hermeticity (not as correct as Bazel). They are the pragmatic choice for most JS/TS monorepos.

---

*The rest of the book unpacks every claim made in this chapter. Chapter 1 asks why build systems exist at all. Chapter 2 explains why the dependency graph's shape determines build speed. Chapters 3–10 go deep on mechanisms. Chapters 11–15 are for building one yourself.*

---

<a id="ch1"></a>
## Chapter 1 — What is a build system?

**One-line pitch:** Source code is not a program. A build system is the machine that turns one into the other — and you don't need one until, suddenly, you desperately do.

---

<a id="ch1-1"></a>
### 1.1 The gap between source code and a running program

A source file is a text file. It contains instructions written for a human reader in a language with rules a human can understand. A running program is a sequence of machine instructions that a processor executes directly. The distance between these two things — the translation from human-readable text to executable binary — is not small. It involves parsing, type checking, code generation, optimization, linking against other compiled code, and packaging into a format the operating system knows how to load.

For most of computing history, this translation was performed manually. The programmer knew which commands to run, in which order, with which flags. This knowledge lived in their head.

As long as there was one programmer, working on one file, this was fine.

It stopped being fine quickly.

---

<a id="ch1-2"></a>
### 1.2 The progression of pain — why build systems exist

The need for a build system does not arrive all at once. It accumulates through a series of specific, concrete problems. Each is tolerable in isolation. Together they become unmanageable. Understanding each stage of the progression is the clearest way to understand what a build system actually does and why each of its features exists.

**Stage 1: one file, no problem.**

```c
// hello.c
#include <stdio.h>
int main() {
    printf("hello, world\n");
    return 0;
}
```

You compile it:

```bash
gcc hello.c -o hello
```

You run it. It works. You need nothing else. No Makefile. No BUILD file. No build system. A build system for a single file would be pure overhead — ceremony with no benefit. This is the baseline: the build system earns its existence by solving problems that emerge as the codebase grows.

**Stage 2: multiple files, manual ordering.**

The program grows. You split it into `main.c`, `utils.c`, and `math.c`. Now you need to compile all three and link them together:

```bash
gcc -c utils.c -o utils.o
gcc -c math.c  -o math.o
gcc -c main.c  -o main.o
gcc utils.o math.o main.o -o program
```

This still works. But two new problems have appeared, even if they're not painful yet.

First, you must remember the correct order. If `main.c` depends on functions in `utils.c` and `math.c`, the libraries must exist before the linker runs. For three files this is obvious. For three hundred files it is not.

Second, if you change only `utils.c`, you are still recompiling `main.c` and `math.c`. Their compiled output is identical to what it was before your change. You're throwing away work for no reason. For three files this costs milliseconds. For three thousand files it costs minutes.

**Stage 3: the silent staleness problem.**

This is where things get genuinely dangerous.

You add a header file, `utils.h`, that both `main.c` and `utils.c` include. It defines a struct:

```c
// utils.h
typedef struct {
    int x;
    int y;
} Point;
```

You change the struct to add a field:

```c
typedef struct {
    int x;
    int y;
    int z;  // new field
} Point;
```

You recompile `utils.c` because you edited it. You do not recompile `main.c` because you forgot it includes `utils.h`. The linker succeeds — both object files compile cleanly. You ship a binary where `main.c` was compiled against the old definition of `Point` and `utils.c` was compiled against the new one. The struct's memory layout is different in two parts of the same binary.

This does not produce a compile error. It may not produce a runtime error immediately. It produces a subtle, intermittent, hard-to-diagnose bug that surfaces weeks later under specific conditions. The build system did not fail — it did exactly what you told it to do. The problem is that what you told it to do was wrong, because you were managing dependencies by hand and you made a mistake.

Managing the full transitive set of dependencies across a codebase of any real size by hand is not possible. Humans make mistakes. Build systems do not.

**Stage 4: multiple engineers, divergent environments.**

A second engineer joins the project. They clone the repository and start working. They change `math.c`. They compile their changes and run their tests. Everything passes. They push their code.

You pull their changes and continue working on `main.c`. You compile `main.c` — which you last compiled yesterday, before their change — and link everything together. Your `math.o` was produced before their change to `math.c`. Your binary contains your new `main.c` linked against their old `math.c`. The binary may be wrong. You may not find out until you test it, or until a user finds the bug, or never, if the interaction between the two changes is subtle enough.

"It works on my machine" is not a flip remark. It is the name of a real failure mode that emerges directly from manual dependency management in a multi-engineer environment. Each engineer's machine has a different build history — a different set of compiled objects reflecting different sequences of changes. The builds are not reproducible across machines.

**Stage 5: multiple platforms.**

The program needs to run on Linux and macOS. The compiler flags are different. The library paths are different. The system headers are in different locations. The dynamic linker behaves differently.

You write a shell script to manage this:

```bash
if [ "$(uname)" = "Linux" ]; then
    CFLAGS="-fPIC -lpthread"
elif [ "$(uname)" = "Darwin" ]; then
    CFLAGS="-dynamiclib"
fi
gcc $CFLAGS -c utils.c -o utils.o
```

This works for two platforms. You add a third. You add a cross-compilation target. You add debug and release build variants. The script grows. Platform detection logic accumulates. Edge cases multiply. A new engineer tries to build on a platform they've never tried before and spends a day debugging the script rather than writing code.

**Stage 6: scale.**

The codebase reaches 500 files. A full recompile takes 20 minutes. You cannot wait 20 minutes every time you change one file. You must recompile incrementally — only the files that changed and the files that depend on the changed files.

But you cannot know which files depend on which other files unless you track the dependency graph explicitly. Without explicit tracking, you have two choices: recompile everything (20 minutes, always correct) or recompile only what you remember changed (fast, sometimes wrong). Neither is acceptable for a large project with many engineers.

**The conclusion: build systems exist to make implicit information explicit.**

At each stage of the progression, the underlying problem is the same: the information needed to build correctly — what depends on what, what changed, what needs to be recompiled, in what order, with what flags — is implicit. It exists in the engineer's head, in documentation, in convention. Human memory is fallible. Documentation goes stale. Convention breaks down as teams grow.

A build system makes this information explicit, in a form the machine can read and reason about. It converts the implicit dependency graph into a declared one. It detects changes mechanically rather than relying on human memory. It manages ordering, parallelism, and caching automatically. The engineer writes what needs to be built. The system works out how.

---

<a id="ch1-3"></a>
### 1.3 The four jobs of a build system

Every build system, from Make to Bazel to Buck2, performs four fundamental jobs. Understanding what these jobs are, and why each one is necessary, maps directly to the progression described above.

**Job 1: Dependency resolution.**

Given a target the engineer wants to build — a binary, a library, a test suite — the build system must determine everything that target depends on, transitively. `main.c` depends on `utils.h` and `math.h`. `utils.h` depends on nothing. `math.h` depends on `vectors.h`. The build system must discover this full graph before it can decide what to build.

This is the answer to Stage 3 and Stage 6. Without explicit dependency resolution, engineers manage the graph by hand and make mistakes. With it, the system knows the complete transitive closure of every dependency and can correctly determine what needs rebuilding when anything changes.

**Job 2: Change detection.**

Given the dependency graph, the build system must determine which parts of it are stale — which artifacts need to be rebuilt because one of their inputs changed since they were last produced.

Naive change detection uses file modification timestamps: if the source file is newer than the object file, recompile. This is what Make uses and it fails in multiple ways (Section 3.2). Correct change detection uses content hashing: if the content of the source file has changed — regardless of when it was touched — the object file is stale and must be rebuilt. Same content, same hash, no rebuild needed.

This is the answer to Stage 2 and Stage 3. The build system knows not just what the graph looks like but whether any node in the graph has actually changed.

**Job 3: Execution.**

Given what needs to be built and in what order, the build system runs the necessary commands — compiler invocations, linker calls, code generation scripts — in the correct dependency order, as much in parallel as the dependency graph permits.

This is the answer to Stage 1 and Stage 6. The build system owns the sequencing and parallelization. The engineer doesn't need to know the correct order or manage which steps can run simultaneously.

**Job 4: Caching.**

Given that the same build action has been run before with identical inputs, the build system can skip running it again and return the previously-computed result. This applies both locally (on the engineer's machine, build-over-build) and remotely (across all engineers on a team, so that work done by one engineer is reused by others).

This is the answer to Stage 2 and the multi-engineer problem of Stage 4. Caching makes builds fast. Remote caching makes the team's collective build history into a shared asset.

---

<a id="ch1-4"></a>
### 1.4 What happens without a build system

The most direct way to appreciate a build system is to observe what software development looks like without one.

In the earliest commercial software projects, build procedures were documented in prose: a README or a runbook that described the sequence of commands needed to produce a working binary. New engineers read the document, typed the commands, and hoped the document was accurate. If the document was wrong — if a step was missing, or the order was wrong, or the flags had changed since the document was last updated — the build failed with a confusing error that required an experienced engineer to diagnose.

Build knowledge was an oral tradition. Senior engineers knew which incantations worked and which failed. Onboarding a new engineer took days or weeks, a significant fraction of which was spent learning the build. This cost was accepted as normal.

Without automated change detection, engineers either rebuilt everything on every change — slow — or rebuilt only what they remembered changing — fast but unreliable. The choice between correctness and speed was made manually, per build, based on the engineer's memory of what they'd changed and what depended on it. Engineers got good at remembering. They also made mistakes.

Without dependency tracking, the bugs described in Stage 3 above — shipping binaries where different translation units were compiled against different versions of a shared header — were common. They were hard to diagnose precisely because they didn't produce compile errors. The code compiled. The binary ran. It just behaved incorrectly in specific circumstances.

As a result, many projects maintained a ritual: the **clean build**. When anything was uncertain, you deleted all compiled artifacts and rebuilt from scratch. This guaranteed correctness — if you compile everything fresh, there can be no stale artifacts — but it meant that fast incremental builds were unavailable as a safety option. The clean build was slow, but it was trusted. The incremental build was fast, but it might be wrong. Teams that had been burned enough times chose slow and trusted.

The build system's promise is: you should never have to choose. The incremental build should be as trustworthy as the clean build, and fast.

---

<a id="ch1-5"></a>
### 1.5 What a build system gives you — the concrete transformation

With a build system, the progression from Stage 1 to Stage 6 looks different at each step.

**Multiple files:** The build system reads the dependency declarations and compiles exactly the files that need compiling, in the correct order, without the engineer specifying the order. A change to `utils.c` recompiles `utils.c`. Files that don't depend on `utils.c` are untouched.

**Headers and transitive dependencies:** The build system tracks header dependencies, either through explicit declarations (Bazel's `hdrs`) or through compiler-generated dependency files (Make with `-MD`). A change to `utils.h` triggers recompilation of every file that includes it, directly or transitively. The silent staleness bug of Stage 3 becomes impossible because the system, not the engineer, is tracking the graph.

**Multiple engineers:** When engineers share a remote cache, each engineer's successful build populates a cache that all other engineers can query. If engineer A compiled `math.c` this morning, engineer B does not compile `math.c` this afternoon — they download engineer A's result. The team's collective build history becomes a shared asset that grows more valuable as the team grows. "It works on my machine" becomes diagnosable: if a build works locally but fails on CI, the discrepancy points to a real dependency that is present in one environment and absent in another — a bug to fix, not a mystery to tolerate.

**Multiple platforms:** The build system owns platform abstraction. Compiler selection, flag management, and library path configuration are expressed as build system configuration rather than shell script conditionals. The engineer declares what to build; the build system determines how to build it for the current platform. Adding a new platform means updating the build system configuration, not auditing every shell script in the repository.

**Scale:** At hundreds or thousands of files, the build system's dependency graph, change detection, and caching make incremental builds fast and correct simultaneously. A change to one file rebuilds exactly the affected subset of the codebase. A change to a widely-used header triggers a larger rebuild — but only as large as the actual transitive impact, no larger.

The build system does not change what work must be done. It changes who is responsible for knowing what that work is. Before: the engineer. After: the system. The engineer's cognitive load drops. The system's correctness guarantees go up. The build becomes something you can trust rather than something you periodically doubt.

---

<a id="ch1-6"></a>
### 1.6 Why multiple files? Why the dependency graph reflects your organization

Everyone accepts that real software lives in multiple files. Nobody stops to ask why. The answer is not obvious — it is the product of at least four distinct forces, and understanding those forces explains why the dependency graph looks the way it does, and why the build system's job is inseparable from the organizational structure of the team that writes the code.

#### The cognitive reason: the file as a unit of human comprehension

The most immediate reason for multiple files is that a human being cannot hold an entire program in working memory at once. Cognitive science research consistently shows that human working memory can actively track roughly seven chunks of information simultaneously. A program of any real complexity has thousands of interdependent concepts.

The file is a unit of cognitive scope. When you open `auth/tokens.c` you are imposing a boundary on what you need to think about. You don't need to hold database connection logic or HTTP routing in your head while reading it. The file tells the reader: everything relevant to token management is here. Everything else is somewhere else.

This is not a technical requirement. It is a psychological one. You could write an entire operating system in one file. The compiler does not care. But no human engineer could navigate, modify, or reason about it. The file boundary is a tool for human comprehension first, and everything else second.

#### The compilation reason: the historical accident that still shapes everything

In the earliest days of compiled languages, programs were split into multiple files for a reason that has nothing to do with human comprehension: early computers did not have enough RAM to compile a large program all at once. The compiler itself needed memory to run. Programs were therefore split into translation units — pieces small enough that the compiler could load one, compile it to an object file, and discard it before loading the next.

This gave us the compilation model that C and C++ still use today: each `.c` or `.cc` file is an independent translation unit. The compiler processes it in isolation, produces an object file, and moves on. The linker combines all the object files at the end.

This 60-year-old memory constraint is the direct origin of the header file. If `main.c` calls a function defined in `utils.c`, the compiler needs to know the function's signature when it compiles `main.c` — even though it has not compiled `utils.c` yet. The header file solves this: it declares the signature without defining the implementation, allowing the compiler to type-check the call in `main.c` while deferring the actual compilation of `utils.c`.

The header file is not an organizational nicety. It is a direct technical consequence of "compile one file at a time," which was itself a direct consequence of 1960s memory constraints. The entire C and C++ build model — the include problem, the need for explicit dependency declarations, the reason Makefile engineers spent so much time tracking transitive header dependencies — flows from this single ancient constraint. Modern languages designed without this constraint (Go, Rust, Java) handle it differently, which is one reason their build times are dramatically faster for equivalent code size.

#### The organizational reason: files as units of parallel work and ownership

Multiple files are how multiple engineers work on the same program simultaneously without constantly overwriting each other's changes.

If the entire program is one file, two engineers editing it simultaneously will produce merge conflicts on almost every commit. As the team grows, contention on a single file becomes the bottleneck for the entire team's ability to ship. This is not a theoretical concern — it is the first concrete pain point that growing teams hit when they try to work in an undivided codebase.

Files are units of ownership. Engineer A owns `auth/tokens.c`. Engineer B owns `network/http.c`. They work simultaneously, commit independently, review each other's changes in isolation. The file boundary is a coordination mechanism as much as a technical one.

This organizational function of the file connects directly to a principle that Melvin Conway articulated in 1967:

> *Any organization that designs a system will produce a design whose structure is a copy of the organization's communication structure.*

Conway's Law is not a prediction — it is closer to a law of physics for software. Engineers design interfaces at the boundaries between their code and other people's code, and those boundaries naturally reflect their communication patterns. Team A and Team B communicate through a defined API. The API reflects the boundary between their teams. The organizational chart and the dependency graph mirror each other.

This has a direct and often underappreciated consequence: **the build system is partly an organizational tool, not just a technical one.**

When Bazel's `visibility` attribute lets a team declare which other packages may depend on their library, it is enforcing an organizational boundary at the technical level:

```python
cc_library(
    name = "token_store",
    srcs = ["token_store.cc"],
    hdrs = ["token_store.h"],
    visibility = ["//auth/..."],  # only the auth package tree may depend on this
)
```

This declaration says: this code is the auth team's internal implementation. Other teams may not take a dependency on it without an explicit decision to change the visibility. The build system enforces the Conway boundary structurally. A dependency that crosses an organizational line in the wrong direction fails the build — not in code review, where it depends on a reviewer noticing, but at the point of declaration, automatically, on every engineer's machine and every CI run.

#### The build performance reason: files as units of change detection

Once you have separate compilation, you have the possibility of incremental builds. If you change only `utils.c`, only `utils.c` needs to be recompiled — along with anything that depends on it. Everything else can reuse its cached compilation result.

This means the file boundary is not just a human or organizational tool. It is the unit of change detection in the build system. The finer-grained your files, the more precisely the build system can identify what changed. A large file that contains authentication logic, database logic, and HTTP logic means any change to any of the three invalidates the compiled output of all of them. Splitting into three files means changing authentication doesn't force a recompile of the database logic.

The decision of how to split code into files is therefore simultaneously:

- A **cognitive** decision: what fits in one mental scope?
- A **compilation** decision: what is one translation unit?
- A **organizational** decision: who owns what?
- A **build performance** decision: what is the unit of change detection and caching?

These four forces do not always agree. A file that is the right size for human comprehension may be too large for efficient incremental builds. A file boundary that matches team ownership may not match compilation efficiency. The tension between them is one of the fundamental design challenges of software architecture at scale — and it plays out directly in how BUILD files and Makefiles are written.

#### The module as the real unit: above the file, below the repository

In practice, neither the individual file nor the entire repository is the right unit of software organization for large systems. The file is too small — it is an implementation detail, not an interface boundary. The repository is too large — it contains everything, which means it expresses no organizational structure at all.

The intermediate unit — the **module**, **package**, or **library** — is the real building block of large software systems.

In Bazel terms, the BUILD file defines a package: a directory of related files that are built, versioned, and depended upon together as a unit. The `cc_library` or `java_library` target is the unit of dependency. When `//auth:client` depends on `//network:http`, it is a team-level dependency between the auth package and the network package — not a file-level dependency between individual source files.

This abstraction separates two things that are often conflated:

- The **organizational unit** (the package, owned by a team, with a declared interface that other packages depend on)
- The **compilation unit** (the individual file, which the compiler processes in isolation)

Internal refactoring within a package — splitting a file, merging two files, reorganizing code — is invisible to external callers as long as the package's declared interface doesn't change. The package's `hdrs` (public headers in C++, exported symbols in Go, `pub` items in Rust) define what external callers may depend on. Everything else is implementation detail, free to change without affecting callers.

The build system enforces this abstraction. Files within a package that are not declared as public are not accessible to code outside the package — the sandbox blocks the access. The package boundary is real and enforced, not merely conventional. This is the technical enforcement of Conway's Law: the organizational boundary between teams is expressed in the BUILD file and enforced by the build system.

#### Why this matters for everything that follows

The dependency graph is not an artifact of build system design. It is the shape of the organization's communication structure, made legible and mechanically enforceable. It reflects:

- How engineers have divided cognitive scope (files)
- How the compiler model requires separate translation units (headers)
- How teams divide ownership (package boundaries)
- How change propagates through the codebase (what depends on what)

The build system's job is to manage the consequences of these organizational choices: to know the full graph, to detect changes accurately, to rebuild exactly what is necessary, and to enforce the boundaries that prevent the graph from degenerating into the tangled, untestable, unrefactorable mess that every large codebase tends toward without active maintenance.

This is why the dependency graph is the central object in build system design, and why the next chapter treats its shape as the primary determinant of build performance. The graph is not just a data structure. It is a map of how the organization works.

---

<a id="ch2"></a>
## Chapter 2 — Why the dependency graph is everything

**One-line pitch:** The shape of your dependency graph sets a hard lower bound on your build speed that no amount of hardware can break — and it determines your entire team's integration throughput, not just your own build time.

---

<a id="ch2-1"></a>
### 2.1 The critical path — the hard floor

Every build has a critical path: the longest chain of sequential dependencies from a source file to the final target. Steps on the critical path cannot be parallelized — each one must wait for the previous one to complete before it can begin. No amount of additional hardware changes this.

Suppose your build has the following structure:

```
A → B → C → D → binary
```

Each step takes 10 seconds. The critical path is 40 seconds. If you add 100 machines, the build still takes 40 seconds — because the machines cannot start step B until step A is done, cannot start C until B is done, and so on. The critical path is not a software limitation or a build system limitation. It is a logical constraint imposed by the dependency structure.

Now suppose your build has this structure instead:

```
A ─┐
B ─┤
C ─┼─→ binary
D ─┤
E ─┘
```

All five steps are independent. They can run in parallel. With five machines, the build takes 10 seconds — the time of the single slowest step. Same total work. Radically different wall-clock time.

The difference between these two cases has nothing to do with the build system, the hardware, the compiler, or any other tool. It is determined entirely by the shape of the dependency graph. A build system that understands the graph can exploit the parallelism in the second case automatically. But no build system can create parallelism where the graph provides none.

This is why the dependency graph is the central object in build system design. Before asking "how do we make this build faster," the correct question is always "what does the dependency graph look like, and how long is the critical path?"

---

<a id="ch2-2"></a>
### 2.2 Wide flat graph vs. deep chain — same work, different build times

Consider two codebases that each require compiling 100 files. The total amount of computation is identical. The build times can differ by two orders of magnitude depending on graph shape.

**Codebase A — wide and flat:**
All 100 files are independent of each other. The build system can compile all 100 simultaneously. With sufficient machines, the wall-clock time approaches the compile time of a single file.

**Codebase B — a deep chain:**
File 1 produces a header that file 2 needs. File 2 produces a header that file 3 needs. And so on for all 100 files. The build system must compile them sequentially. The wall-clock time is the sum of all 100 compile times.

If each file takes 2 seconds to compile, Codebase A builds in ~2 seconds. Codebase B builds in ~200 seconds. Same number of files. Same compiler. Same hardware. The 100x difference comes entirely from graph topology.

In practice, real codebases are neither perfectly flat nor perfectly chained — they are somewhere in between, with clusters of independent work connected by shared dependencies. The build system's job is to exploit every available opportunity for parallelism. The architect's job is to design the graph so those opportunities are as numerous as possible.

---

<a id="ch2-7"></a>
### 2.7 Concurrent changes: independent files vs. a chain

The dependency graph doesn't just affect individual build times. It determines your entire team's throughput when multiple engineers are making concurrent changes.

Consider two scenarios, each with 10,000 engineers each making one change:

**Case 1: all 10,000 engineers change independent files.**

The dependency graph is wide and flat. Each change touches a leaf node — a file with no dependents, or dependents limited to that engineer's area of the codebase. CI can validate all 10,000 changes in parallel. Each validation recompiles one file and runs that file's test suite. The remote cache absorbs most of this: many engineers will have already compiled their file locally and uploaded the result, so CI gets a cache hit.

Wall-clock time for the CI system: roughly the time to validate one change. Team throughput: excellent. Remote cache effectiveness: high.

**Case 2: all 10,000 engineers change files on the same dependency chain.**

Engineer 1 changes file A. Engineer 2 changes file B, which depends on A. Engineer 3 changes file C, which depends on B. And so on.

CI cannot validate these changes in parallel. Validating change 2 requires the output of change 1, because B depends on A. Validating change 3 requires the output of change 2. The 10,000 changes must be validated sequentially. Wall-clock time for CI: 10,000 × validation time per change.

More critically, the remote cache cannot help. Engineer 2 built B locally against the original version of A. Engineer 1 changed A. CI must build B against engineer 1's new version of A — a combination that has never been computed by anyone. The remote cache has no entry for this key. Every node downstream of every change point is a guaranteed cache miss on CI, regardless of how much local building the engineers did.

This is not a build system failure. The build system is working correctly — it is correctly determining that the integrated state of these changes has never been built before, and building it from scratch. The problem is the graph topology. The only fix is to restructure the graph.

---

<a id="ch2-8"></a>
### 2.8 The combinatorial explosion of cache misses

The Case 2 scenario above gets worse the more you examine it.

Suppose engineers 1 through 5 each change one file in a chain A → B → C → D → E → binary.

Engineer 1's local cache has entries for: A_new, B (built against original A), C (built against original B), D (built against original C), E (built against original D).

Engineer 2's local cache has entries for: A (original), B_new (built against original A), C (built against original B), D (built against original C), E (built against original D).

CI must build: A_new, then B_new built against A_new, then C built against B_new-built-against-A_new, then D built against that, then E built against that.

The cache key for each of CI's required artifacts includes the hash of its input, which includes the hash of the artifact below it in the chain, which includes the hash of the artifact below that. Every node in the chain produces a unique key — one that was never computed by any engineer's local build, because no local build ever saw this combination of changes.

Engineer 1 built B against the old A. Engineer 2 built B against the new A. CI needs B built against the new A — and only engineer 2 has that, but engineer 2 didn't change A in the same commit, so their B_new is built against the same old A that engineer 1 used.

Every artifact that CI needs is a cache miss. The deeper the chain, the more misses. The more concurrent changes on the chain, the more distinct versions of each artifact exist — none of which match what CI needs.

Adding more CI machines does not help. The chain is sequential. Each link must wait for the previous link to complete. The problem is topological, not computational. The only real fix is to make the graph flatter.

---

<a id="ch2-3"></a>
### 2.3 Cycles: the dependency pattern that makes the build impossible

The three patterns that destroy build performance are the base library trap, overly broad imports, and circular dependencies. The first two make builds slow. The third makes them impossible.

A cycle exists when A depends on B, B depends on C, and C depends on A. In build terms: to build A you need B first, to build B you need C first, to build C you need A first. There is no valid starting point. The build cannot proceed.

In a small codebase this sounds like something only a careless engineer would write. In a large codebase with hundreds of engineers, cycles emerge gradually and accidentally. Nobody creates one intentionally. They form through a series of individually reasonable decisions that happen to close a loop.

The classic formation: Team A owns a utility library. Team B owns a data processing library. One day, an engineer on Team A adds a convenience function that calls one small thing in Team B's library — it seems harmless, saves duplication. Weeks later, an engineer on Team B needs a helper from Team A's library — also seems harmless. Now there is a cycle. Neither engineer saw it coming because each dependency looked local and reasonable in isolation.

**How Make handles cycles — or fails to.**

Make has no principled cycle detection. What happens when you create a cycle depends on the Make implementation and the shape of the cycle.

For a direct self-dependency (`A depends on A`), GNU Make prints a warning — "Circular A ← A dependency dropped" — and continues, silently ignoring the dependency. For indirect cycles (`A → B → C → A`), Make's behavior is implementation-defined. In practice: it may loop indefinitely, exhaust the process stack and crash, silently drop some dependencies and produce a wrong build, or print a confusing error about a missing file that has nothing to do with the actual problem.

The fundamental reason Make cannot properly detect cycles is that it builds its dependency graph lazily and implicitly — interleaving graph construction with execution. There is no phase where the full graph exists in memory and can be analyzed. Make discovers cycles mid-execution rather than at analysis time, by which point it has already begun building things in an incorrect order.

Recursive Make made this dramatically worse. Each sub-Makefile could only see its own directory's dependencies. A cycle that crossed directory boundaries was invisible to every individual Makefile — each looked correct in isolation. Only the combined system had the cycle, and the combined system had no mechanism to inspect itself as a whole.

The symptoms of a cycle in a large Makefile were varied and confusing: infinite recursion crashing the Make process; partial builds where some targets were silently skipped producing a binary that linked but behaved incorrectly; build errors pointing at missing files that were actually caused by incorrect build order; builds that succeeded sometimes and failed other times depending on the order targets were specified on the command line. Because Make did not tell you "there is a cycle," the engineer diagnosed it by inspection — manually tracing dependency chains through thousands of lines of Makefile, looking for the loop. In a Makefile maintained by multiple teams over years, this could take days.

Experienced Makefile engineers developed layering conventions to prevent cycles — library A may only depend on libraries in layer N-1 or below, never upward — and enforced them through code review and periodic manual dependency audits using external tools (`cinclude2dot`, Doxygen's dependency analysis, commercial tools). These audits were run quarterly or before major refactors, and finding a cycle meant a multi-day effort to restructure code across team boundaries. The convention was social. The enforcement was human. It broke down regularly.

**How Bazel handles cycles — at analysis time, as a hard error.**

Bazel detects cycles before executing any build action. This is possible because Bazel strictly separates the analysis phase (constructing the full dependency graph) from the execution phase (running compilers and linkers). During analysis, Bazel loads all BUILD files, evaluates all rules, and builds the complete graph in memory. A standard depth-first search runs on the completed graph. If any node is encountered twice on the same DFS path, a cycle exists.

The build stops immediately with a clear error identifying the complete cycle:

```
ERROR: cycle in dependency graph:
    //auth:client
    //network:http
    //auth:tokens
    //auth:client   ← (self)
```

The engineer sees exactly which targets form the loop and exactly which edge to remove. The build does not proceed until the cycle is resolved. A cycle is not a warning in Bazel — it is a hard error that fails the build at analysis time, before any compilation begins, on every machine and every CI run.

Beyond detection, Bazel's `visibility` system prevents many cycles from forming in the first place. If `//network:http` declares `visibility = ["//network/..."]`, then `//auth:client` cannot take a dependency on it at all — the dependency would be rejected at declaration time, not at cycle-detection time. The organizational boundary is enforced structurally, before the cycle can form.

The contrast with Make is total. Make discovers cycles mid-execution and responds with confusion. Bazel detects them pre-execution and responds with a precise diagnosis. Make engineers spent days finding cycles by hand. Bazel engineers are told exactly what the cycle is, in under a second, the first time they try to build with it.

This is one of the clearest illustrations of what the shift from task-based to artifact-based systems actually buys: not just performance, but the structural ability to reason about the graph before running it.

---

<a id="ch3"></a>
## Chapter 3 — The two generations of build systems

**One-line pitch:** Every build system ever built chooses between telling the machine *what to run* (tasks) or *what to produce* (artifacts). That one decision explains everything — including why large C/C++ shops used to employ specialists whose entire job was keeping the build working.

---

<a id="ch3-1"></a>
### 3.1 Generation 1: task-based systems

The first generation of build systems — Make (1976), Ant (2000), Maven (2004), Gradle (2007) — all share a fundamental design: you tell the system *what commands to run*, and trust that those commands produce the right outputs.

A Makefile entry looks like this:

```makefile
main.o: main.c utils.h
    gcc -c main.c -o main.o
```

This says: "to produce `main.o`, first ensure `main.c` and `utils.h` are up to date, then run this command." The build system tracks whether targets are up to date by comparing file modification timestamps — if the source is newer than the output, rebuild.

The critical limitation is structural: **the build system has no idea what the command actually does.** It cannot inspect the command, predict its outputs, or verify that the outputs are correct. It just runs the recipe and trusts the result. This means:

- It cannot safely cache outputs and share them across machines — it has no way to know if a cached result is valid for a different environment
- It cannot safely parallelize without the engineer manually declaring which targets are independent
- It cannot detect undeclared dependencies — if `main.c` secretly includes a header that isn't listed, the Makefile is wrong but Make cannot tell you

These are not bugs. They are direct consequences of the design. A system that lets you run arbitrary commands cannot reason about what those commands produce.

---

<a id="ch3-2"></a>
### 3.2 Make in depth — the canonical example

Make deserves detailed examination because its limitations are not immediately obvious — it works well at small scale and its failure modes only emerge as the codebase grows.

**The timestamp model.** Make determines whether a target needs rebuilding by comparing the modification time of the output file against the modification times of its declared inputs. If any input is newer than the output, the target is stale and must be rebuilt. This was a reasonable approximation in 1976, when files were only modified by intentional edits. It fails in the modern development environment in three common ways:

- `git checkout` changes file timestamps to the checkout time, not the original modification time. Switching branches can make every file appear to need rebuilding even if the contents are identical to the previous build.
- System clocks can be wrong, especially in distributed or virtualized environments. A clock skew of a few seconds makes timestamp comparison meaningless.
- Copying files from one location to another preserves or changes timestamps in ways Make does not expect. Build scripts that copy intermediate outputs frequently trigger spurious rebuilds.

**The C/C++ include problem.** A single `.c` file may `#include` dozens of headers, each of which may include further headers. The transitive set of included headers can easily exceed 40 files for a modest translation unit. Every one of these headers is a real dependency: if any of them changes, the `.c` file must recompile. Make has no mechanism to discover these dependencies automatically. The engineer must list every header, for every source file, in the Makefile by hand.

In practice, nobody does this. The result is Makefiles that are silently wrong: they compile correctly on a clean build because everything gets rebuilt, but on an incremental build after a header change, some translation units are not rebuilt even though they should be. The binary is stale and the error may not manifest until runtime.

The standard workaround is to use `gcc -M` or `gcc -MD` to auto-generate dependency files, then `include` those files in the Makefile. This works but is fragile: the dependency generation must happen before the build, the generated files must be tracked correctly, and if a header is deleted the generated dependency file still references it, causing Make to error.

**Recursive Make.** Large C/C++ projects are divided into subdirectories, each with its own Makefile, invoked from a top-level Makefile via `$(MAKE) -C subdir`. This pattern feels modular but creates a fundamental problem: each sub-Makefile executes in isolation. It cannot see the dependency graph of sibling directories. It cannot know whether a dependency in another directory is up to date. It makes its decisions with incomplete information, leading to incorrect parallelism and missed rebuilds.

Peter Miller documented this in exhaustive detail in his 1997 paper "Recursive Make Considered Harmful," which argued that the entire recursive Make pattern was broken by design and should be replaced with a single Makefile covering the entire project. The paper was written 28 years ago. Recursive Make is still in widespread use. The problem never went away because fixing it requires restructuring the entire build.

**Platform variance.** The same Makefile must work on Linux, macOS, and historically on Solaris, HP-UX, and AIX. Compiler paths differ. Library locations differ. Compiler flag names differ (what GCC calls `-fpic`, MSVC calls something else entirely). The Makefile grows a thicket of platform detection conditionals:

```makefile
UNAME := $(shell uname)
ifeq ($(UNAME), Linux)
    CFLAGS += -fPIC
endif
ifeq ($(UNAME), Darwin)
    LDFLAGS += -dynamiclib
endif
```

These conditionals accumulate over years as new platforms are added and old ones are forgotten. By the time a platform is retired, nobody is sure which conditionals are still needed.

**Order-dependent builds.** Because Make's correctness depends on timestamps and explicit dependency declarations, build behavior can depend on the order in which files were created or modified. A clean build works because everything is rebuilt in dependency order. An incremental build may fail because a specific sequence of edits put the build into a state where some targets are considered up to date when they shouldn't be. Diagnosing these failures — "clean build works, incremental build breaks" — is one of the most time-consuming categories of build debugging. The fix is usually to run `make clean` and rebuild from scratch, which defeats the purpose of incremental builds.

---

<a id="ch3-3"></a>
### 3.3 The Build Engineer — a named, essential organizational role

The difficulties described in the previous section were not theoretical. They were experienced by every large software organization that used Make at scale, and the industry's response was predictable: it created a specialized role to manage them.

**Build Engineer** and **Release Engineer** were real job titles, with career ladders, headcount allocations, and at large companies, entire teams. This is not a minor historical footnote. The existence of a whole job category dedicated to keeping the build working is the most concrete evidence available that task-based build systems were failing at their fundamental purpose.

**The role at different scales**

At a small company of 20 to 50 engineers, the build was typically owned by one senior engineer who "also did builds." It was not their entire job, but it consumed a disproportionate fraction of their time. When the build broke — and it broke regularly — everything stopped until they fixed it. Their knowledge of the build infrastructure was deep, undocumented, and irreplaceable. If they left, nobody fully understood the build system for months afterward.

At a mid-size company of 100 to 500 engineers, the role became dedicated headcount — typically one to three people who owned the Makefile infrastructure, the release scripts, the nightly build system, and the packaging and deployment pipelines. They attended architectural planning meetings specifically to flag when proposed changes would increase build complexity. They were the people who understood the platform abstraction layers, the header dependency chains, and which parts of the Makefile were load-bearing versus historical accident.

At large companies — Sun Microsystems, Hewlett-Packard, SGI, IBM, Microsoft, AT&T through the 1980s and 1990s — build engineering was an entire team, sometimes 10 to 20 people, organizationally separate from software development and from QA. They owned dedicated build hardware: rooms full of workstations that existed purely to run builds, because developer machines, with their ad-hoc modifications and locally installed libraries, could not be trusted to produce reproducible results. If you worked at one of these companies and wanted to make a release build, you submitted your code to the build team and they built it on their controlled machines.

**The daily build and the build cop**

One of the core practices that emerged from this era was the **daily build**: every night, a clean build of the entire codebase runs from scratch, and the result is verified. If the build fails, the build is broken, and the failure is announced to the team.

Joel Spolsky wrote about this in the early 2000s as a basic best practice that many software companies still had not implemented. That he felt the need to advocate for it as a best practice — when the alternative was letting the build stay broken for days or weeks — is itself a measure of how dysfunctional the norm was.

Associated with the daily build was the **build cop**: the engineer responsible for keeping the daily build green. In some organizations this was a rotating weekly duty — every engineer took a turn as build cop, which had the beneficial side effect of forcing everyone on the team to understand how the build worked and how it broke. In other organizations it was owned entirely by the build team.

Without these practices, the common pattern was: someone breaks the build, nobody notices immediately, engineers discover the broken build when they try to compile, they work around it with local patches that paper over the underlying problem, the workarounds accumulate, and the actual root cause becomes progressively harder to find.

**Software Configuration Management — the formal discipline**

At larger organizations, the build engineering role sat inside a broader discipline called **Software Configuration Management**, or SCM. SCM owned not just the build but the entire chain from source code to deployed artifact: version control infrastructure, branching and merging strategy, build scripts, release packaging, artifact storage, and environment management.

SCM was recognized as a distinct engineering specialty with its own certifications (CMMi, IEEE standards), its own conferences, and its own career track that did not require being a software developer. A senior SCM engineer might spend their entire career on build and release infrastructure without writing product code.

The tooling this role maintained was elaborate. **IBM Rational ClearCase** was the dominant enterprise version control and build system through the 1990s and into the 2000s. ClearCase's model — dynamic views, derived objects, versioned object bases, configuration specifications — was genuinely powerful. It supported branching strategies, build variant management, and derived object caching in ways that were years ahead of what Git would offer. It also required dedicated administrators the way Oracle required DBAs. Understanding how to configure a ClearCase view, write a config spec, manage a VOB (versioned object base), and debug derived object storage failures was a full-time specialization. Companies hired ClearCase administrators, trained them for months, and depended on them as critical infrastructure. Getting the build configuration right in ClearCase was typically a weeks-long project.

**Perforce** occupied a similar niche at game companies and financial firms, where its performance on large binary files and its strict central server model were valued. The gaming industry in particular — dealing with gigabytes of assets alongside code — needed version control and build engineers who understood the specific constraints of that combination.

**The knowledge transfer problem**

The deepest structural failure of the build-engineer-as-specialist model was the knowledge transfer problem. The build engineer's knowledge was almost entirely undocumented, because it lived in the Makefile itself, and Makefiles at scale are not self-documenting.

Why does this `ifeq` block exist? It works around a compiler bug on HP-UX 10.20 that was never fixed. HP-UX 10.20 has been out of support for fifteen years but nobody has verified that removing the workaround is safe.

Why is this target listed before that one even though the dependency graph doesn't require it? Because in 2003 there was a race condition when they ran in parallel and the fix was to serialize them. The race condition was never diagnosed; the serialization was never reverted.

Why does this library get compiled with `-O1` when everything else uses `-O2`? The original author is gone. There's a comment that says "do not change" with no explanation.

This kind of implicit, unverifiable knowledge accumulated in every large Makefile over years of maintenance. It was carried in the build engineer's head, occasionally in comments, never in documentation. When they left, their replacement spent months in a state of learned helplessness: making changes carefully, watching what broke, and building up a new mental model of the invisible constraints. During this period, builds broke in new ways as the codebase evolved and the inherited workarounds no longer matched the problems they were written to solve.

The build engineer was, in this sense, a single point of failure for the entire engineering organization's ability to ship software. If they were unavailable — sick, on vacation, departed — the organization's release capability was degraded or gone until they returned or a replacement was trained.

**The implicit verdict**

The existence of this role, at this scale, for this long, is the historical verdict on task-based build systems. If a build system requires a specialist to prevent it from breaking, the build system is doing something wrong. The complexity that justified the specialist's existence was not inherent to the problem of building software. It was created by a design that gave the engineer too much power and the system too little.

---

<a id="ch3-4"></a>
### 3.4 Generation 2: artifact-based systems

Artifact-based build systems — Google's Blaze (internal, ~2006), Bazel (open source, 2015), Meta's Buck (2013), Buck2 (2022) — were designed with a different foundational premise: **the build system should know exactly what every build action produces.**

Instead of "run this command," you declare "produce this artifact from these specific inputs":

```python
cc_library(
    name = "utils",
    srcs = ["utils.cc"],
    hdrs = ["utils.h"],
    deps = ["//base:logging"],
)
```

The build system owns the execution. It knows the inputs, it knows the expected output, it runs the action in an isolated sandbox, and it verifies the result. Because it knows the complete input set and the complete output set, it can:

- **Cache safely:** same inputs will always produce the same output, so if it has seen these inputs before it can return the cached result without running the action
- **Parallelize safely:** actions with no shared inputs or outputs cannot interfere with each other, so the system can run them simultaneously without the engineer declaring it
- **Detect missing dependencies:** if a build action tries to read a file that was not declared as an input, the sandbox blocks the access and the build fails — forcing the engineer to declare the dependency explicitly

The core design insight is a transfer of power: the engineer gives up the ability to run arbitrary commands, and in exchange the system takes on responsibility for correctness, caching, and parallelism. The engineer describes *what* to build; the system decides *how*.

This design makes the build engineer's traditional job largely unnecessary. There is no timestamp model to defeat, no platform abstraction layer to hand-maintain, no recursive Make ordering to manage. The system enforces the constraints that the specialist previously enforced through discipline.

---

<a id="ch3-5"></a>
### 3.5 The three properties of a modern build system

Artifact-based systems aim for three properties that task-based systems cannot guarantee:

**Hermeticity.** A build action can only see what it declares. If your `cc_library` rule doesn't list a header as an input, the sandboxed compiler process cannot read it — the filesystem mount is constructed to make it invisible. Undeclared dependencies are not silently tolerated; they are structurally blocked. This is the property that makes the cache correctness argument (Chapter 5) hold.

**Incrementality.** Only rebuild what changed, verified by content hash rather than timestamp. If the content of `utils.cc` is identical to what was compiled last time — same bytes, same compiler flags, same transitive dependencies — the build system returns the cached object file without running the compiler. The change detection is reliable in a way that timestamp comparison is not.

**Reproducibility.** The same inputs always produce the same outputs, on any machine, at any time. A build that passes on one engineer's laptop will produce the same artifacts on any other laptop, on CI, on a colleague's machine. This is a consequence of hermeticity (the build can't see anything it didn't declare) combined with content hashing (the same declared inputs always produce the same cache key).

---

<a id="ch3-6"></a>
### 3.6 Why these properties are hard to achieve simultaneously

Each of these properties is achievable individually. Making all three hold simultaneously for a large, real-world codebase is hard.

Hermeticity requires sandboxing, and sandboxing has overhead. Every action must set up a sandboxed environment — constructing the symlink forest, configuring the filesystem namespace — before executing. For a large build with hundreds of thousands of actions, this overhead accumulates.

Incrementality requires accurate change detection, which requires complete input declarations. Complete input declarations require engineers to list every header, every data file, every tool their action depends on. This is more work than writing a shell command, and getting it wrong degrades incrementality silently: a missing declaration means the cache key is incomplete, which means cache hits may be incorrect.

Reproducibility requires that the build environment contain no hidden variables: no ambient environment variables, no implicitly linked system libraries, no timezone dependencies, no random seeds. Achieving true reproducibility for complex software — especially software that links against system libraries or generates code with timestamps — requires ongoing vigilance and tooling.

The engineering effort to maintain all three properties simultaneously at Google's scale is substantial. It is, however, a one-time engineering cost that is paid by the build system team rather than a recurring tax paid by every engineer every day. That is the core economic argument for artifact-based systems over task-based ones.

---

<a id="ch3-7"></a>
### 3.7 What Bazel changed about the build engineering role

The build engineering role did not disappear when artifact-based systems arrived. It transformed.

Before artifact-based systems, the build engineer's job was primarily reactive: diagnose broken builds, maintain fragile Makefile infrastructure, manage the knowledge that the system could not encode. Their expertise was about working around the system's limitations.

After artifact-based systems, the job became primarily proactive: design the BUILD file architecture, write and maintain Starlark rules, define the toolchain configuration, tune remote cache and RBE performance. Their expertise is about using the system's capabilities effectively.

More importantly, the knowledge is now *in the code* rather than in their head. A new engineer can read the Starlark rules and understand exactly what each build action does and why. The platform abstraction layer is expressed as explicit toolchain rules rather than Makefile conditionals. The ordering constraints are expressed as dependency declarations rather than implicit ordering hacks. When the build engineer leaves, their replacement reads the rules and the BUILD files and understands the system — not in months, but in days.

The single point of failure is gone. What remains is a legitimate systems engineering specialty: the build system as infrastructure, maintained with the same rigor as the production serving infrastructure.

---

<a id="ch5"></a>
## Chapter 5 — Change detection and content hashing

**One-line pitch:** File modification timestamps lie. Content hashes don't. This one switch is what makes reliable incremental builds possible — but the guarantee is only as strong as what you actually hash.

---

<a id="ch5-7"></a>
### 5.7 When is a stale cache hit truly impossible — unpacking the claim

It is common to hear that with a properly designed build system, a stale cache hit is "mathematically impossible." This claim is true, but it has a hidden load-bearing assumption that is worth examining carefully.

**The mathematical argument.** A build action's cache key is the SHA-256 hash of all its declared inputs: source file contents, compiler binary, compiler flags, transitive dependency hashes, relevant environment variables, target platform. If any input changes, its hash changes. If any input hash changes, the composite cache key changes. A different cache key is a cache miss, which triggers a rebuild. For a stale hit to occur — for the system to return a cached result that does not correspond to the current inputs — the new cache key would have to collide with the old one. SHA-256 collision probability is approximately 1 in 2²⁵⁶. For practical purposes, this is zero.

**The hidden assumption.** The argument is valid only if *all inputs are correctly declared*. This word "all" is doing enormous work. If even one input to a build action is not listed in its input declaration — a header it reads, an environment variable it consults, a system library it links against — then the cache key does not capture that input. A change to that undeclared input will not change the cache key. The system will return a cached result that was produced with a different version of the input. This is a stale hit.

The stale hit did not require a SHA-256 collision. It required an incomplete input declaration. The two failure modes are completely different: one is cryptographically infeasible, the other is a mundane engineering mistake.

**The asymmetry.** The hash function is sound. SHA-256 does exactly what it claims to do. The fragile part is the declared input set. If the set is complete, the guarantee holds. If the set is incomplete, the guarantee fails entirely — not probabilistically, but categorically.

**Hermeticity as the structural fix.** The only way to make the "all inputs" assumption hold automatically, rather than relying on engineer discipline, is hermeticity. If a build action runs in a sandbox that physically prevents access to anything not declared as an input, then the set of declared inputs is by construction the set of all inputs — because any undeclared input would have caused the build to fail at access time rather than silently affecting the output.

This is why hermeticity and cache correctness are inseparable. The stale-cache-is-impossible claim is true *given hermeticity*. Without hermeticity, it is an aspiration enforced by discipline.

---

<a id="ch7"></a>
## Chapter 7 — Caching: local, remote, and distributed

**One-line pitch:** The best build is the one you never run — but a wrong cache hit is worse than no cache at all. Correctness and speed are both required, and they have different solutions.

---

<a id="ch7-3"></a>
### 7.3 The two failure modes that look identical from the outside

"It passed locally but failed on CI" is the most common build system complaint in any engineering organization. It looks like a cache problem. It usually isn't.

There are two distinct failure modes that produce this symptom, and they have completely different causes and solutions:

**Failure mode A — stale cache.** The build action's inputs changed, the cache key should have changed, but due to a bug in cache key computation the system returned an old cached result. This is rare in a well-implemented content-addressed system. It requires either a hash collision (negligible probability) or a bug in the key computation logic. When it does happen, the fix is to correct the key computation.

**Failure mode B — environmental leak.** The build action has an undeclared dependency on something in the environment — a system library version, an environment variable, the system compiler rather than the declared compiler, an ambient config file. The build passes locally because the local environment has the right version of the undeclared dependency. CI has a different environment. The build fails on CI not because the cache is wrong but because the build itself was wrong: it was never truly hermetic.

The cache is not the cause of Failure Mode B. The missing input declaration is the cause. In fact, the cache makes the underlying problem *harder to detect*: if the build always ran from scratch, the environmental dependency would surface immediately as an inconsistency across machines. The cache allows the local build to succeed (using the locally-correct undeclared input) and the CI build to fail, making it look like a caching issue when it is a declaration issue.

Distinguishing these two failure modes matters because the fixes are completely different. For Mode A: audit cache key computation. For Mode B: find the undeclared dependency and either declare it explicitly or make the build hermetic so the dependency cannot be accessed implicitly.

---

<a id="ch7-4"></a>
### 7.4 Environmental leaks in depth

Environmental leaks are the most common source of "works on my machine" failures in build systems, and they are more varied than they first appear.

**System library version.** The build action links against `libssl.so` from the system without declaring it as an input. The developer's machine has OpenSSL 3.1. CI has OpenSSL 1.1. The binary links successfully on both but behaves differently at runtime, or fails to link on CI because a symbol changed between versions.

**Ambient PATH.** The build action invokes `python3` by name without declaring which Python binary to use. The developer's machine resolves `python3` to 3.11. CI resolves it to 3.9. The action produces different output depending on which Python runs it.

**Implicit compiler selection.** The build action is declared as a `cc_binary` but the sandbox is not fully hermetic, so the action can see the system compiler at `/usr/bin/gcc`. The developer has GCC 13. CI has GCC 11. Subtle code generation differences change the binary.

**Undeclared header.** The `.cc` file `#include`s a header from a system path that is not declared as a dependency. The header exists on the developer's machine from a previously installed package. It does not exist on a fresh CI worker.

**Ambient configuration files.** A tool reads a configuration file from `~/.config/tool/settings.json` if it exists. The developer's config enables an optimization. CI workers have no such file. The output differs.

In all of these cases, the cache did not cause the problem. The problem was already there — the build was non-reproducible across environments. The cache made it visible by contrasting the locally-cached result with what CI actually builds.

The structural fix is hermeticity: run every build action in a sandbox that has no access to anything not explicitly declared. If the action tries to read `/usr/lib/libssl.so` and that path is not in the declared inputs, the sandbox blocks the read and the build fails immediately — on the developer's machine, visibly, with a clear error. The undeclared dependency is surfaced at the point of creation rather than discovered in CI.

---

<a id="ch7-5"></a>
### 7.5 Remote cache — shared across engineers and CI

The remote cache is where content-addressed build systems become a force multiplier for large teams.

Every engineer's local build populates the remote cache with the artifacts it produces. Every other engineer's build queries the remote cache before running any action. If engineer A compiled `//auth:server` this morning and pushed the result, every other engineer who needs `//auth:server` with the same inputs gets it immediately — without compiling. The action was run once. Everyone benefits.

This works correctly because of content addressing. The cache key for `//auth:server` is derived from the content of its inputs, not from who built it or when. Engineer A's result has the same key as engineer B's result if and only if their inputs were identical. The cache therefore never returns engineer A's result in response to engineer B's query unless the inputs were genuinely identical — in which case the result is, by the guarantee in Section 5.7, correct.

The practical consequence is that a large team populates an increasingly warm cache over the course of a day. Morning engineers build their targets; by afternoon, most incremental builds for any individual engineer are almost entirely cache hits. The team's aggregate build work does not scale linearly with team size — it approaches a ceiling set by the number of distinct input combinations that actually occur. In a stable codebase with active development, that ceiling is far below N×single-engineer-build-time.

This also means the remote cache helps enormously with independent changes (Case 1 from Chapter 2) and cannot help with chain changes (Case 2). When changes touch the same dependency chain, each change produces a new set of input combinations that no other engineer has seen. The cache has no entries for these combinations. Every action on the chain is a miss. The cache is not failing — it is correctly reporting that this integrated state has never been built before.

---

<a id="ch8"></a>
## Chapter 8 — Build systems at team scale

**One-line pitch:** A build system that works perfectly for one engineer can fail catastrophically for 10,000 — not because the build system is wrong, but because the graph shape and coordination model weren't designed for concurrent work.

---

<a id="ch8-2"></a>
### 8.2 The stable interface pattern

The most effective architectural intervention for build performance at team scale is the separation of interface from implementation at every module boundary.

In C/C++ terms: every module has a header file (the interface) and an implementation file (the definition). Everything that uses the module depends on the header. Nothing that uses the module depends on the implementation file. The linker depends on the implementation file; nothing else does.

```
Before:                          After:
                                 
main.c ──────────────┐          main.c ──────────────┐
server.c ─────────── utils.c    server.c ─── utils.h ──── utils.c
auth.c ──────────────┘          auth.c ──────────────┘
```

In the "Before" diagram, `utils.c` is on the critical path of everything. Any change to `utils.c` — even a refactor that doesn't change any externally visible behavior — forces `main.c`, `server.c`, and `auth.c` to recompile. Their object files are inputs that include the hash of `utils.c`'s compilation output, which changed.

In the "After" diagram, `utils.h` is the shared dependency. If the implementation in `utils.c` changes but the interface in `utils.h` does not — the function signatures stay the same, the struct layouts stay the same — then nothing that depends on `utils.h` needs to recompile. The change is invisible to the rest of the graph. Only the linker must re-run, which is typically fast.

This is not a build system feature. It is an architectural decision that the build system can enforce through dependency declarations. The rule is simple: callers depend on interfaces; only the linker depends on implementations. Applied consistently, this rule converts implementation changes — the most common category of change — from graph-wide cache invalidation events into purely local rebuilds.

The same principle applies in other languages: Java interfaces vs. classes, Python module `__init__.py` defining the public API vs. the internal module files, Go package interfaces vs. concrete types. Every language has a mechanism for expressing "this is the contract, this is the implementation" — and exploiting that separation at module boundaries is the most direct way to reduce the blast radius of changes.

---

<a id="ch8-5"></a>
### 8.5 The submit queue — serializing merges to eliminate integration chaos

A submit queue is the build-system-level answer to the following problem: engineer A's change passes CI, engineer B's change passes CI, but when both land on main simultaneously, the combination breaks the build.

This failure mode is structurally guaranteed in any system where CI validates changes against a base commit rather than against the current HEAD. Two changes that are each individually correct can be mutually incompatible — they modify the same function in different ways, they both add a field to the same struct, they both rename a variable. Neither CI run detects this because neither CI run sees the other change.

The submit queue fixes this by enforcing a total order on merges: before your change can land, it must be validated against current HEAD (which includes all changes that landed before yours in the queue). Each validation is incremental from the previous known-good state.

```
Queue position 1: validate A on top of HEAD → pass → merge A
Queue position 2: validate B on top of HEAD+A → pass → merge B
Queue position 3: validate C on top of HEAD+A+B → fail → C is rejected
```

Engineer C's change is rejected not because it was wrong in isolation but because it conflicts with the combination of A and B. The submit queue found this before it reached main.

The cost is throughput: each change must wait for the previous validation to complete before its own validation can start against the correct base. With a 10-minute CI time and 200 engineers all trying to merge simultaneously, the last engineer in the queue waits theoretically 2,000 minutes. In practice, submit queues address this through batching (validating multiple compatible changes together), speculative execution (beginning validation before the upstream change fully completes), and prioritization (urgent fixes jump the queue). But the fundamental tension between correctness (total ordering) and throughput (parallelism) cannot be eliminated — it can only be managed.

Google's TAP (Test Automation Platform) is the most sophisticated implementation of this pattern. TAP batches changes, predicts which changes can be safely parallelized, runs speculative builds, and manages a queue of thousands of changes per day across hundreds of millions of lines of code. Its throughput is only possible because the underlying build system (Blaze/Bazel) makes each incremental validation fast — and even so, the queue is frequently the bottleneck in getting changes to production.

---

<a id="ch9"></a>
## Chapter 9 — Version control philosophy and its build consequences

**One-line pitch:** Your version control model is not independent of your build system. The choice between monorepo and polyrepo, between trunk-based development and branch-centric workflows, between source dependencies and artifact dependencies — each of these choices profoundly shapes what your build system can do and how it must be designed. Getting this wrong creates years of pain that no amount of tooling can fix.

---

<a id="ch9-1"></a>
### 9.1 The foundational question: where does your dependency live?

Every software organization must answer one foundational question about how its components depend on each other: **is the dependency expressed as source code or as a versioned artifact?**

If it is source code, then when you need library X, you check it out from the repository alongside your own code and build it. The version you use is whatever is currently on the shared trunk. Changing X and changing all its callers can happen in the same commit.

If it is a versioned artifact, then library X is published as version 2.4.1 to an artifact registry, and your project declares a dependency on `X >= 2.4.0`. You download the binary artifact at build time. Changing X requires publishing a new version; callers upgrade on their own schedule.

These are not just technical choices. They encode a philosophy about organizational structure, the distribution of coordination costs, and the expected rate of change of shared interfaces. Everything downstream — which build system makes sense, which CI model works, how breaking changes propagate — follows from this choice.

---

<a id="ch9-2"></a>
### 9.2 Google's model: Piper, G4, and the monorepo

Google's internal version control system is **Piper**, a distributed filesystem designed specifically for a codebase that cannot fit in any single machine's memory. Engineers access it through a client called **G4** — the fourth iteration of Google's internal version control tooling.

The defining characteristic of Piper is that it hosts the entire Google codebase — Search, Ads, Maps, YouTube, Android, Cloud, and every internal tool and library — in a single repository. One repo. One trunk. No forks. Engineers across the entire company work in the same codebase and see each other's changes (subject to access controls) as they are committed.

The scale is genuinely unprecedented: billions of files, decades of history, tens of thousands of engineers making changes concurrently every day. No off-the-shelf version control system can handle this. Git, as we will examine in Section 9.7, requires every clone to contain the full history — at Google's scale, a full clone would take weeks and consume terabytes of storage. Piper uses a virtual file system model: you see the entire tree when you browse it, but files are only fetched to your local machine when you actually open them. The G4 client is purpose-built to make this access pattern feel natural.

---

<a id="ch9-3"></a>
### 9.3 What the monorepo makes possible

**Atomic cross-project changes.** The most important consequence of a monorepo is that you can change a library and all its callers in the same commit. The codebase never enters a state where a new library version exists but its callers have not been updated. There is no migration period, no backward compatibility requirement for the migration, no need to maintain two versions simultaneously while the upgrade propagates. You update everything at once or you don't update at all.

This sounds simple. Its implications are enormous. At Google, it means that refactoring a widely-used API — renaming a function, changing a parameter type, splitting a module — can be done in a single coherent operation across the entire codebase. Google has built internal tooling, notably the large-scale change (LSC) infrastructure and a tool called Rosie, specifically to automate these global refactors. Rosie takes a refactoring operation, applies it across the entire codebase, splits the result into manageable chunks, and creates code review requests for each chunk. A change to an API that touches 10,000 call sites can be reviewed, tested, and landed in days. In a polyrepo model, the same change requires coordinating with dozens of teams over months.

**"Live at HEAD" as an organizational principle.** In a monorepo, every team always depends on the current version of every library. There is no version pinning for internal dependencies. If a library changes in a way that breaks a downstream team, the library's team is responsible for updating the downstream callers before the change lands — or the change does not land. The burden of upgrades is placed on the upstream team (who understands the change) rather than the downstream teams (who may not). This inverts the traditional upgrade model, where downstream teams accumulate technical debt by staying on old library versions.

**The build system is the dependency manager.** There is no Maven, no npm, no pip for internal dependencies. Bazel resolves all internal dependencies from source. You declare `deps = ["//libraries/auth:client"]` and Bazel finds the source, determines what needs to be compiled, and builds exactly what is needed. Because everything is source, there are no version numbers, no lockfiles, no dependency resolution algorithms to debug. The dependency graph is just the build graph.

**Visibility into the entire dependency graph.** Because the entire codebase is in one repository and one build system, tooling can analyze the entire global dependency graph. Google knows, for any library, every file in the codebase that depends on it — transitively. This information powers automated impact analysis ("if I change this, what breaks?"), automated large-scale refactoring, build performance optimization, and code health metrics. None of this is possible when the codebase is distributed across thousands of independent repositories.

---

<a id="ch9-4"></a>
### 9.4 What the monorepo requires

The monorepo's advantages come with requirements that are easy to underestimate.

**Custom version control infrastructure.** As noted, Piper is not Git. It is a purpose-built distributed filesystem with a custom client, custom access controls, and custom change management tooling. Building and maintaining this infrastructure is a significant engineering investment. For organizations that cannot make this investment, approximations are possible — Microsoft built VFS for Git (now called Scalar) to handle the Windows codebase, and tools like `git sparse-checkout` can reduce the cost of working in a large Git monorepo — but these are approximations with real limitations.

**A build system designed for the model.** The monorepo model only works smoothly with a build system that understands source-level dependencies and can build exactly the subset of the codebase that any given change touches. Bazel was designed for this model. Running Maven or Gradle in a monorepo creates a system where the build tool tries to manage version numbers for things that have no versions, and produces confusion and fragility. The build system and the VCS model must be matched.

**Cultural and organizational alignment.** "Live at HEAD" means upstream teams are responsible for updating downstream callers when they change interfaces. This requires a culture of ownership and responsibility that not every organization has. In a polyrepo model, you can publish a breaking change, increment the major version, and move on — it is the downstream teams' problem to upgrade when they choose. In a monorepo, breaking changes require immediate, global resolution. Teams must be willing to do this work, and the tools (automated refactoring, large-scale change infrastructure) must exist to make it feasible.

---

<a id="ch9-5"></a>
### 9.5 Amazon's model: polyrepo, service ownership, artifact versioning

Amazon's engineering culture is built around a principle associated with Jeff Bezos: the **two-pizza team**. Teams should be small enough to be fed by two pizzas. Each team owns a service end-to-end — from design through implementation to production operations. Each service has its own repository, its own build system, its own deployment pipeline, and its own on-call rotation.

This is a deeply different organizational philosophy from Google's monorepo model, and it produces deeply different build infrastructure.

In the Amazon model, services communicate through APIs, typically HTTP or gRPC. The API is the interface. The implementation behind the API is the team's private concern. When a team changes their implementation, they deploy a new version of their service. Callers do not recompile or rebuild — they call the same endpoint and the behavior changes behind it.

For libraries — code shared across services rather than deployed as separate services — the dependency model is versioned artifacts. Team A publishes `auth-client-java 2.4.1` to an internal Maven repository. Every service that needs authentication functionality declares a dependency on some version of `auth-client-java`. Each service decides independently when to upgrade.

---

<a id="ch9-6"></a>
### 9.6 What polyrepo makes possible — and what it breaks

**Team autonomy.** A team can deploy their service any time they choose, without coordinating with any other team. Their repository is their domain. They set their own code review standards, their own testing requirements, their own deployment cadence. There is no shared trunk to worry about, no risk that another team's commit breaks their build, no coordination required to ship.

**Independent technology choices.** Each service can choose its own language, framework, and build toolchain. A Java service uses Gradle. A Python service uses pip and pytest. A newer Rust service uses Cargo. There is no central mandate for consistency. Teams adopt new technologies when they choose rather than when the organization agrees to migrate.

**Clear blast radius.** A bug in one service affects that service. It does not affect other services until they upgrade to a version that contains the bug. Services can roll back independently. The isolation is real.

**The diamond dependency problem.** Service A depends on `library-X 1.2` and `library-Y 2.0`. Library-Y 2.0 depends on `library-X 1.3`. Service A now has two versions of library-X in its dependency graph. In Java (Maven/Gradle), the build system picks one version and silently ignores the other — which version it picks depends on which path in the dependency graph reaches it first. This silent resolution may cause runtime failures that are extremely hard to diagnose.

In a monorepo, this problem does not exist. There is exactly one version of library-X in the codebase. If library-Y needs a different version, the version must be updated for everyone simultaneously.

**Cross-service atomic changes are structurally impossible.** If a library API changes in a breaking way, the migration in a polyrepo model looks like: publish `library 3.0` with the new API, maintain `library 2.x` with the old API, coordinate with each downstream team to upgrade (weeks to months), deprecate 2.x once all dependents have migrated, eventually stop maintaining it. During the entire migration period, both the old and new API exist, the library team supports both, and the codebase is in an inconsistent state.

In Google's monorepo, the same migration is: change the library, update all callers in the same commit using automated tooling, land everything at once. Total time: days.

---

<a id="ch9-7"></a>
### 9.7 Git's philosophy and its build consequences

Git was designed by Linus Torvalds in 2005 for Linux kernel development. The Linux kernel development model is: distributed, with no single authoritative central repository; branch-centric, with contributors maintaining separate branches and submitting patches for inclusion; meritocratic, with no formal organizational hierarchy. Git's design reflects this model deeply.

Every Git clone is a complete copy of the entire repository history. Branching is cheap — a branch is just a pointer to a commit. Merging is the primary coordination mechanism. Working offline is fully supported.

These design choices have profound consequences for build systems:

**Git encourages the polyrepo model.** Because each clone is independent and self-contained, the natural unit of Git usage is "one repository per project." Monorepos are possible in Git, and many companies run them successfully at moderate scale. But Git was not designed for them, and this shows. Cloning a Git repository with millions of files and decades of history takes a very long time and consumes enormous storage. Operations that are fast on a small repo — `git status`, `git grep`, `git log` — become slow as the repo grows.

Microsoft's experience is instructive. When they moved the Windows codebase to Git, they created what was at the time the largest Git repository in the world. The standard Git client was completely unusable — basic operations took minutes. They had to build an entirely new virtual filesystem layer, VFS for Git (now called Scalar), that lazily fetches files on demand rather than cloning the full history. This is the same fundamental problem Google solved with Piper, arrived at independently by a different large organization.

**Branch-based workflows create integration latency.** The standard GitHub workflow — create a feature branch, develop on it, open a pull request, merge — means code lives on a branch, diverging from main, for days or weeks. The longer a branch lives, the larger the merge. Large merges are hard to review, easy to break, and expensive to rebuild. The submit queue model (Chapter 8) is the answer to this, but it is architecturally at odds with long-lived feature branches.

**CI is designed around branches, not the integrated result.** Every major CI system built for Git — GitHub Actions, GitLab CI, CircleCI, Jenkins — validates code on its current branch, not on the result of merging that branch with all other concurrent branches. This means two PRs can each pass CI individually while their combination breaks the build. This is the fundamental integration problem the submit queue addresses; it is structural in Git-based CI, not an implementation oversight.

**Git's content addressing is for history, not for builds.** Git uses SHA-1 (transitioning to SHA-256) to address objects in its object store — commits, trees, blobs. This is a content-addressed system, but it is content addressing in the service of version history: the same file contents always produce the same blob hash, making deduplication and integrity verification efficient. It is not content addressing in the service of build caching: Bazel cannot use Git's object store to determine whether a build action needs to re-run. Bazel computes its own content hashes of build inputs, independently of Git. The two systems operate on the same files but address different questions.

---

<a id="ch9-8"></a>
### 9.8 The mismatch problem

The most common painful situation in build system design arises when a company's version control model and build system model are mismatched. This happens in two directions.

**Adopting Bazel in a polyrepo environment.** Bazel is designed for source-level dependencies in a monorepo. Its external dependency management — `WORKSPACE` files, `http_archive` rules, `bazel_dep` in MODULE.bazel — is functional but fights the grain of the design. Teams that adopt Bazel for its build correctness properties while retaining a polyrepo structure with versioned artifact dependencies spend significant engineering effort on configuration that provides none of Bazel's core benefits. They get the syntax and tooling overhead of Bazel without the architectural properties (live at HEAD, atomic cross-project changes, global dependency visibility) that make the overhead worthwhile.

**Adopting a monorepo structure without the build infrastructure.** Some organizations consolidate all their code into a single repository — often motivated by discoverability or developer experience — without investing in a build system designed for monorepos. The result is a repository where every change triggers a full rebuild of everything, because the build system (Maven, Gradle, npm scripts) has no ability to determine what actually changed and what needs rebuilding. The monorepo has all the organizational complexity of a shared codebase with none of the build performance benefits.

In both cases, the solution is the same: align the version control model and the build system model. Either commit to the monorepo model with appropriate infrastructure (Bazel or Buck2, trunk-based development, large-scale change tooling), or commit to the polyrepo model with appropriate infrastructure (versioned artifacts, service-level isolation, independent CI). Hybrid approaches can be made to work, but they require constant ongoing effort to maintain coherence between two systems with different design assumptions.

---

<a id="ch9-9"></a>
### 9.9 The philosophical divide: a comparison

|  | Monorepo (Google model) | Polyrepo (Amazon model) |
|---|---|---|
| Dependency unit | Source at HEAD | Versioned artifact |
| Breaking change | Atomic — fix all callers now | Staged — migrate over time |
| Team autonomy | Lower — shared trunk, shared standards | Higher — independent repos and cadences |
| Large-scale refactoring | Easy — one operation, one commit | Hard — months of cross-team coordination |
| Diamond dependency | Impossible — one version of everything | Structural risk |
| Build system role | Dependency manager and builder | Builder only |
| VCS requirements | Custom or VFS layer | Standard Git |
| Best fit | Tightly coupled systems, shared platforms | Loosely coupled services, high team autonomy |

Neither model is universally correct. Organizations with tightly coupled systems — where one team's change frequently requires updates in other teams' code — benefit enormously from the monorepo's atomic change capability. Organizations with loosely coupled services — where teams deploy independently and interfaces are stable — pay the coordination cost of the monorepo without receiving proportionate benefit.

The critical insight is that this is a choice that must be made deliberately, with full awareness of its build system implications. Companies that fall into a model by accident — ending up with a monorepo because it was convenient, or a polyrepo because nobody thought about it — often spend years dealing with infrastructure mismatches that could have been avoided.

---

<a id="ch10"></a>
## Chapter 10 — Theory meets reality: why even Google has slow builds

**One-line pitch:** All the theory eliminates waste. It does not eliminate work. At Google's scale, even a perfectly waste-free build takes a long time for large changes — and understanding why is as important as understanding the theory.

---

<a id="ch10-1"></a>
### 10.1 What the theory actually promises

The theory of artifact-based build systems, remote caching, content addressing, and hermetic execution promises one precise thing: **build time proportional to the scope of the change, not to the size of the codebase.**

A change to one `.cc` file that does not affect any interface: the build system recompiles that one file, relinks the binaries that depend on it, and runs only the tests whose inputs changed. For a large codebase, this might mean recompiling one file and running a handful of tests. Seconds.

A change to a widely-used interface: the build system must recompile every file that transitively depends on that interface, relink every binary that includes any of those files, and run every test that might be affected. For a core library used throughout a large codebase, this might mean recompiling millions of files and running hundreds of thousands of tests. Hours.

Both of these outcomes are correct. The build system is working exactly as designed. The slow build is slow because the change genuinely affected a large portion of the codebase, and verifying correctness of a large portion of the codebase takes time. The infrastructure makes that time as small as physically possible — but it cannot make it zero, because the work is real.

What the theory eliminates is *unnecessary* work: recompiling files that didn't need to change, running tests whose inputs are identical to a previous run, serializing work that could run in parallel. What it cannot eliminate is *necessary* work, and at Google's scale, even the necessary work is enormous.

---

<a id="ch10-4"></a>
### 10.4 Tests are the real bottleneck, not compilation

TAP — Google's Test Automation Platform — does not just build code. It runs tests. And test execution time has a fundamentally different character from compilation time.

A compiled artifact can be cached indefinitely if its inputs have not changed. If `//auth:server` compiled to the same binary last week as it does today, the cached binary is valid and no compilation is needed. The artifact does not decay.

A test result can only be cached if two conditions hold: the test is deterministic (same inputs always produce the same result), and the test's inputs have not changed since the cached result was produced. For many categories of tests, one or both conditions fail.

**Integration tests** that call real services, read from real databases, or depend on network behavior are not deterministic in the required sense — they may succeed today and fail tomorrow due to infrastructure changes with no relationship to the code being tested. Their results cannot be cached.

**Tests with broad input sets** — tests that depend on the behavior of a complex system — may have input sets that include almost everything in the codebase. A change to any component of the system invalidates the cached test result, forcing a re-run.

**Flaky tests** fail non-deterministically: they pass sometimes and fail sometimes with identical inputs, due to timing races, network variability, or other sources of non-determinism. TAP detects flakiness and retries suspected flaky failures. At Google's scale — millions of test runs per day — even a 0.1% flakiness rate produces thousands of retries daily, each consuming worker time and adding to queue depth.

The practical consequence is that for any change to a widely-used library, TAP must execute hundreds of thousands or millions of tests. Even with perfect parallelism across thousands of workers, the wall-clock time is bounded by the slowest test in the critical path of the test dependency graph, and by the total volume of test work divided by the worker pool size. This is frequently measured in hours for large interface changes, and this is expected, correct behavior — not a failure of the build system.

---

<a id="ch10-8"></a>
### 10.8 Clean builds: when all the theory hits zero

All of the cache efficiency described in the preceding sections assumes a warm cache — a state where at least some of the build's required artifacts have been computed before and their results are available. A clean build destroys this assumption. Cache hits drop to zero. Every advantage the theory provides vanishes.

Clean builds happen more often than engineers expect:

**New engineer onboarding.** A new engineer's first build is necessarily a clean build. Their local cache is empty. The remote cache may have entries for the targets they need, but only if those exact targets were built with the exact same configuration by someone else recently enough that the entries haven't been evicted.

**Toolchain upgrades.** Changing the compiler version, upgrading a core build dependency, or modifying compilation flags changes the cache key for every single target in the codebase simultaneously. One upgrade renders the entire remote cache useless in a single operation. The next build by every engineer who takes the upgrade is a clean build.

**Old branch bases.** Remote caches have finite storage and use LRU eviction. If an engineer is working on a long-lived branch based on a commit from two months ago, the remote cache may have evicted the entries for that commit's build graph. Rebasing onto HEAD restores access to a warm cache; staying on the old base means building cold.

**Correctness verification.** Some build and security workflows intentionally run clean builds to verify that the build is fully reproducible and that no cached result is masking a hermeticity violation. These are correct and important to run; they are also slow.

On a clean build, the effective critical path is the full critical path — the longest sequential dependency chain through the entire build graph. There is nothing cached to skip. And the volume of work in a large codebase, even parallelized across thousands of RBE workers, is bounded by the critical path length and by how many actions fit in each parallel wave.

This is not a build system failure. It is an honest accounting of how much work exists. The build system makes the clean build as fast as physically possible — but "as fast as physically possible" for a billion-line codebase with a 200-step critical path is not fast in absolute terms.

---

<a id="ch10-9"></a>
### 10.9 Mitigations for clean build latency

**Cache priming.** Before rolling out a toolchain upgrade, a CI job builds the entire codebase with the new configuration and populates the remote cache. Engineers who upgrade find a pre-warmed cache waiting for them and get incremental build behavior immediately rather than triggering their own clean builds. This converts a fleet-wide clean build event into a single planned one.

**Pre-built artifact distribution.** For stable base layers — third-party libraries, core platform components — pre-built binary artifacts can be distributed to engineers rather than built from source. Engineers download the pre-built artifact at workspace setup time. This sidesteps the build entirely for those nodes, converting their build cost to a download cost.

**Stratified caching.** Maintain a "blessed" remote cache snapshot corresponding to each release branch. New engineers seed their local cache from the snapshot before their first build. They then only build the delta between the snapshot and their current workspace — which, for a new engineer who has not yet made any changes, may be very small.

**Language choice.** C++ clean build times are dominated by header parsing and template instantiation. A moderately complex C++ translation unit may `#include` hundreds of headers, each of which must be parsed and type-checked, generating enormous amounts of work per file. Go and Rust were explicitly designed to avoid this: Go has fast compilation as a first-order design goal, with separate compilation, a simple syntax, and no header files; Rust compiles at the crate level with well-defined interfaces. Companies building new services in Go or Rust report clean build times that are dramatically faster for equivalent functional complexity. This is not solely a build system question — it is partly a language design question with direct build consequences.

---

<a id="ch11"></a>
## Chapter 11 — Optimizing the build: keeping the graph healthy over time

**One-line pitch:** Designing a good build graph is not enough. The graph degrades continuously through individually rational decisions, and maintaining build health requires ongoing intervention that scales from manual restructuring to automated detection to AI-assisted prescription.

---

<a id="ch11-1"></a>
### 11.1 The build graph degrades — and it's nobody's fault

The naive assumption is that you design your module boundaries once, write your BUILD files correctly, and you're done. This is wrong in practice. The build graph is a living structure, and it degrades over time through a process that is almost thermodynamic.

Every engineer adding code makes locally rational decisions. They add a dependency because they need a function. They put a new file in an existing module because it's convenient and the module is already there. They add a test that imports broadly because writing a narrowly-scoped test takes more thought. They add a utility function to the base library because that's where utility functions go.

No single change is unreasonable. Each one passes code review. Each one is justified by immediate need. But the cumulative effect, over months and years, is predictable and damaging:

- Modules grow. What started as a focused 500-line library with a clear interface accumulates helpers, exceptions, edge cases, and convenience wrappers until it is a 5,000-line package that does too many things.
- Dependency chains deepen. Each new dependency between modules adds an edge to the graph. Enough edges, and the critical path lengthens — not dramatically on any given day, but steadily, quarter over quarter.
- Interfaces widen. A module's public interface starts narrow (three exported functions) and gradually expands as callers need "just one more thing." A wide interface means any implementation change is more likely to touch something a caller depends on, increasing the blast radius.
- The base library trap forms. One widely-used utility module — string helpers, logging, error handling — accumulates imports from everywhere. It becomes a hub node in the dependency graph. Any change to it invalidates a large fraction of the cache. The deeper it sits in the graph, the worse the cascade.

This degradation is invisible until it crosses a pain threshold. Nobody notices the build getting 2 seconds slower per quarter. They notice when it crosses 5 minutes, or 10 minutes — the point where it changes engineer behavior. People start context-switching during builds. They batch changes to avoid rebuilding. CI queues back up. The submit queue becomes the throughput bottleneck. "The build is slow" becomes a standing complaint in team retros.

By then, the graph has been degrading for years, and the fix is not a configuration change — it is a structural reorganization of the codebase.

---

<a id="ch11-2"></a>
### 11.2 The three stages of growth: when different interventions matter

The interventions that improve build performance are different at each stage of codebase growth. Applying the wrong intervention at the wrong stage wastes effort or actively makes things worse.

**Stage 1: Small codebase, small team (1–20 engineers, <100K lines).**

The build graph is naturally flat and fast because there isn't enough code to create deep chains. Clean builds take seconds to low minutes. Incremental builds are near-instant. The remote cache, if it exists, has high hit rates because few people are making concurrent changes to the same areas.

The correct intervention at this stage is: do nothing special. Fine-grained modularization would be pure overhead — the management cost of many small BUILD targets exceeds any caching benefit. Keep modules at the natural cognitive boundary (one module per logical component, as §1.6 describes). Focus on writing correct dependency declarations and avoid the habit of listing overly broad dependencies ("just depend on the whole utils library"), because those habits will be painful to unwind later.

The most valuable discipline at this stage is not performance optimization but hygiene: declare dependencies accurately, keep interfaces narrow, avoid cycles. These are cheap now and expensive to fix later.

**Stage 2: Medium codebase, growing team (20–200 engineers, 100K–5M lines).**

This is where degradation becomes visible. The base library trap is forming. Some incremental builds are noticeably slower than they used to be — not because any single change caused it, but because the graph has accumulated enough edges that changes to central modules cascade further. "My build used to take 30 seconds, now it takes 3 minutes" is the characteristic complaint.

The correct interventions at this stage are targeted and architectural:

- **Split modules at the points of highest change frequency and widest dependency fanout.** The module that every team imports and that changes weekly is the highest-value split target. Separating its interface from its implementation (the stable interface pattern from §8.2) so that implementation changes don't cascade is the single most impactful optimization.
- **Prune unnecessary dependencies.** Over time, modules accumulate dependencies that were once needed but no longer are — a function was called during development, the call was removed, but the `deps` entry stayed. Each unnecessary dependency widens the cache key and may cause spurious rebuilds. Tools like Bazel's `unused_deps` and `buildozer` automate detection and removal.
- **Break deep chains.** If the critical path runs through a sequence of modules A → B → C → D → E, look for opportunities to decouple. Can B's dependency on A be narrowed to an interface? Can C be reorganized so it doesn't need B's full output? Each link removed from the critical path reduces the sequential floor.

At this stage, a single engineer who understands the dependency graph can often identify the top three or four structural problems by inspection and propose targeted fixes. The fixes are disruptive — splitting a module means updating every importer — but they are feasible because the codebase is still small enough for one person to understand the shape.

**Stage 3: Large codebase, many teams (200+ engineers, 5M+ lines).**

At this scale, no single person can hold the full dependency graph in their head. The graph has thousands of nodes and tens of thousands of edges. The degradation is continuous and distributed — happening in ten different parts of the graph simultaneously, driven by ten different teams, none of whom see the global picture.

Manual intervention is no longer sufficient. The correct interventions at this stage are systematic and tool-driven:

- **Automated graph health monitoring.** Track critical path length, maximum dependency fanout, module size distribution, and cache hit rates as metrics over time. Set alerts when they cross thresholds. Google tracks critical path length as a build health metric (§2.6) precisely because it is the canary for graph degradation.
- **Dependency budget enforcement.** Set a maximum number of direct dependencies per target. When a module exceeds the budget, the build system or linter flags it for review. This prevents the gradual widening that causes hub nodes.
- **Automated unused dependency removal.** Run `unused_deps` analysis as part of CI. Dependencies that are declared but never used in the source are removed automatically or flagged for removal.
- **Periodic graph audits.** Quarterly or per-release analysis of the full dependency graph to identify emerging hub nodes, deepening chains, and modules that have grown beyond their original scope. This is the large-scale equivalent of the Makefile audits from §3.3, but tooling-assisted rather than manual.

The transition from Stage 2 to Stage 3 is where many organizations stumble. They had one senior engineer who understood the graph and kept it healthy through targeted intervention. That engineer's intuitive understanding doesn't scale to a graph too large to visualize. The organization needs to shift from person-driven optimization to process-driven optimization, and that shift requires investment in tooling and metrics that competes for resources with feature work.

---

<a id="ch11-3"></a>
### 11.3 Modularization: the primary technique and its real costs

Modularization — splitting coarse-grained build targets into smaller, well-interfaced modules — is the most frequently recommended technique for improving build performance. It works through two mechanisms:

**Reducing blast radius for incremental builds.** A single large module that contains authentication logic, session management, and token storage means any change to any of the three invalidates the entire module's cached output. Splitting into three modules — `//auth:authn`, `//auth:sessions`, `//auth:tokens` — means a change to token storage only invalidates `//auth:tokens`. Callers of `//auth:authn` and `//auth:sessions` are unaffected. The cache key for those targets hasn't changed, so they don't rebuild.

**Exposing parallelism for clean builds.** Three independent modules can be compiled simultaneously on three workers. One monolithic module must be compiled sequentially (or at least, the compiler handles its own internal parallelism, which may be less efficient than the build system's scheduler). The flatter the graph, the more parallel work the scheduler can dispatch.

These are real benefits. But modularization has costs that are rarely discussed honestly.

**BUILD file complexity increases.** Every new target is a new BUILD file entry (or a new entry in an existing BUILD file) with its own `srcs`, `hdrs`, `deps`, and `visibility` declarations. For N modules, there are O(N²) potential dependency edges to manage. At Google's scale, BUILD files are themselves a significant maintenance burden — they are the most frequently edited non-source files in the repository.

**Over-granularization has diminishing and eventually negative returns.** There is a sweet spot for module size. Too coarse, and you lose caching precision and parallelism. Too fine, and the overhead of managing thousands of tiny targets — graph analysis time, cache lookup overhead, BUILD file maintenance, sandboxing setup per action — exceeds the benefit. A module that wraps a single 20-line function is almost certainly too small; the build system's overhead for setting up, caching, and tracking that target is disproportionate to the compilation work it represents.

The sweet spot depends on your language's compilation model. C++ has high per-file compilation cost (header parsing, template instantiation), so finer granularity pays off — each cache hit saves significant work. Go has low per-file compilation cost (fast compiler, no headers), so very fine granularity yields less benefit per split. The optimal module size is not a universal constant; it is a function of your language, your compiler, and your change patterns.

**Module boundaries create coupling surfaces.** When you split a module, you must define which symbols are public (the interface) and which are private (the implementation). Any symbol you export becomes a dependency that callers can rely on. Once callers depend on it, changing it requires updating all callers. A module that was easy to refactor internally, when it was one unit, becomes harder to refactor once its internals have been exported as interfaces for the new sub-modules.

The risk is that splitting a module exposes implementation details that were previously hidden. If `//auth:tokens` was an internal implementation detail of `//auth`, nobody outside the auth team depended on it. Once it's a separate target with `visibility = ["//auth/..."]`, it's still internal — but the pressure to widen that visibility ("we just need one function from tokens, can you make it visible to us?") is real, and each widening creates new coupling.

**The four forces from §1.6 do not always agree on where to split.** The cognitive boundary (what fits in one mental scope?) may suggest a larger module than the build performance boundary (what is the optimal caching unit?). The organizational boundary (who owns this code?) may not align with the compilation boundary (what is an efficient translation unit?). A split that optimizes for build speed may produce modules that are confusing to navigate. A split that optimizes for team ownership may produce modules that are too coarse for efficient caching.

The practical heuristic: split where the change frequency differs. If one part of a module changes weekly and another part changes monthly, they should be separate targets — not because of any abstract principle, but because the weekly-changing part will constantly invalidate the monthly-changing part's cache if they share a target. Change frequency is the most reliable signal for where module boundaries should be.

---

<a id="ch11-4"></a>
### 11.4 Beyond modularization: the full optimization toolkit

Modularization is the most impactful intervention, but it is not the only one. A complete build optimization strategy includes several additional techniques.

**Interface stabilization.** The stable interface pattern (§8.2) is the single most effective technique for reducing incremental build cascades. By separating interface from implementation at every module boundary, implementation changes — the most common category of change — become invisible to callers. The interface changes rarely; the implementation changes frequently; callers depend only on the interface. Applied systematically, this converts most changes from "rebuild everything downstream" to "rebuild this one module and relink."

**Dependency pruning.** Remove dependencies that are declared but not actually used. These are surprisingly common — they accumulate as code is refactored and call sites are removed without updating the BUILD file. Each unnecessary dependency widens the cache key, potentially causing rebuilds when the unused dependency changes. Automated tools (`unused_deps` in Bazel, `gazelle` for Go) can detect and remove these systematically.

**Breaking the base library trap.** If one module — typically a utilities library — is depended on by a large fraction of the codebase, any change to it invalidates a large fraction of the cache. The fix is to split the base library into independent sub-modules: `//base:strings`, `//base:logging`, `//base:errors`, `//base:time`. Each caller then depends only on the specific sub-module it uses. A change to string utilities no longer invalidates the cached output of modules that only use logging. This is modularization applied to the single highest-leverage target.

**Build configuration hygiene.** Variant explosion — building the same code with many different combinations of flags, platforms, and build modes — multiplies the effective size of the build graph. Each variant is a distinct set of cache keys. If you build `debug` and `release` and `asan` and `tsan` for each of `linux-x86_64` and `linux-arm64` and `macos-arm64`, you have 12 variants, each with its own full build graph. Reducing the number of variants that engineers routinely build — for example, by only running sanitizer builds in CI rather than locally — reduces the effective build surface.

**Compiler and language choice.** As discussed in §10.9, the language's compilation model directly affects build time. C++ clean builds are slow because of header parsing and template instantiation. Go builds are fast because Go was designed for fast compilation. Rust builds are slower than Go but faster than C++ for equivalent complexity, and Rust's crate-level compilation produces well-defined cache boundaries. For new code in a polyglot codebase, choosing a language with faster compilation has direct and permanent build performance benefits that compound over the life of the codebase.

**Graph-aware refactoring.** When refactoring code, consider the dependency graph impact explicitly. Moving a widely-used function from a large module to a small, focused module can dramatically reduce build cascade for changes to the large module. Inlining a small module into its sole caller can reduce graph complexity without affecting build time. These are refactoring decisions where the build graph is a first-class consideration alongside code readability and maintainability.

---

<a id="ch11-5"></a>
### 11.5 The restructuring problem: why optimization doesn't happen

The techniques described above are well understood. Any experienced build engineer can identify the top problems in a dependency graph and propose fixes. The problem is rarely diagnosis. The problem is execution.

**Restructuring competes with feature work.** Splitting a module, pruning unused dependencies, breaking a base library into sub-modules — these all require engineering time. The benefit is diffuse: everyone's build gets slightly faster. The cost is concentrated: one engineer (or one team) must do the migration work. In any prioritization framework that compares "ship feature X this quarter" against "make everyone's build 15% faster," the feature usually wins. The build optimization is deferred. It is deferred again next quarter. The degradation continues.

**The coordination cost is high.** Splitting a module in a monorepo means updating every file that imports from it. In a large codebase, a widely-used module may have hundreds or thousands of importers. Each importer must be updated in a way that is compatible with the new module structure. In a monorepo with a submit queue, this is an atomic change — it must all land together, or the build breaks. Coordinating a change that touches hundreds of files across dozens of teams requires tooling, communication, and organizational authority that most teams don't have.

This is precisely why Google built Rosie and the large-scale change (LSC) infrastructure. A proposed refactoring is expressed as an automated transformation. Rosie applies it across the entire codebase, splits the result into reviewable chunks, and creates code review requests for each affected team. Without this infrastructure, the coordination cost of large restructurings is so high that they simply don't happen — and the graph continues to degrade.

**The benefit is hard to measure.** "The build is 15% faster" is hard to translate into business impact. It saves each engineer a few minutes per day. Aggregated across the organization, this is significant — 200 engineers saving 5 minutes per day is 16 engineering-hours daily, or roughly two full-time engineers' worth of productive time recovered. But this aggregate benefit is invisible in any individual team's metrics. It appears nowhere in a product roadmap. It has no champion.

The organizations that maintain healthy build graphs are the ones that treat build performance as infrastructure — with dedicated headcount, explicit SLOs (e.g., "p50 incremental build time under 30 seconds"), and the authority to block changes that degrade those SLOs. This is the modern equivalent of the build engineer role from §3.3 — not a person maintaining fragile Makefiles, but a team maintaining build health metrics and the tooling to enforce them.

---

<a id="ch11-6"></a>
### 11.6 Automated detection: monitoring graph health at scale

At the scale where manual intervention fails, automated detection becomes essential. The key is to treat the dependency graph as a monitored system, the same way you monitor production service latency or error rates.

**Critical path length.** Track the longest sequential chain in the build graph over time. An upward trend means the graph is deepening — new dependencies are being added that extend the critical path. This is the single most important build health metric because the critical path is the hard floor on build time that no amount of hardware can break.

**Dependency fanout.** For each target, track the number of direct and transitive dependents. A target with rapidly growing fanout is becoming a hub node — a future base library trap. Alert when a target's fanout exceeds a threshold, and require explicit approval (from the build health team) to add new dependents beyond that threshold.

**Module size distribution.** Track the lines of code and number of source files per target. A target that grows steadily is accumulating responsibility and should be evaluated for splitting. The distribution of module sizes should be roughly log-normal — a few large modules, many medium ones, many small ones. A shift toward larger modules over time is a signal of insufficient splitting discipline.

**Cache hit rates.** Track the ratio of cache hits to cache misses for incremental builds, both locally and on the remote cache. A declining hit rate — especially if the codebase hasn't grown dramatically — indicates that the graph structure is reducing cacheability. Common causes: modules that are too coarse (so any change invalidates a large target), interfaces that are too unstable (so downstream cache keys change frequently), or unnecessary dependencies (so unrelated changes cascade through the graph).

**Build time percentiles.** Track p50, p90, and p99 incremental and clean build times across the engineering population. A creeping p50 is gradual degradation. A spike in p99 often indicates a specific structural change — a new dependency that deepened the critical path, or a module split that inadvertently created a cycle in the dependency ordering.

These metrics should be visible to the engineering organization the way production latency is visible — on dashboards, with alerts, with clear ownership. The team that owns build health reviews these metrics regularly and initiates interventions when they trend in the wrong direction.

---

<a id="ch11-7"></a>
### 11.7 Where AI changes the economics

The interventions in this chapter have three stages: detect the problem, prescribe a fix, and execute the fix. AI's impact is different at each stage.

**Detection: already automatable, AI is not required.** The metrics described in §11.6 — critical path length, dependency fanout, module size, cache hit rates — are computable from the build graph by static analysis. No machine learning is needed. Any build system that maintains an in-memory graph (Bazel, Buck2) can compute these metrics directly. The detection tooling is engineering, not AI.

**Prescription: where AI adds genuine value.** Knowing that a module is too large is easy. Knowing *where to split it* is hard. The right split point depends on:

- Code semantics: which functions and types form a coherent unit?
- Change patterns: which parts change together and which change independently?
- Dependency structure: which split minimizes cross-module dependencies?
- Team ownership: does the split align with who maintains what?
- Build performance impact: does the split actually improve cache hit rates and parallelism?

These are questions that require understanding the code, not just the graph. An engineer answers them by reading the code, understanding the domain, and exercising judgment. An AI that can read both the code and the graph simultaneously can propose split points that balance all five considerations — something that is tedious and error-prone for a human working on a large module, but natural for a system that can process the full context at once.

Concretely, an AI-assisted build optimization tool could:

- Analyze a flagged module and propose two or three candidate split points, each with a predicted impact on cache hit rates and critical path length.
- Identify groups of functions that always change together (high change co-occurrence) and group them into a proposed sub-module.
- Predict the build performance impact of a proposed restructuring *before* it is applied, by simulating the new graph against historical change patterns.
- Identify upcoming graph degradation from in-progress changes — a pre-commit "build health" check that flags "this change adds a dependency that extends the critical path by 12% and will affect 340 downstream targets."

**Execution: where AI may eventually become essential.** Executing a module split at scale requires updating every importer, modifying BUILD files, adjusting visibility declarations, and landing the change atomically across the codebase. Today, this is done by specialized tooling (Rosie at Google, custom scripts elsewhere) that applies mechanical transformations.

AI can make the execution step smarter. Instead of mechanical find-and-replace, an AI-driven refactoring tool can understand the intent of the split, handle edge cases (callers that use the module in unusual ways, test files that need special treatment, generated code that references the old module), and produce a change that is correct on the first pass rather than requiring rounds of manual fixup.

The deeper implication is economic. The techniques in this chapter have always been understood. The reason they weren't applied consistently is that the cost of execution — the coordination, the manual work, the risk of breakage — exceeded the perceived benefit. If AI reduces the execution cost by an order of magnitude — making a module split a one-hour automated operation rather than a one-week cross-team project — then the cost-benefit equation changes fundamentally. Optimizations that were "not worth the effort" become routine maintenance.

This connects to §16.9 (what AI changes and what it doesn't) and §17.9 in the permanent foundations chapter. The dependency graph is still a DAG. The critical path is still the hard floor. Content addressing still provides the cache correctness guarantee. None of the theory changes. But the *maintenance of a healthy graph* — the ongoing work of keeping the build fast as the codebase grows — is where AI changes the economics most significantly. The foundations are permanent. The maintenance is where the leverage is.

---

## Discussion log

| # | Topic | Chapters affected |
|---|-------|------------------|
| 1 | What is a build system; the four jobs; dependency graph basics | Ch 1, Ch 2 |
| 2 | Why graph shape dominates build time; critical path; change propagation | Ch 2 |
| 3 | Local cache correctness; cache key anatomy; stale cache vs. environmental leak; when it doesn't matter | Ch 5, Ch 7 |
| 4 | Why "stale result is mathematically impossible" — the hidden assumption; SHA-256 collision probability; hermeticity as the real fix | Ch 5 |
| 5 | 10,000 engineers on the critical path; independent files (Case 1) vs. chain (Case 2); combinatorial cache miss explosion; interface split; module ownership; submit queue; trunk-based development | Ch 2, Ch 7, Ch 8 |
| 6 | Theory meets reality: why even Google has slow builds; critical path floor; interface changes as worst case; tests as the real bottleneck; submit queue latency; RBE overhead; non-cacheable work | Ch 10 |
| 7 | Clean builds: cold cache, full critical path exposure, toolchain upgrades, cache priming, language choice | Ch 10 |
| 8 | The Build Engineer as a named organizational role; daily build and build cop; SCM discipline; ClearCase; the knowledge transfer failure mode; what Bazel changed about the role | Ch 3 |
| 9 | VCS philosophy: monorepo vs. polyrepo; Google Piper/G4; Amazon two-pizza teams; Git design and its build consequences; the mismatch problem | Ch 9 |
| 10 | Why multiple files exist: cognitive, compilation, organizational, and build performance reasons; Conway's Law; the module as the real unit; build system as organizational tool | Ch 1 |
| 11 | Cycle detection: how Make fails to handle cycles; Makefile engineer's manual audit burden; how Bazel detects cycles at analysis time as a hard error; visibility preventing cycles structurally | Ch 2, Ch 3 |
| 12 | Optimizing the build: graph degradation as a continuous process; three stages of growth; modularization techniques and costs; the restructuring problem; automated detection; AI-assisted prescription and execution | Ch 11 |

---

<a id="appendix-d"></a>
# Appendix D — A history of build systems: milestones, ideas, and turning points

*A chronological record of the ideas, tools, papers, and cultural shifts that shaped how software gets built. Each entry is chosen not just for its technical significance but for what it changed about how engineers think.*

---

## The pre-history: before automation (before 1976)

**Early 1950s–1960s — manual compilation as the norm**
The first programmers compiled programs by hand. On early batch systems, this meant submitting a deck of punched cards to an operator, waiting hours for the result, and collecting the output the next morning. There was no concept of a "build system" — the build was the programmer, following a procedure in their head or written on paper. Dependency management was a human memory problem.

**1967 — Conway's Law stated**
Melvin Conway, in a paper submitted to the *Harvard Business Review* (rejected) and later published in *Datamation*, observed that organizations designing systems produce designs that mirror their communication structure. Not published as a software engineering insight specifically, but would become one of the most quoted principles in the field. The connection to build system design — that the dependency graph reflects the org chart — was not made explicitly for decades.

**Early 1970s — the link editor and separate compilation**
Separate compilation and linking had been standard in Fortran and COBOL since the late 1950s, but it became the dominant C programming model in the early Unix era. The linker — a program that combined separately compiled object files into a single binary — meant that a multi-file program required multiple compilation commands plus a link command, in the correct order. By the early 1970s, even modest Unix programs required a sequence of commands that was tedious to type and easy to get wrong. The need for automation was felt, if not yet filled.

---

## The Make era (1976–1999)

**1976 — Make invented at Bell Labs**
Stuart Feldman, a researcher at Bell Labs, wrote the first version of Make in a weekend in April 1976. The motivation was practical and immediate: he and a colleague had wasted a morning debugging a program that wasn't recompiling correctly because they had forgotten to run the compiler after editing a source file. Make was the solution: a tool that tracked which files were newer than their compiled outputs and ran the necessary commands automatically. The original implementation was around 300 lines of code. Feldman later received the ACM Software System Award in 2003 for this work.

The core ideas in that first weekend implementation — targets, prerequisites, recipes, timestamp-based change detection — remain structurally identical to what GNU Make does today. The tool evolved; the design barely changed.

**1977 — "Make — A Program for Maintaining Computer Programs" published**
Feldman's paper in *Software — Practice and Experience* is the first formal description of Make and its design rationale. It remains readable today and notable for its honesty: Feldman acknowledges the timestamp model's limitations and the manual nature of dependency declaration. The paper describes a tool that solves a real problem pragmatically, with full awareness that it is not a complete solution.

**1979 — Make ships with Unix Version 7**
Make becomes part of the standard Unix distribution, which means it ships on every major Unix workstation of the era — PDP-11s, VAXes, and eventually the Sun workstations that dominated academic and commercial computing in the 1980s. This distribution decision more than any technical feature made Make the universal build tool for the next two decades. Entire generations of C programmers learned Make because it was simply there, part of the environment, the obvious tool for the job.

**1983 — Make ships with BSD Unix**
BSD's version of Make adds features that become standard, including pattern rules (compile any `.c` file the same way without listing each one). The BSD and GNU Make lineages would diverge and reconverge over the following decades, creating a fragmented ecosystem that is still a source of portability headaches today.

**1985 — GNU Make released**
Richard Stallman and Roland McGrath release GNU Make as part of the GNU project. GNU Make adds substantial new features over the original: automatic variables (`$@`, `$<`, `$^`), more powerful pattern matching, computed variable names, and eventually parallel execution via `-j`. GNU Make becomes the dominant Make implementation on Linux and is the version most engineers encounter today. Its manual runs to over 200 pages — a measure of how much complexity had accumulated around a tool designed in a weekend.

**1988 — Recursive Make becomes standard practice**
By the late 1980s, large C/C++ projects had settled on recursive Make as the standard way to manage multi-directory builds: a top-level Makefile invokes `$(MAKE) -C subdir` for each subdirectory, each of which has its own Makefile. This pattern scales the Makefile to larger projects but — as would be documented in detail nearly a decade later — fundamentally breaks the dependency graph by preventing any Makefile from seeing the whole picture.

**1990 — Makefile generators begin appearing**
As the complexity of platform-specific Makefiles grew, tools to generate Makefiles from higher-level descriptions appeared. Imake (used for X Window System builds) was an early example. The existence of Makefile generators — tools that produce the tool that produces the binary — was an early signal that Make itself had become too complex to write directly at scale.

**Mid 1990s — the Build Engineer role peaks**
The combination of large C/C++ codebases, recursive Make, platform portability requirements, and ClearCase or Perforce version control created enough complexity that dedicated Build Engineer headcount was justified at large companies. Sun, HP, SGI, IBM, and Microsoft all maintained build engineering teams. The role would not begin to decline until artifact-based build systems made it unnecessary a decade later.

**1997 — "Recursive Make Considered Harmful" published**
Peter Miller's paper, distributed as a technical report and later widely circulated on the internet, makes the definitive case that recursive Make is broken by design. The core argument: recursive Make prevents any Makefile from seeing the full dependency graph, leading to missed rebuilds, incorrect parallelism, and unpredictable behavior. Miller proposes replacing all recursive Makefiles with a single non-recursive Makefile covering the entire project. The paper was widely read, frequently cited, and largely ignored in practice — the cost of migrating existing recursive builds was too high. But it correctly identified a fundamental design flaw and named it, which is its lasting contribution.

**1998 — autoconf/automake mature**
The autoconf and automake tools — which generate `configure` scripts and `Makefile.am` files respectively — reach maturity and become the standard build infrastructure for open source Unix software. The `./configure && make && make install` incantation becomes the universal way to build open source software. The autoconf ecosystem is extraordinarily complex, built on M4 macro processing, shell scripting, and generated Makefile fragments. It solves real portability problems but is notoriously difficult to understand or modify. Engineers who master it are genuinely rare.

---

## The first wave of alternatives (2000–2008)

**2000 — Apache Ant released**
James Duncan Davidson releases Ant (Another Neat Tool) as the build system for the Tomcat web server. Ant represents the first serious challenge to Make's dominance: XML-based build files, Java-native, with a pluggable task architecture. Ant's approach — platform-independent, expressed in XML rather than shell commands, with a rich library of built-in tasks for Java compilation, JAR packaging, and testing — made it immediately more productive than Make for Java projects. Ant spread rapidly through the Java ecosystem and became the standard Java build tool through the mid-2000s.

Ant did not solve Make's fundamental problem — tasks could still do anything, so safe caching was still impossible — but it made the problem more manageable for Java by providing a rich, portable task library that eliminated the platform variance issues. The XML build files were verbose but comprehensible to Java engineers who were accustomed to XML configuration.

**2000 — SCons released**
SCons (Software Construction tool) takes a different approach: build files are Python scripts, and the tool provides a proper dependency graph with content-based change detection (using MD5 hashes rather than timestamps). SCons is ahead of its time — content hashing for change detection, automatic header dependency scanning for C/C++, a proper Python API instead of a special-purpose language. It is slower than Make for large builds and never achieves wide adoption, but its design anticipates ideas that would become central to Bazel a decade later.

**2001 — CMake initial release**
Kitware releases CMake, initially for cross-platform builds of the VTK visualization toolkit. CMake is a Makefile generator — it generates native build files (Makefiles, Visual Studio project files, Xcode projects) from platform-independent `CMakeLists.txt` files. CMake addresses the platform variance problem that plagued raw Makefiles without trying to replace Make entirely. It becomes the dominant build system for cross-platform C/C++ projects and remains so today, despite widespread frustration with its syntax and behavior.

**2004 — Apache Maven released**
Maven introduces a radical idea: **convention over configuration**. Rather than describing how to build your project, you describe what your project is (a Java library, a web application) and Maven applies standardized conventions (source goes in `src/main/java`, tests in `src/test/java`, the output is a JAR with a standard name). Maven also introduces **central dependency management**: you declare dependencies by group, artifact, and version; Maven downloads them from Maven Central automatically. This is the first widely-adopted artifact repository model — the dependency is a versioned artifact, not source code.

Maven's dependency management fundamentally changed how Java software was distributed and reused. It created the ecosystem model that npm, pip, Cargo, and Go modules would all follow. Its build model is still task-based and still cannot cache safely, but its dependency management represents a genuine conceptual advance.

**2007 — Gradle development begins**
Hans Dockter begins developing Gradle as an attempt to combine Maven's dependency management with a more flexible, Groovy-based build language. Gradle preserves Maven's repository model and project conventions while allowing arbitrary task logic through a rich plugin API. It will grow to become the dominant Android build system and a significant presence in the broader Java/JVM ecosystem. Like all task-based systems, it cannot provide safe caching without explicit configuration — but its incremental build support and build cache (added later) represent the task-based ecosystem's attempt to close the gap with artifact-based systems.

**2007 — Google Blaze development well underway**
Inside Google, the build system that will eventually become Bazel is already in active use. The exact origin is less well-documented publicly, but by 2007 Blaze was the standard build system for most of Google's codebase. The key design decisions — Starlark (then called BUILD language), the sandbox model, content-addressed remote cache — are already in place. Google's build infrastructure at this point is years ahead of anything available to the outside world.

---

## The artifact-based revolution (2009–2015)

**2009 — Ninja build system conceived**
Evan Martin at Google creates Ninja as a build execution engine designed for speed, not expressiveness. Ninja's build files are not meant to be written by hand — they are generated by higher-level tools like CMake and GN. Ninja's design philosophy is the inverse of Make's: do as little as possible, as fast as possible, with no abstraction. It uses content hashing for change detection, builds an explicit dependency graph before execution, and is designed for correctness and speed rather than flexibility. Ninja becomes the execution backend for Chrome's build system and is later adopted by many other projects. It demonstrates that Make's slowness is a design choice, not an inevitability.

**2010 — "Build Systems à la Carte" intellectual foundation laid**
The academic treatment of build system design — formalizing what properties a build system can and cannot have, what different designs trade off — develops through this period. The ideas that will eventually appear in the 2018 paper of the same name by Mokhov, Mitchell, and Peyton Jones are circulating in the research community. The formalization of "suspending," "restarting," and "early cutoff" as distinct build system properties gives the field a vocabulary for comparing designs rigorously.

**2012 — Buck development begins at Facebook**
Facebook engineers, frustrated with the limitations of their existing build tools for the large Hack and Java codebase, begin building Buck. Buck is explicitly inspired by Blaze: BUILD files, a restricted rule language, content-addressed caching, hermetic builds. It is the first Blaze-inspired system to be developed outside Google. Buck will be open sourced in 2013 and become an important data point that Blaze's design is general, not Google-specific.

**2013 — Buck open sourced**
Facebook open sources Buck at the first Mobile@Scale conference. This is a significant moment: it is the first public implementation of the artifact-based build system model. Engineers outside Google can now see a working implementation of content-addressed caching, hermetic sandboxing, and explicit dependency graphs — concepts that had been visible in descriptions of Blaze but never in code. Buck directly influences the design of Pants, Please, and eventually informs the thinking around Bazel's open source release.

**2013 — Pants build system released**
Twitter, Foursquare, and Square jointly release Pants, another Blaze-inspired build system for Python and JVM projects. Pants represents the beginning of a small ecosystem of Blaze-derived systems, each making different tradeoffs. The existence of multiple independent implementations of the same core design — Buck, Pants, and soon Bazel — validates the design's generality and creates a community of engineers thinking seriously about artifact-based build systems.

**2015 — Bazel open sourced by Google**
Google releases Bazel as open source. This is the most significant event in build system history since Make shipped with Unix Version 7. For the first time, the broader engineering community has access to the production implementation of the ideas that had powered Google's build infrastructure for nearly a decade.

The release generates enormous interest and significant frustration. Bazel is designed for Google's infrastructure — monorepo, Piper, RBE — and many of its rough edges come from the difficulty of adapting it to Git-based, polyrepo environments. The documentation is sparse. The external dependency management story is immature. But the core design — hermetic sandboxes, content-addressed caching, the analysis/execution phase separation, Starlark — is sound and unlike anything else publicly available. Engineers who invest in understanding it emerge with a fundamentally different view of what a build system can be.

---

## The JavaScript ecosystem and the modern era (2016–present)

**2016 — Yarn released by Facebook**
Facebook releases Yarn as a faster, more reliable alternative to npm. Yarn introduces the lockfile as a first-class concept — a complete, deterministic record of every package and version in a dependency tree. The lockfile is content addressing applied to dependency management: the same lockfile always produces the same installed packages, on any machine. This is a concrete step toward reproducibility in the JavaScript ecosystem, which had been plagued by "it works on my machine" dependency issues. npm would later adopt the lockfile concept.

**2017 — Remote Build Execution API (REAPI) proposed**
Google and the Bazel community begin developing the Remote Execution API — a standard protocol for communicating between a build client and a remote execution and caching service. The API defines how a client describes a build action (inputs, command, expected outputs), how it queries the content-addressable store, and how it receives results. The significance of REAPI is that it decouples the build client from the execution infrastructure: any build system that speaks REAPI can use any REAPI-compatible remote cache or executor. Bazel, Buck2, Pants, and other systems all eventually implement REAPI, creating a shared infrastructure ecosystem.

**2018 — "Build Systems à la Carte" paper published**
Andrey Mokhov, Neil Mitchell, and Simon Peyton Jones publish a landmark paper in the proceedings of ICFP. The paper provides the first rigorous formal taxonomy of build system designs, classifying them by which scheduling strategy they use (topological, restarting, suspending) and which rebuilding strategy they use (dirty bit, verifying traces, constructive traces). The paper proves formally what practitioners had observed empirically: different designs make different tradeoffs, and no single design dominates all others. It gives the field a precise vocabulary and a theoretical foundation. The paper is notable for being both academically rigorous and practically relevant — it analyzes Make, Excel, Shake, and Bazel using the same formal framework.

**2018 — Nx released**
Nrwl releases Nx, initially as a set of Angular-specific tools for monorepo management, quickly expanding to support the broader TypeScript and JavaScript ecosystem. Nx introduces the concept of the **project graph** — an explicit dependency graph over monorepo projects — as a first-class citizen, with affected-build detection (only build projects affected by a change), task caching, and distributed execution. Nx brings the core ideas of Bazel's design — graph-aware builds, caching, incremental execution — to the JavaScript ecosystem in a form that is accessible to engineers who would never adopt Bazel.

**2020 — Turborepo development begins**
Jared Palmer begins developing Turborepo as a task runner for JavaScript/TypeScript monorepos. Turborepo's design philosophy is radical minimalism: do one thing (run tasks in the right order with caching), do it fast, and stay out of the way. The `turbo.json` pipeline configuration is intentionally simple. The remote cache integration with Vercel is frictionless. Turborepo will be acquired by Vercel in 2021 and open sourced shortly after, becoming the default choice for new JavaScript monorepos that want caching without the complexity of Nx or the overhead of Bazel.

**2021 — Buck2 development begins at Meta**
Meta engineers begin a ground-up rewrite of Buck, driven by lessons accumulated over eight years of operating Buck1 at scale. The key architectural decisions: write in Rust (for performance and correctness), completely separate the build system core from the rules (so rules can evolve without rebuilding the core), use a single incremental dependency graph that covers all phases (rather than separate graphs for loading, analysis, and execution), and make Starlark evaluation fully lazy and parallel. Buck2's architecture represents the current state of the art in build system design.

**2021 — Turborepo open sourced, acquired by Vercel**
Vercel acquires Turborepo and open sources it. The acquisition signals that build infrastructure is now a commercial product category, not just an internal engineering investment. Vercel integrates Turborepo's remote cache with its deployment platform, making remote caching a default part of the JavaScript deployment experience. The lines between build system, CI system, and deployment platform begin to blur.

**2022 — Buck2 open sourced by Meta**
Meta open sources Buck2. The timing — seven years after Bazel's open source release — means the community receives a more mature, more production-hardened design. Buck2's complete core/rules separation is immediately influential: it demonstrates that a build system can be designed so that the rules (which change frequently as language ecosystems evolve) and the core (which should be stable and correct) are fully independent. The Rust implementation delivers significant performance improvements over Bazel for comparable workloads.

**2023 — Bazel MODULE.bazel (Bzlmod) stabilized**
Bazel's external dependency management system, MODULE.bazel (called Bzlmod), reaches stable status after years of development. Bzlmod replaces the fragile WORKSPACE file model with a first-class module system: each Bazel project declares its dependencies in `MODULE.bazel`, and a central registry resolves the dependency graph. This is Make's dependency declaration problem solved at the package manager level: explicit, machine-readable, version-locked, reproducible. It closes one of the most significant gaps between Bazel and the polyrepo/artifact-versioning world.

---

## Key ideas and the year they crystallized

Beyond individual tools, certain ideas changed how the field thinks about build systems. These are the conceptual turning points:

| Year | Idea | Significance |
|------|------|-------------|
| 1976 | Dependency-driven rebuilding | Make: only rebuild what changed. Eliminated the clean-rebuild-everything ritual for most cases. |
| 1976 | The recipe as the unit of work | Make: each target has a command to run. Simple, flexible, and — as it turned out — too flexible for safe caching. |
| 1997 | Recursive Make is broken by design | Miller: the first precise articulation that Make's scalability model was structurally flawed, not just poorly implemented. |
| ~2006 | Hermetic sandboxing | Blaze: if the build action cannot read undeclared inputs, the cache guarantee becomes structural rather than aspirational. |
| ~2006 | Content-addressed build cache | Blaze: cache key = hash of all declared inputs. Same inputs = same output. Changed caching from a performance optimization to a correctness mechanism. |
| ~2006 | Analysis phase before execution phase | Blaze: build the complete dependency graph first, then execute. Enables pre-execution cycle detection, global scheduling, and correctness verification. |
| ~2006 | Restricted rule language | Blaze/Starlark: a language expressive enough for build rules but not expressive enough for arbitrary side effects. The constraint that makes hermeticity enforceable. |
| 2013 | Remote content-addressable cache | Buck/Bazel: extend the cache from one machine to all machines. Same inputs = same key = shared result across the entire team. |
| 2017 | Standardized remote execution protocol | REAPI: decouple build client from execution infrastructure. Any compliant client can use any compliant service. |
| 2018 | Formal taxonomy of build systems | "Build Systems à la Carte": gave the field a precise vocabulary. Made it possible to compare designs rigorously rather than by intuition. |
| ~2020 | Core/rules separation | Buck2: the build system core and the language-specific rules are fully independent modules. Rules can evolve without core changes; core can be optimized without affecting rules. |
| ~2020 | Single incremental graph | Buck2: no separate loading/analysis/execution graphs. One graph, incrementally maintained, covering the entire build. |

---

## The papers and writings that shaped the field

**Stuart Feldman — "Make: A Program for Maintaining Computer Programs" (1977)**
The original Make paper. Short, readable, honest about limitations. The starting point for everything that follows.

**Peter Miller — "Recursive Make Considered Harmful" (1997)**
The definitive diagnosis of Make's scalability failure. Widely cited, rarely acted upon, but important for naming the problem precisely.

**Turing Award Lecture on Unix (Ritchie and Thompson, 1984)**
Not specifically about build systems, but provides essential context for the Unix culture in which Make was developed and spread — a culture that valued small tools, composition, and simplicity that sometimes tipped into underengineering.

**Andrey Mokhov, Neil Mitchell, Simon Peyton Jones — "Build Systems à la Carte" (2018)**
The academic foundation. Provides formal definitions of build system properties and a taxonomy that makes comparison rigorous. Essential reading for anyone who wants to understand the design space.

**Titus Winters, Tom Manshreck, Hyrum Wright — "Software Engineering at Google" (2020), Chapter 18**
The practitioner's account of building and maintaining software at Google's scale. Chapter 18 covers build systems specifically and is the most accessible description of what Blaze/Bazel actually achieves in production.

**Bazel Blog — "Correctness and Reproducibility" series**
A series of posts from the Bazel team explaining the cache correctness guarantee, hermeticity, and the relationship between sandboxing and reproducibility. Technically precise and practically grounded.

---

*End of Appendix D*

---

<a id="ch13"></a>
## Chapter 13 — Beyond the monorepo giants: Gradle, Android, and the JavaScript ecosystem

**One-line pitch:** The build tools most engineers use daily — Gradle for Android, npm/Turborepo/Nx for JavaScript — were not designed by Google or Meta. They were designed by the community, often under commercial pressure, and they carry the marks of that history. Understanding where they came from explains both their power and their limits.

---

<a id="ch13-1"></a>
### 13.1 Maven, Gradle, and Android — three things that look the same and aren't

Android development involves three systems that are easy to conflate because they share vocabulary, coordinate formats, and surface syntax. They are not the same thing.

**Maven** is a build system and the originator of the artifact repository model. It uses `pom.xml` for build configuration and introduced the `group:artifact:version` coordinate format for declaring dependencies. Maven's primary lasting contribution is not its build model but Maven Central — the public repository of published JVM artifacts — and the coordinate system that made sharing Java libraries frictionless. The build model itself is task-based, tied to a fixed lifecycle of phases, and cannot safely cache outputs.

**Maven Central / `maven.google.com`** is an artifact repository — a server that stores versioned JAR files and makes them downloadable by coordinate. This is infrastructure, not a build tool. It is the warehouse. Any build tool that understands the Maven coordinate format — Gradle, Bazel, sbt, Ivy — can use it as a dependency source. Android's dependencies (AndroidX, Google Play Services, third-party libraries) are published here.

**Gradle** is the actual build system for Android apps. It downloads dependencies from Maven repositories using Maven coordinates, which is why it looks like Maven — but Gradle, not Maven, is compiling the Kotlin and Java, processing resources, running tests, and packaging the APK. The `build.gradle` file is a Gradle build file, not a Maven pom.

The confusion is understandable. Maven invented the coordinate format. Gradle adopted it wholesale because all the libraries were already published that way. Using `implementation 'androidx.appcompat:appcompat:1.6.1'` in a Gradle file looks like Maven syntax because Gradle deliberately chose to be compatible with the Maven ecosystem, not because Gradle is Maven.

The precise relationship: Maven created the warehouse and the address format. Gradle is the truck that drives to the warehouse using those addresses. Android Studio is the loading dock. They are three different things at three different layers.

---

<a id="ch13-2"></a>
### 13.2 Gradle and the Android build: why it is slow and why it is hard to fix

Android builds have a reputation as some of the slowest in the industry. A large Android app doing a clean build can take 10–15 minutes. An incremental build after a significant change can take several minutes even on a fast developer machine. This is not an accident of implementation — it is a structural consequence of how the Android build is designed.

**The historical path.** Early Android development (2008–2013) used Eclipse and Apache Ant. Ant had no dependency management — engineers downloaded JARs manually and checked them into `libs/`. When Google launched Android Studio in 2013, it adopted Gradle as the build system. The choice was deliberate: Gradle had Maven's dependency management (solving the JAR-in-version-control problem), a flexible Groovy DSL (better than Maven's XML), and better support for Android's multi-variant build requirements — debug vs. release builds, different product flavors (free vs. paid), multiple ABI targets.

The Android Gradle Plugin (AGP) is the component that adds Android-specific knowledge to Gradle. AGP is large and complex: it encodes how to compile Kotlin and Java sources, how to process XML resources with `aapt2`, how to run ProGuard or R8 for minification and obfuscation, how to compile DEX bytecode for the Dalvik/ART runtime, how to package the APK or AAB, and how to run instrumentation tests on devices or emulators.

**Why it is slow by design.** The Android build pipeline is a deep sequential chain:

```
Kotlin/Java compilation
    → Annotation processing (Dagger, Room, Hilt, kapt)
        → Bytecode transformation (desugaring, instrumentation)
            → R8/ProGuard minification
                → DEX compilation
                    → Resource processing (aapt2)
                        → APK/AAB packaging
                            → Signing
```

Many of these steps have sequential dependencies on each other. Annotation processors, in particular, are a persistent performance problem: they run as a separate compilation phase, can generate new source files, and their outputs must be compiled again — adding extra rounds to the compilation cycle. Kotlin's `kapt` (Kotlin Annotation Processing Tool) compiles Kotlin to Java stubs before running Java annotation processors, adding significant overhead.

**The task-based ceiling.** Gradle is a Generation 1 system. Its incremental build support — added in later versions — works by having task authors declare their inputs and outputs explicitly. If a task declares its inputs correctly, Gradle can skip it when inputs haven't changed. If a task's input declarations are incomplete or a plugin author made a mistake, Gradle conservatively re-runs the task. There is no hermetic sandboxing to enforce completeness — correctness relies on plugin author discipline.

Gradle's build cache can share task outputs across machines if configured, but because the cache key depends on declared inputs rather than verified inputs, cache poisoning from undeclared dependencies is possible. Most Android teams either don't configure a remote build cache or configure it with limited trust.

The result is that Android engineers experience exactly the failure modes Chapter 3 describes: incremental builds that are sometimes wrong (requiring a `./gradlew clean` to recover), builds that pass locally and fail on CI due to environmental differences, and build times that scale poorly as the app grows.

---

<a id="ch13-3"></a>
### 13.3 What Google uses internally for Android — and why it is different

The gap between what Google uses internally and what the industry uses externally is nowhere more visible than in Android development.

**For Android apps at Google** — Gmail, Maps, YouTube, Google Photos, and dozens of others — the build system is Bazel with `rules_android`. These rules encode the same Android-specific build logic that AGP provides for Gradle: resource processing, DEX compilation, APK packaging, testing. But they do it within Bazel's hermetic, content-addressed model.

The practical consequences:

- Clean builds are rare. The remote build cache is warm for essentially all targets because every engineer and every CI run populates it. An engineer picking up a ticket for the first time compiles almost nothing — they download cached artifacts for every unchanged target.
- Incremental builds after a small change are nearly instant. Bazel rebuilds only the affected actions in the dependency graph and nothing else.
- "Passes locally, fails CI" almost never happens for build correctness reasons. The hermetic sandbox ensures that if the build declares its inputs correctly, the output is identical on every machine.
- Large-scale refactoring is feasible. Google's automated refactoring infrastructure (Rosie, LSC) can modify Android app code across dozens of apps simultaneously because the build system understands the complete dependency graph.

**For the Android platform itself** — the Android OS, system apps, the framework — Google uses a custom build system called **Soong**, which replaced the older Make-based build in AOSP around 2017. Soong uses Blueprint files (a JSON-like declarative format) and is designed specifically for AOSP's requirements: cross-compilation to multiple architectures, building the kernel alongside userspace code, managing the complexity of a full operating system build. Parts of AOSP are being progressively migrated to Bazel via the "Mixed Builds" project.

**Why the external world hasn't followed.** `rules_android` is publicly available and actively maintained by Google. Migrating an existing Android app from Gradle/AGP to Bazel is technically feasible. It is also a substantial investment: the Gradle/AGP ecosystem has mature tooling (Android Studio integration, lint, profiling, instant run/apply changes), a large community, and years of accumulated plugin development. Bazel's Android support, while correct and fast, has weaker IDE integration and requires significant upfront investment to configure.

The result is a persistent two-tier world: Google's internal Android teams build in seconds with Bazel; the external Android community builds in minutes with Gradle, accepting the performance cost because the migration cost is higher.

This gap is not specific to Android. It appears everywhere Google's internal infrastructure differs from what they've made available externally: Piper vs. Git, Blaze vs. Bazel, internal RBE vs. self-hosted or commercial RBE. The tools Google open sources are real and useful, but they are optimized for Google's infrastructure. Using them outside that infrastructure requires bridging a gap that the original design didn't need to consider.

---

---

<a id="ch16"></a>
## Chapter 16 — The permanent foundations: what will not change in the AI era

**One-line pitch:** Every major build system concept maps directly to a foundational computer science topic established decades before Bazel existed. These are mathematical and logical truths about computation — not engineering conventions, not technology fashion. AI doesn't change them. Cloud computing didn't change them. Understanding them is an investment that does not depreciate.

---

<a id="ch16-1"></a>
### 16.1 Build systems are applied computer science, not technology fashion

Software engineering is full of fashions. Frameworks rise and fall. Languages come and go. Deployment patterns cycle through decades of reinvention. Engineers who invest heavily in a specific tool or platform often find their knowledge depreciates as the ecosystem moves on.

Build systems look like they might be in this category. Ant replaced Make. Maven replaced Ant. Gradle replaced Maven for Android. Bazel is replacing Gradle in some organizations. Buck2 is competing with Bazel. Something will come after Buck2. Doesn't this knowledge depreciate too?

No — and understanding why not is the entire point of this chapter.

The tools are implementations. Underneath every tool is a set of ideas, and those ideas are grounded in computer science that was established before any of these tools existed and will remain true after all of them are replaced. When you understand *why* Bazel uses content hashing rather than timestamps, you understand something true about the nature of change detection in any caching system — not just Bazel. When you understand *why* the dependency graph must be acyclic, you understand something true about the structure of computation — not just build systems.

The claim of this chapter: every major build system concept is applied computer science. The CS theory underneath does not change. The tools change. The theory doesn't.

This matters especially in the AI era. The question every engineer should ask before investing time in a domain is: will this knowledge still be valuable in five years? Ten years? For build systems, the answer is: the tool-specific knowledge will depreciate, and the foundational knowledge will not. This chapter maps the two categories explicitly so you know which is which.

---

<a id="ch16-2"></a>
### 16.2 Graph theory: the dependency graph is a DAG, forever

The dependency graph is a **Directed Acyclic Graph**. This is not a design choice — it is a mathematical requirement imposed by the nature of computation itself.

A cycle in the dependency graph makes the build logically impossible. To build A you need B; to build B you need C; to build C you need A. There is no valid starting point. This is not a limitation of any particular build system — it is a statement about what it means for one thing to depend on another. If X depends on Y, then Y must exist before X can be built. This ordering constraint, applied transitively, produces a partial order. A partial order with no cycles is a DAG. This is graph theory, developed by Euler in the 18th century and formalized in the 20th.

Everything else about build systems that matters follows directly from the properties of DAGs:

**Topological sort** is the algorithm that produces a valid build order — a linear ordering of vertices such that for every directed edge u → v, u appears before v. A topological sort of the build graph is a sequence of actions that can be executed without violating any dependency. Kahn's algorithm and DFS-based topological sort are standard undergraduate CS. They are correct for any DAG. They will be correct for any build system that represents dependencies as a DAG, which is all of them.

**The critical path** is the longest path from any source to any sink in the DAG. It is the hard lower bound on build time, regardless of how much parallelism is available. You cannot execute two actions in parallel if one depends on the other, so the sequential chain of the critical path must be executed sequentially. No amount of hardware, no amount of AI, no algorithmic improvement can make the build faster than its critical path. This is a theorem about DAG execution, not a limitation of current technology.

**Change propagation** — determining which targets must be rebuilt given a changed source file — is the problem of finding all nodes reachable from the changed node in the DAG. This is graph reachability: BFS or DFS, O(V+E). Any correct incremental build system must solve exactly this problem. The algorithm is the same regardless of the tool.

**Cycle detection** is DFS with a recursion stack — O(V+E). If you encounter a node already on the current DFS path, you have a cycle. Bazel reports it. Buck2 reports it. Any correct build system must detect cycles before attempting execution, because a cycle makes the build logically impossible. The algorithm doesn't change.

An AI system that writes and builds code still produces code where some files depend on other files. Those dependencies still form a graph. That graph must still be acyclic for a valid build to exist. Topological sort is still the ordering algorithm. Critical path is still the speed floor. These will be true as long as computation has the concept of one thing depending on another.

---

<a id="ch16-3"></a>
### 16.3 Hashing and content addressing: from cryptography and distributed systems

**Content addressing** — identifying data by its content rather than its location — is a principle from cryptographic hash functions. SHA-256 is a function that maps arbitrary input to a 256-bit output with three key properties: determinism (same input → same output always), avalanche effect (one bit change → completely different output), and collision resistance (probability of two different inputs producing the same output ≈ 1/2²⁵⁶).

These properties are what make the build cache correctness guarantee work. The cache key for a build action is the SHA-256 hash of all its declared inputs. If any input changes by even one byte, the hash changes completely. A different hash is a cache miss, which triggers a rebuild. For a stale hit to occur — for the system to return a cached result that corresponds to different inputs — the new hash would have to collide with the old one. The probability is 1 in 2²⁵⁶. For practical purposes: impossible.

This is number theory and cryptography. The specific hash function may change (SHA-256 → SHA-3, or a future quantum-resistant function) but the property being exploited — collision resistance — and its implications for cache correctness are permanent.

**The Merkle tree** is the data structure that makes hashing a large input set efficient. A Merkle tree hashes each leaf node (a single file or value), then hashes each parent node as a function of its children's hashes, recursively up to a root hash. If any leaf changes, the change propagates up to the root. To verify that a large set of files is unchanged, you only need to check the root hash — one comparison, regardless of how many files are in the set. If the root matches, all leaves match.

Git uses Merkle trees for its object store. Bitcoin uses them for block chains. Certificate transparency uses them for audit logs. Bazel uses them for dependency fingerprinting. Ralph Merkle invented them in 1979. They will be the right data structure for this problem for as long as the problem exists.

**The completeness theorem** — the cache correctness guarantee holds if and only if all inputs are declared — is a statement about the limits of any formal system. You cannot prove the absence of undeclared inputs by inspecting declared ones. No algorithm exists that can look at a build action and determine whether it has hidden dependencies without executing it in a controlled environment and observing what it accesses. This is why hermeticity is not just a nice-to-have — it is the only complete solution. By physically blocking access to undeclared inputs, the sandbox makes the completeness requirement automatically satisfied: anything the action can access was declared, because anything not declared is inaccessible. This is applied formal logic. It does not change.

---

<a id="ch16-4"></a>
### 16.4 Parallelism and scheduling theory: Amdahl's Law and the critical path

**Amdahl's Law** (Gene Amdahl, 1967): the maximum speedup from parallelism is limited by the sequential fraction of the work. If the critical path of your build constitutes 10% of the total work, no amount of parallel execution can make the build more than 10× faster than the sequential time for that critical path. More precisely: if fraction *s* of the work must be sequential, the maximum speedup is 1/s regardless of the number of processors.

This is mathematics. It applies to build systems, to scientific computing, to any parallelizable workload. An AI-powered build system, a quantum computer, or any future technology still cannot execute step B before step A if B depends on A. The critical path is the sequential fraction. Amdahl's Law bounds the achievable speedup.

**List scheduling** is the class of algorithms that maintain a ready queue of tasks (tasks whose dependencies have all completed) and dispatch them to available workers as workers become free. Build system schedulers are list schedulers. The optimal list scheduling policy — dispatch the task that will block the most downstream work, i.e., the task on the longest remaining critical path — is an NP-hard problem in general but has efficient approximations. These approximations (dispatching by estimated remaining critical path length) are what Bazel, Buck2, and other systems use. Operations research established these results in the 1950s and 1960s. The problem structure hasn't changed.

**Work stealing** — idle workers steal tasks from the queues of busy workers — is a fundamental parallel scheduling technique that achieves optimal expected runtime for certain classes of task graphs. It was analyzed theoretically in the 1990s (Blumofe and Leiserson) and is implemented in Rust's Rayon, Java's ForkJoinPool, Go's goroutine scheduler, and build system execution engines. The theoretical guarantees hold regardless of whether the tasks being scheduled are compilation steps or AI inference calls.

---

<a id="ch16-5"></a>
### 16.5 Caching theory: the fundamental space-time tradeoff

**Caching** is the universal technique of storing the result of an expensive computation to avoid recomputing it when the same inputs recur. The fundamental tradeoff is always **space vs. time**: more cache storage means fewer cache misses means less recomputation means faster execution. This tradeoff is not specific to build systems. It appears in CPU caches, database query caches, web caches, DNS caches, and build artifact caches. It is a permanent feature of any system where computation is expensive and repetition is common.

**Cache replacement policies** — when the cache is full, which entry to evict — are a well-studied problem with known theoretical properties. LRU (least recently used), LFU (least frequently used), FIFO, ARC (adaptive replacement cache) all make different assumptions about access patterns and provide different performance guarantees. The optimal policy (evict the entry that will be accessed furthest in the future — Bélády's algorithm) is theoretically optimal but requires future knowledge, so it serves as a benchmark rather than an implementation. Build systems use LRU or TTL-based eviction for the remote cache. These are correct choices given the access patterns of build artifacts. The theory behind these choices does not change.

**Cache invalidation** is hard precisely because the general problem — knowing when a cached result no longer reflects the current state of the world — requires tracking all dependencies of the cached computation. Build systems solve this with content addressing, which converts the invalidation problem from "has the world changed in a way that affects this result?" to "has any declared input changed?" — a much more tractable question, answerable by hash comparison. The insight that content addressing makes cache invalidation tractable is permanent. It was true before Bazel and will be true after it.

---

<a id="ch16-6"></a>
### 16.6 Formal language theory: why the BUILD file must be restricted

**Starlark** — the language used in Bazel and Buck2 BUILD files — is deliberately designed as a restricted language. It has no I/O, no randomness, no network access, no global mutable state. In some implementations, it disallows arbitrary recursion. This is not a limitation imposed by implementation difficulty. It is a deliberate application of formal language theory.

The reason: if the build description language is **Turing-complete** (can express any computation), the build system cannot reason about what a build rule will do without executing it. This is the **halting problem** — Alan Turing proved in 1936 that no algorithm exists that can determine, for an arbitrary program and input, whether the program will terminate. A Turing-complete build rule could do anything: loop forever, make network requests, generate random outputs, read environment state that isn't declared. If a rule can do anything, the build system cannot safely cache its outputs, because it cannot statically determine what the rule actually produced or what it depended on.

By restricting the language below Turing-completeness — removing arbitrary recursion and I/O — Starlark becomes **statically analyzable**. The build system can determine all dependencies by reading the BUILD file, without executing any build actions. All outputs can be predicted from inputs. Caching is safe.

This is an application of the **Chomsky hierarchy** — the classification of formal languages by their expressive power and the class of automata needed to recognize them (regular languages, context-free languages, context-sensitive languages, recursively enumerable languages). Starlark sits deliberately below the top of the hierarchy, trading expressive power for analyzability. This tradeoff — expressiveness vs. analyzability — is permanent. Any language powerful enough to describe build rules while remaining amenable to static analysis must make this tradeoff. The specific language changes. The tradeoff doesn't.

The implication for the AI era is direct: an AI system that writes BUILD files still must write them in a language that the build system can analyze. If the AI writes Turing-complete build rules, the build system cannot safely cache them. This constraint is not a limitation of current AI — it is a consequence of the halting problem, which is a theorem about computation itself.

---

<a id="ch16-7"></a>
### 16.7 Distributed systems: the remote cache is a distributed store

The remote build cache is a distributed system, and all the fundamental challenges of distributed systems apply.

**The CAP theorem** (Eric Brewer, 2000, proven by Gilbert and Lynch 2002): a distributed data store cannot simultaneously guarantee consistency (every read sees the most recent write), availability (every request receives a response), and partition tolerance (the system continues operating despite network failures). In the presence of network partitions, you must choose between consistency and availability.

The remote build cache chooses **availability over strong consistency**: a cache miss is always safe (you just rebuild the artifact), but blocking on consistency during a network partition would halt the entire build. So the cache is designed to be available even when some nodes are partitioned. The consequence is that a newly written cache entry may not be immediately visible to all clients — eventual consistency. This is the correct tradeoff for a build cache. The reasoning is an application of CAP, which doesn't change.

**The RBE (Remote Build Execution) protocol** is a specific form of **task queue with memoized results** — a pattern that appears in MapReduce, in distributed stream processing, and in build systems. The pattern: describe a unit of work as a pure function of its inputs (content-addressed by digest), check if the result has been computed before (action cache lookup), execute if not (send to a worker), store the result (write to CAS and action cache), return the result to the requester. This pattern will exist in distributed computing as long as computation is expensive and repetition is common.

**Positive network externalities** in the remote cache — more engineers building populates a warmer cache that benefits everyone — is an instance of a network effect, a concept from economics and network science. As the team grows, the expected cache hit rate for any given action increases because the probability that someone else has already built that exact action grows. This is a mathematical property of the access patterns, not a feature of any specific tool.

---

<a id="ch16-8"></a>
### 16.8 Operating systems: sandboxing is kernel-level least privilege

**Hermetic sandboxing** implements the **principle of least privilege** — a security principle articulated by Saltzer and Schroeder in 1975: every component of a system should be able to access only the information and resources necessary for its legitimate purpose, and no more. For a build action, "necessary resources" means the declared inputs. Nothing else.

Linux implements this for build sandboxes through **mount namespaces** (added in kernel 3.8, 2013): a process can have a different view of the filesystem from the rest of the system. The build system creates a mount namespace where only the declared input files are visible — everything else is absent or replaced with empty directories. A build action that tries to read an undeclared file fails immediately with "file not found," surfacing the missing declaration at build time rather than allowing a silent environmental dependency.

**cgroups** (control groups, Linux kernel) provide resource limiting: a sandboxed build action cannot consume more than a specified amount of CPU, memory, or disk I/O. This prevents one runaway build action from starving others on the same machine.

**seccomp** (secure computing mode) restricts which system calls a sandboxed process can make. A compiler doesn't need to create network sockets. By blocking those system calls, the sandbox prevents undeclared network dependencies.

These specific Linux mechanisms will change over time. The concept they implement — OS-enforced least privilege as a mechanism for hermeticity — will not. Future sandboxing might use WebAssembly modules, hardware enclaves, or virtualization technologies. The requirement is the same: the build action must be physically prevented from accessing undeclared inputs, because social conventions and code review cannot achieve the same guarantee.

---

<a id="ch16-9"></a>
### 16.9 What AI changes — and what it doesn't

AI is changing many aspects of software development: how code is written, how it is reviewed, how tests are generated, how documentation is maintained. It is reasonable to ask whether AI changes build systems in ways that invalidate the foundations described in this chapter.

**What AI changes:**

The authoring of BUILD files. An AI can generate correct BUILD file entries from inspecting source code, eliminating a category of manual work that has always been error-prone. This is a real and valuable change.

The detection of performance problems. An AI that understands build graphs can identify overly broad dependencies, base library traps, and cycle risks before they are committed, surfacing problems earlier.

The optimization of build configurations. An AI can suggest optimal parallelism settings, cache configurations, and dependency graph restructurings based on historical build data.

**What AI does not change:**

The dependency graph is still a DAG. Code still has dependencies. Dependencies still form a graph. The graph must still be acyclic for a valid build order to exist.

The critical path is still the speed floor. Even an AI-written, AI-optimized build cannot execute step B before step A if B depends on A. Amdahl's Law applies regardless of who wrote the build rules.

Content addressing is still the mechanism for cache correctness. The hash of a build action's inputs is still the cache key. A changed input still changes the hash. A stale hit still requires a collision. Collision resistance is still the guarantee.

The halting problem is still unsolved. An AI that writes Turing-complete build rules still cannot be safely cached without running them and observing their outputs. The restriction on the rule language is still necessary.

Hermeticity is still the only complete solution to undeclared dependencies. An AI cannot declare dependencies it doesn't know about. If an AI-generated build rule has a hidden dependency — reads an environment variable it didn't declare, accesses a system library it didn't list — the sandbox is still the only mechanism that converts "convention" into "enforcement."

The tools that implement these concepts will be AI-assisted and eventually AI-driven. The concepts themselves are grounded in mathematics and logic that AI does not supersede.

---

<a id="ch16-10"></a>
### 16.10 Why this knowledge does not depreciate

The senior engineer's implicit question when evaluating any domain: is this a good investment of learning time, or will this knowledge be obsolete in three years?

For build systems, the answer splits cleanly across two layers:

**The tool layer depreciates.** Specific Starlark syntax, specific Bazel flags, specific RBE endpoint configuration, specific Buck2 rule APIs — these change with every release. Tool-specific knowledge has a half-life of years.

**The foundation layer does not depreciate.** The DAG structure of dependencies, topological sort as the ordering algorithm, critical path as the speed floor, content hashing as the mechanism for cache correctness, the halting problem as the reason for restricted rule languages, least privilege as the principle behind sandboxing — these are grounded in mathematics, logic, and computer science theory that has been stable for decades and will remain stable indefinitely.

The build engineer who understands only the tools can configure a Bazel build and debug a BUILD file. When Bazel is replaced, they must start over. The build engineer who understands the foundations can evaluate any build system — existing or future — by asking: how does it represent the dependency graph? how does it compute cache keys? what is its rule language and what can it statically analyze? how does it enforce hermeticity? The answers to these questions immediately reveal the system's strengths and limitations, without needing to read its documentation.

This is the difference between knowledge of a map and knowledge of navigation. The map changes. Navigation doesn't.

The build system landscape will continue to evolve. New tools will emerge. AI will automate parts of the process that currently require human effort. The specific protocols, languages, and platforms will change. But as long as software has dependencies — and it will — the dependency graph will be a DAG. As long as builds are expensive — and they will be — caching will require content addressing. As long as caches must be correct — and they must — the halting problem will require restricted rule languages. As long as builds must be reproducible — and they must — hermetic sandboxing will be necessary.

The foundations of build systems are the foundations of computation. They do not expire.

---

# IDEAS INBOX

*Raw ideas captured as they surface. Not yet organized or placed. Domain tags are approximate.*

---

## How to use this section

Drop any idea here in conversation — a question, an observation, a half-formed thought, something you remember from experience, something you think is missing. It gets captured immediately with a domain tag. When ready to organize, each item gets moved into the right chapter or flagged as a new chapter candidate.

**Domain tags used:**
- `[history]` — historical context, how things were done before
- `[theory]` — first principles, concepts, formal properties
- `[practice]` — real-world behavior, war stories, what actually happens at scale
- `[tooling]` — specific tools: Make, Bazel, Buck2, Turborepo, Nx, etc.
- `[org]` — organizational structure, team dynamics, Conway's Law
- `[perf]` — build performance, latency, throughput
- `[correctness]` — cache correctness, hermeticity, reproducibility
- `[vcs]` — version control, branching, monorepo/polyrepo
- `[reader]` — notes about audience, tone, book structure
- `[new-chapter?]` — may warrant its own chapter

---

## Inbox items

**[tooling] [org]** Chapter 13 should be broadened beyond JavaScript to cover the full "ecosystem build systems" landscape — Gradle/Android, JavaScript (Turborepo/Nx), and potentially iOS (Xcode/Tuist). The Android section now has written content (Sections 14.1–14.3). JavaScript sections still need prose. iOS/Xcode is a gap worth considering: Xcode's build system is opaque, slow, and widely hated — the reasons why map directly to the task-based vs. artifact-based distinction and would resonate strongly with mobile engineers.

**[theory] [new-chapter?]** The build system is fundamentally a dependency manager — everything else (change detection, execution ordering, caching) is a consequence of the dependency graph. Dependencies have two axes of complexity: *granularity* (repository → package → module → file → symbol) and *type* (source code, headers, resource files, generated files, compiler/toolchain, compiler flags, environment variables, system libraries, test fixtures, non-deterministic inputs). The completeness problem — declaring not just the dependencies you know about but all that exist — is the root cause of most build failures. Discussed and partially contested: the framing covers roughly half the book well (cache correctness, hermeticity, Make vs. Bazel) but not the other half (scale, organizational coordination, VCS philosophy). The real unifying thread may be broader: the build system as the place where technical, organizational, and infrastructure complexity all become visible simultaneously. Resolve: consider whether this belongs as a dedicated section in Ch 1 or Ch 2, or as a framing note in the introduction.

---

*End of document*
