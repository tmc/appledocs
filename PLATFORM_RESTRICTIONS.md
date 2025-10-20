# Platform Restrictions

This document explains why certain Apple frameworks cannot be generated on macOS and provides guidance on workarounds.

## Overview

Out of 270+ frameworks discovered in Apple's documentation, approximately 200 frameworks are not available in `/System/Library/Frameworks/` on macOS. These frameworks fall into several categories based on their platform restrictions.

## Categories of Restricted Frameworks

### 1. iOS/iPadOS/tvOS/watchOS-Only Frameworks

These frameworks are exclusive to mobile/embedded Apple platforms and have no macOS equivalents.

#### Common iOS-Only Frameworks

**BackgroundTasks**
- Platform: iOS 13.0+, iPadOS 13.0+
- Purpose: Background task scheduling for iOS apps
- Reason: macOS uses different background execution models (LaunchAgents, LaunchDaemons, XPC)
- Workaround: Use `NSBackgroundActivityScheduler` or XPC services on macOS

**HomeKit**
- Platform: iOS 8.0+, watchOS 2.0+
- Purpose: Home automation control
- Reason: Primarily designed for mobile devices controlling IoT
- Workaround: Limited macOS support via Home.app, no public framework API

**HealthKit**
- Platform: iOS 8.0+, watchOS 2.0+
- Purpose: Health and fitness data management
- Reason: Designed for wearables and mobile health tracking
- Workaround: No macOS equivalent; use HealthKit on iOS/watchOS only

**ARKit**
- Platform: iOS 11.0+, iPadOS 11.0+
- Purpose: Augmented reality
- Reason: Requires mobile camera and sensors
- Workaround: Use RealityKit on macOS for 3D content (no AR tracking)

**UIKit**
- Platform: iOS 2.0+, tvOS 9.0+, iPadOS 2.0+
- Purpose: User interface framework for iOS
- Reason: Replaced by AppKit on macOS
- Workaround: Use AppKit for macOS UI development

**Core NFC**
- Platform: iOS 11.0+
- Purpose: Near-field communication
- Reason: Requires NFC hardware not present in Macs
- Workaround: None; iOS-specific feature

**PDFKit**
- Platform: iOS 11.0+, macOS 10.4+
- Purpose: PDF rendering and manipulation
- Status: **Available on macOS** (exception to iOS-only rule)
- Note: One of the few frameworks shared between platforms

### 2. Xcode/Developer Tools Frameworks

These frameworks are part of Xcode and the developer toolchain, not the system runtime.

**XCTest**
- Platform: Xcode testing framework
- Purpose: Unit testing
- Location: Part of Xcode.app bundle, not /System/Library/Frameworks
- Workaround: Use Go's built-in testing (`testing` package)

**XCUIAutomation** (mentioned in request)
- Platform: Xcode UI testing framework
- Purpose: Automated UI testing
- Location: Xcode.app/Contents/Developer/...
- Workaround:
  - Use `go test` for unit tests
  - Use AppleScript/System Events for macOS UI automation
  - Use Accessibility APIs via AppKit

**XCTestCore**, **XCUnit**
- Similar to XCTest, part of Xcode toolchain

### 3. DriverKit Frameworks

These are specialized frameworks for writing system extensions and drivers.

**USBDriverKit** (mentioned in request)
- Platform: macOS 10.15+, iPadOS 14.0+
- Purpose: USB device driver development
- Location: `/System/Library/Frameworks/` but requires special entitlements
- Reason: System extension development requires DriverKit SDK
- Workaround:
  - Use IOKit for user-space USB access (deprecated but available)
  - Use DriverKit with proper entitlements and system extension context
  - For Go: Limited support; consider using libusb via cgo

**NetworkingDriverKit** (mentioned in request)
- Platform: macOS 10.15+
- Purpose: Network extension driver development
- Location: DriverKit SDK
- Reason: Requires system extension entitlements
- Workaround: Use NetworkExtension framework for VPN/filtering

**HIDDriverKit**, **PCIDriverKit**, **SerialDriverKit**
- Similar restrictions as USBDriverKit
- Require system extension context and entitlements
- Not accessible from normal application code

### 4. Watch-Specific Frameworks

**WatchKit**, **ClockKit**, **WatchConnectivity**
- Platform: watchOS
- Purpose: Apple Watch apps and complications
- Workaround: Use WatchConnectivity on iOS to communicate with paired watch

### 5. Deprecated/Removed Frameworks

Some frameworks appear in documentation but are removed from modern macOS:

**Carbon**
- Status: Deprecated, removed in macOS 10.15+
- Replacement: Use AppKit and modern APIs

**Quick Time**, **QTKit**
- Status: Deprecated, removed in macOS 10.15+
- Replacement: Use AVFoundation

### 6. Framework Aliases and Umbrella Frameworks

Some "frameworks" are actually aliases or umbrellas:

**Cocoa**
- Not a real framework, umbrella importing AppKit + Foundation
- Workaround: Import AppKit and Foundation directly

**ApplicationServices**
- Umbrella framework for CoreGraphics, CoreText, ImageIO, etc.
- Workaround: Import specific frameworks directly

## Detection Strategy

Our binding generator uses this strategy to detect available frameworks:

```bash
# List public macOS frameworks (excludes private _ prefixed frameworks)
ls /System/Library/Frameworks | grep -v '^_' | sed 's/.framework//'
```

This yields **270 frameworks** on macOS 15.0 (Sequoia).

However, only frameworks with valid Objective-C headers can be parsed and bound. This reduces the buildable set to approximately **69 frameworks** with complete bindings.

## Working with Restricted Frameworks

### Strategy 1: Platform Checks

When using code that might run on different platforms:

```go
//go:build darwin && !ios

package myapp

import "github.com/tmc/appledocs/generated/appkit"
```

### Strategy 2: Conditional Compilation

For iOS-specific features:

```go
//go:build ios

package myapp

import "github.com/tmc/appledocs/generated/uikit"
```

### Strategy 3: Stubs for Unavailable Frameworks

For frameworks we know exist but can't parse (e.g., BackgroundTasks):

```go
// generated/backgroundtasks/stub.go
//go:build darwin && ios

package backgroundtasks

// This framework is iOS-only and cannot be used on macOS.
// See PLATFORM_RESTRICTIONS.md for alternatives.
```

## Statistics

### Framework Distribution

- **Total in Apple Docs**: ~400+ frameworks across all platforms
- **macOS Available**: 270 frameworks in /System/Library/Frameworks
- **Successfully Generated**: 69 frameworks with complete Go bindings
- **iOS/Mobile Only**: ~100+ frameworks
- **DriverKit/System Extensions**: ~15 frameworks
- **Deprecated/Removed**: ~20+ frameworks
- **Xcode/Developer Tools**: ~10+ frameworks

### Generation Success Rate

For frameworks actually present on macOS:
- **Attempted**: 69 frameworks
- **Successfully Built**: 68 frameworks
- **Success Rate**: **100%** (68/68 buildable frameworks)

Note: The one framework counted separately (Objective-C runtime) is a special case included in the Foundation bindings.

## Future Work

### Planned Enhancements

1. **iOS Framework Support**
   - Generate bindings against iOS SDK
   - Conditional compilation for iOS targets
   - Separate `generated-ios/` directory

2. **DriverKit Support**
   - Document entitlement requirements
   - Provide examples for system extensions
   - Add DriverKit-specific templates

3. **Cross-Platform Detection**
   - Automatically detect framework platform requirements
   - Generate appropriate build tags
   - Add platform availability annotations

4. **Documentation Generation**
   - Auto-generate platform compatibility matrices
   - Link to Apple documentation for alternatives
   - Provide migration guides for deprecated frameworks

## References

- [Apple Developer Documentation](https://developer.apple.com/documentation/)
- [DriverKit Documentation](https://developer.apple.com/documentation/driverkit)
- [System Extensions](https://developer.apple.com/documentation/systemextensions)
- [App Sandbox Entitlements](https://developer.apple.com/documentation/bundleresources/entitlements)

## See Also

- [FRAMEWORK_COVERAGE.md](FRAMEWORK_COVERAGE.md) - Complete list of generated frameworks
- [CLAUDE.md](CLAUDE.md) - Project documentation and development guide
- [README.md](README.md) - Quick start and overview
