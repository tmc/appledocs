// Code generated from Apple documentation for CoreMIDI. DO NOT EDIT.

package coremidi

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MIDINetworkConnection] class.
var (
	MIDINetworkConnectionClass     _MIDINetworkConnectionClass
	MIDINetworkConnectionClassOnce sync.Once
)

func getMIDINetworkConnectionClass() _MIDINetworkConnectionClass {
	MIDINetworkConnectionClassOnce.Do(func() {
		MIDINetworkConnectionClass = _MIDINetworkConnectionClass{objc.GetClass("MIDINetworkConnection")}
	})
	return MIDINetworkConnectionClass
}

type _MIDINetworkConnectionClass struct {
	class objc.Class
}

// An interface definition for the [MIDINetworkConnection] class.
type IMIDINetworkConnection interface {
	objectivec.IObject
	// properties:
	Host() IMIDINetworkHost
	// methods:
}

// An object that connects a session to a host.


// An object that connects a session to a host.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDINetworkConnection
type MIDINetworkConnection struct {
	objectivec.Object
}

// MIDINetworkConnectionFrom constructs a [MIDINetworkConnection] from an unsafe.Pointer.
//
// An object that connects a session to a host.
func MIDINetworkConnectionFrom(ptr unsafe.Pointer) MIDINetworkConnection {
	return MIDINetworkConnection{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MIDINetworkConnectionClass) Alloc() MIDINetworkConnection {
	rv := objc.Send[MIDINetworkConnection](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MIDINetworkConnectionClass) New() MIDINetworkConnection {
	rv := objc.Send[MIDINetworkConnection](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MIDINetworkConnection) Init() MIDINetworkConnection {
	rv := objc.Send[MIDINetworkConnection](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MIDINetworkConnection) Autorelease() MIDINetworkConnection {
	rv := objc.Send[MIDINetworkConnection](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMIDINetworkConnection creates a new MIDINetworkConnection instance.
func NewMIDINetworkConnection() MIDINetworkConnection {
	return getMIDINetworkConnectionClass().New()
}



// Creates a connection to the specified host.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDINetworkConnection/init(host:)
func NewMIDINetworkConnectionWithHost(host IMIDINetworkHost) MIDINetworkConnection {
	rv := objc.Send[MIDINetworkConnection](objc.ID(getMIDINetworkConnectionClass().class), objc.Sel("connectionWithHost:"), host)
	return rv
}



// Creates a connection to the specified host.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDINetworkConnection/init(host:)
func (mc _MIDINetworkConnectionClass) ConnectionWithHost(host IMIDINetworkHost) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(mc.class), objc.Sel("connectionWithHost:"), host)
	return rv
}


// The host connection.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDINetworkConnection/host
func (m_ MIDINetworkConnection) Host() IMIDINetworkHost {
	rv := objc.Send[MIDINetworkHost](m_.ID, objc.Sel("host"))
	return rv
}


