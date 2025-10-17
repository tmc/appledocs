// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [UnitArea] class.
var UnitAreaClass objc.Class

func init() {
	UnitAreaClass = objc.GetClass("NSUnitArea")
}

type UnitArea struct {
	objc.ID
}

func UnitAreaFrom(ptr unsafe.Pointer) UnitArea {
	return UnitArea{
		ID: objc.ID(ptr),
	}
}



