// Code generated from Apple documentation for NetworkExtension. DO NOT EDIT.

package networkextension

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class NEAppProxyUDPFlow */


/* debug [class_header]: Header for NEAppProxyUDPFlow */
// The class instance for the [NEAppProxyUDPFlow] class.
var (
	NEAppProxyUDPFlowClass     _NEAppProxyUDPFlowClass
	NEAppProxyUDPFlowClassOnce sync.Once
)

func getNEAppProxyUDPFlowClass() _NEAppProxyUDPFlowClass {
	NEAppProxyUDPFlowClassOnce.Do(func() {
		NEAppProxyUDPFlowClass = _NEAppProxyUDPFlowClass{objc.GetClass("NEAppProxyUDPFlow")}
	})
	return NEAppProxyUDPFlowClass
}

type _NEAppProxyUDPFlowClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for NEAppProxyUDPFlow */
// An interface definition for the [NEAppProxyUDPFlow] class.
type INEAppProxyUDPFlow interface {
	INEAppProxyFlow
	
/* debug [class_interface_properties]: Properties for NEAppProxyUDPFlow */
	// properties:
	LocalEndpoint() INWEndpoint
	LocalFlowEndpoint() objectivec.IObject
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for NEAppProxyUDPFlow */
	// methods:
	ReadDatagramsAndFlowEndpointsWithCompletionHandler(completionHandler unsafe.Pointer)
	WriteDatagramsSentByFlowEndpointsCompletionHandler(datagrams []foundation.Data, remoteEndpoints objectivec.IObject, completionHandler unsafe.Pointer)
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for NEAppProxyUDPFlow */
// Alloc allocates a new instance without initialization.
func (nc _NEAppProxyUDPFlowClass) Alloc() NEAppProxyUDPFlow {
	rv := objc.Send[NEAppProxyUDPFlow](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (nc _NEAppProxyUDPFlowClass) New() NEAppProxyUDPFlow {
	rv := objc.Send[NEAppProxyUDPFlow](objc.ID(nc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (n_ NEAppProxyUDPFlow) Init() NEAppProxyUDPFlow {
	rv := objc.Send[NEAppProxyUDPFlow](n_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (n_ NEAppProxyUDPFlow) Autorelease() NEAppProxyUDPFlow {
	rv := objc.Send[NEAppProxyUDPFlow](n_.ID, objc.Sel("autorelease"))
	return rv
}

// NewNEAppProxyUDPFlow creates a new NEAppProxyUDPFlow instance.
func NewNEAppProxyUDPFlow() NEAppProxyUDPFlow {
	return getNEAppProxyUDPFlowClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for NEAppProxyUDPFlow */
// An object for reading and writing data to and from a UDP conversation being proxied by the provider.
//
// App Proxy Providers receive UDP connections to be proxied in the form of objects.


// An object for reading and writing data to and from a UDP conversation being proxied by the provider.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEAppProxyUDPFlow
type NEAppProxyUDPFlow struct {
	NEAppProxyFlow
}

// NEAppProxyUDPFlowFrom constructs a [NEAppProxyUDPFlow] from an unsafe.Pointer.
//
// An object for reading and writing data to and from a UDP conversation being proxied by the provider.
func NEAppProxyUDPFlowFrom(ptr unsafe.Pointer) NEAppProxyUDPFlow {
	return NEAppProxyUDPFlow{
		NEAppProxyFlow: NEAppProxyFlowFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for NEAppProxyUDPFlow *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for NEAppProxyUDPFlow */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for NEAppProxyUDPFlow */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for NEAppProxyUDPFlow */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEAppProxyUDPFlow/readDatagramsAndFlowEndpointsWithCompletionHandler:
func (n_ NEAppProxyUDPFlow) ReadDatagramsAndFlowEndpointsWithCompletionHandler(completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](n_.ID, objc.Sel("readDatagramsAndFlowEndpointsWithCompletionHandler:"), completionHandler)
}/* debug [instance_methods/method]: ReadDatagramsAndFlowEndpointsWithCompletionHandler */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEAppProxyUDPFlow/writeDatagrams:sentByFlowEndpoints:completionHandler:
func (n_ NEAppProxyUDPFlow) WriteDatagramsSentByFlowEndpointsCompletionHandler(datagrams []foundation.Data, remoteEndpoints objectivec.IObject, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](n_.ID, objc.Sel("writeDatagrams:sentByFlowEndpoints:completionHandler:"), datagrams, remoteEndpoints, completionHandler)
}/* debug [instance_methods/method]: WriteDatagramsSentByFlowEndpointsCompletionHandler */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for NEAppProxyUDPFlow */

// An object containing information about the local endpoint of the flow.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEAppProxyUDPFlow/localEndpoint
func (n_ NEAppProxyUDPFlow) LocalEndpoint() INWEndpoint {
	rv := objc.Send[NWEndpoint](n_.ID, objc.Sel("localEndpoint"))
	return rv
}/* debug [instance_properties/getter]: localEndpoint */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEAppProxyUDPFlow/localFlowEndpoint-9a8gj
func (n_ NEAppProxyUDPFlow) LocalFlowEndpoint() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](n_.ID, objc.Sel("localFlowEndpoint"))
	return rv
}/* debug [instance_properties/getter]: localFlowEndpoint */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class NEAppProxyUDPFlow */



