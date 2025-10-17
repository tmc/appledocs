// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [Stepper] class.
var StepperClass objc.Class

func init() {
	StepperClass = objc.GetClass("NSStepper")
}

type Stepper struct {
	objc.ID
}

func StepperFrom(ptr unsafe.Pointer) Stepper {
	return Stepper{
		ID: objc.ID(ptr),
	}
}




