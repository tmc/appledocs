// Code generated from Apple documentation for NetworkExtension. DO NOT EDIT.

package networkextension

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class NETransparentProxyNetworkSettings */


/* debug [class_header]: Header for NETransparentProxyNetworkSettings */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for NETransparentProxyNetworkSettings */
// An interface definition for the [NETransparentProxyNetworkSettings] class.
type INETransparentProxyNetworkSettings interface {
	INETunnelNetworkSettings
	
/* debug [class_interface_properties]: Properties for NETransparentProxyNetworkSettings */
	// properties:
	ExcludedNetworkRules() []NENetworkRule
	SetExcludedNetworkRules(value []NENetworkRule)
	IncludedNetworkRules() []NENetworkRule
	SetIncludedNetworkRules(value []NENetworkRule)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for NETransparentProxyNetworkSettings */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for NETransparentProxyNetworkSettings */
// Alloc allocates a new instance without initialization.
func (nc _NETransparentProxyNetworkSettingsClass) Alloc() NETransparentProxyNetworkSettings {
	rv := objc.Send[NETransparentProxyNetworkSettings](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for NETransparentProxyNetworkSettings */
// A specification of what traffic to route through a transparent proxy.
//
// A proxy network settings object contains two properties: an array of rules to include traffic ( ) and an array of rules to exclude traffic ( ). The exclusion rules take prirority. Therefore, if a given flow matches any of the , evaluation ends and the flow doesn’t route to the proxy. If there’s no match, then evaluation continues and attempts to match the flow against the .


// A specification of what traffic to route through a transparent proxy.
//
// [Full Topic]
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for NETransparentProxyNetworkSettings *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for NETransparentProxyNetworkSettings */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for NETransparentProxyNetworkSettings */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for NETransparentProxyNetworkSettings */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for NETransparentProxyNetworkSettings */

// An array of rules that collectively specify what traffic to not route through the transparent proxy.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NETransparentProxyNetworkSettings/excludedNetworkRules
func (n_ NETransparentProxyNetworkSettings) ExcludedNetworkRules() []NENetworkRule {
	rv := objc.Send[[]NENetworkRule](n_.ID, objc.Sel("excludedNetworkRules"))
	return rv
}/* debug [instance_properties/getter]: excludedNetworkRules */


// An array of rules that collectively specify what traffic to not route through the transparent proxy.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NETransparentProxyNetworkSettings/excludedNetworkRules
func (n_ NETransparentProxyNetworkSettings) SetExcludedNetworkRules(value []NENetworkRule) {
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
}/* debug [instance_properties/setter]: excludedNetworkRules */


// An array of rules that collectively specify what traffic to route through the transparent proxy.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NETransparentProxyNetworkSettings/includedNetworkRules
func (n_ NETransparentProxyNetworkSettings) IncludedNetworkRules() []NENetworkRule {
	rv := objc.Send[[]NENetworkRule](n_.ID, objc.Sel("includedNetworkRules"))
	return rv
}/* debug [instance_properties/getter]: includedNetworkRules */


// An array of rules that collectively specify what traffic to route through the transparent proxy.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NETransparentProxyNetworkSettings/includedNetworkRules
func (n_ NETransparentProxyNetworkSettings) SetIncludedNetworkRules(value []NENetworkRule) {
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
}/* debug [instance_properties/setter]: includedNetworkRules */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class NETransparentProxyNetworkSettings */



