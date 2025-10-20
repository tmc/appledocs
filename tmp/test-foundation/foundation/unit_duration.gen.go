// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

var UnitDurationClass _UnitDurationClass

func init() {
	UnitDurationClass = _UnitDurationClass{objc.GetClass("NSUnitDuration")}
}

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




