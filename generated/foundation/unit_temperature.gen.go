// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [UnitTemperature] class.
var UnitTemperatureClass objc.Class

func init() {
	UnitTemperatureClass = objc.GetClass("NSUnitTemperature")
}

type UnitTemperature struct {
	objc.ID
}

func UnitTemperatureFrom(ptr unsafe.Pointer) UnitTemperature {
	return UnitTemperature{
		ID: objc.ID(ptr),
	}
}



