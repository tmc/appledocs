// Code generated from Apple documentation for AVFAudio. DO NOT EDIT.

package avfaudio

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)





// The class instance for the [AudioInputNode] class.
var (
	AudioInputNodeClass     _AudioInputNodeClass
	AudioInputNodeClassOnce sync.Once
)

func getAudioInputNodeClass() _AudioInputNodeClass {
	AudioInputNodeClassOnce.Do(func() {
		AudioInputNodeClass = _AudioInputNodeClass{objc.GetClass("AVAudioInputNode")}
	})
	return AudioInputNodeClass
}

type _AudioInputNodeClass struct {
	class objc.Class
}





// An interface definition for the [AudioInputNode] class.
type IAudioInputNode interface {
	IAudioIONode
	

	// properties:
	VoiceProcessingAGCEnabled() bool
	SetVoiceProcessingAGCEnabled(value bool)
	VoiceProcessingBypassed() bool
	SetVoiceProcessingBypassed(value bool)
	VoiceProcessingInputMuted() bool
	SetVoiceProcessingInputMuted(value bool)
	VoiceProcessingOtherAudioDuckingConfiguration() objc.IObject /* cross-framework: AVAudioVoiceProcessingOtherAudioDuckingConfiguration */
	SetVoiceProcessingOtherAudioDuckingConfiguration(value objc.IObject /* cross-framework: AVAudioVoiceProcessingOtherAudioDuckingConfiguration */)
	IsVoiceProcessingAGCEnabled() bool
	SetIsVoiceProcessingAGCEnabled(value bool)
	IsVoiceProcessingBypassed() bool
	SetIsVoiceProcessingBypassed(value bool)
	IsVoiceProcessingInputMuted() bool
	SetIsVoiceProcessingInputMuted(value bool)


	

	// methods:
	SetManualRenderingInputPCMFormatInputBlock(format IAVAudioFormat, block AudioIONodeInputBlock /* not a class type */) bool
	SetMutedSpeechActivityEventListener(listenerBlock unsafe.Pointer) bool


}





// Alloc allocates a new instance without initialization.
func (ac _AudioInputNodeClass) Alloc() AudioInputNode {
	rv := objc.Send[AudioInputNode](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (ac _AudioInputNodeClass) New() AudioInputNode {
	rv := objc.Send[AudioInputNode](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ AudioInputNode) Init() AudioInputNode {
	rv := objc.Send[AudioInputNode](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ AudioInputNode) Autorelease() AudioInputNode {
	rv := objc.Send[AudioInputNode](a_.ID, objc.Sel("autorelease"))
	return rv
}

// NewAudioInputNode creates a new AudioInputNode instance.
func NewAudioInputNode() AudioInputNode {
	return getAudioInputNodeClass().New()
}





// An object that connects to the system’s audio input.
//
// This node connects to the system’s audio input when rendering to or from an audio device. In manual rendering mode, this node supplies input data to the engine. This audio node has one element. The format of the input scope reflects: The audio hardware sample rate and channel count when it connects to hardware. The format of the PCM audio data that the node supplies to the engine in manual rendering mode. For more information, see When rendering from an audio device, the input node doesn’t support format conversion. In this case, the format of the output scope must be the same as the input and the formats for all nodes connected to the input chain. In manual rendering mode, the format of the output scope is initially the same as the input, but you may set it to a different format, which converts the node.


// An object that connects to the system’s audio input.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioInputNode
type AudioInputNode struct {
	AudioIONode
}

// AudioInputNodeFrom constructs a [AudioInputNode] from an unsafe.Pointer.
//
// An object that connects to the system’s audio input.
func AudioInputNodeFrom(ptr unsafe.Pointer) AudioInputNode {
	return AudioInputNode{
		AudioIONode: AudioIONodeFrom(ptr),
	}
}




















// Supplies the data through the input node to the engine while operating in the manual rendering mode.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioInputNode/setManualRenderingInputPCMFormat(_:inputBlock:)
func (a_ AudioInputNode) SetManualRenderingInputPCMFormatInputBlock(format IAVAudioFormat, block AudioIONodeInputBlock /* not a class type */) bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("setManualRenderingInputPCMFormat:inputBlock:"), format, block)
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioInputNode/setMutedSpeechActivityEventListener(_:)
func (a_ AudioInputNode) SetMutedSpeechActivityEventListener(listenerBlock unsafe.Pointer) bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("setMutedSpeechActivityEventListener:"), listenerBlock)
	return rv
}







// A Boolean that indicates whether automatic gain control on the processed microphone uplink signal is active.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioInputNode/isVoiceProcessingAGCEnabled
func (a_ AudioInputNode) VoiceProcessingAGCEnabled() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("voiceProcessingAGCEnabled"))
	return rv
}


// A Boolean that indicates whether automatic gain control on the processed microphone uplink signal is active.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioInputNode/isVoiceProcessingAGCEnabled
func (a_ AudioInputNode) SetVoiceProcessingAGCEnabled(value bool) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setVoiceProcessingAGCEnabled:"), value)
}


// A Boolean that indicates whether the node bypasses all microphone uplink processing of the voice-processing unit.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioInputNode/isVoiceProcessingBypassed
func (a_ AudioInputNode) VoiceProcessingBypassed() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("voiceProcessingBypassed"))
	return rv
}


// A Boolean that indicates whether the node bypasses all microphone uplink processing of the voice-processing unit.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioInputNode/isVoiceProcessingBypassed
func (a_ AudioInputNode) SetVoiceProcessingBypassed(value bool) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setVoiceProcessingBypassed:"), value)
}


// A Boolean that indicates whether the input of the voice processing unit is in a muted state.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioInputNode/isVoiceProcessingInputMuted
func (a_ AudioInputNode) VoiceProcessingInputMuted() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("voiceProcessingInputMuted"))
	return rv
}


// A Boolean that indicates whether the input of the voice processing unit is in a muted state.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioInputNode/isVoiceProcessingInputMuted
func (a_ AudioInputNode) SetVoiceProcessingInputMuted(value bool) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setVoiceProcessingInputMuted:"), value)
}


// The ducking configuration of nonvoice audio.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioInputNode/voiceProcessingOtherAudioDuckingConfiguration
func (a_ AudioInputNode) VoiceProcessingOtherAudioDuckingConfiguration() objc.IObject /* cross-framework: AVAudioVoiceProcessingOtherAudioDuckingConfiguration */ {
	rv := objc.Send[objc.ID](a_.ID, objc.Sel("voiceProcessingOtherAudioDuckingConfiguration"))
	return rv
}


// The ducking configuration of nonvoice audio.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioInputNode/voiceProcessingOtherAudioDuckingConfiguration
func (a_ AudioInputNode) SetVoiceProcessingOtherAudioDuckingConfiguration(value objc.IObject /* cross-framework: AVAudioVoiceProcessingOtherAudioDuckingConfiguration */) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setVoiceProcessingOtherAudioDuckingConfiguration:"), value)
}


// A Boolean that indicates whether automatic gain control on the processed microphone uplink signal is active.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avaudioinputnode/isvoiceprocessingagcenabled
func (a_ AudioInputNode) IsVoiceProcessingAGCEnabled() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("isVoiceProcessingAGCEnabled"))
	return rv
}


// A Boolean that indicates whether automatic gain control on the processed microphone uplink signal is active.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avaudioinputnode/isvoiceprocessingagcenabled
func (a_ AudioInputNode) SetIsVoiceProcessingAGCEnabled(value bool) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setIsVoiceProcessingAGCEnabled:"), value)
}


// A Boolean that indicates whether the node bypasses all microphone uplink processing of the voice-processing unit.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avaudioinputnode/isvoiceprocessingbypassed
func (a_ AudioInputNode) IsVoiceProcessingBypassed() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("isVoiceProcessingBypassed"))
	return rv
}


// A Boolean that indicates whether the node bypasses all microphone uplink processing of the voice-processing unit.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avaudioinputnode/isvoiceprocessingbypassed
func (a_ AudioInputNode) SetIsVoiceProcessingBypassed(value bool) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setIsVoiceProcessingBypassed:"), value)
}


// A Boolean that indicates whether the input of the voice processing unit is in a muted state.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avaudioinputnode/isvoiceprocessinginputmuted
func (a_ AudioInputNode) IsVoiceProcessingInputMuted() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("isVoiceProcessingInputMuted"))
	return rv
}


// A Boolean that indicates whether the input of the voice processing unit is in a muted state.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avaudioinputnode/isvoiceprocessinginputmuted
func (a_ AudioInputNode) SetIsVoiceProcessingInputMuted(value bool) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setIsVoiceProcessingInputMuted:"), value)
}








