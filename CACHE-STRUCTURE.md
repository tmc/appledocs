# AppDocs Cache Structure

## Overview

The appledocs cache stores downloaded Apple documentation JSON files to minimize network requests and enable offline access.

## Cache Location

Default: `~/.appledocs/cache/`

## Cache Structure Changes (October 2025)

### Problem

The original cache structure mapped URLs directly to file paths:
```
URL:  https://developer.apple.com/tutorials/data/documentation/foundation/nsstring.json
Path: ~/.appledocs/cache/developer.apple.com/tutorials/data/documentation/foundation/nsstring.json
```

This caused file/directory conflicts when a URL had both:
1. **Content** (the base type documentation)
2. **Children** (nested methods, properties, etc.)

Example conflict:
```
# Both needed to exist:
/foundation/attributedstringprotocol        # File (base type docs)
/foundation/attributedstringprotocol/characters/  # Directory (child path)
```

This resulted in errors like:
```
Error: mkdir /path/attributedstringprotocol: not a directory
```

### Solution

**New cache structure (2025-10-22)**: Always use directory structure with `index.json` files.

```
URL:  https://developer.apple.com/tutorials/data/documentation/foundation/nsstring.json
Path: ~/.appledocs/cache/developer.apple.com/tutorials/data/documentation/foundation/nsstring/index.json
                                                                                           ^^^^^^^^^^
                                                                        (directory)    (file)
```

This allows both the base type and children to coexist:
```
/foundation/attributedstringprotocol/index.json        # Base type docs
/foundation/attributedstringprotocol/characters/index.json  # Child docs
```

### Implementation Details

1. **For documentation URLs** (`/documentation/` or `.json` paths):
   - Strip `.json` extension from directory name
   - Create directory with that name
   - Store content in `index.json` within that directory

2. **For non-documentation URLs**:
   - Use the path as filename (legacy behavior)

3. **ETag files**: Stored as `index.json.etag` alongside `index.json`

4. **Query parameters** (except `language`):
   - Encoded in filename: `index.{encoded-query}.json`

### Code Reference

Implementation: `internal/crawler/methods.go` lines 24-67

```go
// For documentation URLs, always use directory structure with index.json
// This prevents file/directory collisions when a path has both content and children
if strings.Contains(parsed.Path, "/documentation/") || strings.HasSuffix(parsed.Path, ".json") {
    cacheDir = basePath
    if strings.HasSuffix(cacheDir, ".json") {
        cacheDir = strings.TrimSuffix(cacheDir, ".json")
    }
    cachePath = filepath.Join(cacheDir, "index.json")
}
```

## HTML vs JSON Cache Issue

### Problem

Old cache contained HTML files instead of JSON:
```bash
$ file ~/.appledocs/cache/developer.apple.com/documentation/foundation/nsextensioncontext
HTML document text, ASCII text
```

**Count**: ~2,410 HTML files found in Foundation cache alone.

### Root Cause

Apple's documentation has two URL patterns:
- `/documentation/foundation/nsextensioncontext` → Returns HTML (web page)
- `/tutorials/data/documentation/foundation/nsextensioncontext.json` → Returns JSON (API data)

The old crawler cached from the wrong URL pattern.

### Solution

The `ResolveURL` function (in `internal/crawler/utils.go`) now properly converts:
```go
if strings.HasPrefix(relative, "documentation/") {
    return base + "/tutorials/data/" + relative
}
```

This ensures JSON URLs are fetched, not HTML pages.

### Cleanup

A cleanup script is available to remove stale HTML files:
```bash
./cleanup-html-cache.sh ~/.appledocs/cache
```

The script:
1. Finds all HTML files in cache
2. Removes them and their `.etag` companions
3. Lets the crawler re-fetch as JSON on next run

## Migration Path

### From Old to New Cache

1. **Automatic**: The new crawler creates `index.json` files in directories, old files remain untouched
2. **Coexistence**: Both old files and new directories can exist temporarily
3. **Cleanup**: Run cleanup script or manually delete old structure

### Automatic Cleanup with --prune Flag

The easiest way to clean up old cache files is to use the `--prune` flag:

```bash
# Prune old cache files and crawl
appledocs crawl Foundation --prune

# Just prune without crawling much (use short timeout)
appledocs crawl --prune --max-time 1
```

The `--prune` flag automatically removes:
1. **HTML files** (should be JSON)
2. **Old structure files** (when `index.json` exists in the same directory)
3. **Language variant duplicates** (`.language%3Dobjc`, `.language%3Dswift`)

### Manual Cache Refresh

To force the crawler to use the new structure for everything:

```bash
# Delete all .etag files (forces refetch)
find ~/.appledocs/cache -name "*.etag" -type f -delete

# Run crawler (it will create new index.json structure)
appledocs crawl Foundation
```

### Verify New Structure

```bash
# Check for index.json files
find ~/.appledocs/cache -name "index.json" | head -10

# Verify they're JSON not HTML
file ~/.appledocs/cache/developer.apple.com/tutorials/data/documentation/foundation/nsstring/index.json
```

## Cache Performance

The new structure has no performance impact:
- Same number of filesystem operations
- Slightly longer paths (adds `/index.json` to each)
- Eliminates file/directory conflict errors

## Benefits

1. ✅ **No More Conflicts**: Files and directories can coexist
2. ✅ **Correct Content**: Only JSON cached, not HTML
3. ✅ **Future-Proof**: Handles arbitrary nesting depths
4. ✅ **Clear Structure**: `index.json` pattern is intuitive
5. ✅ **Backward Compatible**: Old cache files still work until refreshed

## Statistics (Foundation Crawl - October 22, 2025)

**Before fix**:
- Errors: 370+ file/directory conflicts
- HTML files: 2,410
- Cache hit rate: 93%

**After fix**:
- Errors: 0 file/directory conflicts
- HTML files: 0 (being cleaned up)
- Cache hit rate: ~64% (due to structure change forcing refetch)

## Future Considerations

1. **Cache Version**: Consider adding version marker to detect old cache
2. **Automatic Migration**: Could auto-convert old structure on crawler startup
3. **Compression**: Could compress `index.json` files to save space
4. **Deduplication**: Many files have duplicate content (Swift vs ObjC variants)
