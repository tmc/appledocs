// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/corefoundation"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/vision"
)

/* debug [class.gen.go]: Generating class NSTextBlock */


/* debug [class_header]: Header for NSTextBlock */
// The class instance for the [TextBlock] class.
var (
	TextBlockClass     _TextBlockClass
	TextBlockClassOnce sync.Once
)

func getTextBlockClass() _TextBlockClass {
	TextBlockClassOnce.Do(func() {
		TextBlockClass = _TextBlockClass{objc.GetClass("NSTextBlock")}
	})
	return TextBlockClass
}

type _TextBlockClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for TextBlock */
// An interface definition for the [TextBlock] class.
type ITextBlock interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for TextBlock */
	// properties:
	BackgroundColor() IColor
	SetBackgroundColor(value IColor)
	ContentWidth() float64
	ContentWidthValueType() TextBlockValueType
	VerticalAlignment() TextBlockVerticalAlignment
	SetVerticalAlignment(value TextBlockVerticalAlignment)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for TextBlock */
	// methods:
	BorderColorForEdge(edge RectEdge /* not a class type */) IColor
	BoundsRectForContentRectInRectTextContainerCharacterRange(contentRect Rect /* not a class type */, rect Rect /* not a class type */, textContainer ITextContainer, charRange corefoundation.Range) Rect /* not a class type */
	DrawBackgroundWithFrameInViewCharacterRangeLayoutManager(frameRect Rect /* not a class type */, controlView IView, charRange corefoundation.Range, layoutManager ILayoutManager)
	RectForLayoutAtPointInRectTextContainerCharacterRange(startingPoint vision.Point, rect Rect /* not a class type */, textContainer ITextContainer, charRange corefoundation.Range) Rect /* not a class type */
	SetBorderColor(color IColor)
	SetBorderColorForEdge(color IColor, edge RectEdge /* not a class type */)
	SetContentWidthType(val float64, type_ TextBlockValueType)
	SetValueTypeForDimension(val float64, type_ TextBlockValueType, dimension TextBlockDimension)
	SetWidthTypeForLayer(val float64, type_ TextBlockValueType, layer TextBlockLayer)
	SetWidthTypeForLayerEdge(val float64, type_ TextBlockValueType, layer TextBlockLayer, edge RectEdge /* not a class type */)
	ValueForDimension(dimension TextBlockDimension) float64
	ValueTypeForDimension(dimension TextBlockDimension) TextBlockValueType
	WidthForLayerEdge(layer TextBlockLayer, edge RectEdge /* not a class type */) float64
	WidthValueTypeForLayerEdge(layer TextBlockLayer, edge RectEdge /* not a class type */) TextBlockValueType
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for TextBlock */
// Alloc allocates a new instance without initialization.
func (tc _TextBlockClass) Alloc() TextBlock {
	rv := objc.Send[TextBlock](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (tc _TextBlockClass) New() TextBlock {
	rv := objc.Send[TextBlock](objc.ID(tc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (t_ TextBlock) Init() TextBlock {
	rv := objc.Send[TextBlock](t_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t_ TextBlock) Autorelease() TextBlock {
	rv := objc.Send[TextBlock](t_.ID, objc.Sel("autorelease"))
	return rv
}

// NewTextBlock creates a new TextBlock instance.
func NewTextBlock() TextBlock {
	return getTextBlockClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for TextBlock */
// A block of text laid out in a subregion of the text container.
//
// A text block appears as an attribute of a paragraph, and as part of the paragraph style. The most important subclass of is , which represents a block of text that appears as a cell in a table. The table itself is a object. All objects reference this table, which controls their sizing and positioning.


// A block of text laid out in a subregion of the text container.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextBlock
type TextBlock struct {
	objectivec.Object
}

// TextBlockFrom constructs a [TextBlock] from an unsafe.Pointer.
//
// A block of text laid out in a subregion of the text container.
func TextBlockFrom(ptr unsafe.Pointer) TextBlock {
	return TextBlock{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for TextBlock */
/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for TextBlock */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for TextBlock */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for TextBlock */

// Returns the border color of the specified text block edge.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextBlock/borderColor(for:)
func (t_ TextBlock) BorderColorForEdge(edge RectEdge /* not a class type */) IColor {
	rv := objc.Send[Color](t_.ID, objc.Sel("borderColorForEdge:"), edge)
	return rv
}/* debug [instance_methods/method]: BorderColorForEdge */


// Returns the rectangle the text in the block actually occupies, including padding, borders, and margins.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextBlock/boundsRect(forContentRect:in:textContainer:characterRange:)
func (t_ TextBlock) BoundsRectForContentRectInRectTextContainerCharacterRange(contentRect Rect /* not a class type */, rect Rect /* not a class type */, textContainer ITextContainer, charRange corefoundation.Range) Rect /* not a class type */ {
	rv := objc.Send[Rect](t_.ID, objc.Sel("boundsRectForContentRect:inRect:textContainer:characterRange:"), contentRect, rect, textContainer, charRange)
	return rv
}/* debug [instance_methods/method]: BoundsRectForContentRectInRectTextContainerCharacterRange */


// Called by the layout manager to draw any colors and other decorations before the text is drawn.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextBlock/drawBackground(withFrame:in:characterRange:layoutManager:)
func (t_ TextBlock) DrawBackgroundWithFrameInViewCharacterRangeLayoutManager(frameRect Rect /* not a class type */, controlView IView, charRange corefoundation.Range, layoutManager ILayoutManager) {
	objc.Send[objc.ID](t_.ID, objc.Sel("drawBackgroundWithFrame:inView:characterRange:layoutManager:"), frameRect, controlView, charRange, layoutManager)
}/* debug [instance_methods/method]: DrawBackgroundWithFrameInViewCharacterRangeLayoutManager */


// Returns the rectangle within which glyphs should be laid out for the specified arguments.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextBlock/rectForLayout(at:in:textContainer:characterRange:)
func (t_ TextBlock) RectForLayoutAtPointInRectTextContainerCharacterRange(startingPoint vision.Point, rect Rect /* not a class type */, textContainer ITextContainer, charRange corefoundation.Range) Rect /* not a class type */ {
	rv := objc.Send[Rect](t_.ID, objc.Sel("rectForLayoutAtPoint:inRect:textContainer:characterRange:"), startingPoint, rect, textContainer, charRange)
	return rv
}/* debug [instance_methods/method]: RectForLayoutAtPointInRectTextContainerCharacterRange */


// Sets the color of all borders of the text block.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextBlock/setBorderColor(_:)
func (t_ TextBlock) SetBorderColor(color IColor) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setBorderColor:"), color)
}/* debug [instance_methods/method]: SetBorderColor */


// Sets the border color of the specified edge of the text block.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextBlock/setBorderColor(_:for:)
func (t_ TextBlock) SetBorderColorForEdge(color IColor, edge RectEdge /* not a class type */) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setBorderColor:forEdge:"), color, edge)
}/* debug [instance_methods/method]: SetBorderColorForEdge */


// Sets the width of the text block.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextBlock/setContentWidth(_:type:)
func (t_ TextBlock) SetContentWidthType(val float64, type_ TextBlockValueType) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setContentWidth:type:"), val, type_)
}/* debug [instance_methods/method]: SetContentWidthType */


// Sets a dimension of the text block.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextBlock/setValue(_:type:for:)
func (t_ TextBlock) SetValueTypeForDimension(val float64, type_ TextBlockValueType, dimension TextBlockDimension) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setValue:type:forDimension:"), val, type_, dimension)
}/* debug [instance_methods/method]: SetValueTypeForDimension */


// Sets the width of all edges of a specified layer of the text block.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextBlock/setWidth(_:type:for:)
func (t_ TextBlock) SetWidthTypeForLayer(val float64, type_ TextBlockValueType, layer TextBlockLayer) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setWidth:type:forLayer:"), val, type_, layer)
}/* debug [instance_methods/method]: SetWidthTypeForLayer */


// Sets the width of a specified edge of a specified layer of the text block.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextBlock/setWidth(_:type:for:edge:)
func (t_ TextBlock) SetWidthTypeForLayerEdge(val float64, type_ TextBlockValueType, layer TextBlockLayer, edge RectEdge /* not a class type */) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setWidth:type:forLayer:edge:"), val, type_, layer, edge)
}/* debug [instance_methods/method]: SetWidthTypeForLayerEdge */


// Returns the value of the specified text block dimension.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextBlock/value(for:)
func (t_ TextBlock) ValueForDimension(dimension TextBlockDimension) float64 {
	rv := objc.Send[float64](t_.ID, objc.Sel("valueForDimension:"), dimension)
	return rv
}/* debug [instance_methods/method]: ValueForDimension */


// Returns the value type of the specified text block dimension.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextBlock/valueType(for:)
func (t_ TextBlock) ValueTypeForDimension(dimension TextBlockDimension) TextBlockValueType {
	rv := objc.Send[TextBlockValueType](t_.ID, objc.Sel("valueTypeForDimension:"), dimension)
	return rv
}/* debug [instance_methods/method]: ValueTypeForDimension */


// Returns the width of an edge of a specified layer of the text block.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextBlock/width(for:edge:)
func (t_ TextBlock) WidthForLayerEdge(layer TextBlockLayer, edge RectEdge /* not a class type */) float64 {
	rv := objc.Send[float64](t_.ID, objc.Sel("widthForLayer:edge:"), layer, edge)
	return rv
}/* debug [instance_methods/method]: WidthForLayerEdge */


// Returns the value type of an edge of a specified layer of the text block.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextBlock/widthValueType(for:edge:)
func (t_ TextBlock) WidthValueTypeForLayerEdge(layer TextBlockLayer, edge RectEdge /* not a class type */) TextBlockValueType {
	rv := objc.Send[TextBlockValueType](t_.ID, objc.Sel("widthValueTypeForLayer:edge:"), layer, edge)
	return rv
}/* debug [instance_methods/method]: WidthValueTypeForLayerEdge */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for TextBlock */

// The background color of the text block.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextBlock/backgroundColor
func (t_ TextBlock) BackgroundColor() IColor {
	rv := objc.Send[Color](t_.ID, objc.Sel("backgroundColor"))
	return rv
}/* debug [instance_properties/getter]: backgroundColor */


// The background color of the text block.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextBlock/backgroundColor
func (t_ TextBlock) SetBackgroundColor(value IColor) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setBackgroundColor:"), value)
}/* debug [instance_properties/setter]: backgroundColor */


// The width of the text block.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextBlock/contentWidth
func (t_ TextBlock) ContentWidth() float64 {
	rv := objc.Send[float64](t_.ID, objc.Sel("contentWidth"))
	return rv
}/* debug [instance_properties/getter]: contentWidth */


// The type of value stored for the text block width.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextBlock/contentWidthValueType
func (t_ TextBlock) ContentWidthValueType() TextBlockValueType {
	rv := objc.Send[TextBlockValueType](t_.ID, objc.Sel("contentWidthValueType"))
	return rv
}/* debug [instance_properties/getter]: contentWidthValueType */


// The vertical alignment of the text block.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextBlock/verticalAlignment-swift.property
func (t_ TextBlock) VerticalAlignment() TextBlockVerticalAlignment {
	rv := objc.Send[TextBlockVerticalAlignment](t_.ID, objc.Sel("verticalAlignment"))
	return rv
}/* debug [instance_properties/getter]: verticalAlignment */


// The vertical alignment of the text block.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextBlock/verticalAlignment-swift.property
func (t_ TextBlock) SetVerticalAlignment(value TextBlockVerticalAlignment) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setVerticalAlignment:"), value)
}/* debug [instance_properties/setter]: verticalAlignment */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class NSTextBlock */


