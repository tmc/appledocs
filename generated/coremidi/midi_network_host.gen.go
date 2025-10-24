// Code generated from Apple documentation for CoreMIDI. DO NOT EDIT.

package coremidi

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MIDINetworkHost */


/* debug [class_header]: Header for MIDINetworkHost */
// The class instance for the [MIDINetworkHost] class.
var (
	MIDINetworkHostClass     _MIDINetworkHostClass
	MIDINetworkHostClassOnce sync.Once
)

func getMIDINetworkHostClass() _MIDINetworkHostClass {
	MIDINetworkHostClassOnce.Do(func() {
		MIDINetworkHostClass = _MIDINetworkHostClass{objc.GetClass("MIDINetworkHost")}
	})
	return MIDINetworkHostClass
}

type _MIDINetworkHostClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MIDINetworkHost */
// An interface definition for the [MIDINetworkHost] class.
type IMIDINetworkHost interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MIDINetworkHost */
	// properties:
	Address() objc.IObject /* cross-framework: NSString */
	Name() objc.IObject /* cross-framework: NSString */
	NetServiceDomain() objc.IObject /* cross-framework: NSString */
	NetServiceName() objc.IObject /* cross-framework: NSString */
	Port() uint
	MIDINetworkBonjourServiceType() objc.IObject /* cross-framework: NSString */
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MIDINetworkHost */
	// methods:
	HasSameAddressAs(other IMIDINetworkHost) bool
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MIDINetworkHost */
// Alloc allocates a new instance without initialization.
func (mc _MIDINetworkHostClass) Alloc() MIDINetworkHost {
	rv := objc.Send[MIDINetworkHost](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MIDINetworkHostClass) New() MIDINetworkHost {
	rv := objc.Send[MIDINetworkHost](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MIDINetworkHost) Init() MIDINetworkHost {
	rv := objc.Send[MIDINetworkHost](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MIDINetworkHost) Autorelease() MIDINetworkHost {
	rv := objc.Send[MIDINetworkHost](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMIDINetworkHost creates a new MIDINetworkHost instance.
func NewMIDINetworkHost() MIDINetworkHost {
	return getMIDINetworkHostClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MIDINetworkHost */
// An object that represents the host’s network address.


// An object that represents the host’s network address.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDINetworkHost
type MIDINetworkHost struct {
	objectivec.Object
}

// MIDINetworkHostFrom constructs a [MIDINetworkHost] from an unsafe.Pointer.
//
// An object that represents the host’s network address.
func MIDINetworkHostFrom(ptr unsafe.Pointer) MIDINetworkHost {
	return MIDINetworkHost{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MIDINetworkHost */

// Creates a host with the specified name, adress, and port.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDINetworkHost/init(name:address:port:)
func NewMIDINetworkHostWithNameAddressPort(name objc.IObject /* cross-framework: NSString */, address objc.IObject /* cross-framework: NSString */, port uint) MIDINetworkHost {
	rv := objc.Send[MIDINetworkHost](objc.ID(getMIDINetworkHostClass().class), objc.Sel("hostWithName:address:port:"), name, address, port)
	return rv
}/* debug [class_init_methods/constructor]: NewMIDINetworkHostWithNameAddressPort */


// Creates a host with the specified name and net service.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDINetworkHost/init(name:netService:)
func NewMIDINetworkHostWithNameNetService(name objc.IObject /* cross-framework: NSString */, netService foundation.NetService) MIDINetworkHost {
	rv := objc.Send[MIDINetworkHost](objc.ID(getMIDINetworkHostClass().class), objc.Sel("hostWithName:netService:"), name, netService)
	return rv
}/* debug [class_init_methods/constructor]: NewMIDINetworkHostWithNameNetService */


// Creates a host with the specified name, net service name, and domain.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDINetworkHost/init(name:netServiceName:netServiceDomain:)
func NewMIDINetworkHostWithNameNetServiceNameNetServiceDomain(name objc.IObject /* cross-framework: NSString */, netServiceName objc.IObject /* cross-framework: NSString */, netServiceDomain objc.IObject /* cross-framework: NSString */) MIDINetworkHost {
	rv := objc.Send[MIDINetworkHost](objc.ID(getMIDINetworkHostClass().class), objc.Sel("hostWithName:netServiceName:netServiceDomain:"), name, netServiceName, netServiceDomain)
	return rv
}/* debug [class_init_methods/constructor]: NewMIDINetworkHostWithNameNetServiceNameNetServiceDomain */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MIDINetworkHost */

// Creates a host with the specified name, adress, and port.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDINetworkHost/init(name:address:port:)
func (mc _MIDINetworkHostClass) HostWithNameAddressPort(name objc.IObject /* cross-framework: NSString */, address objc.IObject /* cross-framework: NSString */, port uint) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(mc.class), objc.Sel("hostWithName:address:port:"), name, address, port)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=HostWithNameAddressPort) */


// Creates a host with the specified name and net service.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDINetworkHost/init(name:netService:)
func (mc _MIDINetworkHostClass) HostWithNameNetService(name objc.IObject /* cross-framework: NSString */, netService foundation.NetService) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(mc.class), objc.Sel("hostWithName:netService:"), name, netService)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=HostWithNameNetService) */


// Creates a host with the specified name, net service name, and domain.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDINetworkHost/init(name:netServiceName:netServiceDomain:)
func (mc _MIDINetworkHostClass) HostWithNameNetServiceNameNetServiceDomain(name objc.IObject /* cross-framework: NSString */, netServiceName objc.IObject /* cross-framework: NSString */, netServiceDomain objc.IObject /* cross-framework: NSString */) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(mc.class), objc.Sel("hostWithName:netServiceName:netServiceDomain:"), name, netServiceName, netServiceDomain)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=HostWithNameNetServiceNameNetServiceDomain) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MIDINetworkHost */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MIDINetworkHost */

// Compares this host instance with another to see if they share the same address value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDINetworkHost/hasSameAddress(as:)
func (m_ MIDINetworkHost) HasSameAddressAs(other IMIDINetworkHost) bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("hasSameAddressAs:"), other)
	return rv
}/* debug [instance_methods/method]: HasSameAddressAs */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MIDINetworkHost */

// The host address.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDINetworkHost/address
func (m_ MIDINetworkHost) Address() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("address"))
	return rv
}/* debug [instance_properties/getter]: address */


// The host name.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDINetworkHost/name
func (m_ MIDINetworkHost) Name() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("name"))
	return rv
}/* debug [instance_properties/getter]: name */


// The net service domain.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDINetworkHost/netServiceDomain
func (m_ MIDINetworkHost) NetServiceDomain() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("netServiceDomain"))
	return rv
}/* debug [instance_properties/getter]: netServiceDomain */


// The net service name.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDINetworkHost/netServiceName
func (m_ MIDINetworkHost) NetServiceName() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("netServiceName"))
	return rv
}/* debug [instance_properties/getter]: netServiceName */


// The host port.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDINetworkHost/port
func (m_ MIDINetworkHost) Port() uint {
	rv := objc.Send[uint](m_.ID, objc.Sel("port"))
	return rv
}/* debug [instance_properties/getter]: port */


// The Bonjour service type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coremidi/midinetworkbonjourservicetype
func (m_ MIDINetworkHost) MIDINetworkBonjourServiceType() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("MIDINetworkBonjourServiceType"))
	return rv
}/* debug [instance_properties/getter]: MIDINetworkBonjourServiceType */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MIDINetworkHost */


