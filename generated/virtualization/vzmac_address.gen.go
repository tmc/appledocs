// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [VZMACAddress] class.
var (
	VZMACAddressClass     _VZMACAddressClass
	VZMACAddressClassOnce sync.Once
)

func getVZMACAddressClass() _VZMACAddressClass {
	VZMACAddressClassOnce.Do(func() {
		VZMACAddressClass = _VZMACAddressClass{objc.GetClass("VZMACAddress")}
	})
	return VZMACAddressClass
}

type _VZMACAddressClass struct {
	class objc.Class
}

// An interface definition for the [VZMACAddress] class.
type IVZMACAddress interface {
	objectivec.IObject
	EthernetAddress() unsafe.Pointer
	IsBroadcastAddress() bool
	SetIsBroadcastAddress(value bool)
	IsLocallyAdministeredAddress() bool
	SetIsLocallyAdministeredAddress(value bool)
	IsMulticastAddress() bool
	SetIsMulticastAddress(value bool)
	IsUnicastAddress() bool
	SetIsUnicastAddress(value bool)
	IsUniversallyAdministeredAddress() bool
	SetIsUniversallyAdministeredAddress(value bool)
	String() string
	SetString(value string)
}

// The media access control (MAC) address for a network interface in your virtual machine.
//
// A object contains the hardware address of your network interface. Every network device has a unique 48-bit MAC address that the system uses to route network packets to that device. Call the method to get a local MAC address suitable for use with your network interfaces. Alternatively, you can create a object yourself from a string or structure.
//
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZMACAddress
type VZMACAddress struct {
	objectivec.Object
}

// VZMACAddressFrom constructs a [VZMACAddress] from an unsafe.Pointer.
//
// The media access control (MAC) address for a network interface in your virtual machine.
func VZMACAddressFrom(ptr unsafe.Pointer) VZMACAddress {
	return VZMACAddress{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (vc _VZMACAddressClass) Alloc() VZMACAddress {
	rv := objc.Send[VZMACAddress](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (vc _VZMACAddressClass) New() VZMACAddress {
	rv := objc.Send[VZMACAddress](objc.ID(vc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (v_ VZMACAddress) Init() VZMACAddress {
	rv := objc.Send[VZMACAddress](v_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v_ VZMACAddress) Autorelease() VZMACAddress {
	rv := objc.Send[VZMACAddress](v_.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZMACAddress creates a new VZMACAddress instance.
func NewVZMACAddress() VZMACAddress {
	return getVZMACAddressClass().New()
}




// Creates a MAC address object from a specially formatted string.
//
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZMACAddress/init(string:)
func NewVZMACAddressWithString(string_ string) VZMACAddress {
	instance := getVZMACAddressClass().Alloc()
	rv := objc.Send[VZMACAddress](instance.ID, objc.Sel("initWithString:"), objc.String(string_))
	rv.Autorelease()
	return rv
}


// Returns a valid, random, locally administered, unicast MAC address.
//
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZMACAddress/randomLocallyAdministered()
func (vc _VZMACAddressClass) RandomLocallyAdministeredAddress() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(vc.class), objc.Sel("randomLocallyAdministeredAddress"))
	return rv
}

// The MAC address as an Ethernet data structure.
//
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZMACAddress/ethernetAddress
func (v_ VZMACAddress) EthernetAddress() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](v_.ID, objc.Sel("ethernetAddress"))
	return rv
}

// A Boolean value that indicates whether the address is a broadcast address.
//
// [Full Topic]: https://developer.apple.com/documentation/virtualization/vzmacaddress/isbroadcastaddress
func (v_ VZMACAddress) IsBroadcastAddress() bool {
	rv := objc.Send[bool](v_.ID, objc.Sel("isBroadcastAddress"))
	return rv
}


// SetIsBroadcastAddress sets the value of the isBroadcastAddress property.
// A Boolean value that indicates whether the address is a broadcast address.

//
// [Full Topic]: https://developer.apple.com/documentation/virtualization/vzmacaddress/isbroadcastaddress
func (v_ VZMACAddress) SetIsBroadcastAddress(value bool) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setIsBroadcastAddress:"), value)
}

// A Boolean value that indicates whether the address is a locally administered address (LAA).
//
// [Full Topic]: https://developer.apple.com/documentation/virtualization/vzmacaddress/islocallyadministeredaddress
func (v_ VZMACAddress) IsLocallyAdministeredAddress() bool {
	rv := objc.Send[bool](v_.ID, objc.Sel("isLocallyAdministeredAddress"))
	return rv
}


// SetIsLocallyAdministeredAddress sets the value of the isLocallyAdministeredAddress property.
// A Boolean value that indicates whether the address is a locally administered address (LAA).

//
// [Full Topic]: https://developer.apple.com/documentation/virtualization/vzmacaddress/islocallyadministeredaddress
func (v_ VZMACAddress) SetIsLocallyAdministeredAddress(value bool) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setIsLocallyAdministeredAddress:"), value)
}

// A Boolean value that indicates whether the address is a multicast address.
//
// [Full Topic]: https://developer.apple.com/documentation/virtualization/vzmacaddress/ismulticastaddress
func (v_ VZMACAddress) IsMulticastAddress() bool {
	rv := objc.Send[bool](v_.ID, objc.Sel("isMulticastAddress"))
	return rv
}


// SetIsMulticastAddress sets the value of the isMulticastAddress property.
// A Boolean value that indicates whether the address is a multicast address.

//
// [Full Topic]: https://developer.apple.com/documentation/virtualization/vzmacaddress/ismulticastaddress
func (v_ VZMACAddress) SetIsMulticastAddress(value bool) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setIsMulticastAddress:"), value)
}

// A Boolean value that indicates whether the address is a unicast address.
//
// [Full Topic]: https://developer.apple.com/documentation/virtualization/vzmacaddress/isunicastaddress
func (v_ VZMACAddress) IsUnicastAddress() bool {
	rv := objc.Send[bool](v_.ID, objc.Sel("isUnicastAddress"))
	return rv
}


// SetIsUnicastAddress sets the value of the isUnicastAddress property.
// A Boolean value that indicates whether the address is a unicast address.

//
// [Full Topic]: https://developer.apple.com/documentation/virtualization/vzmacaddress/isunicastaddress
func (v_ VZMACAddress) SetIsUnicastAddress(value bool) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setIsUnicastAddress:"), value)
}

// A Boolean value that indicates whether the address is a universally adminstered address (UAA).
//
// [Full Topic]: https://developer.apple.com/documentation/virtualization/vzmacaddress/isuniversallyadministeredaddress
func (v_ VZMACAddress) IsUniversallyAdministeredAddress() bool {
	rv := objc.Send[bool](v_.ID, objc.Sel("isUniversallyAdministeredAddress"))
	return rv
}


// SetIsUniversallyAdministeredAddress sets the value of the isUniversallyAdministeredAddress property.
// A Boolean value that indicates whether the address is a universally adminstered address (UAA).

//
// [Full Topic]: https://developer.apple.com/documentation/virtualization/vzmacaddress/isuniversallyadministeredaddress
func (v_ VZMACAddress) SetIsUniversallyAdministeredAddress(value bool) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setIsUniversallyAdministeredAddress:"), value)
}

// The MAC address as a formatted string.
//
// [Full Topic]: https://developer.apple.com/documentation/virtualization/vzmacaddress/string
func (v_ VZMACAddress) String() string {
	rv := objc.Send[string](v_.ID, objc.Sel("string"))
	return rv
}


// SetString sets the value of the string property.
// The MAC address as a formatted string.

//
// [Full Topic]: https://developer.apple.com/documentation/virtualization/vzmacaddress/string
func (v_ VZMACAddress) SetString(value string) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setString:"), objc.String(value))
}


