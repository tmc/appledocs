// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

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

// An interface definition for the [TextBlock] class.
type ITextBlock interface {
	objectivec.IObject
	// properties:
	BackgroundColor() IColor
	SetBackgroundColor(value IColor)
	ContentWidth() float64 /* primitive/slice/pointer. */
	SetContentWidth(value float64 /* primitive/slice/pointer. */)
	ContentWidthValueType() unsafe.Pointer
	SetContentWidthValueType(value unsafe.Pointer)
	VerticalAlignment() unsafe.Pointer
	SetVerticalAlignment(value unsafe.Pointer)
	// methods:
	BorderColorForEdge(edge RectEdge /* not a class type */) IColor
	RectForLayoutAtPointInRectTextContainerCharacterRange(startingPoint objc.IObject /* cross-framework Point */, rect objc.IObject /* cross-framework Rect */, textContainer ITextContainer, charRange objc.IObject /* cross-framework Range */) objc.IObject /* cross-framework: Rect */
	SetValueTypeForDimension(val float64 /* primitive/slice/pointer. */, type_ TextBlockValueType /* not a class type */, dimension TextBlockDimension)
	ValueForDimension(dimension TextBlockDimension) float64 /* primitive/slice/pointer. */
	ValueTypeForDimension(dimension TextBlockDimension) TextBlockValueType /* not a class type */
}

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

// Alloc allocates a new instance without initialization.
func (tc _TextBlockClass) Alloc() TextBlock {
	rv := objc.Send[TextBlock](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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



// Returns the border color of the specified text block edge.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextBlock/borderColor(for:)
func (t_ TextBlock) BorderColorForEdge(edge RectEdge /* not a class type */) IColor {
	rv := objc.Send[Color](t_.ID, objc.Sel("borderColorForEdge:"), edge)
	return rv
}


// Returns the rectangle within which glyphs should be laid out for the specified arguments.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextBlock/rectForLayout(at:in:textContainer:characterRange:)
func (t_ TextBlock) RectForLayoutAtPointInRectTextContainerCharacterRange(startingPoint objc.IObject /* cross-framework Point */, rect objc.IObject /* cross-framework Rect */, textContainer ITextContainer, charRange objc.IObject /* cross-framework Range */) objc.IObject /* cross-framework: Rect */ {
	rv := objc.Send[Rect](t_.ID, objc.Sel("rectForLayoutAtPoint:inRect:textContainer:characterRange:"), startingPoint, rect, textContainer, charRange)
	return rv
}


// Sets a dimension of the text block.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextBlock/setValue(_:type:for:)
func (t_ TextBlock) SetValueTypeForDimension(val float64 /* primitive/slice/pointer. */, type_ TextBlockValueType /* not a class type */, dimension TextBlockDimension) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setValue:type:forDimension:"), val, type_, dimension)
}


// Returns the value of the specified text block dimension.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextBlock/value(for:)
func (t_ TextBlock) ValueForDimension(dimension TextBlockDimension) float64 /* primitive/slice/pointer. */ {
	rv := objc.Send[float64](t_.ID, objc.Sel("valueForDimension:"), dimension)
	return rv
}


// Returns the value type of the specified text block dimension.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextBlock/valueType(for:)
func (t_ TextBlock) ValueTypeForDimension(dimension TextBlockDimension) TextBlockValueType /* not a class type */ {
	rv := objc.Send[TextBlockValueType](t_.ID, objc.Sel("valueTypeForDimension:"), dimension)
	return rv
}


// The background color of the text block.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextblock/backgroundcolor
func (t_ TextBlock) BackgroundColor() IColor {
	rv := objc.Send[Color](t_.ID, objc.Sel("backgroundColor"))
	return rv
}


// The background color of the text block.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextblock/backgroundcolor
func (t_ TextBlock) SetBackgroundColor(value IColor) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setBackgroundColor:"), value)
}


// The width of the text block.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextblock/contentwidth
func (t_ TextBlock) ContentWidth() float64 /* primitive/slice/pointer. */ {
	rv := objc.Send[float64](t_.ID, objc.Sel("contentWidth"))
	return rv
}


// The width of the text block.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextblock/contentwidth
func (t_ TextBlock) SetContentWidth(value float64 /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setContentWidth:"), value)
}


// The type of value stored for the text block width.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextblock/contentwidthvaluetype
func (t_ TextBlock) ContentWidthValueType() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](t_.ID, objc.Sel("contentWidthValueType"))
	return rv
}


// The type of value stored for the text block width.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextblock/contentwidthvaluetype
func (t_ TextBlock) SetContentWidthValueType(value unsafe.Pointer) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setContentWidthValueType:"), value)
}


// The vertical alignment of the text block.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextblock/verticalalignment-swift.property
func (t_ TextBlock) VerticalAlignment() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](t_.ID, objc.Sel("verticalAlignment"))
	return rv
}


// The vertical alignment of the text block.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextblock/verticalalignment-swift.property
func (t_ TextBlock) SetVerticalAlignment(value unsafe.Pointer) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setVerticalAlignment:"), value)
}



