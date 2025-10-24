// Code generated from Apple documentation for StoreKitTest. DO NOT EDIT.

package storekittest

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class SKTestSession */

/* debug [class_header]: Header for SKTestSession */
// The class instance for the [TestSession] class.
var (
	TestSessionClass     _TestSessionClass
	TestSessionClassOnce sync.Once
)

func getTestSessionClass() _TestSessionClass {
	TestSessionClassOnce.Do(func() {
		TestSessionClass = _TestSessionClass{objc.GetClass("SKTestSession")}
	})
	return TestSessionClass
}

type _TestSessionClass struct {
	class objc.Class
}

/* debug [class_header]: End header */

/* debug [class_interface]: Interface for TestSession */
// An interface definition for the [TestSession] class.
type ITestSession interface {
	objectivec.IObject

	/* debug [class_interface_properties]: Properties for TestSession */
	// properties:
	AskToBuyEnabled() bool
	SetAskToBuyEnabled(value bool)
	BillingGracePeriodEnabled() bool
	SetBillingGracePeriodEnabled(value bool)
	DisableDialogs() bool
	SetDisableDialogs(value bool)
	FailTransactionsEnabled() bool
	SetFailTransactionsEnabled(value bool)
	FailureError() ErrorCode /* not a class type */
	SetFailureError(value ErrorCode /* not a class type */)
	InterruptedPurchasesEnabled() bool
	SetInterruptedPurchasesEnabled(value bool)
	Locale() foundation.Locale
	SetLocale(value foundation.Locale)
	BillingRetryOnRenewalEnabled() bool
	SetBillingRetryOnRenewalEnabled(value bool)
	Storefront() objc.IObject /* cross-framework: NSString */
	SetStorefront(value objc.IObject /* cross-framework: NSString */)
	TimeRate() TestTimeRate
	SetTimeRate(value TestTimeRate)
	BillingGracePeriodIsEnabled() bool
	SetBillingGracePeriodIsEnabled(value bool)
	ShouldEnterBillingRetryOnRenewal() bool
	SetShouldEnterBillingRetryOnRenewal(value bool)
	/* debug [class_interface_properties]: End properties */

	/* debug [class_interface_methods]: Methods for TestSession */
	// methods:
	AllTransactions() []TestTransaction
	ApproveAskToBuyTransactionWithIdentifierError(identifier uint, error_ unsafe.Pointer) bool
	ClearTransactions()
	ConsentToPriceIncreaseForTransactionWithIdentifierError(identifier uint, error_ unsafe.Pointer) bool
	DeclineAskToBuyTransactionWithIdentifierError(identifier uint, error_ unsafe.Pointer) bool
	DeclinePriceIncreaseForTransactionWithIdentifierError(identifier uint, error_ unsafe.Pointer) bool
	DeleteTransactionWithIdentifierError(identifier uint, error_ unsafe.Pointer) bool
	DisableAutoRenewForTransactionWithIdentifierError(identifier uint, error_ unsafe.Pointer) bool
	EnableAutoRenewForTransactionWithIdentifierError(identifier uint, error_ unsafe.Pointer) bool
	ExpireSubscriptionWithProductIdentifierError(productIdentifier objc.IObject /* cross-framework: NSString */, error_ unsafe.Pointer) bool
	ForceRenewalOfSubscriptionWithProductIdentifierError(productIdentifier objc.IObject /* cross-framework: NSString */, error_ unsafe.Pointer) bool
	RefundTransactionWithIdentifierError(identifier uint, error_ unsafe.Pointer) bool
	RequestPriceIncreaseConsentForTransactionWithIdentifierError(identifier uint, error_ unsafe.Pointer) bool
	ResetToDefaultState()
	ResolveIssueForTransactionWithIdentifierError(identifier uint, error_ unsafe.Pointer) bool
	/* debug [class_interface_methods]: End methods */

}

/* debug [class_interface]: End interface */

/* debug [class_constructors]: Constructors for TestSession */
// Alloc allocates a new instance without initialization.
func (tc _TestSessionClass) Alloc() TestSession {
	rv := objc.Send[TestSession](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (tc _TestSessionClass) New() TestSession {
	rv := objc.Send[TestSession](objc.ID(tc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (t_ TestSession) Init() TestSession {
	rv := objc.Send[TestSession](t_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t_ TestSession) Autorelease() TestSession {
	rv := objc.Send[TestSession](t_.ID, objc.Sel("autorelease"))
	return rv
}

// NewTestSession creates a new TestSession instance.
func NewTestSession() TestSession {
	return getTestSessionClass().New()
}

/* debug [class_constructors]: End constructors */

/* debug [class_struct]: Struct for TestSession */
// The controls and environment configuration you use to test StoreKit transactions in Xcode.
//
// This class controls the settings that the server uses when it processes transactions. Run tests that reconfigure the environment serially, not concurrently, to avoid overwriting each other’s environment settings. The test environment creates an instance each time your test code calls any method of that affects in-app purchases, including: You can manage the transactions in the test environment. To get a list of all transactions in the test environment, call . To delete a single transaction, call . To delete all the transactions, call . Before automating a test session with , you must create a StoreKit configuration file. For more information, see and . Set to to run tests without showing test environment UI.

// The controls and environment configuration you use to test StoreKit transactions in Xcode.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/StoreKitTest/SKTestSession
type TestSession struct {
	objectivec.Object
}

// TestSessionFrom constructs a [TestSession] from an unsafe.Pointer.
//
// The controls and environment configuration you use to test StoreKit transactions in Xcode.
func TestSessionFrom(ptr unsafe.Pointer) TestSession {
	return TestSession{objectivec.Object{objc.ID(ptr)}}
}

/* debug [class_struct]: End struct */

/* debug [class_init_methods]: Init methods for TestSession */

// Initializes the test session with the provided configuration file that you include in your application’s bundle.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/StoreKitTest/SKTestSession/init(configurationFileNamed:)
func NewTestSessionWithConfigurationFileNamedError(filename objc.IObject /* cross-framework: NSString */, error_ unsafe.Pointer) TestSession {
	instance := getTestSessionClass().Alloc()
	rv := objc.Send[TestSession](instance.ID, objc.Sel("initWithConfigurationFileNamed:error:"), filename, error_)
	rv.Autorelease()
	return rv
} /* debug [class_init_methods/constructor]: NewTestSessionWithConfigurationFileNamedError */

// Initializes the test session with a configuration file you provide through a URL.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/StoreKitTest/SKTestSession/init(contentsOf:)
func NewTestSessionWithContentsOfURLError(fileURL objc.IObject /* cross-framework: NSURL */, error_ unsafe.Pointer) TestSession {
	instance := getTestSessionClass().Alloc()
	rv := objc.Send[TestSession](instance.ID, objc.Sel("initWithContentsOfURL:error:"), fileURL, error_)
	rv.Autorelease()
	return rv
} /* debug [class_init_methods/constructor]: NewTestSessionWithContentsOfURLError */

/* debug [class_init_methods]: End init methods */

/* debug [class_methods]: Class methods for TestSession */
/* debug [class_methods]: End class methods */

/* debug [class_properties_class]: Class properties for TestSession */
/* debug [class_properties_class]: End class properties */

/* debug [instance_methods]: Instance methods for TestSession */

// Gets a list of all transactions in the test environment.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/StoreKitTest/SKTestSession/allTransactions()
func (t_ TestSession) AllTransactions() []TestTransaction {
	rv := objc.Send[[]TestTransaction](t_.ID, objc.Sel("allTransactions"))
	return rv
} /* debug [instance_methods/method]: AllTransactions */

// Resolves an Ask to Buy test scenario by approving the transaction.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/StoreKitTest/SKTestSession/approveAskToBuyTransaction(identifier:)
func (t_ TestSession) ApproveAskToBuyTransactionWithIdentifierError(identifier uint, error_ unsafe.Pointer) bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("approveAskToBuyTransactionWithIdentifier:error:"), identifier, error_)
	return rv
} /* debug [instance_methods/method]: ApproveAskToBuyTransactionWithIdentifierError */

// Removes all transactions from the test environment.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/StoreKitTest/SKTestSession/clearTransactions()
func (t_ TestSession) ClearTransactions() {
	objc.Send[objc.ID](t_.ID, objc.Sel("clearTransactions"))
} /* debug [instance_methods/method]: ClearTransactions */

// Simulates a user consenting to a price increase for an auto-renewable subscription.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/StoreKitTest/SKTestSession/consentToPriceIncreaseForTransaction(identifier:)
func (t_ TestSession) ConsentToPriceIncreaseForTransactionWithIdentifierError(identifier uint, error_ unsafe.Pointer) bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("consentToPriceIncreaseForTransactionWithIdentifier:error:"), identifier, error_)
	return rv
} /* debug [instance_methods/method]: ConsentToPriceIncreaseForTransactionWithIdentifierError */

// Resolves an Ask to Buy test scenario by declining the transaction.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/StoreKitTest/SKTestSession/declineAskToBuyTransaction(identifier:)
func (t_ TestSession) DeclineAskToBuyTransactionWithIdentifierError(identifier uint, error_ unsafe.Pointer) bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("declineAskToBuyTransactionWithIdentifier:error:"), identifier, error_)
	return rv
} /* debug [instance_methods/method]: DeclineAskToBuyTransactionWithIdentifierError */

// Simulates a user canceling an auto-renewable subscription by disabling auto-renew.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/StoreKitTest/SKTestSession/declinePriceIncreaseForTransaction(identifier:)
func (t_ TestSession) DeclinePriceIncreaseForTransactionWithIdentifierError(identifier uint, error_ unsafe.Pointer) bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("declinePriceIncreaseForTransactionWithIdentifier:error:"), identifier, error_)
	return rv
} /* debug [instance_methods/method]: DeclinePriceIncreaseForTransactionWithIdentifierError */

// Deletes a specific transaction from the test environment.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/StoreKitTest/SKTestSession/deleteTransaction(identifier:)
func (t_ TestSession) DeleteTransactionWithIdentifierError(identifier uint, error_ unsafe.Pointer) bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("deleteTransactionWithIdentifier:error:"), identifier, error_)
	return rv
} /* debug [instance_methods/method]: DeleteTransactionWithIdentifierError */

// Disables auto-renewing for an auto-renewable subscription in the test environment.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/StoreKitTest/SKTestSession/disableAutoRenewForTransaction(identifier:)
func (t_ TestSession) DisableAutoRenewForTransactionWithIdentifierError(identifier uint, error_ unsafe.Pointer) bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("disableAutoRenewForTransactionWithIdentifier:error:"), identifier, error_)
	return rv
} /* debug [instance_methods/method]: DisableAutoRenewForTransactionWithIdentifierError */

// Enables auto-renewing for an auto-renewable subscription in the test environment.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/StoreKitTest/SKTestSession/enableAutoRenewForTransaction(identifier:)
func (t_ TestSession) EnableAutoRenewForTransactionWithIdentifierError(identifier uint, error_ unsafe.Pointer) bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("enableAutoRenewForTransactionWithIdentifier:error:"), identifier, error_)
	return rv
} /* debug [instance_methods/method]: EnableAutoRenewForTransactionWithIdentifierError */

// Causes the identified auto-renewable subscription to expire immediately in the test environment.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/StoreKitTest/SKTestSession/expireSubscription(productIdentifier:)
func (t_ TestSession) ExpireSubscriptionWithProductIdentifierError(productIdentifier objc.IObject /* cross-framework: NSString */, error_ unsafe.Pointer) bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("expireSubscriptionWithProductIdentifier:error:"), productIdentifier, error_)
	return rv
} /* debug [instance_methods/method]: ExpireSubscriptionWithProductIdentifierError */

// Ends the previous subscription period and begins the next period in the test environment.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/StoreKitTest/SKTestSession/forceRenewalOfSubscription(productIdentifier:)
func (t_ TestSession) ForceRenewalOfSubscriptionWithProductIdentifierError(productIdentifier objc.IObject /* cross-framework: NSString */, error_ unsafe.Pointer) bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("forceRenewalOfSubscriptionWithProductIdentifier:error:"), productIdentifier, error_)
	return rv
} /* debug [instance_methods/method]: ForceRenewalOfSubscriptionWithProductIdentifierError */

// Simulates a refund for an in-app purchase that completes outside of the app.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/StoreKitTest/SKTestSession/refundTransaction(identifier:)
func (t_ TestSession) RefundTransactionWithIdentifierError(identifier uint, error_ unsafe.Pointer) bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("refundTransactionWithIdentifier:error:"), identifier, error_)
	return rv
} /* debug [instance_methods/method]: RefundTransactionWithIdentifierError */

// Simulates a price increase that requires customer consent for an auto-renewable subscription.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/StoreKitTest/SKTestSession/requestPriceIncreaseConsentForTransaction(identifier:)
func (t_ TestSession) RequestPriceIncreaseConsentForTransactionWithIdentifierError(identifier uint, error_ unsafe.Pointer) bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("requestPriceIncreaseConsentForTransactionWithIdentifier:error:"), identifier, error_)
	return rv
} /* debug [instance_methods/method]: RequestPriceIncreaseConsentForTransactionWithIdentifierError */

// Removes all property overrides and resets all test session settings to their default state.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/StoreKitTest/SKTestSession/resetToDefaultState()
func (t_ TestSession) ResetToDefaultState() {
	objc.Send[objc.ID](t_.ID, objc.Sel("resetToDefaultState"))
} /* debug [instance_methods/method]: ResetToDefaultState */

// Simulates resolving an issue when you test interrupted purchases or billing retry scenarios.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/StoreKitTest/SKTestSession/resolveIssueForTransaction(identifier:)
func (t_ TestSession) ResolveIssueForTransactionWithIdentifierError(identifier uint, error_ unsafe.Pointer) bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("resolveIssueForTransactionWithIdentifier:error:"), identifier, error_)
	return rv
} /* debug [instance_methods/method]: ResolveIssueForTransactionWithIdentifierError */

/* debug [instance_methods]: End instance methods */

/* debug [instance_properties]: Instance properties for TestSession */

// A Boolean value that determines whether the testing environment simulates an Ask to Buy scenario.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/StoreKitTest/SKTestSession/askToBuyEnabled
func (t_ TestSession) AskToBuyEnabled() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("askToBuyEnabled"))
	return rv
} /* debug [instance_properties/getter]: askToBuyEnabled */

// A Boolean value that determines whether the testing environment simulates an Ask to Buy scenario.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/StoreKitTest/SKTestSession/askToBuyEnabled
func (t_ TestSession) SetAskToBuyEnabled(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setAskToBuyEnabled:"), value)
} /* debug [instance_properties/setter]: askToBuyEnabled */

// A Boolean value that indicates whether the test environment simulates a billing grace period for auto-renewable subscriptions.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/StoreKitTest/SKTestSession/billingGracePeriodIsEnabled
func (t_ TestSession) BillingGracePeriodEnabled() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("billingGracePeriodEnabled"))
	return rv
} /* debug [instance_properties/getter]: billingGracePeriodEnabled */

// A Boolean value that indicates whether the test environment simulates a billing grace period for auto-renewable subscriptions.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/StoreKitTest/SKTestSession/billingGracePeriodIsEnabled
func (t_ TestSession) SetBillingGracePeriodEnabled(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setBillingGracePeriodEnabled:"), value)
} /* debug [instance_properties/setter]: billingGracePeriodEnabled */

// A Boolean value that determines whether the testing environment disables dialogs during automated testing.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/StoreKitTest/SKTestSession/disableDialogs
func (t_ TestSession) DisableDialogs() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("disableDialogs"))
	return rv
} /* debug [instance_properties/getter]: disableDialogs */

// A Boolean value that determines whether the testing environment disables dialogs during automated testing.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/StoreKitTest/SKTestSession/disableDialogs
func (t_ TestSession) SetDisableDialogs(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setDisableDialogs:"), value)
} /* debug [instance_properties/setter]: disableDialogs */

// A Boolean value that determines whether transactions fail in the testing environment.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/StoreKitTest/SKTestSession/failTransactionsEnabled
func (t_ TestSession) FailTransactionsEnabled() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("failTransactionsEnabled"))
	return rv
} /* debug [instance_properties/getter]: failTransactionsEnabled */

// A Boolean value that determines whether transactions fail in the testing environment.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/StoreKitTest/SKTestSession/failTransactionsEnabled
func (t_ TestSession) SetFailTransactionsEnabled(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setFailTransactionsEnabled:"), value)
} /* debug [instance_properties/setter]: failTransactionsEnabled */

// The error code that transactions return when you enable failing transactions.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/StoreKitTest/SKTestSession/failureError
func (t_ TestSession) FailureError() ErrorCode /* not a class type */ {
	rv := objc.Send[ErrorCode](t_.ID, objc.Sel("failureError"))
	return rv
} /* debug [instance_properties/getter]: failureError */

// The error code that transactions return when you enable failing transactions.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/StoreKitTest/SKTestSession/failureError
func (t_ TestSession) SetFailureError(value ErrorCode /* not a class type */) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setFailureError:"), value)
} /* debug [instance_properties/setter]: failureError */

// A Boolean value that determines whether the test environment simulates an interrupted purchase.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/StoreKitTest/SKTestSession/interruptedPurchasesEnabled
func (t_ TestSession) InterruptedPurchasesEnabled() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("interruptedPurchasesEnabled"))
	return rv
} /* debug [instance_properties/getter]: interruptedPurchasesEnabled */

// A Boolean value that determines whether the test environment simulates an interrupted purchase.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/StoreKitTest/SKTestSession/interruptedPurchasesEnabled
func (t_ TestSession) SetInterruptedPurchasesEnabled(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setInterruptedPurchasesEnabled:"), value)
} /* debug [instance_properties/setter]: interruptedPurchasesEnabled */

// The value that determines the localization metadata the test environment uses.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/StoreKitTest/SKTestSession/locale
func (t_ TestSession) Locale() foundation.Locale {
	rv := objc.Send[foundation.Locale](t_.ID, objc.Sel("locale"))
	return rv
} /* debug [instance_properties/getter]: locale */

// The value that determines the localization metadata the test environment uses.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/StoreKitTest/SKTestSession/locale
func (t_ TestSession) SetLocale(value foundation.Locale) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setLocale:"), value)
} /* debug [instance_properties/setter]: locale */

// A Boolean value that indicates whether the testing environment enters a billing retry state when an auto-renewable subscription renews.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/StoreKitTest/SKTestSession/shouldEnterBillingRetryOnRenewal
func (t_ TestSession) BillingRetryOnRenewalEnabled() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("billingRetryOnRenewalEnabled"))
	return rv
} /* debug [instance_properties/getter]: billingRetryOnRenewalEnabled */

// A Boolean value that indicates whether the testing environment enters a billing retry state when an auto-renewable subscription renews.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/StoreKitTest/SKTestSession/shouldEnterBillingRetryOnRenewal
func (t_ TestSession) SetBillingRetryOnRenewalEnabled(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setBillingRetryOnRenewalEnabled:"), value)
} /* debug [instance_properties/setter]: billingRetryOnRenewalEnabled */

// The three-letter code that represents the region associated with the App Store storefront.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/StoreKitTest/SKTestSession/storefront
func (t_ TestSession) Storefront() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](t_.ID, objc.Sel("storefront"))
	return rv
} /* debug [instance_properties/getter]: storefront */

// The three-letter code that represents the region associated with the App Store storefront.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/StoreKitTest/SKTestSession/storefront
func (t_ TestSession) SetStorefront(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setStorefront:"), value)
} /* debug [instance_properties/setter]: storefront */

// The rate at which time passes for subscriptions in the test environment as compared to real time.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/StoreKitTest/SKTestSession/timeRate-swift.property
func (t_ TestSession) TimeRate() TestTimeRate {
	rv := objc.Send[TestTimeRate](t_.ID, objc.Sel("timeRate"))
	return rv
} /* debug [instance_properties/getter]: timeRate */

// The rate at which time passes for subscriptions in the test environment as compared to real time.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/StoreKitTest/SKTestSession/timeRate-swift.property
func (t_ TestSession) SetTimeRate(value TestTimeRate) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setTimeRate:"), value)
} /* debug [instance_properties/setter]: timeRate */

// A Boolean value that indicates whether the test environment simulates a billing grace period for auto-renewable subscriptions.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/storekittest/sktestsession/billinggraceperiodisenabled
func (t_ TestSession) BillingGracePeriodIsEnabled() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("billingGracePeriodIsEnabled"))
	return rv
} /* debug [instance_properties/getter]: billingGracePeriodIsEnabled */

// A Boolean value that indicates whether the test environment simulates a billing grace period for auto-renewable subscriptions.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/storekittest/sktestsession/billinggraceperiodisenabled
func (t_ TestSession) SetBillingGracePeriodIsEnabled(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setBillingGracePeriodIsEnabled:"), value)
} /* debug [instance_properties/setter]: billingGracePeriodIsEnabled */

// A Boolean value that indicates whether the testing environment enters a billing retry state when an auto-renewable subscription renews.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/storekittest/sktestsession/shouldenterbillingretryonrenewal
func (t_ TestSession) ShouldEnterBillingRetryOnRenewal() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("shouldEnterBillingRetryOnRenewal"))
	return rv
} /* debug [instance_properties/getter]: shouldEnterBillingRetryOnRenewal */

// A Boolean value that indicates whether the testing environment enters a billing retry state when an auto-renewable subscription renews.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/storekittest/sktestsession/shouldenterbillingretryonrenewal
func (t_ TestSession) SetShouldEnterBillingRetryOnRenewal(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setShouldEnterBillingRetryOnRenewal:"), value)
} /* debug [instance_properties/setter]: shouldEnterBillingRetryOnRenewal */

/* debug [instance_properties]: End instance properties */

/* debug [class.gen.go]: End class SKTestSession */
