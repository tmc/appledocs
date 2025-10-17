// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [UnitEnergy] class.
var UnitEnergyClass objc.Class

func init() {
	UnitEnergyClass = objc.GetClass("NSUnitEnergy")
}

type UnitEnergy struct {
	objc.ID
}

func UnitEnergyFrom(ptr unsafe.Pointer) UnitEnergy {
	return UnitEnergy{
		ID: objc.ID(ptr),
	}
}




