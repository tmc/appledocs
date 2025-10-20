// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

var unitPowerClass _UnitPowerClass

func init() {
	unitPowerClass = _UnitPowerClass{objc.GetClass("NSUnitPower")}
}

type _UnitPowerClass struct {
	class objc.Class
}

type UnitPower struct {
	objc.ID
}

func UnitPowerFrom(ptr unsafe.Pointer) UnitPower {
	return UnitPower{
		ID: objc.ID(ptr),
	}
}




