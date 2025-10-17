// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [UnitConcentrationMass] class.
var unitConcentrationMassClass = _UnitConcentrationMassClass{objc.GetClass("NSUnitConcentrationMass")}

type _UnitConcentrationMassClass struct {
	class objc.Class
}

// An interface definition for the [UnitConcentrationMass] class.
type IUnitConcentrationMass interface {
	IDimension
}

// A unit of measure for concentration of mass. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitConcentrationMass

type UnitConcentrationMass struct {
	Dimension
}

// UnitConcentrationMassFrom constructs a [UnitConcentrationMass] from an unsafe.Pointer.
//
// A unit of measure for concentration of mass.
func UnitConcentrationMassFrom(ptr unsafe.Pointer) UnitConcentrationMass {
	return UnitConcentrationMass{
		Dimension: DimensionFrom(ptr),
	}
}



