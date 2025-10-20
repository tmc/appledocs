// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

var unitAccelerationClass _UnitAccelerationClass

func init() {
	unitAccelerationClass = _UnitAccelerationClass{objc.GetClass("NSUnitAcceleration")}
}

type _UnitAccelerationClass struct {
	class objc.Class
}

type UnitAcceleration struct {
	objc.ID
}

func UnitAccelerationFrom(ptr unsafe.Pointer) UnitAcceleration {
	return UnitAcceleration{
		ID: objc.ID(ptr),
	}
}




