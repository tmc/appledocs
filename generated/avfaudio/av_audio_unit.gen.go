// Code generated from Apple documentation for AVFAudio. DO NOT EDIT.

package avfaudio

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)





// The class instance for the [AudioUnit] class.
var (
	AudioUnitClass     _AudioUnitClass
	AudioUnitClassOnce sync.Once
)

func getAudioUnitClass() _AudioUnitClass {
	AudioUnitClassOnce.Do(func() {
		AudioUnitClass = _AudioUnitClass{objc.GetClass("AVAudioUnit")}
	})
	return AudioUnitClass
}

type _AudioUnitClass struct {
	class objc.Class
}





// An interface definition for the [AudioUnit] class.
type IAudioUnit interface {
	IAudioNode
	

	// properties:
	AUAudioUnit() IAudioUnit
	AudioComponentDescription() audiotoolbox.AudioComponentDescription
	AudioUnit() audiotoolbox.AudioUnit
	ManufacturerName() objc.IObject /* cross-framework: NSString */
	Name() objc.IObject /* cross-framework: NSString */
	Version() uint


	

	// methods:
	LoadAudioUnitPresetAtURLError(url objc.IObject /* cross-framework: NSURL */, outError objectivec.IObject) bool


}





// Alloc allocates a new instance without initialization.
func (ac _AudioUnitClass) Alloc() AudioUnit {
	rv := objc.Send[AudioUnit](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (ac _AudioUnitClass) New() AudioUnit {
	rv := objc.Send[AudioUnit](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ AudioUnit) Init() AudioUnit {
	rv := objc.Send[AudioUnit](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ AudioUnit) Autorelease() AudioUnit {
	rv := objc.Send[AudioUnit](a_.ID, objc.Sel("autorelease"))
	return rv
}

// NewAudioUnit creates a new AudioUnit instance.
func NewAudioUnit() AudioUnit {
	return getAudioUnitClass().New()
}





// A subclass of the audio node class that, processes audio either in real time or nonreal time, depending on the type of the audio unit.


// A subclass of the audio node class that, processes audio either in real time or nonreal time, depending on the type of the audio unit.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioUnit
type AudioUnit struct {
	AudioNode
}

// AudioUnitFrom constructs a [AudioUnit] from an unsafe.Pointer.
//
// A subclass of the audio node class that, processes audio either in real time or nonreal time, depending on the type of the audio unit.
func AudioUnitFrom(ptr unsafe.Pointer) AudioUnit {
	return AudioUnit{
		AudioNode: AudioNodeFrom(ptr),
	}
}










// Creates an instance of an audio unit component asynchronously and wraps it in an audio unit class.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioUnit/instantiate(with:options:completionHandler:)
func (ac _AudioUnitClass) InstantiateWithComponentDescriptionOptionsCompletionHandler(audioComponentDescription audiotoolbox.AudioComponentDescription, options objectivec.IObject, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(ac.class), objc.Sel("instantiateWithComponentDescription:options:completionHandler:"), audioComponentDescription, options, completionHandler)
}












// Loads an audio unit using a specified preset.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioUnit/loadPreset(at:)
func (a_ AudioUnit) LoadAudioUnitPresetAtURLError(url objc.IObject /* cross-framework: NSURL */, outError objectivec.IObject) bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("loadAudioUnitPresetAtURL:error:"), url, outError)
	return rv
}







// The audio unit class wrapping or underlying the implementation’s audio unit.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioUnit/auAudioUnit
func (a_ AudioUnit) AUAudioUnit() IAudioUnit {
	rv := objc.Send[AudioUnit](a_.ID, objc.Sel("AUAudioUnit"))
	return rv
}


// The audio component description that represents the underlying Core Audio audio unit.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioUnit/audioComponentDescription
func (a_ AudioUnit) AudioComponentDescription() audiotoolbox.AudioComponentDescription {
	rv := objc.Send[audiotoolbox.AudioComponentDescription](a_.ID, objc.Sel("audioComponentDescription"))
	return rv
}


// The underlying Core Audio audio unit.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioUnit/audioUnit
func (a_ AudioUnit) AudioUnit() audiotoolbox.AudioUnit {
	rv := objc.Send[audiotoolbox.AudioUnit](a_.ID, objc.Sel("audioUnit"))
	return rv
}


// The name of the manufacturer of the audio unit.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioUnit/manufacturerName
func (a_ AudioUnit) ManufacturerName() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](a_.ID, objc.Sel("manufacturerName"))
	return rv
}


// The name of the audio unit.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioUnit/name
func (a_ AudioUnit) Name() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](a_.ID, objc.Sel("name"))
	return rv
}


// The version number of the audio unit.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioUnit/version
func (a_ AudioUnit) Version() uint {
	rv := objc.Send[uint](a_.ID, objc.Sel("version"))
	return rv
}








