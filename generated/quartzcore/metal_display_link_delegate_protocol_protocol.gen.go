// Code generated from Apple documentation for QuartzCore. DO NOT EDIT.

package quartzcore

// PMetalDisplayLinkDelegate is the CAMetalDisplayLinkDelegate protocol interface.
//
// A protocol your app implements to respond to callbacks from Core Animation for a Metal display link.
//
// Availability:
//   - Mac Catalyst 17.0+
//   - iOS 17.0+
//   - iPadOS 17.0+
//   - macOS 14.0+
//   - tvOS 17.0+
//   - visionOS 1.0+
//
// See: doc://com.apple.quartzcore/documentation/QuartzCore/CAMetalDisplayLinkDelegate
type PMetalDisplayLinkDelegate interface {
	// Required methods
	MetalDisplayLinkNeedsUpdate(link IMetalDisplayLink, update IMetalDisplayLinkUpdate)
}
