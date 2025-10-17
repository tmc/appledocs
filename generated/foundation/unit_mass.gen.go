// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [UnitMass] class.
var unitMassClass = _UnitMassClass{objc.GetClass("NSUnitMass")}

type _UnitMassClass struct {
	class objc.Class
}

// A unit of measure for mass. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitMass

type UnitMass struct {
	Dimension
}

// UnitMassFrom constructs a [UnitMass] from an unsafe.Pointer.
//
// A unit of measure for mass.
func UnitMassFrom(ptr unsafe.Pointer) UnitMass {
	return UnitMass{
		Dimension: DimensionFrom(ptr),
	}
}



