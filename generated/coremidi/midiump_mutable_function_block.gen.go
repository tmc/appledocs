// Code generated from Apple documentation for CoreMIDI. DO NOT EDIT.

package coremidi

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [MIDIUMPMutableFunctionBlock] class.
var (
	MIDIUMPMutableFunctionBlockClass     _MIDIUMPMutableFunctionBlockClass
	MIDIUMPMutableFunctionBlockClassOnce sync.Once
)

func getMIDIUMPMutableFunctionBlockClass() _MIDIUMPMutableFunctionBlockClass {
	MIDIUMPMutableFunctionBlockClassOnce.Do(func() {
		MIDIUMPMutableFunctionBlockClass = _MIDIUMPMutableFunctionBlockClass{objc.GetClass("MIDIUMPMutableFunctionBlock")}
	})
	return MIDIUMPMutableFunctionBlockClass
}

type _MIDIUMPMutableFunctionBlockClass struct {
	class objc.Class
}

// An interface definition for the [MIDIUMPMutableFunctionBlock] class.
type IMIDIUMPMutableFunctionBlock interface {
	IMIDIUMPFunctionBlock
	ReconfigureWithFirstGroupDirectionMIDI1InfoUIHintError(firstGroup IMIDIUMPGroupNumber, direction MIDIUMPFunctionBlockDirection, MIDI1Info IMIDIUMPFunctionBlockMIDI1Info, UIHint IMIDIUMPFunctionBlockUIHint, error_ unsafe.Pointer) bool
	SetEnabledError(isEnabled bool, error_ unsafe.Pointer) bool
	SetNameError(name string, error_ unsafe.Pointer) bool
	UMPEndpoint() MIDIUMPMutableEndpoint
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDIUMPMutableFunctionBlock

type MIDIUMPMutableFunctionBlock struct {
	MIDIUMPFunctionBlock
}

// MIDIUMPMutableFunctionBlockFrom constructs a [MIDIUMPMutableFunctionBlock] from an unsafe.Pointer.
func MIDIUMPMutableFunctionBlockFrom(ptr unsafe.Pointer) MIDIUMPMutableFunctionBlock {
	return MIDIUMPMutableFunctionBlock{
		MIDIUMPFunctionBlock: MIDIUMPFunctionBlockFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (mc _MIDIUMPMutableFunctionBlockClass) Alloc() MIDIUMPMutableFunctionBlock {
	rv := objc.Send[MIDIUMPMutableFunctionBlock](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MIDIUMPMutableFunctionBlockClass) New() MIDIUMPMutableFunctionBlock {
	rv := objc.Send[MIDIUMPMutableFunctionBlock](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MIDIUMPMutableFunctionBlock) Init() MIDIUMPMutableFunctionBlock {
	rv := objc.Send[MIDIUMPMutableFunctionBlock](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MIDIUMPMutableFunctionBlock) Autorelease() MIDIUMPMutableFunctionBlock {
	rv := objc.Send[MIDIUMPMutableFunctionBlock](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMIDIUMPMutableFunctionBlock creates a new MIDIUMPMutableFunctionBlock instance.
func NewMIDIUMPMutableFunctionBlock() MIDIUMPMutableFunctionBlock {
	return getMIDIUMPMutableFunctionBlockClass().New()
}




// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDIUMPMutableFunctionBlock/init(name:direction:firstGroup:totalGroupsSpanned:maxSysEx8Streams:midi1Info:uiHint:isEnabled:)

func NewMIDIUMPMutableFunctionBlockWithNameDirectionFirstGroupTotalGroupsSpannedMaxSysEx8StreamsMIDI1InfoUIHintIsEnabled(name string, direction MIDIUMPFunctionBlockDirection, firstGroup IMIDIUMPGroupNumber, totalGroupsSpanned IMIDIUInteger7, maxSysEx8Streams IMIDIUInteger7, MIDI1Info IMIDIUMPFunctionBlockMIDI1Info, UIHint IMIDIUMPFunctionBlockUIHint, isEnabled bool) MIDIUMPMutableFunctionBlock {
	instance := getMIDIUMPMutableFunctionBlockClass().Alloc()
	rv := objc.Send[MIDIUMPMutableFunctionBlock](instance.ID, objc.Sel("initWithName:direction:firstGroup:totalGroupsSpanned:maxSysEx8Streams:MIDI1Info:UIHint:isEnabled:"), objc.String(name), direction, firstGroup, totalGroupsSpanned, maxSysEx8Streams, MIDI1Info, UIHint, isEnabled)
	rv.Autorelease()
	return rv
}




// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDIUMPMutableFunctionBlock/reconfigure(firstGroup:direction:MIDI1Info:UIHint:)

func (m_ MIDIUMPMutableFunctionBlock) ReconfigureWithFirstGroupDirectionMIDI1InfoUIHintError(firstGroup IMIDIUMPGroupNumber, direction MIDIUMPFunctionBlockDirection, MIDI1Info IMIDIUMPFunctionBlockMIDI1Info, UIHint IMIDIUMPFunctionBlockUIHint, error_ unsafe.Pointer) bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("reconfigureWithFirstGroup:direction:MIDI1Info:UIHint:error:"), firstGroup, direction, MIDI1Info, UIHint, error_)
	return rv
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDIUMPMutableFunctionBlock/setEnabled(_:)

func (m_ MIDIUMPMutableFunctionBlock) SetEnabledError(isEnabled bool, error_ unsafe.Pointer) bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("setEnabled:error:"), isEnabled, error_)
	return rv
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDIUMPMutableFunctionBlock/setName(_:)

func (m_ MIDIUMPMutableFunctionBlock) SetNameError(name string, error_ unsafe.Pointer) bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("setName:error:"), objc.String(name), error_)
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDIUMPMutableFunctionBlock/umpEndpoint

func (m_ MIDIUMPMutableFunctionBlock) UMPEndpoint() MIDIUMPMutableEndpoint {
	rv := objc.Send[MIDIUMPMutableEndpoint](m_.ID, objc.Sel("UMPEndpoint"))
	return rv
}


