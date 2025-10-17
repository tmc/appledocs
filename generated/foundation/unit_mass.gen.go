// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [UnitMass] class.
var UnitMassClass = _UnitMassClass{objc.GetClass("NSUnitMass")}

type _UnitMassClass struct {
	class objc.Class
}

type UnitMass struct {
	objc.ID
}

func UnitMassFrom(ptr unsafe.Pointer) UnitMass {
	return UnitMass{
		ID: objc.ID(ptr),
	}
}




