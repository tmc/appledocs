// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

var UnitTemperatureClass _UnitTemperatureClass

func init() {
	UnitTemperatureClass = _UnitTemperatureClass{objc.GetClass("NSUnitTemperature")}
}

type _UnitTemperatureClass struct {
	class objc.Class
}

type UnitTemperature struct {
	objc.ID
}

func UnitTemperatureFrom(ptr unsafe.Pointer) UnitTemperature {
	return UnitTemperature{
		ID: objc.ID(ptr),
	}
}




