// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [HTTPURLResponse] class.
var (
	HTTPURLResponseClass     _HTTPURLResponseClass
	HTTPURLResponseClassOnce sync.Once
)

func getHTTPURLResponseClass() _HTTPURLResponseClass {
	HTTPURLResponseClassOnce.Do(func() {
		HTTPURLResponseClass = _HTTPURLResponseClass{objc.GetClass("NSHTTPURLResponse")}
	})
	return HTTPURLResponseClass
}

type _HTTPURLResponseClass struct {
	class objc.Class
}

// An interface definition for the [HTTPURLResponse] class.
type IHTTPURLResponse interface {
	IURLResponse
	// properties:
	AllHeaderFields() unsafe.Pointer
	SetAllHeaderFields(value unsafe.Pointer)
	StatusCode() int /* primitive/slice/pointer. */
	SetStatusCode(value int /* primitive/slice/pointer. */)
	// methods:
}

// The metadata associated with the response to an HTTP protocol URL load request.
//
// The class is a subclass of that provides methods for accessing information specific to HTTP protocol responses. Whenever you make HTTP URL load requests, any response objects you get back from the , , or class are instances of the class.


// The metadata associated with the response to an HTTP protocol URL load request.
//
// [Full Topic]
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

// Alloc allocates a new instance without initialization.
func (hc _HTTPURLResponseClass) Alloc() HTTPURLResponse {
	rv := objc.Send[HTTPURLResponse](objc.ID(hc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (hc _HTTPURLResponseClass) New() HTTPURLResponse {
	rv := objc.Send[HTTPURLResponse](objc.ID(hc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (h_ HTTPURLResponse) Init() HTTPURLResponse {
	rv := objc.Send[HTTPURLResponse](h_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (h_ HTTPURLResponse) Autorelease() HTTPURLResponse {
	rv := objc.Send[HTTPURLResponse](h_.ID, objc.Sel("autorelease"))
	return rv
}

// NewHTTPURLResponse creates a new HTTPURLResponse instance.
func NewHTTPURLResponse() HTTPURLResponse {
	return getHTTPURLResponseClass().New()
}



// All HTTP header fields of the response.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/httpurlresponse/allheaderfields
func (h_ HTTPURLResponse) AllHeaderFields() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](h_.ID, objc.Sel("allHeaderFields"))
	return rv
}


// All HTTP header fields of the response.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/httpurlresponse/allheaderfields
func (h_ HTTPURLResponse) SetAllHeaderFields(value unsafe.Pointer) {
	objc.Send[objc.ID](h_.ID, objc.Sel("setAllHeaderFields:"), value)
}


// The response’s HTTP status code.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/httpurlresponse/statuscode
func (h_ HTTPURLResponse) StatusCode() int /* primitive/slice/pointer. */ {
	rv := objc.Send[int](h_.ID, objc.Sel("statusCode"))
	return rv
}


// The response’s HTTP status code.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/httpurlresponse/statuscode
func (h_ HTTPURLResponse) SetStatusCode(value int /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](h_.ID, objc.Sel("setStatusCode:"), value)
}



