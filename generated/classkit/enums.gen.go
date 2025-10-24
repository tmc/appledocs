// Code generated from Apple documentation for ClassKit. DO NOT EDIT.

package classkit

/* debug [enums.gen.go]: Generating 4 enums for ClassKit */
// Enum types and constants
/* debug [enums.gen.go]: Processing enum CLSErrorCode (0 cases) */
// CLSErrorCode - Error codes that ClassKit issues.
//
// [Full Topic]: https://developer.apple.com/documentation/ClassKit/CLSError/Code
type CLSErrorCode uint

/* debug [enums.gen.go]: Processing enum CLSBinaryValueType (4 cases) */
// CLSBinaryValueType - The kinds of outcomes that a binary activity item can represent.
//
// [Full Topic]: https://developer.apple.com/documentation/ClassKit/CLSBinaryValueType
type CLSBinaryValueType uint

const (
	// CLSBinaryValueTypeCorrectIncorrect - Correct or incorrect.
	//
	// [Full Topic]: https://developer.apple.com/documentation/ClassKit/CLSBinaryValueType/correctIncorrect
	CLSBinaryValueTypeCorrectIncorrect CLSBinaryValueType = 0
	// CLSBinaryValueTypePassFail - Pass or fail.
	//
	// [Full Topic]: https://developer.apple.com/documentation/ClassKit/CLSBinaryValueType/passFail
	CLSBinaryValueTypePassFail CLSBinaryValueType = 0
	// CLSBinaryValueTypeTrueFalse - True or false.
	//
	// [Full Topic]: https://developer.apple.com/documentation/ClassKit/CLSBinaryValueType/trueFalse
	CLSBinaryValueTypeTrueFalse CLSBinaryValueType = 0
	// CLSBinaryValueTypeYesNo - Yes or no.
	//
	// [Full Topic]: https://developer.apple.com/documentation/ClassKit/CLSBinaryValueType/yesNo
	CLSBinaryValueTypeYesNo CLSBinaryValueType = 0
)

/* debug [enums.gen.go]: Processing enum CLSContextType (18 cases) */
// CLSContextType - The kinds of assignable content a context can contain.
//
// [Full Topic]: https://developer.apple.com/documentation/ClassKit/CLSContextType
type CLSContextType uint

const (
	// CLSContextTypeApp - An app context.
	//
	// [Full Topic]: https://developer.apple.com/documentation/ClassKit/CLSContextType/app
	CLSContextTypeApp CLSContextType = 0
	// CLSContextTypeAudio - An audio context.
	//
	// [Full Topic]: https://developer.apple.com/documentation/ClassKit/CLSContextType/audio
	CLSContextTypeAudio CLSContextType = 0
	// CLSContextTypeBook - A book context.
	//
	// [Full Topic]: https://developer.apple.com/documentation/ClassKit/CLSContextType/book
	CLSContextTypeBook CLSContextType = 0
	// CLSContextTypeChallenge - A challenge context.
	//
	// [Full Topic]: https://developer.apple.com/documentation/ClassKit/CLSContextType/challenge
	CLSContextTypeChallenge CLSContextType = 0
	// CLSContextTypeChapter - A chapter context.
	//
	// [Full Topic]: https://developer.apple.com/documentation/ClassKit/CLSContextType/chapter
	CLSContextTypeChapter CLSContextType = 0
	// CLSContextTypeCourse - A context that represents an entire course.
	//
	// [Full Topic]: https://developer.apple.com/documentation/ClassKit/CLSContextType/course
	CLSContextTypeCourse CLSContextType = 0
	// CLSContextTypeCustom - A context for assignable content that isn’t represented by one of the built-in context types.
	//
	// [Full Topic]: https://developer.apple.com/documentation/ClassKit/CLSContextType/custom
	CLSContextTypeCustom CLSContextType = 0
	// CLSContextTypeDocument - A document context.
	//
	// [Full Topic]: https://developer.apple.com/documentation/ClassKit/CLSContextType/document
	CLSContextTypeDocument CLSContextType = 0
	// CLSContextTypeExercise - An exercise context.
	//
	// [Full Topic]: https://developer.apple.com/documentation/ClassKit/CLSContextType/exercise
	CLSContextTypeExercise CLSContextType = 0
	// CLSContextTypeGame - A game context.
	//
	// [Full Topic]: https://developer.apple.com/documentation/ClassKit/CLSContextType/game
	CLSContextTypeGame CLSContextType = 0
	// CLSContextTypeLesson - A lesson context.
	//
	// [Full Topic]: https://developer.apple.com/documentation/ClassKit/CLSContextType/lesson
	CLSContextTypeLesson CLSContextType = 0
	// CLSContextTypeLevel - A level context.
	//
	// [Full Topic]: https://developer.apple.com/documentation/ClassKit/CLSContextType/level
	CLSContextTypeLevel CLSContextType = 0
	// CLSContextTypeNone - No type is assigned.
	//
	// [Full Topic]: https://developer.apple.com/documentation/ClassKit/CLSContextType/none
	CLSContextTypeNone CLSContextType = 0
	// CLSContextTypePage - A page context.
	//
	// [Full Topic]: https://developer.apple.com/documentation/ClassKit/CLSContextType/page
	CLSContextTypePage CLSContextType = 0
	// CLSContextTypeQuiz - A quiz context.
	//
	// [Full Topic]: https://developer.apple.com/documentation/ClassKit/CLSContextType/quiz
	CLSContextTypeQuiz CLSContextType = 0
	// CLSContextTypeSection - A section context.
	//
	// [Full Topic]: https://developer.apple.com/documentation/ClassKit/CLSContextType/section
	CLSContextTypeSection CLSContextType = 0
	// CLSContextTypeTask - A task context.
	//
	// [Full Topic]: https://developer.apple.com/documentation/ClassKit/CLSContextType/task
	CLSContextTypeTask CLSContextType = 0
	// CLSContextTypeVideo - A video context.
	//
	// [Full Topic]: https://developer.apple.com/documentation/ClassKit/CLSContextType/video
	CLSContextTypeVideo CLSContextType = 0
)

/* debug [enums.gen.go]: Processing enum CLSProgressReportingCapabilityKind (5 cases) */
// CLSProgressReportingCapabilityKind - The available kinds of progress reporting that a context can perform.
//
// [Full Topic]: https://developer.apple.com/documentation/ClassKit/CLSProgressReportingCapability/Kind-swift.enum
type CLSProgressReportingCapabilityKind uint

const (
	// CLSProgressReportingCapabilityKindBinary - A binary outcome for the task, like true or false.
	//
	// [Full Topic]: https://developer.apple.com/documentation/ClassKit/CLSProgressReportingCapability/Kind-swift.enum/binary
	CLSProgressReportingCapabilityKindBinary CLSProgressReportingCapabilityKind = 0
	// CLSProgressReportingCapabilityKindDuration - Time spent performing the task.
	//
	// [Full Topic]: https://developer.apple.com/documentation/ClassKit/CLSProgressReportingCapability/Kind-swift.enum/duration
	CLSProgressReportingCapabilityKindDuration CLSProgressReportingCapabilityKind = 0
	// CLSProgressReportingCapabilityKindPercent - The percentage of the total task that has been completed.
	//
	// [Full Topic]: https://developer.apple.com/documentation/ClassKit/CLSProgressReportingCapability/Kind-swift.enum/percent
	CLSProgressReportingCapabilityKindPercent CLSProgressReportingCapabilityKind = 0
	// CLSProgressReportingCapabilityKindQuantity - A discrete value.
	//
	// [Full Topic]: https://developer.apple.com/documentation/ClassKit/CLSProgressReportingCapability/Kind-swift.enum/quantity
	CLSProgressReportingCapabilityKindQuantity CLSProgressReportingCapabilityKind = 0
	// CLSProgressReportingCapabilityKindScore - A score.
	//
	// [Full Topic]: https://developer.apple.com/documentation/ClassKit/CLSProgressReportingCapability/Kind-swift.enum/score
	CLSProgressReportingCapabilityKindScore CLSProgressReportingCapabilityKind = 0
)


