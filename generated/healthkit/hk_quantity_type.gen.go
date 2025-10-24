// Code generated from Apple documentation for HealthKit. DO NOT EDIT.

package healthkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class HKQuantityType */


/* debug [class_header]: Header for HKQuantityType */
// The class instance for the [HKQuantityType] class.
var (
	HKQuantityTypeClass     _HKQuantityTypeClass
	HKQuantityTypeClassOnce sync.Once
)

func getHKQuantityTypeClass() _HKQuantityTypeClass {
	HKQuantityTypeClassOnce.Do(func() {
		HKQuantityTypeClass = _HKQuantityTypeClass{objc.GetClass("HKQuantityType")}
	})
	return HKQuantityTypeClass
}

type _HKQuantityTypeClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for HKQuantityType */
// An interface definition for the [HKQuantityType] class.
type IHKQuantityType interface {
	IHKSampleType
	
/* debug [class_interface_properties]: Properties for HKQuantityType */
	// properties:
	AggregationStyle() HKQuantityAggregationStyle
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for HKQuantityType */
	// methods:
	IsCompatibleWithUnit(unit IHKUnit) bool
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for HKQuantityType */
// Alloc allocates a new instance without initialization.
func (hc _HKQuantityTypeClass) Alloc() HKQuantityType {
	rv := objc.Send[HKQuantityType](objc.ID(hc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (hc _HKQuantityTypeClass) New() HKQuantityType {
	rv := objc.Send[HKQuantityType](objc.ID(hc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (h_ HKQuantityType) Init() HKQuantityType {
	rv := objc.Send[HKQuantityType](h_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (h_ HKQuantityType) Autorelease() HKQuantityType {
	rv := objc.Send[HKQuantityType](h_.ID, objc.Sel("autorelease"))
	return rv
}

// NewHKQuantityType creates a new HKQuantityType instance.
func NewHKQuantityType() HKQuantityType {
	return getHKQuantityTypeClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for HKQuantityType */
// A type that identifies samples that store numerical values.
//
// The class is a concrete subclass of the class. To create a quantity type instance, use the object type’s convenience method. Use quantity types to: Request permission to read or write matching quantity samples. Create and share matching quantity samples. Query for matching quantity samples.


// A type that identifies samples that store numerical values.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKQuantityType
type HKQuantityType struct {
	HKSampleType
}

// HKQuantityTypeFrom constructs a [HKQuantityType] from an unsafe.Pointer.
//
// A type that identifies samples that store numerical values.
func HKQuantityTypeFrom(ptr unsafe.Pointer) HKQuantityType {
	return HKQuantityType{
		HKSampleType: HKSampleTypeFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for HKQuantityType *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for HKQuantityType */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for HKQuantityType */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for HKQuantityType */

// Returns a Boolean value that indicates whether the quantity type is compatible with the given unit.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKQuantityType/is(compatibleWith:)
func (h_ HKQuantityType) IsCompatibleWithUnit(unit IHKUnit) bool {
	rv := objc.Send[bool](h_.ID, objc.Sel("isCompatibleWithUnit:"), unit)
	return rv
}/* debug [instance_methods/method]: IsCompatibleWithUnit */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for HKQuantityType */

// The aggregation style for the given quantity type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKQuantityType/aggregationStyle
func (h_ HKQuantityType) AggregationStyle() HKQuantityAggregationStyle {
	rv := objc.Send[HKQuantityAggregationStyle](h_.ID, objc.Sel("aggregationStyle"))
	return rv
}/* debug [instance_properties/getter]: aggregationStyle */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class HKQuantityType */



