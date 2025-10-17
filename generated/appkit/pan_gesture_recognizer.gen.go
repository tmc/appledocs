// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [PanGestureRecognizer] class.
var panGestureRecognizerClass = _PanGestureRecognizerClass{objc.GetClass("NSPanGestureRecognizer")}

type _PanGestureRecognizerClass struct {
	class objc.Class
}

// An interface definition for the [PanGestureRecognizer] class.
type IPanGestureRecognizer interface {
	IGestureRecognizer
}

// A continuous gesture recognizer for panning gestures. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPanGestureRecognizer

type PanGestureRecognizer struct {
	GestureRecognizer
}

// PanGestureRecognizerFrom constructs a [PanGestureRecognizer] from an unsafe.Pointer.
//
// A continuous gesture recognizer for panning gestures.
func PanGestureRecognizerFrom(ptr unsafe.Pointer) PanGestureRecognizer {
	return PanGestureRecognizer{
		GestureRecognizer: GestureRecognizerFrom(ptr),
	}
}



