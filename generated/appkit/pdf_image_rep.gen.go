// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [PDFImageRep] class.
var (
	pDFImageRepClass     _PDFImageRepClass
	pDFImageRepClassOnce sync.Once
)

func getPDFImageRepClass() _PDFImageRepClass {
	pDFImageRepClassOnce.Do(func() {
		pDFImageRepClass = _PDFImageRepClass{objc.GetClass("NSPDFImageRep")}
	})
	return pDFImageRepClass
}

type _PDFImageRepClass struct {
	class objc.Class
}

// An interface definition for the [PDFImageRep] class.
type IPDFImageRep interface {
	IImageRep
}

// An object that can render an image from a PDF format data stream.
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
func (fc _PDFImageRepClass) Alloc() PDFImageRep {
	rv := objc.Send[PDFImageRep](objc.ID(fc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (fc _PDFImageRepClass) New() PDFImageRep {
	rv := objc.Send[PDFImageRep](objc.ID(fc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (f_ PDFImageRep) Init() PDFImageRep {
	rv := objc.Send[PDFImageRep](f_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (f_ PDFImageRep) Autorelease() PDFImageRep {
	rv := objc.Send[PDFImageRep](f_.ID, objc.Sel("autorelease"))
	return rv
}

// NewPDFImageRep creates a new PDFImageRep instance.
func NewPDFImageRep() PDFImageRep {
	return getPDFImageRepClass().New()
}




