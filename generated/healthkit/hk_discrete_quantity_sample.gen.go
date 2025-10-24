// Code generated from Apple documentation for HealthKit. DO NOT EDIT.

package healthkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
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
	// properties:
	AverageQuantity() IHKQuantity
	SetAverageQuantity(value IHKQuantity)
	MaximumQuantity() IHKQuantity
	SetMaximumQuantity(value IHKQuantity)
	MinimumQuantity() IHKQuantity
	SetMinimumQuantity(value IHKQuantity)
	MostRecentQuantity() IHKQuantity
	SetMostRecentQuantity(value IHKQuantity)
	MostRecentQuantityDateInterval() objc.IObject /* cross-framework: DateInterval */
	SetMostRecentQuantityDateInterval(value objc.IObject /* cross-framework: DateInterval */)
	HKPredicateKeyPathAverage() objc.IObject /* cross-framework: NSString */
	HKPredicateKeyPathMax() objc.IObject /* cross-framework: NSString */
	HKPredicateKeyPathMin() objc.IObject /* cross-framework: NSString */
	HKPredicateKeyPathMostRecent() objc.IObject /* cross-framework: NSString */
	HKPredicateKeyPathMostRecentDuration() objc.IObject /* cross-framework: NSString */
	HKPredicateKeyPathMostRecentEndDate() objc.IObject /* cross-framework: NSString */
	HKPredicateKeyPathMostRecentStartDate() objc.IObject /* cross-framework: NSString */
	// methods:
}

// A sample that represents a discrete quantity.
//
// A quantity sample contains one or more objects. Each quantity represents a single piece of data with a single numeric value and the value’s associated units. Use these samples to store data representing independent measurements, such as height, heart rate, or temperature. The class is a concrete subclass of the class. Discrete quantity samples are immutable; you set the sample’s properties when you create it, and they cannot change.


// A sample that represents a discrete quantity.
//
// [Full Topic]
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



// The average of all quantities contained by the sample.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkdiscretequantitysample/averagequantity
func (h_ HKDiscreteQuantitySample) AverageQuantity() IHKQuantity {
	rv := objc.Send[HKQuantity](h_.ID, objc.Sel("averageQuantity"))
	return rv
}


// The average of all quantities contained by the sample.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkdiscretequantitysample/averagequantity
func (h_ HKDiscreteQuantitySample) SetAverageQuantity(value IHKQuantity) {
	objc.Send[objc.ID](h_.ID, objc.Sel("setAverageQuantity:"), value)
}


// The maximum quantity contained by the sample.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkdiscretequantitysample/maximumquantity
func (h_ HKDiscreteQuantitySample) MaximumQuantity() IHKQuantity {
	rv := objc.Send[HKQuantity](h_.ID, objc.Sel("maximumQuantity"))
	return rv
}


// The maximum quantity contained by the sample.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkdiscretequantitysample/maximumquantity
func (h_ HKDiscreteQuantitySample) SetMaximumQuantity(value IHKQuantity) {
	objc.Send[objc.ID](h_.ID, objc.Sel("setMaximumQuantity:"), value)
}


// The minimum value contained by the sample.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkdiscretequantitysample/minimumquantity
func (h_ HKDiscreteQuantitySample) MinimumQuantity() IHKQuantity {
	rv := objc.Send[HKQuantity](h_.ID, objc.Sel("minimumQuantity"))
	return rv
}


// The minimum value contained by the sample.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkdiscretequantitysample/minimumquantity
func (h_ HKDiscreteQuantitySample) SetMinimumQuantity(value IHKQuantity) {
	objc.Send[objc.ID](h_.ID, objc.Sel("setMinimumQuantity:"), value)
}


// The most recent quantity contained by the sample.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkdiscretequantitysample/mostrecentquantity
func (h_ HKDiscreteQuantitySample) MostRecentQuantity() IHKQuantity {
	rv := objc.Send[HKQuantity](h_.ID, objc.Sel("mostRecentQuantity"))
	return rv
}


// The most recent quantity contained by the sample.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkdiscretequantitysample/mostrecentquantity
func (h_ HKDiscreteQuantitySample) SetMostRecentQuantity(value IHKQuantity) {
	objc.Send[objc.ID](h_.ID, objc.Sel("setMostRecentQuantity:"), value)
}


// The date interval for the most recent quantity contained by the sample.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkdiscretequantitysample/mostrecentquantitydateinterval
func (h_ HKDiscreteQuantitySample) MostRecentQuantityDateInterval() objc.IObject /* cross-framework: DateInterval */ {
	rv := objc.Send[foundation.DateInterval](h_.ID, objc.Sel("mostRecentQuantityDateInterval"))
	return rv
}


// The date interval for the most recent quantity contained by the sample.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkdiscretequantitysample/mostrecentquantitydateinterval
func (h_ HKDiscreteQuantitySample) SetMostRecentQuantityDateInterval(value objc.IObject /* cross-framework: DateInterval */) {
	objc.Send[objc.ID](h_.ID, objc.Sel("setMostRecentQuantityDateInterval:"), value)
}


// The key path for the sample’s average quantity.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkpredicatekeypathaverage
func (h_ HKDiscreteQuantitySample) HKPredicateKeyPathAverage() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](h_.ID, objc.Sel("HKPredicateKeyPathAverage"))
	return rv
}


// The key path for the sample’s maximum quantity.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkpredicatekeypathmax
func (h_ HKDiscreteQuantitySample) HKPredicateKeyPathMax() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](h_.ID, objc.Sel("HKPredicateKeyPathMax"))
	return rv
}


// The key path for the sample’s minimum quantity.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkpredicatekeypathmin
func (h_ HKDiscreteQuantitySample) HKPredicateKeyPathMin() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](h_.ID, objc.Sel("HKPredicateKeyPathMin"))
	return rv
}


// The key path for the sample’s most recent quantity.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkpredicatekeypathmostrecent
func (h_ HKDiscreteQuantitySample) HKPredicateKeyPathMostRecent() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](h_.ID, objc.Sel("HKPredicateKeyPathMostRecent"))
	return rv
}


// A key path for the duration of the sample’s most recent quantity.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkpredicatekeypathmostrecentduration
func (h_ HKDiscreteQuantitySample) HKPredicateKeyPathMostRecentDuration() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](h_.ID, objc.Sel("HKPredicateKeyPathMostRecentDuration"))
	return rv
}


// The key path for the end date of the sample’s most recent quantity.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkpredicatekeypathmostrecentenddate
func (h_ HKDiscreteQuantitySample) HKPredicateKeyPathMostRecentEndDate() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](h_.ID, objc.Sel("HKPredicateKeyPathMostRecentEndDate"))
	return rv
}


// The key path for the start date of the sample’s most recent quantity.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkpredicatekeypathmostrecentstartdate
func (h_ HKDiscreteQuantitySample) HKPredicateKeyPathMostRecentStartDate() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](h_.ID, objc.Sel("HKPredicateKeyPathMostRecentStartDate"))
	return rv
}



