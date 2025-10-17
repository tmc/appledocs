// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [IndexPath] class.
var IndexPathClass objc.Class

func init() {
	IndexPathClass = objc.GetClass("NSIndexPath")
}

type IndexPath struct {
	objc.ID
}

func IndexPathFrom(ptr unsafe.Pointer) IndexPath {
	return IndexPath{
		ID: objc.ID(ptr),
	}
}




