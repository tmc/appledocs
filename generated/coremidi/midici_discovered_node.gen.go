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



