// Code generated from Apple documentation for PDFKit. DO NOT EDIT.

package pdfkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

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

// An interface definition for the [PDFAnnotationCircle] class.
type IPDFAnnotationCircle interface {
	IPDFAnnotation
}

//
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

// Alloc allocates a new instance without initialization.
func (pc _PDFAnnotationCircleClass) Alloc() PDFAnnotationCircle {
	rv := objc.Send[PDFAnnotationCircle](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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


// Sets the stroke color for the annotation.
//
// [Full Topic]: https://developer.apple.com/documentation/pdfkit/pdfannotation/color
func (p_ PDFAnnotationCircle) Color() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("color"))
	return rv
}


// SetColor sets the value of the color property.
// Sets the stroke color for the annotation.

//
// [Full Topic]: https://developer.apple.com/documentation/pdfkit/pdfannotation/color
func (p_ PDFAnnotationCircle) SetColor(value unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setColor:"), value)
}

// Sets the line width (in points) for the border.
//
// [Full Topic]: https://developer.apple.com/documentation/pdfkit/pdfborder/linewidth
func (p_ PDFAnnotationCircle) LineWidth() float64 {
	rv := objc.Send[float64](p_.ID, objc.Sel("lineWidth"))
	return rv
}


// SetLineWidth sets the value of the lineWidth property.
// Sets the line width (in points) for the border.

//
// [Full Topic]: https://developer.apple.com/documentation/pdfkit/pdfborder/linewidth
func (p_ PDFAnnotationCircle) SetLineWidth(value float64) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setLineWidth:"), value)
}

// Sets the border style.
//
// [Full Topic]: https://developer.apple.com/documentation/pdfkit/pdfborder/style
func (p_ PDFAnnotationCircle) Style() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("style"))
	return rv
}


// SetStyle sets the value of the style property.
// Sets the border style.

//
// [Full Topic]: https://developer.apple.com/documentation/pdfkit/pdfborder/style
func (p_ PDFAnnotationCircle) SetStyle(value unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setStyle:"), value)
}



