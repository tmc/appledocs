// Code generated from Apple documentation for Speech. DO NOT EDIT.

package speech

/* debug [enums.gen.go]: Generating 4 enums for Speech */
// Enum types and constants
/* debug [enums.gen.go]: Processing enum SFSpeechErrorCode (6 cases) */
// SFSpeechErrorCode - Error codes that can be thrown under the Speech framework’s error domain.
//
// [Full Topic]: https://developer.apple.com/documentation/Speech/SFSpeechError/Code
type SFSpeechErrorCode uint

const (
	// SFSpeechErrorCodeAudioReadFailed - The audio file could not be read.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Speech/SFSpeechError/Code/audioReadFailed
	SFSpeechErrorCodeAudioReadFailed SFSpeechErrorCode = 0
	// SFSpeechErrorCodeInternalServiceError - There was an internal error.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Speech/SFSpeechError/Code/internalServiceError
	SFSpeechErrorCodeInternalServiceError SFSpeechErrorCode = 0
	// SFSpeechErrorCodeMalformedSupplementalModel - The custom language model file was malformed.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Speech/SFSpeechError/Code/malformedSupplementalModel
	SFSpeechErrorCodeMalformedSupplementalModel SFSpeechErrorCode = 0
	// SFSpeechErrorCodeMissingParameter - A required parameter is missing/nil.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Speech/SFSpeechError/Code/missingParameter
	SFSpeechErrorCodeMissingParameter SFSpeechErrorCode = 0
	// SFSpeechErrorCodeTimeout - The operation timed out.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Speech/SFSpeechError/Code/timeout
	SFSpeechErrorCodeTimeout SFSpeechErrorCode = 0
	// SFSpeechErrorCodeUndefinedTemplateClassName - The custom language model templates were malformed.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Speech/SFSpeechError/Code/undefinedTemplateClassName
	SFSpeechErrorCodeUndefinedTemplateClassName SFSpeechErrorCode = 0
)

/* debug [enums.gen.go]: Processing enum SFSpeechRecognitionTaskHint (4 cases) */
// SFSpeechRecognitionTaskHint - The type of task for which you are using speech recognition.
//
// [Full Topic]: https://developer.apple.com/documentation/Speech/SFSpeechRecognitionTaskHint
type SFSpeechRecognitionTaskHint uint

const (
	// SFSpeechRecognitionTaskHintConfirmation - A task that uses captured speech for short, confirmation-style requests.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Speech/SFSpeechRecognitionTaskHint/confirmation
	SFSpeechRecognitionTaskHintConfirmation SFSpeechRecognitionTaskHint = 0
	// SFSpeechRecognitionTaskHintDictation - A task that uses captured speech for text entry.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Speech/SFSpeechRecognitionTaskHint/dictation
	SFSpeechRecognitionTaskHintDictation SFSpeechRecognitionTaskHint = 0
	// SFSpeechRecognitionTaskHintSearch - A task that uses captured speech to specify search terms.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Speech/SFSpeechRecognitionTaskHint/search
	SFSpeechRecognitionTaskHintSearch SFSpeechRecognitionTaskHint = 0
	// SFSpeechRecognitionTaskHintUnspecified - An unspecified type of task.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Speech/SFSpeechRecognitionTaskHint/unspecified
	SFSpeechRecognitionTaskHintUnspecified SFSpeechRecognitionTaskHint = 0
)

/* debug [enums.gen.go]: Processing enum SFSpeechRecognitionTaskState (5 cases) */
// SFSpeechRecognitionTaskState - The state of the task associated with the recognition request.
//
// [Full Topic]: https://developer.apple.com/documentation/Speech/SFSpeechRecognitionTaskState
type SFSpeechRecognitionTaskState uint

const (
	// SFSpeechRecognitionTaskStateCanceling - Delivery of recognition results has finished, but audio recording may be ongoing.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Speech/SFSpeechRecognitionTaskState/canceling
	SFSpeechRecognitionTaskStateCanceling SFSpeechRecognitionTaskState = 0
	// SFSpeechRecognitionTaskStateCompleted - Delivery of recognition requests has finished and audio recording has stopped.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Speech/SFSpeechRecognitionTaskState/completed
	SFSpeechRecognitionTaskStateCompleted SFSpeechRecognitionTaskState = 0
	// SFSpeechRecognitionTaskStateFinishing - Audio recording has stopped, but delivery of recognition results may continue.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Speech/SFSpeechRecognitionTaskState/finishing
	SFSpeechRecognitionTaskStateFinishing SFSpeechRecognitionTaskState = 0
	// SFSpeechRecognitionTaskStateRunning - Speech recognition (potentially including audio recording) is in progress.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Speech/SFSpeechRecognitionTaskState/running
	SFSpeechRecognitionTaskStateRunning SFSpeechRecognitionTaskState = 0
	// SFSpeechRecognitionTaskStateStarting - Speech recognition (potentially including audio recording) has not yet started.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Speech/SFSpeechRecognitionTaskState/starting
	SFSpeechRecognitionTaskStateStarting SFSpeechRecognitionTaskState = 0
)

/* debug [enums.gen.go]: Processing enum SFSpeechRecognizerAuthorizationStatus (4 cases) */
// SFSpeechRecognizerAuthorizationStatus - The app’s authorization to perform speech recognition.
//
// [Full Topic]: https://developer.apple.com/documentation/Speech/SFSpeechRecognizerAuthorizationStatus
type SFSpeechRecognizerAuthorizationStatus uint

const (
	// SFSpeechRecognizerAuthorizationStatusAuthorized - The user granted your app’s request to perform speech recognition.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Speech/SFSpeechRecognizerAuthorizationStatus/authorized
	SFSpeechRecognizerAuthorizationStatusAuthorized SFSpeechRecognizerAuthorizationStatus = 0
	// SFSpeechRecognizerAuthorizationStatusDenied - The user denied your app’s request to perform speech recognition.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Speech/SFSpeechRecognizerAuthorizationStatus/denied
	SFSpeechRecognizerAuthorizationStatusDenied SFSpeechRecognizerAuthorizationStatus = 0
	// SFSpeechRecognizerAuthorizationStatusNotDetermined - The app’s authorization status has not yet been determined.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Speech/SFSpeechRecognizerAuthorizationStatus/notDetermined
	SFSpeechRecognizerAuthorizationStatusNotDetermined SFSpeechRecognizerAuthorizationStatus = 0
	// SFSpeechRecognizerAuthorizationStatusRestricted - The device prevents your app from performing speech recognition.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Speech/SFSpeechRecognizerAuthorizationStatus/restricted
	SFSpeechRecognizerAuthorizationStatusRestricted SFSpeechRecognizerAuthorizationStatus = 0
)
