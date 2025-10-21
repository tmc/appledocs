// Code generated from Apple documentation for ReplayKit. DO NOT EDIT.

package replaykit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/appkit"
)

// The class instance for the [RPBroadcastActivityViewController] class.
var (
	RPBroadcastActivityViewControllerClass     _RPBroadcastActivityViewControllerClass
	RPBroadcastActivityViewControllerClassOnce sync.Once
)

func getRPBroadcastActivityViewControllerClass() _RPBroadcastActivityViewControllerClass {
	RPBroadcastActivityViewControllerClassOnce.Do(func() {
		RPBroadcastActivityViewControllerClass = _RPBroadcastActivityViewControllerClass{objc.GetClass("RPBroadcastActivityViewController")}
	})
	return RPBroadcastActivityViewControllerClass
}

type _RPBroadcastActivityViewControllerClass struct {
	class objc.Class
}

// An interface definition for the [RPBroadcastActivityViewController] class.
type IRPBroadcastActivityViewController interface {
	appkit.IViewController
}

// A view controller that displays a user interface where users choose a broadcast service.
//
// The view controller displays the broadcast services currently installed on the device. On iPad, you must present the broadcast activity view controller as a popover.
//
// [Full Topic]: https://developer.apple.com/documentation/ReplayKit/RPBroadcastActivityViewController
type RPBroadcastActivityViewController struct {
	appkit.ViewController
}

// RPBroadcastActivityViewControllerFrom constructs a [RPBroadcastActivityViewController] from an unsafe.Pointer.
//
// A view controller that displays a user interface where users choose a broadcast service.
func RPBroadcastActivityViewControllerFrom(ptr unsafe.Pointer) RPBroadcastActivityViewController {
	return RPBroadcastActivityViewController{
		ViewController: appkit.ViewControllerFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (rc _RPBroadcastActivityViewControllerClass) Alloc() RPBroadcastActivityViewController {
	rv := objc.Send[RPBroadcastActivityViewController](objc.ID(rc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (rc _RPBroadcastActivityViewControllerClass) New() RPBroadcastActivityViewController {
	rv := objc.Send[RPBroadcastActivityViewController](objc.ID(rc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (r_ RPBroadcastActivityViewController) Init() RPBroadcastActivityViewController {
	rv := objc.Send[RPBroadcastActivityViewController](r_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (r_ RPBroadcastActivityViewController) Autorelease() RPBroadcastActivityViewController {
	rv := objc.Send[RPBroadcastActivityViewController](r_.ID, objc.Sel("autorelease"))
	return rv
}

// NewRPBroadcastActivityViewController creates a new RPBroadcastActivityViewController instance.
func NewRPBroadcastActivityViewController() RPBroadcastActivityViewController {
	return getRPBroadcastActivityViewControllerClass().New()
}


// Loads a broadcast activity view controller.
//
// [Full Topic]: https://developer.apple.com/documentation/ReplayKit/RPBroadcastActivityViewController/load(handler:)
func (rc _RPBroadcastActivityViewControllerClass) LoadBroadcastActivityViewControllerWithHandler(handler unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(rc.class), objc.Sel("loadBroadcastActivityViewControllerWithHandler:"), handler)
}

// Loads a broadcast activity view controller with a preferred extension.
//
// [Full Topic]: https://developer.apple.com/documentation/ReplayKit/RPBroadcastActivityViewController/load(withPreferredExtension:handler:)
func (rc _RPBroadcastActivityViewControllerClass) LoadBroadcastActivityViewControllerWithPreferredExtensionHandler(preferredExtension string, handler unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(rc.class), objc.Sel("loadBroadcastActivityViewControllerWithPreferredExtension:handler:"), objc.String(preferredExtension), handler)
}

// The delegate for the broadcast activity view controller.
//
// [Full Topic]: https://developer.apple.com/documentation/ReplayKit/RPBroadcastActivityViewController/delegate
func (r_ RPBroadcastActivityViewController) Delegate() objc.ID {
	rv := objc.Send[objc.ID](r_.ID, objc.Sel("delegate"))
	return rv
}


// SetDelegate sets the value of the delegate property.
// The delegate for the broadcast activity view controller.

//
// [Full Topic]: https://developer.apple.com/documentation/ReplayKit/RPBroadcastActivityViewController/delegate
func (r_ RPBroadcastActivityViewController) SetDelegate(value objc.ID) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setDelegate:"), value)
}


