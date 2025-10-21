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
	HandleAppMessageCompletionHandler(messageData foundation.IData, completionHandler unsafe.Pointer)
	SetTunnelNetworkSettingsCompletionHandler(tunnelNetworkSettings INETunnelNetworkSettings, completionHandler unsafe.Pointer)
}

// An abstract base class shared by NEPacketTunnelProvider and NEAppProxyProvider.
//
// Each instance corresponds to a single tunneling session, with a single associated configuration.
//
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


// Handle messages sent by the tunnel provider extension’s containing app.
//
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NETunnelProvider/handleAppMessage(_:completionHandler:)
func (n_ NETunnelProvider) HandleAppMessageCompletionHandler(messageData foundation.IData, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](n_.ID, objc.Sel("handleAppMessage:completionHandler:"), messageData, completionHandler)
}

// Specify the network settings for the current tunneling session.
//
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NETunnelProvider/setTunnelNetworkSettings(_:completionHandler:)
func (n_ NETunnelProvider) SetTunnelNetworkSettingsCompletionHandler(tunnelNetworkSettings INETunnelNetworkSettings, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setTunnelNetworkSettings:completionHandler:"), tunnelNetworkSettings, completionHandler)
}

// The app rules dictating which apps use the current tunneling session.
//
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NETunnelProvider/appRules
func (n_ NETunnelProvider) AppRules() []NEAppRule {
	rv := objc.Send[[]NEAppRule](n_.ID, objc.Sel("appRules"))
	return rv
}

// The configuration of the current tunneling session.
//
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NETunnelProvider/protocolConfiguration
func (n_ NETunnelProvider) ProtocolConfiguration() NEVPNProtocol {
	rv := objc.Send[NEVPNProtocol](n_.ID, objc.Sel("protocolConfiguration"))
	return rv
}

// Indicate to the system that the tunnel is being re-established.
//
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NETunnelProvider/reasserting
func (n_ NETunnelProvider) Reasserting() bool {
	rv := objc.Send[bool](n_.ID, objc.Sel("reasserting"))
	return rv
}


// SetReasserting sets the value of the reasserting property.
// Indicate to the system that the tunnel is being re-established.

//
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NETunnelProvider/reasserting
func (n_ NETunnelProvider) SetReasserting(value bool) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setReasserting:"), value)
}

// The method by which network traffic is routed to the tunnel.
//
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NETunnelProvider/routingMethod
func (n_ NETunnelProvider) RoutingMethod() NETunnelProviderRoutingMethod {
	rv := objc.Send[NETunnelProviderRoutingMethod](n_.ID, objc.Sel("routingMethod"))
	return rv
}

// The domain used for Tunnel Provider errors.
//
// [Full Topic]: https://developer.apple.com/documentation/networkextension/netunnelprovidererrordomain
func (n_ NETunnelProvider) NETunnelProviderErrorDomain() appkit.string {
	rv := objc.Send[appkit.string](n_.ID, objc.Sel("NETunnelProviderErrorDomain"))
	return rv
}



