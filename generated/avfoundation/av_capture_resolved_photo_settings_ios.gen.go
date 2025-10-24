//go:build darwin && ios

// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// iOS-only methods for CaptureResolvedPhotoSettings


// Retrieves the resolved dimensions of the semantic segmentation mattes that the photo output delivers.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureResolvedPhotoSettings/dimensionsForSemanticSegmentationMatte(ofType:)
func (c_ CaptureResolvedPhotoSettings) DimensionsForSemanticSegmentationMatteOfType(semanticSegmentationMatteType SemanticSegmentationMatteType /* typedef */) VideoDimensions /* not a class type */ {
	rv := objc.Send[VideoDimensions](c_.ID, objc.Sel("dimensionsForSemanticSegmentationMatteOfType:"), semanticSegmentationMatteType)
	return rv
}

// iOS-only properties

// The resolved dimensions of the photo proxy when using deferred photo delivery.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureResolvedPhotoSettings/deferredPhotoProxyDimensions
func (c_ CaptureResolvedPhotoSettings) DeferredPhotoProxyDimensions() VideoDimensions /* not a class type */ {
	rv := objc.Send[VideoDimensions](c_.ID, objc.Sel("deferredPhotoProxyDimensions"))
	return rv
}

// The size, in pixels, of the thumbnail image that the capture delivers.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureResolvedPhotoSettings/embeddedThumbnailDimensions
func (c_ CaptureResolvedPhotoSettings) EmbeddedThumbnailDimensions() VideoDimensions /* not a class type */ {
	rv := objc.Send[VideoDimensions](c_.ID, objc.Sel("embeddedThumbnailDimensions"))
	return rv
}

// A Boolean value that indicates whether the system applies content-aware distortion correction when capturing the photo.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureResolvedPhotoSettings/isContentAwareDistortionCorrectionEnabled
func (c_ CaptureResolvedPhotoSettings) ContentAwareDistortionCorrectionEnabled() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("contentAwareDistortionCorrectionEnabled"))
	return rv
}

// A Boolean value indicating whether this capture combines image data from a dual camera.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureResolvedPhotoSettings/isDualCameraFusionEnabled
func (c_ CaptureResolvedPhotoSettings) DualCameraFusionEnabled() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("dualCameraFusionEnabled"))
	return rv
}

// A Boolean value indicating whether the camera flash fires for this capture.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureResolvedPhotoSettings/isFlashEnabled
func (c_ CaptureResolvedPhotoSettings) FlashEnabled() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("flashEnabled"))
	return rv
}

// A Boolean value indicating whether the camera automatically reduces red-eye when capturing photos.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureResolvedPhotoSettings/isRedEyeReductionEnabled
func (c_ CaptureResolvedPhotoSettings) RedEyeReductionEnabled() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("redEyeReductionEnabled"))
	return rv
}

// A Boolean value indicating whether this capture uses image stabilization.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureResolvedPhotoSettings/isStillImageStabilizationEnabled
func (c_ CaptureResolvedPhotoSettings) StillImageStabilizationEnabled() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("stillImageStabilizationEnabled"))
	return rv
}

// A Boolean value that specifies whether the system automatically uses virtual device image fusion.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureResolvedPhotoSettings/isVirtualDeviceFusionEnabled
func (c_ CaptureResolvedPhotoSettings) VirtualDeviceFusionEnabled() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("virtualDeviceFusionEnabled"))
	return rv
}

// The size, in pixels, of the Live Photo movie content that the capture delivers.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureResolvedPhotoSettings/livePhotoMovieDimensions
func (c_ CaptureResolvedPhotoSettings) LivePhotoMovieDimensions() VideoDimensions /* not a class type */ {
	rv := objc.Send[VideoDimensions](c_.ID, objc.Sel("livePhotoMovieDimensions"))
	return rv
}

// The time range in which to expect the system to deliver the photo to the delegate.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureResolvedPhotoSettings/photoProcessingTimeRange
func (c_ CaptureResolvedPhotoSettings) PhotoProcessingTimeRange() TimeRange /* not a class type */ {
	rv := objc.Send[TimeRange](c_.ID, objc.Sel("photoProcessingTimeRange"))
	return rv
}

// The size, in pixels, of the portrait effects matte that the capture delivers.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureResolvedPhotoSettings/portraitEffectsMatteDimensions
func (c_ CaptureResolvedPhotoSettings) PortraitEffectsMatteDimensions() VideoDimensions /* not a class type */ {
	rv := objc.Send[VideoDimensions](c_.ID, objc.Sel("portraitEffectsMatteDimensions"))
	return rv
}

// The size, in pixels, of the preview image that the system delivers with the capture.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureResolvedPhotoSettings/previewDimensions
func (c_ CaptureResolvedPhotoSettings) PreviewDimensions() VideoDimensions /* not a class type */ {
	rv := objc.Send[VideoDimensions](c_.ID, objc.Sel("previewDimensions"))
	return rv
}

// The size, in pixels, of the RAW-format embedded thumbnail image that the capture delivers.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureResolvedPhotoSettings/rawEmbeddedThumbnailDimensions
func (c_ CaptureResolvedPhotoSettings) RawEmbeddedThumbnailDimensions() VideoDimensions /* not a class type */ {
	rv := objc.Send[VideoDimensions](c_.ID, objc.Sel("rawEmbeddedThumbnailDimensions"))
	return rv
}

// The size, in pixels, of the RAW-format photo image that the capture delivers.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureResolvedPhotoSettings/rawPhotoDimensions
func (c_ CaptureResolvedPhotoSettings) RawPhotoDimensions() VideoDimensions /* not a class type */ {
	rv := objc.Send[VideoDimensions](c_.ID, objc.Sel("rawPhotoDimensions"))
	return rv
}





