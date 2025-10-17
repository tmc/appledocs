// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [UnitPower] class.
var UnitPowerClass objc.Class

func init() {
	UnitPowerClass = objc.GetClass("NSUnitPower")
}

type UnitPower struct {
	objc.ID
}

func UnitPowerFrom(ptr unsafe.Pointer) UnitPower {
	return UnitPower{
		ID: objc.ID(ptr),
	}
}



