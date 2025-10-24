// Code generated from Apple documentation for ParavirtualizedGraphics. DO NOT EDIT.

package paravirtualizedgraphics

/* debug [enums.gen.go]: Generating 1 enums for ParavirtualizedGraphics */
// Enum types and constants
/* debug [enums.gen.go]: Processing enum PGResumeErrorCode (6 cases) */
// PGResumeErrorCode - Error codes for suspend-resume actions.
//
// [Full Topic]: https://developer.apple.com/documentation/ParavirtualizedGraphics/PGResumeErrorCode
type PGResumeErrorCode uint

const (
	// PGResumeErrorCodeIncompatibleDevice - The resume device is missing capabilities that the suspended device provided.
	//
	// [Full Topic]: https://developer.apple.com/documentation/ParavirtualizedGraphics/PGResumeErrorCode/incompatibleDevice
	PGResumeErrorCodeIncompatibleDevice PGResumeErrorCode = 0
	// PGResumeErrorCodeInternalFault - An internal error occurred.
	//
	// [Full Topic]: https://developer.apple.com/documentation/ParavirtualizedGraphics/PGResumeErrorCode/internalFault
	PGResumeErrorCodeInternalFault PGResumeErrorCode = 0
	// PGResumeErrorCodeInvalidContent - The content of the suspend state or the guest memory isn’t valid.
	//
	// [Full Topic]: https://developer.apple.com/documentation/ParavirtualizedGraphics/PGResumeErrorCode/invalidContent
	PGResumeErrorCodeInvalidContent PGResumeErrorCode = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/ParavirtualizedGraphics/PGResumeErrorCode/invalidDisplayPortCount
	PGResumeErrorCodeInvalidDisplayPortCount PGResumeErrorCode = 0
	// PGResumeErrorCodeInvalidGuestVersion - The guest version is incompatible with this framework version.
	//
	// [Full Topic]: https://developer.apple.com/documentation/ParavirtualizedGraphics/PGResumeErrorCode/invalidGuestVersion
	PGResumeErrorCodeInvalidGuestVersion PGResumeErrorCode = 0
	// PGResumeErrorCodeInvalidSuspendStateVersion - The suspend state version is incompatible with this framework version.
	//
	// [Full Topic]: https://developer.apple.com/documentation/ParavirtualizedGraphics/PGResumeErrorCode/invalidSuspendStateVersion
	PGResumeErrorCodeInvalidSuspendStateVersion PGResumeErrorCode = 0
)


