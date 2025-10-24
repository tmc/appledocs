// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class VZVirtioSocketConnection */

/* debug [class_header]: Header for VZVirtioSocketConnection */
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

/* debug [class_header]: End header */

/* debug [class_interface]: Interface for VZVirtioSocketConnection */
// An interface definition for the [VZVirtioSocketConnection] class.
type IVZVirtioSocketConnection interface {
	objectivec.IObject

	/* debug [class_interface_properties]: Properties for VZVirtioSocketConnection */
	// properties:
	DestinationPort() uint32 /* not a class type */
	FileDescriptor() int
	SourcePort() uint32 /* not a class type */
	/* debug [class_interface_properties]: End properties */

	/* debug [class_interface_methods]: Methods for VZVirtioSocketConnection */
	// methods:
	Close()
	/* debug [class_interface_methods]: End methods */

}

/* debug [class_interface]: End interface */

/* debug [class_constructors]: Constructors for VZVirtioSocketConnection */
// Alloc allocates a new instance without initialization.
func (vc _VZVirtioSocketConnectionClass) Alloc() VZVirtioSocketConnection {
	rv := objc.Send[VZVirtioSocketConnection](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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

/* debug [class_constructors]: End constructors */

/* debug [class_struct]: Struct for VZVirtioSocketConnection */
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

/* debug [class_struct]: End struct */

/* debug [class_init_methods]: Init methods for VZVirtioSocketConnection */ /* debug [class_init_methods]: End init methods */

/* debug [class_methods]: Class methods for VZVirtioSocketConnection */
/* debug [class_methods]: End class methods */

/* debug [class_properties_class]: Class properties for VZVirtioSocketConnection */
/* debug [class_properties_class]: End class properties */

/* debug [instance_methods]: Instance methods for VZVirtioSocketConnection */

// Close the file descriptor associated with the socket.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZVirtioSocketConnection/close()
func (v_ VZVirtioSocketConnection) Close() {
	objc.Send[objc.ID](v_.ID, objc.Sel("close"))
} /* debug [instance_methods/method]: Close */

/* debug [instance_methods]: End instance methods */

/* debug [instance_properties]: Instance properties for VZVirtioSocketConnection */

// The destination port number of the connection.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZVirtioSocketConnection/destinationPort
func (v_ VZVirtioSocketConnection) DestinationPort() uint32 /* not a class type */ {
	rv := objc.Send[uint32](v_.ID, objc.Sel("destinationPort"))
	return rv
} /* debug [instance_properties/getter]: destinationPort */

// The file descriptor to use when sending data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZVirtioSocketConnection/fileDescriptor
func (v_ VZVirtioSocketConnection) FileDescriptor() int {
	rv := objc.Send[int](v_.ID, objc.Sel("fileDescriptor"))
	return rv
} /* debug [instance_properties/getter]: fileDescriptor */

// The port number of the system that opened the connection.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZVirtioSocketConnection/sourcePort
func (v_ VZVirtioSocketConnection) SourcePort() uint32 /* not a class type */ {
	rv := objc.Send[uint32](v_.ID, objc.Sel("sourcePort"))
	return rv
} /* debug [instance_properties/getter]: sourcePort */

/* debug [instance_properties]: End instance properties */

/* debug [class.gen.go]: End class VZVirtioSocketConnection */
