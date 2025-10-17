// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [IndexSet] class.
var IndexSetClass objc.Class

func init() {
	IndexSetClass = objc.GetClass("NSIndexSet")
}

type IndexSet struct {
	objc.ID
}

func IndexSetFrom(ptr unsafe.Pointer) IndexSet {
	return IndexSet{
		ID: objc.ID(ptr),
	}
}




