# Apple Docs Performance Dashboard - Implementation Guide

## Executive Summary

This guide provides a comprehensive implementation roadmap for the appledocs performance monitoring dashboard. Based on the analysis of current metrics collection, we recommend a phased hybrid approach that starts with a terminal UI and progressively adds web-based monitoring capabilities.

## Quick Start Implementation

### Phase 1: Terminal Dashboard (Week 1)

Create a basic terminal dashboard that integrates with the existing metrics:

```go
// cmd/dashboard/main.go
package main

import (
    "context"
    "fmt"
    "time"
    
    ui "github.com/gizak/termui/v3"
    "github.com/gizak/termui/v3/widgets"
)

type Dashboard struct {
    metrics *appledocs
    
    // UI widgets
    processedGauge *widgets.Gauge
    cacheHitGauge  *widgets.Gauge
    errorGauge     *widgets.Gauge
    speedPlot      *widgets.Plot
    errorList      *widgets.List
    downloadsList  *widgets.List
}

func NewDashboard(app *appledocs) *Dashboard {
    return &Dashboard{
        metrics:        app,
        processedGauge: widgets.NewGauge(),
        cacheHitGauge:  widgets.NewGauge(),
        errorGauge:     widgets.NewGauge(),
        speedPlot:      widgets.NewPlot(),
        errorList:      widgets.NewList(),
        downloadsList:  widgets.NewList(),
    }
}

func (d *Dashboard) Run(ctx context.Context) error {
    if err := ui.Init(); err != nil {
        return fmt.Errorf("failed to initialize termui: %v", err)
    }
    defer ui.Close()
    
    d.setupLayout()
    d.render()
    
    ticker := time.NewTicker(time.Second)
    defer ticker.Stop()
    
    uiEvents := ui.PollEvents()
    for {
        select {
        case <-ctx.Done():
            return ctx.Err()
        case <-ticker.C:
            d.update()
            d.render()
        case e := <-uiEvents:
            switch e.ID {
            case "q", "<C-c>":
                return nil
            case "r":
                d.refresh()
            }
        }
    }
}
```

### Phase 2: Metrics API (Week 2)

Add HTTP endpoints to expose metrics:

```go
// metrics_api.go
package main

import (
    "encoding/json"
    "net/http"
    "time"
)

type MetricsAPI struct {
    app *appledocs
}

func (api *MetricsAPI) RegisterHandlers(mux *http.ServeMux) {
    mux.HandleFunc("/api/metrics/current", api.handleCurrentMetrics)
    mux.HandleFunc("/api/metrics/stream", api.handleMetricsStream)
    mux.HandleFunc("/api/metrics/history", api.handleHistoricalMetrics)
}

func (api *MetricsAPI) handleCurrentMetrics(w http.ResponseWriter, r *http.Request) {
    metrics := api.app.getEnhancedMetrics()
    
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(metrics)
}

func (api *MetricsAPI) handleMetricsStream(w http.ResponseWriter, r *http.Request) {
    // Upgrade to WebSocket
    conn, err := upgrader.Upgrade(w, r, nil)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }
    defer conn.Close()
    
    ticker := time.NewTicker(time.Second)
    defer ticker.Stop()
    
    for {
        select {
        case <-ticker.C:
            metrics := api.app.getEnhancedMetrics()
            if err := conn.WriteJSON(metrics); err != nil {
                return
            }
        }
    }
}
```

### Phase 3: Web Dashboard (Week 3)

Create a simple web dashboard:

```html
<!-- dashboard/index.html -->
<!DOCTYPE html>
<html>
<head>
    <title>Apple Docs Monitor</title>
    <script src="https://cdn.jsdelivr.net/npm/chart.js"></script>
    <style>
        body {
            font-family: -apple-system, BlinkMacSystemFont, sans-serif;
            background: #0a0a0a;
            color: #fff;
            margin: 0;
            padding: 20px;
        }
        .dashboard-grid {
            display: grid;
            grid-template-columns: repeat(auto-fit, minmax(250px, 1fr));
            gap: 20px;
        }
        .metric-card {
            background: #1a1a1a;
            border-radius: 8px;
            padding: 20px;
            border: 1px solid #3a3a3a;
        }
        .metric-value {
            font-size: 2.5em;
            font-weight: 300;
            margin: 10px 0;
        }
        .metric-label {
            color: #b0b0b0;
            text-transform: uppercase;
            font-size: 0.8em;
            letter-spacing: 1px;
        }
        .trend-up { color: #00ff88; }
        .trend-down { color: #ff3366; }
    </style>
</head>
<body>
    <h1>Apple Docs Performance Monitor</h1>
    
    <div class="dashboard-grid">
        <div class="metric-card">
            <div class="metric-label">Documents Processed</div>
            <div class="metric-value" id="processed">-</div>
            <div class="metric-trend" id="processed-trend"></div>
        </div>
        
        <div class="metric-card">
            <div class="metric-label">Cache Hit Rate</div>
            <div class="metric-value" id="cache-hit-rate">-</div>
            <div class="metric-trend" id="cache-trend"></div>
        </div>
        
        <div class="metric-card">
            <div class="metric-label">Error Rate</div>
            <div class="metric-value" id="error-rate">-</div>
            <div class="metric-trend" id="error-trend"></div>
        </div>
        
        <div class="metric-card">
            <div class="metric-label">Download Speed</div>
            <div class="metric-value" id="download-speed">-</div>
            <div class="metric-trend" id="speed-trend"></div>
        </div>
    </div>
    
    <div style="margin-top: 40px;">
        <canvas id="processingChart" width="400" height="100"></canvas>
    </div>
    
    <script src="dashboard.js"></script>
</body>
</html>
```

```javascript
// dashboard/dashboard.js
class AppleDocsDashboard {
    constructor() {
        this.ws = null;
        this.chart = null;
        this.chartData = [];
        this.maxDataPoints = 60;
        
        this.initWebSocket();
        this.initChart();
    }
    
    initWebSocket() {
        const protocol = window.location.protocol === 'https:' ? 'wss:' : 'ws:';
        this.ws = new WebSocket(`${protocol}//${window.location.host}/api/metrics/stream`);
        
        this.ws.onmessage = (event) => {
            const metrics = JSON.parse(event.data);
            this.updateMetrics(metrics);
        };
        
        this.ws.onerror = (error) => {
            console.error('WebSocket error:', error);
        };
        
        this.ws.onclose = () => {
            // Reconnect after 5 seconds
            setTimeout(() => this.initWebSocket(), 5000);
        };
    }
    
    updateMetrics(metrics) {
        // Update KPI cards
        document.getElementById('processed').textContent = 
            metrics.processed.toLocaleString();
        
        document.getElementById('cache-hit-rate').textContent = 
            `${metrics.cache_hit_rate.toFixed(1)}%`;
        
        const errorRate = (metrics.errors / metrics.processed * 100) || 0;
        document.getElementById('error-rate').textContent = 
            `${errorRate.toFixed(2)}%`;
        
        document.getElementById('download-speed').textContent = 
            `${metrics.download_rate_mbps.toFixed(1)} MB/s`;
        
        // Update chart
        this.updateChart(metrics.processing_rate_per_sec);
    }
    
    initChart() {
        const ctx = document.getElementById('processingChart').getContext('2d');
        this.chart = new Chart(ctx, {
            type: 'line',
            data: {
                labels: [],
                datasets: [{
                    label: 'Documents/sec',
                    data: [],
                    borderColor: '#00ff88',
                    backgroundColor: 'rgba(0, 255, 136, 0.1)',
                    borderWidth: 2,
                    pointRadius: 0,
                    tension: 0.4
                }]
            },
            options: {
                responsive: true,
                maintainAspectRatio: false,
                scales: {
                    y: {
                        beginAtZero: true,
                        grid: { color: '#3a3a3a' },
                        ticks: { color: '#b0b0b0' }
                    },
                    x: {
                        grid: { display: false },
                        ticks: { color: '#b0b0b0' }
                    }
                },
                plugins: {
                    legend: { display: false }
                }
            }
        });
    }
    
    updateChart(value) {
        const now = new Date().toLocaleTimeString();
        
        this.chartData.push(value);
        if (this.chartData.length > this.maxDataPoints) {
            this.chartData.shift();
        }
        
        this.chart.data.labels = this.chartData.map((_, i) => 
            i === this.chartData.length - 1 ? now : ''
        );
        this.chart.data.datasets[0].data = this.chartData;
        this.chart.update('none'); // Disable animation for performance
    }
}

// Initialize dashboard
document.addEventListener('DOMContentLoaded', () => {
    new AppleDocsDashboard();
});
```

## Key Implementation Decisions

### 1. Start with Terminal UI
- **Rationale**: Minimal dependencies, immediate value
- **Timeline**: 3-4 days
- **Benefits**: Works over SSH, low resource usage

### 2. Progressive Enhancement
- **Phase 1**: Terminal dashboard (Week 1)
- **Phase 2**: HTTP API (Week 2)
- **Phase 3**: Web dashboard (Week 3)
- **Phase 4**: Grafana integration (Optional)

### 3. Storage Strategy
- **Short-term**: In-memory ring buffer
- **Medium-term**: SQLite for persistence
- **Long-term**: Time-series database option

### 4. Real-time Updates
- **Terminal**: 1-second refresh rate
- **Web**: WebSocket for live updates
- **API**: REST + WebSocket hybrid

## Recommended Metrics Enhancements

### 1. Add Response Time Percentiles
```go
type ResponseTimeHistogram struct {
    values []time.Duration
    mutex  sync.RWMutex
}

func (h *ResponseTimeHistogram) Add(d time.Duration) {
    h.mutex.Lock()
    h.values = append(h.values, d)
    if len(h.values) > 1000 {
        h.values = h.values[1:]
    }
    h.mutex.Unlock()
}

func (h *ResponseTimeHistogram) Percentile(p float64) time.Duration {
    h.mutex.RLock()
    defer h.mutex.RUnlock()
    
    if len(h.values) == 0 {
        return 0
    }
    
    sorted := make([]time.Duration, len(h.values))
    copy(sorted, h.values)
    sort.Slice(sorted, func(i, j int) bool {
        return sorted[i] < sorted[j]
    })
    
    idx := int(float64(len(sorted)-1) * p / 100)
    return sorted[idx]
}
```

### 2. Track Framework-Specific Metrics
```go
type FrameworkMetrics struct {
    metrics map[string]*FrameworkStats
    mutex   sync.RWMutex
}

type FrameworkStats struct {
    DocumentCount int
    LastUpdated   time.Time
    ErrorCount    int
    AvgSize       int64
}

func (f *FrameworkMetrics) Update(framework string, size int64, err error) {
    f.mutex.Lock()
    defer f.mutex.Unlock()
    
    if f.metrics == nil {
        f.metrics = make(map[string]*FrameworkStats)
    }
    
    stats, exists := f.metrics[framework]
    if !exists {
        stats = &FrameworkStats{}
        f.metrics[framework] = stats
    }
    
    stats.DocumentCount++
    stats.LastUpdated = time.Now()
    if err != nil {
        stats.ErrorCount++
    }
    
    // Update average size
    stats.AvgSize = (stats.AvgSize*int64(stats.DocumentCount-1) + size) / 
                   int64(stats.DocumentCount)
}
```

## Deployment Configuration

### 1. Add Dashboard Flag
```go
// main.go additions
var (
    enableDashboard = flag.Bool("dashboard", false, "enable performance dashboard")
    dashboardPort   = flag.String("dashboard-port", "8080", "dashboard HTTP port")
    terminalUI      = flag.Bool("terminal-ui", false, "show terminal dashboard")
)

func main() {
    flag.Parse()
    
    app := &appledocs{
        // ... existing initialization
    }
    
    if *enableDashboard {
        api := &MetricsAPI{app: app}
        mux := http.NewServeMux()
        api.RegisterHandlers(mux)
        
        // Serve static files
        mux.Handle("/", http.FileServer(http.Dir("dashboard")))
        
        go func() {
            log.Printf("Dashboard available at http://localhost:%s", *dashboardPort)
            log.Fatal(http.ListenAndServe(":"+*dashboardPort, mux))
        }()
    }
    
    if *terminalUI {
        dashboard := NewDashboard(app)
        go dashboard.Run(context.Background())
    }
    
    // ... rest of main function
}
```

### 2. Docker Support
```dockerfile
# Dockerfile
FROM golang:1.21 AS builder
WORKDIR /app
COPY . .
RUN go mod download
RUN go build -o appledocs .

FROM alpine:latest
RUN apk add --no-cache ca-certificates
WORKDIR /root/
COPY --from=builder /app/appledocs .
COPY --from=builder /app/dashboard ./dashboard

EXPOSE 8080
CMD ["./appledocs", "-dashboard", "-dashboard-port", "8080"]
```

### 3. Systemd Service
```ini
[Unit]
Description=Apple Docs Crawler with Dashboard
After=network.target

[Service]
Type=simple
User=appledocs
WorkingDirectory=/opt/appledocs
ExecStart=/opt/appledocs/appledocs -dashboard -terminal-ui
Restart=on-failure
RestartSec=10

# Performance tuning
LimitNOFILE=65536
Environment="GOGC=100"

[Install]
WantedBy=multi-user.target
```

## Testing Strategy

### 1. Unit Tests
```go
func TestMetricsCollection(t *testing.T) {
    app := &appledocs{
        httpErrors: make(map[int]int),
        startTime:  time.Now(),
    }
    
    // Test metric recording
    app.recordResponseTime(100 * time.Millisecond)
    app.recordBytesDownloaded(1024)
    app.recordHTTPError(404)
    
    metrics := app.getEnhancedMetrics()
    
    if metrics.AvgResponseTime != 100*time.Millisecond {
        t.Errorf("Expected avg response time 100ms, got %v", 
                 metrics.AvgResponseTime)
    }
}
```

### 2. Integration Tests
```go
func TestDashboardAPI(t *testing.T) {
    app := &appledocs{
        httpErrors: make(map[int]int),
        startTime:  time.Now(),
    }
    
    api := &MetricsAPI{app: app}
    
    req := httptest.NewRequest("GET", "/api/metrics/current", nil)
    w := httptest.NewRecorder()
    
    api.handleCurrentMetrics(w, req)
    
    if w.Code != http.StatusOK {
        t.Errorf("Expected status 200, got %d", w.Code)
    }
    
    var metrics MetricsSnapshot
    if err := json.NewDecoder(w.Body).Decode(&metrics); err != nil {
        t.Errorf("Failed to decode response: %v", err)
    }
}
```

## Performance Optimization Tips

1. **Use atomic operations** for metric updates to avoid lock contention
2. **Batch metric writes** to reduce storage overhead
3. **Implement sampling** for high-frequency metrics
4. **Use compression** for historical data storage
5. **Cache aggregated values** to reduce computation

## Next Steps

1. **Implement Phase 1** - Terminal dashboard (3-4 days)
2. **Add basic metrics API** - HTTP endpoints (2-3 days)
3. **Create web dashboard** - Simple HTML/JS (1 week)
4. **Add persistence layer** - SQLite storage (2-3 days)
5. **Optional: Grafana integration** - For advanced users (2-3 days)

This implementation guide provides a practical roadmap for adding comprehensive performance monitoring to the appledocs application, starting with minimal changes and progressively adding more sophisticated features.