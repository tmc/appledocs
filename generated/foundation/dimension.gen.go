// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [Dimension] class.
var dimensionClass = _DimensionClass{objc.GetClass("NSDimension")}

type _DimensionClass struct {
	class objc.Class
}

// An interface definition for the [Dimension] class.
type IDimension interface {
	IUnit
}

// An abstract class representing a dimensional unit of measure. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Dimension

type Dimension struct {
	Unit
}

// DimensionFrom constructs a [Dimension] from an unsafe.Pointer.
//
// An abstract class representing a dimensional unit of measure.
func DimensionFrom(ptr unsafe.Pointer) Dimension {
	return Dimension{
		Unit: UnitFrom(ptr),
	}
}



