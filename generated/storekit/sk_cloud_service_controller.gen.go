// Code generated from Apple documentation for StoreKit. DO NOT EDIT.

package storekit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
)

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

// An interface definition for the [CloudServiceController] class.
type ICloudServiceController interface {
	objectivec.IObject
	RequestStorefrontCountryCodeWithCompletionHandler(completionHandler unsafe.Pointer)
}

// An object that determines the current capabilities of a person’s Music library.
//
// Use an object to determine the current capabilities of a customer’s Music library, like whether the device allows playback of Apple Music catalog tracks and the addition of tracks to the library.
//
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

// Alloc allocates a new instance without initialization.
func (cc _CloudServiceControllerClass) Alloc() CloudServiceController {
	rv := objc.Send[CloudServiceController](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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


// Returns the type of authorization the customer has for accessing the Music library on the device.
//
// [Full Topic]: https://developer.apple.com/documentation/StoreKit/SKCloudServiceController/authorizationStatus()
func (cc _CloudServiceControllerClass) AuthorizationStatus() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(cc.class), objc.Sel("authorizationStatus"))
	return rv
}

// Asks the customer for permission to access the Music library on the device.
//
// [Full Topic]: https://developer.apple.com/documentation/StoreKit/SKCloudServiceController/requestAuthorization(_:)
func (cc _CloudServiceControllerClass) RequestAuthorization(completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(cc.class), objc.Sel("requestAuthorization:"), completionHandler)
}

// Gets the country code for the storefront associated with a customer’s iTunes account.
//
// [Full Topic]: https://developer.apple.com/documentation/StoreKit/SKCloudServiceController/requestStorefrontCountryCode(completionHandler:)
func (c_ CloudServiceController) RequestStorefrontCountryCodeWithCompletionHandler(completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("requestStorefrontCountryCodeWithCompletionHandler:"), completionHandler)
}



