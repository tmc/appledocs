// Code generated from Apple documentation for PDFKit. DO NOT EDIT.

package pdfkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/appkit"
	"github.com/tmc/appledocs/generated/foundation"
)

/* debug [class.gen.go]: Generating class PDFAnnotationChoiceWidget */


/* debug [class_header]: Header for PDFAnnotationChoiceWidget */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for PDFAnnotationChoiceWidget */
// An interface definition for the [PDFAnnotationChoiceWidget] class.
type IPDFAnnotationChoiceWidget interface {
	IPDFAnnotation
	
/* debug [class_interface_properties]: Properties for PDFAnnotationChoiceWidget */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for PDFAnnotationChoiceWidget */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for PDFAnnotationChoiceWidget */
// Alloc allocates a new instance without initialization.
func (pc _PDFAnnotationChoiceWidgetClass) Alloc() PDFAnnotationChoiceWidget {
	rv := objc.Send[PDFAnnotationChoiceWidget](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for PDFAnnotationChoiceWidget */
// A object provides user interactivity on a page of a PDF document, in the form of pop-up menus and lists.
//
// inherits general annotation behavior from the class. If you use a object, your application must handle hit testing, unless you are simply using to display content. This is because automatically handles hit testing for you.


// A object provides user interactivity on a page of a PDF document, in the form of pop-up menus and lists.
//
// [Full Topic]
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for PDFAnnotationChoiceWidget *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for PDFAnnotationChoiceWidget */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for PDFAnnotationChoiceWidget */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for PDFAnnotationChoiceWidget */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for PDFAnnotationChoiceWidget */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class PDFAnnotationChoiceWidget */



