// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [TextListElement] class.
var (
	textListElementClass     _TextListElementClass
	textListElementClassOnce sync.Once
)

func getTextListElementClass() _TextListElementClass {
	textListElementClassOnce.Do(func() {
		textListElementClass = _TextListElementClass{objc.GetClass("NSTextListElement")}
	})
	return textListElementClass
}

type _TextListElementClass struct {
	class objc.Class
}

// An interface definition for the [TextListElement] class.
type ITextListElement interface {
	ITextParagraph
}

// A class that represents a text list node. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextListElement
type TextListElement struct {
	TextParagraph
}

// TextListElementFrom constructs a [TextListElement] from an unsafe.Pointer.
//
// A class that represents a text list node.
func TextListElementFrom(ptr unsafe.Pointer) TextListElement {
	return TextListElement{
		TextParagraph: TextParagraphFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (tc _TextListElementClass) Alloc() TextListElement {
	rv := objc.Send[TextListElement](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (tc _TextListElementClass) New() TextListElement {
	rv := objc.Send[TextListElement](objc.ID(tc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (t_ TextListElement) Init() TextListElement {
	rv := objc.Send[TextListElement](t_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t_ TextListElement) Autorelease() TextListElement {
	rv := objc.Send[TextListElement](t_.ID, objc.Sel("autorelease"))
	return rv
}

// NewTextListElement creates a new TextListElement instance.
func NewTextListElement() TextListElement {
	return getTextListElementClass().New()
}




