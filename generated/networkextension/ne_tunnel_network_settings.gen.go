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
	

	// properties:
	DNSSettings() INEDNSSettings
	SetDNSSettings(value INEDNSSettings)
	ProxySettings() INEProxySettings
	SetProxySettings(value INEProxySettings)
	TunnelRemoteAddress() foundation.foundation.INSString


	

	// methods:


}





// Alloc allocates a new instance without initialization.
func (nc _NETunnelNetworkSettingsClass) Alloc() NETunnelNetworkSettings {
	rv := objc.Send[NETunnelNetworkSettings](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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






// Initialize a object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NETunnelNetworkSettings/init(tunnelRemoteAddress:)
func NewNETunnelNetworkSettingsWithTunnelRemoteAddress(address foundation.foundation.INSString) NETunnelNetworkSettings {
	instance := getNETunnelNetworkSettingsClass().Alloc()
	rv := objc.Send[NETunnelNetworkSettings](instance.ID, objc.Sel("initWithTunnelRemoteAddress:"), address)
	rv.Autorelease()
	return rv
}






















// The tunnel DNS settings.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NETunnelNetworkSettings/dnsSettings
func (n_ NETunnelNetworkSettings) DNSSettings() INEDNSSettings {
	rv := objc.Send[NEDNSSettings](n_.ID, objc.Sel("DNSSettings"))
	return rv
}


// The tunnel DNS settings.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NETunnelNetworkSettings/dnsSettings
func (n_ NETunnelNetworkSettings) SetDNSSettings(value INEDNSSettings) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setDNSSettings:"), value)
}


// The tunnel HTTP proxy settings.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NETunnelNetworkSettings/proxySettings
func (n_ NETunnelNetworkSettings) ProxySettings() INEProxySettings {
	rv := objc.Send[NEProxySettings](n_.ID, objc.Sel("proxySettings"))
	return rv
}


// The tunnel HTTP proxy settings.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NETunnelNetworkSettings/proxySettings
func (n_ NETunnelNetworkSettings) SetProxySettings(value INEProxySettings) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setProxySettings:"), value)
}


// The IP address of the tunnel server.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NETunnelNetworkSettings/tunnelRemoteAddress
func (n_ NETunnelNetworkSettings) TunnelRemoteAddress() foundation.foundation.INSString {
	rv := objc.Send[foundation.NSString](n_.ID, objc.Sel("tunnelRemoteAddress"))
	return rv
}







