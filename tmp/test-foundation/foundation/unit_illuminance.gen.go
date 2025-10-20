// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

var UnitIlluminanceClass _UnitIlluminanceClass

func init() {
	UnitIlluminanceClass = _UnitIlluminanceClass{objc.GetClass("NSUnitIlluminance")}
}

type _UnitIlluminanceClass struct {
	class objc.Class
}

type UnitIlluminance struct {
	objc.ID
}

func UnitIlluminanceFrom(ptr unsafe.Pointer) UnitIlluminance {
	return UnitIlluminance{
		ID: objc.ID(ptr),
	}
}




