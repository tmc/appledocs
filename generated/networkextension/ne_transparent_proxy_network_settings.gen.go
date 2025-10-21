// Code generated from Apple documentation for NetworkExtension. DO NOT EDIT.

package networkextension

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [NETransparentProxyNetworkSettings] class.
var (
	NETransparentProxyNetworkSettingsClass     _NETransparentProxyNetworkSettingsClass
	NETransparentProxyNetworkSettingsClassOnce sync.Once
)

func getNETransparentProxyNetworkSettingsClass() _NETransparentProxyNetworkSettingsClass {
	NETransparentProxyNetworkSettingsClassOnce.Do(func() {
		NETransparentProxyNetworkSettingsClass = _NETransparentProxyNetworkSettingsClass{objc.GetClass("NETransparentProxyNetworkSettings")}
	})
	return NETransparentProxyNetworkSettingsClass
}

type _NETransparentProxyNetworkSettingsClass struct {
	class objc.Class
}

// An interface definition for the [NETransparentProxyNetworkSettings] class.
type INETransparentProxyNetworkSettings interface {
	INETunnelNetworkSettings
}

// A specification of what traffic to route through a transparent proxy.
//
// A proxy network settings object contains two properties: an array of rules to include traffic ( ) and an array of rules to exclude traffic ( ). The exclusion rules take prirority. Therefore, if a given flow matches any of the , evaluation ends and the flow doesn’t route to the proxy. If there’s no match, then evaluation continues and attempts to match the flow against the .
//
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NETransparentProxyNetworkSettings
type NETransparentProxyNetworkSettings struct {
	NETunnelNetworkSettings
}

// NETransparentProxyNetworkSettingsFrom constructs a [NETransparentProxyNetworkSettings] from an unsafe.Pointer.
//
// A specification of what traffic to route through a transparent proxy.
func NETransparentProxyNetworkSettingsFrom(ptr unsafe.Pointer) NETransparentProxyNetworkSettings {
	return NETransparentProxyNetworkSettings{
		NETunnelNetworkSettings: NETunnelNetworkSettingsFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (nc _NETransparentProxyNetworkSettingsClass) Alloc() NETransparentProxyNetworkSettings {
	rv := objc.Send[NETransparentProxyNetworkSettings](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (nc _NETransparentProxyNetworkSettingsClass) New() NETransparentProxyNetworkSettings {
	rv := objc.Send[NETransparentProxyNetworkSettings](objc.ID(nc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (n_ NETransparentProxyNetworkSettings) Init() NETransparentProxyNetworkSettings {
	rv := objc.Send[NETransparentProxyNetworkSettings](n_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (n_ NETransparentProxyNetworkSettings) Autorelease() NETransparentProxyNetworkSettings {
	rv := objc.Send[NETransparentProxyNetworkSettings](n_.ID, objc.Sel("autorelease"))
	return rv
}

// NewNETransparentProxyNetworkSettings creates a new NETransparentProxyNetworkSettings instance.
func NewNETransparentProxyNetworkSettings() NETransparentProxyNetworkSettings {
	return getNETransparentProxyNetworkSettingsClass().New()
}


// An array of rules that collectively specify what traffic to not route through the transparent proxy.
//
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NETransparentProxyNetworkSettings/excludedNetworkRules
func (n_ NETransparentProxyNetworkSettings) ExcludedNetworkRules() []NENetworkRule {
	rv := objc.Send[[]NENetworkRule](n_.ID, objc.Sel("excludedNetworkRules"))
	return rv
}


// SetExcludedNetworkRules sets the value of the excludedNetworkRules property.
// An array of rules that collectively specify what traffic to not route through the transparent proxy.

//
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NETransparentProxyNetworkSettings/excludedNetworkRules
func (n_ NETransparentProxyNetworkSettings) SetExcludedNetworkRules(value []NENetworkRule) {
	// Convert Go slice to NSArray
	var nsArray objc.ID
	if len(value) > 0 {
		nsArray = objc.ID(objc.GetClass("NSMutableArray")).Send(objc.Sel("arrayWithCapacity:"), len(value))
		for _, item := range value {
			nsArray.Send(objc.Sel("addObject:"), item)
		}
	} else {
		nsArray = objc.ID(objc.GetClass("NSArray")).Send(objc.Sel("array"))
	}
	objc.Send[objc.ID](n_.ID, objc.Sel("setExcludedNetworkRules:"), nsArray)
}
// An array of rules that collectively specify what traffic to route through the transparent proxy.
//
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NETransparentProxyNetworkSettings/includedNetworkRules
func (n_ NETransparentProxyNetworkSettings) IncludedNetworkRules() []NENetworkRule {
	rv := objc.Send[[]NENetworkRule](n_.ID, objc.Sel("includedNetworkRules"))
	return rv
}


// SetIncludedNetworkRules sets the value of the includedNetworkRules property.
// An array of rules that collectively specify what traffic to route through the transparent proxy.

//
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NETransparentProxyNetworkSettings/includedNetworkRules
func (n_ NETransparentProxyNetworkSettings) SetIncludedNetworkRules(value []NENetworkRule) {
	// Convert Go slice to NSArray
	var nsArray objc.ID
	if len(value) > 0 {
		nsArray = objc.ID(objc.GetClass("NSMutableArray")).Send(objc.Sel("arrayWithCapacity:"), len(value))
		for _, item := range value {
			nsArray.Send(objc.Sel("addObject:"), item)
		}
	} else {
		nsArray = objc.ID(objc.GetClass("NSArray")).Send(objc.Sel("array"))
	}
	objc.Send[objc.ID](n_.ID, objc.Sel("setIncludedNetworkRules:"), nsArray)
}


