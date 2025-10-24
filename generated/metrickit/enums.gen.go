// Code generated from Apple documentation for MetricKit. DO NOT EDIT.

package metrickit

/* debug [enums.gen.go]: Generating 1 enums for MetricKit */
// Enum types and constants
/* debug [enums.gen.go]: Processing enum MXErrorCode (6 cases) */
// MXErrorCode - Error codes for error values from app metrics.
//
// [Full Topic]: https://developer.apple.com/documentation/MetricKit/MXError/Code
type MXErrorCode uint

const (
	// MXErrorLaunchTaskDuplicated - A task with the same ID has already been started.
	//
	// [Full Topic]: https://developer.apple.com/documentation/MetricKit/MXError/Code/launchTaskDuplicated
	MXErrorLaunchTaskDuplicated MXErrorCode = 0
	// MXErrorLaunchTaskInternalFailure - Internal failures happened inside the framework.
	//
	// [Full Topic]: https://developer.apple.com/documentation/MetricKit/MXError/Code/launchTaskInternalFailure
	MXErrorLaunchTaskInternalFailure MXErrorCode = 0
	// MXErrorLaunchTaskInvalidID - The task ID is a   value or exceeds the maximum 128 character length.
	//
	// [Full Topic]: https://developer.apple.com/documentation/MetricKit/MXError/Code/launchTaskInvalidID
	MXErrorLaunchTaskInvalidID MXErrorCode = 0
	// MXErrorLaunchTaskMaxCount - Exceeded the maximum number of tasks.
	//
	// [Full Topic]: https://developer.apple.com/documentation/MetricKit/MXError/Code/launchTaskMaxCount
	MXErrorLaunchTaskMaxCount MXErrorCode = 0
	// MXErrorLaunchTaskPastDeadline - The start call was made too late.
	//
	// [Full Topic]: https://developer.apple.com/documentation/MetricKit/MXError/Code/launchTaskPastDeadline
	MXErrorLaunchTaskPastDeadline MXErrorCode = 0
	// MXErrorLaunchTaskUnknown - The task hasn’t been started or has already been finished.
	//
	// [Full Topic]: https://developer.apple.com/documentation/MetricKit/MXError/Code/launchTaskUnknown
	MXErrorLaunchTaskUnknown MXErrorCode = 0
)


