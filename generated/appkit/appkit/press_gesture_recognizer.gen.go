// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [PressGestureRecognizer] class.
var PressGestureRecognizerClass objc.Class

func init() {
	PressGestureRecognizerClass = objc.GetClass("NSPressGestureRecognizer")
}

type PressGestureRecognizer struct {
	objc.ID
}

func PressGestureRecognizerFrom(ptr unsafe.Pointer) PressGestureRecognizer {
	return PressGestureRecognizer{
		ID: objc.ID(ptr),
	}
}




