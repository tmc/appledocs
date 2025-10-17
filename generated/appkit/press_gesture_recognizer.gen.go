// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [PressGestureRecognizer] class.
var pressGestureRecognizerClass = _PressGestureRecognizerClass{objc.GetClass("NSPressGestureRecognizer")}

type _PressGestureRecognizerClass struct {
	class objc.Class
}

// A discrete gesture recognizer that tracks whether the user holds down a mouse button for a minimum amount of time before releasing it. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPressGestureRecognizer

type PressGestureRecognizer struct {
	GestureRecognizer
}

// PressGestureRecognizerFrom constructs a [PressGestureRecognizer] from an unsafe.Pointer.
//
// A discrete gesture recognizer that tracks whether the user holds down a mouse button for a minimum amount of time before releasing it.
func PressGestureRecognizerFrom(ptr unsafe.Pointer) PressGestureRecognizer {
	return PressGestureRecognizer{
		GestureRecognizer: GestureRecognizerFrom(ptr),
	}
}



