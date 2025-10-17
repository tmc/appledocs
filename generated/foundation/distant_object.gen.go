// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [DistantObject] class.
var DistantObjectClass objc.Class

func init() {
	DistantObjectClass = objc.GetClass("NSDistantObject")
}

type DistantObject struct {
	objc.ID
}

func DistantObjectFrom(ptr unsafe.Pointer) DistantObject {
	return DistantObject{
		ID: objc.ID(ptr),
	}
}



