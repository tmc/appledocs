// Code generated from Apple documentation for PassKit. DO NOT EDIT.

package passkit

// Enum types and constants
// PKAddressField - Billing or shipping address fields.
//
// [Full Topic]: https://developer.apple.com/documentation/PassKit/PKAddressField
type PKAddressField uint

// PKMerchantCapability - Capabilities for processing payment.
//
// [Full Topic]: https://developer.apple.com/documentation/PassKit/PKMerchantCapability
type PKMerchantCapability uint

// PKPassLibraryAuthorizationStatus enum type
//
// [Full Topic]: https://developer.apple.com/documentation/PassKit/PKPassLibrary/AuthorizationStatus
type PKPassLibraryAuthorizationStatus uint

const (
	//
	// [Full Topic]: https://developer.apple.com/documentation/PassKit/PKPassLibrary/AuthorizationStatus/authorized
	PKPassLibraryAuthorizationStatusAuthorized PKPassLibraryAuthorizationStatus = 0
)

// PKPassLibraryCapability enum type
//
// [Full Topic]: https://developer.apple.com/documentation/PassKit/PKPassLibrary/Capability
type PKPassLibraryCapability uint

// PKPaymentAuthorizationStatus - General success and failure status for payment authorization.
//
// [Full Topic]: https://developer.apple.com/documentation/PassKit/PKPaymentAuthorizationStatus
type PKPaymentAuthorizationStatus uint

const (
	// PKPaymentAuthorizationStatusFailure - Merchant failed to authorize the transaction.
	//
	// [Full Topic]: https://developer.apple.com/documentation/PassKit/PKPaymentAuthorizationStatus/failure
	PKPaymentAuthorizationStatusFailure PKPaymentAuthorizationStatus = 0
	// PKPaymentAuthorizationStatusSuccess - Merchant successfully authorized the transaction, or the transaction is expected to succeed.
	//
	// [Full Topic]: https://developer.apple.com/documentation/PassKit/PKPaymentAuthorizationStatus/success
	PKPaymentAuthorizationStatusSuccess PKPaymentAuthorizationStatus = 0
)

// PKShippingContactEditingMode - Constants that indicate whether the shipping mode prevents the user from editing fields of the shipping address.
//
// [Full Topic]: https://developer.apple.com/documentation/PassKit/PKShippingContactEditingMode
type PKShippingContactEditingMode uint

const (
	// PKShippingContactEditingModeEnabled - All fields of the shipping contact on the payment sheet are editable by the user.
	//
	// [Full Topic]: https://developer.apple.com/documentation/PassKit/PKShippingContactEditingMode/enabled
	PKShippingContactEditingModeEnabled PKShippingContactEditingMode = 0
)

// PKShippingType - A complete list of valid shipping types.
//
// [Full Topic]: https://developer.apple.com/documentation/PassKit/PKShippingType
type PKShippingType uint


