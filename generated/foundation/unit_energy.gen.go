// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [UnitEnergy] class.
var UnitEnergyClass = _UnitEnergyClass{objc.GetClass("NSUnitEnergy")}

type _UnitEnergyClass struct {
	class objc.Class
}

type UnitEnergy struct {
	objc.ID
}

func UnitEnergyFrom(ptr unsafe.Pointer) UnitEnergy {
	return UnitEnergy{
		ID: objc.ID(ptr),
	}
}




