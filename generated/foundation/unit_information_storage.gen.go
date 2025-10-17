// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [UnitInformationStorage] class.
var UnitInformationStorageClass objc.Class

func init() {
	UnitInformationStorageClass = objc.GetClass("NSUnitInformationStorage")
}

type UnitInformationStorage struct {
	objc.ID
}

func UnitInformationStorageFrom(ptr unsafe.Pointer) UnitInformationStorage {
	return UnitInformationStorage{
		ID: objc.ID(ptr),
	}
}



