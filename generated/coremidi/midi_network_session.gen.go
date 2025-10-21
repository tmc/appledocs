// Code generated from Apple documentation for CoreMIDI. DO NOT EDIT.

package coremidi

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [MIDINetworkSession] class.
var (
	MIDINetworkSessionClass     _MIDINetworkSessionClass
	MIDINetworkSessionClassOnce sync.Once
)

func getMIDINetworkSessionClass() _MIDINetworkSessionClass {
	MIDINetworkSessionClassOnce.Do(func() {
		MIDINetworkSessionClass = _MIDINetworkSessionClass{objc.GetClass("MIDINetworkSession")}
	})
	return MIDINetworkSessionClass
}

type _MIDINetworkSessionClass struct {
	class objc.Class
}

// An interface definition for the [MIDINetworkSession] class.
type IMIDINetworkSession interface {
	objectivec.IObject
	AddConnection(connection unsafe.Pointer) bool
	RemoveConnection(connection unsafe.Pointer) bool
}

// An object that represents a pairing of a source and destination.
//
// A session can have any number of connections. The system broadcasts output to all connections, and merges input from multiple connections.
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

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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
	return getMIDINetworkSessionClass().New()
}


// Adds a new connection to this session.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDINetworkSession/addConnection(_:)
func (m_ MIDINetworkSession) AddConnection(connection unsafe.Pointer) bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("addConnection:"), connection)
	return rv
}

// Removes a connection from this session.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDINetworkSession/removeConnection(_:)
func (m_ MIDINetworkSession) RemoveConnection(connection unsafe.Pointer) bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("removeConnection:"), connection)
	return rv
}

// A Boolean value that determines whether the session is enabled.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDINetworkSession/isEnabled
func (m_ MIDINetworkSession) Enabled() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("enabled"))
	return rv
}


// SetEnabled sets the value of the enabled property.
// A Boolean value that determines whether the session is enabled.

//
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDINetworkSession/isEnabled
func (m_ MIDINetworkSession) SetEnabled(value bool) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setEnabled:"), value)
}
// The session’s UDP port.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDINetworkSession/networkPort
func (m_ MIDINetworkSession) NetworkPort() uint {
	rv := objc.Send[uint](m_.ID, objc.Sel("networkPort"))
	return rv
}



