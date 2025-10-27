// Code generated from Apple documentation for NetworkExtension. DO NOT EDIT.

package networkextension

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
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
	

	// properties:
	Addresses() []string
	ExcludedRoutes() []NEIPv4Route
	SetExcludedRoutes(value []NEIPv4Route)
	IncludedRoutes() []NEIPv4Route
	SetIncludedRoutes(value []NEIPv4Route)
	Router() foundation.foundation.INSString
	SetRouter(value foundation.foundation.INSString)
	SubnetMasks() []string
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






// Initializes an IPv4 settings object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEIPv4Settings/init(addresses:subnetMasks:)
func NewNEIPv4SettingsWithAddressesSubnetMasks(addresses []string, subnetMasks []string) NEIPv4Settings {
	instance := getNEIPv4SettingsClass().Alloc()
	rv := objc.Send[NEIPv4Settings](instance.ID, objc.Sel("initWithAddresses:subnetMasks:"), addresses, subnetMasks)
	rv.Autorelease()
	return rv
}







// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEIPv4Settings/settingsWithAutomaticAddressing
func (nc _NEIPv4SettingsClass) SettingsWithAutomaticAddressing() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(nc.class), objc.Sel("settingsWithAutomaticAddressing"))
	return rv
}

















// The IPv4 addresses to assign to the TUN interface.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEIPv4Settings/addresses
func (n_ NEIPv4Settings) Addresses() []string {
	rv := objc.Send[[]string](n_.ID, objc.Sel("addresses"))
	return rv
}


// The IPv4 network traffic that the system routes to the primary physical interface, not the TUN interface.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEIPv4Settings/excludedRoutes
func (n_ NEIPv4Settings) ExcludedRoutes() []NEIPv4Route {
	rv := objc.Send[[]NEIPv4Route](n_.ID, objc.Sel("excludedRoutes"))
	return rv
}


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
}


// The IPv4 network traffic that the system routes to the TUN interface.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEIPv4Settings/includedRoutes
func (n_ NEIPv4Settings) IncludedRoutes() []NEIPv4Route {
	rv := objc.Send[[]NEIPv4Route](n_.ID, objc.Sel("includedRoutes"))
	return rv
}


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
}


// The address of the next-hop gateway router represented as a dotted decimal string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEIPv4Settings/router
func (n_ NEIPv4Settings) Router() foundation.foundation.INSString {
	rv := objc.Send[foundation.NSString](n_.ID, objc.Sel("router"))
	return rv
}


// The address of the next-hop gateway router represented as a dotted decimal string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEIPv4Settings/router
func (n_ NEIPv4Settings) SetRouter(value foundation.foundation.INSString) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setRouter:"), value)
}


// The IPv4 network masks to assign to the TUN interface.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEIPv4Settings/subnetMasks
func (n_ NEIPv4Settings) SubnetMasks() []string {
	rv := objc.Send[[]string](n_.ID, objc.Sel("subnetMasks"))
	return rv
}


// The tunnel IP version 4 settings.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/networkextension/nepackettunnelnetworksettings/ipv4settings
func (n_ NEIPv4Settings) Ipv4Settings() INEIPv4Settings {
	rv := objc.Send[NEIPv4Settings](n_.ID, objc.Sel("ipv4Settings"))
	return rv
}


// The tunnel IP version 4 settings.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/networkextension/nepackettunnelnetworksettings/ipv4settings
func (n_ NEIPv4Settings) SetIpv4Settings(value INEIPv4Settings) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setIpv4Settings:"), value)
}


// The tunnel IP version 6 settings.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/networkextension/nepackettunnelnetworksettings/ipv6settings
func (n_ NEIPv4Settings) Ipv6Settings() INEIPv6Settings {
	rv := objc.Send[NEIPv6Settings](n_.ID, objc.Sel("ipv6Settings"))
	return rv
}


// The tunnel IP version 6 settings.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/networkextension/nepackettunnelnetworksettings/ipv6settings
func (n_ NEIPv4Settings) SetIpv6Settings(value INEIPv6Settings) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setIpv6Settings:"), value)
}


// The size of the maximum trasnmission unit, in bytes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/networkextension/nepackettunnelnetworksettings/mtu
func (n_ NEIPv4Settings) Mtu() foundation.foundation.INSNumber {
	rv := objc.Send[foundation.NSNumber](n_.ID, objc.Sel("mtu"))
	return rv
}


// The size of the maximum trasnmission unit, in bytes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/networkextension/nepackettunnelnetworksettings/mtu
func (n_ NEIPv4Settings) SetMtu(value foundation.foundation.INSNumber) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setMtu:"), value)
}


// The number of bytes added to each tunneled packet for storing tunneling protocol headers.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/networkextension/nepackettunnelnetworksettings/tunneloverheadbytes
func (n_ NEIPv4Settings) TunnelOverheadBytes() foundation.foundation.INSNumber {
	rv := objc.Send[foundation.NSNumber](n_.ID, objc.Sel("tunnelOverheadBytes"))
	return rv
}


// The number of bytes added to each tunneled packet for storing tunneling protocol headers.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/networkextension/nepackettunnelnetworksettings/tunneloverheadbytes
func (n_ NEIPv4Settings) SetTunnelOverheadBytes(value foundation.foundation.INSNumber) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setTunnelOverheadBytes:"), value)
}







