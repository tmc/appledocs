// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [Stepper] class.
var (
	stepperClass     _StepperClass
	stepperClassOnce sync.Once
)

func getStepperClass() _StepperClass {
	stepperClassOnce.Do(func() {
		stepperClass = _StepperClass{objc.GetClass("NSStepper")}
	})
	return stepperClass
}

type _StepperClass struct {
	class objc.Class
}

// An interface definition for the [Stepper] class.
type IStepper interface {
	IControl
}

// An interface with up and down arrow buttons for incrementing or decrementing a value. [Full Topic]
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




