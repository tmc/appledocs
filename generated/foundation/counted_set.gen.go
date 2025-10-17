// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [CountedSet] class.
var CountedSetClass = _CountedSetClass{objc.GetClass("NSCountedSet")}

type _CountedSetClass struct {
	class objc.Class
}

type CountedSet struct {
	objc.ID
}

func CountedSetFrom(ptr unsafe.Pointer) CountedSet {
	return CountedSet{
		ID: objc.ID(ptr),
	}
}




