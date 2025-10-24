// Code generated from Apple documentation for PDFKit. DO NOT EDIT.

package pdfkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/appkit"
	"github.com/tmc/appledocs/generated/vision"
)

/* debug [class.gen.go]: Generating class PDFAnnotationLine */


/* debug [class_header]: Header for PDFAnnotationLine */
// The class instance for the [PDFAnnotationLine] class.
var (
	PDFAnnotationLineClass     _PDFAnnotationLineClass
	PDFAnnotationLineClassOnce sync.Once
)

func getPDFAnnotationLineClass() _PDFAnnotationLineClass {
	PDFAnnotationLineClassOnce.Do(func() {
		PDFAnnotationLineClass = _PDFAnnotationLineClass{objc.GetClass("PDFAnnotationLine")}
	})
	return PDFAnnotationLineClass
}

type _PDFAnnotationLineClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for PDFAnnotationLine */
// An interface definition for the [PDFAnnotationLine] class.
type IPDFAnnotationLine interface {
	IPDFAnnotation
	
/* debug [class_interface_properties]: Properties for PDFAnnotationLine */
	// properties:
	Color() appkit.Color
	SetColor(value appkit.Color)
	LineWidth() float64
	SetLineWidth(value float64)
	Style() PDFBorderStyle
	SetStyle(value PDFBorderStyle)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for PDFAnnotationLine */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for PDFAnnotationLine */
// Alloc allocates a new instance without initialization.
func (pc _PDFAnnotationLineClass) Alloc() PDFAnnotationLine {
	rv := objc.Send[PDFAnnotationLine](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (pc _PDFAnnotationLineClass) New() PDFAnnotationLine {
	rv := objc.Send[PDFAnnotationLine](objc.ID(pc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (p_ PDFAnnotationLine) Init() PDFAnnotationLine {
	rv := objc.Send[PDFAnnotationLine](p_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p_ PDFAnnotationLine) Autorelease() PDFAnnotationLine {
	rv := objc.Send[PDFAnnotationLine](p_.ID, objc.Sel("autorelease"))
	return rv
}

// NewPDFAnnotationLine creates a new PDFAnnotationLine instance.
func NewPDFAnnotationLine() PDFAnnotationLine {
	return getPDFAnnotationLineClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for PDFAnnotationLine */
// A object displays a single line on a page.
//
// The and properties of the annotation’s associated object determines the stroke thickness and style. The property of the class determines the stroke color.


// A object displays a single line on a page.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFAnnotationLine
type PDFAnnotationLine struct {
	PDFAnnotation
}

// PDFAnnotationLineFrom constructs a [PDFAnnotationLine] from an unsafe.Pointer.
//
// A object displays a single line on a page.
func PDFAnnotationLineFrom(ptr unsafe.Pointer) PDFAnnotationLine {
	return PDFAnnotationLine{
		PDFAnnotation: PDFAnnotationFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for PDFAnnotationLine *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for PDFAnnotationLine */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for PDFAnnotationLine */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for PDFAnnotationLine */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for PDFAnnotationLine */

// Sets the stroke color for the annotation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/pdfkit/pdfannotation/color
func (p_ PDFAnnotationLine) Color() appkit.Color {
	rv := objc.Send[appkit.Color](p_.ID, objc.Sel("color"))
	return rv
}/* debug [instance_properties/getter]: color */


// Sets the stroke color for the annotation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/pdfkit/pdfannotation/color
func (p_ PDFAnnotationLine) SetColor(value appkit.Color) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setColor:"), value)
}/* debug [instance_properties/setter]: color */


// Sets the line width (in points) for the border.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/pdfkit/pdfborder/linewidth
func (p_ PDFAnnotationLine) LineWidth() float64 {
	rv := objc.Send[float64](p_.ID, objc.Sel("lineWidth"))
	return rv
}/* debug [instance_properties/getter]: lineWidth */


// Sets the line width (in points) for the border.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/pdfkit/pdfborder/linewidth
func (p_ PDFAnnotationLine) SetLineWidth(value float64) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setLineWidth:"), value)
}/* debug [instance_properties/setter]: lineWidth */


// Sets the border style.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/pdfkit/pdfborder/style
func (p_ PDFAnnotationLine) Style() PDFBorderStyle {
	rv := objc.Send[PDFBorderStyle](p_.ID, objc.Sel("style"))
	return rv
}/* debug [instance_properties/getter]: style */


// Sets the border style.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/pdfkit/pdfborder/style
func (p_ PDFAnnotationLine) SetStyle(value PDFBorderStyle) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setStyle:"), value)
}/* debug [instance_properties/setter]: style */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class PDFAnnotationLine */



