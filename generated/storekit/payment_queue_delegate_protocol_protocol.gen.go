// Code generated from Apple documentation for StoreKit. DO NOT EDIT.

package storekit

import (
	"github.com/tmc/appledocs/generated/objc"
)

// PPaymentQueueDelegate is the SKPaymentQueueDelegate protocol interface.
//
// The protocol that provides information needed to complete transactions.
//
// Availability:
//   - Mac Catalyst 13.1+ (Deprecated in 18.0)
//   - iOS 13.0+ (Deprecated in 18.0)
//   - iPadOS 13.0+ (Deprecated in 18.0)
//   - macOS 10.15+ (Deprecated in 15.0)
//   - tvOS 13.0+ (Deprecated in 18.0)
//   - visionOS 1.0+ (Deprecated in 2.0)
//   - watchOS 6.2+ (Deprecated in 11.0)
//
// See: doc://com.apple.storekit/documentation/StoreKit/SKPaymentQueueDelegate
type PPaymentQueueDelegate interface {
	// Optional methods
	PaymentQueueShouldContinueTransactionInStorefront(paymentQueue ISKPaymentQueue, transaction ISKPaymentTransaction, newStorefront objc.IObject /* cross-framework: Storefront */) bool
	HasPaymentQueueShouldContinueTransactionInStorefront() bool
	PaymentQueueShouldShowPriceConsent(paymentQueue ISKPaymentQueue) bool
	HasPaymentQueueShouldShowPriceConsent() bool
}

// PaymentQueueDelegate is a delegate implementation builder for the PPaymentQueueDelegate protocol.
//
// Use this struct to create a custom delegate by setting handler functions for the methods you want to implement.
type PaymentQueueDelegate struct {
	_PaymentQueueShouldContinueTransactionInStorefront func(paymentQueue ISKPaymentQueue, transaction ISKPaymentTransaction, newStorefront objc.IObject /* cross-framework: Storefront */) bool
	_PaymentQueueShouldShowPriceConsent                func(paymentQueue ISKPaymentQueue) bool
}

// SetPaymentQueueShouldContinueTransactionInStorefront sets the handler for the PaymentQueueShouldContinueTransactionInStorefront delegate method.
//
// Asks the delegate whether to continue the transaction if the device’s App Store storefront changes during a transaction.
func (d *PaymentQueueDelegate) SetPaymentQueueShouldContinueTransactionInStorefront(f func(paymentQueue ISKPaymentQueue, transaction ISKPaymentTransaction, newStorefront objc.IObject /* cross-framework: Storefront */) bool) {
	d._PaymentQueueShouldContinueTransactionInStorefront = f
}

// SetPaymentQueueShouldShowPriceConsent sets the handler for the PaymentQueueShouldShowPriceConsent delegate method.
//
// Asks the delegate whether to immediately display a price consent sheet.
func (d *PaymentQueueDelegate) SetPaymentQueueShouldShowPriceConsent(f func(paymentQueue ISKPaymentQueue) bool) {
	d._PaymentQueueShouldShowPriceConsent = f
}

// PaymentQueueShouldContinueTransactionInStorefront implements the PPaymentQueueDelegate interface.
func (d *PaymentQueueDelegate) PaymentQueueShouldContinueTransactionInStorefront(paymentQueue ISKPaymentQueue, transaction ISKPaymentTransaction, newStorefront objc.IObject /* cross-framework: Storefront */) bool {
	if d._PaymentQueueShouldContinueTransactionInStorefront != nil {
		return d._PaymentQueueShouldContinueTransactionInStorefront(paymentQueue, transaction, newStorefront)
	}
	var zero bool
	return zero
}

// HasPaymentQueueShouldContinueTransactionInStorefront returns true if a handler for PaymentQueueShouldContinueTransactionInStorefront has been set.
func (d *PaymentQueueDelegate) HasPaymentQueueShouldContinueTransactionInStorefront() bool {
	return d._PaymentQueueShouldContinueTransactionInStorefront != nil
}

// PaymentQueueShouldShowPriceConsent implements the PPaymentQueueDelegate interface.
func (d *PaymentQueueDelegate) PaymentQueueShouldShowPriceConsent(paymentQueue ISKPaymentQueue) bool {
	if d._PaymentQueueShouldShowPriceConsent != nil {
		return d._PaymentQueueShouldShowPriceConsent(paymentQueue)
	}
	var zero bool
	return zero
}

// HasPaymentQueueShouldShowPriceConsent returns true if a handler for PaymentQueueShouldShowPriceConsent has been set.
func (d *PaymentQueueDelegate) HasPaymentQueueShouldShowPriceConsent() bool {
	return d._PaymentQueueShouldShowPriceConsent != nil
}
