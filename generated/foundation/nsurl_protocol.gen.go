// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [URLProtocol] class.
var (
	URLProtocolClass     _URLProtocolClass
	URLProtocolClassOnce sync.Once
)

func getURLProtocolClass() _URLProtocolClass {
	URLProtocolClassOnce.Do(func() {
		URLProtocolClass = _URLProtocolClass{objc.GetClass("NSURLProtocol")}
	})
	return URLProtocolClass
}

type _URLProtocolClass struct {
	class objc.Class
}

// An interface definition for the [URLProtocol] class.
type IURLProtocol interface {
	objectivec.IObject
	// properties:
	CachedResponse() ICachedURLResponse
	Client() objc.ID
	Request() IURLRequest
	Task() IURLSessionTask
	ProtocolClasses() objc.Class
	SetProtocolClasses(value objc.Class)
	// methods:
	StartLoading()
	StopLoading()
}

// An abstract class that handles the loading of protocol-specific URL data.
//
// Don’t instantiate a subclass directly. Instead, create subclasses for any custom protocols or URL schemes that your app supports. When a download starts, the system creates the appropriate protocol object to handle the corresponding URL request. You define your protocol class and call the class method during your app’s launch time so that the system is aware of your protocol. To support the customization of protocol-specific requests, create extensions to the class to provide any custom API that you need. You can store and retrieve protocol-specific request data by using ’s class methods and . Create a for each request your subclass processes successfully. You may want to create a custom class to provide protocol specific information.


// An abstract class that handles the loading of protocol-specific URL data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLProtocol
type URLProtocol struct {
	objectivec.Object
}

// URLProtocolFrom constructs a [URLProtocol] from an unsafe.Pointer.
//
// An abstract class that handles the loading of protocol-specific URL data.
func URLProtocolFrom(ptr unsafe.Pointer) URLProtocol {
	return URLProtocol{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (uc _URLProtocolClass) Alloc() URLProtocol {
	rv := objc.Send[URLProtocol](objc.ID(uc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (uc _URLProtocolClass) New() URLProtocol {
	rv := objc.Send[URLProtocol](objc.ID(uc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (u_ URLProtocol) Init() URLProtocol {
	rv := objc.Send[URLProtocol](u_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (u_ URLProtocol) Autorelease() URLProtocol {
	rv := objc.Send[URLProtocol](u_.ID, objc.Sel("autorelease"))
	return rv
}

// NewURLProtocol creates a new URLProtocol instance.
func NewURLProtocol() URLProtocol {
	return getURLProtocolClass().New()
}



// Creates a URL protocol instance to handle the request.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLProtocol/init(request:cachedResponse:client:)
func NewURLProtocolWithRequestCachedResponseClient(request IURLRequest, cachedResponse ICachedURLResponse, client objc.IObject) URLProtocol {
	instance := getURLProtocolClass().Alloc()
	rv := objc.Send[URLProtocol](instance.ID, objc.Sel("initWithRequest:cachedResponse:client:"), request, cachedResponse, client)
	rv.Autorelease()
	return rv
}


// Creates a URL protocol instance to handle the task.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLProtocol/init(task:cachedResponse:client:)
func NewURLProtocolWithTaskCachedResponseClient(task IURLSessionTask, cachedResponse ICachedURLResponse, client objc.IObject) URLProtocol {
	instance := getURLProtocolClass().Alloc()
	rv := objc.Send[URLProtocol](instance.ID, objc.Sel("initWithTask:cachedResponse:client:"), task, cachedResponse, client)
	rv.Autorelease()
	return rv
}



// Determines whether the protocol subclass can handle the specified task.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLProtocol/canInit(with:)-18gbo
func (uc _URLProtocolClass) CanInitWithTask(task IURLSessionTask) bool {
	rv := objc.Send[bool](objc.ID(uc.class), objc.Sel("canInitWithTask:"), task)
	return rv
}


// Determines whether the protocol subclass can handle the specified request.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLProtocol/canInit(with:)-76brg
func (uc _URLProtocolClass) CanInitWithRequest(request IURLRequest) bool {
	rv := objc.Send[bool](objc.ID(uc.class), objc.Sel("canInitWithRequest:"), request)
	return rv
}


// Returns a canonical version of the specified request.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLProtocol/canonicalRequest(for:)
func (uc _URLProtocolClass) CanonicalRequestForRequest(request IURLRequest) IURLRequest {
	rv := objc.Send[URLRequest](objc.ID(uc.class), objc.Sel("canonicalRequestForRequest:"), request)
	return rv
}


// Attempts to register a subclass of , making it visible to the URL loading system.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLProtocol/registerClass(_:)
func (uc _URLProtocolClass) RegisterClass(protocolClass objc.Class) bool {
	rv := objc.Send[bool](objc.ID(uc.class), objc.Sel("registerClass:"), protocolClass)
	return rv
}


// Removes the property associated with the specified key in the specified request.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLProtocol/removeProperty(forKey:in:)
func (uc _URLProtocolClass) RemovePropertyForKeyInRequest(key IString, request IMutableURLRequest) {
	objc.Send[objc.ID](objc.ID(uc.class), objc.Sel("removePropertyForKey:inRequest:"), key, request)
}


// A Boolean value indicating whether two requests are equivalent for cache purposes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLProtocol/requestIsCacheEquivalent(_:to:)
func (uc _URLProtocolClass) RequestIsCacheEquivalentToRequest(a IURLRequest, b IURLRequest) bool {
	rv := objc.Send[bool](objc.ID(uc.class), objc.Sel("requestIsCacheEquivalent:toRequest:"), a, b)
	return rv
}


// Unregisters the specified subclass of .
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLProtocol/unregisterClass(_:)
func (uc _URLProtocolClass) UnregisterClass(protocolClass objc.Class) {
	objc.Send[objc.ID](objc.ID(uc.class), objc.Sel("unregisterClass:"), protocolClass)
}


// Starts protocol-specific loading of the request.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLProtocol/startLoading()
func (u_ URLProtocol) StartLoading() {
	objc.Send[objc.ID](u_.ID, objc.Sel("startLoading"))
}


// Stops protocol-specific loading of the request.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLProtocol/stopLoading()
func (u_ URLProtocol) StopLoading() {
	objc.Send[objc.ID](u_.ID, objc.Sel("stopLoading"))
}


// The protocol’s cached response.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLProtocol/cachedResponse
func (u_ URLProtocol) CachedResponse() ICachedURLResponse {
	rv := objc.Send[CachedURLResponse](u_.ID, objc.Sel("cachedResponse"))
	return rv
}


// The object the protocol uses to communicate with the URL loading system.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLProtocol/client
func (u_ URLProtocol) Client() objc.ID {
	rv := objc.Send[objc.ID](u_.ID, objc.Sel("client"))
	return rv
}


// The protocol’s request.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLProtocol/request
func (u_ URLProtocol) Request() IURLRequest {
	rv := objc.Send[URLRequest](u_.ID, objc.Sel("request"))
	return rv
}


// The protocol’s task.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLProtocol/task
func (u_ URLProtocol) Task() IURLSessionTask {
	rv := objc.Send[URLSessionTask](u_.ID, objc.Sel("task"))
	return rv
}


// An array of extra protocol subclasses that handle requests in a session.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/urlsessionconfiguration/protocolclasses
func (u_ URLProtocol) ProtocolClasses() objc.Class {
	rv := objc.Send[objc.Class](u_.ID, objc.Sel("protocolClasses"))
	return rv
}


// An array of extra protocol subclasses that handle requests in a session.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/urlsessionconfiguration/protocolclasses
func (u_ URLProtocol) SetProtocolClasses(value objc.Class) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setProtocolClasses:"), value)
}


