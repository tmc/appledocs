// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [FontPanel] class.
var FontPanelClass objc.Class

func init() {
	FontPanelClass = objc.GetClass("NSFontPanel")
}

type FontPanel struct {
	objc.ID
}

func FontPanelFrom(ptr unsafe.Pointer) FontPanel {
	return FontPanel{
		ID: objc.ID(ptr),
	}
}



