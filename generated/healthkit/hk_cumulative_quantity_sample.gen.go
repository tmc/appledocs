// Code generated from Apple documentation for HealthKit. DO NOT EDIT.

package healthkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class HKCumulativeQuantitySample */


/* debug [class_header]: Header for HKCumulativeQuantitySample */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for HKCumulativeQuantitySample */
// An interface definition for the [HKCumulativeQuantitySample] class.
type IHKCumulativeQuantitySample interface {
	IHKQuantitySample
	
/* debug [class_interface_properties]: Properties for HKCumulativeQuantitySample */
	// properties:
	SumQuantity() IHKQuantity
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for HKCumulativeQuantitySample */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for HKCumulativeQuantitySample */
// Alloc allocates a new instance without initialization.
func (hc _HKCumulativeQuantitySampleClass) Alloc() HKCumulativeQuantitySample {
	rv := objc.Send[HKCumulativeQuantitySample](objc.ID(hc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for HKCumulativeQuantitySample */
// A sample that represents a cumulative quantity.
//
// A quantity sample contains one or more objects. Each quantity represents a single piece of data with a single numeric value and the value’s associated units. Use these samples to store data that accumulates over time, such as step count, active energy burned, or walking distance. The class is a concrete subclass of the class. Cumulative quantity samples are immutable; you set the sample’s properties when you create it, and they cannot change.


// A sample that represents a cumulative quantity.
//
// [Full Topic]
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for HKCumulativeQuantitySample *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for HKCumulativeQuantitySample */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for HKCumulativeQuantitySample */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for HKCumulativeQuantitySample */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for HKCumulativeQuantitySample */

// The sum of all the quantities contained by the sample.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKCumulativeQuantitySample/sumQuantity
func (h_ HKCumulativeQuantitySample) SumQuantity() IHKQuantity {
	rv := objc.Send[HKQuantity](h_.ID, objc.Sel("sumQuantity"))
	return rv
}/* debug [instance_properties/getter]: sumQuantity */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class HKCumulativeQuantitySample */



