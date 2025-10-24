// Code generated from Apple documentation for PDFKit. DO NOT EDIT.

package pdfkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/appkit"
	"github.com/tmc/appledocs/generated/foundation"
)

/* debug [class.gen.go]: Generating class PDFAnnotationMarkup */


/* debug [class_header]: Header for PDFAnnotationMarkup */
// The class instance for the [PDFAnnotationMarkup] class.
var (
	PDFAnnotationMarkupClass     _PDFAnnotationMarkupClass
	PDFAnnotationMarkupClassOnce sync.Once
)

func getPDFAnnotationMarkupClass() _PDFAnnotationMarkupClass {
	PDFAnnotationMarkupClassOnce.Do(func() {
		PDFAnnotationMarkupClass = _PDFAnnotationMarkupClass{objc.GetClass("PDFAnnotationMarkup")}
	})
	return PDFAnnotationMarkupClass
}

type _PDFAnnotationMarkupClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for PDFAnnotationMarkup */
// An interface definition for the [PDFAnnotationMarkup] class.
type IPDFAnnotationMarkup interface {
	IPDFAnnotation
	
/* debug [class_interface_properties]: Properties for PDFAnnotationMarkup */
	// properties:
	Color() appkit.Color
	SetColor(value appkit.Color)
	LineWidth() float64
	SetLineWidth(value float64)
	Style() PDFBorderStyle
	SetStyle(value PDFBorderStyle)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for PDFAnnotationMarkup */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for PDFAnnotationMarkup */
// Alloc allocates a new instance without initialization.
func (pc _PDFAnnotationMarkupClass) Alloc() PDFAnnotationMarkup {
	rv := objc.Send[PDFAnnotationMarkup](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (pc _PDFAnnotationMarkupClass) New() PDFAnnotationMarkup {
	rv := objc.Send[PDFAnnotationMarkup](objc.ID(pc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (p_ PDFAnnotationMarkup) Init() PDFAnnotationMarkup {
	rv := objc.Send[PDFAnnotationMarkup](p_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p_ PDFAnnotationMarkup) Autorelease() PDFAnnotationMarkup {
	rv := objc.Send[PDFAnnotationMarkup](p_.ID, objc.Sel("autorelease"))
	return rv
}

// NewPDFAnnotationMarkup creates a new PDFAnnotationMarkup instance.
func NewPDFAnnotationMarkup() PDFAnnotationMarkup {
	return getPDFAnnotationMarkupClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for PDFAnnotationMarkup */
// A object appears as highlighting, underlining, or a strikethrough style applied to the text of a document.
//
// The and properties of the annotation’s associated object determines the stroke thickness and style. The property of the class determines the stroke color.


// A object appears as highlighting, underlining, or a strikethrough style applied to the text of a document.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFAnnotationMarkup
type PDFAnnotationMarkup struct {
	PDFAnnotation
}

// PDFAnnotationMarkupFrom constructs a [PDFAnnotationMarkup] from an unsafe.Pointer.
//
// A object appears as highlighting, underlining, or a strikethrough style applied to the text of a document.
func PDFAnnotationMarkupFrom(ptr unsafe.Pointer) PDFAnnotationMarkup {
	return PDFAnnotationMarkup{
		PDFAnnotation: PDFAnnotationFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for PDFAnnotationMarkup *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for PDFAnnotationMarkup */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for PDFAnnotationMarkup */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for PDFAnnotationMarkup */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for PDFAnnotationMarkup */

// Sets the stroke color for the annotation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/pdfkit/pdfannotation/color
func (p_ PDFAnnotationMarkup) Color() appkit.Color {
	rv := objc.Send[appkit.Color](p_.ID, objc.Sel("color"))
	return rv
}/* debug [instance_properties/getter]: color */


// Sets the stroke color for the annotation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/pdfkit/pdfannotation/color
func (p_ PDFAnnotationMarkup) SetColor(value appkit.Color) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setColor:"), value)
}/* debug [instance_properties/setter]: color */


// Sets the line width (in points) for the border.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/pdfkit/pdfborder/linewidth
func (p_ PDFAnnotationMarkup) LineWidth() float64 {
	rv := objc.Send[float64](p_.ID, objc.Sel("lineWidth"))
	return rv
}/* debug [instance_properties/getter]: lineWidth */


// Sets the line width (in points) for the border.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/pdfkit/pdfborder/linewidth
func (p_ PDFAnnotationMarkup) SetLineWidth(value float64) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setLineWidth:"), value)
}/* debug [instance_properties/setter]: lineWidth */


// Sets the border style.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/pdfkit/pdfborder/style
func (p_ PDFAnnotationMarkup) Style() PDFBorderStyle {
	rv := objc.Send[PDFBorderStyle](p_.ID, objc.Sel("style"))
	return rv
}/* debug [instance_properties/getter]: style */


// Sets the border style.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/pdfkit/pdfborder/style
func (p_ PDFAnnotationMarkup) SetStyle(value PDFBorderStyle) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setStyle:"), value)
}/* debug [instance_properties/setter]: style */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class PDFAnnotationMarkup */



