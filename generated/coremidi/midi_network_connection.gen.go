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



