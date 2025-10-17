// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [UnitElectricPotentialDifference] class.
var UnitElectricPotentialDifferenceClass objc.Class

func init() {
	UnitElectricPotentialDifferenceClass = objc.GetClass("NSUnitElectricPotentialDifference")
}

type UnitElectricPotentialDifference struct {
	objc.ID
}

func UnitElectricPotentialDifferenceFrom(ptr unsafe.Pointer) UnitElectricPotentialDifference {
	return UnitElectricPotentialDifference{
		ID: objc.ID(ptr),
	}
}



