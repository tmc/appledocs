// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [LayoutYAxisAnchor] class.
var LayoutYAxisAnchorClass objc.Class

func init() {
	LayoutYAxisAnchorClass = objc.GetClass("NSLayoutYAxisAnchor")
}

type LayoutYAxisAnchor struct {
	objc.ID
}

func LayoutYAxisAnchorFrom(ptr unsafe.Pointer) LayoutYAxisAnchor {
	return LayoutYAxisAnchor{
		ID: objc.ID(ptr),
	}
}




