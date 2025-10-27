// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
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
	Action() objc.SEL
	SetAction(value objc.SEL)
	CustomizationLabel() foundation.foundation.INSString
	SetCustomizationLabel(value foundation.foundation.INSString)
	Increment() float64
	SetIncrement(value float64)
	MaxValue() float64
	SetMaxValue(value float64)
	MinValue() float64
	SetMinValue(value float64)
	Target() objc.ID
	SetTarget(value objc.ID)
	Value() float64
	SetValue(value float64)


	

	// methods:


}





// Alloc allocates a new instance without initialization.
func (sc _StepperTouchBarItemClass) Alloc() StepperTouchBarItem {
	rv := objc.Send[StepperTouchBarItem](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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






// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSStepperTouchBarItem/init(identifier:drawingHandler:)
func NewStepperTouchBarItemWithIdentifierDrawingHandler(identifier TouchBarItemIdentifier, drawingHandler unsafe.Pointer) StepperTouchBarItem {
	rv := objc.Send[StepperTouchBarItem](objc.ID(getStepperTouchBarItemClass().class), objc.Sel("stepperTouchBarItemWithIdentifier:drawingHandler:"), identifier, drawingHandler)
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSStepperTouchBarItem/init(identifier:formatter:)
func NewStepperTouchBarItemWithIdentifierFormatter(identifier TouchBarItemIdentifier, formatter foundation.Formatter) StepperTouchBarItem {
	rv := objc.Send[StepperTouchBarItem](objc.ID(getStepperTouchBarItemClass().class), objc.Sel("stepperTouchBarItemWithIdentifier:formatter:"), identifier, formatter)
	return rv
}







// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSStepperTouchBarItem/init(identifier:drawingHandler:)
func (sc _StepperTouchBarItemClass) StepperTouchBarItemWithIdentifierDrawingHandler(identifier TouchBarItemIdentifier, drawingHandler unsafe.Pointer) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(sc.class), objc.Sel("stepperTouchBarItemWithIdentifier:drawingHandler:"), identifier, drawingHandler)
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSStepperTouchBarItem/init(identifier:formatter:)
func (sc _StepperTouchBarItemClass) StepperTouchBarItemWithIdentifierFormatter(identifier TouchBarItemIdentifier, formatter foundation.Formatter) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(sc.class), objc.Sel("stepperTouchBarItemWithIdentifier:formatter:"), identifier, formatter)
	return rv
}

















// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSStepperTouchBarItem/action
func (s_ StepperTouchBarItem) Action() objc.SEL {
	rv := objc.Send[objc.SEL](s_.ID, objc.Sel("action"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSStepperTouchBarItem/action
func (s_ StepperTouchBarItem) SetAction(value objc.SEL) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setAction:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSStepperTouchBarItem/customizationLabel
func (s_ StepperTouchBarItem) CustomizationLabel() foundation.foundation.INSString {
	rv := objc.Send[foundation.NSString](s_.ID, objc.Sel("customizationLabel"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSStepperTouchBarItem/customizationLabel
func (s_ StepperTouchBarItem) SetCustomizationLabel(value foundation.foundation.INSString) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setCustomizationLabel:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSStepperTouchBarItem/increment
func (s_ StepperTouchBarItem) Increment() float64 {
	rv := objc.Send[float64](s_.ID, objc.Sel("increment"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSStepperTouchBarItem/increment
func (s_ StepperTouchBarItem) SetIncrement(value float64) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setIncrement:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSStepperTouchBarItem/maxValue
func (s_ StepperTouchBarItem) MaxValue() float64 {
	rv := objc.Send[float64](s_.ID, objc.Sel("maxValue"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSStepperTouchBarItem/maxValue
func (s_ StepperTouchBarItem) SetMaxValue(value float64) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setMaxValue:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSStepperTouchBarItem/minValue
func (s_ StepperTouchBarItem) MinValue() float64 {
	rv := objc.Send[float64](s_.ID, objc.Sel("minValue"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSStepperTouchBarItem/minValue
func (s_ StepperTouchBarItem) SetMinValue(value float64) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setMinValue:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSStepperTouchBarItem/target
func (s_ StepperTouchBarItem) Target() objc.ID {
	rv := objc.Send[objc.ID](s_.ID, objc.Sel("target"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSStepperTouchBarItem/target
func (s_ StepperTouchBarItem) SetTarget(value objc.ID) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setTarget:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSStepperTouchBarItem/value
func (s_ StepperTouchBarItem) Value() float64 {
	rv := objc.Send[float64](s_.ID, objc.Sel("value"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSStepperTouchBarItem/value
func (s_ StepperTouchBarItem) SetValue(value float64) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setValue:"), value)
}







