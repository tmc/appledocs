// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [HTTPURLResponse] class.
var hTTPURLResponseClass = _HTTPURLResponseClass{objc.GetClass("NSHTTPURLResponse")}

type _HTTPURLResponseClass struct {
	class objc.Class
}

// An interface definition for the [HTTPURLResponse] class.
type IHTTPURLResponse interface {
	IURLResponse
}

// The metadata associated with the response to an HTTP protocol URL load request. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/HTTPURLResponse

type HTTPURLResponse struct {
	URLResponse
}

// HTTPURLResponseFrom constructs a [HTTPURLResponse] from an unsafe.Pointer.
//
// The metadata associated with the response to an HTTP protocol URL load request.
func HTTPURLResponseFrom(ptr unsafe.Pointer) HTTPURLResponse {
	return HTTPURLResponse{
		URLResponse: URLResponseFrom(ptr),
	}
}



