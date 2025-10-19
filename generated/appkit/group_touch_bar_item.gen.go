// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [GroupTouchBarItem] class.
var (
	groupTouchBarItemClass     _GroupTouchBarItemClass
	groupTouchBarItemClassOnce sync.Once
)

func getGroupTouchBarItemClass() _GroupTouchBarItemClass {
	groupTouchBarItemClassOnce.Do(func() {
		groupTouchBarItemClass = _GroupTouchBarItemClass{objc.GetClass("NSGroupTouchBarItem")}
	})
	return groupTouchBarItemClass
}

type _GroupTouchBarItemClass struct {
	class objc.Class
}

// An interface definition for the [GroupTouchBarItem] class.
type IGroupTouchBarItem interface {
	ITouchBarItem
}

// A bar item that provides a bar to contain other items. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSGroupTouchBarItem

type GroupTouchBarItem struct {
	TouchBarItem
}

// GroupTouchBarItemFrom constructs a [GroupTouchBarItem] from an unsafe.Pointer.
//
// A bar item that provides a bar to contain other items.
func GroupTouchBarItemFrom(ptr unsafe.Pointer) GroupTouchBarItem {
	return GroupTouchBarItem{
		TouchBarItem: TouchBarItemFrom(ptr),
	}
}
// Alloc allocates a new instance without initialization.
func (gc _GroupTouchBarItemClass) Alloc() GroupTouchBarItem {
	rv := objc.Send[GroupTouchBarItem](objc.ID(gc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new instance with a +1 retain count.
func (gc _GroupTouchBarItemClass) New() GroupTouchBarItem {
	rv := objc.Send[GroupTouchBarItem](objc.ID(gc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (g_ GroupTouchBarItem) Init() GroupTouchBarItem {
	rv := objc.Send[GroupTouchBarItem](g_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (g_ GroupTouchBarItem) Autorelease() GroupTouchBarItem {
	rv := objc.Send[GroupTouchBarItem](g_.ID, objc.Sel("autorelease"))
	return rv
}

// NewGroupTouchBarItem creates a new GroupTouchBarItem instance.
func NewGroupTouchBarItem() GroupTouchBarItem {
	return getGroupTouchBarItemClass().New()
}




