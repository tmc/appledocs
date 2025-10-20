// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

var UnitElectricChargeClass _UnitElectricChargeClass

func init() {
	UnitElectricChargeClass = _UnitElectricChargeClass{objc.GetClass("NSUnitElectricCharge")}
}

type _UnitElectricChargeClass struct {
	class objc.Class
}

type UnitElectricCharge struct {
	objc.ID
}

func UnitElectricChargeFrom(ptr unsafe.Pointer) UnitElectricCharge {
	return UnitElectricCharge{
		ID: objc.ID(ptr),
	}
}




