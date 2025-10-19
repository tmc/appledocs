// Code generated from Apple documentation for NetworkExtension. DO NOT EDIT.

package networkextension

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [NEPacketTunnelNetworkSettings] class.
var (
	nEPacketTunnelNetworkSettingsClass     _NEPacketTunnelNetworkSettingsClass
	nEPacketTunnelNetworkSettingsClassOnce sync.Once
)

func getNEPacketTunnelNetworkSettingsClass() _NEPacketTunnelNetworkSettingsClass {
	nEPacketTunnelNetworkSettingsClassOnce.Do(func() {
		nEPacketTunnelNetworkSettingsClass = _NEPacketTunnelNetworkSettingsClass{objc.GetClass("NEPacketTunnelNetworkSettings")}
	})
	return nEPacketTunnelNetworkSettingsClass
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




