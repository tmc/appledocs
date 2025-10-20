// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

var IndexSetClass _IndexSetClass

func init() {
	IndexSetClass = _IndexSetClass{objc.GetClass("NSIndexSet")}
}

type _IndexSetClass struct {
	class objc.Class
}

type IndexSet struct {
	objc.ID
}

func IndexSetFrom(ptr unsafe.Pointer) IndexSet {
	return IndexSet{
		ID: objc.ID(ptr),
	}
}




