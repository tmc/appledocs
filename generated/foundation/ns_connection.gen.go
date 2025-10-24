// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [Connection] class.
var (
	ConnectionClass     _ConnectionClass
	ConnectionClassOnce sync.Once
)

func getConnectionClass() _ConnectionClass {
	ConnectionClassOnce.Do(func() {
		ConnectionClass = _ConnectionClass{objc.GetClass("NSConnection")}
	})
	return ConnectionClass
}

type _ConnectionClass struct {
	class objc.Class
}

// An interface definition for the [Connection] class.
type IConnection interface {
	objectivec.IObject
	// properties:
	Delegate() objc.ID
	SetDelegate(value objc.ID)
	IndependentConversationQueueing() bool
	SetIndependentConversationQueueing(value bool)
	LocalObjects() IArray
	MultipleThreadsEnabled() bool
	ReceivePort() IPort
	RemoteObjects() IArray
	ReplyTimeout() float64
	SetReplyTimeout(value float64)
	RequestModes() []string
	RequestTimeout() float64
	SetRequestTimeout(value float64)
	RootObject() objc.ID
	SetRootObject(value objc.ID)
	RootProxy() IDistantObject
	SendPort() IPort
	Statistics() IDictionary
	Valid() bool
	// methods:
}

// An object that manages the communication between objects in different threads or between a thread and a process running on a local or remote system.
//
// Connection objects form the backbone of the distributed objects mechanism and normally operate in the background. You use the methods of explicitly when vending an object to other applications, when accessing such a vended object through a proxy, and when altering default communication parameters. At other times, you simply interact with a vended object or its proxy. A single connection object may be shared by multiple threads and used to access a vended object.


// An object that manages the communication between objects in different threads or between a thread and a process running on a local or remote system.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSConnection
type Connection struct {
	objectivec.Object
}

// ConnectionFrom constructs a [Connection] from an unsafe.Pointer.
//
// An object that manages the communication between objects in different threads or between a thread and a process running on a local or remote system.
func ConnectionFrom(ptr unsafe.Pointer) Connection {
	return Connection{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (cc _ConnectionClass) Alloc() Connection {
	rv := objc.Send[Connection](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (cc _ConnectionClass) New() Connection {
	rv := objc.Send[Connection](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ Connection) Init() Connection {
	rv := objc.Send[Connection](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ Connection) Autorelease() Connection {
	rv := objc.Send[Connection](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewConnection creates a new Connection instance.
func NewConnection() Connection {
	return getConnectionClass().New()
}



// Returns an object initialized with given send and receive ports.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSConnection/initWithReceivePort:sendPort:
func NewConnectionWithReceivePortSendPort(receivePort IPort, sendPort IPort) Connection {
	instance := getConnectionClass().Alloc()
	rv := objc.Send[Connection](instance.ID, objc.Sel("initWithReceivePort:sendPort:"), receivePort, sendPort)
	rv.Autorelease()
	return rv
}



// Returns all valid objects in the process.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSConnection/allConnections
func (cc _ConnectionClass) AllConnections() []Connection {
	rv := objc.Send[[]Connection](objc.ID(cc.class), objc.Sel("allConnections"))
	return rv
}


// Returns an object that communicates using given send and receive ports.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSConnection/connectionWithReceivePort:sendPort:
func (cc _ConnectionClass) ConnectionWithReceivePortSendPort(receivePort IPort, sendPort IPort) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(cc.class), objc.Sel("connectionWithReceivePort:sendPort:"), receivePort, sendPort)
	return rv
}


// Returns the object whose send port links it to the object registered with the default under a given name on a given host.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSConnection/connectionWithRegisteredName:host:
func (cc _ConnectionClass) ConnectionWithRegisteredNameHost(name IString, hostName IString) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(cc.class), objc.Sel("connectionWithRegisteredName:host:"), name, hostName)
	return rv
}


// Returns the object whose send port links it to the object registered under a given name with a given server on a given host.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSConnection/connectionWithRegisteredName:host:usingNameServer:
func (cc _ConnectionClass) ConnectionWithRegisteredNameHostUsingNameServer(name IString, hostName IString, server IPortNameServer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(cc.class), objc.Sel("connectionWithRegisteredName:host:usingNameServer:"), name, hostName, server)
	return rv
}


// Returns a token object representing any conversation in progress in the current thread.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSConnection/currentConversation
func (cc _ConnectionClass) CurrentConversation() objc.ID {
	rv := objc.Send[objc.ID](objc.ID(cc.class), objc.Sel("currentConversation"))
	return rv
}


// Returns the default object for the current thread.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSConnection/defaultConnection
func (cc _ConnectionClass) DefaultConnection() IConnection {
	rv := objc.Send[Connection](objc.ID(cc.class), objc.Sel("defaultConnection"))
	return rv
}


// Returns a proxy for the root object of the object registered with the default under a given name on a given host.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSConnection/rootProxyForConnectionWithRegisteredName:host:
func (cc _ConnectionClass) RootProxyForConnectionWithRegisteredNameHost(name IString, hostName IString) IDistantObject {
	rv := objc.Send[DistantObject](objc.ID(cc.class), objc.Sel("rootProxyForConnectionWithRegisteredName:host:"), name, hostName)
	return rv
}


// Returns a proxy for the root object of the object registered with under on a given host.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSConnection/rootProxyForConnectionWithRegisteredName:host:usingNameServer:
func (cc _ConnectionClass) RootProxyForConnectionWithRegisteredNameHostUsingNameServer(name IString, hostName IString, server IPortNameServer) IDistantObject {
	rv := objc.Send[DistantObject](objc.ID(cc.class), objc.Sel("rootProxyForConnectionWithRegisteredName:host:usingNameServer:"), name, hostName, server)
	return rv
}


// Creates and returns a new connection object representing a vended service on the default system port name server.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSConnection/serviceConnectionWithName:rootObject:
func (cc _ConnectionClass) ServiceConnectionWithNameRootObject(name IString, root objc.IObject) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(cc.class), objc.Sel("serviceConnectionWithName:rootObject:"), name, root)
	return rv
}


// Creates and returns a new connection object representing a vended service on the specified port name server.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSConnection/serviceConnectionWithName:rootObject:usingNameServer:
func (cc _ConnectionClass) ServiceConnectionWithNameRootObjectUsingNameServer(name IString, root objc.IObject, server IPortNameServer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(cc.class), objc.Sel("serviceConnectionWithName:rootObject:usingNameServer:"), name, root, server)
	return rv
}


// The receiver’s delegate.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSConnection/delegate-c.property
func (c_ Connection) Delegate() objc.ID {
	rv := objc.Send[objc.ID](c_.ID, objc.Sel("delegate"))
	return rv
}


// The receiver’s delegate.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSConnection/delegate-c.property
func (c_ Connection) SetDelegate(value objc.ID) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setDelegate:"), value)
}


// A Boolean value that indicates whether the receiver handles remote messages atomically.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSConnection/independentConversationQueueing
func (c_ Connection) IndependentConversationQueueing() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("independentConversationQueueing"))
	return rv
}


// A Boolean value that indicates whether the receiver handles remote messages atomically.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSConnection/independentConversationQueueing
func (c_ Connection) SetIndependentConversationQueueing(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIndependentConversationQueueing:"), value)
}


// The local objects that have been sent over the connection and still have proxies at the other end.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSConnection/localObjects
func (c_ Connection) LocalObjects() IArray {
	rv := objc.Send[Array](c_.ID, objc.Sel("localObjects"))
	return rv
}


// A Boolean value that indicates whether the receiver supports requests from multiple threads.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSConnection/multipleThreadsEnabled
func (c_ Connection) MultipleThreadsEnabled() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("multipleThreadsEnabled"))
	return rv
}


// The port on which the receiver receives incoming network messages.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSConnection/receivePort-c.property
func (c_ Connection) ReceivePort() IPort {
	rv := objc.Send[Port](c_.ID, objc.Sel("receivePort"))
	return rv
}


// The local proxies for remote objects that have been received over the connection but not deallocated yet.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSConnection/remoteObjects
func (c_ Connection) RemoteObjects() IArray {
	rv := objc.Send[Array](c_.ID, objc.Sel("remoteObjects"))
	return rv
}


// The timeout interval for replies to outgoing remote messages.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSConnection/replyTimeout
func (c_ Connection) ReplyTimeout() float64 {
	rv := objc.Send[objc.ID](c_.ID, objc.Sel("replyTimeout"))
	return rv
}


// The timeout interval for replies to outgoing remote messages.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSConnection/replyTimeout
func (c_ Connection) SetReplyTimeout(value float64) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setReplyTimeout:"), value)
}


// The set of request modes the receiver’s receive port is registered for with its object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSConnection/requestModes-c.property
func (c_ Connection) RequestModes() []string {
	rv := objc.Send[[]string](c_.ID, objc.Sel("requestModes"))
	return rv
}


// The timeout interval for outgoing remote messages.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSConnection/requestTimeout
func (c_ Connection) RequestTimeout() float64 {
	rv := objc.Send[objc.ID](c_.ID, objc.Sel("requestTimeout"))
	return rv
}


// The timeout interval for outgoing remote messages.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSConnection/requestTimeout
func (c_ Connection) SetRequestTimeout(value float64) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setRequestTimeout:"), value)
}


// The object that the receiver (or its parent) makes available to other applications or threads.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSConnection/rootObject-c.property
func (c_ Connection) RootObject() objc.ID {
	rv := objc.Send[objc.ID](c_.ID, objc.Sel("rootObject"))
	return rv
}


// The object that the receiver (or its parent) makes available to other applications or threads.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSConnection/rootObject-c.property
func (c_ Connection) SetRootObject(value objc.ID) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setRootObject:"), value)
}


// The proxy for the root object of the receiver’s peer in another application or thread.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSConnection/rootProxy
func (c_ Connection) RootProxy() IDistantObject {
	rv := objc.Send[DistantObject](c_.ID, objc.Sel("rootProxy"))
	return rv
}


// The port that the connection sends outgoing network messages through.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSConnection/sendPort-c.property
func (c_ Connection) SendPort() IPort {
	rv := objc.Send[Port](c_.ID, objc.Sel("sendPort"))
	return rv
}


// A dictionary containing various statistics for the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSConnection/statistics-c.property
func (c_ Connection) Statistics() IDictionary {
	rv := objc.Send[Dictionary](c_.ID, objc.Sel("statistics"))
	return rv
}


// A Boolean value that indicates whether the receiver is known to be valid.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSConnection/valid
func (c_ Connection) Valid() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("valid"))
	return rv
}


