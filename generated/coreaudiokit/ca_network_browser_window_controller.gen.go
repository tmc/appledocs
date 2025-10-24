// Code generated from Apple documentation for CoreAudioKit. DO NOT EDIT.

package coreaudiokit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/appkit"
)

/* debug [class.gen.go]: Generating class CANetworkBrowserWindowController */


/* debug [class_header]: Header for CANetworkBrowserWindowController */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for NetworkBrowserWindowController */
// An interface definition for the [NetworkBrowserWindowController] class.
type INetworkBrowserWindowController interface {
	appkit.IWindowController
	
/* debug [class_interface_properties]: Properties for NetworkBrowserWindowController */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for NetworkBrowserWindowController */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for NetworkBrowserWindowController */
// Alloc allocates a new instance without initialization.
func (nc _NetworkBrowserWindowControllerClass) Alloc() NetworkBrowserWindowController {
	rv := objc.Send[NetworkBrowserWindowController](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for NetworkBrowserWindowController */
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for NetworkBrowserWindowController */
/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for NetworkBrowserWindowController */

// Returns a Boolean value that indicates whether the current machine hardware supports Audio Video Bridging (AVB).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreAudioKit/CANetworkBrowserWindowController/isAVBSupported()
func (nc _NetworkBrowserWindowControllerClass) IsAVBSupported() bool {
	rv := objc.Send[bool](objc.ID(nc.class), objc.Sel("isAVBSupported"))
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=IsAVBSupported) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for NetworkBrowserWindowController */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for NetworkBrowserWindowController */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for NetworkBrowserWindowController */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class CANetworkBrowserWindowController */


