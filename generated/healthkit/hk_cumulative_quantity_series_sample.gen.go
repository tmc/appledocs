// Code generated from Apple documentation for HealthKit. DO NOT EDIT.

package healthkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

/* debug [class.gen.go]: Generating class HKCumulativeQuantitySeriesSample */


/* debug [class_header]: Header for HKCumulativeQuantitySeriesSample */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for HKCumulativeQuantitySeriesSample */
// An interface definition for the [HKCumulativeQuantitySeriesSample] class.
type IHKCumulativeQuantitySeriesSample interface {
	IHKCumulativeQuantitySample
	
/* debug [class_interface_properties]: Properties for HKCumulativeQuantitySeriesSample */
	// properties:
	Sum() IHKQuantity
	HKPredicateKeyPathSum() objc.IObject /* cross-framework: NSString */
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for HKCumulativeQuantitySeriesSample */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for HKCumulativeQuantitySeriesSample */
// Alloc allocates a new instance without initialization.
func (hc _HKCumulativeQuantitySeriesSampleClass) Alloc() HKCumulativeQuantitySeriesSample {
	rv := objc.Send[HKCumulativeQuantitySeriesSample](objc.ID(hc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for HKCumulativeQuantitySeriesSample */
// A sample representing a series of cumulative quantity values.


// A sample representing a series of cumulative quantity values.
//
// [Full Topic]
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for HKCumulativeQuantitySeriesSample *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for HKCumulativeQuantitySeriesSample */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for HKCumulativeQuantitySeriesSample */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for HKCumulativeQuantitySeriesSample */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for HKCumulativeQuantitySeriesSample */

// The sum of all the quantities in the series.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKCumulativeQuantitySeriesSample/sum
func (h_ HKCumulativeQuantitySeriesSample) Sum() IHKQuantity {
	rv := objc.Send[HKQuantity](h_.ID, objc.Sel("sum"))
	return rv
}/* debug [instance_properties/getter]: sum */


// The key path for accessing the sum of a quantity series inside a predicate format string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkpredicatekeypathsum
func (h_ HKCumulativeQuantitySeriesSample) HKPredicateKeyPathSum() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](h_.ID, objc.Sel("HKPredicateKeyPathSum"))
	return rv
}/* debug [instance_properties/getter]: HKPredicateKeyPathSum */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class HKCumulativeQuantitySeriesSample */



