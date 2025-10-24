// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

// PTextFinderBarContainer is the NSTextFinderBarContainer protocol interface.
//
// A protocol that provides a container in which the find bar is displayed.
//
// Availability:
//   - macOS +
//
// See: doc://com.apple.appkit/documentation/AppKit/NSTextFinderBarContainer
type PTextFinderBarContainer interface {
	// Required methods
	FindBarViewDidChangeHeight()
	// Optional methods
	ContentView() View
	HasContentView() bool
}
