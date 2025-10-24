// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (

	"github.com/tmc/appledocs/generated/objc"

	"github.com/tmc/appledocs/generated/vision"
)

// PInputServerMouseTracker is the NSInputServerMouseTracker protocol interface.
//
// Availability:
//   - macOS +
//
// See: doc://com.apple.appkit/documentation/AppKit/NSInputServerMouseTracker
type PInputServerMouseTracker interface {
	// Required methods
	MouseDownOnCharacterIndexAtCoordinateWithModifierClient(index uint, point vision.Point, flags uint, sender objc.IObject) bool/* debug [protocol_interface/required_method]: MouseDownOnCharacterIndexAtCoordinateWithModifierClient */
	MouseDraggedOnCharacterIndexAtCoordinateWithModifierClient(index uint, point vision.Point, flags uint, sender objc.IObject) bool/* debug [protocol_interface/required_method]: MouseDraggedOnCharacterIndexAtCoordinateWithModifierClient */
	MouseUpOnCharacterIndexAtCoordinateWithModifierClient(index uint, point vision.Point, flags uint, sender objc.IObject)/* debug [protocol_interface/required_method]: MouseUpOnCharacterIndexAtCoordinateWithModifierClient */
}
