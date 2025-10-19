// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [Measurement] class.
var measurementClass = _MeasurementClass{objc.GetClass("NSMeasurement")}

type _MeasurementClass struct {
	class objc.Class
}

// An interface definition for the [Measurement] class.
type IMeasurement interface {
	objectivec.IObject
}

// A numeric quantity labeled with a unit of measure, with support for unit conversion and unit-aware calculations. [Full Topic]
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

// New creates and returns a new instance with a +1 retain count.
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
	return measurementClass.New()
}




