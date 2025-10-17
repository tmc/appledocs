# AppKit Examples

This directory contains various examples demonstrating different approaches to using AppKit from Go.

## Examples

### helloworld-generated-bindings
Original implementation using generated bindings with minimal abstractions.
- **Lines:** 249
- **Raw objc calls:** 11
- **Style:** Explicit, verbose
- **Status:** ✅ Working

### helloworld-darwinkit-style
Darwinkit-style API using only generated bindings (no generated code modifications).
- **Lines:** 140
- **Raw objc calls:** 8
- **Style:** Clean, darwinkit-compatible
- **Status:** ✅ Working
- **Compatibility:** 85% with darwinkit

### helloworld-darwinkit
Reference implementation using actual darwinkit library.
- **Lines:** ~90
- **Raw objc calls:** 0
- **Style:** Idiomatic darwinkit
- **Status:** ⚠️ Darwinkit has build issues currently

## Comparison

See [COMPARISON_SUMMARY.md](./COMPARISON_SUMMARY.md) for detailed API comparison.

## Running Examples

```bash
# Generated bindings (original)
cd helloworld-generated-bindings && go run .

# Generated bindings (darwinkit-style)
cd helloworld-darwinkit-style && go run .

# Darwinkit (when working)
cd helloworld-darwinkit && go run .
```

## Key Findings

1. **Our generated bindings work well** - 85% darwinkit API compatibility
2. **Type-safe methods available** - Most common operations have typed methods
3. **Minor gaps** - Missing some convenience functions and constants
4. **Production ready** - Usable today for real applications

See [../DARWINKIT_COMPARISON.md](../DARWINKIT_COMPARISON.md) for full analysis.
