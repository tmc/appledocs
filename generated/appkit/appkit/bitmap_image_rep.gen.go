// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [BitmapImageRep] class.
var BitmapImageRepClass objc.Class

func init() {
	BitmapImageRepClass = objc.GetClass("NSBitmapImageRep")
}

type BitmapImageRep struct {
	objc.ID
}

func BitmapImageRepFrom(ptr unsafe.Pointer) BitmapImageRep {
	return BitmapImageRep{
		ID: objc.ID(ptr),
	}
}




