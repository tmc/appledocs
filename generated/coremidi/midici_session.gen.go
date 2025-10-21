// Code generated from Apple documentation for CoreMIDI. DO NOT EDIT.

package coremidi

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MIDICISession] class.
var (
	MIDICISessionClass     _MIDICISessionClass
	MIDICISessionClassOnce sync.Once
)

func getMIDICISessionClass() _MIDICISessionClass {
	MIDICISessionClassOnce.Do(func() {
		MIDICISessionClass = _MIDICISessionClass{objc.GetClass("MIDICISession")}
	})
	return MIDICISessionClass
}

type _MIDICISessionClass struct {
	class objc.Class
}

// An interface definition for the [MIDICISession] class.
type IMIDICISession interface {
	objectivec.IObject
}

// An object that represents a MIDI-CI session.
//
// A MIDI-CI session is a bidirectional communication path between a MIDI source and destination identified using MIDI-CI discovery. Use a session to manipulate MIDI-CI profiles and to discover device capabilities.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDICISession
type MIDICISession struct {
	objectivec.Object
}

// MIDICISessionFrom constructs a [MIDICISession] from an unsafe.Pointer.
//
// An object that represents a MIDI-CI session.
func MIDICISessionFrom(ptr unsafe.Pointer) MIDICISession {
	return MIDICISession{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MIDICISessionClass) Alloc() MIDICISession {
	rv := objc.Send[MIDICISession](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MIDICISessionClass) New() MIDICISession {
	rv := objc.Send[MIDICISession](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MIDICISession) Init() MIDICISession {
	rv := objc.Send[MIDICISession](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MIDICISession) Autorelease() MIDICISession {
	rv := objc.Send[MIDICISession](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMIDICISession creates a new MIDICISession instance.
func NewMIDICISession() MIDICISession {
	return getMIDICISessionClass().New()
}




// Creates a MIDI-CI session.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDICISession/init(discoveredNode:dataReadyHandler:disconnectHandler:)
func NewMIDICISessionWithDiscoveredNodeDataReadyHandlerDisconnectHandler(discoveredNode unsafe.Pointer, handler unsafe.Pointer, disconnectHandler unsafe.Pointer) MIDICISession {
	instance := getMIDICISessionClass().Alloc()
	rv := objc.Send[MIDICISession](instance.ID, objc.Sel("initWithDiscoveredNode:dataReadyHandler:disconnectHandler:"), discoveredNode, handler, disconnectHandler)
	rv.Autorelease()
	return rv
}


// Information about a MIDI-CI device.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDICISession/deviceInfo
func (m_ MIDICISession) DeviceInfo() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("deviceInfo"))
	return rv
}

// The maximum size of System Exclusive (SysEx) messages.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDICISession/maxSysExSize
func (m_ MIDICISession) MaxSysExSize() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("maxSysExSize"))
	return rv
}

// The MIDI destination with which the session is communicating.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDICISession/midiDestination
func (m_ MIDICISession) MidiDestination() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("midiDestination"))
	return rv
}


