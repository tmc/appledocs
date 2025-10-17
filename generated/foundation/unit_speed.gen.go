// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [UnitSpeed] class.
var UnitSpeedClass objc.Class

func init() {
	UnitSpeedClass = objc.GetClass("NSUnitSpeed")
}

type UnitSpeed struct {
	objc.ID
}

func UnitSpeedFrom(ptr unsafe.Pointer) UnitSpeed {
	return UnitSpeed{
		ID: objc.ID(ptr),
	}
}



