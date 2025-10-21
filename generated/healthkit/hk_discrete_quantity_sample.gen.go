// Code generated from Apple documentation for HealthKit. DO NOT EDIT.

package healthkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [HKDiscreteQuantitySample] class.
var (
	HKDiscreteQuantitySampleClass     _HKDiscreteQuantitySampleClass
	HKDiscreteQuantitySampleClassOnce sync.Once
)

func getHKDiscreteQuantitySampleClass() _HKDiscreteQuantitySampleClass {
	HKDiscreteQuantitySampleClassOnce.Do(func() {
		HKDiscreteQuantitySampleClass = _HKDiscreteQuantitySampleClass{objc.GetClass("HKDiscreteQuantitySample")}
	})
	return HKDiscreteQuantitySampleClass
}

type _HKDiscreteQuantitySampleClass struct {
	class objc.Class
}

// An interface definition for the [HKDiscreteQuantitySample] class.
type IHKDiscreteQuantitySample interface {
	IHKQuantitySample
}

// A sample that represents a discrete quantity.
//
// A quantity sample contains one or more objects. Each quantity represents a single piece of data with a single numeric value and the value’s associated units. Use these samples to store data representing independent measurements, such as height, heart rate, or temperature. The class is a concrete subclass of the class. Discrete quantity samples are immutable; you set the sample’s properties when you create it, and they cannot change.
//
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKDiscreteQuantitySample
type HKDiscreteQuantitySample struct {
	HKQuantitySample
}

// HKDiscreteQuantitySampleFrom constructs a [HKDiscreteQuantitySample] from an unsafe.Pointer.
//
// A sample that represents a discrete quantity.
func HKDiscreteQuantitySampleFrom(ptr unsafe.Pointer) HKDiscreteQuantitySample {
	return HKDiscreteQuantitySample{
		HKQuantitySample: HKQuantitySampleFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (hc _HKDiscreteQuantitySampleClass) Alloc() HKDiscreteQuantitySample {
	rv := objc.Send[HKDiscreteQuantitySample](objc.ID(hc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (hc _HKDiscreteQuantitySampleClass) New() HKDiscreteQuantitySample {
	rv := objc.Send[HKDiscreteQuantitySample](objc.ID(hc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (h_ HKDiscreteQuantitySample) Init() HKDiscreteQuantitySample {
	rv := objc.Send[HKDiscreteQuantitySample](h_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (h_ HKDiscreteQuantitySample) Autorelease() HKDiscreteQuantitySample {
	rv := objc.Send[HKDiscreteQuantitySample](h_.ID, objc.Sel("autorelease"))
	return rv
}

// NewHKDiscreteQuantitySample creates a new HKDiscreteQuantitySample instance.
func NewHKDiscreteQuantitySample() HKDiscreteQuantitySample {
	return getHKDiscreteQuantitySampleClass().New()
}




