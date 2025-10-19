// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [TextLayoutManager] class.
var (
	textLayoutManagerClass     _TextLayoutManagerClass
	textLayoutManagerClassOnce sync.Once
)

func getTextLayoutManagerClass() _TextLayoutManagerClass {
	textLayoutManagerClassOnce.Do(func() {
		textLayoutManagerClass = _TextLayoutManagerClass{objc.GetClass("NSTextLayoutManager")}
	})
	return textLayoutManagerClass
}

type _TextLayoutManagerClass struct {
	class objc.Class
}

// An interface definition for the [TextLayoutManager] class.
type ITextLayoutManager interface {
	objectivec.IObject
}

// The primary class that you use to manage text layout and presentation for custom text displays. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextLayoutManager
type TextLayoutManager struct {
	objectivec.Object
}

// TextLayoutManagerFrom constructs a [TextLayoutManager] from an unsafe.Pointer.
//
// The primary class that you use to manage text layout and presentation for custom text displays.
func TextLayoutManagerFrom(ptr unsafe.Pointer) TextLayoutManager {
	return TextLayoutManager{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (tc _TextLayoutManagerClass) Alloc() TextLayoutManager {
	rv := objc.Send[TextLayoutManager](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (tc _TextLayoutManagerClass) New() TextLayoutManager {
	rv := objc.Send[TextLayoutManager](objc.ID(tc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (t_ TextLayoutManager) Init() TextLayoutManager {
	rv := objc.Send[TextLayoutManager](t_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t_ TextLayoutManager) Autorelease() TextLayoutManager {
	rv := objc.Send[TextLayoutManager](t_.ID, objc.Sel("autorelease"))
	return rv
}

// NewTextLayoutManager creates a new TextLayoutManager instance.
func NewTextLayoutManager() TextLayoutManager {
	return getTextLayoutManagerClass().New()
}




