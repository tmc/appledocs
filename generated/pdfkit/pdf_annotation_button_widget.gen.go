// Code generated from Apple documentation for PDFKit. DO NOT EDIT.

package pdfkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/appkit"
	"github.com/tmc/appledocs/generated/foundation"
)

/* debug [class.gen.go]: Generating class PDFAnnotationButtonWidget */


/* debug [class_header]: Header for PDFAnnotationButtonWidget */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for PDFAnnotationButtonWidget */
// An interface definition for the [PDFAnnotationButtonWidget] class.
type IPDFAnnotationButtonWidget interface {
	IPDFAnnotation
	
/* debug [class_interface_properties]: Properties for PDFAnnotationButtonWidget */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for PDFAnnotationButtonWidget */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for PDFAnnotationButtonWidget */
// Alloc allocates a new instance without initialization.
func (pc _PDFAnnotationButtonWidgetClass) Alloc() PDFAnnotationButtonWidget {
	rv := objc.Send[PDFAnnotationButtonWidget](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for PDFAnnotationButtonWidget */
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for PDFAnnotationButtonWidget *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for PDFAnnotationButtonWidget */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for PDFAnnotationButtonWidget */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for PDFAnnotationButtonWidget */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for PDFAnnotationButtonWidget */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class PDFAnnotationButtonWidget */



