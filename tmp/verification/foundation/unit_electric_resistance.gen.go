// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

var unitElectricResistanceClass _UnitElectricResistanceClass

func init() {
	unitElectricResistanceClass = _UnitElectricResistanceClass{objc.GetClass("NSUnitElectricResistance")}
}

type _UnitElectricResistanceClass struct {
	class objc.Class
}

type UnitElectricResistance struct {
	objc.ID
}

func UnitElectricResistanceFrom(ptr unsafe.Pointer) UnitElectricResistance {
	return UnitElectricResistance{
		ID: objc.ID(ptr),
	}
}




