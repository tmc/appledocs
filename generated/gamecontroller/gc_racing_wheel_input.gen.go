// Code generated from Apple documentation for GameController. DO NOT EDIT.

package gamecontroller

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class GCRacingWheelInput */


/* debug [class_header]: Header for GCRacingWheelInput */
// The class instance for the [GCRacingWheelInput] class.
var (
	GCRacingWheelInputClass     _GCRacingWheelInputClass
	GCRacingWheelInputClassOnce sync.Once
)

func getGCRacingWheelInputClass() _GCRacingWheelInputClass {
	GCRacingWheelInputClassOnce.Do(func() {
		GCRacingWheelInputClass = _GCRacingWheelInputClass{objc.GetClass("GCRacingWheelInput")}
	})
	return GCRacingWheelInputClass
}

type _GCRacingWheelInputClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for GCRacingWheelInput */
// An interface definition for the [GCRacingWheelInput] class.
type IGCRacingWheelInput interface {
	IGCRacingWheelInputState
	
/* debug [class_interface_properties]: Properties for GCRacingWheelInput */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for GCRacingWheelInput */
	// methods:
	Capture() IGCRacingWheelInputState
	NextInputState() unsafe.Pointer
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for GCRacingWheelInput */
// Alloc allocates a new instance without initialization.
func (gc _GCRacingWheelInputClass) Alloc() GCRacingWheelInput {
	rv := objc.Send[GCRacingWheelInput](objc.ID(gc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (gc _GCRacingWheelInputClass) New() GCRacingWheelInput {
	rv := objc.Send[GCRacingWheelInput](objc.ID(gc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (g_ GCRacingWheelInput) Init() GCRacingWheelInput {
	rv := objc.Send[GCRacingWheelInput](g_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (g_ GCRacingWheelInput) Autorelease() GCRacingWheelInput {
	rv := objc.Send[GCRacingWheelInput](g_.ID, objc.Sel("autorelease"))
	return rv
}

// NewGCRacingWheelInput creates a new GCRacingWheelInput instance.
func NewGCRacingWheelInput() GCRacingWheelInput {
	return getGCRacingWheelInputClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for GCRacingWheelInput */
// A controller profile that supports a racing wheel.


// A controller profile that supports a racing wheel.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameController/GCRacingWheelInput
type GCRacingWheelInput struct {
	GCRacingWheelInputState
}

// GCRacingWheelInputFrom constructs a [GCRacingWheelInput] from an unsafe.Pointer.
//
// A controller profile that supports a racing wheel.
func GCRacingWheelInputFrom(ptr unsafe.Pointer) GCRacingWheelInput {
	return GCRacingWheelInput{
		GCRacingWheelInputState: GCRacingWheelInputStateFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for GCRacingWheelInput *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for GCRacingWheelInput */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for GCRacingWheelInput */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for GCRacingWheelInput */

// Returns a snapshot of the racing wheel inputs.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameController/GCRacingWheelInput/capture()
func (g_ GCRacingWheelInput) Capture() IGCRacingWheelInputState {
	rv := objc.Send[GCRacingWheelInputState](g_.ID, objc.Sel("capture"))
	return rv
}/* debug [instance_methods/method]: Capture */


// Returns the next input state of the racing wheel from the queue.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameController/GCRacingWheelInput/nextInputState()
func (g_ GCRacingWheelInput) NextInputState() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](g_.ID, objc.Sel("nextInputState"))
	return rv
}/* debug [instance_methods/method]: NextInputState */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for GCRacingWheelInput */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class GCRacingWheelInput */



