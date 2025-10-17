// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [PurgeableData] class.
var PurgeableDataClass objc.Class

func init() {
	PurgeableDataClass = objc.GetClass("NSPurgeableData")
}

type PurgeableData struct {
	objc.ID
}

func PurgeableDataFrom(ptr unsafe.Pointer) PurgeableData {
	return PurgeableData{
		ID: objc.ID(ptr),
	}
}




