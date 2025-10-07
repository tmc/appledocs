# Project Status

**Last Updated:** October 7, 2025 (Ultra-Think Decision)

## What We Built

A Go package for programmatic access to Apple's developer documentation, plus a crawler tool for generating fresh documentation locally.

### Core Package (`github.com/tmc/appledocs`)

**Files:**
- `types.go` - Type definitions (121 lines)
- `fs.go` - Filesystem and query functions (272 lines)
- `maps.go` - Map-based helpers (200 lines)
- `doc.go` - Package documentation (86 lines)

**Total:** ~650 lines vs 50,000+ lines of auto-generated code

### Two Complementary APIs

#### 1. Typed API (Recommended - 95% of use cases)
```go
fsys, _ := appledocs.Open("docs/")
doc, _ := appledocs.GetSymbol(fsys, "Foundation/NSString")
title := doc.Metadata.Title        // Clean, no casts!
kind := doc.Metadata.SymbolKind    // Compile-time safe
```

#### 2. Map API (For flexibility)
```go
raw, _ := appledocs.LoadMap(fsys, "Foundation/NSString.json")
custom := appledocs.GetString(raw, "custom", "field")
```

### Example Programs (6 Total)

1. **list-frameworks** - Browse available frameworks
2. **list-classes** - Find symbols by type
3. **platform-analysis** - API availability analysis
4. **search-symbols** - Cross-framework search
5. **extract-methods** - Method signature extraction
6. **codegen-demo** - Code generation pattern (DarwinKit-style)

All examples build successfully and demonstrate real-world usage patterns.

## Distribution Research

Comprehensive research document: `docs/ALTERNATIVE_DISTRIBUTION_STRATEGIES.md` (1,917 lines)

### Key Findings

Evaluated 5 distribution strategies:

| Strategy | Score | Cost/Month | Recommendation |
|----------|-------|------------|----------------|
| **OCI Artifacts (GitHub)** | 33/35 | $0 | ⭐ **PRIMARY** |
| **GOPROXY + CDN** | 32/35 | $0-5 | Secondary |
| **SQLite** | 31/35 | $0 | Advanced features |
| **IPFS** | 24/35 | $0 | Not recommended |
| **Git LFS** | 19/35 | $387 | ❌ Avoid |

### Winner: GitHub Container Registry (OCI)

**Why:**
- Free forever (unlimited bandwidth for public packages)
- Industry standard (ML/AI community uses this)
- Best performance (global CDN, <50ms latency)
- Excellent Go integration (go-containerregistry)
- Zero maintenance (GitHub handles infrastructure)
- Layer deduplication (83% savings on version updates)

**Implementation timeline:** 2-3 weeks
**Cost:** $0/month

### Quick Win: GitHub Releases

For immediate MVP:
```bash
tar -cf - docs/ | zstd -19 -o appledocs-v17.tar.zst
gh release create v17.0.0 appledocs-v17.tar.zst
```

Simple download library provides fs.FS interface.

## Data Characteristics

- **Source:** ~314MB uncompressed markdown/JSON
- **Compressed:** ~30MB with gzip/zstd (90.4% reduction)
- **Files:** 63,449 files
- **Frameworks:** 376 frameworks
- **Symbols:** 1,284+ in Foundation alone

## Design Philosophy

**Data-driven radical simplicity:**

1. Analyzed 500 JSON files to understand actual structure
2. Discovered 100% consistent top-level schema
3. Rejected 4,000+ auto-generated types as overkill
4. Built ~650 lines of hand-crafted, maintainable code
5. Provides both type-safe and flexible access patterns

**Key insight:** The JSON IS the schema. Don't replicate it - just provide clean access.

## Testing

All core package tests pass:
```bash
✅ github.com/tmc/appledocs        PASS (11/11 tests)
✅ github.com/tmc/appledocs/reader PASS (11/11 tests) [deprecated, consolidated]
```

## Git History (Clean)

Recent commits:
- `0c4dd86` - Alternative distribution strategy analysis (research)
- `f6f664e` - Distribution strategy documentation
- `eb1c243` - Comprehensive example programs
- `b9dc8ec` - Consolidate reader package into main
- `7535d42` - Update embedding docs and navigation
- `55a0e74` - Type generation and analysis tools
- `2ecb05f` - Simple map-based API
- `8c796c6` - Reader package with fs.FS interface

## Documentation

- `README.md` - User-facing guide with API reference (349 lines)
- `DISTRIBUTION.md` - Distribution strategy overview
- `RADICAL_SIMPLICITY.md` - Design rationale (202 lines)
- `PROGRAMMATIC_ACCESS.md` - Programmatic usage guide
- `SUMMARY.md` - Project summary
- `docs/ALTERNATIVE_DISTRIBUTION_STRATEGIES.md` - Comprehensive research (1,917 lines)
- `docs/EMBEDDING.md` - Embedding considerations
- `examples/README.md` - Example programs guide

## Strategic Direction: Crawler-First

**Decision made:** October 7, 2025 after ultra-think analysis

**Approach:** Ship the crawler tool, not pre-packaged data

See [ULTRA_THINK_DECISION.md](ULTRA_THINK_DECISION.md) for complete rationale.

## Next Steps

### Week 1: CLI Polish

1. **Improve crawler CLI**
   - Better flag handling
   - Progress indicators
   - Framework filtering (`--frameworks Foundation,UIKit`)
   - Format selection (`--format json|markdown`)
   - **Effort:** 3-5 days

### Week 2: Documentation

2. **User Documentation**
   - Update README with crawler-first approach
   - Quick start guide (5-minute setup)
   - DarwinKit integration example
   - CI/CD usage patterns
   - **Effort:** 3-5 days

### Week 3: Release v1.0

3. **Ship v1.0.0**
   - Multi-platform binaries (goreleaser)
   - Pre-crawled tarball (convenience)
   - Release notes and changelog
   - Announcement
   - **Effort:** 3-5 days

### Future Enhancements (Data-Driven)

4. **Only if usage shows need:**
   - Smart caching (ETag-based)
   - Advanced filtering (platform, API level)
   - SQLite output format
   - CDN/OCI distribution (v2.0)

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

## Metrics

- **Lines of code:** ~650 (core package)
- **Example programs:** 6
- **Documentation:** ~5,000 lines
- **Research:** ~2,000 lines
- **Test coverage:** 22 tests (all passing)
- **Implementation time:** ~1 week
- **Compression ratio:** 10.5:1 (314MB → 30MB)
- **Frameworks supported:** 376
- **Binary size impact:** <1MB (without embedded data)

## Related Projects

- [DarwinKit](https://github.com/progrium/darwinkit) - Go bindings for macOS frameworks (target use case)

## License

MIT

---

**Status:** ✅ **Ready for v1.0** - Core complete, crawler exists, strategy finalized

**Next Action:** Polish crawler CLI (Week 1, see ROADMAP.md)

**Key Documents:**
- [ULTRA_THINK_DECISION.md](ULTRA_THINK_DECISION.md) - Why crawler-first
- [ROADMAP.md](ROADMAP.md) - Implementation plan
- [DESIGN_ANALYSIS.md](DESIGN_ANALYSIS.md) - Cost/benefit analysis
