// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [StepperCell] class.
var (
	stepperCellClass     _StepperCellClass
	stepperCellClassOnce sync.Once
)

func getStepperCellClass() _StepperCellClass {
	stepperCellClassOnce.Do(func() {
		stepperCellClass = _StepperCellClass{objc.GetClass("NSStepperCell")}
	})
	return stepperCellClass
}

type _StepperCellClass struct {
	class objc.Class
}

// An interface definition for the [StepperCell] class.
type IStepperCell interface {
	IActionCell
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
// Alloc allocates a new instance without initialization.
func (sc _StepperCellClass) Alloc() StepperCell {
	rv := objc.Send[StepperCell](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new instance with a +1 retain count.
func (sc _StepperCellClass) New() StepperCell {
	rv := objc.Send[StepperCell](objc.ID(sc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (s_ StepperCell) Init() StepperCell {
	rv := objc.Send[StepperCell](s_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s_ StepperCell) Autorelease() StepperCell {
	rv := objc.Send[StepperCell](s_.ID, objc.Sel("autorelease"))
	return rv
}

// NewStepperCell creates a new StepperCell instance.
func NewStepperCell() StepperCell {
	return getStepperCellClass().New()
}




