// Code generated from Apple documentation for PDFKit. DO NOT EDIT.

package pdfkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/appkit"
	"github.com/tmc/appledocs/generated/foundation"
)

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

// An interface definition for the [PDFAnnotationInk] class.
type IPDFAnnotationInk interface {
	IPDFAnnotation
	AddBezierPath(path appkit.IBezierPath)
	Paths() foundation.Array
	RemoveBezierPath(path appkit.IBezierPath)
}

//
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

// Alloc allocates a new instance without initialization.
func (pc _PDFAnnotationInkClass) Alloc() PDFAnnotationInk {
	rv := objc.Send[PDFAnnotationInk](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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


// Adds a Bezier path to an annotation.
//
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFAnnotationInk/add(_:)
func (p_ PDFAnnotationInk) AddBezierPath(path appkit.IBezierPath) {
	objc.Send[objc.ID](p_.ID, objc.Sel("addBezierPath:"), path)
}

// Returns an array containing the Bezier paths that make up an annotation.
//
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFAnnotationInk/paths()
func (p_ PDFAnnotationInk) Paths() foundation.Array {
	rv := objc.Send[foundation.Array](p_.ID, objc.Sel("paths"))
	return rv
}

// Removes a Bezier path from an annotation.
//
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFAnnotationInk/remove(_:)
func (p_ PDFAnnotationInk) RemoveBezierPath(path appkit.IBezierPath) {
	objc.Send[objc.ID](p_.ID, objc.Sel("removeBezierPath:"), path)
}

// Sets the stroke color for the annotation.
//
// [Full Topic]: https://developer.apple.com/documentation/pdfkit/pdfannotation/color
func (p_ PDFAnnotationInk) Color() appkit.Color {
	rv := objc.Send[appkit.Color](p_.ID, objc.Sel("color"))
	return rv
}


// SetColor sets the value of the color property.
// Sets the stroke color for the annotation.

//
// [Full Topic]: https://developer.apple.com/documentation/pdfkit/pdfannotation/color
func (p_ PDFAnnotationInk) SetColor(value appkit.IColor) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setColor:"), value)
}

// Sets the line width (in points) for the border.
//
// [Full Topic]: https://developer.apple.com/documentation/pdfkit/pdfborder/linewidth
func (p_ PDFAnnotationInk) LineWidth() float64 {
	rv := objc.Send[float64](p_.ID, objc.Sel("lineWidth"))
	return rv
}


// SetLineWidth sets the value of the lineWidth property.
// Sets the line width (in points) for the border.

//
// [Full Topic]: https://developer.apple.com/documentation/pdfkit/pdfborder/linewidth
func (p_ PDFAnnotationInk) SetLineWidth(value float64) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setLineWidth:"), value)
}

// Sets the border style.
//
// [Full Topic]: https://developer.apple.com/documentation/pdfkit/pdfborder/style
func (p_ PDFAnnotationInk) Style() PDFBorderStyle {
	rv := objc.Send[PDFBorderStyle](p_.ID, objc.Sel("style"))
	return rv
}


// SetStyle sets the value of the style property.
// Sets the border style.

//
// [Full Topic]: https://developer.apple.com/documentation/pdfkit/pdfborder/style
func (p_ PDFAnnotationInk) SetStyle(value PDFBorderStyle) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setStyle:"), value)
}



