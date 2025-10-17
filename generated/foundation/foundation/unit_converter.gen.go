// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [UnitConverter] class.
var UnitConverterClass objc.Class

func init() {
	UnitConverterClass = objc.GetClass("NSUnitConverter")
}

type UnitConverter struct {
	objc.ID
}

func UnitConverterFrom(ptr unsafe.Pointer) UnitConverter {
	return UnitConverter{
		ID: objc.ID(ptr),
	}
}




