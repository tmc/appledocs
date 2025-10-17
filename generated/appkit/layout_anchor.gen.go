// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [LayoutAnchor] class.
var LayoutAnchorClass objc.Class

func init() {
	LayoutAnchorClass = objc.GetClass("NSLayoutAnchor")
}

type LayoutAnchor struct {
	objc.ID
}

func LayoutAnchorFrom(ptr unsafe.Pointer) LayoutAnchor {
	return LayoutAnchor{
		ID: objc.ID(ptr),
	}
}



