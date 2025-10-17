// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [UnitElectricCharge] class.
var unitElectricChargeClass = _UnitElectricChargeClass{objc.GetClass("NSUnitElectricCharge")}

type _UnitElectricChargeClass struct {
	class objc.Class
}

// A unit of measure for electric charge. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitElectricCharge

type UnitElectricCharge struct {
	Dimension
}

// UnitElectricChargeFrom constructs a [UnitElectricCharge] from an unsafe.Pointer.
//
// A unit of measure for electric charge.
func UnitElectricChargeFrom(ptr unsafe.Pointer) UnitElectricCharge {
	return UnitElectricCharge{
		Dimension: DimensionFrom(ptr),
	}
}



