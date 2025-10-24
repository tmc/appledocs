//go:build darwin && ios

// Code generated from Apple documentation for StoreKit. DO NOT EDIT.

package storekit

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/coretelephony"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// iOS-only methods for PaymentTransaction


// iOS-only properties

// A signed receipt that records all information about a successful payment transaction.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/StoreKit/SKPaymentTransaction/transactionReceipt
func (p_ PaymentTransaction) TransactionReceipt() objc.IObject /* cross-framework: NSData */ {
	rv := objc.Send[foundation.NSData](p_.ID, objc.Sel("transactionReceipt"))
	return rv
}





