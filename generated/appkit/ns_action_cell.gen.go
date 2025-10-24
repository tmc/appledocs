// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class NSActionCell */


/* debug [class_header]: Header for NSActionCell */
// The class instance for the [ActionCell] class.
var (
	ActionCellClass     _ActionCellClass
	ActionCellClassOnce sync.Once
)

func getActionCellClass() _ActionCellClass {
	ActionCellClassOnce.Do(func() {
		ActionCellClass = _ActionCellClass{objc.GetClass("NSActionCell")}
	})
	return ActionCellClass
}

type _ActionCellClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for ActionCell */
// An interface definition for the [ActionCell] class.
type IActionCell interface {
	ICell
	
/* debug [class_interface_properties]: Properties for ActionCell */
	// properties:
	Action() objc.SEL
	SetAction(value objc.SEL)
	Tag() int
	SetTag(value int)
	Target() objc.ID
	SetTarget(value objc.ID)
	ControlView() IView
	SetControlView(value IView)
	FloatValue() float32
	SetFloatValue(value float32)
	IntValue() objectivec.IObject
	SetIntValue(value objectivec.IObject)
	IntegerValue() int
	SetIntegerValue(value int)
	StringValue() objc.IObject /* cross-framework: NSString */
	SetStringValue(value objc.IObject /* cross-framework: NSString */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for ActionCell */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for ActionCell */
// Alloc allocates a new instance without initialization.
func (ac _ActionCellClass) Alloc() ActionCell {
	rv := objc.Send[ActionCell](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (ac _ActionCellClass) New() ActionCell {
	rv := objc.Send[ActionCell](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ ActionCell) Init() ActionCell {
	rv := objc.Send[ActionCell](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ ActionCell) Autorelease() ActionCell {
	rv := objc.Send[ActionCell](a_.ID, objc.Sel("autorelease"))
	return rv
}

// NewActionCell creates a new ActionCell instance.
func NewActionCell() ActionCell {
	return getActionCellClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for ActionCell */
// An active area inside a control.
//
// An does three things: it displays text or an icon; it provides the target object and action method used by its object; and it handles mouse (cursor) tracking by properly highlighting its area and sending action messages to its target based on cursor movement. The of an is the view in which the receiver was last drawn.


// An active area inside a control.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSActionCell
type ActionCell struct {
	Cell
}

// ActionCellFrom constructs a [ActionCell] from an unsafe.Pointer.
//
// An active area inside a control.
func ActionCellFrom(ptr unsafe.Pointer) ActionCell {
	return ActionCell{
		Cell: CellFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for ActionCell *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for ActionCell */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for ActionCell */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for ActionCell */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for ActionCell */

// Returns the receiver’s action-message selector.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSActionCell/action
func (a_ ActionCell) Action() objc.SEL {
	rv := objc.Send[objc.SEL](a_.ID, objc.Sel("action"))
	return rv
}/* debug [instance_properties/getter]: action */


// Returns the receiver’s action-message selector.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSActionCell/action
func (a_ ActionCell) SetAction(value objc.SEL) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setAction:"), value)
}/* debug [instance_properties/setter]: action */


// Returns the receiver’s tag.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSActionCell/tag
func (a_ ActionCell) Tag() int {
	rv := objc.Send[int](a_.ID, objc.Sel("tag"))
	return rv
}/* debug [instance_properties/getter]: tag */


// Returns the receiver’s tag.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSActionCell/tag
func (a_ ActionCell) SetTag(value int) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setTag:"), value)
}/* debug [instance_properties/setter]: tag */


// Returns the receiver’s target object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSActionCell/target
func (a_ ActionCell) Target() objc.ID {
	rv := objc.Send[objc.ID](a_.ID, objc.Sel("target"))
	return rv
}/* debug [instance_properties/getter]: target */


// Returns the receiver’s target object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSActionCell/target
func (a_ ActionCell) SetTarget(value objc.ID) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setTarget:"), value)
}/* debug [instance_properties/setter]: target */


// The view associated with the cell.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscell/controlview
func (a_ ActionCell) ControlView() IView {
	rv := objc.Send[View](a_.ID, objc.Sel("controlView"))
	return rv
}/* debug [instance_properties/getter]: controlView */


// The view associated with the cell.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscell/controlview
func (a_ ActionCell) SetControlView(value IView) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setControlView:"), value)
}/* debug [instance_properties/setter]: controlView */


// The cell’s value as a single-precision floating-point number.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscell/floatvalue
func (a_ ActionCell) FloatValue() float32 {
	rv := objc.Send[float32](a_.ID, objc.Sel("floatValue"))
	return rv
}/* debug [instance_properties/getter]: floatValue */


// The cell’s value as a single-precision floating-point number.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscell/floatvalue
func (a_ ActionCell) SetFloatValue(value float32) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setFloatValue:"), value)
}/* debug [instance_properties/setter]: floatValue */


// The cell’s value as an integer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscell/intvalue
func (a_ ActionCell) IntValue() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](a_.ID, objc.Sel("intValue"))
	return rv
}/* debug [instance_properties/getter]: intValue */


// The cell’s value as an integer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscell/intvalue
func (a_ ActionCell) SetIntValue(value objectivec.IObject) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setIntValue:"), value)
}/* debug [instance_properties/setter]: intValue */


// The cell’s value as an integer value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscell/integervalue
func (a_ ActionCell) IntegerValue() int {
	rv := objc.Send[int](a_.ID, objc.Sel("integerValue"))
	return rv
}/* debug [instance_properties/getter]: integerValue */


// The cell’s value as an integer value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscell/integervalue
func (a_ ActionCell) SetIntegerValue(value int) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setIntegerValue:"), value)
}/* debug [instance_properties/setter]: integerValue */


// The cell’s value as a string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscell/stringvalue
func (a_ ActionCell) StringValue() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](a_.ID, objc.Sel("stringValue"))
	return rv
}/* debug [instance_properties/getter]: stringValue */


// The cell’s value as a string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscell/stringvalue
func (a_ ActionCell) SetStringValue(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setStringValue:"), value)
}/* debug [instance_properties/setter]: stringValue */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class NSActionCell */



