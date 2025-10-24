// Code generated from Apple documentation for PDFKit. DO NOT EDIT.

package pdfkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/appkit"
)

/* debug [class.gen.go]: Generating class PDFAnnotationSquare */


/* debug [class_header]: Header for PDFAnnotationSquare */
// The class instance for the [PDFAnnotationSquare] class.
var (
	PDFAnnotationSquareClass     _PDFAnnotationSquareClass
	PDFAnnotationSquareClassOnce sync.Once
)

func getPDFAnnotationSquareClass() _PDFAnnotationSquareClass {
	PDFAnnotationSquareClassOnce.Do(func() {
		PDFAnnotationSquareClass = _PDFAnnotationSquareClass{objc.GetClass("PDFAnnotationSquare")}
	})
	return PDFAnnotationSquareClass
}

type _PDFAnnotationSquareClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for PDFAnnotationSquare */
// An interface definition for the [PDFAnnotationSquare] class.
type IPDFAnnotationSquare interface {
	IPDFAnnotation
	
/* debug [class_interface_properties]: Properties for PDFAnnotationSquare */
	// properties:
	Color() appkit.Color
	SetColor(value appkit.Color)
	LineWidth() float64
	SetLineWidth(value float64)
	Style() PDFBorderStyle
	SetStyle(value PDFBorderStyle)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for PDFAnnotationSquare */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for PDFAnnotationSquare */
// Alloc allocates a new instance without initialization.
func (pc _PDFAnnotationSquareClass) Alloc() PDFAnnotationSquare {
	rv := objc.Send[PDFAnnotationSquare](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (pc _PDFAnnotationSquareClass) New() PDFAnnotationSquare {
	rv := objc.Send[PDFAnnotationSquare](objc.ID(pc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (p_ PDFAnnotationSquare) Init() PDFAnnotationSquare {
	rv := objc.Send[PDFAnnotationSquare](p_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p_ PDFAnnotationSquare) Autorelease() PDFAnnotationSquare {
	rv := objc.Send[PDFAnnotationSquare](p_.ID, objc.Sel("autorelease"))
	return rv
}

// NewPDFAnnotationSquare creates a new PDFAnnotationSquare instance.
func NewPDFAnnotationSquare() PDFAnnotationSquare {
	return getPDFAnnotationSquareClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for PDFAnnotationSquare */
// A rectangle annotation on a page.


// A rectangle annotation on a page.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFAnnotationSquare
type PDFAnnotationSquare struct {
	PDFAnnotation
}

// PDFAnnotationSquareFrom constructs a [PDFAnnotationSquare] from an unsafe.Pointer.
//
// A rectangle annotation on a page.
func PDFAnnotationSquareFrom(ptr unsafe.Pointer) PDFAnnotationSquare {
	return PDFAnnotationSquare{
		PDFAnnotation: PDFAnnotationFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for PDFAnnotationSquare *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for PDFAnnotationSquare */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for PDFAnnotationSquare */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for PDFAnnotationSquare */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for PDFAnnotationSquare */

// Sets the stroke color for the annotation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/pdfkit/pdfannotation/color
func (p_ PDFAnnotationSquare) Color() appkit.Color {
	rv := objc.Send[appkit.Color](p_.ID, objc.Sel("color"))
	return rv
}/* debug [instance_properties/getter]: color */


// Sets the stroke color for the annotation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/pdfkit/pdfannotation/color
func (p_ PDFAnnotationSquare) SetColor(value appkit.Color) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setColor:"), value)
}/* debug [instance_properties/setter]: color */


// Sets the line width (in points) for the border.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/pdfkit/pdfborder/linewidth
func (p_ PDFAnnotationSquare) LineWidth() float64 {
	rv := objc.Send[float64](p_.ID, objc.Sel("lineWidth"))
	return rv
}/* debug [instance_properties/getter]: lineWidth */


// Sets the line width (in points) for the border.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/pdfkit/pdfborder/linewidth
func (p_ PDFAnnotationSquare) SetLineWidth(value float64) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setLineWidth:"), value)
}/* debug [instance_properties/setter]: lineWidth */


// Sets the border style.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/pdfkit/pdfborder/style
func (p_ PDFAnnotationSquare) Style() PDFBorderStyle {
	rv := objc.Send[PDFBorderStyle](p_.ID, objc.Sel("style"))
	return rv
}/* debug [instance_properties/getter]: style */


// Sets the border style.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/pdfkit/pdfborder/style
func (p_ PDFAnnotationSquare) SetStyle(value PDFBorderStyle) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setStyle:"), value)
}/* debug [instance_properties/setter]: style */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class PDFAnnotationSquare */



