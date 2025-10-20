// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

var unitPressureClass _UnitPressureClass

func init() {
	unitPressureClass = _UnitPressureClass{objc.GetClass("NSUnitPressure")}
}

type _UnitPressureClass struct {
	class objc.Class
}

type UnitPressure struct {
	objc.ID
}

func UnitPressureFrom(ptr unsafe.Pointer) UnitPressure {
	return UnitPressure{
		ID: objc.ID(ptr),
	}
}




