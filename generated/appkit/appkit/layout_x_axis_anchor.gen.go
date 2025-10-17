// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [LayoutXAxisAnchor] class.
var LayoutXAxisAnchorClass objc.Class

func init() {
	LayoutXAxisAnchorClass = objc.GetClass("NSLayoutXAxisAnchor")
}

type LayoutXAxisAnchor struct {
	objc.ID
}

func LayoutXAxisAnchorFrom(ptr unsafe.Pointer) LayoutXAxisAnchor {
	return LayoutXAxisAnchor{
		ID: objc.ID(ptr),
	}
}




