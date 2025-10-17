// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [RotationGestureRecognizer] class.
var RotationGestureRecognizerClass objc.Class

func init() {
	RotationGestureRecognizerClass = objc.GetClass("NSRotationGestureRecognizer")
}

type RotationGestureRecognizer struct {
	objc.ID
}

func RotationGestureRecognizerFrom(ptr unsafe.Pointer) RotationGestureRecognizer {
	return RotationGestureRecognizer{
		ID: objc.ID(ptr),
	}
}




