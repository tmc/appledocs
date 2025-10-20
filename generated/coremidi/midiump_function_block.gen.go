// Code generated from Apple documentation for CoreMIDI. DO NOT EDIT.

package coremidi

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MIDIUMPFunctionBlock] class.
var (
	MIDIUMPFunctionBlockClass     _MIDIUMPFunctionBlockClass
	MIDIUMPFunctionBlockClassOnce sync.Once
)

func getMIDIUMPFunctionBlockClass() _MIDIUMPFunctionBlockClass {
	MIDIUMPFunctionBlockClassOnce.Do(func() {
		MIDIUMPFunctionBlockClass = _MIDIUMPFunctionBlockClass{objc.GetClass("MIDIUMPFunctionBlock")}
	})
	return MIDIUMPFunctionBlockClass
}

type _MIDIUMPFunctionBlockClass struct {
	class objc.Class
}

// An interface definition for the [MIDIUMPFunctionBlock] class.
type IMIDIUMPFunctionBlock interface {
	objectivec.IObject
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

// Alloc allocates a new instance without initialization.
func (mc _MIDIUMPFunctionBlockClass) Alloc() MIDIUMPFunctionBlock {
	rv := objc.Send[MIDIUMPFunctionBlock](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MIDIUMPFunctionBlockClass) New() MIDIUMPFunctionBlock {
	rv := objc.Send[MIDIUMPFunctionBlock](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MIDIUMPFunctionBlock) Init() MIDIUMPFunctionBlock {
	rv := objc.Send[MIDIUMPFunctionBlock](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MIDIUMPFunctionBlock) Autorelease() MIDIUMPFunctionBlock {
	rv := objc.Send[MIDIUMPFunctionBlock](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMIDIUMPFunctionBlock creates a new MIDIUMPFunctionBlock instance.
func NewMIDIUMPFunctionBlock() MIDIUMPFunctionBlock {
	return getMIDIUMPFunctionBlockClass().New()
}


//
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDIUMPFunctionBlock/direction
func (m_ MIDIUMPFunctionBlock) Direction() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("direction"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDIUMPFunctionBlock/firstGroup
func (m_ MIDIUMPFunctionBlock) FirstGroup() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("firstGroup"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDIUMPFunctionBlock/functionBlockID
func (m_ MIDIUMPFunctionBlock) FunctionBlockID() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("functionBlockID"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDIUMPFunctionBlock/isEnabled
func (m_ MIDIUMPFunctionBlock) IsEnabled() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("isEnabled"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDIUMPFunctionBlock/maxSysEx8Streams
func (m_ MIDIUMPFunctionBlock) MaxSysEx8Streams() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("maxSysEx8Streams"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDIUMPFunctionBlock/midi1Info
func (m_ MIDIUMPFunctionBlock) MIDI1Info() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("MIDI1Info"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDIUMPFunctionBlock/midiCIDevice
func (m_ MIDIUMPFunctionBlock) MidiCIDevice() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("midiCIDevice"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDIUMPFunctionBlock/name
func (m_ MIDIUMPFunctionBlock) Name() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("name"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDIUMPFunctionBlock/totalGroupsSpanned
func (m_ MIDIUMPFunctionBlock) TotalGroupsSpanned() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("totalGroupsSpanned"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDIUMPFunctionBlock/uiHint
func (m_ MIDIUMPFunctionBlock) UIHint() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("UIHint"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDIUMPFunctionBlock/umpEndpoint
func (m_ MIDIUMPFunctionBlock) UMPEndpoint() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("UMPEndpoint"))
	return rv
}



