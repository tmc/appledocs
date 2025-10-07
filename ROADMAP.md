# AppleDocs Roadmap - Crawler-First Approach

**Updated:** October 7, 2025 (Ultra-Think Decision)
**Status:** Ready for v1.0
**Approach:** Distribute the crawler, not the data

## Executive Decision

**After deep architectural analysis:** Don't distribute data. Distribute the tool that generates the data.

See [ULTRA_THINK_DECISION.md](ULTRA_THINK_DECISION.md) for complete rationale.

## Current State

✅ **Complete:**
- Core library (650 lines: types, fs, maps)
- Working crawler (cmd/appledocs/)
- Markdown/HTML conversion
- 6 example programs
- Comprehensive documentation

🎯 **Strategy:** Ship the crawler as primary tool

## The Crawler Pattern

### For End Users
```bash
# Install
go install github.com/tmc/appledocs/cmd/appledocs@latest

# Generate docs
appledocs crawl --frameworks Foundation,UIKit
```

### For Code Generators (DarwinKit)
```bash
# In CI/build
appledocs crawl --format json --frameworks Foundation,AppKit
```

### For Convenience
```bash
# Pre-crawled tarball (optional)
curl -L https://github.com/tmc/appledocs/releases/v17.0.0.tar.zst | \
  tar -I zstd -x -C ~/.appledocs/v17
```

## Why This Approach

| Benefit | Description |
|---------|-------------|
| **Always Current** | Crawl latest from Apple anytime |
| **Selective** | Choose frameworks you need |
| **Flexible** | JSON or Markdown output |
| **Simple** | Tool already exists, just needs polish |
| **Ecosystem Fit** | Like protoc, sqlc, DarwinKit |

## Implementation Phases

### Phase 1: CLI Polish (Week 1)

**Goal:** Make crawler production-ready

**Tasks:**
- [ ] Improve CLI flags (use cobra or better flag handling)
- [ ] Add progress indicators (download/conversion progress)
- [ ] Framework filtering (`--frameworks Foundation,UIKit`)
- [ ] Format selection (`--format json|markdown`)
- [ ] Output validation (verify structure)
- [ ] Error handling improvements
- [ ] Add `--update` flag (re-crawl existing)

**Success Criteria:**
- Clear, intuitive CLI
- Shows progress during long operations
- Helpful error messages
- Works reliably on macOS/Linux/Windows

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

### Phase 3: Release Prep (Week 3)

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

## Future Enhancements (Post-v1.0)

### v1.1 - Smart Caching
- ETag support (avoid re-downloading unchanged files)
- Incremental updates
- Cache management commands

### v1.2 - Advanced Filtering
- Filter by platform (iOS only, macOS only)
- Filter by API level
- Exclude deprecated APIs

### v1.3 - Output Formats
- SQLite database output
- JSON streaming
- Custom templates

### v2.0 - CDN/Proxy (If Needed)
- Optional hosted service for pre-crawled data
- Would use OCI for layer deduplication
- Only if crawler bandwidth becomes issue

## What We're NOT Building

❌ **fetch package** - Crawler replaces this
❌ **OCI distribution** - Premature optimization
❌ **Custom cache system** - OS filesystem works fine

## Migration from Previous Plan

### For Users of fetch Package (from Right Claude)
**Before:**
```go
fsys, _ := fetch.Open("17.0.0")
```

**After:**
```bash
appledocs crawl --output ~/.appledocs/v17
```
```go
fsys, _ := appledocs.Open(os.ExpandEnv("$HOME/.appledocs/v17"))
```

### For Future OCI Interest (from Left Claude)
The OCI research is valuable and documented. If we ever need CDN distribution with deduplication, the foundation is there. But that's a v2.0 decision driven by actual usage data.

## Timeline

```
Week 1: CLI Polish
├── Day 1-2: Flag improvements, progress bars
├── Day 3-4: Framework filtering, format selection
└── Day 5: Testing, validation

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

**Total: 3 weeks to v1.0.0**

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

## Resources

- [ULTRA_THINK_DECISION.md](ULTRA_THINK_DECISION.md) - Why this approach
- [DESIGN_ANALYSIS.md](DESIGN_ANALYSIS.md) - Mathematical analysis
- [STATUS.md](STATUS.md) - Current state
- [README.md](README.md) - User guide

## Lessons Learned

1. **Question assumptions** - "Distribution" wasn't the problem
2. **Look at existing tools** - Crawler was already there
3. **Study ecosystem** - DarwinKit showed the pattern
4. **Ultra-think pays off** - Avoided 120+ hours of OCI work

---

**Next Action:** Polish CLI (Phase 1, Day 1)
**Status:** Ready to implement!
