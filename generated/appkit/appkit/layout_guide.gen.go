// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [LayoutGuide] class.
var LayoutGuideClass objc.Class

func init() {
	LayoutGuideClass = objc.GetClass("NSLayoutGuide")
}

type LayoutGuide struct {
	objc.ID
}

func LayoutGuideFrom(ptr unsafe.Pointer) LayoutGuide {
	return LayoutGuide{
		ID: objc.ID(ptr),
	}
}



