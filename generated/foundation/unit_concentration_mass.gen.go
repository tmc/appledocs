// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [UnitConcentrationMass] class.
var UnitConcentrationMassClass = _UnitConcentrationMassClass{objc.GetClass("NSUnitConcentrationMass")}

type _UnitConcentrationMassClass struct {
	class objc.Class
}

type UnitConcentrationMass struct {
	objc.ID
}

func UnitConcentrationMassFrom(ptr unsafe.Pointer) UnitConcentrationMass {
	return UnitConcentrationMass{
		ID: objc.ID(ptr),
	}
}




