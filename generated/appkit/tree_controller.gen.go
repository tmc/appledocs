// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [TreeController] class.
var (
	treeControllerClass     _TreeControllerClass
	treeControllerClassOnce sync.Once
)

func getTreeControllerClass() _TreeControllerClass {
	treeControllerClassOnce.Do(func() {
		treeControllerClass = _TreeControllerClass{objc.GetClass("NSTreeController")}
	})
	return treeControllerClass
}

type _TreeControllerClass struct {
	class objc.Class
}

// An interface definition for the [TreeController] class.
type ITreeController interface {
	IObjectController
}

// A bindings-compatible controller that manages a tree of objects. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTreeController

type TreeController struct {
	ObjectController
}

// TreeControllerFrom constructs a [TreeController] from an unsafe.Pointer.
//
// A bindings-compatible controller that manages a tree of objects.
func TreeControllerFrom(ptr unsafe.Pointer) TreeController {
	return TreeController{
		ObjectController: ObjectControllerFrom(ptr),
	}
}
// Alloc allocates a new instance without initialization.
func (tc _TreeControllerClass) Alloc() TreeController {
	rv := objc.Send[TreeController](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new instance with a +1 retain count.
func (tc _TreeControllerClass) New() TreeController {
	rv := objc.Send[TreeController](objc.ID(tc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (t_ TreeController) Init() TreeController {
	rv := objc.Send[TreeController](t_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t_ TreeController) Autorelease() TreeController {
	rv := objc.Send[TreeController](t_.ID, objc.Sel("autorelease"))
	return rv
}

// NewTreeController creates a new TreeController instance.
func NewTreeController() TreeController {
	return getTreeControllerClass().New()
}




