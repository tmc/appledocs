// Code generated from Apple documentation for AVFAudio. DO NOT EDIT.

package avfaudio

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

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

// An interface definition for the [AudioTime] class.
type IAudioTime interface {
	objectivec.IObject
	ExtrapolateTimeFromAnchor(anchorTime unsafe.Pointer) unsafe.Pointer
}

// An object you use to represent a moment in time.
//
// The object represents a single moment in time in two ways: As host time, using the system’s basic clock with As audio samples at a particular sample rate A single instance contains either or both representations, meaning it might represent only a sample time, a host time, or both. Instances of this class are immutable.
//
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

// Alloc allocates a new instance without initialization.
func (ac _AudioTimeClass) Alloc() AudioTime {
	rv := objc.Send[AudioTime](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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


// Creates an audio time object with the specified timestamp and sample rate.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioTime/init(audioTimeStamp:sampleRate:)
func NewAudioTimeWithAudioTimeStampSampleRate(ts unsafe.Pointer, sampleRate unsafe.Pointer) AudioTime {
	instance := getAudioTimeClass().Alloc()
	rv := objc.Send[AudioTime](instance.ID, objc.Sel("initWithAudioTimeStamp:sampleRate:"), ts, sampleRate)
	rv.Autorelease()
	return rv
}

// Creates an audio time object with the specified host time.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioTime/init(hostTime:)
func NewAudioTimeWithHostTime(hostTime uint64) AudioTime {
	instance := getAudioTimeClass().Alloc()
	rv := objc.Send[AudioTime](instance.ID, objc.Sel("initWithHostTime:"), hostTime)
	rv.Autorelease()
	return rv
}

// Creates an audio time object with the specified host time, sample time, and sample rate.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioTime/init(hostTime:sampleTime:atRate:)
func NewAudioTimeWithHostTimeSampleTimeAtRate(hostTime uint64, sampleTime unsafe.Pointer, sampleRate unsafe.Pointer) AudioTime {
	instance := getAudioTimeClass().Alloc()
	rv := objc.Send[AudioTime](instance.ID, objc.Sel("initWithHostTime:sampleTime:atRate:"), hostTime, sampleTime, sampleRate)
	rv.Autorelease()
	return rv
}

// Creates an audio time object with the specified timestamp and sample rate.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioTime/init(sampleTime:atRate:)
func NewAudioTimeWithSampleTimeAtRate(sampleTime unsafe.Pointer, sampleRate unsafe.Pointer) AudioTime {
	instance := getAudioTimeClass().Alloc()
	rv := objc.Send[AudioTime](instance.ID, objc.Sel("initWithSampleTime:atRate:"), sampleTime, sampleRate)
	rv.Autorelease()
	return rv
}


// Converts seconds to host time.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioTime/hostTime(forSeconds:)
func (ac _AudioTimeClass) HostTimeForSeconds(seconds TimeInterval) uint64 {
	rv := objc.Send[uint64](objc.ID(ac.class), objc.Sel("hostTimeForSeconds:"), seconds)
	return rv
}

// Converts host time to seconds.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioTime/seconds(forHostTime:)
func (ac _AudioTimeClass) SecondsForHostTime(hostTime uint64) TimeInterval {
	rv := objc.Send[TimeInterval](objc.ID(ac.class), objc.Sel("secondsForHostTime:"), hostTime)
	return rv
}

// Creates an audio time object with the specified timestamp and sample rate.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioTime/timeWithAudioTimeStamp:sampleRate:
func (ac _AudioTimeClass) TimeWithAudioTimeStampSampleRate(ts unsafe.Pointer, sampleRate unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(ac.class), objc.Sel("timeWithAudioTimeStamp:sampleRate:"), ts, sampleRate)
	return rv
}

// Creates an audio time object with the specified host time.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioTime/timeWithHostTime:
func (ac _AudioTimeClass) TimeWithHostTime(hostTime uint64) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(ac.class), objc.Sel("timeWithHostTime:"), hostTime)
	return rv
}

// Creates an audio time object with the specified host time, sample time, and sample rate.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioTime/timeWithHostTime:sampleTime:atRate:
func (ac _AudioTimeClass) TimeWithHostTimeSampleTimeAtRate(hostTime uint64, sampleTime unsafe.Pointer, sampleRate unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(ac.class), objc.Sel("timeWithHostTime:sampleTime:atRate:"), hostTime, sampleTime, sampleRate)
	return rv
}

// Creates an audio time object with the specified sample time and sample rate.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioTime/timeWithSampleTime:atRate:
func (ac _AudioTimeClass) TimeWithSampleTimeAtRate(sampleTime unsafe.Pointer, sampleRate unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(ac.class), objc.Sel("timeWithSampleTime:atRate:"), sampleTime, sampleRate)
	return rv
}

// Creates an audio time object by converting between host time and sample time.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioTime/extrapolateTime(fromAnchor:)
func (a_ AudioTime) ExtrapolateTimeFromAnchor(anchorTime unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("extrapolateTimeFromAnchor:"), anchorTime)
	return rv
}

// The time as an audio timestamp.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioTime/audioTimeStamp
func (a_ AudioTime) AudioTimeStamp() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("audioTimeStamp"))
	return rv
}

// The host time.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioTime/hostTime
func (a_ AudioTime) HostTime() uint64 {
	rv := objc.Send[uint64](a_.ID, objc.Sel("hostTime"))
	return rv
}

// A Boolean value that indicates whether the host time value is valid.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioTime/isHostTimeValid
func (a_ AudioTime) HostTimeValid() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("hostTimeValid"))
	return rv
}

// The sampling rate that the sample time property expresses.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioTime/sampleRate
func (a_ AudioTime) SampleRate() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("sampleRate"))
	return rv
}

// The time as a number of audio samples that the current audio device tracks.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioTime/sampleTime
func (a_ AudioTime) SampleTime() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("sampleTime"))
	return rv
}


