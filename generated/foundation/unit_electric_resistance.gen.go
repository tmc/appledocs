// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [UnitElectricResistance] class.
var unitElectricResistanceClass = _UnitElectricResistanceClass{objc.GetClass("NSUnitElectricResistance")}

type _UnitElectricResistanceClass struct {
	class objc.Class
}

// An interface definition for the [UnitElectricResistance] class.
type IUnitElectricResistance interface {
	IDimension
}

// A unit of measure for electric resistance. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitElectricResistance

type UnitElectricResistance struct {
	Dimension
}

// UnitElectricResistanceFrom constructs a [UnitElectricResistance] from an unsafe.Pointer.
//
// A unit of measure for electric resistance.
func UnitElectricResistanceFrom(ptr unsafe.Pointer) UnitElectricResistance {
	return UnitElectricResistance{
		Dimension: DimensionFrom(ptr),
	}
}



