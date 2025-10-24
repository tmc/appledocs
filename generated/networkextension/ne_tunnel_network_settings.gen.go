// Code generated from Apple documentation for NetworkExtension. DO NOT EDIT.

package networkextension

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [NETunnelNetworkSettings] class.
var (
	NETunnelNetworkSettingsClass     _NETunnelNetworkSettingsClass
	NETunnelNetworkSettingsClassOnce sync.Once
)

func getNETunnelNetworkSettingsClass() _NETunnelNetworkSettingsClass {
	NETunnelNetworkSettingsClassOnce.Do(func() {
		NETunnelNetworkSettingsClass = _NETunnelNetworkSettingsClass{objc.GetClass("NETunnelNetworkSettings")}
	})
	return NETunnelNetworkSettingsClass
}

type _NETunnelNetworkSettingsClass struct {
	class objc.Class
}

// An interface definition for the [NETunnelNetworkSettings] class.
type INETunnelNetworkSettings interface {
	objectivec.IObject
	// properties:
	DnsSettings() INEDNSSettings
	SetDnsSettings(value INEDNSSettings)
	ProxySettings() objc.IObject /* cross-framework: NEProxySettings */
	SetProxySettings(value objc.IObject /* cross-framework: NEProxySettings */)
	TunnelRemoteAddress() objc.IObject /* cross-framework: NSString */
	SetTunnelRemoteAddress(value objc.IObject /* cross-framework: NSString */)
	// methods:
}

// The configuration for a tunnel provider’s virtual interface.


// The configuration for a tunnel provider’s virtual interface.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NETunnelNetworkSettings
type NETunnelNetworkSettings struct {
	objectivec.Object
}

// NETunnelNetworkSettingsFrom constructs a [NETunnelNetworkSettings] from an unsafe.Pointer.
//
// The configuration for a tunnel provider’s virtual interface.
func NETunnelNetworkSettingsFrom(ptr unsafe.Pointer) NETunnelNetworkSettings {
	return NETunnelNetworkSettings{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (nc _NETunnelNetworkSettingsClass) Alloc() NETunnelNetworkSettings {
	rv := objc.Send[NETunnelNetworkSettings](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (nc _NETunnelNetworkSettingsClass) New() NETunnelNetworkSettings {
	rv := objc.Send[NETunnelNetworkSettings](objc.ID(nc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (n_ NETunnelNetworkSettings) Init() NETunnelNetworkSettings {
	rv := objc.Send[NETunnelNetworkSettings](n_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (n_ NETunnelNetworkSettings) Autorelease() NETunnelNetworkSettings {
	rv := objc.Send[NETunnelNetworkSettings](n_.ID, objc.Sel("autorelease"))
	return rv
}

// NewNETunnelNetworkSettings creates a new NETunnelNetworkSettings instance.
func NewNETunnelNetworkSettings() NETunnelNetworkSettings {
	return getNETunnelNetworkSettingsClass().New()
}



// The tunnel DNS settings.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/networkextension/netunnelnetworksettings/dnssettings
func (n_ NETunnelNetworkSettings) DnsSettings() INEDNSSettings {
	rv := objc.Send[NEDNSSettings](n_.ID, objc.Sel("dnsSettings"))
	return rv
}


// The tunnel DNS settings.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/networkextension/netunnelnetworksettings/dnssettings
func (n_ NETunnelNetworkSettings) SetDnsSettings(value INEDNSSettings) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setDnsSettings:"), value)
}


// The tunnel HTTP proxy settings.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/networkextension/netunnelnetworksettings/proxysettings
func (n_ NETunnelNetworkSettings) ProxySettings() objc.IObject /* cross-framework: NEProxySettings */ {
	rv := objc.Send[NEProxySettings](n_.ID, objc.Sel("proxySettings"))
	return rv
}


// The tunnel HTTP proxy settings.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/networkextension/netunnelnetworksettings/proxysettings
func (n_ NETunnelNetworkSettings) SetProxySettings(value objc.IObject /* cross-framework: NEProxySettings */) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setProxySettings:"), value)
}


// The IP address of the tunnel server.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/networkextension/netunnelnetworksettings/tunnelremoteaddress
func (n_ NETunnelNetworkSettings) TunnelRemoteAddress() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](n_.ID, objc.Sel("tunnelRemoteAddress"))
	return rv
}


// The IP address of the tunnel server.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/networkextension/netunnelnetworksettings/tunnelremoteaddress
func (n_ NETunnelNetworkSettings) SetTunnelRemoteAddress(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setTunnelRemoteAddress:"), value)
}



