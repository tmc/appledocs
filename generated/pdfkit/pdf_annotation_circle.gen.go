// Code generated from Apple documentation for PDFKit. DO NOT EDIT.

package pdfkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [PDFAnnotationCircle] class.
var (
	PDFAnnotationCircleClass     _PDFAnnotationCircleClass
	PDFAnnotationCircleClassOnce sync.Once
)

func getPDFAnnotationCircleClass() _PDFAnnotationCircleClass {
	PDFAnnotationCircleClassOnce.Do(func() {
		PDFAnnotationCircleClass = _PDFAnnotationCircleClass{objc.GetClass("PDFAnnotationCircle")}
	})
	return PDFAnnotationCircleClass
}

type _PDFAnnotationCircleClass struct {
	class objc.Class
}

// An interface definition for the [PDFAnnotationCircle] class.
type IPDFAnnotationCircle interface {
	IPDFAnnotation
}

//
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFAnnotationCircle
type PDFAnnotationCircle struct {
	PDFAnnotation
}

// PDFAnnotationCircleFrom constructs a [PDFAnnotationCircle] from an unsafe.Pointer.
func PDFAnnotationCircleFrom(ptr unsafe.Pointer) PDFAnnotationCircle {
	return PDFAnnotationCircle{
		PDFAnnotation: PDFAnnotationFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (pc _PDFAnnotationCircleClass) Alloc() PDFAnnotationCircle {
	rv := objc.Send[PDFAnnotationCircle](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (pc _PDFAnnotationCircleClass) New() PDFAnnotationCircle {
	rv := objc.Send[PDFAnnotationCircle](objc.ID(pc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (p_ PDFAnnotationCircle) Init() PDFAnnotationCircle {
	rv := objc.Send[PDFAnnotationCircle](p_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p_ PDFAnnotationCircle) Autorelease() PDFAnnotationCircle {
	rv := objc.Send[PDFAnnotationCircle](p_.ID, objc.Sel("autorelease"))
	return rv
}

// NewPDFAnnotationCircle creates a new PDFAnnotationCircle instance.
func NewPDFAnnotationCircle() PDFAnnotationCircle {
	return getPDFAnnotationCircleClass().New()
}




