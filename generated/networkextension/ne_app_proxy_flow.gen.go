// Code generated from Apple documentation for NetworkExtension. DO NOT EDIT.

package networkextension

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/coretelephony"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class NEAppProxyFlow */


/* debug [class_header]: Header for NEAppProxyFlow */
// The class instance for the [NEAppProxyFlow] class.
var (
	NEAppProxyFlowClass     _NEAppProxyFlowClass
	NEAppProxyFlowClassOnce sync.Once
)

func getNEAppProxyFlowClass() _NEAppProxyFlowClass {
	NEAppProxyFlowClassOnce.Do(func() {
		NEAppProxyFlowClass = _NEAppProxyFlowClass{objc.GetClass("NEAppProxyFlow")}
	})
	return NEAppProxyFlowClass
}

type _NEAppProxyFlowClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for NEAppProxyFlow */
// An interface definition for the [NEAppProxyFlow] class.
type INEAppProxyFlow interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for NEAppProxyFlow */
	// properties:
	IsBound() bool
	MetaData() INEFlowMetaData
	NetworkInterface() objectivec.IObject
	SetNetworkInterface(value objectivec.IObject)
	RemoteHostname() objc.IObject /* cross-framework: NSString */
	NEAppProxyErrorDomain() objc.IObject /* cross-framework: NSString */
	Interface() objectivec.IObject
	SetInterface(value objectivec.IObject)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for NEAppProxyFlow */
	// methods:
	CloseReadWithError(error_ objc.IObject /* cross-framework: Error */)
	CloseWriteWithError(error_ objc.IObject /* cross-framework: Error */)
	OpenWithLocalFlowEndpointCompletionHandler(localEndpoint objectivec.IObject, completionHandler unsafe.Pointer)
	SetMetadata(parameters objectivec.IObject)
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for NEAppProxyFlow */
// Alloc allocates a new instance without initialization.
func (nc _NEAppProxyFlowClass) Alloc() NEAppProxyFlow {
	rv := objc.Send[NEAppProxyFlow](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (nc _NEAppProxyFlowClass) New() NEAppProxyFlow {
	rv := objc.Send[NEAppProxyFlow](objc.ID(nc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (n_ NEAppProxyFlow) Init() NEAppProxyFlow {
	rv := objc.Send[NEAppProxyFlow](n_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (n_ NEAppProxyFlow) Autorelease() NEAppProxyFlow {
	rv := objc.Send[NEAppProxyFlow](n_.ID, objc.Sel("autorelease"))
	return rv
}

// NewNEAppProxyFlow creates a new NEAppProxyFlow instance.
func NewNEAppProxyFlow() NEAppProxyFlow {
	return getNEAppProxyFlowClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for NEAppProxyFlow */
// An abstract base class shared by NEAppProxyTCPFlow and NEAppProxyUDPFlow.
//
// App Proxy Providers receive network connections to be proxied in the form of objects, which are passed to the App Proxy Provider via the method. objects are initially in an unopened state. Before they can be used to transmit network data, they must be opened using the method. When you are finished with a flow, you should call and , and then release the object.


// An abstract base class shared by NEAppProxyTCPFlow and NEAppProxyUDPFlow.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEAppProxyFlow
type NEAppProxyFlow struct {
	objectivec.Object
}

// NEAppProxyFlowFrom constructs a [NEAppProxyFlow] from an unsafe.Pointer.
//
// An abstract base class shared by NEAppProxyTCPFlow and NEAppProxyUDPFlow.
func NEAppProxyFlowFrom(ptr unsafe.Pointer) NEAppProxyFlow {
	return NEAppProxyFlow{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for NEAppProxyFlow *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for NEAppProxyFlow */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for NEAppProxyFlow */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for NEAppProxyFlow */

// Close the flow for further read operations.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEAppProxyFlow/closeReadWithError(_:)
func (n_ NEAppProxyFlow) CloseReadWithError(error_ objc.IObject /* cross-framework: Error */) {
	objc.Send[objc.ID](n_.ID, objc.Sel("closeReadWithError:"), error_)
}/* debug [instance_methods/method]: CloseReadWithError */


// Close the flow for further write operations.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEAppProxyFlow/closeWriteWithError(_:)
func (n_ NEAppProxyFlow) CloseWriteWithError(error_ objc.IObject /* cross-framework: Error */) {
	objc.Send[objc.ID](n_.ID, objc.Sel("closeWriteWithError:"), error_)
}/* debug [instance_methods/method]: CloseWriteWithError */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEAppProxyFlow/openWithLocalFlowEndpoint:completionHandler:
func (n_ NEAppProxyFlow) OpenWithLocalFlowEndpointCompletionHandler(localEndpoint objectivec.IObject, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](n_.ID, objc.Sel("openWithLocalFlowEndpoint:completionHandler:"), localEndpoint, completionHandler)
}/* debug [instance_methods/method]: OpenWithLocalFlowEndpointCompletionHandler */


// Sets the flow’s metadata for use by proxy providers.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEAppProxyFlow/setMetadata(_:)
func (n_ NEAppProxyFlow) SetMetadata(parameters objectivec.IObject) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setMetadata:"), parameters)
}/* debug [instance_methods/method]: SetMetadata */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for NEAppProxyFlow */

// A Boolean value that indicates whether the flow has a binding to a specific interface.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEAppProxyFlow/isBound
func (n_ NEAppProxyFlow) IsBound() bool {
	rv := objc.Send[bool](n_.ID, objc.Sel("isBound"))
	return rv
}/* debug [instance_properties/getter]: isBound */


// A metadata object containing information about the source app of the flow.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEAppProxyFlow/metaData
func (n_ NEAppProxyFlow) MetaData() INEFlowMetaData {
	rv := objc.Send[NEFlowMetaData](n_.ID, objc.Sel("metaData"))
	return rv
}/* debug [instance_properties/getter]: metaData */


// The network interface, if any, used by this flow.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEAppProxyFlow/networkInterface
func (n_ NEAppProxyFlow) NetworkInterface() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](n_.ID, objc.Sel("networkInterface"))
	return rv
}/* debug [instance_properties/getter]: networkInterface */


// The network interface, if any, used by this flow.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEAppProxyFlow/networkInterface
func (n_ NEAppProxyFlow) SetNetworkInterface(value objectivec.IObject) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setNetworkInterface:"), value)
}/* debug [instance_properties/setter]: networkInterface */


// The remote host name for flows created from a hostname.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEAppProxyFlow/remoteHostname
func (n_ NEAppProxyFlow) RemoteHostname() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](n_.ID, objc.Sel("remoteHostname"))
	return rv
}/* debug [instance_properties/getter]: remoteHostname */


// The domain used for app proxy errors.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/networkextension/neappproxyerrordomain
func (n_ NEAppProxyFlow) NEAppProxyErrorDomain() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](n_.ID, objc.Sel("NEAppProxyErrorDomain"))
	return rv
}/* debug [instance_properties/getter]: NEAppProxyErrorDomain */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/networkextension/neappproxyflow/interface
func (n_ NEAppProxyFlow) Interface() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](n_.ID, objc.Sel("interface"))
	return rv
}/* debug [instance_properties/getter]: interface */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/networkextension/neappproxyflow/interface
func (n_ NEAppProxyFlow) SetInterface(value objectivec.IObject) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setInterface:"), value)
}/* debug [instance_properties/setter]: interface */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class NEAppProxyFlow */



