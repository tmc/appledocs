// Code generated from Apple documentation for NetworkExtension. DO NOT EDIT.

package networkextension

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
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
}

// The configuration for a packet tunnel provider’s virtual interface.
//
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
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEPacketTunnelNetworkSettings/ipv4Settings
func (n_ NEPacketTunnelNetworkSettings) IPv4Settings() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](n_.ID, objc.Sel("IPv4Settings"))
	return rv
}


// SetIPv4Settings sets the value of the IPv4Settings property.
// The tunnel IP version 4 settings.

//
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEPacketTunnelNetworkSettings/ipv4Settings
func (n_ NEPacketTunnelNetworkSettings) SetIPv4Settings(value unsafe.Pointer) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setIPv4Settings:"), value)
}

// The tunnel IP version 6 settings.
//
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEPacketTunnelNetworkSettings/ipv6Settings
func (n_ NEPacketTunnelNetworkSettings) IPv6Settings() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](n_.ID, objc.Sel("IPv6Settings"))
	return rv
}


// SetIPv6Settings sets the value of the IPv6Settings property.
// The tunnel IP version 6 settings.

//
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEPacketTunnelNetworkSettings/ipv6Settings
func (n_ NEPacketTunnelNetworkSettings) SetIPv6Settings(value unsafe.Pointer) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setIPv6Settings:"), value)
}

// The size of the maximum trasnmission unit, in bytes.
//
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEPacketTunnelNetworkSettings/mtu
func (n_ NEPacketTunnelNetworkSettings) MTU() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](n_.ID, objc.Sel("MTU"))
	return rv
}


// SetMTU sets the value of the MTU property.
// The size of the maximum trasnmission unit, in bytes.

//
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEPacketTunnelNetworkSettings/mtu
func (n_ NEPacketTunnelNetworkSettings) SetMTU(value unsafe.Pointer) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setMTU:"), value)
}

// The number of bytes added to each tunneled packet for storing tunneling protocol headers.
//
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEPacketTunnelNetworkSettings/tunnelOverheadBytes
func (n_ NEPacketTunnelNetworkSettings) TunnelOverheadBytes() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](n_.ID, objc.Sel("tunnelOverheadBytes"))
	return rv
}


// SetTunnelOverheadBytes sets the value of the tunnelOverheadBytes property.
// The number of bytes added to each tunneled packet for storing tunneling protocol headers.

//
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEPacketTunnelNetworkSettings/tunnelOverheadBytes
func (n_ NEPacketTunnelNetworkSettings) SetTunnelOverheadBytes(value unsafe.Pointer) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setTunnelOverheadBytes:"), value)
}



