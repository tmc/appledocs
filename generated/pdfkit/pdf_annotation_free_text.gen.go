// Code generated from Apple documentation for PDFKit. DO NOT EDIT.

package pdfkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/appkit"
	"github.com/tmc/appledocs/generated/foundation"
)

/* debug [class.gen.go]: Generating class PDFAnnotationFreeText */


/* debug [class_header]: Header for PDFAnnotationFreeText */
// The class instance for the [PDFAnnotationFreeText] class.
var (
	PDFAnnotationFreeTextClass     _PDFAnnotationFreeTextClass
	PDFAnnotationFreeTextClassOnce sync.Once
)

func getPDFAnnotationFreeTextClass() _PDFAnnotationFreeTextClass {
	PDFAnnotationFreeTextClassOnce.Do(func() {
		PDFAnnotationFreeTextClass = _PDFAnnotationFreeTextClass{objc.GetClass("PDFAnnotationFreeText")}
	})
	return PDFAnnotationFreeTextClass
}

type _PDFAnnotationFreeTextClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for PDFAnnotationFreeText */
// An interface definition for the [PDFAnnotationFreeText] class.
type IPDFAnnotationFreeText interface {
	IPDFAnnotation
	
/* debug [class_interface_properties]: Properties for PDFAnnotationFreeText */
	// properties:
	Contents() objc.IObject /* cross-framework: NSString */
	SetContents(value objc.IObject /* cross-framework: NSString */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for PDFAnnotationFreeText */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for PDFAnnotationFreeText */
// Alloc allocates a new instance without initialization.
func (pc _PDFAnnotationFreeTextClass) Alloc() PDFAnnotationFreeText {
	rv := objc.Send[PDFAnnotationFreeText](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (pc _PDFAnnotationFreeTextClass) New() PDFAnnotationFreeText {
	rv := objc.Send[PDFAnnotationFreeText](objc.ID(pc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (p_ PDFAnnotationFreeText) Init() PDFAnnotationFreeText {
	rv := objc.Send[PDFAnnotationFreeText](p_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p_ PDFAnnotationFreeText) Autorelease() PDFAnnotationFreeText {
	rv := objc.Send[PDFAnnotationFreeText](p_.ID, objc.Sel("autorelease"))
	return rv
}

// NewPDFAnnotationFreeText creates a new PDFAnnotationFreeText instance.
func NewPDFAnnotationFreeText() PDFAnnotationFreeText {
	return getPDFAnnotationFreeTextClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for PDFAnnotationFreeText */
// A object displays text on a page.
//
// Unlike a object, a object has no open or closed state; its text is always visible. The text annotation performed in Preview uses . The class’s property lets you get and set the textual content for a object.


// A object displays text on a page.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFAnnotationFreeText
type PDFAnnotationFreeText struct {
	PDFAnnotation
}

// PDFAnnotationFreeTextFrom constructs a [PDFAnnotationFreeText] from an unsafe.Pointer.
//
// A object displays text on a page.
func PDFAnnotationFreeTextFrom(ptr unsafe.Pointer) PDFAnnotationFreeText {
	return PDFAnnotationFreeText{
		PDFAnnotation: PDFAnnotationFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for PDFAnnotationFreeText *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for PDFAnnotationFreeText */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for PDFAnnotationFreeText */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for PDFAnnotationFreeText */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for PDFAnnotationFreeText */

// Returns the textual content (if any) associated with the annotation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/pdfkit/pdfannotation/contents
func (p_ PDFAnnotationFreeText) Contents() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](p_.ID, objc.Sel("contents"))
	return rv
}/* debug [instance_properties/getter]: contents */


// Returns the textual content (if any) associated with the annotation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/pdfkit/pdfannotation/contents
func (p_ PDFAnnotationFreeText) SetContents(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setContents:"), value)
}/* debug [instance_properties/setter]: contents */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class PDFAnnotationFreeText */



