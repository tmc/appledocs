# Apple Docs Mirror Tool - Developer Integration Guide

## Overview

This guide provides comprehensive information for developers who want to integrate appledocs into their applications, extend its functionality, or use it as a library component. It covers both programmatic usage patterns and extension development.

## Table of Contents

- [Using appledocs as a Library](#using-appledocs-as-a-library)
- [Extending with Custom Processors](#extending-with-custom-processors)
- [Integration Patterns](#integration-patterns)
- [Performance Tuning](#performance-tuning)
- [Data Pipeline Integration](#data-pipeline-integration)
- [Custom Output Formats](#custom-output-formats)
- [Metrics and Monitoring Integration](#metrics-and-monitoring-integration)
- [Testing and Quality Assurance](#testing-and-quality-assurance)

## Using appledocs as a Library

### Basic Library Integration

While appledocs is primarily designed as a CLI tool, you can integrate its components into Go applications:

```go
package main

import (
    "context"
    "fmt"
    "log"
    "net/http"
    "time"
    
    "golang.org/x/time/rate"
)

// DocumentationMirror wraps appledocs functionality for library use
type DocumentationMirror struct {
    client      *http.Client
    baseURL     string
    cacheDir    string
    outputDir   string
    rateLimiter *rate.Limiter
}

// NewDocumentationMirror creates a new documentation mirror instance
func NewDocumentationMirror(config Config) *DocumentationMirror {
    return &DocumentationMirror{
        client: &http.Client{
            Timeout: config.Timeout,
        },
        baseURL:     config.BaseURL,
        cacheDir:    config.CacheDir,
        outputDir:   config.OutputDir,
        rateLimiter: rate.NewLimiter(rate.Limit(config.RateLimit), int(config.RateLimit)),
    }
}

// Config holds configuration for the documentation mirror
type Config struct {
    BaseURL     string
    CacheDir    string
    OutputDir   string
    Timeout     time.Duration
    RateLimit   float64
    Concurrency int
}

// DefaultConfig returns sensible defaults
func DefaultConfig() Config {
    return Config{
        BaseURL:     "https://developer.apple.com",
        CacheDir:    ".cache",
        OutputDir:   "output",
        Timeout:     30 * time.Second,
        RateLimit:   10.0,
        Concurrency: 5,
    }
}

// MirrorFramework mirrors documentation for a specific framework
func (dm *DocumentationMirror) MirrorFramework(ctx context.Context, framework string) error {
    log.Printf("Starting mirror for framework: %s", framework)
    
    // Implementation would use appledocs internal functions
    // This is a simplified example showing the interface
    
    entryURL := fmt.Sprintf("%s/tutorials/data/documentation/%s.json", dm.baseURL, framework)
    
    // Use rate limiter
    if err := dm.rateLimiter.Wait(ctx); err != nil {
        return fmt.Errorf("rate limit wait failed: %v", err)
    }
    
    // Fetch and process (simplified)
    resp, err := dm.client.Get(entryURL)
    if err != nil {
        return fmt.Errorf("failed to fetch %s: %v", entryURL, err)
    }
    defer resp.Body.Close()
    
    log.Printf("Successfully mirrored framework: %s", framework)
    return nil
}

// Usage example
func main() {
    config := DefaultConfig()
    mirror := NewDocumentationMirror(config)
    
    ctx := context.Background()
    if err := mirror.MirrorFramework(ctx, "SwiftUI"); err != nil {
        log.Fatalf("Mirror failed: %v", err)
    }
}
```

### Validation Library Integration

```go
package main

import (
    "fmt"
    "log"
    "os"
)

// ValidationService wraps appledocs validation functionality
type ValidationService struct {
    cacheDir string
    enableChecksums bool
}

// NewValidationService creates a validation service
func NewValidationService(cacheDir string, enableChecksums bool) *ValidationService {
    return &ValidationService{
        cacheDir: cacheDir,
        enableChecksums: enableChecksums,
    }
}

// ValidateCache validates the cache directory
func (vs *ValidationService) ValidateCache() (*ValidationReport, error) {
    // Use appledocs validation functions
    var result ValidationResult
    
    if vs.enableChecksums {
        result = ValidateCacheIntegrityWithChecksums(vs.cacheDir)
    } else {
        result = ValidateCache(vs.cacheDir)
    }
    
    report := &ValidationReport{
        Valid:       result.Valid,
        ErrorCount:  len(result.Errors),
        WarningCount: len(result.Warnings),
        Errors:      make([]string, len(result.Errors)),
        Warnings:    make([]string, len(result.Warnings)),
    }
    
    for i, err := range result.Errors {
        report.Errors[i] = err.Error()
    }
    
    for i, warn := range result.Warnings {
        report.Warnings[i] = warn.Error()
    }
    
    return report, nil
}

// ValidationReport provides a structured validation result
type ValidationReport struct {
    Valid        bool     `json:"valid"`
    ErrorCount   int      `json:"error_count"`
    WarningCount int      `json:"warning_count"`
    Errors       []string `json:"errors,omitempty"`
    Warnings     []string `json:"warnings,omitempty"`
}

// Example usage
func main() {
    validator := NewValidationService(".cache", true)
    
    report, err := validator.ValidateCache()
    if err != nil {
        log.Fatalf("Validation failed: %v", err)
    }
    
    fmt.Printf("Cache validation: %s\n", validationStatus(report.Valid))
    fmt.Printf("Errors: %d, Warnings: %d\n", report.ErrorCount, report.WarningCount)
    
    if !report.Valid {
        fmt.Println("Validation errors:")
        for _, err := range report.Errors {
            fmt.Printf("  - %s\n", err)
        }
    }
}

func validationStatus(valid bool) string {
    if valid {
        return "PASSED"
    }
    return "FAILED"
}
```

## Extending with Custom Processors

### Custom Content Processor

```go
package main

import (
    "encoding/json"
    "fmt"
    "strings"
)

// ContentProcessor interface for custom processing
type ContentProcessor interface {
    ProcessDocument(doc *DocJSONData, path string) error
    GetProcessorName() string
}

// APIExtractor extracts API information from documentation
type APIExtractor struct {
    apis map[string]APIInfo
}

// APIInfo holds extracted API information
type APIInfo struct {
    Name        string   `json:"name"`
    Type        string   `json:"type"`
    Platforms   []string `json:"platforms"`
    Abstract    string   `json:"abstract"`
    Deprecated  bool     `json:"deprecated"`
    Beta        bool     `json:"beta"`
}

// NewAPIExtractor creates a new API extractor
func NewAPIExtractor() *APIExtractor {
    return &APIExtractor{
        apis: make(map[string]APIInfo),
    }
}

// ProcessDocument extracts API information from a document
func (ae *APIExtractor) ProcessDocument(doc *DocJSONData, path string) error {
    if doc.Metadata.Title == "" {
        return nil // Skip documents without titles
    }
    
    api := APIInfo{
        Name: doc.Metadata.Title,
        Type: doc.Metadata.Role,
    }
    
    // Extract platform information
    for _, platform := range doc.Metadata.Platforms {
        api.Platforms = append(api.Platforms, platform.Name)
        if platform.Deprecated {
            api.Deprecated = true
        }
        if platform.Beta {
            api.Beta = true
        }
    }
    
    // Extract abstract
    if len(doc.Abstract) > 0 && doc.Abstract[0].Text != "" {
        api.Abstract = doc.Abstract[0].Text
    }
    
    ae.apis[path] = api
    return nil
}

// GetProcessorName returns the processor name
func (ae *APIExtractor) GetProcessorName() string {
    return "APIExtractor"
}

// Export exports the collected API information
func (ae *APIExtractor) Export() ([]byte, error) {
    return json.MarshalIndent(ae.apis, "", "  ")
}

// GetDeprecatedAPIs returns all deprecated APIs
func (ae *APIExtractor) GetDeprecatedAPIs() []APIInfo {
    var deprecated []APIInfo
    for _, api := range ae.apis {
        if api.Deprecated {
            deprecated = append(deprecated, api)
        }
    }
    return deprecated
}

// GetAPIsByPlatform returns APIs for a specific platform
func (ae *APIExtractor) GetAPIsByPlatform(platform string) []APIInfo {
    var apis []APIInfo
    for _, api := range ae.apis {
        for _, p := range api.Platforms {
            if p == platform {
                apis = append(apis, api)
                break
            }
        }
    }
    return apis
}

// Custom processor usage example
func main() {
    extractor := NewAPIExtractor()
    
    // Process documents (integration with appledocs crawling)
    processedCount := 0
    
    // Simulate processing multiple documents
    sampleDocs := []string{
        "tutorials/data/documentation/SwiftUI.json",
        "tutorials/data/documentation/UIKit.json",
    }
    
    for _, docPath := range sampleDocs {
        // In real implementation, load actual documents
        doc := &DocJSONData{
            Metadata: Metadata{
                Title: extractFrameworkName(docPath),
                Role:  "framework",
                Platforms: []Platform{
                    {Name: "iOS", IntroducedAt: "13.0"},
                    {Name: "macOS", IntroducedAt: "10.15"},
                },
            },
            Abstract: []TextContent{
                {Text: fmt.Sprintf("Framework for %s", extractFrameworkName(docPath))},
            },
        }
        
        if err := extractor.ProcessDocument(doc, docPath); err != nil {
            fmt.Printf("Error processing %s: %v\n", docPath, err)
            continue
        }
        processedCount++
    }
    
    fmt.Printf("Processed %d documents with %s\n", processedCount, extractor.GetProcessorName())
    
    // Export results
    data, err := extractor.Export()
    if err != nil {
        fmt.Printf("Export failed: %v\n", err)
        return
    }
    
    fmt.Printf("Extracted API data:\n%s\n", string(data))
    
    // Get deprecated APIs
    deprecated := extractor.GetDeprecatedAPIs()
    fmt.Printf("Found %d deprecated APIs\n", len(deprecated))
    
    // Get iOS specific APIs
    iosAPIs := extractor.GetAPIsByPlatform("iOS")
    fmt.Printf("Found %d iOS APIs\n", len(iosAPIs))
}

func extractFrameworkName(path string) string {
    parts := strings.Split(path, "/")
    if len(parts) > 0 {
        name := parts[len(parts)-1]
        return strings.TrimSuffix(name, ".json")
    }
    return "Unknown"
}
```

### Custom Output Format Generator

```go
package main

import (
    "encoding/csv"
    "fmt"
    "os"
    "path/filepath"
    "strings"
)

// CSVGenerator generates CSV reports from Apple documentation
type CSVGenerator struct {
    outputPath string
    records    [][]string
}

// NewCSVGenerator creates a new CSV generator
func NewCSVGenerator(outputPath string) *CSVGenerator {
    return &CSVGenerator{
        outputPath: outputPath,
        records: [][]string{
            {"Path", "Title", "Type", "Platforms", "Abstract", "Deprecated", "Beta"},
        },
    }
}

// ProcessDocument adds document information to CSV
func (cg *CSVGenerator) ProcessDocument(doc *DocJSONData, path string) error {
    platforms := make([]string, len(doc.Metadata.Platforms))
    deprecated := "false"
    beta := "false"
    
    for i, platform := range doc.Metadata.Platforms {
        platforms[i] = platform.Name
        if platform.Deprecated {
            deprecated = "true"
        }
        if platform.Beta {
            beta = "true"
        }
    }
    
    abstract := ""
    if len(doc.Abstract) > 0 {
        abstract = doc.Abstract[0].Text
    }
    
    record := []string{
        path,
        doc.Metadata.Title,
        doc.Metadata.Role,
        strings.Join(platforms, ";"),
        abstract,
        deprecated,
        beta,
    }
    
    cg.records = append(cg.records, record)
    return nil
}

// WriteCSV writes the collected records to a CSV file
func (cg *CSVGenerator) WriteCSV() error {
    if err := os.MkdirAll(filepath.Dir(cg.outputPath), 0755); err != nil {
        return fmt.Errorf("failed to create output directory: %v", err)
    }
    
    file, err := os.Create(cg.outputPath)
    if err != nil {
        return fmt.Errorf("failed to create CSV file: %v", err)
    }
    defer file.Close()
    
    writer := csv.NewWriter(file)
    defer writer.Flush()
    
    return writer.WriteAll(cg.records)
}

// GetProcessorName returns the processor name
func (cg *CSVGenerator) GetProcessorName() string {
    return "CSVGenerator"
}

// Usage example
func main() {
    generator := NewCSVGenerator("output/documentation_report.csv")
    
    // Process documents (integration point)
    // ... document processing logic ...
    
    if err := generator.WriteCSV(); err != nil {
        fmt.Printf("Failed to write CSV: %v\n", err)
        return
    }
    
    fmt.Printf("CSV report generated: %s\n", generator.outputPath)
}
```

## Integration Patterns

### Microservice Integration

```go
package main

import (
    "context"
    "encoding/json"
    "fmt"
    "net/http"
    "time"
)

// DocumentationService provides microservice interface to appledocs
type DocumentationService struct {
    mirror    *DocumentationMirror
    validator *ValidationService
    port      string
}

// NewDocumentationService creates a new documentation service
func NewDocumentationService(config Config, port string) *DocumentationService {
    return &DocumentationService{
        mirror:    NewDocumentationMirror(config),
        validator: NewValidationService(config.CacheDir, true),
        port:      port,
    }
}

// ServiceStatus represents the service status
type ServiceStatus struct {
    Status    string    `json:"status"`
    LastSync  time.Time `json:"last_sync"`
    CacheSize int64     `json:"cache_size_bytes"`
    Valid     bool      `json:"cache_valid"`
}

// Start starts the HTTP service
func (ds *DocumentationService) Start() error {
    http.HandleFunc("/health", ds.healthHandler)
    http.HandleFunc("/sync", ds.syncHandler)
    http.HandleFunc("/validate", ds.validateHandler)
    http.HandleFunc("/status", ds.statusHandler)
    
    fmt.Printf("Documentation service starting on port %s\n", ds.port)
    return http.ListenAndServe(":"+ds.port, nil)
}

func (ds *DocumentationService) healthHandler(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    
    report, err := ds.validator.ValidateCache()
    if err != nil {
        http.Error(w, fmt.Sprintf("Validation failed: %v", err), http.StatusInternalServerError)
        return
    }
    
    status := ServiceStatus{
        Status:   "healthy",
        LastSync: time.Now(),
        Valid:    report.Valid,
    }
    
    if !report.Valid {
        status.Status = "degraded"
    }
    
    json.NewEncoder(w).Encode(status)
}

func (ds *DocumentationService) syncHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodPost {
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
        return
    }
    
    framework := r.URL.Query().Get("framework")
    if framework == "" {
        http.Error(w, "Framework parameter required", http.StatusBadRequest)
        return
    }
    
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Minute)
    defer cancel()
    
    err := ds.mirror.MirrorFramework(ctx, framework)
    if err != nil {
        http.Error(w, fmt.Sprintf("Sync failed: %v", err), http.StatusInternalServerError)
        return
    }
    
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(map[string]string{
        "status": "success",
        "framework": framework,
    })
}

func (ds *DocumentationService) validateHandler(w http.ResponseWriter, r *http.Request) {
    report, err := ds.validator.ValidateCache()
    if err != nil {
        http.Error(w, fmt.Sprintf("Validation failed: %v", err), http.StatusInternalServerError)
        return
    }
    
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(report)
}

func (ds *DocumentationService) statusHandler(w http.ResponseWriter, r *http.Request) {
    // Implement status reporting
    status := ServiceStatus{
        Status:   "running",
        LastSync: time.Now(),
        Valid:    true,
    }
    
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(status)
}

// Usage
func main() {
    config := DefaultConfig()
    service := NewDocumentationService(config, "8080")
    
    if err := service.Start(); err != nil {
        fmt.Printf("Service failed: %v\n", err)
    }
}
```

### Event-Driven Integration

```go
package main

import (
    "context"
    "encoding/json"
    "fmt"
    "log"
    "time"
)

// Event types
const (
    EventDocumentProcessed = "document.processed"
    EventValidationFailed  = "validation.failed"
    EventSyncCompleted     = "sync.completed"
)

// Event represents a system event
type Event struct {
    Type      string                 `json:"type"`
    Timestamp time.Time              `json:"timestamp"`
    Source    string                 `json:"source"`
    Data      map[string]interface{} `json:"data"`
}

// EventHandler processes events
type EventHandler interface {
    HandleEvent(event Event) error
}

// EventBus manages event distribution
type EventBus struct {
    handlers map[string][]EventHandler
}

// NewEventBus creates a new event bus
func NewEventBus() *EventBus {
    return &EventBus{
        handlers: make(map[string][]EventHandler),
    }
}

// Subscribe adds an event handler
func (eb *EventBus) Subscribe(eventType string, handler EventHandler) {
    eb.handlers[eventType] = append(eb.handlers[eventType], handler)
}

// Publish sends an event to all subscribers
func (eb *EventBus) Publish(event Event) {
    for _, handler := range eb.handlers[event.Type] {
        go func(h EventHandler, e Event) {
            if err := h.HandleEvent(e); err != nil {
                log.Printf("Event handler error: %v", err)
            }
        }(handler, event)
    }
}

// DocumentationEventProducer produces events during documentation processing
type DocumentationEventProducer struct {
    eventBus *EventBus
    mirror   *DocumentationMirror
}

// NewDocumentationEventProducer creates a new event producer
func NewDocumentationEventProducer(eventBus *EventBus, mirror *DocumentationMirror) *DocumentationEventProducer {
    return &DocumentationEventProducer{
        eventBus: eventBus,
        mirror:   mirror,
    }
}

// ProcessWithEvents processes documents and emits events
func (dep *DocumentationEventProducer) ProcessWithEvents(ctx context.Context, framework string) error {
    err := dep.mirror.MirrorFramework(ctx, framework)
    
    event := Event{
        Type:      EventDocumentProcessed,
        Timestamp: time.Now(),
        Source:    "documentation-mirror",
        Data: map[string]interface{}{
            "framework": framework,
            "success":   err == nil,
        },
    }
    
    if err != nil {
        event.Type = EventValidationFailed
        event.Data["error"] = err.Error()
    }
    
    dep.eventBus.Publish(event)
    return err
}

// MetricsHandler collects metrics from events
type MetricsHandler struct {
    metrics map[string]int
}

// NewMetricsHandler creates a new metrics handler
func NewMetricsHandler() *MetricsHandler {
    return &MetricsHandler{
        metrics: make(map[string]int),
    }
}

// HandleEvent processes events for metrics collection
func (mh *MetricsHandler) HandleEvent(event Event) error {
    mh.metrics[event.Type]++
    
    log.Printf("Metrics updated: %s = %d", event.Type, mh.metrics[event.Type])
    return nil
}

// GetMetrics returns current metrics
func (mh *MetricsHandler) GetMetrics() map[string]int {
    return mh.metrics
}

// NotificationHandler sends notifications for important events
type NotificationHandler struct {
    webhookURL string
}

// NewNotificationHandler creates a new notification handler
func NewNotificationHandler(webhookURL string) *NotificationHandler {
    return &NotificationHandler{
        webhookURL: webhookURL,
    }
}

// HandleEvent sends notifications for critical events
func (nh *NotificationHandler) HandleEvent(event Event) error {
    if event.Type == EventValidationFailed {
        return nh.sendNotification(fmt.Sprintf("Validation failed: %v", event.Data["error"]))
    }
    return nil
}

func (nh *NotificationHandler) sendNotification(message string) error {
    // Implement notification logic (webhook, email, etc.)
    log.Printf("NOTIFICATION: %s", message)
    return nil
}

// Usage example
func main() {
    eventBus := NewEventBus()
    
    // Subscribe handlers
    metricsHandler := NewMetricsHandler()
    notificationHandler := NewNotificationHandler("https://hooks.slack.com/...")
    
    eventBus.Subscribe(EventDocumentProcessed, metricsHandler)
    eventBus.Subscribe(EventValidationFailed, metricsHandler)
    eventBus.Subscribe(EventValidationFailed, notificationHandler)
    
    // Create documentation processor with events
    config := DefaultConfig()
    mirror := NewDocumentationMirror(config)
    processor := NewDocumentationEventProducer(eventBus, mirror)
    
    // Process frameworks
    frameworks := []string{"SwiftUI", "UIKit", "Foundation"}
    
    for _, framework := range frameworks {
        ctx := context.Background()
        if err := processor.ProcessWithEvents(ctx, framework); err != nil {
            log.Printf("Processing failed for %s: %v", framework, err)
        }
    }
    
    // Wait for async event processing
    time.Sleep(2 * time.Second)
    
    // Display metrics
    metrics := metricsHandler.GetMetrics()
    fmt.Printf("Final metrics: %+v\n", metrics)
}
```

## Performance Tuning

### Memory-Optimized Processing

```go
package main

import (
    "context"
    "runtime"
    "time"
)

// MemoryOptimizedProcessor handles large documentation sets efficiently
type MemoryOptimizedProcessor struct {
    config       Config
    maxMemoryMB  int
    batchSize    int
    gcInterval   time.Duration
}

// NewMemoryOptimizedProcessor creates a memory-optimized processor
func NewMemoryOptimizedProcessor(config Config) *MemoryOptimizedProcessor {
    return &MemoryOptimizedProcessor{
        config:      config,
        maxMemoryMB: 512, // 512MB memory limit
        batchSize:   100,  // Process 100 documents per batch
        gcInterval:  30 * time.Second,
    }
}

// ProcessInBatches processes documents in memory-efficient batches
func (mop *MemoryOptimizedProcessor) ProcessInBatches(ctx context.Context, documents []string) error {
    // Start memory monitoring
    go mop.memoryMonitor(ctx)
    
    for i := 0; i < len(documents); i += mop.batchSize {
        end := i + mop.batchSize
        if end > len(documents) {
            end = len(documents)
        }
        
        batch := documents[i:end]
        if err := mop.processBatch(ctx, batch); err != nil {
            return fmt.Errorf("batch processing failed: %v", err)
        }
        
        // Force garbage collection between batches
        runtime.GC()
        
        // Check memory usage
        if mop.getMemoryUsageMB() > mop.maxMemoryMB {
            time.Sleep(1 * time.Second) // Brief pause to allow GC
        }
        
        select {
        case <-ctx.Done():
            return ctx.Err()
        default:
        }
    }
    
    return nil
}

func (mop *MemoryOptimizedProcessor) processBatch(ctx context.Context, batch []string) error {
    for _, doc := range batch {
        // Process individual document
        // Implementation would use appledocs functions
        fmt.Printf("Processing: %s\n", doc)
    }
    return nil
}

func (mop *MemoryOptimizedProcessor) memoryMonitor(ctx context.Context) {
    ticker := time.NewTicker(mop.gcInterval)
    defer ticker.Stop()
    
    for {
        select {
        case <-ctx.Done():
            return
        case <-ticker.C:
            memUsage := mop.getMemoryUsageMB()
            if memUsage > mop.maxMemoryMB*80/100 { // 80% threshold
                runtime.GC()
                fmt.Printf("Memory usage high (%d MB), triggered GC\n", memUsage)
            }
        }
    }
}

func (mop *MemoryOptimizedProcessor) getMemoryUsageMB() int {
    var m runtime.MemStats
    runtime.ReadMemStats(&m)
    return int(m.Alloc / 1024 / 1024)
}

// Usage
func main() {
    config := DefaultConfig()
    processor := NewMemoryOptimizedProcessor(config)
    
    // Simulate large document set
    documents := make([]string, 1000)
    for i := range documents {
        documents[i] = fmt.Sprintf("document_%d.json", i)
    }
    
    ctx := context.Background()
    if err := processor.ProcessInBatches(ctx, documents); err != nil {
        fmt.Printf("Processing failed: %v\n", err)
    }
}
```

### Concurrent Processing Patterns

```go
package main

import (
    "context"
    "fmt"
    "sync"
    "time"
)

// ConcurrentProcessor manages concurrent document processing
type ConcurrentProcessor struct {
    workers    int
    queueSize  int
    rateLimit  time.Duration
    errorLimit int
}

// NewConcurrentProcessor creates a concurrent processor
func NewConcurrentProcessor(workers, queueSize int, rateLimit time.Duration, errorLimit int) *ConcurrentProcessor {
    return &ConcurrentProcessor{
        workers:    workers,
        queueSize:  queueSize,
        rateLimit:  rateLimit,
        errorLimit: errorLimit,
    }
}

// ProcessConcurrently processes documents with controlled concurrency
func (cp *ConcurrentProcessor) ProcessConcurrently(ctx context.Context, documents []string) error {
    // Create work queue
    workQueue := make(chan string, cp.queueSize)
    results := make(chan ProcessResult, len(documents))
    
    // Start workers
    var wg sync.WaitGroup
    for i := 0; i < cp.workers; i++ {
        wg.Add(1)
        go cp.worker(ctx, &wg, workQueue, results, i)
    }
    
    // Send work
    go func() {
        defer close(workQueue)
        for _, doc := range documents {
            select {
            case workQueue <- doc:
            case <-ctx.Done():
                return
            }
        }
    }()
    
    // Collect results
    go func() {
        wg.Wait()
        close(results)
    }()
    
    // Process results
    var successCount, errorCount int
    for result := range results {
        if result.Error != nil {
            errorCount++
            fmt.Printf("Error processing %s: %v\n", result.Document, result.Error)
            
            if errorCount > cp.errorLimit {
                return fmt.Errorf("error limit exceeded (%d)", cp.errorLimit)
            }
        } else {
            successCount++
        }
    }
    
    fmt.Printf("Processing complete: %d success, %d errors\n", successCount, errorCount)
    return nil
}

// ProcessResult holds processing results
type ProcessResult struct {
    Document string
    Error    error
    Duration time.Duration
}

func (cp *ConcurrentProcessor) worker(ctx context.Context, wg *sync.WaitGroup, work <-chan string, results chan<- ProcessResult, workerID int) {
    defer wg.Done()
    
    for {
        select {
        case doc, ok := <-work:
            if !ok {
                return
            }
            
            start := time.Now()
            err := cp.processDocument(ctx, doc)
            duration := time.Since(start)
            
            results <- ProcessResult{
                Document: doc,
                Error:    err,
                Duration: duration,
            }
            
            // Rate limiting
            if cp.rateLimit > 0 {
                time.Sleep(cp.rateLimit)
            }
            
        case <-ctx.Done():
            return
        }
    }
}

func (cp *ConcurrentProcessor) processDocument(ctx context.Context, document string) error {
    // Simulate document processing
    time.Sleep(100 * time.Millisecond)
    
    // Simulate occasional errors
    if document == "error_document.json" {
        return fmt.Errorf("simulated processing error")
    }
    
    return nil
}

// Usage
func main() {
    processor := NewConcurrentProcessor(
        5,                    // 5 workers
        100,                  // queue size
        50*time.Millisecond, // rate limit
        10,                   // error limit
    )
    
    documents := []string{
        "doc1.json", "doc2.json", "doc3.json",
        "error_document.json", "doc4.json",
    }
    
    ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
    defer cancel()
    
    if err := processor.ProcessConcurrently(ctx, documents); err != nil {
        fmt.Printf("Processing failed: %v\n", err)
    }
}
```

## Data Pipeline Integration

### Apache Kafka Integration

```go
package main

import (
    "context"
    "encoding/json"
    "fmt"
    "log"
    "time"
)

// DocumentationMessage represents a Kafka message for documentation events
type DocumentationMessage struct {
    EventType   string                 `json:"event_type"`
    Timestamp   time.Time              `json:"timestamp"`
    Framework   string                 `json:"framework,omitempty"`
    Document    string                 `json:"document,omitempty"`
    Metadata    map[string]interface{} `json:"metadata,omitempty"`
    Error       string                 `json:"error,omitempty"`
}

// KafkaIntegration handles Kafka integration for documentation events
type KafkaIntegration struct {
    topicName string
    // In real implementation, include Kafka producer/consumer
}

// NewKafkaIntegration creates a new Kafka integration
func NewKafkaIntegration(topicName string) *KafkaIntegration {
    return &KafkaIntegration{
        topicName: topicName,
    }
}

// PublishDocumentEvent publishes a documentation event to Kafka
func (ki *KafkaIntegration) PublishDocumentEvent(eventType, framework, document string, metadata map[string]interface{}) error {
    message := DocumentationMessage{
        EventType: eventType,
        Timestamp: time.Now(),
        Framework: framework,
        Document:  document,
        Metadata:  metadata,
    }
    
    messageBytes, err := json.Marshal(message)
    if err != nil {
        return fmt.Errorf("failed to marshal message: %v", err)
    }
    
    // In real implementation, publish to Kafka
    log.Printf("Publishing to Kafka topic %s: %s", ki.topicName, string(messageBytes))
    
    return nil
}

// ConsumeDocumentEvents consumes documentation events from Kafka
func (ki *KafkaIntegration) ConsumeDocumentEvents(ctx context.Context, handler func(DocumentationMessage) error) error {
    // In real implementation, consume from Kafka
    // Simulated for demonstration
    for {
        select {
        case <-ctx.Done():
            return ctx.Err()
        default:
            // Simulate receiving a message
            message := DocumentationMessage{
                EventType: "document.processed",
                Timestamp: time.Now(),
                Framework: "SwiftUI",
                Document:  "View.json",
            }
            
            if err := handler(message); err != nil {
                log.Printf("Error handling message: %v", err)
            }
            
            time.Sleep(1 * time.Second) // Simulate delay
        }
    }
}

// DocumentationStreamProcessor processes documentation events from Kafka
type DocumentationStreamProcessor struct {
    kafka     *KafkaIntegration
    processor *DocumentationMirror
}

// NewDocumentationStreamProcessor creates a stream processor
func NewDocumentationStreamProcessor(kafka *KafkaIntegration, processor *DocumentationMirror) *DocumentationStreamProcessor {
    return &DocumentationStreamProcessor{
        kafka:     kafka,
        processor: processor,
    }
}

// ProcessStream processes the documentation event stream
func (dsp *DocumentationStreamProcessor) ProcessStream(ctx context.Context) error {
    return dsp.kafka.ConsumeDocumentEvents(ctx, dsp.handleMessage)
}

func (dsp *DocumentationStreamProcessor) handleMessage(message DocumentationMessage) error {
    switch message.EventType {
    case "framework.requested":
        return dsp.handleFrameworkRequest(message)
    case "validation.requested":
        return dsp.handleValidationRequest(message)
    default:
        log.Printf("Unknown event type: %s", message.EventType)
    }
    return nil
}

func (dsp *DocumentationStreamProcessor) handleFrameworkRequest(message DocumentationMessage) error {
    ctx := context.Background()
    err := dsp.processor.MirrorFramework(ctx, message.Framework)
    
    // Publish result
    eventType := "framework.completed"
    if err != nil {
        eventType = "framework.failed"
    }
    
    metadata := map[string]interface{}{
        "success": err == nil,
    }
    
    return dsp.kafka.PublishDocumentEvent(eventType, message.Framework, "", metadata)
}

func (dsp *DocumentationStreamProcessor) handleValidationRequest(message DocumentationMessage) error {
    // Implement validation handling
    return nil
}

// Usage
func main() {
    kafka := NewKafkaIntegration("documentation-events")
    
    config := DefaultConfig()
    mirror := NewDocumentationMirror(config)
    
    processor := NewDocumentationStreamProcessor(kafka, mirror)
    
    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Minute)
    defer cancel()
    
    // Start stream processing
    if err := processor.ProcessStream(ctx); err != nil {
        log.Printf("Stream processing failed: %v", err)
    }
}
```

### Database Integration

```go
package main

import (
    "database/sql"
    "encoding/json"
    "fmt"
    "time"
    
    _ "github.com/lib/pq" // PostgreSQL driver
)

// DocumentationDB handles database operations for documentation
type DocumentationDB struct {
    db *sql.DB
}

// NewDocumentationDB creates a new database connection
func NewDocumentationDB(connectionString string) (*DocumentationDB, error) {
    db, err := sql.Open("postgres", connectionString)
    if err != nil {
        return nil, fmt.Errorf("failed to connect to database: %v", err)
    }
    
    if err := db.Ping(); err != nil {
        return nil, fmt.Errorf("failed to ping database: %v", err)
    }
    
    return &DocumentationDB{db: db}, nil
}

// CreateTables creates necessary database tables
func (ddb *DocumentationDB) CreateTables() error {
    queries := []string{
        `CREATE TABLE IF NOT EXISTS documents (
            id SERIAL PRIMARY KEY,
            path VARCHAR(500) UNIQUE NOT NULL,
            title VARCHAR(200),
            framework VARCHAR(100),
            doc_type VARCHAR(50),
            platforms TEXT[],
            content JSONB,
            created_at TIMESTAMP DEFAULT NOW(),
            updated_at TIMESTAMP DEFAULT NOW()
        )`,
        `CREATE TABLE IF NOT EXISTS processing_metrics (
            id SERIAL PRIMARY KEY,
            run_id UUID NOT NULL,
            start_time TIMESTAMP NOT NULL,
            end_time TIMESTAMP,
            processed_count INTEGER,
            error_count INTEGER,
            cache_hits INTEGER,
            cache_misses INTEGER,
            metrics JSONB,
            created_at TIMESTAMP DEFAULT NOW()
        )`,
        `CREATE INDEX IF NOT EXISTS idx_documents_framework ON documents(framework)`,
        `CREATE INDEX IF NOT EXISTS idx_documents_type ON documents(doc_type)`,
        `CREATE INDEX IF NOT EXISTS idx_documents_updated ON documents(updated_at)`,
    }
    
    for _, query := range queries {
        if _, err := ddb.db.Exec(query); err != nil {
            return fmt.Errorf("failed to execute query: %v", err)
        }
    }
    
    return nil
}

// DocumentRecord represents a document in the database
type DocumentRecord struct {
    ID        int             `json:"id"`
    Path      string          `json:"path"`
    Title     string          `json:"title"`
    Framework string          `json:"framework"`
    DocType   string          `json:"doc_type"`
    Platforms []string        `json:"platforms"`
    Content   json.RawMessage `json:"content"`
    CreatedAt time.Time       `json:"created_at"`
    UpdatedAt time.Time       `json:"updated_at"`
}

// InsertDocument inserts or updates a document
func (ddb *DocumentationDB) InsertDocument(doc *DocJSONData, path string) error {
    platforms := make([]string, len(doc.Metadata.Platforms))
    for i, p := range doc.Metadata.Platforms {
        platforms[i] = p.Name
    }
    
    contentJSON, err := json.Marshal(doc)
    if err != nil {
        return fmt.Errorf("failed to marshal document content: %v", err)
    }
    
    query := `
        INSERT INTO documents (path, title, framework, doc_type, platforms, content, updated_at)
        VALUES ($1, $2, $3, $4, $5, $6, NOW())
        ON CONFLICT (path) DO UPDATE SET
            title = EXCLUDED.title,
            framework = EXCLUDED.framework,
            doc_type = EXCLUDED.doc_type,
            platforms = EXCLUDED.platforms,
            content = EXCLUDED.content,
            updated_at = NOW()
    `
    
    framework := extractFrameworkFromPath(path)
    
    _, err = ddb.db.Exec(query, path, doc.Metadata.Title, framework, doc.Metadata.Role, platforms, contentJSON)
    return err
}

// GetDocumentsByFramework retrieves documents for a specific framework
func (ddb *DocumentationDB) GetDocumentsByFramework(framework string) ([]DocumentRecord, error) {
    query := `
        SELECT id, path, title, framework, doc_type, platforms, content, created_at, updated_at
        FROM documents
        WHERE framework = $1
        ORDER BY title
    `
    
    rows, err := ddb.db.Query(query, framework)
    if err != nil {
        return nil, err
    }
    defer rows.Close()
    
    var documents []DocumentRecord
    for rows.Next() {
        var doc DocumentRecord
        var platforms string
        
        err := rows.Scan(
            &doc.ID, &doc.Path, &doc.Title, &doc.Framework,
            &doc.DocType, &platforms, &doc.Content,
            &doc.CreatedAt, &doc.UpdatedAt,
        )
        if err != nil {
            return nil, err
        }
        
        // Parse platforms array
        if err := json.Unmarshal([]byte(platforms), &doc.Platforms); err != nil {
            return nil, err
        }
        
        documents = append(documents, doc)
    }
    
    return documents, rows.Err()
}

// InsertMetrics inserts processing metrics
func (ddb *DocumentationDB) InsertMetrics(runID string, metrics *MetricsSnapshot) error {
    metricsJSON, err := json.Marshal(metrics)
    if err != nil {
        return fmt.Errorf("failed to marshal metrics: %v", err)
    }
    
    query := `
        INSERT INTO processing_metrics (
            run_id, start_time, end_time, processed_count,
            error_count, cache_hits, cache_misses, metrics
        ) VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
    `
    
    endTime := metrics.StartTime.Add(metrics.RuntimeDuration)
    
    _, err = ddb.db.Exec(query,
        runID, metrics.StartTime, endTime, metrics.Processed,
        metrics.Errors, metrics.CacheHits, metrics.CacheMisses, metricsJSON,
    )
    
    return err
}

func extractFrameworkFromPath(path string) string {
    // Extract framework name from path
    // Implementation depends on path structure
    return "Unknown"
}

// Usage
func main() {
    db, err := NewDocumentationDB("postgres://user:password@localhost/appledocs?sslmode=disable")
    if err != nil {
        log.Printf("Database connection failed: %v", err)
        return
    }
    defer db.db.Close()
    
    if err := db.CreateTables(); err != nil {
        log.Printf("Table creation failed: %v", err)
        return
    }
    
    // Example usage
    docs, err := db.GetDocumentsByFramework("SwiftUI")
    if err != nil {
        log.Printf("Query failed: %v", err)
        return
    }
    
    fmt.Printf("Found %d SwiftUI documents\n", len(docs))
}
```

## Testing and Quality Assurance

### Integration Test Framework

```go
package main

import (
    "context"
    "os"
    "path/filepath"
    "testing"
    "time"
)

// TestEnvironment sets up isolated test environment
type TestEnvironment struct {
    tempDir   string
    cacheDir  string
    outputDir string
    config    Config
}

// NewTestEnvironment creates a new test environment
func NewTestEnvironment(t *testing.T) *TestEnvironment {
    tempDir := t.TempDir()
    
    return &TestEnvironment{
        tempDir:   tempDir,
        cacheDir:  filepath.Join(tempDir, "cache"),
        outputDir: filepath.Join(tempDir, "output"),
        config: Config{
            BaseURL:     "https://developer.apple.com",
            CacheDir:    filepath.Join(tempDir, "cache"),
            OutputDir:   filepath.Join(tempDir, "output"),
            Timeout:     10 * time.Second,
            RateLimit:   1.0,
            Concurrency: 1,
        },
    }
}

// Setup prepares the test environment
func (te *TestEnvironment) Setup(t *testing.T) error {
    dirs := []string{te.cacheDir, te.outputDir}
    for _, dir := range dirs {
        if err := os.MkdirAll(dir, 0755); err != nil {
            return err
        }
    }
    return nil
}

// TestDocumentationMirror tests the documentation mirror functionality
func TestDocumentationMirror(t *testing.T) {
    env := NewTestEnvironment(t)
    if err := env.Setup(t); err != nil {
        t.Fatalf("Setup failed: %v", err)
    }
    
    mirror := NewDocumentationMirror(env.config)
    
    ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
    defer cancel()
    
    // Test framework mirroring
    t.Run("MirrorFramework", func(t *testing.T) {
        err := mirror.MirrorFramework(ctx, "SwiftUI")
        if err != nil {
            t.Errorf("MirrorFramework failed: %v", err)
        }
        
        // Verify output exists
        expectedPath := filepath.Join(env.outputDir, "tutorials", "data", "documentation", "SwiftUI.json")
        if _, err := os.Stat(expectedPath); os.IsNotExist(err) {
            t.Errorf("Expected output file not found: %s", expectedPath)
        }
    })
}

// TestValidation tests the validation functionality
func TestValidation(t *testing.T) {
    env := NewTestEnvironment(t)
    if err := env.Setup(t); err != nil {
        t.Fatalf("Setup failed: %v", err)
    }
    
    validator := NewValidationService(env.cacheDir, true)
    
    t.Run("ValidateEmptyCache", func(t *testing.T) {
        report, err := validator.ValidateCache()
        if err != nil {
            t.Errorf("ValidateCache failed: %v", err)
        }
        
        if !report.Valid {
            t.Errorf("Empty cache should be valid")
        }
    })
    
    t.Run("ValidateWithChecksums", func(t *testing.T) {
        // Create a test file
        testFile := filepath.Join(env.cacheDir, "test.json")
        testContent := []byte(`{"test": "content"}`)
        
        if err := os.WriteFile(testFile, testContent, 0644); err != nil {
            t.Fatalf("Failed to create test file: %v", err)
        }
        
        // Validate with checksums
        result := ValidateCacheIntegrityWithChecksums(env.cacheDir)
        if !result.Valid {
            t.Errorf("Validation should pass for valid cache")
        }
    })
}

// BenchmarkDocumentProcessing benchmarks document processing performance
func BenchmarkDocumentProcessing(b *testing.B) {
    env := NewTestEnvironment(&testing.T{})
    if err := env.Setup(&testing.T{}); err != nil {
        b.Fatalf("Setup failed: %v", err)
    }
    
    processor := NewMemoryOptimizedProcessor(env.config)
    
    // Create test documents
    documents := make([]string, 100)
    for i := range documents {
        documents[i] = fmt.Sprintf("test_document_%d.json", i)
    }
    
    b.ResetTimer()
    
    for i := 0; i < b.N; i++ {
        ctx := context.Background()
        if err := processor.ProcessInBatches(ctx, documents); err != nil {
            b.Errorf("ProcessInBatches failed: %v", err)
        }
    }
}

// TestConcurrentProcessing tests concurrent processing under load
func TestConcurrentProcessing(t *testing.T) {
    processor := NewConcurrentProcessor(5, 100, 10*time.Millisecond, 5)
    
    documents := make([]string, 1000)
    for i := range documents {
        documents[i] = fmt.Sprintf("document_%d.json", i)
    }
    
    ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
    defer cancel()
    
    if err := processor.ProcessConcurrently(ctx, documents); err != nil {
        t.Errorf("Concurrent processing failed: %v", err)
    }
}
```

This comprehensive developer integration guide provides detailed examples of how to extend and integrate the appledocs tool into various development workflows, from simple library usage to complex distributed systems integration.