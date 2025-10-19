// Code generated from Apple documentation for NetworkExtension. DO NOT EDIT.

package networkextension

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [NETunnelNetworkSettings] class.
var nETunnelNetworkSettingsClass = _NETunnelNetworkSettingsClass{objc.GetClass("NETunnelNetworkSettings")}

type _NETunnelNetworkSettingsClass struct {
	class objc.Class
}

// An interface definition for the [NETunnelNetworkSettings] class.
type INETunnelNetworkSettings interface {
	objectivec.IObject
}

// The configuration for a tunnel provider’s virtual interface. [Full Topic]
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

// New creates and returns a new instance with a +1 retain count.
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
	return nETunnelNetworkSettingsClass.New()
}




