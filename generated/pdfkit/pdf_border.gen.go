// Code generated from Apple documentation for PDFKit. DO NOT EDIT.

package pdfkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [PDFBorder] class.
var (
	PDFBorderClass     _PDFBorderClass
	PDFBorderClassOnce sync.Once
)

func getPDFBorderClass() _PDFBorderClass {
	PDFBorderClassOnce.Do(func() {
		PDFBorderClass = _PDFBorderClass{objc.GetClass("PDFBorder")}
	})
	return PDFBorderClass
}

type _PDFBorderClass struct {
	class objc.Class
}

// An interface definition for the [PDFBorder] class.
type IPDFBorder interface {
	objectivec.IObject
}

// An optional border for an annotation that lies completely within the annotation rectangle.
//
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFBorder
type PDFBorder struct {
	objectivec.Object
}

// PDFBorderFrom constructs a [PDFBorder] from an unsafe.Pointer.
//
// An optional border for an annotation that lies completely within the annotation rectangle.
func PDFBorderFrom(ptr unsafe.Pointer) PDFBorder {
	return PDFBorder{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (pc _PDFBorderClass) Alloc() PDFBorder {
	rv := objc.Send[PDFBorder](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (pc _PDFBorderClass) New() PDFBorder {
	rv := objc.Send[PDFBorder](objc.ID(pc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (p_ PDFBorder) Init() PDFBorder {
	rv := objc.Send[PDFBorder](p_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p_ PDFBorder) Autorelease() PDFBorder {
	rv := objc.Send[PDFBorder](p_.ID, objc.Sel("autorelease"))
	return rv
}

// NewPDFBorder creates a new PDFBorder instance.
func NewPDFBorder() PDFBorder {
	return getPDFBorderClass().New()
}


// A dictionary that contains a deep copy of all border properties.
//
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFBorder/borderKeyValues
func (p_ PDFBorder) BorderKeyValues() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("borderKeyValues"))
	return rv
}

// Gets the dash pattern for the border as an array of NSNumber objects.
//
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFBorder/dashPattern
func (p_ PDFBorder) DashPattern() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("dashPattern"))
	return rv
}


// SetDashPattern sets the value of the dashPattern property.
// Gets the dash pattern for the border as an array of NSNumber objects.

//
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFBorder/dashPattern
func (p_ PDFBorder) SetDashPattern(value unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setDashPattern:"), value)
}

// Sets the line width (in points) for the border.
//
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFBorder/lineWidth
func (p_ PDFBorder) LineWidth() float64 {
	rv := objc.Send[float64](p_.ID, objc.Sel("lineWidth"))
	return rv
}


// SetLineWidth sets the value of the lineWidth property.
// Sets the line width (in points) for the border.

//
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFBorder/lineWidth
func (p_ PDFBorder) SetLineWidth(value float64) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setLineWidth:"), value)
}

// Sets the border style.
//
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFBorder/style
func (p_ PDFBorder) Style() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("style"))
	return rv
}


// SetStyle sets the value of the style property.
// Sets the border style.

//
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFBorder/style
func (p_ PDFBorder) SetStyle(value unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setStyle:"), value)
}



