// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class NSSound */


/* debug [class_header]: Header for NSSound */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for Sound */
// An interface definition for the [Sound] class.
type ISound interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for Sound */
	// properties:
	CurrentTime() float64
	SetCurrentTime(value float64)
	Delegate() unsafe.Pointer
	SetDelegate(value unsafe.Pointer)
	Duration() float64
	Playing() bool
	Loops() bool
	SetLoops(value bool)
	Name() SoundName /* typedef */
	PlaybackDeviceIdentifier() SoundPlaybackDeviceIdentifier /* typedef */
	SetPlaybackDeviceIdentifier(value SoundPlaybackDeviceIdentifier /* typedef */)
	Volume() float32
	SetVolume(value float32)
	IsPlaying() bool
	SetIsPlaying(value bool)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for Sound */
	// methods:
	Pause() bool
	Play() bool
	Resume() bool
	Stop() bool
	WriteToPasteboard(pasteboard IPasteboard)
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for Sound */
// Alloc allocates a new instance without initialization.
func (sc _SoundClass) Alloc() Sound {
	rv := objc.Send[Sound](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for Sound */
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for Sound */

// Returns the instance associated with a given name.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSound/init(named:)
func NewSoundNamed(name SoundName /* typedef */) Sound {
	rv := objc.Send[Sound](objc.ID(getSoundClass().class), objc.Sel("soundNamed:"), name)
	return rv
}/* debug [class_init_methods/constructor]: NewSoundNamed */


// Initializes the receiver with the audio data located at a given filepath.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSound/init(contentsOfFile:byReference:)
func NewSoundWithContentsOfFileByReference(path objc.IObject /* cross-framework: NSString */, byRef bool) Sound {
	instance := getSoundClass().Alloc()
	rv := objc.Send[Sound](instance.ID, objc.Sel("initWithContentsOfFile:byReference:"), path, byRef)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewSoundWithContentsOfFileByReference */


// Initializes the receiver with the audio data located at a given URL.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSound/init(contentsOf:byReference:)
func NewSoundWithContentsOfURLByReference(url objc.IObject /* cross-framework: NSURL */, byRef bool) Sound {
	instance := getSoundClass().Alloc()
	rv := objc.Send[Sound](instance.ID, objc.Sel("initWithContentsOfURL:byReference:"), url, byRef)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewSoundWithContentsOfURLByReference */


// Initializes the receiver with a given audio data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSound/init(data:)
func NewSoundWithData(data objc.IObject /* cross-framework: NSData */) Sound {
	instance := getSoundClass().Alloc()
	rv := objc.Send[Sound](instance.ID, objc.Sel("initWithData:"), data)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewSoundWithData */


// Initializes the receiver with data from a pasteboard. The pasteboard should contain a type returned by . expects the data to have a proper magic number, sound header, and data for the formats it supports.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSound/init(pasteboard:)
func NewSoundWithPasteboard(pasteboard IPasteboard) Sound {
	instance := getSoundClass().Alloc()
	rv := objc.Send[Sound](instance.ID, objc.Sel("initWithPasteboard:"), pasteboard)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewSoundWithPasteboard */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for Sound */

// Indicates whether the receiver can create an instance of itself from the data in a pasteboard.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSound/canInit(with:)
func (sc _SoundClass) CanInitWithPasteboard(pasteboard IPasteboard) bool {
	rv := objc.Send[bool](objc.ID(sc.class), objc.Sel("canInitWithPasteboard:"), pasteboard)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=CanInitWithPasteboard) */


// Returns the instance associated with a given name.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSound/init(named:)
func (sc _SoundClass) SoundNamed(name SoundName /* typedef */) ISound {
	rv := objc.Send[Sound](objc.ID(sc.class), objc.Sel("soundNamed:"), name)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=SoundNamed) */


// Provides the list of file types the class understands.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSound/soundUnfilteredFileTypes
func (sc _SoundClass) SoundUnfilteredFileTypes() foundation.Array {
	rv := objc.Send[foundation.Array](objc.ID(sc.class), objc.Sel("soundUnfilteredFileTypes"))
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=SoundUnfilteredFileTypes) */


// Provides a list of the pasteboard types that the class can accept.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSound/soundUnfilteredPasteboardTypes
func (sc _SoundClass) SoundUnfilteredPasteboardTypes() foundation.Array {
	rv := objc.Send[foundation.Array](objc.ID(sc.class), objc.Sel("soundUnfilteredPasteboardTypes"))
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=SoundUnfilteredPasteboardTypes) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for Sound */

// Provides the file types the class understands.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSound/soundUnfilteredTypes
func (sc _SoundClass) SoundUnfilteredTypes() []string {
	rv := objc.Send[[]string](objc.ID(sc.class), objc.Sel("soundUnfilteredTypes"))
	return rv
}/* debug [class_properties_class/property]: soundUnfilteredTypes */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for Sound */

// Pauses audio playback.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSound/pause()
func (s_ Sound) Pause() bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("pause"))
	return rv
}/* debug [instance_methods/method]: Pause */


// Initiates audio playback.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSound/play()
func (s_ Sound) Play() bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("play"))
	return rv
}/* debug [instance_methods/method]: Play */


// Resumes audio playback.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSound/resume()
func (s_ Sound) Resume() bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("resume"))
	return rv
}/* debug [instance_methods/method]: Resume */


// Concludes audio playback.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSound/stop()
func (s_ Sound) Stop() bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("stop"))
	return rv
}/* debug [instance_methods/method]: Stop */


// Writes the receiver’s data to a pasteboard.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSound/write(to:)
func (s_ Sound) WriteToPasteboard(pasteboard IPasteboard) {
	objc.Send[objc.ID](s_.ID, objc.Sel("writeToPasteboard:"), pasteboard)
}/* debug [instance_methods/method]: WriteToPasteboard */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for Sound */

// The sound’s playback progress, in seconds.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSound/currentTime
func (s_ Sound) CurrentTime() float64 {
	rv := objc.Send[float64](s_.ID, objc.Sel("currentTime"))
	return rv
}/* debug [instance_properties/getter]: currentTime */


// The sound’s playback progress, in seconds.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSound/currentTime
func (s_ Sound) SetCurrentTime(value float64) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setCurrentTime:"), value)
}/* debug [instance_properties/setter]: currentTime */


// The sound’s delegate.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSound/delegate
func (s_ Sound) Delegate() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("delegate"))
	return rv
}/* debug [instance_properties/getter]: delegate */


// The sound’s delegate.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSound/delegate
func (s_ Sound) SetDelegate(value unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setDelegate:"), value)
}/* debug [instance_properties/setter]: delegate */


// The duration of the sound, in seconds.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSound/duration
func (s_ Sound) Duration() float64 {
	rv := objc.Send[float64](s_.ID, objc.Sel("duration"))
	return rv
}/* debug [instance_properties/getter]: duration */


// A Boolean that indicates whether the sound is playing its audio data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSound/isPlaying
func (s_ Sound) Playing() bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("playing"))
	return rv
}/* debug [instance_properties/getter]: playing */


// A Boolean that indicates whether the sound restarts playback when it reaches the end of its content.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSound/loops
func (s_ Sound) Loops() bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("loops"))
	return rv
}/* debug [instance_properties/getter]: loops */


// A Boolean that indicates whether the sound restarts playback when it reaches the end of its content.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSound/loops
func (s_ Sound) SetLoops(value bool) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setLoops:"), value)
}/* debug [instance_properties/setter]: loops */


// The name assigned to the sound.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSound/name-swift.property
func (s_ Sound) Name() SoundName /* typedef */ {
	rv := objc.Send[foundation.NSString](s_.ID, objc.Sel("name"))
	return rv
}/* debug [instance_properties/getter]: name */


// Identifies the sound’s output device
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSound/playbackDeviceIdentifier-swift.property
func (s_ Sound) PlaybackDeviceIdentifier() SoundPlaybackDeviceIdentifier /* typedef */ {
	rv := objc.Send[foundation.NSString](s_.ID, objc.Sel("playbackDeviceIdentifier"))
	return rv
}/* debug [instance_properties/getter]: playbackDeviceIdentifier */


// Identifies the sound’s output device
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSound/playbackDeviceIdentifier-swift.property
func (s_ Sound) SetPlaybackDeviceIdentifier(value SoundPlaybackDeviceIdentifier /* typedef */) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setPlaybackDeviceIdentifier:"), value)
}/* debug [instance_properties/setter]: playbackDeviceIdentifier */


// Provides the file types the class understands.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSound/soundUnfilteredTypes
func (s_ Sound) SoundUnfilteredTypes() []string {
	rv := objc.Send[[]string](s_.ID, objc.Sel("soundUnfilteredTypes"))
	return rv
}/* debug [instance_properties/getter]: soundUnfilteredTypes */


// The volume of the sound.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSound/volume
func (s_ Sound) Volume() float32 {
	rv := objc.Send[float32](s_.ID, objc.Sel("volume"))
	return rv
}/* debug [instance_properties/getter]: volume */


// The volume of the sound.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSound/volume
func (s_ Sound) SetVolume(value float32) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setVolume:"), value)
}/* debug [instance_properties/setter]: volume */


// A Boolean that indicates whether the sound is playing its audio data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssound/isplaying
func (s_ Sound) IsPlaying() bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("isPlaying"))
	return rv
}/* debug [instance_properties/getter]: isPlaying */


// A Boolean that indicates whether the sound is playing its audio data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssound/isplaying
func (s_ Sound) SetIsPlaying(value bool) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setIsPlaying:"), value)
}/* debug [instance_properties/setter]: isPlaying */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class NSSound */


