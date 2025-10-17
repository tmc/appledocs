// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [Scroller] class.
var scrollerClass = _ScrollerClass{objc.GetClass("NSScroller")}

type _ScrollerClass struct {
	class objc.Class
}

// An interface definition for the [Scroller] class.
type IScroller interface {
	IControl
}

// An object that controls scrolling of a document view within a scroll view or other type of container view. [Full Topic]
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



