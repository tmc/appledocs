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

// A bindings-compatible controller that manages a tree of objects.
//
// The class provides selection and sort management. Its primary purpose is to act as the controller when binding and instances to a hierarchical collection of objects. The root content object of the tree can be a single object, or an array of objects. An object requires that you describe how the tree of objects is traversed by specifying the key-path for child objects specified by . All child objects for the tree must be key-value coding compliant for the same child key path. If necessary, you should add properties to your model classes that map the child key name to the appropriate class-specific property name. Child objects can implement a count method (specified to the tree controller using ) that, if provided, returns the number of child objects available. Your model objects are expected to update the value of the count key path in a key-value observing compliant method. Optionally, you can also provide a leaf key path using that specifies a key in your model object that returns if the object is a leaf node, and if it is not. Changes to the leaf node value of the child object should be made in a key-value observing compliant manner. Providing the leaf node key path can improve performance, because it prevents the from having to examine the child object to determine if it is a leaf node. For more information about using NSTreeController in your app, see .
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

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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




