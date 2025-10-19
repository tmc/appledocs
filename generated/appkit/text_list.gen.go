// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [TextList] class.
var (
	textListClass     _TextListClass
	textListClassOnce sync.Once
)

func getTextListClass() _TextListClass {
	textListClassOnce.Do(func() {
		textListClass = _TextListClass{objc.GetClass("NSTextList")}
	})
	return textListClass
}

type _TextListClass struct {
	class objc.Class
}

// An interface definition for the [TextList] class.
type ITextList interface {
	objectivec.IObject
}

// A section of text that forms a single list. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextList

type TextList struct {
	objectivec.Object
}

// TextListFrom constructs a [TextList] from an unsafe.Pointer.
//
// A section of text that forms a single list.
func TextListFrom(ptr unsafe.Pointer) TextList {
	return TextList{objectivec.Object{objc.ID(ptr)}}
}
// Alloc allocates a new instance without initialization.
func (tc _TextListClass) Alloc() TextList {
	rv := objc.Send[TextList](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (tc _TextListClass) New() TextList {
	rv := objc.Send[TextList](objc.ID(tc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (t_ TextList) Init() TextList {
	rv := objc.Send[TextList](t_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t_ TextList) Autorelease() TextList {
	rv := objc.Send[TextList](t_.ID, objc.Sel("autorelease"))
	return rv
}

// NewTextList creates a new TextList instance.
func NewTextList() TextList {
	return getTextListClass().New()
}




