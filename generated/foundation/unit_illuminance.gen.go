// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [UnitIlluminance] class.
var UnitIlluminanceClass objc.Class

func init() {
	UnitIlluminanceClass = objc.GetClass("NSUnitIlluminance")
}

type UnitIlluminance struct {
	objc.ID
}

func UnitIlluminanceFrom(ptr unsafe.Pointer) UnitIlluminance {
	return UnitIlluminance{
		ID: objc.ID(ptr),
	}
}



