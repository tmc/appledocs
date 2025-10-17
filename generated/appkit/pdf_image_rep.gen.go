// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [PDFImageRep] class.
var pDFImageRepClass = _PDFImageRepClass{objc.GetClass("NSPDFImageRep")}

type _PDFImageRepClass struct {
	class objc.Class
}

// An object that can render an image from a PDF format data stream. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPDFImageRep

type PDFImageRep struct {
	ImageRep
}

// PDFImageRepFrom constructs a [PDFImageRep] from an unsafe.Pointer.
//
// An object that can render an image from a PDF format data stream.
func PDFImageRepFrom(ptr unsafe.Pointer) PDFImageRep {
	return PDFImageRep{
		ImageRep: ImageRepFrom(ptr),
	}
}



