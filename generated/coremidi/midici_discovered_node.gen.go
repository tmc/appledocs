// Code generated from Apple documentation for CoreMIDI. DO NOT EDIT.

package coremidi

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MIDICIDiscoveredNode] class.
var mIDICIDiscoveredNodeClass = _MIDICIDiscoveredNodeClass{objc.GetClass("MIDICIDiscoveredNode")}

type _MIDICIDiscoveredNodeClass struct {
	class objc.Class
}

// An interface definition for the [MIDICIDiscoveredNode] class.
type IMIDICIDiscoveredNode interface {
	objectivec.IObject
}

// A discovered MIDI-CI node that represents a MIDI source and destination that respond to capability inquiries. [Full Topic]
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

// New creates and returns a new instance with a +1 retain count.
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
	return mIDICIDiscoveredNodeClass.New()
}




