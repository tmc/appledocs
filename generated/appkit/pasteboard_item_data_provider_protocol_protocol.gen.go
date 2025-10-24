// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

// PPasteboardItemDataProvider is the NSPasteboardItemDataProvider protocol interface.
//
// A set of methods implemented by the data provider of a pasteboard item to provide the data for a particular UTI type.
//
// Availability:
//   - macOS +
//
// See: doc://com.apple.appkit/documentation/AppKit/NSPasteboardItemDataProvider
type PPasteboardItemDataProvider interface {
	// Required methods
	PasteboardItemProvideDataForType(pasteboard IPasteboard, item IPasteboardItem, type_ PasteboardType /* typedef */)/* debug [protocol_interface/required_method]: PasteboardItemProvideDataForType */
	// Optional methods
	PasteboardFinishedWithDataProvider(pasteboard IPasteboard)
	HasPasteboardFinishedWithDataProvider() bool
}
