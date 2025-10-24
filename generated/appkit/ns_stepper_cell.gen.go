// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class NSStepperCell */


/* debug [class_header]: Header for NSStepperCell */
// The class instance for the [StepperCell] class.
var (
	StepperCellClass     _StepperCellClass
	StepperCellClassOnce sync.Once
)

func getStepperCellClass() _StepperCellClass {
	StepperCellClassOnce.Do(func() {
		StepperCellClass = _StepperCellClass{objc.GetClass("NSStepperCell")}
	})
	return StepperCellClass
}

type _StepperCellClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for StepperCell */
// An interface definition for the [StepperCell] class.
type IStepperCell interface {
	IActionCell
	
/* debug [class_interface_properties]: Properties for StepperCell */
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

	
/* debug [class_interface_methods]: Methods for StepperCell */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for StepperCell */
// Alloc allocates a new instance without initialization.
func (sc _StepperCellClass) Alloc() StepperCell {
	rv := objc.Send[StepperCell](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for StepperCell */
// An object controls the appearance and behavior of an object.


// An object controls the appearance and behavior of an object.
//
// [Full Topic]
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for StepperCell *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for StepperCell */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for StepperCell */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for StepperCell */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for StepperCell */

// A Boolean value indicating how the receiver responds to mouse events.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSStepperCell/autorepeat
func (s_ StepperCell) Autorepeat() bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("autorepeat"))
	return rv
}/* debug [instance_properties/getter]: autorepeat */


// A Boolean value indicating how the receiver responds to mouse events.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSStepperCell/autorepeat
func (s_ StepperCell) SetAutorepeat(value bool) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setAutorepeat:"), value)
}/* debug [instance_properties/setter]: autorepeat */


// The amount by which the receiver will change per increment or decrement.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSStepperCell/increment
func (s_ StepperCell) Increment() float64 {
	rv := objc.Send[float64](s_.ID, objc.Sel("increment"))
	return rv
}/* debug [instance_properties/getter]: increment */


// The amount by which the receiver will change per increment or decrement.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSStepperCell/increment
func (s_ StepperCell) SetIncrement(value float64) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setIncrement:"), value)
}/* debug [instance_properties/setter]: increment */


// The maximum value for the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSStepperCell/maxValue
func (s_ StepperCell) MaxValue() float64 {
	rv := objc.Send[float64](s_.ID, objc.Sel("maxValue"))
	return rv
}/* debug [instance_properties/getter]: maxValue */


// The maximum value for the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSStepperCell/maxValue
func (s_ StepperCell) SetMaxValue(value float64) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setMaxValue:"), value)
}/* debug [instance_properties/setter]: maxValue */


// The minimum value for the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSStepperCell/minValue
func (s_ StepperCell) MinValue() float64 {
	rv := objc.Send[float64](s_.ID, objc.Sel("minValue"))
	return rv
}/* debug [instance_properties/getter]: minValue */


// The minimum value for the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSStepperCell/minValue
func (s_ StepperCell) SetMinValue(value float64) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setMinValue:"), value)
}/* debug [instance_properties/setter]: minValue */


// A Boolean value indicating whether the receiver wraps around the minimum and maximum values.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSStepperCell/valueWraps
func (s_ StepperCell) ValueWraps() bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("valueWraps"))
	return rv
}/* debug [instance_properties/getter]: valueWraps */


// A Boolean value indicating whether the receiver wraps around the minimum and maximum values.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSStepperCell/valueWraps
func (s_ StepperCell) SetValueWraps(value bool) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setValueWraps:"), value)
}/* debug [instance_properties/setter]: valueWraps */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class NSStepperCell */



