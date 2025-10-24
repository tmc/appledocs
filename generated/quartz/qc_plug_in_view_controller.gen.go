// Code generated from Apple documentation for Quartz. DO NOT EDIT.

package quartz

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/appkit"
	"github.com/tmc/appledocs/generated/foundation"
)

/* debug [class.gen.go]: Generating class QCPlugInViewController */


/* debug [class_header]: Header for QCPlugInViewController */
// The class instance for the [QCPlugInViewController] class.
var (
	QCPlugInViewControllerClass     _QCPlugInViewControllerClass
	QCPlugInViewControllerClassOnce sync.Once
)

func getQCPlugInViewControllerClass() _QCPlugInViewControllerClass {
	QCPlugInViewControllerClassOnce.Do(func() {
		QCPlugInViewControllerClass = _QCPlugInViewControllerClass{objc.GetClass("QCPlugInViewController")}
	})
	return QCPlugInViewControllerClass
}

type _QCPlugInViewControllerClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for QCPlugInViewController */
// An interface definition for the [QCPlugInViewController] class.
type IQCPlugInViewController interface {
	appkit.IViewController
	
/* debug [class_interface_properties]: Properties for QCPlugInViewController */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for QCPlugInViewController */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for QCPlugInViewController */
// Alloc allocates a new instance without initialization.
func (qc _QCPlugInViewControllerClass) Alloc() QCPlugInViewController {
	rv := objc.Send[QCPlugInViewController](objc.ID(qc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (qc _QCPlugInViewControllerClass) New() QCPlugInViewController {
	rv := objc.Send[QCPlugInViewController](objc.ID(qc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (q_ QCPlugInViewController) Init() QCPlugInViewController {
	rv := objc.Send[QCPlugInViewController](q_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (q_ QCPlugInViewController) Autorelease() QCPlugInViewController {
	rv := objc.Send[QCPlugInViewController](q_.ID, objc.Sel("autorelease"))
	return rv
}

// NewQCPlugInViewController creates a new QCPlugInViewController instance.
func NewQCPlugInViewController() QCPlugInViewController {
	return getQCPlugInViewControllerClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for QCPlugInViewController */
// The class communicates (through Cocoa bindings) between a custom patch and the view used for the internal settings of the custom patch. Only custom patches that use internal settings exposed to the user need to use the class.
//
// You access the internal settings of a custom patch through key-value coding (KVC). All the KVC keys that represent the internal settings of the custom patch must be listed in its method. The view controller for a custom patch expects the nib file class set to the class the view outlet connected to the view that contains the editing controls The controls are bound to the as the target and as the model key path, where is the KVC key for a given internal setting of the custom patch instance.


// The class communicates (through Cocoa bindings) between a custom patch and the view used for the internal settings of the custom patch. Only custom patches that use internal settings exposed to the user need to use the class.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Quartz/QCPlugInViewController
type QCPlugInViewController struct {
	appkit.ViewController
}

// QCPlugInViewControllerFrom constructs a [QCPlugInViewController] from an unsafe.Pointer.
//
// The class communicates (through Cocoa bindings) between a custom patch and the view used for the internal settings of the custom patch. Only custom patches that use internal settings exposed to the user need to use the class.
func QCPlugInViewControllerFrom(ptr unsafe.Pointer) QCPlugInViewController {
	return QCPlugInViewController{
		ViewController: appkit.ViewControllerFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for QCPlugInViewController */

// Creates and initializes a controller for the specified object and nib file.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Quartz/QCPlugInViewController/init(plugIn:viewNibName:)
func NewQCPlugInViewControllerWithPlugInViewNibName(plugIn IQCPlugIn, name objc.IObject /* cross-framework: NSString */) QCPlugInViewController {
	instance := getQCPlugInViewControllerClass().Alloc()
	rv := objc.Send[QCPlugInViewController](instance.ID, objc.Sel("initWithPlugIn:viewNibName:"), plugIn, name)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewQCPlugInViewControllerWithPlugInViewNibName */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for QCPlugInViewController */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for QCPlugInViewController */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for QCPlugInViewController */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for QCPlugInViewController */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class QCPlugInViewController */


