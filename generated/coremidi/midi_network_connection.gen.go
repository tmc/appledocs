// Code generated from Apple documentation for CoreMIDI. DO NOT EDIT.

package coremidi

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MIDINetworkConnection] class.
var mIDINetworkConnectionClass = _MIDINetworkConnectionClass{objc.GetClass("MIDINetworkConnection")}

type _MIDINetworkConnectionClass struct {
	class objc.Class
}

// An interface definition for the [MIDINetworkConnection] class.
type IMIDINetworkConnection interface {
	objectivec.IObject
}

// An object that connects a session to a host. [Full Topic]
//
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

// New creates and returns a new instance with a +1 retain count.
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
	return mIDINetworkConnectionClass.New()
}




