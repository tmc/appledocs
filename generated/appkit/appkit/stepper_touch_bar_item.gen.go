// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [StepperTouchBarItem] class.
var StepperTouchBarItemClass objc.Class

func init() {
	StepperTouchBarItemClass = objc.GetClass("NSStepperTouchBarItem")
}

type StepperTouchBarItem struct {
	objc.ID
}

func StepperTouchBarItemFrom(ptr unsafe.Pointer) StepperTouchBarItem {
	return StepperTouchBarItem{
		ID: objc.ID(ptr),
	}
}




