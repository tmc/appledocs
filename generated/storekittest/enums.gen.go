// Code generated from Apple documentation for StoreKitTest. DO NOT EDIT.

package storekittest

// Enum types and constants
// SKAdTestErrorCode - Enumerated error codes related to ad network testing in the testing environment.
//
// [Full Topic]: https://developer.apple.com/documentation/StoreKitTest/SKAdTestError/Code
type AdTestErrorCode uint

const (
	// AdTestErrorCodeConflictingSource - This error code is unused.
	//
	// [Full Topic]: https://developer.apple.com/documentation/StoreKitTest/SKAdTestError/Code/conflictingSource
	AdTestErrorCodeConflictingSource AdTestErrorCode = 0
	// AdTestErrorCodeExcessivePostbacks - Too many postbacks submitted to the test session.
	//
	// [Full Topic]: https://developer.apple.com/documentation/StoreKitTest/SKAdTestError/Code/excessivePostbacks
	AdTestErrorCodeExcessivePostbacks AdTestErrorCode = 0
	// AdTestErrorCodeInvalidCampaignId - The campaign ID isn’t an integer between one and one hundred.
	//
	// [Full Topic]: https://developer.apple.com/documentation/StoreKitTest/SKAdTestError/Code/invalidCampaignId
	AdTestErrorCodeInvalidCampaignId AdTestErrorCode = 0
	// AdTestErrorCodeInvalidConversionValue - The conversion value isn’t valid, in the testing environment.
	//
	// [Full Topic]: https://developer.apple.com/documentation/StoreKitTest/SKAdTestError/Code/invalidConversionValue
	AdTestErrorCodeInvalidConversionValue AdTestErrorCode = 0
	// AdTestErrorCodeInvalidImpressionId - The impression ID isn’t a valid UUID string.
	//
	// [Full Topic]: https://developer.apple.com/documentation/StoreKitTest/SKAdTestError/Code/invalidImpressionId
	AdTestErrorCodeInvalidImpressionId AdTestErrorCode = 0
	// AdTestErrorCodeInvalidPostbackURL - The URL for the postback isn’t valid, in the testing environment.
	//
	// [Full Topic]: https://developer.apple.com/documentation/StoreKitTest/SKAdTestError/Code/invalidPostbackURL
	AdTestErrorCodeInvalidPostbackURL AdTestErrorCode = 0
	// AdTestErrorCodeInvalidRunnerUpPostback - A non-winning postback is defined with a version prior to version 3, in the testing environment.
	//
	// [Full Topic]: https://developer.apple.com/documentation/StoreKitTest/SKAdTestError/Code/invalidRunnerUpPostback
	AdTestErrorCodeInvalidRunnerUpPostback AdTestErrorCode = 0
	// AdTestErrorCodeInvalidSourceAppAdamId - The app ID is less than zero.
	//
	// [Full Topic]: https://developer.apple.com/documentation/StoreKitTest/SKAdTestError/Code/invalidSourceAppAdamId
	AdTestErrorCodeInvalidSourceAppAdamId AdTestErrorCode = 0
	// AdTestErrorCodeInvalidSourceDomain - The source domain isn’t in the correct format.
	//
	// [Full Topic]: https://developer.apple.com/documentation/StoreKitTest/SKAdTestError/Code/invalidSourceDomain
	AdTestErrorCodeInvalidSourceDomain AdTestErrorCode = 0
	// AdTestErrorCodeInvalidSourceIdentifier - The postback’s identifier isn’t two, three, or four digits.
	//
	// [Full Topic]: https://developer.apple.com/documentation/StoreKitTest/SKAdTestError/Code/invalidSourceIdentifier
	AdTestErrorCodeInvalidSourceIdentifier AdTestErrorCode = 0
	// AdTestErrorCodeInvalidVersion - A postback contains an incorrect version number.
	//
	// [Full Topic]: https://developer.apple.com/documentation/StoreKitTest/SKAdTestError/Code/invalidVersion
	AdTestErrorCodeInvalidVersion AdTestErrorCode = 0
	// AdTestErrorCodeInvalidWinningPostbackCount - The number of winning postbacks isn’t valid, in the testing environment.
	//
	// [Full Topic]: https://developer.apple.com/documentation/StoreKitTest/SKAdTestError/Code/invalidWinningPostbackCount
	AdTestErrorCodeInvalidWinningPostbackCount AdTestErrorCode = 0
	// AdTestErrorCodeMalformedPostbacks - The postback in the testing environment is malformed.
	//
	// [Full Topic]: https://developer.apple.com/documentation/StoreKitTest/SKAdTestError/Code/malformedPostbacks
	AdTestErrorCodeMalformedPostbacks AdTestErrorCode = 0
	// AdTestErrorCodeMisplacedWinnerPostback - A winning postback wasn’t found in the first position, in the testing environment.
	//
	// [Full Topic]: https://developer.apple.com/documentation/StoreKitTest/SKAdTestError/Code/misplacedWinnerPostback
	AdTestErrorCodeMisplacedWinnerPostback AdTestErrorCode = 0
	// AdTestErrorCodeMissingPostbacks - The testing environment doesn’t have any postbacks.
	//
	// [Full Topic]: https://developer.apple.com/documentation/StoreKitTest/SKAdTestError/Code/missingPostbacks
	AdTestErrorCodeMissingPostbacks AdTestErrorCode = 0
	// AdTestErrorCodeMissingSignature - The signature for the ad is missing, in the testing environment.
	//
	// [Full Topic]: https://developer.apple.com/documentation/StoreKitTest/SKAdTestError/Code/missingSignature
	AdTestErrorCodeMissingSignature AdTestErrorCode = 0
	// AdTestErrorCodeMissingWinningPostback - The testing environment is missing a winning postback.
	//
	// [Full Topic]: https://developer.apple.com/documentation/StoreKitTest/SKAdTestError/Code/missingWinningPostback
	AdTestErrorCodeMissingWinningPostback AdTestErrorCode = 0
	// AdTestErrorCodeNoPendingPostbacks - The test session doesn’t have any pending postbacks to send.
	//
	// [Full Topic]: https://developer.apple.com/documentation/StoreKitTest/SKAdTestError/Code/noPendingPostbacks
	AdTestErrorCodeNoPendingPostbacks AdTestErrorCode = 0
	// AdTestErrorCodeSignatureInvalidKey - The public key isn’t a valid cryptographic key, in the testing environment.
	//
	// [Full Topic]: https://developer.apple.com/documentation/StoreKitTest/SKAdTestError/Code/signatureInvalidKey
	AdTestErrorCodeSignatureInvalidKey AdTestErrorCode = 0
	// AdTestErrorCodeSignatureInvalidOrder - The order of the parameters in the signature is invalid, in the testing environment.
	//
	// [Full Topic]: https://developer.apple.com/documentation/StoreKitTest/SKAdTestError/Code/signatureInvalidOrder
	AdTestErrorCodeSignatureInvalidOrder AdTestErrorCode = 0
	// AdTestErrorCodeSignatureMissingAdNetworkId - The signature is missing an ad network identifier, in the testing environment.
	//
	// [Full Topic]: https://developer.apple.com/documentation/StoreKitTest/SKAdTestError/Code/signatureMissingAdNetworkId
	AdTestErrorCodeSignatureMissingAdNetworkId AdTestErrorCode = 0
	// AdTestErrorCodeSignatureMissingAppAdamId - The signature is missing the app item identifier for the advertised app, in the testing environment.
	//
	// [Full Topic]: https://developer.apple.com/documentation/StoreKitTest/SKAdTestError/Code/signatureMissingAppAdamId
	AdTestErrorCodeSignatureMissingAppAdamId AdTestErrorCode = 0
	// AdTestErrorCodeSignatureMissingCampaignId - The signature is missing the campaign identifier, in the testing environment.
	//
	// [Full Topic]: https://developer.apple.com/documentation/StoreKitTest/SKAdTestError/Code/signatureMissingCampaignId
	AdTestErrorCodeSignatureMissingCampaignId AdTestErrorCode = 0
	// AdTestErrorCodeSignatureMissingFidelityType - The signature is missing the fidelity type, in the testing environment.
	//
	// [Full Topic]: https://developer.apple.com/documentation/StoreKitTest/SKAdTestError/Code/signatureMissingFidelityType
	AdTestErrorCodeSignatureMissingFidelityType AdTestErrorCode = 0
	// AdTestErrorCodeSignatureMissingNonce - The signature is missing the nonce, in the testing environment.
	//
	// [Full Topic]: https://developer.apple.com/documentation/StoreKitTest/SKAdTestError/Code/signatureMissingNonce
	AdTestErrorCodeSignatureMissingNonce AdTestErrorCode = 0
	// AdTestErrorCodeSignatureMissingSourceAppAdamId - The signature is missing the source app item identifier, in the testing environment.
	//
	// [Full Topic]: https://developer.apple.com/documentation/StoreKitTest/SKAdTestError/Code/signatureMissingSourceAppAdamId
	AdTestErrorCodeSignatureMissingSourceAppAdamId AdTestErrorCode = 0
	// AdTestErrorCodeSignatureMissingSourceDomain - The signature is missing the source domain, in the testing environment.
	//
	// [Full Topic]: https://developer.apple.com/documentation/StoreKitTest/SKAdTestError/Code/signatureMissingSourceDomain
	AdTestErrorCodeSignatureMissingSourceDomain AdTestErrorCode = 0
	// AdTestErrorCodeSignatureMissingSourceIdentifier - The signature is missing the source identifier, in the testing environment.
	//
	// [Full Topic]: https://developer.apple.com/documentation/StoreKitTest/SKAdTestError/Code/signatureMissingSourceIdentifier
	AdTestErrorCodeSignatureMissingSourceIdentifier AdTestErrorCode = 0
	// AdTestErrorCodeSignatureMissingTimestamp - The signature is missing a timestamp, in the testing environment.
	//
	// [Full Topic]: https://developer.apple.com/documentation/StoreKitTest/SKAdTestError/Code/signatureMissingTimestamp
	AdTestErrorCodeSignatureMissingTimestamp AdTestErrorCode = 0
	// AdTestErrorCodeSignatureUnknownError - An unknown error occurred with the signature in the testing environment.
	//
	// [Full Topic]: https://developer.apple.com/documentation/StoreKitTest/SKAdTestError/Code/signatureUnknownError
	AdTestErrorCodeSignatureUnknownError AdTestErrorCode = 0
	// AdTestErrorCodeSignatureVerificationFailed - The signature verification failed in the testing environment.
	//
	// [Full Topic]: https://developer.apple.com/documentation/StoreKitTest/SKAdTestError/Code/signatureVerificationFailed
	AdTestErrorCodeSignatureVerificationFailed AdTestErrorCode = 0
	// AdTestErrorCodeUnknownError - An unknown error occurred in the testing environment.
	//
	// [Full Topic]: https://developer.apple.com/documentation/StoreKitTest/SKAdTestError/Code/unknownError
	AdTestErrorCodeUnknownError AdTestErrorCode = 0
	// AdTestErrorCodeUnlinkedWinningPostbacks - The postbacks aren’t correctly related to one another.
	//
	// [Full Topic]: https://developer.apple.com/documentation/StoreKitTest/SKAdTestError/Code/unlinkedWinningPostbacks
	AdTestErrorCodeUnlinkedWinningPostbacks AdTestErrorCode = 0
)

// SKTestErrorCode - Error codes in the testing environment.
//
// [Full Topic]: https://developer.apple.com/documentation/StoreKitTest/SKTestError/Code
type TestErrorCode uint

const (
	// TestErrorCodeFileNotFound - The initializer can’t find the file.
	//
	// [Full Topic]: https://developer.apple.com/documentation/StoreKitTest/SKTestError/Code/fileNotFound
	TestErrorCodeFileNotFound TestErrorCode = 0
	// TestErrorCodeInvalidAction - The action is invalid.
	//
	// [Full Topic]: https://developer.apple.com/documentation/StoreKitTest/SKTestError/Code/invalidAction
	TestErrorCodeInvalidAction TestErrorCode = 0
	// TestErrorCodeInvalidProductIdentifier - The product identifier is invalid.
	//
	// [Full Topic]: https://developer.apple.com/documentation/StoreKitTest/SKTestError/Code/invalidProductIdentifier
	TestErrorCodeInvalidProductIdentifier TestErrorCode = 0
	// TestErrorCodeInvalidProductType - The product type is invalid.
	//
	// [Full Topic]: https://developer.apple.com/documentation/StoreKitTest/SKTestError/Code/invalidProductType
	TestErrorCodeInvalidProductType TestErrorCode = 0
	// TestErrorCodeInvalidURL - The URL is invalid.
	//
	// [Full Topic]: https://developer.apple.com/documentation/StoreKitTest/SKTestError/Code/invalidURL
	TestErrorCodeInvalidURL TestErrorCode = 0
	// TestErrorCodeNoSubscriptionFound - The test environment didn’t find a subscription.
	//
	// [Full Topic]: https://developer.apple.com/documentation/StoreKitTest/SKTestError/Code/noSubscriptionFound
	TestErrorCodeNoSubscriptionFound TestErrorCode = 0
	// TestErrorCodeNoTransactionFound - The test environment didn’t find a transaction.
	//
	// [Full Topic]: https://developer.apple.com/documentation/StoreKitTest/SKTestError/Code/noTransactionFound
	TestErrorCodeNoTransactionFound TestErrorCode = 0
	// TestErrorCodeServiceUnavailable - The service isn’t available.
	//
	// [Full Topic]: https://developer.apple.com/documentation/StoreKitTest/SKTestError/Code/serviceUnavailable
	TestErrorCodeServiceUnavailable TestErrorCode = 0
)

// SKTestTimeRate - The values for rates of time passing in the test environment.
//
// [Full Topic]: https://developer.apple.com/documentation/StoreKitTest/SKTestSession/TimeRate-swift.enum
type TestTimeRate uint

const (
	// TestTimeRateMonthlyRenewalEveryThirtySeconds - A rate of time in the test environment in which monthly subscriptions renew every 30 seconds.
	//
	// [Full Topic]: https://developer.apple.com/documentation/StoreKitTest/SKTestSession/TimeRate-swift.enum/monthlyRenewalEveryThirtySeconds
	TestTimeRateMonthlyRenewalEveryThirtySeconds TestTimeRate = 0
)


