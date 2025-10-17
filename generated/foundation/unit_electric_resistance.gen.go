// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [UnitElectricResistance] class.
var UnitElectricResistanceClass objc.Class

func init() {
	UnitElectricResistanceClass = objc.GetClass("NSUnitElectricResistance")
}

type UnitElectricResistance struct {
	objc.ID
}

func UnitElectricResistanceFrom(ptr unsafe.Pointer) UnitElectricResistance {
	return UnitElectricResistance{
		ID: objc.ID(ptr),
	}
}



