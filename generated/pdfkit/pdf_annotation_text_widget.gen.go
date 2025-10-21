// Code generated from Apple documentation for PDFKit. DO NOT EDIT.

package pdfkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [PDFAnnotationTextWidget] class.
var (
	PDFAnnotationTextWidgetClass     _PDFAnnotationTextWidgetClass
	PDFAnnotationTextWidgetClassOnce sync.Once
)

func getPDFAnnotationTextWidgetClass() _PDFAnnotationTextWidgetClass {
	PDFAnnotationTextWidgetClassOnce.Do(func() {
		PDFAnnotationTextWidgetClass = _PDFAnnotationTextWidgetClass{objc.GetClass("PDFAnnotationTextWidget")}
	})
	return PDFAnnotationTextWidgetClass
}

type _PDFAnnotationTextWidgetClass struct {
	class objc.Class
}

// An interface definition for the [PDFAnnotationTextWidget] class.
type IPDFAnnotationTextWidget interface {
	IPDFAnnotation
}

// A object allows you to manage the appearance and content of text fields.
//
// objects support interactive forms in a PDF document. This object is comparable to an editable in Cocoa or an edit text view in Carbon.
//
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFAnnotationTextWidget
type PDFAnnotationTextWidget struct {
	PDFAnnotation
}

// PDFAnnotationTextWidgetFrom constructs a [PDFAnnotationTextWidget] from an unsafe.Pointer.
//
// A object allows you to manage the appearance and content of text fields.
func PDFAnnotationTextWidgetFrom(ptr unsafe.Pointer) PDFAnnotationTextWidget {
	return PDFAnnotationTextWidget{
		PDFAnnotation: PDFAnnotationFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (pc _PDFAnnotationTextWidgetClass) Alloc() PDFAnnotationTextWidget {
	rv := objc.Send[PDFAnnotationTextWidget](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (pc _PDFAnnotationTextWidgetClass) New() PDFAnnotationTextWidget {
	rv := objc.Send[PDFAnnotationTextWidget](objc.ID(pc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (p_ PDFAnnotationTextWidget) Init() PDFAnnotationTextWidget {
	rv := objc.Send[PDFAnnotationTextWidget](p_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p_ PDFAnnotationTextWidget) Autorelease() PDFAnnotationTextWidget {
	rv := objc.Send[PDFAnnotationTextWidget](p_.ID, objc.Sel("autorelease"))
	return rv
}

// NewPDFAnnotationTextWidget creates a new PDFAnnotationTextWidget instance.
func NewPDFAnnotationTextWidget() PDFAnnotationTextWidget {
	return getPDFAnnotationTextWidgetClass().New()
}




