# Framework Examples

This directory contains example programs demonstrating the usage of generated Apple framework bindings.

## Overview

The generated bindings provide documentation and structure for Apple frameworks, but require manual implementation using `purego` to actually call framework functions.

## Example Structure

Each framework has its own example directory:

- `examples/foundation/` - Foundation framework
- `examples/coregraphics/` - CoreGraphics framework
- `examples/security/` - Security framework
- etc.

## Generated Bindings vs. Working Examples

### Generated Bindings (in `generated/frameworks/`)

The code in `generated/frameworks/` contains:
- Documentation extracted from Apple's documentation
- Function signatures as comments
- Type definitions
- Framework loading infrastructure

These are **not ready-to-use** - they serve as documentation and require manual binding.

### Working Examples (in `examples/`)

For actual working code, see these examples:

**Pure Purego Examples** (ports from darwinkit):
- `examples/clickme-purego/` - **Simple button example** (recommended starting point for UI)
- `examples/helloworld-purego/` - WebView example with custom delegate
- `examples/foundation-working/` - **Complete working Foundation example with E2E tests** (recommended starting point for Foundation)

**Other Examples**:
- `examples/drawing-generated-bindings/` - Complex CoreGraphics drawing example
- `examples/security-keychain/` - Security framework example

## Creating New Examples

To create a working example:

1. Import the generated bindings for documentation:
   ```go
   import _ "github.com/tmc/appledocs/generated/frameworks/coregraphics"
   ```

2. Manually bind functions using purego:
   ```go
   var CGContextSetRGBFillColor func(c CGContextRef, r, g, b, a CGFloat)
   purego.RegisterLibFunc(&CGContextSetRGBFillColor, cgLib, "CGContextSetRGBFillColor")
   ```

3. Call the bound functions:
   ```go
   CGContextSetRGBFillColor(ctx, 1.0, 0.0, 0.0, 1.0)
   ```

## Scripts

- `scripts/generate-examples.sh` - Generate example template programs
- `scripts/run-example-tests.sh` - Test that examples compile and run
- `scripts/tidy-examples.sh` - Run `go mod tidy` on all examples

## Testing

Basic tests verify that:
1. Generated bindings can be imported
2. Framework loading succeeds
3. Basic program structure works

Full functional tests require manual implementation of framework calls.

## Future Work

- Auto-generate working bindings (not just documentation)
- Comprehensive test suite for common framework operations
- CI/CD integration for binding validation
