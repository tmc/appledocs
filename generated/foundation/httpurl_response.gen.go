// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [HTTPURLResponse] class.
var HTTPURLResponseClass objc.Class

func init() {
	HTTPURLResponseClass = objc.GetClass("NSHTTPURLResponse")
}

type HTTPURLResponse struct {
	objc.ID
}

func HTTPURLResponseFrom(ptr unsafe.Pointer) HTTPURLResponse {
	return HTTPURLResponse{
		ID: objc.ID(ptr),
	}
}



