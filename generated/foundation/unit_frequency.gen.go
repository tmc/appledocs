// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [UnitFrequency] class.
var UnitFrequencyClass objc.Class

func init() {
	UnitFrequencyClass = objc.GetClass("NSUnitFrequency")
}

type UnitFrequency struct {
	objc.ID
}

func UnitFrequencyFrom(ptr unsafe.Pointer) UnitFrequency {
	return UnitFrequency{
		ID: objc.ID(ptr),
	}
}



