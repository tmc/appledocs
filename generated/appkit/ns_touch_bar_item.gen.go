// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [TouchBarItem] class.
var (
	TouchBarItemClass     _TouchBarItemClass
	TouchBarItemClassOnce sync.Once
)

func getTouchBarItemClass() _TouchBarItemClass {
	TouchBarItemClassOnce.Do(func() {
		TouchBarItemClass = _TouchBarItemClass{objc.GetClass("NSTouchBarItem")}
	})
	return TouchBarItemClass
}

type _TouchBarItemClass struct {
	class objc.Class
}

// An interface definition for the [TouchBarItem] class.
type ITouchBarItem interface {
	objectivec.IObject
}

// A parent class referenced by other AppKit classes.


// A parent class referenced by other AppKit classes. [Full Topic]
type TouchBarItem struct {
	objectivec.Object
}

// TouchBarItemFrom constructs a [TouchBarItem] from an unsafe.Pointer.
//
// A parent class referenced by other AppKit classes.
func TouchBarItemFrom(ptr unsafe.Pointer) TouchBarItem {
	return TouchBarItem{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (tc _TouchBarItemClass) Alloc() TouchBarItem {
	rv := objc.Send[TouchBarItem](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (tc _TouchBarItemClass) New() TouchBarItem {
	rv := objc.Send[TouchBarItem](objc.ID(tc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (t_ TouchBarItem) Init() TouchBarItem {
	rv := objc.Send[TouchBarItem](t_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t_ TouchBarItem) Autorelease() TouchBarItem {
	rv := objc.Send[TouchBarItem](t_.ID, objc.Sel("autorelease"))
	return rv
}

// NewTouchBarItem creates a new TouchBarItem instance.
func NewTouchBarItem() TouchBarItem {
	return getTouchBarItemClass().New()
}




