// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [CapturePhotoSettings] class.
var (
	CapturePhotoSettingsClass     _CapturePhotoSettingsClass
	CapturePhotoSettingsClassOnce sync.Once
)

func getCapturePhotoSettingsClass() _CapturePhotoSettingsClass {
	CapturePhotoSettingsClassOnce.Do(func() {
		CapturePhotoSettingsClass = _CapturePhotoSettingsClass{objc.GetClass("AVCapturePhotoSettings")}
	})
	return CapturePhotoSettingsClass
}

type _CapturePhotoSettingsClass struct {
	class objc.Class
}

// An interface definition for the [CapturePhotoSettings] class.
type ICapturePhotoSettings interface {
	objectivec.IObject
	PortraitEffectsMatteDeliveryEnabled() bool
	SetPortraitEffectsMatteDeliveryEnabled(value bool)
	AvailableEmbeddedThumbnailPhotoCodecTypes() VideoCodecType
	SetAvailableEmbeddedThumbnailPhotoCodecTypes(value VideoCodecType)
	AvailablePreviewPhotoPixelFormatTypes() unsafe.Pointer
	SetAvailablePreviewPhotoPixelFormatTypes(value unsafe.Pointer)
	AvailableRawEmbeddedThumbnailPhotoCodecTypes() VideoCodecType
	SetAvailableRawEmbeddedThumbnailPhotoCodecTypes(value VideoCodecType)
	EmbeddedThumbnailPhotoFormat() string
	SetEmbeddedThumbnailPhotoFormat(value string)
	EmbedsDepthDataInPhoto() bool
	SetEmbedsDepthDataInPhoto(value bool)
	EmbedsPortraitEffectsMatteInPhoto() bool
	SetEmbedsPortraitEffectsMatteInPhoto(value bool)
	EmbedsSemanticSegmentationMattesInPhoto() bool
	SetEmbedsSemanticSegmentationMattesInPhoto(value bool)
	EnabledSemanticSegmentationMatteTypes() unsafe.Pointer
	SetEnabledSemanticSegmentationMatteTypes(value unsafe.Pointer)
	FlashMode() unsafe.Pointer
	SetFlashMode(value unsafe.Pointer)
	Format() string
	SetFormat(value string)
	IsAutoContentAwareDistortionCorrectionEnabled() bool
	SetIsAutoContentAwareDistortionCorrectionEnabled(value bool)
	IsAutoDualCameraFusionEnabled() bool
	SetIsAutoDualCameraFusionEnabled(value bool)
	IsAutoRedEyeReductionEnabled() bool
	SetIsAutoRedEyeReductionEnabled(value bool)
	IsAutoStillImageStabilizationEnabled() bool
	SetIsAutoStillImageStabilizationEnabled(value bool)
	IsAutoVirtualDeviceFusionEnabled() bool
	SetIsAutoVirtualDeviceFusionEnabled(value bool)
	IsCameraCalibrationDataDeliveryEnabled() bool
	SetIsCameraCalibrationDataDeliveryEnabled(value bool)
	IsConstantColorEnabled() bool
	SetIsConstantColorEnabled(value bool)
	IsConstantColorFallbackPhotoDeliveryEnabled() bool
	SetIsConstantColorFallbackPhotoDeliveryEnabled(value bool)
	IsDepthDataDeliveryEnabled() bool
	SetIsDepthDataDeliveryEnabled(value bool)
	IsDepthDataFiltered() bool
	SetIsDepthDataFiltered(value bool)
	IsDualCameraDualPhotoDeliveryEnabled() bool
	SetIsDualCameraDualPhotoDeliveryEnabled(value bool)
	IsHighResolutionPhotoEnabled() bool
	SetIsHighResolutionPhotoEnabled(value bool)
	IsPortraitEffectsMatteDeliveryEnabled() bool
	SetIsPortraitEffectsMatteDeliveryEnabled(value bool)
	IsShutterSoundSuppressionEnabled() bool
	SetIsShutterSoundSuppressionEnabled(value bool)
	LivePhotoMovieFileURL() foundation.URL
	SetLivePhotoMovieFileURL(value foundation.IURL)
	LivePhotoMovieMetadata() AVMetadataItem
	SetLivePhotoMovieMetadata(value IAVMetadataItem)
	LivePhotoVideoCodecType() VideoCodecType
	SetLivePhotoVideoCodecType(value VideoCodecType)
	MaxPhotoDimensions() unsafe.Pointer
	SetMaxPhotoDimensions(value unsafe.Pointer)
	Metadata() string
	SetMetadata(value string)
	PhotoQualityPrioritization() unsafe.Pointer
	SetPhotoQualityPrioritization(value unsafe.Pointer)
	PreviewPhotoFormat() string
	SetPreviewPhotoFormat(value string)
	ProcessedFileType() FileType
	SetProcessedFileType(value FileType)
	RawEmbeddedThumbnailPhotoFormat() string
	SetRawEmbeddedThumbnailPhotoFormat(value string)
	RawFileFormat() string
	SetRawFileFormat(value string)
	RawFileType() FileType
	SetRawFileType(value FileType)
	RawPhotoPixelFormatType() unsafe.Pointer
	SetRawPhotoPixelFormatType(value unsafe.Pointer)
	UniqueID() unsafe.Pointer
	SetUniqueID(value unsafe.Pointer)
	VirtualDeviceConstituentPhotoDeliveryEnabledDevices() AVCaptureDevice
	SetVirtualDeviceConstituentPhotoDeliveryEnabledDevices(value IAVCaptureDevice)
}

// A specification of the features and settings to use for a single photo capture request.
//
// To take a photo, you create and configure a object, then pass it to the method. A instance can include any combination of settings, regardless of whether that combination is valid for a given capture session. When you initiate a capture by passing a photo settings object to the method, the photo capture output validates your settings to ensure deterministic behavior. For example, the setting must specify a value that’s present in the photo output’s array. For detailed validation rules, see each property description below.


// A specification of the features and settings to use for a single photo capture request.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCapturePhotoSettings

type CapturePhotoSettings struct {
	objectivec.Object
}

// CapturePhotoSettingsFrom constructs a [CapturePhotoSettings] from an unsafe.Pointer.
//
// A specification of the features and settings to use for a single photo capture request.
func CapturePhotoSettingsFrom(ptr unsafe.Pointer) CapturePhotoSettings {
	return CapturePhotoSettings{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (cc _CapturePhotoSettingsClass) Alloc() CapturePhotoSettings {
	rv := objc.Send[CapturePhotoSettings](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (cc _CapturePhotoSettingsClass) New() CapturePhotoSettings {
	rv := objc.Send[CapturePhotoSettings](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CapturePhotoSettings) Init() CapturePhotoSettings {
	rv := objc.Send[CapturePhotoSettings](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CapturePhotoSettings) Autorelease() CapturePhotoSettings {
	rv := objc.Send[CapturePhotoSettings](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCapturePhotoSettings creates a new CapturePhotoSettings instance.
func NewCapturePhotoSettings() CapturePhotoSettings {
	return getCapturePhotoSettingsClass().New()
}



// Specifies whether a portrait effects matte should be captured along with the photo.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCapturePhotoSettings/isPortraitEffectsMatteDeliveryEnabled

func (c_ CapturePhotoSettings) PortraitEffectsMatteDeliveryEnabled() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("portraitEffectsMatteDeliveryEnabled"))
	return rv
}


// Specifies whether a portrait effects matte should be captured along with the photo.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCapturePhotoSettings/isPortraitEffectsMatteDeliveryEnabled

func (c_ CapturePhotoSettings) SetPortraitEffectsMatteDeliveryEnabled(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setPortraitEffectsMatteDeliveryEnabled:"), value)
}


// An array of video codec types compatible with the photo settings for embedding thumbnail images in photo file output.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturephotosettings/availableembeddedthumbnailphotocodectypes

func (c_ CapturePhotoSettings) AvailableEmbeddedThumbnailPhotoCodecTypes() VideoCodecType {
	rv := objc.Send[VideoCodecType](c_.ID, objc.Sel("availableEmbeddedThumbnailPhotoCodecTypes"))
	return rv
}


// An array of video codec types compatible with the photo settings for embedding thumbnail images in photo file output.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturephotosettings/availableembeddedthumbnailphotocodectypes

func (c_ CapturePhotoSettings) SetAvailableEmbeddedThumbnailPhotoCodecTypes(value VideoCodecType) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setAvailableEmbeddedThumbnailPhotoCodecTypes:"), value)
}


// An array of available of pixel format types available to specify a preview photo format.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturephotosettings/availablepreviewphotopixelformattypes-30d9

func (c_ CapturePhotoSettings) AvailablePreviewPhotoPixelFormatTypes() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("availablePreviewPhotoPixelFormatTypes"))
	return rv
}


// An array of available of pixel format types available to specify a preview photo format.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturephotosettings/availablepreviewphotopixelformattypes-30d9

func (c_ CapturePhotoSettings) SetAvailablePreviewPhotoPixelFormatTypes(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setAvailablePreviewPhotoPixelFormatTypes:"), value)
}


// An array of video codec types compatible with the photo settings for embedding raw thumbnail images in photo file output.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturephotosettings/availablerawembeddedthumbnailphotocodectypes

func (c_ CapturePhotoSettings) AvailableRawEmbeddedThumbnailPhotoCodecTypes() VideoCodecType {
	rv := objc.Send[VideoCodecType](c_.ID, objc.Sel("availableRawEmbeddedThumbnailPhotoCodecTypes"))
	return rv
}


// An array of video codec types compatible with the photo settings for embedding raw thumbnail images in photo file output.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturephotosettings/availablerawembeddedthumbnailphotocodectypes

func (c_ CapturePhotoSettings) SetAvailableRawEmbeddedThumbnailPhotoCodecTypes(value VideoCodecType) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setAvailableRawEmbeddedThumbnailPhotoCodecTypes:"), value)
}


// A dictionary describing the format for delivery of thumbnail images embedded in photo file output.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturephotosettings/embeddedthumbnailphotoformat

func (c_ CapturePhotoSettings) EmbeddedThumbnailPhotoFormat() string {
	rv := objc.Send[string](c_.ID, objc.Sel("embeddedThumbnailPhotoFormat"))
	return rv
}


// A dictionary describing the format for delivery of thumbnail images embedded in photo file output.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturephotosettings/embeddedthumbnailphotoformat

func (c_ CapturePhotoSettings) SetEmbeddedThumbnailPhotoFormat(value string) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setEmbeddedThumbnailPhotoFormat:"), objc.String(value))
}


// A Boolean value that determines whether any depth data captured with the photo is included when generating output file data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturephotosettings/embedsdepthdatainphoto

func (c_ CapturePhotoSettings) EmbedsDepthDataInPhoto() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("embedsDepthDataInPhoto"))
	return rv
}


// A Boolean value that determines whether any depth data captured with the photo is included when generating output file data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturephotosettings/embedsdepthdatainphoto

func (c_ CapturePhotoSettings) SetEmbedsDepthDataInPhoto(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setEmbedsDepthDataInPhoto:"), value)
}


// Specifies whether the portrait effects matte captured with ths photo should be written to the photo’s file structure.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturephotosettings/embedsportraiteffectsmatteinphoto

func (c_ CapturePhotoSettings) EmbedsPortraitEffectsMatteInPhoto() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("embedsPortraitEffectsMatteInPhoto"))
	return rv
}


// Specifies whether the portrait effects matte captured with ths photo should be written to the photo’s file structure.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturephotosettings/embedsportraiteffectsmatteinphoto

func (c_ CapturePhotoSettings) SetEmbedsPortraitEffectsMatteInPhoto(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setEmbedsPortraitEffectsMatteInPhoto:"), value)
}


// A Boolean value that specifies whether to write the enabled semantic segmentation matte types captured with this photo to the photo’s file structure.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturephotosettings/embedssemanticsegmentationmattesinphoto

func (c_ CapturePhotoSettings) EmbedsSemanticSegmentationMattesInPhoto() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("embedsSemanticSegmentationMattesInPhoto"))
	return rv
}


// A Boolean value that specifies whether to write the enabled semantic segmentation matte types captured with this photo to the photo’s file structure.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturephotosettings/embedssemanticsegmentationmattesinphoto

func (c_ CapturePhotoSettings) SetEmbedsSemanticSegmentationMattesInPhoto(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setEmbedsSemanticSegmentationMattesInPhoto:"), value)
}


// An array of semantic segmentation matte types that the photo render pipeline can deliver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturephotosettings/enabledsemanticsegmentationmattetypes

func (c_ CapturePhotoSettings) EnabledSemanticSegmentationMatteTypes() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("enabledSemanticSegmentationMatteTypes"))
	return rv
}


// An array of semantic segmentation matte types that the photo render pipeline can deliver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturephotosettings/enabledsemanticsegmentationmattetypes

func (c_ CapturePhotoSettings) SetEnabledSemanticSegmentationMatteTypes(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setEnabledSemanticSegmentationMatteTypes:"), value)
}


// A setting for whether to fire the flash when capturing photos.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturephotosettings/flashmode

func (c_ CapturePhotoSettings) FlashMode() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("flashMode"))
	return rv
}


// A setting for whether to fire the flash when capturing photos.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturephotosettings/flashmode

func (c_ CapturePhotoSettings) SetFlashMode(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setFlashMode:"), value)
}


// A dictionary describing the processed format (for example, JPEG) to deliver captured photos in.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturephotosettings/format

func (c_ CapturePhotoSettings) Format() string {
	rv := objc.Send[string](c_.ID, objc.Sel("format"))
	return rv
}


// A dictionary describing the processed format (for example, JPEG) to deliver captured photos in.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturephotosettings/format

func (c_ CapturePhotoSettings) SetFormat(value string) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setFormat:"), objc.String(value))
}


// A Boolean value that specifies whether the photo output, at its discretion, uses content-aware distortion correction on this photo request.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturephotosettings/isautocontentawaredistortioncorrectionenabled

func (c_ CapturePhotoSettings) IsAutoContentAwareDistortionCorrectionEnabled() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("isAutoContentAwareDistortionCorrectionEnabled"))
	return rv
}


// A Boolean value that specifies whether the photo output, at its discretion, uses content-aware distortion correction on this photo request.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturephotosettings/isautocontentawaredistortioncorrectionenabled

func (c_ CapturePhotoSettings) SetIsAutoContentAwareDistortionCorrectionEnabled(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIsAutoContentAwareDistortionCorrectionEnabled:"), value)
}


// A Boolean value that specifies whether captures automatically combine data from a dual camera device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturephotosettings/isautodualcamerafusionenabled

func (c_ CapturePhotoSettings) IsAutoDualCameraFusionEnabled() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("isAutoDualCameraFusionEnabled"))
	return rv
}


// A Boolean value that specifies whether captures automatically combine data from a dual camera device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturephotosettings/isautodualcamerafusionenabled

func (c_ CapturePhotoSettings) SetIsAutoDualCameraFusionEnabled(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIsAutoDualCameraFusionEnabled:"), value)
}


// A Boolean value that indicates whether to use auto red-eye reduction on flash captures.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturephotosettings/isautoredeyereductionenabled

func (c_ CapturePhotoSettings) IsAutoRedEyeReductionEnabled() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("isAutoRedEyeReductionEnabled"))
	return rv
}


// A Boolean value that indicates whether to use auto red-eye reduction on flash captures.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturephotosettings/isautoredeyereductionenabled

func (c_ CapturePhotoSettings) SetIsAutoRedEyeReductionEnabled(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIsAutoRedEyeReductionEnabled:"), value)
}


// A Boolean value that specifies whether captures use automatic image stabilization.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturephotosettings/isautostillimagestabilizationenabled

func (c_ CapturePhotoSettings) IsAutoStillImageStabilizationEnabled() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("isAutoStillImageStabilizationEnabled"))
	return rv
}


// A Boolean value that specifies whether captures use automatic image stabilization.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturephotosettings/isautostillimagestabilizationenabled

func (c_ CapturePhotoSettings) SetIsAutoStillImageStabilizationEnabled(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIsAutoStillImageStabilizationEnabled:"), value)
}


// A Boolean value that specifies whether to use automatic virtual-device image fusion.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturephotosettings/isautovirtualdevicefusionenabled

func (c_ CapturePhotoSettings) IsAutoVirtualDeviceFusionEnabled() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("isAutoVirtualDeviceFusionEnabled"))
	return rv
}


// A Boolean value that specifies whether to use automatic virtual-device image fusion.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturephotosettings/isautovirtualdevicefusionenabled

func (c_ CapturePhotoSettings) SetIsAutoVirtualDeviceFusionEnabled(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIsAutoVirtualDeviceFusionEnabled:"), value)
}


// A Boolean value that determines whether a dual photo capture also delivers camera calibration data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturephotosettings/iscameracalibrationdatadeliveryenabled

func (c_ CapturePhotoSettings) IsCameraCalibrationDataDeliveryEnabled() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("isCameraCalibrationDataDeliveryEnabled"))
	return rv
}


// A Boolean value that determines whether a dual photo capture also delivers camera calibration data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturephotosettings/iscameracalibrationdatadeliveryenabled

func (c_ CapturePhotoSettings) SetIsCameraCalibrationDataDeliveryEnabled(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIsCameraCalibrationDataDeliveryEnabled:"), value)
}


// A Boolean value that indicates whether to capture the photo with constant color.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturephotosettings/isconstantcolorenabled

func (c_ CapturePhotoSettings) IsConstantColorEnabled() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("isConstantColorEnabled"))
	return rv
}


// A Boolean value that indicates whether to capture the photo with constant color.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturephotosettings/isconstantcolorenabled

func (c_ CapturePhotoSettings) SetIsConstantColorEnabled(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIsConstantColorEnabled:"), value)
}


// A Boolean value that indicates whether to deliver a fallback photo when taking a constant color capture.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturephotosettings/isconstantcolorfallbackphotodeliveryenabled

func (c_ CapturePhotoSettings) IsConstantColorFallbackPhotoDeliveryEnabled() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("isConstantColorFallbackPhotoDeliveryEnabled"))
	return rv
}


// A Boolean value that indicates whether to deliver a fallback photo when taking a constant color capture.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturephotosettings/isconstantcolorfallbackphotodeliveryenabled

func (c_ CapturePhotoSettings) SetIsConstantColorFallbackPhotoDeliveryEnabled(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIsConstantColorFallbackPhotoDeliveryEnabled:"), value)
}


// A Boolean value that determines whether the photo output captures depth data along with the photo.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturephotosettings/isdepthdatadeliveryenabled

func (c_ CapturePhotoSettings) IsDepthDataDeliveryEnabled() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("isDepthDataDeliveryEnabled"))
	return rv
}


// A Boolean value that determines whether the photo output captures depth data along with the photo.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturephotosettings/isdepthdatadeliveryenabled

func (c_ CapturePhotoSettings) SetIsDepthDataDeliveryEnabled(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIsDepthDataDeliveryEnabled:"), value)
}


// A Boolean value that determines whether to smooth noise and fill in missing values in depth data output.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturephotosettings/isdepthdatafiltered

func (c_ CapturePhotoSettings) IsDepthDataFiltered() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("isDepthDataFiltered"))
	return rv
}


// A Boolean value that determines whether to smooth noise and fill in missing values in depth data output.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturephotosettings/isdepthdatafiltered

func (c_ CapturePhotoSettings) SetIsDepthDataFiltered(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIsDepthDataFiltered:"), value)
}


// A Boolean value that determines whether a dual camera device delivers images from both cameras.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturephotosettings/isdualcameradualphotodeliveryenabled

func (c_ CapturePhotoSettings) IsDualCameraDualPhotoDeliveryEnabled() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("isDualCameraDualPhotoDeliveryEnabled"))
	return rv
}


// A Boolean value that determines whether a dual camera device delivers images from both cameras.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturephotosettings/isdualcameradualphotodeliveryenabled

func (c_ CapturePhotoSettings) SetIsDualCameraDualPhotoDeliveryEnabled(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIsDualCameraDualPhotoDeliveryEnabled:"), value)
}


// A Boolean value that specifies whether to capture still images at the highest resolution supported by the active device and format.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturephotosettings/ishighresolutionphotoenabled

func (c_ CapturePhotoSettings) IsHighResolutionPhotoEnabled() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("isHighResolutionPhotoEnabled"))
	return rv
}


// A Boolean value that specifies whether to capture still images at the highest resolution supported by the active device and format.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturephotosettings/ishighresolutionphotoenabled

func (c_ CapturePhotoSettings) SetIsHighResolutionPhotoEnabled(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIsHighResolutionPhotoEnabled:"), value)
}


// Specifies whether a portrait effects matte should be captured along with the photo.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturephotosettings/isportraiteffectsmattedeliveryenabled

func (c_ CapturePhotoSettings) IsPortraitEffectsMatteDeliveryEnabled() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("isPortraitEffectsMatteDeliveryEnabled"))
	return rv
}


// Specifies whether a portrait effects matte should be captured along with the photo.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturephotosettings/isportraiteffectsmattedeliveryenabled

func (c_ CapturePhotoSettings) SetIsPortraitEffectsMatteDeliveryEnabled(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIsPortraitEffectsMatteDeliveryEnabled:"), value)
}


// A Boolean value that indicates whether to suppress the built-in shutter sound when capturing a photo.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturephotosettings/isshuttersoundsuppressionenabled

func (c_ CapturePhotoSettings) IsShutterSoundSuppressionEnabled() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("isShutterSoundSuppressionEnabled"))
	return rv
}


// A Boolean value that indicates whether to suppress the built-in shutter sound when capturing a photo.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturephotosettings/isshuttersoundsuppressionenabled

func (c_ CapturePhotoSettings) SetIsShutterSoundSuppressionEnabled(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIsShutterSoundSuppressionEnabled:"), value)
}


// A URL at which to write Live Photo movie output.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturephotosettings/livephotomoviefileurl

func (c_ CapturePhotoSettings) LivePhotoMovieFileURL() foundation.URL {
	rv := objc.Send[foundation.URL](c_.ID, objc.Sel("livePhotoMovieFileURL"))
	return rv
}


// A URL at which to write Live Photo movie output.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturephotosettings/livephotomoviefileurl

func (c_ CapturePhotoSettings) SetLivePhotoMovieFileURL(value foundation.IURL) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setLivePhotoMovieFileURL:"), value)
}


// A dictionary of metadata to include in the Live Photo movie file.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturephotosettings/livephotomoviemetadata

func (c_ CapturePhotoSettings) LivePhotoMovieMetadata() AVMetadataItem {
	rv := objc.Send[AVMetadataItem](c_.ID, objc.Sel("livePhotoMovieMetadata"))
	return rv
}


// A dictionary of metadata to include in the Live Photo movie file.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturephotosettings/livephotomoviemetadata

func (c_ CapturePhotoSettings) SetLivePhotoMovieMetadata(value IAVMetadataItem) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setLivePhotoMovieMetadata:"), value)
}


// The video codec to use for encoding the movie portion of Live Photo output.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturephotosettings/livephotovideocodectype

func (c_ CapturePhotoSettings) LivePhotoVideoCodecType() VideoCodecType {
	rv := objc.Send[VideoCodecType](c_.ID, objc.Sel("livePhotoVideoCodecType"))
	return rv
}


// The video codec to use for encoding the movie portion of Live Photo output.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturephotosettings/livephotovideocodectype

func (c_ CapturePhotoSettings) SetLivePhotoVideoCodecType(value VideoCodecType) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setLivePhotoVideoCodecType:"), value)
}


// The maximum resolution of the photo to capture.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturephotosettings/maxphotodimensions

func (c_ CapturePhotoSettings) MaxPhotoDimensions() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("maxPhotoDimensions"))
	return rv
}


// The maximum resolution of the photo to capture.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturephotosettings/maxphotodimensions

func (c_ CapturePhotoSettings) SetMaxPhotoDimensions(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setMaxPhotoDimensions:"), value)
}


// A dictionary of metadata keys and values to embed in photo file output.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturephotosettings/metadata

func (c_ CapturePhotoSettings) Metadata() string {
	rv := objc.Send[string](c_.ID, objc.Sel("metadata"))
	return rv
}


// A dictionary of metadata keys and values to embed in photo file output.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturephotosettings/metadata

func (c_ CapturePhotoSettings) SetMetadata(value string) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setMetadata:"), objc.String(value))
}


// A setting that indicates how to prioritize photo quality against speed of photo delivery.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturephotosettings/photoqualityprioritization

func (c_ CapturePhotoSettings) PhotoQualityPrioritization() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("photoQualityPrioritization"))
	return rv
}


// A setting that indicates how to prioritize photo quality against speed of photo delivery.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturephotosettings/photoqualityprioritization

func (c_ CapturePhotoSettings) SetPhotoQualityPrioritization(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setPhotoQualityPrioritization:"), value)
}


// A dictionary describing the format for delivery of preview-sized images alongside the main photo.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturephotosettings/previewphotoformat

func (c_ CapturePhotoSettings) PreviewPhotoFormat() string {
	rv := objc.Send[string](c_.ID, objc.Sel("previewPhotoFormat"))
	return rv
}


// A dictionary describing the format for delivery of preview-sized images alongside the main photo.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturephotosettings/previewphotoformat

func (c_ CapturePhotoSettings) SetPreviewPhotoFormat(value string) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setPreviewPhotoFormat:"), objc.String(value))
}


// The container file format for eventual output of the processed image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturephotosettings/processedfiletype

func (c_ CapturePhotoSettings) ProcessedFileType() FileType {
	rv := objc.Send[FileType](c_.ID, objc.Sel("processedFileType"))
	return rv
}


// The container file format for eventual output of the processed image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturephotosettings/processedfiletype

func (c_ CapturePhotoSettings) SetProcessedFileType(value FileType) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setProcessedFileType:"), value)
}


// A dictionary describing the format for delivery of raw thumbnail images embedded in photo file output.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturephotosettings/rawembeddedthumbnailphotoformat

func (c_ CapturePhotoSettings) RawEmbeddedThumbnailPhotoFormat() string {
	rv := objc.Send[string](c_.ID, objc.Sel("rawEmbeddedThumbnailPhotoFormat"))
	return rv
}


// A dictionary describing the format for delivery of raw thumbnail images embedded in photo file output.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturephotosettings/rawembeddedthumbnailphotoformat

func (c_ CapturePhotoSettings) SetRawEmbeddedThumbnailPhotoFormat(value string) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setRawEmbeddedThumbnailPhotoFormat:"), objc.String(value))
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturephotosettings/rawfileformat

func (c_ CapturePhotoSettings) RawFileFormat() string {
	rv := objc.Send[string](c_.ID, objc.Sel("rawFileFormat"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturephotosettings/rawfileformat

func (c_ CapturePhotoSettings) SetRawFileFormat(value string) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setRawFileFormat:"), objc.String(value))
}


// The container file format for eventual output of the RAW image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturephotosettings/rawfiletype

func (c_ CapturePhotoSettings) RawFileType() FileType {
	rv := objc.Send[FileType](c_.ID, objc.Sel("rawFileType"))
	return rv
}


// The container file format for eventual output of the RAW image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturephotosettings/rawfiletype

func (c_ CapturePhotoSettings) SetRawFileType(value FileType) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setRawFileType:"), value)
}


// An identifier for the Bayer RAW pixel format to deliver captured RAW photos in.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturephotosettings/rawphotopixelformattype

func (c_ CapturePhotoSettings) RawPhotoPixelFormatType() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("rawPhotoPixelFormatType"))
	return rv
}


// An identifier for the Bayer RAW pixel format to deliver captured RAW photos in.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturephotosettings/rawphotopixelformattype

func (c_ CapturePhotoSettings) SetRawPhotoPixelFormatType(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setRawPhotoPixelFormatType:"), value)
}


// A unique identifier for this photo settings instance.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturephotosettings/uniqueid

func (c_ CapturePhotoSettings) UniqueID() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("uniqueID"))
	return rv
}


// A unique identifier for this photo settings instance.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturephotosettings/uniqueid

func (c_ CapturePhotoSettings) SetUniqueID(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setUniqueID:"), value)
}


// The constituent devices for which the virtual device should deliver photos.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturephotosettings/virtualdeviceconstituentphotodeliveryenableddevices

func (c_ CapturePhotoSettings) VirtualDeviceConstituentPhotoDeliveryEnabledDevices() AVCaptureDevice {
	rv := objc.Send[AVCaptureDevice](c_.ID, objc.Sel("virtualDeviceConstituentPhotoDeliveryEnabledDevices"))
	return rv
}


// The constituent devices for which the virtual device should deliver photos.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturephotosettings/virtualdeviceconstituentphotodeliveryenableddevices

func (c_ CapturePhotoSettings) SetVirtualDeviceConstituentPhotoDeliveryEnabledDevices(value IAVCaptureDevice) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setVirtualDeviceConstituentPhotoDeliveryEnabledDevices:"), value)
}



