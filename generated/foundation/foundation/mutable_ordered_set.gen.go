// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [MutableOrderedSet] class.
var MutableOrderedSetClass objc.Class

func init() {
	MutableOrderedSetClass = objc.GetClass("NSMutableOrderedSet")
}

type MutableOrderedSet struct {
	objc.ID
}

func MutableOrderedSetFrom(ptr unsafe.Pointer) MutableOrderedSet {
	return MutableOrderedSet{
		ID: objc.ID(ptr),
	}
}




