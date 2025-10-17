// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [PrintPanel] class.
var PrintPanelClass objc.Class

func init() {
	PrintPanelClass = objc.GetClass("NSPrintPanel")
}

type PrintPanel struct {
	objc.ID
}

func PrintPanelFrom(ptr unsafe.Pointer) PrintPanel {
	return PrintPanel{
		ID: objc.ID(ptr),
	}
}



