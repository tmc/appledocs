# Multi-Version macOS API Binding Strategy

## Current State

The `appledocs.Platform` type already captures version information from Apple's docs:
- `IntroducedAt` (e.g., "15.4", "10.15")
- `DeprecatedAt`
- `Beta` flag
- `Deprecated` flag
- `Unavailable` flag

However, the binding generator (`cmd/generate-framework-bindings`) doesn't currently use this information.

## Design Goals

1. **Go-idiomatic**: Use standard Go patterns (build tags, package structure, doc comments)
2. **Backward compatible**: Existing code continues to work
3. **Type-safe**: Leverage compile-time checks where possible
4. **Simple**: Minimize complexity for users
5. **Flexible**: Support various version targeting strategies

## Recommended Approach: Annotated Single Package

Generate one package per framework with version annotations in godoc comments and optional runtime checks.

### Package Structure
```
generated/
└── frameworks/
    └── fskit/                 # Package fskit
        ├── doc.go            # Package docs with min version
        ├── types.gen.go      # Core types
        ├── loader.gen.go     # Framework loading
        └── functions.gen.go  # Functions with version annotations
```

### Generated Code Example

#### doc.go
```go
// Package fskit provides Go bindings for the FSKit framework.
//
// FSKit enables you to implement a file system that runs in user space.
//
// Minimum macOS version: 15.0
// Framework path: /System/Library/Frameworks/FSKit.framework/FSKit
//
// Generated from Apple documentation.
package fskit

const (
	// MinMacOSVersion is the minimum macOS version required for this framework
	MinMacOSVersion = "15.0"

	// FrameworkPath is the system path to the framework
	FrameworkPath = "/System/Library/Frameworks/FSKit.framework/FSKit"
)
```

#### functions.gen.go
```go
// FSUnaryFileSystemCreate creates a new unary file system.
//
// Availability:
//   - macOS 15.4+
//   - iOS 18.0+
//
// Status: Stable
func FSUnaryFileSystemCreate(...) FSUnaryFileSystemRef

// FSExperimentalFeature demonstrates a beta API.
//
// Availability:
//   - macOS 15.4+ (Beta)
//
// Status: Beta - API may change
func FSExperimentalFeature(...)

// FSDeprecatedFunction shows a deprecated API.
//
// Availability:
//   - macOS 15.0+ (Deprecated in 15.4)
//
// Status: Deprecated - Use FSUnaryFileSystemCreate instead
//
// Deprecated: Use FSUnaryFileSystemCreate
func FSDeprecatedFunction(...)
```

## Alternative Approaches (Not Recommended)

### Option 1: Build Tags
```go
//go:build darwin && macos15.4

package fskit
```

**Problems:**
- Requires exact Go version matching
- Difficult to maintain multiple versions
- Build tags don't support version ranges well

### Option 2: Version Subpackages
```
generated/frameworks/fskit/
├── v15_0/  # macOS 15.0 APIs
├── v15_4/  # macOS 15.4 APIs
└── v16_0/  # macOS 16.0 APIs
```

**Problems:**
- Package proliferation
- Import complexity (`fskit/v15_4`)
- Unclear which package to use
- Maintenance burden

### Option 3: Feature Flags
```go
const (
	SupportsFSUnaryFileSystem = runtime.GOOS == "darwin" && macosVersion >= 15.4
)
```

**Problems:**
- Runtime checks (slower)
- Requires version detection code
- No compile-time safety

## Implementation Plan

### Phase 1: Enhance Binding Generator

Modify `cmd/generate-framework-bindings/main.go` to:

1. **Parse platform metadata**:
```go
type ParsedFunction struct {
	Name         string
	ReturnType   string
	Parameters   []Parameter
	Comment      string
	Availability PlatformAvailability  // NEW
}

type PlatformAvailability struct {
	MacOS      VersionInfo
	IOS        VersionInfo
	WatchOS    VersionInfo
	Deprecated bool
	Beta       bool
}

type VersionInfo struct {
	IntroducedAt string
	DeprecatedAt string
}
```

2. **Extract version from JSON**:
```go
func processJSONFile(path string) (*ParsedFunction, error) {
	// ... existing code ...

	// NEW: Extract platform info
	if doc.Metadata.Platforms != nil {
		fn.Availability = extractAvailability(doc.Metadata.Platforms)
	}

	return fn, nil
}
```

3. **Generate enhanced comments**:
```go
func generateFunctionComment(fn *ParsedFunction) string {
	var buf strings.Builder

	// Function description
	fmt.Fprintf(&buf, "// %s %s\n", fn.Name, fn.Comment)
	fmt.Fprintf(&buf, "//\n")

	// Availability section
	if !fn.Availability.IsEmpty() {
		fmt.Fprintf(&buf, "// Availability:\n")

		if fn.Availability.MacOS.IntroducedAt != "" {
			status := ""
			if fn.Availability.Beta {
				status = " (Beta)"
			}
			if fn.Availability.MacOS.DeprecatedAt != "" {
				status = fmt.Sprintf(" (Deprecated in %s)", fn.Availability.MacOS.DeprecatedAt)
			}
			fmt.Fprintf(&buf, "//   - macOS %s+%s\n", fn.Availability.MacOS.IntroducedAt, status)
		}

		// ... similar for iOS, watchOS ...
	}

	// Deprecation warning
	if fn.Availability.Deprecated {
		fmt.Fprintf(&buf, "//\n")
		fmt.Fprintf(&buf, "// Deprecated: %s\n", getDeprecationMessage(fn))
	}

	return buf.String()
}
```

4. **Generate package documentation**:
```go
func generateDocFile(outputDir, framework string, functions []*ParsedFunction) {
	minVersion := findMinimumVersion(functions)

	doc := fmt.Sprintf(`// Package %s provides Go bindings for the %s framework.
//
// Minimum macOS version: %s
// Framework path: /System/Library/Frameworks/%s.framework/%s
//
// Generated from Apple documentation.
package %s

const (
	MinMacOSVersion = "%s"
	FrameworkPath   = "/System/Library/Frameworks/%s.framework/%s"
)
`,
		strings.ToLower(framework),
		framework,
		minVersion,
		framework, framework,
		strings.ToLower(framework),
		minVersion,
		framework, framework,
	)

	os.WriteFile(filepath.Join(outputDir, "doc.go"), []byte(doc), 0644)
}
```

### Phase 2: Optional Runtime Version Checking

For users who want runtime checks:

```go
// version.gen.go
package fskit

import (
	"fmt"
	"os"
	"syscall"
)

// CheckVersion verifies the current macOS version meets minimum requirements.
// Returns an error if the version is too old.
func CheckVersion() error {
	ver, err := macOSVersion()
	if err != nil {
		return fmt.Errorf("failed to get macOS version: %w", err)
	}

	if compareVersion(ver, MinMacOSVersion) < 0 {
		return fmt.Errorf("FSKit requires macOS %s or later (found %s)", MinMacOSVersion, ver)
	}

	return nil
}

// macOSVersion returns the current macOS version string
func macOSVersion() (string, error) {
	var utsname syscall.Utsname
	if err := syscall.Uname(&utsname); err != nil {
		return "", err
	}

	// Parse Darwin kernel version to macOS version
	// Darwin 24.x = macOS 15.x
	// ...
	return parseKernelVersion(utsname.Release[:]), nil
}
```

### Phase 3: Enhanced Example Usage

```go
package main

import (
	"log"

	"github.com/tmc/appledocs/generated/frameworks/fskit"
)

func main() {
	// Optional: Check version at startup
	if err := fskit.CheckVersion(); err != nil {
		log.Fatalf("Version check failed: %v", err)
	}

	// Use FSKit normally - OS handles symbol resolution
	fs := fskit.FSUnaryFileSystemCreate(...)
}
```

## Benefits of This Approach

1. **Simple imports**: Just `import "framework/fskit"`
2. **Clear documentation**: Version info in godoc
3. **IDE-friendly**: Autocomplete shows version requirements
4. **Future-proof**: Easy to add new APIs
5. **Backward compatible**: Existing code works
6. **Minimal overhead**: No runtime checks unless explicitly called
7. **Standard Go**: Uses only built-in Go features

## Migration Path

1. Generate new bindings with version annotations
2. Existing code continues to work (no breaking changes)
3. Users can optionally add `CheckVersion()` calls
4. Documentation improvements happen automatically

## Testing Strategy

```go
// Test that version annotations are generated
func TestVersionAnnotations(t *testing.T) {
	// Generate bindings for FSKit
	// Verify doc.go contains MinMacOSVersion
	// Verify functions have Availability comments
}

// Test version parsing
func TestVersionExtraction(t *testing.T) {
	doc := &Document{
		Metadata: Metadata{
			Platforms: []Platform{
				{Name: "macOS", IntroducedAt: "15.4", Beta: false},
			},
		},
	}

	avail := extractAvailability(doc.Metadata.Platforms)
	assert.Equal(t, "15.4", avail.MacOS.IntroducedAt)
}
```

## Future Enhancements

1. **Version-specific types**: Generate different struct definitions for different OS versions
2. **API evolution tracking**: Show how APIs changed across versions
3. **Migration guides**: Auto-generate upgrade guides
4. **Compatibility matrix**: Table showing what works on which OS version

## Conclusion

This approach provides version awareness without sacrificing simplicity. It's Go-idiomatic, uses standard documentation practices, and gives users clear information about API availability while letting the OS handle the actual symbol resolution.
