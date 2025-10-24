// Code generated from Apple documentation for NetworkExtension. DO NOT EDIT.

package networkextension

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class NEEthernetTunnelNetworkSettings */


/* debug [class_header]: Header for NEEthernetTunnelNetworkSettings */
// The class instance for the [NEEthernetTunnelNetworkSettings] class.
var (
	NEEthernetTunnelNetworkSettingsClass     _NEEthernetTunnelNetworkSettingsClass
	NEEthernetTunnelNetworkSettingsClassOnce sync.Once
)

func getNEEthernetTunnelNetworkSettingsClass() _NEEthernetTunnelNetworkSettingsClass {
	NEEthernetTunnelNetworkSettingsClassOnce.Do(func() {
		NEEthernetTunnelNetworkSettingsClass = _NEEthernetTunnelNetworkSettingsClass{objc.GetClass("NEEthernetTunnelNetworkSettings")}
	})
	return NEEthernetTunnelNetworkSettingsClass
}

type _NEEthernetTunnelNetworkSettingsClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for NEEthernetTunnelNetworkSettings */
// An interface definition for the [NEEthernetTunnelNetworkSettings] class.
type INEEthernetTunnelNetworkSettings interface {
	INEPacketTunnelNetworkSettings
	
/* debug [class_interface_properties]: Properties for NEEthernetTunnelNetworkSettings */
	// properties:
	EthernetAddress() objc.IObject /* cross-framework: NSString */
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for NEEthernetTunnelNetworkSettings */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for NEEthernetTunnelNetworkSettings */
// Alloc allocates a new instance without initialization.
func (nc _NEEthernetTunnelNetworkSettingsClass) Alloc() NEEthernetTunnelNetworkSettings {
	rv := objc.Send[NEEthernetTunnelNetworkSettings](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for NEEthernetTunnelNetworkSettings */
// The network settings for an ethernet-based VPN tunnel.
//
// You use this type with instances to communicate the desired network settings for the packet tunnel to the framework. The framework takes care of applying the contained settings to the system. Instances of this class are thread-safe.


// The network settings for an ethernet-based VPN tunnel.
//
// [Full Topic]
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for NEEthernetTunnelNetworkSettings */

// Creates a settings object with a given tunnel remote address and MAC address.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEEthernetTunnelNetworkSettings/init(tunnelRemoteAddress:ethernetAddress:mtu:)
func NewNEEthernetTunnelNetworkSettingsWithTunnelRemoteAddressEthernetAddressMtu(address objc.IObject /* cross-framework: NSString */, ethernetAddress objc.IObject /* cross-framework: NSString */, mtu int) NEEthernetTunnelNetworkSettings {
	instance := getNEEthernetTunnelNetworkSettingsClass().Alloc()
	rv := objc.Send[NEEthernetTunnelNetworkSettings](instance.ID, objc.Sel("initWithTunnelRemoteAddress:ethernetAddress:mtu:"), address, ethernetAddress, mtu)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewNEEthernetTunnelNetworkSettingsWithTunnelRemoteAddressEthernetAddressMtu */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for NEEthernetTunnelNetworkSettings */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for NEEthernetTunnelNetworkSettings */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for NEEthernetTunnelNetworkSettings */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for NEEthernetTunnelNetworkSettings */

// The ethernet address of the tunnel interface, as a string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEEthernetTunnelNetworkSettings/ethernetAddress
func (n_ NEEthernetTunnelNetworkSettings) EthernetAddress() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](n_.ID, objc.Sel("ethernetAddress"))
	return rv
}/* debug [instance_properties/getter]: ethernetAddress */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class NEEthernetTunnelNetworkSettings */


