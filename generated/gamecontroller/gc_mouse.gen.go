// Code generated from Apple documentation for GameController. DO NOT EDIT.

package gamecontroller

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [GCMouse] class.
var (
	GCMouseClass     _GCMouseClass
	GCMouseClassOnce sync.Once
)

func getGCMouseClass() _GCMouseClass {
	GCMouseClassOnce.Do(func() {
		GCMouseClass = _GCMouseClass{objc.GetClass("GCMouse")}
	})
	return GCMouseClass
}

type _GCMouseClass struct {
	class objc.Class
}

// An interface definition for the [GCMouse] class.
type IGCMouse interface {
	objectivec.IObject
}

// An object that represents a physical mouse connected to a device.
//
// To get a mouse object and its input values, register for the (Swift) or (Objective-C) notification for when a mouse connects to the device. Then register for the (Swift) or (Objective-C) notification for when it becomes the mouse. Alternatively, use the class property or the class method to get a mouse object. Then get the current input values from the mouse object’s controller profile.
//
// [Full Topic]: https://developer.apple.com/documentation/GameController/GCMouse
type GCMouse struct {
	objectivec.Object
}

// GCMouseFrom constructs a [GCMouse] from an unsafe.Pointer.
//
// An object that represents a physical mouse connected to a device.
func GCMouseFrom(ptr unsafe.Pointer) GCMouse {
	return GCMouse{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (gc _GCMouseClass) Alloc() GCMouse {
	rv := objc.Send[GCMouse](objc.ID(gc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (gc _GCMouseClass) New() GCMouse {
	rv := objc.Send[GCMouse](objc.ID(gc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (g_ GCMouse) Init() GCMouse {
	rv := objc.Send[GCMouse](g_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (g_ GCMouse) Autorelease() GCMouse {
	rv := objc.Send[GCMouse](g_.ID, objc.Sel("autorelease"))
	return rv
}

// NewGCMouse creates a new GCMouse instance.
func NewGCMouse() GCMouse {
	return getGCMouseClass().New()
}


// The most recent mouse that the user connects.
//
// [Full Topic]: https://developer.apple.com/documentation/GameController/GCMouse/current
func (gc _GCMouseClass) Current() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(gc.class), objc.Sel("current"))
	return rv
}
// The most recent mouse that the user connects.
//
// [Full Topic]: https://developer.apple.com/documentation/GameController/GCMouse/current
func (g_ GCMouse) Current() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](g_.ID, objc.Sel("current"))
	return rv
}



