// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

// PMTL4CompilerTask is the MTL4CompilerTask protocol interface.
//
// A reference to an asynchronous compilation task that you initiate from a compiler instance.
//
// Availability:
//   - Mac Catalyst 26.0+
//   - iOS 26.0+
//   - iPadOS 26.0+
//   - macOS 26.0+
//   - tvOS 26.0+
//   - visionOS 26.0+
//
// See: doc://com.apple.metal/documentation/Metal/MTL4CompilerTask
type PMTL4CompilerTask interface {
	// Required methods
	WaitUntilCompleted()
}
