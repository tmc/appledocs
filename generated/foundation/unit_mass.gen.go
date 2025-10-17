// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [UnitMass] class.
var UnitMassClass objc.Class

func init() {
	UnitMassClass = objc.GetClass("NSUnitMass")
}

type UnitMass struct {
	objc.ID
}

func UnitMassFrom(ptr unsafe.Pointer) UnitMass {
	return UnitMass{
		ID: objc.ID(ptr),
	}
}



