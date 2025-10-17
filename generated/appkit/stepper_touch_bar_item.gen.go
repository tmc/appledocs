// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [StepperTouchBarItem] class.
var stepperTouchBarItemClass = _StepperTouchBarItemClass{objc.GetClass("NSStepperTouchBarItem")}

type _StepperTouchBarItemClass struct {
	class objc.Class
}

// An interface definition for the [StepperTouchBarItem] class.
type IStepperTouchBarItem interface {
	ITouchBarItem
}

// A bar item that provides a stepper control for incrementing or decrementing a value. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSStepperTouchBarItem

type StepperTouchBarItem struct {
	TouchBarItem
}

// StepperTouchBarItemFrom constructs a [StepperTouchBarItem] from an unsafe.Pointer.
//
// A bar item that provides a stepper control for incrementing or decrementing a value.
func StepperTouchBarItemFrom(ptr unsafe.Pointer) StepperTouchBarItem {
	return StepperTouchBarItem{
		TouchBarItem: TouchBarItemFrom(ptr),
	}
}



