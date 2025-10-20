// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

var UnitConverterClass _UnitConverterClass

func init() {
	UnitConverterClass = _UnitConverterClass{objc.GetClass("NSUnitConverter")}
}

type _UnitConverterClass struct {
	class objc.Class
}

type UnitConverter struct {
	objc.ID
}

func UnitConverterFrom(ptr unsafe.Pointer) UnitConverter {
	return UnitConverter{
		ID: objc.ID(ptr),
	}
}




