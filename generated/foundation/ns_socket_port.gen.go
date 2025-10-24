// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class NSSocketPort */


/* debug [class_header]: Header for NSSocketPort */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for SocketPort */
// An interface definition for the [SocketPort] class.
type ISocketPort interface {
	IPort
	
/* debug [class_interface_properties]: Properties for SocketPort */
	// properties:
	Address() IData
	SetAddress(value IData)
	Protocol() objectivec.IObject
	SetProtocol(value objectivec.IObject)
	ProtocolFamily() objectivec.IObject
	SetProtocolFamily(value objectivec.IObject)
	Socket() objectivec.IObject
	SetSocket(value objectivec.IObject)
	SocketType() objectivec.IObject
	SetSocketType(value objectivec.IObject)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for SocketPort */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for SocketPort */
// Alloc allocates a new instance without initialization.
func (sc _SocketPortClass) Alloc() SocketPort {
	rv := objc.Send[SocketPort](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for SocketPort */
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for SocketPort */
/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for SocketPort */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for SocketPort */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for SocketPort */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for SocketPort */

// The receiver’s socket address structure stored inside an
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/socketport/address
func (s_ SocketPort) Address() IData {
	rv := objc.Send[Data](s_.ID, objc.Sel("address"))
	return rv
}/* debug [instance_properties/getter]: address */


// The receiver’s socket address structure stored inside an
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/socketport/address
func (s_ SocketPort) SetAddress(value IData) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setAddress:"), value)
}/* debug [instance_properties/setter]: address */


// The protocol that the receiver uses for communication.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/socketport/protocol
func (s_ SocketPort) Protocol() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](s_.ID, objc.Sel("protocol"))
	return rv
}/* debug [instance_properties/getter]: protocol */


// The protocol that the receiver uses for communication.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/socketport/protocol
func (s_ SocketPort) SetProtocol(value objectivec.IObject) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setProtocol:"), value)
}/* debug [instance_properties/setter]: protocol */


// The protocol family that the receiver uses for communication.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/socketport/protocolfamily
func (s_ SocketPort) ProtocolFamily() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](s_.ID, objc.Sel("protocolFamily"))
	return rv
}/* debug [instance_properties/getter]: protocolFamily */


// The protocol family that the receiver uses for communication.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/socketport/protocolfamily
func (s_ SocketPort) SetProtocolFamily(value objectivec.IObject) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setProtocolFamily:"), value)
}/* debug [instance_properties/setter]: protocolFamily */


// The receiver’s native socket identifier on the platform.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/socketport/socket
func (s_ SocketPort) Socket() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](s_.ID, objc.Sel("socket"))
	return rv
}/* debug [instance_properties/getter]: socket */


// The receiver’s native socket identifier on the platform.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/socketport/socket
func (s_ SocketPort) SetSocket(value objectivec.IObject) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setSocket:"), value)
}/* debug [instance_properties/setter]: socket */


// The receiver’s socket type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/socketport/sockettype
func (s_ SocketPort) SocketType() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](s_.ID, objc.Sel("socketType"))
	return rv
}/* debug [instance_properties/getter]: socketType */


// The receiver’s socket type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/socketport/sockettype
func (s_ SocketPort) SetSocketType(value objectivec.IObject) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setSocketType:"), value)
}/* debug [instance_properties/setter]: socketType */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class NSSocketPort */


