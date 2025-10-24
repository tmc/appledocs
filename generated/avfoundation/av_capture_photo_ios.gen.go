//go:build darwin && ios

// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// iOS-only methods for CapturePhoto


// Gets a customized representation of the photo data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCapturePhoto/fileDataRepresentation(with:)
func (c_ CapturePhoto) FileDataRepresentationWithCustomizer(customizer unsafe.Pointer) foundation.Data {
	rv := objc.Send[foundation.Data](c_.ID, objc.Sel("fileDataRepresentationWithCustomizer:"), customizer)
	return rv
}

// Extracts and returns the captured photo’s preview image as a Core Graphics image object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCapturePhoto/previewCGImageRepresentation()
func (c_ CapturePhoto) PreviewCGImageRepresentation() ImageRef /* not a class type */ {
	rv := objc.Send[ImageRef](c_.ID, objc.Sel("previewCGImageRepresentation"))
	return rv
}

// Retrieves the semantic segmentation matte associated with this photo.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCapturePhoto/semanticSegmentationMatte(for:)
func (c_ CapturePhoto) SemanticSegmentationMatteForType(semanticSegmentationMatteType SemanticSegmentationMatteType /* typedef */) ISemanticSegmentationMatte {
	rv := objc.Send[SemanticSegmentationMatte](c_.ID, objc.Sel("semanticSegmentationMatteForType:"), semanticSegmentationMatteType)
	return rv
}

// iOS-only properties

// The variations available for bracketed capture settings for this photo.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCapturePhoto/bracketSettings
func (c_ CapturePhoto) BracketSettings() IAVCaptureBracketedStillImageSettings {
	rv := objc.Send[CaptureBracketedStillImageSettings](c_.ID, objc.Sel("bracketSettings"))
	return rv
}

// Calibration information for the camera device that captured the photo.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCapturePhoto/cameraCalibrationData
func (c_ CapturePhoto) CameraCalibrationData() IAVCameraCalibrationData {
	rv := objc.Send[CameraCalibrationData](c_.ID, objc.Sel("cameraCalibrationData"))
	return rv
}

// Depth or disparity map data captured with the photo.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCapturePhoto/depthData
func (c_ CapturePhoto) DepthData() IAVDepthData {
	rv := objc.Send[DepthData](c_.ID, objc.Sel("depthData"))
	return rv
}

// A dictionary describing the data format for a preview-sized image accompanying the captured photo.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCapturePhoto/embeddedThumbnailPhotoFormat
func (c_ CapturePhoto) EmbeddedThumbnailPhotoFormat() foundation.IDictionary {
	rv := objc.Send[foundation.IDictionary](c_.ID, objc.Sel("embeddedThumbnailPhotoFormat"))
	return rv
}

// A Boolean value indicating whether this photo object contains RAW format data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCapturePhoto/isRawPhoto
func (c_ CapturePhoto) RawPhoto() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("rawPhoto"))
	return rv
}

// Information about the use of lens stabilization during bracketed photo capture.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCapturePhoto/lensStabilizationStatus
func (c_ CapturePhoto) LensStabilizationStatus() CaptureLensStabilizationStatus {
	rv := objc.Send[CaptureLensStabilizationStatus](c_.ID, objc.Sel("lensStabilizationStatus"))
	return rv
}

// A dictionary of metadata describing the captured image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCapturePhoto/metadata
func (c_ CapturePhoto) Metadata() foundation.IDictionary {
	rv := objc.Send[foundation.IDictionary](c_.ID, objc.Sel("metadata"))
	return rv
}

// The portrait effects matte captured with the photo.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCapturePhoto/portraitEffectsMatte
func (c_ CapturePhoto) PortraitEffectsMatte() IAVPortraitEffectsMatte {
	rv := objc.Send[PortraitEffectsMatte](c_.ID, objc.Sel("portraitEffectsMatte"))
	return rv
}

// The pixel data for a preview-sized version of the photo, if requested.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCapturePhoto/previewPixelBuffer
func (c_ CapturePhoto) PreviewPixelBuffer() PixelBufferRef /* not a class type */ {
	rv := objc.Send[PixelBufferRef](c_.ID, objc.Sel("previewPixelBuffer"))
	return rv
}

// The 1-based index of this photo in a bracketed capture sequence.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCapturePhoto/sequenceCount
func (c_ CapturePhoto) SequenceCount() int {
	rv := objc.Send[int](c_.ID, objc.Sel("sequenceCount"))
	return rv
}

// The type of device that captured the photo.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCapturePhoto/sourceDeviceType
func (c_ CapturePhoto) SourceDeviceType() CaptureDeviceType /* typedef */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("sourceDeviceType"))
	return rv
}





