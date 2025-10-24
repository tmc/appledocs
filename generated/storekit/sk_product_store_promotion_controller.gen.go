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
	// properties:
	// methods:
}

// A product promotion controller for customizing the order and visibility of In-App Purchases per device.
//
// For information about promoting In-App Purchases, see .


// A product promotion controller for customizing the order and visibility of In-App Purchases per device.
//
// [Full Topic]
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




