// Code generated from Apple documentation for HealthKit. DO NOT EDIT.

package healthkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [HKCumulativeQuantitySeriesSample] class.
var (
	HKCumulativeQuantitySeriesSampleClass     _HKCumulativeQuantitySeriesSampleClass
	HKCumulativeQuantitySeriesSampleClassOnce sync.Once
)

func getHKCumulativeQuantitySeriesSampleClass() _HKCumulativeQuantitySeriesSampleClass {
	HKCumulativeQuantitySeriesSampleClassOnce.Do(func() {
		HKCumulativeQuantitySeriesSampleClass = _HKCumulativeQuantitySeriesSampleClass{objc.GetClass("HKCumulativeQuantitySeriesSample")}
	})
	return HKCumulativeQuantitySeriesSampleClass
}

type _HKCumulativeQuantitySeriesSampleClass struct {
	class objc.Class
}

// An interface definition for the [HKCumulativeQuantitySeriesSample] class.
type IHKCumulativeQuantitySeriesSample interface {
	IHKCumulativeQuantitySample
}

// A sample representing a series of cumulative quantity values.
//
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKCumulativeQuantitySeriesSample
type HKCumulativeQuantitySeriesSample struct {
	HKCumulativeQuantitySample
}

// HKCumulativeQuantitySeriesSampleFrom constructs a [HKCumulativeQuantitySeriesSample] from an unsafe.Pointer.
//
// A sample representing a series of cumulative quantity values.
func HKCumulativeQuantitySeriesSampleFrom(ptr unsafe.Pointer) HKCumulativeQuantitySeriesSample {
	return HKCumulativeQuantitySeriesSample{
		HKCumulativeQuantitySample: HKCumulativeQuantitySampleFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (hc _HKCumulativeQuantitySeriesSampleClass) Alloc() HKCumulativeQuantitySeriesSample {
	rv := objc.Send[HKCumulativeQuantitySeriesSample](objc.ID(hc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (hc _HKCumulativeQuantitySeriesSampleClass) New() HKCumulativeQuantitySeriesSample {
	rv := objc.Send[HKCumulativeQuantitySeriesSample](objc.ID(hc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (h_ HKCumulativeQuantitySeriesSample) Init() HKCumulativeQuantitySeriesSample {
	rv := objc.Send[HKCumulativeQuantitySeriesSample](h_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (h_ HKCumulativeQuantitySeriesSample) Autorelease() HKCumulativeQuantitySeriesSample {
	rv := objc.Send[HKCumulativeQuantitySeriesSample](h_.ID, objc.Sel("autorelease"))
	return rv
}

// NewHKCumulativeQuantitySeriesSample creates a new HKCumulativeQuantitySeriesSample instance.
func NewHKCumulativeQuantitySeriesSample() HKCumulativeQuantitySeriesSample {
	return getHKCumulativeQuantitySeriesSampleClass().New()
}




