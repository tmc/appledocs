// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

var unitLengthClass _UnitLengthClass

func init() {
	unitLengthClass = _UnitLengthClass{objc.GetClass("NSUnitLength")}
}

type _UnitLengthClass struct {
	class objc.Class
}

type UnitLength struct {
	objc.ID
}

func UnitLengthFrom(ptr unsafe.Pointer) UnitLength {
	return UnitLength{
		ID: objc.ID(ptr),
	}
}




