// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [URLResponse] class.
var uRLResponseClass = _URLResponseClass{objc.GetClass("NSURLResponse")}

type _URLResponseClass struct {
	class objc.Class
}

// An interface definition for the [URLResponse] class.
type IURLResponse interface {
	objectivec.IObject
}

// The metadata associated with the response to a URL load request, independent of protocol and URL scheme. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLResponse

type URLResponse struct {
	objectivec.Object
}

// URLResponseFrom constructs a [URLResponse] from an unsafe.Pointer.
//
// The metadata associated with the response to a URL load request, independent of protocol and URL scheme.
func URLResponseFrom(ptr unsafe.Pointer) URLResponse {
	return URLResponse{objectivec.Object{objc.ID(ptr)}}
}



