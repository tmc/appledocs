// Code generated from Apple documentation for ImageCaptureCore. DO NOT EDIT.

package imagecapturecore

// Enum types and constants
// ICMediaPresentation enum type
//
// [Full Topic]: https://developer.apple.com/documentation/ImageCaptureCore/ICMediaPresentation
type ICMediaPresentation uint

const (
	//
	// [Full Topic]: https://developer.apple.com/documentation/ImageCaptureCore/ICMediaPresentation/convertedAssets
	ICMediaPresentationConvertedAssets ICMediaPresentation = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/ImageCaptureCore/ICMediaPresentation/originalAssets
	ICMediaPresentationOriginalAssets ICMediaPresentation = 0
)

// ICReturnConnectionErrorCode enum type
//
// [Full Topic]: https://developer.apple.com/documentation/ImageCaptureCore/ICReturnConnectionError/Code
type ICReturnConnectionErrorCode uint

// ICReturnDownloadErrorCode enum type
//
// [Full Topic]: https://developer.apple.com/documentation/ImageCaptureCore/ICReturnDownloadError/Code
type ICReturnDownloadErrorCode uint

const (
	// ICReturnDownloadFileWritable - The destination file is not writable.
	//
	// [Full Topic]: https://developer.apple.com/documentation/ImageCaptureCore/ICReturnDownloadError/Code/fileWritable
	ICReturnDownloadFileWritable ICReturnDownloadErrorCode = 0
	// ICReturnDownloadPathInvalid - The destination path is invalid.
	//
	// [Full Topic]: https://developer.apple.com/documentation/ImageCaptureCore/ICReturnDownloadError/Code/pathInvalid
	ICReturnDownloadPathInvalid ICReturnDownloadErrorCode = 0
)


