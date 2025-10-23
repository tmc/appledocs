// Code generated from Apple documentation for FSKit. DO NOT EDIT.

package fskit

// Enum types and constants
// FSBlockmapFlags - Flags that describe the behavior of a blockmap operation.
//
// [Full Topic]: https://developer.apple.com/documentation/FSKit/FSBlockmapFlags
type FSBlockmapFlags uint

const (
	// FSBlockmapFlagsWrite - A flag that describes a write operation.
	//
	// [Full Topic]: https://developer.apple.com/documentation/FSKit/FSBlockmapFlags/write
	FSBlockmapFlagsWrite FSBlockmapFlags = 0
)

// FSCompleteIOFlags - Flags that describe the behavior of an I/O completion operation.
//
// [Full Topic]: https://developer.apple.com/documentation/FSKit/FSCompleteIOFlags
type FSCompleteIOFlags uint

const (
	// FSCompleteIOFlagsWrite - A flag that describes a write operation.
	//
	// [Full Topic]: https://developer.apple.com/documentation/FSKit/FSCompleteIOFlags/write
	FSCompleteIOFlagsWrite FSCompleteIOFlags = 0
)

// FSErrorCode - A code that indicates a specific FSKit error.
//
// [Full Topic]: https://developer.apple.com/documentation/FSKit/FSError/Code
type FSErrorCode uint

const (
	// FSErrorStatusOperationInProgress - An operation is in progress.
	//
	// [Full Topic]: https://developer.apple.com/documentation/FSKit/FSError/Code/statusOperationInProgress
	FSErrorStatusOperationInProgress FSErrorCode = 0
)

// FSExtentType - An enumeration of types of extents.
//
// [Full Topic]: https://developer.apple.com/documentation/FSKit/FSExtentType
type FSExtentType uint

const (
	// FSExtentTypeZeroFill - An extent type to indicate uninitialized data.
	//
	// [Full Topic]: https://developer.apple.com/documentation/FSKit/FSExtentType/zeroFill
	FSExtentTypeZeroFill FSExtentType = 0
)

// FSMatchResult - A type that represents the recognition and usability of a probed resource.
//
// [Full Topic]: https://developer.apple.com/documentation/FSKit/FSMatchResult
type FSMatchResult uint

const (
	// FSMatchResultNotRecognized - The probe doesn’t recognize the resource.
	//
	// [Full Topic]: https://developer.apple.com/documentation/FSKit/FSMatchResult/notRecognized
	FSMatchResultNotRecognized FSMatchResult = 0
	// FSMatchResultRecognized - The probe recognizes the resource but can’t use it.
	//
	// [Full Topic]: https://developer.apple.com/documentation/FSKit/FSMatchResult/recognized
	FSMatchResultRecognized FSMatchResult = 0
)


