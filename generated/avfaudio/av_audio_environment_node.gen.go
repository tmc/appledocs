// Code generated from Apple documentation for AVFAudio. DO NOT EDIT.

package avfaudio

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [AudioEnvironmentNode] class.
var (
	AudioEnvironmentNodeClass     _AudioEnvironmentNodeClass
	AudioEnvironmentNodeClassOnce sync.Once
)

func getAudioEnvironmentNodeClass() _AudioEnvironmentNodeClass {
	AudioEnvironmentNodeClassOnce.Do(func() {
		AudioEnvironmentNodeClass = _AudioEnvironmentNodeClass{objc.GetClass("AVAudioEnvironmentNode")}
	})
	return AudioEnvironmentNodeClass
}

type _AudioEnvironmentNodeClass struct {
	class objc.Class
}

// An interface definition for the [AudioEnvironmentNode] class.
type IAudioEnvironmentNode interface {
	IAudioNode
	ApplicableRenderingAlgorithms() []foundation.Number
	ListenerHeadTrackingEnabled() bool
	SetListenerHeadTrackingEnabled(value bool)
	ListenerAngularOrientation() unsafe.Pointer
	SetListenerAngularOrientation(value unsafe.Pointer)
	DistanceAttenuationParameters() AVAudioEnvironmentDistanceAttenuationParameters
	SetDistanceAttenuationParameters(value IAVAudioEnvironmentDistanceAttenuationParameters)
	IsListenerHeadTrackingEnabled() bool
	SetIsListenerHeadTrackingEnabled(value bool)
	ListenerPosition() unsafe.Pointer
	SetListenerPosition(value unsafe.Pointer)
	ListenerVectorOrientation() unsafe.Pointer
	SetListenerVectorOrientation(value unsafe.Pointer)
	NextAvailableInputBus() AudioNodeBus
	SetNextAvailableInputBus(value IAudioNodeBus)
	OutputType() AudioEnvironmentOutputType
	SetOutputType(value AudioEnvironmentOutputType)
	OutputVolume() float32
	SetOutputVolume(value float32)
	ReverbParameters() AVAudioEnvironmentReverbParameters
	SetReverbParameters(value IAVAudioEnvironmentReverbParameters)
	KAudioChannelLayoutTag_AudioUnit_4() unsafe.Pointer
	SetKAudioChannelLayoutTag_AudioUnit_4(value unsafe.Pointer)
	KAudioChannelLayoutTag_AudioUnit_5_0() unsafe.Pointer
	SetKAudioChannelLayoutTag_AudioUnit_5_0(value unsafe.Pointer)
	KAudioChannelLayoutTag_AudioUnit_6_0() unsafe.Pointer
	SetKAudioChannelLayoutTag_AudioUnit_6_0(value unsafe.Pointer)
	KAudioChannelLayoutTag_AudioUnit_7_0() unsafe.Pointer
	SetKAudioChannelLayoutTag_AudioUnit_7_0(value unsafe.Pointer)
	KAudioChannelLayoutTag_AudioUnit_7_0_Front() unsafe.Pointer
	SetKAudioChannelLayoutTag_AudioUnit_7_0_Front(value unsafe.Pointer)
	KAudioChannelLayoutTag_AudioUnit_8() unsafe.Pointer
	SetKAudioChannelLayoutTag_AudioUnit_8(value unsafe.Pointer)
}

// An object that simulates a 3D audio environment.
//
// The class is a mixer node that simulates a 3D audio environment. Any node that conforms to can act as a source node, such as . The environment node has an implicit listener. You set the listener’s position and orientation, and the system then controls the way the user experiences the virtual world. To help characterize the environment, this class defines properties for distance attenuation and reverberation. affects how inputs with different channel configurations render. Spatialization applies only to inputs with a mono channel connection format. This class doesn’t spatialize stereo inputs or support inputs with connection formats of more than two channels. To set the node’s output to a multichannel format, use an that has one of the following :


// An object that simulates a 3D audio environment.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioEnvironmentNode

type AudioEnvironmentNode struct {
	AudioNode
}

// AudioEnvironmentNodeFrom constructs a [AudioEnvironmentNode] from an unsafe.Pointer.
//
// An object that simulates a 3D audio environment.
func AudioEnvironmentNodeFrom(ptr unsafe.Pointer) AudioEnvironmentNode {
	return AudioEnvironmentNode{
		AudioNode: AudioNodeFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (ac _AudioEnvironmentNodeClass) Alloc() AudioEnvironmentNode {
	rv := objc.Send[AudioEnvironmentNode](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ac _AudioEnvironmentNodeClass) New() AudioEnvironmentNode {
	rv := objc.Send[AudioEnvironmentNode](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ AudioEnvironmentNode) Init() AudioEnvironmentNode {
	rv := objc.Send[AudioEnvironmentNode](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ AudioEnvironmentNode) Autorelease() AudioEnvironmentNode {
	rv := objc.Send[AudioEnvironmentNode](a_.ID, objc.Sel("autorelease"))
	return rv
}

// NewAudioEnvironmentNode creates a new AudioEnvironmentNode instance.
func NewAudioEnvironmentNode() AudioEnvironmentNode {
	return getAudioEnvironmentNodeClass().New()
}




// An array of rendering algorithms applicable to the environment node.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioEnvironmentNode/applicableRenderingAlgorithms

func (a_ AudioEnvironmentNode) ApplicableRenderingAlgorithms() []foundation.Number {
	rv := objc.Send[[]foundation.Number](a_.ID, objc.Sel("applicableRenderingAlgorithms"))
	return rv
}


// A Boolean value that indicates whether the listener orientation is automatically rotated based on head orientation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioEnvironmentNode/isListenerHeadTrackingEnabled

func (a_ AudioEnvironmentNode) ListenerHeadTrackingEnabled() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("listenerHeadTrackingEnabled"))
	return rv
}


// A Boolean value that indicates whether the listener orientation is automatically rotated based on head orientation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioEnvironmentNode/isListenerHeadTrackingEnabled

func (a_ AudioEnvironmentNode) SetListenerHeadTrackingEnabled(value bool) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setListenerHeadTrackingEnabled:"), value)
}


// The listener’s angular orientation in the environment.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioEnvironmentNode/listenerAngularOrientation

func (a_ AudioEnvironmentNode) ListenerAngularOrientation() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("listenerAngularOrientation"))
	return rv
}


// The listener’s angular orientation in the environment.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioEnvironmentNode/listenerAngularOrientation

func (a_ AudioEnvironmentNode) SetListenerAngularOrientation(value unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setListenerAngularOrientation:"), value)
}


// The distance attenuation parameters for the environment.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avaudioenvironmentnode/distanceattenuationparameters

func (a_ AudioEnvironmentNode) DistanceAttenuationParameters() AVAudioEnvironmentDistanceAttenuationParameters {
	rv := objc.Send[AVAudioEnvironmentDistanceAttenuationParameters](a_.ID, objc.Sel("distanceAttenuationParameters"))
	return rv
}


// The distance attenuation parameters for the environment.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avaudioenvironmentnode/distanceattenuationparameters

func (a_ AudioEnvironmentNode) SetDistanceAttenuationParameters(value IAVAudioEnvironmentDistanceAttenuationParameters) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setDistanceAttenuationParameters:"), value)
}


// A Boolean value that indicates whether the listener orientation is automatically rotated based on head orientation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avaudioenvironmentnode/islistenerheadtrackingenabled

func (a_ AudioEnvironmentNode) IsListenerHeadTrackingEnabled() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("isListenerHeadTrackingEnabled"))
	return rv
}


// A Boolean value that indicates whether the listener orientation is automatically rotated based on head orientation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avaudioenvironmentnode/islistenerheadtrackingenabled

func (a_ AudioEnvironmentNode) SetIsListenerHeadTrackingEnabled(value bool) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setIsListenerHeadTrackingEnabled:"), value)
}


// The listener’s position in the 3D environment.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avaudioenvironmentnode/listenerposition

func (a_ AudioEnvironmentNode) ListenerPosition() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("listenerPosition"))
	return rv
}


// The listener’s position in the 3D environment.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avaudioenvironmentnode/listenerposition

func (a_ AudioEnvironmentNode) SetListenerPosition(value unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setListenerPosition:"), value)
}


// The listener’s vector orientation in the environment.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avaudioenvironmentnode/listenervectororientation

func (a_ AudioEnvironmentNode) ListenerVectorOrientation() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("listenerVectorOrientation"))
	return rv
}


// The listener’s vector orientation in the environment.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avaudioenvironmentnode/listenervectororientation

func (a_ AudioEnvironmentNode) SetListenerVectorOrientation(value unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setListenerVectorOrientation:"), value)
}


// An unused input bus.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avaudioenvironmentnode/nextavailableinputbus

func (a_ AudioEnvironmentNode) NextAvailableInputBus() AudioNodeBus {
	rv := objc.Send[AudioNodeBus](a_.ID, objc.Sel("nextAvailableInputBus"))
	return rv
}


// An unused input bus.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avaudioenvironmentnode/nextavailableinputbus

func (a_ AudioEnvironmentNode) SetNextAvailableInputBus(value IAudioNodeBus) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setNextAvailableInputBus:"), value)
}


// The type of output hardware.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avaudioenvironmentnode/outputtype

func (a_ AudioEnvironmentNode) OutputType() AudioEnvironmentOutputType {
	rv := objc.Send[AudioEnvironmentOutputType](a_.ID, objc.Sel("outputType"))
	return rv
}


// The type of output hardware.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avaudioenvironmentnode/outputtype

func (a_ AudioEnvironmentNode) SetOutputType(value AudioEnvironmentOutputType) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setOutputType:"), value)
}


// The mixer’s output volume.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avaudioenvironmentnode/outputvolume

func (a_ AudioEnvironmentNode) OutputVolume() float32 {
	rv := objc.Send[float32](a_.ID, objc.Sel("outputVolume"))
	return rv
}


// The mixer’s output volume.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avaudioenvironmentnode/outputvolume

func (a_ AudioEnvironmentNode) SetOutputVolume(value float32) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setOutputVolume:"), value)
}


// The reverb parameters for the environment.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avaudioenvironmentnode/reverbparameters

func (a_ AudioEnvironmentNode) ReverbParameters() AVAudioEnvironmentReverbParameters {
	rv := objc.Send[AVAudioEnvironmentReverbParameters](a_.ID, objc.Sel("reverbParameters"))
	return rv
}


// The reverb parameters for the environment.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avaudioenvironmentnode/reverbparameters

func (a_ AudioEnvironmentNode) SetReverbParameters(value IAVAudioEnvironmentReverbParameters) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setReverbParameters:"), value)
}


// A quadraphonic symmetrical layout, recommended for use by audio units.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreAudioTypes/kAudioChannelLayoutTag_AudioUnit_4

func (a_ AudioEnvironmentNode) KAudioChannelLayoutTag_AudioUnit_4() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("kAudioChannelLayoutTag_AudioUnit_4"))
	return rv
}


// A quadraphonic symmetrical layout, recommended for use by audio units.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreAudioTypes/kAudioChannelLayoutTag_AudioUnit_4

func (a_ AudioEnvironmentNode) SetKAudioChannelLayoutTag_AudioUnit_4(value unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setKAudioChannelLayoutTag_AudioUnit_4:"), value)
}


// A 5-channel surround-based layout, recommended for use by audio units.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreAudioTypes/kAudioChannelLayoutTag_AudioUnit_5_0

func (a_ AudioEnvironmentNode) KAudioChannelLayoutTag_AudioUnit_5_0() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("kAudioChannelLayoutTag_AudioUnit_5_0"))
	return rv
}


// A 5-channel surround-based layout, recommended for use by audio units.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreAudioTypes/kAudioChannelLayoutTag_AudioUnit_5_0

func (a_ AudioEnvironmentNode) SetKAudioChannelLayoutTag_AudioUnit_5_0(value unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setKAudioChannelLayoutTag_AudioUnit_5_0:"), value)
}


// A 6-channel surround-based layout, recommended for use by audio units.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreAudioTypes/kAudioChannelLayoutTag_AudioUnit_6_0

func (a_ AudioEnvironmentNode) KAudioChannelLayoutTag_AudioUnit_6_0() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("kAudioChannelLayoutTag_AudioUnit_6_0"))
	return rv
}


// A 6-channel surround-based layout, recommended for use by audio units.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreAudioTypes/kAudioChannelLayoutTag_AudioUnit_6_0

func (a_ AudioEnvironmentNode) SetKAudioChannelLayoutTag_AudioUnit_6_0(value unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setKAudioChannelLayoutTag_AudioUnit_6_0:"), value)
}


// A 7-channel surround-based layout, recommended for use by audio units.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreAudioTypes/kAudioChannelLayoutTag_AudioUnit_7_0

func (a_ AudioEnvironmentNode) KAudioChannelLayoutTag_AudioUnit_7_0() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("kAudioChannelLayoutTag_AudioUnit_7_0"))
	return rv
}


// A 7-channel surround-based layout, recommended for use by audio units.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreAudioTypes/kAudioChannelLayoutTag_AudioUnit_7_0

func (a_ AudioEnvironmentNode) SetKAudioChannelLayoutTag_AudioUnit_7_0(value unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setKAudioChannelLayoutTag_AudioUnit_7_0:"), value)
}


// An alternate 7-channel surround-based layout, for use by audio units.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreAudioTypes/kAudioChannelLayoutTag_AudioUnit_7_0_Front

func (a_ AudioEnvironmentNode) KAudioChannelLayoutTag_AudioUnit_7_0_Front() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("kAudioChannelLayoutTag_AudioUnit_7_0_Front"))
	return rv
}


// An alternate 7-channel surround-based layout, for use by audio units.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreAudioTypes/kAudioChannelLayoutTag_AudioUnit_7_0_Front

func (a_ AudioEnvironmentNode) SetKAudioChannelLayoutTag_AudioUnit_7_0_Front(value unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setKAudioChannelLayoutTag_AudioUnit_7_0_Front:"), value)
}


// An octagonal symmetrical layout, recommended for use by audio units.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreAudioTypes/kAudioChannelLayoutTag_AudioUnit_8

func (a_ AudioEnvironmentNode) KAudioChannelLayoutTag_AudioUnit_8() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("kAudioChannelLayoutTag_AudioUnit_8"))
	return rv
}


// An octagonal symmetrical layout, recommended for use by audio units.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreAudioTypes/kAudioChannelLayoutTag_AudioUnit_8

func (a_ AudioEnvironmentNode) SetKAudioChannelLayoutTag_AudioUnit_8(value unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setKAudioChannelLayoutTag_AudioUnit_8:"), value)
}


