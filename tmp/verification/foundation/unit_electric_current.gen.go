// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

var unitElectricCurrentClass _UnitElectricCurrentClass

func init() {
	unitElectricCurrentClass = _UnitElectricCurrentClass{objc.GetClass("NSUnitElectricCurrent")}
}

type _UnitElectricCurrentClass struct {
	class objc.Class
}

type UnitElectricCurrent struct {
	objc.ID
}

func UnitElectricCurrentFrom(ptr unsafe.Pointer) UnitElectricCurrent {
	return UnitElectricCurrent{
		ID: objc.ID(ptr),
	}
}




