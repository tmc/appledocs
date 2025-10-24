// Code generated from Apple documentation for SystemExtensions. DO NOT EDIT.

package systemextensions

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class OSSystemExtensionRequest */


/* debug [class_header]: Header for OSSystemExtensionRequest */
// The class instance for the [OSSystemExtensionRequest] class.
var (
	OSSystemExtensionRequestClass     _OSSystemExtensionRequestClass
	OSSystemExtensionRequestClassOnce sync.Once
)

func getOSSystemExtensionRequestClass() _OSSystemExtensionRequestClass {
	OSSystemExtensionRequestClassOnce.Do(func() {
		OSSystemExtensionRequestClass = _OSSystemExtensionRequestClass{objc.GetClass("OSSystemExtensionRequest")}
	})
	return OSSystemExtensionRequestClass
}

type _OSSystemExtensionRequestClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for OSSystemExtensionRequest */
// An interface definition for the [OSSystemExtensionRequest] class.
type IOSSystemExtensionRequest interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for OSSystemExtensionRequest */
	// properties:
	Delegate() unsafe.Pointer
	SetDelegate(value unsafe.Pointer)
	Identifier() objc.IObject /* cross-framework: NSString */
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for OSSystemExtensionRequest */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for OSSystemExtensionRequest */
// Alloc allocates a new instance without initialization.
func (oc _OSSystemExtensionRequestClass) Alloc() OSSystemExtensionRequest {
	rv := objc.Send[OSSystemExtensionRequest](objc.ID(oc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (oc _OSSystemExtensionRequestClass) New() OSSystemExtensionRequest {
	rv := objc.Send[OSSystemExtensionRequest](objc.ID(oc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (o_ OSSystemExtensionRequest) Init() OSSystemExtensionRequest {
	rv := objc.Send[OSSystemExtensionRequest](o_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (o_ OSSystemExtensionRequest) Autorelease() OSSystemExtensionRequest {
	rv := objc.Send[OSSystemExtensionRequest](o_.ID, objc.Sel("autorelease"))
	return rv
}

// NewOSSystemExtensionRequest creates a new OSSystemExtensionRequest instance.
func NewOSSystemExtensionRequest() OSSystemExtensionRequest {
	return getOSSystemExtensionRequestClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for OSSystemExtensionRequest */
// A request to activate or deactivate a system extension.


// A request to activate or deactivate a system extension.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SystemExtensions/OSSystemExtensionRequest
type OSSystemExtensionRequest struct {
	objectivec.Object
}

// OSSystemExtensionRequestFrom constructs a [OSSystemExtensionRequest] from an unsafe.Pointer.
//
// A request to activate or deactivate a system extension.
func OSSystemExtensionRequestFrom(ptr unsafe.Pointer) OSSystemExtensionRequest {
	return OSSystemExtensionRequest{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for OSSystemExtensionRequest *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for OSSystemExtensionRequest */

// Creates a request to activate a System Extension.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SystemExtensions/OSSystemExtensionRequest/activationRequest(forExtensionWithIdentifier:queue:)
func (oc _OSSystemExtensionRequestClass) ActivationRequestForExtensionQueue(identifier objc.IObject /* cross-framework: NSString */, queue unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(oc.class), objc.Sel("activationRequestForExtension:queue:"), identifier, queue)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ActivationRequestForExtensionQueue) */


// Creates a request to deactivate a System Extension.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SystemExtensions/OSSystemExtensionRequest/deactivationRequest(forExtensionWithIdentifier:queue:)
func (oc _OSSystemExtensionRequestClass) DeactivationRequestForExtensionQueue(identifier objc.IObject /* cross-framework: NSString */, queue unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(oc.class), objc.Sel("deactivationRequestForExtension:queue:"), identifier, queue)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=DeactivationRequestForExtensionQueue) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SystemExtensions/OSSystemExtensionRequest/propertiesRequest(forExtensionWithIdentifier:queue:)
func (oc _OSSystemExtensionRequestClass) PropertiesRequestForExtensionQueue(identifier objc.IObject /* cross-framework: NSString */, queue unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(oc.class), objc.Sel("propertiesRequestForExtension:queue:"), identifier, queue)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=PropertiesRequestForExtensionQueue) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for OSSystemExtensionRequest */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for OSSystemExtensionRequest */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for OSSystemExtensionRequest */

// A delegate to receive updates about the progress of a request.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SystemExtensions/OSSystemExtensionRequest/delegate
func (o_ OSSystemExtensionRequest) Delegate() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](o_.ID, objc.Sel("delegate"))
	return rv
}/* debug [instance_properties/getter]: delegate */


// A delegate to receive updates about the progress of a request.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SystemExtensions/OSSystemExtensionRequest/delegate
func (o_ OSSystemExtensionRequest) SetDelegate(value unsafe.Pointer) {
	objc.Send[objc.ID](o_.ID, objc.Sel("setDelegate:"), value)
}/* debug [instance_properties/setter]: delegate */


// The bundle identifier of the target extension.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SystemExtensions/OSSystemExtensionRequest/identifier
func (o_ OSSystemExtensionRequest) Identifier() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](o_.ID, objc.Sel("identifier"))
	return rv
}/* debug [instance_properties/getter]: identifier */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class OSSystemExtensionRequest */



