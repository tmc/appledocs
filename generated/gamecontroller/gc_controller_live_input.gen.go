// Code generated from Apple documentation for GameController. DO NOT EDIT.

package gamecontroller

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class GCControllerLiveInput */


/* debug [class_header]: Header for GCControllerLiveInput */
// The class instance for the [GCControllerLiveInput] class.
var (
	GCControllerLiveInputClass     _GCControllerLiveInputClass
	GCControllerLiveInputClassOnce sync.Once
)

func getGCControllerLiveInputClass() _GCControllerLiveInputClass {
	GCControllerLiveInputClassOnce.Do(func() {
		GCControllerLiveInputClass = _GCControllerLiveInputClass{objc.GetClass("GCControllerLiveInput")}
	})
	return GCControllerLiveInputClass
}

type _GCControllerLiveInputClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for GCControllerLiveInput */
// An interface definition for the [GCControllerLiveInput] class.
type IGCControllerLiveInput interface {
	IGCControllerInputState
	
/* debug [class_interface_properties]: Properties for GCControllerLiveInput */
	// properties:
	UnmappedInput() IGCControllerLiveInput
	Input() IGCControllerLiveInput
	SetInput(value IGCControllerLiveInput)
	Unmapped() IGCControllerLiveInput
	SetUnmapped(value IGCControllerLiveInput)
	InputStateQueueDepth() int
	SetInputStateQueueDepth(value int)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for GCControllerLiveInput */
	// methods:
	Capture() IGCControllerInputState
	NextInputState() unsafe.Pointer
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for GCControllerLiveInput */
// Alloc allocates a new instance without initialization.
func (gc _GCControllerLiveInputClass) Alloc() GCControllerLiveInput {
	rv := objc.Send[GCControllerLiveInput](objc.ID(gc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (gc _GCControllerLiveInputClass) New() GCControllerLiveInput {
	rv := objc.Send[GCControllerLiveInput](objc.ID(gc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (g_ GCControllerLiveInput) Init() GCControllerLiveInput {
	rv := objc.Send[GCControllerLiveInput](g_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (g_ GCControllerLiveInput) Autorelease() GCControllerLiveInput {
	rv := objc.Send[GCControllerLiveInput](g_.ID, objc.Sel("autorelease"))
	return rv
}

// NewGCControllerLiveInput creates a new GCControllerLiveInput instance.
func NewGCControllerLiveInput() GCControllerLiveInput {
	return getGCControllerLiveInputClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for GCControllerLiveInput */
// The input profile for a controller.
//
// Instances of represent the current input state of a controller. You can save snapshots of the input state and receive callbacks when the input state changes. You can also get the elements of the controller and their current input values from instances. Use the  method to save a copy of the current input state. If you want Game Controller to buffer snapshots of the input states for you, use the    property to set the buffer’s queue depth to a value other than . Then use the method to get the snapshots when you’re ready to process input.


// The input profile for a controller.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameController/GCControllerLiveInput
type GCControllerLiveInput struct {
	GCControllerInputState
}

// GCControllerLiveInputFrom constructs a [GCControllerLiveInput] from an unsafe.Pointer.
//
// The input profile for a controller.
func GCControllerLiveInputFrom(ptr unsafe.Pointer) GCControllerLiveInput {
	return GCControllerLiveInput{
		GCControllerInputState: GCControllerInputStateFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for GCControllerLiveInput *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for GCControllerLiveInput */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for GCControllerLiveInput */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for GCControllerLiveInput */

// Returns a snapshot of the physical device inputs.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameController/GCControllerLiveInput/capture()
func (g_ GCControllerLiveInput) Capture() IGCControllerInputState {
	rv := objc.Send[GCControllerInputState](g_.ID, objc.Sel("capture"))
	return rv
}/* debug [instance_methods/method]: Capture */


// Returns the next device input state from the queue.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameController/GCControllerLiveInput/nextInputState()
func (g_ GCControllerLiveInput) NextInputState() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](g_.ID, objc.Sel("nextInputState"))
	return rv
}/* debug [instance_methods/method]: NextInputState */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for GCControllerLiveInput */

// The live input of a controller without any system-level remapping of the controls.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameController/GCControllerLiveInput/unmapped
func (g_ GCControllerLiveInput) UnmappedInput() IGCControllerLiveInput {
	rv := objc.Send[GCControllerLiveInput](g_.ID, objc.Sel("unmappedInput"))
	return rv
}/* debug [instance_properties/getter]: unmappedInput */


// The input profile for the controller.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamecontroller/gccontroller/input
func (g_ GCControllerLiveInput) Input() IGCControllerLiveInput {
	rv := objc.Send[GCControllerLiveInput](g_.ID, objc.Sel("input"))
	return rv
}/* debug [instance_properties/getter]: input */


// The input profile for the controller.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamecontroller/gccontroller/input
func (g_ GCControllerLiveInput) SetInput(value IGCControllerLiveInput) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setInput:"), value)
}/* debug [instance_properties/setter]: input */


// The live input of a controller without any system-level remapping of the controls.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamecontroller/gccontrollerliveinput/unmapped
func (g_ GCControllerLiveInput) Unmapped() IGCControllerLiveInput {
	rv := objc.Send[GCControllerLiveInput](g_.ID, objc.Sel("unmapped"))
	return rv
}/* debug [instance_properties/getter]: unmapped */


// The live input of a controller without any system-level remapping of the controls.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamecontroller/gccontrollerliveinput/unmapped
func (g_ GCControllerLiveInput) SetUnmapped(value IGCControllerLiveInput) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setUnmapped:"), value)
}/* debug [instance_properties/setter]: unmapped */


// The maximum number of input values that the queue stores.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamecontroller/gcdevicephysicalinput/inputstatequeuedepth
func (g_ GCControllerLiveInput) InputStateQueueDepth() int {
	rv := objc.Send[int](g_.ID, objc.Sel("inputStateQueueDepth"))
	return rv
}/* debug [instance_properties/getter]: inputStateQueueDepth */


// The maximum number of input values that the queue stores.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamecontroller/gcdevicephysicalinput/inputstatequeuedepth
func (g_ GCControllerLiveInput) SetInputStateQueueDepth(value int) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setInputStateQueueDepth:"), value)
}/* debug [instance_properties/setter]: inputStateQueueDepth */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class GCControllerLiveInput */



