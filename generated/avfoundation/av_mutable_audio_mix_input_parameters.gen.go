// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/corevideo"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class AVMutableAudioMixInputParameters */


/* debug [class_header]: Header for AVMutableAudioMixInputParameters */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MutableAudioMixInputParameters */
// An interface definition for the [MutableAudioMixInputParameters] class.
type IMutableAudioMixInputParameters interface {
	IAudioMixInputParameters
	
/* debug [class_interface_properties]: Properties for MutableAudioMixInputParameters */
	// properties:
	AudioTapProcessor() objectivec.IObject
	SetAudioTapProcessor(value objectivec.IObject)
	AudioTimePitchAlgorithm() AudioTimePitchAlgorithm /* typedef */
	SetAudioTimePitchAlgorithm(value AudioTimePitchAlgorithm /* typedef */)
	TrackID() PersistentTrackID /* not a class type */
	SetTrackID(value PersistentTrackID /* not a class type */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MutableAudioMixInputParameters */
	// methods:
	SetVolumeAtTime(volume float32, time objc.IObject /* cross-framework: Time */)
	SetVolumeRampFromStartVolumeToEndVolumeTimeRange(startVolume float32, endVolume float32, timeRange TimeRange /* not a class type */)
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MutableAudioMixInputParameters */
// Alloc allocates a new instance without initialization.
func (mc _MutableAudioMixInputParametersClass) Alloc() MutableAudioMixInputParameters {
	rv := objc.Send[MutableAudioMixInputParameters](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MutableAudioMixInputParameters */
// The parameters you use when adding an audio track to a mix.


// The parameters you use when adding an audio track to a mix.
//
// [Full Topic]
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MutableAudioMixInputParameters */

// Creates a mutable input parameters object for a given track.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMutableAudioMixInputParameters/init(track:)
func NewMutableAudioMixInputParametersWithTrack(track IAVAssetTrack) MutableAudioMixInputParameters {
	rv := objc.Send[MutableAudioMixInputParameters](objc.ID(getMutableAudioMixInputParametersClass().class), objc.Sel("audioMixInputParametersWithTrack:"), track)
	return rv
}/* debug [class_init_methods/constructor]: NewMutableAudioMixInputParametersWithTrack */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MutableAudioMixInputParameters */

// Creates a mutable input parameters object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMutableAudioMixInputParameters/audioMixInputParameters
func (mc _MutableAudioMixInputParametersClass) AudioMixInputParameters() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(mc.class), objc.Sel("audioMixInputParameters"))
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=AudioMixInputParameters) */


// Creates a mutable input parameters object for a given track.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMutableAudioMixInputParameters/init(track:)
func (mc _MutableAudioMixInputParametersClass) AudioMixInputParametersWithTrack(track IAVAssetTrack) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(mc.class), objc.Sel("audioMixInputParametersWithTrack:"), track)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=AudioMixInputParametersWithTrack) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MutableAudioMixInputParameters */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MutableAudioMixInputParameters */

// Sets the value of the audio volume starting at the specified time.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMutableAudioMixInputParameters/setVolume(_:at:)
func (m_ MutableAudioMixInputParameters) SetVolumeAtTime(volume float32, time objc.IObject /* cross-framework: Time */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setVolume:atTime:"), volume, time)
}/* debug [instance_methods/method]: SetVolumeAtTime */


// Sets a volume ramp to apply during a specified time range.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMutableAudioMixInputParameters/setVolumeRamp(fromStartVolume:toEndVolume:timeRange:)
func (m_ MutableAudioMixInputParameters) SetVolumeRampFromStartVolumeToEndVolumeTimeRange(startVolume float32, endVolume float32, timeRange TimeRange /* not a class type */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setVolumeRampFromStartVolume:toEndVolume:timeRange:"), startVolume, endVolume, timeRange)
}/* debug [instance_methods/method]: SetVolumeRampFromStartVolumeToEndVolumeTimeRange */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MutableAudioMixInputParameters */

// The audio processing tap associated with the track.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMutableAudioMixInputParameters/audioTapProcessor
func (m_ MutableAudioMixInputParameters) AudioTapProcessor() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](m_.ID, objc.Sel("audioTapProcessor"))
	return rv
}/* debug [instance_properties/getter]: audioTapProcessor */


// The audio processing tap associated with the track.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMutableAudioMixInputParameters/audioTapProcessor
func (m_ MutableAudioMixInputParameters) SetAudioTapProcessor(value objectivec.IObject) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setAudioTapProcessor:"), value)
}/* debug [instance_properties/setter]: audioTapProcessor */


// The processing algorithm used to manage audio pitch for scaled audio edits.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMutableAudioMixInputParameters/audioTimePitchAlgorithm
func (m_ MutableAudioMixInputParameters) AudioTimePitchAlgorithm() AudioTimePitchAlgorithm /* typedef */ {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("audioTimePitchAlgorithm"))
	return rv
}/* debug [instance_properties/getter]: audioTimePitchAlgorithm */


// The processing algorithm used to manage audio pitch for scaled audio edits.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMutableAudioMixInputParameters/audioTimePitchAlgorithm
func (m_ MutableAudioMixInputParameters) SetAudioTimePitchAlgorithm(value AudioTimePitchAlgorithm /* typedef */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setAudioTimePitchAlgorithm:"), value)
}/* debug [instance_properties/setter]: audioTimePitchAlgorithm */


// The identifier of the audio track to which the parameters should be applied.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMutableAudioMixInputParameters/trackID
func (m_ MutableAudioMixInputParameters) TrackID() PersistentTrackID /* not a class type */ {
	rv := objc.Send[PersistentTrackID](m_.ID, objc.Sel("trackID"))
	return rv
}/* debug [instance_properties/getter]: trackID */


// The identifier of the audio track to which the parameters should be applied.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMutableAudioMixInputParameters/trackID
func (m_ MutableAudioMixInputParameters) SetTrackID(value PersistentTrackID /* not a class type */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTrackID:"), value)
}/* debug [instance_properties/setter]: trackID */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class AVMutableAudioMixInputParameters */


