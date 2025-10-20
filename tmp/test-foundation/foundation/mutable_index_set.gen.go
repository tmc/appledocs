// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

var MutableIndexSetClass _MutableIndexSetClass

func init() {
	MutableIndexSetClass = _MutableIndexSetClass{objc.GetClass("NSMutableIndexSet")}
}

type _MutableIndexSetClass struct {
	class objc.Class
}

type MutableIndexSet struct {
	objc.ID
}

func MutableIndexSetFrom(ptr unsafe.Pointer) MutableIndexSet {
	return MutableIndexSet{
		ID: objc.ID(ptr),
	}
}




