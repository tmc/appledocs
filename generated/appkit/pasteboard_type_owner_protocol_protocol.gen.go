// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

// PPasteboardTypeOwner is the NSPasteboardTypeOwner protocol interface.
//
// An object that serves as a data provider for data types that use lazy data fulfillment from a pasteboard request.
//
// Availability:
//   - macOS +
//
// See: doc://com.apple.appkit/documentation/AppKit/NSPasteboardTypeOwner
type PPasteboardTypeOwner interface {
	// Required methods
	PasteboardProvideDataForType(sender IPasteboard, type_ objc.IObject /* cross-framework: PasteboardType */)
	// Optional methods
	PasteboardChangedOwner(sender IPasteboard)
	HasPasteboardChangedOwner() bool
}
