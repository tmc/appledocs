// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [UnitLength] class.
var UnitLengthClass objc.Class

func init() {
	UnitLengthClass = objc.GetClass("NSUnitLength")
}

type UnitLength struct {
	objc.ID
}

func UnitLengthFrom(ptr unsafe.Pointer) UnitLength {
	return UnitLength{
		ID: objc.ID(ptr),
	}
}




