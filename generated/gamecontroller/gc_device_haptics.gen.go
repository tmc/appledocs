// Code generated from Apple documentation for GameController. DO NOT EDIT.

package gamecontroller

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [GCDeviceHaptics] class.
var (
	GCDeviceHapticsClass     _GCDeviceHapticsClass
	GCDeviceHapticsClassOnce sync.Once
)

func getGCDeviceHapticsClass() _GCDeviceHapticsClass {
	GCDeviceHapticsClassOnce.Do(func() {
		GCDeviceHapticsClass = _GCDeviceHapticsClass{objc.GetClass("GCDeviceHaptics")}
	})
	return GCDeviceHapticsClass
}

type _GCDeviceHapticsClass struct {
	class objc.Class
}

// An interface definition for the [GCDeviceHaptics] class.
type IGCDeviceHaptics interface {
	objectivec.IObject
	CreateEngineWithLocality(locality IGCHapticsLocality) unsafe.Pointer
}

// The locations of haptic actuators on a game controller.
//
// Use this class to create a haptic engine with a specified locality. Any patterns you send to that engine play on the specified actuators.
//
// [Full Topic]: https://developer.apple.com/documentation/GameController/GCDeviceHaptics
type GCDeviceHaptics struct {
	objectivec.Object
}

// GCDeviceHapticsFrom constructs a [GCDeviceHaptics] from an unsafe.Pointer.
//
// The locations of haptic actuators on a game controller.
func GCDeviceHapticsFrom(ptr unsafe.Pointer) GCDeviceHaptics {
	return GCDeviceHaptics{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (gc _GCDeviceHapticsClass) Alloc() GCDeviceHaptics {
	rv := objc.Send[GCDeviceHaptics](objc.ID(gc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (gc _GCDeviceHapticsClass) New() GCDeviceHaptics {
	rv := objc.Send[GCDeviceHaptics](objc.ID(gc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (g_ GCDeviceHaptics) Init() GCDeviceHaptics {
	rv := objc.Send[GCDeviceHaptics](g_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (g_ GCDeviceHaptics) Autorelease() GCDeviceHaptics {
	rv := objc.Send[GCDeviceHaptics](g_.ID, objc.Sel("autorelease"))
	return rv
}

// NewGCDeviceHaptics creates a new GCDeviceHaptics instance.
func NewGCDeviceHaptics() GCDeviceHaptics {
	return getGCDeviceHapticsClass().New()
}


// Creates a haptics engine with the specified locality.
//
// [Full Topic]: https://developer.apple.com/documentation/GameController/GCDeviceHaptics/createEngine(withLocality:)
func (g_ GCDeviceHaptics) CreateEngineWithLocality(locality IGCHapticsLocality) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](g_.ID, objc.Sel("createEngineWithLocality:"), locality)
	return rv
}

// The locations of haptic actuators on the device.
//
// [Full Topic]: https://developer.apple.com/documentation/GameController/GCDeviceHaptics/supportedLocalities
func (g_ GCDeviceHaptics) SupportedLocalities() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](g_.ID, objc.Sel("supportedLocalities"))
	return rv
}

// A Boolean value that indicates whether the device supports haptic event playback.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreHaptics/CHHapticDeviceCapability/supportsHaptics
func (g_ GCDeviceHaptics) SupportsHaptics() bool {
	rv := objc.Send[bool](g_.ID, objc.Sel("supportsHaptics"))
	return rv
}


// SetSupportsHaptics sets the value of the supportsHaptics property.
// A Boolean value that indicates whether the device supports haptic event playback.

//
// [Full Topic]: https://developer.apple.com/documentation/CoreHaptics/CHHapticDeviceCapability/supportsHaptics
func (g_ GCDeviceHaptics) SetSupportsHaptics(value bool) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setSupportsHaptics:"), value)
}

// An infinite duration for a haptics event.
//
// [Full Topic]: https://developer.apple.com/documentation/gamecontroller/gchapticdurationinfinite
func (g_ GCDeviceHaptics) GCHapticDurationInfinite() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](g_.ID, objc.Sel("GCHapticDurationInfinite"))
	return rv
}



