// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
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
	ChannelMapping() foundation.Array
	Pause() bool
	Play() bool
	Resume() bool
	SetChannelMapping(channelMapping objectivec.IObject)
	Stop() bool
	WriteToPasteboard(pasteboard IPasteboard)
	CurrentTime() float64
	SetCurrentTime(value float64)
	Delegate() objc.ID
	SetDelegate(value objc.ID)
	Duration() float64
	Playing() bool
	Loops() bool
	SetLoops(value bool)
	Name() SoundName
	PlaybackDeviceIdentifier() SoundPlaybackDeviceIdentifier
	SetPlaybackDeviceIdentifier(value ISoundPlaybackDeviceIdentifier)
	Volume() float32
	SetVolume(value float32)
	IsPlaying() bool
	SetIsPlaying(value bool)
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



// Returns the instance associated with a given name.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSound/init(named:)
func NewSoundNamed(name ISoundName) Sound {
	rv := objc.Send[Sound](objc.ID(getSoundClass().class), objc.Sel("soundNamed:"), name)
	return rv
}


// Initializes the receiver with the audio data located at a given filepath.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSound/init(contentsOfFile:byReference:)
func NewSoundWithContentsOfFileByReference(path string, byRef bool) Sound {
	instance := getSoundClass().Alloc()
	rv := objc.Send[Sound](instance.ID, objc.Sel("initWithContentsOfFile:byReference:"), objc.String(path), byRef)
	rv.Autorelease()
	return rv
}


// Initializes the receiver with the audio data located at a given URL.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSound/init(contentsOf:byReference:)
func NewSoundWithContentsOfURLByReference(url foundation.IURL, byRef bool) Sound {
	instance := getSoundClass().Alloc()
	rv := objc.Send[Sound](instance.ID, objc.Sel("initWithContentsOfURL:byReference:"), url, byRef)
	rv.Autorelease()
	return rv
}


// Initializes the receiver with a given audio data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSound/init(data:)
func NewSoundWithData(data foundation.IData) Sound {
	instance := getSoundClass().Alloc()
	rv := objc.Send[Sound](instance.ID, objc.Sel("initWithData:"), data)
	rv.Autorelease()
	return rv
}


// Initializes the receiver with data from a pasteboard. The pasteboard should contain a type returned by . expects the data to have a proper magic number, sound header, and data for the formats it supports.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSound/init(pasteboard:)
func NewSoundWithPasteboard(pasteboard IPasteboard) Sound {
	instance := getSoundClass().Alloc()
	rv := objc.Send[Sound](instance.ID, objc.Sel("initWithPasteboard:"), pasteboard)
	rv.Autorelease()
	return rv
}



// Indicates whether the receiver can create an instance of itself from the data in a pasteboard.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSound/canInit(with:)
func (sc _SoundClass) CanInitWithPasteboard(pasteboard IPasteboard) bool {
	rv := objc.Send[bool](objc.ID(sc.class), objc.Sel("canInitWithPasteboard:"), pasteboard)
	return rv
}


// Returns the instance associated with a given name.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSound/init(named:)
func (sc _SoundClass) SoundNamed(name ISoundName) Sound {
	rv := objc.Send[Sound](objc.ID(sc.class), objc.Sel("soundNamed:"), name)
	return rv
}


// Provides the list of file types the class understands.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSound/soundUnfilteredFileTypes
func (sc _SoundClass) SoundUnfilteredFileTypes() foundation.Array {
	rv := objc.Send[foundation.Array](objc.ID(sc.class), objc.Sel("soundUnfilteredFileTypes"))
	return rv
}


// Provides a list of the pasteboard types that the class can accept.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSound/soundUnfilteredPasteboardTypes
func (sc _SoundClass) SoundUnfilteredPasteboardTypes() foundation.Array {
	rv := objc.Send[foundation.Array](objc.ID(sc.class), objc.Sel("soundUnfilteredPasteboardTypes"))
	return rv
}


// Provides the file types the class understands.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSound/soundUnfilteredTypes
func (sc _SoundClass) SoundUnfilteredTypes() []string {
	rv := objc.Send[[]string](objc.ID(sc.class), objc.Sel("soundUnfilteredTypes"))
	return rv
}

// Provides the receiver’s channel map.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSound/channelMapping
func (s_ Sound) ChannelMapping() foundation.Array {
	rv := objc.Send[foundation.Array](s_.ID, objc.Sel("channelMapping"))
	return rv
}


// Pauses audio playback.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSound/pause()
func (s_ Sound) Pause() bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("pause"))
	return rv
}


// Initiates audio playback.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSound/play()
func (s_ Sound) Play() bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("play"))
	return rv
}


// Resumes audio playback.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSound/resume()
func (s_ Sound) Resume() bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("resume"))
	return rv
}


// Specifies the receiver’s channel map.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSound/setChannelMapping:
func (s_ Sound) SetChannelMapping(channelMapping objectivec.IObject) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setChannelMapping:"), channelMapping)
}


// Concludes audio playback.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSound/stop()
func (s_ Sound) Stop() bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("stop"))
	return rv
}


// Writes the receiver’s data to a pasteboard.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSound/write(to:)
func (s_ Sound) WriteToPasteboard(pasteboard IPasteboard) {
	objc.Send[objc.ID](s_.ID, objc.Sel("writeToPasteboard:"), pasteboard)
}


// The sound’s playback progress, in seconds.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSound/currentTime
func (s_ Sound) CurrentTime() float64 {
	rv := objc.Send[float64](s_.ID, objc.Sel("currentTime"))
	return rv
}


// The sound’s playback progress, in seconds.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSound/currentTime
func (s_ Sound) SetCurrentTime(value float64) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setCurrentTime:"), value)
}


// The sound’s delegate.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSound/delegate
func (s_ Sound) Delegate() objc.ID {
	rv := objc.Send[objc.ID](s_.ID, objc.Sel("delegate"))
	return rv
}


// The sound’s delegate.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSound/delegate
func (s_ Sound) SetDelegate(value objc.ID) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setDelegate:"), value)
}


// The duration of the sound, in seconds.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSound/duration
func (s_ Sound) Duration() float64 {
	rv := objc.Send[float64](s_.ID, objc.Sel("duration"))
	return rv
}


// A Boolean that indicates whether the sound is playing its audio data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSound/isPlaying
func (s_ Sound) Playing() bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("playing"))
	return rv
}


// A Boolean that indicates whether the sound restarts playback when it reaches the end of its content.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSound/loops
func (s_ Sound) Loops() bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("loops"))
	return rv
}


// A Boolean that indicates whether the sound restarts playback when it reaches the end of its content.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSound/loops
func (s_ Sound) SetLoops(value bool) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setLoops:"), value)
}


// The name assigned to the sound.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSound/name-swift.property
func (s_ Sound) Name() SoundName {
	rv := objc.Send[SoundName](s_.ID, objc.Sel("name"))
	return rv
}


// Identifies the sound’s output device
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSound/playbackDeviceIdentifier-swift.property
func (s_ Sound) PlaybackDeviceIdentifier() SoundPlaybackDeviceIdentifier {
	rv := objc.Send[SoundPlaybackDeviceIdentifier](s_.ID, objc.Sel("playbackDeviceIdentifier"))
	return rv
}


// Identifies the sound’s output device
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSound/playbackDeviceIdentifier-swift.property
func (s_ Sound) SetPlaybackDeviceIdentifier(value ISoundPlaybackDeviceIdentifier) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setPlaybackDeviceIdentifier:"), value)
}


// Provides the file types the class understands.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSound/soundUnfilteredTypes
func (s_ Sound) SoundUnfilteredTypes() []string {
	rv := objc.Send[[]string](s_.ID, objc.Sel("soundUnfilteredTypes"))
	return rv
}


// The volume of the sound.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSound/volume
func (s_ Sound) Volume() float32 {
	rv := objc.Send[float32](s_.ID, objc.Sel("volume"))
	return rv
}


// The volume of the sound.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSound/volume
func (s_ Sound) SetVolume(value float32) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setVolume:"), value)
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


