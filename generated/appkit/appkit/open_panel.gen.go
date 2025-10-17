// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [OpenPanel] class.
var OpenPanelClass objc.Class

func init() {
	OpenPanelClass = objc.GetClass("NSOpenPanel")
}

type OpenPanel struct {
	objc.ID
}

func OpenPanelFrom(ptr unsafe.Pointer) OpenPanel {
	return OpenPanel{
		ID: objc.ID(ptr),
	}
}




