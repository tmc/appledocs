// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [PDFPanel] class.
var PDFPanelClass objc.Class

func init() {
	PDFPanelClass = objc.GetClass("NSPDFPanel")
}

type PDFPanel struct {
	objc.ID
}

func PDFPanelFrom(ptr unsafe.Pointer) PDFPanel {
	return PDFPanel{
		ID: objc.ID(ptr),
	}
}




