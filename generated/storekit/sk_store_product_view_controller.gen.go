// Code generated from Apple documentation for StoreKit. DO NOT EDIT.

package storekit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/appkit"
)

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

// An interface definition for the [StoreProductViewController] class.
type IStoreProductViewController interface {
	appkit.IViewController
	// properties:
	Delegate() StoreProductViewControllerDelegate /* not a class type */
	SetDelegate(value StoreProductViewControllerDelegate /* not a class type */)
	// methods:
}

// A view controller that provides a page where customers can purchase media from the App Store.
//
// To display a store for customers to purchase media from the App Store, follow these steps: Create an object and set its . Indicate a specific product to sell by passing its iTunes item identifier to the method. Present the view controller modally from another view controller in your app. Your delegate dismisses the view controller when the customer completes the purchase. Present the object immediately when someone triggers an interaction, such as tapping a Buy button. Load the product information before presenting the view controller to ensure a seamless user experience. This class ignores settings, and those settings have no impact on the sheet’s presentation. To recommend another app without displaying a full product page, and to recommend an App Clip’s corresponding app from within the App Clip, use .


// A view controller that provides a page where customers can purchase media from the App Store.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/StoreKit/SKStoreProductViewController
type StoreProductViewController struct {
	appkit.ViewController
}

// StoreProductViewControllerFrom constructs a [StoreProductViewController] from an unsafe.Pointer.
//
// A view controller that provides a page where customers can purchase media from the App Store.
func StoreProductViewControllerFrom(ptr unsafe.Pointer) StoreProductViewController {
	return StoreProductViewController{
		ViewController: appkit.ViewControllerFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (sc _StoreProductViewControllerClass) Alloc() StoreProductViewController {
	rv := objc.Send[StoreProductViewController](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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



// The store view controller’s delegate.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/storekit/skstoreproductviewcontroller/delegate
func (s_ StoreProductViewController) Delegate() StoreProductViewControllerDelegate /* not a class type */ {
	rv := objc.Send[StoreProductViewControllerDelegate](s_.ID, objc.Sel("delegate"))
	return rv
}


// The store view controller’s delegate.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/storekit/skstoreproductviewcontroller/delegate
func (s_ StoreProductViewController) SetDelegate(value StoreProductViewControllerDelegate /* not a class type */) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setDelegate:"), value)
}



