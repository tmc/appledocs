// Code generated from Apple documentation for NetworkExtension. DO NOT EDIT.

package networkextension

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class NEFilterSocketFlow */


/* debug [class_header]: Header for NEFilterSocketFlow */
// The class instance for the [NEFilterSocketFlow] class.
var (
	NEFilterSocketFlowClass     _NEFilterSocketFlowClass
	NEFilterSocketFlowClassOnce sync.Once
)

func getNEFilterSocketFlowClass() _NEFilterSocketFlowClass {
	NEFilterSocketFlowClassOnce.Do(func() {
		NEFilterSocketFlowClass = _NEFilterSocketFlowClass{objc.GetClass("NEFilterSocketFlow")}
	})
	return NEFilterSocketFlowClass
}

type _NEFilterSocketFlowClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for NEFilterSocketFlow */
// An interface definition for the [NEFilterSocketFlow] class.
type INEFilterSocketFlow interface {
	INEFilterFlow
	
/* debug [class_interface_properties]: Properties for NEFilterSocketFlow */
	// properties:
	LocalEndpoint() INWEndpoint
	LocalFlowEndpoint() objectivec.IObject
	RemoteEndpoint() INWEndpoint
	RemoteFlowEndpoint() objectivec.IObject
	RemoteHostname() objc.IObject /* cross-framework: NSString */
	SocketFamily() int
	SocketProtocol() int
	SocketType() int
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for NEFilterSocketFlow */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for NEFilterSocketFlow */
// Alloc allocates a new instance without initialization.
func (nc _NEFilterSocketFlowClass) Alloc() NEFilterSocketFlow {
	rv := objc.Send[NEFilterSocketFlow](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (nc _NEFilterSocketFlowClass) New() NEFilterSocketFlow {
	rv := objc.Send[NEFilterSocketFlow](objc.ID(nc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (n_ NEFilterSocketFlow) Init() NEFilterSocketFlow {
	rv := objc.Send[NEFilterSocketFlow](n_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (n_ NEFilterSocketFlow) Autorelease() NEFilterSocketFlow {
	rv := objc.Send[NEFilterSocketFlow](n_.ID, objc.Sel("autorelease"))
	return rv
}

// NewNEFilterSocketFlow creates a new NEFilterSocketFlow instance.
func NewNEFilterSocketFlow() NEFilterSocketFlow {
	return getNEFilterSocketFlowClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for NEFilterSocketFlow */
// A flow of network data that the filter examines.


// A flow of network data that the filter examines.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEFilterSocketFlow
type NEFilterSocketFlow struct {
	NEFilterFlow
}

// NEFilterSocketFlowFrom constructs a [NEFilterSocketFlow] from an unsafe.Pointer.
//
// A flow of network data that the filter examines.
func NEFilterSocketFlowFrom(ptr unsafe.Pointer) NEFilterSocketFlow {
	return NEFilterSocketFlow{
		NEFilterFlow: NEFilterFlowFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for NEFilterSocketFlow *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for NEFilterSocketFlow */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for NEFilterSocketFlow */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for NEFilterSocketFlow */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for NEFilterSocketFlow */

// An object containing details about the socket’s local endpoint.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEFilterSocketFlow/localEndpoint
func (n_ NEFilterSocketFlow) LocalEndpoint() INWEndpoint {
	rv := objc.Send[NWEndpoint](n_.ID, objc.Sel("localEndpoint"))
	return rv
}/* debug [instance_properties/getter]: localEndpoint */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEFilterSocketFlow/localFlowEndpoint-4nt54
func (n_ NEFilterSocketFlow) LocalFlowEndpoint() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](n_.ID, objc.Sel("localFlowEndpoint"))
	return rv
}/* debug [instance_properties/getter]: localFlowEndpoint */


// An object containing details about the socket’s remote endpoint.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEFilterSocketFlow/remoteEndpoint
func (n_ NEFilterSocketFlow) RemoteEndpoint() INWEndpoint {
	rv := objc.Send[NWEndpoint](n_.ID, objc.Sel("remoteEndpoint"))
	return rv
}/* debug [instance_properties/getter]: remoteEndpoint */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEFilterSocketFlow/remoteFlowEndpoint-52dxr
func (n_ NEFilterSocketFlow) RemoteFlowEndpoint() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](n_.ID, objc.Sel("remoteFlowEndpoint"))
	return rv
}/* debug [instance_properties/getter]: remoteFlowEndpoint */


// The flow’s remote hostname, if applicable.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEFilterSocketFlow/remoteHostname
func (n_ NEFilterSocketFlow) RemoteHostname() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](n_.ID, objc.Sel("remoteHostname"))
	return rv
}/* debug [instance_properties/getter]: remoteHostname */


// The protocol family of the socket.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEFilterSocketFlow/socketFamily
func (n_ NEFilterSocketFlow) SocketFamily() int {
	rv := objc.Send[int](n_.ID, objc.Sel("socketFamily"))
	return rv
}/* debug [instance_properties/getter]: socketFamily */


// The protocol of the socket.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEFilterSocketFlow/socketProtocol
func (n_ NEFilterSocketFlow) SocketProtocol() int {
	rv := objc.Send[int](n_.ID, objc.Sel("socketProtocol"))
	return rv
}/* debug [instance_properties/getter]: socketProtocol */


// The type of the socket.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEFilterSocketFlow/socketType
func (n_ NEFilterSocketFlow) SocketType() int {
	rv := objc.Send[int](n_.ID, objc.Sel("socketType"))
	return rv
}/* debug [instance_properties/getter]: socketType */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class NEFilterSocketFlow */



