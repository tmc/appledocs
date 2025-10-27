// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

// PCaptureScope is the MTLCaptureScope protocol interface.
//
// A type that can programmatically customize a GPU frame capture.
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 11.0+
//   - iPadOS 11.0+
//   - macOS 10.13+
//   - tvOS 11.0+
//   - visionOS 1.0+
//
// See: doc://com.apple.metal/documentation/Metal/MTLCaptureScope
type PCaptureScope interface {
	// Required methods
	BeginScope()
	EndScope()
}
