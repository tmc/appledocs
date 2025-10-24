// Code generated from Apple documentation for MailKit. DO NOT EDIT.

package mailkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/appkit"
)

/* debug [class.gen.go]: Generating class MEExtensionViewController */


/* debug [class_header]: Header for MEExtensionViewController */
// The class instance for the [MEExtensionViewController] class.
var (
	MEExtensionViewControllerClass     _MEExtensionViewControllerClass
	MEExtensionViewControllerClassOnce sync.Once
)

func getMEExtensionViewControllerClass() _MEExtensionViewControllerClass {
	MEExtensionViewControllerClassOnce.Do(func() {
		MEExtensionViewControllerClass = _MEExtensionViewControllerClass{objc.GetClass("MEExtensionViewController")}
	})
	return MEExtensionViewControllerClass
}

type _MEExtensionViewControllerClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MEExtensionViewController */
// An interface definition for the [MEExtensionViewController] class.
type IMEExtensionViewController interface {
	appkit.IViewController
	
/* debug [class_interface_properties]: Properties for MEExtensionViewController */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MEExtensionViewController */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MEExtensionViewController */
// Alloc allocates a new instance without initialization.
func (mc _MEExtensionViewControllerClass) Alloc() MEExtensionViewController {
	rv := objc.Send[MEExtensionViewController](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MEExtensionViewControllerClass) New() MEExtensionViewController {
	rv := objc.Send[MEExtensionViewController](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MEExtensionViewController) Init() MEExtensionViewController {
	rv := objc.Send[MEExtensionViewController](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MEExtensionViewController) Autorelease() MEExtensionViewController {
	rv := objc.Send[MEExtensionViewController](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMEExtensionViewController creates a new MEExtensionViewController instance.
func NewMEExtensionViewController() MEExtensionViewController {
	return getMEExtensionViewControllerClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MEExtensionViewController */
// An object that manages a view for compose session and message security handlers.
//
// Create a custom subclass of to provide MailKit with a view that displays: Additional configuration options in a Mail compose window Details about the user who signed a digitally signed email message


// An object that manages a view for compose session and message security handlers.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MailKit/MEExtensionViewController
type MEExtensionViewController struct {
	appkit.ViewController
}

// MEExtensionViewControllerFrom constructs a [MEExtensionViewController] from an unsafe.Pointer.
//
// An object that manages a view for compose session and message security handlers.
func MEExtensionViewControllerFrom(ptr unsafe.Pointer) MEExtensionViewController {
	return MEExtensionViewController{
		ViewController: appkit.ViewControllerFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MEExtensionViewController *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MEExtensionViewController */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MEExtensionViewController */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MEExtensionViewController */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MEExtensionViewController */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MEExtensionViewController */



