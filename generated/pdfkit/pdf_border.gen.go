// Code generated from Apple documentation for PDFKit. DO NOT EDIT.

package pdfkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/appkit"
	"github.com/tmc/appledocs/generated/coregraphics"
	"github.com/tmc/appledocs/generated/objectivec"
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
func (p_ PDFBorder) BorderKeyValues() objc.ID {
	rv := objc.Send[objc.ID](p_.ID, objc.Sel("borderKeyValues"))
	return rv
}

// Gets the dash pattern for the border as an array of NSNumber objects.
//
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFBorder/dashPattern
func (p_ PDFBorder) DashPattern() objc.ID {
	rv := objc.Send[objc.ID](p_.ID, objc.Sel("dashPattern"))
	return rv
}


// SetDashPattern sets the value of the dashPattern property.
// Gets the dash pattern for the border as an array of NSNumber objects.

//
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFBorder/dashPattern
func (p_ PDFBorder) SetDashPattern(value objc.ID) {
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
func (p_ PDFBorder) Style() PDFBorderStyle {
	rv := objc.Send[PDFBorderStyle](p_.ID, objc.Sel("style"))
	return rv
}


// SetStyle sets the value of the style property.
// Sets the border style.

//
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFBorder/style
func (p_ PDFBorder) SetStyle(value PDFBorderStyle) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setStyle:"), value)
}

// The alignment of the free text and text widget annotation’s text content.
//
// [Full Topic]: https://developer.apple.com/documentation/pdfkit/pdfannotation/alignment
func (p_ PDFBorder) Alignment() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("alignment"))
	return rv
}


// SetAlignment sets the value of the alignment property.
// The alignment of the free text and text widget annotation’s text content.

//
// [Full Topic]: https://developer.apple.com/documentation/pdfkit/pdfannotation/alignment
func (p_ PDFBorder) SetAlignment(value unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setAlignment:"), value)
}

// Sets the border style for the annotation.
//
// [Full Topic]: https://developer.apple.com/documentation/pdfkit/pdfannotation/border
func (p_ PDFBorder) Border() PDFBorder {
	rv := objc.Send[PDFBorder](p_.ID, objc.Sel("border"))
	return rv
}


// SetBorder sets the value of the border property.
// Sets the border style for the annotation.

//
// [Full Topic]: https://developer.apple.com/documentation/pdfkit/pdfannotation/border
func (p_ PDFBorder) SetBorder(value IPDFBorder) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setBorder:"), value)
}

// Returns the bounding box for the annotation in page space.
//
// [Full Topic]: https://developer.apple.com/documentation/pdfkit/pdfannotation/bounds
func (p_ PDFBorder) Bounds() coregraphics.CGRect {
	rv := objc.Send[coregraphics.CGRect](p_.ID, objc.Sel("bounds"))
	return rv
}


// SetBounds sets the value of the bounds property.
// Returns the bounding box for the annotation in page space.

//
// [Full Topic]: https://developer.apple.com/documentation/pdfkit/pdfannotation/bounds
func (p_ PDFBorder) SetBounds(value coregraphics.CGRect) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setBounds:"), value)
}

// Sets the stroke color for the annotation.
//
// [Full Topic]: https://developer.apple.com/documentation/pdfkit/pdfannotation/color
func (p_ PDFBorder) Color() appkit.Color {
	rv := objc.Send[appkit.Color](p_.ID, objc.Sel("color"))
	return rv
}


// SetColor sets the value of the color property.
// Sets the stroke color for the annotation.

//
// [Full Topic]: https://developer.apple.com/documentation/pdfkit/pdfannotation/color
func (p_ PDFBorder) SetColor(value appkit.IColor) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setColor:"), value)
}

// Returns the textual content (if any) associated with the annotation.
//
// [Full Topic]: https://developer.apple.com/documentation/pdfkit/pdfannotation/contents
func (p_ PDFBorder) Contents() appkit.string {
	rv := objc.Send[appkit.string](p_.ID, objc.Sel("contents"))
	return rv
}


// SetContents sets the value of the contents property.
// Returns the textual content (if any) associated with the annotation.

//
// [Full Topic]: https://developer.apple.com/documentation/pdfkit/pdfannotation/contents
func (p_ PDFBorder) SetContents(value appkit.string) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setContents:"), value)
}

// The font the annotation uses to display text.
//
// [Full Topic]: https://developer.apple.com/documentation/pdfkit/pdfannotation/font
func (p_ PDFBorder) Font() appkit.Font {
	rv := objc.Send[appkit.Font](p_.ID, objc.Sel("font"))
	return rv
}


// SetFont sets the value of the font property.
// The font the annotation uses to display text.

//
// [Full Topic]: https://developer.apple.com/documentation/pdfkit/pdfannotation/font
func (p_ PDFBorder) SetFont(value appkit.IFont) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setFont:"), value)
}

// The font color the annotation uses to display text.
//
// [Full Topic]: https://developer.apple.com/documentation/pdfkit/pdfannotation/fontcolor
func (p_ PDFBorder) FontColor() appkit.Color {
	rv := objc.Send[appkit.Color](p_.ID, objc.Sel("fontColor"))
	return rv
}


// SetFontColor sets the value of the fontColor property.
// The font color the annotation uses to display text.

//
// [Full Topic]: https://developer.apple.com/documentation/pdfkit/pdfannotation/fontcolor
func (p_ PDFBorder) SetFontColor(value appkit.IColor) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setFontColor:"), value)
}

// Returns a Boolean value that indicates whether the annotation has an appearance stream associated with it.
//
// [Full Topic]: https://developer.apple.com/documentation/pdfkit/pdfannotation/hasappearancestream
func (p_ PDFBorder) HasAppearanceStream() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("hasAppearanceStream"))
	return rv
}


// SetHasAppearanceStream sets the value of the hasAppearanceStream property.
// Returns a Boolean value that indicates whether the annotation has an appearance stream associated with it.

//
// [Full Topic]: https://developer.apple.com/documentation/pdfkit/pdfannotation/hasappearancestream
func (p_ PDFBorder) SetHasAppearanceStream(value bool) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setHasAppearanceStream:"), value)
}

// A Boolean value that indicates whether the annotation is in a highlighted state, such as when the mouse is down on a link annotation.
//
// [Full Topic]: https://developer.apple.com/documentation/pdfkit/pdfannotation/ishighlighted
func (p_ PDFBorder) IsHighlighted() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("isHighlighted"))
	return rv
}


// SetIsHighlighted sets the value of the isHighlighted property.
// A Boolean value that indicates whether the annotation is in a highlighted state, such as when the mouse is down on a link annotation.

//
// [Full Topic]: https://developer.apple.com/documentation/pdfkit/pdfannotation/ishighlighted
func (p_ PDFBorder) SetIsHighlighted(value bool) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setIsHighlighted:"), value)
}



