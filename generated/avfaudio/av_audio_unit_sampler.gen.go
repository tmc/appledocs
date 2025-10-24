// Code generated from Apple documentation for AVFAudio. DO NOT EDIT.

package avfaudio

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)





// The class instance for the [AudioUnitSampler] class.
var (
	AudioUnitSamplerClass     _AudioUnitSamplerClass
	AudioUnitSamplerClassOnce sync.Once
)

func getAudioUnitSamplerClass() _AudioUnitSamplerClass {
	AudioUnitSamplerClassOnce.Do(func() {
		AudioUnitSamplerClass = _AudioUnitSamplerClass{objc.GetClass("AVAudioUnitSampler")}
	})
	return AudioUnitSamplerClass
}

type _AudioUnitSamplerClass struct {
	class objc.Class
}





// An interface definition for the [AudioUnitSampler] class.
type IAudioUnitSampler interface {
	IAudioUnitMIDIInstrument
	

	// properties:
	GlobalTuning() float32
	SetGlobalTuning(value float32)
	MasterGain() float32
	SetMasterGain(value float32)
	OverallGain() float32
	SetOverallGain(value float32)
	StereoPan() float32
	SetStereoPan(value float32)


	

	// methods:
	LoadAudioFilesAtURLsError(audioFiles []foundation.URL, outError objectivec.IObject) bool
	LoadInstrumentAtURLError(instrumentURL objc.IObject /* cross-framework: NSURL */, outError objectivec.IObject) bool
	LoadSoundBankInstrumentAtURLProgramBankMSBBankLSBError(bankURL objc.IObject /* cross-framework: NSURL */, program uint8 /* not a class type */, bankMSB uint8 /* not a class type */, bankLSB uint8 /* not a class type */, outError objectivec.IObject) bool


}





// Alloc allocates a new instance without initialization.
func (ac _AudioUnitSamplerClass) Alloc() AudioUnitSampler {
	rv := objc.Send[AudioUnitSampler](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (ac _AudioUnitSamplerClass) New() AudioUnitSampler {
	rv := objc.Send[AudioUnitSampler](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ AudioUnitSampler) Init() AudioUnitSampler {
	rv := objc.Send[AudioUnitSampler](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ AudioUnitSampler) Autorelease() AudioUnitSampler {
	rv := objc.Send[AudioUnitSampler](a_.ID, objc.Sel("autorelease"))
	return rv
}

// NewAudioUnitSampler creates a new AudioUnitSampler instance.
func NewAudioUnitSampler() AudioUnitSampler {
	return getAudioUnitSamplerClass().New()
}





// An object that you configure with one or more instrument samples, based on Apple’s Sampler audio unit.
//
// An is an for Apple’s Sampler audio unit. You configure the sampler by loading instruments from different types of files. These include an file, DLS, or SF2 sound bank; an EXS24 instrument; a single audio file; or an array of audio files. The output of a is a single stereo bus.


// An object that you configure with one or more instrument samples, based on Apple’s Sampler audio unit.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioUnitSampler
type AudioUnitSampler struct {
	AudioUnitMIDIInstrument
}

// AudioUnitSamplerFrom constructs a [AudioUnitSampler] from an unsafe.Pointer.
//
// An object that you configure with one or more instrument samples, based on Apple’s Sampler audio unit.
func AudioUnitSamplerFrom(ptr unsafe.Pointer) AudioUnitSampler {
	return AudioUnitSampler{
		AudioUnitMIDIInstrument: AudioUnitMIDIInstrumentFrom(ptr),
	}
}




















// Configures the sampler by loading the specified audio files.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioUnitSampler/loadAudioFiles(at:)
func (a_ AudioUnitSampler) LoadAudioFilesAtURLsError(audioFiles []foundation.URL, outError objectivec.IObject) bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("loadAudioFilesAtURLs:error:"), audioFiles, outError)
	return rv
}


// Configures the sampler with the specified instrument file.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioUnitSampler/loadInstrument(at:)
func (a_ AudioUnitSampler) LoadInstrumentAtURLError(instrumentURL objc.IObject /* cross-framework: NSURL */, outError objectivec.IObject) bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("loadInstrumentAtURL:error:"), instrumentURL, outError)
	return rv
}


// Loads a specific instrument from the specified soundbank.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioUnitSampler/loadSoundBankInstrument(at:program:bankMSB:bankLSB:)
func (a_ AudioUnitSampler) LoadSoundBankInstrumentAtURLProgramBankMSBBankLSBError(bankURL objc.IObject /* cross-framework: NSURL */, program uint8 /* not a class type */, bankMSB uint8 /* not a class type */, bankLSB uint8 /* not a class type */, outError objectivec.IObject) bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("loadSoundBankInstrumentAtURL:program:bankMSB:bankLSB:error:"), bankURL, program, bankMSB, bankLSB, outError)
	return rv
}







// An adjustment for the tuning of all the played notes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioUnitSampler/globalTuning
func (a_ AudioUnitSampler) GlobalTuning() float32 {
	rv := objc.Send[float32](a_.ID, objc.Sel("globalTuning"))
	return rv
}


// An adjustment for the tuning of all the played notes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioUnitSampler/globalTuning
func (a_ AudioUnitSampler) SetGlobalTuning(value float32) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setGlobalTuning:"), value)
}


// An adjustment for the gain of all the played notes, in decibels.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioUnitSampler/masterGain
func (a_ AudioUnitSampler) MasterGain() float32 {
	rv := objc.Send[float32](a_.ID, objc.Sel("masterGain"))
	return rv
}


// An adjustment for the gain of all the played notes, in decibels.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioUnitSampler/masterGain
func (a_ AudioUnitSampler) SetMasterGain(value float32) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setMasterGain:"), value)
}


// An adjustment for the gain of all the played notes, in decibels.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioUnitSampler/overallGain
func (a_ AudioUnitSampler) OverallGain() float32 {
	rv := objc.Send[float32](a_.ID, objc.Sel("overallGain"))
	return rv
}


// An adjustment for the gain of all the played notes, in decibels.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioUnitSampler/overallGain
func (a_ AudioUnitSampler) SetOverallGain(value float32) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setOverallGain:"), value)
}


// An adjustment for the stereo panning of all the played notes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioUnitSampler/stereoPan
func (a_ AudioUnitSampler) StereoPan() float32 {
	rv := objc.Send[float32](a_.ID, objc.Sel("stereoPan"))
	return rv
}


// An adjustment for the stereo panning of all the played notes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioUnitSampler/stereoPan
func (a_ AudioUnitSampler) SetStereoPan(value float32) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setStereoPan:"), value)
}








