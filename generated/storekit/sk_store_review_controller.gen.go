// Code generated from Apple documentation for StoreKit. DO NOT EDIT.

package storekit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class SKStoreReviewController */


/* debug [class_header]: Header for SKStoreReviewController */
// The class instance for the [StoreReviewController] class.
var (
	StoreReviewControllerClass     _StoreReviewControllerClass
	StoreReviewControllerClassOnce sync.Once
)

func getStoreReviewControllerClass() _StoreReviewControllerClass {
	StoreReviewControllerClassOnce.Do(func() {
		StoreReviewControllerClass = _StoreReviewControllerClass{objc.GetClass("SKStoreReviewController")}
	})
	return StoreReviewControllerClass
}

type _StoreReviewControllerClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for StoreReviewController */
// An interface definition for the [StoreReviewController] class.
type IStoreReviewController interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for StoreReviewController */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for StoreReviewController */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for StoreReviewController */
// Alloc allocates a new instance without initialization.
func (sc _StoreReviewControllerClass) Alloc() StoreReviewController {
	rv := objc.Send[StoreReviewController](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (sc _StoreReviewControllerClass) New() StoreReviewController {
	rv := objc.Send[StoreReviewController](objc.ID(sc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (s_ StoreReviewController) Init() StoreReviewController {
	rv := objc.Send[StoreReviewController](s_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s_ StoreReviewController) Autorelease() StoreReviewController {
	rv := objc.Send[StoreReviewController](s_.ID, objc.Sel("autorelease"))
	return rv
}

// NewStoreReviewController creates a new StoreReviewController instance.
func NewStoreReviewController() StoreReviewController {
	return getStoreReviewControllerClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for StoreReviewController */
// An object that controls the process of requesting App Store ratings and reviews from customers.
//
// Use the method to indicate when it makes sense within the logic of your app to ask the customer for ratings and reviews.


// An object that controls the process of requesting App Store ratings and reviews from customers.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/StoreKit/SKStoreReviewController
type StoreReviewController struct {
	objectivec.Object
}

// StoreReviewControllerFrom constructs a [StoreReviewController] from an unsafe.Pointer.
//
// An object that controls the process of requesting App Store ratings and reviews from customers.
func StoreReviewControllerFrom(ptr unsafe.Pointer) StoreReviewController {
	return StoreReviewController{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for StoreReviewController *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for StoreReviewController */

// Tells StoreKit to ask the customer to rate or review your app, if appropriate.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/StoreKit/SKStoreReviewController/requestReview()
func (sc _StoreReviewControllerClass) RequestReview() {
	objc.Send[objc.ID](objc.ID(sc.class), objc.Sel("requestReview"))
}/* debug [class_methods/method]: Class method for%!(EXTRA string=RequestReview) */


// Tells StoreKit to ask the customer to rate or review the app, if appropriate, using the specified scene.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/StoreKit/SKStoreReviewController/requestReview(in:)
func (sc _StoreReviewControllerClass) RequestReviewInScene(windowScene WindowScene /* not a class type */) {
	objc.Send[objc.ID](objc.ID(sc.class), objc.Sel("requestReviewInScene:"), windowScene)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=RequestReviewInScene) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for StoreReviewController */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for StoreReviewController */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for StoreReviewController */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class SKStoreReviewController */


