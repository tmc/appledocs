//go:build darwin && ios

// Code generated from Apple documentation for StoreKit. DO NOT EDIT.

package storekit

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

// iOS-only methods for StoreProductViewController


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/StoreKit/SKStoreProductViewController/loadProduct(withParameters:impression:completionBlock:)
func (s_ StoreProductViewController) LoadProductWithParametersImpressionCompletionBlock(parameters foundation.IDictionary, impression ISKAdImpression, block unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID, objc.Sel("loadProductWithParameters:impression:completionBlock:"), parameters, impression, block)
}

// iOS-only properties





