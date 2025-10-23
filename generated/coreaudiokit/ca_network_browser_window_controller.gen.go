// Code generated from Apple documentation for CoreAudioKit. DO NOT EDIT.

package coreaudiokit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/appkit"
)

// The class instance for the [NetworkBrowserWindowController] class.
var (
	NetworkBrowserWindowControllerClass     _NetworkBrowserWindowControllerClass
	NetworkBrowserWindowControllerClassOnce sync.Once
)

func getNetworkBrowserWindowControllerClass() _NetworkBrowserWindowControllerClass {
	NetworkBrowserWindowControllerClassOnce.Do(func() {
		NetworkBrowserWindowControllerClass = _NetworkBrowserWindowControllerClass{objc.GetClass("CANetworkBrowserWindowController")}
	})
	return NetworkBrowserWindowControllerClass
}

type _NetworkBrowserWindowControllerClass struct {
	class objc.Class
}

// An interface definition for the [NetworkBrowserWindowController] class.
type INetworkBrowserWindowController interface {
	appkit.IWindowController
	// properties:
	// methods:
}

// A window controller that displays available network audio devices.


// A window controller that displays available network audio devices.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreAudioKit/CANetworkBrowserWindowController
type NetworkBrowserWindowController struct {
	appkit.WindowController
}

// NetworkBrowserWindowControllerFrom constructs a [NetworkBrowserWindowController] from an unsafe.Pointer.
//
// A window controller that displays available network audio devices.
func NetworkBrowserWindowControllerFrom(ptr unsafe.Pointer) NetworkBrowserWindowController {
	return NetworkBrowserWindowController{
		WindowController: appkit.WindowControllerFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (nc _NetworkBrowserWindowControllerClass) Alloc() NetworkBrowserWindowController {
	rv := objc.Send[NetworkBrowserWindowController](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (nc _NetworkBrowserWindowControllerClass) New() NetworkBrowserWindowController {
	rv := objc.Send[NetworkBrowserWindowController](objc.ID(nc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (n_ NetworkBrowserWindowController) Init() NetworkBrowserWindowController {
	rv := objc.Send[NetworkBrowserWindowController](n_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (n_ NetworkBrowserWindowController) Autorelease() NetworkBrowserWindowController {
	rv := objc.Send[NetworkBrowserWindowController](n_.ID, objc.Sel("autorelease"))
	return rv
}

// NewNetworkBrowserWindowController creates a new NetworkBrowserWindowController instance.
func NewNetworkBrowserWindowController() NetworkBrowserWindowController {
	return getNetworkBrowserWindowControllerClass().New()
}





