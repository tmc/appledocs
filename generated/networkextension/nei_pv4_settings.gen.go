// Code generated from Apple documentation for NetworkExtension. DO NOT EDIT.

package networkextension

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class NEIPv4Settings */


/* debug [class_header]: Header for NEIPv4Settings */
// The class instance for the [NEIPv4Settings] class.
var (
	NEIPv4SettingsClass     _NEIPv4SettingsClass
	NEIPv4SettingsClassOnce sync.Once
)

func getNEIPv4SettingsClass() _NEIPv4SettingsClass {
	NEIPv4SettingsClassOnce.Do(func() {
		NEIPv4SettingsClass = _NEIPv4SettingsClass{objc.GetClass("NEIPv4Settings")}
	})
	return NEIPv4SettingsClass
}

type _NEIPv4SettingsClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for NEIPv4Settings */
// An interface definition for the [NEIPv4Settings] class.
type INEIPv4Settings interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for NEIPv4Settings */
	// properties:
	Addresses() []string
	ExcludedRoutes() []NEIPv4Route
	SetExcludedRoutes(value []NEIPv4Route)
	IncludedRoutes() []NEIPv4Route
	SetIncludedRoutes(value []NEIPv4Route)
	Router() objc.IObject /* cross-framework: NSString */
	SetRouter(value objc.IObject /* cross-framework: NSString */)
	SubnetMasks() []string
	Ipv4Settings() INEIPv4Settings
	SetIpv4Settings(value INEIPv4Settings)
	Ipv6Settings() INEIPv6Settings
	SetIpv6Settings(value INEIPv6Settings)
	Mtu() objc.IObject /* cross-framework: NSNumber */
	SetMtu(value objc.IObject /* cross-framework: NSNumber */)
	TunnelOverheadBytes() objc.IObject /* cross-framework: NSNumber */
	SetTunnelOverheadBytes(value objc.IObject /* cross-framework: NSNumber */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for NEIPv4Settings */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for NEIPv4Settings */
// Alloc allocates a new instance without initialization.
func (nc _NEIPv4SettingsClass) Alloc() NEIPv4Settings {
	rv := objc.Send[NEIPv4Settings](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (nc _NEIPv4SettingsClass) New() NEIPv4Settings {
	rv := objc.Send[NEIPv4Settings](objc.ID(nc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (n_ NEIPv4Settings) Init() NEIPv4Settings {
	rv := objc.Send[NEIPv4Settings](n_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (n_ NEIPv4Settings) Autorelease() NEIPv4Settings {
	rv := objc.Send[NEIPv4Settings](n_.ID, objc.Sel("autorelease"))
	return rv
}

// NewNEIPv4Settings creates a new NEIPv4Settings instance.
func NewNEIPv4Settings() NEIPv4Settings {
	return getNEIPv4SettingsClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for NEIPv4Settings */
// The IPv4 settings of an IP layer network tunnel.
//
// To specify the IPv4 settings of a packet tunnel, set its . property to an instance of this class.


// The IPv4 settings of an IP layer network tunnel.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEIPv4Settings
type NEIPv4Settings struct {
	objectivec.Object
}

// NEIPv4SettingsFrom constructs a [NEIPv4Settings] from an unsafe.Pointer.
//
// The IPv4 settings of an IP layer network tunnel.
func NEIPv4SettingsFrom(ptr unsafe.Pointer) NEIPv4Settings {
	return NEIPv4Settings{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for NEIPv4Settings */

// Initializes an IPv4 settings object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEIPv4Settings/init(addresses:subnetMasks:)
func NewNEIPv4SettingsWithAddressesSubnetMasks(addresses []string, subnetMasks []string) NEIPv4Settings {
	instance := getNEIPv4SettingsClass().Alloc()
	rv := objc.Send[NEIPv4Settings](instance.ID, objc.Sel("initWithAddresses:subnetMasks:"), addresses, subnetMasks)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewNEIPv4SettingsWithAddressesSubnetMasks */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for NEIPv4Settings */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEIPv4Settings/settingsWithAutomaticAddressing
func (nc _NEIPv4SettingsClass) SettingsWithAutomaticAddressing() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(nc.class), objc.Sel("settingsWithAutomaticAddressing"))
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=SettingsWithAutomaticAddressing) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for NEIPv4Settings */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for NEIPv4Settings */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for NEIPv4Settings */

// The IPv4 addresses to assign to the TUN interface.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEIPv4Settings/addresses
func (n_ NEIPv4Settings) Addresses() []string {
	rv := objc.Send[[]string](n_.ID, objc.Sel("addresses"))
	return rv
}/* debug [instance_properties/getter]: addresses */


// The IPv4 network traffic that the system routes to the primary physical interface, not the TUN interface.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEIPv4Settings/excludedRoutes
func (n_ NEIPv4Settings) ExcludedRoutes() []NEIPv4Route {
	rv := objc.Send[[]NEIPv4Route](n_.ID, objc.Sel("excludedRoutes"))
	return rv
}/* debug [instance_properties/getter]: excludedRoutes */


// The IPv4 network traffic that the system routes to the primary physical interface, not the TUN interface.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEIPv4Settings/excludedRoutes
func (n_ NEIPv4Settings) SetExcludedRoutes(value []NEIPv4Route) {
	var nsArray objc.ID
	if len(value) > 0 {
		nsArray = objc.ID(objc.GetClass("NSMutableArray")).Send(objc.Sel("arrayWithCapacity:"), len(value))
		for _, item := range value {
			nsArray.Send(objc.Sel("addObject:"), item)
		}
	} else {
		nsArray = objc.ID(objc.GetClass("NSArray")).Send(objc.Sel("array"))
	}
	objc.Send[objc.ID](n_.ID, objc.Sel("setExcludedRoutes:"), nsArray)
}/* debug [instance_properties/setter]: excludedRoutes */


// The IPv4 network traffic that the system routes to the TUN interface.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEIPv4Settings/includedRoutes
func (n_ NEIPv4Settings) IncludedRoutes() []NEIPv4Route {
	rv := objc.Send[[]NEIPv4Route](n_.ID, objc.Sel("includedRoutes"))
	return rv
}/* debug [instance_properties/getter]: includedRoutes */


// The IPv4 network traffic that the system routes to the TUN interface.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEIPv4Settings/includedRoutes
func (n_ NEIPv4Settings) SetIncludedRoutes(value []NEIPv4Route) {
	var nsArray objc.ID
	if len(value) > 0 {
		nsArray = objc.ID(objc.GetClass("NSMutableArray")).Send(objc.Sel("arrayWithCapacity:"), len(value))
		for _, item := range value {
			nsArray.Send(objc.Sel("addObject:"), item)
		}
	} else {
		nsArray = objc.ID(objc.GetClass("NSArray")).Send(objc.Sel("array"))
	}
	objc.Send[objc.ID](n_.ID, objc.Sel("setIncludedRoutes:"), nsArray)
}/* debug [instance_properties/setter]: includedRoutes */


// The address of the next-hop gateway router represented as a dotted decimal string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEIPv4Settings/router
func (n_ NEIPv4Settings) Router() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](n_.ID, objc.Sel("router"))
	return rv
}/* debug [instance_properties/getter]: router */


// The address of the next-hop gateway router represented as a dotted decimal string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEIPv4Settings/router
func (n_ NEIPv4Settings) SetRouter(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setRouter:"), value)
}/* debug [instance_properties/setter]: router */


// The IPv4 network masks to assign to the TUN interface.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEIPv4Settings/subnetMasks
func (n_ NEIPv4Settings) SubnetMasks() []string {
	rv := objc.Send[[]string](n_.ID, objc.Sel("subnetMasks"))
	return rv
}/* debug [instance_properties/getter]: subnetMasks */


// The tunnel IP version 4 settings.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/networkextension/nepackettunnelnetworksettings/ipv4settings
func (n_ NEIPv4Settings) Ipv4Settings() INEIPv4Settings {
	rv := objc.Send[NEIPv4Settings](n_.ID, objc.Sel("ipv4Settings"))
	return rv
}/* debug [instance_properties/getter]: ipv4Settings */


// The tunnel IP version 4 settings.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/networkextension/nepackettunnelnetworksettings/ipv4settings
func (n_ NEIPv4Settings) SetIpv4Settings(value INEIPv4Settings) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setIpv4Settings:"), value)
}/* debug [instance_properties/setter]: ipv4Settings */


// The tunnel IP version 6 settings.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/networkextension/nepackettunnelnetworksettings/ipv6settings
func (n_ NEIPv4Settings) Ipv6Settings() INEIPv6Settings {
	rv := objc.Send[NEIPv6Settings](n_.ID, objc.Sel("ipv6Settings"))
	return rv
}/* debug [instance_properties/getter]: ipv6Settings */


// The tunnel IP version 6 settings.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/networkextension/nepackettunnelnetworksettings/ipv6settings
func (n_ NEIPv4Settings) SetIpv6Settings(value INEIPv6Settings) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setIpv6Settings:"), value)
}/* debug [instance_properties/setter]: ipv6Settings */


// The size of the maximum trasnmission unit, in bytes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/networkextension/nepackettunnelnetworksettings/mtu
func (n_ NEIPv4Settings) Mtu() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](n_.ID, objc.Sel("mtu"))
	return rv
}/* debug [instance_properties/getter]: mtu */


// The size of the maximum trasnmission unit, in bytes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/networkextension/nepackettunnelnetworksettings/mtu
func (n_ NEIPv4Settings) SetMtu(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setMtu:"), value)
}/* debug [instance_properties/setter]: mtu */


// The number of bytes added to each tunneled packet for storing tunneling protocol headers.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/networkextension/nepackettunnelnetworksettings/tunneloverheadbytes
func (n_ NEIPv4Settings) TunnelOverheadBytes() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](n_.ID, objc.Sel("tunnelOverheadBytes"))
	return rv
}/* debug [instance_properties/getter]: tunnelOverheadBytes */


// The number of bytes added to each tunneled packet for storing tunneling protocol headers.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/networkextension/nepackettunnelnetworksettings/tunneloverheadbytes
func (n_ NEIPv4Settings) SetTunnelOverheadBytes(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setTunnelOverheadBytes:"), value)
}/* debug [instance_properties/setter]: tunnelOverheadBytes */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class NEIPv4Settings */


