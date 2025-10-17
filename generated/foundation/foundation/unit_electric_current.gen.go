// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [UnitElectricCurrent] class.
var UnitElectricCurrentClass objc.Class

func init() {
	UnitElectricCurrentClass = objc.GetClass("NSUnitElectricCurrent")
}

type UnitElectricCurrent struct {
	objc.ID
}

func UnitElectricCurrentFrom(ptr unsafe.Pointer) UnitElectricCurrent {
	return UnitElectricCurrent{
		ID: objc.ID(ptr),
	}
}




