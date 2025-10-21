// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [TextElement] class.
var (
	TextElementClass     _TextElementClass
	TextElementClassOnce sync.Once
)

func getTextElementClass() _TextElementClass {
	TextElementClassOnce.Do(func() {
		TextElementClass = _TextElementClass{objc.GetClass("NSTextElement")}
	})
	return TextElementClass
}

type _TextElementClass struct {
	class objc.Class
}

// An interface definition for the [TextElement] class.
type ITextElement interface {
	objectivec.IObject
}

// An abstract base class that represents the smallest units of text layout such as paragraphs or attachments.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextElement
type TextElement struct {
	objectivec.Object
}

// TextElementFrom constructs a [TextElement] from an unsafe.Pointer.
//
// An abstract base class that represents the smallest units of text layout such as paragraphs or attachments.
func TextElementFrom(ptr unsafe.Pointer) TextElement {
	return TextElement{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (tc _TextElementClass) Alloc() TextElement {
	rv := objc.Send[TextElement](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (tc _TextElementClass) New() TextElement {
	rv := objc.Send[TextElement](objc.ID(tc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (t_ TextElement) Init() TextElement {
	rv := objc.Send[TextElement](t_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t_ TextElement) Autorelease() TextElement {
	rv := objc.Send[TextElement](t_.ID, objc.Sel("autorelease"))
	return rv
}

// NewTextElement creates a new TextElement instance.
func NewTextElement() TextElement {
	return getTextElementClass().New()
}




