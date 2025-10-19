// Code generated from Apple documentation for CoreMIDI. DO NOT EDIT.

package coremidi

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MIDICIResponder] class.
var mIDICIResponderClass = _MIDICIResponderClass{objc.GetClass("MIDICIResponder")}

type _MIDICIResponderClass struct {
	class objc.Class
}

// An interface definition for the [MIDICIResponder] class.
type IMIDICIResponder interface {
	objectivec.IObject
}

// An object that responds to MIDI-CI inquiries from an initiator on behalf of a MIDI client, and handles profile and property exchange operations. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDICIResponder

type MIDICIResponder struct {
	objectivec.Object
}

// MIDICIResponderFrom constructs a [MIDICIResponder] from an unsafe.Pointer.
//
// An object that responds to MIDI-CI inquiries from an initiator on behalf of a MIDI client, and handles profile and property exchange operations.
func MIDICIResponderFrom(ptr unsafe.Pointer) MIDICIResponder {
	return MIDICIResponder{objectivec.Object{objc.ID(ptr)}}
}
// Alloc allocates a new instance without initialization.
func (mc _MIDICIResponderClass) Alloc() MIDICIResponder {
	rv := objc.Send[MIDICIResponder](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new instance with a +1 retain count.
func (mc _MIDICIResponderClass) New() MIDICIResponder {
	rv := objc.Send[MIDICIResponder](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MIDICIResponder) Init() MIDICIResponder {
	rv := objc.Send[MIDICIResponder](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MIDICIResponder) Autorelease() MIDICIResponder {
	rv := objc.Send[MIDICIResponder](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMIDICIResponder creates a new MIDICIResponder instance.
func NewMIDICIResponder() MIDICIResponder {
	return mIDICIResponderClass.New()
}




