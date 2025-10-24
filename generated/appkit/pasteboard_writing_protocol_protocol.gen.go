// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (

	"github.com/tmc/appledocs/generated/objc"
)

// PPasteboardWriting is the NSPasteboardWriting protocol interface.
//
// A set of methods that defines the interface for retrieving a representation of an object that can be written to a pasteboard.
//
// Availability:
//   - macOS +
//
// See: doc://com.apple.appkit/documentation/AppKit/NSPasteboardWriting
type PPasteboardWriting interface {
	// Required methods
	PasteboardPropertyListForType(type_ objc.IObject /* cross-framework: PasteboardType */) objc.ID
	WritableTypesForPasteboard(pasteboard IPasteboard) []string
	// Optional methods
	WritingOptionsForTypePasteboard(type_ objc.IObject /* cross-framework: PasteboardType */, pasteboard IPasteboard) PasteboardWritingOptions
	HasWritingOptionsForTypePasteboard() bool
}
