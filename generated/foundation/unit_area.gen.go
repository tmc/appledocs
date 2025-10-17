// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [UnitArea] class.
var unitAreaClass = _UnitAreaClass{objc.GetClass("NSUnitArea")}

type _UnitAreaClass struct {
	class objc.Class
}

// An interface definition for the [UnitArea] class.
type IUnitArea interface {
	IDimension
}

// A unit of measure for area. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitArea

type UnitArea struct {
	Dimension
}

// UnitAreaFrom constructs a [UnitArea] from an unsafe.Pointer.
//
// A unit of measure for area.
func UnitAreaFrom(ptr unsafe.Pointer) UnitArea {
	return UnitArea{
		Dimension: DimensionFrom(ptr),
	}
}



