// Code generated from Apple documentation for ReplayKit. DO NOT EDIT.

package replaykit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

/* debug [class.gen.go]: Generating class RPBroadcastActivityViewController */


/* debug [class_header]: Header for RPBroadcastActivityViewController */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for RPBroadcastActivityViewController */
// An interface definition for the [RPBroadcastActivityViewController] class.
type IRPBroadcastActivityViewController interface {
	IViewController
	
/* debug [class_interface_properties]: Properties for RPBroadcastActivityViewController */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for RPBroadcastActivityViewController */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for RPBroadcastActivityViewController */
// Alloc allocates a new instance without initialization.
func (rc _RPBroadcastActivityViewControllerClass) Alloc() RPBroadcastActivityViewController {
	rv := objc.Send[RPBroadcastActivityViewController](objc.ID(rc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for RPBroadcastActivityViewController */
// A view controller that displays a user interface where users choose a broadcast service.
//
// The view controller displays the broadcast services currently installed on the device. On iPad, you must present the broadcast activity view controller as a popover.


// A view controller that displays a user interface where users choose a broadcast service.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ReplayKit/RPBroadcastActivityViewController
type RPBroadcastActivityViewController struct {
	ViewController
}

// RPBroadcastActivityViewControllerFrom constructs a [RPBroadcastActivityViewController] from an unsafe.Pointer.
//
// A view controller that displays a user interface where users choose a broadcast service.
func RPBroadcastActivityViewControllerFrom(ptr unsafe.Pointer) RPBroadcastActivityViewController {
	return RPBroadcastActivityViewController{
		ViewController: ViewControllerFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for RPBroadcastActivityViewController *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for RPBroadcastActivityViewController */

// Loads a broadcast activity view controller.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ReplayKit/RPBroadcastActivityViewController/load(handler:)
func (rc _RPBroadcastActivityViewControllerClass) LoadBroadcastActivityViewControllerWithHandler(handler unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(rc.class), objc.Sel("loadBroadcastActivityViewControllerWithHandler:"), handler)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=LoadBroadcastActivityViewControllerWithHandler) */


// Loads a broadcast activity view controller with a preferred extension.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ReplayKit/RPBroadcastActivityViewController/load(withPreferredExtension:handler:)
func (rc _RPBroadcastActivityViewControllerClass) LoadBroadcastActivityViewControllerWithPreferredExtensionHandler(preferredExtension objc.IObject /* cross-framework: NSString */, handler unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(rc.class), objc.Sel("loadBroadcastActivityViewControllerWithPreferredExtension:handler:"), preferredExtension, handler)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=LoadBroadcastActivityViewControllerWithPreferredExtensionHandler) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for RPBroadcastActivityViewController */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for RPBroadcastActivityViewController */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for RPBroadcastActivityViewController */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class RPBroadcastActivityViewController */


