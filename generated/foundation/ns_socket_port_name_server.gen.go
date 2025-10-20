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
}

// A port name server that takes and returns socket ports.
//
// Port removal functionality is supported by the method and should be used to remove invalid socket ports. Unlike the other port name servers, can operate over a network. By registering your socket ports, you make them available to other computers on the local network without hard-coding the TCP port numbers. Clients just need to know the name of the port. is implemented using and registers ports in the local network domain. The registered name of a port must be unique within the local domain, not just the local host. The name server only supports TCP/IP (either IPv4 or IPv6) sockets.
//
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




