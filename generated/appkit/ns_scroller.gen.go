// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [Scroller] class.
var (
	ScrollerClass     _ScrollerClass
	ScrollerClassOnce sync.Once
)

func getScrollerClass() _ScrollerClass {
	ScrollerClassOnce.Do(func() {
		ScrollerClass = _ScrollerClass{objc.GetClass("NSScroller")}
	})
	return ScrollerClass
}

type _ScrollerClass struct {
	class objc.Class
}

// An interface definition for the [Scroller] class.
type IScroller interface {
	IControl
}

// An object that controls scrolling of a document view within a scroll view or other type of container view.
//
// A scroller displays a slot containing a knob that the user can drag directly to the desired location. The knob indicates both the position within the document view and—by varying in size within the slot—the amount visible relative to the size of the document view. Typically, you don’t need to program with scrollers; instead, you configure them with an object in a . Don’t use an scroller when a slider would be more appropriate. An object represents a range of values for something in the application and lets the user choose a setting. A scroller represents the relative position of the visible portion of a view and lets the user choose which portion to view.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScroller
type Scroller struct {
	Control
}

// ScrollerFrom constructs a [Scroller] from an unsafe.Pointer.
//
// An object that controls scrolling of a document view within a scroll view or other type of container view.
func ScrollerFrom(ptr unsafe.Pointer) Scroller {
	return Scroller{
		Control: ControlFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (sc _ScrollerClass) Alloc() Scroller {
	rv := objc.Send[Scroller](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (sc _ScrollerClass) New() Scroller {
	rv := objc.Send[Scroller](objc.ID(sc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (s_ Scroller) Init() Scroller {
	rv := objc.Send[Scroller](s_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s_ Scroller) Autorelease() Scroller {
	rv := objc.Send[Scroller](s_.ID, objc.Sel("autorelease"))
	return rv
}

// NewScroller creates a new Scroller instance.
func NewScroller() Scroller {
	return getScrollerClass().New()
}




