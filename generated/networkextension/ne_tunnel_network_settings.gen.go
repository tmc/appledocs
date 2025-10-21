// Code generated from Apple documentation for NetworkExtension. DO NOT EDIT.

package networkextension

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
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
}

// The configuration for a tunnel provider’s virtual interface.
//
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


// Initialize a object.
//
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NETunnelNetworkSettings/init(tunnelRemoteAddress:)
func NewNETunnelNetworkSettingsWithTunnelRemoteAddress(address string) NETunnelNetworkSettings {
	instance := getNETunnelNetworkSettingsClass().Alloc()
	rv := objc.Send[NETunnelNetworkSettings](instance.ID, objc.Sel("initWithTunnelRemoteAddress:"), objc.String(address))
	rv.Autorelease()
	return rv
}


// The tunnel DNS settings.
//
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NETunnelNetworkSettings/dnsSettings
func (n_ NETunnelNetworkSettings) DNSSettings() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](n_.ID, objc.Sel("DNSSettings"))
	return rv
}


// SetDNSSettings sets the value of the DNSSettings property.
// The tunnel DNS settings.

//
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NETunnelNetworkSettings/dnsSettings
func (n_ NETunnelNetworkSettings) SetDNSSettings(value unsafe.Pointer) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setDNSSettings:"), value)
}
// The tunnel HTTP proxy settings.
//
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NETunnelNetworkSettings/proxySettings
func (n_ NETunnelNetworkSettings) ProxySettings() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](n_.ID, objc.Sel("proxySettings"))
	return rv
}


// SetProxySettings sets the value of the proxySettings property.
// The tunnel HTTP proxy settings.

//
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NETunnelNetworkSettings/proxySettings
func (n_ NETunnelNetworkSettings) SetProxySettings(value unsafe.Pointer) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setProxySettings:"), value)
}
// The IP address of the tunnel server.
//
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NETunnelNetworkSettings/tunnelRemoteAddress
func (n_ NETunnelNetworkSettings) TunnelRemoteAddress() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](n_.ID, objc.Sel("tunnelRemoteAddress"))
	return rv
}


