// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/corevideo"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class AVAudioMixInputParameters */


/* debug [class_header]: Header for AVAudioMixInputParameters */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for AudioMixInputParameters */
// An interface definition for the [AudioMixInputParameters] class.
type IAudioMixInputParameters interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for AudioMixInputParameters */
	// properties:
	AudioTapProcessor() objectivec.IObject
	AudioTimePitchAlgorithm() AudioTimePitchAlgorithm /* typedef */
	TrackID() PersistentTrackID /* not a class type */
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for AudioMixInputParameters */
	// methods:
	GetVolumeRampForTimeStartVolumeEndVolumeTimeRange(time objc.IObject /* cross-framework: Time */, startVolume objectivec.IObject, endVolume objectivec.IObject, timeRange TimeRange /* not a class type */) bool
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for AudioMixInputParameters */
// Alloc allocates a new instance without initialization.
func (ac _AudioMixInputParametersClass) Alloc() AudioMixInputParameters {
	rv := objc.Send[AudioMixInputParameters](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for AudioMixInputParameters */
// An object that represents the parameters that you apply when adding an audio track to a mix.
//
// You use an instance to apply audio volume ramps for an input to an audio mix. Mix parameters are associated with audio tracks via the property. Audio volume is currently supported as a time-varying parameter. has a mutable subclass, . Before the first time at which a volume is set, a volume of 1.0 used; after the last time for which a volume has been set, the last volume is used. Within the time range of a volume ramp, the volume is interpolated between the start volume and end volume of the ramp. For example, setting the volume to 1.0 at time 0 and also setting a volume ramp from a volume of 0.5 to 0.2 with a timeRange of [4.0, 5.0] results in an audio volume parameters that hold the volume constant at 1.0 from 0.0 sec to 4.0 sec, then cause it to jump to 0.5 and descend to 0.2 from 4.0 sec to 9.0 sec, holding constant at 0.2 thereafter. Given that this is an immutable variant of the object, you should not allocate and initialize a version of this class yourself. Other classes may return instances of this class.


// An object that represents the parameters that you apply when adding an audio track to a mix.
//
// [Full Topic]
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for AudioMixInputParameters *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for AudioMixInputParameters */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for AudioMixInputParameters */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for AudioMixInputParameters */

// Retrieves the volume ramp that includes the specified time.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAudioMixInputParameters/getVolumeRamp(for:startVolume:endVolume:timeRange:)
func (a_ AudioMixInputParameters) GetVolumeRampForTimeStartVolumeEndVolumeTimeRange(time objc.IObject /* cross-framework: Time */, startVolume objectivec.IObject, endVolume objectivec.IObject, timeRange TimeRange /* not a class type */) bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("getVolumeRampForTime:startVolume:endVolume:timeRange:"), time, startVolume, endVolume, timeRange)
	return rv
}/* debug [instance_methods/method]: GetVolumeRampForTimeStartVolumeEndVolumeTimeRange */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for AudioMixInputParameters */

// The audio processing tap associated with the track.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAudioMixInputParameters/audioTapProcessor
func (a_ AudioMixInputParameters) AudioTapProcessor() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](a_.ID, objc.Sel("audioTapProcessor"))
	return rv
}/* debug [instance_properties/getter]: audioTapProcessor */


// The processing algorithm used to manage audio pitch for scaled audio edits.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAudioMixInputParameters/audioTimePitchAlgorithm
func (a_ AudioMixInputParameters) AudioTimePitchAlgorithm() AudioTimePitchAlgorithm /* typedef */ {
	rv := objc.Send[foundation.NSString](a_.ID, objc.Sel("audioTimePitchAlgorithm"))
	return rv
}/* debug [instance_properties/getter]: audioTimePitchAlgorithm */


// The identifier of the audio track to which the parameters should be applied.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAudioMixInputParameters/trackID
func (a_ AudioMixInputParameters) TrackID() PersistentTrackID /* not a class type */ {
	rv := objc.Send[PersistentTrackID](a_.ID, objc.Sel("trackID"))
	return rv
}/* debug [instance_properties/getter]: trackID */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class AVAudioMixInputParameters */



