// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [StepperCell] class.
var stepperCellClass = _StepperCellClass{objc.GetClass("NSStepperCell")}

type _StepperCellClass struct {
	class objc.Class
}

// An object controls the appearance and behavior of an object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSStepperCell

type StepperCell struct {
	ActionCell
}

// StepperCellFrom constructs a [StepperCell] from an unsafe.Pointer.
//
// An object controls the appearance and behavior of an object.
func StepperCellFrom(ptr unsafe.Pointer) StepperCell {
	return StepperCell{
		ActionCell: ActionCellFrom(ptr),
	}
}



