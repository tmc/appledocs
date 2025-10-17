// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [PurgeableData] class.
var PurgeableDataClass = _PurgeableDataClass{objc.GetClass("NSPurgeableData")}

type _PurgeableDataClass struct {
	class objc.Class
}

type PurgeableData struct {
	objc.ID
}

func PurgeableDataFrom(ptr unsafe.Pointer) PurgeableData {
	return PurgeableData{
		ID: objc.ID(ptr),
	}
}




