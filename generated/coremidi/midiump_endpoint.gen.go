// Code generated from Apple documentation for CoreMIDI. DO NOT EDIT.

package coremidi

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MIDIUMPEndpoint] class.
var mIDIUMPEndpointClass = _MIDIUMPEndpointClass{objc.GetClass("MIDIUMPEndpoint")}

type _MIDIUMPEndpointClass struct {
	class objc.Class
}

//
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDIUMPEndpoint

type MIDIUMPEndpoint struct {
	objectivec.Object
}

// MIDIUMPEndpointFrom constructs a [MIDIUMPEndpoint] from an unsafe.Pointer.
func MIDIUMPEndpointFrom(ptr unsafe.Pointer) MIDIUMPEndpoint {
	return MIDIUMPEndpoint{objectivec.Object{objc.ID(ptr)}}
}



