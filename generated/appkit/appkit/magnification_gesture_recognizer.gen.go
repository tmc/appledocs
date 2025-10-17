// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [MagnificationGestureRecognizer] class.
var MagnificationGestureRecognizerClass objc.Class

func init() {
	MagnificationGestureRecognizerClass = objc.GetClass("NSMagnificationGestureRecognizer")
}

type MagnificationGestureRecognizer struct {
	objc.ID
}

func MagnificationGestureRecognizerFrom(ptr unsafe.Pointer) MagnificationGestureRecognizer {
	return MagnificationGestureRecognizer{
		ID: objc.ID(ptr),
	}
}




