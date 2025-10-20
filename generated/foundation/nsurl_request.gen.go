// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [URLRequest] class.
var (
	URLRequestClass     _URLRequestClass
	URLRequestClassOnce sync.Once
)

func getURLRequestClass() _URLRequestClass {
	URLRequestClassOnce.Do(func() {
		URLRequestClass = _URLRequestClass{objc.GetClass("NSURLRequest")}
	})
	return URLRequestClass
}

type _URLRequestClass struct {
	class objc.Class
}

// An interface definition for the [URLRequest] class.
type IURLRequest interface {
	objectivec.IObject
}

// A URL load request that is independent of protocol or URL scheme.
//
// Use this type in Swift when you need reference semantics or other Foundation-specific behavior. encapsulates two essential properties of a load request: the URL to load and the policies used to load it. In addition, for HTTP and HTTPS requests, includes the HTTP method ( , , and so on) and the HTTP headers. Finally, custom protocols can support custom properties as explained in . only represents information about the request. Use other classes, such as , to send the request to a server. See and for an introduction to these techniques. The mutable subclass of is .
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

// Alloc allocates a new instance without initialization.
func (uc _URLRequestClass) Alloc() URLRequest {
	rv := objc.Send[URLRequest](objc.ID(uc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (uc _URLRequestClass) New() URLRequest {
	rv := objc.Send[URLRequest](objc.ID(uc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (u_ URLRequest) Init() URLRequest {
	rv := objc.Send[URLRequest](u_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (u_ URLRequest) Autorelease() URLRequest {
	rv := objc.Send[URLRequest](u_.ID, objc.Sel("autorelease"))
	return rv
}

// NewURLRequest creates a new URLRequest instance.
func NewURLRequest() URLRequest {
	return getURLRequestClass().New()
}


// A Boolean value that indicates whether connections may use the network when the user has specified Low Data Mode.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURLRequest/allowsConstrainedNetworkAccess
func (u_ URLRequest) AllowsConstrainedNetworkAccess() bool {
	rv := objc.Send[bool](u_.ID, objc.Sel("allowsConstrainedNetworkAccess"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURLRequest/allowsPersistentDNS
func (u_ URLRequest) AllowsPersistentDNS() bool {
	rv := objc.Send[bool](u_.ID, objc.Sel("allowsPersistentDNS"))
	return rv
}

// A Boolean value that indicates whether the default cookie handling will be used for this request.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURLRequest/httpShouldHandleCookies
func (u_ URLRequest) HTTPShouldHandleCookies() bool {
	rv := objc.Send[bool](u_.ID, objc.Sel("HTTPShouldHandleCookies"))
	return rv
}

// The network service type of the request.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURLRequest/networkServiceType-swift.property
func (u_ URLRequest) NetworkServiceType() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](u_.ID, objc.Sel("networkServiceType"))
	return rv
}



