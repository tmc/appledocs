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
	// properties:
	Address() string /* primitive/slice/pointer. */
	Name() string /* primitive/slice/pointer. */
	NetServiceDomain() string /* primitive/slice/pointer. */
	NetServiceName() string /* primitive/slice/pointer. */
	MIDINetworkBonjourServiceType() string /* primitive/slice/pointer. */
	Port() int /* primitive/slice/pointer. */
	SetPort(value int /* primitive/slice/pointer. */)
	// methods:
	HasSameAddressAs(other IMIDINetworkHost) bool /* primitive/slice/pointer. */
}

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



// Creates a host with the specified name, adress, and port.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDINetworkHost/init(name:address:port:)
func NewMIDINetworkHostWithNameAddressPort(name string /* primitive/slice/pointer. */, address string /* primitive/slice/pointer. */, port uint /* primitive/slice/pointer. */) MIDINetworkHost {
	rv := objc.Send[MIDINetworkHost](objc.ID(getMIDINetworkHostClass().class), objc.Sel("hostWithName:address:port:"), objc.String(name), objc.String(address), port)
	return rv
}


// Creates a host with the specified name and net service.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDINetworkHost/init(name:netService:)
func NewMIDINetworkHostWithNameNetService(name string /* primitive/slice/pointer. */, netService NetService /* not a class type */) MIDINetworkHost {
	rv := objc.Send[MIDINetworkHost](objc.ID(getMIDINetworkHostClass().class), objc.Sel("hostWithName:netService:"), objc.String(name), netService)
	return rv
}



// Creates a host with the specified name, adress, and port.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDINetworkHost/init(name:address:port:)
func (mc _MIDINetworkHostClass) HostWithNameAddressPort(name string /* primitive/slice/pointer. */, address string /* primitive/slice/pointer. */, port uint /* primitive/slice/pointer. */) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(mc.class), objc.Sel("hostWithName:address:port:"), objc.String(name), objc.String(address), port)
	return rv
}


// Creates a host with the specified name and net service.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDINetworkHost/init(name:netService:)
func (mc _MIDINetworkHostClass) HostWithNameNetService(name string /* primitive/slice/pointer. */, netService NetService /* not a class type */) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(mc.class), objc.Sel("hostWithName:netService:"), objc.String(name), netService)
	return rv
}


// Compares this host instance with another to see if they share the same address value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDINetworkHost/hasSameAddress(as:)
func (m_ MIDINetworkHost) HasSameAddressAs(other IMIDINetworkHost) bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](m_.ID, objc.Sel("hasSameAddressAs:"), other)
	return rv
}


// The host address.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDINetworkHost/address
func (m_ MIDINetworkHost) Address() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](m_.ID, objc.Sel("address"))
	return rv
}


// The host name.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDINetworkHost/name
func (m_ MIDINetworkHost) Name() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](m_.ID, objc.Sel("name"))
	return rv
}


// The net service domain.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDINetworkHost/netServiceDomain
func (m_ MIDINetworkHost) NetServiceDomain() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](m_.ID, objc.Sel("netServiceDomain"))
	return rv
}


// The net service name.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDINetworkHost/netServiceName
func (m_ MIDINetworkHost) NetServiceName() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](m_.ID, objc.Sel("netServiceName"))
	return rv
}


// The Bonjour service type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coremidi/midinetworkbonjourservicetype
func (m_ MIDINetworkHost) MIDINetworkBonjourServiceType() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](m_.ID, objc.Sel("MIDINetworkBonjourServiceType"))
	return rv
}


// The host port.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coremidi/midinetworkhost/port
func (m_ MIDINetworkHost) Port() int /* primitive/slice/pointer. */ {
	rv := objc.Send[int](m_.ID, objc.Sel("port"))
	return rv
}


// The host port.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coremidi/midinetworkhost/port
func (m_ MIDINetworkHost) SetPort(value int /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setPort:"), value)
}


