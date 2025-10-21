// Code generated from Apple documentation for PDFKit. DO NOT EDIT.

package pdfkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [PDFAnnotationSquare] class.
var (
	PDFAnnotationSquareClass     _PDFAnnotationSquareClass
	PDFAnnotationSquareClassOnce sync.Once
)

func getPDFAnnotationSquareClass() _PDFAnnotationSquareClass {
	PDFAnnotationSquareClassOnce.Do(func() {
		PDFAnnotationSquareClass = _PDFAnnotationSquareClass{objc.GetClass("PDFAnnotationSquare")}
	})
	return PDFAnnotationSquareClass
}

type _PDFAnnotationSquareClass struct {
	class objc.Class
}

// An interface definition for the [PDFAnnotationSquare] class.
type IPDFAnnotationSquare interface {
	IPDFAnnotation
}

// A rectangle annotation on a page.
//
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFAnnotationSquare
type PDFAnnotationSquare struct {
	PDFAnnotation
}

// PDFAnnotationSquareFrom constructs a [PDFAnnotationSquare] from an unsafe.Pointer.
//
// A rectangle annotation on a page.
func PDFAnnotationSquareFrom(ptr unsafe.Pointer) PDFAnnotationSquare {
	return PDFAnnotationSquare{
		PDFAnnotation: PDFAnnotationFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (pc _PDFAnnotationSquareClass) Alloc() PDFAnnotationSquare {
	rv := objc.Send[PDFAnnotationSquare](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (pc _PDFAnnotationSquareClass) New() PDFAnnotationSquare {
	rv := objc.Send[PDFAnnotationSquare](objc.ID(pc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (p_ PDFAnnotationSquare) Init() PDFAnnotationSquare {
	rv := objc.Send[PDFAnnotationSquare](p_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p_ PDFAnnotationSquare) Autorelease() PDFAnnotationSquare {
	rv := objc.Send[PDFAnnotationSquare](p_.ID, objc.Sel("autorelease"))
	return rv
}

// NewPDFAnnotationSquare creates a new PDFAnnotationSquare instance.
func NewPDFAnnotationSquare() PDFAnnotationSquare {
	return getPDFAnnotationSquareClass().New()
}




