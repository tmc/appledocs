// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [AudioMixInputParameters] class.
var (
	AudioMixInputParametersClass     _AudioMixInputParametersClass
	AudioMixInputParametersClassOnce sync.Once
)

func getAudioMixInputParametersClass() _AudioMixInputParametersClass {
	AudioMixInputParametersClassOnce.Do(func() {
		AudioMixInputParametersClass = _AudioMixInputParametersClass{objc.GetClass("AVAudioMixInputParameters")}
	})
	return AudioMixInputParametersClass
}

type _AudioMixInputParametersClass struct {
	class objc.Class
}

// An interface definition for the [AudioMixInputParameters] class.
type IAudioMixInputParameters interface {
	objectivec.IObject
	AudioTapProcessor() unsafe.Pointer
	SetAudioTapProcessor(value unsafe.Pointer)
	AudioTimePitchAlgorithm() AudioTimePitchAlgorithm
	SetAudioTimePitchAlgorithm(value IAudioTimePitchAlgorithm)
	TrackID() unsafe.Pointer
	SetTrackID(value unsafe.Pointer)
}

// An object that represents the parameters that you apply when adding an audio track to a mix.
//
// You use an instance to apply audio volume ramps for an input to an audio mix. Mix parameters are associated with audio tracks via the property. Audio volume is currently supported as a time-varying parameter. has a mutable subclass, . Before the first time at which a volume is set, a volume of 1.0 used; after the last time for which a volume has been set, the last volume is used. Within the time range of a volume ramp, the volume is interpolated between the start volume and end volume of the ramp. For example, setting the volume to 1.0 at time 0 and also setting a volume ramp from a volume of 0.5 to 0.2 with a timeRange of [4.0, 5.0] results in an audio volume parameters that hold the volume constant at 1.0 from 0.0 sec to 4.0 sec, then cause it to jump to 0.5 and descend to 0.2 from 4.0 sec to 9.0 sec, holding constant at 0.2 thereafter. Given that this is an immutable variant of the object, you should not allocate and initialize a version of this class yourself. Other classes may return instances of this class.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAudioMixInputParameters
type AudioMixInputParameters struct {
	objectivec.Object
}

// AudioMixInputParametersFrom constructs a [AudioMixInputParameters] from an unsafe.Pointer.
//
// An object that represents the parameters that you apply when adding an audio track to a mix.
func AudioMixInputParametersFrom(ptr unsafe.Pointer) AudioMixInputParameters {
	return AudioMixInputParameters{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (ac _AudioMixInputParametersClass) Alloc() AudioMixInputParameters {
	rv := objc.Send[AudioMixInputParameters](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ac _AudioMixInputParametersClass) New() AudioMixInputParameters {
	rv := objc.Send[AudioMixInputParameters](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ AudioMixInputParameters) Init() AudioMixInputParameters {
	rv := objc.Send[AudioMixInputParameters](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ AudioMixInputParameters) Autorelease() AudioMixInputParameters {
	rv := objc.Send[AudioMixInputParameters](a_.ID, objc.Sel("autorelease"))
	return rv
}

// NewAudioMixInputParameters creates a new AudioMixInputParameters instance.
func NewAudioMixInputParameters() AudioMixInputParameters {
	return getAudioMixInputParametersClass().New()
}


// The audio processing tap associated with the track.
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avaudiomixinputparameters/audiotapprocessor
func (a_ AudioMixInputParameters) AudioTapProcessor() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("audioTapProcessor"))
	return rv
}


// SetAudioTapProcessor sets the value of the audioTapProcessor property.
// The audio processing tap associated with the track.

//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avaudiomixinputparameters/audiotapprocessor
func (a_ AudioMixInputParameters) SetAudioTapProcessor(value unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setAudioTapProcessor:"), value)
}

// The processing algorithm used to manage audio pitch for scaled audio edits.
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avaudiomixinputparameters/audiotimepitchalgorithm
func (a_ AudioMixInputParameters) AudioTimePitchAlgorithm() AudioTimePitchAlgorithm {
	rv := objc.Send[AudioTimePitchAlgorithm](a_.ID, objc.Sel("audioTimePitchAlgorithm"))
	return rv
}


// SetAudioTimePitchAlgorithm sets the value of the audioTimePitchAlgorithm property.
// The processing algorithm used to manage audio pitch for scaled audio edits.

//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avaudiomixinputparameters/audiotimepitchalgorithm
func (a_ AudioMixInputParameters) SetAudioTimePitchAlgorithm(value IAudioTimePitchAlgorithm) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setAudioTimePitchAlgorithm:"), value)
}

// The identifier of the audio track to which the parameters should be applied.
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avaudiomixinputparameters/trackid
func (a_ AudioMixInputParameters) TrackID() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("trackID"))
	return rv
}


// SetTrackID sets the value of the trackID property.
// The identifier of the audio track to which the parameters should be applied.

//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avaudiomixinputparameters/trackid
func (a_ AudioMixInputParameters) SetTrackID(value unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setTrackID:"), value)
}



