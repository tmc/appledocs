// Code generated from Apple documentation for PDFKit. DO NOT EDIT.

package pdfkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [PDFAnnotationText] class.
var (
	PDFAnnotationTextClass     _PDFAnnotationTextClass
	PDFAnnotationTextClassOnce sync.Once
)

func getPDFAnnotationTextClass() _PDFAnnotationTextClass {
	PDFAnnotationTextClassOnce.Do(func() {
		PDFAnnotationTextClass = _PDFAnnotationTextClass{objc.GetClass("PDFAnnotationText")}
	})
	return PDFAnnotationTextClass
}

type _PDFAnnotationTextClass struct {
	class objc.Class
}

// An interface definition for the [PDFAnnotationText] class.
type IPDFAnnotationText interface {
	IPDFAnnotation
}

// A object displays as an icon (such as a “sticky note”) attached to a specified point in the PDF document.
//
// Each object has a object associated with it. In its closed state, the annotation appears as an icon. In its open state, it displays as a pop-up window containing the text of the note. Note that your application must do the work to put up a window containing the text in response to a . Currently, text annotations do not scale and rotate with the page.
//
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFAnnotationText
type PDFAnnotationText struct {
	PDFAnnotation
}

// PDFAnnotationTextFrom constructs a [PDFAnnotationText] from an unsafe.Pointer.
//
// A object displays as an icon (such as a “sticky note”) attached to a specified point in the PDF document.
func PDFAnnotationTextFrom(ptr unsafe.Pointer) PDFAnnotationText {
	return PDFAnnotationText{
		PDFAnnotation: PDFAnnotationFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (pc _PDFAnnotationTextClass) Alloc() PDFAnnotationText {
	rv := objc.Send[PDFAnnotationText](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (pc _PDFAnnotationTextClass) New() PDFAnnotationText {
	rv := objc.Send[PDFAnnotationText](objc.ID(pc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (p_ PDFAnnotationText) Init() PDFAnnotationText {
	rv := objc.Send[PDFAnnotationText](p_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p_ PDFAnnotationText) Autorelease() PDFAnnotationText {
	rv := objc.Send[PDFAnnotationText](p_.ID, objc.Sel("autorelease"))
	return rv
}

// NewPDFAnnotationText creates a new PDFAnnotationText instance.
func NewPDFAnnotationText() PDFAnnotationText {
	return getPDFAnnotationTextClass().New()
}




