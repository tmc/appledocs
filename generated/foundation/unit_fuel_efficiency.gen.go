// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [UnitFuelEfficiency] class.
var unitFuelEfficiencyClass = _UnitFuelEfficiencyClass{objc.GetClass("NSUnitFuelEfficiency")}

type _UnitFuelEfficiencyClass struct {
	class objc.Class
}

// A unit of measure for fuel efficiency. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitFuelEfficiency

type UnitFuelEfficiency struct {
	Dimension
}

// UnitFuelEfficiencyFrom constructs a [UnitFuelEfficiency] from an unsafe.Pointer.
//
// A unit of measure for fuel efficiency.
func UnitFuelEfficiencyFrom(ptr unsafe.Pointer) UnitFuelEfficiency {
	return UnitFuelEfficiency{
		Dimension: DimensionFrom(ptr),
	}
}



