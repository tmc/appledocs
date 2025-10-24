// Code generated from Apple documentation for PDFKit. DO NOT EDIT.

package pdfkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/appkit"
	"github.com/tmc/appledocs/generated/foundation"
)

/* debug [class.gen.go]: Generating class PDFAnnotationInk */


/* debug [class_header]: Header for PDFAnnotationInk */
// The class instance for the [PDFAnnotationInk] class.
var (
	PDFAnnotationInkClass     _PDFAnnotationInkClass
	PDFAnnotationInkClassOnce sync.Once
)

func getPDFAnnotationInkClass() _PDFAnnotationInkClass {
	PDFAnnotationInkClassOnce.Do(func() {
		PDFAnnotationInkClass = _PDFAnnotationInkClass{objc.GetClass("PDFAnnotationInk")}
	})
	return PDFAnnotationInkClass
}

type _PDFAnnotationInkClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for PDFAnnotationInk */
// An interface definition for the [PDFAnnotationInk] class.
type IPDFAnnotationInk interface {
	IPDFAnnotation
	
/* debug [class_interface_properties]: Properties for PDFAnnotationInk */
	// properties:
	Color() appkit.Color
	SetColor(value appkit.Color)
	LineWidth() float64
	SetLineWidth(value float64)
	Style() PDFBorderStyle
	SetStyle(value PDFBorderStyle)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for PDFAnnotationInk */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for PDFAnnotationInk */
// Alloc allocates a new instance without initialization.
func (pc _PDFAnnotationInkClass) Alloc() PDFAnnotationInk {
	rv := objc.Send[PDFAnnotationInk](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (pc _PDFAnnotationInkClass) New() PDFAnnotationInk {
	rv := objc.Send[PDFAnnotationInk](objc.ID(pc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (p_ PDFAnnotationInk) Init() PDFAnnotationInk {
	rv := objc.Send[PDFAnnotationInk](p_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p_ PDFAnnotationInk) Autorelease() PDFAnnotationInk {
	rv := objc.Send[PDFAnnotationInk](p_.ID, objc.Sel("autorelease"))
	return rv
}

// NewPDFAnnotationInk creates a new PDFAnnotationInk instance.
func NewPDFAnnotationInk() PDFAnnotationInk {
	return getPDFAnnotationInkClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for PDFAnnotationInk */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFAnnotationInk
type PDFAnnotationInk struct {
	PDFAnnotation
}

// PDFAnnotationInkFrom constructs a [PDFAnnotationInk] from an unsafe.Pointer.
func PDFAnnotationInkFrom(ptr unsafe.Pointer) PDFAnnotationInk {
	return PDFAnnotationInk{
		PDFAnnotation: PDFAnnotationFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for PDFAnnotationInk *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for PDFAnnotationInk */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for PDFAnnotationInk */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for PDFAnnotationInk */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for PDFAnnotationInk */

// Sets the stroke color for the annotation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/pdfkit/pdfannotation/color
func (p_ PDFAnnotationInk) Color() appkit.Color {
	rv := objc.Send[appkit.Color](p_.ID, objc.Sel("color"))
	return rv
}/* debug [instance_properties/getter]: color */


// Sets the stroke color for the annotation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/pdfkit/pdfannotation/color
func (p_ PDFAnnotationInk) SetColor(value appkit.Color) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setColor:"), value)
}/* debug [instance_properties/setter]: color */


// Sets the line width (in points) for the border.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/pdfkit/pdfborder/linewidth
func (p_ PDFAnnotationInk) LineWidth() float64 {
	rv := objc.Send[float64](p_.ID, objc.Sel("lineWidth"))
	return rv
}/* debug [instance_properties/getter]: lineWidth */


// Sets the line width (in points) for the border.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/pdfkit/pdfborder/linewidth
func (p_ PDFAnnotationInk) SetLineWidth(value float64) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setLineWidth:"), value)
}/* debug [instance_properties/setter]: lineWidth */


// Sets the border style.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/pdfkit/pdfborder/style
func (p_ PDFAnnotationInk) Style() PDFBorderStyle {
	rv := objc.Send[PDFBorderStyle](p_.ID, objc.Sel("style"))
	return rv
}/* debug [instance_properties/getter]: style */


// Sets the border style.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/pdfkit/pdfborder/style
func (p_ PDFAnnotationInk) SetStyle(value PDFBorderStyle) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setStyle:"), value)
}/* debug [instance_properties/setter]: style */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class PDFAnnotationInk */



