// Code generated from Apple documentation for StoreKitTest. DO NOT EDIT.

package storekittest

/* debug [enums.gen.go]: Generating 3 enums for StoreKitTest */
// Enum types and constants
/* debug [enums.gen.go]: Processing enum SKAdTestErrorCode (33 cases) */
// SKAdTestErrorCode - Enumerated error codes related to ad network testing in the testing environment.
//
// [Full Topic]: https://developer.apple.com/documentation/StoreKitTest/SKAdTestError/Code
type SKAdTestErrorCode uint

const (
	// SKAdTestErrorCodeConflictingSource - This error code is unused.
	//
	// [Full Topic]: https://developer.apple.com/documentation/StoreKitTest/SKAdTestError/Code/conflictingSource
	SKAdTestErrorCodeConflictingSource SKAdTestErrorCode = 0
	// SKAdTestErrorCodeExcessivePostbacks - Too many postbacks submitted to the test session.
	//
	// [Full Topic]: https://developer.apple.com/documentation/StoreKitTest/SKAdTestError/Code/excessivePostbacks
	SKAdTestErrorCodeExcessivePostbacks SKAdTestErrorCode = 0
	// SKAdTestErrorCodeInvalidCampaignId - The campaign ID isn’t an integer between one and one hundred.
	//
	// [Full Topic]: https://developer.apple.com/documentation/StoreKitTest/SKAdTestError/Code/invalidCampaignId
	SKAdTestErrorCodeInvalidCampaignId SKAdTestErrorCode = 0
	// SKAdTestErrorCodeInvalidConversionValue - The conversion value isn’t valid, in the testing environment.
	//
	// [Full Topic]: https://developer.apple.com/documentation/StoreKitTest/SKAdTestError/Code/invalidConversionValue
	SKAdTestErrorCodeInvalidConversionValue SKAdTestErrorCode = 0
	// SKAdTestErrorCodeInvalidImpressionId - The impression ID isn’t a valid UUID string.
	//
	// [Full Topic]: https://developer.apple.com/documentation/StoreKitTest/SKAdTestError/Code/invalidImpressionId
	SKAdTestErrorCodeInvalidImpressionId SKAdTestErrorCode = 0
	// SKAdTestErrorCodeInvalidPostbackURL - The URL for the postback isn’t valid, in the testing environment.
	//
	// [Full Topic]: https://developer.apple.com/documentation/StoreKitTest/SKAdTestError/Code/invalidPostbackURL
	SKAdTestErrorCodeInvalidPostbackURL SKAdTestErrorCode = 0
	// SKAdTestErrorCodeInvalidRunnerUpPostback - A non-winning postback is defined with a version prior to version 3, in the testing environment.
	//
	// [Full Topic]: https://developer.apple.com/documentation/StoreKitTest/SKAdTestError/Code/invalidRunnerUpPostback
	SKAdTestErrorCodeInvalidRunnerUpPostback SKAdTestErrorCode = 0
	// SKAdTestErrorCodeInvalidSourceAppAdamId - The app ID is less than zero.
	//
	// [Full Topic]: https://developer.apple.com/documentation/StoreKitTest/SKAdTestError/Code/invalidSourceAppAdamId
	SKAdTestErrorCodeInvalidSourceAppAdamId SKAdTestErrorCode = 0
	// SKAdTestErrorCodeInvalidSourceDomain - The source domain isn’t in the correct format.
	//
	// [Full Topic]: https://developer.apple.com/documentation/StoreKitTest/SKAdTestError/Code/invalidSourceDomain
	SKAdTestErrorCodeInvalidSourceDomain SKAdTestErrorCode = 0
	// SKAdTestErrorCodeInvalidSourceIdentifier - The postback’s identifier isn’t two, three, or four digits.
	//
	// [Full Topic]: https://developer.apple.com/documentation/StoreKitTest/SKAdTestError/Code/invalidSourceIdentifier
	SKAdTestErrorCodeInvalidSourceIdentifier SKAdTestErrorCode = 0
	// SKAdTestErrorCodeInvalidVersion - A postback contains an incorrect version number.
	//
	// [Full Topic]: https://developer.apple.com/documentation/StoreKitTest/SKAdTestError/Code/invalidVersion
	SKAdTestErrorCodeInvalidVersion SKAdTestErrorCode = 0
	// SKAdTestErrorCodeInvalidWinningPostbackCount - The number of winning postbacks isn’t valid, in the testing environment.
	//
	// [Full Topic]: https://developer.apple.com/documentation/StoreKitTest/SKAdTestError/Code/invalidWinningPostbackCount
	SKAdTestErrorCodeInvalidWinningPostbackCount SKAdTestErrorCode = 0
	// SKAdTestErrorCodeMalformedPostbacks - The postback in the testing environment is malformed.
	//
	// [Full Topic]: https://developer.apple.com/documentation/StoreKitTest/SKAdTestError/Code/malformedPostbacks
	SKAdTestErrorCodeMalformedPostbacks SKAdTestErrorCode = 0
	// SKAdTestErrorCodeMisplacedWinnerPostback - A winning postback wasn’t found in the first position, in the testing environment.
	//
	// [Full Topic]: https://developer.apple.com/documentation/StoreKitTest/SKAdTestError/Code/misplacedWinnerPostback
	SKAdTestErrorCodeMisplacedWinnerPostback SKAdTestErrorCode = 0
	// SKAdTestErrorCodeMissingPostbacks - The testing environment doesn’t have any postbacks.
	//
	// [Full Topic]: https://developer.apple.com/documentation/StoreKitTest/SKAdTestError/Code/missingPostbacks
	SKAdTestErrorCodeMissingPostbacks SKAdTestErrorCode = 0
	// SKAdTestErrorCodeMissingSignature - The signature for the ad is missing, in the testing environment.
	//
	// [Full Topic]: https://developer.apple.com/documentation/StoreKitTest/SKAdTestError/Code/missingSignature
	SKAdTestErrorCodeMissingSignature SKAdTestErrorCode = 0
	// SKAdTestErrorCodeMissingWinningPostback - The testing environment is missing a winning postback.
	//
	// [Full Topic]: https://developer.apple.com/documentation/StoreKitTest/SKAdTestError/Code/missingWinningPostback
	SKAdTestErrorCodeMissingWinningPostback SKAdTestErrorCode = 0
	// SKAdTestErrorCodeNoPendingPostbacks - The test session doesn’t have any pending postbacks to send.
	//
	// [Full Topic]: https://developer.apple.com/documentation/StoreKitTest/SKAdTestError/Code/noPendingPostbacks
	SKAdTestErrorCodeNoPendingPostbacks SKAdTestErrorCode = 0
	// SKAdTestErrorCodeSignatureInvalidKey - The public key isn’t a valid cryptographic key, in the testing environment.
	//
	// [Full Topic]: https://developer.apple.com/documentation/StoreKitTest/SKAdTestError/Code/signatureInvalidKey
	SKAdTestErrorCodeSignatureInvalidKey SKAdTestErrorCode = 0
	// SKAdTestErrorCodeSignatureInvalidOrder - The order of the parameters in the signature is invalid, in the testing environment.
	//
	// [Full Topic]: https://developer.apple.com/documentation/StoreKitTest/SKAdTestError/Code/signatureInvalidOrder
	SKAdTestErrorCodeSignatureInvalidOrder SKAdTestErrorCode = 0
	// SKAdTestErrorCodeSignatureMissingAdNetworkId - The signature is missing an ad network identifier, in the testing environment.
	//
	// [Full Topic]: https://developer.apple.com/documentation/StoreKitTest/SKAdTestError/Code/signatureMissingAdNetworkId
	SKAdTestErrorCodeSignatureMissingAdNetworkId SKAdTestErrorCode = 0
	// SKAdTestErrorCodeSignatureMissingAppAdamId - The signature is missing the app item identifier for the advertised app, in the testing environment.
	//
	// [Full Topic]: https://developer.apple.com/documentation/StoreKitTest/SKAdTestError/Code/signatureMissingAppAdamId
	SKAdTestErrorCodeSignatureMissingAppAdamId SKAdTestErrorCode = 0
	// SKAdTestErrorCodeSignatureMissingCampaignId - The signature is missing the campaign identifier, in the testing environment.
	//
	// [Full Topic]: https://developer.apple.com/documentation/StoreKitTest/SKAdTestError/Code/signatureMissingCampaignId
	SKAdTestErrorCodeSignatureMissingCampaignId SKAdTestErrorCode = 0
	// SKAdTestErrorCodeSignatureMissingFidelityType - The signature is missing the fidelity type, in the testing environment.
	//
	// [Full Topic]: https://developer.apple.com/documentation/StoreKitTest/SKAdTestError/Code/signatureMissingFidelityType
	SKAdTestErrorCodeSignatureMissingFidelityType SKAdTestErrorCode = 0
	// SKAdTestErrorCodeSignatureMissingNonce - The signature is missing the nonce, in the testing environment.
	//
	// [Full Topic]: https://developer.apple.com/documentation/StoreKitTest/SKAdTestError/Code/signatureMissingNonce
	SKAdTestErrorCodeSignatureMissingNonce SKAdTestErrorCode = 0
	// SKAdTestErrorCodeSignatureMissingSourceAppAdamId - The signature is missing the source app item identifier, in the testing environment.
	//
	// [Full Topic]: https://developer.apple.com/documentation/StoreKitTest/SKAdTestError/Code/signatureMissingSourceAppAdamId
	SKAdTestErrorCodeSignatureMissingSourceAppAdamId SKAdTestErrorCode = 0
	// SKAdTestErrorCodeSignatureMissingSourceDomain - The signature is missing the source domain, in the testing environment.
	//
	// [Full Topic]: https://developer.apple.com/documentation/StoreKitTest/SKAdTestError/Code/signatureMissingSourceDomain
	SKAdTestErrorCodeSignatureMissingSourceDomain SKAdTestErrorCode = 0
	// SKAdTestErrorCodeSignatureMissingSourceIdentifier - The signature is missing the source identifier, in the testing environment.
	//
	// [Full Topic]: https://developer.apple.com/documentation/StoreKitTest/SKAdTestError/Code/signatureMissingSourceIdentifier
	SKAdTestErrorCodeSignatureMissingSourceIdentifier SKAdTestErrorCode = 0
	// SKAdTestErrorCodeSignatureMissingTimestamp - The signature is missing a timestamp, in the testing environment.
	//
	// [Full Topic]: https://developer.apple.com/documentation/StoreKitTest/SKAdTestError/Code/signatureMissingTimestamp
	SKAdTestErrorCodeSignatureMissingTimestamp SKAdTestErrorCode = 0
	// SKAdTestErrorCodeSignatureUnknownError - An unknown error occurred with the signature in the testing environment.
	//
	// [Full Topic]: https://developer.apple.com/documentation/StoreKitTest/SKAdTestError/Code/signatureUnknownError
	SKAdTestErrorCodeSignatureUnknownError SKAdTestErrorCode = 0
	// SKAdTestErrorCodeSignatureVerificationFailed - The signature verification failed in the testing environment.
	//
	// [Full Topic]: https://developer.apple.com/documentation/StoreKitTest/SKAdTestError/Code/signatureVerificationFailed
	SKAdTestErrorCodeSignatureVerificationFailed SKAdTestErrorCode = 0
	// SKAdTestErrorCodeUnknownError - An unknown error occurred in the testing environment.
	//
	// [Full Topic]: https://developer.apple.com/documentation/StoreKitTest/SKAdTestError/Code/unknownError
	SKAdTestErrorCodeUnknownError SKAdTestErrorCode = 0
	// SKAdTestErrorCodeUnlinkedWinningPostbacks - The postbacks aren’t correctly related to one another.
	//
	// [Full Topic]: https://developer.apple.com/documentation/StoreKitTest/SKAdTestError/Code/unlinkedWinningPostbacks
	SKAdTestErrorCodeUnlinkedWinningPostbacks SKAdTestErrorCode = 0
)

/* debug [enums.gen.go]: Processing enum SKTestErrorCode (8 cases) */
// SKTestErrorCode - Error codes in the testing environment.
//
// [Full Topic]: https://developer.apple.com/documentation/StoreKitTest/SKTestError/Code
type SKTestErrorCode uint

const (
	// SKTestErrorCodeFileNotFound - The initializer can’t find the file.
	//
	// [Full Topic]: https://developer.apple.com/documentation/StoreKitTest/SKTestError/Code/fileNotFound
	SKTestErrorCodeFileNotFound SKTestErrorCode = 0
	// SKTestErrorCodeInvalidAction - The action is invalid.
	//
	// [Full Topic]: https://developer.apple.com/documentation/StoreKitTest/SKTestError/Code/invalidAction
	SKTestErrorCodeInvalidAction SKTestErrorCode = 0
	// SKTestErrorCodeInvalidProductIdentifier - The product identifier is invalid.
	//
	// [Full Topic]: https://developer.apple.com/documentation/StoreKitTest/SKTestError/Code/invalidProductIdentifier
	SKTestErrorCodeInvalidProductIdentifier SKTestErrorCode = 0
	// SKTestErrorCodeInvalidProductType - The product type is invalid.
	//
	// [Full Topic]: https://developer.apple.com/documentation/StoreKitTest/SKTestError/Code/invalidProductType
	SKTestErrorCodeInvalidProductType SKTestErrorCode = 0
	// SKTestErrorCodeInvalidURL - The URL is invalid.
	//
	// [Full Topic]: https://developer.apple.com/documentation/StoreKitTest/SKTestError/Code/invalidURL
	SKTestErrorCodeInvalidURL SKTestErrorCode = 0
	// SKTestErrorCodeNoSubscriptionFound - The test environment didn’t find a subscription.
	//
	// [Full Topic]: https://developer.apple.com/documentation/StoreKitTest/SKTestError/Code/noSubscriptionFound
	SKTestErrorCodeNoSubscriptionFound SKTestErrorCode = 0
	// SKTestErrorCodeNoTransactionFound - The test environment didn’t find a transaction.
	//
	// [Full Topic]: https://developer.apple.com/documentation/StoreKitTest/SKTestError/Code/noTransactionFound
	SKTestErrorCodeNoTransactionFound SKTestErrorCode = 0
	// SKTestErrorCodeServiceUnavailable - The service isn’t available.
	//
	// [Full Topic]: https://developer.apple.com/documentation/StoreKitTest/SKTestError/Code/serviceUnavailable
	SKTestErrorCodeServiceUnavailable SKTestErrorCode = 0
)

/* debug [enums.gen.go]: Processing enum SKTestTimeRate (18 cases) */
// SKTestTimeRate - The values for rates of time passing in the test environment.
//
// [Full Topic]: https://developer.apple.com/documentation/StoreKitTest/SKTestSession/TimeRate-swift.enum
type SKTestTimeRate uint

const (
	// SKTestTimeRateFiveMinutesIsOneDay - A rate of time in which 5 minutes in the test environment represents one day.
	//
	// [Full Topic]: https://developer.apple.com/documentation/StoreKitTest/SKTestSession/TimeRate-swift.enum/fiveMinutesIsOneDay
	SKTestTimeRateFiveMinutesIsOneDay SKTestTimeRate = 0
	// SKTestTimeRateMonthlyRenewalEveryFifteenMinutes - A rate of time in the test environment in which monthly subscriptions renew every 15 minutes.
	//
	// [Full Topic]: https://developer.apple.com/documentation/StoreKitTest/SKTestSession/TimeRate-swift.enum/monthlyRenewalEveryFifteenMinutes
	SKTestTimeRateMonthlyRenewalEveryFifteenMinutes SKTestTimeRate = 0
	// SKTestTimeRateMonthlyRenewalEveryFiveMinutes - A rate of time in the test environment in which monthly subscriptions renew every 5 minutes.
	//
	// [Full Topic]: https://developer.apple.com/documentation/StoreKitTest/SKTestSession/TimeRate-swift.enum/monthlyRenewalEveryFiveMinutes
	SKTestTimeRateMonthlyRenewalEveryFiveMinutes SKTestTimeRate = 0
	// SKTestTimeRateMonthlyRenewalEveryHour - A rate of time in the test environment in which monthly subscriptions renew every hour.
	//
	// [Full Topic]: https://developer.apple.com/documentation/StoreKitTest/SKTestSession/TimeRate-swift.enum/monthlyRenewalEveryHour
	SKTestTimeRateMonthlyRenewalEveryHour SKTestTimeRate = 0
	// SKTestTimeRateMonthlyRenewalEveryThirtyMinutes - A rate of time in the test environment in which monthly subscriptions renew every 30 minutes.
	//
	// [Full Topic]: https://developer.apple.com/documentation/StoreKitTest/SKTestSession/TimeRate-swift.enum/monthlyRenewalEveryThirtyMinutes
	SKTestTimeRateMonthlyRenewalEveryThirtyMinutes SKTestTimeRate = 0
	// SKTestTimeRateMonthlyRenewalEveryThirtySeconds - A rate of time in the test environment in which monthly subscriptions renew every 30 seconds.
	//
	// [Full Topic]: https://developer.apple.com/documentation/StoreKitTest/SKTestSession/TimeRate-swift.enum/monthlyRenewalEveryThirtySeconds
	SKTestTimeRateMonthlyRenewalEveryThirtySeconds SKTestTimeRate = 0
	// SKTestTimeRateOneHourIsOneDay - A rate of time in which 1 hour in the test environment represents one day.
	//
	// [Full Topic]: https://developer.apple.com/documentation/StoreKitTest/SKTestSession/TimeRate-swift.enum/oneHourIsOneDay
	SKTestTimeRateOneHourIsOneDay SKTestTimeRate = 0
	// SKTestTimeRateOneMinuteIsOneDay - A rate of time in which 1 minute in the test environment represents one day.
	//
	// [Full Topic]: https://developer.apple.com/documentation/StoreKitTest/SKTestSession/TimeRate-swift.enum/oneMinuteIsOneDay
	SKTestTimeRateOneMinuteIsOneDay SKTestTimeRate = 0
	// SKTestTimeRateOneRenewalEveryFifteenMinutes - A rate of time in the test environment in which subscriptions of any time length renew every 15 minutes.
	//
	// [Full Topic]: https://developer.apple.com/documentation/StoreKitTest/SKTestSession/TimeRate-swift.enum/oneRenewalEveryFifteenMinutes
	SKTestTimeRateOneRenewalEveryFifteenMinutes SKTestTimeRate = 0
	// SKTestTimeRateOneRenewalEveryFiveMinutes - A rate of time in the test environment in which subscriptions of any time length renew every 5 minutes.
	//
	// [Full Topic]: https://developer.apple.com/documentation/StoreKitTest/SKTestSession/TimeRate-swift.enum/oneRenewalEveryFiveMinutes
	SKTestTimeRateOneRenewalEveryFiveMinutes SKTestTimeRate = 0
	// SKTestTimeRateOneRenewalEveryMinute - A rate of time in the test environment in which subscriptions of any time length renew every minute.
	//
	// [Full Topic]: https://developer.apple.com/documentation/StoreKitTest/SKTestSession/TimeRate-swift.enum/oneRenewalEveryMinute
	SKTestTimeRateOneRenewalEveryMinute SKTestTimeRate = 0
	// SKTestTimeRateOneRenewalEveryTenSeconds - A rate of time in the test environment in which subscriptions of any time length renew every 10 seconds.
	//
	// [Full Topic]: https://developer.apple.com/documentation/StoreKitTest/SKTestSession/TimeRate-swift.enum/oneRenewalEveryTenSeconds
	SKTestTimeRateOneRenewalEveryTenSeconds SKTestTimeRate = 0
	// SKTestTimeRateOneRenewalEveryThirtySeconds - A rate of time in the test environment in which subscriptions of any time length renew every 30 seconds.
	//
	// [Full Topic]: https://developer.apple.com/documentation/StoreKitTest/SKTestSession/TimeRate-swift.enum/oneRenewalEveryThirtySeconds
	SKTestTimeRateOneRenewalEveryThirtySeconds SKTestTimeRate = 0
	// SKTestTimeRateOneRenewalEveryTwoSeconds - A rate of time in the test environment in which subscriptions of any time length renew every 2 seconds.
	//
	// [Full Topic]: https://developer.apple.com/documentation/StoreKitTest/SKTestSession/TimeRate-swift.enum/oneRenewalEveryTwoSeconds
	SKTestTimeRateOneRenewalEveryTwoSeconds SKTestTimeRate = 0
	// SKTestTimeRateOneSecondIsOneDay - A rate of time in which 1 second in the test environment represents one day.
	//
	// [Full Topic]: https://developer.apple.com/documentation/StoreKitTest/SKTestSession/TimeRate-swift.enum/oneSecondIsOneDay
	SKTestTimeRateOneSecondIsOneDay SKTestTimeRate = 0
	// SKTestTimeRateRealTime - A rate of time in which the test environment runs in real time.
	//
	// [Full Topic]: https://developer.apple.com/documentation/StoreKitTest/SKTestSession/TimeRate-swift.enum/realTime
	SKTestTimeRateRealTime SKTestTimeRate = 0
	// SKTestTimeRateThirtyMinutesIsOneDay - A rate of time in which 30 minutes in the test environment represents one day.
	//
	// [Full Topic]: https://developer.apple.com/documentation/StoreKitTest/SKTestSession/TimeRate-swift.enum/thirtyMinutesIsOneDay
	SKTestTimeRateThirtyMinutesIsOneDay SKTestTimeRate = 0
	// SKTestTimeRateThirtySecondsIsOneDay - A rate of time in which 30 seconds in the test environment represents one day.
	//
	// [Full Topic]: https://developer.apple.com/documentation/StoreKitTest/SKTestSession/TimeRate-swift.enum/thirtySecondsIsOneDay
	SKTestTimeRateThirtySecondsIsOneDay SKTestTimeRate = 0
)


