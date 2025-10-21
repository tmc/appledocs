// Code generated from Apple documentation for MediaExtension. DO NOT EDIT.

package mediaextension

// Enum types and constants
// MEError - An enumeration that models media extension error codes.
//
// [Full Topic]: https://developer.apple.com/documentation/MediaExtension/MEError-swift.struct/Code
type MEError uint

const (
// MEErrorAllocationFailure - An error code that indicates the extension can’t allocate memory.
//
	// [Full Topic]: https://developer.apple.com/documentation/MediaExtension/MEError-swift.struct/Code/allocationFailure
MEErrorAllocationFailure MEError = 0
// MEErrorEndOfStream - An error code that indicates the extension reached the end of the source file.
//
	// [Full Topic]: https://developer.apple.com/documentation/MediaExtension/MEError-swift.struct/Code/endOfStream
MEErrorEndOfStream MEError = 0
// MEErrorInternalFailure - An error code that indicates the extension encountered an internal operation failure, such as code loading.
//
	// [Full Topic]: https://developer.apple.com/documentation/MediaExtension/MEError-swift.struct/Code/internalFailure
MEErrorInternalFailure MEError = 0
// MEErrorLocationNotAvailable - An error code that indicates specific sample isn’t contiguous, spans more than one file, or is for some other reason unsuitable for reading directly from a file.
//
	// [Full Topic]: https://developer.apple.com/documentation/MediaExtension/MEError-swift.struct/Code/locationNotAvailable
MEErrorLocationNotAvailable MEError = 0
// MEErrorNoSamples - An error code that indicates there are no samples in the track or a request to load a sample buffer fails.
//
	// [Full Topic]: https://developer.apple.com/documentation/MediaExtension/MEError-swift.struct/Code/noSamples
MEErrorNoSamples MEError = 0
// MEErrorPermissionDenied - An error code that indicates the extension received a request to perform an invalid operation on a byte source.
//
	// [Full Topic]: https://developer.apple.com/documentation/MediaExtension/MEError-swift.struct/Code/permissionDenied
MEErrorPermissionDenied MEError = 0
// MEErrorPropertyNotSupported - An error code that indicates the extension encountered a property it doesn’t support reading and writing to.
//
	// [Full Topic]: https://developer.apple.com/documentation/MediaExtension/MEError-swift.struct/Code/propertyNotSupported
MEErrorPropertyNotSupported MEError = 0
// MEErrorReferenceMissing - An error code that indicates the decoder received a request to decode a sample without decoding the required reference frame dependencies first.
//
	// [Full Topic]: https://developer.apple.com/documentation/MediaExtension/MEError-swift.struct/Code/referenceMissing
MEErrorReferenceMissing MEError = 0
// MEErrorUnsupportedFeature - An error code that indicates the extension doesn’t support an aspect of the media.
//
	// [Full Topic]: https://developer.apple.com/documentation/MediaExtension/MEError-swift.struct/Code/unsupportedFeature
MEErrorUnsupportedFeature MEError = 0
)


