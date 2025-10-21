// Code generated from Apple documentation for CoreMIDI. DO NOT EDIT.

package coremidi

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

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

// An interface definition for the [MIDINetworkHost] class.
type IMIDINetworkHost interface {
	objectivec.IObject
}

// An object that represents the host’s network address.
//
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

// Alloc allocates a new instance without initialization.
func (mc _MIDINetworkHostClass) Alloc() MIDINetworkHost {
	rv := objc.Send[MIDINetworkHost](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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




// Creates a host with the specified name, net service name, and domain.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDINetworkHost/init(name:netServiceName:netServiceDomain:)
func NewMIDINetworkHostWithNameNetServiceNameNetServiceDomain(name string, netServiceName string, netServiceDomain string) MIDINetworkHost {
	rv := objc.Send[MIDINetworkHost](objc.ID(getMIDINetworkHostClass().class), objc.Sel("hostWithName:netServiceName:netServiceDomain:"), objc.String(name), objc.String(netServiceName), objc.String(netServiceDomain))
	return rv
}


// Creates a host with the specified name, net service name, and domain.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDINetworkHost/init(name:netServiceName:netServiceDomain:)
func (mc _MIDINetworkHostClass) HostWithNameNetServiceNameNetServiceDomain(name string, netServiceName string, netServiceDomain string) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(mc.class), objc.Sel("hostWithName:netServiceName:netServiceDomain:"), objc.String(name), objc.String(netServiceName), objc.String(netServiceDomain))
	return rv
}

// The host port.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDINetworkHost/port
func (m_ MIDINetworkHost) Port() uint {
	rv := objc.Send[uint](m_.ID, objc.Sel("port"))
	return rv
}

// The Bonjour service type.
//
// [Full Topic]: https://developer.apple.com/documentation/coremidi/midinetworkbonjourservicetype
func (m_ MIDINetworkHost) MIDINetworkBonjourServiceType() string {
	rv := objc.Send[string](m_.ID, objc.Sel("MIDINetworkBonjourServiceType"))
	return rv
}

// The host address.
//
// [Full Topic]: https://developer.apple.com/documentation/coremidi/midinetworkhost/address
func (m_ MIDINetworkHost) Address() string {
	rv := objc.Send[string](m_.ID, objc.Sel("address"))
	return rv
}


// SetAddress sets the value of the address property.
// The host address.

//
// [Full Topic]: https://developer.apple.com/documentation/coremidi/midinetworkhost/address
func (m_ MIDINetworkHost) SetAddress(value string) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setAddress:"), objc.String(value))
}

// The host name.
//
// [Full Topic]: https://developer.apple.com/documentation/coremidi/midinetworkhost/name
func (m_ MIDINetworkHost) Name() string {
	rv := objc.Send[string](m_.ID, objc.Sel("name"))
	return rv
}


// SetName sets the value of the name property.
// The host name.

//
// [Full Topic]: https://developer.apple.com/documentation/coremidi/midinetworkhost/name
func (m_ MIDINetworkHost) SetName(value string) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setName:"), objc.String(value))
}

// The net service domain.
//
// [Full Topic]: https://developer.apple.com/documentation/coremidi/midinetworkhost/netservicedomain
func (m_ MIDINetworkHost) NetServiceDomain() string {
	rv := objc.Send[string](m_.ID, objc.Sel("netServiceDomain"))
	return rv
}


// SetNetServiceDomain sets the value of the netServiceDomain property.
// The net service domain.

//
// [Full Topic]: https://developer.apple.com/documentation/coremidi/midinetworkhost/netservicedomain
func (m_ MIDINetworkHost) SetNetServiceDomain(value string) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setNetServiceDomain:"), objc.String(value))
}

// The net service name.
//
// [Full Topic]: https://developer.apple.com/documentation/coremidi/midinetworkhost/netservicename
func (m_ MIDINetworkHost) NetServiceName() string {
	rv := objc.Send[string](m_.ID, objc.Sel("netServiceName"))
	return rv
}


// SetNetServiceName sets the value of the netServiceName property.
// The net service name.

//
// [Full Topic]: https://developer.apple.com/documentation/coremidi/midinetworkhost/netservicename
func (m_ MIDINetworkHost) SetNetServiceName(value string) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setNetServiceName:"), objc.String(value))
}


