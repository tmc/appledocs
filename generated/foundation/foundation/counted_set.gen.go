// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [CountedSet] class.
var CountedSetClass objc.Class

func init() {
	CountedSetClass = objc.GetClass("NSCountedSet")
}

type CountedSet struct {
	objc.ID
}

func CountedSetFrom(ptr unsafe.Pointer) CountedSet {
	return CountedSet{
		ID: objc.ID(ptr),
	}
}




