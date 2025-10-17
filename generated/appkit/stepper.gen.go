// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [Stepper] class.
var stepperClass = _StepperClass{objc.GetClass("NSStepper")}

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



