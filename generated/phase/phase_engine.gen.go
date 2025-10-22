// Code generated from Apple documentation for PHASE. DO NOT EDIT.

package phase

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/avfaudio"
	"github.com/tmc/appledocs/generated/objectivec"
)

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

// An interface definition for the [PHASEEngine] class.
type IPHASEEngine interface {
	objectivec.IObject
	Pause()
	StartAndReturnError(error_ unsafe.Pointer) bool
	Stop()
	Update()
	ActiveGroupPreset() PHASEGroupPreset
	AssetRegistry() PHASEAssetRegistry
	DefaultMedium() PHASEMedium
	SetDefaultMedium(value IPHASEMedium)
	DefaultReverbPreset() PHASEReverbPreset
	SetDefaultReverbPreset(value IPHASEReverbPreset)
	Duckers() []PHASEDucker
	Groups() unsafe.Pointer
	LastRenderTime() avfaudio.AudioTime
	OutputSpatializationMode() PHASESpatializationMode
	SetOutputSpatializationMode(value PHASESpatializationMode)
	RenderingState() PHASERenderingState
	RootObject() PHASEObject
	SoundEvents() []PHASESoundEvent
	UnitsPerMeter() float64
	SetUnitsPerMeter(value float64)
	UnitsPerSecond() float64
	SetUnitsPerSecond(value float64)
}

// An object that manages audio assets, controls playback, and configures environmental effects.
//
// Before using PHASE, an app creates an instance of this object. Apps access all of the framework’s functionality through engine functions or properties, or through other PHASE classes into which you pass the engine object.
//
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

// Alloc allocates a new instance without initialization.
func (pc _PHASEEngineClass) Alloc() PHASEEngine {
	rv := objc.Send[PHASEEngine](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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




// Creates an engine updated by the app or framework.
//
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASEEngine/init(updateMode:)
func NewPHASEEngineWithUpdateMode(updateMode PHASEUpdateMode) PHASEEngine {
	instance := getPHASEEngineClass().Alloc()
	rv := objc.Send[PHASEEngine](instance.ID, objc.Sel("initWithUpdateMode:"), updateMode)
	rv.Autorelease()
	return rv
}



// Creates a new engine that has both update and rendering modes.
//
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASEEngine/init(updateMode:renderingMode:)
func NewPHASEEngineWithUpdateModeRenderingMode(updateMode PHASEUpdateMode, renderingMode PHASERenderingMode) PHASEEngine {
	instance := getPHASEEngineClass().Alloc()
	rv := objc.Send[PHASEEngine](instance.ID, objc.Sel("initWithUpdateMode:renderingMode:"), updateMode, renderingMode)
	rv.Autorelease()
	return rv
}


// Pauses all audio playback.
//
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASEEngine/pause()
func (p_ PHASEEngine) Pause() {
	objc.Send[objc.ID](p_.ID, objc.Sel("pause"))
}

// Starts or resumes all audio playback.
//
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASEEngine/start()
func (p_ PHASEEngine) StartAndReturnError(error_ unsafe.Pointer) bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("startAndReturnError:"), error_)
	return rv
}

// Stops all audio playback.
//
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASEEngine/stop()
func (p_ PHASEEngine) Stop() {
	objc.Send[objc.ID](p_.ID, objc.Sel("stop"))
}

// Processes app commands and increments framework processing.
//
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASEEngine/update()
func (p_ PHASEEngine) Update() {
	objc.Send[objc.ID](p_.ID, objc.Sel("update"))
}

// The settings that define playback for a group of sounds.
//
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASEEngine/activeGroupPreset
func (p_ PHASEEngine) ActiveGroupPreset() PHASEGroupPreset {
	rv := objc.Send[PHASEGroupPreset](p_.ID, objc.Sel("activeGroupPreset"))
	return rv
}

// An object that loads and unloads audio resources.
//
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASEEngine/assetRegistry
func (p_ PHASEEngine) AssetRegistry() PHASEAssetRegistry {
	rv := objc.Send[PHASEAssetRegistry](p_.ID, objc.Sel("assetRegistry"))
	return rv
}

// The physical matter through which sound travels.
//
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASEEngine/defaultMedium
func (p_ PHASEEngine) DefaultMedium() PHASEMedium {
	rv := objc.Send[PHASEMedium](p_.ID, objc.Sel("defaultMedium"))
	return rv
}


// SetDefaultMedium sets the value of the defaultMedium property.
// The physical matter through which sound travels.

//
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASEEngine/defaultMedium
func (p_ PHASEEngine) SetDefaultMedium(value IPHASEMedium) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setDefaultMedium:"), value)
}

// The environmental surroundings that determine how sound resonates.
//
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASEEngine/defaultReverbPreset
func (p_ PHASEEngine) DefaultReverbPreset() PHASEReverbPreset {
	rv := objc.Send[PHASEReverbPreset](p_.ID, objc.Sel("defaultReverbPreset"))
	return rv
}


// SetDefaultReverbPreset sets the value of the defaultReverbPreset property.
// The environmental surroundings that determine how sound resonates.

//
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASEEngine/defaultReverbPreset
func (p_ PHASEEngine) SetDefaultReverbPreset(value IPHASEReverbPreset) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setDefaultReverbPreset:"), value)
}

// An array of objects that reduce the volume of simultaneously playing sounds.
//
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASEEngine/duckers
func (p_ PHASEEngine) Duckers() []PHASEDucker {
	rv := objc.Send[[]PHASEDucker](p_.ID, objc.Sel("duckers"))
	return rv
}

// A list of named groups that contain sounds the app operates on collectively.
//
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASEEngine/groups
func (p_ PHASEEngine) Groups() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("groups"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASEEngine/lastRenderTime
func (p_ PHASEEngine) LastRenderTime() avfaudio.AudioTime {
	rv := objc.Send[avfaudio.AudioTime](p_.ID, objc.Sel("lastRenderTime"))
	return rv
}

// The mode the engine implements to create a 3D sound experience.
//
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASEEngine/outputSpatializationMode
func (p_ PHASEEngine) OutputSpatializationMode() PHASESpatializationMode {
	rv := objc.Send[PHASESpatializationMode](p_.ID, objc.Sel("outputSpatializationMode"))
	return rv
}


// SetOutputSpatializationMode sets the value of the outputSpatializationMode property.
// The mode the engine implements to create a 3D sound experience.

//
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASEEngine/outputSpatializationMode
func (p_ PHASEEngine) SetOutputSpatializationMode(value PHASESpatializationMode) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setOutputSpatializationMode:"), value)
}

// The status of the engine’s audio playback.
//
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASEEngine/renderingState
func (p_ PHASEEngine) RenderingState() PHASERenderingState {
	rv := objc.Send[PHASERenderingState](p_.ID, objc.Sel("renderingState"))
	return rv
}

// The main object to which the app adds child objects.
//
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASEEngine/rootObject
func (p_ PHASEEngine) RootObject() PHASEObject {
	rv := objc.Send[PHASEObject](p_.ID, objc.Sel("rootObject"))
	return rv
}

// A collection of the sounds that play under various runtime circumstances.
//
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASEEngine/soundEvents
func (p_ PHASEEngine) SoundEvents() []PHASESoundEvent {
	rv := objc.Send[[]PHASESoundEvent](p_.ID, objc.Sel("soundEvents"))
	return rv
}

// A conversion factor from meters to your app’s preferred unit of measurement.
//
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASEEngine/unitsPerMeter
func (p_ PHASEEngine) UnitsPerMeter() float64 {
	rv := objc.Send[float64](p_.ID, objc.Sel("unitsPerMeter"))
	return rv
}


// SetUnitsPerMeter sets the value of the unitsPerMeter property.
// A conversion factor from meters to your app’s preferred unit of measurement.

//
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASEEngine/unitsPerMeter
func (p_ PHASEEngine) SetUnitsPerMeter(value float64) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setUnitsPerMeter:"), value)
}

// A conversion factor from seconds to your app’s preferred unit of time.
//
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASEEngine/unitsPerSecond
func (p_ PHASEEngine) UnitsPerSecond() float64 {
	rv := objc.Send[float64](p_.ID, objc.Sel("unitsPerSecond"))
	return rv
}


// SetUnitsPerSecond sets the value of the unitsPerSecond property.
// A conversion factor from seconds to your app’s preferred unit of time.

//
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASEEngine/unitsPerSecond
func (p_ PHASEEngine) SetUnitsPerSecond(value float64) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setUnitsPerSecond:"), value)
}


