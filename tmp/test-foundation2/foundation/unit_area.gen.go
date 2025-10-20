// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

var unitAreaClass _UnitAreaClass

func init() {
	unitAreaClass = _UnitAreaClass{objc.GetClass("NSUnitArea")}
}

type _UnitAreaClass struct {
	class objc.Class
}

type UnitArea struct {
	objc.ID
}

func UnitAreaFrom(ptr unsafe.Pointer) UnitArea {
	return UnitArea{
		ID: objc.ID(ptr),
	}
}




