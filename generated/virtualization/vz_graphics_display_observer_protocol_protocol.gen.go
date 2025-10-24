// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

// PVZGraphicsDisplayObserver is the VZGraphicsDisplayObserver protocol interface.
//
// A protocol you implement to observe state changes in graphic displays.
//
// Availability:
//   - macOS 14.0+
//
// See: doc://com.apple.virtualization/documentation/Virtualization/VZGraphicsDisplayObserver
type PVZGraphicsDisplayObserver interface {
	// Optional methods
	DisplayDidBeginReconfiguration(display IVZGraphicsDisplay)
	HasDisplayDidBeginReconfiguration() bool
	DisplayDidEndReconfiguration(display IVZGraphicsDisplay)
	HasDisplayDidEndReconfiguration() bool
}
