// Code generated from Apple documentation for StoreKit. DO NOT EDIT.

package storekit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class SKStoreProductViewController */

/* debug [class_header]: Header for SKStoreProductViewController */
// The class instance for the [StoreProductViewController] class.
var (
	StoreProductViewControllerClass     _StoreProductViewControllerClass
	StoreProductViewControllerClassOnce sync.Once
)

func getStoreProductViewControllerClass() _StoreProductViewControllerClass {
	StoreProductViewControllerClassOnce.Do(func() {
		StoreProductViewControllerClass = _StoreProductViewControllerClass{objc.GetClass("SKStoreProductViewController")}
	})
	return StoreProductViewControllerClass
}

type _StoreProductViewControllerClass struct {
	class objc.Class
}

/* debug [class_header]: End header */

/* debug [class_interface]: Interface for StoreProductViewController */
// An interface definition for the [StoreProductViewController] class.
type IStoreProductViewController interface {
	IViewController

	/* debug [class_interface_properties]: Properties for StoreProductViewController */
	// properties:
	Delegate() unsafe.Pointer
	SetDelegate(value unsafe.Pointer)
	/* debug [class_interface_properties]: End properties */

	/* debug [class_interface_methods]: Methods for StoreProductViewController */
	// methods:
	LoadProductWithParametersCompletionBlock(parameters foundation.IDictionary, block unsafe.Pointer)
	/* debug [class_interface_methods]: End methods */

}

/* debug [class_interface]: End interface */

/* debug [class_constructors]: Constructors for StoreProductViewController */
// Alloc allocates a new instance without initialization.
func (sc _StoreProductViewControllerClass) Alloc() StoreProductViewController {
	rv := objc.Send[StoreProductViewController](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (sc _StoreProductViewControllerClass) New() StoreProductViewController {
	rv := objc.Send[StoreProductViewController](objc.ID(sc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (s_ StoreProductViewController) Init() StoreProductViewController {
	rv := objc.Send[StoreProductViewController](s_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s_ StoreProductViewController) Autorelease() StoreProductViewController {
	rv := objc.Send[StoreProductViewController](s_.ID, objc.Sel("autorelease"))
	return rv
}

// NewStoreProductViewController creates a new StoreProductViewController instance.
func NewStoreProductViewController() StoreProductViewController {
	return getStoreProductViewControllerClass().New()
}

/* debug [class_constructors]: End constructors */

/* debug [class_struct]: Struct for StoreProductViewController */
// A view controller that provides a page where customers can purchase media from the App Store.
//
// To display a store for customers to purchase media from the App Store, follow these steps: Create an object and set its . Indicate a specific product to sell by passing its iTunes item identifier to the method. Present the view controller modally from another view controller in your app. Your delegate dismisses the view controller when the customer completes the purchase. Present the object immediately when someone triggers an interaction, such as tapping a Buy button. Load the product information before presenting the view controller to ensure a seamless user experience. This class ignores settings, and those settings have no impact on the sheet’s presentation. To recommend another app without displaying a full product page, and to recommend an App Clip’s corresponding app from within the App Clip, use .

// A view controller that provides a page where customers can purchase media from the App Store.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/StoreKit/SKStoreProductViewController
type StoreProductViewController struct {
	ViewController
}

// StoreProductViewControllerFrom constructs a [StoreProductViewController] from an unsafe.Pointer.
//
// A view controller that provides a page where customers can purchase media from the App Store.
func StoreProductViewControllerFrom(ptr unsafe.Pointer) StoreProductViewController {
	return StoreProductViewController{
		ViewController: ViewControllerFrom(ptr),
	}
}

/* debug [class_struct]: End struct */

/* debug [class_init_methods]: Init methods for StoreProductViewController */ /* debug [class_init_methods]: End init methods */

/* debug [class_methods]: Class methods for StoreProductViewController */
/* debug [class_methods]: End class methods */

/* debug [class_properties_class]: Class properties for StoreProductViewController */
/* debug [class_properties_class]: End class properties */

/* debug [instance_methods]: Instance methods for StoreProductViewController */

// Loads a new product screen to display.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/StoreKit/SKStoreProductViewController/loadProduct(withParameters:completionBlock:)
func (s_ StoreProductViewController) LoadProductWithParametersCompletionBlock(parameters foundation.IDictionary, block unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID, objc.Sel("loadProductWithParameters:completionBlock:"), parameters, block)
} /* debug [instance_methods/method]: LoadProductWithParametersCompletionBlock */

/* debug [instance_methods]: End instance methods */

/* debug [instance_properties]: Instance properties for StoreProductViewController */

// The store view controller’s delegate.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/StoreKit/SKStoreProductViewController/delegate
func (s_ StoreProductViewController) Delegate() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("delegate"))
	return rv
} /* debug [instance_properties/getter]: delegate */

// The store view controller’s delegate.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/StoreKit/SKStoreProductViewController/delegate
func (s_ StoreProductViewController) SetDelegate(value unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setDelegate:"), value)
} /* debug [instance_properties/setter]: delegate */

/* debug [instance_properties]: End instance properties */

/* debug [class.gen.go]: End class SKStoreProductViewController */
