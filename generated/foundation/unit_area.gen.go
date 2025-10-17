// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [UnitArea] class.
var UnitAreaClass = _UnitAreaClass{objc.GetClass("NSUnitArea")}

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




