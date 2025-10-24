// Code generated from Apple documentation for AVFAudio. DO NOT EDIT.

package avfaudio

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class AVAudioIONode */


/* debug [class_header]: Header for AVAudioIONode */
// The class instance for the [AudioIONode] class.
var (
	AudioIONodeClass     _AudioIONodeClass
	AudioIONodeClassOnce sync.Once
)

func getAudioIONodeClass() _AudioIONodeClass {
	AudioIONodeClassOnce.Do(func() {
		AudioIONodeClass = _AudioIONodeClass{objc.GetClass("AVAudioIONode")}
	})
	return AudioIONodeClass
}

type _AudioIONodeClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for AudioIONode */
// An interface definition for the [AudioIONode] class.
type IAudioIONode interface {
	IAudioNode
	
/* debug [class_interface_properties]: Properties for AudioIONode */
	// properties:
	AudioUnit() audiotoolbox.AudioUnit
	VoiceProcessingEnabled() bool
	PresentationLatency() float64
	IsVoiceProcessingEnabled() bool
	SetIsVoiceProcessingEnabled(value bool)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for AudioIONode */
	// methods:
	SetVoiceProcessingEnabledError(enabled bool, outError objectivec.IObject) bool
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for AudioIONode */
// Alloc allocates a new instance without initialization.
func (ac _AudioIONodeClass) Alloc() AudioIONode {
	rv := objc.Send[AudioIONode](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (ac _AudioIONodeClass) New() AudioIONode {
	rv := objc.Send[AudioIONode](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ AudioIONode) Init() AudioIONode {
	rv := objc.Send[AudioIONode](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ AudioIONode) Autorelease() AudioIONode {
	rv := objc.Send[AudioIONode](a_.ID, objc.Sel("autorelease"))
	return rv
}

// NewAudioIONode creates a new AudioIONode instance.
func NewAudioIONode() AudioIONode {
	return getAudioIONodeClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for AudioIONode */
// An object that performs audio input or output in the engine.
//
// When rendering to and from an audio device in macOS, and communicate with the system’s default input and output devices. In iOS, they communicate with the devices appropriate to the app’s category, configurations, and user actions, such as connecting or disconnecting external devices. In the manual rendering mode, and perform the input and output in the engine in response to the client’s request.


// An object that performs audio input or output in the engine.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioIONode
type AudioIONode struct {
	AudioNode
}

// AudioIONodeFrom constructs a [AudioIONode] from an unsafe.Pointer.
//
// An object that performs audio input or output in the engine.
func AudioIONodeFrom(ptr unsafe.Pointer) AudioIONode {
	return AudioIONode{
		AudioNode: AudioNodeFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for AudioIONode *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for AudioIONode */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for AudioIONode */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for AudioIONode */

// Enables or disables voice processing on the I/O node.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioIONode/setVoiceProcessingEnabled(_:)
func (a_ AudioIONode) SetVoiceProcessingEnabledError(enabled bool, outError objectivec.IObject) bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("setVoiceProcessingEnabled:error:"), enabled, outError)
	return rv
}/* debug [instance_methods/method]: SetVoiceProcessingEnabledError */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for AudioIONode */

// The node’s underlying audio unit, if any.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioIONode/audioUnit
func (a_ AudioIONode) AudioUnit() audiotoolbox.AudioUnit {
	rv := objc.Send[audiotoolbox.AudioUnit](a_.ID, objc.Sel("audioUnit"))
	return rv
}/* debug [instance_properties/getter]: audioUnit */


// A Boolean value that indicates whether voice processing is in an enabled state.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioIONode/isVoiceProcessingEnabled
func (a_ AudioIONode) VoiceProcessingEnabled() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("voiceProcessingEnabled"))
	return rv
}/* debug [instance_properties/getter]: voiceProcessingEnabled */


// The presentation or hardware latency, applicable when rendering to or from an audio device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioIONode/presentationLatency
func (a_ AudioIONode) PresentationLatency() float64 {
	rv := objc.Send[float64](a_.ID, objc.Sel("presentationLatency"))
	return rv
}/* debug [instance_properties/getter]: presentationLatency */


// A Boolean value that indicates whether voice processing is in an enabled state.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avaudioionode/isvoiceprocessingenabled
func (a_ AudioIONode) IsVoiceProcessingEnabled() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("isVoiceProcessingEnabled"))
	return rv
}/* debug [instance_properties/getter]: isVoiceProcessingEnabled */


// A Boolean value that indicates whether voice processing is in an enabled state.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avaudioionode/isvoiceprocessingenabled
func (a_ AudioIONode) SetIsVoiceProcessingEnabled(value bool) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setIsVoiceProcessingEnabled:"), value)
}/* debug [instance_properties/setter]: isVoiceProcessingEnabled */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class AVAudioIONode */



