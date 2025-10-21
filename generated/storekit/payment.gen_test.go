// Code generated from Apple documentation for StoreKit. DO NOT EDIT.

package storekit_test

import (
	"github.com/tmc/appledocs/generated/storekit"
)

// Suppress unused import errors
var _ = storekit.NewPayment

// ExampleNewPaymentWithProduct demonstrates how to create a Payment instance using NewPaymentWithProduct.
// Returns a new payment for the specified product.
func ExampleNewPaymentWithProduct() {
	_ = storekit.NewPaymentWithProduct(
		storekit.SKProduct{}, // product SKProduct
	)
	// Output:
}
