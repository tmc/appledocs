// Code generated from Apple documentation for InputMethodKit. DO NOT EDIT.

package inputmethodkit

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"

	"github.com/tmc/appledocs/generated/vision"
)

// PIMKMouseHandling is the IMKMouseHandling protocol interface.
//
// The   protocol defines methods that your input method can implement to handle mouse events.
//
// Availability:
//   - macOS 10.5+
//
// See: doc://com.apple.inputmethodkit/documentation/InputMethodKit/IMKMouseHandling
type PIMKMouseHandling interface {
	// Required methods
	MouseDownOnCharacterIndexCoordinateWithModifierContinueTrackingClient(index uint, point vision.Point, flags uint, keepTracking unsafe.Pointer, sender objc.IObject) bool/* debug [protocol_interface/required_method]: MouseDownOnCharacterIndexCoordinateWithModifierContinueTrackingClient */
	MouseMovedOnCharacterIndexCoordinateWithModifierClient(index uint, point vision.Point, flags uint, sender objc.IObject) bool/* debug [protocol_interface/required_method]: MouseMovedOnCharacterIndexCoordinateWithModifierClient */
	MouseUpOnCharacterIndexCoordinateWithModifierClient(index uint, point vision.Point, flags uint, sender objc.IObject) bool/* debug [protocol_interface/required_method]: MouseUpOnCharacterIndexCoordinateWithModifierClient */
}
