// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (

	"github.com/tmc/appledocs/generated/objc"

	"github.com/tmc/appledocs/generated/corefoundation"

	"github.com/tmc/appledocs/generated/objectivec"
)

// PInputServerMouseTracker is the NSInputServerMouseTracker protocol interface.
//
// Availability:
//   - macOS +
//
// See: doc://com.apple.appkit/documentation/AppKit/NSInputServerMouseTracker
type PInputServerMouseTracker interface {
	// Required methods
	MouseDownOnCharacterIndexAtCoordinateWithModifierClient(index uint, point corefoundation.CGPoint, flags uint, sender objectivec.IObject) bool
	MouseDraggedOnCharacterIndexAtCoordinateWithModifierClient(index uint, point corefoundation.CGPoint, flags uint, sender objectivec.IObject) bool
	MouseUpOnCharacterIndexAtCoordinateWithModifierClient(index uint, point corefoundation.CGPoint, flags uint, sender objectivec.IObject)
}
