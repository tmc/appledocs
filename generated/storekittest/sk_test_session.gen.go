// Code generated from Apple documentation for StoreKitTest. DO NOT EDIT.

package storekittest

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

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

// An interface definition for the [TestSession] class.
type ITestSession interface {
	objectivec.IObject
	AllTransactions() []TestTransaction
	ApproveAskToBuyTransactionWithIdentifierError(identifier uint, error_ unsafe.Pointer) bool
	BuyProductWithIdentifierError(productIdentifier string, error_ unsafe.Pointer) bool
	ClearTransactions()
	ConsentToPriceIncreaseForTransactionWithIdentifierError(identifier uint, error_ unsafe.Pointer) bool
	DeclineAskToBuyTransactionWithIdentifierError(identifier uint, error_ unsafe.Pointer) bool
	DeclinePriceIncreaseForTransactionWithIdentifierError(identifier uint, error_ unsafe.Pointer) bool
	DeleteTransactionWithIdentifierError(identifier uint, error_ unsafe.Pointer) bool
	DisableAutoRenewForTransactionWithIdentifierError(identifier uint, error_ unsafe.Pointer) bool
	EnableAutoRenewForTransactionWithIdentifierError(identifier uint, error_ unsafe.Pointer) bool
	ExpireSubscriptionWithProductIdentifierError(productIdentifier string, error_ unsafe.Pointer) bool
	ForceRenewalOfSubscriptionWithProductIdentifierError(productIdentifier string, error_ unsafe.Pointer) bool
	RefundTransactionWithIdentifierError(identifier uint, error_ unsafe.Pointer) bool
	RequestPriceIncreaseConsentForTransactionWithIdentifierError(identifier uint, error_ unsafe.Pointer) bool
	ResetToDefaultState()
	ResolveIssueForTransactionWithIdentifierError(identifier uint, error_ unsafe.Pointer) bool
}

// The controls and environment configuration you use to test StoreKit transactions in Xcode.
//
// This class controls the settings that the server uses when it processes transactions. Run tests that reconfigure the environment serially, not concurrently, to avoid overwriting each other’s environment settings. The test environment creates an instance each time your test code calls any method of that affects in-app purchases, including: You can manage the transactions in the test environment. To get a list of all transactions in the test environment, call . To delete a single transaction, call . To delete all the transactions, call . Before automating a test session with , you must create a StoreKit configuration file. For more information, see and . Set to to run tests without showing test environment UI.
//
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

// Alloc allocates a new instance without initialization.
func (tc _TestSessionClass) Alloc() TestSession {
	rv := objc.Send[TestSession](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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




// Initializes the test session with the provided configuration file that you include in your application’s bundle.
//
// [Full Topic]: https://developer.apple.com/documentation/StoreKitTest/SKTestSession/init(configurationFileNamed:)
func NewTestSessionWithConfigurationFileNamedError(filename string, error_ unsafe.Pointer) TestSession {
	instance := getTestSessionClass().Alloc()
	rv := objc.Send[TestSession](instance.ID, objc.Sel("initWithConfigurationFileNamed:error:"), objc.String(filename), error_)
	rv.Autorelease()
	return rv
}



// Initializes the test session with a configuration file you provide through a URL.
//
// [Full Topic]: https://developer.apple.com/documentation/StoreKitTest/SKTestSession/init(contentsOf:)
func NewTestSessionWithContentsOfURLError(fileURL unsafe.Pointer, error_ unsafe.Pointer) TestSession {
	instance := getTestSessionClass().Alloc()
	rv := objc.Send[TestSession](instance.ID, objc.Sel("initWithContentsOfURL:error:"), fileURL, error_)
	rv.Autorelease()
	return rv
}


// Gets a list of all transactions in the test environment.
//
// [Full Topic]: https://developer.apple.com/documentation/StoreKitTest/SKTestSession/allTransactions()
func (t_ TestSession) AllTransactions() []TestTransaction {
	rv := objc.Send[[]TestTransaction](t_.ID, objc.Sel("allTransactions"))
	return rv
}

// Resolves an Ask to Buy test scenario by approving the transaction.
//
// [Full Topic]: https://developer.apple.com/documentation/StoreKitTest/SKTestSession/approveAskToBuyTransaction(identifier:)
func (t_ TestSession) ApproveAskToBuyTransactionWithIdentifierError(identifier uint, error_ unsafe.Pointer) bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("approveAskToBuyTransactionWithIdentifier:error:"), identifier, error_)
	return rv
}

// Simulates buying an in-app purchase or subscription outside the app.
//
// [Full Topic]: https://developer.apple.com/documentation/StoreKitTest/SKTestSession/buyProduct(productIdentifier:)
func (t_ TestSession) BuyProductWithIdentifierError(productIdentifier string, error_ unsafe.Pointer) bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("buyProductWithIdentifier:error:"), objc.String(productIdentifier), error_)
	return rv
}

// Removes all transactions from the test environment.
//
// [Full Topic]: https://developer.apple.com/documentation/StoreKitTest/SKTestSession/clearTransactions()
func (t_ TestSession) ClearTransactions() {
	objc.Send[objc.ID](t_.ID, objc.Sel("clearTransactions"))
}

// Simulates a user consenting to a price increase for an auto-renewable subscription.
//
// [Full Topic]: https://developer.apple.com/documentation/StoreKitTest/SKTestSession/consentToPriceIncreaseForTransaction(identifier:)
func (t_ TestSession) ConsentToPriceIncreaseForTransactionWithIdentifierError(identifier uint, error_ unsafe.Pointer) bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("consentToPriceIncreaseForTransactionWithIdentifier:error:"), identifier, error_)
	return rv
}

// Resolves an Ask to Buy test scenario by declining the transaction.
//
// [Full Topic]: https://developer.apple.com/documentation/StoreKitTest/SKTestSession/declineAskToBuyTransaction(identifier:)
func (t_ TestSession) DeclineAskToBuyTransactionWithIdentifierError(identifier uint, error_ unsafe.Pointer) bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("declineAskToBuyTransactionWithIdentifier:error:"), identifier, error_)
	return rv
}

// Simulates a user canceling an auto-renewable subscription by disabling auto-renew.
//
// [Full Topic]: https://developer.apple.com/documentation/StoreKitTest/SKTestSession/declinePriceIncreaseForTransaction(identifier:)
func (t_ TestSession) DeclinePriceIncreaseForTransactionWithIdentifierError(identifier uint, error_ unsafe.Pointer) bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("declinePriceIncreaseForTransactionWithIdentifier:error:"), identifier, error_)
	return rv
}

// Deletes a specific transaction from the test environment.
//
// [Full Topic]: https://developer.apple.com/documentation/StoreKitTest/SKTestSession/deleteTransaction(identifier:)
func (t_ TestSession) DeleteTransactionWithIdentifierError(identifier uint, error_ unsafe.Pointer) bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("deleteTransactionWithIdentifier:error:"), identifier, error_)
	return rv
}

// Disables auto-renewing for an auto-renewable subscription in the test environment.
//
// [Full Topic]: https://developer.apple.com/documentation/StoreKitTest/SKTestSession/disableAutoRenewForTransaction(identifier:)
func (t_ TestSession) DisableAutoRenewForTransactionWithIdentifierError(identifier uint, error_ unsafe.Pointer) bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("disableAutoRenewForTransactionWithIdentifier:error:"), identifier, error_)
	return rv
}

// Enables auto-renewing for an auto-renewable subscription in the test environment.
//
// [Full Topic]: https://developer.apple.com/documentation/StoreKitTest/SKTestSession/enableAutoRenewForTransaction(identifier:)
func (t_ TestSession) EnableAutoRenewForTransactionWithIdentifierError(identifier uint, error_ unsafe.Pointer) bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("enableAutoRenewForTransactionWithIdentifier:error:"), identifier, error_)
	return rv
}

// Causes the identified auto-renewable subscription to expire immediately in the test environment.
//
// [Full Topic]: https://developer.apple.com/documentation/StoreKitTest/SKTestSession/expireSubscription(productIdentifier:)
func (t_ TestSession) ExpireSubscriptionWithProductIdentifierError(productIdentifier string, error_ unsafe.Pointer) bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("expireSubscriptionWithProductIdentifier:error:"), objc.String(productIdentifier), error_)
	return rv
}

// Ends the previous subscription period and begins the next period in the test environment.
//
// [Full Topic]: https://developer.apple.com/documentation/StoreKitTest/SKTestSession/forceRenewalOfSubscription(productIdentifier:)
func (t_ TestSession) ForceRenewalOfSubscriptionWithProductIdentifierError(productIdentifier string, error_ unsafe.Pointer) bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("forceRenewalOfSubscriptionWithProductIdentifier:error:"), objc.String(productIdentifier), error_)
	return rv
}

// Simulates a refund for an in-app purchase that completes outside of the app.
//
// [Full Topic]: https://developer.apple.com/documentation/StoreKitTest/SKTestSession/refundTransaction(identifier:)
func (t_ TestSession) RefundTransactionWithIdentifierError(identifier uint, error_ unsafe.Pointer) bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("refundTransactionWithIdentifier:error:"), identifier, error_)
	return rv
}

// Simulates a price increase that requires customer consent for an auto-renewable subscription.
//
// [Full Topic]: https://developer.apple.com/documentation/StoreKitTest/SKTestSession/requestPriceIncreaseConsentForTransaction(identifier:)
func (t_ TestSession) RequestPriceIncreaseConsentForTransactionWithIdentifierError(identifier uint, error_ unsafe.Pointer) bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("requestPriceIncreaseConsentForTransactionWithIdentifier:error:"), identifier, error_)
	return rv
}

// Removes all property overrides and resets all test session settings to their default state.
//
// [Full Topic]: https://developer.apple.com/documentation/StoreKitTest/SKTestSession/resetToDefaultState()
func (t_ TestSession) ResetToDefaultState() {
	objc.Send[objc.ID](t_.ID, objc.Sel("resetToDefaultState"))
}

// Simulates resolving an issue when you test interrupted purchases or billing retry scenarios.
//
// [Full Topic]: https://developer.apple.com/documentation/StoreKitTest/SKTestSession/resolveIssueForTransaction(identifier:)
func (t_ TestSession) ResolveIssueForTransactionWithIdentifierError(identifier uint, error_ unsafe.Pointer) bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("resolveIssueForTransactionWithIdentifier:error:"), identifier, error_)
	return rv
}

// A Boolean value that determines whether the testing environment simulates an Ask to Buy scenario.
//
// [Full Topic]: https://developer.apple.com/documentation/StoreKitTest/SKTestSession/askToBuyEnabled
func (t_ TestSession) AskToBuyEnabled() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("askToBuyEnabled"))
	return rv
}


// SetAskToBuyEnabled sets the value of the askToBuyEnabled property.
// A Boolean value that determines whether the testing environment simulates an Ask to Buy scenario.

//
// [Full Topic]: https://developer.apple.com/documentation/StoreKitTest/SKTestSession/askToBuyEnabled
func (t_ TestSession) SetAskToBuyEnabled(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setAskToBuyEnabled:"), value)
}

// A Boolean value that indicates whether the test environment simulates a billing grace period for auto-renewable subscriptions.
//
// [Full Topic]: https://developer.apple.com/documentation/StoreKitTest/SKTestSession/billingGracePeriodIsEnabled
func (t_ TestSession) BillingGracePeriodEnabled() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("billingGracePeriodEnabled"))
	return rv
}


// SetBillingGracePeriodEnabled sets the value of the billingGracePeriodEnabled property.
// A Boolean value that indicates whether the test environment simulates a billing grace period for auto-renewable subscriptions.

//
// [Full Topic]: https://developer.apple.com/documentation/StoreKitTest/SKTestSession/billingGracePeriodIsEnabled
func (t_ TestSession) SetBillingGracePeriodEnabled(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setBillingGracePeriodEnabled:"), value)
}

// A Boolean value that determines whether the testing environment disables dialogs during automated testing.
//
// [Full Topic]: https://developer.apple.com/documentation/StoreKitTest/SKTestSession/disableDialogs
func (t_ TestSession) DisableDialogs() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("disableDialogs"))
	return rv
}


// SetDisableDialogs sets the value of the disableDialogs property.
// A Boolean value that determines whether the testing environment disables dialogs during automated testing.

//
// [Full Topic]: https://developer.apple.com/documentation/StoreKitTest/SKTestSession/disableDialogs
func (t_ TestSession) SetDisableDialogs(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setDisableDialogs:"), value)
}

// A Boolean value that determines whether transactions fail in the testing environment.
//
// [Full Topic]: https://developer.apple.com/documentation/StoreKitTest/SKTestSession/failTransactionsEnabled
func (t_ TestSession) FailTransactionsEnabled() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("failTransactionsEnabled"))
	return rv
}


// SetFailTransactionsEnabled sets the value of the failTransactionsEnabled property.
// A Boolean value that determines whether transactions fail in the testing environment.

//
// [Full Topic]: https://developer.apple.com/documentation/StoreKitTest/SKTestSession/failTransactionsEnabled
func (t_ TestSession) SetFailTransactionsEnabled(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setFailTransactionsEnabled:"), value)
}

// The error code that transactions return when you enable failing transactions.
//
// [Full Topic]: https://developer.apple.com/documentation/StoreKitTest/SKTestSession/failureError
func (t_ TestSession) FailureError() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](t_.ID, objc.Sel("failureError"))
	return rv
}


// SetFailureError sets the value of the failureError property.
// The error code that transactions return when you enable failing transactions.

//
// [Full Topic]: https://developer.apple.com/documentation/StoreKitTest/SKTestSession/failureError
func (t_ TestSession) SetFailureError(value unsafe.Pointer) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setFailureError:"), value)
}

// A Boolean value that determines whether the test environment simulates an interrupted purchase.
//
// [Full Topic]: https://developer.apple.com/documentation/StoreKitTest/SKTestSession/interruptedPurchasesEnabled
func (t_ TestSession) InterruptedPurchasesEnabled() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("interruptedPurchasesEnabled"))
	return rv
}


// SetInterruptedPurchasesEnabled sets the value of the interruptedPurchasesEnabled property.
// A Boolean value that determines whether the test environment simulates an interrupted purchase.

//
// [Full Topic]: https://developer.apple.com/documentation/StoreKitTest/SKTestSession/interruptedPurchasesEnabled
func (t_ TestSession) SetInterruptedPurchasesEnabled(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setInterruptedPurchasesEnabled:"), value)
}

// The value that determines the localization metadata the test environment uses.
//
// [Full Topic]: https://developer.apple.com/documentation/StoreKitTest/SKTestSession/locale
func (t_ TestSession) Locale() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](t_.ID, objc.Sel("locale"))
	return rv
}


// SetLocale sets the value of the locale property.
// The value that determines the localization metadata the test environment uses.

//
// [Full Topic]: https://developer.apple.com/documentation/StoreKitTest/SKTestSession/locale
func (t_ TestSession) SetLocale(value unsafe.Pointer) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setLocale:"), value)
}

// A Boolean value that indicates whether the testing environment enters a billing retry state when an auto-renewable subscription renews.
//
// [Full Topic]: https://developer.apple.com/documentation/StoreKitTest/SKTestSession/shouldEnterBillingRetryOnRenewal
func (t_ TestSession) BillingRetryOnRenewalEnabled() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("billingRetryOnRenewalEnabled"))
	return rv
}


// SetBillingRetryOnRenewalEnabled sets the value of the billingRetryOnRenewalEnabled property.
// A Boolean value that indicates whether the testing environment enters a billing retry state when an auto-renewable subscription renews.

//
// [Full Topic]: https://developer.apple.com/documentation/StoreKitTest/SKTestSession/shouldEnterBillingRetryOnRenewal
func (t_ TestSession) SetBillingRetryOnRenewalEnabled(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setBillingRetryOnRenewalEnabled:"), value)
}

// The three-letter code that represents the region associated with the App Store storefront.
//
// [Full Topic]: https://developer.apple.com/documentation/StoreKitTest/SKTestSession/storefront
func (t_ TestSession) Storefront() string {
	rv := objc.Send[string](t_.ID, objc.Sel("storefront"))
	return rv
}


// SetStorefront sets the value of the storefront property.
// The three-letter code that represents the region associated with the App Store storefront.

//
// [Full Topic]: https://developer.apple.com/documentation/StoreKitTest/SKTestSession/storefront
func (t_ TestSession) SetStorefront(value string) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setStorefront:"), objc.String(value))
}

// The rate at which time passes for subscriptions in the test environment as compared to real time.
//
// [Full Topic]: https://developer.apple.com/documentation/StoreKitTest/SKTestSession/timeRate-swift.property
func (t_ TestSession) TimeRate() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](t_.ID, objc.Sel("timeRate"))
	return rv
}


// SetTimeRate sets the value of the timeRate property.
// The rate at which time passes for subscriptions in the test environment as compared to real time.

//
// [Full Topic]: https://developer.apple.com/documentation/StoreKitTest/SKTestSession/timeRate-swift.property
func (t_ TestSession) SetTimeRate(value unsafe.Pointer) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setTimeRate:"), value)
}


