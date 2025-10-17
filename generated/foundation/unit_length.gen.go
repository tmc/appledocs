// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [UnitLength] class.
var unitLengthClass = _UnitLengthClass{objc.GetClass("NSUnitLength")}

type _UnitLengthClass struct {
	class objc.Class
}

// A unit of measure for length. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitLength

type UnitLength struct {
	Dimension
}

// UnitLengthFrom constructs a [UnitLength] from an unsafe.Pointer.
//
// A unit of measure for length.
func UnitLengthFrom(ptr unsafe.Pointer) UnitLength {
	return UnitLength{
		Dimension: DimensionFrom(ptr),
	}
}



