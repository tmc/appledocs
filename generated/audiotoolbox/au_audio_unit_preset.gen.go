// Code generated from Apple documentation for AudioToolbox. DO NOT EDIT.

package audiotoolbox

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
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
	// properties:
	Name() objc.IObject /* cross-framework: NSString */
	SetName(value objc.IObject /* cross-framework: NSString */)
	Number() int
	SetNumber(value int)
	FullState() objc.IObject /* cross-framework: NSString */
	SetFullState(value objc.IObject /* cross-framework: NSString */)
	FullStateForDocument() objc.IObject /* cross-framework: NSString */
	SetFullStateForDocument(value objc.IObject /* cross-framework: NSString */)
	// methods:
}

// A class that describes an interface for custom parameter settings provided by the audio unit developer.
//
// These presets often produce a useful sound or starting point. For more details on working with Audio Unit presets, see Note that the version 3 property is bridged to the version 2 API. Similarly, the version 3 property is bridged to the version 2 API.


// A class that describes an interface for custom parameter settings provided by the audio unit developer.
//
// [Full Topic]
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
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUAudioUnitPreset/name
func (a_ AudioUnitPreset) Name() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](a_.ID, objc.Sel("name"))
	return rv
}


// The preset’s name.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUAudioUnitPreset/name
func (a_ AudioUnitPreset) SetName(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setName:"), value)
}


// The preset’s unique numeric identifier.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUAudioUnitPreset/number
func (a_ AudioUnitPreset) Number() int {
	rv := objc.Send[int](a_.ID, objc.Sel("number"))
	return rv
}


// The preset’s unique numeric identifier.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUAudioUnitPreset/number
func (a_ AudioUnitPreset) SetNumber(value int) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setNumber:"), value)
}


// A persistable snapshot of the audio unit’s properties and parameters, suitable for saving as a user preset.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/audiotoolbox/auaudiounit/fullstate
func (a_ AudioUnitPreset) FullState() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](a_.ID, objc.Sel("fullState"))
	return rv
}


// A persistable snapshot of the audio unit’s properties and parameters, suitable for saving as a user preset.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/audiotoolbox/auaudiounit/fullstate
func (a_ AudioUnitPreset) SetFullState(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setFullState:"), value)
}


// A persistable snapshot of the audio unit’s properties and parameters, suitable for saving in a user’s document.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/audiotoolbox/auaudiounit/fullstatefordocument
func (a_ AudioUnitPreset) FullStateForDocument() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](a_.ID, objc.Sel("fullStateForDocument"))
	return rv
}


// A persistable snapshot of the audio unit’s properties and parameters, suitable for saving in a user’s document.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/audiotoolbox/auaudiounit/fullstatefordocument
func (a_ AudioUnitPreset) SetFullStateForDocument(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setFullStateForDocument:"), value)
}



