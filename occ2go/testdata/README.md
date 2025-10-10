# occ2go Test Organization

This directory contains scripttest-based integration tests for the `occ2go` command-line tool.

## Quick Reference

```bash
# Run all baseline tests (default, CI-safe)
go test ./occ2go

# Run aspirational tests (expected to document future goals)
go test ./occ2go -aspirational

# Verbose output
go test -v ./occ2go

# Run specific test
go test -v ./occ2go -run TestScripts/parse_class
```

## Test Organization

### Baseline Tests (`testdata/*.txt`)

Located in `testdata/*.txt` - these verify current functionality and **must pass** before commits.

| Test | Purpose | Status |
|------|---------|--------|
| `parse_real_function` | Parse CGContextMoveToPoint function | ✅ PASS |
| `parse_function_complex` | Parse CGColorSpaceRelease (void return, single param) | ✅ PASS |
| `parse_class` | Parse NSApplication class | ✅ PASS |
| `error_nonexistent_file` | Error handling for missing files | ✅ PASS |

### Aspirational Tests (`testdata/aspirational/*.txt`)

Located in `testdata/aspirational/*.txt` - these define future goals and document what needs to be implemented.

| Test | Goal | Status |
|------|------|--------|
| `parse_enum` | Parse CGBlendMode enum | ❌ Not yet implemented |
| `parse_typedef` | Parse CGBitmapContextReleaseDataCallback typedef | ❌ Not yet implemented |

## Test Philosophy

**Baseline tests** validate what works now:
- They must always pass
- They run in CI
- Breaking them requires investigation
- They document current capabilities

**Aspirational tests** define what we're building toward:
- They're expected to fail (for now)
- They skip by default (`-aspirational` to run)
- They serve as executable specifications
- They document unsupported symbol types
- Passing them is the definition of "done"

## Adding Tests

### When to Add a Baseline Test

- Feature is implemented and working
- You want to prevent regressions
- Verifying a bug fix

**Example:**
```bash
# testdata/my_feature.txt
env INPUT_DIR=$HOME/.appledocs/cache/developer.apple.com/tutorials/data/documentation
exec occ2go $INPUT_DIR/Framework/Symbol.json
stdout 'expected output'
```

### When to Add an Aspirational Test

- Defining a new feature's requirements
- Setting acceptance criteria
- Documenting a goal before implementation

**Example:**
```bash
# testdata/aspirational/future_feature.txt
# This test defines the goal for XYZ feature
# Currently fails because: [reason]
# Will pass when: [criteria]
! exec occ2go $INPUT_DIR/Framework/UnsupportedSymbol.json
stderr 'unsupported symbol type'
```

## Scripttest Syntax

Tests use [rsc.io/script/scripttest](https://pkg.go.dev/rsc.io/script/scripttest):

| Command | Purpose | Example |
|---------|---------|---------|
| `exec` | Run command, expect success | `exec occ2go file.json` |
| `! exec` | Run command, expect failure | `! exec occ2go nonexistent.json` |
| `stdout` | Check stdout contains pattern | `stdout 'func CGContextMoveToPoint'` |
| `stderr` | Check stderr contains pattern | `stderr 'unsupported symbol type'` |
| `env` | Set environment variable | `env INPUT_DIR=/path/to/cache` |

## CI Integration

The default `go test` run is CI-safe:
- Only runs baseline tests (aspirational skipped)
- All tests should pass
- Fast execution (~0.3 seconds)

```yaml
# .github/workflows/test.yml
- name: Test occ2go
  run: go test ./occ2go
```

## Development Workflow

1. **Writing new feature**:
   - Write aspirational test first (TDD style)
   - Implement feature
   - Test passes → move to baseline

2. **Fixing bug**:
   - Add baseline test that reproduces bug (should fail)
   - Fix bug
   - Test passes → commit both

3. **Refactoring**:
   - Run baseline tests frequently
   - They catch regressions immediately
   - All tests passing = refactor is safe

## Test Status

Last updated: 2025-10-10

**Baseline**: 4/4 passing ✅
**Aspirational**: 2 tests documenting future enum/typedef support ⏭️
**Total**: All tests behaving as expected ✅

## Examples

### Successful Parse (Baseline)

```bash
$ occ2go ~/.appledocs/cache/.../CoreGraphics/CGContextMoveToPoint.json
// Begins a new subpath at the point you specify.
//
// [Full Topic]: doc://com.apple.coregraphics/documentation/CoreGraphics/CGContextMoveToPoint
func CGContextMoveToPoint(c CGContextRef, x CGFloat, y CGFloat) {
	// TODO: Implementation
}
```

### Expected Failure (Aspirational)

```bash
$ occ2go ~/.appledocs/cache/.../CoreGraphics/CGBlendMode.json
time=... level=ERROR msg="Failed to parse document" error="unsupported symbol type: c:@E@CGBlendMode"
```

This is expected! The aspirational test documents that enums aren't supported yet.
