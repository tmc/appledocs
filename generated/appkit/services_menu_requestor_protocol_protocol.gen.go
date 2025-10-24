// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

// PServicesMenuRequestor is the NSServicesMenuRequestor protocol interface.
//
// A set of methods that support interaction with items users can share through a sharing service.
//
// Availability:
//   - macOS +
//
// See: doc://com.apple.appkit/documentation/AppKit/NSServicesMenuRequestor
type PServicesMenuRequestor interface {
	// Optional methods
	ReadSelectionFromPasteboard(pboard IPasteboard) bool
	HasReadSelectionFromPasteboard() bool
	WriteSelectionToPasteboardTypes(pboard IPasteboard, types []string) bool
	HasWriteSelectionToPasteboardTypes() bool
}
