# AppleDocs Roadmap

**Status:** Ready for v1.0
**Last Updated:** October 7, 2025
**Approach:** Crawler-first distribution

## TL;DR

**Don't distribute data. Distribute the tool that generates the data.**

## Table of Contents

1. [Executive Summary](#executive-summary)
2. [Current State](#current-state)
3. [The Decision](#the-decision)
4. [Why Crawler-First](#why-crawler-first)
5. [Implementation Plan](#implementation-plan)
6. [Distribution Strategy](#distribution-strategy)
7. [Design Philosophy](#design-philosophy)
8. [Future Enhancements](#future-enhancements)
9. [What We're Not Building](#what-were-not-building)

## Executive Summary

AppleDocs provides programmatic access to Apple's developer documentation from Go. After thorough analysis, we chose a **crawler-first approach**: distribute the tool that generates documentation, not pre-packaged data.

### Core Package (650 lines)
- `types.go` - Type definitions (121 lines)
- `fs.go` - Filesystem and query functions (272 lines)
- `maps.go` - Map-based helpers (200 lines)
- `doc.go` - Package documentation (86 lines)

### Crawler Tool
- `cmd/appledocs/main.go` - Web crawler (1,962 lines)
- Downloads from Apple's servers
- Converts to Markdown/HTML
- Already 90% complete

## Current State

### ✅ Complete

**Core Library:**
- Two complementary APIs (typed + map-based)
- fs.FS-based design (works with any filesystem)
- 6 example programs
- 22 tests (all passing)

**Crawler:**
- Downloads from developer.apple.com
- Markdown/HTML conversion
- Rate limiting and concurrency
- Error handling

**Data Characteristics:**
- 314MB uncompressed → 15MB compressed (zstd, 95.2% reduction)
- 63,449 files
- 376 frameworks
- 1,284+ symbols in Foundation alone

**Documentation:**
- Comprehensive README
- Distribution research (1,917 lines)
- Strategic decision documentation
- 6 working examples

### 🎯 Ready For

**v1.0 Release:**
- Crawler exists and works
- Path is clear
- Just needs polish

## The Decision

### Three Approaches Considered

#### ❌ Approach 1: GitHub Releases MVP
```go
// 272-line fetch package
fsys, _ := fetch.Open("17.0.0")  // Downloads 15MB tarball
```

**Why rejected:**
- Pre-crawled data gets stale
- 15MB download every version
- Doesn't serve code generators well
- Adds complexity

#### ❌ Approach 2: OCI Distribution
```go
// 800+ line OCI package
fsys, _ := oci.Pull(ctx, "ghcr.io/tmc/appledocs:v17")
```

**Why rejected:**
- Massive complexity (OCI layer management)
- 120+ hours implementation
- Solving wrong problem (dedup for ~1x/year updates)
- Still distributes stale data

#### ✅ Approach 3: Crawler Pattern (Winner)
```bash
# Install the tool
go install github.com/tmc/appledocs/cmd/appledocs@latest

# Generate fresh data
appledocs crawl --frameworks Foundation,UIKit --format json
```

**Why chosen:**
- ✅ Already 90% implemented
- ✅ Always current (from Apple's servers)
- ✅ Selective (choose frameworks)
- ✅ Flexible (JSON or Markdown)
- ✅ Zero distribution complexity
- ✅ Matches ecosystem pattern (protoc, sqlc, DarwinKit)

## Why Crawler-First

### Technical Reasons

1. **Solves actual problem:** Code generators need fresh, filtered data
2. **Leverages existing work:** 90% done already
3. **Future-proof:** Works with any Apple docs version
4. **Flexible:** Users control what/when/how to crawl

### Philosophical Reasons

1. **Radical simplicity:** Tool does one thing well
2. **User agency:** Users generate their own data
3. **Transparency:** Clear what's happening (crawling Apple)
4. **Composability:** Works with standard Unix tools

### Ecosystem Alignment

- **Like protoc:** Generate code from schemas
- **Like sqlc:** Generate code from SQL
- **Like DarwinKit:** Generate bindings from docs
- **Like go generate:** Tool-driven code generation

### The "Russ Cox Test"

**Can you explain it in one sentence?**

❌ "Uses OCI artifacts with content-addressed layer deduplication..."
❌ "Downloads from GitHub Releases with zstd compression..."
✅ **"Crawl Apple's docs, save locally, read with Go"**

## Implementation Plan

### Phase 1: Metadata Encoding & CLI Polish (Week 1)

**Goal:** Encode crawled docs into Go API for efficient metadata access, then polish crawler

**Part A: Metadata Encoding (Days 1-3)**

After crawling, we must encode the documentation metadata into an efficient Go API structure. This allows fast queries without parsing 63,449 JSON files.

**Tasks:**
- [ ] Design metadata index structure
  - Framework → Classes/Protocols mapping
  - Symbol name → File path lookup
  - Platform availability index
  - Symbol kind categorization
- [ ] Generate index from crawled data
  - Parse all JSON files once
  - Extract metadata (title, kind, externalID, platforms)
  - Build in-memory structures
- [ ] Encode index into Go
  - Generate `index.go` with constants/maps
  - Or: Generate SQLite database for queries
  - Or: Generate protobuf/msgpack for fast loading
- [ ] Add query functions to API
  - `ListFrameworks()` - Fast framework enumeration
  - `ListSymbols(framework)` - Symbol listing without file I/O
  - `SearchSymbols(pattern)` - Search by name
  - `GetSymbolPath(framework, symbol)` - Resolve to file path

**Rationale:**
- **Problem:** Opening 63,449 files to list frameworks is slow
- **Solution:** Pre-index metadata during crawl, encode into efficient structure
- **Benefit:** Sub-millisecond queries for common operations

**Example Generated Code:**
```go
// index.go - Generated during crawl
package appledocs

var Frameworks = []string{
    "ARKit", "AVFAudio", "AVFoundation", // ... 376 total
}

var SymbolIndex = map[string]string{
    "Foundation/NSString": "Foundation/NSString.json",
    "Foundation/NSArray": "Foundation/NSArray.json",
    // ... 1000s of symbols
}

var FrameworkSymbols = map[string][]string{
    "Foundation": {"NSString", "NSArray", "NSData", /* ... */},
    "UIKit": {"UIView", "UIViewController", /* ... */},
}
```

**Part B: CLI Polish (Days 4-5)**

**Tasks:**
- [ ] Improve CLI flags (use cobra or better flag handling)
- [ ] Add progress indicators (download/conversion progress)
- [ ] Framework filtering (`--frameworks Foundation,UIKit`)
- [ ] Format selection (`--format json|markdown`)
- [ ] Output validation (verify structure)
- [ ] Generate metadata index (invoke indexing code)
- [ ] Error handling improvements
- [ ] Add `--update` flag (re-crawl existing)

**Success Criteria:**
- Metadata index generated alongside crawled data
- Fast queries without parsing all files
- Clear, intuitive CLI
- Shows progress during long operations
- Helpful error messages
- Works reliably on macOS/Linux/Windows

**Effort:** 5 days (3 for indexing, 2 for CLI polish)

### Phase 2: Documentation (Week 2)

**Goal:** Users can easily adopt the tool

**Tasks:**
- [ ] Update README with crawler-first approach
- [ ] Quick start guide (5-minute setup)
- [ ] DarwinKit integration example
- [ ] CI/CD usage patterns
- [ ] Caching strategies
- [ ] Troubleshooting guide
- [ ] FAQ

**Success Criteria:**
- New user can get started in <5 minutes
- DarwinKit integration documented
- Common issues covered

**Effort:** 3-5 days

### Phase 3: Release v1.0 (Week 3)

**Goal:** Ship v1.0.0

**Tasks:**
- [ ] Set up goreleaser config
- [ ] Build multi-platform binaries
  - darwin/amd64
  - darwin/arm64
  - linux/amd64
  - linux/arm64
  - windows/amd64
- [ ] Create pre-crawled tarball (convenience)
- [ ] Write release notes
- [ ] Update CHANGELOG
- [ ] Tag v1.0.0
- [ ] Publish release
- [ ] Announcement (blog post / reddit)

**Success Criteria:**
- Binaries work on all platforms
- Installation is simple
- Documentation is complete
- v1.0.0 tagged and published

**Effort:** 3-5 days

**Total Timeline:** 3 weeks to v1.0.0

**Note on Metadata Encoding:** The index generation is critical for performance. Without it, every query would require scanning thousands of JSON files. The index acts as a fast lookup layer, making operations like `ListFrameworks()` and `SearchSymbols()` instant.

## Distribution Strategy: Clean Consumption

The goal is **dead-simple consumption** of Apple documentation data.

### Three Consumption Patterns

#### Pattern 1: Zero-Setup with Pre-Built Index (Recommended)

Download pre-crawled data with metadata index already built:

```bash
# One command to get started
curl -L https://github.com/tmc/appledocs/releases/v17.0.0.tar.zst | \
  tar -I zstd -x -C ~/.appledocs/v17

# Includes:
# - All 376 frameworks (63,449 JSON files)
# - Pre-built metadata index
# - Ready to query instantly
```

**Go usage:**
```go
import "github.com/tmc/appledocs"

func main() {
    // One line to open
    docs := appledocs.MustOpen("~/.appledocs/v17")

    // Instant queries (no file scanning)
    frameworks := docs.ListFrameworks()              // []string
    classes := docs.ListClasses("Foundation")         // []string
    doc := docs.GetSymbol("Foundation/NSString")      // *Document

    fmt.Println(doc.Title, doc.Abstract)
}
```

**Why this is clean:**
- ✅ Single curl command to get data
- ✅ No crawling required
- ✅ Metadata index pre-built
- ✅ Works immediately

#### Pattern 2: Fresh Crawl (Always Current)

For users who want latest data from Apple:

```bash
# Install crawler
go install github.com/tmc/appledocs/cmd/appledocs@latest

# Crawl (builds metadata index automatically)
appledocs crawl --output ~/.appledocs/v17

# Done - same API as Pattern 1
```

**Go usage:** Same as Pattern 1 - the API doesn't care if data came from tarball or crawl.

**Why this is clean:**
- ✅ One command to install
- ✅ One command to crawl
- ✅ Metadata index auto-generated
- ✅ Same consumption API

#### Pattern 3: Selective Frameworks (CI/CD)

For code generators that only need specific frameworks:

```bash
# In Makefile or CI config
appledocs crawl \
  --frameworks Foundation,UIKit,AppKit \
  --output ./docs \
  --format json
```

**Go usage:**
```go
docs := appledocs.MustOpen("./docs")

// Query works even with subset of frameworks
for _, class := range docs.ListClasses("Foundation") {
    doc := docs.GetSymbol("Foundation/" + class)
    generateBinding(doc)
}
```

**Why this is clean:**
- ✅ Selective download (smaller, faster)
- ✅ Same API works with partial data
- ✅ Perfect for CI pipelines

### Key Design Principle: Uniform API

**No matter how you get the data, the consumption API is identical:**

```go
// Pattern 1: Pre-built tarball
docs := appledocs.MustOpen("~/.appledocs/v17")

// Pattern 2: Fresh crawl
docs := appledocs.MustOpen("~/.appledocs/v17")

// Pattern 3: Selective frameworks
docs := appledocs.MustOpen("./docs")

// All three work identically:
frameworks := docs.ListFrameworks()
classes := docs.ListClasses("Foundation")
doc := docs.GetSymbol("Foundation/NSString")
```

### What Makes This Clean

1. **Zero configuration** - No environment variables, no config files
2. **One-line setup** - Single curl or appledocs command
3. **Instant queries** - Metadata index pre-built or auto-generated
4. **Uniform API** - Same code works regardless of data source
5. **No network calls** - Everything local after initial setup
6. **No hidden magic** - Clear what's happening at each step

### Clean API Design

**Single entry point, intuitive methods:**

```go
// One function to rule them all
docs := appledocs.MustOpen(path)

// Discovery APIs (instant via metadata index)
frameworks := docs.ListFrameworks()           // []string
classes := docs.ListClasses(framework)        // []string
protocols := docs.ListProtocols(framework)    // []string
methods := docs.ListMethods(class)            // []string

// Lookup APIs (one file read per call)
doc := docs.GetSymbol(path)                   // *Document
info := docs.GetSymbolInfo(path)              // *SymbolInfo

// Search APIs (using index)
results := docs.Search(pattern)               // []SearchResult
filtered := docs.Filter(criteria)             // []string
```

**Example: Complete workflow**
```go
import "github.com/tmc/appledocs"

func main() {
    // Open data (from any source)
    docs := appledocs.MustOpen("~/.appledocs/v17")

    // Discover what exists
    for _, fw := range docs.ListFrameworks() {
        fmt.Println("Framework:", fw)

        // List all classes in framework
        for _, class := range docs.ListClasses(fw) {
            // Get full documentation
            doc := docs.GetSymbol(fw + "/" + class)

            fmt.Printf("  %s: %s\n", doc.Title, doc.Abstract)

            // Process as needed
            generateCode(doc)
        }
    }
}
```

**Why this API is clean:**
- ✅ One type to import: `appledocs.Docs`
- ✅ Predictable method names: `List*`, `Get*`, `Search*`
- ✅ No fs.FS exposure (internal detail)
- ✅ No manual JSON parsing
- ✅ Obvious what each method does
- ✅ Works identically regardless of data source

## Design Philosophy

### Data-Driven Radical Simplicity

**Key insight:** The JSON IS the schema. Don't replicate it in Go types.

After analyzing 500+ JSON files, we discovered:
1. 100% consistent top-level structure
2. References use uniform map structure
3. Metadata follows predictable patterns

**Decision:** Build ~650 lines of hand-crafted, maintainable code instead of 4,000+ auto-generated types.

### What the Data Shows

**Document Structure (100% consistent):**
```
Every document has exactly these fields:
- identifier {interfaceLanguage, url}
- kind: "symbol" | "article"
- metadata {role, title, modules, ...}
- hierarchy {paths: [][]string}
- references: map[string]Reference
- schemaVersion {major, minor, patch}
- legalNotices {copyright, privacyPolicy, termsOfUse}
```

**Symbol Kinds (from actual data):**
- property (207)
- var (119)
- method (46)
- struct (37)
- init (16)
- class (16)
- enum (11)
- protocol (4)
- module (4)

### Programmatic Access Example

**The clean way to access Apple documentation data:**

```go
import "github.com/tmc/appledocs"

func main() {
    // Open (works with any source: tarball, crawl, selective)
    docs := appledocs.MustOpen("~/.appledocs/v17")

    // Get a specific symbol
    doc := docs.GetSymbol("Foundation/NSString")

    // Access structured data (type-safe)
    fmt.Println("Title:", doc.Title)
    fmt.Println("Kind:", doc.Kind)            // "class"
    fmt.Println("ExternalID:", doc.ExternalID)  // "c:objc(cs)NSString"
    fmt.Println("Abstract:", doc.Abstract)

    // Platform availability
    for _, platform := range doc.Platforms {
        fmt.Printf("- %s (since %s)\n", platform.Name, platform.IntroducedAt)
    }

    // Process methods
    for _, method := range docs.ListMethods("Foundation/NSString") {
        methodDoc := docs.GetSymbol(method)
        fmt.Printf("Method: %s\n", methodDoc.Title)
        generateBinding(methodDoc)
    }
}
```

**Key point:** No manual JSON parsing, no type assertions, no fs.FS juggling. Just clean, obvious method calls.

## Future Enhancements

### v1.1 - Smart Caching
- ETag support (avoid re-downloading unchanged files)
- Incremental updates
- Cache management commands

### v1.2 - Advanced Filtering
- Filter by platform (iOS only, macOS only)
- Filter by API level
- Exclude deprecated APIs

### v1.3 - Advanced Indexing
- Full-text search index
- Cross-reference index (find all usages)
- Inheritance hierarchy index
- Protocol conformance index

### v2.0 - CDN/Proxy (If Needed)
- Optional hosted service for pre-crawled data
- Would use OCI for layer deduplication
- Only if crawler bandwidth becomes issue
- Data-driven decision based on actual usage

### Optional Future Features

**Low Priority:**
1. Declaration truncation handling for very long signatures
2. Multi-line formatting for complex method signatures
3. REST API section support (restParameters, restEndpoints)
4. Properties section support
5. Mentions section rendering

**Future Possibilities:**
1. Language variant switching UI
2. Dark mode support
3. Search functionality
4. Right sidebar "On This Page" TOC
5. Automated navigation JSON generation

## What We're Not Building

❌ **fetch package** (272 lines) - Crawler replaces this
❌ **OCI distribution** (800+ lines) - Premature optimization
❌ **Custom cache system** - OS filesystem works fine
❌ **Data modules** - Users generate their own data

### Migration from Previous Plan

**For Users of fetch Package (Right Claude's work):**

Before:
```go
fsys, _ := fetch.Open("17.0.0")
```

After:
```bash
appledocs crawl --output ~/.appledocs/v17
```
```go
fsys, _ := appledocs.Open(os.ExpandEnv("$HOME/.appledocs/v17"))
```

**For Future OCI Interest (Left Claude's research):**

The OCI research is valuable and documented. If we ever need CDN distribution with deduplication, the foundation is there. But that's a v2.0 decision driven by actual usage data.

## Success Metrics

### Phase 1 Success
- [ ] Crawler runs reliably on all platforms
- [ ] Progress feedback is clear
- [ ] Framework filtering works
- [ ] Output is validated

### Phase 2 Success
- [ ] User can go from zero to working in <5 minutes
- [ ] DarwinKit integration tested
- [ ] Common questions answered

### v1.0 Success
- [ ] 10+ users successfully using the tool
- [ ] DarwinKit integration confirmed
- [ ] Zero critical bugs
- [ ] Documentation complete

## Lessons Learned

### What We Discovered

1. **15MB compression breakthrough** - Expected 30MB, achieved 15MB (95.2% reduction)
2. **DarwinKit pattern** - Users generate data themselves, don't distribute it
3. **Crawler already exists** - 90% done, just needs polish
4. **Ultra-think pays off** - Prevented 120+ hours of unnecessary OCI work

### Mistakes We Avoided

1. **Assuming distribution was the problem** - Generation is the solution
2. **Optimizing for size** - 15MB is fine, freshness matters more
3. **Building for speculation** - DarwinKit shows the real pattern
4. **Perfect before ship** - Ship crawler, iterate based on reality

### What Worked

- Parallel implementation (MVP + OCI research)
- Ultra-think analysis caught wrong direction early
- Questioning assumptions led to better solution
- Multi-Claude collaboration validated ideas

## Use Cases

### Code Generation (Primary)
```go
// DarwinKit-style bindings generation
frameworks, _ := appledocs.ListFrameworks(fsys)
for _, fw := range frameworks {
    doc, _ := appledocs.GetFramework(fsys, fw)
    for id, ref := range doc.References {
        if ref.SymbolKind == "class" {
            generateGoClass(ref.Title, doc)
        }
    }
}
```

### Documentation Search
```go
matches, _ := appledocs.SearchSymbols(fsys, "Foundation", "string")
for _, match := range matches {
    info, _ := appledocs.GetSymbolInfo(fsys, "Foundation/"+match)
    fmt.Printf("%s: %s\n", info.Title, info.Abstract)
}
```

### API Analysis
```go
symbols, _ := appledocs.ListSymbols(fsys, "Foundation")
for _, symbol := range symbols {
    doc, _ := appledocs.GetSymbol(fsys, "Foundation/"+symbol)
    checkPlatformSupport(doc.Metadata.Platforms)
}
```

## Performance Characteristics

- **Lazy loading:** Files read on-demand via `fs.FS`
- **No parsing overhead:** Direct JSON unmarshaling
- **Memory efficient:** Don't load entire corpus
- **Fast queries:** Direct map access, no ORM overhead

## Metrics

- **Lines of code:** ~650 (core package)
- **Example programs:** 6
- **Documentation:** ~5,000 lines
- **Research:** ~2,000 lines
- **Test coverage:** 22 tests (all passing)
- **Implementation time:** ~1 week
- **Compression ratio:** 21:1 (314MB → 15MB)
- **Frameworks supported:** 376
- **Binary size impact:** <1MB (without embedded data)

## Related Projects

- [DarwinKit](https://github.com/progrium/darwinkit) - Go bindings for macOS frameworks (primary use case)
- [protoc](https://github.com/protocolbuffers/protobuf) - Protocol buffer compiler (similar pattern)
- [sqlc](https://github.com/sqlc-dev/sqlc) - SQL to Go code generator (similar pattern)

## Timeline

```
Week 1: Metadata Encoding & CLI Polish
├── Day 1-3: Design and implement metadata indexing
│   ├── Design index schema
│   ├── Build indexer (parse all JSON once)
│   └── Generate index.go or index.db
├── Day 4-5: CLI improvements
│   ├── Flag handling, progress bars
│   └── Framework filtering, format selection

Week 2: Documentation
├── Day 1-2: README, quick start
├── Day 3: Integration examples
└── Day 4-5: FAQ, troubleshooting

Week 3: Release
├── Day 1-2: goreleaser setup, builds
├── Day 3: Testing, pre-crawled tarball
├── Day 4: Release prep, tag
└── Day 5: Publish, announce
```

## Conclusion

**Ship the crawler. Skip the distribution complexity.**

The answer was in the codebase all along - we already built the right tool. We just needed to recognize it.

---

**Status:** ✅ Ready for v1.0
**Confidence:** Very high (based on ecosystem patterns + real user needs)
**Risk:** Low (crawler already works, just needs polish)
**Effort:** ~40 hours vs 120+ for alternatives
**Next Action:** Polish CLI (Phase 1, Day 1)
