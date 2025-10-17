// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [URLRequest] class.
var URLRequestClass = _URLRequestClass{objc.GetClass("NSURLRequest")}

type _URLRequestClass struct {
	class objc.Class
}

type URLRequest struct {
	objc.ID
}

func URLRequestFrom(ptr unsafe.Pointer) URLRequest {
	return URLRequest{
		ID: objc.ID(ptr),
	}
}




