// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [StepperCell] class.
var StepperCellClass objc.Class

func init() {
	StepperCellClass = objc.GetClass("NSStepperCell")
}

type StepperCell struct {
	objc.ID
}

func StepperCellFrom(ptr unsafe.Pointer) StepperCell {
	return StepperCell{
		ID: objc.ID(ptr),
	}
}



