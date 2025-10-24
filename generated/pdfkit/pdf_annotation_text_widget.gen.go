// Code generated from Apple documentation for PDFKit. DO NOT EDIT.

package pdfkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/appkit"
	"github.com/tmc/appledocs/generated/foundation"
)

/* debug [class.gen.go]: Generating class PDFAnnotationTextWidget */


/* debug [class_header]: Header for PDFAnnotationTextWidget */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for PDFAnnotationTextWidget */
// An interface definition for the [PDFAnnotationTextWidget] class.
type IPDFAnnotationTextWidget interface {
	IPDFAnnotation
	
/* debug [class_interface_properties]: Properties for PDFAnnotationTextWidget */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for PDFAnnotationTextWidget */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for PDFAnnotationTextWidget */
// Alloc allocates a new instance without initialization.
func (pc _PDFAnnotationTextWidgetClass) Alloc() PDFAnnotationTextWidget {
	rv := objc.Send[PDFAnnotationTextWidget](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for PDFAnnotationTextWidget */
// A object allows you to manage the appearance and content of text fields.
//
// objects support interactive forms in a PDF document. This object is comparable to an editable in Cocoa or an edit text view in Carbon.


// A object allows you to manage the appearance and content of text fields.
//
// [Full Topic]
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for PDFAnnotationTextWidget *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for PDFAnnotationTextWidget */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for PDFAnnotationTextWidget */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for PDFAnnotationTextWidget */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for PDFAnnotationTextWidget */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class PDFAnnotationTextWidget */



