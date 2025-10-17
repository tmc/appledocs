// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [UnitFuelEfficiency] class.
var UnitFuelEfficiencyClass objc.Class

func init() {
	UnitFuelEfficiencyClass = objc.GetClass("NSUnitFuelEfficiency")
}

type UnitFuelEfficiency struct {
	objc.ID
}

func UnitFuelEfficiencyFrom(ptr unsafe.Pointer) UnitFuelEfficiency {
	return UnitFuelEfficiency{
		ID: objc.ID(ptr),
	}
}




