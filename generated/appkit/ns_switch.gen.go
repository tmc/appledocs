// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [Switch_] class.
var (
	Switch_Class     _Switch_Class
	Switch_ClassOnce sync.Once
)

func getSwitch_Class() _Switch_Class {
	Switch_ClassOnce.Do(func() {
		Switch_Class = _Switch_Class{objc.GetClass("NSSwitch")}
	})
	return Switch_Class
}

type _Switch_Class struct {
	class objc.Class
}

// An interface definition for the [Switch_] class.
type ISwitch_ interface {
	IControl
}

// A control that offers a binary choice.
//
// The class provides a simple interface for displaying and toggling a Boolean state, such as on/off. A switch toggles its and sends its when clicked, activated through the keyboard, or tapped in the Touch Bar. also allows dragging between states, and if is , the switch sends its for each change in position during the drag. doesn’t use an instance of to provide its functionality. The class property and instance property both return , and they ignore attempts to set a non- value. For design guidance, see Human Interface Guidelines > .
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSwitch
type Switch_ struct {
	Control
}

// Switch_From constructs a [Switch_] from an unsafe.Pointer.
//
// A control that offers a binary choice.
func Switch_From(ptr unsafe.Pointer) Switch_ {
	return Switch_{
		Control: ControlFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (sc _Switch_Class) Alloc() Switch_ {
	rv := objc.Send[Switch_](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (sc _Switch_Class) New() Switch_ {
	rv := objc.Send[Switch_](objc.ID(sc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (s_ Switch_) Init() Switch_ {
	rv := objc.Send[Switch_](s_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s_ Switch_) Autorelease() Switch_ {
	rv := objc.Send[Switch_](s_.ID, objc.Sel("autorelease"))
	return rv
}

// NewSwitch_ creates a new Switch_ instance.
func NewSwitch_() Switch_ {
	return getSwitch_Class().New()
}
