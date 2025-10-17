// Code generated from Apple documentation for CoreMIDI. DO NOT EDIT.

package coremidi

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MIDIUMPFunctionBlock] class.
var mIDIUMPFunctionBlockClass = _MIDIUMPFunctionBlockClass{objc.GetClass("MIDIUMPFunctionBlock")}

type _MIDIUMPFunctionBlockClass struct {
	class objc.Class
}

//
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDIUMPFunctionBlock

type MIDIUMPFunctionBlock struct {
	objectivec.Object
}

// MIDIUMPFunctionBlockFrom constructs a [MIDIUMPFunctionBlock] from an unsafe.Pointer.
func MIDIUMPFunctionBlockFrom(ptr unsafe.Pointer) MIDIUMPFunctionBlock {
	return MIDIUMPFunctionBlock{objectivec.Object{objc.ID(ptr)}}
}



