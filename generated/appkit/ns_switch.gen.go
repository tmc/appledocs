// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class NSSwitch */


/* debug [class_header]: Header for NSSwitch */
// The class instance for the [Switch] class.
var (
	SwitchClass     _SwitchClass
	SwitchClassOnce sync.Once
)

func getSwitchClass() _SwitchClass {
	SwitchClassOnce.Do(func() {
		SwitchClass = _SwitchClass{objc.GetClass("NSSwitch")}
	})
	return SwitchClass
}

type _SwitchClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for Switch */
// An interface definition for the [Switch] class.
type ISwitch interface {
	IControl
	
/* debug [class_interface_properties]: Properties for Switch */
	// properties:
	State() ControlStateValue /* typedef */
	SetState(value ControlStateValue /* typedef */)
	Action() objectivec.IObject
	SetAction(value objectivec.IObject)
	Cell() ICell
	SetCell(value ICell)
	IsContinuous() bool
	SetIsContinuous(value bool)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for Switch */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for Switch */
// Alloc allocates a new instance without initialization.
func (sc _SwitchClass) Alloc() Switch {
	rv := objc.Send[Switch](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (sc _SwitchClass) New() Switch {
	rv := objc.Send[Switch](objc.ID(sc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (s_ Switch) Init() Switch {
	rv := objc.Send[Switch](s_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s_ Switch) Autorelease() Switch {
	rv := objc.Send[Switch](s_.ID, objc.Sel("autorelease"))
	return rv
}

// NewSwitch creates a new Switch instance.
func NewSwitch() Switch {
	return getSwitchClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for Switch */
// A control that offers a binary choice.
//
// The class provides a simple interface for displaying and toggling a Boolean state, such as on/off. A switch toggles its and sends its when clicked, activated through the keyboard, or tapped in the Touch Bar. also allows dragging between states, and if is , the switch sends its for each change in position during the drag. doesn’t use an instance of to provide its functionality. The class property and instance property both return , and they ignore attempts to set a non- value. For design guidance, see Human Interface Guidelines > .


// A control that offers a binary choice.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSwitch
type Switch struct {
	Control
}

// SwitchFrom constructs a [Switch] from an unsafe.Pointer.
//
// A control that offers a binary choice.
func SwitchFrom(ptr unsafe.Pointer) Switch {
	return Switch{
		Control: ControlFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for Switch *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for Switch */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for Switch */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for Switch */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for Switch */

// The current position of the switch.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSwitch/state
func (s_ Switch) State() ControlStateValue /* typedef */ {
	rv := objc.Send[int](s_.ID, objc.Sel("state"))
	return rv
}/* debug [instance_properties/getter]: state */


// The current position of the switch.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSwitch/state
func (s_ Switch) SetState(value ControlStateValue /* typedef */) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setState:"), value)
}/* debug [instance_properties/setter]: state */


// The default action-message selector associated with the control.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscontrol/action
func (s_ Switch) Action() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](s_.ID, objc.Sel("action"))
	return rv
}/* debug [instance_properties/getter]: action */


// The default action-message selector associated with the control.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscontrol/action
func (s_ Switch) SetAction(value objectivec.IObject) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setAction:"), value)
}/* debug [instance_properties/setter]: action */


// The receiver’s cell object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscontrol/cell
func (s_ Switch) Cell() ICell {
	rv := objc.Send[Cell](s_.ID, objc.Sel("cell"))
	return rv
}/* debug [instance_properties/getter]: cell */


// The receiver’s cell object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscontrol/cell
func (s_ Switch) SetCell(value ICell) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setCell:"), value)
}/* debug [instance_properties/setter]: cell */


// A Boolean value indicating whether the receiver’s cell sends its action message continuously to its target during mouse tracking.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscontrol/iscontinuous
func (s_ Switch) IsContinuous() bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("isContinuous"))
	return rv
}/* debug [instance_properties/getter]: isContinuous */


// A Boolean value indicating whether the receiver’s cell sends its action message continuously to its target during mouse tracking.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscontrol/iscontinuous
func (s_ Switch) SetIsContinuous(value bool) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setIsContinuous:"), value)
}/* debug [instance_properties/setter]: isContinuous */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class NSSwitch */



