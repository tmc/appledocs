// Code generated from Apple documentation for GameController. DO NOT EDIT.

package gamecontroller

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [GCControllerButtonInput] class.
var (
	GCControllerButtonInputClass     _GCControllerButtonInputClass
	GCControllerButtonInputClassOnce sync.Once
)

func getGCControllerButtonInputClass() _GCControllerButtonInputClass {
	GCControllerButtonInputClassOnce.Do(func() {
		GCControllerButtonInputClass = _GCControllerButtonInputClass{objc.GetClass("GCControllerButtonInput")}
	})
	return GCControllerButtonInputClass
}

type _GCControllerButtonInputClass struct {
	class objc.Class
}

// An interface definition for the [GCControllerButtonInput] class.
type IGCControllerButtonInput interface {
	IGCControllerElement
}

// A control element that represents a button touch or press.
//
// A object represents a button on a controller that can report either analog or digital values.
//
// [Full Topic]: https://developer.apple.com/documentation/GameController/GCControllerButtonInput
type GCControllerButtonInput struct {
	GCControllerElement
}

// GCControllerButtonInputFrom constructs a [GCControllerButtonInput] from an unsafe.Pointer.
//
// A control element that represents a button touch or press.
func GCControllerButtonInputFrom(ptr unsafe.Pointer) GCControllerButtonInput {
	return GCControllerButtonInput{
		GCControllerElement: GCControllerElementFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (gc _GCControllerButtonInputClass) Alloc() GCControllerButtonInput {
	rv := objc.Send[GCControllerButtonInput](objc.ID(gc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (gc _GCControllerButtonInputClass) New() GCControllerButtonInput {
	rv := objc.Send[GCControllerButtonInput](objc.ID(gc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (g_ GCControllerButtonInput) Init() GCControllerButtonInput {
	rv := objc.Send[GCControllerButtonInput](g_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (g_ GCControllerButtonInput) Autorelease() GCControllerButtonInput {
	rv := objc.Send[GCControllerButtonInput](g_.ID, objc.Sel("autorelease"))
	return rv
}

// NewGCControllerButtonInput creates a new GCControllerButtonInput instance.
func NewGCControllerButtonInput() GCControllerButtonInput {
	return getGCControllerButtonInputClass().New()
}


// A Boolean value that indicates whether the user is touching the button.
//
// [Full Topic]: https://developer.apple.com/documentation/GameController/GCControllerButtonInput/isTouched
func (g_ GCControllerButtonInput) Touched() bool {
	rv := objc.Send[bool](g_.ID, objc.Sel("touched"))
	return rv
}

// The block that the element calls when the user touches the button.
//
// [Full Topic]: https://developer.apple.com/documentation/GameController/GCControllerButtonInput/touchedChangedHandler
func (g_ GCControllerButtonInput) TouchedChangedHandler() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](g_.ID, objc.Sel("touchedChangedHandler"))
	return rv
}


// SetTouchedChangedHandler sets the value of the touchedChangedHandler property.
// The block that the element calls when the user touches the button.

//
// [Full Topic]: https://developer.apple.com/documentation/GameController/GCControllerButtonInput/touchedChangedHandler
func (g_ GCControllerButtonInput) SetTouchedChangedHandler(value unsafe.Pointer) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setTouchedChangedHandler:"), value)
}
// The block that the element calls when the user changes the level of pressure on the button.
//
// [Full Topic]: https://developer.apple.com/documentation/GameController/GCControllerButtonInput/valueChangedHandler
func (g_ GCControllerButtonInput) ValueChangedHandler() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](g_.ID, objc.Sel("valueChangedHandler"))
	return rv
}


// SetValueChangedHandler sets the value of the valueChangedHandler property.
// The block that the element calls when the user changes the level of pressure on the button.

//
// [Full Topic]: https://developer.apple.com/documentation/GameController/GCControllerButtonInput/valueChangedHandler
func (g_ GCControllerButtonInput) SetValueChangedHandler(value unsafe.Pointer) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setValueChangedHandler:"), value)
}


