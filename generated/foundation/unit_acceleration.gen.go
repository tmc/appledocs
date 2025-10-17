// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [UnitAcceleration] class.
var UnitAccelerationClass objc.Class

func init() {
	UnitAccelerationClass = objc.GetClass("NSUnitAcceleration")
}

type UnitAcceleration struct {
	objc.ID
}

func UnitAccelerationFrom(ptr unsafe.Pointer) UnitAcceleration {
	return UnitAcceleration{
		ID: objc.ID(ptr),
	}
}



