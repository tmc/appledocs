// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [UnitDispersion] class.
var unitDispersionClass = _UnitDispersionClass{objc.GetClass("NSUnitDispersion")}

type _UnitDispersionClass struct {
	class objc.Class
}

// An interface definition for the [UnitDispersion] class.
type IUnitDispersion interface {
	IDimension
}

// A unit of measure for specific quantities of dispersion. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitDispersion

type UnitDispersion struct {
	Dimension
}

// UnitDispersionFrom constructs a [UnitDispersion] from an unsafe.Pointer.
//
// A unit of measure for specific quantities of dispersion.
func UnitDispersionFrom(ptr unsafe.Pointer) UnitDispersion {
	return UnitDispersion{
		Dimension: DimensionFrom(ptr),
	}
}



