// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [DraggingItem] class.
var (
	draggingItemClass     _DraggingItemClass
	draggingItemClassOnce sync.Once
)

func getDraggingItemClass() _DraggingItemClass {
	draggingItemClassOnce.Do(func() {
		draggingItemClass = _DraggingItemClass{objc.GetClass("NSDraggingItem")}
	})
	return draggingItemClass
}

type _DraggingItemClass struct {
	class objc.Class
}

// An interface definition for the [DraggingItem] class.
type IDraggingItem interface {
	objectivec.IObject
}

// A single dragged item within a dragging session. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDraggingItem
type DraggingItem struct {
	objectivec.Object
}

// DraggingItemFrom constructs a [DraggingItem] from an unsafe.Pointer.
//
// A single dragged item within a dragging session.
func DraggingItemFrom(ptr unsafe.Pointer) DraggingItem {
	return DraggingItem{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (dc _DraggingItemClass) Alloc() DraggingItem {
	rv := objc.Send[DraggingItem](objc.ID(dc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (dc _DraggingItemClass) New() DraggingItem {
	rv := objc.Send[DraggingItem](objc.ID(dc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (d_ DraggingItem) Init() DraggingItem {
	rv := objc.Send[DraggingItem](d_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (d_ DraggingItem) Autorelease() DraggingItem {
	rv := objc.Send[DraggingItem](d_.ID, objc.Sel("autorelease"))
	return rv
}

// NewDraggingItem creates a new DraggingItem instance.
func NewDraggingItem() DraggingItem {
	return getDraggingItemClass().New()
}




