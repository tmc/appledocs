// Code generated from Apple documentation for AVFAudio. DO NOT EDIT.

package avfaudio

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [AudioEnvironmentDistanceAttenuationParameters] class.
var (
	AudioEnvironmentDistanceAttenuationParametersClass     _AudioEnvironmentDistanceAttenuationParametersClass
	AudioEnvironmentDistanceAttenuationParametersClassOnce sync.Once
)

func getAudioEnvironmentDistanceAttenuationParametersClass() _AudioEnvironmentDistanceAttenuationParametersClass {
	AudioEnvironmentDistanceAttenuationParametersClassOnce.Do(func() {
		AudioEnvironmentDistanceAttenuationParametersClass = _AudioEnvironmentDistanceAttenuationParametersClass{objc.GetClass("AVAudioEnvironmentDistanceAttenuationParameters")}
	})
	return AudioEnvironmentDistanceAttenuationParametersClass
}

type _AudioEnvironmentDistanceAttenuationParametersClass struct {
	class objc.Class
}

// An interface definition for the [AudioEnvironmentDistanceAttenuationParameters] class.
type IAudioEnvironmentDistanceAttenuationParameters interface {
	objectivec.IObject
	DistanceAttenuationModel() unsafe.Pointer
	SetDistanceAttenuationModel(value unsafe.Pointer)
	MaximumDistance() float32
	SetMaximumDistance(value float32)
	ReferenceDistance() float32
	SetReferenceDistance(value float32)
	RolloffFactor() float32
	SetRolloffFactor(value float32)
}

// An object that specifies the amount of attenuation distance, the gradual loss in audio intensity, and other characteristics.


// An object that specifies the amount of attenuation distance, the gradual loss in audio intensity, and other characteristics.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioEnvironmentDistanceAttenuationParameters
type AudioEnvironmentDistanceAttenuationParameters struct {
	objectivec.Object
}

// AudioEnvironmentDistanceAttenuationParametersFrom constructs a [AudioEnvironmentDistanceAttenuationParameters] from an unsafe.Pointer.
//
// An object that specifies the amount of attenuation distance, the gradual loss in audio intensity, and other characteristics.
func AudioEnvironmentDistanceAttenuationParametersFrom(ptr unsafe.Pointer) AudioEnvironmentDistanceAttenuationParameters {
	return AudioEnvironmentDistanceAttenuationParameters{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (ac _AudioEnvironmentDistanceAttenuationParametersClass) Alloc() AudioEnvironmentDistanceAttenuationParameters {
	rv := objc.Send[AudioEnvironmentDistanceAttenuationParameters](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ac _AudioEnvironmentDistanceAttenuationParametersClass) New() AudioEnvironmentDistanceAttenuationParameters {
	rv := objc.Send[AudioEnvironmentDistanceAttenuationParameters](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ AudioEnvironmentDistanceAttenuationParameters) Init() AudioEnvironmentDistanceAttenuationParameters {
	rv := objc.Send[AudioEnvironmentDistanceAttenuationParameters](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ AudioEnvironmentDistanceAttenuationParameters) Autorelease() AudioEnvironmentDistanceAttenuationParameters {
	rv := objc.Send[AudioEnvironmentDistanceAttenuationParameters](a_.ID, objc.Sel("autorelease"))
	return rv
}

// NewAudioEnvironmentDistanceAttenuationParameters creates a new AudioEnvironmentDistanceAttenuationParameters instance.
func NewAudioEnvironmentDistanceAttenuationParameters() AudioEnvironmentDistanceAttenuationParameters {
	return getAudioEnvironmentDistanceAttenuationParametersClass().New()
}



// The distance attenuation model that describes the drop-off in gain as the source moves away from the listener.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avaudioenvironmentdistanceattenuationparameters/distanceattenuationmodel
func (a_ AudioEnvironmentDistanceAttenuationParameters) DistanceAttenuationModel() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("distanceAttenuationModel"))
	return rv
}


// The distance attenuation model that describes the drop-off in gain as the source moves away from the listener.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avaudioenvironmentdistanceattenuationparameters/distanceattenuationmodel
func (a_ AudioEnvironmentDistanceAttenuationParameters) SetDistanceAttenuationModel(value unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setDistanceAttenuationModel:"), value)
}


// The distance beyond which the node applies no further attenuation, in meters.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avaudioenvironmentdistanceattenuationparameters/maximumdistance
func (a_ AudioEnvironmentDistanceAttenuationParameters) MaximumDistance() float32 {
	rv := objc.Send[float32](a_.ID, objc.Sel("maximumDistance"))
	return rv
}


// The distance beyond which the node applies no further attenuation, in meters.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avaudioenvironmentdistanceattenuationparameters/maximumdistance
func (a_ AudioEnvironmentDistanceAttenuationParameters) SetMaximumDistance(value float32) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setMaximumDistance:"), value)
}


// The minimum distance at which the node applies attenuation, in meters.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avaudioenvironmentdistanceattenuationparameters/referencedistance
func (a_ AudioEnvironmentDistanceAttenuationParameters) ReferenceDistance() float32 {
	rv := objc.Send[float32](a_.ID, objc.Sel("referenceDistance"))
	return rv
}


// The minimum distance at which the node applies attenuation, in meters.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avaudioenvironmentdistanceattenuationparameters/referencedistance
func (a_ AudioEnvironmentDistanceAttenuationParameters) SetReferenceDistance(value float32) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setReferenceDistance:"), value)
}


// A factor that determines the attenuation curve.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avaudioenvironmentdistanceattenuationparameters/rollofffactor
func (a_ AudioEnvironmentDistanceAttenuationParameters) RolloffFactor() float32 {
	rv := objc.Send[float32](a_.ID, objc.Sel("rolloffFactor"))
	return rv
}


// A factor that determines the attenuation curve.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avaudioenvironmentdistanceattenuationparameters/rollofffactor
func (a_ AudioEnvironmentDistanceAttenuationParameters) SetRolloffFactor(value float32) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setRolloffFactor:"), value)
}



