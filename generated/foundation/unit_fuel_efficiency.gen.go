// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [UnitFuelEfficiency] class.
var UnitFuelEfficiencyClass = _UnitFuelEfficiencyClass{objc.GetClass("NSUnitFuelEfficiency")}

type _UnitFuelEfficiencyClass struct {
	class objc.Class
}

type UnitFuelEfficiency struct {
	objc.ID
}

func UnitFuelEfficiencyFrom(ptr unsafe.Pointer) UnitFuelEfficiency {
	return UnitFuelEfficiency{
		ID: objc.ID(ptr),
	}
}




