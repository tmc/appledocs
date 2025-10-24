// Code generated from Apple documentation for HealthKit. DO NOT EDIT.

package healthkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

/* debug [class.gen.go]: Generating class HKDiscreteQuantitySample */


/* debug [class_header]: Header for HKDiscreteQuantitySample */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for HKDiscreteQuantitySample */
// An interface definition for the [HKDiscreteQuantitySample] class.
type IHKDiscreteQuantitySample interface {
	IHKQuantitySample
	
/* debug [class_interface_properties]: Properties for HKDiscreteQuantitySample */
	// properties:
	AverageQuantity() IHKQuantity
	MaximumQuantity() IHKQuantity
	MinimumQuantity() IHKQuantity
	MostRecentQuantity() IHKQuantity
	MostRecentQuantityDateInterval() foundation.DateInterval
	HKPredicateKeyPathAverage() objc.IObject /* cross-framework: NSString */
	HKPredicateKeyPathMax() objc.IObject /* cross-framework: NSString */
	HKPredicateKeyPathMin() objc.IObject /* cross-framework: NSString */
	HKPredicateKeyPathMostRecent() objc.IObject /* cross-framework: NSString */
	HKPredicateKeyPathMostRecentDuration() objc.IObject /* cross-framework: NSString */
	HKPredicateKeyPathMostRecentEndDate() objc.IObject /* cross-framework: NSString */
	HKPredicateKeyPathMostRecentStartDate() objc.IObject /* cross-framework: NSString */
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for HKDiscreteQuantitySample */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for HKDiscreteQuantitySample */
// Alloc allocates a new instance without initialization.
func (hc _HKDiscreteQuantitySampleClass) Alloc() HKDiscreteQuantitySample {
	rv := objc.Send[HKDiscreteQuantitySample](objc.ID(hc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for HKDiscreteQuantitySample */
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for HKDiscreteQuantitySample *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for HKDiscreteQuantitySample */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for HKDiscreteQuantitySample */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for HKDiscreteQuantitySample */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for HKDiscreteQuantitySample */

// The average of all quantities contained by the sample.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKDiscreteQuantitySample/averageQuantity
func (h_ HKDiscreteQuantitySample) AverageQuantity() IHKQuantity {
	rv := objc.Send[HKQuantity](h_.ID, objc.Sel("averageQuantity"))
	return rv
}/* debug [instance_properties/getter]: averageQuantity */


// The maximum quantity contained by the sample.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKDiscreteQuantitySample/maximumQuantity
func (h_ HKDiscreteQuantitySample) MaximumQuantity() IHKQuantity {
	rv := objc.Send[HKQuantity](h_.ID, objc.Sel("maximumQuantity"))
	return rv
}/* debug [instance_properties/getter]: maximumQuantity */


// The minimum value contained by the sample.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKDiscreteQuantitySample/minimumQuantity
func (h_ HKDiscreteQuantitySample) MinimumQuantity() IHKQuantity {
	rv := objc.Send[HKQuantity](h_.ID, objc.Sel("minimumQuantity"))
	return rv
}/* debug [instance_properties/getter]: minimumQuantity */


// The most recent quantity contained by the sample.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKDiscreteQuantitySample/mostRecentQuantity
func (h_ HKDiscreteQuantitySample) MostRecentQuantity() IHKQuantity {
	rv := objc.Send[HKQuantity](h_.ID, objc.Sel("mostRecentQuantity"))
	return rv
}/* debug [instance_properties/getter]: mostRecentQuantity */


// The date interval for the most recent quantity contained by the sample.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKDiscreteQuantitySample/mostRecentQuantityDateInterval
func (h_ HKDiscreteQuantitySample) MostRecentQuantityDateInterval() foundation.DateInterval {
	rv := objc.Send[foundation.DateInterval](h_.ID, objc.Sel("mostRecentQuantityDateInterval"))
	return rv
}/* debug [instance_properties/getter]: mostRecentQuantityDateInterval */


// The key path for the sample’s average quantity.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkpredicatekeypathaverage
func (h_ HKDiscreteQuantitySample) HKPredicateKeyPathAverage() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](h_.ID, objc.Sel("HKPredicateKeyPathAverage"))
	return rv
}/* debug [instance_properties/getter]: HKPredicateKeyPathAverage */


// The key path for the sample’s maximum quantity.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkpredicatekeypathmax
func (h_ HKDiscreteQuantitySample) HKPredicateKeyPathMax() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](h_.ID, objc.Sel("HKPredicateKeyPathMax"))
	return rv
}/* debug [instance_properties/getter]: HKPredicateKeyPathMax */


// The key path for the sample’s minimum quantity.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkpredicatekeypathmin
func (h_ HKDiscreteQuantitySample) HKPredicateKeyPathMin() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](h_.ID, objc.Sel("HKPredicateKeyPathMin"))
	return rv
}/* debug [instance_properties/getter]: HKPredicateKeyPathMin */


// The key path for the sample’s most recent quantity.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkpredicatekeypathmostrecent
func (h_ HKDiscreteQuantitySample) HKPredicateKeyPathMostRecent() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](h_.ID, objc.Sel("HKPredicateKeyPathMostRecent"))
	return rv
}/* debug [instance_properties/getter]: HKPredicateKeyPathMostRecent */


// A key path for the duration of the sample’s most recent quantity.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkpredicatekeypathmostrecentduration
func (h_ HKDiscreteQuantitySample) HKPredicateKeyPathMostRecentDuration() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](h_.ID, objc.Sel("HKPredicateKeyPathMostRecentDuration"))
	return rv
}/* debug [instance_properties/getter]: HKPredicateKeyPathMostRecentDuration */


// The key path for the end date of the sample’s most recent quantity.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkpredicatekeypathmostrecentenddate
func (h_ HKDiscreteQuantitySample) HKPredicateKeyPathMostRecentEndDate() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](h_.ID, objc.Sel("HKPredicateKeyPathMostRecentEndDate"))
	return rv
}/* debug [instance_properties/getter]: HKPredicateKeyPathMostRecentEndDate */


// The key path for the start date of the sample’s most recent quantity.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkpredicatekeypathmostrecentstartdate
func (h_ HKDiscreteQuantitySample) HKPredicateKeyPathMostRecentStartDate() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](h_.ID, objc.Sel("HKPredicateKeyPathMostRecentStartDate"))
	return rv
}/* debug [instance_properties/getter]: HKPredicateKeyPathMostRecentStartDate */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class HKDiscreteQuantitySample */



