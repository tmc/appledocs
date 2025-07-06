# Apple Docs Dashboard - Data Flow Architecture

## Overview

This document describes the data flow architecture for the appledocs performance monitoring system, covering metrics collection, processing, storage, and delivery to dashboard consumers.

## High-Level Architecture

```
┌─────────────────────────────────────────────────────────────────────┐
│                         Metrics Collection Layer                      │
├───────────────┬───────────────┬──────────────┬────────────────────────┤
│ HTTP Client   │ Cache Manager │ URL Processor│ Content Classifier   │
│ • Response    │ • Hit/Miss    │ • Queue Depth│ • Framework Count    │
│ • Errors      │ • Size        │ • Processing │ • Class Count        │
│ • Timing      │ • Age         │   Rate       │ • Method Count       │
└───────┬───────┴───────┬───────┴──────┬───────┴────────┬──────────────┘
        │               │              │                │
        ▼               ▼              ▼                ▼
┌─────────────────────────────────────────────────────────────────────┐
│                      Metrics Aggregation Engine                      │
│  ┌─────────────┐  ┌──────────────┐  ┌──────────────┐  ┌──────────┐ │
│  │ Time-Series │  │  Counters    │  │  Histograms  │  │  Gauges  │ │
│  │   Buffer    │  │  & Rates     │  │ (Latencies)  │  │ (Levels) │ │
│  └─────────────┘  └──────────────┘  └──────────────┘  └──────────┘ │
└───────────────────────────┬─────────────────────────────────────────┘
                           │
                           ▼
┌─────────────────────────────────────────────────────────────────────┐
│                         Storage Layer                                 │
│  ┌──────────────┐  ┌──────────────┐  ┌──────────────┐              │
│  │  In-Memory   │  │    SQLite    │  │  Time-Series │              │
│  │ Ring Buffer  │  │   Database   │  │   Storage    │              │
│  └──────────────┘  └──────────────┘  └──────────────┘              │
└───────────────────────────┬─────────────────────────────────────────┘
                           │
                           ▼
┌─────────────────────────────────────────────────────────────────────┐
│                      API & Delivery Layer                            │
│  ┌──────────────┐  ┌──────────────┐  ┌──────────────┐  ┌─────────┐ │
│  │  REST API    │  │  WebSocket   │  │  Prometheus  │  │Terminal │ │
│  │   Server     │  │    Server    │  │   Exporter   │  │   UI    │ │
│  └──────────────┘  └──────────────┘  └──────────────┘  └─────────┘ │
└─────────────────────────────────────────────────────────────────────┘
```

## Detailed Component Architecture

### 1. Metrics Collection Pipeline

```go
// Core metrics collection pipeline
type MetricsPipeline struct {
    collectors []MetricsCollector
    processors []MetricsProcessor
    sinks      []MetricsSink
}

// Collector interface
type MetricsCollector interface {
    Collect(ctx context.Context) (*RawMetric, error)
    Name() string
}

// Processor interface
type MetricsProcessor interface {
    Process(metric *RawMetric) (*ProcessedMetric, error)
}

// Sink interface
type MetricsSink interface {
    Write(metric *ProcessedMetric) error
}
```

#### Collection Points

1. **HTTP Operations**
   ```go
   type HTTPMetricsCollector struct {
       requestStart    time.Time
       requestEnd      time.Time
       statusCode      int
       bytesTransferred int64
       url             string
   }
   ```

2. **Cache Operations**
   ```go
   type CacheMetricsCollector struct {
       operation  string // "hit", "miss", "evict"
       key        string
       size       int64
       timestamp  time.Time
   }
   ```

3. **Processing Operations**
   ```go
   type ProcessingMetricsCollector struct {
       documentType string
       startTime    time.Time
       endTime      time.Time
       success      bool
       errorType    string
   }
   ```

### 2. Data Aggregation Strategy

#### Real-time Aggregation (1-second window)
```go
type RealtimeAggregator struct {
    window    time.Duration
    buffer    *RingBuffer
    mutex     sync.RWMutex
}

func (a *RealtimeAggregator) Aggregate() AggregatedMetrics {
    a.mutex.RLock()
    defer a.mutex.RUnlock()
    
    return AggregatedMetrics{
        Timestamp:        time.Now(),
        RequestsPerSec:   a.buffer.Rate(),
        AvgResponseTime:  a.buffer.Average("response_time"),
        CurrentBandwidth: a.buffer.Sum("bytes") / a.window.Seconds(),
    }
}
```

#### Time-Window Aggregation
```
┌─────────────────────────────────────────────────────┐
│ Raw Metrics (1ms)                                   │
│ • Individual HTTP requests                          │
│ • Cache operations                                  │
│ • Processing events                                 │
└─────────────────────┬───────────────────────────────┘
                     │
                     ▼
┌─────────────────────────────────────────────────────┐
│ 1-Second Aggregates                                 │
│ • Requests/sec                                      │
│ • Current bandwidth                                 │
│ • Active connections                                │
└─────────────────────┬───────────────────────────────┘
                     │
                     ▼
┌─────────────────────────────────────────────────────┐
│ 1-Minute Aggregates                                 │
│ • Average rates                                     │
│ • Error percentages                                 │
│ • Cache hit ratios                                  │
└─────────────────────┬───────────────────────────────┘
                     │
                     ▼
┌─────────────────────────────────────────────────────┐
│ 1-Hour Aggregates                                   │
│ • Hourly summaries                                  │
│ • Trend calculations                                │
│ • Anomaly detection                                 │
└─────────────────────────────────────────────────────┘
```

### 3. Storage Architecture

#### In-Memory Storage (Ring Buffer)
```go
type RingBuffer struct {
    data     []MetricPoint
    capacity int
    head     int
    tail     int
    mutex    sync.RWMutex
}

type MetricPoint struct {
    Timestamp time.Time
    Values    map[string]float64
}

// Efficient O(1) insert and read
func (rb *RingBuffer) Insert(point MetricPoint) {
    rb.mutex.Lock()
    defer rb.mutex.Unlock()
    
    rb.data[rb.head] = point
    rb.head = (rb.head + 1) % rb.capacity
    if rb.head == rb.tail {
        rb.tail = (rb.tail + 1) % rb.capacity
    }
}
```

#### SQLite Schema
```sql
-- High-frequency metrics (1-second resolution)
CREATE TABLE metrics_realtime (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    timestamp INTEGER NOT NULL,
    metric_type TEXT NOT NULL,
    value REAL NOT NULL,
    tags TEXT, -- JSON encoded tags
    INDEX idx_timestamp_type (timestamp, metric_type)
) WITHOUT ROWID;

-- Aggregated metrics (1-minute resolution)
CREATE TABLE metrics_aggregated (
    timestamp INTEGER NOT NULL,
    metric_type TEXT NOT NULL,
    count INTEGER,
    sum REAL,
    min REAL,
    max REAL,
    avg REAL,
    p50 REAL,
    p95 REAL,
    p99 REAL,
    PRIMARY KEY (timestamp, metric_type)
) WITHOUT ROWID;

-- Session metadata
CREATE TABLE crawl_sessions (
    session_id TEXT PRIMARY KEY,
    start_time INTEGER NOT NULL,
    end_time INTEGER,
    total_processed INTEGER,
    total_errors INTEGER,
    configuration TEXT -- JSON
);
```

#### Time-Series Optimization
```go
// Circular buffer with time-based partitioning
type TimeSeriesStore struct {
    partitions map[int64]*Partition // Unix hour -> Partition
    retention  time.Duration
}

type Partition struct {
    hour     int64
    metrics  []CompressedMetric
    index    map[string][]int // metric_name -> positions
}

// Compression for storage efficiency
type CompressedMetric struct {
    BaseTime   int32    // Seconds since partition start
    MetricID   uint16   // Mapped metric name
    Value      float32  // Compressed value
    Tags       []byte   // Compressed tags
}
```

### 4. Real-time Streaming Architecture

#### WebSocket Protocol
```go
type MetricsWebSocketHub struct {
    clients    map[*Client]bool
    broadcast  chan MetricUpdate
    register   chan *Client
    unregister chan *Client
}

type MetricUpdate struct {
    Type      string                 `json:"type"`
    Timestamp time.Time              `json:"timestamp"`
    Data      map[string]interface{} `json:"data"`
}

// Client subscription management
type Client struct {
    conn         *websocket.Conn
    send         chan []byte
    subscriptions map[string]bool // Metric types to receive
}
```

#### Message Protocol
```json
// Client -> Server (Subscribe)
{
    "action": "subscribe",
    "metrics": ["realtime.requests", "cache.hit_rate", "errors.rate"]
}

// Server -> Client (Update)
{
    "type": "metric.update",
    "timestamp": "2024-01-15T12:34:56Z",
    "data": {
        "realtime.requests": 45.2,
        "cache.hit_rate": 0.873,
        "errors.rate": 0.004
    }
}

// Server -> Client (Alert)
{
    "type": "alert",
    "severity": "warning",
    "message": "Cache hit rate below threshold",
    "metric": "cache.hit_rate",
    "value": 0.65,
    "threshold": 0.70
}
```

### 5. API Endpoints

#### REST API Structure
```yaml
/api/v1/metrics:
  /current:
    GET: Get current snapshot of all metrics
  
  /history:
    GET: Get historical metrics
    parameters:
      - from: timestamp
      - to: timestamp
      - resolution: 1s|1m|1h|1d
      - metrics: comma-separated list
  
  /stream:
    GET: WebSocket endpoint for real-time updates
  
  /export:
    GET: Export metrics in various formats
    parameters:
      - format: json|csv|prometheus
      - period: last_hour|last_day|last_week|custom

/api/v1/sessions:
  GET: List crawl sessions
  /{session_id}:
    GET: Get session details
    /metrics:
      GET: Get metrics for specific session

/api/v1/alerts:
  GET: Get alert configuration
  POST: Create/update alert rules
  /{alert_id}:
    DELETE: Remove alert rule
```

### 6. Data Processing Pipeline

```go
// Pipeline stages
type MetricsPipeline struct {
    stages []PipelineStage
}

type PipelineStage interface {
    Process(ctx context.Context, in <-chan Metric) <-chan Metric
}

// Example pipeline
pipeline := NewMetricsPipeline(
    NewCollectionStage(),      // Collect raw metrics
    NewValidationStage(),      // Validate and clean
    NewEnrichmentStage(),      // Add calculated fields
    NewAggregationStage(),     // Time-window aggregation
    NewPersistenceStage(),     // Store to database
    NewBroadcastStage(),       // Send to subscribers
)
```

#### Data Enrichment
```go
type EnrichmentStage struct {
    enrichers []MetricEnricher
}

type MetricEnricher interface {
    Enrich(metric *Metric) error
}

// Example enrichers
type RateCalculator struct{}
func (r *RateCalculator) Enrich(m *Metric) error {
    if m.Type == "counter" {
        m.Fields["rate"] = calculateRate(m)
    }
    return nil
}

type PercentileCalculator struct{}
func (p *PercentileCalculator) Enrich(m *Metric) error {
    if m.Type == "histogram" {
        m.Fields["p50"] = calculatePercentile(m, 50)
        m.Fields["p95"] = calculatePercentile(m, 95)
        m.Fields["p99"] = calculatePercentile(m, 99)
    }
    return nil
}
```

### 7. Performance Optimization

#### Batch Processing
```go
type BatchProcessor struct {
    batchSize    int
    flushInterval time.Duration
    buffer       []Metric
    mu           sync.Mutex
}

func (b *BatchProcessor) Add(metric Metric) {
    b.mu.Lock()
    b.buffer = append(b.buffer, metric)
    
    if len(b.buffer) >= b.batchSize {
        b.flush()
    }
    b.mu.Unlock()
}

func (b *BatchProcessor) flush() {
    if len(b.buffer) == 0 {
        return
    }
    
    // Process batch
    batch := b.buffer
    b.buffer = make([]Metric, 0, b.batchSize)
    
    go b.processBatch(batch)
}
```

#### Memory Management
```go
// Object pooling for frequent allocations
var metricPool = sync.Pool{
    New: func() interface{} {
        return &Metric{
            Fields: make(map[string]float64, 10),
            Tags:   make(map[string]string, 5),
        }
    },
}

// Reuse metric objects
func getMetric() *Metric {
    return metricPool.Get().(*Metric)
}

func putMetric(m *Metric) {
    m.Reset()
    metricPool.Put(m)
}
```

### 8. Data Retention & Cleanup

```go
type RetentionManager struct {
    policies []RetentionPolicy
    ticker   *time.Ticker
}

type RetentionPolicy struct {
    MetricType string
    Resolution time.Duration
    Retention  time.Duration
}

func (r *RetentionManager) cleanup() {
    for _, policy := range r.policies {
        cutoff := time.Now().Add(-policy.Retention)
        r.deleteOldMetrics(policy.MetricType, cutoff)
    }
}

// Efficient deletion with minimal locking
func (r *RetentionManager) deleteOldMetrics(metricType string, cutoff time.Time) {
    // Use database partitioning for efficient bulk deletes
    query := `DELETE FROM metrics_aggregated 
              WHERE metric_type = ? AND timestamp < ?`
    r.db.Exec(query, metricType, cutoff.Unix())
}
```

### 9. Export Formats

#### Prometheus Format
```go
func (e *PrometheusExporter) Export(metrics []Metric) string {
    var output strings.Builder
    
    for _, metric := range metrics {
        // TYPE declaration
        fmt.Fprintf(&output, "# TYPE %s %s\n", metric.Name, metric.Type)
        
        // HELP text
        if metric.Help != "" {
            fmt.Fprintf(&output, "# HELP %s %s\n", metric.Name, metric.Help)
        }
        
        // Metric line
        labels := formatLabels(metric.Tags)
        fmt.Fprintf(&output, "%s%s %f %d\n", 
            metric.Name, labels, metric.Value, metric.Timestamp.Unix())
    }
    
    return output.String()
}
```

#### JSON Streaming
```go
func (e *JSONExporter) StreamExport(w io.Writer, metrics <-chan Metric) error {
    encoder := json.NewEncoder(w)
    
    // Start JSON array
    w.Write([]byte("["))
    first := true
    
    for metric := range metrics {
        if !first {
            w.Write([]byte(","))
        }
        first = false
        
        if err := encoder.Encode(metric); err != nil {
            return err
        }
    }
    
    // End JSON array
    w.Write([]byte("]"))
    return nil
}
```

### 10. Alert Processing

```go
type AlertEngine struct {
    rules      []AlertRule
    evalPeriod time.Duration
    notifiers  []Notifier
}

type AlertRule struct {
    ID         string
    MetricName string
    Condition  string // e.g., "> 90", "< 50"
    Duration   time.Duration
    Severity   string
}

func (a *AlertEngine) evaluate(metric Metric) {
    for _, rule := range a.rules {
        if rule.matches(metric) {
            alert := Alert{
                Rule:      rule,
                Value:     metric.Value,
                Timestamp: time.Now(),
            }
            a.notify(alert)
        }
    }
}
```

## Data Flow Scenarios

### Scenario 1: Real-time Monitoring
```
User opens dashboard → WebSocket connection established → 
Subscribe to metrics → Server sends updates every second → 
Client renders charts → Connection maintained until close
```

### Scenario 2: Historical Analysis
```
User requests 24h data → API validates request → 
Query aggregated data from SQLite → Apply resolution sampling → 
Compress response → Send to client → Client renders analysis
```

### Scenario 3: Alert Triggered
```
Metric exceeds threshold → Alert engine evaluates → 
Alert created → Notifications sent → Dashboard updated → 
Alert logged to database → User acknowledges
```

## Security & Performance Considerations

### Security
- Rate limiting on API endpoints
- WebSocket connection limits
- Input validation on all metrics
- Sanitization of user-provided tags

### Performance
- Metric batching to reduce overhead
- Compression for network transmission
- Indexed storage for fast queries
- Caching of frequently accessed data

This data flow architecture provides a scalable, efficient foundation for real-time monitoring and historical analysis of the appledocs crawler performance.