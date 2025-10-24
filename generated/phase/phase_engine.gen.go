// Code generated from Apple documentation for PHASE. DO NOT EDIT.

package phase

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/avfaudio"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class PHASEEngine */


/* debug [class_header]: Header for PHASEEngine */
// The class instance for the [PHASEEngine] class.
var (
	PHASEEngineClass     _PHASEEngineClass
	PHASEEngineClassOnce sync.Once
)

func getPHASEEngineClass() _PHASEEngineClass {
	PHASEEngineClassOnce.Do(func() {
		PHASEEngineClass = _PHASEEngineClass{objc.GetClass("PHASEEngine")}
	})
	return PHASEEngineClass
}

type _PHASEEngineClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for PHASEEngine */
// An interface definition for the [PHASEEngine] class.
type IPHASEEngine interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for PHASEEngine */
	// properties:
	ActiveGroupPreset() IPHASEGroupPreset
	AssetRegistry() IPHASEAssetRegistry
	DefaultMedium() IPHASEMedium
	SetDefaultMedium(value IPHASEMedium)
	DefaultReverbPreset() PHASEReverbPreset
	SetDefaultReverbPreset(value PHASEReverbPreset)
	Duckers() []PHASEDucker
	Groups() foundation.IDictionary
	LastRenderTime() avfaudio.AudioTime
	OutputSpatializationMode() PHASESpatializationMode
	SetOutputSpatializationMode(value PHASESpatializationMode)
	RenderingState() PHASERenderingState
	RootObject() IPHASEObject
	SoundEvents() []PHASESoundEvent
	UnitsPerMeter() float64
	SetUnitsPerMeter(value float64)
	UnitsPerSecond() float64
	SetUnitsPerSecond(value float64)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for PHASEEngine */
	// methods:
	Pause()
	StartAndReturnError(error_ unsafe.Pointer) bool
	Stop()
	Update()
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for PHASEEngine */
// Alloc allocates a new instance without initialization.
func (pc _PHASEEngineClass) Alloc() PHASEEngine {
	rv := objc.Send[PHASEEngine](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (pc _PHASEEngineClass) New() PHASEEngine {
	rv := objc.Send[PHASEEngine](objc.ID(pc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (p_ PHASEEngine) Init() PHASEEngine {
	rv := objc.Send[PHASEEngine](p_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p_ PHASEEngine) Autorelease() PHASEEngine {
	rv := objc.Send[PHASEEngine](p_.ID, objc.Sel("autorelease"))
	return rv
}

// NewPHASEEngine creates a new PHASEEngine instance.
func NewPHASEEngine() PHASEEngine {
	return getPHASEEngineClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for PHASEEngine */
// An object that manages audio assets, controls playback, and configures environmental effects.
//
// Before using PHASE, an app creates an instance of this object. Apps access all of the framework’s functionality through engine functions or properties, or through other PHASE classes into which you pass the engine object.


// An object that manages audio assets, controls playback, and configures environmental effects.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASEEngine
type PHASEEngine struct {
	objectivec.Object
}

// PHASEEngineFrom constructs a [PHASEEngine] from an unsafe.Pointer.
//
// An object that manages audio assets, controls playback, and configures environmental effects.
func PHASEEngineFrom(ptr unsafe.Pointer) PHASEEngine {
	return PHASEEngine{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for PHASEEngine */

// Creates an engine updated by the app or framework.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASEEngine/init(updateMode:)
func NewPHASEEngineWithUpdateMode(updateMode PHASEUpdateMode) PHASEEngine {
	instance := getPHASEEngineClass().Alloc()
	rv := objc.Send[PHASEEngine](instance.ID, objc.Sel("initWithUpdateMode:"), updateMode)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewPHASEEngineWithUpdateMode */


// Creates a new engine that has both update and rendering modes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASEEngine/init(updateMode:renderingMode:)
func NewPHASEEngineWithUpdateModeRenderingMode(updateMode PHASEUpdateMode, renderingMode PHASERenderingMode) PHASEEngine {
	instance := getPHASEEngineClass().Alloc()
	rv := objc.Send[PHASEEngine](instance.ID, objc.Sel("initWithUpdateMode:renderingMode:"), updateMode, renderingMode)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewPHASEEngineWithUpdateModeRenderingMode */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for PHASEEngine */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for PHASEEngine */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for PHASEEngine */

// Pauses all audio playback.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASEEngine/pause()
func (p_ PHASEEngine) Pause() {
	objc.Send[objc.ID](p_.ID, objc.Sel("pause"))
}/* debug [instance_methods/method]: Pause */


// Starts or resumes all audio playback.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASEEngine/start()
func (p_ PHASEEngine) StartAndReturnError(error_ unsafe.Pointer) bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("startAndReturnError:"), error_)
	return rv
}/* debug [instance_methods/method]: StartAndReturnError */


// Stops all audio playback.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASEEngine/stop()
func (p_ PHASEEngine) Stop() {
	objc.Send[objc.ID](p_.ID, objc.Sel("stop"))
}/* debug [instance_methods/method]: Stop */


// Processes app commands and increments framework processing.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASEEngine/update()
func (p_ PHASEEngine) Update() {
	objc.Send[objc.ID](p_.ID, objc.Sel("update"))
}/* debug [instance_methods/method]: Update */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for PHASEEngine */

// The settings that define playback for a group of sounds.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASEEngine/activeGroupPreset
func (p_ PHASEEngine) ActiveGroupPreset() IPHASEGroupPreset {
	rv := objc.Send[PHASEGroupPreset](p_.ID, objc.Sel("activeGroupPreset"))
	return rv
}/* debug [instance_properties/getter]: activeGroupPreset */


// An object that loads and unloads audio resources.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASEEngine/assetRegistry
func (p_ PHASEEngine) AssetRegistry() IPHASEAssetRegistry {
	rv := objc.Send[PHASEAssetRegistry](p_.ID, objc.Sel("assetRegistry"))
	return rv
}/* debug [instance_properties/getter]: assetRegistry */


// The physical matter through which sound travels.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASEEngine/defaultMedium
func (p_ PHASEEngine) DefaultMedium() IPHASEMedium {
	rv := objc.Send[PHASEMedium](p_.ID, objc.Sel("defaultMedium"))
	return rv
}/* debug [instance_properties/getter]: defaultMedium */


// The physical matter through which sound travels.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASEEngine/defaultMedium
func (p_ PHASEEngine) SetDefaultMedium(value IPHASEMedium) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setDefaultMedium:"), value)
}/* debug [instance_properties/setter]: defaultMedium */


// The environmental surroundings that determine how sound resonates.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASEEngine/defaultReverbPreset
func (p_ PHASEEngine) DefaultReverbPreset() PHASEReverbPreset {
	rv := objc.Send[PHASEReverbPreset](p_.ID, objc.Sel("defaultReverbPreset"))
	return rv
}/* debug [instance_properties/getter]: defaultReverbPreset */


// The environmental surroundings that determine how sound resonates.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASEEngine/defaultReverbPreset
func (p_ PHASEEngine) SetDefaultReverbPreset(value PHASEReverbPreset) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setDefaultReverbPreset:"), value)
}/* debug [instance_properties/setter]: defaultReverbPreset */


// An array of objects that reduce the volume of simultaneously playing sounds.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASEEngine/duckers
func (p_ PHASEEngine) Duckers() []PHASEDucker {
	rv := objc.Send[[]PHASEDucker](p_.ID, objc.Sel("duckers"))
	return rv
}/* debug [instance_properties/getter]: duckers */


// A list of named groups that contain sounds the app operates on collectively.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASEEngine/groups
func (p_ PHASEEngine) Groups() foundation.IDictionary {
	rv := objc.Send[foundation.IDictionary](p_.ID, objc.Sel("groups"))
	return rv
}/* debug [instance_properties/getter]: groups */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASEEngine/lastRenderTime
func (p_ PHASEEngine) LastRenderTime() avfaudio.AudioTime {
	rv := objc.Send[avfaudio.AudioTime](p_.ID, objc.Sel("lastRenderTime"))
	return rv
}/* debug [instance_properties/getter]: lastRenderTime */


// The mode the engine implements to create a 3D sound experience.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASEEngine/outputSpatializationMode
func (p_ PHASEEngine) OutputSpatializationMode() PHASESpatializationMode {
	rv := objc.Send[PHASESpatializationMode](p_.ID, objc.Sel("outputSpatializationMode"))
	return rv
}/* debug [instance_properties/getter]: outputSpatializationMode */


// The mode the engine implements to create a 3D sound experience.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASEEngine/outputSpatializationMode
func (p_ PHASEEngine) SetOutputSpatializationMode(value PHASESpatializationMode) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setOutputSpatializationMode:"), value)
}/* debug [instance_properties/setter]: outputSpatializationMode */


// The status of the engine’s audio playback.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASEEngine/renderingState
func (p_ PHASEEngine) RenderingState() PHASERenderingState {
	rv := objc.Send[PHASERenderingState](p_.ID, objc.Sel("renderingState"))
	return rv
}/* debug [instance_properties/getter]: renderingState */


// The main object to which the app adds child objects.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASEEngine/rootObject
func (p_ PHASEEngine) RootObject() IPHASEObject {
	rv := objc.Send[PHASEObject](p_.ID, objc.Sel("rootObject"))
	return rv
}/* debug [instance_properties/getter]: rootObject */


// A collection of the sounds that play under various runtime circumstances.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASEEngine/soundEvents
func (p_ PHASEEngine) SoundEvents() []PHASESoundEvent {
	rv := objc.Send[[]PHASESoundEvent](p_.ID, objc.Sel("soundEvents"))
	return rv
}/* debug [instance_properties/getter]: soundEvents */


// A conversion factor from meters to your app’s preferred unit of measurement.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASEEngine/unitsPerMeter
func (p_ PHASEEngine) UnitsPerMeter() float64 {
	rv := objc.Send[float64](p_.ID, objc.Sel("unitsPerMeter"))
	return rv
}/* debug [instance_properties/getter]: unitsPerMeter */


// A conversion factor from meters to your app’s preferred unit of measurement.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASEEngine/unitsPerMeter
func (p_ PHASEEngine) SetUnitsPerMeter(value float64) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setUnitsPerMeter:"), value)
}/* debug [instance_properties/setter]: unitsPerMeter */


// A conversion factor from seconds to your app’s preferred unit of time.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASEEngine/unitsPerSecond
func (p_ PHASEEngine) UnitsPerSecond() float64 {
	rv := objc.Send[float64](p_.ID, objc.Sel("unitsPerSecond"))
	return rv
}/* debug [instance_properties/getter]: unitsPerSecond */


// A conversion factor from seconds to your app’s preferred unit of time.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASEEngine/unitsPerSecond
func (p_ PHASEEngine) SetUnitsPerSecond(value float64) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setUnitsPerSecond:"), value)
}/* debug [instance_properties/setter]: unitsPerSecond */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class PHASEEngine */


