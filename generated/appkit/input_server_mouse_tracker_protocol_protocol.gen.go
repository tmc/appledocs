// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (

	"github.com/tmc/appledocs/generated/objc"
)

// PInputServerMouseTracker is the NSInputServerMouseTracker protocol interface.
//
// Availability:
//   - macOS +
//
// See: doc://com.apple.appkit/documentation/AppKit/NSInputServerMouseTracker
type PInputServerMouseTracker interface {
	// Required methods
	MouseDownOnCharacterIndexAtCoordinateWithModifierClient(index uint, point objc.IObject /* cross-framework: Point */, flags uint, sender objc.IObject) bool
	MouseDraggedOnCharacterIndexAtCoordinateWithModifierClient(index uint, point objc.IObject /* cross-framework: Point */, flags uint, sender objc.IObject) bool
	MouseUpOnCharacterIndexAtCoordinateWithModifierClient(index uint, point objc.IObject /* cross-framework: Point */, flags uint, sender objc.IObject)
}
