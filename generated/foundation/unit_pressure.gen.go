// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [UnitPressure] class.
var unitPressureClass = _UnitPressureClass{objc.GetClass("NSUnitPressure")}

type _UnitPressureClass struct {
	class objc.Class
}

// A unit of measure for pressure. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitPressure

type UnitPressure struct {
	Dimension
}

// UnitPressureFrom constructs a [UnitPressure] from an unsafe.Pointer.
//
// A unit of measure for pressure.
func UnitPressureFrom(ptr unsafe.Pointer) UnitPressure {
	return UnitPressure{
		Dimension: DimensionFrom(ptr),
	}
}



