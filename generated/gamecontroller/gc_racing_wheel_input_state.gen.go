// Code generated from Apple documentation for GameController. DO NOT EDIT.

package gamecontroller

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class GCRacingWheelInputState */


/* debug [class_header]: Header for GCRacingWheelInputState */
// The class instance for the [GCRacingWheelInputState] class.
var (
	GCRacingWheelInputStateClass     _GCRacingWheelInputStateClass
	GCRacingWheelInputStateClassOnce sync.Once
)

func getGCRacingWheelInputStateClass() _GCRacingWheelInputStateClass {
	GCRacingWheelInputStateClassOnce.Do(func() {
		GCRacingWheelInputStateClass = _GCRacingWheelInputStateClass{objc.GetClass("GCRacingWheelInputState")}
	})
	return GCRacingWheelInputStateClass
}

type _GCRacingWheelInputStateClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for GCRacingWheelInputState */
// An interface definition for the [GCRacingWheelInputState] class.
type IGCRacingWheelInputState interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for GCRacingWheelInputState */
	// properties:
	AcceleratorPedal() unsafe.Pointer
	BrakePedal() unsafe.Pointer
	ClutchPedal() unsafe.Pointer
	Shifter() IGCGearShifterElement
	Wheel() IGCSteeringWheelElement
	GCInputPedalAccelerator() objc.IObject /* cross-framework: NSString */
	SetGCInputPedalAccelerator(value objc.IObject /* cross-framework: NSString */)
	GCInputPedalBrake() objc.IObject /* cross-framework: NSString */
	SetGCInputPedalBrake(value objc.IObject /* cross-framework: NSString */)
	GCInputPedalClutch() objc.IObject /* cross-framework: NSString */
	SetGCInputPedalClutch(value objc.IObject /* cross-framework: NSString */)
	GCInputShifter() objc.IObject /* cross-framework: NSString */
	SetGCInputShifter(value objc.IObject /* cross-framework: NSString */)
	GCInputSteeringWheel() objc.IObject /* cross-framework: NSString */
	SetGCInputSteeringWheel(value objc.IObject /* cross-framework: NSString */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for GCRacingWheelInputState */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for GCRacingWheelInputState */
// Alloc allocates a new instance without initialization.
func (gc _GCRacingWheelInputStateClass) Alloc() GCRacingWheelInputState {
	rv := objc.Send[GCRacingWheelInputState](objc.ID(gc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (gc _GCRacingWheelInputStateClass) New() GCRacingWheelInputState {
	rv := objc.Send[GCRacingWheelInputState](objc.ID(gc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (g_ GCRacingWheelInputState) Init() GCRacingWheelInputState {
	rv := objc.Send[GCRacingWheelInputState](g_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (g_ GCRacingWheelInputState) Autorelease() GCRacingWheelInputState {
	rv := objc.Send[GCRacingWheelInputState](g_.ID, objc.Sel("autorelease"))
	return rv
}

// NewGCRacingWheelInputState creates a new GCRacingWheelInputState instance.
func NewGCRacingWheelInputState() GCRacingWheelInputState {
	return getGCRacingWheelInputStateClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for GCRacingWheelInputState */
// The input for the wheel of a racing wheel controller.


// The input for the wheel of a racing wheel controller.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameController/GCRacingWheelInputState
type GCRacingWheelInputState struct {
	objectivec.Object
}

// GCRacingWheelInputStateFrom constructs a [GCRacingWheelInputState] from an unsafe.Pointer.
//
// The input for the wheel of a racing wheel controller.
func GCRacingWheelInputStateFrom(ptr unsafe.Pointer) GCRacingWheelInputState {
	return GCRacingWheelInputState{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for GCRacingWheelInputState *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for GCRacingWheelInputState */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for GCRacingWheelInputState */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for GCRacingWheelInputState */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for GCRacingWheelInputState */

// The controller’s accelerator pedal element.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameController/GCRacingWheelInputState/acceleratorPedal
func (g_ GCRacingWheelInputState) AcceleratorPedal() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](g_.ID, objc.Sel("acceleratorPedal"))
	return rv
}/* debug [instance_properties/getter]: acceleratorPedal */


// The controller’s brake pedal element.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameController/GCRacingWheelInputState/brakePedal
func (g_ GCRacingWheelInputState) BrakePedal() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](g_.ID, objc.Sel("brakePedal"))
	return rv
}/* debug [instance_properties/getter]: brakePedal */


// The controller’s clutch element.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameController/GCRacingWheelInputState/clutchPedal
func (g_ GCRacingWheelInputState) ClutchPedal() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](g_.ID, objc.Sel("clutchPedal"))
	return rv
}/* debug [instance_properties/getter]: clutchPedal */


// The controller’s gear shift element.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameController/GCRacingWheelInputState/shifter
func (g_ GCRacingWheelInputState) Shifter() IGCGearShifterElement {
	rv := objc.Send[GCGearShifterElement](g_.ID, objc.Sel("shifter"))
	return rv
}/* debug [instance_properties/getter]: shifter */


// The controller’s wheel element.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameController/GCRacingWheelInputState/wheel
func (g_ GCRacingWheelInputState) Wheel() IGCSteeringWheelElement {
	rv := objc.Send[GCSteeringWheelElement](g_.ID, objc.Sel("wheel"))
	return rv
}/* debug [instance_properties/getter]: wheel */


// The name of the accelerator element.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamecontroller/gcinputpedalaccelerator-6kg6u
func (g_ GCRacingWheelInputState) GCInputPedalAccelerator() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](g_.ID, objc.Sel("GCInputPedalAccelerator"))
	return rv
}/* debug [instance_properties/getter]: GCInputPedalAccelerator */


// The name of the accelerator element.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamecontroller/gcinputpedalaccelerator-6kg6u
func (g_ GCRacingWheelInputState) SetGCInputPedalAccelerator(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setGCInputPedalAccelerator:"), value)
}/* debug [instance_properties/setter]: GCInputPedalAccelerator */


// The name of the brake element.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamecontroller/gcinputpedalbrake-6wpdc
func (g_ GCRacingWheelInputState) GCInputPedalBrake() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](g_.ID, objc.Sel("GCInputPedalBrake"))
	return rv
}/* debug [instance_properties/getter]: GCInputPedalBrake */


// The name of the brake element.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamecontroller/gcinputpedalbrake-6wpdc
func (g_ GCRacingWheelInputState) SetGCInputPedalBrake(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setGCInputPedalBrake:"), value)
}/* debug [instance_properties/setter]: GCInputPedalBrake */


// The name of the clutch element.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamecontroller/gcinputpedalclutch-82gwe
func (g_ GCRacingWheelInputState) GCInputPedalClutch() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](g_.ID, objc.Sel("GCInputPedalClutch"))
	return rv
}/* debug [instance_properties/getter]: GCInputPedalClutch */


// The name of the clutch element.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamecontroller/gcinputpedalclutch-82gwe
func (g_ GCRacingWheelInputState) SetGCInputPedalClutch(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setGCInputPedalClutch:"), value)
}/* debug [instance_properties/setter]: GCInputPedalClutch */


// The name of the shifter element.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamecontroller/gcinputshifter-6miga
func (g_ GCRacingWheelInputState) GCInputShifter() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](g_.ID, objc.Sel("GCInputShifter"))
	return rv
}/* debug [instance_properties/getter]: GCInputShifter */


// The name of the shifter element.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamecontroller/gcinputshifter-6miga
func (g_ GCRacingWheelInputState) SetGCInputShifter(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setGCInputShifter:"), value)
}/* debug [instance_properties/setter]: GCInputShifter */


// The name of the steering wheel element.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamecontroller/gcinputsteeringwheel-26283
func (g_ GCRacingWheelInputState) GCInputSteeringWheel() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](g_.ID, objc.Sel("GCInputSteeringWheel"))
	return rv
}/* debug [instance_properties/getter]: GCInputSteeringWheel */


// The name of the steering wheel element.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamecontroller/gcinputsteeringwheel-26283
func (g_ GCRacingWheelInputState) SetGCInputSteeringWheel(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setGCInputSteeringWheel:"), value)
}/* debug [instance_properties/setter]: GCInputSteeringWheel */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class GCRacingWheelInputState */



