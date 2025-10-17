// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [UnitFrequency] class.
var unitFrequencyClass = _UnitFrequencyClass{objc.GetClass("NSUnitFrequency")}

type _UnitFrequencyClass struct {
	class objc.Class
}

// An interface definition for the [UnitFrequency] class.
type IUnitFrequency interface {
	IDimension
}

// A unit of measure for frequency. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitFrequency

type UnitFrequency struct {
	Dimension
}

// UnitFrequencyFrom constructs a [UnitFrequency] from an unsafe.Pointer.
//
// A unit of measure for frequency.
func UnitFrequencyFrom(ptr unsafe.Pointer) UnitFrequency {
	return UnitFrequency{
		Dimension: DimensionFrom(ptr),
	}
}



