// Code generated from Apple documentation for AVFAudio. DO NOT EDIT.

package avfaudio

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/audiotoolbox"
	"github.com/tmc/appledocs/generated/foundation"
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
	AuAudioUnit() IAudioUnit
	SetAuAudioUnit(value IAudioUnit)
	AudioComponentDescription() objc.IObject /* cross-framework: AudioComponentDescription */
	SetAudioComponentDescription(value objc.IObject /* cross-framework: AudioComponentDescription */)
	AudioUnit() audiotoolbox.IAudioUnit
	SetAudioUnit(value audiotoolbox.IAudioUnit)
	ManufacturerName() objc.IObject /* cross-framework: NSString */
	SetManufacturerName(value objc.IObject /* cross-framework: NSString */)
	Name() objc.IObject /* cross-framework: NSString */
	SetName(value objc.IObject /* cross-framework: NSString */)
	Version() int
	SetVersion(value int)
	// methods:
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

// Alloc allocates a new instance without initialization.
func (ac _AudioUnitClass) Alloc() AudioUnit {
	rv := objc.Send[AudioUnit](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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



// Creates an instance of an audio unit component asynchronously and wraps it in an audio unit class.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioUnit/instantiate(with:options:completionHandler:)
func (ac _AudioUnitClass) InstantiateWithComponentDescriptionOptionsCompletionHandler(audioComponentDescription objc.IObject /* cross-framework: AudioComponentDescription */, options unsafe.Pointer, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(ac.class), objc.Sel("instantiateWithComponentDescription:options:completionHandler:"), audioComponentDescription, options, completionHandler)
}


// The audio unit class wrapping or underlying the implementation’s audio unit.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avaudiounit/auaudiounit
func (a_ AudioUnit) AuAudioUnit() IAudioUnit {
	rv := objc.Send[AudioUnit](a_.ID, objc.Sel("auAudioUnit"))
	return rv
}


// The audio unit class wrapping or underlying the implementation’s audio unit.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avaudiounit/auaudiounit
func (a_ AudioUnit) SetAuAudioUnit(value IAudioUnit) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setAuAudioUnit:"), value)
}


// The audio component description that represents the underlying Core Audio audio unit.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avaudiounit/audiocomponentdescription
func (a_ AudioUnit) AudioComponentDescription() objc.IObject /* cross-framework: AudioComponentDescription */ {
	rv := objc.Send[audiotoolbox.AudioComponentDescription](a_.ID, objc.Sel("audioComponentDescription"))
	return rv
}


// The audio component description that represents the underlying Core Audio audio unit.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avaudiounit/audiocomponentdescription
func (a_ AudioUnit) SetAudioComponentDescription(value objc.IObject /* cross-framework: AudioComponentDescription */) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setAudioComponentDescription:"), value)
}


// The underlying Core Audio audio unit.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avaudiounit/audiounit
func (a_ AudioUnit) AudioUnit() audiotoolbox.IAudioUnit {
	rv := objc.Send[audiotoolbox.AudioUnit](a_.ID, objc.Sel("audioUnit"))
	return rv
}


// The underlying Core Audio audio unit.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avaudiounit/audiounit
func (a_ AudioUnit) SetAudioUnit(value audiotoolbox.IAudioUnit) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setAudioUnit:"), value)
}


// The name of the manufacturer of the audio unit.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avaudiounit/manufacturername
func (a_ AudioUnit) ManufacturerName() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](a_.ID, objc.Sel("manufacturerName"))
	return rv
}


// The name of the manufacturer of the audio unit.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avaudiounit/manufacturername
func (a_ AudioUnit) SetManufacturerName(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setManufacturerName:"), value)
}


// The name of the audio unit.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avaudiounit/name
func (a_ AudioUnit) Name() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](a_.ID, objc.Sel("name"))
	return rv
}


// The name of the audio unit.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avaudiounit/name
func (a_ AudioUnit) SetName(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setName:"), value)
}


// The version number of the audio unit.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avaudiounit/version
func (a_ AudioUnit) Version() int {
	rv := objc.Send[int](a_.ID, objc.Sel("version"))
	return rv
}


// The version number of the audio unit.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avaudiounit/version
func (a_ AudioUnit) SetVersion(value int) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setVersion:"), value)
}



