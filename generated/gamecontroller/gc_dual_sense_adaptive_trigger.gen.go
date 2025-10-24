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
	// properties:
	ArmPosition() float32
	SetArmPosition(value float32)
	Mode() unsafe.Pointer
	SetMode(value unsafe.Pointer)
	Status() unsafe.Pointer
	SetStatus(value unsafe.Pointer)
	// methods:
	SetModeSlopeFeedbackWithStartPositionEndPositionStartStrengthEndStrength(startPosition float32, endPosition float32, startStrength float32, endStrength float32)
	SetModeVibrationWithAmplitudesFrequency(positionalAmplitudes unsafe.Pointer, frequency float32)
}

// A class that encapsulates the features of a DualSense adaptive trigger.
//
// A object allows you to specify a dynamic resistance force that the DualSense controller applies when the user pulls the trigger. For example, set the resistance to give the user the feeling of pulling back on a bow string, firing a weapon, or pulling a lever.


// A class that encapsulates the features of a DualSense adaptive trigger.
//
// [Full Topic]
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



// Sets the mode to provide feedback when the user tilts the trigger between the start and the end positions.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameController/GCDualSenseAdaptiveTrigger/setModeSlopeFeedback(startPosition:endPosition:startStrength:endStrength:)
func (g_ GCDualSenseAdaptiveTrigger) SetModeSlopeFeedbackWithStartPositionEndPositionStartStrengthEndStrength(startPosition float32, endPosition float32, startStrength float32, endStrength float32) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setModeSlopeFeedbackWithStartPosition:endPosition:startStrength:endStrength:"), startPosition, endPosition, startStrength, endStrength)
}


// Sets the mode to vibrate with the specified amplitudes for each possible trigger position.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameController/GCDualSenseAdaptiveTrigger/setModeVibration(amplitudes:frequency:)
func (g_ GCDualSenseAdaptiveTrigger) SetModeVibrationWithAmplitudesFrequency(positionalAmplitudes unsafe.Pointer, frequency float32) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setModeVibrationWithAmplitudes:frequency:"), positionalAmplitudes, frequency)
}


// The position of the trigger’s arm.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamecontroller/gcdualsenseadaptivetrigger/armposition
func (g_ GCDualSenseAdaptiveTrigger) ArmPosition() float32 {
	rv := objc.Send[float32](g_.ID, objc.Sel("armPosition"))
	return rv
}


// The position of the trigger’s arm.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamecontroller/gcdualsenseadaptivetrigger/armposition
func (g_ GCDualSenseAdaptiveTrigger) SetArmPosition(value float32) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setArmPosition:"), value)
}


// The current configuration of the adaptive trigger.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamecontroller/gcdualsenseadaptivetrigger/mode-swift.property
func (g_ GCDualSenseAdaptiveTrigger) Mode() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](g_.ID, objc.Sel("mode"))
	return rv
}


// The current configuration of the adaptive trigger.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamecontroller/gcdualsenseadaptivetrigger/mode-swift.property
func (g_ GCDualSenseAdaptiveTrigger) SetMode(value unsafe.Pointer) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setMode:"), value)
}


// The current status of the adaptive trigger and whether it’s applying effects.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamecontroller/gcdualsenseadaptivetrigger/status-swift.property
func (g_ GCDualSenseAdaptiveTrigger) Status() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](g_.ID, objc.Sel("status"))
	return rv
}


// The current status of the adaptive trigger and whether it’s applying effects.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamecontroller/gcdualsenseadaptivetrigger/status-swift.property
func (g_ GCDualSenseAdaptiveTrigger) SetStatus(value unsafe.Pointer) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setStatus:"), value)
}



