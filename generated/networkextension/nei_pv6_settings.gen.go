// Code generated from Apple documentation for NetworkExtension. DO NOT EDIT.

package networkextension

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)





// The class instance for the [NEIPv6Settings] class.
var (
	NEIPv6SettingsClass     _NEIPv6SettingsClass
	NEIPv6SettingsClassOnce sync.Once
)

func getNEIPv6SettingsClass() _NEIPv6SettingsClass {
	NEIPv6SettingsClassOnce.Do(func() {
		NEIPv6SettingsClass = _NEIPv6SettingsClass{objc.GetClass("NEIPv6Settings")}
	})
	return NEIPv6SettingsClass
}

type _NEIPv6SettingsClass struct {
	class objc.Class
}





// An interface definition for the [NEIPv6Settings] class.
type INEIPv6Settings interface {
	objectivec.IObject
	

	// properties:
	Addresses() []string
	ExcludedRoutes() []NEIPv6Route
	SetExcludedRoutes(value []NEIPv6Route)
	IncludedRoutes() []NEIPv6Route
	SetIncludedRoutes(value []NEIPv6Route)
	NetworkPrefixLengths() []foundation.Number
	Ipv4Settings() INEIPv4Settings
	SetIpv4Settings(value INEIPv4Settings)
	Ipv6Settings() INEIPv6Settings
	SetIpv6Settings(value INEIPv6Settings)
	Mtu() foundation.foundation.INSNumber
	SetMtu(value foundation.foundation.INSNumber)
	TunnelOverheadBytes() foundation.foundation.INSNumber
	SetTunnelOverheadBytes(value foundation.foundation.INSNumber)


	

	// methods:


}





// Alloc allocates a new instance without initialization.
func (nc _NEIPv6SettingsClass) Alloc() NEIPv6Settings {
	rv := objc.Send[NEIPv6Settings](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (nc _NEIPv6SettingsClass) New() NEIPv6Settings {
	rv := objc.Send[NEIPv6Settings](objc.ID(nc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (n_ NEIPv6Settings) Init() NEIPv6Settings {
	rv := objc.Send[NEIPv6Settings](n_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (n_ NEIPv6Settings) Autorelease() NEIPv6Settings {
	rv := objc.Send[NEIPv6Settings](n_.ID, objc.Sel("autorelease"))
	return rv
}

// NewNEIPv6Settings creates a new NEIPv6Settings instance.
func NewNEIPv6Settings() NEIPv6Settings {
	return getNEIPv6SettingsClass().New()
}





// The IPv6 settings of an IP layer network tunnel.
//
// To specify the IPv6 settings of a packet tunnel, set its . property to an instance of this class.


// The IPv6 settings of an IP layer network tunnel.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEIPv6Settings
type NEIPv6Settings struct {
	objectivec.Object
}

// NEIPv6SettingsFrom constructs a [NEIPv6Settings] from an unsafe.Pointer.
//
// The IPv6 settings of an IP layer network tunnel.
func NEIPv6SettingsFrom(ptr unsafe.Pointer) NEIPv6Settings {
	return NEIPv6Settings{objectivec.Object{objc.ID(ptr)}}
}






// Initializes the IPv6 settings object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEIPv6Settings/init(addresses:networkPrefixLengths:)
func NewNEIPv6SettingsWithAddressesNetworkPrefixLengths(addresses []string, networkPrefixLengths []foundation.Number) NEIPv6Settings {
	instance := getNEIPv6SettingsClass().Alloc()
	rv := objc.Send[NEIPv6Settings](instance.ID, objc.Sel("initWithAddresses:networkPrefixLengths:"), addresses, networkPrefixLengths)
	rv.Autorelease()
	return rv
}







// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEIPv6Settings/settingsWithAutomaticAddressing
func (nc _NEIPv6SettingsClass) SettingsWithAutomaticAddressing() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(nc.class), objc.Sel("settingsWithAutomaticAddressing"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEIPv6Settings/settingsWithLinkLocalAddressing
func (nc _NEIPv6SettingsClass) SettingsWithLinkLocalAddressing() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(nc.class), objc.Sel("settingsWithLinkLocalAddressing"))
	return rv
}

















// The IPv6 addresses to assign to the TUN interface.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEIPv6Settings/addresses
func (n_ NEIPv6Settings) Addresses() []string {
	rv := objc.Send[[]string](n_.ID, objc.Sel("addresses"))
	return rv
}


// The IPv6 network traffic that the system routes to the primary physical interface, not the TUN interface.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEIPv6Settings/excludedRoutes
func (n_ NEIPv6Settings) ExcludedRoutes() []NEIPv6Route {
	rv := objc.Send[[]NEIPv6Route](n_.ID, objc.Sel("excludedRoutes"))
	return rv
}


// The IPv6 network traffic that the system routes to the primary physical interface, not the TUN interface.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEIPv6Settings/excludedRoutes
func (n_ NEIPv6Settings) SetExcludedRoutes(value []NEIPv6Route) {
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


// The IPv6 network traffic that the system routes to the TUN interface.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEIPv6Settings/includedRoutes
func (n_ NEIPv6Settings) IncludedRoutes() []NEIPv6Route {
	rv := objc.Send[[]NEIPv6Route](n_.ID, objc.Sel("includedRoutes"))
	return rv
}


// The IPv6 network traffic that the system routes to the TUN interface.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEIPv6Settings/includedRoutes
func (n_ NEIPv6Settings) SetIncludedRoutes(value []NEIPv6Route) {
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


// The IPv6 network prefix lengths to assign to the TUN interface.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEIPv6Settings/networkPrefixLengths
func (n_ NEIPv6Settings) NetworkPrefixLengths() []foundation.Number {
	rv := objc.Send[[]foundation.Number](n_.ID, objc.Sel("networkPrefixLengths"))
	return rv
}


// The tunnel IP version 4 settings.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/networkextension/nepackettunnelnetworksettings/ipv4settings
func (n_ NEIPv6Settings) Ipv4Settings() INEIPv4Settings {
	rv := objc.Send[NEIPv4Settings](n_.ID, objc.Sel("ipv4Settings"))
	return rv
}


// The tunnel IP version 4 settings.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/networkextension/nepackettunnelnetworksettings/ipv4settings
func (n_ NEIPv6Settings) SetIpv4Settings(value INEIPv4Settings) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setIpv4Settings:"), value)
}


// The tunnel IP version 6 settings.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/networkextension/nepackettunnelnetworksettings/ipv6settings
func (n_ NEIPv6Settings) Ipv6Settings() INEIPv6Settings {
	rv := objc.Send[NEIPv6Settings](n_.ID, objc.Sel("ipv6Settings"))
	return rv
}


// The tunnel IP version 6 settings.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/networkextension/nepackettunnelnetworksettings/ipv6settings
func (n_ NEIPv6Settings) SetIpv6Settings(value INEIPv6Settings) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setIpv6Settings:"), value)
}


// The size of the maximum trasnmission unit, in bytes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/networkextension/nepackettunnelnetworksettings/mtu
func (n_ NEIPv6Settings) Mtu() foundation.foundation.INSNumber {
	rv := objc.Send[foundation.NSNumber](n_.ID, objc.Sel("mtu"))
	return rv
}


// The size of the maximum trasnmission unit, in bytes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/networkextension/nepackettunnelnetworksettings/mtu
func (n_ NEIPv6Settings) SetMtu(value foundation.foundation.INSNumber) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setMtu:"), value)
}


// The number of bytes added to each tunneled packet for storing tunneling protocol headers.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/networkextension/nepackettunnelnetworksettings/tunneloverheadbytes
func (n_ NEIPv6Settings) TunnelOverheadBytes() foundation.foundation.INSNumber {
	rv := objc.Send[foundation.NSNumber](n_.ID, objc.Sel("tunnelOverheadBytes"))
	return rv
}


// The number of bytes added to each tunneled packet for storing tunneling protocol headers.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/networkextension/nepackettunnelnetworksettings/tunneloverheadbytes
func (n_ NEIPv6Settings) SetTunnelOverheadBytes(value foundation.foundation.INSNumber) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setTunnelOverheadBytes:"), value)
}







