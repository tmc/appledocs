// Code generated from Apple documentation for AVFAudio. DO NOT EDIT.

package avfaudio

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)





// The class instance for the [AudioUnitEQFilterParameters] class.
var (
	AudioUnitEQFilterParametersClass     _AudioUnitEQFilterParametersClass
	AudioUnitEQFilterParametersClassOnce sync.Once
)

func getAudioUnitEQFilterParametersClass() _AudioUnitEQFilterParametersClass {
	AudioUnitEQFilterParametersClassOnce.Do(func() {
		AudioUnitEQFilterParametersClass = _AudioUnitEQFilterParametersClass{objc.GetClass("AVAudioUnitEQFilterParameters")}
	})
	return AudioUnitEQFilterParametersClass
}

type _AudioUnitEQFilterParametersClass struct {
	class objc.Class
}





// An interface definition for the [AudioUnitEQFilterParameters] class.
type IAudioUnitEQFilterParameters interface {
	objectivec.IObject
	

	// properties:
	Bandwidth() float32
	SetBandwidth(value float32)
	Bypass() bool
	SetBypass(value bool)
	FilterType() AudioUnitEQFilterType
	SetFilterType(value AudioUnitEQFilterType)
	Frequency() float32
	SetFrequency(value float32)
	Gain() float32
	SetGain(value float32)
	Bands() IAVAudioUnitEQFilterParameters
	SetBands(value IAVAudioUnitEQFilterParameters)
	GlobalGain() float32
	SetGlobalGain(value float32)


	

	// methods:


}





// Alloc allocates a new instance without initialization.
func (ac _AudioUnitEQFilterParametersClass) Alloc() AudioUnitEQFilterParameters {
	rv := objc.Send[AudioUnitEQFilterParameters](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (ac _AudioUnitEQFilterParametersClass) New() AudioUnitEQFilterParameters {
	rv := objc.Send[AudioUnitEQFilterParameters](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ AudioUnitEQFilterParameters) Init() AudioUnitEQFilterParameters {
	rv := objc.Send[AudioUnitEQFilterParameters](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ AudioUnitEQFilterParameters) Autorelease() AudioUnitEQFilterParameters {
	rv := objc.Send[AudioUnitEQFilterParameters](a_.ID, objc.Sel("autorelease"))
	return rv
}

// NewAudioUnitEQFilterParameters creates a new AudioUnitEQFilterParameters instance.
func NewAudioUnitEQFilterParameters() AudioUnitEQFilterParameters {
	return getAudioUnitEQFilterParametersClass().New()
}





// An object that encapsulates the parameters that the equalizer uses.


// An object that encapsulates the parameters that the equalizer uses.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioUnitEQFilterParameters
type AudioUnitEQFilterParameters struct {
	objectivec.Object
}

// AudioUnitEQFilterParametersFrom constructs a [AudioUnitEQFilterParameters] from an unsafe.Pointer.
//
// An object that encapsulates the parameters that the equalizer uses.
func AudioUnitEQFilterParametersFrom(ptr unsafe.Pointer) AudioUnitEQFilterParameters {
	return AudioUnitEQFilterParameters{objectivec.Object{objc.ID(ptr)}}
}

























// The bandwidth of the equalizer filter, in octaves.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioUnitEQFilterParameters/bandwidth
func (a_ AudioUnitEQFilterParameters) Bandwidth() float32 {
	rv := objc.Send[float32](a_.ID, objc.Sel("bandwidth"))
	return rv
}


// The bandwidth of the equalizer filter, in octaves.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioUnitEQFilterParameters/bandwidth
func (a_ AudioUnitEQFilterParameters) SetBandwidth(value float32) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setBandwidth:"), value)
}


// The bypass state of the equalizer filter band.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioUnitEQFilterParameters/bypass
func (a_ AudioUnitEQFilterParameters) Bypass() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("bypass"))
	return rv
}


// The bypass state of the equalizer filter band.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioUnitEQFilterParameters/bypass
func (a_ AudioUnitEQFilterParameters) SetBypass(value bool) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setBypass:"), value)
}


// The equalizer filter type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioUnitEQFilterParameters/filterType
func (a_ AudioUnitEQFilterParameters) FilterType() AudioUnitEQFilterType {
	rv := objc.Send[AudioUnitEQFilterType](a_.ID, objc.Sel("filterType"))
	return rv
}


// The equalizer filter type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioUnitEQFilterParameters/filterType
func (a_ AudioUnitEQFilterParameters) SetFilterType(value AudioUnitEQFilterType) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setFilterType:"), value)
}


// The frequency of the equalizer filter, in hertz.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioUnitEQFilterParameters/frequency
func (a_ AudioUnitEQFilterParameters) Frequency() float32 {
	rv := objc.Send[float32](a_.ID, objc.Sel("frequency"))
	return rv
}


// The frequency of the equalizer filter, in hertz.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioUnitEQFilterParameters/frequency
func (a_ AudioUnitEQFilterParameters) SetFrequency(value float32) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setFrequency:"), value)
}


// The gain of the equalizer filter, in decibels.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioUnitEQFilterParameters/gain
func (a_ AudioUnitEQFilterParameters) Gain() float32 {
	rv := objc.Send[float32](a_.ID, objc.Sel("gain"))
	return rv
}


// The gain of the equalizer filter, in decibels.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioUnitEQFilterParameters/gain
func (a_ AudioUnitEQFilterParameters) SetGain(value float32) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setGain:"), value)
}


// An array of equalizer filter parameters.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avaudiouniteq/bands
func (a_ AudioUnitEQFilterParameters) Bands() IAVAudioUnitEQFilterParameters {
	rv := objc.Send[AudioUnitEQFilterParameters](a_.ID, objc.Sel("bands"))
	return rv
}


// An array of equalizer filter parameters.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avaudiouniteq/bands
func (a_ AudioUnitEQFilterParameters) SetBands(value IAVAudioUnitEQFilterParameters) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setBands:"), value)
}


// The overall gain adjustment that the audio unit applies to the signal, in decibels.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avaudiouniteq/globalgain
func (a_ AudioUnitEQFilterParameters) GlobalGain() float32 {
	rv := objc.Send[float32](a_.ID, objc.Sel("globalGain"))
	return rv
}


// The overall gain adjustment that the audio unit applies to the signal, in decibels.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avaudiouniteq/globalgain
func (a_ AudioUnitEQFilterParameters) SetGlobalGain(value float32) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setGlobalGain:"), value)
}








