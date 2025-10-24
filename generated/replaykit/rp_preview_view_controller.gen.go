// Code generated from Apple documentation for ReplayKit. DO NOT EDIT.

package replaykit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class RPPreviewViewController */


/* debug [class_header]: Header for RPPreviewViewController */
// The class instance for the [RPPreviewViewController] class.
var (
	RPPreviewViewControllerClass     _RPPreviewViewControllerClass
	RPPreviewViewControllerClassOnce sync.Once
)

func getRPPreviewViewControllerClass() _RPPreviewViewControllerClass {
	RPPreviewViewControllerClassOnce.Do(func() {
		RPPreviewViewControllerClass = _RPPreviewViewControllerClass{objc.GetClass("RPPreviewViewController")}
	})
	return RPPreviewViewControllerClass
}

type _RPPreviewViewControllerClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for RPPreviewViewController */
// An interface definition for the [RPPreviewViewController] class.
type IRPPreviewViewController interface {
	IViewController
	
/* debug [class_interface_properties]: Properties for RPPreviewViewController */
	// properties:
	PreviewControllerDelegate() unsafe.Pointer
	SetPreviewControllerDelegate(value unsafe.Pointer)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for RPPreviewViewController */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for RPPreviewViewController */
// Alloc allocates a new instance without initialization.
func (rc _RPPreviewViewControllerClass) Alloc() RPPreviewViewController {
	rv := objc.Send[RPPreviewViewController](objc.ID(rc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (rc _RPPreviewViewControllerClass) New() RPPreviewViewController {
	rv := objc.Send[RPPreviewViewController](objc.ID(rc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (r_ RPPreviewViewController) Init() RPPreviewViewController {
	rv := objc.Send[RPPreviewViewController](r_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (r_ RPPreviewViewController) Autorelease() RPPreviewViewController {
	rv := objc.Send[RPPreviewViewController](r_.ID, objc.Sel("autorelease"))
	return rv
}

// NewRPPreviewViewController creates a new RPPreviewViewController instance.
func NewRPPreviewViewController() RPPreviewViewController {
	return getRPPreviewViewControllerClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for RPPreviewViewController */
// An object that displays a user interface where users preview and edit a screen recording that you create with ReplayKit.
//
// Upon completion of a successful recording, the preview view controller is passed into the completion handler for .


// An object that displays a user interface where users preview and edit a screen recording that you create with ReplayKit.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ReplayKit/RPPreviewViewController
type RPPreviewViewController struct {
	ViewController
}

// RPPreviewViewControllerFrom constructs a [RPPreviewViewController] from an unsafe.Pointer.
//
// An object that displays a user interface where users preview and edit a screen recording that you create with ReplayKit.
func RPPreviewViewControllerFrom(ptr unsafe.Pointer) RPPreviewViewController {
	return RPPreviewViewController{
		ViewController: ViewControllerFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for RPPreviewViewController *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for RPPreviewViewController */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for RPPreviewViewController */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for RPPreviewViewController */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for RPPreviewViewController */

// The preview view controller’s delegate.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ReplayKit/RPPreviewViewController/previewControllerDelegate
func (r_ RPPreviewViewController) PreviewControllerDelegate() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](r_.ID, objc.Sel("previewControllerDelegate"))
	return rv
}/* debug [instance_properties/getter]: previewControllerDelegate */


// The preview view controller’s delegate.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ReplayKit/RPPreviewViewController/previewControllerDelegate
func (r_ RPPreviewViewController) SetPreviewControllerDelegate(value unsafe.Pointer) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setPreviewControllerDelegate:"), value)
}/* debug [instance_properties/setter]: previewControllerDelegate */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class RPPreviewViewController */


