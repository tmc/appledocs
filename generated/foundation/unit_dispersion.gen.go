// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [UnitDispersion] class.
var UnitDispersionClass = _UnitDispersionClass{objc.GetClass("NSUnitDispersion")}

type _UnitDispersionClass struct {
	class objc.Class
}

type UnitDispersion struct {
	objc.ID
}

func UnitDispersionFrom(ptr unsafe.Pointer) UnitDispersion {
	return UnitDispersion{
		ID: objc.ID(ptr),
	}
}




