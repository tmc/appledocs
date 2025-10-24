// Code generated from Apple documentation for VisionKit. DO NOT EDIT.

package visionkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)





// The class instance for the [ViewController] class.
var (
	ViewControllerClass     _ViewControllerClass
	ViewControllerClassOnce sync.Once
)

func getViewControllerClass() _ViewControllerClass {
	ViewControllerClassOnce.Do(func() {
		ViewControllerClass = _ViewControllerClass{objc.GetClass("UIViewController")}
	})
	return ViewControllerClass
}

type _ViewControllerClass struct {
	class objc.Class
}





// An interface definition for the [ViewController] class.
type IViewController interface {
	objectivec.IObject
	

	// properties:


	

	// methods:


}





// Alloc allocates a new instance without initialization.
func (vc _ViewControllerClass) Alloc() ViewController {
	rv := objc.Send[ViewController](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (vc _ViewControllerClass) New() ViewController {
	rv := objc.Send[ViewController](objc.ID(vc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (v_ ViewController) Init() ViewController {
	rv := objc.Send[ViewController](v_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v_ ViewController) Autorelease() ViewController {
	rv := objc.Send[ViewController](v_.ID, objc.Sel("autorelease"))
	return rv
}

// NewViewController creates a new ViewController instance.
func NewViewController() ViewController {
	return getViewControllerClass().New()
}





// A parent class referenced by other VisionKit classes.


// A parent class referenced by other VisionKit classes. [Full Topic]
type ViewController struct {
	objectivec.Object
}

// ViewControllerFrom constructs a [ViewController] from an unsafe.Pointer.
//
// A parent class referenced by other VisionKit classes.
func ViewControllerFrom(ptr unsafe.Pointer) ViewController {
	return ViewController{objectivec.Object{objc.ID(ptr)}}
}































