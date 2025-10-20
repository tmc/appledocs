// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

var MutableURLRequestClass _MutableURLRequestClass

func init() {
	MutableURLRequestClass = _MutableURLRequestClass{objc.GetClass("NSMutableURLRequest")}
}

type _MutableURLRequestClass struct {
	class objc.Class
}

type MutableURLRequest struct {
	objc.ID
}

func MutableURLRequestFrom(ptr unsafe.Pointer) MutableURLRequest {
	return MutableURLRequest{
		ID: objc.ID(ptr),
	}
}




