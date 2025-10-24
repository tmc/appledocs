// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (

	"github.com/tmc/appledocs/generated/objc"
)

// PUserInterfaceValidations is the NSUserInterfaceValidations protocol interface.
//
// A protocol that a custom class can adopt to manage the enabled state of a UI element.
//
// Availability:
//   - macOS +
//
// See: doc://com.apple.appkit/documentation/AppKit/NSUserInterfaceValidations
type PUserInterfaceValidations interface {
	// Required methods
	ValidateUserInterfaceItem(item objc.IObject) bool
}
