// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

var MapTableClass _MapTableClass

func init() {
	MapTableClass = _MapTableClass{objc.GetClass("NSMapTable")}
}

type _MapTableClass struct {
	class objc.Class
}

type MapTable struct {
	objc.ID
}

func MapTableFrom(ptr unsafe.Pointer) MapTable {
	return MapTable{
		ID: objc.ID(ptr),
	}
}




