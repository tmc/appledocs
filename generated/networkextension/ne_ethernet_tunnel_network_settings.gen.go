// Code generated from Apple documentation for NetworkExtension. DO NOT EDIT.

package networkextension

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [NEEthernetTunnelNetworkSettings] class.
var (
	nEEthernetTunnelNetworkSettingsClass     _NEEthernetTunnelNetworkSettingsClass
	nEEthernetTunnelNetworkSettingsClassOnce sync.Once
)

func getNEEthernetTunnelNetworkSettingsClass() _NEEthernetTunnelNetworkSettingsClass {
	nEEthernetTunnelNetworkSettingsClassOnce.Do(func() {
		nEEthernetTunnelNetworkSettingsClass = _NEEthernetTunnelNetworkSettingsClass{objc.GetClass("NEEthernetTunnelNetworkSettings")}
	})
	return nEEthernetTunnelNetworkSettingsClass
}

type _NEEthernetTunnelNetworkSettingsClass struct {
	class objc.Class
}

// An interface definition for the [NEEthernetTunnelNetworkSettings] class.
type INEEthernetTunnelNetworkSettings interface {
	INEPacketTunnelNetworkSettings
}

// The network settings for an ethernet-based VPN tunnel. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEEthernetTunnelNetworkSettings
type NEEthernetTunnelNetworkSettings struct {
	NEPacketTunnelNetworkSettings
}

// NEEthernetTunnelNetworkSettingsFrom constructs a [NEEthernetTunnelNetworkSettings] from an unsafe.Pointer.
//
// The network settings for an ethernet-based VPN tunnel.
func NEEthernetTunnelNetworkSettingsFrom(ptr unsafe.Pointer) NEEthernetTunnelNetworkSettings {
	return NEEthernetTunnelNetworkSettings{
		NEPacketTunnelNetworkSettings: NEPacketTunnelNetworkSettingsFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (nc _NEEthernetTunnelNetworkSettingsClass) Alloc() NEEthernetTunnelNetworkSettings {
	rv := objc.Send[NEEthernetTunnelNetworkSettings](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (nc _NEEthernetTunnelNetworkSettingsClass) New() NEEthernetTunnelNetworkSettings {
	rv := objc.Send[NEEthernetTunnelNetworkSettings](objc.ID(nc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (n_ NEEthernetTunnelNetworkSettings) Init() NEEthernetTunnelNetworkSettings {
	rv := objc.Send[NEEthernetTunnelNetworkSettings](n_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (n_ NEEthernetTunnelNetworkSettings) Autorelease() NEEthernetTunnelNetworkSettings {
	rv := objc.Send[NEEthernetTunnelNetworkSettings](n_.ID, objc.Sel("autorelease"))
	return rv
}

// NewNEEthernetTunnelNetworkSettings creates a new NEEthernetTunnelNetworkSettings instance.
func NewNEEthernetTunnelNetworkSettings() NEEthernetTunnelNetworkSettings {
	return getNEEthernetTunnelNetworkSettingsClass().New()
}




