// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [WindowTab] class.
var (
	windowTabClass     _WindowTabClass
	windowTabClassOnce sync.Once
)

func getWindowTabClass() _WindowTabClass {
	windowTabClassOnce.Do(func() {
		windowTabClass = _WindowTabClass{objc.GetClass("NSWindowTab")}
	})
	return windowTabClass
}

type _WindowTabClass struct {
	class objc.Class
}

// An interface definition for the [WindowTab] class.
type IWindowTab interface {
	objectivec.IObject
}

// A tab associated with a window that is part of a tabbing group. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindowTab
type WindowTab struct {
	objectivec.Object
}

// WindowTabFrom constructs a [WindowTab] from an unsafe.Pointer.
//
// A tab associated with a window that is part of a tabbing group.
func WindowTabFrom(ptr unsafe.Pointer) WindowTab {
	return WindowTab{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (wc _WindowTabClass) Alloc() WindowTab {
	rv := objc.Send[WindowTab](objc.ID(wc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (wc _WindowTabClass) New() WindowTab {
	rv := objc.Send[WindowTab](objc.ID(wc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (w_ WindowTab) Init() WindowTab {
	rv := objc.Send[WindowTab](w_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (w_ WindowTab) Autorelease() WindowTab {
	rv := objc.Send[WindowTab](w_.ID, objc.Sel("autorelease"))
	return rv
}

// NewWindowTab creates a new WindowTab instance.
func NewWindowTab() WindowTab {
	return getWindowTabClass().New()
}




