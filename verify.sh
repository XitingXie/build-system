#!/usr/bin/env bash
# verify.sh — end-to-end verification of the toy build system.
#
# Tests exercised:
#   1. Fresh build: every target shows [build]
#   2. Cached build: every target shows [cached] (no source changes)
#   3. Incremental rebuild: modifying a source file invalidates only the
#      affected targets; unrelated targets stay [cached]
#   4. query deps / query rdeps
#   5. graph command
#   6. build clean
#   7. build //... (wildcard — build all targets)
#
# Exit codes: 0 = all checks passed, non-zero = first failure.

set -euo pipefail

REPO_ROOT="$(cd "$(dirname "$0")" && pwd)"
BIN="$REPO_ROOT/.verify-build-bin"

# ── colour helpers ────────────────────────────────────────────────────────────
GREEN='\033[0;32m'
RED='\033[0;31m'
BOLD='\033[1m'
RESET='\033[0m'

pass() { echo -e "${GREEN}  PASS${RESET}  $1"; }
fail() { echo -e "${RED}  FAIL${RESET}  $1"; exit 1; }
header() { echo -e "\n${BOLD}=== $1 ===${RESET}"; }

# ── helpers ───────────────────────────────────────────────────────────────────
# count_lines_matching <pattern> <output>
count_lines_matching() { echo "$2" | grep -c "$1" || true; }

run_build() {
    # Always run from REPO_ROOT so the binary can find BUILD files via os.Getwd().
    (cd "$REPO_ROOT" && "$BIN" "$@" 2>&1)
}

# ── step 0: compile the binary ────────────────────────────────────────────────
header "Step 0: Compile build binary"
(cd "$REPO_ROOT" && go build -o "$BIN" ./cmd/build)
pass "binary compiled → $BIN"

# ── step 1: clean slate ───────────────────────────────────────────────────────
header "Step 1: Clean cache"
run_build clean
pass "cache cleared"

# ── step 2: fresh build ───────────────────────────────────────────────────────
header "Step 2: Fresh build of //examples/greeting:greeting"
OUT=$(run_build //examples/greeting:greeting)
echo "$OUT"

built=$(count_lines_matching "\[build\]" "$OUT")
cached=$(count_lines_matching "\[cached\]" "$OUT")

[ "$built" -ge 3 ] \
    || fail "Expected ≥3 [build] lines on first build, got $built"
[ "$cached" -eq 0 ] \
    || fail "Expected 0 [cached] lines on first build, got $cached"
pass "all targets built from scratch ($built [build], $cached [cached])"

# Verify the output file was actually written.
GREETING_OUT="$REPO_ROOT/out/examples/greeting/greeting.out"
[ -f "$GREETING_OUT" ] \
    || fail "Output file not found: $GREETING_OUT"
CONTENT=$(cat "$GREETING_OUT")
echo "  greeting.out contents: $(echo "$CONTENT" | tr '\n' ' ')"
echo "$CONTENT" | grep -qi "hello" \
    || fail "greeting.out does not contain 'Hello'"
echo "$CONTENT" | grep -qi "world" \
    || fail "greeting.out does not contain 'World'"
pass "output file has correct content"

# ── step 3: cached build ──────────────────────────────────────────────────────
header "Step 3: Cached build (no changes)"
OUT=$(run_build //examples/greeting:greeting)
echo "$OUT"

built=$(count_lines_matching "\[build\]" "$OUT")
cached=$(count_lines_matching "\[cached\]" "$OUT")

[ "$built" -eq 0 ] \
    || fail "Expected 0 [build] lines on cached build, got $built"
[ "$cached" -ge 3 ] \
    || fail "Expected ≥3 [cached] lines on cached build, got $cached"
pass "all targets served from cache ($cached [cached], $built [build])"

# ── step 4: incremental rebuild after source change ───────────────────────────
header "Step 4: Incremental rebuild after modifying hello.txt"
echo "Hi" > "$REPO_ROOT/examples/greeting/hello.txt"

OUT=$(run_build //examples/greeting:greeting)
echo "$OUT"

# hello and greeting must be rebuilt; world must stay cached.
echo "$OUT" | grep -q "\[build\].*hello" \
    || fail "Expected //examples/greeting:hello to be rebuilt"
echo "$OUT" | grep -q "\[build\].*greeting" \
    || fail "Expected //examples/greeting:greeting to be rebuilt"
echo "$OUT" | grep -q "\[cached\].*world" \
    || fail "Expected //examples/greeting:world to remain cached"
pass "only affected targets rebuilt"

# Restore hello.txt.
echo "Hello" > "$REPO_ROOT/examples/greeting/hello.txt"

# ── step 5: cross-package build ───────────────────────────────────────────────
header "Step 5: Cross-package build //examples/upper:upper"
OUT=$(run_build //examples/upper:upper)
echo "$OUT"

UPPER_OUT="$REPO_ROOT/out/examples/upper/upper.out"
[ -f "$UPPER_OUT" ] || fail "Output file not found: $UPPER_OUT"
CONTENT=$(cat "$UPPER_OUT")
echo "  upper.out contents: $(echo "$CONTENT" | tr '\n' ' ')"
echo "$CONTENT" | grep -q "HELLO" \
    || fail "upper.out does not contain 'HELLO'"
echo "$CONTENT" | grep -q "WORLD" \
    || fail "upper.out does not contain 'WORLD'"
pass "cross-package build produced correct uppercased output"

# ── step 6: query deps ────────────────────────────────────────────────────────
header "Step 6: query deps //examples/greeting:greeting"
OUT=$(run_build query deps //examples/greeting:greeting)
echo "$OUT"
echo "$OUT" | grep -q "//examples/greeting:hello" \
    || fail "deps missing //examples/greeting:hello"
echo "$OUT" | grep -q "//examples/greeting:world" \
    || fail "deps missing //examples/greeting:world"
pass "query deps returned correct dependencies"

# ── step 7: query rdeps ───────────────────────────────────────────────────────
header "Step 7: query rdeps //examples/greeting:hello"
OUT=$(run_build query rdeps //examples/greeting:hello)
echo "$OUT"
echo "$OUT" | grep -q "//examples/greeting:greeting" \
    || fail "rdeps missing //examples/greeting:greeting"
pass "query rdeps returned correct reverse dependencies"

# ── step 8: graph ─────────────────────────────────────────────────────────────
header "Step 8: graph //examples/greeting:greeting"
OUT=$(run_build graph //examples/greeting:greeting)
echo "$OUT"
echo "$OUT" | grep -q "//examples/greeting:greeting" \
    || fail "graph output missing root target"
echo "$OUT" | grep -q "//examples/greeting:hello" \
    || fail "graph output missing dep //examples/greeting:hello"
pass "graph command produced dependency tree"

# ── step 9: build all (//...) ─────────────────────────────────────────────────
header "Step 9: build //... (wildcard)"
run_build clean
OUT=$(run_build //...)
echo "$OUT"
built=$(count_lines_matching "\[build\]" "$OUT")
[ "$built" -ge 4 ] \
    || fail "Expected ≥4 [build] lines for //..., got $built"
pass "wildcard build built $built targets"

# ── step 10: clean ────────────────────────────────────────────────────────────
header "Step 10: Final clean"
OUT=$(run_build clean)
echo "$OUT"
echo "$OUT" | grep -qi "cleared" || fail "clean did not report success"
pass "cache cleaned"

# ── done ──────────────────────────────────────────────────────────────────────
echo ""
echo -e "${GREEN}${BOLD}All checks passed.${RESET}"

# Clean up the temporary binary.
rm -f "$BIN"
