// Code generated from Apple documentation for HealthKit. DO NOT EDIT.

package healthkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [HKVisionPrism] class.
var (
	HKVisionPrismClass     _HKVisionPrismClass
	HKVisionPrismClassOnce sync.Once
)

func getHKVisionPrismClass() _HKVisionPrismClass {
	HKVisionPrismClassOnce.Do(func() {
		HKVisionPrismClass = _HKVisionPrismClass{objc.GetClass("HKVisionPrism")}
	})
	return HKVisionPrismClass
}

type _HKVisionPrismClass struct {
	class objc.Class
}

// An interface definition for the [HKVisionPrism] class.
type IHKVisionPrism interface {
	objectivec.IObject
}

// Prescription data for eye alignment.
//
// To include prism information in a glasses prescription, start by creating an object. Then, pass this value to the ’s initializer. Finally, create the glasses prescription and save it to the HealthKit store.
//
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKVisionPrism
type HKVisionPrism struct {
	objectivec.Object
}

// HKVisionPrismFrom constructs a [HKVisionPrism] from an unsafe.Pointer.
//
// Prescription data for eye alignment.
func HKVisionPrismFrom(ptr unsafe.Pointer) HKVisionPrism {
	return HKVisionPrism{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (hc _HKVisionPrismClass) Alloc() HKVisionPrism {
	rv := objc.Send[HKVisionPrism](objc.ID(hc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (hc _HKVisionPrismClass) New() HKVisionPrism {
	rv := objc.Send[HKVisionPrism](objc.ID(hc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (h_ HKVisionPrism) Init() HKVisionPrism {
	rv := objc.Send[HKVisionPrism](h_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (h_ HKVisionPrism) Autorelease() HKVisionPrism {
	rv := objc.Send[HKVisionPrism](h_.ID, objc.Sel("autorelease"))
	return rv
}

// NewHKVisionPrism creates a new HKVisionPrism instance.
func NewHKVisionPrism() HKVisionPrism {
	return getHKVisionPrismClass().New()
}




