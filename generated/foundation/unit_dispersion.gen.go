// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [UnitDispersion] class.
var UnitDispersionClass objc.Class

func init() {
	UnitDispersionClass = objc.GetClass("NSUnitDispersion")
}

type UnitDispersion struct {
	objc.ID
}

func UnitDispersionFrom(ptr unsafe.Pointer) UnitDispersion {
	return UnitDispersion{
		ID: objc.ID(ptr),
	}
}



