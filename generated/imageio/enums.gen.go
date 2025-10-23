// Code generated from Apple documentation for ImageIO. DO NOT EDIT.

package imageio

// Enum types and constants
// CGImageAnimationStatus - Constants that indicate the result of animating an image sequence.
//
// [Full Topic]: https://developer.apple.com/documentation/ImageIO/CGImageAnimationStatus
type CGImageAnimationStatus uint

// CGImageMetadataErrors - Constants for errors that occur when getting or setting metadata information.
//
// [Full Topic]: https://developer.apple.com/documentation/ImageIO/CGImageMetadataErrors
type CGImageMetadataErrors uint

const (
	// kCGImageMetadataErrorBadArgument - An error that indicates a parameter was malformed or contained invalid data.
	//
	// [Full Topic]: https://developer.apple.com/documentation/ImageIO/CGImageMetadataErrors/badArgument
	kCGImageMetadataErrorBadArgument CGImageMetadataErrors = 0
	// kCGImageMetadataErrorUnknown - An error that indicates an unknown condition occurred.
	//
	// [Full Topic]: https://developer.apple.com/documentation/ImageIO/CGImageMetadataErrors/unknown
	kCGImageMetadataErrorUnknown CGImageMetadataErrors = 0
)

// CGImagePropertyOrientation - A value describing the intended display orientation for an image.
//
// [Full Topic]: https://developer.apple.com/documentation/ImageIO/CGImagePropertyOrientation
type CGImagePropertyOrientation uint

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


