// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [URLProtocol] class.
var URLProtocolClass objc.Class

func init() {
	URLProtocolClass = objc.GetClass("NSURLProtocol")
}

type URLProtocol struct {
	objc.ID
}

func URLProtocolFrom(ptr unsafe.Pointer) URLProtocol {
	return URLProtocol{
		ID: objc.ID(ptr),
	}
}



