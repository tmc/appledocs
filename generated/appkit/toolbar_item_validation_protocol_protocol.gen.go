// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

// PToolbarItemValidation is the NSToolbarItemValidation protocol interface.
//
// Validation of a toolbar item.
//
// Availability:
//   - macOS +
//
// See: doc://com.apple.appkit/documentation/AppKit/NSToolbarItemValidation
type PToolbarItemValidation interface {
	// Required methods
	ValidateToolbarItem(item IToolbarItem) bool/* debug [protocol_interface/required_method]: ValidateToolbarItem */
}
