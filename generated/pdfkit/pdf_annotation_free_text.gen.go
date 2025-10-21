// Code generated from Apple documentation for PDFKit. DO NOT EDIT.

package pdfkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [PDFAnnotationFreeText] class.
var (
	PDFAnnotationFreeTextClass     _PDFAnnotationFreeTextClass
	PDFAnnotationFreeTextClassOnce sync.Once
)

func getPDFAnnotationFreeTextClass() _PDFAnnotationFreeTextClass {
	PDFAnnotationFreeTextClassOnce.Do(func() {
		PDFAnnotationFreeTextClass = _PDFAnnotationFreeTextClass{objc.GetClass("PDFAnnotationFreeText")}
	})
	return PDFAnnotationFreeTextClass
}

type _PDFAnnotationFreeTextClass struct {
	class objc.Class
}

// An interface definition for the [PDFAnnotationFreeText] class.
type IPDFAnnotationFreeText interface {
	IPDFAnnotation
}

// A object displays text on a page.
//
// Unlike a object, a object has no open or closed state; its text is always visible. The text annotation performed in Preview uses . The class’s property lets you get and set the textual content for a object.
//
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFAnnotationFreeText
type PDFAnnotationFreeText struct {
	PDFAnnotation
}

// PDFAnnotationFreeTextFrom constructs a [PDFAnnotationFreeText] from an unsafe.Pointer.
//
// A object displays text on a page.
func PDFAnnotationFreeTextFrom(ptr unsafe.Pointer) PDFAnnotationFreeText {
	return PDFAnnotationFreeText{
		PDFAnnotation: PDFAnnotationFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (pc _PDFAnnotationFreeTextClass) Alloc() PDFAnnotationFreeText {
	rv := objc.Send[PDFAnnotationFreeText](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (pc _PDFAnnotationFreeTextClass) New() PDFAnnotationFreeText {
	rv := objc.Send[PDFAnnotationFreeText](objc.ID(pc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (p_ PDFAnnotationFreeText) Init() PDFAnnotationFreeText {
	rv := objc.Send[PDFAnnotationFreeText](p_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p_ PDFAnnotationFreeText) Autorelease() PDFAnnotationFreeText {
	rv := objc.Send[PDFAnnotationFreeText](p_.ID, objc.Sel("autorelease"))
	return rv
}

// NewPDFAnnotationFreeText creates a new PDFAnnotationFreeText instance.
func NewPDFAnnotationFreeText() PDFAnnotationFreeText {
	return getPDFAnnotationFreeTextClass().New()
}




