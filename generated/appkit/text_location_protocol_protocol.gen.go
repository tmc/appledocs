// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (

	"github.com/tmc/appledocs/generated/objc"
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
	Compare(location objc.IObject) ComparisonResult
}
