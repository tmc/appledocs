// Code generated from Apple documentation for StoreKit. DO NOT EDIT.

package storekit

/* debug [enums.gen.go]: Generating 11 enums for StoreKit */
// Enum types and constants
/* debug [enums.gen.go]: Processing enum SKOverlayPosition (2 cases) */
// SKOverlayPosition - Constants that identify the position of an overlay on the screen.
//
// [Full Topic]: https://developer.apple.com/documentation/StoreKit/SKOverlay/Position
type SKOverlayPosition uint

const (
	// SKOverlayPositionBottom - Specifies that the overlay is at the bottom of the screen.
	//
	// [Full Topic]: https://developer.apple.com/documentation/StoreKit/SKOverlay/Position/bottom
	SKOverlayPositionBottom SKOverlayPosition = 0
	// SKOverlayPositionBottomRaised - Specifies that the overlay is at a raised position at the bottom of the screen.
	//
	// [Full Topic]: https://developer.apple.com/documentation/StoreKit/SKOverlay/Position/bottomRaised
	SKOverlayPositionBottomRaised SKOverlayPosition = 0
)

/* debug [enums.gen.go]: Processing enum SKANError (12 cases) */
// SKANError - Constants that indicate the type of error for an ad network attribution operation.
//
// [Full Topic]: https://developer.apple.com/documentation/StoreKit/SKANError-swift.struct/Code
type SKANError uint

const (
	// SKANErrorAdNetworkIdMissing - The ad network identifier in the ad impression doesn’t match the value in the information property list.
	//
	// [Full Topic]: https://developer.apple.com/documentation/StoreKit/SKANError-swift.struct/Code/adNetworkIdMissing
	SKANErrorAdNetworkIdMissing SKANError = 0
	// SKANErrorImpressionMissingRequiredValue - A required value is missing from a view-through ad impression.
	//
	// [Full Topic]: https://developer.apple.com/documentation/StoreKit/SKANError-swift.struct/Code/impressionMissingRequiredValue
	SKANErrorImpressionMissingRequiredValue SKANError = 0
	// SKANErrorImpressionNotFound - The system can’t find the ad impression.
	//
	// [Full Topic]: https://developer.apple.com/documentation/StoreKit/SKANError-swift.struct/Code/impressionNotFound
	SKANErrorImpressionNotFound SKANError = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/StoreKit/SKANError-swift.struct/Code/impressionTooShort
	SKANErrorImpressionTooShort SKANError = 0
	// SKANErrorInvalidAdvertisedAppId - The App Store ID of the advertised app is invalid.
	//
	// [Full Topic]: https://developer.apple.com/documentation/StoreKit/SKANError-swift.struct/Code/invalidAdvertisedAppId
	SKANErrorInvalidAdvertisedAppId SKANError = 0
	// SKANErrorInvalidCampaignId - The campaign identifier that you provided is invalid.
	//
	// [Full Topic]: https://developer.apple.com/documentation/StoreKit/SKANError-swift.struct/Code/invalidCampaignId
	SKANErrorInvalidCampaignId SKANError = 0
	// SKANErrorInvalidConversionValue - The conversion value is invalid.
	//
	// [Full Topic]: https://developer.apple.com/documentation/StoreKit/SKANError-swift.struct/Code/invalidConversionValue
	SKANErrorInvalidConversionValue SKANError = 0
	// SKANErrorInvalidSourceAppId - The App Store ID of the app displaying the ad is invalid.
	//
	// [Full Topic]: https://developer.apple.com/documentation/StoreKit/SKANError-swift.struct/Code/invalidSourceAppId
	SKANErrorInvalidSourceAppId SKANError = 0
	// SKANErrorInvalidVersion - The SKAdNetwork version number is invalid.
	//
	// [Full Topic]: https://developer.apple.com/documentation/StoreKit/SKANError-swift.struct/Code/invalidVersion
	SKANErrorInvalidVersion SKANError = 0
	// SKANErrorMismatchedSourceAppId - The source app identifier in the ad impression doesn’t match the app identifier in the source app.
	//
	// [Full Topic]: https://developer.apple.com/documentation/StoreKit/SKANError-swift.struct/Code/mismatchedSourceAppId
	SKANErrorMismatchedSourceAppId SKANError = 0
	// SKANErrorUnknown - An unknown error occurred.
	//
	// [Full Topic]: https://developer.apple.com/documentation/StoreKit/SKANError-swift.struct/Code/unknown
	SKANErrorUnknown SKANError = 0
	// SKANErrorUnsupported - Your app attempted to use functionality that isn’t supported in the specified version.
	//
	// [Full Topic]: https://developer.apple.com/documentation/StoreKit/SKANError-swift.struct/Code/unsupported
	SKANErrorUnsupported SKANError = 0
)

/* debug [enums.gen.go]: Processing enum SKCloudServiceAuthorizationStatus (4 cases) */
// SKCloudServiceAuthorizationStatus - Constants that indicate the type of authorization the customer has for accessing the Music library.
//
// [Full Topic]: https://developer.apple.com/documentation/StoreKit/SKCloudServiceAuthorizationStatus
type SKCloudServiceAuthorizationStatus uint

const (
	// SKCloudServiceAuthorizationStatusAuthorized - The user authorizes playback of Apple Music tracks and the addition of tracks to their music library.
	//
	// [Full Topic]: https://developer.apple.com/documentation/StoreKit/SKCloudServiceAuthorizationStatus/authorized
	SKCloudServiceAuthorizationStatusAuthorized SKCloudServiceAuthorizationStatus = 0
	// SKCloudServiceAuthorizationStatusDenied - The user does not authorize any access to their music library.
	//
	// [Full Topic]: https://developer.apple.com/documentation/StoreKit/SKCloudServiceAuthorizationStatus/denied
	SKCloudServiceAuthorizationStatusDenied SKCloudServiceAuthorizationStatus = 0
	// SKCloudServiceAuthorizationStatusNotDetermined - The authorization type cannot be determined.
	//
	// [Full Topic]: https://developer.apple.com/documentation/StoreKit/SKCloudServiceAuthorizationStatus/notDetermined
	SKCloudServiceAuthorizationStatusNotDetermined SKCloudServiceAuthorizationStatus = 0
	// SKCloudServiceAuthorizationStatusRestricted - Access to the music library is restricted in a way that the user cannot change, so your app should not prompt for authorization. An example of this situation is if the device is in an education mode.
	//
	// [Full Topic]: https://developer.apple.com/documentation/StoreKit/SKCloudServiceAuthorizationStatus/restricted
	SKCloudServiceAuthorizationStatusRestricted SKCloudServiceAuthorizationStatus = 0
)

/* debug [enums.gen.go]: Processing enum SKCloudServiceCapability (4 cases) */
// SKCloudServiceCapability - Constants that specify the current capabilities of the customer’s Music library on the device.
//
// [Full Topic]: https://developer.apple.com/documentation/StoreKit/SKCloudServiceCapability
type SKCloudServiceCapability uint

const (
	// SKCloudServiceCapabilityAddToCloudMusicLibrary - The device allows tracks to be added to the user’s music library.
	//
	// [Full Topic]: https://developer.apple.com/documentation/StoreKit/SKCloudServiceCapability/addToCloudMusicLibrary
	SKCloudServiceCapabilityAddToCloudMusicLibrary SKCloudServiceCapability = 0
	// SKCloudServiceCapabilityMusicCatalogPlayback - The device allows playback of Apple Music catalog tracks.
	//
	// [Full Topic]: https://developer.apple.com/documentation/StoreKit/SKCloudServiceCapability/musicCatalogPlayback
	SKCloudServiceCapabilityMusicCatalogPlayback SKCloudServiceCapability = 0
	// SKCloudServiceCapabilityMusicCatalogSubscriptionEligible - The device allows subscription to the Apple Music catalog.
	//
	// [Full Topic]: https://developer.apple.com/documentation/StoreKit/SKCloudServiceCapability/musicCatalogSubscriptionEligible
	SKCloudServiceCapabilityMusicCatalogSubscriptionEligible SKCloudServiceCapability = 0
	// SKCloudServiceCapabilityNone - The device does not allow playback of Apple Music content or the addition of tracks to the music library.
	//
	// [Full Topic]: https://developer.apple.com/documentation/StoreKit/SKCloudServiceCapability/SKCloudServiceCapabilityNone
	SKCloudServiceCapabilityNone SKCloudServiceCapability = 0
)

/* debug [enums.gen.go]: Processing enum SKDownloadState (6 cases) */
// SKDownloadState - The states that a download operation can be in.
//
// [Full Topic]: https://developer.apple.com/documentation/StoreKit/SKDownloadState
type SKDownloadState uint

const (
	// SKDownloadStateActive - Indicates that the content is currently being downloaded.
	//
	// [Full Topic]: https://developer.apple.com/documentation/StoreKit/SKDownloadState/active
	SKDownloadStateActive SKDownloadState = 0
	// SKDownloadStateCancelled - Indicates that your app canceled the download.
	//
	// [Full Topic]: https://developer.apple.com/documentation/StoreKit/SKDownloadState/cancelled
	SKDownloadStateCancelled SKDownloadState = 0
	// SKDownloadStateFailed - Indicates that an error occurred while the file was being downloaded.
	//
	// [Full Topic]: https://developer.apple.com/documentation/StoreKit/SKDownloadState/failed
	SKDownloadStateFailed SKDownloadState = 0
	// SKDownloadStateFinished - Indicates that the content was successfully downloaded.
	//
	// [Full Topic]: https://developer.apple.com/documentation/StoreKit/SKDownloadState/finished
	SKDownloadStateFinished SKDownloadState = 0
	// SKDownloadStatePaused - Indicates that your app paused the download.
	//
	// [Full Topic]: https://developer.apple.com/documentation/StoreKit/SKDownloadState/paused
	SKDownloadStatePaused SKDownloadState = 0
	// SKDownloadStateWaiting - Indicates that the download has not started yet.
	//
	// [Full Topic]: https://developer.apple.com/documentation/StoreKit/SKDownloadState/waiting
	SKDownloadStateWaiting SKDownloadState = 0
)

/* debug [enums.gen.go]: Processing enum SKErrorCode (21 cases) */
// SKErrorCode - Error codes for StoreKit errors.
//
// [Full Topic]: https://developer.apple.com/documentation/StoreKit/SKError/Code
type SKErrorCode uint

const (
	// SKErrorClientInvalid - Error code indicating that the client is not allowed to perform the attempted action.
	//
	// [Full Topic]: https://developer.apple.com/documentation/StoreKit/SKError/Code/clientInvalid
	SKErrorClientInvalid SKErrorCode = 0
	// SKErrorCloudServiceNetworkConnectionFailed - Error code indicating that the device could not connect to the network.
	//
	// [Full Topic]: https://developer.apple.com/documentation/StoreKit/SKError/Code/cloudServiceNetworkConnectionFailed
	SKErrorCloudServiceNetworkConnectionFailed SKErrorCode = 0
	// SKErrorCloudServicePermissionDenied - Error code indicating that the user has not allowed access to Cloud service information.
	//
	// [Full Topic]: https://developer.apple.com/documentation/StoreKit/SKError/Code/cloudServicePermissionDenied
	SKErrorCloudServicePermissionDenied SKErrorCode = 0
	// SKErrorCloudServiceRevoked - Error code indicating that the user has revoked permission to use this cloud service.
	//
	// [Full Topic]: https://developer.apple.com/documentation/StoreKit/SKError/Code/cloudServiceRevoked
	SKErrorCloudServiceRevoked SKErrorCode = 0
	// SKErrorIneligibleForOffer - An error code that indicates the user is ineligible for the subscription offer.
	//
	// [Full Topic]: https://developer.apple.com/documentation/StoreKit/SKError/Code/ineligibleForOffer
	SKErrorIneligibleForOffer SKErrorCode = 0
	// SKErrorInvalidOfferIdentifier - Error code indicating that the offer identifier is invalid.
	//
	// [Full Topic]: https://developer.apple.com/documentation/StoreKit/SKError/Code/invalidOfferIdentifier
	SKErrorInvalidOfferIdentifier SKErrorCode = 0
	// SKErrorInvalidOfferPrice - Error code indicating that the price you specified in App Store Connect is no longer valid.
	//
	// [Full Topic]: https://developer.apple.com/documentation/StoreKit/SKError/Code/invalidOfferPrice
	SKErrorInvalidOfferPrice SKErrorCode = 0
	// SKErrorInvalidSignature - Error code indicating that the signature in a payment discount isn’t valid.
	//
	// [Full Topic]: https://developer.apple.com/documentation/StoreKit/SKError/Code/invalidSignature
	SKErrorInvalidSignature SKErrorCode = 0
	// SKErrorMissingOfferParams - Error code indicating that parameters are missing in a payment discount.
	//
	// [Full Topic]: https://developer.apple.com/documentation/StoreKit/SKError/Code/missingOfferParams
	SKErrorMissingOfferParams SKErrorCode = 0
	// SKErrorOverlayCancelled - An error code that indicates the cancellation of an overlay.
	//
	// [Full Topic]: https://developer.apple.com/documentation/StoreKit/SKError/Code/overlayCancelled
	SKErrorOverlayCancelled SKErrorCode = 0
	// SKErrorOverlayInvalidConfiguration - An error code that indicates the overlay’s configuration is invalid.
	//
	// [Full Topic]: https://developer.apple.com/documentation/StoreKit/SKError/Code/overlayInvalidConfiguration
	SKErrorOverlayInvalidConfiguration SKErrorCode = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/StoreKit/SKError/Code/overlayPresentedInBackgroundScene
	SKErrorOverlayPresentedInBackgroundScene SKErrorCode = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/StoreKit/SKError/Code/overlayTimeout
	SKErrorOverlayTimeout SKErrorCode = 0
	// SKErrorPaymentCancelled - Error code indicating that the user canceled a payment request.
	//
	// [Full Topic]: https://developer.apple.com/documentation/StoreKit/SKError/Code/paymentCancelled
	SKErrorPaymentCancelled SKErrorCode = 0
	// SKErrorPaymentInvalid - Error code indicating that one of the payment parameters wasn’t recognized by the App Store.
	//
	// [Full Topic]: https://developer.apple.com/documentation/StoreKit/SKError/Code/paymentInvalid
	SKErrorPaymentInvalid SKErrorCode = 0
	// SKErrorPaymentNotAllowed - Error code indicating that the user is not allowed to authorize payments.
	//
	// [Full Topic]: https://developer.apple.com/documentation/StoreKit/SKError/Code/paymentNotAllowed
	SKErrorPaymentNotAllowed SKErrorCode = 0
	// SKErrorPrivacyAcknowledgementRequired - Error code indicating that the user has not yet acknowledged Apple’s privacy policy for Apple Music.
	//
	// [Full Topic]: https://developer.apple.com/documentation/StoreKit/SKError/Code/privacyAcknowledgementRequired
	SKErrorPrivacyAcknowledgementRequired SKErrorCode = 0
	// SKErrorStoreProductNotAvailable - Error code indicating that the requested product is not available in the store.
	//
	// [Full Topic]: https://developer.apple.com/documentation/StoreKit/SKError/Code/storeProductNotAvailable
	SKErrorStoreProductNotAvailable SKErrorCode = 0
	// SKErrorUnauthorizedRequestData - Error code indicating that the app is attempting to use a property for which it does not have the required entitlement.
	//
	// [Full Topic]: https://developer.apple.com/documentation/StoreKit/SKError/Code/unauthorizedRequestData
	SKErrorUnauthorizedRequestData SKErrorCode = 0
	// SKErrorUnknown - Error code indicating that an unknown or unexpected error occurred.
	//
	// [Full Topic]: https://developer.apple.com/documentation/StoreKit/SKError/Code/unknown
	SKErrorUnknown SKErrorCode = 0
	// SKErrorUnsupportedPlatform - An error code that indicates the current platform doesn’t support overlays.
	//
	// [Full Topic]: https://developer.apple.com/documentation/StoreKit/SKError/Code/unsupportedPlatform
	SKErrorUnsupportedPlatform SKErrorCode = 0
)

/* debug [enums.gen.go]: Processing enum SKPaymentTransactionState (5 cases) */
// SKPaymentTransactionState - Values representing the state of a transaction.
//
// [Full Topic]: https://developer.apple.com/documentation/StoreKit/SKPaymentTransactionState
type SKPaymentTransactionState uint

const (
	// SKPaymentTransactionStateDeferred - A transaction that is in the queue, but its final status is pending external action such as Ask to Buy.
	//
	// [Full Topic]: https://developer.apple.com/documentation/StoreKit/SKPaymentTransactionState/deferred
	SKPaymentTransactionStateDeferred SKPaymentTransactionState = 0
	// SKPaymentTransactionStateFailed - A failed transaction.
	//
	// [Full Topic]: https://developer.apple.com/documentation/StoreKit/SKPaymentTransactionState/failed
	SKPaymentTransactionStateFailed SKPaymentTransactionState = 0
	// SKPaymentTransactionStatePurchased - A successfully processed transaction.
	//
	// [Full Topic]: https://developer.apple.com/documentation/StoreKit/SKPaymentTransactionState/purchased
	SKPaymentTransactionStatePurchased SKPaymentTransactionState = 0
	// SKPaymentTransactionStatePurchasing - A transaction that is being processed by the App Store.
	//
	// [Full Topic]: https://developer.apple.com/documentation/StoreKit/SKPaymentTransactionState/purchasing
	SKPaymentTransactionStatePurchasing SKPaymentTransactionState = 0
	// SKPaymentTransactionStateRestored - A transaction that restores content previously purchased by the user.
	//
	// [Full Topic]: https://developer.apple.com/documentation/StoreKit/SKPaymentTransactionState/restored
	SKPaymentTransactionStateRestored SKPaymentTransactionState = 0
)

/* debug [enums.gen.go]: Processing enum SKProductPeriodUnit (4 cases) */
// SKProductPeriodUnit - Values representing the duration of an interval, from a day up to a year.
//
// [Full Topic]: https://developer.apple.com/documentation/StoreKit/SKProduct/PeriodUnit
type SKProductPeriodUnit uint

const (
	// SKProductPeriodUnitDay - An interval lasting one day.
	//
	// [Full Topic]: https://developer.apple.com/documentation/StoreKit/SKProduct/PeriodUnit/day
	SKProductPeriodUnitDay SKProductPeriodUnit = 0
	// SKProductPeriodUnitMonth - An interval lasting one month.
	//
	// [Full Topic]: https://developer.apple.com/documentation/StoreKit/SKProduct/PeriodUnit/month
	SKProductPeriodUnitMonth SKProductPeriodUnit = 0
	// SKProductPeriodUnitWeek - An interval lasting one week.
	//
	// [Full Topic]: https://developer.apple.com/documentation/StoreKit/SKProduct/PeriodUnit/week
	SKProductPeriodUnitWeek SKProductPeriodUnit = 0
	// SKProductPeriodUnitYear - An interval lasting one year.
	//
	// [Full Topic]: https://developer.apple.com/documentation/StoreKit/SKProduct/PeriodUnit/year
	SKProductPeriodUnitYear SKProductPeriodUnit = 0
)

/* debug [enums.gen.go]: Processing enum SKProductDiscountPaymentMode (3 cases) */
// SKProductDiscountPaymentMode - Values representing the payment modes for a product discount.
//
// [Full Topic]: https://developer.apple.com/documentation/StoreKit/SKProductDiscount/PaymentMode-swift.enum
type SKProductDiscountPaymentMode uint

const (
	// SKProductDiscountPaymentModeFreeTrial - A constant that indicates that the payment mode is a free trial.
	//
	// [Full Topic]: https://developer.apple.com/documentation/StoreKit/SKProductDiscount/PaymentMode-swift.enum/freeTrial
	SKProductDiscountPaymentModeFreeTrial SKProductDiscountPaymentMode = 0
	// SKProductDiscountPaymentModePayAsYouGo - A constant that indicates a product discount that applies over a single billing period or multiple billing periods.
	//
	// [Full Topic]: https://developer.apple.com/documentation/StoreKit/SKProductDiscount/PaymentMode-swift.enum/payAsYouGo
	SKProductDiscountPaymentModePayAsYouGo SKProductDiscountPaymentMode = 0
	// SKProductDiscountPaymentModePayUpFront - A constant that indicates that the system applies the product discount up front.
	//
	// [Full Topic]: https://developer.apple.com/documentation/StoreKit/SKProductDiscount/PaymentMode-swift.enum/payUpFront
	SKProductDiscountPaymentModePayUpFront SKProductDiscountPaymentMode = 0
)

/* debug [enums.gen.go]: Processing enum SKProductDiscountType (2 cases) */
// SKProductDiscountType - Values representing the types of discount offers an app can present.
//
// [Full Topic]: https://developer.apple.com/documentation/StoreKit/SKProductDiscount/Type-swift.enum
type SKProductDiscountType uint

const (
	// SKProductDiscountTypeIntroductory - A constant indicating the discount type is an introductory offer.
	//
	// [Full Topic]: https://developer.apple.com/documentation/StoreKit/SKProductDiscount/Type-swift.enum/introductory
	SKProductDiscountTypeIntroductory SKProductDiscountType = 0
	// SKProductDiscountTypeSubscription - A constant indicating the discount type is a promotional offer.
	//
	// [Full Topic]: https://developer.apple.com/documentation/StoreKit/SKProductDiscount/Type-swift.enum/subscription
	SKProductDiscountTypeSubscription SKProductDiscountType = 0
)

/* debug [enums.gen.go]: Processing enum SKProductStorePromotionVisibility (3 cases) */
// SKProductStorePromotionVisibility - The visibility settings that determine if an in-app purchase is visible on a device.
//
// [Full Topic]: https://developer.apple.com/documentation/StoreKit/SKProductStorePromotionVisibility
type SKProductStorePromotionVisibility uint

const (
	// SKProductStorePromotionVisibilityDefault - Indicates product visibility is the same as the default value set in App Store Connect.
	//
	// [Full Topic]: https://developer.apple.com/documentation/StoreKit/SKProductStorePromotionVisibility/default
	SKProductStorePromotionVisibilityDefault SKProductStorePromotionVisibility = 0
	// SKProductStorePromotionVisibilityHide - Indicates product is hidden.
	//
	// [Full Topic]: https://developer.apple.com/documentation/StoreKit/SKProductStorePromotionVisibility/hide
	SKProductStorePromotionVisibilityHide SKProductStorePromotionVisibility = 0
	// SKProductStorePromotionVisibilityShow - Indicates product is shown.
	//
	// [Full Topic]: https://developer.apple.com/documentation/StoreKit/SKProductStorePromotionVisibility/show
	SKProductStorePromotionVisibilityShow SKProductStorePromotionVisibility = 0
)
