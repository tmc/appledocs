// Code generated from Apple documentation for GameController. DO NOT EDIT.

package gamecontroller

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/appkit"
)

// The class instance for the [GCEventViewController] class.
var (
	GCEventViewControllerClass     _GCEventViewControllerClass
	GCEventViewControllerClassOnce sync.Once
)

func getGCEventViewControllerClass() _GCEventViewControllerClass {
	GCEventViewControllerClassOnce.Do(func() {
		GCEventViewControllerClass = _GCEventViewControllerClass{objc.GetClass("GCEventViewController")}
	})
	return GCEventViewControllerClass
}

type _GCEventViewControllerClass struct {
	class objc.Class
}

// An interface definition for the [GCEventViewController] class.
type IGCEventViewController interface {
	appkit.IViewController
	// properties:
	ControllerUserInteractionEnabled() bool /* primitive/slice/pointer. */
	SetControllerUserInteractionEnabled(value bool /* primitive/slice/pointer. */)
	// methods:
}

// A view controller that delivers input either from the responder chain to views, or from game controllers to profiles.
//
// On systems, such as tvOS, where the player uses the game controller to both navigate the system interface and play your game, use a object as the root view controller to selectively receive input directly from the game controller. You can’t simultaneously process input through the responder chain and Game Controller input elements. By default the system delivers input events to your app using the responder chain. To get the input values through the game controller objects, set a object as the root view controller. The view controller delivers the input for its views and their subviews to the game controller’s profile. To switch back to the responder chain, set the view controller’s property to .


// A view controller that delivers input either from the responder chain to views, or from game controllers to profiles.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameController/GCEventViewController
type GCEventViewController struct {
	appkit.ViewController
}

// GCEventViewControllerFrom constructs a [GCEventViewController] from an unsafe.Pointer.
//
// A view controller that delivers input either from the responder chain to views, or from game controllers to profiles.
func GCEventViewControllerFrom(ptr unsafe.Pointer) GCEventViewController {
	return GCEventViewController{
		ViewController: appkit.ViewControllerFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (gc _GCEventViewControllerClass) Alloc() GCEventViewController {
	rv := objc.Send[GCEventViewController](objc.ID(gc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (gc _GCEventViewControllerClass) New() GCEventViewController {
	rv := objc.Send[GCEventViewController](objc.ID(gc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (g_ GCEventViewController) Init() GCEventViewController {
	rv := objc.Send[GCEventViewController](g_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (g_ GCEventViewController) Autorelease() GCEventViewController {
	rv := objc.Send[GCEventViewController](g_.ID, objc.Sel("autorelease"))
	return rv
}

// NewGCEventViewController creates a new GCEventViewController instance.
func NewGCEventViewController() GCEventViewController {
	return getGCEventViewControllerClass().New()
}



// A Boolean value that indicates whether the system delivers game controller input to profile objects or to views using the responder chain.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameController/GCEventViewController/controllerUserInteractionEnabled
func (g_ GCEventViewController) ControllerUserInteractionEnabled() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](g_.ID, objc.Sel("controllerUserInteractionEnabled"))
	return rv
}


// A Boolean value that indicates whether the system delivers game controller input to profile objects or to views using the responder chain.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameController/GCEventViewController/controllerUserInteractionEnabled
func (g_ GCEventViewController) SetControllerUserInteractionEnabled(value bool /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setControllerUserInteractionEnabled:"), value)
}



