// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [URLRequest] class.
var URLRequestClass objc.Class

func init() {
	URLRequestClass = objc.GetClass("NSURLRequest")
}

type URLRequest struct {
	objc.ID
}

func URLRequestFrom(ptr unsafe.Pointer) URLRequest {
	return URLRequest{
		ID: objc.ID(ptr),
	}
}



