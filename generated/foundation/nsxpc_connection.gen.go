// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [XPCConnection] class.
var (
	XPCConnectionClass     _XPCConnectionClass
	XPCConnectionClassOnce sync.Once
)

func getXPCConnectionClass() _XPCConnectionClass {
	XPCConnectionClassOnce.Do(func() {
		XPCConnectionClass = _XPCConnectionClass{objc.GetClass("NSXPCConnection")}
	})
	return XPCConnectionClass
}

type _XPCConnectionClass struct {
	class objc.Class
}

// An interface definition for the [XPCConnection] class.
type IXPCConnection interface {
	objectivec.IObject
	Activate()
	Invalidate()
	RemoteObjectProxyWithErrorHandler(handler unsafe.Pointer) objc.ID
	Resume()
	ScheduleSendBarrierBlock(block unsafe.Pointer)
	SetCodeSigningRequirement(requirement string)
	Suspend()
	SynchronousRemoteObjectProxyWithErrorHandler(handler unsafe.Pointer) objc.ID
}

// A bidirectional communication channel between two processes.
//
// This class is the primary means of creating and configuring the communication mechanism between two processes. Each process has one instance of this class to represent the endpoint in the communication channel.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSXPCConnection
type XPCConnection struct {
	objectivec.Object
}

// XPCConnectionFrom constructs a [XPCConnection] from an unsafe.Pointer.
//
// A bidirectional communication channel between two processes.
func XPCConnectionFrom(ptr unsafe.Pointer) XPCConnection {
	return XPCConnection{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (xc _XPCConnectionClass) Alloc() XPCConnection {
	rv := objc.Send[XPCConnection](objc.ID(xc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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
	return getXPCConnectionClass().New()
}

// Initializes an object to connect to an object in another process, identified by an object.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSXPCConnection/init(listenerEndpoint:)
func NewXPCConnectionWithListenerEndpoint(endpoint unsafe.Pointer) XPCConnection {
	instance := getXPCConnectionClass().Alloc()
	rv := objc.Send[XPCConnection](instance.ID, objc.Sel("initWithListenerEndpoint:"), endpoint)
	rv.Autorelease()
	return rv
}

// Initializes an object to connect to a LaunchAgent or LaunchDaemon with a name advertised in a .
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSXPCConnection/init(machServiceName:options:)
func NewXPCConnectionWithMachServiceNameOptions(name string, options unsafe.Pointer) XPCConnection {
	instance := getXPCConnectionClass().Alloc()
	rv := objc.Send[XPCConnection](instance.ID, objc.Sel("initWithMachServiceName:options:"), objc.String(name), options)
	rv.Autorelease()
	return rv
}

// Initializes an object to connect to an object in an XPC service, identified by a service name.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSXPCConnection/init(serviceName:)
func NewXPCConnectionWithServiceName(serviceName string) XPCConnection {
	instance := getXPCConnectionClass().Alloc()
	rv := objc.Send[XPCConnection](instance.ID, objc.Sel("initWithServiceName:"), objc.String(serviceName))
	rv.Autorelease()
	return rv
}

// Returns the current connection, in the context of a call to a method on your exported object.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSXPCConnection/current()
func (xc _XPCConnectionClass) CurrentConnection() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(xc.class), objc.Sel("currentConnection"))
	return rv
}

// Activates the connection.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSXPCConnection/activate()
func (x_ XPCConnection) Activate() {
	objc.Send[objc.ID](x_.ID, objc.Sel("activate"))
}

// Invalidates the connection.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSXPCConnection/invalidate()
func (x_ XPCConnection) Invalidate() {
	objc.Send[objc.ID](x_.ID, objc.Sel("invalidate"))
}

// Returns a proxy for the remote object (that is, the object exported from the other side of this connection) with the specified error handler.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSXPCConnection/remoteObjectProxyWithErrorHandler(_:)
func (x_ XPCConnection) RemoteObjectProxyWithErrorHandler(handler unsafe.Pointer) objc.ID {
	rv := objc.Send[objc.ID](x_.ID, objc.Sel("remoteObjectProxyWithErrorHandler:"), handler)
	return rv
}

// Starts or resumes handling of messages on a connection.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSXPCConnection/resume()
func (x_ XPCConnection) Resume() {
	objc.Send[objc.ID](x_.ID, objc.Sel("resume"))
}

// Add a barrier block to execute on the connection.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSXPCConnection/scheduleSendBarrierBlock(_:)
func (x_ XPCConnection) ScheduleSendBarrierBlock(block unsafe.Pointer) {
	objc.Send[objc.ID](x_.ID, objc.Sel("scheduleSendBarrierBlock:"), block)
}

// Sets the code signing requirement for this connection.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSXPCConnection/setCodeSigningRequirement(_:)
func (x_ XPCConnection) SetCodeSigningRequirement(requirement string) {
	objc.Send[objc.ID](x_.ID, objc.Sel("setCodeSigningRequirement:"), objc.String(requirement))
}

// Suspends the connection.
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

// The BSM audit session identifier for the connecting process.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSXPCConnection/auditSessionIdentifier
func (x_ XPCConnection) AuditSessionIdentifier() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](x_.ID, objc.Sel("auditSessionIdentifier"))
	return rv
}

// The effective group ID (EGID) of the connecting process.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSXPCConnection/effectiveGroupIdentifier
func (x_ XPCConnection) EffectiveGroupIdentifier() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](x_.ID, objc.Sel("effectiveGroupIdentifier"))
	return rv
}

// The effective user ID (EUID) of the connecting process.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSXPCConnection/effectiveUserIdentifier
func (x_ XPCConnection) EffectiveUserIdentifier() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](x_.ID, objc.Sel("effectiveUserIdentifier"))
	return rv
}

// If the connection was created with an object, returns the endpoint object used.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSXPCConnection/endpoint
func (x_ XPCConnection) Endpoint() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](x_.ID, objc.Sel("endpoint"))
	return rv
}

// The object that describes the protocol for the exported object on this connection.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSXPCConnection/exportedInterface
func (x_ XPCConnection) ExportedInterface() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](x_.ID, objc.Sel("exportedInterface"))
	return rv
}

// SetExportedInterface sets the value of the exportedInterface property.
// The object that describes the protocol for the exported object on this connection.

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSXPCConnection/exportedInterface
func (x_ XPCConnection) SetExportedInterface(value unsafe.Pointer) {
	objc.Send[objc.ID](x_.ID, objc.Sel("setExportedInterface:"), value)
}

// An exported object for the connection.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSXPCConnection/exportedObject
func (x_ XPCConnection) ExportedObject() objc.ID {
	rv := objc.Send[objc.ID](x_.ID, objc.Sel("exportedObject"))
	return rv
}

// SetExportedObject sets the value of the exportedObject property.
// An exported object for the connection.

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSXPCConnection/exportedObject
func (x_ XPCConnection) SetExportedObject(value objc.ID) {
	objc.Send[objc.ID](x_.ID, objc.Sel("setExportedObject:"), value)
}

// The process ID (PID) of the connecting process.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSXPCConnection/processIdentifier
func (x_ XPCConnection) ProcessIdentifier() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](x_.ID, objc.Sel("processIdentifier"))
	return rv
}

// Defines the object that describes the protocol for the object represented by the .
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSXPCConnection/remoteObjectInterface
func (x_ XPCConnection) RemoteObjectInterface() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](x_.ID, objc.Sel("remoteObjectInterface"))
	return rv
}

// SetRemoteObjectInterface sets the value of the remoteObjectInterface property.
// Defines the object that describes the protocol for the object represented by the .

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSXPCConnection/remoteObjectInterface
func (x_ XPCConnection) SetRemoteObjectInterface(value unsafe.Pointer) {
	objc.Send[objc.ID](x_.ID, objc.Sel("setRemoteObjectInterface:"), value)
}

// Returns a proxy for the remote object (that is, the from the other side of this connection).
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSXPCConnection/remoteObjectProxy
func (x_ XPCConnection) RemoteObjectProxy() objc.ID {
	rv := objc.Send[objc.ID](x_.ID, objc.Sel("remoteObjectProxy"))
	return rv
}

// The name of the XPC service that this connection was configured to connect to.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSXPCConnection/serviceName
func (x_ XPCConnection) ServiceName() string {
	rv := objc.Send[string](x_.ID, objc.Sel("serviceName"))
	return rv
}
