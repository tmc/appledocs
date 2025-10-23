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
	// properties:
	AuditSessionIdentifier() unsafe.Pointer
	EffectiveGroupIdentifier() unsafe.Pointer
	EffectiveUserIdentifier() unsafe.Pointer
	Endpoint() IXPCListenerEndpoint
	ExportedInterface() IXPCInterface
	SetExportedInterface(value IXPCInterface)
	ExportedObject() objc.ID
	SetExportedObject(value objc.ID)
	InterruptionHandler() unsafe.Pointer
	SetInterruptionHandler(value unsafe.Pointer)
	InvalidationHandler() unsafe.Pointer
	SetInvalidationHandler(value unsafe.Pointer)
	ProcessIdentifier() unsafe.Pointer
	RemoteObjectInterface() IXPCInterface
	SetRemoteObjectInterface(value IXPCInterface)
	RemoteObjectProxy() objc.ID
	ServiceName() IString
	NSXPCConnectionCodeSigningRequirementFailure() int /* primitive/slice/pointer. */
	SetNSXPCConnectionCodeSigningRequirementFailure(value int /* primitive/slice/pointer. */)
	NSXPCConnectionErrorMaximum() int /* primitive/slice/pointer. */
	SetNSXPCConnectionErrorMaximum(value int /* primitive/slice/pointer. */)
	NSXPCConnectionErrorMinimum() int /* primitive/slice/pointer. */
	SetNSXPCConnectionErrorMinimum(value int /* primitive/slice/pointer. */)
	NSXPCConnectionInterrupted() int /* primitive/slice/pointer. */
	SetNSXPCConnectionInterrupted(value int /* primitive/slice/pointer. */)
	NSXPCConnectionInvalid() int /* primitive/slice/pointer. */
	SetNSXPCConnectionInvalid(value int /* primitive/slice/pointer. */)
	NSXPCConnectionReplyInvalid() int /* primitive/slice/pointer. */
	SetNSXPCConnectionReplyInvalid(value int /* primitive/slice/pointer. */)
	// methods:
	Activate()
	Invalidate()
	RemoteObjectProxyWithErrorHandler(handler unsafe.Pointer) objc.ID
	Resume()
	ScheduleSendBarrierBlock(block unsafe.Pointer)
	SetCodeSigningRequirement(requirement IString)
	Suspend()
	SynchronousRemoteObjectProxyWithErrorHandler(handler unsafe.Pointer) objc.ID
}

// A bidirectional communication channel between two processes.
//
// This class is the primary means of creating and configuring the communication mechanism between two processes. Each process has one instance of this class to represent the endpoint in the communication channel.


// A bidirectional communication channel between two processes.
//
// [Full Topic]
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
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSXPCConnection/init(listenerEndpoint:)
func NewXPCConnectionWithListenerEndpoint(endpoint IXPCListenerEndpoint) XPCConnection {
	instance := getXPCConnectionClass().Alloc()
	rv := objc.Send[XPCConnection](instance.ID, objc.Sel("initWithListenerEndpoint:"), endpoint)
	rv.Autorelease()
	return rv
}


// Initializes an object to connect to a LaunchAgent or LaunchDaemon with a name advertised in a .
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSXPCConnection/init(machServiceName:options:)
func NewXPCConnectionWithMachServiceNameOptions(name IString, options XPCConnectionOptions) XPCConnection {
	instance := getXPCConnectionClass().Alloc()
	rv := objc.Send[XPCConnection](instance.ID, objc.Sel("initWithMachServiceName:options:"), name, options)
	rv.Autorelease()
	return rv
}


// Initializes an object to connect to an object in an XPC service, identified by a service name.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSXPCConnection/init(serviceName:)
func NewXPCConnectionWithServiceName(serviceName IString) XPCConnection {
	instance := getXPCConnectionClass().Alloc()
	rv := objc.Send[XPCConnection](instance.ID, objc.Sel("initWithServiceName:"), serviceName)
	rv.Autorelease()
	return rv
}



// Returns the current connection, in the context of a call to a method on your exported object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSXPCConnection/current()
func (xc _XPCConnectionClass) CurrentConnection() IXPCConnection {
	rv := objc.Send[XPCConnection](objc.ID(xc.class), objc.Sel("currentConnection"))
	return rv
}


// Activates the connection.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSXPCConnection/activate()
func (x_ XPCConnection) Activate() {
	objc.Send[objc.ID](x_.ID, objc.Sel("activate"))
}


// Invalidates the connection.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSXPCConnection/invalidate()
func (x_ XPCConnection) Invalidate() {
	objc.Send[objc.ID](x_.ID, objc.Sel("invalidate"))
}


// Returns a proxy for the remote object (that is, the object exported from the other side of this connection) with the specified error handler.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSXPCConnection/remoteObjectProxyWithErrorHandler(_:)
func (x_ XPCConnection) RemoteObjectProxyWithErrorHandler(handler unsafe.Pointer) objc.ID {
	rv := objc.Send[objc.ID](x_.ID, objc.Sel("remoteObjectProxyWithErrorHandler:"), handler)
	return rv
}


// Starts or resumes handling of messages on a connection.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSXPCConnection/resume()
func (x_ XPCConnection) Resume() {
	objc.Send[objc.ID](x_.ID, objc.Sel("resume"))
}


// Add a barrier block to execute on the connection.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSXPCConnection/scheduleSendBarrierBlock(_:)
func (x_ XPCConnection) ScheduleSendBarrierBlock(block unsafe.Pointer) {
	objc.Send[objc.ID](x_.ID, objc.Sel("scheduleSendBarrierBlock:"), block)
}


// Sets the code signing requirement for this connection.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSXPCConnection/setCodeSigningRequirement(_:)
func (x_ XPCConnection) SetCodeSigningRequirement(requirement IString) {
	objc.Send[objc.ID](x_.ID, objc.Sel("setCodeSigningRequirement:"), requirement)
}


// Suspends the connection.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSXPCConnection/suspend()
func (x_ XPCConnection) Suspend() {
	objc.Send[objc.ID](x_.ID, objc.Sel("suspend"))
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSXPCConnection/synchronousRemoteObjectProxyWithErrorHandler(_:)
func (x_ XPCConnection) SynchronousRemoteObjectProxyWithErrorHandler(handler unsafe.Pointer) objc.ID {
	rv := objc.Send[objc.ID](x_.ID, objc.Sel("synchronousRemoteObjectProxyWithErrorHandler:"), handler)
	return rv
}


// The BSM audit session identifier for the connecting process.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSXPCConnection/auditSessionIdentifier
func (x_ XPCConnection) AuditSessionIdentifier() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](x_.ID, objc.Sel("auditSessionIdentifier"))
	return rv
}


// The effective group ID (EGID) of the connecting process.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSXPCConnection/effectiveGroupIdentifier
func (x_ XPCConnection) EffectiveGroupIdentifier() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](x_.ID, objc.Sel("effectiveGroupIdentifier"))
	return rv
}


// The effective user ID (EUID) of the connecting process.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSXPCConnection/effectiveUserIdentifier
func (x_ XPCConnection) EffectiveUserIdentifier() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](x_.ID, objc.Sel("effectiveUserIdentifier"))
	return rv
}


// If the connection was created with an object, returns the endpoint object used.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSXPCConnection/endpoint
func (x_ XPCConnection) Endpoint() IXPCListenerEndpoint {
	rv := objc.Send[XPCListenerEndpoint](x_.ID, objc.Sel("endpoint"))
	return rv
}


// The object that describes the protocol for the exported object on this connection.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSXPCConnection/exportedInterface
func (x_ XPCConnection) ExportedInterface() IXPCInterface {
	rv := objc.Send[XPCInterface](x_.ID, objc.Sel("exportedInterface"))
	return rv
}


// The object that describes the protocol for the exported object on this connection.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSXPCConnection/exportedInterface
func (x_ XPCConnection) SetExportedInterface(value IXPCInterface) {
	objc.Send[objc.ID](x_.ID, objc.Sel("setExportedInterface:"), value)
}


// An exported object for the connection.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSXPCConnection/exportedObject
func (x_ XPCConnection) ExportedObject() objc.ID {
	rv := objc.Send[objc.ID](x_.ID, objc.Sel("exportedObject"))
	return rv
}


// An exported object for the connection.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSXPCConnection/exportedObject
func (x_ XPCConnection) SetExportedObject(value objc.ID) {
	objc.Send[objc.ID](x_.ID, objc.Sel("setExportedObject:"), value)
}


// An interruption handler that is called if the remote process exits or crashes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSXPCConnection/interruptionHandler
func (x_ XPCConnection) InterruptionHandler() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](x_.ID, objc.Sel("interruptionHandler"))
	return rv
}


// An interruption handler that is called if the remote process exits or crashes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSXPCConnection/interruptionHandler
func (x_ XPCConnection) SetInterruptionHandler(value unsafe.Pointer) {
	objc.Send[objc.ID](x_.ID, objc.Sel("setInterruptionHandler:"), value)
}


// An invalidation handler that is called if the connection can not be formed or the connection has terminated and may not be re-established.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSXPCConnection/invalidationHandler
func (x_ XPCConnection) InvalidationHandler() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](x_.ID, objc.Sel("invalidationHandler"))
	return rv
}


// An invalidation handler that is called if the connection can not be formed or the connection has terminated and may not be re-established.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSXPCConnection/invalidationHandler
func (x_ XPCConnection) SetInvalidationHandler(value unsafe.Pointer) {
	objc.Send[objc.ID](x_.ID, objc.Sel("setInvalidationHandler:"), value)
}


// The process ID (PID) of the connecting process.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSXPCConnection/processIdentifier
func (x_ XPCConnection) ProcessIdentifier() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](x_.ID, objc.Sel("processIdentifier"))
	return rv
}


// Defines the object that describes the protocol for the object represented by the .
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSXPCConnection/remoteObjectInterface
func (x_ XPCConnection) RemoteObjectInterface() IXPCInterface {
	rv := objc.Send[XPCInterface](x_.ID, objc.Sel("remoteObjectInterface"))
	return rv
}


// Defines the object that describes the protocol for the object represented by the .
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSXPCConnection/remoteObjectInterface
func (x_ XPCConnection) SetRemoteObjectInterface(value IXPCInterface) {
	objc.Send[objc.ID](x_.ID, objc.Sel("setRemoteObjectInterface:"), value)
}


// Returns a proxy for the remote object (that is, the from the other side of this connection).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSXPCConnection/remoteObjectProxy
func (x_ XPCConnection) RemoteObjectProxy() objc.ID {
	rv := objc.Send[objc.ID](x_.ID, objc.Sel("remoteObjectProxy"))
	return rv
}


// The name of the XPC service that this connection was configured to connect to.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSXPCConnection/serviceName
func (x_ XPCConnection) ServiceName() IString {
	rv := objc.Send[String](x_.ID, objc.Sel("serviceName"))
	return rv
}


// A code-signing requirement check failed.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsxpcconnectioncodesigningrequirementfailure-swift.var
func (x_ XPCConnection) NSXPCConnectionCodeSigningRequirementFailure() int /* primitive/slice/pointer. */ {
	rv := objc.Send[int](x_.ID, objc.Sel("NSXPCConnectionCodeSigningRequirementFailure"))
	return rv
}


// A code-signing requirement check failed.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsxpcconnectioncodesigningrequirementfailure-swift.var
func (x_ XPCConnection) SetNSXPCConnectionCodeSigningRequirementFailure(value int /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](x_.ID, objc.Sel("setNSXPCConnectionCodeSigningRequirementFailure:"), value)
}


// The upper bounds of XPC connection error code values.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsxpcconnectionerrormaximum-swift.var
func (x_ XPCConnection) NSXPCConnectionErrorMaximum() int /* primitive/slice/pointer. */ {
	rv := objc.Send[int](x_.ID, objc.Sel("NSXPCConnectionErrorMaximum"))
	return rv
}


// The upper bounds of XPC connection error code values.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsxpcconnectionerrormaximum-swift.var
func (x_ XPCConnection) SetNSXPCConnectionErrorMaximum(value int /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](x_.ID, objc.Sel("setNSXPCConnectionErrorMaximum:"), value)
}


// The lower bounds of XPC connection error code values.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsxpcconnectionerrorminimum-swift.var
func (x_ XPCConnection) NSXPCConnectionErrorMinimum() int /* primitive/slice/pointer. */ {
	rv := objc.Send[int](x_.ID, objc.Sel("NSXPCConnectionErrorMinimum"))
	return rv
}


// The lower bounds of XPC connection error code values.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsxpcconnectionerrorminimum-swift.var
func (x_ XPCConnection) SetNSXPCConnectionErrorMinimum(value int /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](x_.ID, objc.Sel("setNSXPCConnectionErrorMinimum:"), value)
}


// The XPC connection was interrupted.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsxpcconnectioninterrupted-swift.var
func (x_ XPCConnection) NSXPCConnectionInterrupted() int /* primitive/slice/pointer. */ {
	rv := objc.Send[int](x_.ID, objc.Sel("NSXPCConnectionInterrupted"))
	return rv
}


// The XPC connection was interrupted.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsxpcconnectioninterrupted-swift.var
func (x_ XPCConnection) SetNSXPCConnectionInterrupted(value int /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](x_.ID, objc.Sel("setNSXPCConnectionInterrupted:"), value)
}


// The XPC connection was invalid.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsxpcconnectioninvalid-swift.var
func (x_ XPCConnection) NSXPCConnectionInvalid() int /* primitive/slice/pointer. */ {
	rv := objc.Send[int](x_.ID, objc.Sel("NSXPCConnectionInvalid"))
	return rv
}


// The XPC connection was invalid.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsxpcconnectioninvalid-swift.var
func (x_ XPCConnection) SetNSXPCConnectionInvalid(value int /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](x_.ID, objc.Sel("setNSXPCConnectionInvalid:"), value)
}


// The XPC connection reply was invalid.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsxpcconnectionreplyinvalid-swift.var
func (x_ XPCConnection) NSXPCConnectionReplyInvalid() int /* primitive/slice/pointer. */ {
	rv := objc.Send[int](x_.ID, objc.Sel("NSXPCConnectionReplyInvalid"))
	return rv
}


// The XPC connection reply was invalid.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsxpcconnectionreplyinvalid-swift.var
func (x_ XPCConnection) SetNSXPCConnectionReplyInvalid(value int /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](x_.ID, objc.Sel("setNSXPCConnectionReplyInvalid:"), value)
}


