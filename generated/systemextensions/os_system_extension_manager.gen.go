// Code generated from Apple documentation for SystemExtensions. DO NOT EDIT.

package systemextensions

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class OSSystemExtensionManager */


/* debug [class_header]: Header for OSSystemExtensionManager */
// The class instance for the [OSSystemExtensionManager] class.
var (
	OSSystemExtensionManagerClass     _OSSystemExtensionManagerClass
	OSSystemExtensionManagerClassOnce sync.Once
)

func getOSSystemExtensionManagerClass() _OSSystemExtensionManagerClass {
	OSSystemExtensionManagerClassOnce.Do(func() {
		OSSystemExtensionManagerClass = _OSSystemExtensionManagerClass{objc.GetClass("OSSystemExtensionManager")}
	})
	return OSSystemExtensionManagerClass
}

type _OSSystemExtensionManagerClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for OSSystemExtensionManager */
// An interface definition for the [OSSystemExtensionManager] class.
type IOSSystemExtensionManager interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for OSSystemExtensionManager */
	// properties:
	Delegate() objc.IObject /* cross-framework: OSSystemExtensionRequestDelegate */
	SetDelegate(value objc.IObject /* cross-framework: OSSystemExtensionRequestDelegate */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for OSSystemExtensionManager */
	// methods:
	SubmitRequest(request IOSSystemExtensionRequest)
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for OSSystemExtensionManager */
// Alloc allocates a new instance without initialization.
func (oc _OSSystemExtensionManagerClass) Alloc() OSSystemExtensionManager {
	rv := objc.Send[OSSystemExtensionManager](objc.ID(oc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (oc _OSSystemExtensionManagerClass) New() OSSystemExtensionManager {
	rv := objc.Send[OSSystemExtensionManager](objc.ID(oc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (o_ OSSystemExtensionManager) Init() OSSystemExtensionManager {
	rv := objc.Send[OSSystemExtensionManager](o_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (o_ OSSystemExtensionManager) Autorelease() OSSystemExtensionManager {
	rv := objc.Send[OSSystemExtensionManager](o_.ID, objc.Sel("autorelease"))
	return rv
}

// NewOSSystemExtensionManager creates a new OSSystemExtensionManager instance.
func NewOSSystemExtensionManager() OSSystemExtensionManager {
	return getOSSystemExtensionManagerClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for OSSystemExtensionManager */
// A type that facilitates activation and deactivation of system extensions.
//
// Create an instance of with the class methods on that type, and submit it to the shared instance of the extension manager with . Set the on the request to receive the result of the activation or deactivation. The delegate also receives notifications if the user needs to authorize the extension or if a version conflict occurs.


// A type that facilitates activation and deactivation of system extensions.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SystemExtensions/OSSystemExtensionManager
type OSSystemExtensionManager struct {
	objectivec.Object
}

// OSSystemExtensionManagerFrom constructs a [OSSystemExtensionManager] from an unsafe.Pointer.
//
// A type that facilitates activation and deactivation of system extensions.
func OSSystemExtensionManagerFrom(ptr unsafe.Pointer) OSSystemExtensionManager {
	return OSSystemExtensionManager{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for OSSystemExtensionManager *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for OSSystemExtensionManager */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for OSSystemExtensionManager */

// The shared instance of the extension manager.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SystemExtensions/OSSystemExtensionManager/shared
func (oc _OSSystemExtensionManagerClass) SharedManager() OSSystemExtensionManager {
	rv := objc.Send[OSSystemExtensionManager](objc.ID(oc.class), objc.Sel("sharedManager"))
	return rv
}/* debug [class_properties_class/property]: sharedManager */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for OSSystemExtensionManager */

// Submits a system extension request to the manager.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SystemExtensions/OSSystemExtensionManager/submitRequest(_:)
func (o_ OSSystemExtensionManager) SubmitRequest(request IOSSystemExtensionRequest) {
	objc.Send[objc.ID](o_.ID, objc.Sel("submitRequest:"), request)
}/* debug [instance_methods/method]: SubmitRequest */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for OSSystemExtensionManager */

// The shared instance of the extension manager.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SystemExtensions/OSSystemExtensionManager/shared
func (o_ OSSystemExtensionManager) SharedManager() IOSSystemExtensionManager {
	rv := objc.Send[OSSystemExtensionManager](o_.ID, objc.Sel("sharedManager"))
	return rv
}/* debug [instance_properties/getter]: sharedManager */


// A delegate to receive updates about the progress of a request.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/systemextensions/ossystemextensionrequest/delegate
func (o_ OSSystemExtensionManager) Delegate() objc.IObject /* cross-framework: OSSystemExtensionRequestDelegate */ {
	rv := objc.Send[objc.ID](o_.ID, objc.Sel("delegate"))
	return rv
}/* debug [instance_properties/getter]: delegate */


// A delegate to receive updates about the progress of a request.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/systemextensions/ossystemextensionrequest/delegate
func (o_ OSSystemExtensionManager) SetDelegate(value objc.IObject /* cross-framework: OSSystemExtensionRequestDelegate */) {
	objc.Send[objc.ID](o_.ID, objc.Sel("setDelegate:"), value)
}/* debug [instance_properties/setter]: delegate */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class OSSystemExtensionManager */



