// Code generated from Apple documentation for CoreMIDI. DO NOT EDIT.

package coremidi

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MIDICIDiscoveredNode] class.
var (
	MIDICIDiscoveredNodeClass     _MIDICIDiscoveredNodeClass
	MIDICIDiscoveredNodeClassOnce sync.Once
)

func getMIDICIDiscoveredNodeClass() _MIDICIDiscoveredNodeClass {
	MIDICIDiscoveredNodeClassOnce.Do(func() {
		MIDICIDiscoveredNodeClass = _MIDICIDiscoveredNodeClass{objc.GetClass("MIDICIDiscoveredNode")}
	})
	return MIDICIDiscoveredNodeClass
}

type _MIDICIDiscoveredNodeClass struct {
	class objc.Class
}

// An interface definition for the [MIDICIDiscoveredNode] class.
type IMIDICIDiscoveredNode interface {
	objectivec.IObject
}

// A discovered MIDI-CI node that represents a MIDI source and destination that respond to capability inquiries.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDICIDiscoveredNode
type MIDICIDiscoveredNode struct {
	objectivec.Object
}

// MIDICIDiscoveredNodeFrom constructs a [MIDICIDiscoveredNode] from an unsafe.Pointer.
//
// A discovered MIDI-CI node that represents a MIDI source and destination that respond to capability inquiries.
func MIDICIDiscoveredNodeFrom(ptr unsafe.Pointer) MIDICIDiscoveredNode {
	return MIDICIDiscoveredNode{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MIDICIDiscoveredNodeClass) Alloc() MIDICIDiscoveredNode {
	rv := objc.Send[MIDICIDiscoveredNode](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MIDICIDiscoveredNodeClass) New() MIDICIDiscoveredNode {
	rv := objc.Send[MIDICIDiscoveredNode](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MIDICIDiscoveredNode) Init() MIDICIDiscoveredNode {
	rv := objc.Send[MIDICIDiscoveredNode](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MIDICIDiscoveredNode) Autorelease() MIDICIDiscoveredNode {
	rv := objc.Send[MIDICIDiscoveredNode](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMIDICIDiscoveredNode creates a new MIDICIDiscoveredNode instance.
func NewMIDICIDiscoveredNode() MIDICIDiscoveredNode {
	return getMIDICIDiscoveredNodeClass().New()
}


// The node’s MIDI destination.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDICIDiscoveredNode/destination
func (m_ MIDICIDiscoveredNode) Destination() MIDIEntityRef {
	rv := objc.Send[MIDIEntityRef](m_.ID, objc.Sel("destination"))
	return rv
}

// The available MIDI-CI device information.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDICIDiscoveredNode/deviceInfo
func (m_ MIDICIDiscoveredNode) DeviceInfo() MIDICIDeviceInfo {
	rv := objc.Send[MIDICIDeviceInfo](m_.ID, objc.Sel("deviceInfo"))
	return rv
}

// The maximum size of a System Exclusive (SysEx) message this node supports.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDICIDiscoveredNode/maximumSysExSize
func (m_ MIDICIDiscoveredNode) MaximumSysExSize() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("maximumSysExSize"))
	return rv
}

// A Boolean value that indicates whether this node supports MIDI-CI profiles.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDICIDiscoveredNode/supportsProfiles
func (m_ MIDICIDiscoveredNode) SupportsProfiles() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("supportsProfiles"))
	return rv
}

// A Boolean value that indicates whether this node supports MIDI-CI properties.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDICIDiscoveredNode/supportsProperties
func (m_ MIDICIDiscoveredNode) SupportsProperties() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("supportsProperties"))
	return rv
}



