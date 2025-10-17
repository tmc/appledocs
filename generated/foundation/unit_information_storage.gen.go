// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [UnitInformationStorage] class.
var unitInformationStorageClass = _UnitInformationStorageClass{objc.GetClass("NSUnitInformationStorage")}

type _UnitInformationStorageClass struct {
	class objc.Class
}

// A unit of measure for quantities of information. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitInformationStorage

type UnitInformationStorage struct {
	Dimension
}

// UnitInformationStorageFrom constructs a [UnitInformationStorage] from an unsafe.Pointer.
//
// A unit of measure for quantities of information.
func UnitInformationStorageFrom(ptr unsafe.Pointer) UnitInformationStorage {
	return UnitInformationStorage{
		Dimension: DimensionFrom(ptr),
	}
}



