// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"
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
	ValidateUserInterfaceItem(item unsafe.Pointer) bool/* debug [protocol_interface/required_method]: ValidateUserInterfaceItem */
}
