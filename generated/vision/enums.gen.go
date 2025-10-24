// Code generated from Apple documentation for Vision. DO NOT EDIT.

package vision

/* debug [enums.gen.go]: Generating 13 enums for Vision */
// Enum types and constants
/* debug [enums.gen.go]: Processing enum VNBarcodeCompositeType (5 cases) */
// VNBarcodeCompositeType - Composite types for barcode requests.
//
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNBarcodeCompositeType
type VNBarcodeCompositeType uint

const (
	// VNBarcodeCompositeTypeGS1TypeA - A type that represents trade items in bulk.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Vision/VNBarcodeCompositeType/gs1TypeA
	VNBarcodeCompositeTypeGS1TypeA VNBarcodeCompositeType = 0
	// VNBarcodeCompositeTypeGS1TypeB - A type that represents trade items by piece.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Vision/VNBarcodeCompositeType/gs1TypeB
	VNBarcodeCompositeTypeGS1TypeB VNBarcodeCompositeType = 0
	// VNBarcodeCompositeTypeGS1TypeC - A type that represents trade items in varying quantity.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Vision/VNBarcodeCompositeType/gs1TypeC
	VNBarcodeCompositeTypeGS1TypeC VNBarcodeCompositeType = 0
	// VNBarcodeCompositeTypeLinked - A type that represents a linked composite type.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Vision/VNBarcodeCompositeType/linked
	VNBarcodeCompositeTypeLinked VNBarcodeCompositeType = 0
	// VNBarcodeCompositeTypeNone - A type that represents no composite type.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Vision/VNBarcodeCompositeType/none
	VNBarcodeCompositeTypeNone VNBarcodeCompositeType = 0
)

/* debug [enums.gen.go]: Processing enum VNChirality (3 cases) */
// VNChirality - Constants that the define the chirality, or handedness, of a pose.
//
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNChirality
type VNChirality uint

const (
	// VNChiralityLeft - Indicates a left-handed pose.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Vision/VNChirality/left
	VNChiralityLeft VNChirality = 0
	// VNChiralityRight - Indicates a right-handed pose.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Vision/VNChirality/right
	VNChiralityRight VNChirality = 0
	// VNChiralityUnknown - Indicates that the pose chirality is unknown.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Vision/VNChirality/unknown
	VNChiralityUnknown VNChirality = 0
)

/* debug [enums.gen.go]: Processing enum VNElementType (3 cases) */
// VNElementType - An enumeration of the type of element in feature print data.
//
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNElementType
type VNElementType uint

const (
	// VNElementTypeDouble - The elements are double-precision floating-point numbers.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Vision/VNElementType/double
	VNElementTypeDouble VNElementType = 0
	// VNElementTypeFloat - The elements are floating-point numbers.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Vision/VNElementType/float
	VNElementTypeFloat VNElementType = 0
	// VNElementTypeUnknown - The element type isn’t known.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Vision/VNElementType/unknown
	VNElementTypeUnknown VNElementType = 0
)

/* debug [enums.gen.go]: Processing enum VNErrorCode (24 cases) */
// VNErrorCode - Constants that identify errors from the framework.
//
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNErrorCode
type VNErrorCode uint

const (
	// VNErrorDataUnavailable - The data isn’t available.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Vision/VNErrorCode/dataUnavailable
	VNErrorDataUnavailable VNErrorCode = 0
	// VNErrorInternalError - An internal error occurred within the framework.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Vision/VNErrorCode/internalError
	VNErrorInternalError VNErrorCode = 0
	// VNErrorInvalidArgument - An app passed an invalid parameter to a request.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Vision/VNErrorCode/invalidArgument
	VNErrorInvalidArgument VNErrorCode = 0
	// VNErrorInvalidFormat - The format of the image is invalid.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Vision/VNErrorCode/invalidFormat
	VNErrorInvalidFormat VNErrorCode = 0
	// VNErrorInvalidImage - The image is invalid.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Vision/VNErrorCode/invalidImage
	VNErrorInvalidImage VNErrorCode = 0
	// VNErrorInvalidModel - The Core ML model is incompatible with the request.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Vision/VNErrorCode/invalidModel
	VNErrorInvalidModel VNErrorCode = 0
	// VNErrorInvalidOperation - An app requested an unsupported operation.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Vision/VNErrorCode/invalidOperation
	VNErrorInvalidOperation VNErrorCode = 0
	// VNErrorInvalidOption - An app specified an invalid option on a request.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Vision/VNErrorCode/invalidOption
	VNErrorInvalidOption VNErrorCode = 0
	// VNErrorIOError - An I/O error for an image, image sequence, or Core ML model.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Vision/VNErrorCode/ioError
	VNErrorIOError VNErrorCode = 0
	// VNErrorMissingOption - A request is missing a required option.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Vision/VNErrorCode/missingOption
	VNErrorMissingOption VNErrorCode = 0
	// VNErrorNotImplemented - The method isn’t implemented in the underlying model.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Vision/VNErrorCode/notImplemented
	VNErrorNotImplemented VNErrorCode = 0
	// VNErrorOK - The operation finished without error.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Vision/VNErrorCode/OK
	VNErrorOK VNErrorCode = 0
	// VNErrorOperationFailed - The requested operation failed.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Vision/VNErrorCode/operationFailed
	VNErrorOperationFailed VNErrorCode = 0
	// VNErrorOutOfBoundsError - An app attempted to access data that’s out-of-bounds.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Vision/VNErrorCode/outOfBoundsError
	VNErrorOutOfBoundsError VNErrorCode = 0
	// VNErrorOutOfMemory - The system doesn’t have enough memory to complete the request.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Vision/VNErrorCode/outOfMemory
	VNErrorOutOfMemory VNErrorCode = 0
	// VNErrorRequestCancelled - An app canceled the request.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Vision/VNErrorCode/requestCancelled
	VNErrorRequestCancelled VNErrorCode = 0
	// VNErrorTimeout - The requested operation timed out.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Vision/VNErrorCode/timeout
	VNErrorTimeout VNErrorCode = 0
	// VNErrorTimeStampNotFound - The system can’t find a timestamp.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Vision/VNErrorCode/timeStampNotFound
	VNErrorTimeStampNotFound VNErrorCode = 0
	// VNErrorTuriCoreErrorCode - An error occurred during Create ML training due to an invalid transformation or image.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Vision/VNErrorCode/turiCoreErrorCode
	VNErrorTuriCoreErrorCode VNErrorCode = 0
	// VNErrorUnknownError - An unidentified error occurred.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Vision/VNErrorCode/unknownError
	VNErrorUnknownError VNErrorCode = 0
	// VNErrorUnsupportedComputeDevice - An app requested an unsupported compute device.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Vision/VNErrorCode/unsupportedComputeDevice
	VNErrorUnsupportedComputeDevice VNErrorCode = 0
	// VNErrorUnsupportedComputeStage - An app requested an unsupported compute stage.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Vision/VNErrorCode/unsupportedComputeStage
	VNErrorUnsupportedComputeStage VNErrorCode = 0
	// VNErrorUnsupportedRequest - An app attempted an unsupported request.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Vision/VNErrorCode/unsupportedRequest
	VNErrorUnsupportedRequest VNErrorCode = 0
	// VNErrorUnsupportedRevision - An app specified an unsupported request revision.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Vision/VNErrorCode/unsupportedRevision
	VNErrorUnsupportedRevision VNErrorCode = 0
)

/* debug [enums.gen.go]: Processing enum VNGenerateOpticalFlowRequestComputationAccuracy (4 cases) */
// VNGenerateOpticalFlowRequestComputationAccuracy - The supported optical flow accuracy levels.
//
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNGenerateOpticalFlowRequest/ComputationAccuracy-swift.enum
type VNGenerateOpticalFlowRequestComputationAccuracy uint

const (
	// VNGenerateOpticalFlowRequestComputationAccuracyHigh - High accuracy.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Vision/VNGenerateOpticalFlowRequest/ComputationAccuracy-swift.enum/high
	VNGenerateOpticalFlowRequestComputationAccuracyHigh VNGenerateOpticalFlowRequestComputationAccuracy = 0
	// VNGenerateOpticalFlowRequestComputationAccuracyLow - Low accuracy.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Vision/VNGenerateOpticalFlowRequest/ComputationAccuracy-swift.enum/low
	VNGenerateOpticalFlowRequestComputationAccuracyLow VNGenerateOpticalFlowRequestComputationAccuracy = 0
	// VNGenerateOpticalFlowRequestComputationAccuracyMedium - Medium accuracy.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Vision/VNGenerateOpticalFlowRequest/ComputationAccuracy-swift.enum/medium
	VNGenerateOpticalFlowRequestComputationAccuracyMedium VNGenerateOpticalFlowRequestComputationAccuracy = 0
	// VNGenerateOpticalFlowRequestComputationAccuracyVeryHigh - Very high accuracy.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Vision/VNGenerateOpticalFlowRequest/ComputationAccuracy-swift.enum/veryHigh
	VNGenerateOpticalFlowRequestComputationAccuracyVeryHigh VNGenerateOpticalFlowRequestComputationAccuracy = 0
)

/* debug [enums.gen.go]: Processing enum VNGeneratePersonSegmentationRequestQualityLevel (3 cases) */
// VNGeneratePersonSegmentationRequestQualityLevel - Constants that define the levels of quality for a person segmentation request.
//
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNGeneratePersonSegmentationRequest/QualityLevel-swift.enum
type VNGeneratePersonSegmentationRequestQualityLevel uint

const (
	// VNGeneratePersonSegmentationRequestQualityLevelAccurate - Prefers image quality over performance.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Vision/VNGeneratePersonSegmentationRequest/QualityLevel-swift.enum/accurate
	VNGeneratePersonSegmentationRequestQualityLevelAccurate VNGeneratePersonSegmentationRequestQualityLevel = 0
	// VNGeneratePersonSegmentationRequestQualityLevelBalanced - Prefers processing that balances image quality and performance.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Vision/VNGeneratePersonSegmentationRequest/QualityLevel-swift.enum/balanced
	VNGeneratePersonSegmentationRequestQualityLevelBalanced VNGeneratePersonSegmentationRequestQualityLevel = 0
	// VNGeneratePersonSegmentationRequestQualityLevelFast - Prefers performance over image quality.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Vision/VNGeneratePersonSegmentationRequest/QualityLevel-swift.enum/fast
	VNGeneratePersonSegmentationRequestQualityLevelFast VNGeneratePersonSegmentationRequestQualityLevel = 0
)

/* debug [enums.gen.go]: Processing enum VNHumanBodyPose3DObservationHeightEstimation (2 cases) */
// VNHumanBodyPose3DObservationHeightEstimation - Constants that identify body height estimation techniques.
//
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNHumanBodyPose3DObservation/HeightEstimation-swift.enum
type VNHumanBodyPose3DObservationHeightEstimation uint

const (
	// VNHumanBodyPose3DObservationHeightEstimationMeasured - A technique that uses LiDAR depth data to measure body height, in meters.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Vision/VNHumanBodyPose3DObservation/HeightEstimation-swift.enum/measured
	VNHumanBodyPose3DObservationHeightEstimationMeasured VNHumanBodyPose3DObservationHeightEstimation = 0
	// VNHumanBodyPose3DObservationHeightEstimationReference - A technique that uses a reference height.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Vision/VNHumanBodyPose3DObservation/HeightEstimation-swift.enum/reference
	VNHumanBodyPose3DObservationHeightEstimationReference VNHumanBodyPose3DObservationHeightEstimation = 0
)

/* debug [enums.gen.go]: Processing enum VNImageCropAndScaleOption (5 cases) */
// VNImageCropAndScaleOption - Options that define how Vision crops and scales an input-image.
//
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNImageCropAndScaleOption
type VNImageCropAndScaleOption uint

const (
	// VNImageCropAndScaleOptionCenterCrop - An option that scales the image to fit its shorter side within the input dimensions, while preserving its aspect ratio, and center-crops the image.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Vision/VNImageCropAndScaleOption/centerCrop
	VNImageCropAndScaleOptionCenterCrop VNImageCropAndScaleOption = 0
	// VNImageCropAndScaleOptionScaleFill - An option that scales the image to fill the input dimensions, resizing it if necessary.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Vision/VNImageCropAndScaleOption/scaleFill
	VNImageCropAndScaleOptionScaleFill VNImageCropAndScaleOption = 0
	// VNImageCropAndScaleOptionScaleFillRotate90CCW - An option that rotates the image 90 degrees counterclockwise and then scales it to fill the input dimensions.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Vision/VNImageCropAndScaleOption/scaleFillRotate90CCW
	VNImageCropAndScaleOptionScaleFillRotate90CCW VNImageCropAndScaleOption = 0
	// VNImageCropAndScaleOptionScaleFit - An option that scales the image to fit its longer side within the input dimensions, while preserving its aspect ratio, and center-crops the image.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Vision/VNImageCropAndScaleOption/scaleFit
	VNImageCropAndScaleOptionScaleFit VNImageCropAndScaleOption = 0
	// VNImageCropAndScaleOptionScaleFitRotate90CCW - An option that rotates the image 90 degrees counterclockwise and then scales it, while preserving its aspect ratio, to fit on the long side.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Vision/VNImageCropAndScaleOption/scaleFitRotate90CCW
	VNImageCropAndScaleOptionScaleFitRotate90CCW VNImageCropAndScaleOption = 0
)

/* debug [enums.gen.go]: Processing enum VNPointsClassification (3 cases) */
// VNPointsClassification - The set of classifications that describe how to interpret the points the region provides.
//
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNPointsClassification
type VNPointsClassification uint

const (
	//
	// [Full Topic]: https://developer.apple.com/documentation/Vision/VNPointsClassification/closedPath
	VNPointsClassificationClosedPath VNPointsClassification = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/Vision/VNPointsClassification/disconnected
	VNPointsClassificationDisconnected VNPointsClassification = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/Vision/VNPointsClassification/openPath
	VNPointsClassificationOpenPath VNPointsClassification = 0
)

/* debug [enums.gen.go]: Processing enum VNRequestFaceLandmarksConstellation (3 cases) */
// VNRequestFaceLandmarksConstellation - An enumeration of face landmarks in a constellation object.
//
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNRequestFaceLandmarksConstellation
type VNRequestFaceLandmarksConstellation uint

const (
	// VNRequestFaceLandmarksConstellation65Points - A constellation with 65 points.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Vision/VNRequestFaceLandmarksConstellation/constellation65Points
	VNRequestFaceLandmarksConstellation65Points VNRequestFaceLandmarksConstellation = 0
	// VNRequestFaceLandmarksConstellation76Points - A constellation with 76 points.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Vision/VNRequestFaceLandmarksConstellation/constellation76Points
	VNRequestFaceLandmarksConstellation76Points VNRequestFaceLandmarksConstellation = 0
	// VNRequestFaceLandmarksConstellationNotDefined - An undefined constellation.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Vision/VNRequestFaceLandmarksConstellation/constellationNotDefined
	VNRequestFaceLandmarksConstellationNotDefined VNRequestFaceLandmarksConstellation = 0
)

/* debug [enums.gen.go]: Processing enum VNRequestTextRecognitionLevel (2 cases) */
// VNRequestTextRecognitionLevel - Constants that identify the performance and accuracy of the text recognition.
//
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNRequestTextRecognitionLevel
type VNRequestTextRecognitionLevel uint

const (
	// VNRequestTextRecognitionLevelAccurate - Accurate text recognition takes more time to produce a more comprehensive result.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Vision/VNRequestTextRecognitionLevel/accurate
	VNRequestTextRecognitionLevelAccurate VNRequestTextRecognitionLevel = 0
	// VNRequestTextRecognitionLevelFast - Fast text recognition returns results more quickly at the expense of accuracy.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Vision/VNRequestTextRecognitionLevel/fast
	VNRequestTextRecognitionLevelFast VNRequestTextRecognitionLevel = 0
)

/* debug [enums.gen.go]: Processing enum VNRequestTrackingLevel (2 cases) */
// VNRequestTrackingLevel - An enumeration of tracking priorities.
//
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNRequestTrackingLevel
type VNRequestTrackingLevel uint

const (
	// VNRequestTrackingLevelAccurate - Tracking level that favors location accuracy over speed.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Vision/VNRequestTrackingLevel/accurate
	VNRequestTrackingLevelAccurate VNRequestTrackingLevel = 0
	// VNRequestTrackingLevelFast - Tracking level that favors speed over location accuracy.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Vision/VNRequestTrackingLevel/fast
	VNRequestTrackingLevelFast VNRequestTrackingLevel = 0
)

/* debug [enums.gen.go]: Processing enum VNTrackOpticalFlowRequestComputationAccuracy (4 cases) */
// VNTrackOpticalFlowRequestComputationAccuracy - Computational accuracy options.
//
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNTrackOpticalFlowRequest/ComputationAccuracy-swift.enum
type VNTrackOpticalFlowRequestComputationAccuracy uint

const (
	// VNTrackOpticalFlowRequestComputationAccuracyHigh - An option that indicates a high level of computational accuracy.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Vision/VNTrackOpticalFlowRequest/ComputationAccuracy-swift.enum/high
	VNTrackOpticalFlowRequestComputationAccuracyHigh VNTrackOpticalFlowRequestComputationAccuracy = 0
	// VNTrackOpticalFlowRequestComputationAccuracyLow - An option that indicates a low level of computational accuracy.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Vision/VNTrackOpticalFlowRequest/ComputationAccuracy-swift.enum/low
	VNTrackOpticalFlowRequestComputationAccuracyLow VNTrackOpticalFlowRequestComputationAccuracy = 0
	// VNTrackOpticalFlowRequestComputationAccuracyMedium - An option that indicates a moderate level of computational accuracy.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Vision/VNTrackOpticalFlowRequest/ComputationAccuracy-swift.enum/medium
	VNTrackOpticalFlowRequestComputationAccuracyMedium VNTrackOpticalFlowRequestComputationAccuracy = 0
	// VNTrackOpticalFlowRequestComputationAccuracyVeryHigh - An option that indicates a very high level of computational accuracy.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Vision/VNTrackOpticalFlowRequest/ComputationAccuracy-swift.enum/veryHigh
	VNTrackOpticalFlowRequestComputationAccuracyVeryHigh VNTrackOpticalFlowRequestComputationAccuracy = 0
)


