// Code generated from Apple documentation for PDFKit. DO NOT EDIT.

package pdfkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/appkit"
)

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

// An interface definition for the [PDFAnnotationMarkup] class.
type IPDFAnnotationMarkup interface {
	IPDFAnnotation
}

// A object appears as highlighting, underlining, or a strikethrough style applied to the text of a document.
//
// The and properties of the annotation’s associated object determines the stroke thickness and style. The property of the class determines the stroke color.
//
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

// Alloc allocates a new instance without initialization.
func (pc _PDFAnnotationMarkupClass) Alloc() PDFAnnotationMarkup {
	rv := objc.Send[PDFAnnotationMarkup](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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


// Sets the stroke color for the annotation.
//
// [Full Topic]: https://developer.apple.com/documentation/pdfkit/pdfannotation/color
func (p_ PDFAnnotationMarkup) Color() appkit.Color {
	rv := objc.Send[appkit.Color](p_.ID, objc.Sel("color"))
	return rv
}


// SetColor sets the value of the color property.
// Sets the stroke color for the annotation.

//
// [Full Topic]: https://developer.apple.com/documentation/pdfkit/pdfannotation/color
func (p_ PDFAnnotationMarkup) SetColor(value appkit.IColor) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setColor:"), value)
}

// Sets the line width (in points) for the border.
//
// [Full Topic]: https://developer.apple.com/documentation/pdfkit/pdfborder/linewidth
func (p_ PDFAnnotationMarkup) LineWidth() float64 {
	rv := objc.Send[float64](p_.ID, objc.Sel("lineWidth"))
	return rv
}


// SetLineWidth sets the value of the lineWidth property.
// Sets the line width (in points) for the border.

//
// [Full Topic]: https://developer.apple.com/documentation/pdfkit/pdfborder/linewidth
func (p_ PDFAnnotationMarkup) SetLineWidth(value float64) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setLineWidth:"), value)
}

// Sets the border style.
//
// [Full Topic]: https://developer.apple.com/documentation/pdfkit/pdfborder/style
func (p_ PDFAnnotationMarkup) Style() PDFBorderStyle {
	rv := objc.Send[PDFBorderStyle](p_.ID, objc.Sel("style"))
	return rv
}


// SetStyle sets the value of the style property.
// Sets the border style.

//
// [Full Topic]: https://developer.apple.com/documentation/pdfkit/pdfborder/style
func (p_ PDFAnnotationMarkup) SetStyle(value PDFBorderStyle) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setStyle:"), value)
}



