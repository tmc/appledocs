// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [UnitDuration] class.
var unitDurationClass = _UnitDurationClass{objc.GetClass("NSUnitDuration")}

type _UnitDurationClass struct {
	class objc.Class
}

// A unit of measure for a duration of time. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitDuration

type UnitDuration struct {
	Dimension
}

// UnitDurationFrom constructs a [UnitDuration] from an unsafe.Pointer.
//
// A unit of measure for a duration of time.
func UnitDurationFrom(ptr unsafe.Pointer) UnitDuration {
	return UnitDuration{
		Dimension: DimensionFrom(ptr),
	}
}



