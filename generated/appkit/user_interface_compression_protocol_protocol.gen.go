// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

// PUserInterfaceCompression is the NSUserInterfaceCompression protocol interface.
//
// A protocol that describes how a UI control should redisplay when space is restricted.
//
// Availability:
//   - macOS +
//
// See: doc://com.apple.appkit/documentation/AppKit/NSUserInterfaceCompression
type PUserInterfaceCompression interface {
	// Required methods
	CompressWithPrioritizedCompressionOptions(prioritizedOptions []UserInterfaceCompressionOptions)
	MinimumSizeWithPrioritizedCompressionOptions(prioritizedOptions []UserInterfaceCompressionOptions) corefoundation.Size
}
