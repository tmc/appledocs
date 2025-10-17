// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [UnitTemperature] class.
var unitTemperatureClass = _UnitTemperatureClass{objc.GetClass("NSUnitTemperature")}

type _UnitTemperatureClass struct {
	class objc.Class
}

// An interface definition for the [UnitTemperature] class.
type IUnitTemperature interface {
	IDimension
}

// A unit of measure for temperature. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitTemperature

type UnitTemperature struct {
	Dimension
}

// UnitTemperatureFrom constructs a [UnitTemperature] from an unsafe.Pointer.
//
// A unit of measure for temperature.
func UnitTemperatureFrom(ptr unsafe.Pointer) UnitTemperature {
	return UnitTemperature{
		Dimension: DimensionFrom(ptr),
	}
}



