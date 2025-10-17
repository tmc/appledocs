// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [UnitAngle] class.
var UnitAngleClass objc.Class

func init() {
	UnitAngleClass = objc.GetClass("NSUnitAngle")
}

type UnitAngle struct {
	objc.ID
}

func UnitAngleFrom(ptr unsafe.Pointer) UnitAngle {
	return UnitAngle{
		ID: objc.ID(ptr),
	}
}



