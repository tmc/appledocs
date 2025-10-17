// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [ColorList] class.
var ColorListClass objc.Class

func init() {
	ColorListClass = objc.GetClass("NSColorList")
}

type ColorList struct {
	objc.ID
}

func ColorListFrom(ptr unsafe.Pointer) ColorList {
	return ColorList{
		ID: objc.ID(ptr),
	}
}



