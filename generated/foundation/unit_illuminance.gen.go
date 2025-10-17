// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [UnitIlluminance] class.
var unitIlluminanceClass = _UnitIlluminanceClass{objc.GetClass("NSUnitIlluminance")}

type _UnitIlluminanceClass struct {
	class objc.Class
}

// A unit of measure for illuminance. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitIlluminance

type UnitIlluminance struct {
	Dimension
}

// UnitIlluminanceFrom constructs a [UnitIlluminance] from an unsafe.Pointer.
//
// A unit of measure for illuminance.
func UnitIlluminanceFrom(ptr unsafe.Pointer) UnitIlluminance {
	return UnitIlluminance{
		Dimension: DimensionFrom(ptr),
	}
}



