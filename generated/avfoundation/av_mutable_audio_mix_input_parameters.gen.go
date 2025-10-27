// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
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
	

	// properties:
	AudioTapProcessor() objectivec.IObject
	SetAudioTapProcessor(value objectivec.IObject)
	AudioTimePitchAlgorithm() AudioTimePitchAlgorithm
	SetAudioTimePitchAlgorithm(value AudioTimePitchAlgorithm)
	TrackID() PersistentTrackID /* not a class type */
	SetTrackID(value PersistentTrackID /* not a class type */)


	

	// methods:
	SetVolumeAtTime(volume float32, time objectivec.IObject)
	SetVolumeRampFromStartVolumeToEndVolumeTimeRange(startVolume float32, endVolume float32, timeRange objectivec.IObject)


}





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






// Creates a mutable input parameters object for a given track.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMutableAudioMixInputParameters/init(track:)
func NewMutableAudioMixInputParametersWithTrack(track IAVAssetTrack) MutableAudioMixInputParameters {
	rv := objc.Send[MutableAudioMixInputParameters](objc.ID(getMutableAudioMixInputParametersClass().class), objc.Sel("audioMixInputParametersWithTrack:"), track)
	return rv
}







// Creates a mutable input parameters object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMutableAudioMixInputParameters/audioMixInputParameters
func (mc _MutableAudioMixInputParametersClass) AudioMixInputParameters() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(mc.class), objc.Sel("audioMixInputParameters"))
	return rv
}


// Creates a mutable input parameters object for a given track.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMutableAudioMixInputParameters/init(track:)
func (mc _MutableAudioMixInputParametersClass) AudioMixInputParametersWithTrack(track IAVAssetTrack) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(mc.class), objc.Sel("audioMixInputParametersWithTrack:"), track)
	return rv
}












// Sets the value of the audio volume starting at the specified time.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMutableAudioMixInputParameters/setVolume(_:at:)
func (m_ MutableAudioMixInputParameters) SetVolumeAtTime(volume float32, time objectivec.IObject) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setVolume:atTime:"), volume, time)
}


// Sets a volume ramp to apply during a specified time range.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMutableAudioMixInputParameters/setVolumeRamp(fromStartVolume:toEndVolume:timeRange:)
func (m_ MutableAudioMixInputParameters) SetVolumeRampFromStartVolumeToEndVolumeTimeRange(startVolume float32, endVolume float32, timeRange objectivec.IObject) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setVolumeRampFromStartVolume:toEndVolume:timeRange:"), startVolume, endVolume, timeRange)
}







// The audio processing tap associated with the track.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMutableAudioMixInputParameters/audioTapProcessor
func (m_ MutableAudioMixInputParameters) AudioTapProcessor() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](m_.ID, objc.Sel("audioTapProcessor"))
	return rv
}


// The audio processing tap associated with the track.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMutableAudioMixInputParameters/audioTapProcessor
func (m_ MutableAudioMixInputParameters) SetAudioTapProcessor(value objectivec.IObject) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setAudioTapProcessor:"), value)
}


// The processing algorithm used to manage audio pitch for scaled audio edits.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMutableAudioMixInputParameters/audioTimePitchAlgorithm
func (m_ MutableAudioMixInputParameters) AudioTimePitchAlgorithm() AudioTimePitchAlgorithm {
	rv := objc.Send[AudioTimePitchAlgorithm](m_.ID, objc.Sel("audioTimePitchAlgorithm"))
	return rv
}


// The processing algorithm used to manage audio pitch for scaled audio edits.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMutableAudioMixInputParameters/audioTimePitchAlgorithm
func (m_ MutableAudioMixInputParameters) SetAudioTimePitchAlgorithm(value AudioTimePitchAlgorithm) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setAudioTimePitchAlgorithm:"), value)
}


// The identifier of the audio track to which the parameters should be applied.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMutableAudioMixInputParameters/trackID
func (m_ MutableAudioMixInputParameters) TrackID() PersistentTrackID /* not a class type */ {
	rv := objc.Send[PersistentTrackID](m_.ID, objc.Sel("trackID"))
	return rv
}


// The identifier of the audio track to which the parameters should be applied.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMutableAudioMixInputParameters/trackID
func (m_ MutableAudioMixInputParameters) SetTrackID(value PersistentTrackID /* not a class type */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTrackID:"), value)
}







