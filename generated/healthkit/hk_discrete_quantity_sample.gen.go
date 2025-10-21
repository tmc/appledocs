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


// A key path for the duration of the sample’s most recent quantity.
//
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkpredicatekeypathmostrecentduration
func (h_ HKDiscreteQuantitySample) HKPredicateKeyPathMostRecentDuration() string {
	rv := objc.Send[string](h_.ID, objc.Sel("HKPredicateKeyPathMostRecentDuration"))
	return rv
}

// The key path for the sample’s most recent quantity.
//
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkpredicatekeypathmostrecent
func (h_ HKDiscreteQuantitySample) HKPredicateKeyPathMostRecent() string {
	rv := objc.Send[string](h_.ID, objc.Sel("HKPredicateKeyPathMostRecent"))
	return rv
}

// The key path for the sample’s average quantity.
//
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkpredicatekeypathaverage
func (h_ HKDiscreteQuantitySample) HKPredicateKeyPathAverage() string {
	rv := objc.Send[string](h_.ID, objc.Sel("HKPredicateKeyPathAverage"))
	return rv
}

// The maximum quantity contained by the sample.
//
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkdiscretequantitysample/maximumquantity
func (h_ HKDiscreteQuantitySample) MaximumQuantity() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](h_.ID, objc.Sel("maximumQuantity"))
	return rv
}


// SetMaximumQuantity sets the value of the maximumQuantity property.
// The maximum quantity contained by the sample.

//
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkdiscretequantitysample/maximumquantity
func (h_ HKDiscreteQuantitySample) SetMaximumQuantity(value unsafe.Pointer) {
	objc.Send[objc.ID](h_.ID, objc.Sel("setMaximumQuantity:"), value)
}

// The most recent quantity contained by the sample.
//
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkdiscretequantitysample/mostrecentquantity
func (h_ HKDiscreteQuantitySample) MostRecentQuantity() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](h_.ID, objc.Sel("mostRecentQuantity"))
	return rv
}


// SetMostRecentQuantity sets the value of the mostRecentQuantity property.
// The most recent quantity contained by the sample.

//
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkdiscretequantitysample/mostrecentquantity
func (h_ HKDiscreteQuantitySample) SetMostRecentQuantity(value unsafe.Pointer) {
	objc.Send[objc.ID](h_.ID, objc.Sel("setMostRecentQuantity:"), value)
}

// The minimum value contained by the sample.
//
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkdiscretequantitysample/minimumquantity
func (h_ HKDiscreteQuantitySample) MinimumQuantity() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](h_.ID, objc.Sel("minimumQuantity"))
	return rv
}


// SetMinimumQuantity sets the value of the minimumQuantity property.
// The minimum value contained by the sample.

//
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkdiscretequantitysample/minimumquantity
func (h_ HKDiscreteQuantitySample) SetMinimumQuantity(value unsafe.Pointer) {
	objc.Send[objc.ID](h_.ID, objc.Sel("setMinimumQuantity:"), value)
}

// The date interval for the most recent quantity contained by the sample.
//
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkdiscretequantitysample/mostrecentquantitydateinterval
func (h_ HKDiscreteQuantitySample) MostRecentQuantityDateInterval() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](h_.ID, objc.Sel("mostRecentQuantityDateInterval"))
	return rv
}


// SetMostRecentQuantityDateInterval sets the value of the mostRecentQuantityDateInterval property.
// The date interval for the most recent quantity contained by the sample.

//
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkdiscretequantitysample/mostrecentquantitydateinterval
func (h_ HKDiscreteQuantitySample) SetMostRecentQuantityDateInterval(value unsafe.Pointer) {
	objc.Send[objc.ID](h_.ID, objc.Sel("setMostRecentQuantityDateInterval:"), value)
}

// The average of all quantities contained by the sample.
//
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkdiscretequantitysample/averagequantity
func (h_ HKDiscreteQuantitySample) AverageQuantity() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](h_.ID, objc.Sel("averageQuantity"))
	return rv
}


// SetAverageQuantity sets the value of the averageQuantity property.
// The average of all quantities contained by the sample.

//
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkdiscretequantitysample/averagequantity
func (h_ HKDiscreteQuantitySample) SetAverageQuantity(value unsafe.Pointer) {
	objc.Send[objc.ID](h_.ID, objc.Sel("setAverageQuantity:"), value)
}

// The key path for the start date of the sample’s most recent quantity.
//
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkpredicatekeypathmostrecentstartdate
func (h_ HKDiscreteQuantitySample) HKPredicateKeyPathMostRecentStartDate() string {
	rv := objc.Send[string](h_.ID, objc.Sel("HKPredicateKeyPathMostRecentStartDate"))
	return rv
}

// The key path for the end date of the sample’s most recent quantity.
//
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkpredicatekeypathmostrecentenddate
func (h_ HKDiscreteQuantitySample) HKPredicateKeyPathMostRecentEndDate() string {
	rv := objc.Send[string](h_.ID, objc.Sel("HKPredicateKeyPathMostRecentEndDate"))
	return rv
}

// The key path for the sample’s minimum quantity.
//
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkpredicatekeypathmin
func (h_ HKDiscreteQuantitySample) HKPredicateKeyPathMin() string {
	rv := objc.Send[string](h_.ID, objc.Sel("HKPredicateKeyPathMin"))
	return rv
}

// The key path for the sample’s maximum quantity.
//
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkpredicatekeypathmax
func (h_ HKDiscreteQuantitySample) HKPredicateKeyPathMax() string {
	rv := objc.Send[string](h_.ID, objc.Sel("HKPredicateKeyPathMax"))
	return rv
}



