// Code generated from Apple documentation for AudioToolbox. DO NOT EDIT.

package audiotoolbox

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class AUAudioUnitPreset */


/* debug [class_header]: Header for AUAudioUnitPreset */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for AudioUnitPreset */
// An interface definition for the [AudioUnitPreset] class.
type IAudioUnitPreset interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for AudioUnitPreset */
	// properties:
	Name() objc.IObject /* cross-framework: NSString */
	SetName(value objc.IObject /* cross-framework: NSString */)
	Number() int
	SetNumber(value int)
	FullState() objc.IObject /* cross-framework: NSString */
	SetFullState(value objc.IObject /* cross-framework: NSString */)
	FullStateForDocument() objc.IObject /* cross-framework: NSString */
	SetFullStateForDocument(value objc.IObject /* cross-framework: NSString */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for AudioUnitPreset */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for AudioUnitPreset */
// Alloc allocates a new instance without initialization.
func (ac _AudioUnitPresetClass) Alloc() AudioUnitPreset {
	rv := objc.Send[AudioUnitPreset](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for AudioUnitPreset */
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for AudioUnitPreset *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for AudioUnitPreset */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for AudioUnitPreset */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for AudioUnitPreset */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for AudioUnitPreset */

// The preset’s name.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUAudioUnitPreset/name
func (a_ AudioUnitPreset) Name() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](a_.ID, objc.Sel("name"))
	return rv
}/* debug [instance_properties/getter]: name */


// The preset’s name.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUAudioUnitPreset/name
func (a_ AudioUnitPreset) SetName(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setName:"), value)
}/* debug [instance_properties/setter]: name */


// The preset’s unique numeric identifier.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUAudioUnitPreset/number
func (a_ AudioUnitPreset) Number() int {
	rv := objc.Send[int](a_.ID, objc.Sel("number"))
	return rv
}/* debug [instance_properties/getter]: number */


// The preset’s unique numeric identifier.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUAudioUnitPreset/number
func (a_ AudioUnitPreset) SetNumber(value int) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setNumber:"), value)
}/* debug [instance_properties/setter]: number */


// A persistable snapshot of the audio unit’s properties and parameters, suitable for saving as a user preset.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/audiotoolbox/auaudiounit/fullstate
func (a_ AudioUnitPreset) FullState() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](a_.ID, objc.Sel("fullState"))
	return rv
}/* debug [instance_properties/getter]: fullState */


// A persistable snapshot of the audio unit’s properties and parameters, suitable for saving as a user preset.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/audiotoolbox/auaudiounit/fullstate
func (a_ AudioUnitPreset) SetFullState(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setFullState:"), value)
}/* debug [instance_properties/setter]: fullState */


// A persistable snapshot of the audio unit’s properties and parameters, suitable for saving in a user’s document.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/audiotoolbox/auaudiounit/fullstatefordocument
func (a_ AudioUnitPreset) FullStateForDocument() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](a_.ID, objc.Sel("fullStateForDocument"))
	return rv
}/* debug [instance_properties/getter]: fullStateForDocument */


// A persistable snapshot of the audio unit’s properties and parameters, suitable for saving in a user’s document.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/audiotoolbox/auaudiounit/fullstatefordocument
func (a_ AudioUnitPreset) SetFullStateForDocument(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setFullStateForDocument:"), value)
}/* debug [instance_properties/setter]: fullStateForDocument */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class AUAudioUnitPreset */



