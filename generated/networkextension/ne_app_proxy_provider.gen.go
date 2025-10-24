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

/* debug [class.gen.go]: Generating class NEAppProxyProvider */


/* debug [class_header]: Header for NEAppProxyProvider */
// The class instance for the [NEAppProxyProvider] class.
var (
	NEAppProxyProviderClass     _NEAppProxyProviderClass
	NEAppProxyProviderClassOnce sync.Once
)

func getNEAppProxyProviderClass() _NEAppProxyProviderClass {
	NEAppProxyProviderClassOnce.Do(func() {
		NEAppProxyProviderClass = _NEAppProxyProviderClass{objc.GetClass("NEAppProxyProvider")}
	})
	return NEAppProxyProviderClass
}

type _NEAppProxyProviderClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for NEAppProxyProvider */
// An interface definition for the [NEAppProxyProvider] class.
type INEAppProxyProvider interface {
	INETunnelProvider
	
/* debug [class_interface_properties]: Properties for NEAppProxyProvider */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for NEAppProxyProvider */
	// methods:
	CancelProxyWithError(error_ objc.IObject /* cross-framework: Error */)
	HandleNewFlow(flow INEAppProxyFlow) bool
	HandleNewUDPFlowInitialRemoteFlowEndpoint(flow INEAppProxyUDPFlow, remoteEndpoint objectivec.IObject) bool
	StartProxyWithOptionsCompletionHandler(options foundation.IDictionary, completionHandler unsafe.Pointer)
	StopProxyWithReasonCompletionHandler(reason NEProviderStopReason, completionHandler unsafe.Pointer)
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for NEAppProxyProvider */
// Alloc allocates a new instance without initialization.
func (nc _NEAppProxyProviderClass) Alloc() NEAppProxyProvider {
	rv := objc.Send[NEAppProxyProvider](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (nc _NEAppProxyProviderClass) New() NEAppProxyProvider {
	rv := objc.Send[NEAppProxyProvider](objc.ID(nc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (n_ NEAppProxyProvider) Init() NEAppProxyProvider {
	rv := objc.Send[NEAppProxyProvider](n_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (n_ NEAppProxyProvider) Autorelease() NEAppProxyProvider {
	rv := objc.Send[NEAppProxyProvider](n_.ID, objc.Sel("autorelease"))
	return rv
}

// NewNEAppProxyProvider creates a new NEAppProxyProvider instance.
func NewNEAppProxyProvider() NEAppProxyProvider {
	return getNEAppProxyProviderClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for NEAppProxyProvider */
// The principal class for an app proxy provider app extension.
//
// The class provides access to flows of network data in the form of objects. Each object corresponds to a socket opened by an app that matches the app rules specified in the current App Proxy configuration. Your App Proxy Provider acts as a transparent network proxy for the flows of network data that it receives.


// The principal class for an app proxy provider app extension.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEAppProxyProvider
type NEAppProxyProvider struct {
	NETunnelProvider
}

// NEAppProxyProviderFrom constructs a [NEAppProxyProvider] from an unsafe.Pointer.
//
// The principal class for an app proxy provider app extension.
func NEAppProxyProviderFrom(ptr unsafe.Pointer) NEAppProxyProvider {
	return NEAppProxyProvider{
		NETunnelProvider: NETunnelProviderFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for NEAppProxyProvider *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for NEAppProxyProvider */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for NEAppProxyProvider */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for NEAppProxyProvider */

// Stop the network proxy from the App Proxy Provider.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEAppProxyProvider/cancelProxyWithError(_:)
func (n_ NEAppProxyProvider) CancelProxyWithError(error_ objc.IObject /* cross-framework: Error */) {
	objc.Send[objc.ID](n_.ID, objc.Sel("cancelProxyWithError:"), error_)
}/* debug [instance_methods/method]: CancelProxyWithError */


// Handle a new flow of network data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEAppProxyProvider/handleNewFlow(_:)
func (n_ NEAppProxyProvider) HandleNewFlow(flow INEAppProxyFlow) bool {
	rv := objc.Send[bool](n_.ID, objc.Sel("handleNewFlow:"), flow)
	return rv
}/* debug [instance_methods/method]: HandleNewFlow */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEAppProxyProvider/handleNewUDPFlow:initialRemoteFlowEndpoint:
func (n_ NEAppProxyProvider) HandleNewUDPFlowInitialRemoteFlowEndpoint(flow INEAppProxyUDPFlow, remoteEndpoint objectivec.IObject) bool {
	rv := objc.Send[bool](n_.ID, objc.Sel("handleNewUDPFlow:initialRemoteFlowEndpoint:"), flow, remoteEndpoint)
	return rv
}/* debug [instance_methods/method]: HandleNewUDPFlowInitialRemoteFlowEndpoint */


// Start the network proxy.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEAppProxyProvider/startProxy(options:completionHandler:)
func (n_ NEAppProxyProvider) StartProxyWithOptionsCompletionHandler(options foundation.IDictionary, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](n_.ID, objc.Sel("startProxyWithOptions:completionHandler:"), options, completionHandler)
}/* debug [instance_methods/method]: StartProxyWithOptionsCompletionHandler */


// Stop the network proxy.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEAppProxyProvider/stopProxy(with:completionHandler:)
func (n_ NEAppProxyProvider) StopProxyWithReasonCompletionHandler(reason NEProviderStopReason, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](n_.ID, objc.Sel("stopProxyWithReason:completionHandler:"), reason, completionHandler)
}/* debug [instance_methods/method]: StopProxyWithReasonCompletionHandler */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for NEAppProxyProvider */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class NEAppProxyProvider */



