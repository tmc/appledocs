// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [UnitElectricCharge] class.
var UnitElectricChargeClass objc.Class

func init() {
	UnitElectricChargeClass = objc.GetClass("NSUnitElectricCharge")
}

type UnitElectricCharge struct {
	objc.ID
}

func UnitElectricChargeFrom(ptr unsafe.Pointer) UnitElectricCharge {
	return UnitElectricCharge{
		ID: objc.ID(ptr),
	}
}



