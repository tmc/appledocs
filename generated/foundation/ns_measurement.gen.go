// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

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

// An interface definition for the [Measurement] class.
type IMeasurement interface {
	objectivec.IObject
	CanBeConvertedToUnit(unit IUnit) bool
}

// A numeric quantity labeled with a unit of measure, with support for unit conversion and unit-aware calculations.
//
// Use this object in Swift when you need reference semantics or other Foundation-specific behavior. An object represents a quantity and unit of measure. The class provides a programmatic interface to converting measurements into different units, as well as calculating the sum or difference between two measurements. objects are initialized with an object and value. objects are immutable, and cannot be changed after being created. You can use the class to create localized string representations of objects.
//
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

// Alloc allocates a new instance without initialization.
func (mc _MeasurementClass) Alloc() Measurement {
	rv := objc.Send[Measurement](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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




// Initializes a new measurement with a specified double-precision floating-point value and unit.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMeasurement/init(doubleValue:unit:)
func NewMeasurementWithDoubleValueUnit(doubleValue float64, unit unsafe.Pointer) Measurement {
	instance := getMeasurementClass().Alloc()
	rv := objc.Send[Measurement](instance.ID, objc.Sel("initWithDoubleValue:unit:"), doubleValue, unit)
	rv.Autorelease()
	return rv
}


// Indicates whether the measurement can be converted to the given unit.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMeasurement/canBeConverted(to:)
func (m_ Measurement) CanBeConvertedToUnit(unit IUnit) bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("canBeConvertedToUnit:"), unit)
	return rv
}

// The unit of measure.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMeasurement/unit
func (m_ Measurement) Unit() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("unit"))
	return rv
}

// The measurement value, represented as a double-precision floating-point number.
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmeasurement/doublevalue
func (m_ Measurement) DoubleValue() float64 {
	rv := objc.Send[float64](m_.ID, objc.Sel("doubleValue"))
	return rv
}


// SetDoubleValue sets the value of the doubleValue property.
// The measurement value, represented as a double-precision floating-point number.

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmeasurement/doublevalue
func (m_ Measurement) SetDoubleValue(value float64) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setDoubleValue:"), value)
}


