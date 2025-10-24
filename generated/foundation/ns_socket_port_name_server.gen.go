// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class NSSocketPortNameServer */


/* debug [class_header]: Header for NSSocketPortNameServer */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for SocketPortNameServer */
// An interface definition for the [SocketPortNameServer] class.
type ISocketPortNameServer interface {
	IPortNameServer
	
/* debug [class_interface_properties]: Properties for SocketPortNameServer */
	// properties:
	DefaultNameServerPortNumber() uint16 /* not a class type */
	SetDefaultNameServerPortNumber(value uint16 /* not a class type */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for SocketPortNameServer */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for SocketPortNameServer */
// Alloc allocates a new instance without initialization.
func (sc _SocketPortNameServerClass) Alloc() SocketPortNameServer {
	rv := objc.Send[SocketPortNameServer](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for SocketPortNameServer */
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for SocketPortNameServer *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for SocketPortNameServer */

// Returns the shared socket port name server.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSSocketPortNameServer/sharedInstance
func (sc _SocketPortNameServerClass) SharedInstance() objc.ID {
	rv := objc.Send[objc.ID](objc.ID(sc.class), objc.Sel("sharedInstance"))
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=SharedInstance) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for SocketPortNameServer */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for SocketPortNameServer */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for SocketPortNameServer */

// Returns the port number used to contact the name server.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSSocketPortNameServer/defaultNameServerPortNumber
func (s_ SocketPortNameServer) DefaultNameServerPortNumber() uint16 /* not a class type */ {
	rv := objc.Send[uint16](s_.ID, objc.Sel("defaultNameServerPortNumber"))
	return rv
}/* debug [instance_properties/getter]: defaultNameServerPortNumber */


// Returns the port number used to contact the name server.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSSocketPortNameServer/defaultNameServerPortNumber
func (s_ SocketPortNameServer) SetDefaultNameServerPortNumber(value uint16 /* not a class type */) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setDefaultNameServerPortNumber:"), value)
}/* debug [instance_properties/setter]: defaultNameServerPortNumber */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class NSSocketPortNameServer */



