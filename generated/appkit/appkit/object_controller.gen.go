// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [ObjectController] class.
var (
	objectControllerClass     _ObjectControllerClass
	objectControllerClassOnce sync.Once
)

func getObjectControllerClass() _ObjectControllerClass {
	objectControllerClassOnce.Do(func() {
		objectControllerClass = _ObjectControllerClass{objc.GetClass("NSObjectController")}
	})
	return objectControllerClass
}

type _ObjectControllerClass struct {
	class objc.Class
}

// An interface definition for the [ObjectController] class.
type IObjectController interface {
	IController
}

// A controller that can manage an object’s properties referenced by key-value paths. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSObjectController

type ObjectController struct {
	Controller
}

// ObjectControllerFrom constructs a [ObjectController] from an unsafe.Pointer.
//
// A controller that can manage an object’s properties referenced by key-value paths.
func ObjectControllerFrom(ptr unsafe.Pointer) ObjectController {
	return ObjectController{
		Controller: ControllerFrom(ptr),
	}
}
// Alloc allocates a new instance without initialization.
func (oc _ObjectControllerClass) Alloc() ObjectController {
	rv := objc.Send[ObjectController](objc.ID(oc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new instance with a +1 retain count.
func (oc _ObjectControllerClass) New() ObjectController {
	rv := objc.Send[ObjectController](objc.ID(oc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (o_ ObjectController) Init() ObjectController {
	rv := objc.Send[ObjectController](o_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (o_ ObjectController) Autorelease() ObjectController {
	rv := objc.Send[ObjectController](o_.ID, objc.Sel("autorelease"))
	return rv
}

// NewObjectController creates a new ObjectController instance.
func NewObjectController() ObjectController {
	return getObjectControllerClass().New()
}




