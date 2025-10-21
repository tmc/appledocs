// Code generated from Apple documentation for PDFKit. DO NOT EDIT.

package pdfkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [PDFAnnotationLink] class.
var (
	PDFAnnotationLinkClass     _PDFAnnotationLinkClass
	PDFAnnotationLinkClassOnce sync.Once
)

func getPDFAnnotationLinkClass() _PDFAnnotationLinkClass {
	PDFAnnotationLinkClassOnce.Do(func() {
		PDFAnnotationLinkClass = _PDFAnnotationLinkClass{objc.GetClass("PDFAnnotationLink")}
	})
	return PDFAnnotationLinkClass
}

type _PDFAnnotationLinkClass struct {
	class objc.Class
}

// An interface definition for the [PDFAnnotationLink] class.
type IPDFAnnotationLink interface {
	IPDFAnnotation
}

//
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFAnnotationLink
type PDFAnnotationLink struct {
	PDFAnnotation
}

// PDFAnnotationLinkFrom constructs a [PDFAnnotationLink] from an unsafe.Pointer.
func PDFAnnotationLinkFrom(ptr unsafe.Pointer) PDFAnnotationLink {
	return PDFAnnotationLink{
		PDFAnnotation: PDFAnnotationFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (pc _PDFAnnotationLinkClass) Alloc() PDFAnnotationLink {
	rv := objc.Send[PDFAnnotationLink](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (pc _PDFAnnotationLinkClass) New() PDFAnnotationLink {
	rv := objc.Send[PDFAnnotationLink](objc.ID(pc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (p_ PDFAnnotationLink) Init() PDFAnnotationLink {
	rv := objc.Send[PDFAnnotationLink](p_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p_ PDFAnnotationLink) Autorelease() PDFAnnotationLink {
	rv := objc.Send[PDFAnnotationLink](p_.ID, objc.Sel("autorelease"))
	return rv
}

// NewPDFAnnotationLink creates a new PDFAnnotationLink instance.
func NewPDFAnnotationLink() PDFAnnotationLink {
	return getPDFAnnotationLinkClass().New()
}




