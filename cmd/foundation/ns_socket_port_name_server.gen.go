// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [SocketPortNameServer] class.
var (
	SocketPortNameServerClass     _SocketPortNameServerClass
	SocketPortNameServerClassOnce sync.Once
)

func getSocketPortNameServerClass() _SocketPortNameServerClass {
	SocketPortNameServerClassOnce.Do(func() {
		SocketPortNameServerClass = _SocketPortNameServerClass{objc.GetClass("NSSocketPortNameServer")}
	})
	return SocketPortNameServerClass
}

type _SocketPortNameServerClass struct {
	class objc.Class
}

// An interface definition for the [SocketPortNameServer] class.
type ISocketPortNameServer interface {
	IPortNameServer
	PortForName(name string) Port
	PortForNameHost(name string, host string) Port
	PortForNameHostNameServerPortNumber(name string, host string, portNumber unsafe.Pointer) Port
	RegisterPortName(port IPort, name string) bool
	RegisterPortNameNameServerPortNumber(port IPort, name string, portNumber unsafe.Pointer) bool
	RemovePortForName(name string) bool
	DefaultNameServerPortNumber() unsafe.Pointer
	SetDefaultNameServerPortNumber(value unsafe.Pointer)
}

// A port name server that takes and returns socket ports.
//
// Port removal functionality is supported by the method and should be used to remove invalid socket ports. Unlike the other port name servers, can operate over a network. By registering your socket ports, you make them available to other computers on the local network without hard-coding the TCP port numbers. Clients just need to know the name of the port. is implemented using and registers ports in the local network domain. The registered name of a port must be unique within the local domain, not just the local host. The name server only supports TCP/IP (either IPv4 or IPv6) sockets.


// A port name server that takes and returns socket ports.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSSocketPortNameServer
type SocketPortNameServer struct {
	PortNameServer
}

// SocketPortNameServerFrom constructs a [SocketPortNameServer] from an unsafe.Pointer.
//
// A port name server that takes and returns socket ports.
func SocketPortNameServerFrom(ptr unsafe.Pointer) SocketPortNameServer {
	return SocketPortNameServer{
		PortNameServer: PortNameServerFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (sc _SocketPortNameServerClass) Alloc() SocketPortNameServer {
	rv := objc.Send[SocketPortNameServer](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (sc _SocketPortNameServerClass) New() SocketPortNameServer {
	rv := objc.Send[SocketPortNameServer](objc.ID(sc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (s_ SocketPortNameServer) Init() SocketPortNameServer {
	rv := objc.Send[SocketPortNameServer](s_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s_ SocketPortNameServer) Autorelease() SocketPortNameServer {
	rv := objc.Send[SocketPortNameServer](s_.ID, objc.Sel("autorelease"))
	return rv
}

// NewSocketPortNameServer creates a new SocketPortNameServer instance.
func NewSocketPortNameServer() SocketPortNameServer {
	return getSocketPortNameServerClass().New()
}



// Returns the shared socket port name server.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSSocketPortNameServer/sharedInstance
func (sc _SocketPortNameServerClass) SharedInstance() objc.ID {
	rv := objc.Send[objc.ID](objc.ID(sc.class), objc.Sel("sharedInstance"))
	return rv
}


// Looks up and returns the port registered under the specified name on the local host.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSSocketPortNameServer/portForName:
func (s_ SocketPortNameServer) PortForName(name string) Port {
	rv := objc.Send[Port](s_.ID, objc.Sel("portForName:"), objc.String(name))
	return rv
}


// Looks up and returns the port registered under the specified name on a specified host.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSSocketPortNameServer/portForName:host:
func (s_ SocketPortNameServer) PortForNameHost(name string, host string) Port {
	rv := objc.Send[Port](s_.ID, objc.Sel("portForName:host:"), objc.String(name), objc.String(host))
	return rv
}


// Looks up and returns the port registered under the specified name on a specified host.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSSocketPortNameServer/portForName:host:nameServerPortNumber:
func (s_ SocketPortNameServer) PortForNameHostNameServerPortNumber(name string, host string, portNumber unsafe.Pointer) Port {
	rv := objc.Send[Port](s_.ID, objc.Sel("portForName:host:nameServerPortNumber:"), objc.String(name), objc.String(host), portNumber)
	return rv
}


// Registers a given port as a network service with the specified name in the local domain.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSSocketPortNameServer/registerPort:name:
func (s_ SocketPortNameServer) RegisterPortName(port IPort, name string) bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("registerPort:name:"), port, objc.String(name))
	return rv
}


// Registers a given port as a network service with the specified name in the local domain.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSSocketPortNameServer/registerPort:name:nameServerPortNumber:
func (s_ SocketPortNameServer) RegisterPortNameNameServerPortNumber(port IPort, name string, portNumber unsafe.Pointer) bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("registerPort:name:nameServerPortNumber:"), port, objc.String(name), portNumber)
	return rv
}


// Unregisters the port for a given name on the local host.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSSocketPortNameServer/removePortForName:
func (s_ SocketPortNameServer) RemovePortForName(name string) bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("removePortForName:"), objc.String(name))
	return rv
}


// Returns the port number used to contact the name server.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSSocketPortNameServer/defaultNameServerPortNumber
func (s_ SocketPortNameServer) DefaultNameServerPortNumber() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("defaultNameServerPortNumber"))
	return rv
}


// Returns the port number used to contact the name server.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSSocketPortNameServer/defaultNameServerPortNumber
func (s_ SocketPortNameServer) SetDefaultNameServerPortNumber(value unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setDefaultNameServerPortNumber:"), value)
}



