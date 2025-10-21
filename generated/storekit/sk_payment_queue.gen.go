// Code generated from Apple documentation for StoreKit. DO NOT EDIT.

package storekit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [PaymentQueue] class.
var (
	PaymentQueueClass     _PaymentQueueClass
	PaymentQueueClassOnce sync.Once
)

func getPaymentQueueClass() _PaymentQueueClass {
	PaymentQueueClassOnce.Do(func() {
		PaymentQueueClass = _PaymentQueueClass{objc.GetClass("SKPaymentQueue")}
	})
	return PaymentQueueClass
}

type _PaymentQueueClass struct {
	class objc.Class
}

// An interface definition for the [PaymentQueue] class.
type IPaymentQueue interface {
	objectivec.IObject
	AddPayment(payment unsafe.Pointer)
	AddTransactionObserver(observer objc.ID)
	CancelDownloads(downloads unsafe.Pointer)
	FinishTransaction(transaction unsafe.Pointer)
	PauseDownloads(downloads unsafe.Pointer)
	PresentCodeRedemptionSheet()
	RemoveTransactionObserver(observer objc.ID)
	RestoreCompletedTransactions()
	RestoreCompletedTransactionsWithApplicationUsername(username string)
	ResumeDownloads(downloads unsafe.Pointer)
	ShowPriceConsentIfNeeded()
	StartDownloads(downloads unsafe.Pointer)
}

// A queue of payment transactions for the App Store to process.
//
// The payment queue communicates with the App Store and presents a user interface so that the user can authorize payment. The contents of the queue are persistent between launches of your app. To process a payment, first add at least one observer object ( ) to the queue (see ). Then, add a payment object ( ) for the item the user wants to purchase. Each time you add a payment object, the queue creates a transaction object ( ) to process that payment and enqueues it to be processed. After payment is fulfilled, the queue updates the transaction object and then calls any observer objects to provide them the updated transaction. Your observer should process the transaction and then remove it from the queue. The exact mechanism you use to process a processed transaction depends on the design of your app and the product being purchased. Here are a few common examples: If the product is a feature already built into your app, your app enables the feature to process the transaction. If the product includes downloadable content provided by the App Store, your app retrieves the objects from the transaction and ask the payment queue to download them. You provide the actual content files to be served by the App Store to App Store Connect when you create the product information. If the product represents downloadable content provided by your own server, your app might open a network connection to your server and download the content from there. For more information on designing the payment processing portion of your app, see .
//
// [Full Topic]: https://developer.apple.com/documentation/StoreKit/SKPaymentQueue
type PaymentQueue struct {
	objectivec.Object
}

// PaymentQueueFrom constructs a [PaymentQueue] from an unsafe.Pointer.
//
// A queue of payment transactions for the App Store to process.
func PaymentQueueFrom(ptr unsafe.Pointer) PaymentQueue {
	return PaymentQueue{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (pc _PaymentQueueClass) Alloc() PaymentQueue {
	rv := objc.Send[PaymentQueue](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (pc _PaymentQueueClass) New() PaymentQueue {
	rv := objc.Send[PaymentQueue](objc.ID(pc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (p_ PaymentQueue) Init() PaymentQueue {
	rv := objc.Send[PaymentQueue](p_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p_ PaymentQueue) Autorelease() PaymentQueue {
	rv := objc.Send[PaymentQueue](p_.ID, objc.Sel("autorelease"))
	return rv
}

// NewPaymentQueue creates a new PaymentQueue instance.
func NewPaymentQueue() PaymentQueue {
	return getPaymentQueueClass().New()
}


// Adds a payment request to the queue.
//
// [Full Topic]: https://developer.apple.com/documentation/StoreKit/SKPaymentQueue/add(_:)-4vct1
func (p_ PaymentQueue) AddPayment(payment unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("addPayment:"), payment)
}

// Adds an observer to the payment queue.
//
// [Full Topic]: https://developer.apple.com/documentation/StoreKit/SKPaymentQueue/add(_:)-5ciz2
func (p_ PaymentQueue) AddTransactionObserver(observer objc.ID) {
	objc.Send[objc.ID](p_.ID, objc.Sel("addTransactionObserver:"), observer)
}

// Removes a set of downloads from the download list.
//
// [Full Topic]: https://developer.apple.com/documentation/StoreKit/SKPaymentQueue/cancel(_:)
func (p_ PaymentQueue) CancelDownloads(downloads unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("cancelDownloads:"), downloads)
}

// Notifies the App Store that the app finished processing the transaction.
//
// [Full Topic]: https://developer.apple.com/documentation/StoreKit/SKPaymentQueue/finishTransaction(_:)
func (p_ PaymentQueue) FinishTransaction(transaction unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("finishTransaction:"), transaction)
}

// Pauses a set of downloads.
//
// [Full Topic]: https://developer.apple.com/documentation/StoreKit/SKPaymentQueue/pause(_:)
func (p_ PaymentQueue) PauseDownloads(downloads unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("pauseDownloads:"), downloads)
}

// Displays a sheet that enables customers to redeem subscription offer codes that you configure in App Store Connect.
//
// [Full Topic]: https://developer.apple.com/documentation/StoreKit/SKPaymentQueue/presentCodeRedemptionSheet()
func (p_ PaymentQueue) PresentCodeRedemptionSheet() {
	objc.Send[objc.ID](p_.ID, objc.Sel("presentCodeRedemptionSheet"))
}

// Removes an observer from the payment queue.
//
// [Full Topic]: https://developer.apple.com/documentation/StoreKit/SKPaymentQueue/remove(_:)
func (p_ PaymentQueue) RemoveTransactionObserver(observer objc.ID) {
	objc.Send[objc.ID](p_.ID, objc.Sel("removeTransactionObserver:"), observer)
}

// Asks the payment queue to restore previously completed purchases.
//
// [Full Topic]: https://developer.apple.com/documentation/StoreKit/SKPaymentQueue/restoreCompletedTransactions()
func (p_ PaymentQueue) RestoreCompletedTransactions() {
	objc.Send[objc.ID](p_.ID, objc.Sel("restoreCompletedTransactions"))
}

// Asks the payment queue to restore previously completed purchases, providing an opaque identifier for the user’s account.
//
// [Full Topic]: https://developer.apple.com/documentation/StoreKit/SKPaymentQueue/restoreCompletedTransactions(withApplicationUsername:)
func (p_ PaymentQueue) RestoreCompletedTransactionsWithApplicationUsername(username string) {
	objc.Send[objc.ID](p_.ID, objc.Sel("restoreCompletedTransactionsWithApplicationUsername:"), objc.String(username))
}

// Resumes a set of downloads.
//
// [Full Topic]: https://developer.apple.com/documentation/StoreKit/SKPaymentQueue/resume(_:)
func (p_ PaymentQueue) ResumeDownloads(downloads unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("resumeDownloads:"), downloads)
}

// Asks the system to display the price consent sheet if the user hasn’t yet responded to a subscription price increase.
//
// [Full Topic]: https://developer.apple.com/documentation/StoreKit/SKPaymentQueue/showPriceConsentIfNeeded()
func (p_ PaymentQueue) ShowPriceConsentIfNeeded() {
	objc.Send[objc.ID](p_.ID, objc.Sel("showPriceConsentIfNeeded"))
}

// Adds a set of downloads to the download list.
//
// [Full Topic]: https://developer.apple.com/documentation/StoreKit/SKPaymentQueue/start(_:)
func (p_ PaymentQueue) StartDownloads(downloads unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("startDownloads:"), downloads)
}

// A delegate that provides information needed to complete transactions.
//
// [Full Topic]: https://developer.apple.com/documentation/StoreKit/SKPaymentQueue/delegate
func (p_ PaymentQueue) Delegate() objc.ID {
	rv := objc.Send[objc.ID](p_.ID, objc.Sel("delegate"))
	return rv
}


// SetDelegate sets the value of the delegate property.
// A delegate that provides information needed to complete transactions.

//
// [Full Topic]: https://developer.apple.com/documentation/StoreKit/SKPaymentQueue/delegate
func (p_ PaymentQueue) SetDelegate(value objc.ID) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setDelegate:"), value)
}
// Returns an array of pending transactions.
//
// [Full Topic]: https://developer.apple.com/documentation/StoreKit/SKPaymentQueue/transactions
func (p_ PaymentQueue) Transactions() []PaymentTransaction {
	rv := objc.Send[[]PaymentTransaction](p_.ID, objc.Sel("transactions"))
	return rv
}



