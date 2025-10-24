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

/* debug [class.gen.go]: Generating class PDFBorder */


/* debug [class_header]: Header for PDFBorder */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for PDFBorder */
// An interface definition for the [PDFBorder] class.
type IPDFBorder interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for PDFBorder */
	// properties:
	BorderKeyValues() objc.IObject /* cross-framework: NSDictionary */
	DashPattern() objc.IObject /* cross-framework: NSArray */
	SetDashPattern(value objc.IObject /* cross-framework: NSArray */)
	LineWidth() float64
	SetLineWidth(value float64)
	Style() PDFBorderStyle
	SetStyle(value PDFBorderStyle)
	Alignment() TextAlignment /* not a class type */
	SetAlignment(value TextAlignment /* not a class type */)
	Border() IPDFBorder
	SetBorder(value IPDFBorder)
	Bounds() corefoundation.CGRect
	SetBounds(value corefoundation.CGRect)
	Color() appkit.Color
	SetColor(value appkit.Color)
	Contents() objc.IObject /* cross-framework: NSString */
	SetContents(value objc.IObject /* cross-framework: NSString */)
	Font() appkit.Font
	SetFont(value appkit.Font)
	FontColor() appkit.Color
	SetFontColor(value appkit.Color)
	HasAppearanceStream() bool
	SetHasAppearanceStream(value bool)
	IsHighlighted() bool
	SetIsHighlighted(value bool)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for PDFBorder */
	// methods:
	DrawInRect(rect Rect /* not a class type */)
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for PDFBorder */
// Alloc allocates a new instance without initialization.
func (pc _PDFBorderClass) Alloc() PDFBorder {
	rv := objc.Send[PDFBorder](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for PDFBorder */
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for PDFBorder *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for PDFBorder */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for PDFBorder */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for PDFBorder */

// Draws the border.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFBorder/draw(in:)
func (p_ PDFBorder) DrawInRect(rect Rect /* not a class type */) {
	objc.Send[objc.ID](p_.ID, objc.Sel("drawInRect:"), rect)
}/* debug [instance_methods/method]: DrawInRect */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for PDFBorder */

// A dictionary that contains a deep copy of all border properties.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFBorder/borderKeyValues
func (p_ PDFBorder) BorderKeyValues() objc.IObject /* cross-framework: NSDictionary */ {
	rv := objc.Send[foundation.NSDictionary](p_.ID, objc.Sel("borderKeyValues"))
	return rv
}/* debug [instance_properties/getter]: borderKeyValues */


// Gets the dash pattern for the border as an array of NSNumber objects.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFBorder/dashPattern
func (p_ PDFBorder) DashPattern() objc.IObject /* cross-framework: NSArray */ {
	rv := objc.Send[foundation.NSArray](p_.ID, objc.Sel("dashPattern"))
	return rv
}/* debug [instance_properties/getter]: dashPattern */


// Gets the dash pattern for the border as an array of NSNumber objects.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFBorder/dashPattern
func (p_ PDFBorder) SetDashPattern(value objc.IObject /* cross-framework: NSArray */) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setDashPattern:"), value)
}/* debug [instance_properties/setter]: dashPattern */


// Sets the line width (in points) for the border.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFBorder/lineWidth
func (p_ PDFBorder) LineWidth() float64 {
	rv := objc.Send[float64](p_.ID, objc.Sel("lineWidth"))
	return rv
}/* debug [instance_properties/getter]: lineWidth */


// Sets the line width (in points) for the border.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFBorder/lineWidth
func (p_ PDFBorder) SetLineWidth(value float64) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setLineWidth:"), value)
}/* debug [instance_properties/setter]: lineWidth */


// Sets the border style.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFBorder/style
func (p_ PDFBorder) Style() PDFBorderStyle {
	rv := objc.Send[PDFBorderStyle](p_.ID, objc.Sel("style"))
	return rv
}/* debug [instance_properties/getter]: style */


// Sets the border style.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFBorder/style
func (p_ PDFBorder) SetStyle(value PDFBorderStyle) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setStyle:"), value)
}/* debug [instance_properties/setter]: style */


// The alignment of the free text and text widget annotation’s text content.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/pdfkit/pdfannotation/alignment
func (p_ PDFBorder) Alignment() TextAlignment /* not a class type */ {
	rv := objc.Send[TextAlignment](p_.ID, objc.Sel("alignment"))
	return rv
}/* debug [instance_properties/getter]: alignment */


// The alignment of the free text and text widget annotation’s text content.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/pdfkit/pdfannotation/alignment
func (p_ PDFBorder) SetAlignment(value TextAlignment /* not a class type */) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setAlignment:"), value)
}/* debug [instance_properties/setter]: alignment */


// Sets the border style for the annotation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/pdfkit/pdfannotation/border
func (p_ PDFBorder) Border() IPDFBorder {
	rv := objc.Send[PDFBorder](p_.ID, objc.Sel("border"))
	return rv
}/* debug [instance_properties/getter]: border */


// Sets the border style for the annotation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/pdfkit/pdfannotation/border
func (p_ PDFBorder) SetBorder(value IPDFBorder) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setBorder:"), value)
}/* debug [instance_properties/setter]: border */


// Returns the bounding box for the annotation in page space.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/pdfkit/pdfannotation/bounds
func (p_ PDFBorder) Bounds() corefoundation.CGRect {
	rv := objc.Send[corefoundation.CGRect](p_.ID, objc.Sel("bounds"))
	return rv
}/* debug [instance_properties/getter]: bounds */


// Returns the bounding box for the annotation in page space.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/pdfkit/pdfannotation/bounds
func (p_ PDFBorder) SetBounds(value corefoundation.CGRect) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setBounds:"), value)
}/* debug [instance_properties/setter]: bounds */


// Sets the stroke color for the annotation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/pdfkit/pdfannotation/color
func (p_ PDFBorder) Color() appkit.Color {
	rv := objc.Send[appkit.Color](p_.ID, objc.Sel("color"))
	return rv
}/* debug [instance_properties/getter]: color */


// Sets the stroke color for the annotation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/pdfkit/pdfannotation/color
func (p_ PDFBorder) SetColor(value appkit.Color) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setColor:"), value)
}/* debug [instance_properties/setter]: color */


// Returns the textual content (if any) associated with the annotation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/pdfkit/pdfannotation/contents
func (p_ PDFBorder) Contents() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](p_.ID, objc.Sel("contents"))
	return rv
}/* debug [instance_properties/getter]: contents */


// Returns the textual content (if any) associated with the annotation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/pdfkit/pdfannotation/contents
func (p_ PDFBorder) SetContents(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setContents:"), value)
}/* debug [instance_properties/setter]: contents */


// The font the annotation uses to display text.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/pdfkit/pdfannotation/font
func (p_ PDFBorder) Font() appkit.Font {
	rv := objc.Send[appkit.Font](p_.ID, objc.Sel("font"))
	return rv
}/* debug [instance_properties/getter]: font */


// The font the annotation uses to display text.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/pdfkit/pdfannotation/font
func (p_ PDFBorder) SetFont(value appkit.Font) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setFont:"), value)
}/* debug [instance_properties/setter]: font */


// The font color the annotation uses to display text.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/pdfkit/pdfannotation/fontcolor
func (p_ PDFBorder) FontColor() appkit.Color {
	rv := objc.Send[appkit.Color](p_.ID, objc.Sel("fontColor"))
	return rv
}/* debug [instance_properties/getter]: fontColor */


// The font color the annotation uses to display text.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/pdfkit/pdfannotation/fontcolor
func (p_ PDFBorder) SetFontColor(value appkit.Color) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setFontColor:"), value)
}/* debug [instance_properties/setter]: fontColor */


// Returns a Boolean value that indicates whether the annotation has an appearance stream associated with it.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/pdfkit/pdfannotation/hasappearancestream
func (p_ PDFBorder) HasAppearanceStream() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("hasAppearanceStream"))
	return rv
}/* debug [instance_properties/getter]: hasAppearanceStream */


// Returns a Boolean value that indicates whether the annotation has an appearance stream associated with it.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/pdfkit/pdfannotation/hasappearancestream
func (p_ PDFBorder) SetHasAppearanceStream(value bool) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setHasAppearanceStream:"), value)
}/* debug [instance_properties/setter]: hasAppearanceStream */


// A Boolean value that indicates whether the annotation is in a highlighted state, such as when the mouse is down on a link annotation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/pdfkit/pdfannotation/ishighlighted
func (p_ PDFBorder) IsHighlighted() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("isHighlighted"))
	return rv
}/* debug [instance_properties/getter]: isHighlighted */


// A Boolean value that indicates whether the annotation is in a highlighted state, such as when the mouse is down on a link annotation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/pdfkit/pdfannotation/ishighlighted
func (p_ PDFBorder) SetIsHighlighted(value bool) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setIsHighlighted:"), value)
}/* debug [instance_properties/setter]: isHighlighted */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class PDFBorder */



