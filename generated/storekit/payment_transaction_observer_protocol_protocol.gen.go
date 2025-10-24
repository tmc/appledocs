// Code generated from Apple documentation for StoreKit. DO NOT EDIT.

package storekit

import (

	"github.com/tmc/appledocs/generated/coretelephony"
)

// PPaymentTransactionObserver is the SKPaymentTransactionObserver protocol interface.
//
// A set of methods that process transactions, unlock purchased functionality, and continue promoted In-App Purchases.
//
// Availability:
//   - Mac Catalyst 13.1+ (Deprecated in 18.0)
//   - iOS 3.0+ (Deprecated in 18.0)
//   - iPadOS 3.0+ (Deprecated in 18.0)
//   - macOS 10.7+ (Deprecated in 15.0)
//   - tvOS 9.0+ (Deprecated in 18.0)
//   - visionOS 1.0+ (Deprecated in 2.0)
//   - watchOS 6.2+ (Deprecated in 11.0)
//
// See: doc://com.apple.storekit/documentation/StoreKit/SKPaymentTransactionObserver
type PPaymentTransactionObserver interface {
	// Required methods
	PaymentQueueUpdatedTransactions(queue ISKPaymentQueue, transactions []PaymentTransaction)/* debug [protocol_interface/required_method]: PaymentQueueUpdatedTransactions */
	// Optional methods
	PaymentQueueDidRevokeEntitlementsForProductIdentifiers(queue ISKPaymentQueue, productIdentifiers []string)
	HasPaymentQueueDidRevokeEntitlementsForProductIdentifiers() bool
	PaymentQueueRemovedTransactions(queue ISKPaymentQueue, transactions []PaymentTransaction)
	HasPaymentQueueRemovedTransactions() bool
	PaymentQueueRestoreCompletedTransactionsFailedWithError(queue ISKPaymentQueue, error_ objc.IObject /* cross-framework: Error */)
	HasPaymentQueueRestoreCompletedTransactionsFailedWithError() bool
	PaymentQueueShouldAddStorePaymentForProduct(queue ISKPaymentQueue, payment ISKPayment, product ISKProduct) bool
	HasPaymentQueueShouldAddStorePaymentForProduct() bool
	PaymentQueueUpdatedDownloads(queue ISKPaymentQueue, downloads []Download)
	HasPaymentQueueUpdatedDownloads() bool
	PaymentQueueDidChangeStorefront(queue ISKPaymentQueue)
	HasPaymentQueueDidChangeStorefront() bool
	PaymentQueueRestoreCompletedTransactionsFinished(queue ISKPaymentQueue)
	HasPaymentQueueRestoreCompletedTransactionsFinished() bool
}
