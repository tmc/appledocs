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
	textBlockClass     _TextBlockClass
	textBlockClassOnce sync.Once
)

func getTextBlockClass() _TextBlockClass {
	textBlockClassOnce.Do(func() {
		textBlockClass = _TextBlockClass{objc.GetClass("NSTextBlock")}
	})
	return textBlockClass
}

type _TextBlockClass struct {
	class objc.Class
}

// An interface definition for the [TextBlock] class.
type ITextBlock interface {
	objectivec.IObject
}

// A block of text laid out in a subregion of the text container. [Full Topic]
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




