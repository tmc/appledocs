// Code generated from Apple documentation for PDFKit. DO NOT EDIT.

package pdfkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [PDFAnnotationChoiceWidget] class.
var (
	PDFAnnotationChoiceWidgetClass     _PDFAnnotationChoiceWidgetClass
	PDFAnnotationChoiceWidgetClassOnce sync.Once
)

func getPDFAnnotationChoiceWidgetClass() _PDFAnnotationChoiceWidgetClass {
	PDFAnnotationChoiceWidgetClassOnce.Do(func() {
		PDFAnnotationChoiceWidgetClass = _PDFAnnotationChoiceWidgetClass{objc.GetClass("PDFAnnotationChoiceWidget")}
	})
	return PDFAnnotationChoiceWidgetClass
}

type _PDFAnnotationChoiceWidgetClass struct {
	class objc.Class
}

// An interface definition for the [PDFAnnotationChoiceWidget] class.
type IPDFAnnotationChoiceWidget interface {
	IPDFAnnotation
}

// A object provides user interactivity on a page of a PDF document, in the form of pop-up menus and lists.
//
// inherits general annotation behavior from the class. If you use a object, your application must handle hit testing, unless you are simply using to display content. This is because automatically handles hit testing for you.
//
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFAnnotationChoiceWidget
type PDFAnnotationChoiceWidget struct {
	PDFAnnotation
}

// PDFAnnotationChoiceWidgetFrom constructs a [PDFAnnotationChoiceWidget] from an unsafe.Pointer.
//
// A object provides user interactivity on a page of a PDF document, in the form of pop-up menus and lists.
func PDFAnnotationChoiceWidgetFrom(ptr unsafe.Pointer) PDFAnnotationChoiceWidget {
	return PDFAnnotationChoiceWidget{
		PDFAnnotation: PDFAnnotationFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (pc _PDFAnnotationChoiceWidgetClass) Alloc() PDFAnnotationChoiceWidget {
	rv := objc.Send[PDFAnnotationChoiceWidget](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (pc _PDFAnnotationChoiceWidgetClass) New() PDFAnnotationChoiceWidget {
	rv := objc.Send[PDFAnnotationChoiceWidget](objc.ID(pc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (p_ PDFAnnotationChoiceWidget) Init() PDFAnnotationChoiceWidget {
	rv := objc.Send[PDFAnnotationChoiceWidget](p_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p_ PDFAnnotationChoiceWidget) Autorelease() PDFAnnotationChoiceWidget {
	rv := objc.Send[PDFAnnotationChoiceWidget](p_.ID, objc.Sel("autorelease"))
	return rv
}

// NewPDFAnnotationChoiceWidget creates a new PDFAnnotationChoiceWidget instance.
func NewPDFAnnotationChoiceWidget() PDFAnnotationChoiceWidget {
	return getPDFAnnotationChoiceWidgetClass().New()
}




