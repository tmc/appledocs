// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [Panel] class.
var PanelClass objc.Class

func init() {
	PanelClass = objc.GetClass("NSPanel")
}

type Panel struct {
	objc.ID
}

func PanelFrom(ptr unsafe.Pointer) Panel {
	return Panel{
		ID: objc.ID(ptr),
	}
}



