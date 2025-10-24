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

/* debug [class.gen.go]: Generating class NEDNSProxyProvider */


/* debug [class_header]: Header for NEDNSProxyProvider */
// The class instance for the [NEDNSProxyProvider] class.
var (
	NEDNSProxyProviderClass     _NEDNSProxyProviderClass
	NEDNSProxyProviderClassOnce sync.Once
)

func getNEDNSProxyProviderClass() _NEDNSProxyProviderClass {
	NEDNSProxyProviderClassOnce.Do(func() {
		NEDNSProxyProviderClass = _NEDNSProxyProviderClass{objc.GetClass("NEDNSProxyProvider")}
	})
	return NEDNSProxyProviderClass
}

type _NEDNSProxyProviderClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for NEDNSProxyProvider */
// An interface definition for the [NEDNSProxyProvider] class.
type INEDNSProxyProvider interface {
	INEProvider
	
/* debug [class_interface_properties]: Properties for NEDNSProxyProvider */
	// properties:
	SystemDNSSettings() []NEDNSSettings
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for NEDNSProxyProvider */
	// methods:
	CancelProxyWithError(error_ objc.IObject /* cross-framework: Error */)
	HandleNewFlow(flow INEAppProxyFlow) bool
	HandleNewUDPFlowInitialRemoteFlowEndpoint(flow INEAppProxyUDPFlow, remoteEndpoint objectivec.IObject) bool
	StartProxyWithOptionsCompletionHandler(options foundation.IDictionary, completionHandler unsafe.Pointer)
	StopProxyWithReasonCompletionHandler(reason NEProviderStopReason, completionHandler unsafe.Pointer)
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for NEDNSProxyProvider */
// Alloc allocates a new instance without initialization.
func (nc _NEDNSProxyProviderClass) Alloc() NEDNSProxyProvider {
	rv := objc.Send[NEDNSProxyProvider](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (nc _NEDNSProxyProviderClass) New() NEDNSProxyProvider {
	rv := objc.Send[NEDNSProxyProvider](objc.ID(nc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (n_ NEDNSProxyProvider) Init() NEDNSProxyProvider {
	rv := objc.Send[NEDNSProxyProvider](n_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (n_ NEDNSProxyProvider) Autorelease() NEDNSProxyProvider {
	rv := objc.Send[NEDNSProxyProvider](n_.ID, objc.Sel("autorelease"))
	return rv
}

// NewNEDNSProxyProvider creates a new NEDNSProxyProvider instance.
func NewNEDNSProxyProvider() NEDNSProxyProvider {
	return getNEDNSProxyProviderClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for NEDNSProxyProvider */
// The principal class for a DNS proxy provider app extension.
//
// A DNS proxy allows your app to intercept all DNS traffic generated on a device. You can use this capability to provide services like DNS traffic encryption, typically by redirecting DNS traffic to your own server. You usually do this in the context of managed devices, such as those owned by a school or an enterprise. You create a DNS proxy as an app extension based on a custom subclass of the class. Once active, the proxy receives access to flows of DNS traffic in the form of instances. Each flow corresponds to a socket opened by an app to UDP port 53 or TCP port 53. Your DNS proxy provider acts as a transparent DNS proxy for the flows of network data that it receives. When you subclass , you must provide implementations for the following methods:


// The principal class for a DNS proxy provider app extension.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEDNSProxyProvider
type NEDNSProxyProvider struct {
	NEProvider
}

// NEDNSProxyProviderFrom constructs a [NEDNSProxyProvider] from an unsafe.Pointer.
//
// The principal class for a DNS proxy provider app extension.
func NEDNSProxyProviderFrom(ptr unsafe.Pointer) NEDNSProxyProvider {
	return NEDNSProxyProvider{
		NEProvider: NEProviderFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for NEDNSProxyProvider *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for NEDNSProxyProvider */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for NEDNSProxyProvider */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for NEDNSProxyProvider */

// Cancels the DNS proxy.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEDNSProxyProvider/cancelProxyWithError(_:)
func (n_ NEDNSProxyProvider) CancelProxyWithError(error_ objc.IObject /* cross-framework: Error */) {
	objc.Send[objc.ID](n_.ID, objc.Sel("cancelProxyWithError:"), error_)
}/* debug [instance_methods/method]: CancelProxyWithError */


// Handles a new flow of DNS traffic.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEDNSProxyProvider/handleNewFlow(_:)
func (n_ NEDNSProxyProvider) HandleNewFlow(flow INEAppProxyFlow) bool {
	rv := objc.Send[bool](n_.ID, objc.Sel("handleNewFlow:"), flow)
	return rv
}/* debug [instance_methods/method]: HandleNewFlow */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEDNSProxyProvider/handleNewUDPFlow:initialRemoteFlowEndpoint:
func (n_ NEDNSProxyProvider) HandleNewUDPFlowInitialRemoteFlowEndpoint(flow INEAppProxyUDPFlow, remoteEndpoint objectivec.IObject) bool {
	rv := objc.Send[bool](n_.ID, objc.Sel("handleNewUDPFlow:initialRemoteFlowEndpoint:"), flow, remoteEndpoint)
	return rv
}/* debug [instance_methods/method]: HandleNewUDPFlowInitialRemoteFlowEndpoint */


// Starts the DNS proxy.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEDNSProxyProvider/startProxy(options:completionHandler:)
func (n_ NEDNSProxyProvider) StartProxyWithOptionsCompletionHandler(options foundation.IDictionary, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](n_.ID, objc.Sel("startProxyWithOptions:completionHandler:"), options, completionHandler)
}/* debug [instance_methods/method]: StartProxyWithOptionsCompletionHandler */


// Stops the DNS proxy.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEDNSProxyProvider/stopProxy(with:completionHandler:)
func (n_ NEDNSProxyProvider) StopProxyWithReasonCompletionHandler(reason NEProviderStopReason, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](n_.ID, objc.Sel("stopProxyWithReason:completionHandler:"), reason, completionHandler)
}/* debug [instance_methods/method]: StopProxyWithReasonCompletionHandler */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for NEDNSProxyProvider */

// The current system DNS settings.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEDNSProxyProvider/systemDNSSettings
func (n_ NEDNSProxyProvider) SystemDNSSettings() []NEDNSSettings {
	rv := objc.Send[[]NEDNSSettings](n_.ID, objc.Sel("systemDNSSettings"))
	return rv
}/* debug [instance_properties/getter]: systemDNSSettings */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class NEDNSProxyProvider */



