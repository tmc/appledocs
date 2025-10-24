// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class NSStepper */


/* debug [class_header]: Header for NSStepper */
// The class instance for the [Stepper] class.
var (
	StepperClass     _StepperClass
	StepperClassOnce sync.Once
)

func getStepperClass() _StepperClass {
	StepperClassOnce.Do(func() {
		StepperClass = _StepperClass{objc.GetClass("NSStepper")}
	})
	return StepperClass
}

type _StepperClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for Stepper */
// An interface definition for the [Stepper] class.
type IStepper interface {
	IControl
	
/* debug [class_interface_properties]: Properties for Stepper */
	// properties:
	Autorepeat() bool
	SetAutorepeat(value bool)
	Increment() float64
	SetIncrement(value float64)
	MaxValue() float64
	SetMaxValue(value float64)
	MinValue() float64
	SetMinValue(value float64)
	ValueWraps() bool
	SetValueWraps(value bool)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for Stepper */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for Stepper */
// Alloc allocates a new instance without initialization.
func (sc _StepperClass) Alloc() Stepper {
	rv := objc.Send[Stepper](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (sc _StepperClass) New() Stepper {
	rv := objc.Send[Stepper](objc.ID(sc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (s_ Stepper) Init() Stepper {
	rv := objc.Send[Stepper](s_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s_ Stepper) Autorelease() Stepper {
	rv := objc.Send[Stepper](s_.ID, objc.Sel("autorelease"))
	return rv
}

// NewStepper creates a new Stepper instance.
func NewStepper() Stepper {
	return getStepperClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for Stepper */
// An interface with up and down arrow buttons for incrementing or decrementing a value.
//
// A stepper consists of two small arrows that can increment and decrement a value that appears beside it, such as a date or time. The illustration below shows a stepper to the right of a text field, which would show the stepper’s value. The class uses the class to implement its user interface.


// An interface with up and down arrow buttons for incrementing or decrementing a value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSStepper
type Stepper struct {
	Control
}

// StepperFrom constructs a [Stepper] from an unsafe.Pointer.
//
// An interface with up and down arrow buttons for incrementing or decrementing a value.
func StepperFrom(ptr unsafe.Pointer) Stepper {
	return Stepper{
		Control: ControlFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for Stepper *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for Stepper */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for Stepper */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for Stepper */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for Stepper */

// A Boolean value that indicates how the stepper responds to mouse events.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSStepper/autorepeat
func (s_ Stepper) Autorepeat() bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("autorepeat"))
	return rv
}/* debug [instance_properties/getter]: autorepeat */


// A Boolean value that indicates how the stepper responds to mouse events.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSStepper/autorepeat
func (s_ Stepper) SetAutorepeat(value bool) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setAutorepeat:"), value)
}/* debug [instance_properties/setter]: autorepeat */


// The amount by which the receiver changes with each increment or decrement.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSStepper/increment
func (s_ Stepper) Increment() float64 {
	rv := objc.Send[float64](s_.ID, objc.Sel("increment"))
	return rv
}/* debug [instance_properties/getter]: increment */


// The amount by which the receiver changes with each increment or decrement.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSStepper/increment
func (s_ Stepper) SetIncrement(value float64) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setIncrement:"), value)
}/* debug [instance_properties/setter]: increment */


// The stepper’s maximum value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSStepper/maxValue
func (s_ Stepper) MaxValue() float64 {
	rv := objc.Send[float64](s_.ID, objc.Sel("maxValue"))
	return rv
}/* debug [instance_properties/getter]: maxValue */


// The stepper’s maximum value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSStepper/maxValue
func (s_ Stepper) SetMaxValue(value float64) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setMaxValue:"), value)
}/* debug [instance_properties/setter]: maxValue */


// The stepper’s minimum value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSStepper/minValue
func (s_ Stepper) MinValue() float64 {
	rv := objc.Send[float64](s_.ID, objc.Sel("minValue"))
	return rv
}/* debug [instance_properties/getter]: minValue */


// The stepper’s minimum value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSStepper/minValue
func (s_ Stepper) SetMinValue(value float64) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setMinValue:"), value)
}/* debug [instance_properties/setter]: minValue */


// A Boolean value that indicates whether the stepper wraps around the minimum and maximum values.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSStepper/valueWraps
func (s_ Stepper) ValueWraps() bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("valueWraps"))
	return rv
}/* debug [instance_properties/getter]: valueWraps */


// A Boolean value that indicates whether the stepper wraps around the minimum and maximum values.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSStepper/valueWraps
func (s_ Stepper) SetValueWraps(value bool) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setValueWraps:"), value)
}/* debug [instance_properties/setter]: valueWraps */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class NSStepper */



