// Code generated from Apple documentation for PDFKit. DO NOT EDIT.

package pdfkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/coregraphics"
)

// The class instance for the [PDFAnnotation] class.
var (
	PDFAnnotationClass     _PDFAnnotationClass
	PDFAnnotationClassOnce sync.Once
)

func getPDFAnnotationClass() _PDFAnnotationClass {
	PDFAnnotationClassOnce.Do(func() {
		PDFAnnotationClass = _PDFAnnotationClass{objc.GetClass("PDFAnnotation")}
	})
	return PDFAnnotationClass
}

type _PDFAnnotationClass struct {
	class objc.Class
}

// An interface definition for the [PDFAnnotation] class.
type IPDFAnnotation interface {
	objectivec.IObject
	AddBezierPath(path unsafe.Pointer)
	DrawWithBox(box unsafe.Pointer)
	DrawWithBoxInContext(box unsafe.Pointer, context CGContextRef)
	RemoveBezierPath(path unsafe.Pointer)
	RemoveValueForAnnotationKey(key unsafe.Pointer)
	SetBooleanForAnnotationKey(value bool, key unsafe.Pointer) bool
	SetRectForAnnotationKey(value Rect, key unsafe.Pointer) bool
	SetValueForAnnotationKey(value objc.ID, key unsafe.Pointer) bool
	ValueForAnnotationKey(key unsafe.Pointer) objc.ID
}

// An annotation in a PDF document.
//
// In addition to its primary textual content, a PDF file can contain annotations that represent links, form elements, highlighting circles, textual notes, and so on. Each annotation has a specific location on a page and may offer interactivity with the user.
//
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFAnnotation
type PDFAnnotation struct {
	objectivec.Object
}

// PDFAnnotationFrom constructs a [PDFAnnotation] from an unsafe.Pointer.
//
// An annotation in a PDF document.
func PDFAnnotationFrom(ptr unsafe.Pointer) PDFAnnotation {
	return PDFAnnotation{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (pc _PDFAnnotationClass) Alloc() PDFAnnotation {
	rv := objc.Send[PDFAnnotation](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (pc _PDFAnnotationClass) New() PDFAnnotation {
	rv := objc.Send[PDFAnnotation](objc.ID(pc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (p_ PDFAnnotation) Init() PDFAnnotation {
	rv := objc.Send[PDFAnnotation](p_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p_ PDFAnnotation) Autorelease() PDFAnnotation {
	rv := objc.Send[PDFAnnotation](p_.ID, objc.Sel("autorelease"))
	return rv
}

// NewPDFAnnotation creates a new PDFAnnotation instance.
func NewPDFAnnotation() PDFAnnotation {
	return getPDFAnnotationClass().New()
}




// Creates a PDF annotation with the specified bounds, type, and optional properties.
//
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFAnnotation/init(bounds:forType:withProperties:)
func NewPDFAnnotationWithBoundsForTypeWithProperties(bounds Rect, annotationType unsafe.Pointer, properties objc.ID) PDFAnnotation {
	instance := getPDFAnnotationClass().Alloc()
	rv := objc.Send[PDFAnnotation](instance.ID, objc.Sel("initWithBounds:forType:withProperties:"), bounds, annotationType, properties)
	rv.Autorelease()
	return rv
}


// Returns a line style that corresponds to the specified name.
//
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFAnnotation/lineStyle(fromName:)
func (pc _PDFAnnotationClass) LineStyleFromName(name string) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(pc.class), objc.Sel("lineStyleFromName:"), objc.String(name))
	return rv
}

// Returns the name of the line style, which matches the definition in the Adobe PDF Specification.
//
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFAnnotation/name(for:)
func (pc _PDFAnnotationClass) NameForLineStyle(style unsafe.Pointer) string {
	rv := objc.Send[string](objc.ID(pc.class), objc.Sel("nameForLineStyle:"), style)
	return rv
}

// Adds a bezier path to the ink annotation.
//
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFAnnotation/add(_:)
func (p_ PDFAnnotation) AddBezierPath(path unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("addBezierPath:"), path)
}

// Draws the annotation on its associated page.
//
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFAnnotation/draw(with:)
func (p_ PDFAnnotation) DrawWithBox(box unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("drawWithBox:"), box)
}

// Draws the annotation in a graphics context using page-space coordinates relative to the origin of the specified box.
//
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFAnnotation/draw(with:in:)
func (p_ PDFAnnotation) DrawWithBoxInContext(box unsafe.Pointer, context CGContextRef) {
	objc.Send[objc.ID](p_.ID, objc.Sel("drawWithBox:inContext:"), box, context)
}

// Removes a bezier path from an ink annotation.
//
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFAnnotation/remove(_:)
func (p_ PDFAnnotation) RemoveBezierPath(path unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("removeBezierPath:"), path)
}

// Removes a value from the annotation’s dictionary.
//
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFAnnotation/removeValue(forAnnotationKey:)
func (p_ PDFAnnotation) RemoveValueForAnnotationKey(key unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("removeValueForAnnotationKey:"), key)
}

// Sets a Boolean value in the annotation’s dictionary.
//
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFAnnotation/setBoolean(_:forAnnotationKey:)
func (p_ PDFAnnotation) SetBooleanForAnnotationKey(value bool, key unsafe.Pointer) bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("setBoolean:forAnnotationKey:"), value, key)
	return rv
}

// Sets a rectangle value in the annotation’s dictionary.
//
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFAnnotation/setRect(_:forAnnotationKey:)
func (p_ PDFAnnotation) SetRectForAnnotationKey(value Rect, key unsafe.Pointer) bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("setRect:forAnnotationKey:"), value, key)
	return rv
}

// Sets a value in the annotation’s dictionary.
//
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFAnnotation/setValue(_:forAnnotationKey:)
func (p_ PDFAnnotation) SetValueForAnnotationKey(value objc.ID, key unsafe.Pointer) bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("setValue:forAnnotationKey:"), value, key)
	return rv
}

// Returns a deep copy of the key-value pairs of properties for the specified key.
//
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFAnnotation/value(forAnnotationKey:)
func (p_ PDFAnnotation) ValueForAnnotationKey(key unsafe.Pointer) objc.ID {
	rv := objc.Send[objc.ID](p_.ID, objc.Sel("valueForAnnotationKey:"), key)
	return rv
}

// An object that represents an action for a PDF element, such as a link annotation.
//
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFAnnotation/action
func (p_ PDFAnnotation) Action() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("action"))
	return rv
}


// SetAction sets the value of the action property.
// An object that represents an action for a PDF element, such as a link annotation.

//
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFAnnotation/action
func (p_ PDFAnnotation) SetAction(value unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setAction:"), value)
}

// The alignment of the free text and text widget annotation’s text content.
//
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFAnnotation/alignment
func (p_ PDFAnnotation) Alignment() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("alignment"))
	return rv
}


// SetAlignment sets the value of the alignment property.
// The alignment of the free text and text widget annotation’s text content.

//
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFAnnotation/alignment
func (p_ PDFAnnotation) SetAlignment(value unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setAlignment:"), value)
}

// A Boolean value that indicates whether clicking or tapping a selected radio button toggles it to an unselected state.
//
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFAnnotation/allowsToggleToOff
func (p_ PDFAnnotation) AllowsToggleToOff() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("allowsToggleToOff"))
	return rv
}


// SetAllowsToggleToOff sets the value of the allowsToggleToOff property.
// A Boolean value that indicates whether clicking or tapping a selected radio button toggles it to an unselected state.

//
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFAnnotation/allowsToggleToOff
func (p_ PDFAnnotation) SetAllowsToggleToOff(value bool) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setAllowsToggleToOff:"), value)
}

// A dictionary that contains a deep copy of the widget’s properties.
//
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFAnnotation/annotationKeyValues
func (p_ PDFAnnotation) AnnotationKeyValues() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("annotationKeyValues"))
	return rv
}

// The color of the widget’s background.
//
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFAnnotation/backgroundColor
func (p_ PDFAnnotation) BackgroundColor() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("backgroundColor"))
	return rv
}


// SetBackgroundColor sets the value of the backgroundColor property.
// The color of the widget’s background.

//
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFAnnotation/backgroundColor
func (p_ PDFAnnotation) SetBackgroundColor(value unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setBackgroundColor:"), value)
}

// Sets the border style for the annotation.
//
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFAnnotation/border
func (p_ PDFAnnotation) Border() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("border"))
	return rv
}


// SetBorder sets the value of the border property.
// Sets the border style for the annotation.

//
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFAnnotation/border
func (p_ PDFAnnotation) SetBorder(value unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setBorder:"), value)
}

// Returns the bounding box for the annotation in page space.
//
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFAnnotation/bounds
func (p_ PDFAnnotation) Bounds() coregraphics.CGRect {
	rv := objc.Send[coregraphics.CGRect](p_.ID, objc.Sel("bounds"))
	return rv
}


// SetBounds sets the value of the bounds property.
// Returns the bounding box for the annotation in page space.

//
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFAnnotation/bounds
func (p_ PDFAnnotation) SetBounds(value coregraphics.CGRect) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setBounds:"), value)
}

// The current state of the button widget annotation.
//
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFAnnotation/buttonWidgetState
func (p_ PDFAnnotation) ButtonWidgetState() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("buttonWidgetState"))
	return rv
}


// SetButtonWidgetState sets the value of the buttonWidgetState property.
// The current state of the button widget annotation.

//
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFAnnotation/buttonWidgetState
func (p_ PDFAnnotation) SetButtonWidgetState(value unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setButtonWidgetState:"), value)
}

// A string value that differentiates button widgets in the same group, such as to identify mutually exclusive radio buttons from each other.
//
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFAnnotation/buttonWidgetStateString
func (p_ PDFAnnotation) ButtonWidgetStateString() string {
	rv := objc.Send[string](p_.ID, objc.Sel("buttonWidgetStateString"))
	return rv
}


// SetButtonWidgetStateString sets the value of the buttonWidgetStateString property.
// A string value that differentiates button widgets in the same group, such as to identify mutually exclusive radio buttons from each other.

//
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFAnnotation/buttonWidgetStateString
func (p_ PDFAnnotation) SetButtonWidgetStateString(value string) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setButtonWidgetStateString:"), objc.String(value))
}

// The title of push button widget annotations.
//
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFAnnotation/caption
func (p_ PDFAnnotation) Caption() string {
	rv := objc.Send[string](p_.ID, objc.Sel("caption"))
	return rv
}


// SetCaption sets the value of the caption property.
// The title of push button widget annotations.

//
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFAnnotation/caption
func (p_ PDFAnnotation) SetCaption(value string) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setCaption:"), objc.String(value))
}

// An array of strings that specifies the options in either a list or a pop-up menu.
//
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFAnnotation/choices
func (p_ PDFAnnotation) Choices() []string {
	rv := objc.Send[[]string](p_.ID, objc.Sel("choices"))
	return rv
}


// SetChoices sets the value of the choices property.
// An array of strings that specifies the options in either a list or a pop-up menu.

//
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFAnnotation/choices
func (p_ PDFAnnotation) SetChoices(value []string) {
	// Convert Go slice to NSArray
	var nsArray objc.ID
	if len(value) > 0 {
		nsArray = objc.ID(objc.GetClass("NSMutableArray")).Send(objc.Sel("arrayWithCapacity:"), len(value))
		for _, item := range value {
			nsArray.Send(objc.Sel("addObject:"), item)
		}
	} else {
		nsArray = objc.ID(objc.GetClass("NSArray")).Send(objc.Sel("array"))
	}
	objc.Send[objc.ID](p_.ID, objc.Sel("setChoices:"), nsArray)
}

// Sets the stroke color for the annotation.
//
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFAnnotation/color
func (p_ PDFAnnotation) Color() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("color"))
	return rv
}


// SetColor sets the value of the color property.
// Sets the stroke color for the annotation.

//
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFAnnotation/color
func (p_ PDFAnnotation) SetColor(value unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setColor:"), value)
}

// Returns the textual content (if any) associated with the annotation.
//
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFAnnotation/contents
func (p_ PDFAnnotation) Contents() string {
	rv := objc.Send[string](p_.ID, objc.Sel("contents"))
	return rv
}


// SetContents sets the value of the contents property.
// Returns the textual content (if any) associated with the annotation.

//
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFAnnotation/contents
func (p_ PDFAnnotation) SetContents(value string) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setContents:"), objc.String(value))
}

// The destination for a link annotation.
//
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFAnnotation/destination
func (p_ PDFAnnotation) Destination() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("destination"))
	return rv
}


// SetDestination sets the value of the destination property.
// The destination for a link annotation.

//
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFAnnotation/destination
func (p_ PDFAnnotation) SetDestination(value unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setDestination:"), value)
}

// The style of the line annotation’s ending point, such as square or filled arrowhead.
//
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFAnnotation/endLineStyle
func (p_ PDFAnnotation) EndLineStyle() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("endLineStyle"))
	return rv
}


// SetEndLineStyle sets the value of the endLineStyle property.
// The style of the line annotation’s ending point, such as square or filled arrowhead.

//
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFAnnotation/endLineStyle
func (p_ PDFAnnotation) SetEndLineStyle(value unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setEndLineStyle:"), value)
}

// The point where a line ends, in annotation-space coordinates.
//
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFAnnotation/endPoint
func (p_ PDFAnnotation) EndPoint() Point {
	rv := objc.Send[Point](p_.ID, objc.Sel("endPoint"))
	return rv
}


// SetEndPoint sets the value of the endPoint property.
// The point where a line ends, in annotation-space coordinates.

//
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFAnnotation/endPoint
func (p_ PDFAnnotation) SetEndPoint(value Point) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setEndPoint:"), value)
}

// The widget identifier for form annotation actions and behaviors.
//
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFAnnotation/fieldName
func (p_ PDFAnnotation) FieldName() string {
	rv := objc.Send[string](p_.ID, objc.Sel("fieldName"))
	return rv
}


// SetFieldName sets the value of the fieldName property.
// The widget identifier for form annotation actions and behaviors.

//
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFAnnotation/fieldName
func (p_ PDFAnnotation) SetFieldName(value string) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setFieldName:"), objc.String(value))
}

// The font the annotation uses to display text.
//
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFAnnotation/font
func (p_ PDFAnnotation) Font() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("font"))
	return rv
}


// SetFont sets the value of the font property.
// The font the annotation uses to display text.

//
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFAnnotation/font
func (p_ PDFAnnotation) SetFont(value unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setFont:"), value)
}

// The font color the annotation uses to display text.
//
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFAnnotation/fontColor
func (p_ PDFAnnotation) FontColor() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("fontColor"))
	return rv
}


// SetFontColor sets the value of the fontColor property.
// The font color the annotation uses to display text.

//
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFAnnotation/fontColor
func (p_ PDFAnnotation) SetFontColor(value unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setFontColor:"), value)
}

// Returns a Boolean value that indicates whether the annotation has an appearance stream associated with it.
//
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFAnnotation/hasAppearanceStream
func (p_ PDFAnnotation) HasAppearanceStream() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("hasAppearanceStream"))
	return rv
}

// A Boolean value that indicates whether the annotation divides the text widget’s bounds into equally spaced segments, such as in a form entry field.
//
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFAnnotation/hasComb
func (p_ PDFAnnotation) Comb() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("comb"))
	return rv
}


// SetComb sets the value of the comb property.
// A Boolean value that indicates whether the annotation divides the text widget’s bounds into equally spaced segments, such as in a form entry field.

//
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFAnnotation/hasComb
func (p_ PDFAnnotation) SetComb(value bool) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setComb:"), value)
}

// The type of icon to display for a pop-up text annotation.
//
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFAnnotation/iconType
func (p_ PDFAnnotation) IconType() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("iconType"))
	return rv
}


// SetIconType sets the value of the iconType property.
// The type of icon to display for a pop-up text annotation.

//
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFAnnotation/iconType
func (p_ PDFAnnotation) SetIconType(value unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setIconType:"), value)
}

// The fill color for drawing a circle, line, or square annotation.
//
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFAnnotation/interiorColor
func (p_ PDFAnnotation) InteriorColor() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("interiorColor"))
	return rv
}


// SetInteriorColor sets the value of the interiorColor property.
// The fill color for drawing a circle, line, or square annotation.

//
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFAnnotation/interiorColor
func (p_ PDFAnnotation) SetInteriorColor(value unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setInteriorColor:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFAnnotation/isActivatableTextField
func (p_ PDFAnnotation) ActivatableTextField() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("activatableTextField"))
	return rv
}

// A Boolean value that indicates whether the annotation is in a highlighted state, such as when the mouse is down on a link annotation.
//
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFAnnotation/isHighlighted
func (p_ PDFAnnotation) Highlighted() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("highlighted"))
	return rv
}


// SetHighlighted sets the value of the highlighted property.
// A Boolean value that indicates whether the annotation is in a highlighted state, such as when the mouse is down on a link annotation.

//
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFAnnotation/isHighlighted
func (p_ PDFAnnotation) SetHighlighted(value bool) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setHighlighted:"), value)
}

// A Boolean value that indicates whether the choice widget annotation is a list or a pop-up menu.
//
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFAnnotation/isListChoice
func (p_ PDFAnnotation) ListChoice() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("listChoice"))
	return rv
}


// SetListChoice sets the value of the listChoice property.
// A Boolean value that indicates whether the choice widget annotation is a list or a pop-up menu.

//
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFAnnotation/isListChoice
func (p_ PDFAnnotation) SetListChoice(value bool) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setListChoice:"), value)
}

// A Boolean value that indicates whether the text widget annotation displays multiple lines.
//
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFAnnotation/isMultiline
func (p_ PDFAnnotation) Multiline() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("multiline"))
	return rv
}


// SetMultiline sets the value of the multiline property.
// A Boolean value that indicates whether the text widget annotation displays multiple lines.

//
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFAnnotation/isMultiline
func (p_ PDFAnnotation) SetMultiline(value bool) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setMultiline:"), value)
}

// A Boolean value that indicates whether the pop-up annotation is in an opened state, displaying its text content, or in a closed state, displaying an icon.
//
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFAnnotation/isOpen
func (p_ PDFAnnotation) Open() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("open"))
	return rv
}


// SetOpen sets the value of the open property.
// A Boolean value that indicates whether the pop-up annotation is in an opened state, displaying its text content, or in a closed state, displaying an icon.

//
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFAnnotation/isOpen
func (p_ PDFAnnotation) SetOpen(value bool) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setOpen:"), value)
}

// A Boolean value that indicates whether the text widget annotation displays a password field using bullet characters.
//
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFAnnotation/isPasswordField
func (p_ PDFAnnotation) IsPasswordField() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("isPasswordField"))
	return rv
}

// A Boolean value that determines whether the widget is editable.
//
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFAnnotation/isReadOnly
func (p_ PDFAnnotation) ReadOnly() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("readOnly"))
	return rv
}


// SetReadOnly sets the value of the readOnly property.
// A Boolean value that determines whether the widget is editable.

//
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFAnnotation/isReadOnly
func (p_ PDFAnnotation) SetReadOnly(value bool) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setReadOnly:"), value)
}

// The markup type that the annotation displays, either highlight, strikethrough, underline, or redact.
//
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFAnnotation/markupType
func (p_ PDFAnnotation) MarkupType() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("markupType"))
	return rv
}


// SetMarkupType sets the value of the markupType property.
// The markup type that the annotation displays, either highlight, strikethrough, underline, or redact.

//
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFAnnotation/markupType
func (p_ PDFAnnotation) SetMarkupType(value unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setMarkupType:"), value)
}

// The maximum number of characters the text widget annotation allows.
//
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFAnnotation/maximumLength
func (p_ PDFAnnotation) MaximumLength() int {
	rv := objc.Send[int](p_.ID, objc.Sel("maximumLength"))
	return rv
}


// SetMaximumLength sets the value of the maximumLength property.
// The maximum number of characters the text widget annotation allows.

//
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFAnnotation/maximumLength
func (p_ PDFAnnotation) SetMaximumLength(value int) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setMaximumLength:"), value)
}

// Returns the modification date of the annotation.
//
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFAnnotation/modificationDate
func (p_ PDFAnnotation) ModificationDate() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("modificationDate"))
	return rv
}


// SetModificationDate sets the value of the modificationDate property.
// Returns the modification date of the annotation.

//
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFAnnotation/modificationDate
func (p_ PDFAnnotation) SetModificationDate(value unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setModificationDate:"), value)
}

// Returns the page that the annotation is associated with.
//
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFAnnotation/page
func (p_ PDFAnnotation) Page() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("page"))
	return rv
}


// SetPage sets the value of the page property.
// Returns the page that the annotation is associated with.

//
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFAnnotation/page
func (p_ PDFAnnotation) SetPage(value unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setPage:"), value)
}

// An array of bezier paths, in annotation-space coordinates, that compose the annotation.
//
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFAnnotation/paths
func (p_ PDFAnnotation) Paths() []UIBezierPath {
	rv := objc.Send[[]UIBezierPath](p_.ID, objc.Sel("paths"))
	return rv
}

// Returns the pop-up annotation associated with an annotation.
//
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFAnnotation/popup
func (p_ PDFAnnotation) Popup() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("popup"))
	return rv
}


// SetPopup sets the value of the popup property.
// Returns the pop-up annotation associated with an annotation.

//
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFAnnotation/popup
func (p_ PDFAnnotation) SetPopup(value unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setPopup:"), value)
}

// An array of values that represents the points bounding the marked-up text.
//
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFAnnotation/quadrilateralPoints
func (p_ PDFAnnotation) QuadrilateralPoints() []avfoundation.NSValue {
	rv := objc.Send[[]avfoundation.NSValue](p_.ID, objc.Sel("quadrilateralPoints"))
	return rv
}


// SetQuadrilateralPoints sets the value of the quadrilateralPoints property.
// An array of values that represents the points bounding the marked-up text.

//
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFAnnotation/quadrilateralPoints
func (p_ PDFAnnotation) SetQuadrilateralPoints(value []avfoundation.NSValue) {
	// Convert Go slice to NSArray
	var nsArray objc.ID
	if len(value) > 0 {
		nsArray = objc.ID(objc.GetClass("NSMutableArray")).Send(objc.Sel("arrayWithCapacity:"), len(value))
		for _, item := range value {
			nsArray.Send(objc.Sel("addObject:"), item)
		}
	} else {
		nsArray = objc.ID(objc.GetClass("NSArray")).Send(objc.Sel("array"))
	}
	objc.Send[objc.ID](p_.ID, objc.Sel("setQuadrilateralPoints:"), nsArray)
}

// A Boolean value that indicates whether radio buttons in a group turn on and off in unison.
//
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFAnnotation/radiosInUnison
func (p_ PDFAnnotation) RadiosInUnison() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("radiosInUnison"))
	return rv
}


// SetRadiosInUnison sets the value of the radiosInUnison property.
// A Boolean value that indicates whether radio buttons in a group turn on and off in unison.

//
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFAnnotation/radiosInUnison
func (p_ PDFAnnotation) SetRadiosInUnison(value bool) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setRadiosInUnison:"), value)
}

// Returns a Boolean value indicating whether the annotation should be displayed.
//
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFAnnotation/shouldDisplay
func (p_ PDFAnnotation) ShouldDisplay() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("shouldDisplay"))
	return rv
}


// SetShouldDisplay sets the value of the shouldDisplay property.
// Returns a Boolean value indicating whether the annotation should be displayed.

//
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFAnnotation/shouldDisplay
func (p_ PDFAnnotation) SetShouldDisplay(value bool) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setShouldDisplay:"), value)
}

// Returns a Boolean value indicating whether the annotation should appear when the document is printed.
//
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFAnnotation/shouldPrint
func (p_ PDFAnnotation) ShouldPrint() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("shouldPrint"))
	return rv
}


// SetShouldPrint sets the value of the shouldPrint property.
// Returns a Boolean value indicating whether the annotation should appear when the document is printed.

//
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFAnnotation/shouldPrint
func (p_ PDFAnnotation) SetShouldPrint(value bool) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setShouldPrint:"), value)
}

// The name of the stamp, a text or graphics annotation that emulates a rubber stamp effect.
//
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFAnnotation/stampName
func (p_ PDFAnnotation) StampName() string {
	rv := objc.Send[string](p_.ID, objc.Sel("stampName"))
	return rv
}


// SetStampName sets the value of the stampName property.
// The name of the stamp, a text or graphics annotation that emulates a rubber stamp effect.

//
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFAnnotation/stampName
func (p_ PDFAnnotation) SetStampName(value string) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setStampName:"), objc.String(value))
}

// The style of the line annotation’s starting point, such as square or filled arrowhead.
//
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFAnnotation/startLineStyle
func (p_ PDFAnnotation) StartLineStyle() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("startLineStyle"))
	return rv
}


// SetStartLineStyle sets the value of the startLineStyle property.
// The style of the line annotation’s starting point, such as square or filled arrowhead.

//
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFAnnotation/startLineStyle
func (p_ PDFAnnotation) SetStartLineStyle(value unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setStartLineStyle:"), value)
}

// The point where a line begins, in annotation-space coordinates.
//
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFAnnotation/startPoint
func (p_ PDFAnnotation) StartPoint() Point {
	rv := objc.Send[Point](p_.ID, objc.Sel("startPoint"))
	return rv
}


// SetStartPoint sets the value of the startPoint property.
// The point where a line begins, in annotation-space coordinates.

//
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFAnnotation/startPoint
func (p_ PDFAnnotation) SetStartPoint(value Point) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setStartPoint:"), value)
}

// Returns the type of the annotation.
//
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFAnnotation/type
func (p_ PDFAnnotation) Type() string {
	rv := objc.Send[string](p_.ID, objc.Sel("type"))
	return rv
}


// SetType sets the value of the type property.
// Returns the type of the annotation.

//
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFAnnotation/type
func (p_ PDFAnnotation) SetType(value string) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setType:"), objc.String(value))
}

// A URL for a link annotation.
//
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFAnnotation/url
func (p_ PDFAnnotation) URL() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("URL"))
	return rv
}


// SetURL sets the value of the URL property.
// A URL for a link annotation.

//
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFAnnotation/url
func (p_ PDFAnnotation) SetURL(value unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setURL:"), value)
}

// Returns the name of the user who created the annotation.
//
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFAnnotation/userName
func (p_ PDFAnnotation) UserName() string {
	rv := objc.Send[string](p_.ID, objc.Sel("userName"))
	return rv
}


// SetUserName sets the value of the userName property.
// Returns the name of the user who created the annotation.

//
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFAnnotation/userName
func (p_ PDFAnnotation) SetUserName(value string) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setUserName:"), objc.String(value))
}

// An array of strings that specifies the export values for items in a list or a pop-up menu.
//
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFAnnotation/values
func (p_ PDFAnnotation) Values() []string {
	rv := objc.Send[[]string](p_.ID, objc.Sel("values"))
	return rv
}


// SetValues sets the value of the values property.
// An array of strings that specifies the export values for items in a list or a pop-up menu.

//
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFAnnotation/values
func (p_ PDFAnnotation) SetValues(value []string) {
	// Convert Go slice to NSArray
	var nsArray objc.ID
	if len(value) > 0 {
		nsArray = objc.ID(objc.GetClass("NSMutableArray")).Send(objc.Sel("arrayWithCapacity:"), len(value))
		for _, item := range value {
			nsArray.Send(objc.Sel("addObject:"), item)
		}
	} else {
		nsArray = objc.ID(objc.GetClass("NSArray")).Send(objc.Sel("array"))
	}
	objc.Send[objc.ID](p_.ID, objc.Sel("setValues:"), nsArray)
}

// The type of button widget control, either radio button, push button, or checkbox.
//
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFAnnotation/widgetControlType
func (p_ PDFAnnotation) WidgetControlType() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("widgetControlType"))
	return rv
}


// SetWidgetControlType sets the value of the widgetControlType property.
// The type of button widget control, either radio button, push button, or checkbox.

//
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFAnnotation/widgetControlType
func (p_ PDFAnnotation) SetWidgetControlType(value unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setWidgetControlType:"), value)
}

// The string value that the widget reverts to when performing a reset form action.
//
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFAnnotation/widgetDefaultStringValue
func (p_ PDFAnnotation) WidgetDefaultStringValue() string {
	rv := objc.Send[string](p_.ID, objc.Sel("widgetDefaultStringValue"))
	return rv
}


// SetWidgetDefaultStringValue sets the value of the widgetDefaultStringValue property.
// The string value that the widget reverts to when performing a reset form action.

//
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFAnnotation/widgetDefaultStringValue
func (p_ PDFAnnotation) SetWidgetDefaultStringValue(value string) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setWidgetDefaultStringValue:"), objc.String(value))
}

// The type of widget annotation, such as button, choice, or text.
//
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFAnnotation/widgetFieldType
func (p_ PDFAnnotation) WidgetFieldType() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("widgetFieldType"))
	return rv
}


// SetWidgetFieldType sets the value of the widgetFieldType property.
// The type of widget annotation, such as button, choice, or text.

//
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFAnnotation/widgetFieldType
func (p_ PDFAnnotation) SetWidgetFieldType(value unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setWidgetFieldType:"), value)
}

// The string value of the widget annotation.
//
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFAnnotation/widgetStringValue
func (p_ PDFAnnotation) WidgetStringValue() string {
	rv := objc.Send[string](p_.ID, objc.Sel("widgetStringValue"))
	return rv
}


// SetWidgetStringValue sets the value of the widgetStringValue property.
// The string value of the widget annotation.

//
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFAnnotation/widgetStringValue
func (p_ PDFAnnotation) SetWidgetStringValue(value string) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setWidgetStringValue:"), objc.String(value))
}


