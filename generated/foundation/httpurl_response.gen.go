// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [HTTPURLResponse] class.
var (
	hTTPURLResponseClass     _HTTPURLResponseClass
	hTTPURLResponseClassOnce sync.Once
)

func getHTTPURLResponseClass() _HTTPURLResponseClass {
	hTTPURLResponseClassOnce.Do(func() {
		hTTPURLResponseClass = _HTTPURLResponseClass{objc.GetClass("NSHTTPURLResponse")}
	})
	return hTTPURLResponseClass
}

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




