// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [ArrayController] class.
var (
	ArrayControllerClass     _ArrayControllerClass
	ArrayControllerClassOnce sync.Once
)

func getArrayControllerClass() _ArrayControllerClass {
	ArrayControllerClassOnce.Do(func() {
		ArrayControllerClass = _ArrayControllerClass{objc.GetClass("NSArrayController")}
	})
	return ArrayControllerClass
}

type _ArrayControllerClass struct {
	class objc.Class
}

// An interface definition for the [ArrayController] class.
type IArrayController interface {
	IObjectController
}

// A bindings-compatible controller that manages a collection of objects.
//
// Typically the collection that an manages is an array, however, if the controller manages a relationship of a managed object (see ) the collection may be a set. provides selection management and sorting capabilities.
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

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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
	return getArrayControllerClass().New()
}

// An array of key paths that trigger automatic content sorting or filtering
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSArrayController/automaticRearrangementKeyPaths
func (a_ ArrayController) AutomaticRearrangementKeyPaths() []string {
	rv := objc.Send[[]string](a_.ID, objc.Sel("automaticRearrangementKeyPaths"))
	return rv
}
