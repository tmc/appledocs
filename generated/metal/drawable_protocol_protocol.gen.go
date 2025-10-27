// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

// PDrawable is the MTLDrawable protocol interface.
//
// A displayable resource that can be rendered or written to.
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 8.0+
//   - iPadOS 8.0+
//   - macOS 10.11+
//   - tvOS +
//   - visionOS 1.0+
//
// See: doc://com.apple.metal/documentation/Metal/MTLDrawable
type PDrawable interface {
	// Required methods
	AddPresentedHandler(block DrawablePresentedHandler /* not a class type */)
	Present()
	PresentAfterMinimumDuration(duration float64)
	PresentAtTime(presentationTime float64)
}
