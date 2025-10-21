// Code generated from Apple documentation for CoreMIDI. DO NOT EDIT.

package coremidi

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [MIDICIProfile] class.
var (
	MIDICIProfileClass     _MIDICIProfileClass
	MIDICIProfileClassOnce sync.Once
)

func getMIDICIProfileClass() _MIDICIProfileClass {
	MIDICIProfileClassOnce.Do(func() {
		MIDICIProfileClass = _MIDICIProfileClass{objc.GetClass("MIDICIProfile")}
	})
	return MIDICIProfileClass
}

type _MIDICIProfileClass struct {
	class objc.Class
}

// An interface definition for the [MIDICIProfile] class.
type IMIDICIProfile interface {
	objectivec.IObject
}

// A mapping of MIDI messages to specific sounds and synthesis behaviors, such as General MIDI, a drawbar organ, and so on.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDICIProfile
type MIDICIProfile struct {
	objectivec.Object
}

// MIDICIProfileFrom constructs a [MIDICIProfile] from an unsafe.Pointer.
//
// A mapping of MIDI messages to specific sounds and synthesis behaviors, such as General MIDI, a drawbar organ, and so on.
func MIDICIProfileFrom(ptr unsafe.Pointer) MIDICIProfile {
	return MIDICIProfile{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MIDICIProfileClass) Alloc() MIDICIProfile {
	rv := objc.Send[MIDICIProfile](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MIDICIProfileClass) New() MIDICIProfile {
	rv := objc.Send[MIDICIProfile](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MIDICIProfile) Init() MIDICIProfile {
	rv := objc.Send[MIDICIProfile](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MIDICIProfile) Autorelease() MIDICIProfile {
	rv := objc.Send[MIDICIProfile](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMIDICIProfile creates a new MIDICIProfile instance.
func NewMIDICIProfile() MIDICIProfile {
	return getMIDICIProfileClass().New()
}




