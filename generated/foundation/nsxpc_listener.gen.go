// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [XPCListener] class.
var (
	XPCListenerClass     _XPCListenerClass
	XPCListenerClassOnce sync.Once
)

func getXPCListenerClass() _XPCListenerClass {
	XPCListenerClassOnce.Do(func() {
		XPCListenerClass = _XPCListenerClass{objc.GetClass("NSXPCListener")}
	})
	return XPCListenerClass
}

type _XPCListenerClass struct {
	class objc.Class
}

// An interface definition for the [XPCListener] class.
type IXPCListener interface {
	objectivec.IObject
	// properties:
	Delegate() objc.ID
	SetDelegate(value objc.ID)
	Endpoint() IXPCListenerEndpoint
	// methods:
	Activate()
	Invalidate()
	Resume()
	SetConnectionCodeSigningRequirement(requirement IString)
	Suspend()
}

// A listener that waits for new incoming connections, configures them, and accepts or rejects them.
//
// Each XPC service, launchd agent, or launchd daemon typically has at least one object that listens for connections to a specified service name. Each listener must have a delegate that conforms to the protocol. When the listener receives a new connection request, it creates a new object, then asks the delegate to inspect, configure, and resume the connection object by calling the delegate’s method.


// A listener that waits for new incoming connections, configures them, and accepts or rejects them.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSXPCListener
type XPCListener struct {
	objectivec.Object
}

// XPCListenerFrom constructs a [XPCListener] from an unsafe.Pointer.
//
// A listener that waits for new incoming connections, configures them, and accepts or rejects them.
func XPCListenerFrom(ptr unsafe.Pointer) XPCListener {
	return XPCListener{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (xc _XPCListenerClass) Alloc() XPCListener {
	rv := objc.Send[XPCListener](objc.ID(xc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (xc _XPCListenerClass) New() XPCListener {
	rv := objc.Send[XPCListener](objc.ID(xc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (x_ XPCListener) Init() XPCListener {
	rv := objc.Send[XPCListener](x_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (x_ XPCListener) Autorelease() XPCListener {
	rv := objc.Send[XPCListener](x_.ID, objc.Sel("autorelease"))
	return rv
}

// NewXPCListener creates a new XPCListener instance.
func NewXPCListener() XPCListener {
	return getXPCListenerClass().New()
}



// Initializes a listener in a LaunchAgent or LaunchDaemon which has a name advertised in a file.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSXPCListener/init(machServiceName:)
func NewXPCListenerWithMachServiceName(name IString) XPCListener {
	instance := getXPCListenerClass().Alloc()
	rv := objc.Send[XPCListener](instance.ID, objc.Sel("initWithMachServiceName:"), name)
	rv.Autorelease()
	return rv
}



// Returns a new anonymous listener connection.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSXPCListener/anonymous()
func (xc _XPCListenerClass) AnonymousListener() IXPCListener {
	rv := objc.Send[XPCListener](objc.ID(xc.class), objc.Sel("anonymousListener"))
	return rv
}


// Returns the singleton listener used to listen for incoming connections in an XPC service.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSXPCListener/service()
func (xc _XPCListenerClass) ServiceListener() IXPCListener {
	rv := objc.Send[XPCListener](objc.ID(xc.class), objc.Sel("serviceListener"))
	return rv
}


// Activates the listener.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSXPCListener/activate()
func (x_ XPCListener) Activate() {
	objc.Send[objc.ID](x_.ID, objc.Sel("activate"))
}


// Invalidates the listener.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSXPCListener/invalidate()
func (x_ XPCListener) Invalidate() {
	objc.Send[objc.ID](x_.ID, objc.Sel("invalidate"))
}


// Starts processing of incoming requests.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSXPCListener/resume()
func (x_ XPCListener) Resume() {
	objc.Send[objc.ID](x_.ID, objc.Sel("resume"))
}


// Sets the code signing requirement for connections to this listener.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSXPCListener/setConnectionCodeSigningRequirement(_:)
func (x_ XPCListener) SetConnectionCodeSigningRequirement(requirement IString) {
	objc.Send[objc.ID](x_.ID, objc.Sel("setConnectionCodeSigningRequirement:"), requirement)
}


// Suspends the listener.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSXPCListener/suspend()
func (x_ XPCListener) Suspend() {
	objc.Send[objc.ID](x_.ID, objc.Sel("suspend"))
}


// The delegate for the listener.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSXPCListener/delegate
func (x_ XPCListener) Delegate() objc.ID {
	rv := objc.Send[objc.ID](x_.ID, objc.Sel("delegate"))
	return rv
}


// The delegate for the listener.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSXPCListener/delegate
func (x_ XPCListener) SetDelegate(value objc.ID) {
	objc.Send[objc.ID](x_.ID, objc.Sel("setDelegate:"), value)
}


// Returns an endpoint object that may be sent over an existing connection.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSXPCListener/endpoint
func (x_ XPCListener) Endpoint() IXPCListenerEndpoint {
	rv := objc.Send[XPCListenerEndpoint](x_.ID, objc.Sel("endpoint"))
	return rv
}


