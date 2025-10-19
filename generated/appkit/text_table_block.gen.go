// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [TextTableBlock] class.
var textTableBlockClass = _TextTableBlockClass{objc.GetClass("NSTextTableBlock")}

type _TextTableBlockClass struct {
	class objc.Class
}

// An interface definition for the [TextTableBlock] class.
type ITextTableBlock interface {
	ITextBlock
}

// A text block that appears as a cell in a text table. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextTableBlock

type TextTableBlock struct {
	TextBlock
}

// TextTableBlockFrom constructs a [TextTableBlock] from an unsafe.Pointer.
//
// A text block that appears as a cell in a text table.
func TextTableBlockFrom(ptr unsafe.Pointer) TextTableBlock {
	return TextTableBlock{
		TextBlock: TextBlockFrom(ptr),
	}
}
// Alloc allocates a new instance without initialization.
func (tc _TextTableBlockClass) Alloc() TextTableBlock {
	rv := objc.Send[TextTableBlock](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new instance with a +1 retain count.
func (tc _TextTableBlockClass) New() TextTableBlock {
	rv := objc.Send[TextTableBlock](objc.ID(tc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (t_ TextTableBlock) Init() TextTableBlock {
	rv := objc.Send[TextTableBlock](t_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t_ TextTableBlock) Autorelease() TextTableBlock {
	rv := objc.Send[TextTableBlock](t_.ID, objc.Sel("autorelease"))
	return rv
}

// NewTextTableBlock creates a new TextTableBlock instance.
func NewTextTableBlock() TextTableBlock {
	return textTableBlockClass.New()
}




