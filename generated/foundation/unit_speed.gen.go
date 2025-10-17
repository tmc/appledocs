// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [UnitSpeed] class.
var UnitSpeedClass = _UnitSpeedClass{objc.GetClass("NSUnitSpeed")}

type _UnitSpeedClass struct {
	class objc.Class
}

type UnitSpeed struct {
	objc.ID
}

func UnitSpeedFrom(ptr unsafe.Pointer) UnitSpeed {
	return UnitSpeed{
		ID: objc.ID(ptr),
	}
}




