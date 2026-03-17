package main

import (
	"math"
	"sort"
)

// ── view types ────────────────────────────────────────────────────────────────

type Summary struct {
	TotalBuilds   int
	SuccessRate   float64
	AvgDurationMs int64
	CacheHitRate  float64
	TotalActions  int
}

type RecentBuild struct {
	When       string
	Target     string
	DurationMs int64
	Actions    int
	CacheHits  int
	Success    bool
}

type TargetStat struct {
	Label         string
	Runs          int
	AvgDurationMs int64
	MinDurationMs int64
	MaxDurationMs int64
	CacheHitPct   float64
	SuccessPct    float64
}

type DashData struct {
	Summary      Summary
	RecentBuilds []RecentBuild
	TargetStats  []TargetStat
	GeneratedAt  string
}

// aggregate joins builds + actions and computes dashboard view models.
func aggregate(builds []BuildRow, actions []ActionRow, generatedAt string) DashData {
	data := DashData{GeneratedAt: generatedAt}
	if len(builds) == 0 {
		return data
	}

	// Index actions by build ID for O(1) lookup.
	actionsByBuild := map[int64][]ActionRow{}
	for _, a := range actions {
		actionsByBuild[a.BuildID] = append(actionsByBuild[a.BuildID], a)
	}

	// ── Summary ──────────────────────────────────────────────────────────────
	var totalDur int64
	var successes, totalActions, totalCacheHits int
	for _, b := range builds {
		totalDur += b.DurationMs
		if b.Success {
			successes++
		}
		for _, a := range actionsByBuild[b.ID] {
			totalActions++
			if a.CacheHit {
				totalCacheHits++
			}
		}
	}
	n := len(builds)
	cacheHitRate := 0.0
	if totalActions > 0 {
		cacheHitRate = float64(totalCacheHits) / float64(totalActions) * 100
	}
	data.Summary = Summary{
		TotalBuilds:   n,
		SuccessRate:   float64(successes) / float64(n) * 100,
		AvgDurationMs: totalDur / int64(n),
		CacheHitRate:  cacheHitRate,
		TotalActions:  totalActions,
	}

	// ── Recent builds (last 20, newest first) ─────────────────────────────
	start := 0
	if len(builds) > 20 {
		start = len(builds) - 20
	}
	for i := len(builds) - 1; i >= start; i-- {
		b := builds[i]
		hits := 0
		acts := actionsByBuild[b.ID]
		for _, a := range acts {
			if a.CacheHit {
				hits++
			}
		}
		data.RecentBuilds = append(data.RecentBuilds, RecentBuild{
			When:       b.StartedAt.Format("01-02 15:04:05"),
			Target:     b.Target,
			DurationMs: b.DurationMs,
			Actions:    len(acts),
			CacheHits:  hits,
			Success:    b.Success,
		})
	}

	// ── Per-target stats across all action records ────────────────────────
	type accum struct {
		totalDur int64
		minDur   int64
		maxDur   int64
		runs     int
		hits     int
		ok       int
	}
	byLabel := map[string]*accum{}
	for _, a := range actions {
		acc := byLabel[a.Label]
		if acc == nil {
			acc = &accum{minDur: math.MaxInt64}
			byLabel[a.Label] = acc
		}
		acc.runs++
		acc.totalDur += a.DurationMs
		if a.DurationMs < acc.minDur {
			acc.minDur = a.DurationMs
		}
		if a.DurationMs > acc.maxDur {
			acc.maxDur = a.DurationMs
		}
		if a.CacheHit {
			acc.hits++
		}
		if a.ExitCode == 0 {
			acc.ok++
		}
	}
	for label, acc := range byLabel {
		data.TargetStats = append(data.TargetStats, TargetStat{
			Label:         label,
			Runs:          acc.runs,
			AvgDurationMs: acc.totalDur / int64(acc.runs),
			MinDurationMs: acc.minDur,
			MaxDurationMs: acc.maxDur,
			CacheHitPct:   math.Round(float64(acc.hits) / float64(acc.runs) * 100),
			SuccessPct:    math.Round(float64(acc.ok) / float64(acc.runs) * 100),
		})
	}
	// Sort by avg duration descending (slowest first).
	sort.Slice(data.TargetStats, func(i, j int) bool {
		return data.TargetStats[i].AvgDurationMs > data.TargetStats[j].AvgDurationMs
	})

	return data
}
