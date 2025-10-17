// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [OrderedSet] class.
var OrderedSetClass objc.Class

func init() {
	OrderedSetClass = objc.GetClass("NSOrderedSet")
}

type OrderedSet struct {
	objc.ID
}

func OrderedSetFrom(ptr unsafe.Pointer) OrderedSet {
	return OrderedSet{
		ID: objc.ID(ptr),
	}
}



