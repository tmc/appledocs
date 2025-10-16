# Testing Guide

## Quick Reference

```bash
# Run all passing tests (default, CI-safe)
go test ./cmd/generate-framework-bindings

# Run aspirational tests (expected to fail, documents goals)
go test ./cmd/generate-framework-bindings -aspirational

# Verbose output
go test -v ./cmd/generate-framework-bindings

# Run specific test
go test -v ./cmd/generate-framework-bindings -run TestScripts/basic_generation
```

## Test Organization

### Baseline Tests (Always Pass ✅)
Located in `testdata/*.txt` - these verify current functionality and **must pass** before commits.

| Test | Purpose | Status |
|------|---------|--------|
| basic_generation | End-to-end CoreGraphics generation | ✅ PASS |
| error_handling | Invalid inputs and error cases | ✅ PASS |
| framework_options | Different frameworks support | ✅ PASS |
| incremental_generation | Idempotency verification | ✅ PASS |
| output_validation | Code structure validation | ✅ PASS |
| parser_validation | Function signature parsing | ✅ PASS |

### Aspirational Tests (Expected to Fail ❌)
Located in `testdata/aspirational/*.txt` - these define future goals and API targets.

| Test | Goal | Status |
|------|------|--------|
| api_compatibility_nsapplication | NSApplication API parity with darwinkit | ❌ Not yet implemented |
| appkit_filtered_generation | AppKit generation with filtering | ❌ In development |

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
- Passing them is the definition of "done"

## Adding Tests

### When to Add a Baseline Test
- Feature is implemented and working
- You want to prevent regressions
- Verifying a bug fix

**Example:**
```bash
# testdata/my_feature.txt
env INPUT_DIR=$HOME/.appledocs/cache/...
exec generate-framework-bindings -framework CoreGraphics -output test_out
exists test_out/coregraphics/functions.gen.go
grep 'MyNewFeature' test_out/coregraphics/functions.gen.go
```

### When to Add an Aspirational Test
- Defining a new feature's API
- Setting acceptance criteria
- Documenting a goal before implementation

**Example:**
```bash
# testdata/aspirational/future_feature.txt
# This test defines the goal for XYZ feature
# Currently fails because: [reason]
# Will pass when: [criteria]
exec generate-framework-bindings -new-flag
grep 'ExpectedOutput' ...
```

## Scripttest Syntax

Tests use [rsc.io/script/scripttest](https://pkg.go.dev/rsc.io/script/scripttest):

| Command | Purpose | Example |
|---------|---------|---------|
| `exec` | Run command, expect success | `exec generate-framework-bindings -framework CoreGraphics` |
| `! exec` | Run command, expect failure | `! exec generate-framework-bindings -invalid-arg` |
| `exists` | Check file exists | `exists output/functions.gen.go` |
| `grep` | Check file contains pattern | `grep 'package coregraphics' output/doc.go` |
| `! grep` | Check file doesn't contain | `! grep 'Error' output/log.txt` |
| `stderr` | Check stderr contains | `stderr 'Generated bindings'` |
| `stdout` | Check stdout contains | `stdout 'Success'` |
| `env` | Set environment variable | `env INPUT_DIR=/path/to/cache` |

## CI Integration

The default `go test` run is CI-safe:
- Only runs baseline tests (aspirational skipped)
- All tests should pass
- Fast execution (~2 seconds)

```yaml
# .github/workflows/test.yml
- name: Test
  run: go test ./cmd/generate-framework-bindings
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

**Baseline**: 6/6 passing ✅  
**Aspirational**: 0/2 passing ❌ (by design)  
**Total**: All tests behaving as expected ✅
