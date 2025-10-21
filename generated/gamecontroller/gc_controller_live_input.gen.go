// Code generated from Apple documentation for GameController. DO NOT EDIT.

package gamecontroller

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

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

// An interface definition for the [GCControllerLiveInput] class.
type IGCControllerLiveInput interface {
	objectivec.IObject
	Capture() unsafe.Pointer
	NextInputState() unsafe.Pointer
}

// The input profile for a controller.
//
// Instances of represent the current input state of a controller. You can save snapshots of the input state and receive callbacks when the input state changes. You can also get the elements of the controller and their current input values from instances. Use the  method to save a copy of the current input state. If you want Game Controller to buffer snapshots of the input states for you, use the    property to set the buffer’s queue depth to a value other than . Then use the method to get the snapshots when you’re ready to process input.
//
// [Full Topic]: https://developer.apple.com/documentation/GameController/GCControllerLiveInput
type GCControllerLiveInput struct {
	objectivec.Object
}

// GCControllerLiveInputFrom constructs a [GCControllerLiveInput] from an unsafe.Pointer.
//
// The input profile for a controller.
func GCControllerLiveInputFrom(ptr unsafe.Pointer) GCControllerLiveInput {
	return GCControllerLiveInput{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (gc _GCControllerLiveInputClass) Alloc() GCControllerLiveInput {
	rv := objc.Send[GCControllerLiveInput](objc.ID(gc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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


// Returns a snapshot of the physical device inputs.
//
// [Full Topic]: https://developer.apple.com/documentation/GameController/GCControllerLiveInput/capture()
func (g_ GCControllerLiveInput) Capture() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](g_.ID, objc.Sel("capture"))
	return rv
}

// Returns the next device input state from the queue.
//
// [Full Topic]: https://developer.apple.com/documentation/GameController/GCControllerLiveInput/nextInputState()
func (g_ GCControllerLiveInput) NextInputState() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](g_.ID, objc.Sel("nextInputState"))
	return rv
}

// The input profile for the controller.
//
// [Full Topic]: https://developer.apple.com/documentation/gamecontroller/gccontroller/input
func (g_ GCControllerLiveInput) Input() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](g_.ID, objc.Sel("input"))
	return rv
}


// SetInput sets the value of the input property.
// The input profile for the controller.

//
// [Full Topic]: https://developer.apple.com/documentation/gamecontroller/gccontroller/input
func (g_ GCControllerLiveInput) SetInput(value unsafe.Pointer) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setInput:"), value)
}

// The live input of a controller without any system-level remapping of the controls.
//
// [Full Topic]: https://developer.apple.com/documentation/gamecontroller/gccontrollerliveinput/unmapped
func (g_ GCControllerLiveInput) Unmapped() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](g_.ID, objc.Sel("unmapped"))
	return rv
}


// SetUnmapped sets the value of the unmapped property.
// The live input of a controller without any system-level remapping of the controls.

//
// [Full Topic]: https://developer.apple.com/documentation/gamecontroller/gccontrollerliveinput/unmapped
func (g_ GCControllerLiveInput) SetUnmapped(value unsafe.Pointer) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setUnmapped:"), value)
}

// The maximum number of input values that the queue stores.
//
// [Full Topic]: https://developer.apple.com/documentation/gamecontroller/gcdevicephysicalinput/inputstatequeuedepth
func (g_ GCControllerLiveInput) InputStateQueueDepth() int {
	rv := objc.Send[int](g_.ID, objc.Sel("inputStateQueueDepth"))
	return rv
}


// SetInputStateQueueDepth sets the value of the inputStateQueueDepth property.
// The maximum number of input values that the queue stores.

//
// [Full Topic]: https://developer.apple.com/documentation/gamecontroller/gcdevicephysicalinput/inputstatequeuedepth
func (g_ GCControllerLiveInput) SetInputStateQueueDepth(value int) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setInputStateQueueDepth:"), value)
}

// The live input of a controller without any system-level remapping of the controls.
//
// [Full Topic]: https://developer.apple.com/documentation/GameController/GCControllerLiveInput/unmapped
func (g_ GCControllerLiveInput) UnmappedInput() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](g_.ID, objc.Sel("unmappedInput"))
	return rv
}



