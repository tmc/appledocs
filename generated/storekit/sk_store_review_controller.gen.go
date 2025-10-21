// Code generated from Apple documentation for StoreKit. DO NOT EDIT.

package storekit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
)

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

// An interface definition for the [StoreReviewController] class.
type IStoreReviewController interface {
	objectivec.IObject
}

// An object that controls the process of requesting App Store ratings and reviews from customers.
//
// Use the method to indicate when it makes sense within the logic of your app to ask the customer for ratings and reviews.
//
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

// Alloc allocates a new instance without initialization.
func (sc _StoreReviewControllerClass) Alloc() StoreReviewController {
	rv := objc.Send[StoreReviewController](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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


// Tells StoreKit to ask the customer to rate or review your app, if appropriate.
//
// [Full Topic]: https://developer.apple.com/documentation/StoreKit/SKStoreReviewController/requestReview()
func (sc _StoreReviewControllerClass) RequestReview() {
	objc.Send[objc.ID](objc.ID(sc.class), objc.Sel("requestReview"))
}

// Tells StoreKit to ask the customer to rate or review the app, if appropriate, using the specified scene.
//
// [Full Topic]: https://developer.apple.com/documentation/StoreKit/SKStoreReviewController/requestReview(in:)
func (sc _StoreReviewControllerClass) RequestReviewInScene(windowScene unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(sc.class), objc.Sel("requestReviewInScene:"), windowScene)
}



