// Code generated from Apple documentation for HealthKit. DO NOT EDIT.

package healthkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [HKQuantity] class.
var (
	HKQuantityClass     _HKQuantityClass
	HKQuantityClassOnce sync.Once
)

func getHKQuantityClass() _HKQuantityClass {
	HKQuantityClassOnce.Do(func() {
		HKQuantityClass = _HKQuantityClass{objc.GetClass("HKQuantity")}
	})
	return HKQuantityClass
}

type _HKQuantityClass struct {
	class objc.Class
}

// An interface definition for the [HKQuantity] class.
type IHKQuantity interface {
	objectivec.IObject
	// properties:
	// methods:
}

// An object that stores a value for a given unit.
//
// HealthKit uses quantity objects to store numerical data. When you create a quantity, you provide both the quantity’s value and unit. Quantities are immutable objects: Their values are set when the object is first created and cannot change.


// An object that stores a value for a given unit.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKQuantity
type HKQuantity struct {
	objectivec.Object
}

// HKQuantityFrom constructs a [HKQuantity] from an unsafe.Pointer.
//
// An object that stores a value for a given unit.
func HKQuantityFrom(ptr unsafe.Pointer) HKQuantity {
	return HKQuantity{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (hc _HKQuantityClass) Alloc() HKQuantity {
	rv := objc.Send[HKQuantity](objc.ID(hc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (hc _HKQuantityClass) New() HKQuantity {
	rv := objc.Send[HKQuantity](objc.ID(hc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (h_ HKQuantity) Init() HKQuantity {
	rv := objc.Send[HKQuantity](h_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (h_ HKQuantity) Autorelease() HKQuantity {
	rv := objc.Send[HKQuantity](h_.ID, objc.Sel("autorelease"))
	return rv
}

// NewHKQuantity creates a new HKQuantity instance.
func NewHKQuantity() HKQuantity {
	return getHKQuantityClass().New()
}




