// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [UnitPressure] class.
var UnitPressureClass objc.Class

func init() {
	UnitPressureClass = objc.GetClass("NSUnitPressure")
}

type UnitPressure struct {
	objc.ID
}

func UnitPressureFrom(ptr unsafe.Pointer) UnitPressure {
	return UnitPressure{
		ID: objc.ID(ptr),
	}
}




