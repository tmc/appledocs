// Code generated from Apple documentation for NetworkExtension. DO NOT EDIT.

package networkextension

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

/* debug [class.gen.go]: Generating class NEPacketTunnelNetworkSettings */


/* debug [class_header]: Header for NEPacketTunnelNetworkSettings */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for NEPacketTunnelNetworkSettings */
// An interface definition for the [NEPacketTunnelNetworkSettings] class.
type INEPacketTunnelNetworkSettings interface {
	INETunnelNetworkSettings
	
/* debug [class_interface_properties]: Properties for NEPacketTunnelNetworkSettings */
	// properties:
	IPv4Settings() INEIPv4Settings
	SetIPv4Settings(value INEIPv4Settings)
	IPv6Settings() INEIPv6Settings
	SetIPv6Settings(value INEIPv6Settings)
	MTU() objc.IObject /* cross-framework: NSNumber */
	SetMTU(value objc.IObject /* cross-framework: NSNumber */)
	TunnelOverheadBytes() objc.IObject /* cross-framework: NSNumber */
	SetTunnelOverheadBytes(value objc.IObject /* cross-framework: NSNumber */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for NEPacketTunnelNetworkSettings */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for NEPacketTunnelNetworkSettings */
// Alloc allocates a new instance without initialization.
func (nc _NEPacketTunnelNetworkSettingsClass) Alloc() NEPacketTunnelNetworkSettings {
	rv := objc.Send[NEPacketTunnelNetworkSettings](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for NEPacketTunnelNetworkSettings */
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for NEPacketTunnelNetworkSettings *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for NEPacketTunnelNetworkSettings */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for NEPacketTunnelNetworkSettings */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for NEPacketTunnelNetworkSettings */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for NEPacketTunnelNetworkSettings */

// The tunnel IP version 4 settings.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEPacketTunnelNetworkSettings/ipv4Settings
func (n_ NEPacketTunnelNetworkSettings) IPv4Settings() INEIPv4Settings {
	rv := objc.Send[NEIPv4Settings](n_.ID, objc.Sel("IPv4Settings"))
	return rv
}/* debug [instance_properties/getter]: IPv4Settings */


// The tunnel IP version 4 settings.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEPacketTunnelNetworkSettings/ipv4Settings
func (n_ NEPacketTunnelNetworkSettings) SetIPv4Settings(value INEIPv4Settings) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setIPv4Settings:"), value)
}/* debug [instance_properties/setter]: IPv4Settings */


// The tunnel IP version 6 settings.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEPacketTunnelNetworkSettings/ipv6Settings
func (n_ NEPacketTunnelNetworkSettings) IPv6Settings() INEIPv6Settings {
	rv := objc.Send[NEIPv6Settings](n_.ID, objc.Sel("IPv6Settings"))
	return rv
}/* debug [instance_properties/getter]: IPv6Settings */


// The tunnel IP version 6 settings.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEPacketTunnelNetworkSettings/ipv6Settings
func (n_ NEPacketTunnelNetworkSettings) SetIPv6Settings(value INEIPv6Settings) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setIPv6Settings:"), value)
}/* debug [instance_properties/setter]: IPv6Settings */


// The size of the maximum trasnmission unit, in bytes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEPacketTunnelNetworkSettings/mtu
func (n_ NEPacketTunnelNetworkSettings) MTU() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](n_.ID, objc.Sel("MTU"))
	return rv
}/* debug [instance_properties/getter]: MTU */


// The size of the maximum trasnmission unit, in bytes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEPacketTunnelNetworkSettings/mtu
func (n_ NEPacketTunnelNetworkSettings) SetMTU(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setMTU:"), value)
}/* debug [instance_properties/setter]: MTU */


// The number of bytes added to each tunneled packet for storing tunneling protocol headers.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEPacketTunnelNetworkSettings/tunnelOverheadBytes
func (n_ NEPacketTunnelNetworkSettings) TunnelOverheadBytes() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](n_.ID, objc.Sel("tunnelOverheadBytes"))
	return rv
}/* debug [instance_properties/getter]: tunnelOverheadBytes */


// The number of bytes added to each tunneled packet for storing tunneling protocol headers.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEPacketTunnelNetworkSettings/tunnelOverheadBytes
func (n_ NEPacketTunnelNetworkSettings) SetTunnelOverheadBytes(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setTunnelOverheadBytes:"), value)
}/* debug [instance_properties/setter]: tunnelOverheadBytes */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class NEPacketTunnelNetworkSettings */



