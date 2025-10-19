// Code generated from Apple documentation for CoreMIDI. DO NOT EDIT.

package coremidi

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MIDINetworkSession] class.
var mIDINetworkSessionClass = _MIDINetworkSessionClass{objc.GetClass("MIDINetworkSession")}

type _MIDINetworkSessionClass struct {
	class objc.Class
}

// An interface definition for the [MIDINetworkSession] class.
type IMIDINetworkSession interface {
	objectivec.IObject
}

// An object that represents a pairing of a source and destination. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDINetworkSession

type MIDINetworkSession struct {
	objectivec.Object
}

// MIDINetworkSessionFrom constructs a [MIDINetworkSession] from an unsafe.Pointer.
//
// An object that represents a pairing of a source and destination.
func MIDINetworkSessionFrom(ptr unsafe.Pointer) MIDINetworkSession {
	return MIDINetworkSession{objectivec.Object{objc.ID(ptr)}}
}
// Alloc allocates a new instance without initialization.
func (mc _MIDINetworkSessionClass) Alloc() MIDINetworkSession {
	rv := objc.Send[MIDINetworkSession](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new instance with a +1 retain count.
func (mc _MIDINetworkSessionClass) New() MIDINetworkSession {
	rv := objc.Send[MIDINetworkSession](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MIDINetworkSession) Init() MIDINetworkSession {
	rv := objc.Send[MIDINetworkSession](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MIDINetworkSession) Autorelease() MIDINetworkSession {
	rv := objc.Send[MIDINetworkSession](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMIDINetworkSession creates a new MIDINetworkSession instance.
func NewMIDINetworkSession() MIDINetworkSession {
	return mIDINetworkSessionClass.New()
}




