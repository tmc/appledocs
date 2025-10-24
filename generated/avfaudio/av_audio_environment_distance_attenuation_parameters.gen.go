// Code generated from Apple documentation for AVFAudio. DO NOT EDIT.

package avfaudio

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class AVAudioEnvironmentDistanceAttenuationParameters */


/* debug [class_header]: Header for AVAudioEnvironmentDistanceAttenuationParameters */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for AudioEnvironmentDistanceAttenuationParameters */
// An interface definition for the [AudioEnvironmentDistanceAttenuationParameters] class.
type IAudioEnvironmentDistanceAttenuationParameters interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for AudioEnvironmentDistanceAttenuationParameters */
	// properties:
	DistanceAttenuationModel() AudioEnvironmentDistanceAttenuationModel
	SetDistanceAttenuationModel(value AudioEnvironmentDistanceAttenuationModel)
	MaximumDistance() float32
	SetMaximumDistance(value float32)
	ReferenceDistance() float32
	SetReferenceDistance(value float32)
	RolloffFactor() float32
	SetRolloffFactor(value float32)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for AudioEnvironmentDistanceAttenuationParameters */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for AudioEnvironmentDistanceAttenuationParameters */
// Alloc allocates a new instance without initialization.
func (ac _AudioEnvironmentDistanceAttenuationParametersClass) Alloc() AudioEnvironmentDistanceAttenuationParameters {
	rv := objc.Send[AudioEnvironmentDistanceAttenuationParameters](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for AudioEnvironmentDistanceAttenuationParameters */
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for AudioEnvironmentDistanceAttenuationParameters *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for AudioEnvironmentDistanceAttenuationParameters */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for AudioEnvironmentDistanceAttenuationParameters */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for AudioEnvironmentDistanceAttenuationParameters */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for AudioEnvironmentDistanceAttenuationParameters */

// The distance attenuation model that describes the drop-off in gain as the source moves away from the listener.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioEnvironmentDistanceAttenuationParameters/distanceAttenuationModel
func (a_ AudioEnvironmentDistanceAttenuationParameters) DistanceAttenuationModel() AudioEnvironmentDistanceAttenuationModel {
	rv := objc.Send[AudioEnvironmentDistanceAttenuationModel](a_.ID, objc.Sel("distanceAttenuationModel"))
	return rv
}/* debug [instance_properties/getter]: distanceAttenuationModel */


// The distance attenuation model that describes the drop-off in gain as the source moves away from the listener.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioEnvironmentDistanceAttenuationParameters/distanceAttenuationModel
func (a_ AudioEnvironmentDistanceAttenuationParameters) SetDistanceAttenuationModel(value AudioEnvironmentDistanceAttenuationModel) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setDistanceAttenuationModel:"), value)
}/* debug [instance_properties/setter]: distanceAttenuationModel */


// The distance beyond which the node applies no further attenuation, in meters.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioEnvironmentDistanceAttenuationParameters/maximumDistance
func (a_ AudioEnvironmentDistanceAttenuationParameters) MaximumDistance() float32 {
	rv := objc.Send[float32](a_.ID, objc.Sel("maximumDistance"))
	return rv
}/* debug [instance_properties/getter]: maximumDistance */


// The distance beyond which the node applies no further attenuation, in meters.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioEnvironmentDistanceAttenuationParameters/maximumDistance
func (a_ AudioEnvironmentDistanceAttenuationParameters) SetMaximumDistance(value float32) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setMaximumDistance:"), value)
}/* debug [instance_properties/setter]: maximumDistance */


// The minimum distance at which the node applies attenuation, in meters.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioEnvironmentDistanceAttenuationParameters/referenceDistance
func (a_ AudioEnvironmentDistanceAttenuationParameters) ReferenceDistance() float32 {
	rv := objc.Send[float32](a_.ID, objc.Sel("referenceDistance"))
	return rv
}/* debug [instance_properties/getter]: referenceDistance */


// The minimum distance at which the node applies attenuation, in meters.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioEnvironmentDistanceAttenuationParameters/referenceDistance
func (a_ AudioEnvironmentDistanceAttenuationParameters) SetReferenceDistance(value float32) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setReferenceDistance:"), value)
}/* debug [instance_properties/setter]: referenceDistance */


// A factor that determines the attenuation curve.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioEnvironmentDistanceAttenuationParameters/rolloffFactor
func (a_ AudioEnvironmentDistanceAttenuationParameters) RolloffFactor() float32 {
	rv := objc.Send[float32](a_.ID, objc.Sel("rolloffFactor"))
	return rv
}/* debug [instance_properties/getter]: rolloffFactor */


// A factor that determines the attenuation curve.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioEnvironmentDistanceAttenuationParameters/rolloffFactor
func (a_ AudioEnvironmentDistanceAttenuationParameters) SetRolloffFactor(value float32) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setRolloffFactor:"), value)
}/* debug [instance_properties/setter]: rolloffFactor */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class AVAudioEnvironmentDistanceAttenuationParameters */



