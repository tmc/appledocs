// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

// PDockTilePlugIn is the NSDockTilePlugIn protocol interface.
//
// A set of methods implemented by plug-ins that allow an app’s Dock tile to be customized while the app is not running.
//
// Availability:
//   - macOS +
//
// See: doc://com.apple.appkit/documentation/AppKit/NSDockTilePlugIn
type PDockTilePlugIn interface {
	// Optional methods
	DockMenu() Menu
	HasDockMenu() bool
}
