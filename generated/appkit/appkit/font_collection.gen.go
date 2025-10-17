// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [FontCollection] class.
var FontCollectionClass objc.Class

func init() {
	FontCollectionClass = objc.GetClass("NSFontCollection")
}

type FontCollection struct {
	objc.ID
}

func FontCollectionFrom(ptr unsafe.Pointer) FontCollection {
	return FontCollection{
		ID: objc.ID(ptr),
	}
}




