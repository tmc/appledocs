// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [Stepper] class.
var (
	StepperClass     _StepperClass
	StepperClassOnce sync.Once
)

func getStepperClass() _StepperClass {
	StepperClassOnce.Do(func() {
		StepperClass = _StepperClass{objc.GetClass("NSStepper")}
	})
	return StepperClass
}

type _StepperClass struct {
	class objc.Class
}

// An interface definition for the [Stepper] class.
type IStepper interface {
	IControl
	ValueWraps() bool
	SetValueWraps(value bool)
	Autorepeat() bool
	SetAutorepeat(value bool)
	Increment() float64
	SetIncrement(value float64)
	MaxValue() float64
	SetMaxValue(value float64)
	MinValue() float64
	SetMinValue(value float64)
}

// An interface with up and down arrow buttons for incrementing or decrementing a value.
//
// A stepper consists of two small arrows that can increment and decrement a value that appears beside it, such as a date or time. The illustration below shows a stepper to the right of a text field, which would show the stepper’s value. The class uses the class to implement its user interface.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSStepper
type Stepper struct {
	Control
}

// StepperFrom constructs a [Stepper] from an unsafe.Pointer.
//
// An interface with up and down arrow buttons for incrementing or decrementing a value.
func StepperFrom(ptr unsafe.Pointer) Stepper {
	return Stepper{
		Control: ControlFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (sc _StepperClass) Alloc() Stepper {
	rv := objc.Send[Stepper](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (sc _StepperClass) New() Stepper {
	rv := objc.Send[Stepper](objc.ID(sc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (s_ Stepper) Init() Stepper {
	rv := objc.Send[Stepper](s_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s_ Stepper) Autorelease() Stepper {
	rv := objc.Send[Stepper](s_.ID, objc.Sel("autorelease"))
	return rv
}

// NewStepper creates a new Stepper instance.
func NewStepper() Stepper {
	return getStepperClass().New()
}


// A Boolean value that indicates whether the stepper wraps around the minimum and maximum values.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSStepper/valueWraps
func (s_ Stepper) ValueWraps() bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("valueWraps"))
	return rv
}


// SetValueWraps sets the value of the valueWraps property.
// A Boolean value that indicates whether the stepper wraps around the minimum and maximum values.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSStepper/valueWraps
func (s_ Stepper) SetValueWraps(value bool) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setValueWraps:"), value)
}

// A Boolean value that indicates how the stepper responds to mouse events.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsstepper/autorepeat
func (s_ Stepper) Autorepeat() bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("autorepeat"))
	return rv
}


// SetAutorepeat sets the value of the autorepeat property.
// A Boolean value that indicates how the stepper responds to mouse events.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsstepper/autorepeat
func (s_ Stepper) SetAutorepeat(value bool) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setAutorepeat:"), value)
}

// The amount by which the receiver changes with each increment or decrement.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsstepper/increment
func (s_ Stepper) Increment() float64 {
	rv := objc.Send[float64](s_.ID, objc.Sel("increment"))
	return rv
}


// SetIncrement sets the value of the increment property.
// The amount by which the receiver changes with each increment or decrement.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsstepper/increment
func (s_ Stepper) SetIncrement(value float64) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setIncrement:"), value)
}

// The stepper’s maximum value.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsstepper/maxvalue
func (s_ Stepper) MaxValue() float64 {
	rv := objc.Send[float64](s_.ID, objc.Sel("maxValue"))
	return rv
}


// SetMaxValue sets the value of the maxValue property.
// The stepper’s maximum value.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsstepper/maxvalue
func (s_ Stepper) SetMaxValue(value float64) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setMaxValue:"), value)
}

// The stepper’s minimum value.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsstepper/minvalue
func (s_ Stepper) MinValue() float64 {
	rv := objc.Send[float64](s_.ID, objc.Sel("minValue"))
	return rv
}


// SetMinValue sets the value of the minValue property.
// The stepper’s minimum value.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsstepper/minvalue
func (s_ Stepper) SetMinValue(value float64) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setMinValue:"), value)
}



