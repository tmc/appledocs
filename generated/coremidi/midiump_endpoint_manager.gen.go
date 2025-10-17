// Code generated from Apple documentation for CoreMIDI. DO NOT EDIT.

package coremidi

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MIDIUMPEndpointManager] class.
var mIDIUMPEndpointManagerClass = _MIDIUMPEndpointManagerClass{objc.GetClass("MIDIUMPEndpointManager")}

type _MIDIUMPEndpointManagerClass struct {
	class objc.Class
}

//
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDIUMPEndpointManager

type MIDIUMPEndpointManager struct {
	objectivec.Object
}

// MIDIUMPEndpointManagerFrom constructs a [MIDIUMPEndpointManager] from an unsafe.Pointer.
func MIDIUMPEndpointManagerFrom(ptr unsafe.Pointer) MIDIUMPEndpointManager {
	return MIDIUMPEndpointManager{objectivec.Object{objc.ID(ptr)}}
}



