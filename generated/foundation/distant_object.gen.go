// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [DistantObject] class.
var DistantObjectClass = _DistantObjectClass{objc.GetClass("NSDistantObject")}

type _DistantObjectClass struct {
	class objc.Class
}

type DistantObject struct {
	objc.ID
}

func DistantObjectFrom(ptr unsafe.Pointer) DistantObject {
	return DistantObject{
		ID: objc.ID(ptr),
	}
}




