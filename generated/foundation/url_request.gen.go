// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [URLRequest] class.
var uRLRequestClass = _URLRequestClass{objc.GetClass("NSURLRequest")}

type _URLRequestClass struct {
	class objc.Class
}

// A URL load request that is independent of protocol or URL scheme. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURLRequest

type URLRequest struct {
	objectivec.Object
}

// URLRequestFrom constructs a [URLRequest] from an unsafe.Pointer.
//
// A URL load request that is independent of protocol or URL scheme.
func URLRequestFrom(ptr unsafe.Pointer) URLRequest {
	return URLRequest{objectivec.Object{objc.ID(ptr)}}
}



