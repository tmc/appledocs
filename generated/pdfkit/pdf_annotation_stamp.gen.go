// Code generated from Apple documentation for PDFKit. DO NOT EDIT.

package pdfkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [PDFAnnotationStamp] class.
var (
	PDFAnnotationStampClass     _PDFAnnotationStampClass
	PDFAnnotationStampClassOnce sync.Once
)

func getPDFAnnotationStampClass() _PDFAnnotationStampClass {
	PDFAnnotationStampClassOnce.Do(func() {
		PDFAnnotationStampClass = _PDFAnnotationStampClass{objc.GetClass("PDFAnnotationStamp")}
	})
	return PDFAnnotationStampClass
}

type _PDFAnnotationStampClass struct {
	class objc.Class
}

// An interface definition for the [PDFAnnotationStamp] class.
type IPDFAnnotationStamp interface {
	IPDFAnnotation
}

// A object allows you to display a word or phrase, such as “Confidential,” in a PDF page.
//
// A object should have an appearance stream associated with it; otherwise, nothing useful is rendered.
//
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFAnnotationStamp
type PDFAnnotationStamp struct {
	PDFAnnotation
}

// PDFAnnotationStampFrom constructs a [PDFAnnotationStamp] from an unsafe.Pointer.
//
// A object allows you to display a word or phrase, such as “Confidential,” in a PDF page.
func PDFAnnotationStampFrom(ptr unsafe.Pointer) PDFAnnotationStamp {
	return PDFAnnotationStamp{
		PDFAnnotation: PDFAnnotationFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (pc _PDFAnnotationStampClass) Alloc() PDFAnnotationStamp {
	rv := objc.Send[PDFAnnotationStamp](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (pc _PDFAnnotationStampClass) New() PDFAnnotationStamp {
	rv := objc.Send[PDFAnnotationStamp](objc.ID(pc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (p_ PDFAnnotationStamp) Init() PDFAnnotationStamp {
	rv := objc.Send[PDFAnnotationStamp](p_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p_ PDFAnnotationStamp) Autorelease() PDFAnnotationStamp {
	rv := objc.Send[PDFAnnotationStamp](p_.ID, objc.Sel("autorelease"))
	return rv
}

// NewPDFAnnotationStamp creates a new PDFAnnotationStamp instance.
func NewPDFAnnotationStamp() PDFAnnotationStamp {
	return getPDFAnnotationStampClass().New()
}




