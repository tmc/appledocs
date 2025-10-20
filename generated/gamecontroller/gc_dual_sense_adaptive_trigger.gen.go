// Code generated from Apple documentation for GameController. DO NOT EDIT.

package gamecontroller

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [GCDualSenseAdaptiveTrigger] class.
var (
	GCDualSenseAdaptiveTriggerClass     _GCDualSenseAdaptiveTriggerClass
	GCDualSenseAdaptiveTriggerClassOnce sync.Once
)

func getGCDualSenseAdaptiveTriggerClass() _GCDualSenseAdaptiveTriggerClass {
	GCDualSenseAdaptiveTriggerClassOnce.Do(func() {
		GCDualSenseAdaptiveTriggerClass = _GCDualSenseAdaptiveTriggerClass{objc.GetClass("GCDualSenseAdaptiveTrigger")}
	})
	return GCDualSenseAdaptiveTriggerClass
}

type _GCDualSenseAdaptiveTriggerClass struct {
	class objc.Class
}

// An interface definition for the [GCDualSenseAdaptiveTrigger] class.
type IGCDualSenseAdaptiveTrigger interface {
	IGCControllerButtonInput
	SetModeVibrationWithAmplitudesFrequency(positionalAmplitudes unsafe.Pointer, frequency unsafe.Pointer)
}

// A class that encapsulates the features of a DualSense adaptive trigger.
//
// A object allows you to specify a dynamic resistance force that the DualSense controller applies when the user pulls the trigger. For example, set the resistance to give the user the feeling of pulling back on a bow string, firing a weapon, or pulling a lever.
//
// [Full Topic]: https://developer.apple.com/documentation/GameController/GCDualSenseAdaptiveTrigger
type GCDualSenseAdaptiveTrigger struct {
	GCControllerButtonInput
}

// GCDualSenseAdaptiveTriggerFrom constructs a [GCDualSenseAdaptiveTrigger] from an unsafe.Pointer.
//
// A class that encapsulates the features of a DualSense adaptive trigger.
func GCDualSenseAdaptiveTriggerFrom(ptr unsafe.Pointer) GCDualSenseAdaptiveTrigger {
	return GCDualSenseAdaptiveTrigger{
		GCControllerButtonInput: GCControllerButtonInputFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (gc _GCDualSenseAdaptiveTriggerClass) Alloc() GCDualSenseAdaptiveTrigger {
	rv := objc.Send[GCDualSenseAdaptiveTrigger](objc.ID(gc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (gc _GCDualSenseAdaptiveTriggerClass) New() GCDualSenseAdaptiveTrigger {
	rv := objc.Send[GCDualSenseAdaptiveTrigger](objc.ID(gc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (g_ GCDualSenseAdaptiveTrigger) Init() GCDualSenseAdaptiveTrigger {
	rv := objc.Send[GCDualSenseAdaptiveTrigger](g_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (g_ GCDualSenseAdaptiveTrigger) Autorelease() GCDualSenseAdaptiveTrigger {
	rv := objc.Send[GCDualSenseAdaptiveTrigger](g_.ID, objc.Sel("autorelease"))
	return rv
}

// NewGCDualSenseAdaptiveTrigger creates a new GCDualSenseAdaptiveTrigger instance.
func NewGCDualSenseAdaptiveTrigger() GCDualSenseAdaptiveTrigger {
	return getGCDualSenseAdaptiveTriggerClass().New()
}


// Sets the mode to vibrate with the specified amplitudes for each possible trigger position.
//
// [Full Topic]: https://developer.apple.com/documentation/GameController/GCDualSenseAdaptiveTrigger/setModeVibration(amplitudes:frequency:)
func (g_ GCDualSenseAdaptiveTrigger) SetModeVibrationWithAmplitudesFrequency(positionalAmplitudes unsafe.Pointer, frequency unsafe.Pointer) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setModeVibrationWithAmplitudes:frequency:"), positionalAmplitudes, frequency)
}



