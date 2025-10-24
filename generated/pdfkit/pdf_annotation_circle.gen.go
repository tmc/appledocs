// Code generated from Apple documentation for PDFKit. DO NOT EDIT.

package pdfkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/appkit"
)

/* debug [class.gen.go]: Generating class PDFAnnotationCircle */


/* debug [class_header]: Header for PDFAnnotationCircle */
// The class instance for the [PDFAnnotationCircle] class.
var (
	PDFAnnotationCircleClass     _PDFAnnotationCircleClass
	PDFAnnotationCircleClassOnce sync.Once
)

func getPDFAnnotationCircleClass() _PDFAnnotationCircleClass {
	PDFAnnotationCircleClassOnce.Do(func() {
		PDFAnnotationCircleClass = _PDFAnnotationCircleClass{objc.GetClass("PDFAnnotationCircle")}
	})
	return PDFAnnotationCircleClass
}

type _PDFAnnotationCircleClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for PDFAnnotationCircle */
// An interface definition for the [PDFAnnotationCircle] class.
type IPDFAnnotationCircle interface {
	IPDFAnnotation
	
/* debug [class_interface_properties]: Properties for PDFAnnotationCircle */
	// properties:
	Color() appkit.Color
	SetColor(value appkit.Color)
	LineWidth() float64
	SetLineWidth(value float64)
	Style() PDFBorderStyle
	SetStyle(value PDFBorderStyle)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for PDFAnnotationCircle */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for PDFAnnotationCircle */
// Alloc allocates a new instance without initialization.
func (pc _PDFAnnotationCircleClass) Alloc() PDFAnnotationCircle {
	rv := objc.Send[PDFAnnotationCircle](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (pc _PDFAnnotationCircleClass) New() PDFAnnotationCircle {
	rv := objc.Send[PDFAnnotationCircle](objc.ID(pc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (p_ PDFAnnotationCircle) Init() PDFAnnotationCircle {
	rv := objc.Send[PDFAnnotationCircle](p_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p_ PDFAnnotationCircle) Autorelease() PDFAnnotationCircle {
	rv := objc.Send[PDFAnnotationCircle](p_.ID, objc.Sel("autorelease"))
	return rv
}

// NewPDFAnnotationCircle creates a new PDFAnnotationCircle instance.
func NewPDFAnnotationCircle() PDFAnnotationCircle {
	return getPDFAnnotationCircleClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for PDFAnnotationCircle */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFAnnotationCircle
type PDFAnnotationCircle struct {
	PDFAnnotation
}

// PDFAnnotationCircleFrom constructs a [PDFAnnotationCircle] from an unsafe.Pointer.
func PDFAnnotationCircleFrom(ptr unsafe.Pointer) PDFAnnotationCircle {
	return PDFAnnotationCircle{
		PDFAnnotation: PDFAnnotationFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for PDFAnnotationCircle *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for PDFAnnotationCircle */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for PDFAnnotationCircle */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for PDFAnnotationCircle */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for PDFAnnotationCircle */

// Sets the stroke color for the annotation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/pdfkit/pdfannotation/color
func (p_ PDFAnnotationCircle) Color() appkit.Color {
	rv := objc.Send[appkit.Color](p_.ID, objc.Sel("color"))
	return rv
}/* debug [instance_properties/getter]: color */


// Sets the stroke color for the annotation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/pdfkit/pdfannotation/color
func (p_ PDFAnnotationCircle) SetColor(value appkit.Color) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setColor:"), value)
}/* debug [instance_properties/setter]: color */


// Sets the line width (in points) for the border.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/pdfkit/pdfborder/linewidth
func (p_ PDFAnnotationCircle) LineWidth() float64 {
	rv := objc.Send[float64](p_.ID, objc.Sel("lineWidth"))
	return rv
}/* debug [instance_properties/getter]: lineWidth */


// Sets the line width (in points) for the border.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/pdfkit/pdfborder/linewidth
func (p_ PDFAnnotationCircle) SetLineWidth(value float64) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setLineWidth:"), value)
}/* debug [instance_properties/setter]: lineWidth */


// Sets the border style.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/pdfkit/pdfborder/style
func (p_ PDFAnnotationCircle) Style() PDFBorderStyle {
	rv := objc.Send[PDFBorderStyle](p_.ID, objc.Sel("style"))
	return rv
}/* debug [instance_properties/getter]: style */


// Sets the border style.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/pdfkit/pdfborder/style
func (p_ PDFAnnotationCircle) SetStyle(value PDFBorderStyle) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setStyle:"), value)
}/* debug [instance_properties/setter]: style */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class PDFAnnotationCircle */



