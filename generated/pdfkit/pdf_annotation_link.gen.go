// Code generated from Apple documentation for PDFKit. DO NOT EDIT.

package pdfkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

/* debug [class.gen.go]: Generating class PDFAnnotationLink */


/* debug [class_header]: Header for PDFAnnotationLink */
// The class instance for the [PDFAnnotationLink] class.
var (
	PDFAnnotationLinkClass     _PDFAnnotationLinkClass
	PDFAnnotationLinkClassOnce sync.Once
)

func getPDFAnnotationLinkClass() _PDFAnnotationLinkClass {
	PDFAnnotationLinkClassOnce.Do(func() {
		PDFAnnotationLinkClass = _PDFAnnotationLinkClass{objc.GetClass("PDFAnnotationLink")}
	})
	return PDFAnnotationLinkClass
}

type _PDFAnnotationLinkClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for PDFAnnotationLink */
// An interface definition for the [PDFAnnotationLink] class.
type IPDFAnnotationLink interface {
	IPDFAnnotation
	
/* debug [class_interface_properties]: Properties for PDFAnnotationLink */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for PDFAnnotationLink */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for PDFAnnotationLink */
// Alloc allocates a new instance without initialization.
func (pc _PDFAnnotationLinkClass) Alloc() PDFAnnotationLink {
	rv := objc.Send[PDFAnnotationLink](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (pc _PDFAnnotationLinkClass) New() PDFAnnotationLink {
	rv := objc.Send[PDFAnnotationLink](objc.ID(pc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (p_ PDFAnnotationLink) Init() PDFAnnotationLink {
	rv := objc.Send[PDFAnnotationLink](p_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p_ PDFAnnotationLink) Autorelease() PDFAnnotationLink {
	rv := objc.Send[PDFAnnotationLink](p_.ID, objc.Sel("autorelease"))
	return rv
}

// NewPDFAnnotationLink creates a new PDFAnnotationLink instance.
func NewPDFAnnotationLink() PDFAnnotationLink {
	return getPDFAnnotationLinkClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for PDFAnnotationLink */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFAnnotationLink
type PDFAnnotationLink struct {
	PDFAnnotation
}

// PDFAnnotationLinkFrom constructs a [PDFAnnotationLink] from an unsafe.Pointer.
func PDFAnnotationLinkFrom(ptr unsafe.Pointer) PDFAnnotationLink {
	return PDFAnnotationLink{
		PDFAnnotation: PDFAnnotationFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for PDFAnnotationLink *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for PDFAnnotationLink */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for PDFAnnotationLink */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for PDFAnnotationLink */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for PDFAnnotationLink */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class PDFAnnotationLink */



