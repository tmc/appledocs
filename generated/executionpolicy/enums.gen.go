// Code generated from Apple documentation for ExecutionPolicy. DO NOT EDIT.

package executionpolicy

/* debug [enums.gen.go]: Generating 2 enums for ExecutionPolicy */
// Enum types and constants
/* debug [enums.gen.go]: Processing enum EPError (2 cases) */
// EPError enum type
//
// [Full Topic]: https://developer.apple.com/documentation/ExecutionPolicy/EPError-swift.struct/Code
type EPError uint

const (
	//
	// [Full Topic]: https://developer.apple.com/documentation/ExecutionPolicy/EPError-swift.struct/Code/generic
	EPErrorGeneric EPError = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/ExecutionPolicy/EPError-swift.struct/Code/notADeveloperTool
	EPErrorNotADeveloperTool EPError = 0
)

/* debug [enums.gen.go]: Processing enum EPDeveloperToolStatus (4 cases) */
// EPDeveloperToolStatus enum type
//
// [Full Topic]: https://developer.apple.com/documentation/ExecutionPolicy/EPDeveloperToolStatus
type EPDeveloperToolStatus uint

const (
	//
	// [Full Topic]: https://developer.apple.com/documentation/ExecutionPolicy/EPDeveloperToolStatus/authorized
	EPDeveloperToolStatusAuthorized EPDeveloperToolStatus = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/ExecutionPolicy/EPDeveloperToolStatus/denied
	EPDeveloperToolStatusDenied EPDeveloperToolStatus = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/ExecutionPolicy/EPDeveloperToolStatus/notDetermined
	EPDeveloperToolStatusNotDetermined EPDeveloperToolStatus = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/ExecutionPolicy/EPDeveloperToolStatus/restricted
	EPDeveloperToolStatusRestricted EPDeveloperToolStatus = 0
)


