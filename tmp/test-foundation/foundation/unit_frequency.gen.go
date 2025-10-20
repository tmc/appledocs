// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

var UnitFrequencyClass _UnitFrequencyClass

func init() {
	UnitFrequencyClass = _UnitFrequencyClass{objc.GetClass("NSUnitFrequency")}
}

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




