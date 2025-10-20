// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

var UnitVolumeClass _UnitVolumeClass

func init() {
	UnitVolumeClass = _UnitVolumeClass{objc.GetClass("NSUnitVolume")}
}

type _UnitVolumeClass struct {
	class objc.Class
}

type UnitVolume struct {
	objc.ID
}

func UnitVolumeFrom(ptr unsafe.Pointer) UnitVolume {
	return UnitVolume{
		ID: objc.ID(ptr),
	}
}




