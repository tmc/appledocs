// Code generated from Apple documentation for MediaExtension. DO NOT EDIT.

package mediaextension

/* debug [enums.gen.go]: Generating 4 enums for MediaExtension */
// Enum types and constants
/* debug [enums.gen.go]: Processing enum MEError (12 cases) */
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
	// MEErrorInvalidParameter - An error code that indicates the extension received an invalid parameter.
	//
	// [Full Topic]: https://developer.apple.com/documentation/MediaExtension/MEError-swift.struct/Code/invalidParameter
	MEErrorInvalidParameter MEError = 0
	// MEErrorLocationNotAvailable - An error code that indicates specific sample isn’t contiguous, spans more than one file, or is for some other reason unsuitable for reading directly from a file.
	//
	// [Full Topic]: https://developer.apple.com/documentation/MediaExtension/MEError-swift.struct/Code/locationNotAvailable
	MEErrorLocationNotAvailable MEError = 0
	// MEErrorNoSamples - An error code that indicates there are no samples in the track or a request to load a sample buffer fails.
	//
	// [Full Topic]: https://developer.apple.com/documentation/MediaExtension/MEError-swift.struct/Code/noSamples
	MEErrorNoSamples MEError = 0
	// MEErrorNoSuchEdit - An error code that indicates the plug-in track reader received a request to return an edit that’s out of range.
	//
	// [Full Topic]: https://developer.apple.com/documentation/MediaExtension/MEError-swift.struct/Code/noSuchEdit
	MEErrorNoSuchEdit MEError = 0
	// MEErrorParsingFailure - An error code that indicates the extension encountered an error while parsing the media.
	//
	// [Full Topic]: https://developer.apple.com/documentation/MediaExtension/MEError-swift.struct/Code/parsingFailure
	MEErrorParsingFailure MEError = 0
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

/* debug [enums.gen.go]: Processing enum MEFileInfoFragmentsStatus (3 cases) */
// MEFileInfoFragmentsStatus - An enumeration that describes if a media asset contains or supports fragments.
//
// [Full Topic]: https://developer.apple.com/documentation/MediaExtension/MEFileInfo/FragmentsStatus-swift.enum
type MEFileInfoFragmentsStatus uint

const (
	// MEFileInfoContainsFragments - The file is extendable by fragments and contains at least one fragment.
	//
	// [Full Topic]: https://developer.apple.com/documentation/MediaExtension/MEFileInfo/FragmentsStatus-swift.enum/containsFragments
	MEFileInfoContainsFragments MEFileInfoFragmentsStatus = 0
	// MEFileInfoCouldContainButDoesNotContainFragments - The file is extendable by fragments, but doesn’t contain any fragments.
	//
	// [Full Topic]: https://developer.apple.com/documentation/MediaExtension/MEFileInfo/FragmentsStatus-swift.enum/couldContainButDoesNotContainFragments
	MEFileInfoCouldContainButDoesNotContainFragments MEFileInfoFragmentsStatus = 0
	// MEFileInfoCouldNotContainFragments - The file isn’t extendable by fragments.
	//
	// [Full Topic]: https://developer.apple.com/documentation/MediaExtension/MEFileInfo/FragmentsStatus-swift.enum/couldNotContainFragments
	MEFileInfoCouldNotContainFragments MEFileInfoFragmentsStatus = 0
)

/* debug [enums.gen.go]: Processing enum MEDecodeFrameStatus (2 cases) */
// MEDecodeFrameStatus - A type that represents a non-error status related to a frame decode operation.
//
// [Full Topic]: https://developer.apple.com/documentation/MediaExtension/MEDecodeFrameStatus
type MEDecodeFrameStatus uint

const (
	// MEDecodeFrameFrameDropped - A frame decode operation status that indicates the system dropped the output of the frame for a reason other than an error.
	//
	// [Full Topic]: https://developer.apple.com/documentation/MediaExtension/MEDecodeFrameStatus/frameDropped
	MEDecodeFrameFrameDropped MEDecodeFrameStatus = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/MediaExtension/MEDecodeFrameStatus/MEDecodeFrameNoStatus
	MEDecodeFrameNoStatus MEDecodeFrameStatus = 0
)

/* debug [enums.gen.go]: Processing enum MEFormatReaderParseAdditionalFragmentsStatus (3 cases) */
// MEFormatReaderParseAdditionalFragmentsStatus - Informational status flags that the format reader returns after parsing additional fragments.
//
// [Full Topic]: https://developer.apple.com/documentation/MediaExtension/MEFormatReaderParseAdditionalFragmentsStatus
type MEFormatReaderParseAdditionalFragmentsStatus uint

const (
	// MEFormatReaderParseAdditionalFragmentsStatusFragmentAdded - Indicates that the format reader received one or more fragments.
	//
	// [Full Topic]: https://developer.apple.com/documentation/MediaExtension/MEFormatReaderParseAdditionalFragmentsStatus/fragmentAdded
	MEFormatReaderParseAdditionalFragmentsStatusFragmentAdded MEFormatReaderParseAdditionalFragmentsStatus = 0
	// MEFormatReaderParseAdditionalFragmentsStatusFragmentsComplete - Indicates that the format reader can’t receive any more fragments.
	//
	// [Full Topic]: https://developer.apple.com/documentation/MediaExtension/MEFormatReaderParseAdditionalFragmentsStatus/fragmentsComplete
	MEFormatReaderParseAdditionalFragmentsStatusFragmentsComplete MEFormatReaderParseAdditionalFragmentsStatus = 0
	// MEFormatReaderParseAdditionalFragmentsStatusSizeIncreased - Indicates that the format reader file size increased.
	//
	// [Full Topic]: https://developer.apple.com/documentation/MediaExtension/MEFormatReaderParseAdditionalFragmentsStatus/sizeIncreased
	MEFormatReaderParseAdditionalFragmentsStatusSizeIncreased MEFormatReaderParseAdditionalFragmentsStatus = 0
)


