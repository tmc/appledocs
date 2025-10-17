// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [MutableSet] class.
var MutableSetClass objc.Class

func init() {
	MutableSetClass = objc.GetClass("NSMutableSet")
}

type MutableSet struct {
	objc.ID
}

func MutableSetFrom(ptr unsafe.Pointer) MutableSet {
	return MutableSet{
		ID: objc.ID(ptr),
	}
}



