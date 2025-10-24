// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (

	"github.com/tmc/appledocs/generated/objc"
)

// PPasteboardReading is the NSPasteboardReading protocol interface.
//
// A set of methods that defines the interface for initializing an object from a pasteboard.
//
// Availability:
//   - macOS +
//
// See: doc://com.apple.appkit/documentation/AppKit/NSPasteboardReading
type PPasteboardReading interface {
	// Required methods
	InitWithPasteboardPropertyListOfType(propertyList objc.IObject, type_ objc.IObject /* cross-framework: PasteboardType */) objc.ID
}
