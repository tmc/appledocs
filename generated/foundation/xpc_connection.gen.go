// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [XPCConnection] class.
var XPCConnectionClass = _XPCConnectionClass{objc.GetClass("NSXPCConnection")}

type _XPCConnectionClass struct {
	class objc.Class
}

type XPCConnection struct {
	objc.ID
}

func XPCConnectionFrom(ptr unsafe.Pointer) XPCConnection {
	return XPCConnection{
		ID: objc.ID(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (xc _XPCConnectionClass) Alloc() XPCConnection {
	rv := objc.Send[XPCConnection](objc.ID(xc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new instance with a +1 retain count.
func (xc _XPCConnectionClass) New() XPCConnection {
	rv := objc.Send[XPCConnection](objc.ID(xc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (x_ XPCConnection) Init() XPCConnection {
	rv := objc.Send[XPCConnection](x_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (x_ XPCConnection) Autorelease() XPCConnection {
	rv := objc.Send[XPCConnection](x_.ID, objc.Sel("autorelease"))
	return rv
}

// NewXPCConnection creates a new XPCConnection instance.
func NewXPCConnection() XPCConnection {
	return XPCConnectionClass.New()
}
// Initializes an object to connect to an object in another process, identified by an object. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSXPCConnection/init(listenerEndpoint:)
func NewXPCConnectionWithListenerEndpoint(endpoint unsafe.Pointer) XPCConnection {
	instance := XPCConnectionClass.Alloc()
	rv := objc.Send[XPCConnection](instance.ID, objc.Sel("initWithListenerEndpoint:"), endpoint)
	rv.Autorelease()
	return rv
}
// Initializes an object to connect to a LaunchAgent or LaunchDaemon with a name advertised in a . [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSXPCConnection/init(machServiceName:options:)
func NewXPCConnectionWithMachServiceNameOptions(name string, options unsafe.Pointer) XPCConnection {
	instance := XPCConnectionClass.Alloc()
	rv := objc.Send[XPCConnection](instance.ID, objc.Sel("initWithMachServiceName:options:"), name, options)
	rv.Autorelease()
	return rv
}
// Initializes an object to connect to an object in an XPC service, identified by a service name. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSXPCConnection/init(serviceName:)
func NewXPCConnectionWithServiceName(serviceName string) XPCConnection {
	instance := XPCConnectionClass.Alloc()
	rv := objc.Send[XPCConnection](instance.ID, objc.Sel("initWithServiceName:"), serviceName)
	rv.Autorelease()
	return rv
}


// Returns the current connection, in the context of a call to a method on your exported object. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSXPCConnection/current()
func (xc _XPCConnectionClass) CurrentConnection() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(xc.class), objc.Sel("currentConnection"))
	return rv
}
// Activates the connection. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSXPCConnection/activate()
func (x_ XPCConnection) Activate() {
	objc.Send[objc.ID](x_.ID, objc.Sel("activate"))
}
// Invalidates the connection. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSXPCConnection/invalidate()
func (x_ XPCConnection) Invalidate() {
	objc.Send[objc.ID](x_.ID, objc.Sel("invalidate"))
}
// Returns a proxy for the remote object (that is, the object exported from the other side of this connection) with the specified error handler. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSXPCConnection/remoteObjectProxyWithErrorHandler(_:)
func (x_ XPCConnection) RemoteObjectProxyWithErrorHandler(handler unsafe.Pointer) objc.ID {
	rv := objc.Send[objc.ID](x_.ID, objc.Sel("remoteObjectProxyWithErrorHandler:"), handler)
	return rv
}
// Starts or resumes handling of messages on a connection. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSXPCConnection/resume()
func (x_ XPCConnection) Resume() {
	objc.Send[objc.ID](x_.ID, objc.Sel("resume"))
}
// Add a barrier block to execute on the connection. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSXPCConnection/scheduleSendBarrierBlock(_:)
func (x_ XPCConnection) ScheduleSendBarrierBlock(block unsafe.Pointer) {
	objc.Send[objc.ID](x_.ID, objc.Sel("scheduleSendBarrierBlock:"), block)
}
// Sets the code signing requirement for this connection. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSXPCConnection/setCodeSigningRequirement(_:)
func (x_ XPCConnection) SetCodeSigningRequirement(requirement string) {
	objc.Send[objc.ID](x_.ID, objc.Sel("setCodeSigningRequirement:"), requirement)
}
// Suspends the connection. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSXPCConnection/suspend()
func (x_ XPCConnection) Suspend() {
	objc.Send[objc.ID](x_.ID, objc.Sel("suspend"))
}
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSXPCConnection/synchronousRemoteObjectProxyWithErrorHandler(_:)
func (x_ XPCConnection) SynchronousRemoteObjectProxyWithErrorHandler(handler unsafe.Pointer) objc.ID {
	rv := objc.Send[objc.ID](x_.ID, objc.Sel("synchronousRemoteObjectProxyWithErrorHandler:"), handler)
	return rv
}


