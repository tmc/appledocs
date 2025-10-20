// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

var UnitAngleClass _UnitAngleClass

func init() {
	UnitAngleClass = _UnitAngleClass{objc.GetClass("NSUnitAngle")}
}

type _UnitAngleClass struct {
	class objc.Class
}

type UnitAngle struct {
	objc.ID
}

func UnitAngleFrom(ptr unsafe.Pointer) UnitAngle {
	return UnitAngle{
		ID: objc.ID(ptr),
	}
}




