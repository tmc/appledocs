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
}

// A control that offers a binary choice.
//
// The class provides a simple interface for displaying and toggling a Boolean state, such as on/off. A switch toggles its and sends its when clicked, activated through the keyboard, or tapped in the Touch Bar. also allows dragging between states, and if is , the switch sends its for each change in position during the drag. doesn’t use an instance of to provide its functionality. The class property and instance property both return , and they ignore attempts to set a non- value. For design guidance, see Human Interface Guidelines > .
//
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


// The default action-message selector associated with the control.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscontrol/action
func (s_ Switch) Action() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("action"))
	return rv
}


// SetAction sets the value of the action property.
// The default action-message selector associated with the control.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscontrol/action
func (s_ Switch) SetAction(value unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setAction:"), value)
}

// The receiver’s cell object.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscontrol/cell
func (s_ Switch) Cell() NSCell {
	rv := objc.Send[NSCell](s_.ID, objc.Sel("cell"))
	return rv
}


// SetCell sets the value of the cell property.
// The receiver’s cell object.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscontrol/cell
func (s_ Switch) SetCell(value ICell) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setCell:"), value)
}

// A Boolean value indicating whether the receiver’s cell sends its action message continuously to its target during mouse tracking.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscontrol/iscontinuous
func (s_ Switch) IsContinuous() bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("isContinuous"))
	return rv
}


// SetIsContinuous sets the value of the isContinuous property.
// A Boolean value indicating whether the receiver’s cell sends its action message continuously to its target during mouse tracking.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscontrol/iscontinuous
func (s_ Switch) SetIsContinuous(value bool) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setIsContinuous:"), value)
}

// The current position of the switch.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsswitch/state
func (s_ Switch) State() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("state"))
	return rv
}


// SetState sets the value of the state property.
// The current position of the switch.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsswitch/state
func (s_ Switch) SetState(value unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setState:"), value)
}



