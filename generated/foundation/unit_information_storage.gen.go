// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [UnitInformationStorage] class.
var UnitInformationStorageClass = _UnitInformationStorageClass{objc.GetClass("NSUnitInformationStorage")}

type _UnitInformationStorageClass struct {
	class objc.Class
}

type UnitInformationStorage struct {
	objc.ID
}

func UnitInformationStorageFrom(ptr unsafe.Pointer) UnitInformationStorage {
	return UnitInformationStorage{
		ID: objc.ID(ptr),
	}
}




