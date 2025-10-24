// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
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
	// properties:
	Action() unsafe.Pointer
	SetAction(value unsafe.Pointer)
	CustomizationLabel() objc.IObject /* cross-framework: NSString */
	SetCustomizationLabel(value objc.IObject /* cross-framework: NSString */)
	Increment() float64
	SetIncrement(value float64)
	MaxValue() float64
	SetMaxValue(value float64)
	MinValue() float64
	SetMinValue(value float64)
	Target() unsafe.Pointer
	SetTarget(value unsafe.Pointer)
	Value() float64
	SetValue(value float64)
	// methods:
}

// A bar item that provides a stepper control for incrementing or decrementing a value.


// A bar item that provides a stepper control for incrementing or decrementing a value.
//
// [Full Topic]
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



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssteppertouchbaritem/action
func (s_ StepperTouchBarItem) Action() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("action"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssteppertouchbaritem/action
func (s_ StepperTouchBarItem) SetAction(value unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setAction:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssteppertouchbaritem/customizationlabel
func (s_ StepperTouchBarItem) CustomizationLabel() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](s_.ID, objc.Sel("customizationLabel"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssteppertouchbaritem/customizationlabel
func (s_ StepperTouchBarItem) SetCustomizationLabel(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setCustomizationLabel:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssteppertouchbaritem/increment
func (s_ StepperTouchBarItem) Increment() float64 {
	rv := objc.Send[float64](s_.ID, objc.Sel("increment"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssteppertouchbaritem/increment
func (s_ StepperTouchBarItem) SetIncrement(value float64) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setIncrement:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssteppertouchbaritem/maxvalue
func (s_ StepperTouchBarItem) MaxValue() float64 {
	rv := objc.Send[float64](s_.ID, objc.Sel("maxValue"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssteppertouchbaritem/maxvalue
func (s_ StepperTouchBarItem) SetMaxValue(value float64) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setMaxValue:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssteppertouchbaritem/minvalue
func (s_ StepperTouchBarItem) MinValue() float64 {
	rv := objc.Send[float64](s_.ID, objc.Sel("minValue"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssteppertouchbaritem/minvalue
func (s_ StepperTouchBarItem) SetMinValue(value float64) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setMinValue:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssteppertouchbaritem/target
func (s_ StepperTouchBarItem) Target() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("target"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssteppertouchbaritem/target
func (s_ StepperTouchBarItem) SetTarget(value unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setTarget:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssteppertouchbaritem/value
func (s_ StepperTouchBarItem) Value() float64 {
	rv := objc.Send[float64](s_.ID, objc.Sel("value"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssteppertouchbaritem/value
func (s_ StepperTouchBarItem) SetValue(value float64) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setValue:"), value)
}



