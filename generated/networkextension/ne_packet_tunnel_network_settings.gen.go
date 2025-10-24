// Code generated from Apple documentation for NetworkExtension. DO NOT EDIT.

package networkextension

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [NEPacketTunnelNetworkSettings] class.
var (
	NEPacketTunnelNetworkSettingsClass     _NEPacketTunnelNetworkSettingsClass
	NEPacketTunnelNetworkSettingsClassOnce sync.Once
)

func getNEPacketTunnelNetworkSettingsClass() _NEPacketTunnelNetworkSettingsClass {
	NEPacketTunnelNetworkSettingsClassOnce.Do(func() {
		NEPacketTunnelNetworkSettingsClass = _NEPacketTunnelNetworkSettingsClass{objc.GetClass("NEPacketTunnelNetworkSettings")}
	})
	return NEPacketTunnelNetworkSettingsClass
}

type _NEPacketTunnelNetworkSettingsClass struct {
	class objc.Class
}

// An interface definition for the [NEPacketTunnelNetworkSettings] class.
type INEPacketTunnelNetworkSettings interface {
	INETunnelNetworkSettings
	// properties:
	IPv4Settings() objc.IObject /* cross-framework: NEIPv4Settings */
	SetIPv4Settings(value objc.IObject /* cross-framework: NEIPv4Settings */)
	Ipv6Settings() objc.IObject /* cross-framework: NEIPv6Settings */
	SetIpv6Settings(value objc.IObject /* cross-framework: NEIPv6Settings */)
	Mtu() objc.IObject /* cross-framework: NSNumber */
	SetMtu(value objc.IObject /* cross-framework: NSNumber */)
	TunnelOverheadBytes() objc.IObject /* cross-framework: NSNumber */
	SetTunnelOverheadBytes(value objc.IObject /* cross-framework: NSNumber */)
	// methods:
}

// The configuration for a packet tunnel provider’s virtual interface.


// The configuration for a packet tunnel provider’s virtual interface.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEPacketTunnelNetworkSettings
type NEPacketTunnelNetworkSettings struct {
	NETunnelNetworkSettings
}

// NEPacketTunnelNetworkSettingsFrom constructs a [NEPacketTunnelNetworkSettings] from an unsafe.Pointer.
//
// The configuration for a packet tunnel provider’s virtual interface.
func NEPacketTunnelNetworkSettingsFrom(ptr unsafe.Pointer) NEPacketTunnelNetworkSettings {
	return NEPacketTunnelNetworkSettings{
		NETunnelNetworkSettings: NETunnelNetworkSettingsFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (nc _NEPacketTunnelNetworkSettingsClass) Alloc() NEPacketTunnelNetworkSettings {
	rv := objc.Send[NEPacketTunnelNetworkSettings](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (nc _NEPacketTunnelNetworkSettingsClass) New() NEPacketTunnelNetworkSettings {
	rv := objc.Send[NEPacketTunnelNetworkSettings](objc.ID(nc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (n_ NEPacketTunnelNetworkSettings) Init() NEPacketTunnelNetworkSettings {
	rv := objc.Send[NEPacketTunnelNetworkSettings](n_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (n_ NEPacketTunnelNetworkSettings) Autorelease() NEPacketTunnelNetworkSettings {
	rv := objc.Send[NEPacketTunnelNetworkSettings](n_.ID, objc.Sel("autorelease"))
	return rv
}

// NewNEPacketTunnelNetworkSettings creates a new NEPacketTunnelNetworkSettings instance.
func NewNEPacketTunnelNetworkSettings() NEPacketTunnelNetworkSettings {
	return getNEPacketTunnelNetworkSettingsClass().New()
}



// The tunnel IP version 4 settings.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEPacketTunnelNetworkSettings/ipv4Settings
func (n_ NEPacketTunnelNetworkSettings) IPv4Settings() objc.IObject /* cross-framework: NEIPv4Settings */ {
	rv := objc.Send[NEIPv4Settings](n_.ID, objc.Sel("IPv4Settings"))
	return rv
}


// The tunnel IP version 4 settings.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEPacketTunnelNetworkSettings/ipv4Settings
func (n_ NEPacketTunnelNetworkSettings) SetIPv4Settings(value objc.IObject /* cross-framework: NEIPv4Settings */) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setIPv4Settings:"), value)
}


// The tunnel IP version 6 settings.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/networkextension/nepackettunnelnetworksettings/ipv6settings
func (n_ NEPacketTunnelNetworkSettings) Ipv6Settings() objc.IObject /* cross-framework: NEIPv6Settings */ {
	rv := objc.Send[NEIPv6Settings](n_.ID, objc.Sel("ipv6Settings"))
	return rv
}


// The tunnel IP version 6 settings.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/networkextension/nepackettunnelnetworksettings/ipv6settings
func (n_ NEPacketTunnelNetworkSettings) SetIpv6Settings(value objc.IObject /* cross-framework: NEIPv6Settings */) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setIpv6Settings:"), value)
}


// The size of the maximum trasnmission unit, in bytes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/networkextension/nepackettunnelnetworksettings/mtu
func (n_ NEPacketTunnelNetworkSettings) Mtu() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](n_.ID, objc.Sel("mtu"))
	return rv
}


// The size of the maximum trasnmission unit, in bytes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/networkextension/nepackettunnelnetworksettings/mtu
func (n_ NEPacketTunnelNetworkSettings) SetMtu(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setMtu:"), value)
}


// The number of bytes added to each tunneled packet for storing tunneling protocol headers.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/networkextension/nepackettunnelnetworksettings/tunneloverheadbytes
func (n_ NEPacketTunnelNetworkSettings) TunnelOverheadBytes() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](n_.ID, objc.Sel("tunnelOverheadBytes"))
	return rv
}


// The number of bytes added to each tunneled packet for storing tunneling protocol headers.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/networkextension/nepackettunnelnetworksettings/tunneloverheadbytes
func (n_ NEPacketTunnelNetworkSettings) SetTunnelOverheadBytes(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setTunnelOverheadBytes:"), value)
}



