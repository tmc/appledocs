// Code generated from Apple documentation for ImageCaptureCore. DO NOT EDIT.

package imagecapturecore

// Enum types and constants
// ICLegacyReturnCode enum type
//
// [Full Topic]: https://developer.apple.com/documentation/ImageCaptureCore/ICLegacyReturn/Code
type ICLegacyReturnCode uint

const (
//
	// [Full Topic]: https://developer.apple.com/documentation/ImageCaptureCore/ICLegacyReturn/Code/cannotYieldDevice
ICLegacyReturnCodeCannotYieldDevice ICLegacyReturnCode = 0
//
	// [Full Topic]: https://developer.apple.com/documentation/ImageCaptureCore/ICLegacyReturn/Code/deviceLocationIDNotFoundErr
ICLegacyReturnCodeDeviceLocationIDNotFoundErr ICLegacyReturnCode = 0
//
	// [Full Topic]: https://developer.apple.com/documentation/ImageCaptureCore/ICLegacyReturn/Code/deviceNotFoundErr
ICLegacyReturnCodeDeviceNotFoundErr ICLegacyReturnCode = 0
//
	// [Full Topic]: https://developer.apple.com/documentation/ImageCaptureCore/ICLegacyReturn/Code/deviceUnsupportedErr
ICLegacyReturnCodeDeviceUnsupportedErr ICLegacyReturnCode = 0
//
	// [Full Topic]: https://developer.apple.com/documentation/ImageCaptureCore/ICLegacyReturn/Code/frameworkInternalErr
ICLegacyReturnCodeFrameworkInternalErr ICLegacyReturnCode = 0
//
	// [Full Topic]: https://developer.apple.com/documentation/ImageCaptureCore/ICLegacyReturn/Code/invalidObjectErr
ICLegacyReturnCodeInvalidObjectErr ICLegacyReturnCode = 0
//
	// [Full Topic]: https://developer.apple.com/documentation/ImageCaptureCore/ICLegacyReturn/Code/invalidPropertyErr
ICLegacyReturnCodeInvalidPropertyErr ICLegacyReturnCode = 0
//
	// [Full Topic]: https://developer.apple.com/documentation/ImageCaptureCore/ICLegacyReturn/Code/propertyTypeNotFoundErr
ICLegacyReturnCodePropertyTypeNotFoundErr ICLegacyReturnCode = 0
)

// ICReturnConnectionErrorCode enum type
//
// [Full Topic]: https://developer.apple.com/documentation/ImageCaptureCore/ICReturnConnectionError/Code
type ICReturnConnectionErrorCode uint

// ICReturnObjectErrorCode enum type
//
// [Full Topic]: https://developer.apple.com/documentation/ImageCaptureCore/ICReturnObjectError/Code
type ICReturnObjectErrorCode uint

// ICReturnThumbnailErrorCode enum type
//
// [Full Topic]: https://developer.apple.com/documentation/ImageCaptureCore/ICReturnThumbnailError/Code
type ICReturnThumbnailErrorCode uint

const (
// ICReturnThumbnailAlreadyFetching - Item thumbnail request is being serviced.
//
	// [Full Topic]: https://developer.apple.com/documentation/ImageCaptureCore/ICReturnThumbnailError/Code/alreadyFetching
ICReturnThumbnailAlreadyFetching ICReturnThumbnailErrorCode = 0
// ICReturnThumbnailCanceled - Item thumbnail request has been canceled.
//
	// [Full Topic]: https://developer.apple.com/documentation/ImageCaptureCore/ICReturnThumbnailError/Code/canceled
ICReturnThumbnailCanceled ICReturnThumbnailErrorCode = 0
// ICReturnThumbnailInvalid - Item thumbnail request completed with invalid result.
//
	// [Full Topic]: https://developer.apple.com/documentation/ImageCaptureCore/ICReturnThumbnailError/Code/invalid
ICReturnThumbnailInvalid ICReturnThumbnailErrorCode = 0
// ICReturnThumbnailNotAvailable - Item does not have thumbnail available.
//
	// [Full Topic]: https://developer.apple.com/documentation/ImageCaptureCore/ICReturnThumbnailError/Code/notAvailable
ICReturnThumbnailNotAvailable ICReturnThumbnailErrorCode = 0
)

// ICScannerBitDepth - The number of bits per channel in the scanned image.
//
// [Full Topic]: https://developer.apple.com/documentation/ImageCaptureCore/ICScannerBitDepth
type ICScannerBitDepth uint

const (
//
	// [Full Topic]: https://developer.apple.com/documentation/ImageCaptureCore/ICScannerBitDepth/depth16Bits
ICScannerBitDepth16Bits ICScannerBitDepth = 0
//
	// [Full Topic]: https://developer.apple.com/documentation/ImageCaptureCore/ICScannerBitDepth/depth1Bit
ICScannerBitDepth1Bit ICScannerBitDepth = 0
//
	// [Full Topic]: https://developer.apple.com/documentation/ImageCaptureCore/ICScannerBitDepth/depth8Bits
ICScannerBitDepth8Bits ICScannerBitDepth = 0
)

// ICScannerPixelDataType - The pixel data types.
//
// [Full Topic]: https://developer.apple.com/documentation/ImageCaptureCore/ICScannerPixelDataType
type ICScannerPixelDataType uint


