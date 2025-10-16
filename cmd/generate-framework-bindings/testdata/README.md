# Test Organization

This directory contains scripttest-based integration tests for the framework bindings generator.

## Test Categories

### Baseline Tests (`testdata/*.txt`)

These tests **MUST always pass** and verify the current functionality:

- ✅ **basic_generation.txt** - End-to-end CoreGraphics generation
- ✅ **error_handling.txt** - Invalid inputs and error cases
- ✅ **framework_options.txt** - Different frameworks (Foundation, CoreGraphics)
- ✅ **incremental_generation.txt** - Idempotency and reproducibility
- ✅ **output_validation.txt** - Code structure and compilation
- ✅ **parser_validation.txt** - Function signature handling

**Run**: `go test ./cmd/generate-framework-bindings`

All baseline tests should pass in CI and before any commit.

### Aspirational Tests (`testdata/aspirational/*.txt`)

These tests **document future goals** and are expected to fail until implemented:

- ❌ **api_compatibility_nsapplication.txt** - NSApplication API compatibility with darwinkit
- ❌ **appkit_filtered_generation.txt** - AppKit generation with filtering

**Run**: `go test ./cmd/generate-framework-bindings -aspirational`

Aspirational tests define the target API and serve as:
- **Specifications** for what needs to be built
- **Acceptance criteria** for new features
- **Documentation** of the end goal

## Usage

### Run all passing tests (CI-safe):
```bash
go test -v ./cmd/generate-framework-bindings
```

### Run aspirational tests (expected to fail):
```bash
go test -v ./cmd/generate-framework-bindings -aspirational
```

## Adding New Tests

### Baseline Test
If the generator already supports a feature, add a test to `testdata/`:
```bash
testdata/my_new_feature.txt
```

### Aspirational Test
If defining a future feature or API goal, add to `testdata/aspirational/`:
```bash
testdata/aspirational/my_future_feature.txt
```

## Scripttest Format

Tests use [rsc.io/script/scripttest](https://pkg.go.dev/rsc.io/script/scripttest) format:

```bash
# Comments describe what the test does
env INPUT_DIR=$HOME/.appledocs/cache/...
exec generate-framework-bindings -framework CoreGraphics -output test_output
stderr 'Generated bindings'
exists test_output/coregraphics/functions.gen.go
grep 'package coregraphics' test_output/coregraphics/functions.gen.go
```

## Test Status

**Baseline**: 6 tests, all passing ✅
**Aspirational**: 2 tests, both failing as expected ❌

Last updated: 2025-10-10
