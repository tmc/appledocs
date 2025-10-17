// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [ClickGestureRecognizer] class.
var ClickGestureRecognizerClass objc.Class

func init() {
	ClickGestureRecognizerClass = objc.GetClass("NSClickGestureRecognizer")
}

type ClickGestureRecognizer struct {
	objc.ID
}

func ClickGestureRecognizerFrom(ptr unsafe.Pointer) ClickGestureRecognizer {
	return ClickGestureRecognizer{
		ID: objc.ID(ptr),
	}
}




