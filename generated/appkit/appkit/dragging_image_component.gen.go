// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [DraggingImageComponent] class.
var DraggingImageComponentClass objc.Class

func init() {
	DraggingImageComponentClass = objc.GetClass("NSDraggingImageComponent")
}

type DraggingImageComponent struct {
	objc.ID
}

func DraggingImageComponentFrom(ptr unsafe.Pointer) DraggingImageComponent {
	return DraggingImageComponent{
		ID: objc.ID(ptr),
	}
}




