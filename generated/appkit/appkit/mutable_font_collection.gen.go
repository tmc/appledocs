// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [MutableFontCollection] class.
var MutableFontCollectionClass objc.Class

func init() {
	MutableFontCollectionClass = objc.GetClass("NSMutableFontCollection")
}

type MutableFontCollection struct {
	objc.ID
}

func MutableFontCollectionFrom(ptr unsafe.Pointer) MutableFontCollection {
	return MutableFontCollection{
		ID: objc.ID(ptr),
	}
}



