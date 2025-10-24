// Code generated from Apple documentation for NetworkExtension. DO NOT EDIT.

package networkextension

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

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

// An interface definition for the [NETunnelProvider] class.
type INETunnelProvider interface {
	INEProvider
	// properties:
	AppRules() INEAppRule
	SetAppRules(value INEAppRule)
	ProtocolConfiguration() INEVPNProtocol
	SetProtocolConfiguration(value INEVPNProtocol)
	Reasserting() bool
	SetReasserting(value bool)
	RoutingMethod() unsafe.Pointer
	SetRoutingMethod(value unsafe.Pointer)
	NETunnelProviderErrorDomain() objc.IObject /* cross-framework: NSString */
	// methods:
	SetTunnelNetworkSettingsCompletionHandler(tunnelNetworkSettings INETunnelNetworkSettings, completionHandler unsafe.Pointer)
}

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

// Alloc allocates a new instance without initialization.
func (nc _NETunnelProviderClass) Alloc() NETunnelProvider {
	rv := objc.Send[NETunnelProvider](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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



// Specify the network settings for the current tunneling session.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NETunnelProvider/setTunnelNetworkSettings(_:completionHandler:)
func (n_ NETunnelProvider) SetTunnelNetworkSettingsCompletionHandler(tunnelNetworkSettings INETunnelNetworkSettings, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setTunnelNetworkSettings:completionHandler:"), tunnelNetworkSettings, completionHandler)
}


// The app rules dictating which apps use the current tunneling session.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/networkextension/netunnelprovider/apprules
func (n_ NETunnelProvider) AppRules() INEAppRule {
	rv := objc.Send[NEAppRule](n_.ID, objc.Sel("appRules"))
	return rv
}


// The app rules dictating which apps use the current tunneling session.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/networkextension/netunnelprovider/apprules
func (n_ NETunnelProvider) SetAppRules(value INEAppRule) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setAppRules:"), value)
}


// The configuration of the current tunneling session.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/networkextension/netunnelprovider/protocolconfiguration
func (n_ NETunnelProvider) ProtocolConfiguration() INEVPNProtocol {
	rv := objc.Send[NEVPNProtocol](n_.ID, objc.Sel("protocolConfiguration"))
	return rv
}


// The configuration of the current tunneling session.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/networkextension/netunnelprovider/protocolconfiguration
func (n_ NETunnelProvider) SetProtocolConfiguration(value INEVPNProtocol) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setProtocolConfiguration:"), value)
}


// Indicate to the system that the tunnel is being re-established.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/networkextension/netunnelprovider/reasserting
func (n_ NETunnelProvider) Reasserting() bool {
	rv := objc.Send[bool](n_.ID, objc.Sel("reasserting"))
	return rv
}


// Indicate to the system that the tunnel is being re-established.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/networkextension/netunnelprovider/reasserting
func (n_ NETunnelProvider) SetReasserting(value bool) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setReasserting:"), value)
}


// The method by which network traffic is routed to the tunnel.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/networkextension/netunnelprovider/routingmethod
func (n_ NETunnelProvider) RoutingMethod() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](n_.ID, objc.Sel("routingMethod"))
	return rv
}


// The method by which network traffic is routed to the tunnel.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/networkextension/netunnelprovider/routingmethod
func (n_ NETunnelProvider) SetRoutingMethod(value unsafe.Pointer) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setRoutingMethod:"), value)
}


// The domain used for Tunnel Provider errors.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/networkextension/netunnelprovidererrordomain
func (n_ NETunnelProvider) NETunnelProviderErrorDomain() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](n_.ID, objc.Sel("NETunnelProviderErrorDomain"))
	return rv
}



