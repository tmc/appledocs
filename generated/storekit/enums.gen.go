// Code generated from Apple documentation for StoreKit. DO NOT EDIT.

package storekit

// Enum types and constants
// SKANError - Constants that indicate the type of error for an ad network attribution operation.
//
// [Full Topic]: https://developer.apple.com/documentation/StoreKit/SKANError-swift.struct/Code
type ANError uint

const (
// ANErrorAdNetworkIdMissing - The ad network identifier in the ad impression doesn’t match the value in the information property list.
//
	// [Full Topic]: https://developer.apple.com/documentation/StoreKit/SKANError-swift.struct/Code/adNetworkIdMissing
ANErrorAdNetworkIdMissing ANError = 0
// ANErrorImpressionMissingRequiredValue - A required value is missing from a view-through ad impression.
//
	// [Full Topic]: https://developer.apple.com/documentation/StoreKit/SKANError-swift.struct/Code/impressionMissingRequiredValue
ANErrorImpressionMissingRequiredValue ANError = 0
// ANErrorImpressionNotFound - The system can’t find the ad impression.
//
	// [Full Topic]: https://developer.apple.com/documentation/StoreKit/SKANError-swift.struct/Code/impressionNotFound
ANErrorImpressionNotFound ANError = 0
//
	// [Full Topic]: https://developer.apple.com/documentation/StoreKit/SKANError-swift.struct/Code/impressionTooShort
ANErrorImpressionTooShort ANError = 0
// ANErrorInvalidAdvertisedAppId - The App Store ID of the advertised app is invalid.
//
	// [Full Topic]: https://developer.apple.com/documentation/StoreKit/SKANError-swift.struct/Code/invalidAdvertisedAppId
ANErrorInvalidAdvertisedAppId ANError = 0
// ANErrorInvalidCampaignId - The campaign identifier that you provided is invalid.
//
	// [Full Topic]: https://developer.apple.com/documentation/StoreKit/SKANError-swift.struct/Code/invalidCampaignId
ANErrorInvalidCampaignId ANError = 0
// ANErrorInvalidConversionValue - The conversion value is invalid.
//
	// [Full Topic]: https://developer.apple.com/documentation/StoreKit/SKANError-swift.struct/Code/invalidConversionValue
ANErrorInvalidConversionValue ANError = 0
// ANErrorInvalidSourceAppId - The App Store ID of the app displaying the ad is invalid.
//
	// [Full Topic]: https://developer.apple.com/documentation/StoreKit/SKANError-swift.struct/Code/invalidSourceAppId
ANErrorInvalidSourceAppId ANError = 0
// ANErrorInvalidVersion - The SKAdNetwork version number is invalid.
//
	// [Full Topic]: https://developer.apple.com/documentation/StoreKit/SKANError-swift.struct/Code/invalidVersion
ANErrorInvalidVersion ANError = 0
// ANErrorMismatchedSourceAppId - The source app identifier in the ad impression doesn’t match the app identifier in the source app.
//
	// [Full Topic]: https://developer.apple.com/documentation/StoreKit/SKANError-swift.struct/Code/mismatchedSourceAppId
ANErrorMismatchedSourceAppId ANError = 0
// ANErrorUnknown - An unknown error occurred.
//
	// [Full Topic]: https://developer.apple.com/documentation/StoreKit/SKANError-swift.struct/Code/unknown
ANErrorUnknown ANError = 0
// ANErrorUnsupported - Your app attempted to use functionality that isn’t supported in the specified version.
//
	// [Full Topic]: https://developer.apple.com/documentation/StoreKit/SKANError-swift.struct/Code/unsupported
ANErrorUnsupported ANError = 0
)

// SKCloudServiceAuthorizationStatus - Constants that indicate the type of authorization the customer has for accessing the Music library.
//
// [Full Topic]: https://developer.apple.com/documentation/StoreKit/SKCloudServiceAuthorizationStatus
type CloudServiceAuthorizationStatus uint

const (
// CloudServiceAuthorizationStatusAuthorized - The user authorizes playback of Apple Music tracks and the addition of tracks to their music library.
//
	// [Full Topic]: https://developer.apple.com/documentation/StoreKit/SKCloudServiceAuthorizationStatus/authorized
CloudServiceAuthorizationStatusAuthorized CloudServiceAuthorizationStatus = 0
// CloudServiceAuthorizationStatusDenied - The user does not authorize any access to their music library.
//
	// [Full Topic]: https://developer.apple.com/documentation/StoreKit/SKCloudServiceAuthorizationStatus/denied
CloudServiceAuthorizationStatusDenied CloudServiceAuthorizationStatus = 0
// CloudServiceAuthorizationStatusNotDetermined - The authorization type cannot be determined.
//
	// [Full Topic]: https://developer.apple.com/documentation/StoreKit/SKCloudServiceAuthorizationStatus/notDetermined
CloudServiceAuthorizationStatusNotDetermined CloudServiceAuthorizationStatus = 0
// CloudServiceAuthorizationStatusRestricted - Access to the music library is restricted in a way that the user cannot change, so your app should not prompt for authorization. An example of this situation is if the device is in an education mode.
//
	// [Full Topic]: https://developer.apple.com/documentation/StoreKit/SKCloudServiceAuthorizationStatus/restricted
CloudServiceAuthorizationStatusRestricted CloudServiceAuthorizationStatus = 0
)

// SKErrorCode - Error codes for StoreKit errors.
//
// [Full Topic]: https://developer.apple.com/documentation/StoreKit/SKError/Code
type ErrorCode uint

const (
// ErrorCloudServiceRevoked - Error code indicating that the user has revoked permission to use this cloud service.
//
	// [Full Topic]: https://developer.apple.com/documentation/StoreKit/SKError/Code/cloudServiceRevoked
ErrorCloudServiceRevoked ErrorCode = 0
// ErrorIneligibleForOffer - An error code that indicates the user is ineligible for the subscription offer.
//
	// [Full Topic]: https://developer.apple.com/documentation/StoreKit/SKError/Code/ineligibleForOffer
ErrorIneligibleForOffer ErrorCode = 0
//
	// [Full Topic]: https://developer.apple.com/documentation/StoreKit/SKError/Code/overlayTimeout
ErrorOverlayTimeout ErrorCode = 0
// ErrorPaymentCancelled - Error code indicating that the user canceled a payment request.
//
	// [Full Topic]: https://developer.apple.com/documentation/StoreKit/SKError/Code/paymentCancelled
ErrorPaymentCancelled ErrorCode = 0
)

// SKOverlayPosition - Constants that identify the position of an overlay on the screen.
//
// [Full Topic]: https://developer.apple.com/documentation/StoreKit/SKOverlay/Position
type OverlayPosition uint

const (
// OverlayPositionBottom - Specifies that the overlay is at the bottom of the screen.
//
	// [Full Topic]: https://developer.apple.com/documentation/StoreKit/SKOverlay/Position/bottom
OverlayPositionBottom OverlayPosition = 0
// OverlayPositionBottomRaised - Specifies that the overlay is at a raised position at the bottom of the screen.
//
	// [Full Topic]: https://developer.apple.com/documentation/StoreKit/SKOverlay/Position/bottomRaised
OverlayPositionBottomRaised OverlayPosition = 0
)

// SKPaymentTransactionState - Values representing the state of a transaction.
//
// [Full Topic]: https://developer.apple.com/documentation/StoreKit/SKPaymentTransactionState
type PaymentTransactionState uint

const (
// PaymentTransactionStateDeferred - A transaction that is in the queue, but its final status is pending external action such as Ask to Buy.
//
	// [Full Topic]: https://developer.apple.com/documentation/StoreKit/SKPaymentTransactionState/deferred
PaymentTransactionStateDeferred PaymentTransactionState = 0
// PaymentTransactionStateFailed - A failed transaction.
//
	// [Full Topic]: https://developer.apple.com/documentation/StoreKit/SKPaymentTransactionState/failed
PaymentTransactionStateFailed PaymentTransactionState = 0
// PaymentTransactionStatePurchased - A successfully processed transaction.
//
	// [Full Topic]: https://developer.apple.com/documentation/StoreKit/SKPaymentTransactionState/purchased
PaymentTransactionStatePurchased PaymentTransactionState = 0
// PaymentTransactionStatePurchasing - A transaction that is being processed by the App Store.
//
	// [Full Topic]: https://developer.apple.com/documentation/StoreKit/SKPaymentTransactionState/purchasing
PaymentTransactionStatePurchasing PaymentTransactionState = 0
// PaymentTransactionStateRestored - A transaction that restores content previously purchased by the user.
//
	// [Full Topic]: https://developer.apple.com/documentation/StoreKit/SKPaymentTransactionState/restored
PaymentTransactionStateRestored PaymentTransactionState = 0
)

// SKProductDiscountPaymentMode - Values representing the payment modes for a product discount.
//
// [Full Topic]: https://developer.apple.com/documentation/StoreKit/SKProductDiscount/PaymentMode-swift.enum
type ProductDiscountPaymentMode uint

const (
// ProductDiscountPaymentModeFreeTrial - A constant that indicates that the payment mode is a free trial.
//
	// [Full Topic]: https://developer.apple.com/documentation/StoreKit/SKProductDiscount/PaymentMode-swift.enum/freeTrial
ProductDiscountPaymentModeFreeTrial ProductDiscountPaymentMode = 0
// ProductDiscountPaymentModePayAsYouGo - A constant that indicates a product discount that applies over a single billing period or multiple billing periods.
//
	// [Full Topic]: https://developer.apple.com/documentation/StoreKit/SKProductDiscount/PaymentMode-swift.enum/payAsYouGo
ProductDiscountPaymentModePayAsYouGo ProductDiscountPaymentMode = 0
// ProductDiscountPaymentModePayUpFront - A constant that indicates that the system applies the product discount up front.
//
	// [Full Topic]: https://developer.apple.com/documentation/StoreKit/SKProductDiscount/PaymentMode-swift.enum/payUpFront
ProductDiscountPaymentModePayUpFront ProductDiscountPaymentMode = 0
)

// SKProductDiscountType - Values representing the types of discount offers an app can present.
//
// [Full Topic]: https://developer.apple.com/documentation/StoreKit/SKProductDiscount/Type-swift.enum
type ProductDiscountType uint


