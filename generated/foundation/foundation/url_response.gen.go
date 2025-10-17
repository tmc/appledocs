// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [URLResponse] class.
var URLResponseClass objc.Class

func init() {
	URLResponseClass = objc.GetClass("NSURLResponse")
}

type URLResponse struct {
	objc.ID
}

func URLResponseFrom(ptr unsafe.Pointer) URLResponse {
	return URLResponse{
		ID: objc.ID(ptr),
	}
}




