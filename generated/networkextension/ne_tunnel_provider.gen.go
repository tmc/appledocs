// Code generated from Apple documentation for NetworkExtension. DO NOT EDIT.

package networkextension

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

/* debug [class.gen.go]: Generating class NETunnelProvider */


/* debug [class_header]: Header for NETunnelProvider */
// The class instance for the [NETunnelProvider] class.
var (
	NETunnelProviderClass     _NETunnelProviderClass
	NETunnelProviderClassOnce sync.Once
)

func getNETunnelProviderClass() _NETunnelProviderClass {
	NETunnelProviderClassOnce.Do(func() {
		NETunnelProviderClass = _NETunnelProviderClass{objc.GetClass("NETunnelProvider")}
	})
	return NETunnelProviderClass
}

type _NETunnelProviderClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for NETunnelProvider */
// An interface definition for the [NETunnelProvider] class.
type INETunnelProvider interface {
	INEProvider
	
/* debug [class_interface_properties]: Properties for NETunnelProvider */
	// properties:
	AppRules() []NEAppRule
	ProtocolConfiguration() INEVPNProtocol
	Reasserting() bool
	SetReasserting(value bool)
	RoutingMethod() NETunnelProviderRoutingMethod
	NETunnelProviderErrorDomain() objc.IObject /* cross-framework: NSString */
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for NETunnelProvider */
	// methods:
	HandleAppMessageCompletionHandler(messageData objc.IObject /* cross-framework: NSData */, completionHandler unsafe.Pointer)
	SetTunnelNetworkSettingsCompletionHandler(tunnelNetworkSettings INETunnelNetworkSettings, completionHandler unsafe.Pointer)
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for NETunnelProvider */
// Alloc allocates a new instance without initialization.
func (nc _NETunnelProviderClass) Alloc() NETunnelProvider {
	rv := objc.Send[NETunnelProvider](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (nc _NETunnelProviderClass) New() NETunnelProvider {
	rv := objc.Send[NETunnelProvider](objc.ID(nc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (n_ NETunnelProvider) Init() NETunnelProvider {
	rv := objc.Send[NETunnelProvider](n_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (n_ NETunnelProvider) Autorelease() NETunnelProvider {
	rv := objc.Send[NETunnelProvider](n_.ID, objc.Sel("autorelease"))
	return rv
}

// NewNETunnelProvider creates a new NETunnelProvider instance.
func NewNETunnelProvider() NETunnelProvider {
	return getNETunnelProviderClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for NETunnelProvider */
// An abstract base class shared by NEPacketTunnelProvider and NEAppProxyProvider.
//
// Each instance corresponds to a single tunneling session, with a single associated configuration.


// An abstract base class shared by NEPacketTunnelProvider and NEAppProxyProvider.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NETunnelProvider
type NETunnelProvider struct {
	NEProvider
}

// NETunnelProviderFrom constructs a [NETunnelProvider] from an unsafe.Pointer.
//
// An abstract base class shared by NEPacketTunnelProvider and NEAppProxyProvider.
func NETunnelProviderFrom(ptr unsafe.Pointer) NETunnelProvider {
	return NETunnelProvider{
		NEProvider: NEProviderFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for NETunnelProvider *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for NETunnelProvider */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for NETunnelProvider */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for NETunnelProvider */

// Handle messages sent by the tunnel provider extension’s containing app.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NETunnelProvider/handleAppMessage(_:completionHandler:)
func (n_ NETunnelProvider) HandleAppMessageCompletionHandler(messageData objc.IObject /* cross-framework: NSData */, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](n_.ID, objc.Sel("handleAppMessage:completionHandler:"), messageData, completionHandler)
}/* debug [instance_methods/method]: HandleAppMessageCompletionHandler */


// Specify the network settings for the current tunneling session.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NETunnelProvider/setTunnelNetworkSettings(_:completionHandler:)
func (n_ NETunnelProvider) SetTunnelNetworkSettingsCompletionHandler(tunnelNetworkSettings INETunnelNetworkSettings, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setTunnelNetworkSettings:completionHandler:"), tunnelNetworkSettings, completionHandler)
}/* debug [instance_methods/method]: SetTunnelNetworkSettingsCompletionHandler */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for NETunnelProvider */

// The app rules dictating which apps use the current tunneling session.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NETunnelProvider/appRules
func (n_ NETunnelProvider) AppRules() []NEAppRule {
	rv := objc.Send[[]NEAppRule](n_.ID, objc.Sel("appRules"))
	return rv
}/* debug [instance_properties/getter]: appRules */


// The configuration of the current tunneling session.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NETunnelProvider/protocolConfiguration
func (n_ NETunnelProvider) ProtocolConfiguration() INEVPNProtocol {
	rv := objc.Send[NEVPNProtocol](n_.ID, objc.Sel("protocolConfiguration"))
	return rv
}/* debug [instance_properties/getter]: protocolConfiguration */


// Indicate to the system that the tunnel is being re-established.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NETunnelProvider/reasserting
func (n_ NETunnelProvider) Reasserting() bool {
	rv := objc.Send[bool](n_.ID, objc.Sel("reasserting"))
	return rv
}/* debug [instance_properties/getter]: reasserting */


// Indicate to the system that the tunnel is being re-established.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NETunnelProvider/reasserting
func (n_ NETunnelProvider) SetReasserting(value bool) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setReasserting:"), value)
}/* debug [instance_properties/setter]: reasserting */


// The method by which network traffic is routed to the tunnel.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NETunnelProvider/routingMethod
func (n_ NETunnelProvider) RoutingMethod() NETunnelProviderRoutingMethod {
	rv := objc.Send[NETunnelProviderRoutingMethod](n_.ID, objc.Sel("routingMethod"))
	return rv
}/* debug [instance_properties/getter]: routingMethod */


// The domain used for Tunnel Provider errors.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/networkextension/netunnelprovidererrordomain
func (n_ NETunnelProvider) NETunnelProviderErrorDomain() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](n_.ID, objc.Sel("NETunnelProviderErrorDomain"))
	return rv
}/* debug [instance_properties/getter]: NETunnelProviderErrorDomain */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class NETunnelProvider */



