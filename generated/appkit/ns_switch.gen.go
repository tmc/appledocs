// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

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

// An interface definition for the [Switch] class.
type ISwitch interface {
	IControl
	// properties:
	State() objc.IObject /* cross-framework: ControlStateValue */
	SetState(value objc.IObject /* cross-framework: ControlStateValue */)
	Action() unsafe.Pointer
	SetAction(value unsafe.Pointer)
	Cell() ICell
	SetCell(value ICell)
	IsContinuous() bool
	SetIsContinuous(value bool)
	// methods:
}

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

// Alloc allocates a new instance without initialization.
func (sc _SwitchClass) Alloc() Switch {
	rv := objc.Send[Switch](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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



// The current position of the switch.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSwitch/state
func (s_ Switch) State() objc.IObject /* cross-framework: ControlStateValue */ {
	rv := objc.Send[objc.ID](s_.ID, objc.Sel("state"))
	return rv
}


// The current position of the switch.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSwitch/state
func (s_ Switch) SetState(value objc.IObject /* cross-framework: ControlStateValue */) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setState:"), value)
}


// The default action-message selector associated with the control.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscontrol/action
func (s_ Switch) Action() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("action"))
	return rv
}


// The default action-message selector associated with the control.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscontrol/action
func (s_ Switch) SetAction(value unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setAction:"), value)
}


// The receiver’s cell object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscontrol/cell
func (s_ Switch) Cell() ICell {
	rv := objc.Send[Cell](s_.ID, objc.Sel("cell"))
	return rv
}


// The receiver’s cell object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscontrol/cell
func (s_ Switch) SetCell(value ICell) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setCell:"), value)
}


// A Boolean value indicating whether the receiver’s cell sends its action message continuously to its target during mouse tracking.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscontrol/iscontinuous
func (s_ Switch) IsContinuous() bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("isContinuous"))
	return rv
}


// A Boolean value indicating whether the receiver’s cell sends its action message continuously to its target during mouse tracking.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscontrol/iscontinuous
func (s_ Switch) SetIsContinuous(value bool) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setIsContinuous:"), value)
}



