// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

var mutableSetClass _MutableSetClass

func init() {
	mutableSetClass = _MutableSetClass{objc.GetClass("NSMutableSet")}
}

type _MutableSetClass struct {
	class objc.Class
}

type MutableSet struct {
	objc.ID
}

func MutableSetFrom(ptr unsafe.Pointer) MutableSet {
	return MutableSet{
		ID: objc.ID(ptr),
	}
}




