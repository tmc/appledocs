# Apple Docs Mirror Tool - API Reference

## Overview

The `appledocs` tool is a comprehensive solution for crawling, caching, and mirroring Apple's developer documentation. It provides multiple output formats (JSON, HTML, Markdown) and includes validation, metrics collection, and a GraphQL API server.

## Table of Contents

- [Command-Line Interface](#command-line-interface)
- [Public API Types](#public-api-types)
- [Validation System](#validation-system)
- [Metrics and Monitoring](#metrics-and-monitoring)
- [GraphQL API Server](#graphql-api-server)
- [Configuration Options](#configuration-options)

## Command-Line Interface

### Core Commands

#### Basic Crawling
```bash
# Basic crawl operation
./appledocs -mode crawl

# Crawl with specific framework
./appledocs -mode crawl -entry-point SwiftUI

# Crawl specific class documentation
./appledocs -mode crawl -entry-point "SwiftUI/View"
```

#### High-Performance Operations
```bash
# High-concurrency crawling
./appledocs -mode crawl -concurrency 20 -rate-limit 50

# Fast crawling with rate limiting
./appledocs -mode crawl -concurrency 10 -rate-limit 25 -delay 100ms

# Force refresh all cached content
./appledocs -mode crawl -force -concurrency 15
```

#### Content Generation
```bash
# Generate all outputs (crawl + HTML + Markdown)
./appledocs -mode all

# Generate only HTML index
./appledocs -mode html

# Generate only Markdown documentation
./appledocs -mode markdown -md-output ./docs

# Generate specialized EndpointSecurity documentation
make endpointsecurity
```

#### URL Discovery and Analysis
```bash
# Discover all available documentation URLs
./appledocs -print-urls

# Discover URLs for specific framework
./appledocs -print-urls -entry-point SwiftUI

# Discover URLs for specific class
./appledocs -print-urls -entry-point "SecurityFoundation/SFAuthorization"
```

### Advanced Options

#### Performance Tuning
```bash
# Optimize for bandwidth-limited environments
./appledocs -mode crawl -concurrency 5 -rate-limit 5 -timeout 60s

# Optimize for fast local networks
./appledocs -mode crawl -concurrency 50 -rate-limit 100 -timeout 10s

# Skip symbol-level documentation to reduce size
./appledocs -mode crawl -skip-symbols
```

#### Validation and Monitoring
```bash
# Enable comprehensive cache validation
./appledocs -mode crawl -validate-cache -checksum-validation

# Export detailed metrics
./appledocs -mode crawl -export-metrics metrics.json

# Verbose logging for debugging
./appledocs -mode crawl -verbose -log-level debug
```

#### Directory and Output Management
```bash
# Custom directories
./appledocs -mode all \
  -output ./apple-docs \
  -cache ./cache \
  -md-output ./markdown-docs

# Exclude specific paths
./appledocs -mode crawl -exclude-paths "en-US/docs/Mozilla,legacy"
```

### Flag Reference

#### Core Operation Flags
| Flag | Type | Default | Description |
|------|------|---------|-------------|
| `-mode` | string | `crawl` | Operation mode: `crawl`, `html`, `markdown`, or `all` |
| `-base` | string | `https://developer.apple.com` | Base URL for Apple docs |
| `-entry-point` | string | `/tutorials/data/documentation/technologies.json` | Starting path for crawling |

#### Performance Flags
| Flag | Type | Default | Description |
|------|------|---------|-------------|
| `-concurrency` | int | `1` | Number of concurrent downloads |
| `-rate-limit` | float64 | `10.0` | Requests per second (0 = no limit) |
| `-delay` | duration | `0` | Delay between URL processing |
| `-timeout` | duration | `30s` | HTTP request timeout |
| `-max-time` | duration | `1h` | Maximum program runtime |

#### Output Control Flags
| Flag | Type | Default | Description |
|------|------|---------|-------------|
| `-output` | string | `output` | Directory for mirrored content |
| `-cache` | string | `.cache` | Directory for HTTP cache |
| `-md-output` | string | `markdown` | Directory for Markdown output |
| `-pretty` | bool | `true` | Pretty-print JSON files |

#### Validation Flags
| Flag | Type | Default | Description |
|------|------|---------|-------------|
| `-validate-cache` | bool | `false` | Validate cache integrity on startup |
| `-checksum-validation` | bool | `false` | Enable checksum-based validation |
| `-export-metrics` | string | `` | Export metrics to JSON file |

#### Filtering Flags
| Flag | Type | Default | Description |
|------|------|---------|-------------|
| `-skip-symbols` | bool | `false` | Skip symbol-level documentation |
| `-exclude-paths` | string | `en-US/docs/Mozilla` | Comma-separated paths to exclude |
| `-print-urls` | bool | `false` | Only print discovered URLs and exit |

#### Logging Flags
| Flag | Type | Default | Description |
|------|------|---------|-------------|
| `-verbose` | bool | `false` | Enable verbose logging |
| `-log-level` | string | `info` | Log level: `debug`, `info`, `warn`, `error` |

## Public API Types

### Core Application Types

#### `appledocs` struct
Main application state container:
```go
type appledocs struct {
    client         *http.Client
    visitedURLs    map[string]bool
    visitedMutex   sync.RWMutex
    jsonEntries    []JSONFileEntry
    entriesMutex   sync.Mutex
    processedCount int
    badURLs        map[string]bool
    urlDepths      map[string]int
    rateLimiter    *rate.Limiter
    
    // Metrics fields
    cacheHits      int
    cacheMisses    int
    errors         int
    skippedURLs    int
    skippedSymbols int
    statsMutex     sync.Mutex
}
```

#### `JSONFileEntry` struct
Represents a discovered JSON documentation file:
```go
type JSONFileEntry struct {
    Path string // Relative path to the file
    URL  string // Original URL
}
```

### Documentation Types

#### `DocJSONData` struct
Root structure for Apple documentation JSON:
```go
type DocJSONData struct {
    Abstract               []TextContent         `json:"abstract,omitempty"`
    Hierarchy              Hierarchy             `json:"hierarchy,omitempty"`
    Identifier             Identifier            `json:"identifier,omitempty"`
    Kind                   string                `json:"kind,omitempty"`
    Metadata               Metadata              `json:"metadata,omitempty"`
    PrimaryContentSections []ContentSection      `json:"primaryContentSections,omitempty"`
    References             map[string]Reference  `json:"references,omitempty"`
    TopicSections          []TopicSection        `json:"topicSections,omitempty"`
    RelationshipsSections  []RelationshipSection `json:"relationshipsSections,omitempty"`
    SeeAlsoSections        []SeeAlsoSection      `json:"seeAlsoSections,omitempty"`
}
```

#### `Metadata` struct
Document metadata:
```go
type Metadata struct {
    ExternalID     string     `json:"externalID,omitempty"`
    Fragments      []Fragment `json:"fragments,omitempty"`
    Modules        []Module   `json:"modules,omitempty"`
    Role           string     `json:"role,omitempty"`
    RoleHeading    string     `json:"roleHeading,omitempty"`
    SymbolKind     string     `json:"symbolKind,omitempty"`
    Title          string     `json:"title,omitempty"`
    Platforms      []Platform `json:"platforms,omitempty"`
    NavigatorTitle []Fragment `json:"navigatorTitle,omitempty"`
}
```

#### `Platform` struct
Platform compatibility information:
```go
type Platform struct {
    Beta         bool   `json:"beta,omitempty"`
    Deprecated   bool   `json:"deprecated,omitempty"`
    IntroducedAt string `json:"introducedAt,omitempty"`
    Name         string `json:"name,omitempty"`
    Unavailable  bool   `json:"unavailable,omitempty"`
}
```

### Tree and HTML Generation Types

#### `TreeNode` struct
File tree representation for HTML generation:
```go
type TreeNode struct {
    Name     string
    Path     string
    IsDir    bool
    Children []*TreeNode
}
```

## Validation System

### Validation Types

#### `ValidationResult` struct
```go
type ValidationResult struct {
    Valid    bool
    Errors   []ValidationError
    Warnings []ValidationError
}
```

#### `ValidationError` struct
```go
type ValidationError struct {
    Field   string
    Value   interface{}
    Message string
    Path    string
}
```

### Validation Functions

#### Command-Line Validation
```go
func ValidateCommandLineFlags() ValidationResult
```
Validates all command-line flags for correctness and consistency.

#### Data Integrity Validation
```go
func ValidateDataIntegrity(filePath string, data []byte) ValidationResult
```
Performs comprehensive validation of downloaded data including:
- File size checks
- Binary data detection
- JSON structure validation
- Apple documentation schema validation

#### Cache Validation
```go
func ValidateCache(cacheDir string) ValidationResult
```
Validates cache directory integrity including:
- Directory existence
- File accessibility
- Empty file detection
- Age-based warnings

#### Enhanced Checksum Validation
```go
func ValidateCacheIntegrityWithChecksums(cacheDir string) ValidationResult
```
Advanced validation using SHA-256 checksums for:
- Corruption detection
- Tampering verification
- Metadata management

### Checksum Management

#### `ChecksumManager` struct
```go
type ChecksumManager struct {
    metadataPath string
    metadata     *CacheMetadata
}
```

#### `FileChecksum` struct
```go
type FileChecksum struct {
    Path     string    `json:"path"`
    SHA256   string    `json:"sha256"`
    Size     int64     `json:"size"`
    ModTime  time.Time `json:"modTime"`
    Verified time.Time `json:"verified"`
}
```

## Metrics and Monitoring

### Metrics Types

#### `MetricsSnapshot` struct
Comprehensive metrics collection:
```go
type MetricsSnapshot struct {
    // Basic metrics
    Processed       int           `json:"processed"`
    CacheHits       int           `json:"cache_hits"`
    CacheMisses     int           `json:"cache_misses"`
    Errors          int           `json:"errors"`
    SkippedURLs     int           `json:"skipped_urls"`
    SkippedSymbols  int           `json:"skipped_symbols"`
    
    // Enhanced metrics
    StartTime              time.Time     `json:"start_time"`
    RuntimeDuration        time.Duration `json:"runtime_duration"`
    TotalBytesDownloaded   int64         `json:"total_bytes_downloaded"`
    TotalBytesFromCache    int64         `json:"total_bytes_from_cache"`
    AvgResponseTime        time.Duration `json:"avg_response_time"`
    RequestCount           int           `json:"request_count"`
    HTTPErrors             map[int]int   `json:"http_errors"`
    RetryCount             int           `json:"retry_count"`
    FrameworkCount         int           `json:"framework_count"`
    ClassCount             int           `json:"class_count"`
    MethodCount            int           `json:"method_count"`
    
    // Calculated metrics
    CacheHitRate           float64       `json:"cache_hit_rate"`
    DownloadRate           float64       `json:"download_rate_mbps"`
    ProcessingRate         float64       `json:"processing_rate_per_sec"`
    EstimatedTimeRemaining time.Duration `json:"estimated_time_remaining"`
}
```

### Metrics Collection Functions

#### Enhanced Metrics
```go
func (app *appledocs) getEnhancedMetrics() MetricsSnapshot
```
Returns comprehensive metrics snapshot with calculated rates and estimates.

#### Metrics Recording
```go
func (app *appledocs) recordResponseTime(duration time.Duration)
func (app *appledocs) recordBytesDownloaded(bytes int64)
func (app *appledocs) recordBytesFromCache(bytes int64)
func (app *appledocs) recordHTTPError(statusCode int)
func (app *appledocs) recordContentType(path string)
```

## GraphQL API Server

### Server Components

The GraphQL API server (`cmd/appledocs-gql`) provides programmatic access to cached documentation.

#### Document Service
```go
type DocumentService struct {
    cacheDir string
}

func NewDocumentService(cacheDir string) *DocumentService
func (s *DocumentService) GetDocumentByPath(path string) (*Document, error)
func (s *DocumentService) Search(query string, limit int) ([]*Document, error)
```

#### Document Type
```go
type Document struct {
    ID       string                 `json:"id"`
    Path     string                 `json:"path"`
    Title    string                 `json:"title,omitempty"`
    Abstract string                 `json:"abstract,omitempty"`
    Metadata map[string]interface{} `json:"metadata,omitempty"`
    Content  map[string]interface{} `json:"content,omitempty"`
}
```

### REST API Endpoints

#### Document Retrieval
```
GET /api/document?path=<path>
```
Retrieves a specific document by path.

#### Search
```
GET /api/search?q=<query>&limit=<limit>
```
Searches documents by query string.

#### Frameworks List
```
GET /api/frameworks
```
Lists all top-level frameworks.

### GraphQL Schema

```graphql
type Document {
    id: String!
    path: String!
    title: String
    abstract: String
    metadata: JSON
    content: JSON
}

type SearchResult {
    results: [Document!]!
    count: Int!
    query: String!
    limit: Int!
}

type Query {
    document(path: String!): Document
    search(query: String!, limit: Int = 10): SearchResult
    frameworks: [Document!]!
}
```

## Configuration Options

### Environment Variables

While the tool primarily uses command-line flags, certain behaviors can be influenced by environment:

- Cache location: Use `-cache` flag or ensure `.cache` directory is writable
- Output permissions: Ensure output directories have proper write permissions
- Network timeouts: Use `-timeout` flag for HTTP operations

### Best Practices Configuration

#### Development Environment
```bash
./appledocs -mode crawl \
  -concurrency 5 \
  -rate-limit 10 \
  -verbose \
  -log-level debug \
  -validate-cache
```

#### Production Environment
```bash
./appledocs -mode crawl \
  -concurrency 20 \
  -rate-limit 50 \
  -timeout 30s \
  -max-time 2h \
  -export-metrics metrics.json \
  -checksum-validation
```

#### Large-Scale Deployment
```bash
./appledocs -mode crawl \
  -concurrency 50 \
  -rate-limit 100 \
  -skip-symbols \
  -exclude-paths "legacy,deprecated" \
  -export-metrics /var/log/appledocs-metrics.json
```

### Output Directory Structure

```
output/
├── index.html                    # Generated HTML index
├── tutorials/
│   └── data/
│       └── documentation/
│           ├── technologies.json # Root technologies file
│           ├── SwiftUI.json     # Framework documentation
│           ├── SwiftUI/         # Framework classes/symbols
│           └── ...
└── design/                      # Design guidelines (if crawled)

markdown/
├── index.md                     # Markdown index
├── tutorials/
│   └── data/
│       └── documentation/
│           ├── SwiftUI.md       # Framework documentation
│           ├── EndpointSecurity.md
│           └── ...
└── EndpointSecurity.md         # Specialized reference

.cache/
├── developer.apple.com/         # Cached HTTP responses
│   └── tutorials/
│       └── data/
└── known-bad-urls.txt          # URLs to skip
```

## Error Handling

### Common Error Types

1. **Network Errors**: Handled with exponential backoff retry
2. **Rate Limiting**: Managed by built-in rate limiter
3. **Validation Errors**: Reported but don't stop processing
4. **File System Errors**: Atomic writes prevent corruption
5. **JSON Parsing Errors**: Gracefully handled with detailed logging

### Error Recovery

- Bad URLs are tracked and skipped on subsequent runs
- Cache corruption is detected and files are re-downloaded
- Network timeouts trigger automatic retries
- Graceful shutdown preserves partial progress

## Integration Patterns

### CI/CD Integration

```bash
# Daily documentation update
./appledocs -mode crawl -validate-cache -export-metrics daily-metrics.json

# Generate documentation artifacts
./appledocs -mode all -md-output ./docs

# Validate and commit changes
git add output/ markdown/
git commit -m "Update Apple documentation $(date +%Y-%m-%d)"
```

### Monitoring Integration

```bash
# Export metrics for monitoring systems
./appledocs -mode crawl -export-metrics /tmp/metrics.json

# Parse metrics for alerting
jq '.errors' /tmp/metrics.json | xargs -I {} echo "Errors: {}"
```

This API reference provides comprehensive documentation for all public interfaces, configuration options, and usage patterns of the appledocs tool.