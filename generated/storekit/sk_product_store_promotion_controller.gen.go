// Code generated from Apple documentation for StoreKit. DO NOT EDIT.

package storekit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class SKProductStorePromotionController */


/* debug [class_header]: Header for SKProductStorePromotionController */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for ProductStorePromotionController */
// An interface definition for the [ProductStorePromotionController] class.
type IProductStorePromotionController interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for ProductStorePromotionController */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for ProductStorePromotionController */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for ProductStorePromotionController */
// Alloc allocates a new instance without initialization.
func (pc _ProductStorePromotionControllerClass) Alloc() ProductStorePromotionController {
	rv := objc.Send[ProductStorePromotionController](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for ProductStorePromotionController */
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for ProductStorePromotionController *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for ProductStorePromotionController */

// Returns the default product store promotion controller.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/StoreKit/SKProductStorePromotionController/default()
func (pc _ProductStorePromotionControllerClass) DefaultController() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(pc.class), objc.Sel("defaultController"))
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=DefaultController) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for ProductStorePromotionController */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for ProductStorePromotionController */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for ProductStorePromotionController */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class SKProductStorePromotionController */



