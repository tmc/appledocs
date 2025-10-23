// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [Sound] class.
var (
	SoundClass     _SoundClass
	SoundClassOnce sync.Once
)

func getSoundClass() _SoundClass {
	SoundClassOnce.Do(func() {
		SoundClass = _SoundClass{objc.GetClass("NSSound")}
	})
	return SoundClass
}

type _SoundClass struct {
	class objc.Class
}

// An interface definition for the [Sound] class.
type ISound interface {
	objectivec.IObject
	CurrentTime() unsafe.Pointer
	SetCurrentTime(value unsafe.Pointer)
	Delegate() unsafe.Pointer
	SetDelegate(value unsafe.Pointer)
	Duration() unsafe.Pointer
	SetDuration(value unsafe.Pointer)
	IsPlaying() bool
	SetIsPlaying(value bool)
	Loops() bool
	SetLoops(value bool)
	Name() unsafe.Pointer
	SetName(value unsafe.Pointer)
	PlaybackDeviceIdentifier() unsafe.Pointer
	SetPlaybackDeviceIdentifier(value unsafe.Pointer)
	Volume() float32
	SetVolume(value float32)
}

// A simple interface for loading and playing audio files.
//
// You create a sound object with an audio file or data, which can be in any format that Core Audio supports. Customize the sound by configuring its properties, such as setting its playback volume and looping behavior. Call the sound’s method to begin playback. The system executes this call asynchronously so that it doesn’t interrupt the functioning of your app. If you want to play the system beep sound, use the (Swift) or (Objective-C) function.


// A simple interface for loading and playing audio files.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSound
type Sound struct {
	objectivec.Object
}

// SoundFrom constructs a [Sound] from an unsafe.Pointer.
//
// A simple interface for loading and playing audio files.
func SoundFrom(ptr unsafe.Pointer) Sound {
	return Sound{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (sc _SoundClass) Alloc() Sound {
	rv := objc.Send[Sound](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (sc _SoundClass) New() Sound {
	rv := objc.Send[Sound](objc.ID(sc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (s_ Sound) Init() Sound {
	rv := objc.Send[Sound](s_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s_ Sound) Autorelease() Sound {
	rv := objc.Send[Sound](s_.ID, objc.Sel("autorelease"))
	return rv
}

// NewSound creates a new Sound instance.
func NewSound() Sound {
	return getSoundClass().New()
}



// The sound’s playback progress, in seconds.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssound/currenttime
func (s_ Sound) CurrentTime() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("currentTime"))
	return rv
}


// The sound’s playback progress, in seconds.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssound/currenttime
func (s_ Sound) SetCurrentTime(value unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setCurrentTime:"), value)
}


// The sound’s delegate.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssound/delegate
func (s_ Sound) Delegate() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("delegate"))
	return rv
}


// The sound’s delegate.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssound/delegate
func (s_ Sound) SetDelegate(value unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setDelegate:"), value)
}


// The duration of the sound, in seconds.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssound/duration
func (s_ Sound) Duration() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("duration"))
	return rv
}


// The duration of the sound, in seconds.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssound/duration
func (s_ Sound) SetDuration(value unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setDuration:"), value)
}


// A Boolean that indicates whether the sound is playing its audio data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssound/isplaying
func (s_ Sound) IsPlaying() bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("isPlaying"))
	return rv
}


// A Boolean that indicates whether the sound is playing its audio data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssound/isplaying
func (s_ Sound) SetIsPlaying(value bool) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setIsPlaying:"), value)
}


// A Boolean that indicates whether the sound restarts playback when it reaches the end of its content.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssound/loops
func (s_ Sound) Loops() bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("loops"))
	return rv
}


// A Boolean that indicates whether the sound restarts playback when it reaches the end of its content.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssound/loops
func (s_ Sound) SetLoops(value bool) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setLoops:"), value)
}


// The name assigned to the sound.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssound/name-swift.property
func (s_ Sound) Name() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("name"))
	return rv
}


// The name assigned to the sound.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssound/name-swift.property
func (s_ Sound) SetName(value unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setName:"), value)
}


// Identifies the sound’s output device
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssound/playbackdeviceidentifier-swift.property
func (s_ Sound) PlaybackDeviceIdentifier() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("playbackDeviceIdentifier"))
	return rv
}


// Identifies the sound’s output device
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssound/playbackdeviceidentifier-swift.property
func (s_ Sound) SetPlaybackDeviceIdentifier(value unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setPlaybackDeviceIdentifier:"), value)
}


// The volume of the sound.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssound/volume
func (s_ Sound) Volume() float32 {
	rv := objc.Send[float32](s_.ID, objc.Sel("volume"))
	return rv
}


// The volume of the sound.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssound/volume
func (s_ Sound) SetVolume(value float32) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setVolume:"), value)
}



