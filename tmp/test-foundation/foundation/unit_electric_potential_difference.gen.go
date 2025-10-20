// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

var UnitElectricPotentialDifferenceClass _UnitElectricPotentialDifferenceClass

func init() {
	UnitElectricPotentialDifferenceClass = _UnitElectricPotentialDifferenceClass{objc.GetClass("NSUnitElectricPotentialDifference")}
}

type _UnitElectricPotentialDifferenceClass struct {
	class objc.Class
}

type UnitElectricPotentialDifference struct {
	objc.ID
}

func UnitElectricPotentialDifferenceFrom(ptr unsafe.Pointer) UnitElectricPotentialDifference {
	return UnitElectricPotentialDifference{
		ID: objc.ID(ptr),
	}
}




