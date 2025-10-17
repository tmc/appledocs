// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [LayoutDimension] class.
var LayoutDimensionClass objc.Class

func init() {
	LayoutDimensionClass = objc.GetClass("NSLayoutDimension")
}

type LayoutDimension struct {
	objc.ID
}

func LayoutDimensionFrom(ptr unsafe.Pointer) LayoutDimension {
	return LayoutDimension{
		ID: objc.ID(ptr),
	}
}




