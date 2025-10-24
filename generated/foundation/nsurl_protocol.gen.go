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
	SetCachedResponse(value ICachedURLResponse)
	Client() unsafe.Pointer
	SetClient(value unsafe.Pointer)
	Request() IURLRequest
	SetRequest(value IURLRequest)
	Task() IURLSessionTask
	SetTask(value IURLSessionTask)
	ProtocolClasses() objc.Class
	SetProtocolClasses(value objc.Class)
	// methods:
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



// The protocol’s cached response.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/urlprotocol/cachedresponse
func (u_ URLProtocol) CachedResponse() ICachedURLResponse {
	rv := objc.Send[CachedURLResponse](u_.ID, objc.Sel("cachedResponse"))
	return rv
}


// The protocol’s cached response.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/urlprotocol/cachedresponse
func (u_ URLProtocol) SetCachedResponse(value ICachedURLResponse) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setCachedResponse:"), value)
}


// The object the protocol uses to communicate with the URL loading system.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/urlprotocol/client
func (u_ URLProtocol) Client() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](u_.ID, objc.Sel("client"))
	return rv
}


// The object the protocol uses to communicate with the URL loading system.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/urlprotocol/client
func (u_ URLProtocol) SetClient(value unsafe.Pointer) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setClient:"), value)
}


// The protocol’s request.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/urlprotocol/request
func (u_ URLProtocol) Request() IURLRequest {
	rv := objc.Send[URLRequest](u_.ID, objc.Sel("request"))
	return rv
}


// The protocol’s request.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/urlprotocol/request
func (u_ URLProtocol) SetRequest(value IURLRequest) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setRequest:"), value)
}


// The protocol’s task.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/urlprotocol/task
func (u_ URLProtocol) Task() IURLSessionTask {
	rv := objc.Send[URLSessionTask](u_.ID, objc.Sel("task"))
	return rv
}


// The protocol’s task.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/urlprotocol/task
func (u_ URLProtocol) SetTask(value IURLSessionTask) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setTask:"), value)
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



