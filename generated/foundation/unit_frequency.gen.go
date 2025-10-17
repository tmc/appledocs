// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [UnitFrequency] class.
var UnitFrequencyClass = _UnitFrequencyClass{objc.GetClass("NSUnitFrequency")}

type _UnitFrequencyClass struct {
	class objc.Class
}

type UnitFrequency struct {
	objc.ID
}

func UnitFrequencyFrom(ptr unsafe.Pointer) UnitFrequency {
	return UnitFrequency{
		ID: objc.ID(ptr),
	}
}




