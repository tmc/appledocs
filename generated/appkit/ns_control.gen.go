// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [Control] class.
var (
	ControlClass     _ControlClass
	ControlClassOnce sync.Once
)

func getControlClass() _ControlClass {
	ControlClassOnce.Do(func() {
		ControlClass = _ControlClass{objc.GetClass("NSControl")}
	})
	return ControlClass
}

type _ControlClass struct {
	class objc.Class
}

// An interface definition for the [Control] class.
type IControl interface {
	objectivec.IObject
	// properties:
	// methods:
}

// A parent class referenced by other AppKit classes.


// A parent class referenced by other AppKit classes. [Full Topic]
type Control struct {
	objectivec.Object
}

// ControlFrom constructs a [Control] from an unsafe.Pointer.
//
// A parent class referenced by other AppKit classes.
func ControlFrom(ptr unsafe.Pointer) Control {
	return Control{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (cc _ControlClass) Alloc() Control {
	rv := objc.Send[Control](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (cc _ControlClass) New() Control {
	rv := objc.Send[Control](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ Control) Init() Control {
	rv := objc.Send[Control](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ Control) Autorelease() Control {
	rv := objc.Send[Control](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewControl creates a new Control instance.
func NewControl() Control {
	return getControlClass().New()
}




