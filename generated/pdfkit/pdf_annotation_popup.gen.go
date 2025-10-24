// Code generated from Apple documentation for PDFKit. DO NOT EDIT.

package pdfkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class PDFAnnotationPopup */


/* debug [class_header]: Header for PDFAnnotationPopup */
// The class instance for the [PDFAnnotationPopup] class.
var (
	PDFAnnotationPopupClass     _PDFAnnotationPopupClass
	PDFAnnotationPopupClassOnce sync.Once
)

func getPDFAnnotationPopupClass() _PDFAnnotationPopupClass {
	PDFAnnotationPopupClassOnce.Do(func() {
		PDFAnnotationPopupClass = _PDFAnnotationPopupClass{objc.GetClass("PDFAnnotationPopup")}
	})
	return PDFAnnotationPopupClass
}

type _PDFAnnotationPopupClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for PDFAnnotationPopup */
// An interface definition for the [PDFAnnotationPopup] class.
type IPDFAnnotationPopup interface {
	IPDFAnnotation
	
/* debug [class_interface_properties]: Properties for PDFAnnotationPopup */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for PDFAnnotationPopup */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for PDFAnnotationPopup */
// Alloc allocates a new instance without initialization.
func (pc _PDFAnnotationPopupClass) Alloc() PDFAnnotationPopup {
	rv := objc.Send[PDFAnnotationPopup](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (pc _PDFAnnotationPopupClass) New() PDFAnnotationPopup {
	rv := objc.Send[PDFAnnotationPopup](objc.ID(pc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (p_ PDFAnnotationPopup) Init() PDFAnnotationPopup {
	rv := objc.Send[PDFAnnotationPopup](p_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p_ PDFAnnotationPopup) Autorelease() PDFAnnotationPopup {
	rv := objc.Send[PDFAnnotationPopup](p_.ID, objc.Sel("autorelease"))
	return rv
}

// NewPDFAnnotationPopup creates a new PDFAnnotationPopup instance.
func NewPDFAnnotationPopup() PDFAnnotationPopup {
	return getPDFAnnotationPopupClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for PDFAnnotationPopup */
// A object provides user interactivity on a PDF page in the form of a pop-up menu.


// A object provides user interactivity on a PDF page in the form of a pop-up menu.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFAnnotationPopup
type PDFAnnotationPopup struct {
	PDFAnnotation
}

// PDFAnnotationPopupFrom constructs a [PDFAnnotationPopup] from an unsafe.Pointer.
//
// A object provides user interactivity on a PDF page in the form of a pop-up menu.
func PDFAnnotationPopupFrom(ptr unsafe.Pointer) PDFAnnotationPopup {
	return PDFAnnotationPopup{
		PDFAnnotation: PDFAnnotationFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for PDFAnnotationPopup *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for PDFAnnotationPopup */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for PDFAnnotationPopup */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for PDFAnnotationPopup */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for PDFAnnotationPopup */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class PDFAnnotationPopup */



