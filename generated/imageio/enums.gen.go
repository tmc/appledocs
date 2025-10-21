// Code generated from Apple documentation for ImageIO. DO NOT EDIT.

package imageio

// Enum types and constants
// CGImageAnimationStatus - Constants that indicate the result of animating an image sequence.
//
// [Full Topic]: https://developer.apple.com/documentation/ImageIO/CGImageAnimationStatus
type ImageAnimationStatus uint

const (
	//
	// [Full Topic]: https://developer.apple.com/documentation/ImageIO/CGImageAnimationStatus/allocationFailure
	kCGImageAnimationStatus_AllocationFailure ImageAnimationStatus = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/ImageIO/CGImageAnimationStatus/corruptInputImage
	kCGImageAnimationStatus_CorruptInputImage ImageAnimationStatus = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/ImageIO/CGImageAnimationStatus/incompleteInputImage
	kCGImageAnimationStatus_IncompleteInputImage ImageAnimationStatus = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/ImageIO/CGImageAnimationStatus/parameterError
	kCGImageAnimationStatus_ParameterError ImageAnimationStatus = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/ImageIO/CGImageAnimationStatus/unsupportedFormat
	kCGImageAnimationStatus_UnsupportedFormat ImageAnimationStatus = 0
)

// CGImageMetadataErrors - Constants for errors that occur when getting or setting metadata information.
//
// [Full Topic]: https://developer.apple.com/documentation/ImageIO/CGImageMetadataErrors
type ImageMetadataErrors uint

const (
	// kCGImageMetadataErrorBadArgument - An error that indicates a parameter was malformed or contained invalid data.
	//
	// [Full Topic]: https://developer.apple.com/documentation/ImageIO/CGImageMetadataErrors/badArgument
	kCGImageMetadataErrorBadArgument ImageMetadataErrors = 0
	// kCGImageMetadataErrorPrefixConflict - An error that indicates an attempt to register a namespace with a prefix that is different than the namespace’s existing prefix.
	//
	// [Full Topic]: https://developer.apple.com/documentation/ImageIO/CGImageMetadataErrors/prefixConflict
	kCGImageMetadataErrorPrefixConflict ImageMetadataErrors = 0
	// kCGImageMetadataErrorUnknown - An error that indicates an unknown condition occurred.
	//
	// [Full Topic]: https://developer.apple.com/documentation/ImageIO/CGImageMetadataErrors/unknown
	kCGImageMetadataErrorUnknown ImageMetadataErrors = 0
)

// CGImageMetadataType - Constants that indicate the XMP type for a metadata tag.
//
// [Full Topic]: https://developer.apple.com/documentation/ImageIO/CGImageMetadataType
type ImageMetadataType uint

const (
	// kCGImageMetadataTypeAlternateArray - An ordered array, in which all elements are alternates for the same value.
	//
	// [Full Topic]: https://developer.apple.com/documentation/ImageIO/CGImageMetadataType/alternateArray
	kCGImageMetadataTypeAlternateArray ImageMetadataType = 0
	// kCGImageMetadataTypeAlternateText - An alternate array, in which all elements are localized strings for the same value.
	//
	// [Full Topic]: https://developer.apple.com/documentation/ImageIO/CGImageMetadataType/alternateText
	kCGImageMetadataTypeAlternateText ImageMetadataType = 0
	// kCGImageMetadataTypeArrayOrdered - An array that preserves the order of items.
	//
	// [Full Topic]: https://developer.apple.com/documentation/ImageIO/CGImageMetadataType/arrayOrdered
	kCGImageMetadataTypeArrayOrdered ImageMetadataType = 0
	// kCGImageMetadataTypeArrayUnordered - An array that doesn’t preserve the order of items.
	//
	// [Full Topic]: https://developer.apple.com/documentation/ImageIO/CGImageMetadataType/arrayUnordered
	kCGImageMetadataTypeArrayUnordered ImageMetadataType = 0
	// kCGImageMetadataTypeDefault - The default type for new tags.
	//
	// [Full Topic]: https://developer.apple.com/documentation/ImageIO/CGImageMetadataType/default
	kCGImageMetadataTypeDefault ImageMetadataType = 0
	// kCGImageMetadataTypeInvalid - An invalid metadata type.
	//
	// [Full Topic]: https://developer.apple.com/documentation/ImageIO/CGImageMetadataType/invalid
	kCGImageMetadataTypeInvalid ImageMetadataType = 0
	// kCGImageMetadataTypeString - A string value.
	//
	// [Full Topic]: https://developer.apple.com/documentation/ImageIO/CGImageMetadataType/string
	kCGImageMetadataTypeString ImageMetadataType = 0
	// kCGImageMetadataTypeStructure - A collection of keys and values.
	//
	// [Full Topic]: https://developer.apple.com/documentation/ImageIO/CGImageMetadataType/structure
	kCGImageMetadataTypeStructure ImageMetadataType = 0
)

// CGImagePropertyOrientation - A value describing the intended display orientation for an image.
//
// [Full Topic]: https://developer.apple.com/documentation/ImageIO/CGImagePropertyOrientation
type ImagePropertyOrientation uint

const (
	// kCGImagePropertyOrientationDown - The encoded image data is rotated 180° from the image’s intended display orientation.
	//
	// [Full Topic]: https://developer.apple.com/documentation/ImageIO/CGImagePropertyOrientation/down
	kCGImagePropertyOrientationDown ImagePropertyOrientation = 0
	// kCGImagePropertyOrientationDownMirrored - The encoded image data is vertically flipped from the image’s intended display orientation.
	//
	// [Full Topic]: https://developer.apple.com/documentation/ImageIO/CGImagePropertyOrientation/downMirrored
	kCGImagePropertyOrientationDownMirrored ImagePropertyOrientation = 0
	// kCGImagePropertyOrientationLeft - The encoded image data is rotated 90° clockwise from the image’s intended display orientation.
	//
	// [Full Topic]: https://developer.apple.com/documentation/ImageIO/CGImagePropertyOrientation/left
	kCGImagePropertyOrientationLeft ImagePropertyOrientation = 0
	// kCGImagePropertyOrientationLeftMirrored - The encoded image data is horizontally flipped and rotated 90° counter-clockwise from the image’s intended display orientation.
	//
	// [Full Topic]: https://developer.apple.com/documentation/ImageIO/CGImagePropertyOrientation/leftMirrored
	kCGImagePropertyOrientationLeftMirrored ImagePropertyOrientation = 0
	// kCGImagePropertyOrientationRight - The encoded image data is rotated 90° clockwise from the image’s intended display orientation.
	//
	// [Full Topic]: https://developer.apple.com/documentation/ImageIO/CGImagePropertyOrientation/right
	kCGImagePropertyOrientationRight ImagePropertyOrientation = 0
	// kCGImagePropertyOrientationRightMirrored - The encoded image data is horizontally flipped and rotated 90° clockwise from the image’s intended display orientation.
	//
	// [Full Topic]: https://developer.apple.com/documentation/ImageIO/CGImagePropertyOrientation/rightMirrored
	kCGImagePropertyOrientationRightMirrored ImagePropertyOrientation = 0
	// kCGImagePropertyOrientationUp - The encoded image data matches the image’s intended display orientation.
	//
	// [Full Topic]: https://developer.apple.com/documentation/ImageIO/CGImagePropertyOrientation/up
	kCGImagePropertyOrientationUp ImagePropertyOrientation = 0
	// kCGImagePropertyOrientationUpMirrored - The encoded image data is horizontally flipped from the image’s intended display orientation.
	//
	// [Full Topic]: https://developer.apple.com/documentation/ImageIO/CGImagePropertyOrientation/upMirrored
	kCGImagePropertyOrientationUpMirrored ImagePropertyOrientation = 0
)

// CGImagePropertyTGACompression enum type
//
// [Full Topic]: https://developer.apple.com/documentation/ImageIO/CGImagePropertyTGACompression
type ImagePropertyTGACompression uint

const (
	//
	// [Full Topic]: https://developer.apple.com/documentation/ImageIO/CGImagePropertyTGACompression/tgaCompressionNone
	kCGImageTGACompressionNone ImagePropertyTGACompression = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/ImageIO/CGImagePropertyTGACompression/tgaCompressionRLE
	kCGImageTGACompressionRLE ImagePropertyTGACompression = 0
)

// CGImageSourceStatus - The set of status values for images and image sources.
//
// [Full Topic]: https://developer.apple.com/documentation/ImageIO/CGImageSourceStatus
type ImageSourceStatus uint

const (
	// kCGImageStatusComplete - The operation is complete.
	//
	// [Full Topic]: https://developer.apple.com/documentation/ImageIO/CGImageSourceStatus/statusComplete
	kCGImageStatusComplete ImageSourceStatus = 0
	// kCGImageStatusIncomplete - The operation is not complete
	//
	// [Full Topic]: https://developer.apple.com/documentation/ImageIO/CGImageSourceStatus/statusIncomplete
	kCGImageStatusIncomplete ImageSourceStatus = 0
	// kCGImageStatusInvalidData - The data is not valid.
	//
	// [Full Topic]: https://developer.apple.com/documentation/ImageIO/CGImageSourceStatus/statusInvalidData
	kCGImageStatusInvalidData ImageSourceStatus = 0
	// kCGImageStatusReadingHeader - The image source is reading the header.
	//
	// [Full Topic]: https://developer.apple.com/documentation/ImageIO/CGImageSourceStatus/statusReadingHeader
	kCGImageStatusReadingHeader ImageSourceStatus = 0
	// kCGImageStatusUnexpectedEOF - The end of the file occurred unexpectedly.
	//
	// [Full Topic]: https://developer.apple.com/documentation/ImageIO/CGImageSourceStatus/statusUnexpectedEOF
	kCGImageStatusUnexpectedEOF ImageSourceStatus = 0
	// kCGImageStatusUnknownType - The image is an unknown type.
	//
	// [Full Topic]: https://developer.apple.com/documentation/ImageIO/CGImageSourceStatus/statusUnknownType
	kCGImageStatusUnknownType ImageSourceStatus = 0
)


