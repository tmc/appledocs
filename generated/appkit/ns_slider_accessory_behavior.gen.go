// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [SliderAccessoryBehavior] class.
var (
	SliderAccessoryBehaviorClass     _SliderAccessoryBehaviorClass
	SliderAccessoryBehaviorClassOnce sync.Once
)

func getSliderAccessoryBehaviorClass() _SliderAccessoryBehaviorClass {
	SliderAccessoryBehaviorClassOnce.Do(func() {
		SliderAccessoryBehaviorClass = _SliderAccessoryBehaviorClass{objc.GetClass("NSSliderAccessoryBehavior")}
	})
	return SliderAccessoryBehaviorClass
}

type _SliderAccessoryBehaviorClass struct {
	class objc.Class
}

// An interface definition for the [SliderAccessoryBehavior] class.
type ISliderAccessoryBehavior interface {
	objectivec.IObject
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSliderAccessoryBehavior
type SliderAccessoryBehavior struct {
	objectivec.Object
}

// SliderAccessoryBehaviorFrom constructs a [SliderAccessoryBehavior] from an unsafe.Pointer.
func SliderAccessoryBehaviorFrom(ptr unsafe.Pointer) SliderAccessoryBehavior {
	return SliderAccessoryBehavior{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (sc _SliderAccessoryBehaviorClass) Alloc() SliderAccessoryBehavior {
	rv := objc.Send[SliderAccessoryBehavior](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (sc _SliderAccessoryBehaviorClass) New() SliderAccessoryBehavior {
	rv := objc.Send[SliderAccessoryBehavior](objc.ID(sc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (s_ SliderAccessoryBehavior) Init() SliderAccessoryBehavior {
	rv := objc.Send[SliderAccessoryBehavior](s_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s_ SliderAccessoryBehavior) Autorelease() SliderAccessoryBehavior {
	rv := objc.Send[SliderAccessoryBehavior](s_.ID, objc.Sel("autorelease"))
	return rv
}

// NewSliderAccessoryBehavior creates a new SliderAccessoryBehavior instance.
func NewSliderAccessoryBehavior() SliderAccessoryBehavior {
	return getSliderAccessoryBehaviorClass().New()
}




