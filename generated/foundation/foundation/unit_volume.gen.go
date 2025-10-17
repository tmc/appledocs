// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [UnitVolume] class.
var UnitVolumeClass objc.Class

func init() {
	UnitVolumeClass = objc.GetClass("NSUnitVolume")
}

type UnitVolume struct {
	objc.ID
}

func UnitVolumeFrom(ptr unsafe.Pointer) UnitVolume {
	return UnitVolume{
		ID: objc.ID(ptr),
	}
}




