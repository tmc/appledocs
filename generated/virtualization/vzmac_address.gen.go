// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class VZMACAddress */


/* debug [class_header]: Header for VZMACAddress */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for VZMACAddress */
// An interface definition for the [VZMACAddress] class.
type IVZMACAddress interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for VZMACAddress */
	// properties:
	EthernetAddress() objectivec.IObject
	IsBroadcastAddress() bool
	IsLocallyAdministeredAddress() bool
	IsMulticastAddress() bool
	IsUnicastAddress() bool
	IsUniversallyAdministeredAddress() bool
	String() objc.IObject /* cross-framework: NSString */
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for VZMACAddress */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for VZMACAddress */
// Alloc allocates a new instance without initialization.
func (vc _VZMACAddressClass) Alloc() VZMACAddress {
	rv := objc.Send[VZMACAddress](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for VZMACAddress */
// The media access control (MAC) address for a network interface in your virtual machine.
//
// A object contains the hardware address of your network interface. Every network device has a unique 48-bit MAC address that the system uses to route network packets to that device. Call the method to get a local MAC address suitable for use with your network interfaces. Alternatively, you can create a object yourself from a string or structure.


// The media access control (MAC) address for a network interface in your virtual machine.
//
// [Full Topic]
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for VZMACAddress */

// Creates a MAC address from the specified 48-bit Ethernet address.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZMACAddress/init(ethernetAddress:)
func NewVZMACAddressWithEthernetAddress(ethernetAddress objectivec.IObject) VZMACAddress {
	instance := getVZMACAddressClass().Alloc()
	rv := objc.Send[VZMACAddress](instance.ID, objc.Sel("initWithEthernetAddress:"), ethernetAddress)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewVZMACAddressWithEthernetAddress */


// Creates a MAC address object from a specially formatted string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZMACAddress/init(string:)
func NewVZMACAddressWithString(string_ objc.IObject /* cross-framework: NSString */) VZMACAddress {
	instance := getVZMACAddressClass().Alloc()
	rv := objc.Send[VZMACAddress](instance.ID, objc.Sel("initWithString:"), string_)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewVZMACAddressWithString */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for VZMACAddress */

// Returns a valid, random, locally administered, unicast MAC address.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZMACAddress/randomLocallyAdministered()
func (vc _VZMACAddressClass) RandomLocallyAdministeredAddress() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(vc.class), objc.Sel("randomLocallyAdministeredAddress"))
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=RandomLocallyAdministeredAddress) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for VZMACAddress */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for VZMACAddress */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for VZMACAddress */

// The MAC address as an Ethernet data structure.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZMACAddress/ethernetAddress
func (v_ VZMACAddress) EthernetAddress() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](v_.ID, objc.Sel("ethernetAddress"))
	return rv
}/* debug [instance_properties/getter]: ethernetAddress */


// A Boolean value that indicates whether the address is a broadcast address.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZMACAddress/isBroadcastAddress
func (v_ VZMACAddress) IsBroadcastAddress() bool {
	rv := objc.Send[bool](v_.ID, objc.Sel("isBroadcastAddress"))
	return rv
}/* debug [instance_properties/getter]: isBroadcastAddress */


// A Boolean value that indicates whether the address is a locally administered address (LAA).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZMACAddress/isLocallyAdministeredAddress
func (v_ VZMACAddress) IsLocallyAdministeredAddress() bool {
	rv := objc.Send[bool](v_.ID, objc.Sel("isLocallyAdministeredAddress"))
	return rv
}/* debug [instance_properties/getter]: isLocallyAdministeredAddress */


// A Boolean value that indicates whether the address is a multicast address.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZMACAddress/isMulticastAddress
func (v_ VZMACAddress) IsMulticastAddress() bool {
	rv := objc.Send[bool](v_.ID, objc.Sel("isMulticastAddress"))
	return rv
}/* debug [instance_properties/getter]: isMulticastAddress */


// A Boolean value that indicates whether the address is a unicast address.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZMACAddress/isUnicastAddress
func (v_ VZMACAddress) IsUnicastAddress() bool {
	rv := objc.Send[bool](v_.ID, objc.Sel("isUnicastAddress"))
	return rv
}/* debug [instance_properties/getter]: isUnicastAddress */


// A Boolean value that indicates whether the address is a universally adminstered address (UAA).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZMACAddress/isUniversallyAdministeredAddress
func (v_ VZMACAddress) IsUniversallyAdministeredAddress() bool {
	rv := objc.Send[bool](v_.ID, objc.Sel("isUniversallyAdministeredAddress"))
	return rv
}/* debug [instance_properties/getter]: isUniversallyAdministeredAddress */


// The MAC address as a formatted string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZMACAddress/string
func (v_ VZMACAddress) String() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](v_.ID, objc.Sel("string"))
	return rv
}/* debug [instance_properties/getter]: string */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class VZMACAddress */


