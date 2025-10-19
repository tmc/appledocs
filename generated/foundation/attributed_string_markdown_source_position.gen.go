// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [AttributedStringMarkdownSourcePosition] class.
var (
	attributedStringMarkdownSourcePositionClass     _AttributedStringMarkdownSourcePositionClass
	attributedStringMarkdownSourcePositionClassOnce sync.Once
)

func getAttributedStringMarkdownSourcePositionClass() _AttributedStringMarkdownSourcePositionClass {
	attributedStringMarkdownSourcePositionClassOnce.Do(func() {
		attributedStringMarkdownSourcePositionClass = _AttributedStringMarkdownSourcePositionClass{objc.GetClass("NSAttributedStringMarkdownSourcePosition")}
	})
	return attributedStringMarkdownSourcePositionClass
}

type _AttributedStringMarkdownSourcePositionClass struct {
	class objc.Class
}

// An interface definition for the [AttributedStringMarkdownSourcePosition] class.
type IAttributedStringMarkdownSourcePosition interface {
	objectivec.IObject
}

// The position of attributed string text in its original Markdown source string. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSAttributedStringMarkdownSourcePosition
type AttributedStringMarkdownSourcePosition struct {
	objectivec.Object
}

// AttributedStringMarkdownSourcePositionFrom constructs a [AttributedStringMarkdownSourcePosition] from an unsafe.Pointer.
//
// The position of attributed string text in its original Markdown source string.
func AttributedStringMarkdownSourcePositionFrom(ptr unsafe.Pointer) AttributedStringMarkdownSourcePosition {
	return AttributedStringMarkdownSourcePosition{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (ac _AttributedStringMarkdownSourcePositionClass) Alloc() AttributedStringMarkdownSourcePosition {
	rv := objc.Send[AttributedStringMarkdownSourcePosition](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ac _AttributedStringMarkdownSourcePositionClass) New() AttributedStringMarkdownSourcePosition {
	rv := objc.Send[AttributedStringMarkdownSourcePosition](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ AttributedStringMarkdownSourcePosition) Init() AttributedStringMarkdownSourcePosition {
	rv := objc.Send[AttributedStringMarkdownSourcePosition](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ AttributedStringMarkdownSourcePosition) Autorelease() AttributedStringMarkdownSourcePosition {
	rv := objc.Send[AttributedStringMarkdownSourcePosition](a_.ID, objc.Sel("autorelease"))
	return rv
}

// NewAttributedStringMarkdownSourcePosition creates a new AttributedStringMarkdownSourcePosition instance.
func NewAttributedStringMarkdownSourcePosition() AttributedStringMarkdownSourcePosition {
	return getAttributedStringMarkdownSourcePositionClass().New()
}




