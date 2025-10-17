// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [MapTable] class.
var MapTableClass objc.Class

func init() {
	MapTableClass = objc.GetClass("NSMapTable")
}

type MapTable struct {
	objc.ID
}

func MapTableFrom(ptr unsafe.Pointer) MapTable {
	return MapTable{
		ID: objc.ID(ptr),
	}
}



