// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [VZVirtioSocketConnection] class.
var (
	VZVirtioSocketConnectionClass     _VZVirtioSocketConnectionClass
	VZVirtioSocketConnectionClassOnce sync.Once
)

func getVZVirtioSocketConnectionClass() _VZVirtioSocketConnectionClass {
	VZVirtioSocketConnectionClassOnce.Do(func() {
		VZVirtioSocketConnectionClass = _VZVirtioSocketConnectionClass{objc.GetClass("VZVirtioSocketConnection")}
	})
	return VZVirtioSocketConnectionClass
}

type _VZVirtioSocketConnectionClass struct {
	class objc.Class
}

// An interface definition for the [VZVirtioSocketConnection] class.
type IVZVirtioSocketConnection interface {
	objectivec.IObject
	// properties:
	DestinationPort() uint32 /* not a class type */
	FileDescriptor() int
	SourcePort() uint32 /* not a class type */
	// methods:
	Close()
}

// A port-based connection between the guest operating system and the host computer.
//
// A object contains the port information for the guest operating system and host computer. You don’t create connection objects directly. When the guest operating system initiates a connection, the virtual machine creates the connection object and passes it to the appropriate object, which forwards the object to its delegate. When the virtual machine opens a connection to a guest port, the method (Objective-C) or method (Swift) pass the connection object to your completion handler.


// A port-based connection between the guest operating system and the host computer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZVirtioSocketConnection
type VZVirtioSocketConnection struct {
	objectivec.Object
}

// VZVirtioSocketConnectionFrom constructs a [VZVirtioSocketConnection] from an unsafe.Pointer.
//
// A port-based connection between the guest operating system and the host computer.
func VZVirtioSocketConnectionFrom(ptr unsafe.Pointer) VZVirtioSocketConnection {
	return VZVirtioSocketConnection{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (vc _VZVirtioSocketConnectionClass) Alloc() VZVirtioSocketConnection {
	rv := objc.Send[VZVirtioSocketConnection](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (vc _VZVirtioSocketConnectionClass) New() VZVirtioSocketConnection {
	rv := objc.Send[VZVirtioSocketConnection](objc.ID(vc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (v_ VZVirtioSocketConnection) Init() VZVirtioSocketConnection {
	rv := objc.Send[VZVirtioSocketConnection](v_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v_ VZVirtioSocketConnection) Autorelease() VZVirtioSocketConnection {
	rv := objc.Send[VZVirtioSocketConnection](v_.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZVirtioSocketConnection creates a new VZVirtioSocketConnection instance.
func NewVZVirtioSocketConnection() VZVirtioSocketConnection {
	return getVZVirtioSocketConnectionClass().New()
}



// Close the file descriptor associated with the socket.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZVirtioSocketConnection/close()
func (v_ VZVirtioSocketConnection) Close() {
	objc.Send[objc.ID](v_.ID, objc.Sel("close"))
}


// The destination port number of the connection.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZVirtioSocketConnection/destinationPort
func (v_ VZVirtioSocketConnection) DestinationPort() uint32 /* not a class type */ {
	rv := objc.Send[uint32](v_.ID, objc.Sel("destinationPort"))
	return rv
}


// The file descriptor to use when sending data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZVirtioSocketConnection/fileDescriptor
func (v_ VZVirtioSocketConnection) FileDescriptor() int {
	rv := objc.Send[int](v_.ID, objc.Sel("fileDescriptor"))
	return rv
}


// The port number of the system that opened the connection.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZVirtioSocketConnection/sourcePort
func (v_ VZVirtioSocketConnection) SourcePort() uint32 /* not a class type */ {
	rv := objc.Send[uint32](v_.ID, objc.Sel("sourcePort"))
	return rv
}



