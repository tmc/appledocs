// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [UnitElectricPotentialDifference] class.
var unitElectricPotentialDifferenceClass = _UnitElectricPotentialDifferenceClass{objc.GetClass("NSUnitElectricPotentialDifference")}

type _UnitElectricPotentialDifferenceClass struct {
	class objc.Class
}

// An interface definition for the [UnitElectricPotentialDifference] class.
type IUnitElectricPotentialDifference interface {
	IDimension
}

// A unit of measure for electric potential difference. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitElectricPotentialDifference

type UnitElectricPotentialDifference struct {
	Dimension
}

// UnitElectricPotentialDifferenceFrom constructs a [UnitElectricPotentialDifference] from an unsafe.Pointer.
//
// A unit of measure for electric potential difference.
func UnitElectricPotentialDifferenceFrom(ptr unsafe.Pointer) UnitElectricPotentialDifference {
	return UnitElectricPotentialDifference{
		Dimension: DimensionFrom(ptr),
	}
}



