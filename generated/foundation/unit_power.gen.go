// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [UnitPower] class.
var unitPowerClass = _UnitPowerClass{objc.GetClass("NSUnitPower")}

type _UnitPowerClass struct {
	class objc.Class
}

// An interface definition for the [UnitPower] class.
type IUnitPower interface {
	IDimension
}

// A unit of measure for power. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitPower

type UnitPower struct {
	Dimension
}

// UnitPowerFrom constructs a [UnitPower] from an unsafe.Pointer.
//
// A unit of measure for power.
func UnitPowerFrom(ptr unsafe.Pointer) UnitPower {
	return UnitPower{
		Dimension: DimensionFrom(ptr),
	}
}



