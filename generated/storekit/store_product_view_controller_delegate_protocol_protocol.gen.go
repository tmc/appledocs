// Code generated from Apple documentation for StoreKit. DO NOT EDIT.

package storekit

// PStoreProductViewControllerDelegate is the SKStoreProductViewControllerDelegate protocol interface.
//
// A protocol to call when the customer dismisses the store screen.
//
// Availability:
//   - Mac Catalyst 13.0+
//   - iOS 6.0+
//   - iPadOS 6.0+
//   - macOS 11.0+
//
// See: doc://com.apple.storekit/documentation/StoreKit/SKStoreProductViewControllerDelegate
type PStoreProductViewControllerDelegate interface {
	// Optional methods
	ProductViewControllerDidFinish(viewController ISKStoreProductViewController)
	HasProductViewControllerDidFinish() bool
}

// StoreProductViewControllerDelegate is a delegate implementation builder for the PStoreProductViewControllerDelegate protocol.
//
// Use this struct to create a custom delegate by setting handler functions for the methods you want to implement.
type StoreProductViewControllerDelegate struct {
	_ProductViewControllerDidFinish func(viewController ISKStoreProductViewController)
}

// SetProductViewControllerDidFinish sets the handler for the ProductViewControllerDidFinish delegate method.
//
// Called when the user dismisses the store screen.
func (d *StoreProductViewControllerDelegate) SetProductViewControllerDidFinish(f func(viewController ISKStoreProductViewController)) {
	d._ProductViewControllerDidFinish = f
}

// ProductViewControllerDidFinish implements the PStoreProductViewControllerDelegate interface.
func (d *StoreProductViewControllerDelegate) ProductViewControllerDidFinish(viewController ISKStoreProductViewController) {
	if d._ProductViewControllerDidFinish != nil {
		d._ProductViewControllerDidFinish(viewController)
	}
}

// HasProductViewControllerDidFinish returns true if a handler for ProductViewControllerDidFinish has been set.
func (d *StoreProductViewControllerDelegate) HasProductViewControllerDidFinish() bool {
	return d._ProductViewControllerDidFinish != nil
}
