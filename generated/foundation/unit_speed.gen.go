// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [UnitSpeed] class.
var unitSpeedClass = _UnitSpeedClass{objc.GetClass("NSUnitSpeed")}

type _UnitSpeedClass struct {
	class objc.Class
}

// An interface definition for the [UnitSpeed] class.
type IUnitSpeed interface {
	IDimension
}

// A unit of measure for speed. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitSpeed

type UnitSpeed struct {
	Dimension
}

// UnitSpeedFrom constructs a [UnitSpeed] from an unsafe.Pointer.
//
// A unit of measure for speed.
func UnitSpeedFrom(ptr unsafe.Pointer) UnitSpeed {
	return UnitSpeed{
		Dimension: DimensionFrom(ptr),
	}
}



