// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

// PMenuItemValidation is the NSMenuItemValidation protocol interface.
//
// Availability:
//   - macOS +
//
// See: doc://com.apple.appkit/documentation/AppKit/NSMenuItemValidation
type PMenuItemValidation interface {
	// Required methods
	ValidateMenuItem(menuItem IMenuItem) bool/* debug [protocol_interface/required_method]: ValidateMenuItem */
}
