// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [UnitVolume] class.
var unitVolumeClass = _UnitVolumeClass{objc.GetClass("NSUnitVolume")}

type _UnitVolumeClass struct {
	class objc.Class
}

// An interface definition for the [UnitVolume] class.
type IUnitVolume interface {
	IDimension
}

// A unit of measure for volume. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitVolume

type UnitVolume struct {
	Dimension
}

// UnitVolumeFrom constructs a [UnitVolume] from an unsafe.Pointer.
//
// A unit of measure for volume.
func UnitVolumeFrom(ptr unsafe.Pointer) UnitVolume {
	return UnitVolume{
		Dimension: DimensionFrom(ptr),
	}
}



