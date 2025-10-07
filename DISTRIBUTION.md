# Distribution Strategy

This document describes how to distribute Apple documentation data using Go modules.

## Architecture

The project is split into three modules:

### 1. Core Library (`github.com/tmc/appledocs`)

The main library with types and query functions. **No data included.**

```bash
go get github.com/tmc/appledocs@latest
```

Size: ~100KB

### 2. Data Module (`github.com/tmc/appledocs-data`)

Complete Apple documentation dataset, versioned by SDK release.

```bash
go get github.com/tmc/appledocs-data/v17@latest  # iOS 17 / macOS 14
```

Size: ~2GB uncompressed (future: 200-400MB compressed)

### 3. Fetch Module (`github.com/tmc/appledocs-fetch`) [Future]

On-demand framework loader with caching.

```bash
go get github.com/tmc/appledocs-fetch@latest
```

## Usage Patterns

### Pattern 1: Local Documentation

Download docs manually, use library to read them:

```go
import "github.com/tmc/appledocs"

func main() {
    // Use locally downloaded docs
    fsys, _ := appledocs.Open("./output/tutorials/data/documentation")
    doc, _ := appledocs.GetSymbol(fsys, "Foundation/NSString")
}
```

**Best for:** Development, custom doc sources

### Pattern 2: Embedded Data Module

Import the data module to embed docs in your binary:

```go
import (
    "github.com/tmc/appledocs"
    data "github.com/tmc/appledocs-data/v17"
)

func main() {
    // Use embedded documentation
    fsys := data.FS()
    doc, _ := appledocs.GetSymbol(fsys, "Foundation/NSString")
}
```

**Best for:** Offline tools, single-binary distribution

**Note:** Your binary will be 200-400MB larger

### Pattern 3: On-Demand Fetching [Future]

Download only needed frameworks, cache locally:

```go
import "github.com/tmc/appledocs-fetch"

func main() {
    // Downloads Foundation docs on first use, caches to ~/.cache/appledocs/
    fsys, _ := fetch.Framework("Foundation", "17.0")
    doc, _ := appledocs.GetSymbol(fsys, "Foundation/NSString")
}
```

**Best for:** CI/CD, selective framework usage

## Versioning

Data modules use semantic versioning tied to SDK releases:

- `github.com/tmc/appledocs-data/v15` - iOS 15 / macOS 12
- `github.com/tmc/appledocs-data/v16` - iOS 16 / macOS 13
- `github.com/tmc/appledocs-data/v17` - iOS 17 / macOS 14

The core library (`appledocs`) works with any version.

## Creating a Data Module

### Structure

```
github.com/tmc/appledocs-data/
├── v15/
│   ├── go.mod                    # module github.com/tmc/appledocs-data/v15
│   ├── docs.go                   # Exports fs.FS
│   └── data/
│       └── tutorials/data/documentation/
├── v16/
│   ├── go.mod
│   ├── docs.go
│   └── data/
└── v17/
    ├── go.mod
    ├── docs.go
    └── data/
```

### Example `docs.go`

```go
// Package appledocsdata provides embedded Apple documentation for iOS 17/macOS 14.
package appledocsdata

import (
	"embed"
	"io/fs"
)

// Version information
const (
	IOSVersion   = "17.0"
	MacOSVersion = "14.0"
	SDKDate      = "2023-09-18"
)

//go:embed data/tutorials/data/documentation
var docsFS embed.FS

// FS returns the embedded documentation filesystem.
// Use with github.com/tmc/appledocs to query the documentation.
func FS() fs.FS {
	sub, _ := fs.Sub(docsFS, "data/tutorials/data/documentation")
	return sub
}
```

### Example `go.mod`

```go
module github.com/tmc/appledocs-data/v17

go 1.21

// No dependencies - just data
```

## Repository Organization

### Option A: Monorepo

```
github.com/tmc/appledocs/
├── go.mod                        # Core library
├── types.go
├── fs.go
├── maps.go
└── data/
    ├── v15/
    │   └── go.mod                # Submodule
    ├── v16/
    │   └── go.mod
    └── v17/
        └── go.mod
```

**Pros:** Single repo, easier maintenance
**Cons:** Large repo size, all versions in one place

### Option B: Separate Repositories

```
github.com/tmc/appledocs/         # Core library
github.com/tmc/appledocs-data/    # Data only
```

**Pros:** Keeps main repo small, data repo can be large
**Cons:** Multiple repos to manage

**Recommendation:** Option B - Keep repos separate

## Size Optimization

Current: ~2GB uncompressed JSON

Future optimizations:
1. **Gzip compression**: 2GB → ~400MB (80% reduction)
2. **Deduplication**: Share common reference data
3. **Binary format**: JSON → protobuf/msgpack
4. **Lazy loading**: Decompress on-demand

Target: 200-400MB per SDK version

## Go Module Proxy Benefits

When published to pkg.go.dev:

1. **Caching**: Proxy caches modules, reducing download times
2. **Integrity**: Checksums ensure data integrity
3. **Discovery**: Users can find via `go get`
4. **Versioning**: Immutable versions, no breaking changes
5. **Offline**: Proxy provides mirrors

## Distribution Checklist

### Phase 1: Core Library (✅ Done)
- [x] Create core library without embedded data
- [x] Support any fs.FS as input
- [x] Document local usage patterns

### Phase 2: Data Module
- [ ] Create `appledocs-data` repository
- [ ] Add iOS 17/macOS 14 documentation
- [ ] Version as v17.0.0
- [ ] Publish to pkg.go.dev
- [ ] Document embedding usage

### Phase 3: Compression
- [ ] Implement gzip compression/decompression
- [ ] Update data module with compressed data
- [ ] Benchmark decompression performance
- [ ] Document size savings

### Phase 4: On-Demand Fetching
- [ ] Create `appledocs-fetch` module
- [ ] Implement download + cache logic
- [ ] Support cache directory configuration
- [ ] Add version selection
- [ ] Document fetching usage

### Phase 5: Historical Versions
- [ ] Add v16 (iOS 16/macOS 13)
- [ ] Add v15 (iOS 15/macOS 12)
- [ ] Document version support policy
- [ ] Provide migration guide

## Size Comparison

| Approach | Size | Download Time (100Mbps) |
|----------|------|------------------------|
| Current (uncompressed) | 2.0 GB | ~3 minutes |
| Gzip compressed | ~400 MB | ~30 seconds |
| Selective (per framework) | ~5-50 MB | <5 seconds |
| On-demand (cached) | 0 (initial) | As needed |

## Recommendations

1. **Start simple**: Create data module with current uncompressed JSON
2. **Version by SDK**: Use v15, v16, v17 for major releases
3. **Add compression**: Once basic distribution works
4. **Add selective loading**: After compression is proven
5. **Consider CDN**: For very large-scale usage

## Example: Publishing Data Module

```bash
# 1. Create data repository
mkdir appledocs-data
cd appledocs-data

# 2. Create v17 submodule
mkdir v17
cd v17

# 3. Initialize module
go mod init github.com/tmc/appledocs-data/v17

# 4. Copy documentation
cp -r ~/appledocs/output/tutorials/data/documentation data/

# 5. Create docs.go (see above)
cat > docs.go << 'EOF'
package appledocsdata
// ... (content from above)
EOF

# 6. Test it works
go test

# 7. Commit and tag
git add .
git commit -m "feat: Add iOS 17/macOS 14 documentation"
git tag v17.0.0
git push origin main v17.0.0

# 8. Users can now use it:
# go get github.com/tmc/appledocs-data/v17@v17.0.0
```

## Questions?

- **Why separate modules?** Keeps core library small, allows version selection
- **Why not embed in main module?** 2GB embedded would make every import slow
- **Why version by SDK?** Clear mapping to Apple's release cycle
- **Can I use multiple versions?** Yes, import different versions as needed
- **What about storage?** Go proxy caches, users download once
