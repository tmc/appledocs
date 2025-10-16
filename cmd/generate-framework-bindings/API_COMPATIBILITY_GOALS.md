# API Compatibility Goals for NSApplication

## Overview

This document defines the goal for achieving API compatibility with darwinkit's NSApplication bindings.

## Current State

The `generate-framework-bindings` tool currently:
- ✅ Generates C function bindings using purego
- ✅ Generates type definitions for CoreGraphics (CGRect, CGPoint, etc.)
- ❌ Does NOT generate Objective-C class bindings
- ❌ Does NOT generate method wrappers for classes

## Target API (darwinkit)

Darwinkit provides a Go-friendly API for NSApplication:

```go
type Application struct {
    Responder
}

// Class methods
func Application_SharedApplication() Application
func NewApplication() Application
func ApplicationFrom(ptr unsafe.Pointer) Application

// Instance methods (~100+ methods)
func (a_ Application) Run()
func (a_ Application) Terminate(sender objc.IObject)
func (a_ Application) ActivateIgnoringOtherApps(flag bool)
func (a_ Application) Delegate() ApplicationDelegateObject
func (a_ Application) SetDelegate(value PApplicationDelegate)
func (a_ Application) MainWindow() Window
func (a_ Application) KeyWindow() Window
func (a_ Application) Windows() []Window
func (a_ Application) Hide(sender objc.IObject)
func (a_ Application) Unhide(sender objc.IObject)
// ... ~90 more methods
```

## What Needs to be Implemented

### 1. Class Type Generation

Generate Go struct types for Objective-C classes:
- Parse class hierarchy from Apple documentation
- Generate struct with embedded parent type
- Handle class inheritance properly

### 2. Method Parsing

Parse Objective-C methods from Apple documentation:
- Instance methods: `- (void)run`
- Class methods: `+ (NSApplication *)sharedApplication`
- Method parameters and return types
- Selector names

### 3. Method Wrapper Generation

Generate Go method wrappers that call into objc runtime:
- Convert Go types to Objective-C types
- Handle method dispatch via objc_msgSend
- Wrap return values
- Handle nil receivers safely

### 4. Constructor Methods

Generate constructor functions:
- `NewApplication()` - alloc/init pattern
- `Application_SharedApplication()` - class method wrappers
- `ApplicationFrom(ptr)` - wrap existing pointer

## Test Suite

The scripttest `api_compatibility_nsapplication.txt` verifies API compatibility:

**Current Status**: ❌ FAILING - classes.gen.go not generated at all

**Failure Point**: The generator finds 0 classes in AppKit, so no classes.gen.go file is created.

**Root Cause**: The current generator only parses functions, not classes. The occ2go.ParseDocument likely returns nil for classes.

## Next Steps

1. **Debug class parsing**: Why does `occ2go.ParseDocument` return nil for NSApplication?
   - Check if the AppKit cache contains class documentation
   - Verify the parser recognizes `"kind": "class"` in JSON

2. **Implement class generation**: Once classes are parsed:
   - Update templates to generate proper Go struct types
   - Generate method wrappers using objc runtime
   - Handle inheritance and method overrides

3. **Pass the test**: The test will pass when:
   ```bash
   RUN_SCRIPTTEST=1 go test ./cmd/generate-framework-bindings -run TestScripts/api_compatibility_nsapplication
   ```
   succeeds with all assertions passing.

## References

- Darwinkit Application API: https://pkg.go.dev/github.com/progrium/darwinkit/macos/appkit#Application
- Apple NSApplication docs: https://developer.apple.com/documentation/appkit/nsapplication
- Test file: `cmd/generate-framework-bindings/testdata/api_compatibility_nsapplication.txt`
