// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [MutableURLRequest] class.
var mutableURLRequestClass = _MutableURLRequestClass{objc.GetClass("NSMutableURLRequest")}

type _MutableURLRequestClass struct {
	class objc.Class
}

// A mutable URL load request that is independent of protocol or URL scheme. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMutableURLRequest

type MutableURLRequest struct {
	URLRequest
}

// MutableURLRequestFrom constructs a [MutableURLRequest] from an unsafe.Pointer.
//
// A mutable URL load request that is independent of protocol or URL scheme.
func MutableURLRequestFrom(ptr unsafe.Pointer) MutableURLRequest {
	return MutableURLRequest{
		URLRequest: URLRequestFrom(ptr),
	}
}



