// dashboard is a standalone web server that reads build metrics recorded by
// the toy build system and presents them as an HTML dashboard.
//
// Usage:
//
//	go run . [--addr :8080] [--db ~/.cache/build-system/db]
//
// The dashboard is independent of the build system binary — it only reads
// the two JSONL files that the build system writes.
package main

import (
	"flag"
	"fmt"
	"html/template"
	"net/http"
	"os"
	"time"
)

func main() {
	addr := flag.String("addr", ":8080", "listen address")
	dbDir := flag.String("db", defaultDBDir(), "path to the build-system db directory")
	flag.Parse()

	if err := serve(*addr, *dbDir); err != nil {
		fmt.Fprintf(os.Stderr, "dashboard: %v\n", err)
		os.Exit(1)
	}
}

func serve(addr, dbDir string) error {
	tmpl := template.Must(
		template.New("dash").Funcs(template.FuncMap{
			"pct": func(f float64) string { return fmt.Sprintf("%.1f%%", f) },
			"ms": func(ms int64) string {
				if ms < 1000 {
					return fmt.Sprintf("%dms", ms)
				}
				return fmt.Sprintf("%.2fs", float64(ms)/1000)
			},
			"bar": func(pct float64) int { // pixel width 0–80
				if pct < 0 {
					return 0
				}
				if pct > 100 {
					return 80
				}
				return int(pct * 80 / 100)
			},
			// muldiv computes hits/total*100 safely (returns 0 when total==0).
			"muldiv": func(hits, total int) float64 {
				if total == 0 {
					return 0
				}
				return float64(hits) / float64(total) * 100
			},
		}).Parse(dashHTML),
	)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		builds, err := loadBuilds(dbDir)
		if err != nil {
			http.Error(w, "load builds: "+err.Error(), 500)
			return
		}
		actions, err := loadActions(dbDir)
		if err != nil {
			http.Error(w, "load actions: "+err.Error(), 500)
			return
		}
		data := aggregate(builds, actions, time.Now().Format("2006-01-02 15:04:05"))
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		if err := tmpl.Execute(w, data); err != nil {
			http.Error(w, err.Error(), 500)
		}
	})

	// JSON API endpoints — useful for external tooling or charting libraries.
	http.HandleFunc("/api/builds", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, dbDir+"/builds.jsonl")
	})
	http.HandleFunc("/api/actions", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, dbDir+"/actions.jsonl")
	})

	fmt.Printf("Build dashboard → http://localhost%s\n", addr)
	fmt.Printf("DB directory   → %s\n", dbDir)
	fmt.Printf("API            → /api/builds  /api/actions\n")
	fmt.Println("Press Ctrl-C to stop.")
	return http.ListenAndServe(addr, nil)
}

// ── HTML template ─────────────────────────────────────────────────────────────

const dashHTML = `<!DOCTYPE html>
<html lang="en">
<head>
<meta charset="utf-8">
<meta name="viewport" content="width=device-width,initial-scale=1">
<title>Build Dashboard</title>
<style>
*{box-sizing:border-box;margin:0;padding:0}
body{font-family:-apple-system,BlinkMacSystemFont,"Segoe UI",Roboto,sans-serif;
     background:#f0f2f5;color:#111827;font-size:14px;line-height:1.5}
a{color:inherit;text-decoration:none}

/* ── header ── */
header{background:#111827;color:#f9fafb;padding:0 24px;height:56px;
       display:flex;align-items:center;justify-content:space-between;
       box-shadow:0 1px 4px rgba(0,0,0,.3)}
header h1{font-size:16px;font-weight:600;letter-spacing:.3px;display:flex;align-items:center;gap:8px}
header .meta{font-size:12px;color:#9ca3af}
header .refresh{font-size:12px;color:#60a5fa;cursor:pointer;padding:4px 10px;
                border:1px solid #374151;border-radius:6px;background:none;color:#9ca3af}
header .refresh:hover{background:#1f2937;color:#f9fafb}

/* ── layout ── */
.page{max-width:1280px;margin:24px auto;padding:0 20px}

/* ── cards ── */
.cards{display:grid;grid-template-columns:repeat(auto-fit,minmax(200px,1fr));gap:16px;margin-bottom:24px}
.card{background:#fff;border-radius:10px;padding:20px 22px;
      box-shadow:0 1px 3px rgba(0,0,0,.07),0 1px 2px rgba(0,0,0,.05)}
.card .lbl{font-size:11px;font-weight:600;text-transform:uppercase;letter-spacing:.7px;
           color:#6b7280;margin-bottom:8px}
.card .val{font-size:30px;font-weight:700;color:#111827;line-height:1}
.card .sub{font-size:12px;color:#9ca3af;margin-top:6px}

/* ── section ── */
.section{background:#fff;border-radius:10px;
         box-shadow:0 1px 3px rgba(0,0,0,.07);margin-bottom:24px;overflow:hidden}
.sec-hdr{padding:14px 20px;border-bottom:1px solid #f3f4f6;
         font-weight:600;font-size:13px;color:#374151;
         display:flex;align-items:center;justify-content:space-between}
.sec-hdr .hint{font-size:11px;font-weight:400;color:#9ca3af}

/* ── table ── */
table{width:100%;border-collapse:collapse}
th{padding:9px 16px;text-align:left;font-size:11px;font-weight:600;
   text-transform:uppercase;letter-spacing:.5px;color:#6b7280;
   background:#f9fafb;border-bottom:1px solid #e5e7eb}
td{padding:10px 16px;border-bottom:1px solid #f3f4f6;font-size:13px;color:#374151}
tr:last-child td{border-bottom:none}
tbody tr:hover td{background:#f9fafb}
code{font-family:"SF Mono","Fira Code",monospace;font-size:12px;
     background:#f3f4f6;padding:2px 5px;border-radius:4px}

/* ── badges ── */
.badge{display:inline-flex;align-items:center;padding:2px 8px;
       border-radius:20px;font-size:11px;font-weight:600;white-space:nowrap}
.ok  {background:#d1fae5;color:#065f46}
.fail{background:#fee2e2;color:#991b1b}
.warn{background:#fef3c7;color:#92400e}

/* ── bar ── */
.bar-wrap{display:inline-block;background:#e5e7eb;border-radius:3px;
          height:6px;width:80px;vertical-align:middle;margin-left:6px;overflow:hidden}
.bar-fill{height:100%;border-radius:3px;background:#3b82f6}
.bar-fill.green{background:#10b981}

/* ── empty state ── */
.empty{padding:40px;text-align:center;color:#9ca3af}
.empty p{margin-top:8px;font-size:13px}
</style>
</head>
<body>
<header>
  <h1>
    <svg width="18" height="18" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
      <polyline points="22 12 18 12 15 21 9 3 6 12 2 12"/>
    </svg>
    Build Dashboard
  </h1>
  <div style="display:flex;align-items:center;gap:16px">
    <span class="meta">Updated {{.GeneratedAt}}</span>
    <button class="refresh" onclick="location.reload()">&#8635; Refresh</button>
  </div>
</header>

<div class="page">

{{if eq .Summary.TotalBuilds 0}}
<div class="section">
  <div class="empty">
    <svg width="40" height="40" viewBox="0 0 24 24" fill="none" stroke="#d1d5db" stroke-width="1.5">
      <rect x="3" y="3" width="18" height="18" rx="2"/>
      <path d="M3 9h18M9 21V9"/>
    </svg>
    <p>No build data yet.</p>
    <p>Run <code>build &lt;target&gt;</code> from your workspace to start recording.</p>
  </div>
</div>
{{else}}

<!-- ── Summary cards ── -->
<div class="cards">
  <div class="card">
    <div class="lbl">Total Builds</div>
    <div class="val">{{.Summary.TotalBuilds}}</div>
    <div class="sub">all time</div>
  </div>
  <div class="card">
    <div class="lbl">Success Rate</div>
    <div class="val">{{pct .Summary.SuccessRate}}</div>
    <div class="sub">of builds succeeded</div>
  </div>
  <div class="card">
    <div class="lbl">Avg Build Time</div>
    <div class="val">{{ms .Summary.AvgDurationMs}}</div>
    <div class="sub">wall time</div>
  </div>
  <div class="card">
    <div class="lbl">Cache Hit Rate</div>
    <div class="val">{{pct .Summary.CacheHitRate}}</div>
    <div class="sub">{{.Summary.TotalActions}} total actions</div>
  </div>
</div>

<!-- ── Recent builds ── -->
<div class="section">
  <div class="sec-hdr">
    Recent Builds
    <span class="hint">last 20, newest first</span>
  </div>
  <table>
    <thead><tr>
      <th>When</th>
      <th>Target</th>
      <th>Duration</th>
      <th>Actions</th>
      <th>Cache Hits</th>
      <th>Status</th>
    </tr></thead>
    <tbody>
    {{range .RecentBuilds}}
    <tr>
      <td style="color:#6b7280">{{.When}}</td>
      <td><code>{{.Target}}</code></td>
      <td>{{ms .DurationMs}}</td>
      <td>{{.Actions}}</td>
      <td>
        {{.CacheHits}}/{{.Actions}}
        {{if gt .Actions 0}}
        <span class="bar-wrap">
          <span class="bar-fill green" style="width:{{bar (muldiv .CacheHits .Actions)}}px"></span>
        </span>
        {{end}}
      </td>
      <td>
        {{if .Success}}<span class="badge ok">&#10003; OK</span>
        {{else}}<span class="badge fail">&#10007; FAIL</span>{{end}}
      </td>
    </tr>
    {{end}}
    </tbody>
  </table>
</div>

<!-- ── Per-target stats ── -->
<div class="section">
  <div class="sec-hdr">
    Target Performance
    <span class="hint">sorted by avg duration (slowest first)</span>
  </div>
  <table>
    <thead><tr>
      <th>Target</th>
      <th>Runs</th>
      <th>Avg</th>
      <th>Min</th>
      <th>Max</th>
      <th>Cache Hit %</th>
      <th>Success %</th>
    </tr></thead>
    <tbody>
    {{range .TargetStats}}
    <tr>
      <td><code>{{.Label}}</code></td>
      <td>{{.Runs}}</td>
      <td>{{ms .AvgDurationMs}}</td>
      <td style="color:#6b7280">{{ms .MinDurationMs}}</td>
      <td style="color:#6b7280">{{ms .MaxDurationMs}}</td>
      <td>
        {{pct .CacheHitPct}}
        <span class="bar-wrap">
          <span class="bar-fill" style="width:{{bar .CacheHitPct}}px"></span>
        </span>
      </td>
      <td>
        {{if eq .SuccessPct 100.0}}
          <span class="badge ok">{{pct .SuccessPct}}</span>
        {{else if eq .SuccessPct 0.0}}
          <span class="badge fail">{{pct .SuccessPct}}</span>
        {{else}}
          <span class="badge warn">{{pct .SuccessPct}}</span>
        {{end}}
      </td>
    </tr>
    {{end}}
    </tbody>
  </table>
</div>

{{end}}
</div>
</body>
</html>
`
