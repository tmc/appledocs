// Code generated from Apple documentation for ClassKit. DO NOT EDIT.

package classkit

// Enum types and constants
// CLSBinaryValueType - The kinds of outcomes that a binary activity item can represent.
//
// [Full Topic]: https://developer.apple.com/documentation/ClassKit/CLSBinaryValueType
type SBinaryValueType uint

const (
// SBinaryValueTypeCorrectIncorrect - Correct or incorrect.
//
	// [Full Topic]: https://developer.apple.com/documentation/ClassKit/CLSBinaryValueType/correctIncorrect
SBinaryValueTypeCorrectIncorrect SBinaryValueType = 0
// SBinaryValueTypePassFail - Pass or fail.
//
	// [Full Topic]: https://developer.apple.com/documentation/ClassKit/CLSBinaryValueType/passFail
SBinaryValueTypePassFail SBinaryValueType = 0
// SBinaryValueTypeTrueFalse - True or false.
//
	// [Full Topic]: https://developer.apple.com/documentation/ClassKit/CLSBinaryValueType/trueFalse
SBinaryValueTypeTrueFalse SBinaryValueType = 0
// SBinaryValueTypeYesNo - Yes or no.
//
	// [Full Topic]: https://developer.apple.com/documentation/ClassKit/CLSBinaryValueType/yesNo
SBinaryValueTypeYesNo SBinaryValueType = 0
)

// CLSContextType - The kinds of assignable content a context can contain.
//
// [Full Topic]: https://developer.apple.com/documentation/ClassKit/CLSContextType
type SContextType uint

const (
// SContextTypeApp - An app context.
//
	// [Full Topic]: https://developer.apple.com/documentation/ClassKit/CLSContextType/app
SContextTypeApp SContextType = 0
// SContextTypeAudio - An audio context.
//
	// [Full Topic]: https://developer.apple.com/documentation/ClassKit/CLSContextType/audio
SContextTypeAudio SContextType = 0
// SContextTypeBook - A book context.
//
	// [Full Topic]: https://developer.apple.com/documentation/ClassKit/CLSContextType/book
SContextTypeBook SContextType = 0
// SContextTypeChallenge - A challenge context.
//
	// [Full Topic]: https://developer.apple.com/documentation/ClassKit/CLSContextType/challenge
SContextTypeChallenge SContextType = 0
// SContextTypeChapter - A chapter context.
//
	// [Full Topic]: https://developer.apple.com/documentation/ClassKit/CLSContextType/chapter
SContextTypeChapter SContextType = 0
// SContextTypeCourse - A context that represents an entire course.
//
	// [Full Topic]: https://developer.apple.com/documentation/ClassKit/CLSContextType/course
SContextTypeCourse SContextType = 0
// SContextTypeCustom - A context for assignable content that isn’t represented by one of the built-in context types.
//
	// [Full Topic]: https://developer.apple.com/documentation/ClassKit/CLSContextType/custom
SContextTypeCustom SContextType = 0
// SContextTypeDocument - A document context.
//
	// [Full Topic]: https://developer.apple.com/documentation/ClassKit/CLSContextType/document
SContextTypeDocument SContextType = 0
// SContextTypeExercise - An exercise context.
//
	// [Full Topic]: https://developer.apple.com/documentation/ClassKit/CLSContextType/exercise
SContextTypeExercise SContextType = 0
// SContextTypeGame - A game context.
//
	// [Full Topic]: https://developer.apple.com/documentation/ClassKit/CLSContextType/game
SContextTypeGame SContextType = 0
// SContextTypeLesson - A lesson context.
//
	// [Full Topic]: https://developer.apple.com/documentation/ClassKit/CLSContextType/lesson
SContextTypeLesson SContextType = 0
// SContextTypeLevel - A level context.
//
	// [Full Topic]: https://developer.apple.com/documentation/ClassKit/CLSContextType/level
SContextTypeLevel SContextType = 0
// SContextTypeNone - No type is assigned.
//
	// [Full Topic]: https://developer.apple.com/documentation/ClassKit/CLSContextType/none
SContextTypeNone SContextType = 0
// SContextTypePage - A page context.
//
	// [Full Topic]: https://developer.apple.com/documentation/ClassKit/CLSContextType/page
SContextTypePage SContextType = 0
// SContextTypeQuiz - A quiz context.
//
	// [Full Topic]: https://developer.apple.com/documentation/ClassKit/CLSContextType/quiz
SContextTypeQuiz SContextType = 0
// SContextTypeSection - A section context.
//
	// [Full Topic]: https://developer.apple.com/documentation/ClassKit/CLSContextType/section
SContextTypeSection SContextType = 0
// SContextTypeTask - A task context.
//
	// [Full Topic]: https://developer.apple.com/documentation/ClassKit/CLSContextType/task
SContextTypeTask SContextType = 0
// SContextTypeVideo - A video context.
//
	// [Full Topic]: https://developer.apple.com/documentation/ClassKit/CLSContextType/video
SContextTypeVideo SContextType = 0
)

// CLSErrorCode - Error codes that ClassKit issues.
//
// [Full Topic]: https://developer.apple.com/documentation/ClassKit/CLSError/Code
type SErrorCode uint

const (
// SErrorCodeAuthorizationDenied - The app isn’t authorized to perform the requested operation.
//
	// [Full Topic]: https://developer.apple.com/documentation/ClassKit/CLSError/Code/authorizationDenied
SErrorCodeAuthorizationDenied SErrorCode = 0
// SErrorCodeClassKitUnavailable - ClassKit isn’t available on this device.
//
	// [Full Topic]: https://developer.apple.com/documentation/ClassKit/CLSError/Code/classKitUnavailable
SErrorCodeClassKitUnavailable SErrorCode = 0
// SErrorCodeDatabaseInaccessible - ClassKit isn’t accessible because the device is locked.
//
	// [Full Topic]: https://developer.apple.com/documentation/ClassKit/CLSError/Code/databaseInaccessible
SErrorCodeDatabaseInaccessible SErrorCode = 0
//
	// [Full Topic]: https://developer.apple.com/documentation/ClassKit/CLSError/Code/invalidAccountCredentials
SErrorCodeInvalidAccountCredentials SErrorCode = 0
// SErrorCodeInvalidArgument - An invalid argument was provided to the API.
//
	// [Full Topic]: https://developer.apple.com/documentation/ClassKit/CLSError/Code/invalidArgument
SErrorCodeInvalidArgument SErrorCode = 0
// SErrorCodeInvalidCreate - An attempt to save a new object that already exists in the data store failed.
//
	// [Full Topic]: https://developer.apple.com/documentation/ClassKit/CLSError/Code/invalidCreate
SErrorCodeInvalidCreate SErrorCode = 0
// SErrorCodeInvalidModification - An attempt to modify a read-only object failed.
//
	// [Full Topic]: https://developer.apple.com/documentation/ClassKit/CLSError/Code/invalidModification
SErrorCodeInvalidModification SErrorCode = 0
// SErrorCodeInvalidUpdate - ClassKit failed to save an updated object in the data store.
//
	// [Full Topic]: https://developer.apple.com/documentation/ClassKit/CLSError/Code/invalidUpdate
SErrorCodeInvalidUpdate SErrorCode = 0
// SErrorCodeLimits - A limit has been exceeded.
//
	// [Full Topic]: https://developer.apple.com/documentation/ClassKit/CLSError/Code/limits
SErrorCodeLimits SErrorCode = 0
// SErrorCodeNone - No error.
//
	// [Full Topic]: https://developer.apple.com/documentation/ClassKit/CLSError/Code/none
SErrorCodeNone SErrorCode = 0
// SErrorCodePartialFailure - ClassKit encountered more than one error.
//
	// [Full Topic]: https://developer.apple.com/documentation/ClassKit/CLSError/Code/partialFailure
SErrorCodePartialFailure SErrorCode = 0
)

// CLSProgressReportingCapabilityKind - The available kinds of progress reporting that a context can perform.
//
// [Full Topic]: https://developer.apple.com/documentation/ClassKit/CLSProgressReportingCapability/Kind-swift.enum
type SProgressReportingCapabilityKind uint

const (
// SProgressReportingCapabilityKindBinary - A binary outcome for the task, like true or false.
//
	// [Full Topic]: https://developer.apple.com/documentation/ClassKit/CLSProgressReportingCapability/Kind-swift.enum/binary
SProgressReportingCapabilityKindBinary SProgressReportingCapabilityKind = 0
// SProgressReportingCapabilityKindDuration - Time spent performing the task.
//
	// [Full Topic]: https://developer.apple.com/documentation/ClassKit/CLSProgressReportingCapability/Kind-swift.enum/duration
SProgressReportingCapabilityKindDuration SProgressReportingCapabilityKind = 0
// SProgressReportingCapabilityKindPercent - The percentage of the total task that has been completed.
//
	// [Full Topic]: https://developer.apple.com/documentation/ClassKit/CLSProgressReportingCapability/Kind-swift.enum/percent
SProgressReportingCapabilityKindPercent SProgressReportingCapabilityKind = 0
// SProgressReportingCapabilityKindQuantity - A discrete value.
//
	// [Full Topic]: https://developer.apple.com/documentation/ClassKit/CLSProgressReportingCapability/Kind-swift.enum/quantity
SProgressReportingCapabilityKindQuantity SProgressReportingCapabilityKind = 0
// SProgressReportingCapabilityKindScore - A score.
//
	// [Full Topic]: https://developer.apple.com/documentation/ClassKit/CLSProgressReportingCapability/Kind-swift.enum/score
SProgressReportingCapabilityKindScore SProgressReportingCapabilityKind = 0
)


