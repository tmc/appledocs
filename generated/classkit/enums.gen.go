// Code generated from Apple documentation for ClassKit. DO NOT EDIT.

package classkit

// Enum types and constants
// CLSBinaryValueType - The kinds of outcomes that a binary activity item can represent.
//
// [Full Topic]: https://developer.apple.com/documentation/ClassKit/CLSBinaryValueType
type CLSBinaryValueType uint

// CLSContextType - The kinds of assignable content a context can contain.
//
// [Full Topic]: https://developer.apple.com/documentation/ClassKit/CLSContextType
type CLSContextType uint

const (
	// CLSContextTypeChapter - A chapter context.
	//
	// [Full Topic]: https://developer.apple.com/documentation/ClassKit/CLSContextType/chapter
	CLSContextTypeChapter CLSContextType = 0
	// CLSContextTypeQuiz - A quiz context.
	//
	// [Full Topic]: https://developer.apple.com/documentation/ClassKit/CLSContextType/quiz
	CLSContextTypeQuiz CLSContextType = 0
)

// CLSErrorCode - Error codes that ClassKit issues.
//
// [Full Topic]: https://developer.apple.com/documentation/ClassKit/CLSError/Code
type CLSErrorCode uint

const (
	// CLSErrorCodeAuthorizationDenied - The app isn’t authorized to perform the requested operation.
	//
	// [Full Topic]: https://developer.apple.com/documentation/ClassKit/CLSError/Code/authorizationDenied
	CLSErrorCodeAuthorizationDenied CLSErrorCode = 0
	// CLSErrorCodeClassKitUnavailable - ClassKit isn’t available on this device.
	//
	// [Full Topic]: https://developer.apple.com/documentation/ClassKit/CLSError/Code/classKitUnavailable
	CLSErrorCodeClassKitUnavailable CLSErrorCode = 0
	// CLSErrorCodeDatabaseInaccessible - ClassKit isn’t accessible because the device is locked.
	//
	// [Full Topic]: https://developer.apple.com/documentation/ClassKit/CLSError/Code/databaseInaccessible
	CLSErrorCodeDatabaseInaccessible CLSErrorCode = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/ClassKit/CLSError/Code/invalidAccountCredentials
	CLSErrorCodeInvalidAccountCredentials CLSErrorCode = 0
	// CLSErrorCodeInvalidArgument - An invalid argument was provided to the API.
	//
	// [Full Topic]: https://developer.apple.com/documentation/ClassKit/CLSError/Code/invalidArgument
	CLSErrorCodeInvalidArgument CLSErrorCode = 0
	// CLSErrorCodeInvalidCreate - An attempt to save a new object that already exists in the data store failed.
	//
	// [Full Topic]: https://developer.apple.com/documentation/ClassKit/CLSError/Code/invalidCreate
	CLSErrorCodeInvalidCreate CLSErrorCode = 0
	// CLSErrorCodeInvalidModification - An attempt to modify a read-only object failed.
	//
	// [Full Topic]: https://developer.apple.com/documentation/ClassKit/CLSError/Code/invalidModification
	CLSErrorCodeInvalidModification CLSErrorCode = 0
	// CLSErrorCodeInvalidUpdate - ClassKit failed to save an updated object in the data store.
	//
	// [Full Topic]: https://developer.apple.com/documentation/ClassKit/CLSError/Code/invalidUpdate
	CLSErrorCodeInvalidUpdate CLSErrorCode = 0
	// CLSErrorCodeLimits - A limit has been exceeded.
	//
	// [Full Topic]: https://developer.apple.com/documentation/ClassKit/CLSError/Code/limits
	CLSErrorCodeLimits CLSErrorCode = 0
	// CLSErrorCodeNone - No error.
	//
	// [Full Topic]: https://developer.apple.com/documentation/ClassKit/CLSError/Code/none
	CLSErrorCodeNone CLSErrorCode = 0
	// CLSErrorCodePartialFailure - ClassKit encountered more than one error.
	//
	// [Full Topic]: https://developer.apple.com/documentation/ClassKit/CLSError/Code/partialFailure
	CLSErrorCodePartialFailure CLSErrorCode = 0
)


