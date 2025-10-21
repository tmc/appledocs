// Code generated from Apple documentation for StoreKit. DO NOT EDIT.

package storekit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [ProductStorePromotionController] class.
var (
	ProductStorePromotionControllerClass     _ProductStorePromotionControllerClass
	ProductStorePromotionControllerClassOnce sync.Once
)

func getProductStorePromotionControllerClass() _ProductStorePromotionControllerClass {
	ProductStorePromotionControllerClassOnce.Do(func() {
		ProductStorePromotionControllerClass = _ProductStorePromotionControllerClass{objc.GetClass("SKProductStorePromotionController")}
	})
	return ProductStorePromotionControllerClass
}

type _ProductStorePromotionControllerClass struct {
	class objc.Class
}

// An interface definition for the [ProductStorePromotionController] class.
type IProductStorePromotionController interface {
	objectivec.IObject
	FetchStorePromotionOrderWithCompletionHandler(completionHandler unsafe.Pointer)
	FetchStorePromotionVisibilityForProductCompletionHandler(product unsafe.Pointer, completionHandler unsafe.Pointer)
	UpdateStorePromotionOrderCompletionHandler(promotionOrder unsafe.Pointer, completionHandler unsafe.Pointer)
	UpdateStorePromotionVisibilityForProductCompletionHandler(promotionVisibility unsafe.Pointer, product unsafe.Pointer, completionHandler unsafe.Pointer)
}

// A product promotion controller for customizing the order and visibility of In-App Purchases per device.
//
// For information about promoting In-App Purchases, see .
//
// [Full Topic]: https://developer.apple.com/documentation/StoreKit/SKProductStorePromotionController
type ProductStorePromotionController struct {
	objectivec.Object
}

// ProductStorePromotionControllerFrom constructs a [ProductStorePromotionController] from an unsafe.Pointer.
//
// A product promotion controller for customizing the order and visibility of In-App Purchases per device.
func ProductStorePromotionControllerFrom(ptr unsafe.Pointer) ProductStorePromotionController {
	return ProductStorePromotionController{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (pc _ProductStorePromotionControllerClass) Alloc() ProductStorePromotionController {
	rv := objc.Send[ProductStorePromotionController](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (pc _ProductStorePromotionControllerClass) New() ProductStorePromotionController {
	rv := objc.Send[ProductStorePromotionController](objc.ID(pc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (p_ ProductStorePromotionController) Init() ProductStorePromotionController {
	rv := objc.Send[ProductStorePromotionController](p_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p_ ProductStorePromotionController) Autorelease() ProductStorePromotionController {
	rv := objc.Send[ProductStorePromotionController](p_.ID, objc.Sel("autorelease"))
	return rv
}

// NewProductStorePromotionController creates a new ProductStorePromotionController instance.
func NewProductStorePromotionController() ProductStorePromotionController {
	return getProductStorePromotionControllerClass().New()
}


// Reads the product order override that determines the promoted product order on this device.
//
// [Full Topic]: https://developer.apple.com/documentation/StoreKit/SKProductStorePromotionController/fetchStorePromotionOrder(completionHandler:)
func (p_ ProductStorePromotionController) FetchStorePromotionOrderWithCompletionHandler(completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("fetchStorePromotionOrderWithCompletionHandler:"), completionHandler)
}

// Reads the visibility setting of a promoted product in the App Store for this device.
//
// [Full Topic]: https://developer.apple.com/documentation/StoreKit/SKProductStorePromotionController/fetchStorePromotionVisibility(for:completionHandler:)
func (p_ ProductStorePromotionController) FetchStorePromotionVisibilityForProductCompletionHandler(product unsafe.Pointer, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("fetchStorePromotionVisibilityForProduct:completionHandler:"), product, completionHandler)
}

// Overrides the promoted product order on this device.
//
// [Full Topic]: https://developer.apple.com/documentation/StoreKit/SKProductStorePromotionController/update(storePromotionOrder:completionHandler:)
func (p_ ProductStorePromotionController) UpdateStorePromotionOrderCompletionHandler(promotionOrder unsafe.Pointer, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("updateStorePromotionOrder:completionHandler:"), promotionOrder, completionHandler)
}

// Updates the visibility of the product on the App Store, per device.
//
// [Full Topic]: https://developer.apple.com/documentation/StoreKit/SKProductStorePromotionController/update(storePromotionVisibility:for:completionHandler:)
func (p_ ProductStorePromotionController) UpdateStorePromotionVisibilityForProductCompletionHandler(promotionVisibility unsafe.Pointer, product unsafe.Pointer, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("updateStorePromotionVisibility:forProduct:completionHandler:"), promotionVisibility, product, completionHandler)
}



