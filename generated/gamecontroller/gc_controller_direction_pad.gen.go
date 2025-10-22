// Code generated from Apple documentation for GameController. DO NOT EDIT.

package gamecontroller

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [GCControllerDirectionPad] class.
var (
	GCControllerDirectionPadClass     _GCControllerDirectionPadClass
	GCControllerDirectionPadClassOnce sync.Once
)

func getGCControllerDirectionPadClass() _GCControllerDirectionPadClass {
	GCControllerDirectionPadClassOnce.Do(func() {
		GCControllerDirectionPadClass = _GCControllerDirectionPadClass{objc.GetClass("GCControllerDirectionPad")}
	})
	return GCControllerDirectionPadClass
}

type _GCControllerDirectionPadClass struct {
	class objc.Class
}

// An interface definition for the [GCControllerDirectionPad] class.
type IGCControllerDirectionPad interface {
	IGCControllerElement
	XAxis() GCControllerAxisInput
	YAxis() GCControllerAxisInput
	Down() GCControllerButtonInput
	SetDown(value IGCControllerButtonInput)
	Left() GCControllerButtonInput
	SetLeft(value IGCControllerButtonInput)
	Right() GCControllerButtonInput
	SetRight(value IGCControllerButtonInput)
	Up() GCControllerButtonInput
	SetUp(value IGCControllerButtonInput)
	ValueChangedHandler() unsafe.Pointer
	SetValueChangedHandler(value unsafe.Pointer)
}

// A control element associated with a directional pad or a thumbstick.
//
// You get the input values for this element from its subelements. You can use either the and properties to get coordinates, or the , , , and buttons that simulate directional pad buttons.
//
// [Full Topic]: https://developer.apple.com/documentation/GameController/GCControllerDirectionPad
type GCControllerDirectionPad struct {
	GCControllerElement
}

// GCControllerDirectionPadFrom constructs a [GCControllerDirectionPad] from an unsafe.Pointer.
//
// A control element associated with a directional pad or a thumbstick.
func GCControllerDirectionPadFrom(ptr unsafe.Pointer) GCControllerDirectionPad {
	return GCControllerDirectionPad{
		GCControllerElement: GCControllerElementFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (gc _GCControllerDirectionPadClass) Alloc() GCControllerDirectionPad {
	rv := objc.Send[GCControllerDirectionPad](objc.ID(gc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (gc _GCControllerDirectionPadClass) New() GCControllerDirectionPad {
	rv := objc.Send[GCControllerDirectionPad](objc.ID(gc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (g_ GCControllerDirectionPad) Init() GCControllerDirectionPad {
	rv := objc.Send[GCControllerDirectionPad](g_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (g_ GCControllerDirectionPad) Autorelease() GCControllerDirectionPad {
	rv := objc.Send[GCControllerDirectionPad](g_.ID, objc.Sel("autorelease"))
	return rv
}

// NewGCControllerDirectionPad creates a new GCControllerDirectionPad instance.
func NewGCControllerDirectionPad() GCControllerDirectionPad {
	return getGCControllerDirectionPadClass().New()
}


// The x-axis element of the directional pad.
//
// [Full Topic]: https://developer.apple.com/documentation/GameController/GCControllerDirectionPad/xAxis
func (g_ GCControllerDirectionPad) XAxis() GCControllerAxisInput {
	rv := objc.Send[GCControllerAxisInput](g_.ID, objc.Sel("xAxis"))
	return rv
}

// The y-axis element of the directional pad.
//
// [Full Topic]: https://developer.apple.com/documentation/GameController/GCControllerDirectionPad/yAxis
func (g_ GCControllerDirectionPad) YAxis() GCControllerAxisInput {
	rv := objc.Send[GCControllerAxisInput](g_.ID, objc.Sel("yAxis"))
	return rv
}

// The button element used for the negative y-axis direction.
//
// [Full Topic]: https://developer.apple.com/documentation/gamecontroller/gccontrollerdirectionpad/down
func (g_ GCControllerDirectionPad) Down() GCControllerButtonInput {
	rv := objc.Send[GCControllerButtonInput](g_.ID, objc.Sel("down"))
	return rv
}


// SetDown sets the value of the down property.
// The button element used for the negative y-axis direction.

//
// [Full Topic]: https://developer.apple.com/documentation/gamecontroller/gccontrollerdirectionpad/down
func (g_ GCControllerDirectionPad) SetDown(value IGCControllerButtonInput) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setDown:"), value)
}

// The button element that changes the negative x-axis.
//
// [Full Topic]: https://developer.apple.com/documentation/gamecontroller/gccontrollerdirectionpad/left
func (g_ GCControllerDirectionPad) Left() GCControllerButtonInput {
	rv := objc.Send[GCControllerButtonInput](g_.ID, objc.Sel("left"))
	return rv
}


// SetLeft sets the value of the left property.
// The button element that changes the negative x-axis.

//
// [Full Topic]: https://developer.apple.com/documentation/gamecontroller/gccontrollerdirectionpad/left
func (g_ GCControllerDirectionPad) SetLeft(value IGCControllerButtonInput) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setLeft:"), value)
}

// The button element that changes the positive x-axis.
//
// [Full Topic]: https://developer.apple.com/documentation/gamecontroller/gccontrollerdirectionpad/right
func (g_ GCControllerDirectionPad) Right() GCControllerButtonInput {
	rv := objc.Send[GCControllerButtonInput](g_.ID, objc.Sel("right"))
	return rv
}


// SetRight sets the value of the right property.
// The button element that changes the positive x-axis.

//
// [Full Topic]: https://developer.apple.com/documentation/gamecontroller/gccontrollerdirectionpad/right
func (g_ GCControllerDirectionPad) SetRight(value IGCControllerButtonInput) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setRight:"), value)
}

// The button element that changes the positive y-axis.
//
// [Full Topic]: https://developer.apple.com/documentation/gamecontroller/gccontrollerdirectionpad/up
func (g_ GCControllerDirectionPad) Up() GCControllerButtonInput {
	rv := objc.Send[GCControllerButtonInput](g_.ID, objc.Sel("up"))
	return rv
}


// SetUp sets the value of the up property.
// The button element that changes the positive y-axis.

//
// [Full Topic]: https://developer.apple.com/documentation/gamecontroller/gccontrollerdirectionpad/up
func (g_ GCControllerDirectionPad) SetUp(value IGCControllerButtonInput) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setUp:"), value)
}

// The block that the directional pad calls when the user changes its values.
//
// [Full Topic]: https://developer.apple.com/documentation/gamecontroller/gccontrollerdirectionpad/valuechangedhandler
func (g_ GCControllerDirectionPad) ValueChangedHandler() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](g_.ID, objc.Sel("valueChangedHandler"))
	return rv
}


// SetValueChangedHandler sets the value of the valueChangedHandler property.
// The block that the directional pad calls when the user changes its values.

//
// [Full Topic]: https://developer.apple.com/documentation/gamecontroller/gccontrollerdirectionpad/valuechangedhandler
func (g_ GCControllerDirectionPad) SetValueChangedHandler(value unsafe.Pointer) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setValueChangedHandler:"), value)
}



