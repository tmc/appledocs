// Code generated from Apple documentation for CoreMIDI. DO NOT EDIT.

package coremidi

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MIDIUMPCIProfile] class.
var (
	MIDIUMPCIProfileClass     _MIDIUMPCIProfileClass
	MIDIUMPCIProfileClassOnce sync.Once
)

func getMIDIUMPCIProfileClass() _MIDIUMPCIProfileClass {
	MIDIUMPCIProfileClassOnce.Do(func() {
		MIDIUMPCIProfileClass = _MIDIUMPCIProfileClass{objc.GetClass("MIDIUMPCIProfile")}
	})
	return MIDIUMPCIProfileClass
}

type _MIDIUMPCIProfileClass struct {
	class objc.Class
}

// An interface definition for the [MIDIUMPCIProfile] class.
type IMIDIUMPCIProfile interface {
	objectivec.IObject
	SetProfileStateEnabledChannelCountError(isEnabled bool, enabledChannelCount unsafe.Pointer, error_ unsafe.Pointer) bool
}

//
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDIUMPCIProfile
type MIDIUMPCIProfile struct {
	objectivec.Object
}

// MIDIUMPCIProfileFrom constructs a [MIDIUMPCIProfile] from an unsafe.Pointer.
func MIDIUMPCIProfileFrom(ptr unsafe.Pointer) MIDIUMPCIProfile {
	return MIDIUMPCIProfile{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MIDIUMPCIProfileClass) Alloc() MIDIUMPCIProfile {
	rv := objc.Send[MIDIUMPCIProfile](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MIDIUMPCIProfileClass) New() MIDIUMPCIProfile {
	rv := objc.Send[MIDIUMPCIProfile](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MIDIUMPCIProfile) Init() MIDIUMPCIProfile {
	rv := objc.Send[MIDIUMPCIProfile](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MIDIUMPCIProfile) Autorelease() MIDIUMPCIProfile {
	rv := objc.Send[MIDIUMPCIProfile](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMIDIUMPCIProfile creates a new MIDIUMPCIProfile instance.
func NewMIDIUMPCIProfile() MIDIUMPCIProfile {
	return getMIDIUMPCIProfileClass().New()
}


//
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDIUMPCIProfile/setProfileState(_:enabledChannelCount:)
func (m_ MIDIUMPCIProfile) SetProfileStateEnabledChannelCountError(isEnabled bool, enabledChannelCount unsafe.Pointer, error_ unsafe.Pointer) bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("setProfileState:enabledChannelCount:error:"), isEnabled, enabledChannelCount, error_)
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDIUMPCIProfile/enabledChannelCount
func (m_ MIDIUMPCIProfile) EnabledChannelCount() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("enabledChannelCount"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDIUMPCIProfile/firstChannel
func (m_ MIDIUMPCIProfile) FirstChannel() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("firstChannel"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDIUMPCIProfile/groupOffset
func (m_ MIDIUMPCIProfile) GroupOffset() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("groupOffset"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDIUMPCIProfile/isEnabled
func (m_ MIDIUMPCIProfile) IsEnabled() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("isEnabled"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDIUMPCIProfile/name
func (m_ MIDIUMPCIProfile) Name() string {
	rv := objc.Send[string](m_.ID, objc.Sel("name"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDIUMPCIProfile/profileID
func (m_ MIDIUMPCIProfile) ProfileID() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("profileID"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDIUMPCIProfile/profileType
func (m_ MIDIUMPCIProfile) ProfileType() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("profileType"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDIUMPCIProfile/totalChannelCount
func (m_ MIDIUMPCIProfile) TotalChannelCount() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("totalChannelCount"))
	return rv
}



