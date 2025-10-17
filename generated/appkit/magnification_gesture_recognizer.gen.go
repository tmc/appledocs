// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [MagnificationGestureRecognizer] class.
var magnificationGestureRecognizerClass = _MagnificationGestureRecognizerClass{objc.GetClass("NSMagnificationGestureRecognizer")}

type _MagnificationGestureRecognizerClass struct {
	class objc.Class
}

// A continuous gesture recognizer that tracks a pinch gesture that magnifies content. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMagnificationGestureRecognizer

type MagnificationGestureRecognizer struct {
	GestureRecognizer
}

// MagnificationGestureRecognizerFrom constructs a [MagnificationGestureRecognizer] from an unsafe.Pointer.
//
// A continuous gesture recognizer that tracks a pinch gesture that magnifies content.
func MagnificationGestureRecognizerFrom(ptr unsafe.Pointer) MagnificationGestureRecognizer {
	return MagnificationGestureRecognizer{
		GestureRecognizer: GestureRecognizerFrom(ptr),
	}
}



