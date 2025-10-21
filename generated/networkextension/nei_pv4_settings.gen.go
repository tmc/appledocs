// Code generated from Apple documentation for NetworkExtension. DO NOT EDIT.

package networkextension

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

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

// An interface definition for the [NEIPv4Settings] class.
type INEIPv4Settings interface {
	objectivec.IObject
}

// The IPv4 settings of an IP layer network tunnel.
//
// To specify the IPv4 settings of a packet tunnel, set its . property to an instance of this class.
//
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

// Alloc allocates a new instance without initialization.
func (nc _NEIPv4SettingsClass) Alloc() NEIPv4Settings {
	rv := objc.Send[NEIPv4Settings](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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


// The IPv4 network traffic that the system routes to the primary physical interface, not the TUN interface.
//
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEIPv4Settings/excludedRoutes
func (n_ NEIPv4Settings) ExcludedRoutes() []NEIPv4Route {
	rv := objc.Send[[]NEIPv4Route](n_.ID, objc.Sel("excludedRoutes"))
	return rv
}


// SetExcludedRoutes sets the value of the excludedRoutes property.
// The IPv4 network traffic that the system routes to the primary physical interface, not the TUN interface.

//
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEIPv4Settings/excludedRoutes
func (n_ NEIPv4Settings) SetExcludedRoutes(value []NEIPv4Route) {
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
	objc.Send[objc.ID](n_.ID, objc.Sel("setExcludedRoutes:"), nsArray)
}

// The IPv4 network traffic that the system routes to the TUN interface.
//
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEIPv4Settings/includedRoutes
func (n_ NEIPv4Settings) IncludedRoutes() []NEIPv4Route {
	rv := objc.Send[[]NEIPv4Route](n_.ID, objc.Sel("includedRoutes"))
	return rv
}


// SetIncludedRoutes sets the value of the includedRoutes property.
// The IPv4 network traffic that the system routes to the TUN interface.

//
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEIPv4Settings/includedRoutes
func (n_ NEIPv4Settings) SetIncludedRoutes(value []NEIPv4Route) {
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
	objc.Send[objc.ID](n_.ID, objc.Sel("setIncludedRoutes:"), nsArray)
}

// The IPv4 addresses to assign to the TUN interface.
//
// [Full Topic]: https://developer.apple.com/documentation/networkextension/neipv4settings/addresses
func (n_ NEIPv4Settings) Addresses() appkit.string {
	rv := objc.Send[appkit.string](n_.ID, objc.Sel("addresses"))
	return rv
}


// SetAddresses sets the value of the addresses property.
// The IPv4 addresses to assign to the TUN interface.

//
// [Full Topic]: https://developer.apple.com/documentation/networkextension/neipv4settings/addresses
func (n_ NEIPv4Settings) SetAddresses(value appkit.string) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setAddresses:"), value)
}

// The address of the next-hop gateway router represented as a dotted decimal string.
//
// [Full Topic]: https://developer.apple.com/documentation/networkextension/neipv4settings/router
func (n_ NEIPv4Settings) Router() appkit.string {
	rv := objc.Send[appkit.string](n_.ID, objc.Sel("router"))
	return rv
}


// SetRouter sets the value of the router property.
// The address of the next-hop gateway router represented as a dotted decimal string.

//
// [Full Topic]: https://developer.apple.com/documentation/networkextension/neipv4settings/router
func (n_ NEIPv4Settings) SetRouter(value appkit.string) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setRouter:"), value)
}

// The IPv4 network masks to assign to the TUN interface.
//
// [Full Topic]: https://developer.apple.com/documentation/networkextension/neipv4settings/subnetmasks
func (n_ NEIPv4Settings) SubnetMasks() appkit.string {
	rv := objc.Send[appkit.string](n_.ID, objc.Sel("subnetMasks"))
	return rv
}


// SetSubnetMasks sets the value of the subnetMasks property.
// The IPv4 network masks to assign to the TUN interface.

//
// [Full Topic]: https://developer.apple.com/documentation/networkextension/neipv4settings/subnetmasks
func (n_ NEIPv4Settings) SetSubnetMasks(value appkit.string) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setSubnetMasks:"), value)
}

// The tunnel IP version 4 settings.
//
// [Full Topic]: https://developer.apple.com/documentation/networkextension/nepackettunnelnetworksettings/ipv4settings
func (n_ NEIPv4Settings) Ipv4Settings() NEIPv4Settings {
	rv := objc.Send[NEIPv4Settings](n_.ID, objc.Sel("ipv4Settings"))
	return rv
}


// SetIpv4Settings sets the value of the ipv4Settings property.
// The tunnel IP version 4 settings.

//
// [Full Topic]: https://developer.apple.com/documentation/networkextension/nepackettunnelnetworksettings/ipv4settings
func (n_ NEIPv4Settings) SetIpv4Settings(value INEIPv4Settings) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setIpv4Settings:"), value)
}

// The tunnel IP version 6 settings.
//
// [Full Topic]: https://developer.apple.com/documentation/networkextension/nepackettunnelnetworksettings/ipv6settings
func (n_ NEIPv4Settings) Ipv6Settings() NEIPv6Settings {
	rv := objc.Send[NEIPv6Settings](n_.ID, objc.Sel("ipv6Settings"))
	return rv
}


// SetIpv6Settings sets the value of the ipv6Settings property.
// The tunnel IP version 6 settings.

//
// [Full Topic]: https://developer.apple.com/documentation/networkextension/nepackettunnelnetworksettings/ipv6settings
func (n_ NEIPv4Settings) SetIpv6Settings(value INEIPv6Settings) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setIpv6Settings:"), value)
}

// The size of the maximum trasnmission unit, in bytes.
//
// [Full Topic]: https://developer.apple.com/documentation/networkextension/nepackettunnelnetworksettings/mtu
func (n_ NEIPv4Settings) Mtu() foundation.Number {
	rv := objc.Send[foundation.Number](n_.ID, objc.Sel("mtu"))
	return rv
}


// SetMtu sets the value of the mtu property.
// The size of the maximum trasnmission unit, in bytes.

//
// [Full Topic]: https://developer.apple.com/documentation/networkextension/nepackettunnelnetworksettings/mtu
func (n_ NEIPv4Settings) SetMtu(value foundation.INumber) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setMtu:"), value)
}

// The number of bytes added to each tunneled packet for storing tunneling protocol headers.
//
// [Full Topic]: https://developer.apple.com/documentation/networkextension/nepackettunnelnetworksettings/tunneloverheadbytes
func (n_ NEIPv4Settings) TunnelOverheadBytes() foundation.Number {
	rv := objc.Send[foundation.Number](n_.ID, objc.Sel("tunnelOverheadBytes"))
	return rv
}


// SetTunnelOverheadBytes sets the value of the tunnelOverheadBytes property.
// The number of bytes added to each tunneled packet for storing tunneling protocol headers.

//
// [Full Topic]: https://developer.apple.com/documentation/networkextension/nepackettunnelnetworksettings/tunneloverheadbytes
func (n_ NEIPv4Settings) SetTunnelOverheadBytes(value foundation.INumber) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setTunnelOverheadBytes:"), value)
}



