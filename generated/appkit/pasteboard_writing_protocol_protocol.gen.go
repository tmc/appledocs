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
	PasteboardPropertyListForType(type_ PasteboardType /* typedef */) objc.ID/* debug [protocol_interface/required_method]: PasteboardPropertyListForType */
	WritableTypesForPasteboard(pasteboard IPasteboard) []string/* debug [protocol_interface/required_method]: WritableTypesForPasteboard */
	// Optional methods
	WritingOptionsForTypePasteboard(type_ PasteboardType /* typedef */, pasteboard IPasteboard) PasteboardWritingOptions
	HasWritingOptionsForTypePasteboard() bool
}
