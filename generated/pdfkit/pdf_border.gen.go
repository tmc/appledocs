// Code generated from Apple documentation for PDFKit. DO NOT EDIT.

package pdfkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/appkit"
	"github.com/tmc/appledocs/generated/corefoundation"
	"github.com/tmc/appledocs/generated/foundation"
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
	// properties:
	Alignment() TextAlignment /* not a class type */
	SetAlignment(value TextAlignment /* not a class type */)
	Border() IPDFBorder
	SetBorder(value IPDFBorder)
	Bounds() objc.IObject /* cross-framework: Rect */
	SetBounds(value objc.IObject /* cross-framework: Rect */)
	Color() objc.IObject /* cross-framework: Color */
	SetColor(value objc.IObject /* cross-framework: Color */)
	Contents() objc.IObject /* cross-framework: NSString */
	SetContents(value objc.IObject /* cross-framework: NSString */)
	Font() objc.IObject /* cross-framework: Font */
	SetFont(value objc.IObject /* cross-framework: Font */)
	FontColor() objc.IObject /* cross-framework: Color */
	SetFontColor(value objc.IObject /* cross-framework: Color */)
	HasAppearanceStream() bool
	SetHasAppearanceStream(value bool)
	IsHighlighted() bool
	SetIsHighlighted(value bool)
	BorderKeyValues() unsafe.Pointer
	SetBorderKeyValues(value unsafe.Pointer)
	DashPattern() unsafe.Pointer
	SetDashPattern(value unsafe.Pointer)
	LineWidth() float64
	SetLineWidth(value float64)
	Style() unsafe.Pointer
	SetStyle(value unsafe.Pointer)
	// methods:
}

// An optional border for an annotation that lies completely within the annotation rectangle.


// An optional border for an annotation that lies completely within the annotation rectangle.
//
// [Full Topic]
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



// The alignment of the free text and text widget annotation’s text content.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/pdfkit/pdfannotation/alignment
func (p_ PDFBorder) Alignment() TextAlignment /* not a class type */ {
	rv := objc.Send[TextAlignment](p_.ID, objc.Sel("alignment"))
	return rv
}


// The alignment of the free text and text widget annotation’s text content.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/pdfkit/pdfannotation/alignment
func (p_ PDFBorder) SetAlignment(value TextAlignment /* not a class type */) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setAlignment:"), value)
}


// Sets the border style for the annotation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/pdfkit/pdfannotation/border
func (p_ PDFBorder) Border() IPDFBorder {
	rv := objc.Send[PDFBorder](p_.ID, objc.Sel("border"))
	return rv
}


// Sets the border style for the annotation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/pdfkit/pdfannotation/border
func (p_ PDFBorder) SetBorder(value IPDFBorder) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setBorder:"), value)
}


// Returns the bounding box for the annotation in page space.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/pdfkit/pdfannotation/bounds
func (p_ PDFBorder) Bounds() objc.IObject /* cross-framework: Rect */ {
	rv := objc.Send[corefoundation.Rect](p_.ID, objc.Sel("bounds"))
	return rv
}


// Returns the bounding box for the annotation in page space.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/pdfkit/pdfannotation/bounds
func (p_ PDFBorder) SetBounds(value objc.IObject /* cross-framework: Rect */) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setBounds:"), value)
}


// Sets the stroke color for the annotation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/pdfkit/pdfannotation/color
func (p_ PDFBorder) Color() objc.IObject /* cross-framework: Color */ {
	rv := objc.Send[appkit.Color](p_.ID, objc.Sel("color"))
	return rv
}


// Sets the stroke color for the annotation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/pdfkit/pdfannotation/color
func (p_ PDFBorder) SetColor(value objc.IObject /* cross-framework: Color */) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setColor:"), value)
}


// Returns the textual content (if any) associated with the annotation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/pdfkit/pdfannotation/contents
func (p_ PDFBorder) Contents() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](p_.ID, objc.Sel("contents"))
	return rv
}


// Returns the textual content (if any) associated with the annotation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/pdfkit/pdfannotation/contents
func (p_ PDFBorder) SetContents(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setContents:"), value)
}


// The font the annotation uses to display text.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/pdfkit/pdfannotation/font
func (p_ PDFBorder) Font() objc.IObject /* cross-framework: Font */ {
	rv := objc.Send[appkit.Font](p_.ID, objc.Sel("font"))
	return rv
}


// The font the annotation uses to display text.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/pdfkit/pdfannotation/font
func (p_ PDFBorder) SetFont(value objc.IObject /* cross-framework: Font */) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setFont:"), value)
}


// The font color the annotation uses to display text.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/pdfkit/pdfannotation/fontcolor
func (p_ PDFBorder) FontColor() objc.IObject /* cross-framework: Color */ {
	rv := objc.Send[appkit.Color](p_.ID, objc.Sel("fontColor"))
	return rv
}


// The font color the annotation uses to display text.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/pdfkit/pdfannotation/fontcolor
func (p_ PDFBorder) SetFontColor(value objc.IObject /* cross-framework: Color */) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setFontColor:"), value)
}


// Returns a Boolean value that indicates whether the annotation has an appearance stream associated with it.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/pdfkit/pdfannotation/hasappearancestream
func (p_ PDFBorder) HasAppearanceStream() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("hasAppearanceStream"))
	return rv
}


// Returns a Boolean value that indicates whether the annotation has an appearance stream associated with it.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/pdfkit/pdfannotation/hasappearancestream
func (p_ PDFBorder) SetHasAppearanceStream(value bool) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setHasAppearanceStream:"), value)
}


// A Boolean value that indicates whether the annotation is in a highlighted state, such as when the mouse is down on a link annotation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/pdfkit/pdfannotation/ishighlighted
func (p_ PDFBorder) IsHighlighted() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("isHighlighted"))
	return rv
}


// A Boolean value that indicates whether the annotation is in a highlighted state, such as when the mouse is down on a link annotation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/pdfkit/pdfannotation/ishighlighted
func (p_ PDFBorder) SetIsHighlighted(value bool) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setIsHighlighted:"), value)
}


// A dictionary that contains a deep copy of all border properties.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/pdfkit/pdfborder/borderkeyvalues
func (p_ PDFBorder) BorderKeyValues() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("borderKeyValues"))
	return rv
}


// A dictionary that contains a deep copy of all border properties.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/pdfkit/pdfborder/borderkeyvalues
func (p_ PDFBorder) SetBorderKeyValues(value unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setBorderKeyValues:"), value)
}


// Gets the dash pattern for the border as an array of NSNumber objects.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/pdfkit/pdfborder/dashpattern
func (p_ PDFBorder) DashPattern() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("dashPattern"))
	return rv
}


// Gets the dash pattern for the border as an array of NSNumber objects.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/pdfkit/pdfborder/dashpattern
func (p_ PDFBorder) SetDashPattern(value unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setDashPattern:"), value)
}


// Sets the line width (in points) for the border.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/pdfkit/pdfborder/linewidth
func (p_ PDFBorder) LineWidth() float64 {
	rv := objc.Send[float64](p_.ID, objc.Sel("lineWidth"))
	return rv
}


// Sets the line width (in points) for the border.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/pdfkit/pdfborder/linewidth
func (p_ PDFBorder) SetLineWidth(value float64) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setLineWidth:"), value)
}


// Sets the border style.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/pdfkit/pdfborder/style
func (p_ PDFBorder) Style() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("style"))
	return rv
}


// Sets the border style.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/pdfkit/pdfborder/style
func (p_ PDFBorder) SetStyle(value unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setStyle:"), value)
}



