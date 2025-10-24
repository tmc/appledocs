// Code generated from Apple documentation for HealthKit. DO NOT EDIT.

package healthkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [HKUnit] class.
var (
	HKUnitClass     _HKUnitClass
	HKUnitClassOnce sync.Once
)

func getHKUnitClass() _HKUnitClass {
	HKUnitClassOnce.Do(func() {
		HKUnitClass = _HKUnitClass{objc.GetClass("HKUnit")}
	})
	return HKUnitClass
}

type _HKUnitClass struct {
	class objc.Class
}

// An interface definition for the [HKUnit] class.
type IHKUnit interface {
	objectivec.IObject
	// properties:
	UnitString() objc.IObject /* cross-framework: NSString */
	SetUnitString(value objc.IObject /* cross-framework: NSString */)
	HKUnitMolarMassBloodGlucose() float64
	SetHKUnitMolarMassBloodGlucose(value float64)
	// methods:
}

// A class for managing the units of measure within HealthKit.
//
// The unit class supports most standard SI units (meters, seconds, and grams), SI units with prefixes (centimeters, milliseconds and kilograms) and equivalent non-SI units (feet, minutes, and pounds). HealthKit also supports creating complex units by mathematically combining existing units. You use units when working with HealthKit quantities. Quantities store both the value (as a data type) and its corresponding unit. You can then request the value from the quantity in any compatible units. For more information on working with quantities, see .


// A class for managing the units of measure within HealthKit.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKUnit
type HKUnit struct {
	objectivec.Object
}

// HKUnitFrom constructs a [HKUnit] from an unsafe.Pointer.
//
// A class for managing the units of measure within HealthKit.
func HKUnitFrom(ptr unsafe.Pointer) HKUnit {
	return HKUnit{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (hc _HKUnitClass) Alloc() HKUnit {
	rv := objc.Send[HKUnit](objc.ID(hc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (hc _HKUnitClass) New() HKUnit {
	rv := objc.Send[HKUnit](objc.ID(hc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (h_ HKUnit) Init() HKUnit {
	rv := objc.Send[HKUnit](h_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (h_ HKUnit) Autorelease() HKUnit {
	rv := objc.Send[HKUnit](h_.ID, objc.Sel("autorelease"))
	return rv
}

// NewHKUnit creates a new HKUnit instance.
func NewHKUnit() HKUnit {
	return getHKUnitClass().New()
}



// A string representation of the unit object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkunit/unitstring
func (h_ HKUnit) UnitString() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](h_.ID, objc.Sel("unitString"))
	return rv
}


// A string representation of the unit object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkunit/unitstring
func (h_ HKUnit) SetUnitString(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](h_.ID, objc.Sel("setUnitString:"), value)
}


// The molecular mass of blood glucose, typically used to create mole units for blood glucose.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkunitmolarmassbloodglucose
func (h_ HKUnit) HKUnitMolarMassBloodGlucose() float64 {
	rv := objc.Send[float64](h_.ID, objc.Sel("HKUnitMolarMassBloodGlucose"))
	return rv
}


// The molecular mass of blood glucose, typically used to create mole units for blood glucose.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkunitmolarmassbloodglucose
func (h_ HKUnit) SetHKUnitMolarMassBloodGlucose(value float64) {
	objc.Send[objc.ID](h_.ID, objc.Sel("setHKUnitMolarMassBloodGlucose:"), value)
}



