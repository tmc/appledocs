// Code generated from Apple documentation for StoreKit. DO NOT EDIT.

package storekit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class SKCloudServiceController */


/* debug [class_header]: Header for SKCloudServiceController */
// The class instance for the [CloudServiceController] class.
var (
	CloudServiceControllerClass     _CloudServiceControllerClass
	CloudServiceControllerClassOnce sync.Once
)

func getCloudServiceControllerClass() _CloudServiceControllerClass {
	CloudServiceControllerClassOnce.Do(func() {
		CloudServiceControllerClass = _CloudServiceControllerClass{objc.GetClass("SKCloudServiceController")}
	})
	return CloudServiceControllerClass
}

type _CloudServiceControllerClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for CloudServiceController */
// An interface definition for the [CloudServiceController] class.
type ICloudServiceController interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for CloudServiceController */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for CloudServiceController */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for CloudServiceController */
// Alloc allocates a new instance without initialization.
func (cc _CloudServiceControllerClass) Alloc() CloudServiceController {
	rv := objc.Send[CloudServiceController](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _CloudServiceControllerClass) New() CloudServiceController {
	rv := objc.Send[CloudServiceController](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CloudServiceController) Init() CloudServiceController {
	rv := objc.Send[CloudServiceController](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CloudServiceController) Autorelease() CloudServiceController {
	rv := objc.Send[CloudServiceController](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCloudServiceController creates a new CloudServiceController instance.
func NewCloudServiceController() CloudServiceController {
	return getCloudServiceControllerClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for CloudServiceController */
// An object that determines the current capabilities of a person’s Music library.
//
// Use an object to determine the current capabilities of a customer’s Music library, like whether the device allows playback of Apple Music catalog tracks and the addition of tracks to the library.


// An object that determines the current capabilities of a person’s Music library.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/StoreKit/SKCloudServiceController
type CloudServiceController struct {
	objectivec.Object
}

// CloudServiceControllerFrom constructs a [CloudServiceController] from an unsafe.Pointer.
//
// An object that determines the current capabilities of a person’s Music library.
func CloudServiceControllerFrom(ptr unsafe.Pointer) CloudServiceController {
	return CloudServiceController{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for CloudServiceController *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for CloudServiceController */

// Returns the type of authorization the customer has for accessing the Music library on the device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/StoreKit/SKCloudServiceController/authorizationStatus()
func (cc _CloudServiceControllerClass) AuthorizationStatus() CloudServiceAuthorizationStatus {
	rv := objc.Send[CloudServiceAuthorizationStatus](objc.ID(cc.class), objc.Sel("authorizationStatus"))
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=AuthorizationStatus) */


// Asks the customer for permission to access the Music library on the device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/StoreKit/SKCloudServiceController/requestAuthorization(_:)
func (cc _CloudServiceControllerClass) RequestAuthorization(completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(cc.class), objc.Sel("requestAuthorization:"), completionHandler)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=RequestAuthorization) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for CloudServiceController */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for CloudServiceController */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for CloudServiceController */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class SKCloudServiceController */


