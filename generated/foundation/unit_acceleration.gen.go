// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [UnitAcceleration] class.
var unitAccelerationClass = _UnitAccelerationClass{objc.GetClass("NSUnitAcceleration")}

type _UnitAccelerationClass struct {
	class objc.Class
}

// A unit of measure for acceleration. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitAcceleration

type UnitAcceleration struct {
	Dimension
}

// UnitAccelerationFrom constructs a [UnitAcceleration] from an unsafe.Pointer.
//
// A unit of measure for acceleration.
func UnitAccelerationFrom(ptr unsafe.Pointer) UnitAcceleration {
	return UnitAcceleration{
		Dimension: DimensionFrom(ptr),
	}
}



