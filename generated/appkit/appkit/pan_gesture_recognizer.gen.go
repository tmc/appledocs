// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [PanGestureRecognizer] class.
var PanGestureRecognizerClass objc.Class

func init() {
	PanGestureRecognizerClass = objc.GetClass("NSPanGestureRecognizer")
}

type PanGestureRecognizer struct {
	objc.ID
}

func PanGestureRecognizerFrom(ptr unsafe.Pointer) PanGestureRecognizer {
	return PanGestureRecognizer{
		ID: objc.ID(ptr),
	}
}




