// Code generated from Apple documentation for AVFAudio. DO NOT EDIT.

package avfaudio

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [AudioSessionCapability] class.
var (
	AudioSessionCapabilityClass     _AudioSessionCapabilityClass
	AudioSessionCapabilityClassOnce sync.Once
)

func getAudioSessionCapabilityClass() _AudioSessionCapabilityClass {
	AudioSessionCapabilityClassOnce.Do(func() {
		AudioSessionCapabilityClass = _AudioSessionCapabilityClass{objc.GetClass("AVAudioSessionCapability")}
	})
	return AudioSessionCapabilityClass
}

type _AudioSessionCapabilityClass struct {
	class objc.Class
}

// An interface definition for the [AudioSessionCapability] class.
type IAudioSessionCapability interface {
	objectivec.IObject
	Enabled() bool
	Supported() bool
	IsEnabled() bool
	SetIsEnabled(value bool)
	IsSupported() bool
	SetIsSupported(value bool)
}

// Describes whether a specific capability is supported and if that capability is currently enabled


// Describes whether a specific capability is supported and if that capability is currently enabled
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioSessionCapability
type AudioSessionCapability struct {
	objectivec.Object
}

// AudioSessionCapabilityFrom constructs a [AudioSessionCapability] from an unsafe.Pointer.
//
// Describes whether a specific capability is supported and if that capability is currently enabled
func AudioSessionCapabilityFrom(ptr unsafe.Pointer) AudioSessionCapability {
	return AudioSessionCapability{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (ac _AudioSessionCapabilityClass) Alloc() AudioSessionCapability {
	rv := objc.Send[AudioSessionCapability](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ac _AudioSessionCapabilityClass) New() AudioSessionCapability {
	rv := objc.Send[AudioSessionCapability](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ AudioSessionCapability) Init() AudioSessionCapability {
	rv := objc.Send[AudioSessionCapability](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ AudioSessionCapability) Autorelease() AudioSessionCapability {
	rv := objc.Send[AudioSessionCapability](a_.ID, objc.Sel("autorelease"))
	return rv
}

// NewAudioSessionCapability creates a new AudioSessionCapability instance.
func NewAudioSessionCapability() AudioSessionCapability {
	return getAudioSessionCapabilityClass().New()
}



// A Boolean value that indicates whether the capability is enabled.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioSessionCapability/isEnabled
func (a_ AudioSessionCapability) Enabled() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("enabled"))
	return rv
}


// A Boolean value that indicates whether the capability is supported.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioSessionCapability/isSupported
func (a_ AudioSessionCapability) Supported() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("supported"))
	return rv
}


// A Boolean value that indicates whether the capability is enabled.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avaudiosessioncapability/isenabled
func (a_ AudioSessionCapability) IsEnabled() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("isEnabled"))
	return rv
}


// A Boolean value that indicates whether the capability is enabled.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avaudiosessioncapability/isenabled
func (a_ AudioSessionCapability) SetIsEnabled(value bool) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setIsEnabled:"), value)
}


// A Boolean value that indicates whether the capability is supported.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avaudiosessioncapability/issupported
func (a_ AudioSessionCapability) IsSupported() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("isSupported"))
	return rv
}


// A Boolean value that indicates whether the capability is supported.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avaudiosessioncapability/issupported
func (a_ AudioSessionCapability) SetIsSupported(value bool) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setIsSupported:"), value)
}



