// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [ClickGestureRecognizer] class.
var clickGestureRecognizerClass = _ClickGestureRecognizerClass{objc.GetClass("NSClickGestureRecognizer")}

type _ClickGestureRecognizerClass struct {
	class objc.Class
}

// An interface definition for the [ClickGestureRecognizer] class.
type IClickGestureRecognizer interface {
	IGestureRecognizer
}

// A discrete gesture recognizer that tracks a specified number of mouse clicks. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSClickGestureRecognizer

type ClickGestureRecognizer struct {
	GestureRecognizer
}

// ClickGestureRecognizerFrom constructs a [ClickGestureRecognizer] from an unsafe.Pointer.
//
// A discrete gesture recognizer that tracks a specified number of mouse clicks.
func ClickGestureRecognizerFrom(ptr unsafe.Pointer) ClickGestureRecognizer {
	return ClickGestureRecognizer{
		GestureRecognizer: GestureRecognizerFrom(ptr),
	}
}



