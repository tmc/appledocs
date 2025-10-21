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
}

// A block of text laid out in a subregion of the text container.
//
// A text block appears as an attribute of a paragraph, and as part of the paragraph style. The most important subclass of is , which represents a block of text that appears as a cell in a table. The table itself is a object. All objects reference this table, which controls their sizing and positioning.
//
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


// The background color of the text block.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextblock/backgroundcolor
func (t_ TextBlock) BackgroundColor() NSColor {
	rv := objc.Send[NSColor](t_.ID, objc.Sel("backgroundColor"))
	return rv
}


// SetBackgroundColor sets the value of the backgroundColor property.
// The background color of the text block.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextblock/backgroundcolor
func (t_ TextBlock) SetBackgroundColor(value IColor) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setBackgroundColor:"), value)
}

// The width of the text block.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextblock/contentwidth
func (t_ TextBlock) ContentWidth() float64 {
	rv := objc.Send[float64](t_.ID, objc.Sel("contentWidth"))
	return rv
}


// SetContentWidth sets the value of the contentWidth property.
// The width of the text block.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextblock/contentwidth
func (t_ TextBlock) SetContentWidth(value float64) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setContentWidth:"), value)
}

// The type of value stored for the text block width.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextblock/contentwidthvaluetype
func (t_ TextBlock) ContentWidthValueType() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](t_.ID, objc.Sel("contentWidthValueType"))
	return rv
}


// SetContentWidthValueType sets the value of the contentWidthValueType property.
// The type of value stored for the text block width.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextblock/contentwidthvaluetype
func (t_ TextBlock) SetContentWidthValueType(value unsafe.Pointer) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setContentWidthValueType:"), value)
}

// The vertical alignment of the text block.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextblock/verticalalignment-swift.property
func (t_ TextBlock) VerticalAlignment() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](t_.ID, objc.Sel("verticalAlignment"))
	return rv
}


// SetVerticalAlignment sets the value of the verticalAlignment property.
// The vertical alignment of the text block.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextblock/verticalalignment-swift.property
func (t_ TextBlock) SetVerticalAlignment(value unsafe.Pointer) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setVerticalAlignment:"), value)
}



