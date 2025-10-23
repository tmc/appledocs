// Code generated from Apple documentation for CoreAudioKit. DO NOT EDIT.

package coreaudiokit

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
		ViewControllerClass = _ViewControllerClass{objc.GetClass("AUViewController")}
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

// The base class to extend when creating a custom user interface for an audio unit.
//
// This class doesn’t add any new methods or properties to its superclass, but it does conform to the protocol. A host app can access the view controller by calling the method on the corresponding object.


// The base class to extend when creating a custom user interface for an audio unit.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreAudioKit/AUViewController
type ViewController struct {
	objectivec.Object
}

// ViewControllerFrom constructs a [ViewController] from an unsafe.Pointer.
//
// The base class to extend when creating a custom user interface for an audio unit.
func ViewControllerFrom(ptr unsafe.Pointer) ViewController {
	return ViewController{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (vc _ViewControllerClass) Alloc() ViewController {
	rv := objc.Send[ViewController](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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




