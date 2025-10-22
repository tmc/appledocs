// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [Controller] class.
var (
	ControllerClass     _ControllerClass
	ControllerClassOnce sync.Once
)

func getControllerClass() _ControllerClass {
	ControllerClassOnce.Do(func() {
		ControllerClass = _ControllerClass{objc.GetClass("NSController")}
	})
	return ControllerClass
}

type _ControllerClass struct {
	class objc.Class
}

// An interface definition for the [Controller] class.
type IController interface {
	objectivec.IObject
	IsEditing() bool
	SetIsEditing(value bool)
}

// An abstract class that implements the and informal protocols required for controller classes.


// An abstract class that implements the and informal protocols required for controller classes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSController

type Controller struct {
	objectivec.Object
}

// ControllerFrom constructs a [Controller] from an unsafe.Pointer.
//
// An abstract class that implements the and informal protocols required for controller classes.
func ControllerFrom(ptr unsafe.Pointer) Controller {
	return Controller{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (cc _ControllerClass) Alloc() Controller {
	rv := objc.Send[Controller](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (cc _ControllerClass) New() Controller {
	rv := objc.Send[Controller](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ Controller) Init() Controller {
	rv := objc.Send[Controller](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ Controller) Autorelease() Controller {
	rv := objc.Send[Controller](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewController creates a new Controller instance.
func NewController() Controller {
	return getControllerClass().New()
}




// A Boolean value indicating if any editors are registered with the controller.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscontroller/isediting

func (c_ Controller) IsEditing() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("isEditing"))
	return rv
}


// A Boolean value indicating if any editors are registered with the controller.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscontroller/isediting

func (c_ Controller) SetIsEditing(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIsEditing:"), value)
}


