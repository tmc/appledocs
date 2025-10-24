// Code generated from Apple documentation for AVFAudio. DO NOT EDIT.

package avfaudio

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class AVAudioTime */


/* debug [class_header]: Header for AVAudioTime */
// The class instance for the [AudioTime] class.
var (
	AudioTimeClass     _AudioTimeClass
	AudioTimeClassOnce sync.Once
)

func getAudioTimeClass() _AudioTimeClass {
	AudioTimeClassOnce.Do(func() {
		AudioTimeClass = _AudioTimeClass{objc.GetClass("AVAudioTime")}
	})
	return AudioTimeClass
}

type _AudioTimeClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for AudioTime */
// An interface definition for the [AudioTime] class.
type IAudioTime interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for AudioTime */
	// properties:
	AudioTimeStamp() objc.IObject
	HostTime() uint64
	HostTimeValid() bool
	SampleTimeValid() bool
	SampleRate() float64
	SampleTime() AudioFramePosition /* typedef */
	IsHostTimeValid() bool
	SetIsHostTimeValid(value bool)
	IsSampleTimeValid() bool
	SetIsSampleTimeValid(value bool)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for AudioTime */
	// methods:
	ExtrapolateTimeFromAnchor(anchorTime IAVAudioTime) IAudioTime
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for AudioTime */
// Alloc allocates a new instance without initialization.
func (ac _AudioTimeClass) Alloc() AudioTime {
	rv := objc.Send[AudioTime](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (ac _AudioTimeClass) New() AudioTime {
	rv := objc.Send[AudioTime](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ AudioTime) Init() AudioTime {
	rv := objc.Send[AudioTime](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ AudioTime) Autorelease() AudioTime {
	rv := objc.Send[AudioTime](a_.ID, objc.Sel("autorelease"))
	return rv
}

// NewAudioTime creates a new AudioTime instance.
func NewAudioTime() AudioTime {
	return getAudioTimeClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for AudioTime */
// An object you use to represent a moment in time.
//
// The object represents a single moment in time in two ways: As host time, using the system’s basic clock with As audio samples at a particular sample rate A single instance contains either or both representations, meaning it might represent only a sample time, a host time, or both. Instances of this class are immutable.


// An object you use to represent a moment in time.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioTime
type AudioTime struct {
	objectivec.Object
}

// AudioTimeFrom constructs a [AudioTime] from an unsafe.Pointer.
//
// An object you use to represent a moment in time.
func AudioTimeFrom(ptr unsafe.Pointer) AudioTime {
	return AudioTime{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for AudioTime */

// Creates an audio time object with the specified timestamp and sample rate.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioTime/init(audioTimeStamp:sampleRate:)
func NewAudioTimeWithAudioTimeStampSampleRate(ts objc.IObject, sampleRate float64) AudioTime {
	instance := getAudioTimeClass().Alloc()
	rv := objc.Send[AudioTime](instance.ID, objc.Sel("initWithAudioTimeStamp:sampleRate:"), ts, sampleRate)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewAudioTimeWithAudioTimeStampSampleRate */


// Creates an audio time object with the specified host time.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioTime/init(hostTime:)
func NewAudioTimeWithHostTime(hostTime uint64) AudioTime {
	instance := getAudioTimeClass().Alloc()
	rv := objc.Send[AudioTime](instance.ID, objc.Sel("initWithHostTime:"), hostTime)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewAudioTimeWithHostTime */


// Creates an audio time object with the specified host time, sample time, and sample rate.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioTime/init(hostTime:sampleTime:atRate:)
func NewAudioTimeWithHostTimeSampleTimeAtRate(hostTime uint64, sampleTime AudioFramePosition /* typedef */, sampleRate float64) AudioTime {
	instance := getAudioTimeClass().Alloc()
	rv := objc.Send[AudioTime](instance.ID, objc.Sel("initWithHostTime:sampleTime:atRate:"), hostTime, sampleTime, sampleRate)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewAudioTimeWithHostTimeSampleTimeAtRate */


// Creates an audio time object with the specified timestamp and sample rate.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioTime/init(sampleTime:atRate:)
func NewAudioTimeWithSampleTimeAtRate(sampleTime AudioFramePosition /* typedef */, sampleRate float64) AudioTime {
	instance := getAudioTimeClass().Alloc()
	rv := objc.Send[AudioTime](instance.ID, objc.Sel("initWithSampleTime:atRate:"), sampleTime, sampleRate)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewAudioTimeWithSampleTimeAtRate */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for AudioTime */

// Converts seconds to host time.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioTime/hostTime(forSeconds:)
func (ac _AudioTimeClass) HostTimeForSeconds(seconds float64) uint64 {
	rv := objc.Send[uint64](objc.ID(ac.class), objc.Sel("hostTimeForSeconds:"), seconds)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=HostTimeForSeconds) */


// Converts host time to seconds.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioTime/seconds(forHostTime:)
func (ac _AudioTimeClass) SecondsForHostTime(hostTime uint64) float64 {
	rv := objc.Send[float64](objc.ID(ac.class), objc.Sel("secondsForHostTime:"), hostTime)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=SecondsForHostTime) */


// Creates an audio time object with the specified timestamp and sample rate.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioTime/timeWithAudioTimeStamp:sampleRate:
func (ac _AudioTimeClass) TimeWithAudioTimeStampSampleRate(ts objc.IObject, sampleRate float64) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(ac.class), objc.Sel("timeWithAudioTimeStamp:sampleRate:"), ts, sampleRate)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=TimeWithAudioTimeStampSampleRate) */


// Creates an audio time object with the specified host time.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioTime/timeWithHostTime:
func (ac _AudioTimeClass) TimeWithHostTime(hostTime uint64) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(ac.class), objc.Sel("timeWithHostTime:"), hostTime)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=TimeWithHostTime) */


// Creates an audio time object with the specified host time, sample time, and sample rate.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioTime/timeWithHostTime:sampleTime:atRate:
func (ac _AudioTimeClass) TimeWithHostTimeSampleTimeAtRate(hostTime uint64, sampleTime AudioFramePosition /* typedef */, sampleRate float64) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(ac.class), objc.Sel("timeWithHostTime:sampleTime:atRate:"), hostTime, sampleTime, sampleRate)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=TimeWithHostTimeSampleTimeAtRate) */


// Creates an audio time object with the specified sample time and sample rate.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioTime/timeWithSampleTime:atRate:
func (ac _AudioTimeClass) TimeWithSampleTimeAtRate(sampleTime AudioFramePosition /* typedef */, sampleRate float64) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(ac.class), objc.Sel("timeWithSampleTime:atRate:"), sampleTime, sampleRate)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=TimeWithSampleTimeAtRate) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for AudioTime */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for AudioTime */

// Creates an audio time object by converting between host time and sample time.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioTime/extrapolateTime(fromAnchor:)
func (a_ AudioTime) ExtrapolateTimeFromAnchor(anchorTime IAVAudioTime) IAudioTime {
	rv := objc.Send[AudioTime](a_.ID, objc.Sel("extrapolateTimeFromAnchor:"), anchorTime)
	return rv
}/* debug [instance_methods/method]: ExtrapolateTimeFromAnchor */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for AudioTime */

// The time as an audio timestamp.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioTime/audioTimeStamp
func (a_ AudioTime) AudioTimeStamp() objc.IObject {
	rv := objc.Send[objc.ID](a_.ID, objc.Sel("audioTimeStamp"))
	return rv
}/* debug [instance_properties/getter]: audioTimeStamp */


// The host time.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioTime/hostTime
func (a_ AudioTime) HostTime() uint64 {
	rv := objc.Send[uint64](a_.ID, objc.Sel("hostTime"))
	return rv
}/* debug [instance_properties/getter]: hostTime */


// A Boolean value that indicates whether the host time value is valid.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioTime/isHostTimeValid
func (a_ AudioTime) HostTimeValid() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("hostTimeValid"))
	return rv
}/* debug [instance_properties/getter]: hostTimeValid */


// A Boolean value that indicates whether the sample time and sample rate properties are in a valid state.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioTime/isSampleTimeValid
func (a_ AudioTime) SampleTimeValid() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("sampleTimeValid"))
	return rv
}/* debug [instance_properties/getter]: sampleTimeValid */


// The sampling rate that the sample time property expresses.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioTime/sampleRate
func (a_ AudioTime) SampleRate() float64 {
	rv := objc.Send[float64](a_.ID, objc.Sel("sampleRate"))
	return rv
}/* debug [instance_properties/getter]: sampleRate */


// The time as a number of audio samples that the current audio device tracks.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioTime/sampleTime
func (a_ AudioTime) SampleTime() AudioFramePosition /* typedef */ {
	rv := objc.Send[int64](a_.ID, objc.Sel("sampleTime"))
	return rv
}/* debug [instance_properties/getter]: sampleTime */


// A Boolean value that indicates whether the host time value is valid.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avaudiotime/ishosttimevalid
func (a_ AudioTime) IsHostTimeValid() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("isHostTimeValid"))
	return rv
}/* debug [instance_properties/getter]: isHostTimeValid */


// A Boolean value that indicates whether the host time value is valid.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avaudiotime/ishosttimevalid
func (a_ AudioTime) SetIsHostTimeValid(value bool) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setIsHostTimeValid:"), value)
}/* debug [instance_properties/setter]: isHostTimeValid */


// A Boolean value that indicates whether the sample time and sample rate properties are in a valid state.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avaudiotime/issampletimevalid
func (a_ AudioTime) IsSampleTimeValid() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("isSampleTimeValid"))
	return rv
}/* debug [instance_properties/getter]: isSampleTimeValid */


// A Boolean value that indicates whether the sample time and sample rate properties are in a valid state.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avaudiotime/issampletimevalid
func (a_ AudioTime) SetIsSampleTimeValid(value bool) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setIsSampleTimeValid:"), value)
}/* debug [instance_properties/setter]: isSampleTimeValid */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class AVAudioTime */


