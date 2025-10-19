// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [ArrayController] class.
var arrayControllerClass = _ArrayControllerClass{objc.GetClass("NSArrayController")}

type _ArrayControllerClass struct {
	class objc.Class
}

// An interface definition for the [ArrayController] class.
type IArrayController interface {
	IObjectController
}

// A bindings-compatible controller that manages a collection of objects. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSArrayController

type ArrayController struct {
	ObjectController
}

// ArrayControllerFrom constructs a [ArrayController] from an unsafe.Pointer.
//
// A bindings-compatible controller that manages a collection of objects.
func ArrayControllerFrom(ptr unsafe.Pointer) ArrayController {
	return ArrayController{
		ObjectController: ObjectControllerFrom(ptr),
	}
}
// Alloc allocates a new instance without initialization.
func (ac _ArrayControllerClass) Alloc() ArrayController {
	rv := objc.Send[ArrayController](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new instance with a +1 retain count.
func (ac _ArrayControllerClass) New() ArrayController {
	rv := objc.Send[ArrayController](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ ArrayController) Init() ArrayController {
	rv := objc.Send[ArrayController](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ ArrayController) Autorelease() ArrayController {
	rv := objc.Send[ArrayController](a_.ID, objc.Sel("autorelease"))
	return rv
}

// NewArrayController creates a new ArrayController instance.
func NewArrayController() ArrayController {
	return arrayControllerClass.New()
}




