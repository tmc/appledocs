// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [MutableURLRequest] class.
var (
	MutableURLRequestClass     _MutableURLRequestClass
	MutableURLRequestClassOnce sync.Once
)

func getMutableURLRequestClass() _MutableURLRequestClass {
	MutableURLRequestClassOnce.Do(func() {
		MutableURLRequestClass = _MutableURLRequestClass{objc.GetClass("NSMutableURLRequest")}
	})
	return MutableURLRequestClass
}

type _MutableURLRequestClass struct {
	class objc.Class
}

// An interface definition for the [MutableURLRequest] class.
type IMutableURLRequest interface {
	IURLRequest
	BindToHotspotHelperCommand(command unsafe.Pointer)
}

// A mutable URL load request that is independent of protocol or URL scheme.
//
// In Swift, this object bridges to and you use when you need reference semantics or other Foundation-specific behavior. is a subclass of that allows you to change the request’s properties. only represents information about the request. Use other classes, such as , to send the request to a server. See and for an introduction to these techniques. Classes that create a network operation based on a request make a deep copy of that request. Thus, changing the request after creating a network operation has no effect on the ongoing operation. For example, if you use to create a data task from a request, and then later change the request, the data task continues using the original request.
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

// Alloc allocates a new instance without initialization.
func (mc _MutableURLRequestClass) Alloc() MutableURLRequest {
	rv := objc.Send[MutableURLRequest](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MutableURLRequestClass) New() MutableURLRequest {
	rv := objc.Send[MutableURLRequest](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MutableURLRequest) Init() MutableURLRequest {
	rv := objc.Send[MutableURLRequest](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MutableURLRequest) Autorelease() MutableURLRequest {
	rv := objc.Send[MutableURLRequest](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMutableURLRequest creates a new MutableURLRequest instance.
func NewMutableURLRequest() MutableURLRequest {
	return getMutableURLRequestClass().New()
}


// Binds a URL request to the network interface associated with the hotspot helper command instance.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMutableURLRequest/bind(to:)
func (m_ MutableURLRequest) BindToHotspotHelperCommand(command unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("bindToHotspotHelperCommand:"), command)
}

// The request’s cache policy.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMutableURLRequest/cachePolicy
func (m_ MutableURLRequest) CachePolicy() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("cachePolicy"))
	return rv
}


// SetCachePolicy sets the value of the cachePolicy property.
// The request’s cache policy.

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMutableURLRequest/cachePolicy
func (m_ MutableURLRequest) SetCachePolicy(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setCachePolicy:"), value)
}
// The HTTP request method.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMutableURLRequest/httpMethod
func (m_ MutableURLRequest) HTTPMethod() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("HTTPMethod"))
	return rv
}


// SetHTTPMethod sets the value of the HTTPMethod property.
// The HTTP request method.

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMutableURLRequest/httpMethod
func (m_ MutableURLRequest) SetHTTPMethod(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setHTTPMethod:"), value)
}
// The URL being requested.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMutableURLRequest/url
func (m_ MutableURLRequest) URL() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("URL"))
	return rv
}


// SetURL sets the value of the URL property.
// The URL being requested.

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMutableURLRequest/url
func (m_ MutableURLRequest) SetURL(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setURL:"), value)
}


