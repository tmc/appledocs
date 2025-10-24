// Code generated from Apple documentation for ImageIO. DO NOT EDIT.

package imageio

/* debug [enums.gen.go]: Generating 6 enums for ImageIO */
// Enum types and constants
/* debug [enums.gen.go]: Processing enum CGImageAnimationStatus (5 cases) */
// CGImageAnimationStatus - Constants that indicate the result of animating an image sequence.
//
// [Full Topic]: https://developer.apple.com/documentation/ImageIO/CGImageAnimationStatus
type CGImageAnimationStatus uint

const (
	//
	// [Full Topic]: https://developer.apple.com/documentation/ImageIO/CGImageAnimationStatus/allocationFailure
	kCGImageAnimationStatus_AllocationFailure CGImageAnimationStatus = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/ImageIO/CGImageAnimationStatus/corruptInputImage
	kCGImageAnimationStatus_CorruptInputImage CGImageAnimationStatus = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/ImageIO/CGImageAnimationStatus/incompleteInputImage
	kCGImageAnimationStatus_IncompleteInputImage CGImageAnimationStatus = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/ImageIO/CGImageAnimationStatus/parameterError
	kCGImageAnimationStatus_ParameterError CGImageAnimationStatus = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/ImageIO/CGImageAnimationStatus/unsupportedFormat
	kCGImageAnimationStatus_UnsupportedFormat CGImageAnimationStatus = 0
)

/* debug [enums.gen.go]: Processing enum CGImageMetadataErrors (5 cases) */
// CGImageMetadataErrors - Constants for errors that occur when getting or setting metadata information.
//
// [Full Topic]: https://developer.apple.com/documentation/ImageIO/CGImageMetadataErrors
type CGImageMetadataErrors uint

const (
	// kCGImageMetadataErrorBadArgument - An error that indicates a parameter was malformed or contained invalid data.
	//
	// [Full Topic]: https://developer.apple.com/documentation/ImageIO/CGImageMetadataErrors/badArgument
	kCGImageMetadataErrorBadArgument CGImageMetadataErrors = 0
	// kCGImageMetadataErrorConflictingArguments - An error that indicates an attempt to save conflicting metadata values.
	//
	// [Full Topic]: https://developer.apple.com/documentation/ImageIO/CGImageMetadataErrors/conflictingArguments
	kCGImageMetadataErrorConflictingArguments CGImageMetadataErrors = 0
	// kCGImageMetadataErrorPrefixConflict - An error that indicates an attempt to register a namespace with a prefix that is different than the namespace’s existing prefix.
	//
	// [Full Topic]: https://developer.apple.com/documentation/ImageIO/CGImageMetadataErrors/prefixConflict
	kCGImageMetadataErrorPrefixConflict CGImageMetadataErrors = 0
	// kCGImageMetadataErrorUnknown - An error that indicates an unknown condition occurred.
	//
	// [Full Topic]: https://developer.apple.com/documentation/ImageIO/CGImageMetadataErrors/unknown
	kCGImageMetadataErrorUnknown CGImageMetadataErrors = 0
	// kCGImageMetadataErrorUnsupportedFormat - An error that indicates the metadata was in an unsupported format.
	//
	// [Full Topic]: https://developer.apple.com/documentation/ImageIO/CGImageMetadataErrors/unsupportedFormat
	kCGImageMetadataErrorUnsupportedFormat CGImageMetadataErrors = 0
)

/* debug [enums.gen.go]: Processing enum CGImageMetadataType (8 cases) */
// CGImageMetadataType - Constants that indicate the XMP type for a metadata tag.
//
// [Full Topic]: https://developer.apple.com/documentation/ImageIO/CGImageMetadataType
type CGImageMetadataType uint

const (
	// kCGImageMetadataTypeAlternateArray - An ordered array, in which all elements are alternates for the same value.
	//
	// [Full Topic]: https://developer.apple.com/documentation/ImageIO/CGImageMetadataType/alternateArray
	kCGImageMetadataTypeAlternateArray CGImageMetadataType = 0
	// kCGImageMetadataTypeAlternateText - An alternate array, in which all elements are localized strings for the same value.
	//
	// [Full Topic]: https://developer.apple.com/documentation/ImageIO/CGImageMetadataType/alternateText
	kCGImageMetadataTypeAlternateText CGImageMetadataType = 0
	// kCGImageMetadataTypeArrayOrdered - An array that preserves the order of items.
	//
	// [Full Topic]: https://developer.apple.com/documentation/ImageIO/CGImageMetadataType/arrayOrdered
	kCGImageMetadataTypeArrayOrdered CGImageMetadataType = 0
	// kCGImageMetadataTypeArrayUnordered - An array that doesn’t preserve the order of items.
	//
	// [Full Topic]: https://developer.apple.com/documentation/ImageIO/CGImageMetadataType/arrayUnordered
	kCGImageMetadataTypeArrayUnordered CGImageMetadataType = 0
	// kCGImageMetadataTypeDefault - The default type for new tags.
	//
	// [Full Topic]: https://developer.apple.com/documentation/ImageIO/CGImageMetadataType/default
	kCGImageMetadataTypeDefault CGImageMetadataType = 0
	// kCGImageMetadataTypeInvalid - An invalid metadata type.
	//
	// [Full Topic]: https://developer.apple.com/documentation/ImageIO/CGImageMetadataType/invalid
	kCGImageMetadataTypeInvalid CGImageMetadataType = 0
	// kCGImageMetadataTypeString - A string value.
	//
	// [Full Topic]: https://developer.apple.com/documentation/ImageIO/CGImageMetadataType/string
	kCGImageMetadataTypeString CGImageMetadataType = 0
	// kCGImageMetadataTypeStructure - A collection of keys and values.
	//
	// [Full Topic]: https://developer.apple.com/documentation/ImageIO/CGImageMetadataType/structure
	kCGImageMetadataTypeStructure CGImageMetadataType = 0
)

/* debug [enums.gen.go]: Processing enum CGImagePropertyOrientation (8 cases) */
// CGImagePropertyOrientation - A value describing the intended display orientation for an image.
//
// [Full Topic]: https://developer.apple.com/documentation/ImageIO/CGImagePropertyOrientation
type CGImagePropertyOrientation uint

const (
	// kCGImagePropertyOrientationDown - The encoded image data is rotated 180° from the image’s intended display orientation.
	//
	// [Full Topic]: https://developer.apple.com/documentation/ImageIO/CGImagePropertyOrientation/down
	kCGImagePropertyOrientationDown CGImagePropertyOrientation = 0
	// kCGImagePropertyOrientationDownMirrored - The encoded image data is vertically flipped from the image’s intended display orientation.
	//
	// [Full Topic]: https://developer.apple.com/documentation/ImageIO/CGImagePropertyOrientation/downMirrored
	kCGImagePropertyOrientationDownMirrored CGImagePropertyOrientation = 0
	// kCGImagePropertyOrientationLeft - The encoded image data is rotated 90° clockwise from the image’s intended display orientation.
	//
	// [Full Topic]: https://developer.apple.com/documentation/ImageIO/CGImagePropertyOrientation/left
	kCGImagePropertyOrientationLeft CGImagePropertyOrientation = 0
	// kCGImagePropertyOrientationLeftMirrored - The encoded image data is horizontally flipped and rotated 90° counter-clockwise from the image’s intended display orientation.
	//
	// [Full Topic]: https://developer.apple.com/documentation/ImageIO/CGImagePropertyOrientation/leftMirrored
	kCGImagePropertyOrientationLeftMirrored CGImagePropertyOrientation = 0
	// kCGImagePropertyOrientationRight - The encoded image data is rotated 90° clockwise from the image’s intended display orientation.
	//
	// [Full Topic]: https://developer.apple.com/documentation/ImageIO/CGImagePropertyOrientation/right
	kCGImagePropertyOrientationRight CGImagePropertyOrientation = 0
	// kCGImagePropertyOrientationRightMirrored - The encoded image data is horizontally flipped and rotated 90° clockwise from the image’s intended display orientation.
	//
	// [Full Topic]: https://developer.apple.com/documentation/ImageIO/CGImagePropertyOrientation/rightMirrored
	kCGImagePropertyOrientationRightMirrored CGImagePropertyOrientation = 0
	// kCGImagePropertyOrientationUp - The encoded image data matches the image’s intended display orientation.
	//
	// [Full Topic]: https://developer.apple.com/documentation/ImageIO/CGImagePropertyOrientation/up
	kCGImagePropertyOrientationUp CGImagePropertyOrientation = 0
	// kCGImagePropertyOrientationUpMirrored - The encoded image data is horizontally flipped from the image’s intended display orientation.
	//
	// [Full Topic]: https://developer.apple.com/documentation/ImageIO/CGImagePropertyOrientation/upMirrored
	kCGImagePropertyOrientationUpMirrored CGImagePropertyOrientation = 0
)

/* debug [enums.gen.go]: Processing enum CGImagePropertyTGACompression (2 cases) */
// CGImagePropertyTGACompression enum type
//
// [Full Topic]: https://developer.apple.com/documentation/ImageIO/CGImagePropertyTGACompression
type CGImagePropertyTGACompression uint

const (
	//
	// [Full Topic]: https://developer.apple.com/documentation/ImageIO/CGImagePropertyTGACompression/tgaCompressionNone
	kCGImageTGACompressionNone CGImagePropertyTGACompression = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/ImageIO/CGImagePropertyTGACompression/tgaCompressionRLE
	kCGImageTGACompressionRLE CGImagePropertyTGACompression = 0
)

/* debug [enums.gen.go]: Processing enum CGImageSourceStatus (6 cases) */
// CGImageSourceStatus - The set of status values for images and image sources.
//
// [Full Topic]: https://developer.apple.com/documentation/ImageIO/CGImageSourceStatus
type CGImageSourceStatus uint

const (
	// kCGImageStatusComplete - The operation is complete.
	//
	// [Full Topic]: https://developer.apple.com/documentation/ImageIO/CGImageSourceStatus/statusComplete
	kCGImageStatusComplete CGImageSourceStatus = 0
	// kCGImageStatusIncomplete - The operation is not complete
	//
	// [Full Topic]: https://developer.apple.com/documentation/ImageIO/CGImageSourceStatus/statusIncomplete
	kCGImageStatusIncomplete CGImageSourceStatus = 0
	// kCGImageStatusInvalidData - The data is not valid.
	//
	// [Full Topic]: https://developer.apple.com/documentation/ImageIO/CGImageSourceStatus/statusInvalidData
	kCGImageStatusInvalidData CGImageSourceStatus = 0
	// kCGImageStatusReadingHeader - The image source is reading the header.
	//
	// [Full Topic]: https://developer.apple.com/documentation/ImageIO/CGImageSourceStatus/statusReadingHeader
	kCGImageStatusReadingHeader CGImageSourceStatus = 0
	// kCGImageStatusUnexpectedEOF - The end of the file occurred unexpectedly.
	//
	// [Full Topic]: https://developer.apple.com/documentation/ImageIO/CGImageSourceStatus/statusUnexpectedEOF
	kCGImageStatusUnexpectedEOF CGImageSourceStatus = 0
	// kCGImageStatusUnknownType - The image is an unknown type.
	//
	// [Full Topic]: https://developer.apple.com/documentation/ImageIO/CGImageSourceStatus/statusUnknownType
	kCGImageStatusUnknownType CGImageSourceStatus = 0
)


