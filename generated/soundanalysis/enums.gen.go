// Code generated from Apple documentation for SoundAnalysis. DO NOT EDIT.

package soundanalysis

/* debug [enums.gen.go]: Generating 2 enums for SoundAnalysis */
// Enum types and constants
/* debug [enums.gen.go]: Processing enum SNErrorCode (5 cases) */
// SNErrorCode - The enumerated error codes that the Sound Analysis framework produces.
//
// [Full Topic]: https://developer.apple.com/documentation/SoundAnalysis/SNError/Code
type SNErrorCode uint

const (
	// SNErrorCodeInvalidFile - An error that indicates an audio file is invalid.
	//
	// [Full Topic]: https://developer.apple.com/documentation/SoundAnalysis/SNError/Code/invalidFile
	SNErrorCodeInvalidFile SNErrorCode = 0
	// SNErrorCodeInvalidFormat - An error that indicates the audio data’s format isn’t valid.
	//
	// [Full Topic]: https://developer.apple.com/documentation/SoundAnalysis/SNError/Code/invalidFormat
	SNErrorCodeInvalidFormat SNErrorCode = 0
	// SNErrorCodeInvalidModel - An error that indicates the sound classifier’s underlying Core ML model is an invalid model type.
	//
	// [Full Topic]: https://developer.apple.com/documentation/SoundAnalysis/SNError/Code/invalidModel
	SNErrorCodeInvalidModel SNErrorCode = 0
	// SNErrorCodeOperationFailed - An error that occurs when the framework fails to analyze audio.
	//
	// [Full Topic]: https://developer.apple.com/documentation/SoundAnalysis/SNError/Code/operationFailed
	SNErrorCodeOperationFailed SNErrorCode = 0
	// SNErrorCodeUnknownError - An error that represents a failure that no other error handles.
	//
	// [Full Topic]: https://developer.apple.com/documentation/SoundAnalysis/SNError/Code/unknownError
	SNErrorCodeUnknownError SNErrorCode = 0
)

/* debug [enums.gen.go]: Processing enum SNTimeDurationConstraintType (2 cases) */
// SNTimeDurationConstraintType - Defines the types a time duration constraint uses.
//
// [Full Topic]: https://developer.apple.com/documentation/SoundAnalysis/SNTimeDurationConstraintType
type SNTimeDurationConstraintType int

const (
	// SNTimeDurationConstraintTypeEnumerated - A constraint type that uses an array of time durations to define what a request’s underlying sound classifier accepts.
	//
	// [Full Topic]: https://developer.apple.com/documentation/SoundAnalysis/SNTimeDurationConstraintType/SNTimeDurationConstraintTypeEnumerated
	SNTimeDurationConstraintTypeEnumerated SNTimeDurationConstraintType = 0
	// SNTimeDurationConstraintTypeRange - A constraint type that uses a time duration range to define what a request’s underlying sound classifier accepts.
	//
	// [Full Topic]: https://developer.apple.com/documentation/SoundAnalysis/SNTimeDurationConstraintType/SNTimeDurationConstraintTypeRange
	SNTimeDurationConstraintTypeRange SNTimeDurationConstraintType = 0
)
