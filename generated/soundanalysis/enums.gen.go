// Code generated from Apple documentation for SoundAnalysis. DO NOT EDIT.

package soundanalysis

// Enum types and constants
// SNErrorCode - The enumerated error codes that the Sound Analysis framework produces.
//
// [Full Topic]: https://developer.apple.com/documentation/SoundAnalysis/SNError/Code
type SNErrorCode uint

const (
	// SNErrorCodeInvalidModel - An error that indicates the sound classifier’s underlying Core ML model is an invalid model type.
	//
	// [Full Topic]: https://developer.apple.com/documentation/SoundAnalysis/SNError/Code/invalidModel
	SNErrorCodeInvalidModel SNErrorCode = 0
	// SNErrorCodeOperationFailed - An error that occurs when the framework fails to analyze audio.
	//
	// [Full Topic]: https://developer.apple.com/documentation/SoundAnalysis/SNError/Code/operationFailed
	SNErrorCodeOperationFailed SNErrorCode = 0
)


