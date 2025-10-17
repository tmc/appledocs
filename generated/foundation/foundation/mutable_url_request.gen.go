// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [MutableURLRequest] class.
var MutableURLRequestClass objc.Class

func init() {
	MutableURLRequestClass = objc.GetClass("NSMutableURLRequest")
}

type MutableURLRequest struct {
	objc.ID
}

func MutableURLRequestFrom(ptr unsafe.Pointer) MutableURLRequest {
	return MutableURLRequest{
		ID: objc.ID(ptr),
	}
}




