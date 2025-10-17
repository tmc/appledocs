// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [RotationGestureRecognizer] class.
var rotationGestureRecognizerClass = _RotationGestureRecognizerClass{objc.GetClass("NSRotationGestureRecognizer")}

type _RotationGestureRecognizerClass struct {
	class objc.Class
}

// An interface definition for the [RotationGestureRecognizer] class.
type IRotationGestureRecognizer interface {
	IGestureRecognizer
}

// A continuous gesture recognizer that tracks two trackpad touches moving opposite each other in a circular motion. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSRotationGestureRecognizer

type RotationGestureRecognizer struct {
	GestureRecognizer
}

// RotationGestureRecognizerFrom constructs a [RotationGestureRecognizer] from an unsafe.Pointer.
//
// A continuous gesture recognizer that tracks two trackpad touches moving opposite each other in a circular motion.
func RotationGestureRecognizerFrom(ptr unsafe.Pointer) RotationGestureRecognizer {
	return RotationGestureRecognizer{
		GestureRecognizer: GestureRecognizerFrom(ptr),
	}
}



