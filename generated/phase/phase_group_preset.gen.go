// Code generated from Apple documentation for PHASE. DO NOT EDIT.

package phase

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [PHASEGroupPreset] class.
var (
	PHASEGroupPresetClass     _PHASEGroupPresetClass
	PHASEGroupPresetClassOnce sync.Once
)

func getPHASEGroupPresetClass() _PHASEGroupPresetClass {
	PHASEGroupPresetClassOnce.Do(func() {
		PHASEGroupPresetClass = _PHASEGroupPresetClass{objc.GetClass("PHASEGroupPreset")}
	})
	return PHASEGroupPresetClass
}

type _PHASEGroupPresetClass struct {
	class objc.Class
}

// An interface definition for the [PHASEGroupPreset] class.
type IPHASEGroupPreset interface {
	objectivec.IObject
	// properties:
	Settings() foundation.IDictionary
	TimeToReset() float64
	TimeToTarget() float64
	// methods:
	Activate()
	ActivateWithTimeToTargetOverride(timeToTargetOverride float64)
	Deactivate()
	DeactivateWithTimeToResetOverride(timeToResetOverride float64)
}

// A collection of settings for groups.
//
// Group presets pair groups with audio settings that your app can apply to specific sounds at a particular time in your app’s life cycle. This class enables many predefined group settings to take effect all at once.


// A collection of settings for groups.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASEGroupPreset
type PHASEGroupPreset struct {
	objectivec.Object
}

// PHASEGroupPresetFrom constructs a [PHASEGroupPreset] from an unsafe.Pointer.
//
// A collection of settings for groups.
func PHASEGroupPresetFrom(ptr unsafe.Pointer) PHASEGroupPreset {
	return PHASEGroupPreset{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (pc _PHASEGroupPresetClass) Alloc() PHASEGroupPreset {
	rv := objc.Send[PHASEGroupPreset](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (pc _PHASEGroupPresetClass) New() PHASEGroupPreset {
	rv := objc.Send[PHASEGroupPreset](objc.ID(pc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (p_ PHASEGroupPreset) Init() PHASEGroupPreset {
	rv := objc.Send[PHASEGroupPreset](p_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p_ PHASEGroupPreset) Autorelease() PHASEGroupPreset {
	rv := objc.Send[PHASEGroupPreset](p_.ID, objc.Sel("autorelease"))
	return rv
}

// NewPHASEGroupPreset creates a new PHASEGroupPreset instance.
func NewPHASEGroupPreset() PHASEGroupPreset {
	return getPHASEGroupPresetClass().New()
}



// Creates a group preset with the designated engine, settings, and fade parameters.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASEGroupPreset/init(engine:settings:timeToTarget:timeToReset:)
func NewPHASEGroupPresetWithEngineSettingsTimeToTargetTimeToReset(engine IPHASEEngine, settings foundation.IDictionary, timeToTarget float64, timeToReset float64) PHASEGroupPreset {
	instance := getPHASEGroupPresetClass().Alloc()
	rv := objc.Send[PHASEGroupPreset](instance.ID, objc.Sel("initWithEngine:settings:timeToTarget:timeToReset:"), engine, settings, timeToTarget, timeToReset)
	rv.Autorelease()
	return rv
}



// Applies settings to the designated groups.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASEGroupPreset/activate()
func (p_ PHASEGroupPreset) Activate() {
	objc.Send[objc.ID](p_.ID, objc.Sel("activate"))
}


// Applies settings with an overriden fade duration.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASEGroupPreset/activate(timeToTargetOverride:)
func (p_ PHASEGroupPreset) ActivateWithTimeToTargetOverride(timeToTargetOverride float64) {
	objc.Send[objc.ID](p_.ID, objc.Sel("activateWithTimeToTargetOverride:"), timeToTargetOverride)
}


// Reverts settings for the preset’s groups.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASEGroupPreset/deactivate()
func (p_ PHASEGroupPreset) Deactivate() {
	objc.Send[objc.ID](p_.ID, objc.Sel("deactivate"))
}


// Reverts settings for the preset’s groups using a timed adjustment.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASEGroupPreset/deactivate(timeToResetOverride:)
func (p_ PHASEGroupPreset) DeactivateWithTimeToResetOverride(timeToResetOverride float64) {
	objc.Send[objc.ID](p_.ID, objc.Sel("deactivateWithTimeToResetOverride:"), timeToResetOverride)
}


// A dictionary with preset setting values and group objects as keys.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASEGroupPreset/settings
func (p_ PHASEGroupPreset) Settings() foundation.IDictionary {
	rv := objc.Send[foundation.IDictionary](p_.ID, objc.Sel("settings"))
	return rv
}


// A duration in which the framework restores the group’s original state.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASEGroupPreset/timeToReset
func (p_ PHASEGroupPreset) TimeToReset() float64 {
	rv := objc.Send[float64](p_.ID, objc.Sel("timeToReset"))
	return rv
}


// A duration in which the engine fades the settings from their original value to their new value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASEGroupPreset/timeToTarget
func (p_ PHASEGroupPreset) TimeToTarget() float64 {
	rv := objc.Send[float64](p_.ID, objc.Sel("timeToTarget"))
	return rv
}


