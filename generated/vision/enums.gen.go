// Code generated from Apple documentation for Vision. DO NOT EDIT.

package vision

// Enum types and constants
// VNBarcodeCompositeType - Composite types for barcode requests.
//
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNBarcodeCompositeType
type BarcodeCompositeType uint

// VNChirality - Constants that the define the chirality, or handedness, of a pose.
//
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNChirality
type Chirality uint

// VNErrorCode - Constants that identify errors from the framework.
//
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNErrorCode
type ErrorCode uint

// VNGeneratePersonSegmentationRequestQualityLevel - Constants that define the levels of quality for a person segmentation request.
//
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNGeneratePersonSegmentationRequest/QualityLevel-swift.enum
type GeneratePersonSegmentationRequestQualityLevel uint

// VNImageCropAndScaleOption - Options that define how Vision crops and scales an input-image.
//
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNImageCropAndScaleOption
type ImageCropAndScaleOption uint

const (
// ImageCropAndScaleOptionCenterCrop - An option that scales the image to fit its shorter side within the input dimensions, while preserving its aspect ratio, and center-crops the image.
//
	// [Full Topic]: https://developer.apple.com/documentation/Vision/VNImageCropAndScaleOption/centerCrop
ImageCropAndScaleOptionCenterCrop ImageCropAndScaleOption = 0
// ImageCropAndScaleOptionScaleFill - An option that scales the image to fill the input dimensions, resizing it if necessary.
//
	// [Full Topic]: https://developer.apple.com/documentation/Vision/VNImageCropAndScaleOption/scaleFill
ImageCropAndScaleOptionScaleFill ImageCropAndScaleOption = 0
// ImageCropAndScaleOptionScaleFillRotate90CCW - An option that rotates the image 90 degrees counterclockwise and then scales it to fill the input dimensions.
//
	// [Full Topic]: https://developer.apple.com/documentation/Vision/VNImageCropAndScaleOption/scaleFillRotate90CCW
ImageCropAndScaleOptionScaleFillRotate90CCW ImageCropAndScaleOption = 0
// ImageCropAndScaleOptionScaleFit - An option that scales the image to fit its longer side within the input dimensions, while preserving its aspect ratio, and center-crops the image.
//
	// [Full Topic]: https://developer.apple.com/documentation/Vision/VNImageCropAndScaleOption/scaleFit
ImageCropAndScaleOptionScaleFit ImageCropAndScaleOption = 0
// ImageCropAndScaleOptionScaleFitRotate90CCW - An option that rotates the image 90 degrees counterclockwise and then scales it, while preserving its aspect ratio, to fit on the long side.
//
	// [Full Topic]: https://developer.apple.com/documentation/Vision/VNImageCropAndScaleOption/scaleFitRotate90CCW
ImageCropAndScaleOptionScaleFitRotate90CCW ImageCropAndScaleOption = 0
)

// VNRequestTextRecognitionLevel - Constants that identify the performance and accuracy of the text recognition.
//
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNRequestTextRecognitionLevel
type RequestTextRecognitionLevel uint

const (
// RequestTextRecognitionLevelAccurate - Accurate text recognition takes more time to produce a more comprehensive result.
//
	// [Full Topic]: https://developer.apple.com/documentation/Vision/VNRequestTextRecognitionLevel/accurate
RequestTextRecognitionLevelAccurate RequestTextRecognitionLevel = 0
// RequestTextRecognitionLevelFast - Fast text recognition returns results more quickly at the expense of accuracy.
//
	// [Full Topic]: https://developer.apple.com/documentation/Vision/VNRequestTextRecognitionLevel/fast
RequestTextRecognitionLevelFast RequestTextRecognitionLevel = 0
)


