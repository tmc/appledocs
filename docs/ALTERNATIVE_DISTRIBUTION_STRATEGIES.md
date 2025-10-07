# Alternative Distribution Strategies for Apple Documentation Data

**Research Date:** October 2025
**Data Size:** ~314MB uncompressed JSON (63,449 files)
**Compression Ratio:** ~90% (314MB → 30MB with gzip)

## Executive Summary

This document evaluates alternative distribution strategies beyond the Go module approaches already documented in `DISTRIBUTION.md`. After researching modern distribution technologies, we recommend implementing **three complementary strategies** in this order:

1. **OCI Artifacts** (Primary recommendation) - Best overall solution
2. **Custom GOPROXY + CDN** (Hybrid approach) - Best Go ecosystem integration
3. **Content-Addressed Storage with IPFS** (Future-looking) - Best for decentralization

## Current State

The `github.com/tmc/appledocs` project has:
- Core library using `fs.FS` abstraction (~100KB)
- ~314MB of Apple documentation JSON (compresses to ~30MB)
- Already designed: embedded data module, on-demand fetch module
- Excellent compression characteristics (90% size reduction)

## Evaluation Criteria

Each approach is evaluated on:
1. **Size Efficiency** - Storage and bandwidth optimization
2. **Download Speed** - Time to first byte and total transfer
3. **Versioning** - SDK version management capabilities
4. **Go Integration** - Compatibility with Go tooling and philosophy
5. **Offline Support** - Local caching and offline usage
6. **Cost** - Infrastructure and operational costs
7. **Implementation Complexity** - Development and maintenance effort

---

## Strategy 1: OCI Artifacts (RECOMMENDED)

### Overview

Use OCI (Open Container Initiative) registries to distribute documentation as container artifacts. This leverages the same infrastructure used for Docker images but for arbitrary data.

### Why OCI Artifacts?

The OCI Image Specification v1.1 (released February 2024) added first-class support for distributing arbitrary artifacts beyond container images. This is now the standard way to distribute ML models, WebAssembly modules, and large datasets.

### Architecture

```
github.com/tmc/appledocs (Go library)
           ↓
    ghcr.io/tmc/appledocs-data:v17
    ghcr.io/tmc/appledocs-data:v16
    ghcr.io/tmc/appledocs-data:v15
           ↓ (pulls compressed layers)
    Local registry cache
           ↓
    Go application via fs.FS
```

### Implementation Design

```go
// Package ocidata provides OCI artifact-based documentation loading
package ocidata

import (
    "context"
    "io/fs"

    "github.com/google/go-containerregistry/pkg/crane"
    "github.com/google/go-containerregistry/pkg/v1/tarball"
)

// Pull downloads and extracts documentation from OCI registry
func Pull(ctx context.Context, ref string) (fs.FS, error) {
    // ref format: "ghcr.io/tmc/appledocs-data:v17"

    // 1. Pull image layers
    img, err := crane.Pull(ref)
    if err != nil {
        return nil, err
    }

    // 2. Extract to cache directory
    cacheDir := getCacheDir() // ~/.cache/appledocs/v17/
    if err := tarball.Extract(img, cacheDir); err != nil {
        return nil, err
    }

    // 3. Return fs.FS from cache
    return os.DirFS(cacheDir), nil
}

// Usage in client code
func main() {
    fsys, _ := ocidata.Pull(context.Background(),
        "ghcr.io/tmc/appledocs-data:v17")
    doc, _ := appledocs.GetSymbol(fsys, "Foundation/NSString")
}
```

### Publishing Workflow

```bash
# 1. Create OCI artifact with compressed layers
crane append \
  -f <(tar czf - docs/tutorials/data/documentation) \
  -t ghcr.io/tmc/appledocs-data:v17 \
  scratch

# 2. Add metadata
crane mutate ghcr.io/tmc/appledocs-data:v17 \
  --label "org.opencontainers.image.title=Apple Documentation iOS 17" \
  --label "org.opencontainers.image.version=17.0.0" \
  --label "sdk.ios.version=17.0" \
  --label "sdk.macos.version=14.0"

# 3. Push to registry
crane push ghcr.io/tmc/appledocs-data:v17

# GitHub Actions automation
- name: Publish OCI Artifact
  run: |
    echo ${{ secrets.GITHUB_TOKEN }} | crane auth login ghcr.io -u ${{ github.actor }} --password-stdin
    crane append -f <(tar czf - docs/) -t ghcr.io/tmc/appledocs-data:v$VERSION scratch
    crane push ghcr.io/tmc/appledocs-data:v$VERSION
```

### Evaluation

| Criterion | Score | Notes |
|-----------|-------|-------|
| **Size Efficiency** | ⭐⭐⭐⭐⭐ | Built-in layer compression, deduplication across versions |
| **Download Speed** | ⭐⭐⭐⭐⭐ | CDN-backed registries (GitHub, Docker Hub), parallel layer downloads |
| **Versioning** | ⭐⭐⭐⭐⭐ | Perfect - tags, digests, manifest lists for multi-platform |
| **Go Integration** | ⭐⭐⭐⭐ | Excellent libraries (go-containerregistry), not native to `go get` |
| **Offline Support** | ⭐⭐⭐⭐⭐ | Automatic local caching via Docker/registry cache |
| **Cost** | ⭐⭐⭐⭐⭐ | Free on GitHub Container Registry, Docker Hub (public) |
| **Complexity** | ⭐⭐⭐⭐ | Well-documented, mature ecosystem |

### Pros

1. **Industry Standard**: OCI is the standard for distributing large artifacts in 2024+
2. **Excellent Compression**: Built-in layer compression and deduplication
3. **Free Infrastructure**: GitHub Container Registry (ghcr.io) is free for public projects
4. **Global CDN**: Automatic worldwide distribution via registry mirrors
5. **Layer Sharing**: Multiple SDK versions can share common data (deduplication)
6. **Mature Tooling**: crane, skopeo, oras for CLI; go-containerregistry for Go
7. **Metadata Rich**: Labels, annotations for SDK versions, platform info
8. **Offline Capable**: Works with local registry, skopeo copy for air-gapped
9. **Integrity**: Content-addressed layers with SHA256 verification
10. **Multi-arch Support**: Single manifest list can serve different platforms

### Cons

1. **Not Native Go**: Requires additional library, not integrated with `go get`
2. **Registry Dependency**: Needs OCI registry running (though many free options)
3. **Learning Curve**: Developers unfamiliar with containers need to learn OCI concepts
4. **Binary Size**: Adding go-containerregistry adds ~5MB to binary

### Cost Analysis

- **GitHub Container Registry (ghcr.io)**: Free for public repos, unlimited bandwidth
- **Docker Hub**: Free tier includes unlimited public pulls
- **Self-hosted**: Harbor (free, open-source) or cloud registry (~$5-50/month)

**Verdict**: Essentially free for this use case

### Implementation Roadmap

**Phase 1: MVP (Week 1)**
- [ ] Create `github.com/tmc/appledocs-oci` module
- [ ] Implement Pull() function using go-containerregistry
- [ ] Publish v17 artifact to ghcr.io
- [ ] Add documentation and examples

**Phase 2: Optimization (Week 2)**
- [ ] Implement smart caching (check local first)
- [ ] Add progress bars for downloads
- [ ] Support manifest lists (future multi-platform)
- [ ] Implement layer deduplication for version updates

**Phase 3: Polish (Week 3)**
- [ ] Add GitHub Actions for automated publishing
- [ ] Implement `appledocs pull v17` CLI command
- [ ] Support air-gapped workflows (export/import)
- [ ] Performance benchmarks and documentation

---

## Strategy 2: Custom GOPROXY + CDN (HYBRID)

### Overview

Create a custom Go module proxy that serves documentation as Go modules, backed by a CDN for global distribution. This combines the simplicity of `go get` with the performance of CDN delivery.

### Architecture

```
┌──────────────────────────────────┐
│   User runs: go get ...data/v17  │
└────────────────┬─────────────────┘
                 ↓
┌────────────────────────────────────┐
│  GOPROXY=https://proxy.appledocs  │
│           .dev                     │
│  (Custom proxy server)             │
└────────────┬───────────────────────┘
             ↓
┌────────────────────────────────────┐
│  CloudFront CDN                    │
│  (Caches compressed modules)       │
└────────────┬───────────────────────┘
             ↓
┌────────────────────────────────────┐
│  S3 Bucket: appledocs-modules      │
│  ├── v17/@v/v17.0.0.zip (30MB)    │
│  ├── v16/@v/v16.0.0.zip            │
│  └── v15/@v/v15.0.0.zip            │
└────────────────────────────────────┘
```

### Implementation Design

#### Custom Proxy Server

```go
// server.go - Custom GOPROXY implementation
package main

import (
    "net/http"
    "github.com/goproxy/goproxy"
)

func main() {
    // Create proxy with S3+CloudFront backend
    proxy := &goproxy.Goproxy{
        // Fetch from S3 via CloudFront
        Fetcher: &cloudFrontFetcher{
            baseURL: "https://d111111abcdef8.cloudfront.net",
        },

        // Cache locally for 24 hours
        Cacher: &goproxy.DirCacher{
            Dir: "/var/cache/goproxy",
            MaxAge: 24 * time.Hour,
        },
    }

    http.ListenAndServe(":8080", proxy)
}

type cloudFrontFetcher struct {
    baseURL string
}

func (f *cloudFrontFetcher) Fetch(ctx context.Context, mod, ver string) (io.ReadCloser, error) {
    // Fetch from: https://cdn/.../v17/@v/v17.0.0.zip
    url := fmt.Sprintf("%s/%s/@v/%s.zip", f.baseURL, mod, ver)
    resp, err := http.Get(url)
    if err != nil {
        return nil, err
    }
    return resp.Body, nil
}
```

#### Publishing Pipeline

```bash
#!/bin/bash
# publish.sh - Publish module to S3+CloudFront

VERSION="v17.0.0"
MODULE="github.com/tmc/appledocs-data/v17"

# 1. Create module directory
mkdir -p "/tmp/module/data/tutorials/data/documentation"
cp -r docs/tutorials/data/documentation/* "/tmp/module/data/tutorials/data/documentation/"

# 2. Create go.mod
cat > /tmp/module/go.mod <<EOF
module $MODULE

go 1.21
EOF

# 3. Create data.go
cat > /tmp/module/data.go <<EOF
//go:build !nodata

package appledocsdata

import (
    "embed"
    "io/fs"
)

//go:embed data/tutorials/data/documentation
var docsFS embed.FS

func FS() fs.FS {
    sub, _ := fs.Sub(docsFS, "data/tutorials/data/documentation")
    return sub
}
EOF

# 4. Create zip (with compression)
cd /tmp/module
zip -r -9 "/tmp/${VERSION}.zip" .

# 5. Upload to S3
aws s3 cp "/tmp/${VERSION}.zip" \
    "s3://appledocs-modules/v17/@v/${VERSION}.zip" \
    --content-type "application/zip" \
    --metadata "sdk-version=17.0,sdk-date=2023-09-18"

# 6. Upload metadata files
echo "$VERSION" > /tmp/latest
aws s3 cp /tmp/latest "s3://appledocs-modules/v17/@latest"

# 7. Create and upload .info file
cat > /tmp/${VERSION}.info <<EOF
{
  "Version": "$VERSION",
  "Time": "$(date -Iseconds)"
}
EOF
aws s3 cp "/tmp/${VERSION}.info" \
    "s3://appledocs-modules/v17/@v/${VERSION}.info"

# 8. Invalidate CloudFront cache
aws cloudfront create-invalidation \
    --distribution-id E1234ABCD \
    --paths "/v17/*"
```

#### Client Usage

```bash
# Set custom proxy
export GOPROXY=https://proxy.appledocs.dev,https://proxy.golang.org,direct

# Use as normal Go module
go get github.com/tmc/appledocs-data/v17@latest
```

```go
import data "github.com/tmc/appledocs-data/v17"

func main() {
    fsys := data.FS()
    doc, _ := appledocs.GetSymbol(fsys, "Foundation/NSString")
}
```

### Evaluation

| Criterion | Score | Notes |
|-----------|-------|-------|
| **Size Efficiency** | ⭐⭐⭐⭐⭐ | Zip compression, CloudFront compression |
| **Download Speed** | ⭐⭐⭐⭐⭐ | CloudFront edge caching worldwide |
| **Versioning** | ⭐⭐⭐⭐⭐ | Native Go module versioning |
| **Go Integration** | ⭐⭐⭐⭐⭐ | Perfect - works with `go get` |
| **Offline Support** | ⭐⭐⭐⭐⭐ | Standard Go module cache |
| **Cost** | ⭐⭐⭐⭐ | S3 (~$1-5/month) + CloudFront (~$5-20/month) |
| **Complexity** | ⭐⭐⭐ | Requires server operation, AWS setup |

### Pros

1. **Native Go Experience**: Works with `go get`, `go.mod`, all standard tooling
2. **Transparent**: Users don't need to learn new tools
3. **Global Performance**: CloudFront provides <100ms latency worldwide
4. **Automatic Caching**: Go's module cache handles local storage
5. **Version Immutability**: Once published, versions never change (Go guarantee)
6. **Familiar Workflow**: Standard Go module semantics
7. **Compression**: Multiple layers (zip + CloudFront gzip)
8. **Easy Discovery**: Shows up in pkg.go.dev
9. **Access Control**: Can implement private module access via proxy
10. **Metrics**: CloudFront provides detailed analytics

### Cons

1. **Infrastructure Cost**: ~$10-30/month for S3 + CloudFront
2. **Operational Overhead**: Need to maintain proxy server
3. **Dependency**: Relies on custom proxy availability
4. **Fallback Complexity**: Need graceful degradation if proxy down
5. **Module Size**: Still embeds 30MB+ in binary when used

### Cost Analysis (Monthly)

```
S3 Storage:
  30MB × 3 versions = 90MB
  $0.023/GB × 0.09GB = $0.002/month

S3 Requests (estimate 10K downloads/month):
  10K GET requests × $0.0004/1000 = $0.004/month

Data Transfer Out (to CloudFront):
  10K × 30MB = 300GB
  First 10TB free to CloudFront = $0/month

CloudFront:
  300GB × $0.085/GB = $25.50/month
  10K requests × $0.01/10000 = $0.01/month

Total: ~$25-30/month (at 10K downloads/month)
```

For low usage (<100 downloads/month): ~$1/month
For high usage (100K downloads/month): ~$250/month

**Note**: Could reduce costs with CloudFlare (free tier includes CDN)

### Implementation Roadmap

**Phase 1: Infrastructure (Week 1)**
- [ ] Set up S3 bucket with versioned storage
- [ ] Configure CloudFront distribution
- [ ] Deploy goproxy server (Lambda or Cloud Run for cost efficiency)
- [ ] Set up DNS: proxy.appledocs.dev

**Phase 2: Automation (Week 2)**
- [ ] GitHub Actions workflow for publishing
- [ ] Automated version tagging
- [ ] Health checks and monitoring
- [ ] Documentation for contributors

**Phase 3: Enhancement (Week 3)**
- [ ] Add analytics dashboard
- [ ] Implement rate limiting
- [ ] Set up alerts for errors
- [ ] Performance optimization

### Alternative: CloudFlare-based (Low Cost)

Use CloudFlare R2 (S3-compatible) + CloudFlare CDN = $0/month for reasonable usage:

```
CloudFlare R2:
  - 10GB storage free
  - Egress to CloudFlare CDN: FREE (no egress fees!)
  - First 1M reads/month free

CloudFlare CDN:
  - Free tier: unlimited bandwidth for cache hits
  - Global edge network

Total cost: $0/month for most usage patterns
```

**Recommendation**: Use CloudFlare R2 + CloudFlare CDN for cost optimization

---

## Strategy 3: Content-Addressed Storage with IPFS

### Overview

Use IPFS (InterPlanetary File System) for decentralized, content-addressed distribution. Each SDK version is published to IPFS and pinned for availability.

### Architecture

```
Developer publishes:
  docs/ → IPFS → CID: bafybeid...

Pinning services:
  - Pinata (primary)
  - Web3.Storage (backup)
  - Self-hosted node

User retrieves:
  IPFS Gateway → HTTP → Local IPFS node → Cache
```

### Implementation Design

```go
// ipfsdata/ipfs.go
package ipfsdata

import (
    "context"
    "io/fs"
    "net/http"

    "github.com/ipfs/go-ipfs-api"
)

// Known CIDs for each version
var versionCIDs = map[string]string{
    "v17": "bafybeigdyrzt5sfp7udm7hu76uh7y26nf3efuylqabf3oclgtqy55fbzdi",
    "v16": "bafybeib2w3xkhd4xpfocl72ypocjfrvq7kvvbhwqjsqvwvvvvvvvvvvvvv",
}

// Load downloads documentation via IPFS
func Load(ctx context.Context, version string) (fs.FS, error) {
    cid := versionCIDs[version]

    // Try local IPFS node first
    if fsys, err := loadFromLocalNode(ctx, cid); err == nil {
        return fsys, nil
    }

    // Fallback to public gateway
    return loadFromGateway(ctx, cid)
}

func loadFromLocalNode(ctx context.Context, cid string) (fs.FS, error) {
    shell := ipfsapi.NewShell("localhost:5001")

    // Download to cache
    cacheDir := fmt.Sprintf("~/.cache/appledocs/%s", cid)
    if err := shell.Get(cid, cacheDir); err != nil {
        return nil, err
    }

    return os.DirFS(cacheDir), nil
}

func loadFromGateway(ctx context.Context, cid string) (fs.FS, error) {
    // Use public gateway as fallback
    gateways := []string{
        "https://w3s.link/ipfs/",      // web3.storage
        "https://gateway.pinata.cloud/ipfs/",
        "https://ipfs.io/ipfs/",        // public gateway
    }

    for _, gateway := range gateways {
        url := gateway + cid
        if fsys, err := downloadAndCache(ctx, url, cid); err == nil {
            return fsys, nil
        }
    }

    return nil, fmt.Errorf("failed to load from any gateway")
}
```

### Publishing Workflow

```bash
#!/bin/bash
# publish-ipfs.sh

VERSION="v17"

# 1. Add to IPFS
CID=$(ipfs add -r -Q docs/tutorials/data/documentation)
echo "Published: $CID"

# 2. Pin to pinning services
# Pinata
curl -X POST "https://api.pinata.cloud/pinning/pinByHash" \
  -H "Authorization: Bearer $PINATA_JWT" \
  -H "Content-Type: application/json" \
  -d "{
    \"hashToPin\": \"$CID\",
    \"pinataMetadata\": {
      \"name\": \"appledocs-$VERSION\",
      \"keyvalues\": {
        \"version\": \"$VERSION\",
        \"sdk\": \"iOS 17\"
      }
    }
  }"

# Web3.Storage
w3 up docs/tutorials/data/documentation --name "appledocs-$VERSION"

# 3. Update version mapping in code
echo "Update versionCIDs map with: \"$VERSION\": \"$CID\""

# 4. Publish CID to DNS (optional, for discoverability)
# Set TXT record: _dnslink.appledocs.dev → dnslink=/ipfs/$CID
```

### Evaluation

| Criterion | Score | Notes |
|-----------|-------|-------|
| **Size Efficiency** | ⭐⭐⭐⭐ | Content-addressed deduplication |
| **Download Speed** | ⭐⭐⭐ | Varies by gateway/peer availability |
| **Versioning** | ⭐⭐⭐⭐⭐ | Each version has unique CID |
| **Go Integration** | ⭐⭐ | Requires custom code, not `go get` |
| **Offline Support** | ⭐⭐⭐⭐⭐ | Strong peer-to-peer caching |
| **Cost** | ⭐⭐⭐⭐ | Pinning services ~$0-20/month |
| **Complexity** | ⭐⭐ | Requires IPFS knowledge |

### Pros

1. **Decentralized**: No single point of failure
2. **Content Addressing**: Immutable, verifiable data
3. **Peer Caching**: Users become providers automatically
4. **Censorship Resistant**: Cannot be taken down
5. **Deduplication**: Across all IPFS content globally
6. **Low Cost**: Pinning services free tier often sufficient
7. **Future Proof**: Web3-native, growing ecosystem
8. **Bandwidth Sharing**: Popular content served by many peers
9. **Integrity**: CID cryptographically verifies content
10. **Cool Factor**: Cutting-edge technology

### Cons

1. **Performance Variability**: Depends on peer availability and gateway reliability
2. **Gateway Dependency**: Most users will use HTTP gateways (not true P2P)
3. **Limited Adoption**: Not standard in Go ecosystem
4. **Complex Setup**: Requires running IPFS node for best experience
5. **Pinning Required**: Need to pay pinning service or run own node
6. **Cold Start**: First retrieval can be slow without good peers
7. **Go Library Size**: IPFS libraries are large (~20MB added to binary)
8. **Reliability Concerns**: Public gateways can be slow or down
9. **No Native Compression**: Need to compress before adding to IPFS
10. **Learning Curve**: Steep for developers unfamiliar with IPFS

### Cost Analysis

**Pinning Services:**
- **Pinata**: Free tier (1GB), Pro $20/month (unlimited)
- **Web3.Storage**: Free tier (10GB)
- **Filebase**: $5.99/month for 1TB

**For 90MB total (3 versions × 30MB):**
- Pinata free tier: $0/month
- Web3.Storage: $0/month
- Self-hosted IPFS node: Server cost (~$5-20/month)

**Verdict**: Can be free or very low cost

### Implementation Roadmap

**Phase 1: Proof of Concept (Week 1-2)**
- [ ] Publish v17 to IPFS
- [ ] Pin to Pinata and Web3.Storage
- [ ] Implement gateway-based loader
- [ ] Test download performance from various locations

**Phase 2: Optimization (Week 2-3)**
- [ ] Implement CAR (Content Addressable Archive) format for faster loading
- [ ] Add multiple gateway fallbacks
- [ ] Implement smart caching
- [ ] Add local IPFS node detection

**Phase 3: Polish (Week 3-4)**
- [ ] Create CLI tool for IPFS publishing
- [ ] Implement DNSLink for version resolution
- [ ] Performance benchmarks vs other strategies
- [ ] Comprehensive documentation

### Hybrid Approach: IPFS + HTTP Gateway

Best of both worlds:

```go
// Try IPFS first, fallback to CDN
func Load(version string) (fs.FS, error) {
    // 1. Try local IPFS node (fast, P2P)
    if fsys, err := loadFromIPFS(version); err == nil {
        return fsys, nil
    }

    // 2. Try IPFS gateway (medium, HTTP)
    if fsys, err := loadFromGateway(version); err == nil {
        return fsys, nil
    }

    // 3. Fallback to CloudFront (reliable, fast)
    return loadFromCDN(version)
}
```

---

## Strategy 4: Git LFS (Not Recommended)

### Overview

Use Git Large File Storage to store documentation in the Git repository with pointer files.

### Evaluation

| Criterion | Score | Notes |
|-----------|-------|-------|
| **Size Efficiency** | ⭐⭐⭐ | LFS storage charges can accumulate |
| **Download Speed** | ⭐⭐⭐ | Good, but not as fast as CDN |
| **Versioning** | ⭐⭐⭐⭐ | Git versioning works well |
| **Go Integration** | ⭐⭐ | Requires git clone, not go get |
| **Offline Support** | ⭐⭐⭐ | Once cloned, works offline |
| **Cost** | ⭐⭐ | GitHub: $5/50GB, bandwidth charges |
| **Complexity** | ⭐⭐⭐⭐ | Simple, familiar Git workflow |

### Why Not Recommended

1. **Cost**: GitHub charges $0.0875/GB for bandwidth after free tier (50GB/month)
   - 1,000 downloads × 30MB = 30GB = $2.63/month
   - 10,000 downloads = 300GB = $26.25/month

2. **Not Go Native**: Requires `git clone`, doesn't work with `go get`

3. **LFS Complexity**: Users need Git LFS installed and configured

4. **Bandwidth Limits**: Can hit limits quickly with popular projects

5. **Cloning Overhead**: Downloads all data, not selective

6. **Better Alternatives**: OCI and GOPROXY offer better cost and performance

**Verdict**: Skip this approach - OCI artifacts provide all benefits without the costs

---

## Strategy 5: SQLite + Virtual Filesystem (Interesting Alternative)

### Overview

Package all documentation into a single SQLite database file, use Go's fs.FS via SQLite VFS.

### Architecture

```
docs.sqlite (30MB compressed)
  ├── Table: files
  │   ├── path TEXT PRIMARY KEY
  │   ├── content BLOB (JSON compressed)
  │   └── metadata JSON
  └── Indexes for fast lookup

Go code:
  → Open docs.sqlite
  → Query via SQL or fs.FS interface
  → In-memory caching for hot paths
```

### Implementation Design

```go
// sqlitedata/sqlite.go
package sqlitedata

import (
    "database/sql"
    "io/fs"

    _ "modernc.org/sqlite" // Pure Go SQLite
)

type sqliteFS struct {
    db *sql.DB
}

func Open(dbPath string) (fs.FS, error) {
    db, err := sql.Open("sqlite", dbPath)
    if err != nil {
        return nil, err
    }

    return &sqliteFS{db: db}, nil
}

func (fsys *sqliteFS) Open(name string) (fs.File, error) {
    var content []byte
    err := fsys.db.QueryRow(
        "SELECT decompress(content) FROM files WHERE path = ?",
        name,
    ).Scan(&content)

    if err != nil {
        return nil, err
    }

    return &sqliteFile{
        name:    name,
        content: content,
        offset:  0,
    }, nil
}

// Create database from JSON files
func Create(outputDB string, inputDir string) error {
    db, _ := sql.Open("sqlite", outputDB)

    db.Exec(`
        CREATE TABLE files (
            path TEXT PRIMARY KEY,
            content BLOB,  -- zstd compressed
            size INTEGER,
            modified INTEGER
        );
        CREATE INDEX idx_path ON files(path);
    `)

    // Walk directory and insert files
    filepath.WalkDir(inputDir, func(path string, d fs.DirEntry, err error) error {
        if d.IsDir() {
            return nil
        }

        content, _ := os.ReadFile(path)
        compressed := compress(content) // zstd

        db.Exec(
            "INSERT INTO files (path, content, size, modified) VALUES (?, ?, ?, ?)",
            path, compressed, len(content), d.Info().ModTime().Unix(),
        )

        return nil
    })

    db.Exec("VACUUM") // Compact database
    return nil
}
```

### Publishing

```bash
# Create SQLite database
appledocs-sqlite create \
  --input docs/tutorials/data/documentation \
  --output appledocs-v17.sqlite

# Publish to CDN or registry
aws s3 cp appledocs-v17.sqlite s3://appledocs/v17/docs.sqlite
# or
crane append -f appledocs-v17.sqlite -t ghcr.io/tmc/appledocs-data:v17-sqlite
```

### Evaluation

| Criterion | Score | Notes |
|-----------|-------|-------|
| **Size Efficiency** | ⭐⭐⭐⭐⭐ | Single file, excellent compression with zstd |
| **Download Speed** | ⭐⭐⭐⭐⭐ | Single file download, no overhead |
| **Versioning** | ⭐⭐⭐⭐ | Version by filename/tag |
| **Go Integration** | ⭐⭐⭐⭐ | Pure Go SQLite, fs.FS compatible |
| **Offline Support** | ⭐⭐⭐⭐⭐ | Single file, easy to cache |
| **Cost** | ⭐⭐⭐⭐⭐ | Minimal - single file to host |
| **Complexity** | ⭐⭐⭐ | Requires SQLite VFS implementation |

### Pros

1. **Single File**: Entire dataset in one ~30MB file
2. **Fast Queries**: SQL indexes for path lookups
3. **Excellent Compression**: zstd compression per-file + SQLite pages
4. **Pure Go**: Using modernc.org/sqlite (no CGO)
5. **Easy Distribution**: Just one file to download
6. **Atomic Updates**: Download new version as single transaction
7. **Efficient Selects**: Can load only needed files via SQL
8. **Metadata Queries**: SQL for advanced searches across docs
9. **Small Binary Impact**: Pure Go SQLite is reasonable size
10. **Proven Technology**: SQLite is battle-tested

### Cons

1. **Custom VFS**: Need to implement fs.FS wrapper (moderate complexity)
2. **Memory Usage**: SQLite cache may use more RAM than needed
3. **Less Standard**: Not a common pattern in Go ecosystem
4. **Write Performance**: Immutable data doesn't benefit from DB
5. **Overkill**: SQL queries not needed for simple file access

### Use Cases Where This Excels

1. **Advanced Queries**: Search across all docs for symbols
2. **Incremental Updates**: Download only changed files (delta updates)
3. **Embedded Systems**: Single file easier to manage
4. **Metadata Rich**: Complex queries on SDK metadata

### Implementation Roadmap

**Phase 1: Prototype (Week 1)**
- [ ] Implement fs.FS wrapper for SQLite
- [ ] Create conversion tool (JSON dir → SQLite)
- [ ] Benchmark performance vs directory access
- [ ] Test compression ratios (zstd vs gzip)

**Phase 2: Integration (Week 2)**
- [ ] Publish v17 as SQLite file
- [ ] Implement download-and-open helper
- [ ] Add caching strategy
- [ ] Performance optimization

**Phase 3: Advanced Features (Week 3)**
- [ ] Implement delta updates (download only changed files)
- [ ] Add full-text search across all docs
- [ ] Metadata query examples
- [ ] Documentation and benchmarks

---

## Strategy Comparison Matrix

| Strategy | Size | Speed | Versioning | Go Integration | Offline | Cost | Complexity | **Score** |
|----------|------|-------|------------|----------------|---------|------|------------|-----------|
| **OCI Artifacts** | ⭐⭐⭐⭐⭐ | ⭐⭐⭐⭐⭐ | ⭐⭐⭐⭐⭐ | ⭐⭐⭐⭐ | ⭐⭐⭐⭐⭐ | ⭐⭐⭐⭐⭐ | ⭐⭐⭐⭐ | **33/35** |
| **GOPROXY+CDN** | ⭐⭐⭐⭐⭐ | ⭐⭐⭐⭐⭐ | ⭐⭐⭐⭐⭐ | ⭐⭐⭐⭐⭐ | ⭐⭐⭐⭐⭐ | ⭐⭐⭐⭐ | ⭐⭐⭐ | **32/35** |
| **IPFS** | ⭐⭐⭐⭐ | ⭐⭐⭐ | ⭐⭐⭐⭐⭐ | ⭐⭐ | ⭐⭐⭐⭐⭐ | ⭐⭐⭐⭐ | ⭐⭐ | **24/35** |
| **SQLite** | ⭐⭐⭐⭐⭐ | ⭐⭐⭐⭐⭐ | ⭐⭐⭐⭐ | ⭐⭐⭐⭐ | ⭐⭐⭐⭐⭐ | ⭐⭐⭐⭐⭐ | ⭐⭐⭐ | **31/35** |
| **Git LFS** | ⭐⭐⭐ | ⭐⭐⭐ | ⭐⭐⭐⭐ | ⭐⭐ | ⭐⭐⭐ | ⭐⭐ | ⭐⭐⭐⭐ | **19/35** |

---

## Recommended Implementation Strategy

### Phase 1: OCI Artifacts (Primary - Weeks 1-2)

**Why First:**
- Industry standard for large artifacts in 2024
- Free infrastructure (GitHub Container Registry)
- Best overall balance of features
- Fastest time to value

**Deliverables:**
1. `github.com/tmc/appledocs-oci` module
2. Published artifacts: `ghcr.io/tmc/appledocs-data:v17`
3. Documentation and examples
4. GitHub Actions automation

**Success Metrics:**
- Download time <30 seconds worldwide
- Zero infrastructure cost
- Works offline after first download

### Phase 2: GOPROXY+CDN (Secondary - Weeks 3-4)

**Why Second:**
- Perfect Go ecosystem integration
- Transparent to users (just `go get`)
- Complements OCI for users who prefer standard Go workflow

**Implementation:**
- Use CloudFlare R2 + CloudFlare CDN (free tier)
- Deploy goproxy on Cloud Run (free tier: 2M requests/month)
- Total cost: $0-5/month

**Deliverables:**
1. Custom GOPROXY at `proxy.appledocs.dev`
2. Published modules compatible with `go get`
3. Automated publishing pipeline
4. Fallback to public proxy.golang.org

### Phase 3: Hybrid Approach (Optimization - Week 5)

**Best of All Worlds:**

```go
// appledocs/loader.go - Smart loader with fallbacks
package appledocs

func LoadDocs(version string, opts ...Option) (fs.FS, error) {
    cfg := defaultConfig()
    for _, opt := range opts {
        opt(cfg)
    }

    // Strategy 1: Check local cache first
    if fsys, err := loadFromCache(version); err == nil {
        return fsys, nil
    }

    // Strategy 2: Try OCI registry (fast, reliable)
    if !cfg.skipOCI {
        if fsys, err := ocidata.Pull(ctx, cfg.ociRef(version)); err == nil {
            cacheLocally(version, fsys)
            return fsys, nil
        }
    }

    // Strategy 3: Try custom GOPROXY (Go-native)
    if !cfg.skipProxy {
        if fsys, err := proxydata.Get(version); err == nil {
            cacheLocally(version, fsys)
            return fsys, nil
        }
    }

    // Strategy 4: Fallback to embedded data (if available)
    if cfg.allowEmbedded {
        return embeddedData(version)
    }

    return nil, fmt.Errorf("failed to load version %s", version)
}

// Usage:
fsys, _ := appledocs.LoadDocs("v17",
    appledocs.PreferOCI(),
    appledocs.AllowOffline(),
)
```

---

## Future Considerations

### WebAssembly Distribution

**Opportunity:** Compile docs loader to WASM for browser usage

```javascript
// Browser-based documentation viewer
import { loadDocs } from 'appledocs.wasm';

const docs = await loadDocs('v17');
const nsstring = docs.getSymbol('Foundation/NSString');
```

**Benefits:**
- In-browser documentation search
- No server-side processing
- Progressive web app capabilities

### Delta Updates

**Opportunity:** Download only changed files between versions

```go
// Update from v16 to v17, download only diffs
fsys, _ := appledocs.Update("v16", "v17")
```

**Implementation:**
- Use content addressing (SHA256 of each file)
- Maintain manifest of file hashes per version
- Download only files with changed hashes
- Typical SDK update: ~10-20% of files change → ~3-6MB download

**Savings:**
- Full v17 download: 30MB
- Delta v16→v17: ~5MB (83% reduction)

### Peer-to-Peer Distribution

**Opportunity:** Users share documentation with each other

```go
// Enable P2P sharing in corporate network
fsys, _ := appledocs.LoadDocs("v17",
    appledocs.EnableP2P("lan://"),
)
```

**Benefits:**
- Corporate environments save bandwidth
- Faster downloads within same network
- Resilient to internet outages

**Technologies:**
- BitTorrent/WebTorrent
- IPFS
- Custom P2P protocol

---

## Migration Path from Current State

### Step 1: Prepare Data (Week 1)

```bash
# Current: 314MB uncompressed
cd docs/tutorials/data

# Verify compression ratio
tar czf /tmp/docs.tar.gz documentation/
ls -lh /tmp/docs.tar.gz  # ~30MB

# Create metadata manifest
find documentation/ -type f -name "*.json" \
  -exec sh -c 'echo "{\"path\": \"$1\", \"sha256\": \"$(sha256sum $1 | cut -d" " -f1)\", \"size\": $(wc -c < $1)}" ' _ {} \; \
  | jq -s '.' > manifest.json
```

### Step 2: Publish to OCI (Week 1)

```bash
# Install crane
go install github.com/google/go-containerregistry/cmd/crane@latest

# Authenticate to GitHub Container Registry
echo $GITHUB_TOKEN | crane auth login ghcr.io -u USERNAME --password-stdin

# Create and push artifact
tar czf - documentation/ | crane append \
  -f - \
  -t ghcr.io/tmc/appledocs-data:v17 \
  scratch

crane push ghcr.io/tmc/appledocs-data:v17
```

### Step 3: Create OCI Loader Module (Week 2)

```bash
# Create new module
mkdir appledocs-oci
cd appledocs-oci

go mod init github.com/tmc/appledocs-oci
go get github.com/google/go-containerregistry/pkg/crane

# Implement loader (see code examples above)
# ...

# Publish
git tag v0.1.0
git push origin v0.1.0
```

### Step 4: Update Documentation (Week 2)

Add to main README.md:

```markdown
## Distribution Options

### Option 1: OCI Registry (Recommended)

```bash
go get github.com/tmc/appledocs-oci
```

```go
import (
    "github.com/tmc/appledocs"
    "github.com/tmc/appledocs-oci"
)

func main() {
    fsys, _ := ocidata.Pull(context.Background(),
        "ghcr.io/tmc/appledocs-data:v17")
    doc, _ := appledocs.GetSymbol(fsys, "Foundation/NSString")
}
```

Download size: ~30MB
Requires: Internet connection (cached after first download)
```

### Step 5: Add GOPROXY Support (Weeks 3-4)

See detailed implementation in Strategy 2 above.

### Step 6: Deprecate Old Patterns (Week 5)

Update `DISTRIBUTION.md`:

```markdown
## Deprecated Approaches

### ❌ Embedded Data Module (deprecated)

**Why deprecated:**
- Adds 200-400MB to binary size
- Slower compilation
- Harder to update independently

**Replaced by:** OCI artifacts and GOPROXY

### ❌ Manual Download (deprecated)

**Why deprecated:**
- Manual process
- No version management
- No automatic updates

**Replaced by:** Automated loaders
```

---

## Conclusion and Recommendations

### Primary Recommendation: OCI Artifacts

**Implement OCI artifacts first** for these reasons:

1. **Industry Standard**: OCI is the 2024+ standard for distributing large artifacts
2. **Free Infrastructure**: GitHub Container Registry is free and reliable
3. **Best Performance**: Global CDN, layer caching, parallel downloads
4. **Future Proof**: Growing ecosystem, increasing adoption
5. **Proven at Scale**: Used by major projects (ML models, WebAssembly, etc.)

**Timeline:** 2-3 weeks for full implementation

### Secondary Recommendation: Custom GOPROXY

**Add custom GOPROXY** as complementary approach:

1. **Go Native**: Works seamlessly with `go get` and standard tooling
2. **Low Cost**: Free tier options available (CloudFlare R2)
3. **Familiar**: Zero learning curve for Go developers
4. **Discoverable**: Shows up in pkg.go.dev

**Timeline:** 2-3 weeks after OCI is stable

### Future Exploration: SQLite

**Consider SQLite** for advanced use cases:

1. **Single File**: Entire dataset in one portable file
2. **Queryable**: SQL enables powerful cross-doc searches
3. **Efficient**: Excellent compression and fast random access
4. **Interesting**: Could enable new features (metadata queries, etc.)

**Timeline:** Experimental, evaluate after OCI+GOPROXY are stable

### Do Not Implement

1. **Git LFS**: Higher cost, worse performance than OCI
2. **Pure IPFS**: Too immature for production, gateway reliability issues
3. **BitTorrent**: No clear advantage over OCI, adds complexity

---

## Next Steps

1. **This Week:**
   - Review this document with stakeholders
   - Choose primary implementation (recommend: OCI)
   - Set up GitHub Container Registry
   - Create `appledocs-oci` repository

2. **Week 1-2:**
   - Implement OCI loader module
   - Publish v17 to ghcr.io
   - Write documentation and examples
   - Set up GitHub Actions automation

3. **Week 3-4:**
   - Evaluate OCI performance and adoption
   - Decide on secondary strategy (GOPROXY or SQLite)
   - Begin secondary implementation
   - Gather user feedback

4. **Month 2:**
   - Optimize based on real-world usage
   - Consider hybrid approaches
   - Explore advanced features (delta updates, etc.)
   - Publish blog post about the approach

---

## References

### OCI Artifacts
- [OCI Image Specification v1.1](https://opencontainers.org/posts/blog/2024-03-13-image-and-distribution-1-1/)
- [ORAS (OCI Registry as Storage)](https://oras.land/)
- [go-containerregistry](https://github.com/google/go-containerregistry)
- [Using OCI Artifacts for AI Models](https://www.docker.com/blog/oci-artifacts-for-ai-model-packaging/)

### Go Module Proxy
- [GOPROXY Protocol](https://go.dev/ref/mod#goproxy-protocol)
- [goproxy/goproxy](https://github.com/goproxy/goproxy)
- [Athens Proxy](https://docs.gomods.io/)

### Content-Addressed Storage
- [IPFS Documentation](https://docs.ipfs.tech/)
- [Content Addressing Explained](https://proto.school/content-addressing)
- [Restic Backup](https://restic.net/)

### CDN and Distribution
- [CloudFlare R2](https://www.cloudflare.com/products/r2/)
- [Amazon S3 + CloudFront](https://aws.amazon.com/blogs/networking-and-content-delivery/amazon-s3-amazon-cloudfront-a-match-made-in-the-cloud/)

### SQLite VFS
- [SQLite VFS](https://www.sqlite.org/vfs.html)
- [modernc.org/sqlite](https://pkg.go.dev/modernc.org/sqlite)
- [Go SQLite VFS Support](https://github.com/mattn/go-sqlite3/issues/968)

---

## Appendix: Cost Comparison

### Annual Cost Projections (10K downloads/month)

| Strategy | Infrastructure | Bandwidth | Storage | Total/Year |
|----------|----------------|-----------|---------|------------|
| **OCI (GitHub)** | $0 | $0 | $0 | **$0** |
| **GOPROXY+CloudFront** | $120 | $180 | $1 | **$301** |
| **GOPROXY+CloudFlare** | $60 | $0 | $0 | **$60** |
| **IPFS (Pinata)** | $0 | $0 | $0 | **$0** |
| **Git LFS (GitHub)** | $60 | $315 | $12 | **$387** |

### At Scale (100K downloads/month)

| Strategy | Total/Year |
|----------|------------|
| **OCI (GitHub)** | **$0** |
| **GOPROXY+CloudFlare** | **$120** (Cloud Run) |
| **IPFS (Pinata)** | **$0** (free tier) |
| **Git LFS (GitHub)** | **$3,870** |

**Clear Winner:** OCI artifacts with GitHub Container Registry

---

## Appendix: Performance Benchmarks (Projected)

### Download Time (100Mbps connection)

| Strategy | Time to First Byte | Full Download | Location |
|----------|-------------------|---------------|----------|
| OCI (ghcr.io) | 50ms | 3s | Global |
| GOPROXY+CloudFront | 30ms | 3s | Global |
| GOPROXY+CloudFlare | 20ms | 3s | Global |
| IPFS (gateway) | 500ms | 10-30s | Variable |
| Git LFS | 200ms | 8s | US |

### Deduplication Savings (v16 → v17 update)

| Strategy | Full Download | Delta Only | Savings |
|----------|--------------|------------|---------|
| OCI (layer sharing) | 30MB | 5MB | 83% |
| GOPROXY (full module) | 30MB | 30MB | 0% |
| IPFS (block sharing) | 30MB | 8MB | 73% |
| Git LFS | 30MB | 30MB | 0% |

**Note:** OCI layer deduplication provides significant savings for users downloading multiple versions

---

**Document Version:** 1.0
**Author:** Claude (Anthropic)
**Date:** October 2025
**Status:** Research Complete - Ready for Implementation
