// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class NSStepperTouchBarItem */


/* debug [class_header]: Header for NSStepperTouchBarItem */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for StepperTouchBarItem */
// An interface definition for the [StepperTouchBarItem] class.
type IStepperTouchBarItem interface {
	ITouchBarItem
	
/* debug [class_interface_properties]: Properties for StepperTouchBarItem */
	// properties:
	Action() objc.SEL
	SetAction(value objc.SEL)
	CustomizationLabel() objc.IObject /* cross-framework: NSString */
	SetCustomizationLabel(value objc.IObject /* cross-framework: NSString */)
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
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for StepperTouchBarItem */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for StepperTouchBarItem */
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for StepperTouchBarItem */
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for StepperTouchBarItem */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSStepperTouchBarItem/init(identifier:drawingHandler:)
func NewStepperTouchBarItemWithIdentifierDrawingHandler(identifier TouchBarItemIdentifier /* typedef */, drawingHandler unsafe.Pointer) StepperTouchBarItem {
	rv := objc.Send[StepperTouchBarItem](objc.ID(getStepperTouchBarItemClass().class), objc.Sel("stepperTouchBarItemWithIdentifier:drawingHandler:"), identifier, drawingHandler)
	return rv
}/* debug [class_init_methods/constructor]: NewStepperTouchBarItemWithIdentifierDrawingHandler */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSStepperTouchBarItem/init(identifier:formatter:)
func NewStepperTouchBarItemWithIdentifierFormatter(identifier TouchBarItemIdentifier /* typedef */, formatter objectivec.IObject) StepperTouchBarItem {
	rv := objc.Send[StepperTouchBarItem](objc.ID(getStepperTouchBarItemClass().class), objc.Sel("stepperTouchBarItemWithIdentifier:formatter:"), identifier, formatter)
	return rv
}/* debug [class_init_methods/constructor]: NewStepperTouchBarItemWithIdentifierFormatter */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for StepperTouchBarItem */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSStepperTouchBarItem/init(identifier:drawingHandler:)
func (sc _StepperTouchBarItemClass) StepperTouchBarItemWithIdentifierDrawingHandler(identifier TouchBarItemIdentifier /* typedef */, drawingHandler unsafe.Pointer) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(sc.class), objc.Sel("stepperTouchBarItemWithIdentifier:drawingHandler:"), identifier, drawingHandler)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=StepperTouchBarItemWithIdentifierDrawingHandler) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSStepperTouchBarItem/init(identifier:formatter:)
func (sc _StepperTouchBarItemClass) StepperTouchBarItemWithIdentifierFormatter(identifier TouchBarItemIdentifier /* typedef */, formatter objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(sc.class), objc.Sel("stepperTouchBarItemWithIdentifier:formatter:"), identifier, formatter)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=StepperTouchBarItemWithIdentifierFormatter) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for StepperTouchBarItem */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for StepperTouchBarItem */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for StepperTouchBarItem */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSStepperTouchBarItem/action
func (s_ StepperTouchBarItem) Action() objc.SEL {
	rv := objc.Send[objc.SEL](s_.ID, objc.Sel("action"))
	return rv
}/* debug [instance_properties/getter]: action */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSStepperTouchBarItem/action
func (s_ StepperTouchBarItem) SetAction(value objc.SEL) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setAction:"), value)
}/* debug [instance_properties/setter]: action */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSStepperTouchBarItem/customizationLabel
func (s_ StepperTouchBarItem) CustomizationLabel() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](s_.ID, objc.Sel("customizationLabel"))
	return rv
}/* debug [instance_properties/getter]: customizationLabel */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSStepperTouchBarItem/customizationLabel
func (s_ StepperTouchBarItem) SetCustomizationLabel(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setCustomizationLabel:"), value)
}/* debug [instance_properties/setter]: customizationLabel */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSStepperTouchBarItem/increment
func (s_ StepperTouchBarItem) Increment() float64 {
	rv := objc.Send[float64](s_.ID, objc.Sel("increment"))
	return rv
}/* debug [instance_properties/getter]: increment */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSStepperTouchBarItem/increment
func (s_ StepperTouchBarItem) SetIncrement(value float64) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setIncrement:"), value)
}/* debug [instance_properties/setter]: increment */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSStepperTouchBarItem/maxValue
func (s_ StepperTouchBarItem) MaxValue() float64 {
	rv := objc.Send[float64](s_.ID, objc.Sel("maxValue"))
	return rv
}/* debug [instance_properties/getter]: maxValue */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSStepperTouchBarItem/maxValue
func (s_ StepperTouchBarItem) SetMaxValue(value float64) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setMaxValue:"), value)
}/* debug [instance_properties/setter]: maxValue */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSStepperTouchBarItem/minValue
func (s_ StepperTouchBarItem) MinValue() float64 {
	rv := objc.Send[float64](s_.ID, objc.Sel("minValue"))
	return rv
}/* debug [instance_properties/getter]: minValue */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSStepperTouchBarItem/minValue
func (s_ StepperTouchBarItem) SetMinValue(value float64) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setMinValue:"), value)
}/* debug [instance_properties/setter]: minValue */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSStepperTouchBarItem/target
func (s_ StepperTouchBarItem) Target() objc.ID {
	rv := objc.Send[objc.ID](s_.ID, objc.Sel("target"))
	return rv
}/* debug [instance_properties/getter]: target */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSStepperTouchBarItem/target
func (s_ StepperTouchBarItem) SetTarget(value objc.ID) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setTarget:"), value)
}/* debug [instance_properties/setter]: target */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSStepperTouchBarItem/value
func (s_ StepperTouchBarItem) Value() float64 {
	rv := objc.Send[float64](s_.ID, objc.Sel("value"))
	return rv
}/* debug [instance_properties/getter]: value */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSStepperTouchBarItem/value
func (s_ StepperTouchBarItem) SetValue(value float64) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setValue:"), value)
}/* debug [instance_properties/setter]: value */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class NSStepperTouchBarItem */


