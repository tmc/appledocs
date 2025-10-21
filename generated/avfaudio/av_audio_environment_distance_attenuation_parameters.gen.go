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
}

// An object that specifies the amount of attenuation distance, the gradual loss in audio intensity, and other characteristics.
//
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


// A factor that determines the attenuation curve.
//
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avaudioenvironmentdistanceattenuationparameters/rollofffactor
func (a_ AudioEnvironmentDistanceAttenuationParameters) RolloffFactor() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("rolloffFactor"))
	return rv
}


// SetRolloffFactor sets the value of the rolloffFactor property.
// A factor that determines the attenuation curve.

//
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avaudioenvironmentdistanceattenuationparameters/rollofffactor
func (a_ AudioEnvironmentDistanceAttenuationParameters) SetRolloffFactor(value unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setRolloffFactor:"), value)
}

// The distance attenuation model that describes the drop-off in gain as the source moves away from the listener.
//
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avaudioenvironmentdistanceattenuationparameters/distanceattenuationmodel
func (a_ AudioEnvironmentDistanceAttenuationParameters) DistanceAttenuationModel() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("distanceAttenuationModel"))
	return rv
}


// SetDistanceAttenuationModel sets the value of the distanceAttenuationModel property.
// The distance attenuation model that describes the drop-off in gain as the source moves away from the listener.

//
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avaudioenvironmentdistanceattenuationparameters/distanceattenuationmodel
func (a_ AudioEnvironmentDistanceAttenuationParameters) SetDistanceAttenuationModel(value unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setDistanceAttenuationModel:"), value)
}

// The distance beyond which the node applies no further attenuation, in meters.
//
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avaudioenvironmentdistanceattenuationparameters/maximumdistance
func (a_ AudioEnvironmentDistanceAttenuationParameters) MaximumDistance() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("maximumDistance"))
	return rv
}


// SetMaximumDistance sets the value of the maximumDistance property.
// The distance beyond which the node applies no further attenuation, in meters.

//
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avaudioenvironmentdistanceattenuationparameters/maximumdistance
func (a_ AudioEnvironmentDistanceAttenuationParameters) SetMaximumDistance(value unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setMaximumDistance:"), value)
}

// The minimum distance at which the node applies attenuation, in meters.
//
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avaudioenvironmentdistanceattenuationparameters/referencedistance
func (a_ AudioEnvironmentDistanceAttenuationParameters) ReferenceDistance() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("referenceDistance"))
	return rv
}


// SetReferenceDistance sets the value of the referenceDistance property.
// The minimum distance at which the node applies attenuation, in meters.

//
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avaudioenvironmentdistanceattenuationparameters/referencedistance
func (a_ AudioEnvironmentDistanceAttenuationParameters) SetReferenceDistance(value unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setReferenceDistance:"), value)
}



