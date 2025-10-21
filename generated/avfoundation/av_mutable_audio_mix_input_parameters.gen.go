// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [MutableAudioMixInputParameters] class.
var (
	MutableAudioMixInputParametersClass     _MutableAudioMixInputParametersClass
	MutableAudioMixInputParametersClassOnce sync.Once
)

func getMutableAudioMixInputParametersClass() _MutableAudioMixInputParametersClass {
	MutableAudioMixInputParametersClassOnce.Do(func() {
		MutableAudioMixInputParametersClass = _MutableAudioMixInputParametersClass{objc.GetClass("AVMutableAudioMixInputParameters")}
	})
	return MutableAudioMixInputParametersClass
}

type _MutableAudioMixInputParametersClass struct {
	class objc.Class
}

// An interface definition for the [MutableAudioMixInputParameters] class.
type IMutableAudioMixInputParameters interface {
	IAudioMixInputParameters
}

// The parameters you use when adding an audio track to a mix.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMutableAudioMixInputParameters
type MutableAudioMixInputParameters struct {
	AudioMixInputParameters
}

// MutableAudioMixInputParametersFrom constructs a [MutableAudioMixInputParameters] from an unsafe.Pointer.
//
// The parameters you use when adding an audio track to a mix.
func MutableAudioMixInputParametersFrom(ptr unsafe.Pointer) MutableAudioMixInputParameters {
	return MutableAudioMixInputParameters{
		AudioMixInputParameters: AudioMixInputParametersFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (mc _MutableAudioMixInputParametersClass) Alloc() MutableAudioMixInputParameters {
	rv := objc.Send[MutableAudioMixInputParameters](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MutableAudioMixInputParametersClass) New() MutableAudioMixInputParameters {
	rv := objc.Send[MutableAudioMixInputParameters](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MutableAudioMixInputParameters) Init() MutableAudioMixInputParameters {
	rv := objc.Send[MutableAudioMixInputParameters](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MutableAudioMixInputParameters) Autorelease() MutableAudioMixInputParameters {
	rv := objc.Send[MutableAudioMixInputParameters](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMutableAudioMixInputParameters creates a new MutableAudioMixInputParameters instance.
func NewMutableAudioMixInputParameters() MutableAudioMixInputParameters {
	return getMutableAudioMixInputParametersClass().New()
}


// The audio processing tap associated with the track.
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avmutableaudiomixinputparameters/audiotapprocessor
func (m_ MutableAudioMixInputParameters) AudioTapProcessor() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("audioTapProcessor"))
	return rv
}


// SetAudioTapProcessor sets the value of the audioTapProcessor property.
// The audio processing tap associated with the track.

//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avmutableaudiomixinputparameters/audiotapprocessor
func (m_ MutableAudioMixInputParameters) SetAudioTapProcessor(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setAudioTapProcessor:"), value)
}

// The processing algorithm used to manage audio pitch for scaled audio edits.
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avmutableaudiomixinputparameters/audiotimepitchalgorithm
func (m_ MutableAudioMixInputParameters) AudioTimePitchAlgorithm() AudioTimePitchAlgorithm {
	rv := objc.Send[AudioTimePitchAlgorithm](m_.ID, objc.Sel("audioTimePitchAlgorithm"))
	return rv
}


// SetAudioTimePitchAlgorithm sets the value of the audioTimePitchAlgorithm property.
// The processing algorithm used to manage audio pitch for scaled audio edits.

//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avmutableaudiomixinputparameters/audiotimepitchalgorithm
func (m_ MutableAudioMixInputParameters) SetAudioTimePitchAlgorithm(value IAudioTimePitchAlgorithm) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setAudioTimePitchAlgorithm:"), value)
}

// The identifier of the audio track to which the parameters should be applied.
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avmutableaudiomixinputparameters/trackid
func (m_ MutableAudioMixInputParameters) TrackID() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("trackID"))
	return rv
}


// SetTrackID sets the value of the trackID property.
// The identifier of the audio track to which the parameters should be applied.

//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avmutableaudiomixinputparameters/trackid
func (m_ MutableAudioMixInputParameters) SetTrackID(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTrackID:"), value)
}



