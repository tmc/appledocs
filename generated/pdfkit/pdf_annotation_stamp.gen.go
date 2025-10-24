// Code generated from Apple documentation for PDFKit. DO NOT EDIT.

package pdfkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

/* debug [class.gen.go]: Generating class PDFAnnotationStamp */


/* debug [class_header]: Header for PDFAnnotationStamp */
// The class instance for the [PDFAnnotationStamp] class.
var (
	PDFAnnotationStampClass     _PDFAnnotationStampClass
	PDFAnnotationStampClassOnce sync.Once
)

func getPDFAnnotationStampClass() _PDFAnnotationStampClass {
	PDFAnnotationStampClassOnce.Do(func() {
		PDFAnnotationStampClass = _PDFAnnotationStampClass{objc.GetClass("PDFAnnotationStamp")}
	})
	return PDFAnnotationStampClass
}

type _PDFAnnotationStampClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for PDFAnnotationStamp */
// An interface definition for the [PDFAnnotationStamp] class.
type IPDFAnnotationStamp interface {
	IPDFAnnotation
	
/* debug [class_interface_properties]: Properties for PDFAnnotationStamp */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for PDFAnnotationStamp */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for PDFAnnotationStamp */
// Alloc allocates a new instance without initialization.
func (pc _PDFAnnotationStampClass) Alloc() PDFAnnotationStamp {
	rv := objc.Send[PDFAnnotationStamp](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (pc _PDFAnnotationStampClass) New() PDFAnnotationStamp {
	rv := objc.Send[PDFAnnotationStamp](objc.ID(pc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (p_ PDFAnnotationStamp) Init() PDFAnnotationStamp {
	rv := objc.Send[PDFAnnotationStamp](p_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p_ PDFAnnotationStamp) Autorelease() PDFAnnotationStamp {
	rv := objc.Send[PDFAnnotationStamp](p_.ID, objc.Sel("autorelease"))
	return rv
}

// NewPDFAnnotationStamp creates a new PDFAnnotationStamp instance.
func NewPDFAnnotationStamp() PDFAnnotationStamp {
	return getPDFAnnotationStampClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for PDFAnnotationStamp */
// A object allows you to display a word or phrase, such as “Confidential,” in a PDF page.
//
// A object should have an appearance stream associated with it; otherwise, nothing useful is rendered.


// A object allows you to display a word or phrase, such as “Confidential,” in a PDF page.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFAnnotationStamp
type PDFAnnotationStamp struct {
	PDFAnnotation
}

// PDFAnnotationStampFrom constructs a [PDFAnnotationStamp] from an unsafe.Pointer.
//
// A object allows you to display a word or phrase, such as “Confidential,” in a PDF page.
func PDFAnnotationStampFrom(ptr unsafe.Pointer) PDFAnnotationStamp {
	return PDFAnnotationStamp{
		PDFAnnotation: PDFAnnotationFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for PDFAnnotationStamp *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for PDFAnnotationStamp */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for PDFAnnotationStamp */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for PDFAnnotationStamp */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for PDFAnnotationStamp */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class PDFAnnotationStamp */



