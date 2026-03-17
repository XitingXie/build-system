// Package metrics persists build and action timing data to a two-table
// JSONL "database" under ~/.cache/build-system/db/.
//
// Schema
//
//	builds.jsonl  — one BuildRow per line  (FK: BuildRow.ID)
//	actions.jsonl — one ActionRow per line (FK: ActionRow.BuildID → BuildRow.ID)
package metrics

import (
	"encoding/json"
	"os"
	"path/filepath"
	"sync"
	"time"
)

// BuildRow is one record in the builds table.
type BuildRow struct {
	ID         int64     `json:"id"`          // unix-nano timestamp, unique per build
	StartedAt  time.Time `json:"started_at"`
	DurationMs int64     `json:"duration_ms"`
	Target     string    `json:"target"`
	Success    bool      `json:"success"`
}

// ActionRow is one record in the actions table.
type ActionRow struct {
	BuildID    int64  `json:"build_id"` // FK → BuildRow.ID
	Label      string `json:"label"`
	CacheHit   bool   `json:"cache_hit"`
	DurationMs int64  `json:"duration_ms"`
	ExitCode   int    `json:"exit_code"`
}

// DB is a lightweight append-only JSONL store with two tables.
type DB struct {
	dir string
	mu  sync.Mutex
}

// DefaultDir returns the directory used by the build CLI.
func DefaultDir() string {
	home, _ := os.UserHomeDir()
	return filepath.Join(home, ".cache", "build-system", "db")
}

// Open opens (or creates) the database in dir.
func Open(dir string) (*DB, error) {
	if err := os.MkdirAll(dir, 0755); err != nil {
		return nil, err
	}
	return &DB{dir: dir}, nil
}

// InsertBuild appends a BuildRow and all associated ActionRows atomically
// (both writes are serialised under the same mutex).
func (db *DB) InsertBuild(b BuildRow, actions []ActionRow) error {
	db.mu.Lock()
	defer db.mu.Unlock()

	if err := appendJSON(filepath.Join(db.dir, "builds.jsonl"), b); err != nil {
		return err
	}
	for _, a := range actions {
		a.BuildID = b.ID
		if err := appendJSON(filepath.Join(db.dir, "actions.jsonl"), a); err != nil {
			return err
		}
	}
	return nil
}

// LoadBuilds reads every BuildRow from builds.jsonl.
func LoadBuilds(dir string) ([]BuildRow, error) {
	return readJSONL[BuildRow](filepath.Join(dir, "builds.jsonl"))
}

// LoadActions reads every ActionRow from actions.jsonl.
func LoadActions(dir string) ([]ActionRow, error) {
	return readJSONL[ActionRow](filepath.Join(dir, "actions.jsonl"))
}

// ── helpers ───────────────────────────────────────────────────────────────────

func appendJSON(path string, v any) error {
	f, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer f.Close()
	return json.NewEncoder(f).Encode(v)
}

func readJSONL[T any](path string) ([]T, error) {
	f, err := os.Open(path)
	if os.IsNotExist(err) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	defer f.Close()

	var rows []T
	dec := json.NewDecoder(f)
	for dec.More() {
		var row T
		if err := dec.Decode(&row); err != nil {
			break // skip malformed tail
		}
		rows = append(rows, row)
	}
	return rows, nil
}
