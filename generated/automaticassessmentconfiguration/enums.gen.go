// Code generated from Apple documentation for AutomaticAssessmentConfiguration. DO NOT EDIT.

package automaticassessmentconfiguration

// Enum types and constants
// AEAutocorrectMode - The set of autocorrect features that you can enable during an assessment.
//
// [Full Topic]: https://developer.apple.com/documentation/AutomaticAssessmentConfiguration/AEAssessmentConfiguration/AutocorrectMode-swift.struct
type AEAutocorrectMode uint

// AEAssessmentErrorCode - Error codes that the framework returns if a session fails.
//
// [Full Topic]: https://developer.apple.com/documentation/AutomaticAssessmentConfiguration/AEAssessmentError/Code
type AEAssessmentErrorCode uint

const (
	// AEAssessmentErrorConfigurationUpdatesNotSupported - An active session fails to update its configuration because configuration updates are not supported by the current device or platform.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AutomaticAssessmentConfiguration/AEAssessmentError/Code/configurationUpdatesNotSupported
	AEAssessmentErrorConfigurationUpdatesNotSupported AEAssessmentErrorCode = 0
	// AEAssessmentErrorMultipleParticipantsNotSupported - A session fails to begin or update with a configuration that contains one or more participant applications because mulitple participant configurations are not supported by the device or platform.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AutomaticAssessmentConfiguration/AEAssessmentError/Code/multipleParticipantsNotSupported
	AEAssessmentErrorMultipleParticipantsNotSupported AEAssessmentErrorCode = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/AutomaticAssessmentConfiguration/AEAssessmentError/Code/requiredParticipantsNotAvailable
	AEAssessmentErrorRequiredParticipantsNotAvailable AEAssessmentErrorCode = 0
	// AEAssessmentErrorUnknown - The session encountered an unknown error.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AutomaticAssessmentConfiguration/AEAssessmentError/Code/unknown
	AEAssessmentErrorUnknown AEAssessmentErrorCode = 0
	// AEAssessmentErrorUnsupportedPlatform - The feature isn’t supported on this platform.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AutomaticAssessmentConfiguration/AEAssessmentError/Code/unsupportedPlatform
	AEAssessmentErrorUnsupportedPlatform AEAssessmentErrorCode = 0
)


