// Code generated from Apple documentation for AudioToolbox. DO NOT EDIT.

package audiotoolbox

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [AudioUnitPreset] class.
var (
	AudioUnitPresetClass     _AudioUnitPresetClass
	AudioUnitPresetClassOnce sync.Once
)

func getAudioUnitPresetClass() _AudioUnitPresetClass {
	AudioUnitPresetClassOnce.Do(func() {
		AudioUnitPresetClass = _AudioUnitPresetClass{objc.GetClass("AUAudioUnitPreset")}
	})
	return AudioUnitPresetClass
}

type _AudioUnitPresetClass struct {
	class objc.Class
}

// An interface definition for the [AudioUnitPreset] class.
type IAudioUnitPreset interface {
	objectivec.IObject
}

// A class that describes an interface for custom parameter settings provided by the audio unit developer.
//
// These presets often produce a useful sound or starting point. For more details on working with Audio Unit presets, see Note that the version 3 property is bridged to the version 2 API. Similarly, the version 3 property is bridged to the version 2 API.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUAudioUnitPreset
type AudioUnitPreset struct {
	objectivec.Object
}

// AudioUnitPresetFrom constructs a [AudioUnitPreset] from an unsafe.Pointer.
//
// A class that describes an interface for custom parameter settings provided by the audio unit developer.
func AudioUnitPresetFrom(ptr unsafe.Pointer) AudioUnitPreset {
	return AudioUnitPreset{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (ac _AudioUnitPresetClass) Alloc() AudioUnitPreset {
	rv := objc.Send[AudioUnitPreset](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ac _AudioUnitPresetClass) New() AudioUnitPreset {
	rv := objc.Send[AudioUnitPreset](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ AudioUnitPreset) Init() AudioUnitPreset {
	rv := objc.Send[AudioUnitPreset](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ AudioUnitPreset) Autorelease() AudioUnitPreset {
	rv := objc.Send[AudioUnitPreset](a_.ID, objc.Sel("autorelease"))
	return rv
}

// NewAudioUnitPreset creates a new AudioUnitPreset instance.
func NewAudioUnitPreset() AudioUnitPreset {
	return getAudioUnitPresetClass().New()
}


// The preset’s name.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUAudioUnitPreset/name
func (a_ AudioUnitPreset) Name() string {
	rv := objc.Send[string](a_.ID, objc.Sel("name"))
	return rv
}


// SetName sets the value of the name property.
// The preset’s name.

//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUAudioUnitPreset/name
func (a_ AudioUnitPreset) SetName(value string) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setName:"), objc.String(value))
}

// The preset’s unique numeric identifier.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUAudioUnitPreset/number
func (a_ AudioUnitPreset) Number() int {
	rv := objc.Send[int](a_.ID, objc.Sel("number"))
	return rv
}


// SetNumber sets the value of the number property.
// The preset’s unique numeric identifier.

//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUAudioUnitPreset/number
func (a_ AudioUnitPreset) SetNumber(value int) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setNumber:"), value)
}



