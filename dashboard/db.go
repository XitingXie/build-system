// Package main — db.go defines the data types that mirror the build system's
// metrics schema and provides functions to load them from the JSONL files.
//
// Schema (files under ~/.cache/build-system/db/):
//
//	builds.jsonl  — one BuildRow per line
//	actions.jsonl — one ActionRow per line, FK: ActionRow.BuildID → BuildRow.ID
package main

import (
	"encoding/json"
	"os"
	"path/filepath"
	"time"
)

// BuildRow mirrors metrics.BuildRow in the build system module.
type BuildRow struct {
	ID         int64     `json:"id"`
	StartedAt  time.Time `json:"started_at"`
	DurationMs int64     `json:"duration_ms"`
	Target     string    `json:"target"`
	Success    bool      `json:"success"`
}

// ActionRow mirrors metrics.ActionRow in the build system module.
type ActionRow struct {
	BuildID    int64  `json:"build_id"`
	Label      string `json:"label"`
	CacheHit   bool   `json:"cache_hit"`
	DurationMs int64  `json:"duration_ms"`
	ExitCode   int    `json:"exit_code"`
}

func defaultDBDir() string {
	home, _ := os.UserHomeDir()
	return filepath.Join(home, ".cache", "build-system", "db")
}

func loadBuilds(dir string) ([]BuildRow, error) {
	return readJSONL[BuildRow](filepath.Join(dir, "builds.jsonl"))
}

func loadActions(dir string) ([]ActionRow, error) {
	return readJSONL[ActionRow](filepath.Join(dir, "actions.jsonl"))
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
			break
		}
		rows = append(rows, row)
	}
	return rows, nil
}
