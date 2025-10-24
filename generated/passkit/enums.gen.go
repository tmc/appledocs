// Code generated from Apple documentation for PassKit. DO NOT EDIT.

package passkit

/* debug [enums.gen.go]: Generating 38 enums for PassKit */
// Enum types and constants
/* debug [enums.gen.go]: Processing enum PKPassLibraryAuthorizationStatus (4 cases) */
// PKPassLibraryAuthorizationStatus enum type
//
// [Full Topic]: https://developer.apple.com/documentation/PassKit/PKPassLibrary/AuthorizationStatus
type PKPassLibraryAuthorizationStatus uint

const (
	//
	// [Full Topic]: https://developer.apple.com/documentation/PassKit/PKPassLibrary/AuthorizationStatus/authorized
	PKPassLibraryAuthorizationStatusAuthorized PKPassLibraryAuthorizationStatus = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/PassKit/PKPassLibrary/AuthorizationStatus/denied
	PKPassLibraryAuthorizationStatusDenied PKPassLibraryAuthorizationStatus = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/PassKit/PKPassLibrary/AuthorizationStatus/notDetermined
	PKPassLibraryAuthorizationStatusNotDetermined PKPassLibraryAuthorizationStatus = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/PassKit/PKPassLibrary/AuthorizationStatus/restricted
	PKPassLibraryAuthorizationStatusRestricted PKPassLibraryAuthorizationStatus = 0
)

/* debug [enums.gen.go]: Processing enum PKPassLibraryCapability (1 cases) */
// PKPassLibraryCapability enum type
//
// [Full Topic]: https://developer.apple.com/documentation/PassKit/PKPassLibrary/Capability
type PKPassLibraryCapability uint

const (
	//
	// [Full Topic]: https://developer.apple.com/documentation/PassKit/PKPassLibrary/Capability/backgroundAddPasses
	PKPassLibraryCapabilityBackgroundAddPasses PKPassLibraryCapability = 0
)

/* debug [enums.gen.go]: Processing enum PKAddIdentityDocumentType (3 cases) */
// PKAddIdentityDocumentType - Classifications that reflect the type of identity document.
//
// [Full Topic]: https://developer.apple.com/documentation/PassKit/PKAddIdentityDocumentType
type PKAddIdentityDocumentType uint

const (
	// PKAddIdentityDocumentTypeIDCard - A generic pass that represents a person’s identification.
	//
	// [Full Topic]: https://developer.apple.com/documentation/PassKit/PKAddIdentityDocumentType/idCard
	PKAddIdentityDocumentTypeIDCard PKAddIdentityDocumentType = 0
	// PKAddIdentityDocumentTypeMDL - A pass that represents a driver’s license or government-issued identification.
	//
	// [Full Topic]: https://developer.apple.com/documentation/PassKit/PKAddIdentityDocumentType/mDL
	PKAddIdentityDocumentTypeMDL PKAddIdentityDocumentType = 0
	// PKAddIdentityDocumentTypePhotoID - A pass to use for personal identification.
	//
	// [Full Topic]: https://developer.apple.com/documentation/PassKit/PKAddIdentityDocumentType/photoID
	PKAddIdentityDocumentTypePhotoID PKAddIdentityDocumentType = 0
)

/* debug [enums.gen.go]: Processing enum PKAddPassButtonStyle (2 cases) */
// PKAddPassButtonStyle - The appearance of the buttons that can be created using the 
//
// [Full Topic]: https://developer.apple.com/documentation/PassKit/PKAddPassButtonStyle
type PKAddPassButtonStyle uint

const (
	// PKAddPassButtonStyleBlack - A black button with white lettering.
	//
	// [Full Topic]: https://developer.apple.com/documentation/PassKit/PKAddPassButtonStyle/black
	PKAddPassButtonStyleBlack PKAddPassButtonStyle = 0
	// PKAddPassButtonStyleBlackOutline - A black button with a light outline.
	//
	// [Full Topic]: https://developer.apple.com/documentation/PassKit/PKAddPassButtonStyle/blackOutline
	PKAddPassButtonStyleBlackOutline PKAddPassButtonStyle = 0
)

/* debug [enums.gen.go]: Processing enum PKAddPaymentPassError (3 cases) */
// PKAddPaymentPassError - Error codes for adding payment passes.
//
// [Full Topic]: https://developer.apple.com/documentation/PassKit/PKAddPaymentPassError
type PKAddPaymentPassError uint

const (
	// PKAddPaymentPassErrorSystemCancelled - The system canceled the request to add a card to Apple Pay.
	//
	// [Full Topic]: https://developer.apple.com/documentation/PassKit/PKAddPaymentPassError/systemCancelled
	PKAddPaymentPassErrorSystemCancelled PKAddPaymentPassError = 0
	// PKAddPaymentPassErrorUnsupported - The app cannot add cards to Apple Pay.
	//
	// [Full Topic]: https://developer.apple.com/documentation/PassKit/PKAddPaymentPassError/unsupported
	PKAddPaymentPassErrorUnsupported PKAddPaymentPassError = 0
	// PKAddPaymentPassErrorUserCancelled - The user canceled the request to add a card to Apple Pay.
	//
	// [Full Topic]: https://developer.apple.com/documentation/PassKit/PKAddPaymentPassError/userCancelled
	PKAddPaymentPassErrorUserCancelled PKAddPaymentPassError = 0
)

/* debug [enums.gen.go]: Processing enum PKAddPaymentPassStyle (2 cases) */
// PKAddPaymentPassStyle - The type of payment pass.
//
// [Full Topic]: https://developer.apple.com/documentation/PassKit/PKAddPaymentPassStyle
type PKAddPaymentPassStyle uint

const (
	// PKAddPaymentPassStyleAccess - A pass that authorizes the user to access a location or resource.
	//
	// [Full Topic]: https://developer.apple.com/documentation/PassKit/PKAddPaymentPassStyle/access
	PKAddPaymentPassStyleAccess PKAddPaymentPassStyle = 0
	// PKAddPaymentPassStylePayment - A pass used by a customer for purchasing.
	//
	// [Full Topic]: https://developer.apple.com/documentation/PassKit/PKAddPaymentPassStyle/payment
	PKAddPaymentPassStylePayment PKAddPaymentPassStyle = 0
)

/* debug [enums.gen.go]: Processing enum PKAddressField (6 cases) */
// PKAddressField - Billing or shipping address fields.
//
// [Full Topic]: https://developer.apple.com/documentation/PassKit/PKAddressField
type PKAddressField uint

const (
	// PKAddressFieldAll - All fields.
	//
	// [Full Topic]: https://developer.apple.com/documentation/PassKit/PKAddressField/all
	PKAddressFieldAll PKAddressField = 0
	// PKAddressFieldEmail - The buyer’s email address.
	//
	// [Full Topic]: https://developer.apple.com/documentation/PassKit/PKAddressField/email
	PKAddressFieldEmail PKAddressField = 0
	// PKAddressFieldName - The buyer’s first and last name.
	//
	// [Full Topic]: https://developer.apple.com/documentation/PassKit/PKAddressField/name
	PKAddressFieldName PKAddressField = 0
	// PKAddressFieldPhone - The buyer’s telephone number.
	//
	// [Full Topic]: https://developer.apple.com/documentation/PassKit/PKAddressField/phone
	PKAddressFieldPhone PKAddressField = 0
	// PKAddressFieldNone - No fields.
	//
	// [Full Topic]: https://developer.apple.com/documentation/PassKit/PKAddressField/PKAddressFieldNone
	PKAddressFieldNone PKAddressField = 0
	// PKAddressFieldPostalAddress - The buyer’s full street address, including name, street, city, state or province, postal code, and country or region.
	//
	// [Full Topic]: https://developer.apple.com/documentation/PassKit/PKAddressField/postalAddress
	PKAddressFieldPostalAddress PKAddressField = 0
)

/* debug [enums.gen.go]: Processing enum PKAddSecureElementPassErrorCode (8 cases) */
// PKAddSecureElementPassErrorCode - Error codes for problems that occur when you add a secure element passes.
//
// [Full Topic]: https://developer.apple.com/documentation/PassKit/PKAddSecureElementPassError/Code
type PKAddSecureElementPassErrorCode uint

const (
	// PKAddSecureElementPassDeviceNotReadyError - The reader for the pass isn’t ready to start pairing.
	//
	// [Full Topic]: https://developer.apple.com/documentation/PassKit/PKAddSecureElementPassError/Code/deviceNotReadyError
	PKAddSecureElementPassDeviceNotReadyError PKAddSecureElementPassErrorCode = 0
	// PKAddSecureElementPassDeviceNotSupportedError - The reader for the pass isn’t supported or has an invalid version.
	//
	// [Full Topic]: https://developer.apple.com/documentation/PassKit/PKAddSecureElementPassError/Code/deviceNotSupportedError
	PKAddSecureElementPassDeviceNotSupportedError PKAddSecureElementPassErrorCode = 0
	// PKAddSecureElementPassGenericError - Represents the default error case.
	//
	// [Full Topic]: https://developer.apple.com/documentation/PassKit/PKAddSecureElementPassError/Code/genericError
	PKAddSecureElementPassGenericError PKAddSecureElementPassErrorCode = 0
	// PKAddSecureElementPassInvalidConfigurationError - The configuration for the pass is invalid for either Wallet or the reader.
	//
	// [Full Topic]: https://developer.apple.com/documentation/PassKit/PKAddSecureElementPassError/Code/invalidConfigurationError
	PKAddSecureElementPassInvalidConfigurationError PKAddSecureElementPassErrorCode = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/PassKit/PKAddSecureElementPassError/Code/osVersionNotSupportedError
	PKAddSecureElementPassOSVersionNotSupportedError PKAddSecureElementPassErrorCode = 0
	// PKAddSecureElementPassUnavailableError - Provisioning for secure element passes isn’t available on the device, or the app is missing the entitlement.
	//
	// [Full Topic]: https://developer.apple.com/documentation/PassKit/PKAddSecureElementPassError/Code/unavailableError
	PKAddSecureElementPassUnavailableError PKAddSecureElementPassErrorCode = 0
	// PKAddSecureElementPassUserCanceledError - The user canceled adding the pass.
	//
	// [Full Topic]: https://developer.apple.com/documentation/PassKit/PKAddSecureElementPassError/Code/userCanceledError
	PKAddSecureElementPassUserCanceledError PKAddSecureElementPassErrorCode = 0
	// PKAddSecureElementPassUnknownError - The system canceled adding the pass due to an unknown failure.
	//
	// [Full Topic]: https://developer.apple.com/documentation/PassKit/PKAddSecureElementPassErrorCode/PKAddSecureElementPassUnknownError
	PKAddSecureElementPassUnknownError PKAddSecureElementPassErrorCode = 0
)

/* debug [enums.gen.go]: Processing enum PKAddShareablePassConfigurationPrimaryAction (2 cases) */
// PKAddShareablePassConfigurationPrimaryAction - The kind of add action that the system performs with a pass.
//
// [Full Topic]: https://developer.apple.com/documentation/PassKit/PKAddShareablePassConfigurationPrimaryAction
type PKAddShareablePassConfigurationPrimaryAction uint

const (
	// PKAddShareablePassConfigurationPrimaryActionAdd - A constant that indicates the system adds a pass to a device.
	//
	// [Full Topic]: https://developer.apple.com/documentation/PassKit/PKAddShareablePassConfigurationPrimaryAction/add
	PKAddShareablePassConfigurationPrimaryActionAdd PKAddShareablePassConfigurationPrimaryAction = 0
	// PKAddShareablePassConfigurationPrimaryActionShare - A constant that indicates the system shares the pass with another user.
	//
	// [Full Topic]: https://developer.apple.com/documentation/PassKit/PKAddShareablePassConfigurationPrimaryAction/share
	PKAddShareablePassConfigurationPrimaryActionShare PKAddShareablePassConfigurationPrimaryAction = 0
)

/* debug [enums.gen.go]: Processing enum PKApplePayLaterAvailability (3 cases) */
// PKApplePayLaterAvailability - Values you use to enable or disable Apple Pay Later for a specific transaction.
//
// [Full Topic]: https://developer.apple.com/documentation/PassKit/PKApplePayLaterAvailability
type PKApplePayLaterAvailability uint

const (
	// PKApplePayLaterAvailable - Apple Pay Later is available.
	//
	// [Full Topic]: https://developer.apple.com/documentation/PassKit/PKApplePayLaterAvailability/available
	PKApplePayLaterAvailable PKApplePayLaterAvailability = 0
	// PKApplePayLaterUnavailableItemIneligible - Apple Pay Later is unavailable because one or more ineligible or prohibited items are in the shopping cart, such as gift cards.
	//
	// [Full Topic]: https://developer.apple.com/documentation/PassKit/PKApplePayLaterAvailability/unavailableItemIneligible
	PKApplePayLaterUnavailableItemIneligible PKApplePayLaterAvailability = 0
	// PKApplePayLaterUnavailableRecurringTransaction - Apple Pay Later is unavailable because there’s a recurring payment or subscription in the shopping cart.
	//
	// [Full Topic]: https://developer.apple.com/documentation/PassKit/PKApplePayLaterAvailability/unavailableRecurringTransaction
	PKApplePayLaterUnavailableRecurringTransaction PKApplePayLaterAvailability = 0
)

/* debug [enums.gen.go]: Processing enum PKAutomaticPassPresentationSuppressionResult (5 cases) */
// PKAutomaticPassPresentationSuppressionResult - The result of an attempt to suppress automatic pass presentation.
//
// [Full Topic]: https://developer.apple.com/documentation/PassKit/PKAutomaticPassPresentationSuppressionResult
type PKAutomaticPassPresentationSuppressionResult uint

const (
	// PKAutomaticPassPresentationSuppressionResultAlreadyPresenting - The device is already presenting passes.
	//
	// [Full Topic]: https://developer.apple.com/documentation/PassKit/PKAutomaticPassPresentationSuppressionResult/alreadyPresenting
	PKAutomaticPassPresentationSuppressionResultAlreadyPresenting PKAutomaticPassPresentationSuppressionResult = 0
	// PKAutomaticPassPresentationSuppressionResultCancelled - The system canceled the suppression before calling the response handler.
	//
	// [Full Topic]: https://developer.apple.com/documentation/PassKit/PKAutomaticPassPresentationSuppressionResult/cancelled
	PKAutomaticPassPresentationSuppressionResultCancelled PKAutomaticPassPresentationSuppressionResult = 0
	// PKAutomaticPassPresentationSuppressionResultDenied - The user prevented the suppression, or an internal error occurred.
	//
	// [Full Topic]: https://developer.apple.com/documentation/PassKit/PKAutomaticPassPresentationSuppressionResult/denied
	PKAutomaticPassPresentationSuppressionResultDenied PKAutomaticPassPresentationSuppressionResult = 0
	// PKAutomaticPassPresentationSuppressionResultNotSupported - The device doesn’t support the suppression of automatic pass presentation.
	//
	// [Full Topic]: https://developer.apple.com/documentation/PassKit/PKAutomaticPassPresentationSuppressionResult/notSupported
	PKAutomaticPassPresentationSuppressionResultNotSupported PKAutomaticPassPresentationSuppressionResult = 0
	// PKAutomaticPassPresentationSuppressionResultSuccess - Suppression of automatic presentation successful.
	//
	// [Full Topic]: https://developer.apple.com/documentation/PassKit/PKAutomaticPassPresentationSuppressionResult/success
	PKAutomaticPassPresentationSuppressionResultSuccess PKAutomaticPassPresentationSuppressionResult = 0
)

/* debug [enums.gen.go]: Processing enum PKBarcodeEventConfigurationDataType (3 cases) */
// PKBarcodeEventConfigurationDataType enum type
//
// [Full Topic]: https://developer.apple.com/documentation/PassKit/PKBarcodeEventConfigurationDataType
type PKBarcodeEventConfigurationDataType uint

const (
	//
	// [Full Topic]: https://developer.apple.com/documentation/PassKit/PKBarcodeEventConfigurationDataType/signingCertificate
	PKBarcodeEventConfigurationDataTypeSigningCertificate PKBarcodeEventConfigurationDataType = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/PassKit/PKBarcodeEventConfigurationDataType/signingKeyMaterial
	PKBarcodeEventConfigurationDataTypeSigningKeyMaterial PKBarcodeEventConfigurationDataType = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/PassKit/PKBarcodeEventConfigurationDataType/unknown
	PKBarcodeEventConfigurationDataTypeUnknown PKBarcodeEventConfigurationDataType = 0
)

/* debug [enums.gen.go]: Processing enum PKDisbursementErrorCode (3 cases) */
// PKDisbursementErrorCode - Values that describe errors that can occur while processing the disbursement.
//
// [Full Topic]: https://developer.apple.com/documentation/PassKit/PKDisbursementError/Code
type PKDisbursementErrorCode uint

const (
	// PKDisbursementRecipientContactInvalidError - The recipient’s contact information wasn’t valid.
	//
	// [Full Topic]: https://developer.apple.com/documentation/PassKit/PKDisbursementError/Code/recipientContactInvalidError
	PKDisbursementRecipientContactInvalidError PKDisbursementErrorCode = 0
	// PKDisbursementUnknownError - An unknown error occurred.
	//
	// [Full Topic]: https://developer.apple.com/documentation/PassKit/PKDisbursementError/Code/unknownError
	PKDisbursementUnknownError PKDisbursementErrorCode = 0
	// PKDisbursementUnsupportedCardError - The framework doesn’t support the card the individual presented.
	//
	// [Full Topic]: https://developer.apple.com/documentation/PassKit/PKDisbursementError/Code/unsupportedCardError
	PKDisbursementUnsupportedCardError PKDisbursementErrorCode = 0
)

/* debug [enums.gen.go]: Processing enum PKIdentityButtonLabel (4 cases) */
// PKIdentityButtonLabel - A type that indicates the available labels for an identity button.
//
// [Full Topic]: https://developer.apple.com/documentation/PassKit/PKIdentityButton/Label
type PKIdentityButtonLabel uint

const (
	// PKIdentityButtonLabelContinue - A label with the text continue with Apple Wallet.
	//
	// [Full Topic]: https://developer.apple.com/documentation/PassKit/PKIdentityButton/Label/continue
	PKIdentityButtonLabelContinue PKIdentityButtonLabel = 0
	// PKIdentityButtonLabelVerify - A label with the text verify with Apple Wallet.
	//
	// [Full Topic]: https://developer.apple.com/documentation/PassKit/PKIdentityButton/Label/verify
	PKIdentityButtonLabelVerify PKIdentityButtonLabel = 0
	// PKIdentityButtonLabelVerifyAge - A label with the text verify age with Apple Wallet.
	//
	// [Full Topic]: https://developer.apple.com/documentation/PassKit/PKIdentityButton/Label/verifyAge
	PKIdentityButtonLabelVerifyAge PKIdentityButtonLabel = 0
	// PKIdentityButtonLabelVerifyIdentity - A label with the text verify identity with Apple Wallet.
	//
	// [Full Topic]: https://developer.apple.com/documentation/PassKit/PKIdentityButton/Label/verifyIdentity
	PKIdentityButtonLabelVerifyIdentity PKIdentityButtonLabel = 0
)

/* debug [enums.gen.go]: Processing enum PKIdentityButtonStyle (2 cases) */
// PKIdentityButtonStyle - A type that indicates the available appearances for an identity button.
//
// [Full Topic]: https://developer.apple.com/documentation/PassKit/PKIdentityButton/Style
type PKIdentityButtonStyle uint

const (
	// PKIdentityButtonStyleBlack - A style that represents a black button with white lettering.
	//
	// [Full Topic]: https://developer.apple.com/documentation/PassKit/PKIdentityButton/Style/black
	PKIdentityButtonStyleBlack PKIdentityButtonStyle = 0
	// PKIdentityButtonStyleBlackOutline - A style that represents a black button with a light outline.
	//
	// [Full Topic]: https://developer.apple.com/documentation/PassKit/PKIdentityButton/Style/blackOutline
	PKIdentityButtonStyleBlackOutline PKIdentityButtonStyle = 0
)

/* debug [enums.gen.go]: Processing enum PKIdentityError (9 cases) */
// PKIdentityError - Error codes for identity operations.
//
// [Full Topic]: https://developer.apple.com/documentation/PassKit/PKIdentityError-swift.struct/Code
type PKIdentityError uint

const (
	// PKIdentityErrorCancelled - An error that indicates the user cancels the presented sheet.
	//
	// [Full Topic]: https://developer.apple.com/documentation/PassKit/PKIdentityError-swift.struct/Code/cancelled
	PKIdentityErrorCancelled PKIdentityError = 0
	// PKIdentityErrorInvalidElement - An error that indicates an element the app requests isn’t valid.
	//
	// [Full Topic]: https://developer.apple.com/documentation/PassKit/PKIdentityError-swift.struct/Code/invalidElement
	PKIdentityErrorInvalidElement PKIdentityError = 0
	// PKIdentityErrorInvalidNonce - An error that indicates the number is too large or unsuitable.
	//
	// [Full Topic]: https://developer.apple.com/documentation/PassKit/PKIdentityError-swift.struct/Code/invalidNonce
	PKIdentityErrorInvalidNonce PKIdentityError = 0
	// PKIdentityErrorNetworkUnavailable - An error that indicates a network isn’t available.
	//
	// [Full Topic]: https://developer.apple.com/documentation/PassKit/PKIdentityError-swift.struct/Code/networkUnavailable
	PKIdentityErrorNetworkUnavailable PKIdentityError = 0
	// PKIdentityErrorNoElementsRequested - An error that indicates the elements aren’t supported.
	//
	// [Full Topic]: https://developer.apple.com/documentation/PassKit/PKIdentityError-swift.struct/Code/noElementsRequested
	PKIdentityErrorNoElementsRequested PKIdentityError = 0
	// PKIdentityErrorNotSupported - An error that indicates the request originates from a device the framework doesn’t support.
	//
	// [Full Topic]: https://developer.apple.com/documentation/PassKit/PKIdentityError-swift.struct/Code/notSupported
	PKIdentityErrorNotSupported PKIdentityError = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/PassKit/PKIdentityError-swift.struct/Code/regionNotSupported
	PKIdentityErrorRegionNotSupported PKIdentityError = 0
	// PKIdentityErrorRequestAlreadyInProgress - An error that indicates a request is already in progress.
	//
	// [Full Topic]: https://developer.apple.com/documentation/PassKit/PKIdentityError-swift.struct/Code/requestAlreadyInProgress
	PKIdentityErrorRequestAlreadyInProgress PKIdentityError = 0
	// PKIdentityErrorUnknown - An error that indicates an unknown error.
	//
	// [Full Topic]: https://developer.apple.com/documentation/PassKit/PKIdentityError-swift.struct/Code/unknown
	PKIdentityErrorUnknown PKIdentityError = 0
)

/* debug [enums.gen.go]: Processing enum PKIssuerProvisioningExtensionAuthorizationResult (2 cases) */
// PKIssuerProvisioningExtensionAuthorizationResult - A value that indicates the result of authorizing the addition of a payment card.
//
// [Full Topic]: https://developer.apple.com/documentation/PassKit/PKIssuerProvisioningExtensionAuthorizationResult
type PKIssuerProvisioningExtensionAuthorizationResult uint

const (
	// PKIssuerProvisioningExtensionAuthorizationResultAuthorized - A result that indicates the user successfully authorized adding the payment pass.
	//
	// [Full Topic]: https://developer.apple.com/documentation/PassKit/PKIssuerProvisioningExtensionAuthorizationResult/authorized
	PKIssuerProvisioningExtensionAuthorizationResultAuthorized PKIssuerProvisioningExtensionAuthorizationResult = 0
	// PKIssuerProvisioningExtensionAuthorizationResultCanceled - A result that indicates the user canceled authorization or wasn’t authorized to add the payment card.
	//
	// [Full Topic]: https://developer.apple.com/documentation/PassKit/PKIssuerProvisioningExtensionAuthorizationResult/canceled
	PKIssuerProvisioningExtensionAuthorizationResultCanceled PKIssuerProvisioningExtensionAuthorizationResult = 0
)

/* debug [enums.gen.go]: Processing enum PKMerchantCapability (5 cases) */
// PKMerchantCapability - Capabilities for processing payment.
//
// [Full Topic]: https://developer.apple.com/documentation/PassKit/PKMerchantCapability
type PKMerchantCapability uint

const (
	// PKMerchantCapabilityCredit - Support for credit cards.
	//
	// [Full Topic]: https://developer.apple.com/documentation/PassKit/PKMerchantCapability/credit
	PKMerchantCapabilityCredit PKMerchantCapability = 0
	// PKMerchantCapabilityDebit - Support for debit cards.
	//
	// [Full Topic]: https://developer.apple.com/documentation/PassKit/PKMerchantCapability/debit
	PKMerchantCapabilityDebit PKMerchantCapability = 0
	// PKMerchantCapabilityEMV - Support for the EMV protocol.
	//
	// [Full Topic]: https://developer.apple.com/documentation/PassKit/PKMerchantCapability/emv
	PKMerchantCapabilityEMV PKMerchantCapability = 0
	// PKMerchantCapabilityInstantFundsOut - The value that indicates the merchant supports disbursing funds using Instant Funds Out.
	//
	// [Full Topic]: https://developer.apple.com/documentation/PassKit/PKMerchantCapability/instantFundsOut
	PKMerchantCapabilityInstantFundsOut PKMerchantCapability = 0
	// PKMerchantCapability3DS - Support for the 3-D Secure protocol.
	//
	// [Full Topic]: https://developer.apple.com/documentation/PassKit/PKMerchantCapability/threeDSecure
	PKMerchantCapability3DS PKMerchantCapability = 0
)

/* debug [enums.gen.go]: Processing enum PKPassKitErrorCode (5 cases) */
// PKPassKitErrorCode - Errors that the PassKit framework uses.
//
// [Full Topic]: https://developer.apple.com/documentation/PassKit/PKPassKitError/Code
type PKPassKitErrorCode uint

const (
	// PKInvalidDataError - Invalid pass data.
	//
	// [Full Topic]: https://developer.apple.com/documentation/PassKit/PKPassKitError/Code/invalidDataError
	PKInvalidDataError PKPassKitErrorCode = 0
	// PKInvalidSignature - Invalid pass signature.
	//
	// [Full Topic]: https://developer.apple.com/documentation/PassKit/PKPassKitError/Code/invalidSignature
	PKInvalidSignature PKPassKitErrorCode = 0
	// PKNotEntitledError - Error caused by absence of the required entitlements for the given operation.
	//
	// [Full Topic]: https://developer.apple.com/documentation/PassKit/PKPassKitError/Code/notEntitledError
	PKNotEntitledError PKPassKitErrorCode = 0
	// PKUnknownError - Unknown error.
	//
	// [Full Topic]: https://developer.apple.com/documentation/PassKit/PKPassKitError/Code/unknownError
	PKUnknownError PKPassKitErrorCode = 0
	// PKUnsupportedVersionError - Unsupported pass version.
	//
	// [Full Topic]: https://developer.apple.com/documentation/PassKit/PKPassKitError/Code/unsupportedVersionError
	PKUnsupportedVersionError PKPassKitErrorCode = 0
)

/* debug [enums.gen.go]: Processing enum PKPassLibraryAddPassesStatus (3 cases) */
// PKPassLibraryAddPassesStatus - Statuses that PassKit uses when it adds passes to the pass library.
//
// [Full Topic]: https://developer.apple.com/documentation/PassKit/PKPassLibraryAddPassesStatus
type PKPassLibraryAddPassesStatus uint

const (
	// PKPassLibraryDidAddPasses - A status that occurs when the user successfully adds one or more passes.
	//
	// [Full Topic]: https://developer.apple.com/documentation/PassKit/PKPassLibraryAddPassesStatus/didAddPasses
	PKPassLibraryDidAddPasses PKPassLibraryAddPassesStatus = 0
	// PKPassLibraryDidCancelAddPasses - A status that occurs when the user cancels the addition of passes.
	//
	// [Full Topic]: https://developer.apple.com/documentation/PassKit/PKPassLibraryAddPassesStatus/didCancelAddPasses
	PKPassLibraryDidCancelAddPasses PKPassLibraryAddPassesStatus = 0
	// PKPassLibraryShouldReviewPasses - A status that occurs when the app prompts the user to review the passes.
	//
	// [Full Topic]: https://developer.apple.com/documentation/PassKit/PKPassLibraryAddPassesStatus/shouldReviewPasses
	PKPassLibraryShouldReviewPasses PKPassLibraryAddPassesStatus = 0
)

/* debug [enums.gen.go]: Processing enum PKPassType (4 cases) */
// PKPassType - Types of passes.
//
// [Full Topic]: https://developer.apple.com/documentation/PassKit/PKPassType
type PKPassType uint

const (
	// PKPassTypeAny - A nonspecific pass type.
	//
	// [Full Topic]: https://developer.apple.com/documentation/PassKit/PKPassType/any
	PKPassTypeAny PKPassType = 0
	// PKPassTypeBarcode - A pass that represents a barcode.
	//
	// [Full Topic]: https://developer.apple.com/documentation/PassKit/PKPassType/barcode
	PKPassTypeBarcode PKPassType = 0
	// PKPassTypePayment - A pass that represents a credit or debit card
	//
	// [Full Topic]: https://developer.apple.com/documentation/PassKit/PKPassType/payment
	PKPassTypePayment PKPassType = 0
	// PKPassTypeSecureElement - A pass that represents a credential that the device stores in the Secure Element.
	//
	// [Full Topic]: https://developer.apple.com/documentation/PassKit/PKPassType/secureElement
	PKPassTypeSecureElement PKPassType = 0
)

/* debug [enums.gen.go]: Processing enum PKPayLaterAction (2 cases) */
// PKPayLaterAction - Values you use to set the Apple Pay Later action.
//
// [Full Topic]: https://developer.apple.com/documentation/PassKit/PKPayLaterAction
type PKPayLaterAction uint

const (
	// PKPayLaterActionCalculator - An action the provides the standard price breakdown calculator for Apple Pay Later.
	//
	// [Full Topic]: https://developer.apple.com/documentation/PassKit/PKPayLaterAction/calculator
	PKPayLaterActionCalculator PKPayLaterAction = 0
	// PKPayLaterActionLearnMore - An action that displays a button that a person can tap to learn more about Apple Pay Later.
	//
	// [Full Topic]: https://developer.apple.com/documentation/PassKit/PKPayLaterAction/learnMore
	PKPayLaterActionLearnMore PKPayLaterAction = 0
)

/* debug [enums.gen.go]: Processing enum PKPayLaterDisplayStyle (4 cases) */
// PKPayLaterDisplayStyle - Values you use to style an Apple Pay Later visual merchandising widget.
//
// [Full Topic]: https://developer.apple.com/documentation/PassKit/PKPayLaterDisplayStyle
type PKPayLaterDisplayStyle uint

const (
	// PKPayLaterDisplayStyleBadge - Displays a badge that uses an Apple Pay Later icon.
	//
	// [Full Topic]: https://developer.apple.com/documentation/PassKit/PKPayLaterDisplayStyle/badge
	PKPayLaterDisplayStyleBadge PKPayLaterDisplayStyle = 0
	// PKPayLaterDisplayStyleCheckout - A style used inside of a checkout view that presents Apple Pay and other payment options to the customer.
	//
	// [Full Topic]: https://developer.apple.com/documentation/PassKit/PKPayLaterDisplayStyle/checkout
	PKPayLaterDisplayStyleCheckout PKPayLaterDisplayStyle = 0
	// PKPayLaterDisplayStylePrice - A style that shows the Apple Pay Later view beneath a product’s price.
	//
	// [Full Topic]: https://developer.apple.com/documentation/PassKit/PKPayLaterDisplayStyle/price
	PKPayLaterDisplayStylePrice PKPayLaterDisplayStyle = 0
	// PKPayLaterDisplayStyleStandard - The standard Apple Pay Later visual merchandising widget style.
	//
	// [Full Topic]: https://developer.apple.com/documentation/PassKit/PKPayLaterDisplayStyle/standard
	PKPayLaterDisplayStyleStandard PKPayLaterDisplayStyle = 0
)

/* debug [enums.gen.go]: Processing enum PKPaymentAuthorizationStatus (8 cases) */
// PKPaymentAuthorizationStatus - General success and failure status for payment authorization.
//
// [Full Topic]: https://developer.apple.com/documentation/PassKit/PKPaymentAuthorizationStatus
type PKPaymentAuthorizationStatus uint

const (
	// PKPaymentAuthorizationStatusFailure - Merchant failed to authorize the transaction.
	//
	// [Full Topic]: https://developer.apple.com/documentation/PassKit/PKPaymentAuthorizationStatus/failure
	PKPaymentAuthorizationStatusFailure PKPaymentAuthorizationStatus = 0
	// PKPaymentAuthorizationStatusInvalidBillingPostalAddress - Invalid or unusable billing address.
	//
	// [Full Topic]: https://developer.apple.com/documentation/PassKit/PKPaymentAuthorizationStatus/invalidBillingPostalAddress
	PKPaymentAuthorizationStatusInvalidBillingPostalAddress PKPaymentAuthorizationStatus = 0
	// PKPaymentAuthorizationStatusInvalidShippingContact - Invalid or incomplete shipping contact.
	//
	// [Full Topic]: https://developer.apple.com/documentation/PassKit/PKPaymentAuthorizationStatus/invalidShippingContact
	PKPaymentAuthorizationStatusInvalidShippingContact PKPaymentAuthorizationStatus = 0
	// PKPaymentAuthorizationStatusInvalidShippingPostalAddress - Invalid or unusable shipping address.
	//
	// [Full Topic]: https://developer.apple.com/documentation/PassKit/PKPaymentAuthorizationStatus/invalidShippingPostalAddress
	PKPaymentAuthorizationStatusInvalidShippingPostalAddress PKPaymentAuthorizationStatus = 0
	// PKPaymentAuthorizationStatusPINIncorrect - Incorrect PIN entered.
	//
	// [Full Topic]: https://developer.apple.com/documentation/PassKit/PKPaymentAuthorizationStatus/pinIncorrect
	PKPaymentAuthorizationStatusPINIncorrect PKPaymentAuthorizationStatus = 0
	// PKPaymentAuthorizationStatusPINLockout - PIN retry limit exceeded.
	//
	// [Full Topic]: https://developer.apple.com/documentation/PassKit/PKPaymentAuthorizationStatus/pinLockout
	PKPaymentAuthorizationStatusPINLockout PKPaymentAuthorizationStatus = 0
	// PKPaymentAuthorizationStatusPINRequired - Transaction requires PIN entry.
	//
	// [Full Topic]: https://developer.apple.com/documentation/PassKit/PKPaymentAuthorizationStatus/pinRequired
	PKPaymentAuthorizationStatusPINRequired PKPaymentAuthorizationStatus = 0
	// PKPaymentAuthorizationStatusSuccess - Merchant successfully authorized the transaction, or the transaction is expected to succeed.
	//
	// [Full Topic]: https://developer.apple.com/documentation/PassKit/PKPaymentAuthorizationStatus/success
	PKPaymentAuthorizationStatusSuccess PKPaymentAuthorizationStatus = 0
)

/* debug [enums.gen.go]: Processing enum PKPaymentButtonStyle (4 cases) */
// PKPaymentButtonStyle - A type that indicates the available appearances for an Apple Pay button.
//
// [Full Topic]: https://developer.apple.com/documentation/PassKit/PKPaymentButtonStyle
type PKPaymentButtonStyle uint

const (
	// PKPaymentButtonStyleAutomatic - A button that automatically changes its appearance when the user switches between Light Mode and Dark Mode.
	//
	// [Full Topic]: https://developer.apple.com/documentation/PassKit/PKPaymentButtonStyle/automatic
	PKPaymentButtonStyleAutomatic PKPaymentButtonStyle = 0
	// PKPaymentButtonStyleBlack - A black button with white lettering.
	//
	// [Full Topic]: https://developer.apple.com/documentation/PassKit/PKPaymentButtonStyle/black
	PKPaymentButtonStyleBlack PKPaymentButtonStyle = 0
	// PKPaymentButtonStyleWhite - A white button with black lettering.
	//
	// [Full Topic]: https://developer.apple.com/documentation/PassKit/PKPaymentButtonStyle/white
	PKPaymentButtonStyleWhite PKPaymentButtonStyle = 0
	// PKPaymentButtonStyleWhiteOutline - A white button with black lettering and a black outline.
	//
	// [Full Topic]: https://developer.apple.com/documentation/PassKit/PKPaymentButtonStyle/whiteOutline
	PKPaymentButtonStyleWhiteOutline PKPaymentButtonStyle = 0
)

/* debug [enums.gen.go]: Processing enum PKPaymentButtonType (17 cases) */
// PKPaymentButtonType - The Apple Pay button types you can display to initiate Apple Pay transactions.
//
// [Full Topic]: https://developer.apple.com/documentation/PassKit/PKPaymentButtonType
type PKPaymentButtonType uint

const (
	// PKPaymentButtonTypeAddMoney - An Apple Pay button useful for adding money to a card, account, or payment system.
	//
	// [Full Topic]: https://developer.apple.com/documentation/PassKit/PKPaymentButtonType/addMoney
	PKPaymentButtonTypeAddMoney PKPaymentButtonType = 0
	// PKPaymentButtonTypeBook - An Apple Pay button useful for booking trips, flights, or other experiences.
	//
	// [Full Topic]: https://developer.apple.com/documentation/PassKit/PKPaymentButtonType/book
	PKPaymentButtonTypeBook PKPaymentButtonType = 0
	// PKPaymentButtonTypeBuy - An Apple Pay button useful for product purchases.
	//
	// [Full Topic]: https://developer.apple.com/documentation/PassKit/PKPaymentButtonType/buy
	PKPaymentButtonTypeBuy PKPaymentButtonType = 0
	// PKPaymentButtonTypeCheckout - An Apple Pay button useful for purchase experiences that include other payment buttons that start with “Check out”.
	//
	// [Full Topic]: https://developer.apple.com/documentation/PassKit/PKPaymentButtonType/checkout
	PKPaymentButtonTypeCheckout PKPaymentButtonType = 0
	// PKPaymentButtonTypeContinue - An Apple Pay button useful for general purchases.
	//
	// [Full Topic]: https://developer.apple.com/documentation/PassKit/PKPaymentButtonType/continue
	PKPaymentButtonTypeContinue PKPaymentButtonType = 0
	// PKPaymentButtonTypeContribute - An Apple Pay button useful to help people contribute money to projects, causes, organizations, and other entities.
	//
	// [Full Topic]: https://developer.apple.com/documentation/PassKit/PKPaymentButtonType/contribute
	PKPaymentButtonTypeContribute PKPaymentButtonType = 0
	// PKPaymentButtonTypeDonate - An Apple Pay button used by approved nonprofit organization that lets people make donations.
	//
	// [Full Topic]: https://developer.apple.com/documentation/PassKit/PKPaymentButtonType/donate
	PKPaymentButtonTypeDonate PKPaymentButtonType = 0
	// PKPaymentButtonTypeInStore - An Apple Pay button useful for paying bills or invoices.
	//
	// [Full Topic]: https://developer.apple.com/documentation/PassKit/PKPaymentButtonType/inStore
	PKPaymentButtonTypeInStore PKPaymentButtonType = 0
	// PKPaymentButtonTypeOrder - An Apple Pay button useful for placing orders for such as like meals or flowers.
	//
	// [Full Topic]: https://developer.apple.com/documentation/PassKit/PKPaymentButtonType/order
	PKPaymentButtonTypeOrder PKPaymentButtonType = 0
	// PKPaymentButtonTypePlain - An Apple Pay button with the Apple Pay logo only, useful when an additional call to action isn’t needed.
	//
	// [Full Topic]: https://developer.apple.com/documentation/PassKit/PKPaymentButtonType/plain
	PKPaymentButtonTypePlain PKPaymentButtonType = 0
	// PKPaymentButtonTypeReload - An Apple Pay button useful for adding money to a card, account, or payment system.
	//
	// [Full Topic]: https://developer.apple.com/documentation/PassKit/PKPaymentButtonType/reload
	PKPaymentButtonTypeReload PKPaymentButtonType = 0
	// PKPaymentButtonTypeRent - An Apple Pay button useful for renting items such as cars or scooters.
	//
	// [Full Topic]: https://developer.apple.com/documentation/PassKit/PKPaymentButtonType/rent
	PKPaymentButtonTypeRent PKPaymentButtonType = 0
	// PKPaymentButtonTypeSetUp - An Apple Pay button useful for prompting the user to set up a card.
	//
	// [Full Topic]: https://developer.apple.com/documentation/PassKit/PKPaymentButtonType/setUp
	PKPaymentButtonTypeSetUp PKPaymentButtonType = 0
	// PKPaymentButtonTypeSubscribe - An Apple Pay button useful for purchasing a subscription such as a gym membership or meal-kit delivery service.
	//
	// [Full Topic]: https://developer.apple.com/documentation/PassKit/PKPaymentButtonType/subscribe
	PKPaymentButtonTypeSubscribe PKPaymentButtonType = 0
	// PKPaymentButtonTypeSupport - An Apple Pay button useful supporting people give money to projects, causes, organizations, and other entities.
	//
	// [Full Topic]: https://developer.apple.com/documentation/PassKit/PKPaymentButtonType/support
	PKPaymentButtonTypeSupport PKPaymentButtonType = 0
	// PKPaymentButtonTypeTip - An Apple Pay button useful useful for letting people tip for goods or services.
	//
	// [Full Topic]: https://developer.apple.com/documentation/PassKit/PKPaymentButtonType/tip
	PKPaymentButtonTypeTip PKPaymentButtonType = 0
	// PKPaymentButtonTypeTopUp - An Apple Pay button useful for adding money to a card, account, or payment system.
	//
	// [Full Topic]: https://developer.apple.com/documentation/PassKit/PKPaymentButtonType/topUp
	PKPaymentButtonTypeTopUp PKPaymentButtonType = 0
)

/* debug [enums.gen.go]: Processing enum PKPaymentErrorCode (6 cases) */
// PKPaymentErrorCode - An error code that you provide to indicate problems with address or contact information on an Apple Pay sheet.
//
// [Full Topic]: https://developer.apple.com/documentation/PassKit/PKPaymentError/Code
type PKPaymentErrorCode uint

const (
	// PKPaymentBillingContactInvalidError - The error code that indicates an invalid billing address or billing name.
	//
	// [Full Topic]: https://developer.apple.com/documentation/PassKit/PKPaymentError/Code/billingContactInvalidError
	PKPaymentBillingContactInvalidError PKPaymentErrorCode = 0
	// PKPaymentCouponCodeExpiredError - The error code that indicates an expired coupon.
	//
	// [Full Topic]: https://developer.apple.com/documentation/PassKit/PKPaymentError/Code/couponCodeExpiredError
	PKPaymentCouponCodeExpiredError PKPaymentErrorCode = 0
	// PKPaymentCouponCodeInvalidError - The error code that indicates an invalid coupon.
	//
	// [Full Topic]: https://developer.apple.com/documentation/PassKit/PKPaymentError/Code/couponCodeInvalidError
	PKPaymentCouponCodeInvalidError PKPaymentErrorCode = 0
	// PKPaymentShippingAddressUnserviceableError - The error code that indicates an unserviceable shipping address.
	//
	// [Full Topic]: https://developer.apple.com/documentation/PassKit/PKPaymentError/Code/shippingAddressUnserviceableError
	PKPaymentShippingAddressUnserviceableError PKPaymentErrorCode = 0
	// PKPaymentShippingContactInvalidError - The error code that indicates an invalid shipping address, email, phone, or name.
	//
	// [Full Topic]: https://developer.apple.com/documentation/PassKit/PKPaymentError/Code/shippingContactInvalidError
	PKPaymentShippingContactInvalidError PKPaymentErrorCode = 0
	// PKPaymentUnknownError - The error code that indicates an unknown error.
	//
	// [Full Topic]: https://developer.apple.com/documentation/PassKit/PKPaymentError/Code/unknownError
	PKPaymentUnknownError PKPaymentErrorCode = 0
)

/* debug [enums.gen.go]: Processing enum PKPaymentMethodType (6 cases) */
// PKPaymentMethodType - The type of cards available in Apple Pay.
//
// [Full Topic]: https://developer.apple.com/documentation/PassKit/PKPaymentMethodType
type PKPaymentMethodType uint

const (
	// PKPaymentMethodTypeCredit - A credit card.
	//
	// [Full Topic]: https://developer.apple.com/documentation/PassKit/PKPaymentMethodType/credit
	PKPaymentMethodTypeCredit PKPaymentMethodType = 0
	// PKPaymentMethodTypeDebit - A debit card.
	//
	// [Full Topic]: https://developer.apple.com/documentation/PassKit/PKPaymentMethodType/debit
	PKPaymentMethodTypeDebit PKPaymentMethodType = 0
	// PKPaymentMethodTypeEMoney - An electronic money card.
	//
	// [Full Topic]: https://developer.apple.com/documentation/PassKit/PKPaymentMethodType/eMoney
	PKPaymentMethodTypeEMoney PKPaymentMethodType = 0
	// PKPaymentMethodTypePrepaid - A prepaid card.
	//
	// [Full Topic]: https://developer.apple.com/documentation/PassKit/PKPaymentMethodType/prepaid
	PKPaymentMethodTypePrepaid PKPaymentMethodType = 0
	// PKPaymentMethodTypeStore - A store card.
	//
	// [Full Topic]: https://developer.apple.com/documentation/PassKit/PKPaymentMethodType/store
	PKPaymentMethodTypeStore PKPaymentMethodType = 0
	// PKPaymentMethodTypeUnknown - The card’s type is unknown.
	//
	// [Full Topic]: https://developer.apple.com/documentation/PassKit/PKPaymentMethodType/unknown
	PKPaymentMethodTypeUnknown PKPaymentMethodType = 0
)

/* debug [enums.gen.go]: Processing enum PKPaymentPassActivationState (5 cases) */
// PKPaymentPassActivationState - Cases that indicate payment pass activation states.
//
// [Full Topic]: https://developer.apple.com/documentation/PassKit/PKPaymentPassActivationState
type PKPaymentPassActivationState uint

const (
	// PKPaymentPassActivationStateActivated - Active and ready for payment use.
	//
	// [Full Topic]: https://developer.apple.com/documentation/PassKit/PKPaymentPassActivationState/activated
	PKPaymentPassActivationStateActivated PKPaymentPassActivationState = 0
	// PKPaymentPassActivationStateActivating - Not ready for use but activation is in progress.
	//
	// [Full Topic]: https://developer.apple.com/documentation/PassKit/PKPaymentPassActivationState/activating
	PKPaymentPassActivationStateActivating PKPaymentPassActivationState = 0
	// PKPaymentPassActivationStateDeactivated - Not active because the issuer disabled the account associated with the device.
	//
	// [Full Topic]: https://developer.apple.com/documentation/PassKit/PKPaymentPassActivationState/deactivated
	PKPaymentPassActivationStateDeactivated PKPaymentPassActivationState = 0
	// PKPaymentPassActivationStateRequiresActivation - Not active but may be activated by the issuer.
	//
	// [Full Topic]: https://developer.apple.com/documentation/PassKit/PKPaymentPassActivationState/requiresActivation
	PKPaymentPassActivationStateRequiresActivation PKPaymentPassActivationState = 0
	// PKPaymentPassActivationStateSuspended - Not active and can’t be activated.
	//
	// [Full Topic]: https://developer.apple.com/documentation/PassKit/PKPaymentPassActivationState/suspended
	PKPaymentPassActivationStateSuspended PKPaymentPassActivationState = 0
)

/* debug [enums.gen.go]: Processing enum PKPaymentSummaryItemType (2 cases) */
// PKPaymentSummaryItemType - Constants that describe the type of the payment summary item, such as final or pending.
//
// [Full Topic]: https://developer.apple.com/documentation/PassKit/PKPaymentSummaryItemType
type PKPaymentSummaryItemType uint

const (
	// PKPaymentSummaryItemTypeFinal - A summary item that represents a known, final cost.
	//
	// [Full Topic]: https://developer.apple.com/documentation/PassKit/PKPaymentSummaryItemType/final
	PKPaymentSummaryItemTypeFinal PKPaymentSummaryItemType = 0
	// PKPaymentSummaryItemTypePending - A summary item that represents an estimated or unknown cost.
	//
	// [Full Topic]: https://developer.apple.com/documentation/PassKit/PKPaymentSummaryItemType/pending
	PKPaymentSummaryItemTypePending PKPaymentSummaryItemType = 0
)

/* debug [enums.gen.go]: Processing enum PKRadioTechnology (3 cases) */
// PKRadioTechnology - Constants that describe the type of wireless radio technology that a pass uses.
//
// [Full Topic]: https://developer.apple.com/documentation/PassKit/PKRadioTechnology
type PKRadioTechnology uint

const (
	// PKRadioTechnologyBluetooth - An identifier that indicates the Bluetooth radio frequency communication technology.
	//
	// [Full Topic]: https://developer.apple.com/documentation/PassKit/PKRadioTechnology/bluetooth
	PKRadioTechnologyBluetooth PKRadioTechnology = 0
	// PKRadioTechnologyNFC - An identifier that indicates the near field communication (NFC) radio frequency communication technology.
	//
	// [Full Topic]: https://developer.apple.com/documentation/PassKit/PKRadioTechnology/NFC
	PKRadioTechnologyNFC PKRadioTechnology = 0
	// PKRadioTechnologyNone - An identifier that indicates the pass doesn’t use radio frequency communication.
	//
	// [Full Topic]: https://developer.apple.com/documentation/PassKit/PKRadioTechnology/PKRadioTechnologyNone
	PKRadioTechnologyNone PKRadioTechnology = 0
)

/* debug [enums.gen.go]: Processing enum PKSecureElementPassActivationState (5 cases) */
// PKSecureElementPassActivationState - The activation states of a Secure Element pass.
//
// [Full Topic]: https://developer.apple.com/documentation/PassKit/PKSecureElementPass/PassActivationState-swift.enum
type PKSecureElementPassActivationState uint

const (
	// PKSecureElementPassActivationStateActivated - The pass is active and ready to use.
	//
	// [Full Topic]: https://developer.apple.com/documentation/PassKit/PKSecureElementPass/PassActivationState-swift.enum/activated
	PKSecureElementPassActivationStateActivated PKSecureElementPassActivationState = 0
	// PKSecureElementPassActivationStateActivating - The pass isn’t ready to use, but activation is in progress
	//
	// [Full Topic]: https://developer.apple.com/documentation/PassKit/PKSecureElementPass/PassActivationState-swift.enum/activating
	PKSecureElementPassActivationStateActivating PKSecureElementPassActivationState = 0
	// PKSecureElementPassActivationStateDeactivated - The issuer has deactivated the pass.
	//
	// [Full Topic]: https://developer.apple.com/documentation/PassKit/PKSecureElementPass/PassActivationState-swift.enum/deactivated
	PKSecureElementPassActivationStateDeactivated PKSecureElementPassActivationState = 0
	// PKSecureElementPassActivationStateRequiresActivation - The pass requires activation by the issuer.
	//
	// [Full Topic]: https://developer.apple.com/documentation/PassKit/PKSecureElementPass/PassActivationState-swift.enum/requiresActivation
	PKSecureElementPassActivationStateRequiresActivation PKSecureElementPassActivationState = 0
	// PKSecureElementPassActivationStateSuspended - The user or the issuer has suspended the pass and it isn’t available to use.
	//
	// [Full Topic]: https://developer.apple.com/documentation/PassKit/PKSecureElementPass/PassActivationState-swift.enum/suspended
	PKSecureElementPassActivationStateSuspended PKSecureElementPassActivationState = 0
)

/* debug [enums.gen.go]: Processing enum PKShareSecureElementPassErrorCode (2 cases) */
// PKShareSecureElementPassErrorCode enum type
//
// [Full Topic]: https://developer.apple.com/documentation/PassKit/PKShareSecureElementPassError/Code
type PKShareSecureElementPassErrorCode uint

const (
	//
	// [Full Topic]: https://developer.apple.com/documentation/PassKit/PKShareSecureElementPassError/Code/setupError
	PKShareSecureElementPassSetupError PKShareSecureElementPassErrorCode = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/PassKit/PKShareSecureElementPassError/Code/unknownError
	PKShareSecureElementPassUnknownError PKShareSecureElementPassErrorCode = 0
)

/* debug [enums.gen.go]: Processing enum PKShareSecureElementPassResult (3 cases) */
// PKShareSecureElementPassResult enum type
//
// [Full Topic]: https://developer.apple.com/documentation/PassKit/PKShareSecureElementPassResult
type PKShareSecureElementPassResult uint

const (
	//
	// [Full Topic]: https://developer.apple.com/documentation/PassKit/PKShareSecureElementPassResult/canceled
	PKShareSecureElementPassResultCanceled PKShareSecureElementPassResult = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/PassKit/PKShareSecureElementPassResult/failed
	PKShareSecureElementPassResultFailed PKShareSecureElementPassResult = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/PassKit/PKShareSecureElementPassResult/shared
	PKShareSecureElementPassResultShared PKShareSecureElementPassResult = 0
)

/* debug [enums.gen.go]: Processing enum PKShippingContactEditingMode (3 cases) */
// PKShippingContactEditingMode - Constants that indicate whether the shipping mode prevents the user from editing fields of the shipping address.
//
// [Full Topic]: https://developer.apple.com/documentation/PassKit/PKShippingContactEditingMode
type PKShippingContactEditingMode uint

const (
	// PKShippingContactEditingModeAvailable - The value that indicates Apple Pay Later is available.
	//
	// [Full Topic]: https://developer.apple.com/documentation/PassKit/PKShippingContactEditingMode/available
	PKShippingContactEditingModeAvailable PKShippingContactEditingMode = 0
	// PKShippingContactEditingModeEnabled - All fields of the shipping contact on the payment sheet are editable by the user.
	//
	// [Full Topic]: https://developer.apple.com/documentation/PassKit/PKShippingContactEditingMode/enabled
	PKShippingContactEditingModeEnabled PKShippingContactEditingMode = 0
	// PKShippingContactEditingModeStorePickup - The shipping contact on the payment sheet represents a pickup address and isn’t editable by the user.
	//
	// [Full Topic]: https://developer.apple.com/documentation/PassKit/PKShippingContactEditingMode/storePickup
	PKShippingContactEditingModeStorePickup PKShippingContactEditingMode = 0
)

/* debug [enums.gen.go]: Processing enum PKShippingType (4 cases) */
// PKShippingType - A complete list of valid shipping types.
//
// [Full Topic]: https://developer.apple.com/documentation/PassKit/PKShippingType
type PKShippingType uint

const (
	// PKShippingTypeDelivery - Delivering the purchase by the seller (for example, pizza, flower, or furniture delivery).
	//
	// [Full Topic]: https://developer.apple.com/documentation/PassKit/PKShippingType/delivery
	PKShippingTypeDelivery PKShippingType = 0
	// PKShippingTypeServicePickup - Picking up an item from the provided address by the service (for example, transportation or shipping services that provide home pickup).
	//
	// [Full Topic]: https://developer.apple.com/documentation/PassKit/PKShippingType/servicePickup
	PKShippingTypeServicePickup PKShippingType = 0
	// PKShippingTypeShipping - Shipping the purchase to the provided address using a third-party shipping company. This is the default shipping type.
	//
	// [Full Topic]: https://developer.apple.com/documentation/PassKit/PKShippingType/shipping
	PKShippingTypeShipping PKShippingType = 0
	// PKShippingTypeStorePickup - Store pickup of the purchase from the seller’s store.
	//
	// [Full Topic]: https://developer.apple.com/documentation/PassKit/PKShippingType/storePickup
	PKShippingTypeStorePickup PKShippingType = 0
)

/* debug [enums.gen.go]: Processing enum PKVehicleConnectionErrorCode (3 cases) */
// PKVehicleConnectionErrorCode enum type
//
// [Full Topic]: https://developer.apple.com/documentation/PassKit/PKVehicleConnectionErrorCode
type PKVehicleConnectionErrorCode uint

const (
	//
	// [Full Topic]: https://developer.apple.com/documentation/PassKit/PKVehicleConnectionErrorCode/sessionNotActive
	PKVehicleConnectionErrorCodeSessionNotActive PKVehicleConnectionErrorCode = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/PassKit/PKVehicleConnectionErrorCode/sessionUnableToStart
	PKVehicleConnectionErrorCodeSessionUnableToStart PKVehicleConnectionErrorCode = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/PassKit/PKVehicleConnectionErrorCode/unknown
	PKVehicleConnectionErrorCodeUnknown PKVehicleConnectionErrorCode = 0
)

/* debug [enums.gen.go]: Processing enum PKVehicleConnectionSessionConnectionState (4 cases) */
// PKVehicleConnectionSessionConnectionState enum type
//
// [Full Topic]: https://developer.apple.com/documentation/PassKit/PKVehicleConnectionSessionConnectionState
type PKVehicleConnectionSessionConnectionState uint

const (
	//
	// [Full Topic]: https://developer.apple.com/documentation/PassKit/PKVehicleConnectionSessionConnectionState/connected
	PKVehicleConnectionSessionConnectionStateConnected PKVehicleConnectionSessionConnectionState = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/PassKit/PKVehicleConnectionSessionConnectionState/connecting
	PKVehicleConnectionSessionConnectionStateConnecting PKVehicleConnectionSessionConnectionState = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/PassKit/PKVehicleConnectionSessionConnectionState/disconnected
	PKVehicleConnectionSessionConnectionStateDisconnected PKVehicleConnectionSessionConnectionState = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/PassKit/PKVehicleConnectionSessionConnectionState/failedToConnect
	PKVehicleConnectionSessionConnectionStateFailedToConnect PKVehicleConnectionSessionConnectionState = 0
)


