// Code generated from Apple documentation for CoreMIDI. DO NOT EDIT.

package coremidi

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
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


