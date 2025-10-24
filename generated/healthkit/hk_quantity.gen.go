// Code generated from Apple documentation for HealthKit. DO NOT EDIT.

package healthkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class HKQuantity */


/* debug [class_header]: Header for HKQuantity */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for HKQuantity */
// An interface definition for the [HKQuantity] class.
type IHKQuantity interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for HKQuantity */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for HKQuantity */
	// methods:
	Compare(quantity IHKQuantity) ComparisonResult /* not a class type */
	DoubleValueForUnit(unit IHKUnit) float64
	IsCompatibleWithUnit(unit IHKUnit) bool
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for HKQuantity */
// Alloc allocates a new instance without initialization.
func (hc _HKQuantityClass) Alloc() HKQuantity {
	rv := objc.Send[HKQuantity](objc.ID(hc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for HKQuantity */
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for HKQuantity */

// Instantiates and returns a new quantity object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKQuantity/init(unit:doubleValue:)
func NewHKQuantityWithUnitDoubleValue(unit IHKUnit, value float64) HKQuantity {
	rv := objc.Send[HKQuantity](objc.ID(getHKQuantityClass().class), objc.Sel("quantityWithUnit:doubleValue:"), unit, value)
	return rv
}/* debug [class_init_methods/constructor]: NewHKQuantityWithUnitDoubleValue */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for HKQuantity */

// Instantiates and returns a new quantity object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKQuantity/init(unit:doubleValue:)
func (hc _HKQuantityClass) QuantityWithUnitDoubleValue(unit IHKUnit, value float64) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(hc.class), objc.Sel("quantityWithUnit:doubleValue:"), unit, value)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=QuantityWithUnitDoubleValue) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for HKQuantity */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for HKQuantity */

// Compares two values after converting them to the same units.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKQuantity/compare(_:)
func (h_ HKQuantity) Compare(quantity IHKQuantity) ComparisonResult /* not a class type */ {
	rv := objc.Send[ComparisonResult](h_.ID, objc.Sel("compare:"), quantity)
	return rv
}/* debug [instance_methods/method]: Compare */


// Returns the quantity’s value in the provided unit.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKQuantity/doubleValue(for:)
func (h_ HKQuantity) DoubleValueForUnit(unit IHKUnit) float64 {
	rv := objc.Send[float64](h_.ID, objc.Sel("doubleValueForUnit:"), unit)
	return rv
}/* debug [instance_methods/method]: DoubleValueForUnit */


// Returns a boolean value indicating whether the quantity is compatible with the provided unit.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKQuantity/is(compatibleWith:)
func (h_ HKQuantity) IsCompatibleWithUnit(unit IHKUnit) bool {
	rv := objc.Send[bool](h_.ID, objc.Sel("isCompatibleWithUnit:"), unit)
	return rv
}/* debug [instance_methods/method]: IsCompatibleWithUnit */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for HKQuantity */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class HKQuantity */


