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



