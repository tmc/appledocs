// Code generated from Apple documentation for GameController. DO NOT EDIT.

package gamecontroller

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class GCDualSenseAdaptiveTrigger */


/* debug [class_header]: Header for GCDualSenseAdaptiveTrigger */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for GCDualSenseAdaptiveTrigger */
// An interface definition for the [GCDualSenseAdaptiveTrigger] class.
type IGCDualSenseAdaptiveTrigger interface {
	IGCControllerButtonInput
	
/* debug [class_interface_properties]: Properties for GCDualSenseAdaptiveTrigger */
	// properties:
	ArmPosition() float32
	Mode() GCDualSenseAdaptiveTriggerMode
	Status() GCDualSenseAdaptiveTriggerStatus
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for GCDualSenseAdaptiveTrigger */
	// methods:
	SetModeFeedbackWithResistiveStrengths(positionalResistiveStrengths unsafe.Pointer)
	SetModeFeedbackWithStartPositionResistiveStrength(startPosition float32, resistiveStrength float32)
	SetModeOff()
	SetModeSlopeFeedbackWithStartPositionEndPositionStartStrengthEndStrength(startPosition float32, endPosition float32, startStrength float32, endStrength float32)
	SetModeVibrationWithAmplitudesFrequency(positionalAmplitudes unsafe.Pointer, frequency float32)
	SetModeVibrationWithStartPositionAmplitudeFrequency(startPosition float32, amplitude float32, frequency float32)
	SetModeWeaponWithStartPositionEndPositionResistiveStrength(startPosition float32, endPosition float32, resistiveStrength float32)
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for GCDualSenseAdaptiveTrigger */
// Alloc allocates a new instance without initialization.
func (gc _GCDualSenseAdaptiveTriggerClass) Alloc() GCDualSenseAdaptiveTrigger {
	rv := objc.Send[GCDualSenseAdaptiveTrigger](objc.ID(gc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for GCDualSenseAdaptiveTrigger */
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for GCDualSenseAdaptiveTrigger *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for GCDualSenseAdaptiveTrigger */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for GCDualSenseAdaptiveTrigger */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for GCDualSenseAdaptiveTrigger */

// Sets the mode to provide feedback with the specified strengths for each possible trigger position.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameController/GCDualSenseAdaptiveTrigger/setModeFeedback(resistiveStrengths:)
func (g_ GCDualSenseAdaptiveTrigger) SetModeFeedbackWithResistiveStrengths(positionalResistiveStrengths unsafe.Pointer) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setModeFeedbackWithResistiveStrengths:"), positionalResistiveStrengths)
}/* debug [instance_methods/method]: SetModeFeedbackWithResistiveStrengths */


// Sets the mode to provide feedback when the user depresses the trigger at the start position or at a greater value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameController/GCDualSenseAdaptiveTrigger/setModeFeedbackWithStartPosition(_:resistiveStrength:)
func (g_ GCDualSenseAdaptiveTrigger) SetModeFeedbackWithStartPositionResistiveStrength(startPosition float32, resistiveStrength float32) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setModeFeedbackWithStartPosition:resistiveStrength:"), startPosition, resistiveStrength)
}/* debug [instance_methods/method]: SetModeFeedbackWithStartPositionResistiveStrength */


// Sets the mode to off and stops any trigger effect.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameController/GCDualSenseAdaptiveTrigger/setModeOff()
func (g_ GCDualSenseAdaptiveTrigger) SetModeOff() {
	objc.Send[objc.ID](g_.ID, objc.Sel("setModeOff"))
}/* debug [instance_methods/method]: SetModeOff */


// Sets the mode to provide feedback when the user tilts the trigger between the start and the end positions.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameController/GCDualSenseAdaptiveTrigger/setModeSlopeFeedback(startPosition:endPosition:startStrength:endStrength:)
func (g_ GCDualSenseAdaptiveTrigger) SetModeSlopeFeedbackWithStartPositionEndPositionStartStrengthEndStrength(startPosition float32, endPosition float32, startStrength float32, endStrength float32) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setModeSlopeFeedbackWithStartPosition:endPosition:startStrength:endStrength:"), startPosition, endPosition, startStrength, endStrength)
}/* debug [instance_methods/method]: SetModeSlopeFeedbackWithStartPositionEndPositionStartStrengthEndStrength */


// Sets the mode to vibrate with the specified amplitudes for each possible trigger position.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameController/GCDualSenseAdaptiveTrigger/setModeVibration(amplitudes:frequency:)
func (g_ GCDualSenseAdaptiveTrigger) SetModeVibrationWithAmplitudesFrequency(positionalAmplitudes unsafe.Pointer, frequency float32) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setModeVibrationWithAmplitudes:frequency:"), positionalAmplitudes, frequency)
}/* debug [instance_methods/method]: SetModeVibrationWithAmplitudesFrequency */


// Sets the mode to vibrate when the user depresses the trigger at the start position or at a greater value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameController/GCDualSenseAdaptiveTrigger/setModeVibrationWithStartPosition(_:amplitude:frequency:)
func (g_ GCDualSenseAdaptiveTrigger) SetModeVibrationWithStartPositionAmplitudeFrequency(startPosition float32, amplitude float32, frequency float32) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setModeVibrationWithStartPosition:amplitude:frequency:"), startPosition, amplitude, frequency)
}/* debug [instance_methods/method]: SetModeVibrationWithStartPositionAmplitudeFrequency */


// Sets the mode to provide feedback when the user depresses the trigger between the start and the end positions.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameController/GCDualSenseAdaptiveTrigger/setModeWeaponWithStartPosition(_:endPosition:resistiveStrength:)
func (g_ GCDualSenseAdaptiveTrigger) SetModeWeaponWithStartPositionEndPositionResistiveStrength(startPosition float32, endPosition float32, resistiveStrength float32) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setModeWeaponWithStartPosition:endPosition:resistiveStrength:"), startPosition, endPosition, resistiveStrength)
}/* debug [instance_methods/method]: SetModeWeaponWithStartPositionEndPositionResistiveStrength */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for GCDualSenseAdaptiveTrigger */

// The position of the trigger’s arm.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameController/GCDualSenseAdaptiveTrigger/armPosition
func (g_ GCDualSenseAdaptiveTrigger) ArmPosition() float32 {
	rv := objc.Send[float32](g_.ID, objc.Sel("armPosition"))
	return rv
}/* debug [instance_properties/getter]: armPosition */


// The current configuration of the adaptive trigger.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameController/GCDualSenseAdaptiveTrigger/mode-swift.property
func (g_ GCDualSenseAdaptiveTrigger) Mode() GCDualSenseAdaptiveTriggerMode {
	rv := objc.Send[GCDualSenseAdaptiveTriggerMode](g_.ID, objc.Sel("mode"))
	return rv
}/* debug [instance_properties/getter]: mode */


// The current status of the adaptive trigger and whether it’s applying effects.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameController/GCDualSenseAdaptiveTrigger/status-swift.property
func (g_ GCDualSenseAdaptiveTrigger) Status() GCDualSenseAdaptiveTriggerStatus {
	rv := objc.Send[GCDualSenseAdaptiveTriggerStatus](g_.ID, objc.Sel("status"))
	return rv
}/* debug [instance_properties/getter]: status */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class GCDualSenseAdaptiveTrigger */



