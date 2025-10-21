// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [StepperTouchBarItem] class.
var (
	StepperTouchBarItemClass     _StepperTouchBarItemClass
	StepperTouchBarItemClassOnce sync.Once
)

func getStepperTouchBarItemClass() _StepperTouchBarItemClass {
	StepperTouchBarItemClassOnce.Do(func() {
		StepperTouchBarItemClass = _StepperTouchBarItemClass{objc.GetClass("NSStepperTouchBarItem")}
	})
	return StepperTouchBarItemClass
}

type _StepperTouchBarItemClass struct {
	class objc.Class
}

// An interface definition for the [StepperTouchBarItem] class.
type IStepperTouchBarItem interface {
	ITouchBarItem
}

// A bar item that provides a stepper control for incrementing or decrementing a value.
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


//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSStepperTouchBarItem/init(identifier:formatter:)
func NewStepperTouchBarItemWithIdentifierFormatter(identifier unsafe.Pointer, formatter unsafe.Pointer) StepperTouchBarItem {
	rv := objc.Send[StepperTouchBarItem](objc.ID(getStepperTouchBarItemClass().class), objc.Sel("stepperTouchBarItemWithIdentifier:formatter:"), identifier, formatter)
	return rv
}


//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSStepperTouchBarItem/init(identifier:formatter:)
func (sc _StepperTouchBarItemClass) StepperTouchBarItemWithIdentifierFormatter(identifier unsafe.Pointer, formatter unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(sc.class), objc.Sel("stepperTouchBarItemWithIdentifier:formatter:"), identifier, formatter)
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSStepperTouchBarItem/maxValue
func (s_ StepperTouchBarItem) MaxValue() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("maxValue"))
	return rv
}


// SetMaxValue sets the value of the maxValue property.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSStepperTouchBarItem/maxValue
func (s_ StepperTouchBarItem) SetMaxValue(value unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setMaxValue:"), value)
}


