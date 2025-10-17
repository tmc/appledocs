// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [UnitConcentrationMass] class.
var UnitConcentrationMassClass objc.Class

func init() {
	UnitConcentrationMassClass = objc.GetClass("NSUnitConcentrationMass")
}

type UnitConcentrationMass struct {
	objc.ID
}

func UnitConcentrationMassFrom(ptr unsafe.Pointer) UnitConcentrationMass {
	return UnitConcentrationMass{
		ID: objc.ID(ptr),
	}
}




