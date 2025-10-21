// Code generated from Apple documentation for PassKit. DO NOT EDIT.

package passkit

// Enum types and constants
// PKAddPassButtonStyle - The appearance of the buttons that can be created using the 
//
// [Full Topic]: https://developer.apple.com/documentation/PassKit/PKAddPassButtonStyle
type AddPassButtonStyle uint

const (
	// AddPassButtonStyleBlack - A black button with white lettering.
	//
	// [Full Topic]: https://developer.apple.com/documentation/PassKit/PKAddPassButtonStyle/black
	AddPassButtonStyleBlack AddPassButtonStyle = 0
	// AddPassButtonStyleBlackOutline - A black button with a light outline.
	//
	// [Full Topic]: https://developer.apple.com/documentation/PassKit/PKAddPassButtonStyle/blackOutline
	AddPassButtonStyleBlackOutline AddPassButtonStyle = 0
)

// PKAddPaymentPassError - Error codes for adding payment passes.
//
// [Full Topic]: https://developer.apple.com/documentation/PassKit/PKAddPaymentPassError
type AddPaymentPassError uint

// PKAutomaticPassPresentationSuppressionResult - The result of an attempt to suppress automatic pass presentation.
//
// [Full Topic]: https://developer.apple.com/documentation/PassKit/PKAutomaticPassPresentationSuppressionResult
type AutomaticPassPresentationSuppressionResult uint

const (
	// AutomaticPassPresentationSuppressionResultAlreadyPresenting - The device is already presenting passes.
	//
	// [Full Topic]: https://developer.apple.com/documentation/PassKit/PKAutomaticPassPresentationSuppressionResult/alreadyPresenting
	AutomaticPassPresentationSuppressionResultAlreadyPresenting AutomaticPassPresentationSuppressionResult = 0
	// AutomaticPassPresentationSuppressionResultCancelled - The system canceled the suppression before calling the response handler.
	//
	// [Full Topic]: https://developer.apple.com/documentation/PassKit/PKAutomaticPassPresentationSuppressionResult/cancelled
	AutomaticPassPresentationSuppressionResultCancelled AutomaticPassPresentationSuppressionResult = 0
	// AutomaticPassPresentationSuppressionResultDenied - The user prevented the suppression, or an internal error occurred.
	//
	// [Full Topic]: https://developer.apple.com/documentation/PassKit/PKAutomaticPassPresentationSuppressionResult/denied
	AutomaticPassPresentationSuppressionResultDenied AutomaticPassPresentationSuppressionResult = 0
	// AutomaticPassPresentationSuppressionResultNotSupported - The device doesn’t support the suppression of automatic pass presentation.
	//
	// [Full Topic]: https://developer.apple.com/documentation/PassKit/PKAutomaticPassPresentationSuppressionResult/notSupported
	AutomaticPassPresentationSuppressionResultNotSupported AutomaticPassPresentationSuppressionResult = 0
	// AutomaticPassPresentationSuppressionResultSuccess - Suppression of automatic presentation successful.
	//
	// [Full Topic]: https://developer.apple.com/documentation/PassKit/PKAutomaticPassPresentationSuppressionResult/success
	AutomaticPassPresentationSuppressionResultSuccess AutomaticPassPresentationSuppressionResult = 0
)

// PKIssuerProvisioningExtensionAuthorizationResult - A value that indicates the result of authorizing the addition of a payment card.
//
// [Full Topic]: https://developer.apple.com/documentation/PassKit/PKIssuerProvisioningExtensionAuthorizationResult
type IssuerProvisioningExtensionAuthorizationResult uint

const (
	// IssuerProvisioningExtensionAuthorizationResultAuthorized - A result that indicates the user successfully authorized adding the payment pass.
	//
	// [Full Topic]: https://developer.apple.com/documentation/PassKit/PKIssuerProvisioningExtensionAuthorizationResult/authorized
	IssuerProvisioningExtensionAuthorizationResultAuthorized IssuerProvisioningExtensionAuthorizationResult = 0
)

// PKPassLibraryAuthorizationStatus enum type
//
// [Full Topic]: https://developer.apple.com/documentation/PassKit/PKPassLibrary/AuthorizationStatus
type PassLibraryAuthorizationStatus uint

const (
	//
	// [Full Topic]: https://developer.apple.com/documentation/PassKit/PKPassLibrary/AuthorizationStatus/authorized
	PassLibraryAuthorizationStatusAuthorized PassLibraryAuthorizationStatus = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/PassKit/PKPassLibrary/AuthorizationStatus/denied
	PassLibraryAuthorizationStatusDenied PassLibraryAuthorizationStatus = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/PassKit/PKPassLibrary/AuthorizationStatus/notDetermined
	PassLibraryAuthorizationStatusNotDetermined PassLibraryAuthorizationStatus = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/PassKit/PKPassLibrary/AuthorizationStatus/restricted
	PassLibraryAuthorizationStatusRestricted PassLibraryAuthorizationStatus = 0
)

// PKPassLibraryCapability enum type
//
// [Full Topic]: https://developer.apple.com/documentation/PassKit/PKPassLibrary/Capability
type PassLibraryCapability uint

const (
	//
	// [Full Topic]: https://developer.apple.com/documentation/PassKit/PKPassLibrary/Capability/backgroundAddPasses
	PassLibraryCapabilityBackgroundAddPasses PassLibraryCapability = 0
)

// PKPassLibraryAddPassesStatus - Statuses that PassKit uses when it adds passes to the pass library.
//
// [Full Topic]: https://developer.apple.com/documentation/PassKit/PKPassLibraryAddPassesStatus
type PassLibraryAddPassesStatus uint

const (
	// PassLibraryDidAddPasses - A status that occurs when the user successfully adds one or more passes.
	//
	// [Full Topic]: https://developer.apple.com/documentation/PassKit/PKPassLibraryAddPassesStatus/didAddPasses
	PassLibraryDidAddPasses PassLibraryAddPassesStatus = 0
	// PassLibraryDidCancelAddPasses - A status that occurs when the user cancels the addition of passes.
	//
	// [Full Topic]: https://developer.apple.com/documentation/PassKit/PKPassLibraryAddPassesStatus/didCancelAddPasses
	PassLibraryDidCancelAddPasses PassLibraryAddPassesStatus = 0
	// PassLibraryShouldReviewPasses - A status that occurs when the app prompts the user to review the passes.
	//
	// [Full Topic]: https://developer.apple.com/documentation/PassKit/PKPassLibraryAddPassesStatus/shouldReviewPasses
	PassLibraryShouldReviewPasses PassLibraryAddPassesStatus = 0
)

// PKPassType - Types of passes.
//
// [Full Topic]: https://developer.apple.com/documentation/PassKit/PKPassType
type PassType uint

// PKPaymentAuthorizationStatus - General success and failure status for payment authorization.
//
// [Full Topic]: https://developer.apple.com/documentation/PassKit/PKPaymentAuthorizationStatus
type PaymentAuthorizationStatus uint

const (
	// PaymentAuthorizationStatusFailure - Merchant failed to authorize the transaction.
	//
	// [Full Topic]: https://developer.apple.com/documentation/PassKit/PKPaymentAuthorizationStatus/failure
	PaymentAuthorizationStatusFailure PaymentAuthorizationStatus = 0
	// PaymentAuthorizationStatusInvalidBillingPostalAddress - Invalid or unusable billing address.
	//
	// [Full Topic]: https://developer.apple.com/documentation/PassKit/PKPaymentAuthorizationStatus/invalidBillingPostalAddress
	PaymentAuthorizationStatusInvalidBillingPostalAddress PaymentAuthorizationStatus = 0
	// PaymentAuthorizationStatusInvalidShippingContact - Invalid or incomplete shipping contact.
	//
	// [Full Topic]: https://developer.apple.com/documentation/PassKit/PKPaymentAuthorizationStatus/invalidShippingContact
	PaymentAuthorizationStatusInvalidShippingContact PaymentAuthorizationStatus = 0
	// PaymentAuthorizationStatusInvalidShippingPostalAddress - Invalid or unusable shipping address.
	//
	// [Full Topic]: https://developer.apple.com/documentation/PassKit/PKPaymentAuthorizationStatus/invalidShippingPostalAddress
	PaymentAuthorizationStatusInvalidShippingPostalAddress PaymentAuthorizationStatus = 0
	// PaymentAuthorizationStatusPINIncorrect - Incorrect PIN entered.
	//
	// [Full Topic]: https://developer.apple.com/documentation/PassKit/PKPaymentAuthorizationStatus/pinIncorrect
	PaymentAuthorizationStatusPINIncorrect PaymentAuthorizationStatus = 0
	// PaymentAuthorizationStatusPINLockout - PIN retry limit exceeded.
	//
	// [Full Topic]: https://developer.apple.com/documentation/PassKit/PKPaymentAuthorizationStatus/pinLockout
	PaymentAuthorizationStatusPINLockout PaymentAuthorizationStatus = 0
	// PaymentAuthorizationStatusPINRequired - Transaction requires PIN entry.
	//
	// [Full Topic]: https://developer.apple.com/documentation/PassKit/PKPaymentAuthorizationStatus/pinRequired
	PaymentAuthorizationStatusPINRequired PaymentAuthorizationStatus = 0
	// PaymentAuthorizationStatusSuccess - Merchant successfully authorized the transaction, or the transaction is expected to succeed.
	//
	// [Full Topic]: https://developer.apple.com/documentation/PassKit/PKPaymentAuthorizationStatus/success
	PaymentAuthorizationStatusSuccess PaymentAuthorizationStatus = 0
)

// PKPaymentErrorCode - An error code that you provide to indicate problems with address or contact information on an Apple Pay sheet.
//
// [Full Topic]: https://developer.apple.com/documentation/PassKit/PKPaymentError/Code
type PaymentErrorCode uint

// PKPaymentSummaryItemType - Constants that describe the type of the payment summary item, such as final or pending.
//
// [Full Topic]: https://developer.apple.com/documentation/PassKit/PKPaymentSummaryItemType
type PaymentSummaryItemType uint

// PKSecureElementPassActivationState - The activation states of a Secure Element pass.
//
// [Full Topic]: https://developer.apple.com/documentation/PassKit/PKSecureElementPass/PassActivationState-swift.enum
type SecureElementPassActivationState uint

const (
	// SecureElementPassActivationStateRequiresActivation - The pass requires activation by the issuer.
	//
	// [Full Topic]: https://developer.apple.com/documentation/PassKit/PKSecureElementPass/PassActivationState-swift.enum/requiresActivation
	SecureElementPassActivationStateRequiresActivation SecureElementPassActivationState = 0
)

// PKShareSecureElementPassResult enum type
//
// [Full Topic]: https://developer.apple.com/documentation/PassKit/PKShareSecureElementPassResult
type ShareSecureElementPassResult uint

const (
	//
	// [Full Topic]: https://developer.apple.com/documentation/PassKit/PKShareSecureElementPassResult/canceled
	ShareSecureElementPassResultCanceled ShareSecureElementPassResult = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/PassKit/PKShareSecureElementPassResult/failed
	ShareSecureElementPassResultFailed ShareSecureElementPassResult = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/PassKit/PKShareSecureElementPassResult/shared
	ShareSecureElementPassResultShared ShareSecureElementPassResult = 0
)

// PKShippingContactEditingMode - Constants that indicate whether the shipping mode prevents the user from editing fields of the shipping address.
//
// [Full Topic]: https://developer.apple.com/documentation/PassKit/PKShippingContactEditingMode
type ShippingContactEditingMode uint

const (
	// ShippingContactEditingModeStorePickup - The shipping contact on the payment sheet represents a pickup address and isn’t editable by the user.
	//
	// [Full Topic]: https://developer.apple.com/documentation/PassKit/PKShippingContactEditingMode/storePickup
	ShippingContactEditingModeStorePickup ShippingContactEditingMode = 0
)


