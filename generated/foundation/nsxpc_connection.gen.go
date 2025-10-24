// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class NSXPCConnection */


/* debug [class_header]: Header for NSXPCConnection */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for XPCConnection */
// An interface definition for the [XPCConnection] class.
type IXPCConnection interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for XPCConnection */
	// properties:
	AuditSessionIdentifier() objectivec.IObject
	EffectiveGroupIdentifier() objectivec.IObject
	EffectiveUserIdentifier() objectivec.IObject
	Endpoint() IXPCListenerEndpoint
	ExportedInterface() IXPCInterface
	SetExportedInterface(value IXPCInterface)
	ExportedObject() objc.ID
	SetExportedObject(value objc.ID)
	InterruptionHandler() unsafe.Pointer
	SetInterruptionHandler(value unsafe.Pointer)
	InvalidationHandler() unsafe.Pointer
	SetInvalidationHandler(value unsafe.Pointer)
	ProcessIdentifier() objectivec.IObject
	RemoteObjectInterface() IXPCInterface
	SetRemoteObjectInterface(value IXPCInterface)
	RemoteObjectProxy() objc.ID
	ServiceName() IString
	NSXPCConnectionCodeSigningRequirementFailure() int
	SetNSXPCConnectionCodeSigningRequirementFailure(value int)
	NSXPCConnectionErrorMaximum() int
	SetNSXPCConnectionErrorMaximum(value int)
	NSXPCConnectionErrorMinimum() int
	SetNSXPCConnectionErrorMinimum(value int)
	NSXPCConnectionInterrupted() int
	SetNSXPCConnectionInterrupted(value int)
	NSXPCConnectionInvalid() int
	SetNSXPCConnectionInvalid(value int)
	NSXPCConnectionReplyInvalid() int
	SetNSXPCConnectionReplyInvalid(value int)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for XPCConnection */
	// methods:
	Activate()
	Invalidate()
	RemoteObjectProxyWithErrorHandler(handler unsafe.Pointer) objc.ID
	Resume()
	ScheduleSendBarrierBlock(block unsafe.Pointer)
	SetCodeSigningRequirement(requirement IString)
	Suspend()
	SynchronousRemoteObjectProxyWithErrorHandler(handler unsafe.Pointer) objc.ID
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for XPCConnection */
// Alloc allocates a new instance without initialization.
func (xc _XPCConnectionClass) Alloc() XPCConnection {
	rv := objc.Send[XPCConnection](objc.ID(xc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for XPCConnection */
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for XPCConnection */

// Initializes an object to connect to an object in another process, identified by an object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSXPCConnection/init(listenerEndpoint:)
func NewXPCConnectionWithListenerEndpoint(endpoint IXPCListenerEndpoint) XPCConnection {
	instance := getXPCConnectionClass().Alloc()
	rv := objc.Send[XPCConnection](instance.ID, objc.Sel("initWithListenerEndpoint:"), endpoint)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewXPCConnectionWithListenerEndpoint */


// Initializes an object to connect to a LaunchAgent or LaunchDaemon with a name advertised in a .
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSXPCConnection/init(machServiceName:options:)
func NewXPCConnectionWithMachServiceNameOptions(name IString, options XPCConnectionOptions) XPCConnection {
	instance := getXPCConnectionClass().Alloc()
	rv := objc.Send[XPCConnection](instance.ID, objc.Sel("initWithMachServiceName:options:"), name, options)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewXPCConnectionWithMachServiceNameOptions */


// Initializes an object to connect to an object in an XPC service, identified by a service name.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSXPCConnection/init(serviceName:)
func NewXPCConnectionWithServiceName(serviceName IString) XPCConnection {
	instance := getXPCConnectionClass().Alloc()
	rv := objc.Send[XPCConnection](instance.ID, objc.Sel("initWithServiceName:"), serviceName)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewXPCConnectionWithServiceName */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for XPCConnection */

// Returns the current connection, in the context of a call to a method on your exported object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSXPCConnection/current()
func (xc _XPCConnectionClass) CurrentConnection() IXPCConnection {
	rv := objc.Send[XPCConnection](objc.ID(xc.class), objc.Sel("currentConnection"))
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=CurrentConnection) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for XPCConnection */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for XPCConnection */

// Activates the connection.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSXPCConnection/activate()
func (x_ XPCConnection) Activate() {
	objc.Send[objc.ID](x_.ID, objc.Sel("activate"))
}/* debug [instance_methods/method]: Activate */


// Invalidates the connection.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSXPCConnection/invalidate()
func (x_ XPCConnection) Invalidate() {
	objc.Send[objc.ID](x_.ID, objc.Sel("invalidate"))
}/* debug [instance_methods/method]: Invalidate */


// Returns a proxy for the remote object (that is, the object exported from the other side of this connection) with the specified error handler.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSXPCConnection/remoteObjectProxyWithErrorHandler(_:)
func (x_ XPCConnection) RemoteObjectProxyWithErrorHandler(handler unsafe.Pointer) objc.ID {
	rv := objc.Send[objc.ID](x_.ID, objc.Sel("remoteObjectProxyWithErrorHandler:"), handler)
	return rv
}/* debug [instance_methods/method]: RemoteObjectProxyWithErrorHandler */


// Starts or resumes handling of messages on a connection.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSXPCConnection/resume()
func (x_ XPCConnection) Resume() {
	objc.Send[objc.ID](x_.ID, objc.Sel("resume"))
}/* debug [instance_methods/method]: Resume */


// Add a barrier block to execute on the connection.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSXPCConnection/scheduleSendBarrierBlock(_:)
func (x_ XPCConnection) ScheduleSendBarrierBlock(block unsafe.Pointer) {
	objc.Send[objc.ID](x_.ID, objc.Sel("scheduleSendBarrierBlock:"), block)
}/* debug [instance_methods/method]: ScheduleSendBarrierBlock */


// Sets the code signing requirement for this connection.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSXPCConnection/setCodeSigningRequirement(_:)
func (x_ XPCConnection) SetCodeSigningRequirement(requirement IString) {
	objc.Send[objc.ID](x_.ID, objc.Sel("setCodeSigningRequirement:"), requirement)
}/* debug [instance_methods/method]: SetCodeSigningRequirement */


// Suspends the connection.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSXPCConnection/suspend()
func (x_ XPCConnection) Suspend() {
	objc.Send[objc.ID](x_.ID, objc.Sel("suspend"))
}/* debug [instance_methods/method]: Suspend */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSXPCConnection/synchronousRemoteObjectProxyWithErrorHandler(_:)
func (x_ XPCConnection) SynchronousRemoteObjectProxyWithErrorHandler(handler unsafe.Pointer) objc.ID {
	rv := objc.Send[objc.ID](x_.ID, objc.Sel("synchronousRemoteObjectProxyWithErrorHandler:"), handler)
	return rv
}/* debug [instance_methods/method]: SynchronousRemoteObjectProxyWithErrorHandler */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for XPCConnection */

// The BSM audit session identifier for the connecting process.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSXPCConnection/auditSessionIdentifier
func (x_ XPCConnection) AuditSessionIdentifier() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](x_.ID, objc.Sel("auditSessionIdentifier"))
	return rv
}/* debug [instance_properties/getter]: auditSessionIdentifier */


// The effective group ID (EGID) of the connecting process.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSXPCConnection/effectiveGroupIdentifier
func (x_ XPCConnection) EffectiveGroupIdentifier() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](x_.ID, objc.Sel("effectiveGroupIdentifier"))
	return rv
}/* debug [instance_properties/getter]: effectiveGroupIdentifier */


// The effective user ID (EUID) of the connecting process.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSXPCConnection/effectiveUserIdentifier
func (x_ XPCConnection) EffectiveUserIdentifier() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](x_.ID, objc.Sel("effectiveUserIdentifier"))
	return rv
}/* debug [instance_properties/getter]: effectiveUserIdentifier */


// If the connection was created with an object, returns the endpoint object used.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSXPCConnection/endpoint
func (x_ XPCConnection) Endpoint() IXPCListenerEndpoint {
	rv := objc.Send[XPCListenerEndpoint](x_.ID, objc.Sel("endpoint"))
	return rv
}/* debug [instance_properties/getter]: endpoint */


// The object that describes the protocol for the exported object on this connection.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSXPCConnection/exportedInterface
func (x_ XPCConnection) ExportedInterface() IXPCInterface {
	rv := objc.Send[XPCInterface](x_.ID, objc.Sel("exportedInterface"))
	return rv
}/* debug [instance_properties/getter]: exportedInterface */


// The object that describes the protocol for the exported object on this connection.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSXPCConnection/exportedInterface
func (x_ XPCConnection) SetExportedInterface(value IXPCInterface) {
	objc.Send[objc.ID](x_.ID, objc.Sel("setExportedInterface:"), value)
}/* debug [instance_properties/setter]: exportedInterface */


// An exported object for the connection.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSXPCConnection/exportedObject
func (x_ XPCConnection) ExportedObject() objc.ID {
	rv := objc.Send[objc.ID](x_.ID, objc.Sel("exportedObject"))
	return rv
}/* debug [instance_properties/getter]: exportedObject */


// An exported object for the connection.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSXPCConnection/exportedObject
func (x_ XPCConnection) SetExportedObject(value objc.ID) {
	objc.Send[objc.ID](x_.ID, objc.Sel("setExportedObject:"), value)
}/* debug [instance_properties/setter]: exportedObject */


// An interruption handler that is called if the remote process exits or crashes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSXPCConnection/interruptionHandler
func (x_ XPCConnection) InterruptionHandler() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](x_.ID, objc.Sel("interruptionHandler"))
	return rv
}/* debug [instance_properties/getter]: interruptionHandler */


// An interruption handler that is called if the remote process exits or crashes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSXPCConnection/interruptionHandler
func (x_ XPCConnection) SetInterruptionHandler(value unsafe.Pointer) {
	objc.Send[objc.ID](x_.ID, objc.Sel("setInterruptionHandler:"), value)
}/* debug [instance_properties/setter]: interruptionHandler */


// An invalidation handler that is called if the connection can not be formed or the connection has terminated and may not be re-established.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSXPCConnection/invalidationHandler
func (x_ XPCConnection) InvalidationHandler() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](x_.ID, objc.Sel("invalidationHandler"))
	return rv
}/* debug [instance_properties/getter]: invalidationHandler */


// An invalidation handler that is called if the connection can not be formed or the connection has terminated and may not be re-established.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSXPCConnection/invalidationHandler
func (x_ XPCConnection) SetInvalidationHandler(value unsafe.Pointer) {
	objc.Send[objc.ID](x_.ID, objc.Sel("setInvalidationHandler:"), value)
}/* debug [instance_properties/setter]: invalidationHandler */


// The process ID (PID) of the connecting process.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSXPCConnection/processIdentifier
func (x_ XPCConnection) ProcessIdentifier() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](x_.ID, objc.Sel("processIdentifier"))
	return rv
}/* debug [instance_properties/getter]: processIdentifier */


// Defines the object that describes the protocol for the object represented by the .
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSXPCConnection/remoteObjectInterface
func (x_ XPCConnection) RemoteObjectInterface() IXPCInterface {
	rv := objc.Send[XPCInterface](x_.ID, objc.Sel("remoteObjectInterface"))
	return rv
}/* debug [instance_properties/getter]: remoteObjectInterface */


// Defines the object that describes the protocol for the object represented by the .
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSXPCConnection/remoteObjectInterface
func (x_ XPCConnection) SetRemoteObjectInterface(value IXPCInterface) {
	objc.Send[objc.ID](x_.ID, objc.Sel("setRemoteObjectInterface:"), value)
}/* debug [instance_properties/setter]: remoteObjectInterface */


// Returns a proxy for the remote object (that is, the from the other side of this connection).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSXPCConnection/remoteObjectProxy
func (x_ XPCConnection) RemoteObjectProxy() objc.ID {
	rv := objc.Send[objc.ID](x_.ID, objc.Sel("remoteObjectProxy"))
	return rv
}/* debug [instance_properties/getter]: remoteObjectProxy */


// The name of the XPC service that this connection was configured to connect to.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSXPCConnection/serviceName
func (x_ XPCConnection) ServiceName() IString {
	rv := objc.Send[String](x_.ID, objc.Sel("serviceName"))
	return rv
}/* debug [instance_properties/getter]: serviceName */


// A code-signing requirement check failed.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsxpcconnectioncodesigningrequirementfailure-swift.var
func (x_ XPCConnection) NSXPCConnectionCodeSigningRequirementFailure() int {
	rv := objc.Send[int](x_.ID, objc.Sel("NSXPCConnectionCodeSigningRequirementFailure"))
	return rv
}/* debug [instance_properties/getter]: NSXPCConnectionCodeSigningRequirementFailure */


// A code-signing requirement check failed.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsxpcconnectioncodesigningrequirementfailure-swift.var
func (x_ XPCConnection) SetNSXPCConnectionCodeSigningRequirementFailure(value int) {
	objc.Send[objc.ID](x_.ID, objc.Sel("setNSXPCConnectionCodeSigningRequirementFailure:"), value)
}/* debug [instance_properties/setter]: NSXPCConnectionCodeSigningRequirementFailure */


// The upper bounds of XPC connection error code values.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsxpcconnectionerrormaximum-swift.var
func (x_ XPCConnection) NSXPCConnectionErrorMaximum() int {
	rv := objc.Send[int](x_.ID, objc.Sel("NSXPCConnectionErrorMaximum"))
	return rv
}/* debug [instance_properties/getter]: NSXPCConnectionErrorMaximum */


// The upper bounds of XPC connection error code values.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsxpcconnectionerrormaximum-swift.var
func (x_ XPCConnection) SetNSXPCConnectionErrorMaximum(value int) {
	objc.Send[objc.ID](x_.ID, objc.Sel("setNSXPCConnectionErrorMaximum:"), value)
}/* debug [instance_properties/setter]: NSXPCConnectionErrorMaximum */


// The lower bounds of XPC connection error code values.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsxpcconnectionerrorminimum-swift.var
func (x_ XPCConnection) NSXPCConnectionErrorMinimum() int {
	rv := objc.Send[int](x_.ID, objc.Sel("NSXPCConnectionErrorMinimum"))
	return rv
}/* debug [instance_properties/getter]: NSXPCConnectionErrorMinimum */


// The lower bounds of XPC connection error code values.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsxpcconnectionerrorminimum-swift.var
func (x_ XPCConnection) SetNSXPCConnectionErrorMinimum(value int) {
	objc.Send[objc.ID](x_.ID, objc.Sel("setNSXPCConnectionErrorMinimum:"), value)
}/* debug [instance_properties/setter]: NSXPCConnectionErrorMinimum */


// The XPC connection was interrupted.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsxpcconnectioninterrupted-swift.var
func (x_ XPCConnection) NSXPCConnectionInterrupted() int {
	rv := objc.Send[int](x_.ID, objc.Sel("NSXPCConnectionInterrupted"))
	return rv
}/* debug [instance_properties/getter]: NSXPCConnectionInterrupted */


// The XPC connection was interrupted.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsxpcconnectioninterrupted-swift.var
func (x_ XPCConnection) SetNSXPCConnectionInterrupted(value int) {
	objc.Send[objc.ID](x_.ID, objc.Sel("setNSXPCConnectionInterrupted:"), value)
}/* debug [instance_properties/setter]: NSXPCConnectionInterrupted */


// The XPC connection was invalid.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsxpcconnectioninvalid-swift.var
func (x_ XPCConnection) NSXPCConnectionInvalid() int {
	rv := objc.Send[int](x_.ID, objc.Sel("NSXPCConnectionInvalid"))
	return rv
}/* debug [instance_properties/getter]: NSXPCConnectionInvalid */


// The XPC connection was invalid.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsxpcconnectioninvalid-swift.var
func (x_ XPCConnection) SetNSXPCConnectionInvalid(value int) {
	objc.Send[objc.ID](x_.ID, objc.Sel("setNSXPCConnectionInvalid:"), value)
}/* debug [instance_properties/setter]: NSXPCConnectionInvalid */


// The XPC connection reply was invalid.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsxpcconnectionreplyinvalid-swift.var
func (x_ XPCConnection) NSXPCConnectionReplyInvalid() int {
	rv := objc.Send[int](x_.ID, objc.Sel("NSXPCConnectionReplyInvalid"))
	return rv
}/* debug [instance_properties/getter]: NSXPCConnectionReplyInvalid */


// The XPC connection reply was invalid.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsxpcconnectionreplyinvalid-swift.var
func (x_ XPCConnection) SetNSXPCConnectionReplyInvalid(value int) {
	objc.Send[objc.ID](x_.ID, objc.Sel("setNSXPCConnectionReplyInvalid:"), value)
}/* debug [instance_properties/setter]: NSXPCConnectionReplyInvalid */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class NSXPCConnection */


