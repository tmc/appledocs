// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [SocketPort] class.
var (
	SocketPortClass     _SocketPortClass
	SocketPortClassOnce sync.Once
)

func getSocketPortClass() _SocketPortClass {
	SocketPortClassOnce.Do(func() {
		SocketPortClass = _SocketPortClass{objc.GetClass("NSSocketPort")}
	})
	return SocketPortClass
}

type _SocketPortClass struct {
	class objc.Class
}

// An interface definition for the [SocketPort] class.
type ISocketPort interface {
	IPort
	Address() IData
	Protocol() int
	ProtocolFamily() int
	Socket() SocketNativeHandle
	SocketType() int
}

// A port that represents a BSD socket.
//
// A object can be used as an endpoint for distributed object connections. Companion classes, and , allow for local (on the same machine) communication only. The class allows for both local and remote communication, but may be more expensive than the others for the local case.


// A port that represents a BSD socket.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/SocketPort
type SocketPort struct {
	Port
}

// SocketPortFrom constructs a [SocketPort] from an unsafe.Pointer.
//
// A port that represents a BSD socket.
func SocketPortFrom(ptr unsafe.Pointer) SocketPort {
	return SocketPort{
		Port: PortFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (sc _SocketPortClass) Alloc() SocketPort {
	rv := objc.Send[SocketPort](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (sc _SocketPortClass) New() SocketPort {
	rv := objc.Send[SocketPort](objc.ID(sc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (s_ SocketPort) Init() SocketPort {
	rv := objc.Send[SocketPort](s_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s_ SocketPort) Autorelease() SocketPort {
	rv := objc.Send[SocketPort](s_.ID, objc.Sel("autorelease"))
	return rv
}

// NewSocketPort creates a new SocketPort instance.
func NewSocketPort() SocketPort {
	return getSocketPortClass().New()
}



// Initializes the receiver as a remote socket with the provided arguments.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/SocketPort/init(remoteWithProtocolFamily:socketType:protocol:address:)
func NewSocketPortRemoteWithProtocolFamilySocketTypeProtocolAddress(family int, type_ int, protocol_ int, address IData) SocketPort {
	instance := getSocketPortClass().Alloc()
	rv := objc.Send[SocketPort](instance.ID, objc.Sel("initRemoteWithProtocolFamily:socketType:protocol:address:"), family, type_, protocol_, address)
	rv.Autorelease()
	return rv
}


// Initializes the receiver as a TCP/IP socket of type that can connect to a remote host on a specified port.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/SocketPort/init(remoteWithTCPPort:host:)
func NewSocketPortRemoteWithTCPPortHost(port unsafe.Pointer, hostName string) SocketPort {
	instance := getSocketPortClass().Alloc()
	rv := objc.Send[SocketPort](instance.ID, objc.Sel("initRemoteWithTCPPort:host:"), port, objc.String(hostName))
	rv.Autorelease()
	return rv
}


// Initializes the receiver as a local socket with the provided arguments.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/SocketPort/init(protocolFamily:socketType:protocol:address:)
func NewSocketPortWithProtocolFamilySocketTypeProtocolAddress(family int, type_ int, protocol_ int, address IData) SocketPort {
	instance := getSocketPortClass().Alloc()
	rv := objc.Send[SocketPort](instance.ID, objc.Sel("initWithProtocolFamily:socketType:protocol:address:"), family, type_, protocol_, address)
	rv.Autorelease()
	return rv
}


// Initializes the receiver with a previously created local socket.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/SocketPort/init(protocolFamily:socketType:protocol:socket:)
func NewSocketPortWithProtocolFamilySocketTypeProtocolSocket(family int, type_ int, protocol_ int, sock SocketNativeHandle) SocketPort {
	instance := getSocketPortClass().Alloc()
	rv := objc.Send[SocketPort](instance.ID, objc.Sel("initWithProtocolFamily:socketType:protocol:socket:"), family, type_, protocol_, sock)
	rv.Autorelease()
	return rv
}


// Initializes the receiver as a local TCP/IP socket of type , listening on a specified port number.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/SocketPort/init(tcpPort:)
func NewSocketPortWithTCPPort(port unsafe.Pointer) SocketPort {
	instance := getSocketPortClass().Alloc()
	rv := objc.Send[SocketPort](instance.ID, objc.Sel("initWithTCPPort:"), port)
	rv.Autorelease()
	return rv
}



// The receiver’s socket address structure stored inside an object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/SocketPort/address
func (s_ SocketPort) Address() IData {
	rv := objc.Send[Data](s_.ID, objc.Sel("address"))
	return rv
}


// The protocol that the receiver uses for communication.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/SocketPort/protocol
func (s_ SocketPort) Protocol() int {
	rv := objc.Send[int](s_.ID, objc.Sel("protocol"))
	return rv
}


// The protocol family that the receiver uses for communication.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/SocketPort/protocolFamily
func (s_ SocketPort) ProtocolFamily() int {
	rv := objc.Send[int](s_.ID, objc.Sel("protocolFamily"))
	return rv
}


// The receiver’s native socket identifier on the platform.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/SocketPort/socket
func (s_ SocketPort) Socket() SocketNativeHandle {
	rv := objc.Send[SocketNativeHandle](s_.ID, objc.Sel("socket"))
	return rv
}


// The receiver’s socket type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/SocketPort/socketType
func (s_ SocketPort) SocketType() int {
	rv := objc.Send[int](s_.ID, objc.Sel("socketType"))
	return rv
}


