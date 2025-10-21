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
}

// A port that represents a BSD socket.
//
// A object can be used as an endpoint for distributed object connections. Companion classes, and , allow for local (on the same machine) communication only. The class allows for both local and remote communication, but may be more expensive than the others for the local case.
//
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




// Initializes the receiver as a local socket with the provided arguments.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/SocketPort/init(protocolFamily:socketType:protocol:address:)
func NewSocketPortWithProtocolFamilySocketTypeProtocolAddress(family unsafe.Pointer, type_ unsafe.Pointer, protocol_ unsafe.Pointer, address unsafe.Pointer) SocketPort {
	instance := getSocketPortClass().Alloc()
	rv := objc.Send[SocketPort](instance.ID, objc.Sel("initWithProtocolFamily:socketType:protocol:address:"), family, type_, protocol_, address)
	rv.Autorelease()
	return rv
}


// The protocol family that the receiver uses for communication.
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/socketport/protocolfamily
func (s_ SocketPort) ProtocolFamily() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("protocolFamily"))
	return rv
}


// SetProtocolFamily sets the value of the protocolFamily property.
// The protocol family that the receiver uses for communication.

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/socketport/protocolfamily
func (s_ SocketPort) SetProtocolFamily(value unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setProtocolFamily:"), value)
}

// The receiver’s socket type.
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/socketport/sockettype
func (s_ SocketPort) SocketType() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("socketType"))
	return rv
}


// SetSocketType sets the value of the socketType property.
// The receiver’s socket type.

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/socketport/sockettype
func (s_ SocketPort) SetSocketType(value unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setSocketType:"), value)
}

// The protocol that the receiver uses for communication.
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/socketport/protocol
func (s_ SocketPort) `protocol`() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("`protocol`"))
	return rv
}


// Set`protocol` sets the value of the `protocol` property.
// The protocol that the receiver uses for communication.

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/socketport/protocol
func (s_ SocketPort) Set`protocol`(value unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID, objc.Sel("set`protocol`:"), value)
}

// The receiver’s socket address structure stored inside an object.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/SocketPort/address
func (s_ SocketPort) Address() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("address"))
	return rv
}

// The protocol that the receiver uses for communication.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/SocketPort/protocol
func (s_ SocketPort) Protocol_() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("protocol"))
	return rv
}

// The receiver’s native socket identifier on the platform.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/SocketPort/socket
func (s_ SocketPort) Socket() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("socket"))
	return rv
}


