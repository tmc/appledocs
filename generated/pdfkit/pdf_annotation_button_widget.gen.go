// Code generated from Apple documentation for PDFKit. DO NOT EDIT.

package pdfkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [PDFAnnotationButtonWidget] class.
var (
	PDFAnnotationButtonWidgetClass     _PDFAnnotationButtonWidgetClass
	PDFAnnotationButtonWidgetClassOnce sync.Once
)

func getPDFAnnotationButtonWidgetClass() _PDFAnnotationButtonWidgetClass {
	PDFAnnotationButtonWidgetClassOnce.Do(func() {
		PDFAnnotationButtonWidgetClass = _PDFAnnotationButtonWidgetClass{objc.GetClass("PDFAnnotationButtonWidget")}
	})
	return PDFAnnotationButtonWidgetClass
}

type _PDFAnnotationButtonWidgetClass struct {
	class objc.Class
}

// An interface definition for the [PDFAnnotationButtonWidget] class.
type IPDFAnnotationButtonWidget interface {
	IPDFAnnotation
	// properties:
	// methods:
}

// A object provides user interactivity on a page of a PDF document. There are three types of buttons available: push button, radio button, and checkbox.
//
// inherits general annotation behavior from the class. If you use a object, your application must handle hit testing, unless you are simply using to display content. This is because automatically handles hit testing for you.


// A object provides user interactivity on a page of a PDF document. There are three types of buttons available: push button, radio button, and checkbox.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFAnnotationButtonWidget
type PDFAnnotationButtonWidget struct {
	PDFAnnotation
}

// PDFAnnotationButtonWidgetFrom constructs a [PDFAnnotationButtonWidget] from an unsafe.Pointer.
//
// A object provides user interactivity on a page of a PDF document. There are three types of buttons available: push button, radio button, and checkbox.
func PDFAnnotationButtonWidgetFrom(ptr unsafe.Pointer) PDFAnnotationButtonWidget {
	return PDFAnnotationButtonWidget{
		PDFAnnotation: PDFAnnotationFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (pc _PDFAnnotationButtonWidgetClass) Alloc() PDFAnnotationButtonWidget {
	rv := objc.Send[PDFAnnotationButtonWidget](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (pc _PDFAnnotationButtonWidgetClass) New() PDFAnnotationButtonWidget {
	rv := objc.Send[PDFAnnotationButtonWidget](objc.ID(pc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (p_ PDFAnnotationButtonWidget) Init() PDFAnnotationButtonWidget {
	rv := objc.Send[PDFAnnotationButtonWidget](p_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p_ PDFAnnotationButtonWidget) Autorelease() PDFAnnotationButtonWidget {
	rv := objc.Send[PDFAnnotationButtonWidget](p_.ID, objc.Sel("autorelease"))
	return rv
}

// NewPDFAnnotationButtonWidget creates a new PDFAnnotationButtonWidget instance.
func NewPDFAnnotationButtonWidget() PDFAnnotationButtonWidget {
	return getPDFAnnotationButtonWidgetClass().New()
}




