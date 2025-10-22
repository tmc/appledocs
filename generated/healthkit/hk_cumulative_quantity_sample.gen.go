// Code generated from Apple documentation for HealthKit. DO NOT EDIT.

package healthkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [HKCumulativeQuantitySample] class.
var (
	HKCumulativeQuantitySampleClass     _HKCumulativeQuantitySampleClass
	HKCumulativeQuantitySampleClassOnce sync.Once
)

func getHKCumulativeQuantitySampleClass() _HKCumulativeQuantitySampleClass {
	HKCumulativeQuantitySampleClassOnce.Do(func() {
		HKCumulativeQuantitySampleClass = _HKCumulativeQuantitySampleClass{objc.GetClass("HKCumulativeQuantitySample")}
	})
	return HKCumulativeQuantitySampleClass
}

type _HKCumulativeQuantitySampleClass struct {
	class objc.Class
}

// An interface definition for the [HKCumulativeQuantitySample] class.
type IHKCumulativeQuantitySample interface {
	IHKQuantitySample
	SumQuantity() HKQuantity
	SetSumQuantity(value IHKQuantity)
}

// A sample that represents a cumulative quantity.
//
// A quantity sample contains one or more objects. Each quantity represents a single piece of data with a single numeric value and the value’s associated units. Use these samples to store data that accumulates over time, such as step count, active energy burned, or walking distance. The class is a concrete subclass of the class. Cumulative quantity samples are immutable; you set the sample’s properties when you create it, and they cannot change.
//
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKCumulativeQuantitySample
type HKCumulativeQuantitySample struct {
	HKQuantitySample
}

// HKCumulativeQuantitySampleFrom constructs a [HKCumulativeQuantitySample] from an unsafe.Pointer.
//
// A sample that represents a cumulative quantity.
func HKCumulativeQuantitySampleFrom(ptr unsafe.Pointer) HKCumulativeQuantitySample {
	return HKCumulativeQuantitySample{
		HKQuantitySample: HKQuantitySampleFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (hc _HKCumulativeQuantitySampleClass) Alloc() HKCumulativeQuantitySample {
	rv := objc.Send[HKCumulativeQuantitySample](objc.ID(hc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (hc _HKCumulativeQuantitySampleClass) New() HKCumulativeQuantitySample {
	rv := objc.Send[HKCumulativeQuantitySample](objc.ID(hc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (h_ HKCumulativeQuantitySample) Init() HKCumulativeQuantitySample {
	rv := objc.Send[HKCumulativeQuantitySample](h_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (h_ HKCumulativeQuantitySample) Autorelease() HKCumulativeQuantitySample {
	rv := objc.Send[HKCumulativeQuantitySample](h_.ID, objc.Sel("autorelease"))
	return rv
}

// NewHKCumulativeQuantitySample creates a new HKCumulativeQuantitySample instance.
func NewHKCumulativeQuantitySample() HKCumulativeQuantitySample {
	return getHKCumulativeQuantitySampleClass().New()
}


// The sum of all the quantities contained by the sample.
//
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkcumulativequantitysample/sumquantity
func (h_ HKCumulativeQuantitySample) SumQuantity() HKQuantity {
	rv := objc.Send[HKQuantity](h_.ID, objc.Sel("sumQuantity"))
	return rv
}


// SetSumQuantity sets the value of the sumQuantity property.
// The sum of all the quantities contained by the sample.

//
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkcumulativequantitysample/sumquantity
func (h_ HKCumulativeQuantitySample) SetSumQuantity(value IHKQuantity) {
	objc.Send[objc.ID](h_.ID, objc.Sel("setSumQuantity:"), value)
}



