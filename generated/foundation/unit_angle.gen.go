// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [UnitAngle] class.
var unitAngleClass = _UnitAngleClass{objc.GetClass("NSUnitAngle")}

type _UnitAngleClass struct {
	class objc.Class
}

// A unit of measure for planar angle and rotation. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitAngle

type UnitAngle struct {
	Dimension
}

// UnitAngleFrom constructs a [UnitAngle] from an unsafe.Pointer.
//
// A unit of measure for planar angle and rotation.
func UnitAngleFrom(ptr unsafe.Pointer) UnitAngle {
	return UnitAngle{
		Dimension: DimensionFrom(ptr),
	}
}



