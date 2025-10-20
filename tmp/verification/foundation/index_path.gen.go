// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

var indexPathClass _IndexPathClass

func init() {
	indexPathClass = _IndexPathClass{objc.GetClass("NSIndexPath")}
}

type _IndexPathClass struct {
	class objc.Class
}

type IndexPath struct {
	objc.ID
}

func IndexPathFrom(ptr unsafe.Pointer) IndexPath {
	return IndexPath{
		ID: objc.ID(ptr),
	}
}




