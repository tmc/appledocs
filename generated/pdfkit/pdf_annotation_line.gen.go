// Code generated from Apple documentation for PDFKit. DO NOT EDIT.

package pdfkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/appkit"
)

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

// An interface definition for the [PDFAnnotationLine] class.
type IPDFAnnotationLine interface {
	IPDFAnnotation
	Color() appkit.Color
	SetColor(value appkit.IColor)
	LineWidth() float64
	SetLineWidth(value float64)
	Style() PDFBorderStyle
	SetStyle(value PDFBorderStyle)
}

// A object displays a single line on a page.
//
// The and properties of the annotation’s associated object determines the stroke thickness and style. The property of the class determines the stroke color.
//
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

// Alloc allocates a new instance without initialization.
func (pc _PDFAnnotationLineClass) Alloc() PDFAnnotationLine {
	rv := objc.Send[PDFAnnotationLine](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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


// Sets the stroke color for the annotation.
//
// [Full Topic]: https://developer.apple.com/documentation/pdfkit/pdfannotation/color
func (p_ PDFAnnotationLine) Color() appkit.Color {
	rv := objc.Send[appkit.Color](p_.ID, objc.Sel("color"))
	return rv
}


// SetColor sets the value of the color property.
// Sets the stroke color for the annotation.

//
// [Full Topic]: https://developer.apple.com/documentation/pdfkit/pdfannotation/color
func (p_ PDFAnnotationLine) SetColor(value appkit.IColor) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setColor:"), value)
}

// Sets the line width (in points) for the border.
//
// [Full Topic]: https://developer.apple.com/documentation/pdfkit/pdfborder/linewidth
func (p_ PDFAnnotationLine) LineWidth() float64 {
	rv := objc.Send[float64](p_.ID, objc.Sel("lineWidth"))
	return rv
}


// SetLineWidth sets the value of the lineWidth property.
// Sets the line width (in points) for the border.

//
// [Full Topic]: https://developer.apple.com/documentation/pdfkit/pdfborder/linewidth
func (p_ PDFAnnotationLine) SetLineWidth(value float64) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setLineWidth:"), value)
}

// Sets the border style.
//
// [Full Topic]: https://developer.apple.com/documentation/pdfkit/pdfborder/style
func (p_ PDFAnnotationLine) Style() PDFBorderStyle {
	rv := objc.Send[PDFBorderStyle](p_.ID, objc.Sel("style"))
	return rv
}


// SetStyle sets the value of the style property.
// Sets the border style.

//
// [Full Topic]: https://developer.apple.com/documentation/pdfkit/pdfborder/style
func (p_ PDFAnnotationLine) SetStyle(value PDFBorderStyle) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setStyle:"), value)
}



