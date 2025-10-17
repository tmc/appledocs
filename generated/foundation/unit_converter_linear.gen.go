// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [UnitConverterLinear] class.
var UnitConverterLinearClass objc.Class

func init() {
	UnitConverterLinearClass = objc.GetClass("NSUnitConverterLinear")
}

type UnitConverterLinear struct {
	objc.ID
}

func UnitConverterLinearFrom(ptr unsafe.Pointer) UnitConverterLinear {
	return UnitConverterLinear{
		ID: objc.ID(ptr),
	}
}



