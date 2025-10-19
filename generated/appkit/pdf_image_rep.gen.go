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

// An interface definition for the [PDFImageRep] class.
type IPDFImageRep interface {
	IImageRep
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
// Alloc allocates a new instance without initialization.
func (pc _PDFImageRepClass) Alloc() PDFImageRep {
	rv := objc.Send[PDFImageRep](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new instance with a +1 retain count.
func (pc _PDFImageRepClass) New() PDFImageRep {
	rv := objc.Send[PDFImageRep](objc.ID(pc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (p_ PDFImageRep) Init() PDFImageRep {
	rv := objc.Send[PDFImageRep](p_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p_ PDFImageRep) Autorelease() PDFImageRep {
	rv := objc.Send[PDFImageRep](p_.ID, objc.Sel("autorelease"))
	return rv
}

// NewPDFImageRep creates a new PDFImageRep instance.
func NewPDFImageRep() PDFImageRep {
	return pDFImageRepClass.New()
}




