// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

// PExtensionRequestHandling is the NSExtensionRequestHandling protocol interface.
//
// The interface an app extension uses to respond to a request from a host app.
//
// Availability:
//   - Mac Catalyst 13.0+
//   - iOS 2.0+
//   - iPadOS 2.0+
//   - macOS 10.0+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 2.0+
//
// See: doc://com.apple.foundation/documentation/Foundation/NSExtensionRequestHandling
type PExtensionRequestHandling interface {
	// Required methods
	BeginRequestWithExtensionContext(context IExtensionContext)
}
