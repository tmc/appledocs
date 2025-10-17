// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [MutableIndexSet] class.
var MutableIndexSetClass = _MutableIndexSetClass{objc.GetClass("NSMutableIndexSet")}

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




