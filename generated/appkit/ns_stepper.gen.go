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




