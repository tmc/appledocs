// Code generated from Apple documentation for NetworkExtension. DO NOT EDIT.

package networkextension

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class NEAppProxyTCPFlow */


/* debug [class_header]: Header for NEAppProxyTCPFlow */
// The class instance for the [NEAppProxyTCPFlow] class.
var (
	NEAppProxyTCPFlowClass     _NEAppProxyTCPFlowClass
	NEAppProxyTCPFlowClassOnce sync.Once
)

func getNEAppProxyTCPFlowClass() _NEAppProxyTCPFlowClass {
	NEAppProxyTCPFlowClassOnce.Do(func() {
		NEAppProxyTCPFlowClass = _NEAppProxyTCPFlowClass{objc.GetClass("NEAppProxyTCPFlow")}
	})
	return NEAppProxyTCPFlowClass
}

type _NEAppProxyTCPFlowClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for NEAppProxyTCPFlow */
// An interface definition for the [NEAppProxyTCPFlow] class.
type INEAppProxyTCPFlow interface {
	INEAppProxyFlow
	
/* debug [class_interface_properties]: Properties for NEAppProxyTCPFlow */
	// properties:
	RemoteEndpoint() INWEndpoint
	RemoteFlowEndpoint() objectivec.IObject
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for NEAppProxyTCPFlow */
	// methods:
	ReadDataWithCompletionHandler(completionHandler unsafe.Pointer)
	WriteDataWithCompletionHandler(data objc.IObject /* cross-framework: NSData */, completionHandler unsafe.Pointer)
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for NEAppProxyTCPFlow */
// Alloc allocates a new instance without initialization.
func (nc _NEAppProxyTCPFlowClass) Alloc() NEAppProxyTCPFlow {
	rv := objc.Send[NEAppProxyTCPFlow](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (nc _NEAppProxyTCPFlowClass) New() NEAppProxyTCPFlow {
	rv := objc.Send[NEAppProxyTCPFlow](objc.ID(nc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (n_ NEAppProxyTCPFlow) Init() NEAppProxyTCPFlow {
	rv := objc.Send[NEAppProxyTCPFlow](n_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (n_ NEAppProxyTCPFlow) Autorelease() NEAppProxyTCPFlow {
	rv := objc.Send[NEAppProxyTCPFlow](n_.ID, objc.Sel("autorelease"))
	return rv
}

// NewNEAppProxyTCPFlow creates a new NEAppProxyTCPFlow instance.
func NewNEAppProxyTCPFlow() NEAppProxyTCPFlow {
	return getNEAppProxyTCPFlowClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for NEAppProxyTCPFlow */
// An object for reading and writing data to and from a TCP connection being proxied by the provider.
//
// App Proxy Providers receive TCP connections to be proxied in the form of objects.


// An object for reading and writing data to and from a TCP connection being proxied by the provider.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEAppProxyTCPFlow
type NEAppProxyTCPFlow struct {
	NEAppProxyFlow
}

// NEAppProxyTCPFlowFrom constructs a [NEAppProxyTCPFlow] from an unsafe.Pointer.
//
// An object for reading and writing data to and from a TCP connection being proxied by the provider.
func NEAppProxyTCPFlowFrom(ptr unsafe.Pointer) NEAppProxyTCPFlow {
	return NEAppProxyTCPFlow{
		NEAppProxyFlow: NEAppProxyFlowFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for NEAppProxyTCPFlow *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for NEAppProxyTCPFlow */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for NEAppProxyTCPFlow */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for NEAppProxyTCPFlow */

// Read data from the flow.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEAppProxyTCPFlow/readData(completionHandler:)
func (n_ NEAppProxyTCPFlow) ReadDataWithCompletionHandler(completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](n_.ID, objc.Sel("readDataWithCompletionHandler:"), completionHandler)
}/* debug [instance_methods/method]: ReadDataWithCompletionHandler */


// Write data to the flow.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEAppProxyTCPFlow/write(_:withCompletionHandler:)
func (n_ NEAppProxyTCPFlow) WriteDataWithCompletionHandler(data objc.IObject /* cross-framework: NSData */, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](n_.ID, objc.Sel("writeData:withCompletionHandler:"), data, completionHandler)
}/* debug [instance_methods/method]: WriteDataWithCompletionHandler */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for NEAppProxyTCPFlow */

// An object containing information about the intended remote endpoint of the flow.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEAppProxyTCPFlow/remoteEndpoint
func (n_ NEAppProxyTCPFlow) RemoteEndpoint() INWEndpoint {
	rv := objc.Send[NWEndpoint](n_.ID, objc.Sel("remoteEndpoint"))
	return rv
}/* debug [instance_properties/getter]: remoteEndpoint */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEAppProxyTCPFlow/remoteFlowEndpoint-9lvob
func (n_ NEAppProxyTCPFlow) RemoteFlowEndpoint() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](n_.ID, objc.Sel("remoteFlowEndpoint"))
	return rv
}/* debug [instance_properties/getter]: remoteFlowEndpoint */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class NEAppProxyTCPFlow */



