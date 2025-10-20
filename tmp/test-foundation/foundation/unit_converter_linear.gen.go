// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

var UnitConverterLinearClass _UnitConverterLinearClass

func init() {
	UnitConverterLinearClass = _UnitConverterLinearClass{objc.GetClass("NSUnitConverterLinear")}
}

type _UnitConverterLinearClass struct {
	class objc.Class
}

type UnitConverterLinear struct {
	objc.ID
}

func UnitConverterLinearFrom(ptr unsafe.Pointer) UnitConverterLinear {
	return UnitConverterLinear{
		ID: objc.ID(ptr),
	}
}




