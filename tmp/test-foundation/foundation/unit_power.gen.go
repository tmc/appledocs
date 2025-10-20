// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

var UnitPowerClass _UnitPowerClass

func init() {
	UnitPowerClass = _UnitPowerClass{objc.GetClass("NSUnitPower")}
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




