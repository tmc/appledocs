// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [PDFImageRep] class.
var PDFImageRepClass objc.Class

func init() {
	PDFImageRepClass = objc.GetClass("NSPDFImageRep")
}

type PDFImageRep struct {
	objc.ID
}

func PDFImageRepFrom(ptr unsafe.Pointer) PDFImageRep {
	return PDFImageRep{
		ID: objc.ID(ptr),
	}
}



