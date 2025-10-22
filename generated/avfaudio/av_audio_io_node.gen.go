// Code generated from Apple documentation for AVFAudio. DO NOT EDIT.

package avfaudio

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

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

// An interface definition for the [AudioIONode] class.
type IAudioIONode interface {
	IAudioNode
	AudioUnit() AudioUnit
	SetAudioUnit(value IAudioUnit)
	IsVoiceProcessingEnabled() bool
	SetIsVoiceProcessingEnabled(value bool)
	PresentationLatency() unsafe.Pointer
	SetPresentationLatency(value unsafe.Pointer)
}

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

// Alloc allocates a new instance without initialization.
func (ac _AudioIONodeClass) Alloc() AudioIONode {
	rv := objc.Send[AudioIONode](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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



// The node’s underlying audio unit, if any.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avaudioionode/audiounit

func (a_ AudioIONode) AudioUnit() AudioUnit {
	rv := objc.Send[AudioUnit](a_.ID, objc.Sel("audioUnit"))
	return rv
}


// The node’s underlying audio unit, if any.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avaudioionode/audiounit

func (a_ AudioIONode) SetAudioUnit(value IAudioUnit) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setAudioUnit:"), value)
}


// A Boolean value that indicates whether voice processing is in an enabled state.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avaudioionode/isvoiceprocessingenabled

func (a_ AudioIONode) IsVoiceProcessingEnabled() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("isVoiceProcessingEnabled"))
	return rv
}


// A Boolean value that indicates whether voice processing is in an enabled state.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avaudioionode/isvoiceprocessingenabled

func (a_ AudioIONode) SetIsVoiceProcessingEnabled(value bool) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setIsVoiceProcessingEnabled:"), value)
}


// The presentation or hardware latency, applicable when rendering to or from an audio device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avaudioionode/presentationlatency

func (a_ AudioIONode) PresentationLatency() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("presentationLatency"))
	return rv
}


// The presentation or hardware latency, applicable when rendering to or from an audio device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avaudioionode/presentationlatency

func (a_ AudioIONode) SetPresentationLatency(value unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setPresentationLatency:"), value)
}



