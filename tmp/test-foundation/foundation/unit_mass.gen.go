// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

var UnitMassClass _UnitMassClass

func init() {
	UnitMassClass = _UnitMassClass{objc.GetClass("NSUnitMass")}
}

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




