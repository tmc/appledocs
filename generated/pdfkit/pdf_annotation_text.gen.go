// Code generated from Apple documentation for PDFKit. DO NOT EDIT.

package pdfkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class PDFAnnotationText */


/* debug [class_header]: Header for PDFAnnotationText */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for PDFAnnotationText */
// An interface definition for the [PDFAnnotationText] class.
type IPDFAnnotationText interface {
	IPDFAnnotation
	
/* debug [class_interface_properties]: Properties for PDFAnnotationText */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for PDFAnnotationText */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for PDFAnnotationText */
// Alloc allocates a new instance without initialization.
func (pc _PDFAnnotationTextClass) Alloc() PDFAnnotationText {
	rv := objc.Send[PDFAnnotationText](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for PDFAnnotationText */
// A object displays as an icon (such as a “sticky note”) attached to a specified point in the PDF document.
//
// Each object has a object associated with it. In its closed state, the annotation appears as an icon. In its open state, it displays as a pop-up window containing the text of the note. Note that your application must do the work to put up a window containing the text in response to a . Currently, text annotations do not scale and rotate with the page.


// A object displays as an icon (such as a “sticky note”) attached to a specified point in the PDF document.
//
// [Full Topic]
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for PDFAnnotationText *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for PDFAnnotationText */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for PDFAnnotationText */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for PDFAnnotationText */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for PDFAnnotationText */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class PDFAnnotationText */



