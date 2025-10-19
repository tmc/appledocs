// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [Switch_] class.
var (
	switch_Class     _Switch_Class
	switch_ClassOnce sync.Once
)

func getSwitch_Class() _Switch_Class {
	switch_ClassOnce.Do(func() {
		switch_Class = _Switch_Class{objc.GetClass("NSSwitch")}
	})
	return switch_Class
}

type _Switch_Class struct {
	class objc.Class
}

// An interface definition for the [Switch_] class.
type ISwitch_ interface {
	IControl
}

// A control that offers a binary choice. [Full Topic]
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




