// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [DraggingImageComponent] class.
var (
	draggingImageComponentClass     _DraggingImageComponentClass
	draggingImageComponentClassOnce sync.Once
)

func getDraggingImageComponentClass() _DraggingImageComponentClass {
	draggingImageComponentClassOnce.Do(func() {
		draggingImageComponentClass = _DraggingImageComponentClass{objc.GetClass("NSDraggingImageComponent")}
	})
	return draggingImageComponentClass
}

type _DraggingImageComponentClass struct {
	class objc.Class
}

// An interface definition for the [DraggingImageComponent] class.
type IDraggingImageComponent interface {
	objectivec.IObject
}

// A single object in a dragging item. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDraggingImageComponent

type DraggingImageComponent struct {
	objectivec.Object
}

// DraggingImageComponentFrom constructs a [DraggingImageComponent] from an unsafe.Pointer.
//
// A single object in a dragging item.
func DraggingImageComponentFrom(ptr unsafe.Pointer) DraggingImageComponent {
	return DraggingImageComponent{objectivec.Object{objc.ID(ptr)}}
}
// Alloc allocates a new instance without initialization.
func (dc _DraggingImageComponentClass) Alloc() DraggingImageComponent {
	rv := objc.Send[DraggingImageComponent](objc.ID(dc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (dc _DraggingImageComponentClass) New() DraggingImageComponent {
	rv := objc.Send[DraggingImageComponent](objc.ID(dc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (d_ DraggingImageComponent) Init() DraggingImageComponent {
	rv := objc.Send[DraggingImageComponent](d_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (d_ DraggingImageComponent) Autorelease() DraggingImageComponent {
	rv := objc.Send[DraggingImageComponent](d_.ID, objc.Sel("autorelease"))
	return rv
}

// NewDraggingImageComponent creates a new DraggingImageComponent instance.
func NewDraggingImageComponent() DraggingImageComponent {
	return getDraggingImageComponentClass().New()
}




