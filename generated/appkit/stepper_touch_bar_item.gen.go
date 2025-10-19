// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [StepperTouchBarItem] class.
var (
	stepperTouchBarItemClass     _StepperTouchBarItemClass
	stepperTouchBarItemClassOnce sync.Once
)

func getStepperTouchBarItemClass() _StepperTouchBarItemClass {
	stepperTouchBarItemClassOnce.Do(func() {
		stepperTouchBarItemClass = _StepperTouchBarItemClass{objc.GetClass("NSStepperTouchBarItem")}
	})
	return stepperTouchBarItemClass
}

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

// Alloc allocates a new instance without initialization.
func (sc _StepperTouchBarItemClass) Alloc() StepperTouchBarItem {
	rv := objc.Send[StepperTouchBarItem](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (sc _StepperTouchBarItemClass) New() StepperTouchBarItem {
	rv := objc.Send[StepperTouchBarItem](objc.ID(sc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (s_ StepperTouchBarItem) Init() StepperTouchBarItem {
	rv := objc.Send[StepperTouchBarItem](s_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s_ StepperTouchBarItem) Autorelease() StepperTouchBarItem {
	rv := objc.Send[StepperTouchBarItem](s_.ID, objc.Sel("autorelease"))
	return rv
}

// NewStepperTouchBarItem creates a new StepperTouchBarItem instance.
func NewStepperTouchBarItem() StepperTouchBarItem {
	return getStepperTouchBarItemClass().New()
}




