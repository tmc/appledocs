// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [UnitDuration] class.
var UnitDurationClass objc.Class

func init() {
	UnitDurationClass = objc.GetClass("NSUnitDuration")
}

type UnitDuration struct {
	objc.ID
}

func UnitDurationFrom(ptr unsafe.Pointer) UnitDuration {
	return UnitDuration{
		ID: objc.ID(ptr),
	}
}




