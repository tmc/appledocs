// Code generated from Apple documentation for AVFAudio. DO NOT EDIT.

package avfaudio

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
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
	LoadAudioUnitPresetAtURLError(url foundation.IURL, outError unsafe.Pointer) bool
	AudioUnit() AudioUnit
	AuAudioUnit() AudioUnit
	SetAuAudioUnit(value IAudioUnit)
	AudioComponentDescription() unsafe.Pointer
	SetAudioComponentDescription(value unsafe.Pointer)
	ManufacturerName() string
	SetManufacturerName(value string)
	Name() string
	SetName(value string)
	Version() int
	SetVersion(value int)
}

// A subclass of the audio node class that, processes audio either in real time or nonreal time, depending on the type of the audio unit.
//
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


// Loads an audio unit using a specified preset.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioUnit/loadPreset(at:)
func (a_ AudioUnit) LoadAudioUnitPresetAtURLError(url foundation.IURL, outError unsafe.Pointer) bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("loadAudioUnitPresetAtURL:error:"), url, outError)
	return rv
}

// The underlying Core Audio audio unit.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioUnit/audioUnit
func (a_ AudioUnit) AudioUnit() AudioUnit {
	rv := objc.Send[AudioUnit](a_.ID, objc.Sel("audioUnit"))
	return rv
}

// The audio unit class wrapping or underlying the implementation’s audio unit.
//
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avaudiounit/auaudiounit
func (a_ AudioUnit) AuAudioUnit() AudioUnit {
	rv := objc.Send[AudioUnit](a_.ID, objc.Sel("auAudioUnit"))
	return rv
}


// SetAuAudioUnit sets the value of the auAudioUnit property.
// The audio unit class wrapping or underlying the implementation’s audio unit.

//
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avaudiounit/auaudiounit
func (a_ AudioUnit) SetAuAudioUnit(value IAudioUnit) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setAuAudioUnit:"), value)
}

// The audio component description that represents the underlying Core Audio audio unit.
//
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avaudiounit/audiocomponentdescription
func (a_ AudioUnit) AudioComponentDescription() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("audioComponentDescription"))
	return rv
}


// SetAudioComponentDescription sets the value of the audioComponentDescription property.
// The audio component description that represents the underlying Core Audio audio unit.

//
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avaudiounit/audiocomponentdescription
func (a_ AudioUnit) SetAudioComponentDescription(value unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setAudioComponentDescription:"), value)
}

// The name of the manufacturer of the audio unit.
//
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avaudiounit/manufacturername
func (a_ AudioUnit) ManufacturerName() string {
	rv := objc.Send[string](a_.ID, objc.Sel("manufacturerName"))
	return rv
}


// SetManufacturerName sets the value of the manufacturerName property.
// The name of the manufacturer of the audio unit.

//
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avaudiounit/manufacturername
func (a_ AudioUnit) SetManufacturerName(value string) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setManufacturerName:"), objc.String(value))
}

// The name of the audio unit.
//
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avaudiounit/name
func (a_ AudioUnit) Name() string {
	rv := objc.Send[string](a_.ID, objc.Sel("name"))
	return rv
}


// SetName sets the value of the name property.
// The name of the audio unit.

//
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avaudiounit/name
func (a_ AudioUnit) SetName(value string) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setName:"), objc.String(value))
}

// The version number of the audio unit.
//
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avaudiounit/version
func (a_ AudioUnit) Version() int {
	rv := objc.Send[int](a_.ID, objc.Sel("version"))
	return rv
}


// SetVersion sets the value of the version property.
// The version number of the audio unit.

//
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avaudiounit/version
func (a_ AudioUnit) SetVersion(value int) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setVersion:"), value)
}



