// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

var MutableOrderedSetClass _MutableOrderedSetClass

func init() {
	MutableOrderedSetClass = _MutableOrderedSetClass{objc.GetClass("NSMutableOrderedSet")}
}

type _MutableOrderedSetClass struct {
	class objc.Class
}

type MutableOrderedSet struct {
	objc.ID
}

func MutableOrderedSetFrom(ptr unsafe.Pointer) MutableOrderedSet {
	return MutableOrderedSet{
		ID: objc.ID(ptr),
	}
}




