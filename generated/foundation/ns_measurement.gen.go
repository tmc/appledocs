// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class NSMeasurement */


/* debug [class_header]: Header for NSMeasurement */
// The class instance for the [Measurement] class.
var (
	MeasurementClass     _MeasurementClass
	MeasurementClassOnce sync.Once
)

func getMeasurementClass() _MeasurementClass {
	MeasurementClassOnce.Do(func() {
		MeasurementClass = _MeasurementClass{objc.GetClass("NSMeasurement")}
	})
	return MeasurementClass
}

type _MeasurementClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for Measurement */
// An interface definition for the [Measurement] class.
type IMeasurement interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for Measurement */
	// properties:
	DoubleValue() float64
	Unit() objectivec.IObject
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for Measurement */
	// methods:
	MeasurementByAddingMeasurement(measurement unsafe.Pointer) unsafe.Pointer
	CanBeConvertedToUnit(unit IUnit) bool
	MeasurementByConvertingToUnit(unit IUnit) IMeasurement
	MeasurementBySubtractingMeasurement(measurement unsafe.Pointer) unsafe.Pointer
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for Measurement */
// Alloc allocates a new instance without initialization.
func (mc _MeasurementClass) Alloc() Measurement {
	rv := objc.Send[Measurement](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MeasurementClass) New() Measurement {
	rv := objc.Send[Measurement](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ Measurement) Init() Measurement {
	rv := objc.Send[Measurement](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ Measurement) Autorelease() Measurement {
	rv := objc.Send[Measurement](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMeasurement creates a new Measurement instance.
func NewMeasurement() Measurement {
	return getMeasurementClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for Measurement */
// A numeric quantity labeled with a unit of measure, with support for unit conversion and unit-aware calculations.
//
// Use this object in Swift when you need reference semantics or other Foundation-specific behavior. An object represents a quantity and unit of measure. The class provides a programmatic interface to converting measurements into different units, as well as calculating the sum or difference between two measurements. objects are initialized with an object and value. objects are immutable, and cannot be changed after being created. You can use the class to create localized string representations of objects.


// A numeric quantity labeled with a unit of measure, with support for unit conversion and unit-aware calculations.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMeasurement
type Measurement struct {
	objectivec.Object
}

// MeasurementFrom constructs a [Measurement] from an unsafe.Pointer.
//
// A numeric quantity labeled with a unit of measure, with support for unit conversion and unit-aware calculations.
func MeasurementFrom(ptr unsafe.Pointer) Measurement {
	return Measurement{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for Measurement */

// Initializes a new measurement with a specified double-precision floating-point value and unit.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMeasurement/init(doubleValue:unit:)
func NewMeasurementWithDoubleValueUnit(doubleValue float64, unit objectivec.IObject) Measurement {
	instance := getMeasurementClass().Alloc()
	rv := objc.Send[Measurement](instance.ID, objc.Sel("initWithDoubleValue:unit:"), doubleValue, unit)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewMeasurementWithDoubleValueUnit */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for Measurement */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for Measurement */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for Measurement */

// Returns a new measurement by adding the receiver to the specified measurement.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMeasurement/adding(_:)
func (m_ Measurement) MeasurementByAddingMeasurement(measurement unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("measurementByAddingMeasurement:"), measurement)
	return rv
}/* debug [instance_methods/method]: MeasurementByAddingMeasurement */


// Indicates whether the measurement can be converted to the given unit.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMeasurement/canBeConverted(to:)
func (m_ Measurement) CanBeConvertedToUnit(unit IUnit) bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("canBeConvertedToUnit:"), unit)
	return rv
}/* debug [instance_methods/method]: CanBeConvertedToUnit */


// Returns a measurement created by converting the receiver to the specified unit.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMeasurement/converting(to:)
func (m_ Measurement) MeasurementByConvertingToUnit(unit IUnit) IMeasurement {
	rv := objc.Send[Measurement](m_.ID, objc.Sel("measurementByConvertingToUnit:"), unit)
	return rv
}/* debug [instance_methods/method]: MeasurementByConvertingToUnit */


// Returns a new measurement by subtracting the specified measurement from the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMeasurement/subtracting(_:)
func (m_ Measurement) MeasurementBySubtractingMeasurement(measurement unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("measurementBySubtractingMeasurement:"), measurement)
	return rv
}/* debug [instance_methods/method]: MeasurementBySubtractingMeasurement */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for Measurement */

// The measurement value, represented as a double-precision floating-point number.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMeasurement/doubleValue
func (m_ Measurement) DoubleValue() float64 {
	rv := objc.Send[float64](m_.ID, objc.Sel("doubleValue"))
	return rv
}/* debug [instance_properties/getter]: doubleValue */


// The unit of measure.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMeasurement/unit
func (m_ Measurement) Unit() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](m_.ID, objc.Sel("unit"))
	return rv
}/* debug [instance_properties/getter]: unit */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class NSMeasurement */


