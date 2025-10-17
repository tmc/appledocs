// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [PDFPanel] class.
var pDFPanelClass = _PDFPanelClass{objc.GetClass("NSPDFPanel")}

type _PDFPanelClass struct {
	class objc.Class
}

// A Save or Export as PDF panel that’s consistent with the macOS user interface. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPDFPanel

type PDFPanel struct {
	objectivec.Object
}

// PDFPanelFrom constructs a [PDFPanel] from an unsafe.Pointer.
//
// A Save or Export as PDF panel that’s consistent with the macOS user interface.
func PDFPanelFrom(ptr unsafe.Pointer) PDFPanel {
	return PDFPanel{objectivec.Object{objc.ID(ptr)}}
}



