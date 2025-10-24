// Code generated from Apple documentation for StoreKitTest. DO NOT EDIT.

package storekittest

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class SKTestTransaction */


/* debug [class_header]: Header for SKTestTransaction */
// The class instance for the [TestTransaction] class.
var (
	TestTransactionClass     _TestTransactionClass
	TestTransactionClassOnce sync.Once
)

func getTestTransactionClass() _TestTransactionClass {
	TestTransactionClassOnce.Do(func() {
		TestTransactionClass = _TestTransactionClass{objc.GetClass("SKTestTransaction")}
	})
	return TestTransactionClass
}

type _TestTransactionClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for TestTransaction */
// An interface definition for the [TestTransaction] class.
type ITestTransaction interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for TestTransaction */
	// properties:
	AutoRenewingEnabled() bool
	CancelDate() objc.IObject /* cross-framework: NSDate */
	ExpirationDate() objc.IObject /* cross-framework: NSDate */
	HasPurchaseIssue() bool
	Identifier() uint
	PendingPriceIncreaseConsent() bool
	OriginalTransactionIdentifier() uint
	PendingAskToBuyConfirmation() bool
	ProductIdentifier() objc.IObject /* cross-framework: NSString */
	PurchaseDate() objc.IObject /* cross-framework: NSDate */
	State() PaymentTransactionState /* not a class type */
	IsPendingPriceIncreaseConsent() bool
	SetIsPendingPriceIncreaseConsent(value bool)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for TestTransaction */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for TestTransaction */
// Alloc allocates a new instance without initialization.
func (tc _TestTransactionClass) Alloc() TestTransaction {
	rv := objc.Send[TestTransaction](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (tc _TestTransactionClass) New() TestTransaction {
	rv := objc.Send[TestTransaction](objc.ID(tc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (t_ TestTransaction) Init() TestTransaction {
	rv := objc.Send[TestTransaction](t_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t_ TestTransaction) Autorelease() TestTransaction {
	rv := objc.Send[TestTransaction](t_.ID, objc.Sel("autorelease"))
	return rv
}

// NewTestTransaction creates a new TestTransaction instance.
func NewTestTransaction() TestTransaction {
	return getTestTransactionClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for TestTransaction */
// A transaction that occurs in the testing environment.
//
// The test transaction represents the test environment’s knowledge of the transaction, including its identifier and the transaction’s state. It represents all the transaction-related configurations you control manually in Xcode for interrupted purchases, Ask to Buy scenarios, and changes to a subscription’s auto-renew state. The test environment creates an instance each time your test code calls any method of that affects in-app purchases.


// A transaction that occurs in the testing environment.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/StoreKitTest/SKTestTransaction
type TestTransaction struct {
	objectivec.Object
}

// TestTransactionFrom constructs a [TestTransaction] from an unsafe.Pointer.
//
// A transaction that occurs in the testing environment.
func TestTransactionFrom(ptr unsafe.Pointer) TestTransaction {
	return TestTransaction{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for TestTransaction *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for TestTransaction */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for TestTransaction */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for TestTransaction */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for TestTransaction */

// A Boolean value that indicates whether automatic renewal is enabled for the subscription.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/StoreKitTest/SKTestTransaction/autoRenewingEnabled
func (t_ TestTransaction) AutoRenewingEnabled() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("autoRenewingEnabled"))
	return rv
}/* debug [instance_properties/getter]: autoRenewingEnabled */


// The date when the system refunded the transaction.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/StoreKitTest/SKTestTransaction/cancelDate
func (t_ TestTransaction) CancelDate() objc.IObject /* cross-framework: NSDate */ {
	rv := objc.Send[foundation.NSDate](t_.ID, objc.Sel("cancelDate"))
	return rv
}/* debug [instance_properties/getter]: cancelDate */


// The date a subscription expires.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/StoreKitTest/SKTestTransaction/expirationDate
func (t_ TestTransaction) ExpirationDate() objc.IObject /* cross-framework: NSDate */ {
	rv := objc.Send[foundation.NSDate](t_.ID, objc.Sel("expirationDate"))
	return rv
}/* debug [instance_properties/getter]: expirationDate */


// A Boolean value that indicates whether you resolve this transaction using the test framework functions.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/StoreKitTest/SKTestTransaction/hasPurchaseIssue
func (t_ TestTransaction) HasPurchaseIssue() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("hasPurchaseIssue"))
	return rv
}/* debug [instance_properties/getter]: hasPurchaseIssue */


// The identifier of the transaction in the testing environment.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/StoreKitTest/SKTestTransaction/identifier
func (t_ TestTransaction) Identifier() uint {
	rv := objc.Send[uint](t_.ID, objc.Sel("identifier"))
	return rv
}/* debug [instance_properties/getter]: identifier */


// A Boolean value that indicates whether the auto-renewable subscription has a price increase that’s awaiting user consent in the test environment.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/StoreKitTest/SKTestTransaction/isPendingPriceIncreaseConsent
func (t_ TestTransaction) PendingPriceIncreaseConsent() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("pendingPriceIncreaseConsent"))
	return rv
}/* debug [instance_properties/getter]: pendingPriceIncreaseConsent */


// The identifier of the original transaction.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/StoreKitTest/SKTestTransaction/originalTransactionIdentifier
func (t_ TestTransaction) OriginalTransactionIdentifier() uint {
	rv := objc.Send[uint](t_.ID, objc.Sel("originalTransactionIdentifier"))
	return rv
}/* debug [instance_properties/getter]: originalTransactionIdentifier */


// A Boolean value that indicates whether the transaction is awaiting an Ask to Buy confirmation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/StoreKitTest/SKTestTransaction/pendingAskToBuyConfirmation
func (t_ TestTransaction) PendingAskToBuyConfirmation() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("pendingAskToBuyConfirmation"))
	return rv
}/* debug [instance_properties/getter]: pendingAskToBuyConfirmation */


// An identifier that uniquely represents a product, which you provide in the StoreKit configuration file.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/StoreKitTest/SKTestTransaction/productIdentifier
func (t_ TestTransaction) ProductIdentifier() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](t_.ID, objc.Sel("productIdentifier"))
	return rv
}/* debug [instance_properties/getter]: productIdentifier */


// The date of purchase for the transaction.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/StoreKitTest/SKTestTransaction/purchaseDate
func (t_ TestTransaction) PurchaseDate() objc.IObject /* cross-framework: NSDate */ {
	rv := objc.Send[foundation.NSDate](t_.ID, objc.Sel("purchaseDate"))
	return rv
}/* debug [instance_properties/getter]: purchaseDate */


// The state of the transaction in the test environment.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/StoreKitTest/SKTestTransaction/state
func (t_ TestTransaction) State() PaymentTransactionState /* not a class type */ {
	rv := objc.Send[PaymentTransactionState](t_.ID, objc.Sel("state"))
	return rv
}/* debug [instance_properties/getter]: state */


// A Boolean value that indicates whether the auto-renewable subscription has a price increase that’s awaiting user consent in the test environment.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/storekittest/sktesttransaction/ispendingpriceincreaseconsent
func (t_ TestTransaction) IsPendingPriceIncreaseConsent() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("isPendingPriceIncreaseConsent"))
	return rv
}/* debug [instance_properties/getter]: isPendingPriceIncreaseConsent */


// A Boolean value that indicates whether the auto-renewable subscription has a price increase that’s awaiting user consent in the test environment.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/storekittest/sktesttransaction/ispendingpriceincreaseconsent
func (t_ TestTransaction) SetIsPendingPriceIncreaseConsent(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setIsPendingPriceIncreaseConsent:"), value)
}/* debug [instance_properties/setter]: isPendingPriceIncreaseConsent */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class SKTestTransaction */






