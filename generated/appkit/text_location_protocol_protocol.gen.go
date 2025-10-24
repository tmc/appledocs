// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"
)

// PTextLocation is the NSTextLocation protocol interface.
//
// An interface you implement that represents an abstract location inside your document’s content.
//
// Availability:
//   - macOS 12.0+
//
// See: doc://com.apple.appkit/documentation/AppKit/NSTextLocation
type PTextLocation interface {
	// Required methods
	Compare(location unsafe.Pointer) ComparisonResult/* debug [protocol_interface/required_method]: Compare */
}
