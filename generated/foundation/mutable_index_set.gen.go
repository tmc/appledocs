// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [MutableIndexSet] class.
var MutableIndexSetClass objc.Class

func init() {
	MutableIndexSetClass = objc.GetClass("NSMutableIndexSet")
}

type MutableIndexSet struct {
	objc.ID
}

func MutableIndexSetFrom(ptr unsafe.Pointer) MutableIndexSet {
	return MutableIndexSet{
		ID: objc.ID(ptr),
	}
}



