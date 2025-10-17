// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [UnitEnergy] class.
var unitEnergyClass = _UnitEnergyClass{objc.GetClass("NSUnitEnergy")}

type _UnitEnergyClass struct {
	class objc.Class
}

// An interface definition for the [UnitEnergy] class.
type IUnitEnergy interface {
	IDimension
}

// A unit of measure for energy. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitEnergy

type UnitEnergy struct {
	Dimension
}

// UnitEnergyFrom constructs a [UnitEnergy] from an unsafe.Pointer.
//
// A unit of measure for energy.
func UnitEnergyFrom(ptr unsafe.Pointer) UnitEnergy {
	return UnitEnergy{
		Dimension: DimensionFrom(ptr),
	}
}



