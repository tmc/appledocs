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
	AudioTimeStamp() unsafe.Pointer
	SetAudioTimeStamp(value unsafe.Pointer)
	HostTime() uint64
	SetHostTime(value uint64)
	IsHostTimeValid() bool
	SetIsHostTimeValid(value bool)
	IsSampleTimeValid() bool
	SetIsSampleTimeValid(value bool)
	SampleRate() float64
	SetSampleRate(value float64)
	SampleTime() unsafe.Pointer
	SetSampleTime(value unsafe.Pointer)
}

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



// The time as an audio timestamp.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avaudiotime/audiotimestamp
func (a_ AudioTime) AudioTimeStamp() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("audioTimeStamp"))
	return rv
}


// The time as an audio timestamp.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avaudiotime/audiotimestamp
func (a_ AudioTime) SetAudioTimeStamp(value unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setAudioTimeStamp:"), value)
}


// The host time.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avaudiotime/hosttime
func (a_ AudioTime) HostTime() uint64 {
	rv := objc.Send[uint64](a_.ID, objc.Sel("hostTime"))
	return rv
}


// The host time.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avaudiotime/hosttime
func (a_ AudioTime) SetHostTime(value uint64) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setHostTime:"), value)
}


// A Boolean value that indicates whether the host time value is valid.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avaudiotime/ishosttimevalid
func (a_ AudioTime) IsHostTimeValid() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("isHostTimeValid"))
	return rv
}


// A Boolean value that indicates whether the host time value is valid.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avaudiotime/ishosttimevalid
func (a_ AudioTime) SetIsHostTimeValid(value bool) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setIsHostTimeValid:"), value)
}


// A Boolean value that indicates whether the sample time and sample rate properties are in a valid state.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avaudiotime/issampletimevalid
func (a_ AudioTime) IsSampleTimeValid() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("isSampleTimeValid"))
	return rv
}


// A Boolean value that indicates whether the sample time and sample rate properties are in a valid state.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avaudiotime/issampletimevalid
func (a_ AudioTime) SetIsSampleTimeValid(value bool) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setIsSampleTimeValid:"), value)
}


// The sampling rate that the sample time property expresses.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avaudiotime/samplerate
func (a_ AudioTime) SampleRate() float64 {
	rv := objc.Send[float64](a_.ID, objc.Sel("sampleRate"))
	return rv
}


// The sampling rate that the sample time property expresses.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avaudiotime/samplerate
func (a_ AudioTime) SetSampleRate(value float64) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setSampleRate:"), value)
}


// The time as a number of audio samples that the current audio device tracks.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avaudiotime/sampletime
func (a_ AudioTime) SampleTime() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("sampleTime"))
	return rv
}


// The time as a number of audio samples that the current audio device tracks.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avaudiotime/sampletime
func (a_ AudioTime) SetSampleTime(value unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setSampleTime:"), value)
}



