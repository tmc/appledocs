// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

var OrderedSetClass _OrderedSetClass

func init() {
	OrderedSetClass = _OrderedSetClass{objc.GetClass("NSOrderedSet")}
}

type _OrderedSetClass struct {
	class objc.Class
}

type OrderedSet struct {
	objc.ID
}

func OrderedSetFrom(ptr unsafe.Pointer) OrderedSet {
	return OrderedSet{
		ID: objc.ID(ptr),
	}
}




