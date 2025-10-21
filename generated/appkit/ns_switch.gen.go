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
