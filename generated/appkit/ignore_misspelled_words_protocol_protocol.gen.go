// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (

	"github.com/tmc/appledocs/generated/objc"
)

// PIgnoreMisspelledWords is the NSIgnoreMisspelledWords protocol interface.
//
// A protocol that enables the Ignore button in the Spelling panel to function properly.
//
// Availability:
//   - macOS +
//
// See: doc://com.apple.appkit/documentation/AppKit/NSIgnoreMisspelledWords
type PIgnoreMisspelledWords interface {
	// Required methods
	IgnoreSpelling(sender objc.IObject)/* debug [protocol_interface/required_method]: IgnoreSpelling */
}
