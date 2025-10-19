// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [TextTab] class.
var (
	textTabClass     _TextTabClass
	textTabClassOnce sync.Once
)

func getTextTabClass() _TextTabClass {
	textTabClassOnce.Do(func() {
		textTabClass = _TextTabClass{objc.GetClass("NSTextTab")}
	})
	return textTabClass
}

type _TextTabClass struct {
	class objc.Class
}

// An interface definition for the [TextTab] class.
type ITextTab interface {
	objectivec.IObject
}

// A tab in a paragraph. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextTab

type TextTab struct {
	objectivec.Object
}

// TextTabFrom constructs a [TextTab] from an unsafe.Pointer.
//
// A tab in a paragraph.
func TextTabFrom(ptr unsafe.Pointer) TextTab {
	return TextTab{objectivec.Object{objc.ID(ptr)}}
}
// Alloc allocates a new instance without initialization.
func (tc _TextTabClass) Alloc() TextTab {
	rv := objc.Send[TextTab](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (tc _TextTabClass) New() TextTab {
	rv := objc.Send[TextTab](objc.ID(tc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (t_ TextTab) Init() TextTab {
	rv := objc.Send[TextTab](t_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t_ TextTab) Autorelease() TextTab {
	rv := objc.Send[TextTab](t_.ID, objc.Sel("autorelease"))
	return rv
}

// NewTextTab creates a new TextTab instance.
func NewTextTab() TextTab {
	return getTextTabClass().New()
}




