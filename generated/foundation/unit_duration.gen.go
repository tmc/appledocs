// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [UnitDuration] class.
var UnitDurationClass = _UnitDurationClass{objc.GetClass("NSUnitDuration")}

type _UnitDurationClass struct {
	class objc.Class
}

type UnitDuration struct {
	objc.ID
}

func UnitDurationFrom(ptr unsafe.Pointer) UnitDuration {
	return UnitDuration{
		ID: objc.ID(ptr),
	}
}




