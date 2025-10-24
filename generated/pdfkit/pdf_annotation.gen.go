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

/* debug [class.gen.go]: Generating class PDFAnnotation */


/* debug [class_header]: Header for PDFAnnotation */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for PDFAnnotation */
// An interface definition for the [PDFAnnotation] class.
type IPDFAnnotation interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for PDFAnnotation */
	// properties:
	Action() IPDFAction
	SetAction(value IPDFAction)
	Alignment() TextAlignment /* not a class type */
	SetAlignment(value TextAlignment /* not a class type */)
	AllowsToggleToOff() bool
	SetAllowsToggleToOff(value bool)
	AnnotationKeyValues() objc.IObject /* cross-framework: NSDictionary */
	BackgroundColor() appkit.Color
	SetBackgroundColor(value appkit.Color)
	Border() IPDFBorder
	SetBorder(value IPDFBorder)
	Bounds() corefoundation.CGRect
	SetBounds(value corefoundation.CGRect)
	ButtonWidgetState() PDFWidgetCellState
	SetButtonWidgetState(value PDFWidgetCellState)
	ButtonWidgetStateString() objc.IObject /* cross-framework: NSString */
	SetButtonWidgetStateString(value objc.IObject /* cross-framework: NSString */)
	Caption() objc.IObject /* cross-framework: NSString */
	SetCaption(value objc.IObject /* cross-framework: NSString */)
	Choices() []string
	SetChoices(value []string)
	Color() appkit.Color
	SetColor(value appkit.Color)
	Contents() objc.IObject /* cross-framework: NSString */
	SetContents(value objc.IObject /* cross-framework: NSString */)
	Destination() IPDFDestination
	SetDestination(value IPDFDestination)
	EndLineStyle() PDFLineStyle
	SetEndLineStyle(value PDFLineStyle)
	EndPoint() corefoundation.CGPoint
	SetEndPoint(value corefoundation.CGPoint)
	FieldName() objc.IObject /* cross-framework: NSString */
	SetFieldName(value objc.IObject /* cross-framework: NSString */)
	Font() appkit.Font
	SetFont(value appkit.Font)
	FontColor() appkit.Color
	SetFontColor(value appkit.Color)
	HasAppearanceStream() bool
	Comb() bool
	SetComb(value bool)
	IconType() PDFTextAnnotationIconType
	SetIconType(value PDFTextAnnotationIconType)
	InteriorColor() appkit.Color
	SetInteriorColor(value appkit.Color)
	ActivatableTextField() bool
	Highlighted() bool
	SetHighlighted(value bool)
	ListChoice() bool
	SetListChoice(value bool)
	Multiline() bool
	SetMultiline(value bool)
	Open() bool
	SetOpen(value bool)
	IsPasswordField() bool
	ReadOnly() bool
	SetReadOnly(value bool)
	MarkupType() PDFMarkupType
	SetMarkupType(value PDFMarkupType)
	MaximumLength() int
	SetMaximumLength(value int)
	ModificationDate() objc.IObject /* cross-framework: NSDate */
	SetModificationDate(value objc.IObject /* cross-framework: NSDate */)
	MouseUpAction() IPDFAction
	SetMouseUpAction(value IPDFAction)
	Page() IPDFPage
	SetPage(value IPDFPage)
	Paths() []appkit.BezierPath
	Popup() IPDFAnnotation
	SetPopup(value IPDFAnnotation)
	QuadrilateralPoints() []foundation.Value
	SetQuadrilateralPoints(value []foundation.Value)
	RadiosInUnison() bool
	SetRadiosInUnison(value bool)
	ShouldDisplay() bool
	SetShouldDisplay(value bool)
	ShouldPrint() bool
	SetShouldPrint(value bool)
	StampName() objc.IObject /* cross-framework: NSString */
	SetStampName(value objc.IObject /* cross-framework: NSString */)
	StartLineStyle() PDFLineStyle
	SetStartLineStyle(value PDFLineStyle)
	StartPoint() corefoundation.CGPoint
	SetStartPoint(value corefoundation.CGPoint)
	ToolTip() objc.IObject /* cross-framework: NSString */
	Type() objc.IObject /* cross-framework: NSString */
	SetType(value objc.IObject /* cross-framework: NSString */)
	URL() objc.IObject /* cross-framework: NSURL */
	SetURL(value objc.IObject /* cross-framework: NSURL */)
	UserName() objc.IObject /* cross-framework: NSString */
	SetUserName(value objc.IObject /* cross-framework: NSString */)
	Values() []string
	SetValues(value []string)
	WidgetControlType() PDFWidgetControlType
	SetWidgetControlType(value PDFWidgetControlType)
	WidgetDefaultStringValue() objc.IObject /* cross-framework: NSString */
	SetWidgetDefaultStringValue(value objc.IObject /* cross-framework: NSString */)
	WidgetFieldType() PDFAnnotationWidgetSubtype /* typedef */
	SetWidgetFieldType(value PDFAnnotationWidgetSubtype /* typedef */)
	WidgetStringValue() objc.IObject /* cross-framework: NSString */
	SetWidgetStringValue(value objc.IObject /* cross-framework: NSString */)
	HasComb() bool
	SetHasComb(value bool)
	IsActivatableTextField() bool
	SetIsActivatableTextField(value bool)
	IsHighlighted() bool
	SetIsHighlighted(value bool)
	IsListChoice() bool
	SetIsListChoice(value bool)
	IsMultiline() bool
	SetIsMultiline(value bool)
	IsOpen() bool
	SetIsOpen(value bool)
	IsReadOnly() bool
	SetIsReadOnly(value bool)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for PDFAnnotation */
	// methods:
	AddBezierPath(path appkit.BezierPath)
	DrawWithBoxInContext(box PDFDisplayBox, context ContextRef /* not a class type */)
	RemoveBezierPath(path appkit.BezierPath)
	RemoveValueForAnnotationKey(key PDFAnnotationKey /* typedef */)
	SetBooleanForAnnotationKey(value bool, key PDFAnnotationKey /* typedef */) bool
	SetRectForAnnotationKey(value corefoundation.CGRect, key PDFAnnotationKey /* typedef */) bool
	SetValueForAnnotationKey(value objc.IObject, key PDFAnnotationKey /* typedef */) bool
	ValueForAnnotationKey(key PDFAnnotationKey /* typedef */) objc.ID
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for PDFAnnotation */
// Alloc allocates a new instance without initialization.
func (pc _PDFAnnotationClass) Alloc() PDFAnnotation {
	rv := objc.Send[PDFAnnotation](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for PDFAnnotation */
// An annotation in a PDF document.
//
// In addition to its primary textual content, a PDF file can contain annotations that represent links, form elements, highlighting circles, textual notes, and so on. Each annotation has a specific location on a page and may offer interactivity with the user.


// An annotation in a PDF document.
//
// [Full Topic]
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for PDFAnnotation */

// Creates a PDF annotation object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFAnnotation/init(bounds:)
func NewPDFAnnotationWithBounds(bounds Rect /* not a class type */) PDFAnnotation {
	instance := getPDFAnnotationClass().Alloc()
	rv := objc.Send[PDFAnnotation](instance.ID, objc.Sel("initWithBounds:"), bounds)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewPDFAnnotationWithBounds */


// Creates a PDF annotation with the specified bounds, type, and optional properties.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFAnnotation/init(bounds:forType:withProperties:)
func NewPDFAnnotationWithBoundsForTypeWithProperties(bounds corefoundation.CGRect, annotationType PDFAnnotationSubtype /* typedef */, properties objc.IObject /* cross-framework: NSDictionary */) PDFAnnotation {
	instance := getPDFAnnotationClass().Alloc()
	rv := objc.Send[PDFAnnotation](instance.ID, objc.Sel("initWithBounds:forType:withProperties:"), bounds, annotationType, properties)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewPDFAnnotationWithBoundsForTypeWithProperties */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFAnnotation/init(dictionary:forPage:)
func NewPDFAnnotationWithDictionaryForPage(dictionary objc.IObject /* cross-framework: NSDictionary */, page IPDFPage) PDFAnnotation {
	instance := getPDFAnnotationClass().Alloc()
	rv := objc.Send[PDFAnnotation](instance.ID, objc.Sel("initWithDictionary:forPage:"), dictionary, page)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewPDFAnnotationWithDictionaryForPage */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for PDFAnnotation */

// Returns a line style that corresponds to the specified name.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFAnnotation/lineStyle(fromName:)
func (pc _PDFAnnotationClass) LineStyleFromName(name objc.IObject /* cross-framework: NSString */) PDFLineStyle {
	rv := objc.Send[PDFLineStyle](objc.ID(pc.class), objc.Sel("lineStyleFromName:"), name)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=LineStyleFromName) */


// Returns the name of the line style, which matches the definition in the Adobe PDF Specification.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFAnnotation/name(for:)
func (pc _PDFAnnotationClass) NameForLineStyle(style PDFLineStyle) foundation.String {
	rv := objc.Send[foundation.String](objc.ID(pc.class), objc.Sel("nameForLineStyle:"), style)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=NameForLineStyle) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for PDFAnnotation */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for PDFAnnotation */

// Adds a bezier path to the ink annotation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFAnnotation/add(_:)
func (p_ PDFAnnotation) AddBezierPath(path appkit.BezierPath) {
	objc.Send[objc.ID](p_.ID, objc.Sel("addBezierPath:"), path)
}/* debug [instance_methods/method]: AddBezierPath */


// Draws the annotation in a graphics context using page-space coordinates relative to the origin of the specified box.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFAnnotation/draw(with:in:)
func (p_ PDFAnnotation) DrawWithBoxInContext(box PDFDisplayBox, context ContextRef /* not a class type */) {
	objc.Send[objc.ID](p_.ID, objc.Sel("drawWithBox:inContext:"), box, context)
}/* debug [instance_methods/method]: DrawWithBoxInContext */


// Removes a bezier path from an ink annotation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFAnnotation/remove(_:)
func (p_ PDFAnnotation) RemoveBezierPath(path appkit.BezierPath) {
	objc.Send[objc.ID](p_.ID, objc.Sel("removeBezierPath:"), path)
}/* debug [instance_methods/method]: RemoveBezierPath */


// Removes a value from the annotation’s dictionary.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFAnnotation/removeValue(forAnnotationKey:)
func (p_ PDFAnnotation) RemoveValueForAnnotationKey(key PDFAnnotationKey /* typedef */) {
	objc.Send[objc.ID](p_.ID, objc.Sel("removeValueForAnnotationKey:"), key)
}/* debug [instance_methods/method]: RemoveValueForAnnotationKey */


// Sets a Boolean value in the annotation’s dictionary.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFAnnotation/setBoolean(_:forAnnotationKey:)
func (p_ PDFAnnotation) SetBooleanForAnnotationKey(value bool, key PDFAnnotationKey /* typedef */) bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("setBoolean:forAnnotationKey:"), value, key)
	return rv
}/* debug [instance_methods/method]: SetBooleanForAnnotationKey */


// Sets a rectangle value in the annotation’s dictionary.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFAnnotation/setRect(_:forAnnotationKey:)
func (p_ PDFAnnotation) SetRectForAnnotationKey(value corefoundation.CGRect, key PDFAnnotationKey /* typedef */) bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("setRect:forAnnotationKey:"), value, key)
	return rv
}/* debug [instance_methods/method]: SetRectForAnnotationKey */


// Sets a value in the annotation’s dictionary.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFAnnotation/setValue(_:forAnnotationKey:)
func (p_ PDFAnnotation) SetValueForAnnotationKey(value objc.IObject, key PDFAnnotationKey /* typedef */) bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("setValue:forAnnotationKey:"), value, key)
	return rv
}/* debug [instance_methods/method]: SetValueForAnnotationKey */


// Returns a deep copy of the key-value pairs of properties for the specified key.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFAnnotation/value(forAnnotationKey:)
func (p_ PDFAnnotation) ValueForAnnotationKey(key PDFAnnotationKey /* typedef */) objc.ID {
	rv := objc.Send[objc.ID](p_.ID, objc.Sel("valueForAnnotationKey:"), key)
	return rv
}/* debug [instance_methods/method]: ValueForAnnotationKey */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for PDFAnnotation */

// An object that represents an action for a PDF element, such as a link annotation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFAnnotation/action
func (p_ PDFAnnotation) Action() IPDFAction {
	rv := objc.Send[PDFAction](p_.ID, objc.Sel("action"))
	return rv
}/* debug [instance_properties/getter]: action */


// An object that represents an action for a PDF element, such as a link annotation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFAnnotation/action
func (p_ PDFAnnotation) SetAction(value IPDFAction) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setAction:"), value)
}/* debug [instance_properties/setter]: action */


// The alignment of the free text and text widget annotation’s text content.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFAnnotation/alignment
func (p_ PDFAnnotation) Alignment() TextAlignment /* not a class type */ {
	rv := objc.Send[TextAlignment](p_.ID, objc.Sel("alignment"))
	return rv
}/* debug [instance_properties/getter]: alignment */


// The alignment of the free text and text widget annotation’s text content.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFAnnotation/alignment
func (p_ PDFAnnotation) SetAlignment(value TextAlignment /* not a class type */) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setAlignment:"), value)
}/* debug [instance_properties/setter]: alignment */


// A Boolean value that indicates whether clicking or tapping a selected radio button toggles it to an unselected state.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFAnnotation/allowsToggleToOff
func (p_ PDFAnnotation) AllowsToggleToOff() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("allowsToggleToOff"))
	return rv
}/* debug [instance_properties/getter]: allowsToggleToOff */


// A Boolean value that indicates whether clicking or tapping a selected radio button toggles it to an unselected state.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFAnnotation/allowsToggleToOff
func (p_ PDFAnnotation) SetAllowsToggleToOff(value bool) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setAllowsToggleToOff:"), value)
}/* debug [instance_properties/setter]: allowsToggleToOff */


// A dictionary that contains a deep copy of the widget’s properties.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFAnnotation/annotationKeyValues
func (p_ PDFAnnotation) AnnotationKeyValues() objc.IObject /* cross-framework: NSDictionary */ {
	rv := objc.Send[foundation.NSDictionary](p_.ID, objc.Sel("annotationKeyValues"))
	return rv
}/* debug [instance_properties/getter]: annotationKeyValues */


// The color of the widget’s background.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFAnnotation/backgroundColor
func (p_ PDFAnnotation) BackgroundColor() appkit.Color {
	rv := objc.Send[appkit.Color](p_.ID, objc.Sel("backgroundColor"))
	return rv
}/* debug [instance_properties/getter]: backgroundColor */


// The color of the widget’s background.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFAnnotation/backgroundColor
func (p_ PDFAnnotation) SetBackgroundColor(value appkit.Color) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setBackgroundColor:"), value)
}/* debug [instance_properties/setter]: backgroundColor */


// Sets the border style for the annotation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFAnnotation/border
func (p_ PDFAnnotation) Border() IPDFBorder {
	rv := objc.Send[PDFBorder](p_.ID, objc.Sel("border"))
	return rv
}/* debug [instance_properties/getter]: border */


// Sets the border style for the annotation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFAnnotation/border
func (p_ PDFAnnotation) SetBorder(value IPDFBorder) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setBorder:"), value)
}/* debug [instance_properties/setter]: border */


// Returns the bounding box for the annotation in page space.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFAnnotation/bounds
func (p_ PDFAnnotation) Bounds() corefoundation.CGRect {
	rv := objc.Send[corefoundation.CGRect](p_.ID, objc.Sel("bounds"))
	return rv
}/* debug [instance_properties/getter]: bounds */


// Returns the bounding box for the annotation in page space.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFAnnotation/bounds
func (p_ PDFAnnotation) SetBounds(value corefoundation.CGRect) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setBounds:"), value)
}/* debug [instance_properties/setter]: bounds */


// The current state of the button widget annotation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFAnnotation/buttonWidgetState
func (p_ PDFAnnotation) ButtonWidgetState() PDFWidgetCellState {
	rv := objc.Send[PDFWidgetCellState](p_.ID, objc.Sel("buttonWidgetState"))
	return rv
}/* debug [instance_properties/getter]: buttonWidgetState */


// The current state of the button widget annotation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFAnnotation/buttonWidgetState
func (p_ PDFAnnotation) SetButtonWidgetState(value PDFWidgetCellState) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setButtonWidgetState:"), value)
}/* debug [instance_properties/setter]: buttonWidgetState */


// A string value that differentiates button widgets in the same group, such as to identify mutually exclusive radio buttons from each other.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFAnnotation/buttonWidgetStateString
func (p_ PDFAnnotation) ButtonWidgetStateString() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](p_.ID, objc.Sel("buttonWidgetStateString"))
	return rv
}/* debug [instance_properties/getter]: buttonWidgetStateString */


// A string value that differentiates button widgets in the same group, such as to identify mutually exclusive radio buttons from each other.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFAnnotation/buttonWidgetStateString
func (p_ PDFAnnotation) SetButtonWidgetStateString(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setButtonWidgetStateString:"), value)
}/* debug [instance_properties/setter]: buttonWidgetStateString */


// The title of push button widget annotations.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFAnnotation/caption
func (p_ PDFAnnotation) Caption() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](p_.ID, objc.Sel("caption"))
	return rv
}/* debug [instance_properties/getter]: caption */


// The title of push button widget annotations.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFAnnotation/caption
func (p_ PDFAnnotation) SetCaption(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setCaption:"), value)
}/* debug [instance_properties/setter]: caption */


// An array of strings that specifies the options in either a list or a pop-up menu.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFAnnotation/choices
func (p_ PDFAnnotation) Choices() []string {
	rv := objc.Send[[]string](p_.ID, objc.Sel("choices"))
	return rv
}/* debug [instance_properties/getter]: choices */


// An array of strings that specifies the options in either a list or a pop-up menu.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFAnnotation/choices
func (p_ PDFAnnotation) SetChoices(value []string) {
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
}/* debug [instance_properties/setter]: choices */


// Sets the stroke color for the annotation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFAnnotation/color
func (p_ PDFAnnotation) Color() appkit.Color {
	rv := objc.Send[appkit.Color](p_.ID, objc.Sel("color"))
	return rv
}/* debug [instance_properties/getter]: color */


// Sets the stroke color for the annotation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFAnnotation/color
func (p_ PDFAnnotation) SetColor(value appkit.Color) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setColor:"), value)
}/* debug [instance_properties/setter]: color */


// Returns the textual content (if any) associated with the annotation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFAnnotation/contents
func (p_ PDFAnnotation) Contents() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](p_.ID, objc.Sel("contents"))
	return rv
}/* debug [instance_properties/getter]: contents */


// Returns the textual content (if any) associated with the annotation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFAnnotation/contents
func (p_ PDFAnnotation) SetContents(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setContents:"), value)
}/* debug [instance_properties/setter]: contents */


// The destination for a link annotation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFAnnotation/destination
func (p_ PDFAnnotation) Destination() IPDFDestination {
	rv := objc.Send[PDFDestination](p_.ID, objc.Sel("destination"))
	return rv
}/* debug [instance_properties/getter]: destination */


// The destination for a link annotation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFAnnotation/destination
func (p_ PDFAnnotation) SetDestination(value IPDFDestination) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setDestination:"), value)
}/* debug [instance_properties/setter]: destination */


// The style of the line annotation’s ending point, such as square or filled arrowhead.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFAnnotation/endLineStyle
func (p_ PDFAnnotation) EndLineStyle() PDFLineStyle {
	rv := objc.Send[PDFLineStyle](p_.ID, objc.Sel("endLineStyle"))
	return rv
}/* debug [instance_properties/getter]: endLineStyle */


// The style of the line annotation’s ending point, such as square or filled arrowhead.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFAnnotation/endLineStyle
func (p_ PDFAnnotation) SetEndLineStyle(value PDFLineStyle) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setEndLineStyle:"), value)
}/* debug [instance_properties/setter]: endLineStyle */


// The point where a line ends, in annotation-space coordinates.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFAnnotation/endPoint
func (p_ PDFAnnotation) EndPoint() corefoundation.CGPoint {
	rv := objc.Send[corefoundation.CGPoint](p_.ID, objc.Sel("endPoint"))
	return rv
}/* debug [instance_properties/getter]: endPoint */


// The point where a line ends, in annotation-space coordinates.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFAnnotation/endPoint
func (p_ PDFAnnotation) SetEndPoint(value corefoundation.CGPoint) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setEndPoint:"), value)
}/* debug [instance_properties/setter]: endPoint */


// The widget identifier for form annotation actions and behaviors.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFAnnotation/fieldName
func (p_ PDFAnnotation) FieldName() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](p_.ID, objc.Sel("fieldName"))
	return rv
}/* debug [instance_properties/getter]: fieldName */


// The widget identifier for form annotation actions and behaviors.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFAnnotation/fieldName
func (p_ PDFAnnotation) SetFieldName(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setFieldName:"), value)
}/* debug [instance_properties/setter]: fieldName */


// The font the annotation uses to display text.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFAnnotation/font
func (p_ PDFAnnotation) Font() appkit.Font {
	rv := objc.Send[appkit.Font](p_.ID, objc.Sel("font"))
	return rv
}/* debug [instance_properties/getter]: font */


// The font the annotation uses to display text.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFAnnotation/font
func (p_ PDFAnnotation) SetFont(value appkit.Font) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setFont:"), value)
}/* debug [instance_properties/setter]: font */


// The font color the annotation uses to display text.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFAnnotation/fontColor
func (p_ PDFAnnotation) FontColor() appkit.Color {
	rv := objc.Send[appkit.Color](p_.ID, objc.Sel("fontColor"))
	return rv
}/* debug [instance_properties/getter]: fontColor */


// The font color the annotation uses to display text.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFAnnotation/fontColor
func (p_ PDFAnnotation) SetFontColor(value appkit.Color) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setFontColor:"), value)
}/* debug [instance_properties/setter]: fontColor */


// Returns a Boolean value that indicates whether the annotation has an appearance stream associated with it.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFAnnotation/hasAppearanceStream
func (p_ PDFAnnotation) HasAppearanceStream() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("hasAppearanceStream"))
	return rv
}/* debug [instance_properties/getter]: hasAppearanceStream */


// A Boolean value that indicates whether the annotation divides the text widget’s bounds into equally spaced segments, such as in a form entry field.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFAnnotation/hasComb
func (p_ PDFAnnotation) Comb() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("comb"))
	return rv
}/* debug [instance_properties/getter]: comb */


// A Boolean value that indicates whether the annotation divides the text widget’s bounds into equally spaced segments, such as in a form entry field.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFAnnotation/hasComb
func (p_ PDFAnnotation) SetComb(value bool) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setComb:"), value)
}/* debug [instance_properties/setter]: comb */


// The type of icon to display for a pop-up text annotation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFAnnotation/iconType
func (p_ PDFAnnotation) IconType() PDFTextAnnotationIconType {
	rv := objc.Send[PDFTextAnnotationIconType](p_.ID, objc.Sel("iconType"))
	return rv
}/* debug [instance_properties/getter]: iconType */


// The type of icon to display for a pop-up text annotation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFAnnotation/iconType
func (p_ PDFAnnotation) SetIconType(value PDFTextAnnotationIconType) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setIconType:"), value)
}/* debug [instance_properties/setter]: iconType */


// The fill color for drawing a circle, line, or square annotation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFAnnotation/interiorColor
func (p_ PDFAnnotation) InteriorColor() appkit.Color {
	rv := objc.Send[appkit.Color](p_.ID, objc.Sel("interiorColor"))
	return rv
}/* debug [instance_properties/getter]: interiorColor */


// The fill color for drawing a circle, line, or square annotation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFAnnotation/interiorColor
func (p_ PDFAnnotation) SetInteriorColor(value appkit.Color) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setInteriorColor:"), value)
}/* debug [instance_properties/setter]: interiorColor */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFAnnotation/isActivatableTextField
func (p_ PDFAnnotation) ActivatableTextField() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("activatableTextField"))
	return rv
}/* debug [instance_properties/getter]: activatableTextField */


// A Boolean value that indicates whether the annotation is in a highlighted state, such as when the mouse is down on a link annotation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFAnnotation/isHighlighted
func (p_ PDFAnnotation) Highlighted() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("highlighted"))
	return rv
}/* debug [instance_properties/getter]: highlighted */


// A Boolean value that indicates whether the annotation is in a highlighted state, such as when the mouse is down on a link annotation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFAnnotation/isHighlighted
func (p_ PDFAnnotation) SetHighlighted(value bool) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setHighlighted:"), value)
}/* debug [instance_properties/setter]: highlighted */


// A Boolean value that indicates whether the choice widget annotation is a list or a pop-up menu.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFAnnotation/isListChoice
func (p_ PDFAnnotation) ListChoice() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("listChoice"))
	return rv
}/* debug [instance_properties/getter]: listChoice */


// A Boolean value that indicates whether the choice widget annotation is a list or a pop-up menu.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFAnnotation/isListChoice
func (p_ PDFAnnotation) SetListChoice(value bool) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setListChoice:"), value)
}/* debug [instance_properties/setter]: listChoice */


// A Boolean value that indicates whether the text widget annotation displays multiple lines.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFAnnotation/isMultiline
func (p_ PDFAnnotation) Multiline() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("multiline"))
	return rv
}/* debug [instance_properties/getter]: multiline */


// A Boolean value that indicates whether the text widget annotation displays multiple lines.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFAnnotation/isMultiline
func (p_ PDFAnnotation) SetMultiline(value bool) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setMultiline:"), value)
}/* debug [instance_properties/setter]: multiline */


// A Boolean value that indicates whether the pop-up annotation is in an opened state, displaying its text content, or in a closed state, displaying an icon.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFAnnotation/isOpen
func (p_ PDFAnnotation) Open() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("open"))
	return rv
}/* debug [instance_properties/getter]: open */


// A Boolean value that indicates whether the pop-up annotation is in an opened state, displaying its text content, or in a closed state, displaying an icon.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFAnnotation/isOpen
func (p_ PDFAnnotation) SetOpen(value bool) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setOpen:"), value)
}/* debug [instance_properties/setter]: open */


// A Boolean value that indicates whether the text widget annotation displays a password field using bullet characters.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFAnnotation/isPasswordField
func (p_ PDFAnnotation) IsPasswordField() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("isPasswordField"))
	return rv
}/* debug [instance_properties/getter]: isPasswordField */


// A Boolean value that determines whether the widget is editable.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFAnnotation/isReadOnly
func (p_ PDFAnnotation) ReadOnly() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("readOnly"))
	return rv
}/* debug [instance_properties/getter]: readOnly */


// A Boolean value that determines whether the widget is editable.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFAnnotation/isReadOnly
func (p_ PDFAnnotation) SetReadOnly(value bool) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setReadOnly:"), value)
}/* debug [instance_properties/setter]: readOnly */


// The markup type that the annotation displays, either highlight, strikethrough, underline, or redact.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFAnnotation/markupType
func (p_ PDFAnnotation) MarkupType() PDFMarkupType {
	rv := objc.Send[PDFMarkupType](p_.ID, objc.Sel("markupType"))
	return rv
}/* debug [instance_properties/getter]: markupType */


// The markup type that the annotation displays, either highlight, strikethrough, underline, or redact.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFAnnotation/markupType
func (p_ PDFAnnotation) SetMarkupType(value PDFMarkupType) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setMarkupType:"), value)
}/* debug [instance_properties/setter]: markupType */


// The maximum number of characters the text widget annotation allows.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFAnnotation/maximumLength
func (p_ PDFAnnotation) MaximumLength() int {
	rv := objc.Send[int](p_.ID, objc.Sel("maximumLength"))
	return rv
}/* debug [instance_properties/getter]: maximumLength */


// The maximum number of characters the text widget annotation allows.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFAnnotation/maximumLength
func (p_ PDFAnnotation) SetMaximumLength(value int) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setMaximumLength:"), value)
}/* debug [instance_properties/setter]: maximumLength */


// Returns the modification date of the annotation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFAnnotation/modificationDate
func (p_ PDFAnnotation) ModificationDate() objc.IObject /* cross-framework: NSDate */ {
	rv := objc.Send[foundation.NSDate](p_.ID, objc.Sel("modificationDate"))
	return rv
}/* debug [instance_properties/getter]: modificationDate */


// Returns the modification date of the annotation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFAnnotation/modificationDate
func (p_ PDFAnnotation) SetModificationDate(value objc.IObject /* cross-framework: NSDate */) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setModificationDate:"), value)
}/* debug [instance_properties/setter]: modificationDate */


// The action to perform when a user releases the mouse button within an annotation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFAnnotation/mouseUpAction
func (p_ PDFAnnotation) MouseUpAction() IPDFAction {
	rv := objc.Send[PDFAction](p_.ID, objc.Sel("mouseUpAction"))
	return rv
}/* debug [instance_properties/getter]: mouseUpAction */


// The action to perform when a user releases the mouse button within an annotation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFAnnotation/mouseUpAction
func (p_ PDFAnnotation) SetMouseUpAction(value IPDFAction) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setMouseUpAction:"), value)
}/* debug [instance_properties/setter]: mouseUpAction */


// Returns the page that the annotation is associated with.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFAnnotation/page
func (p_ PDFAnnotation) Page() IPDFPage {
	rv := objc.Send[PDFPage](p_.ID, objc.Sel("page"))
	return rv
}/* debug [instance_properties/getter]: page */


// Returns the page that the annotation is associated with.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFAnnotation/page
func (p_ PDFAnnotation) SetPage(value IPDFPage) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setPage:"), value)
}/* debug [instance_properties/setter]: page */


// An array of bezier paths, in annotation-space coordinates, that compose the annotation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFAnnotation/paths
func (p_ PDFAnnotation) Paths() []appkit.BezierPath {
	rv := objc.Send[[]appkit.BezierPath](p_.ID, objc.Sel("paths"))
	return rv
}/* debug [instance_properties/getter]: paths */


// Returns the pop-up annotation associated with an annotation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFAnnotation/popup
func (p_ PDFAnnotation) Popup() IPDFAnnotation {
	rv := objc.Send[PDFAnnotation](p_.ID, objc.Sel("popup"))
	return rv
}/* debug [instance_properties/getter]: popup */


// Returns the pop-up annotation associated with an annotation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFAnnotation/popup
func (p_ PDFAnnotation) SetPopup(value IPDFAnnotation) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setPopup:"), value)
}/* debug [instance_properties/setter]: popup */


// An array of values that represents the points bounding the marked-up text.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFAnnotation/quadrilateralPoints
func (p_ PDFAnnotation) QuadrilateralPoints() []foundation.Value {
	rv := objc.Send[[]foundation.Value](p_.ID, objc.Sel("quadrilateralPoints"))
	return rv
}/* debug [instance_properties/getter]: quadrilateralPoints */


// An array of values that represents the points bounding the marked-up text.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFAnnotation/quadrilateralPoints
func (p_ PDFAnnotation) SetQuadrilateralPoints(value []foundation.Value) {
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
}/* debug [instance_properties/setter]: quadrilateralPoints */


// A Boolean value that indicates whether radio buttons in a group turn on and off in unison.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFAnnotation/radiosInUnison
func (p_ PDFAnnotation) RadiosInUnison() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("radiosInUnison"))
	return rv
}/* debug [instance_properties/getter]: radiosInUnison */


// A Boolean value that indicates whether radio buttons in a group turn on and off in unison.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFAnnotation/radiosInUnison
func (p_ PDFAnnotation) SetRadiosInUnison(value bool) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setRadiosInUnison:"), value)
}/* debug [instance_properties/setter]: radiosInUnison */


// Returns a Boolean value indicating whether the annotation should be displayed.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFAnnotation/shouldDisplay
func (p_ PDFAnnotation) ShouldDisplay() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("shouldDisplay"))
	return rv
}/* debug [instance_properties/getter]: shouldDisplay */


// Returns a Boolean value indicating whether the annotation should be displayed.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFAnnotation/shouldDisplay
func (p_ PDFAnnotation) SetShouldDisplay(value bool) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setShouldDisplay:"), value)
}/* debug [instance_properties/setter]: shouldDisplay */


// Returns a Boolean value indicating whether the annotation should appear when the document is printed.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFAnnotation/shouldPrint
func (p_ PDFAnnotation) ShouldPrint() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("shouldPrint"))
	return rv
}/* debug [instance_properties/getter]: shouldPrint */


// Returns a Boolean value indicating whether the annotation should appear when the document is printed.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFAnnotation/shouldPrint
func (p_ PDFAnnotation) SetShouldPrint(value bool) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setShouldPrint:"), value)
}/* debug [instance_properties/setter]: shouldPrint */


// The name of the stamp, a text or graphics annotation that emulates a rubber stamp effect.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFAnnotation/stampName
func (p_ PDFAnnotation) StampName() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](p_.ID, objc.Sel("stampName"))
	return rv
}/* debug [instance_properties/getter]: stampName */


// The name of the stamp, a text or graphics annotation that emulates a rubber stamp effect.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFAnnotation/stampName
func (p_ PDFAnnotation) SetStampName(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setStampName:"), value)
}/* debug [instance_properties/setter]: stampName */


// The style of the line annotation’s starting point, such as square or filled arrowhead.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFAnnotation/startLineStyle
func (p_ PDFAnnotation) StartLineStyle() PDFLineStyle {
	rv := objc.Send[PDFLineStyle](p_.ID, objc.Sel("startLineStyle"))
	return rv
}/* debug [instance_properties/getter]: startLineStyle */


// The style of the line annotation’s starting point, such as square or filled arrowhead.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFAnnotation/startLineStyle
func (p_ PDFAnnotation) SetStartLineStyle(value PDFLineStyle) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setStartLineStyle:"), value)
}/* debug [instance_properties/setter]: startLineStyle */


// The point where a line begins, in annotation-space coordinates.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFAnnotation/startPoint
func (p_ PDFAnnotation) StartPoint() corefoundation.CGPoint {
	rv := objc.Send[corefoundation.CGPoint](p_.ID, objc.Sel("startPoint"))
	return rv
}/* debug [instance_properties/getter]: startPoint */


// The point where a line begins, in annotation-space coordinates.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFAnnotation/startPoint
func (p_ PDFAnnotation) SetStartPoint(value corefoundation.CGPoint) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setStartPoint:"), value)
}/* debug [instance_properties/setter]: startPoint */


// Returns text for display as a help tag.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFAnnotation/toolTip
func (p_ PDFAnnotation) ToolTip() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](p_.ID, objc.Sel("toolTip"))
	return rv
}/* debug [instance_properties/getter]: toolTip */


// Returns the type of the annotation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFAnnotation/type
func (p_ PDFAnnotation) Type() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](p_.ID, objc.Sel("type"))
	return rv
}/* debug [instance_properties/getter]: type */


// Returns the type of the annotation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFAnnotation/type
func (p_ PDFAnnotation) SetType(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setType:"), value)
}/* debug [instance_properties/setter]: type */


// A URL for a link annotation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFAnnotation/url
func (p_ PDFAnnotation) URL() objc.IObject /* cross-framework: NSURL */ {
	rv := objc.Send[foundation.NSURL](p_.ID, objc.Sel("URL"))
	return rv
}/* debug [instance_properties/getter]: URL */


// A URL for a link annotation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFAnnotation/url
func (p_ PDFAnnotation) SetURL(value objc.IObject /* cross-framework: NSURL */) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setURL:"), value)
}/* debug [instance_properties/setter]: URL */


// Returns the name of the user who created the annotation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFAnnotation/userName
func (p_ PDFAnnotation) UserName() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](p_.ID, objc.Sel("userName"))
	return rv
}/* debug [instance_properties/getter]: userName */


// Returns the name of the user who created the annotation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFAnnotation/userName
func (p_ PDFAnnotation) SetUserName(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setUserName:"), value)
}/* debug [instance_properties/setter]: userName */


// An array of strings that specifies the export values for items in a list or a pop-up menu.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFAnnotation/values
func (p_ PDFAnnotation) Values() []string {
	rv := objc.Send[[]string](p_.ID, objc.Sel("values"))
	return rv
}/* debug [instance_properties/getter]: values */


// An array of strings that specifies the export values for items in a list or a pop-up menu.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFAnnotation/values
func (p_ PDFAnnotation) SetValues(value []string) {
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
}/* debug [instance_properties/setter]: values */


// The type of button widget control, either radio button, push button, or checkbox.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFAnnotation/widgetControlType
func (p_ PDFAnnotation) WidgetControlType() PDFWidgetControlType {
	rv := objc.Send[PDFWidgetControlType](p_.ID, objc.Sel("widgetControlType"))
	return rv
}/* debug [instance_properties/getter]: widgetControlType */


// The type of button widget control, either radio button, push button, or checkbox.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFAnnotation/widgetControlType
func (p_ PDFAnnotation) SetWidgetControlType(value PDFWidgetControlType) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setWidgetControlType:"), value)
}/* debug [instance_properties/setter]: widgetControlType */


// The string value that the widget reverts to when performing a reset form action.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFAnnotation/widgetDefaultStringValue
func (p_ PDFAnnotation) WidgetDefaultStringValue() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](p_.ID, objc.Sel("widgetDefaultStringValue"))
	return rv
}/* debug [instance_properties/getter]: widgetDefaultStringValue */


// The string value that the widget reverts to when performing a reset form action.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFAnnotation/widgetDefaultStringValue
func (p_ PDFAnnotation) SetWidgetDefaultStringValue(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setWidgetDefaultStringValue:"), value)
}/* debug [instance_properties/setter]: widgetDefaultStringValue */


// The type of widget annotation, such as button, choice, or text.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFAnnotation/widgetFieldType
func (p_ PDFAnnotation) WidgetFieldType() PDFAnnotationWidgetSubtype /* typedef */ {
	rv := objc.Send[objc.ID](p_.ID, objc.Sel("widgetFieldType"))
	return rv
}/* debug [instance_properties/getter]: widgetFieldType */


// The type of widget annotation, such as button, choice, or text.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFAnnotation/widgetFieldType
func (p_ PDFAnnotation) SetWidgetFieldType(value PDFAnnotationWidgetSubtype /* typedef */) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setWidgetFieldType:"), value)
}/* debug [instance_properties/setter]: widgetFieldType */


// The string value of the widget annotation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFAnnotation/widgetStringValue
func (p_ PDFAnnotation) WidgetStringValue() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](p_.ID, objc.Sel("widgetStringValue"))
	return rv
}/* debug [instance_properties/getter]: widgetStringValue */


// The string value of the widget annotation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFAnnotation/widgetStringValue
func (p_ PDFAnnotation) SetWidgetStringValue(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setWidgetStringValue:"), value)
}/* debug [instance_properties/setter]: widgetStringValue */


// A Boolean value that indicates whether the annotation divides the text widget’s bounds into equally spaced segments, such as in a form entry field.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/pdfkit/pdfannotation/hascomb
func (p_ PDFAnnotation) HasComb() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("hasComb"))
	return rv
}/* debug [instance_properties/getter]: hasComb */


// A Boolean value that indicates whether the annotation divides the text widget’s bounds into equally spaced segments, such as in a form entry field.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/pdfkit/pdfannotation/hascomb
func (p_ PDFAnnotation) SetHasComb(value bool) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setHasComb:"), value)
}/* debug [instance_properties/setter]: hasComb */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/pdfkit/pdfannotation/isactivatabletextfield
func (p_ PDFAnnotation) IsActivatableTextField() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("isActivatableTextField"))
	return rv
}/* debug [instance_properties/getter]: isActivatableTextField */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/pdfkit/pdfannotation/isactivatabletextfield
func (p_ PDFAnnotation) SetIsActivatableTextField(value bool) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setIsActivatableTextField:"), value)
}/* debug [instance_properties/setter]: isActivatableTextField */


// A Boolean value that indicates whether the annotation is in a highlighted state, such as when the mouse is down on a link annotation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/pdfkit/pdfannotation/ishighlighted
func (p_ PDFAnnotation) IsHighlighted() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("isHighlighted"))
	return rv
}/* debug [instance_properties/getter]: isHighlighted */


// A Boolean value that indicates whether the annotation is in a highlighted state, such as when the mouse is down on a link annotation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/pdfkit/pdfannotation/ishighlighted
func (p_ PDFAnnotation) SetIsHighlighted(value bool) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setIsHighlighted:"), value)
}/* debug [instance_properties/setter]: isHighlighted */


// A Boolean value that indicates whether the choice widget annotation is a list or a pop-up menu.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/pdfkit/pdfannotation/islistchoice
func (p_ PDFAnnotation) IsListChoice() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("isListChoice"))
	return rv
}/* debug [instance_properties/getter]: isListChoice */


// A Boolean value that indicates whether the choice widget annotation is a list or a pop-up menu.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/pdfkit/pdfannotation/islistchoice
func (p_ PDFAnnotation) SetIsListChoice(value bool) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setIsListChoice:"), value)
}/* debug [instance_properties/setter]: isListChoice */


// A Boolean value that indicates whether the text widget annotation displays multiple lines.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/pdfkit/pdfannotation/ismultiline
func (p_ PDFAnnotation) IsMultiline() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("isMultiline"))
	return rv
}/* debug [instance_properties/getter]: isMultiline */


// A Boolean value that indicates whether the text widget annotation displays multiple lines.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/pdfkit/pdfannotation/ismultiline
func (p_ PDFAnnotation) SetIsMultiline(value bool) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setIsMultiline:"), value)
}/* debug [instance_properties/setter]: isMultiline */


// A Boolean value that indicates whether the pop-up annotation is in an opened state, displaying its text content, or in a closed state, displaying an icon.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/pdfkit/pdfannotation/isopen
func (p_ PDFAnnotation) IsOpen() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("isOpen"))
	return rv
}/* debug [instance_properties/getter]: isOpen */


// A Boolean value that indicates whether the pop-up annotation is in an opened state, displaying its text content, or in a closed state, displaying an icon.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/pdfkit/pdfannotation/isopen
func (p_ PDFAnnotation) SetIsOpen(value bool) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setIsOpen:"), value)
}/* debug [instance_properties/setter]: isOpen */


// A Boolean value that determines whether the widget is editable.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/pdfkit/pdfannotation/isreadonly
func (p_ PDFAnnotation) IsReadOnly() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("isReadOnly"))
	return rv
}/* debug [instance_properties/getter]: isReadOnly */


// A Boolean value that determines whether the widget is editable.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/pdfkit/pdfannotation/isreadonly
func (p_ PDFAnnotation) SetIsReadOnly(value bool) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setIsReadOnly:"), value)
}/* debug [instance_properties/setter]: isReadOnly */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class PDFAnnotation */


