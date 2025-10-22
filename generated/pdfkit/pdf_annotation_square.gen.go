// Code generated from Apple documentation for PDFKit. DO NOT EDIT.

package pdfkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/appkit"
)

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

// An interface definition for the [PDFAnnotationSquare] class.
type IPDFAnnotationSquare interface {
	IPDFAnnotation
	Color() appkit.Color
	SetColor(value appkit.IColor)
	LineWidth() float64
	SetLineWidth(value float64)
	Style() PDFBorderStyle
	SetStyle(value PDFBorderStyle)
}

// A rectangle annotation on a page.
//
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

// Alloc allocates a new instance without initialization.
func (pc _PDFAnnotationSquareClass) Alloc() PDFAnnotationSquare {
	rv := objc.Send[PDFAnnotationSquare](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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


// Sets the stroke color for the annotation.
//
// [Full Topic]: https://developer.apple.com/documentation/pdfkit/pdfannotation/color
func (p_ PDFAnnotationSquare) Color() appkit.Color {
	rv := objc.Send[appkit.Color](p_.ID, objc.Sel("color"))
	return rv
}


// SetColor sets the value of the color property.
// Sets the stroke color for the annotation.

//
// [Full Topic]: https://developer.apple.com/documentation/pdfkit/pdfannotation/color
func (p_ PDFAnnotationSquare) SetColor(value appkit.IColor) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setColor:"), value)
}

// Sets the line width (in points) for the border.
//
// [Full Topic]: https://developer.apple.com/documentation/pdfkit/pdfborder/linewidth
func (p_ PDFAnnotationSquare) LineWidth() float64 {
	rv := objc.Send[float64](p_.ID, objc.Sel("lineWidth"))
	return rv
}


// SetLineWidth sets the value of the lineWidth property.
// Sets the line width (in points) for the border.

//
// [Full Topic]: https://developer.apple.com/documentation/pdfkit/pdfborder/linewidth
func (p_ PDFAnnotationSquare) SetLineWidth(value float64) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setLineWidth:"), value)
}

// Sets the border style.
//
// [Full Topic]: https://developer.apple.com/documentation/pdfkit/pdfborder/style
func (p_ PDFAnnotationSquare) Style() PDFBorderStyle {
	rv := objc.Send[PDFBorderStyle](p_.ID, objc.Sel("style"))
	return rv
}


// SetStyle sets the value of the style property.
// Sets the border style.

//
// [Full Topic]: https://developer.apple.com/documentation/pdfkit/pdfborder/style
func (p_ PDFAnnotationSquare) SetStyle(value PDFBorderStyle) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setStyle:"), value)
}



