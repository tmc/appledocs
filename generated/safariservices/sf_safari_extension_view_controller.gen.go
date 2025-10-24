// Code generated from Apple documentation for SafariServices. DO NOT EDIT.

package safariservices

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class SFSafariExtensionViewController */


/* debug [class_header]: Header for SFSafariExtensionViewController */
// The class instance for the [SFSafariExtensionViewController] class.
var (
	SFSafariExtensionViewControllerClass     _SFSafariExtensionViewControllerClass
	SFSafariExtensionViewControllerClassOnce sync.Once
)

func getSFSafariExtensionViewControllerClass() _SFSafariExtensionViewControllerClass {
	SFSafariExtensionViewControllerClassOnce.Do(func() {
		SFSafariExtensionViewControllerClass = _SFSafariExtensionViewControllerClass{objc.GetClass("SFSafariExtensionViewController")}
	})
	return SFSafariExtensionViewControllerClass
}

type _SFSafariExtensionViewControllerClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for SFSafariExtensionViewController */
// An interface definition for the [SFSafariExtensionViewController] class.
type ISFSafariExtensionViewController interface {
	IViewController
	
/* debug [class_interface_properties]: Properties for SFSafariExtensionViewController */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for SFSafariExtensionViewController */
	// methods:
	DismissPopover()
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for SFSafariExtensionViewController */
// Alloc allocates a new instance without initialization.
func (sc _SFSafariExtensionViewControllerClass) Alloc() SFSafariExtensionViewController {
	rv := objc.Send[SFSafariExtensionViewController](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (sc _SFSafariExtensionViewControllerClass) New() SFSafariExtensionViewController {
	rv := objc.Send[SFSafariExtensionViewController](objc.ID(sc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (s_ SFSafariExtensionViewController) Init() SFSafariExtensionViewController {
	rv := objc.Send[SFSafariExtensionViewController](s_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s_ SFSafariExtensionViewController) Autorelease() SFSafariExtensionViewController {
	rv := objc.Send[SFSafariExtensionViewController](s_.ID, objc.Sel("autorelease"))
	return rv
}

// NewSFSafariExtensionViewController creates a new SFSafariExtensionViewController instance.
func NewSFSafariExtensionViewController() SFSafariExtensionViewController {
	return getSFSafariExtensionViewControllerClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for SFSafariExtensionViewController */
// The view controller for a popover associated with your app extension.
//
// If your toolbar item has a popover, your popover view controller should be a subclass of this class. As with other macOS development, typically you want to add your own outlets and actions to the view controller, and provide an XIB file for its user interface. Your view controller’s contents must use Auto Layout.


// The view controller for a popover associated with your app extension.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SafariServices/SFSafariExtensionViewController
type SFSafariExtensionViewController struct {
	ViewController
}

// SFSafariExtensionViewControllerFrom constructs a [SFSafariExtensionViewController] from an unsafe.Pointer.
//
// The view controller for a popover associated with your app extension.
func SFSafariExtensionViewControllerFrom(ptr unsafe.Pointer) SFSafariExtensionViewController {
	return SFSafariExtensionViewController{
		ViewController: ViewControllerFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for SFSafariExtensionViewController *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for SFSafariExtensionViewController */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for SFSafariExtensionViewController */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for SFSafariExtensionViewController */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SafariServices/SFSafariExtensionViewController/dismissPopover()
func (s_ SFSafariExtensionViewController) DismissPopover() {
	objc.Send[objc.ID](s_.ID, objc.Sel("dismissPopover"))
}/* debug [instance_methods/method]: DismissPopover */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for SFSafariExtensionViewController */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class SFSafariExtensionViewController */



