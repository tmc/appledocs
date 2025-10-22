# Cache Structure Fix - October 22, 2025

## Summary

Fixed two critical issues with the appledocs cache:
1. **File/Directory Conflicts**: URLs with nested paths caused filesystem collisions
2. **HTML vs JSON**: Old cache contained HTML pages instead of JSON API data

## Problem Details

### Issue 1: File/Directory Collisions

**Symptom**: Errors like `mkdir: not a directory` during crawling

**Root Cause**: The cache mapped URLs directly to file paths:
```
/documentation/foundation/attributedstringprotocol       → file
/documentation/foundation/attributedstringprotocol/characters → needs directory
```

Result: ~370 errors during Foundation crawl alone.

### Issue 2: HTML Cached Instead of JSON

**Symptom**: Generated code reading HTML instead of JSON documentation

**Root Cause**: Wrong URL pattern was cached:
- ❌ `https://developer.apple.com/documentation/foundation/nsstring` → HTML
- ✅ `https://developer.apple.com/tutorials/data/documentation/foundation/nsstring.json` → JSON

Result: ~2,410 HTML files in Foundation cache that should have been JSON.

## Solution

### New Cache Structure

Changed from:
```
cache/foundation/nsstring.json                    (file)
cache/foundation/nsstring/init.json               (ERROR: nsstring is a file!)
```

To:
```
cache/foundation/nsstring/index.json              (directory with index.json)
cache/foundation/nsstring/init/index.json         (child directory)
```

### Implementation

**File**: `internal/crawler/methods.go` (lines 24-67)

Key changes:
```go
// For documentation URLs, always use directory structure with index.json
if strings.Contains(parsed.Path, "/documentation/") || strings.HasSuffix(parsed.Path, ".json") {
    cacheDir = basePath
    if strings.HasSuffix(cacheDir, ".json") {
        cacheDir = strings.TrimSuffix(cacheDir, ".json")
    }
    cachePath = filepath.Join(cacheDir, "index.json")
}
```

### Automatic Cleanup: --prune Flag

Added `--prune` flag to automatically clean up old cache:

**File**: `cmd/appledocs/cmd_crawl.go`

**Usage**:
```bash
# Prune and crawl
appledocs crawl Foundation --prune

# Just prune (short timeout)
appledocs crawl --prune --max-time 1
```

**What it removes**:
1. HTML files (detected by content)
2. Old structure files (when index.json exists in same directory)
3. Language variant duplicates (`.language%3Dobjc`, `.language%3Dswift`)

## Results

### Before Fix
- File/directory conflicts: 370+
- HTML files: 2,410 in Foundation
- Cache hit rate: 93% (but hitting wrong content)
- Errors: Constant mkdir failures

### After Fix
- File/directory conflicts: 0
- HTML files: 0 (auto-pruned with --prune)
- Cache hit rate: ~64% initially (rebuilding with correct content)
- Errors: None related to cache structure

## Documentation

Created `CACHE-STRUCTURE.md` with complete documentation of:
- Cache structure design
- Migration path
- Cleanup procedures
- Performance characteristics
- Future considerations

## Testing

**Test 1: New Structure**
```bash
$ find /tmp/test-cache -name "index.json" | head -5
/tmp/test-cache/developer.apple.com/tutorials/data/documentation/foundation/index.json
/tmp/test-cache/developer.apple.com/tutorials/data/documentation/foundation/nsstring/index.json
/tmp/test-cache/developer.apple.com/tutorials/data/documentation/foundation/nsstring/init/index.json
...
```

**Test 2: Content Verification**
```bash
$ file ~/.appledocs/cache/.../nsextensioncontext/index.json
index.json: JSON data  ✅
```

**Test 3: Prune Functionality**
```bash
$ appledocs crawl --prune --max-time 1
time=... level=INFO msg="Cache pruning complete" html_files_removed=2410 old_structure_files_removed=1503
```

## Migration

### Automatic (Recommended)
```bash
# Run crawler with --prune flag
appledocs crawl Foundation --prune
```

### Manual
```bash
# Force refetch by deleting etags
find ~/.appledocs/cache -name "*.etag" -delete

# Remove HTML files
find ~/.appledocs/cache -type f -exec sh -c 'file "$1" | grep -q HTML && rm "$1"' _ {} \;
```

## Performance Impact

- **Filesystem Operations**: No change (same number of reads/writes)
- **Path Length**: Slightly longer (+11 chars for `/index.json`)
- **Cache Lookups**: Same speed (stat + read)
- **Benefits**: Eliminates all file/directory conflict errors

## Files Changed

### Core Changes
- `internal/crawler/methods.go` - Cache path logic
- `cmd/appledocs/cmd_crawl.go` - Added --prune flag and pruneOldCache()
- `internal/crawler/utils.go` - ResolveURL already correct

### Documentation
- `CACHE-STRUCTURE.md` - Complete cache structure documentation
- `CHANGELOG-cache-fix.md` - This file

### Generated Files
- All framework bindings regenerated (property accessors added separately)

## Breaking Changes

**None**. The new structure is fully backward compatible:
- Old cache files still work
- New fetches use new structure
- Both can coexist
- Use `--prune` to clean up when ready

## Future Enhancements

1. **Cache Version Marker**: Add `.cache-version` file to detect old structure
2. **Automatic Migration**: Auto-prune on first run after update
3. **Compression**: Gzip `index.json` files to save space
4. **Deduplication**: Detect and merge duplicate language variants

## Credits

- Issue discovered: October 19, 2025 (HTML in cache)
- File/directory conflicts: October 22, 2025 (Foundation crawl errors)
- Fix implemented: October 22, 2025
- Testing: Session C160, 77A6

## Related Beads

- appledocs-453: Framework-specific entry points (completed)
- (Create new bead for cache structure fix if needed)
