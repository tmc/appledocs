// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [UnitElectricCurrent] class.
var unitElectricCurrentClass = _UnitElectricCurrentClass{objc.GetClass("NSUnitElectricCurrent")}

type _UnitElectricCurrentClass struct {
	class objc.Class
}

// An interface definition for the [UnitElectricCurrent] class.
type IUnitElectricCurrent interface {
	IDimension
}

// A unit of measure for electric current. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitElectricCurrent

type UnitElectricCurrent struct {
	Dimension
}

// UnitElectricCurrentFrom constructs a [UnitElectricCurrent] from an unsafe.Pointer.
//
// A unit of measure for electric current.
func UnitElectricCurrentFrom(ptr unsafe.Pointer) UnitElectricCurrent {
	return UnitElectricCurrent{
		Dimension: DimensionFrom(ptr),
	}
}



