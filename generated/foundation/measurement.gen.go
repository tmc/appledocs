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



