//go:build darwin && ios

// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// iOS-only methods for CapturePhotoSettings


// iOS-only properties

// An array of pixel format types compatible with the photo settings for delivery of preview-sized images.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCapturePhotoSettings/availablePreviewPhotoPixelFormatTypes-2vfwu
func (c_ CapturePhotoSettings) AvailablePreviewPhotoPixelFormatTypes() []foundation.Number {
	rv := objc.Send[[]foundation.Number](c_.ID, objc.Sel("availablePreviewPhotoPixelFormatTypes"))
	return rv
}

// An array of video codec types compatible with the photo settings for embedding thumbnail images in photo file output.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCapturePhotoSettings/availableEmbeddedThumbnailPhotoCodecTypes
func (c_ CapturePhotoSettings) AvailableEmbeddedThumbnailPhotoCodecTypes() []string {
	rv := objc.Send[[]string](c_.ID, objc.Sel("availableEmbeddedThumbnailPhotoCodecTypes"))
	return rv
}

// An array of video codec types compatible with the photo settings for embedding raw thumbnail images in photo file output.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCapturePhotoSettings/availableRawEmbeddedThumbnailPhotoCodecTypes
func (c_ CapturePhotoSettings) AvailableRawEmbeddedThumbnailPhotoCodecTypes() []string {
	rv := objc.Send[[]string](c_.ID, objc.Sel("availableRawEmbeddedThumbnailPhotoCodecTypes"))
	return rv
}

// A dictionary describing the format for delivery of thumbnail images embedded in photo file output.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCapturePhotoSettings/embeddedThumbnailPhotoFormat
func (c_ CapturePhotoSettings) EmbeddedThumbnailPhotoFormat() foundation.IDictionary {
	rv := objc.Send[foundation.IDictionary](c_.ID, objc.Sel("embeddedThumbnailPhotoFormat"))
	return rv
}
func (c_ CapturePhotoSettings) SetEmbeddedThumbnailPhotoFormat(value foundation.IDictionary) {
	c_.ID.Send(objc.RegisterName("setEmbeddedThumbnailPhotoFormat:"), value)
}

// A Boolean value that determines whether any depth data captured with the photo is included when generating output file data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCapturePhotoSettings/embedsDepthDataInPhoto
func (c_ CapturePhotoSettings) EmbedsDepthDataInPhoto() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("embedsDepthDataInPhoto"))
	return rv
}
func (c_ CapturePhotoSettings) SetEmbedsDepthDataInPhoto(value bool) {
	c_.ID.Send(objc.RegisterName("setEmbedsDepthDataInPhoto:"), value)
}

// Specifies whether the portrait effects matte captured with ths photo should be written to the photo’s file structure.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCapturePhotoSettings/embedsPortraitEffectsMatteInPhoto
func (c_ CapturePhotoSettings) EmbedsPortraitEffectsMatteInPhoto() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("embedsPortraitEffectsMatteInPhoto"))
	return rv
}
func (c_ CapturePhotoSettings) SetEmbedsPortraitEffectsMatteInPhoto(value bool) {
	c_.ID.Send(objc.RegisterName("setEmbedsPortraitEffectsMatteInPhoto:"), value)
}

// A Boolean value that specifies whether to write the enabled semantic segmentation matte types captured with this photo to the photo’s file structure.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCapturePhotoSettings/embedsSemanticSegmentationMattesInPhoto
func (c_ CapturePhotoSettings) EmbedsSemanticSegmentationMattesInPhoto() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("embedsSemanticSegmentationMattesInPhoto"))
	return rv
}
func (c_ CapturePhotoSettings) SetEmbedsSemanticSegmentationMattesInPhoto(value bool) {
	c_.ID.Send(objc.RegisterName("setEmbedsSemanticSegmentationMattesInPhoto:"), value)
}

// An array of semantic segmentation matte types that the photo render pipeline can deliver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCapturePhotoSettings/enabledSemanticSegmentationMatteTypes
func (c_ CapturePhotoSettings) EnabledSemanticSegmentationMatteTypes() []string {
	rv := objc.Send[[]string](c_.ID, objc.Sel("enabledSemanticSegmentationMatteTypes"))
	return rv
}
func (c_ CapturePhotoSettings) SetEnabledSemanticSegmentationMatteTypes(value []string) {
	c_.ID.Send(objc.RegisterName("setEnabledSemanticSegmentationMatteTypes:"), value)
}

// A Boolean value that specifies whether the photo output, at its discretion, uses content-aware distortion correction on this photo request.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCapturePhotoSettings/isAutoContentAwareDistortionCorrectionEnabled
func (c_ CapturePhotoSettings) AutoContentAwareDistortionCorrectionEnabled() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("autoContentAwareDistortionCorrectionEnabled"))
	return rv
}
func (c_ CapturePhotoSettings) SetAutoContentAwareDistortionCorrectionEnabled(value bool) {
	c_.ID.Send(objc.RegisterName("setAutoContentAwareDistortionCorrectionEnabled:"), value)
}

// A Boolean value that specifies whether captures automatically combine data from a dual camera device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCapturePhotoSettings/isAutoDualCameraFusionEnabled
func (c_ CapturePhotoSettings) AutoDualCameraFusionEnabled() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("autoDualCameraFusionEnabled"))
	return rv
}
func (c_ CapturePhotoSettings) SetAutoDualCameraFusionEnabled(value bool) {
	c_.ID.Send(objc.RegisterName("setAutoDualCameraFusionEnabled:"), value)
}

// A Boolean value that indicates whether to use auto red-eye reduction on flash captures.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCapturePhotoSettings/isAutoRedEyeReductionEnabled
func (c_ CapturePhotoSettings) AutoRedEyeReductionEnabled() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("autoRedEyeReductionEnabled"))
	return rv
}
func (c_ CapturePhotoSettings) SetAutoRedEyeReductionEnabled(value bool) {
	c_.ID.Send(objc.RegisterName("setAutoRedEyeReductionEnabled:"), value)
}

// A Boolean value that specifies whether captures use automatic image stabilization.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCapturePhotoSettings/isAutoStillImageStabilizationEnabled
func (c_ CapturePhotoSettings) AutoStillImageStabilizationEnabled() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("autoStillImageStabilizationEnabled"))
	return rv
}
func (c_ CapturePhotoSettings) SetAutoStillImageStabilizationEnabled(value bool) {
	c_.ID.Send(objc.RegisterName("setAutoStillImageStabilizationEnabled:"), value)
}

// A Boolean value that specifies whether to use automatic virtual-device image fusion.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCapturePhotoSettings/isAutoVirtualDeviceFusionEnabled
func (c_ CapturePhotoSettings) AutoVirtualDeviceFusionEnabled() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("autoVirtualDeviceFusionEnabled"))
	return rv
}
func (c_ CapturePhotoSettings) SetAutoVirtualDeviceFusionEnabled(value bool) {
	c_.ID.Send(objc.RegisterName("setAutoVirtualDeviceFusionEnabled:"), value)
}

// A Boolean value that determines whether a dual photo capture also delivers camera calibration data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCapturePhotoSettings/isCameraCalibrationDataDeliveryEnabled
func (c_ CapturePhotoSettings) CameraCalibrationDataDeliveryEnabled() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("cameraCalibrationDataDeliveryEnabled"))
	return rv
}
func (c_ CapturePhotoSettings) SetCameraCalibrationDataDeliveryEnabled(value bool) {
	c_.ID.Send(objc.RegisterName("setCameraCalibrationDataDeliveryEnabled:"), value)
}

// A Boolean value that determines whether the photo output captures depth data along with the photo.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCapturePhotoSettings/isDepthDataDeliveryEnabled
func (c_ CapturePhotoSettings) DepthDataDeliveryEnabled() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("depthDataDeliveryEnabled"))
	return rv
}
func (c_ CapturePhotoSettings) SetDepthDataDeliveryEnabled(value bool) {
	c_.ID.Send(objc.RegisterName("setDepthDataDeliveryEnabled:"), value)
}

// A Boolean value that determines whether to smooth noise and fill in missing values in depth data output.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCapturePhotoSettings/isDepthDataFiltered
func (c_ CapturePhotoSettings) DepthDataFiltered() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("depthDataFiltered"))
	return rv
}
func (c_ CapturePhotoSettings) SetDepthDataFiltered(value bool) {
	c_.ID.Send(objc.RegisterName("setDepthDataFiltered:"), value)
}

// A Boolean value that determines whether a dual camera device delivers images from both cameras.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCapturePhotoSettings/isDualCameraDualPhotoDeliveryEnabled
func (c_ CapturePhotoSettings) DualCameraDualPhotoDeliveryEnabled() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("dualCameraDualPhotoDeliveryEnabled"))
	return rv
}
func (c_ CapturePhotoSettings) SetDualCameraDualPhotoDeliveryEnabled(value bool) {
	c_.ID.Send(objc.RegisterName("setDualCameraDualPhotoDeliveryEnabled:"), value)
}

// Specifies whether a portrait effects matte should be captured along with the photo.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCapturePhotoSettings/isPortraitEffectsMatteDeliveryEnabled
func (c_ CapturePhotoSettings) PortraitEffectsMatteDeliveryEnabled() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("portraitEffectsMatteDeliveryEnabled"))
	return rv
}
func (c_ CapturePhotoSettings) SetPortraitEffectsMatteDeliveryEnabled(value bool) {
	c_.ID.Send(objc.RegisterName("setPortraitEffectsMatteDeliveryEnabled:"), value)
}

// A URL at which to write Live Photo movie output.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCapturePhotoSettings/livePhotoMovieFileURL
func (c_ CapturePhotoSettings) LivePhotoMovieFileURL() objc.IObject /* cross-framework: NSURL */ {
	rv := objc.Send[foundation.NSURL](c_.ID, objc.Sel("livePhotoMovieFileURL"))
	return rv
}
func (c_ CapturePhotoSettings) SetLivePhotoMovieFileURL(value objc.IObject /* cross-framework: NSURL */) {
	c_.ID.Send(objc.RegisterName("setLivePhotoMovieFileURL:"), value)
}

// A dictionary of metadata to include in the Live Photo movie file.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCapturePhotoSettings/livePhotoMovieMetadata
func (c_ CapturePhotoSettings) LivePhotoMovieMetadata() []MetadataItem {
	rv := objc.Send[[]MetadataItem](c_.ID, objc.Sel("livePhotoMovieMetadata"))
	return rv
}
func (c_ CapturePhotoSettings) SetLivePhotoMovieMetadata(value []MetadataItem) {
	c_.ID.Send(objc.RegisterName("setLivePhotoMovieMetadata:"), value)
}

// The video codec to use for encoding the movie portion of Live Photo output.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCapturePhotoSettings/livePhotoVideoCodecType
func (c_ CapturePhotoSettings) LivePhotoVideoCodecType() VideoCodecType /* typedef */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("livePhotoVideoCodecType"))
	return rv
}
func (c_ CapturePhotoSettings) SetLivePhotoVideoCodecType(value VideoCodecType /* typedef */) {
	c_.ID.Send(objc.RegisterName("setLivePhotoVideoCodecType:"), value)
}

// A dictionary of metadata keys and values to embed in photo file output.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCapturePhotoSettings/metadata
func (c_ CapturePhotoSettings) Metadata() foundation.IDictionary {
	rv := objc.Send[foundation.IDictionary](c_.ID, objc.Sel("metadata"))
	return rv
}
func (c_ CapturePhotoSettings) SetMetadata(value foundation.IDictionary) {
	c_.ID.Send(objc.RegisterName("setMetadata:"), value)
}

// A dictionary describing the format for delivery of preview-sized images alongside the main photo.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCapturePhotoSettings/previewPhotoFormat
func (c_ CapturePhotoSettings) PreviewPhotoFormat() foundation.IDictionary {
	rv := objc.Send[foundation.IDictionary](c_.ID, objc.Sel("previewPhotoFormat"))
	return rv
}
func (c_ CapturePhotoSettings) SetPreviewPhotoFormat(value foundation.IDictionary) {
	c_.ID.Send(objc.RegisterName("setPreviewPhotoFormat:"), value)
}

// A dictionary describing the format for delivery of raw thumbnail images embedded in photo file output.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCapturePhotoSettings/rawEmbeddedThumbnailPhotoFormat
func (c_ CapturePhotoSettings) RawEmbeddedThumbnailPhotoFormat() foundation.IDictionary {
	rv := objc.Send[foundation.IDictionary](c_.ID, objc.Sel("rawEmbeddedThumbnailPhotoFormat"))
	return rv
}
func (c_ CapturePhotoSettings) SetRawEmbeddedThumbnailPhotoFormat(value foundation.IDictionary) {
	c_.ID.Send(objc.RegisterName("setRawEmbeddedThumbnailPhotoFormat:"), value)
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCapturePhotoSettings/rawFileFormat
func (c_ CapturePhotoSettings) RawFileFormat() foundation.IDictionary {
	rv := objc.Send[foundation.IDictionary](c_.ID, objc.Sel("rawFileFormat"))
	return rv
}
func (c_ CapturePhotoSettings) SetRawFileFormat(value foundation.IDictionary) {
	c_.ID.Send(objc.RegisterName("setRawFileFormat:"), value)
}

// The container file format for eventual output of the RAW image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCapturePhotoSettings/rawFileType
func (c_ CapturePhotoSettings) RawFileType() FileType /* typedef */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("rawFileType"))
	return rv
}

// An identifier for the Bayer RAW pixel format to deliver captured RAW photos in.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCapturePhotoSettings/rawPhotoPixelFormatType
func (c_ CapturePhotoSettings) RawPhotoPixelFormatType() uint32 /* not a class type */ {
	rv := objc.Send[uint32](c_.ID, objc.Sel("rawPhotoPixelFormatType"))
	return rv
}

// The constituent devices for which the virtual device should deliver photos.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCapturePhotoSettings/virtualDeviceConstituentPhotoDeliveryEnabledDevices
func (c_ CapturePhotoSettings) VirtualDeviceConstituentPhotoDeliveryEnabledDevices() []CaptureDevice {
	rv := objc.Send[[]CaptureDevice](c_.ID, objc.Sel("virtualDeviceConstituentPhotoDeliveryEnabledDevices"))
	return rv
}
func (c_ CapturePhotoSettings) SetVirtualDeviceConstituentPhotoDeliveryEnabledDevices(value []CaptureDevice) {
	c_.ID.Send(objc.RegisterName("setVirtualDeviceConstituentPhotoDeliveryEnabledDevices:"), value)
}




