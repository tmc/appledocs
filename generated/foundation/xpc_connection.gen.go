// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [XPCConnection] class.
var XPCConnectionClass objc.Class

func init() {
	XPCConnectionClass = objc.GetClass("NSXPCConnection")
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
func (xc XPCConnection) Alloc() XPCConnection {
	ret := objc.ID(XPCConnectionClass).Send(objc.RegisterName("alloc"))
	return XPCConnection{ret}
}

// Init initializes the instance.
func (x_ XPCConnection) Init() XPCConnection {
	ret := x_.ID.Send(objc.RegisterName("init"))
	return XPCConnection{ret}
}
// Initializes an   object to connect to an   object in another process, identified by an   object. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NSXPCConnection/init(listenerEndpoint:)
func NewXPCConnectionWithListenerEndpoint(endpoint unsafe.Pointer) XPCConnection {
	instance := XPCConnection{}.Alloc()
	sel := objc.RegisterName("initWithListenerEndpoint:")
	ret := instance.ID.Send(sel, endpoint)
	instance = XPCConnection{ret}
	instance.ID = instance.ID.Send(objc.RegisterName("autorelease"))
	return instance
}
// Initializes an   object to connect to a LaunchAgent or LaunchDaemon with a name advertised in a  . [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NSXPCConnection/init(machServiceName:options:)
func NewXPCConnectionWithMachServiceNameOptions(name string, options unsafe.Pointer) XPCConnection {
	instance := XPCConnection{}.Alloc()
	sel := objc.RegisterName("initWithMachServiceName:options:")
	ret := instance.ID.Send(sel, name, options)
	instance = XPCConnection{ret}
	instance.ID = instance.ID.Send(objc.RegisterName("autorelease"))
	return instance
}
// Initializes an   object to connect to an   object in an XPC service, identified by a service name. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NSXPCConnection/init(serviceName:)
func NewXPCConnectionWithServiceName(serviceName string) XPCConnection {
	instance := XPCConnection{}.Alloc()
	sel := objc.RegisterName("initWithServiceName:")
	ret := instance.ID.Send(sel, serviceName)
	instance = XPCConnection{ret}
	instance.ID = instance.ID.Send(objc.RegisterName("autorelease"))
	return instance
}


// Returns the current connection, in the context of a call to a method on your exported object. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NSXPCConnection/current()
func (xc XPCConnection) CurrentConnection() unsafe.Pointer {
	sel := objc.RegisterName("currentConnection")
	ret := objc.ID(XPCConnectionClass).Send(sel)
	return unsafe.Pointer(ret)
}
// Activates the connection. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NSXPCConnection/activate()
func (x_ XPCConnection) Activate() {
	sel := objc.RegisterName("activate")
	x_.ID.Send(sel)
}
// Invalidates the connection. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NSXPCConnection/invalidate()
func (x_ XPCConnection) Invalidate() {
	sel := objc.RegisterName("invalidate")
	x_.ID.Send(sel)
}
// Returns a proxy for the remote object (that is, the object exported from the other side of this connection) with the specified error handler. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NSXPCConnection/remoteObjectProxyWithErrorHandler(_:)
func (x_ XPCConnection) RemoteObjectProxyWithErrorHandler(handler unsafe.Pointer) objc.ID {
	sel := objc.RegisterName("remoteObjectProxyWithErrorHandler:")
	ret := x_.ID.Send(sel, handler)
	return ret
}
// Starts or resumes handling of messages on a connection. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NSXPCConnection/resume()
func (x_ XPCConnection) Resume() {
	sel := objc.RegisterName("resume")
	x_.ID.Send(sel)
}
// Add a barrier block to execute on the connection. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NSXPCConnection/scheduleSendBarrierBlock(_:)
func (x_ XPCConnection) ScheduleSendBarrierBlock(block unsafe.Pointer) {
	sel := objc.RegisterName("scheduleSendBarrierBlock:")
	x_.ID.Send(sel, block)
}
// Sets the code signing requirement for this connection. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NSXPCConnection/setCodeSigningRequirement(_:)
func (x_ XPCConnection) SetCodeSigningRequirement(requirement string) {
	sel := objc.RegisterName("setCodeSigningRequirement:")
	x_.ID.Send(sel, requirement)
}
// Suspends the connection. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NSXPCConnection/suspend()
func (x_ XPCConnection) Suspend() {
	sel := objc.RegisterName("suspend")
	x_.ID.Send(sel)
}
//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NSXPCConnection/synchronousRemoteObjectProxyWithErrorHandler(_:)
func (x_ XPCConnection) SynchronousRemoteObjectProxyWithErrorHandler(handler unsafe.Pointer) objc.ID {
	sel := objc.RegisterName("synchronousRemoteObjectProxyWithErrorHandler:")
	ret := x_.ID.Send(sel, handler)
	return ret
}

