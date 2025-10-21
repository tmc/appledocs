// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [TextParagraph] class.
var (
	TextParagraphClass     _TextParagraphClass
	TextParagraphClassOnce sync.Once
)

func getTextParagraphClass() _TextParagraphClass {
	TextParagraphClassOnce.Do(func() {
		TextParagraphClass = _TextParagraphClass{objc.GetClass("NSTextParagraph")}
	})
	return TextParagraphClass
}

type _TextParagraphClass struct {
	class objc.Class
}

// An interface definition for the [TextParagraph] class.
type ITextParagraph interface {
	ITextElement
}

// A class that represents a single paragraph backed by an attributed string as the contents.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextParagraph
type TextParagraph struct {
	TextElement
}

// TextParagraphFrom constructs a [TextParagraph] from an unsafe.Pointer.
//
// A class that represents a single paragraph backed by an attributed string as the contents.
func TextParagraphFrom(ptr unsafe.Pointer) TextParagraph {
	return TextParagraph{
		TextElement: TextElementFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (tc _TextParagraphClass) Alloc() TextParagraph {
	rv := objc.Send[TextParagraph](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (tc _TextParagraphClass) New() TextParagraph {
	rv := objc.Send[TextParagraph](objc.ID(tc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (t_ TextParagraph) Init() TextParagraph {
	rv := objc.Send[TextParagraph](t_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t_ TextParagraph) Autorelease() TextParagraph {
	rv := objc.Send[TextParagraph](t_.ID, objc.Sel("autorelease"))
	return rv
}

// NewTextParagraph creates a new TextParagraph instance.
func NewTextParagraph() TextParagraph {
	return getTextParagraphClass().New()
}




