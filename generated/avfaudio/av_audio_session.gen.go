// Code generated from Apple documentation for AVFAudio. DO NOT EDIT.

package avfaudio

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class AVAudioSession */


/* debug [class_header]: Header for AVAudioSession */
// The class instance for the [AudioSession] class.
var (
	AudioSessionClass     _AudioSessionClass
	AudioSessionClassOnce sync.Once
)

func getAudioSessionClass() _AudioSessionClass {
	AudioSessionClassOnce.Do(func() {
		AudioSessionClass = _AudioSessionClass{objc.GetClass("AVAudioSession")}
	})
	return AudioSessionClass
}

type _AudioSessionClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for AudioSession */
// An interface definition for the [AudioSession] class.
type IAudioSession interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for AudioSession */
	// properties:
	IsOtherAudioPlaying() bool
	SetIsOtherAudioPlaying(value bool)
	IsOutputMuted() bool
	SetIsOutputMuted(value bool)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for AudioSession */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for AudioSession */
// Alloc allocates a new instance without initialization.
func (ac _AudioSessionClass) Alloc() AudioSession {
	rv := objc.Send[AudioSession](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (ac _AudioSessionClass) New() AudioSession {
	rv := objc.Send[AudioSession](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ AudioSession) Init() AudioSession {
	rv := objc.Send[AudioSession](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ AudioSession) Autorelease() AudioSession {
	rv := objc.Send[AudioSession](a_.ID, objc.Sel("autorelease"))
	return rv
}

// NewAudioSession creates a new AudioSession instance.
func NewAudioSession() AudioSession {
	return getAudioSessionClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for AudioSession */
// An object that communicates to the system how you intend to use audio in your app.
//
// An audio session acts as an intermediary between your app and the operating system — and, in turn, the underlying audio hardware. You use an audio session to communicate to the operating system the general nature of your app’s audio without detailing the specific behavior or required interactions with the audio hardware. You delegate the management of those details to the audio session, which ensures that the operating system can best manage the user’s audio experience. All iOS, tvOS, and watchOS apps have a default audio session that comes preconfigured with the following behavior: It supports audio playback, but disallows audio recording. When the app plays audio, it silences any other background audio. In iOS, setting the Ring/Silent switch to silent mode silences any audio the app is playing. In iOS, locking a device silences the app’s audio. Although the default audio session provides useful behavior, it generally doesn’t provide the audio behavior a media app needs. To change the default behavior, you configure your app’s audio session category. There are six possible categories you can use, but is the one that playback apps most commonly use. This category indicates that audio playback is a central feature of your app. When you specify this category, your app’s audio continues with the Ring/Silent switch set to silent mode (iOS only). Using this category, you can also play background audio if you’re using the Audio, AirPlay, and Picture in Picture background mode. For more information, see . You use an object to configure your app’s audio session. This class is a singleton object used to set the audio session’s category, mode, and other configurations. You can interact with the audio session throughout your app’s life cycle, but it’s often useful to perform this configuration at app launch, as shown in the following example. The audio session uses this configuration when you activate the session using the or method.


// An object that communicates to the system how you intend to use audio in your app.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioSession
type AudioSession struct {
	objectivec.Object
}

// AudioSessionFrom constructs a [AudioSession] from an unsafe.Pointer.
//
// An object that communicates to the system how you intend to use audio in your app.
func AudioSessionFrom(ptr unsafe.Pointer) AudioSession {
	return AudioSession{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for AudioSession */
/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for AudioSession */

// Returns the shared audio session instance.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioSession/sharedInstance()
func (ac _AudioSessionClass) SharedInstance() IAudioSession {
	rv := objc.Send[AudioSession](objc.ID(ac.class), objc.Sel("sharedInstance"))
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=SharedInstance) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for AudioSession */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for AudioSession */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for AudioSession */

// A Boolean value that indicates whether another app is playing audio.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avaudiosession/isotheraudioplaying
func (a_ AudioSession) IsOtherAudioPlaying() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("isOtherAudioPlaying"))
	return rv
}/* debug [instance_properties/getter]: isOtherAudioPlaying */


// A Boolean value that indicates whether another app is playing audio.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avaudiosession/isotheraudioplaying
func (a_ AudioSession) SetIsOtherAudioPlaying(value bool) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setIsOtherAudioPlaying:"), value)
}/* debug [instance_properties/setter]: isOtherAudioPlaying */


// A Boolean value that indicates whether audio output is in a muted state.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avaudiosession/isoutputmuted
func (a_ AudioSession) IsOutputMuted() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("isOutputMuted"))
	return rv
}/* debug [instance_properties/getter]: isOutputMuted */


// A Boolean value that indicates whether audio output is in a muted state.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avaudiosession/isoutputmuted
func (a_ AudioSession) SetIsOutputMuted(value bool) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setIsOutputMuted:"), value)
}/* debug [instance_properties/setter]: isOutputMuted */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class AVAudioSession */


