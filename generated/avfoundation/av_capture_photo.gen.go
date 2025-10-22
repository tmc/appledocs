// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/coregraphics"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [CapturePhoto] class.
var (
	CapturePhotoClass     _CapturePhotoClass
	CapturePhotoClassOnce sync.Once
)

func getCapturePhotoClass() _CapturePhotoClass {
	CapturePhotoClassOnce.Do(func() {
		CapturePhotoClass = _CapturePhotoClass{objc.GetClass("AVCapturePhoto")}
	})
	return CapturePhotoClass
}

type _CapturePhotoClass struct {
	class objc.Class
}

// An interface definition for the [CapturePhoto] class.
type ICapturePhoto interface {
	objectivec.IObject
	CGImageRepresentation() coregraphics.CGImageRef
	FileDataRepresentation() foundation.Data
	RawPhoto() bool
	PixelBuffer() unsafe.Pointer
	BracketSettings() unsafe.Pointer
	SetBracketSettings(value unsafe.Pointer)
	CameraCalibrationData() AVCameraCalibrationData
	SetCameraCalibrationData(value IAVCameraCalibrationData)
	ConstantColorCenterWeightedMeanConfidenceLevel() float32
	SetConstantColorCenterWeightedMeanConfidenceLevel(value float32)
	ConstantColorConfidenceMap() unsafe.Pointer
	SetConstantColorConfidenceMap(value unsafe.Pointer)
	DepthData() AVDepthData
	SetDepthData(value IAVDepthData)
	EmbeddedThumbnailPhotoFormat() string
	SetEmbeddedThumbnailPhotoFormat(value string)
	IsConstantColorFallbackPhoto() bool
	SetIsConstantColorFallbackPhoto(value bool)
	IsRawPhoto() bool
	SetIsRawPhoto(value bool)
	LensStabilizationStatus() unsafe.Pointer
	SetLensStabilizationStatus(value unsafe.Pointer)
	Metadata() string
	SetMetadata(value string)
	PhotoCount() int
	SetPhotoCount(value int)
	PortraitEffectsMatte() AVPortraitEffectsMatte
	SetPortraitEffectsMatte(value IAVPortraitEffectsMatte)
	PreviewPixelBuffer() unsafe.Pointer
	SetPreviewPixelBuffer(value unsafe.Pointer)
	ResolvedSettings() AVCaptureResolvedPhotoSettings
	SetResolvedSettings(value IAVCaptureResolvedPhotoSettings)
	SequenceCount() int
	SetSequenceCount(value int)
	SourceDeviceType() unsafe.Pointer
	SetSourceDeviceType(value unsafe.Pointer)
	Timestamp() unsafe.Pointer
	SetTimestamp(value unsafe.Pointer)
}

// A container for image data from a photo capture output.
//
// When you capture photos with the class, your delegate object receives each resulting image and related data in the form of an object. This object is an immutable wrapper from which you can retrieve various results of the photo capture. In addition to the photo image pixel buffer, an AVCapturePhoto object can also contain a preview-sized pixel buffer, capture metadata, and, on supported devices, depth data and camera calibration data. From an object, you can generate data appropriate for writing to a file, such as HEVC encoded image data containerized in the HEIC file format and including a preview image, depth data and other attachments. An instance wraps a single image result. For example, if you request a bracketed capture of three images, your callback is called three times, each time delivering a single object.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCapturePhoto
type CapturePhoto struct {
	objectivec.Object
}

// CapturePhotoFrom constructs a [CapturePhoto] from an unsafe.Pointer.
//
// A container for image data from a photo capture output.
func CapturePhotoFrom(ptr unsafe.Pointer) CapturePhoto {
	return CapturePhoto{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (cc _CapturePhotoClass) Alloc() CapturePhoto {
	rv := objc.Send[CapturePhoto](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (cc _CapturePhotoClass) New() CapturePhoto {
	rv := objc.Send[CapturePhoto](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CapturePhoto) Init() CapturePhoto {
	rv := objc.Send[CapturePhoto](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CapturePhoto) Autorelease() CapturePhoto {
	rv := objc.Send[CapturePhoto](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCapturePhoto creates a new CapturePhoto instance.
func NewCapturePhoto() CapturePhoto {
	return getCapturePhotoClass().New()
}


// Extracts and returns the captured photo’s primary image as a Core Graphics image object.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCapturePhoto/cgImageRepresentation()
func (c_ CapturePhoto) CGImageRepresentation() coregraphics.CGImageRef {
	rv := objc.Send[coregraphics.CGImageRef](c_.ID, objc.Sel("CGImageRepresentation"))
	return rv
}

// Generates and returns a flat data representation of the photo and its attachments.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCapturePhoto/fileDataRepresentation()
func (c_ CapturePhoto) FileDataRepresentation() foundation.Data {
	rv := objc.Send[foundation.Data](c_.ID, objc.Sel("fileDataRepresentation"))
	return rv
}

// A Boolean value indicating whether this photo object contains RAW format data.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCapturePhoto/isRawPhoto
func (c_ CapturePhoto) RawPhoto() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("rawPhoto"))
	return rv
}

// The uncompressed or RAW image sample buffer for the photo, if requested.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCapturePhoto/pixelBuffer
func (c_ CapturePhoto) PixelBuffer() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("pixelBuffer"))
	return rv
}

// The variations available for bracketed capture settings for this photo.
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturephoto/bracketsettings
func (c_ CapturePhoto) BracketSettings() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("bracketSettings"))
	return rv
}


// SetBracketSettings sets the value of the bracketSettings property.
// The variations available for bracketed capture settings for this photo.

//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturephoto/bracketsettings
func (c_ CapturePhoto) SetBracketSettings(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setBracketSettings:"), value)
}

// Calibration information for the camera device that captured the photo.
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturephoto/cameracalibrationdata
func (c_ CapturePhoto) CameraCalibrationData() AVCameraCalibrationData {
	rv := objc.Send[AVCameraCalibrationData](c_.ID, objc.Sel("cameraCalibrationData"))
	return rv
}


// SetCameraCalibrationData sets the value of the cameraCalibrationData property.
// Calibration information for the camera device that captured the photo.

//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturephoto/cameracalibrationdata
func (c_ CapturePhoto) SetCameraCalibrationData(value IAVCameraCalibrationData) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setCameraCalibrationData:"), value)
}

// A score that summarizes the overall confidence level of a constant color photo.
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturephoto/constantcolorcenterweightedmeanconfidencelevel
func (c_ CapturePhoto) ConstantColorCenterWeightedMeanConfidenceLevel() float32 {
	rv := objc.Send[float32](c_.ID, objc.Sel("constantColorCenterWeightedMeanConfidenceLevel"))
	return rv
}


// SetConstantColorCenterWeightedMeanConfidenceLevel sets the value of the constantColorCenterWeightedMeanConfidenceLevel property.
// A score that summarizes the overall confidence level of a constant color photo.

//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturephoto/constantcolorcenterweightedmeanconfidencelevel
func (c_ CapturePhoto) SetConstantColorCenterWeightedMeanConfidenceLevel(value float32) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setConstantColorCenterWeightedMeanConfidenceLevel:"), value)
}

// A pixel buffer where each pixel value indicates how fully the system achieves the constant color effect in the corresponding region of the photo.
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturephoto/constantcolorconfidencemap
func (c_ CapturePhoto) ConstantColorConfidenceMap() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("constantColorConfidenceMap"))
	return rv
}


// SetConstantColorConfidenceMap sets the value of the constantColorConfidenceMap property.
// A pixel buffer where each pixel value indicates how fully the system achieves the constant color effect in the corresponding region of the photo.

//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturephoto/constantcolorconfidencemap
func (c_ CapturePhoto) SetConstantColorConfidenceMap(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setConstantColorConfidenceMap:"), value)
}

// Depth or disparity map data captured with the photo.
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturephoto/depthdata
func (c_ CapturePhoto) DepthData() AVDepthData {
	rv := objc.Send[AVDepthData](c_.ID, objc.Sel("depthData"))
	return rv
}


// SetDepthData sets the value of the depthData property.
// Depth or disparity map data captured with the photo.

//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturephoto/depthdata
func (c_ CapturePhoto) SetDepthData(value IAVDepthData) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setDepthData:"), value)
}

// A dictionary describing the data format for a preview-sized image accompanying the captured photo.
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturephoto/embeddedthumbnailphotoformat
func (c_ CapturePhoto) EmbeddedThumbnailPhotoFormat() string {
	rv := objc.Send[string](c_.ID, objc.Sel("embeddedThumbnailPhotoFormat"))
	return rv
}


// SetEmbeddedThumbnailPhotoFormat sets the value of the embeddedThumbnailPhotoFormat property.
// A dictionary describing the data format for a preview-sized image accompanying the captured photo.

//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturephoto/embeddedthumbnailphotoformat
func (c_ CapturePhoto) SetEmbeddedThumbnailPhotoFormat(value string) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setEmbeddedThumbnailPhotoFormat:"), objc.String(value))
}

// A Boolean value that Indicates whether this photo is a fallback photo for a constant color capture.
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturephoto/isconstantcolorfallbackphoto
func (c_ CapturePhoto) IsConstantColorFallbackPhoto() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("isConstantColorFallbackPhoto"))
	return rv
}


// SetIsConstantColorFallbackPhoto sets the value of the isConstantColorFallbackPhoto property.
// A Boolean value that Indicates whether this photo is a fallback photo for a constant color capture.

//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturephoto/isconstantcolorfallbackphoto
func (c_ CapturePhoto) SetIsConstantColorFallbackPhoto(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIsConstantColorFallbackPhoto:"), value)
}

// A Boolean value indicating whether this photo object contains RAW format data.
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturephoto/israwphoto
func (c_ CapturePhoto) IsRawPhoto() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("isRawPhoto"))
	return rv
}


// SetIsRawPhoto sets the value of the isRawPhoto property.
// A Boolean value indicating whether this photo object contains RAW format data.

//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturephoto/israwphoto
func (c_ CapturePhoto) SetIsRawPhoto(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIsRawPhoto:"), value)
}

// Information about the use of lens stabilization during bracketed photo capture.
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturephoto/lensstabilizationstatus
func (c_ CapturePhoto) LensStabilizationStatus() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("lensStabilizationStatus"))
	return rv
}


// SetLensStabilizationStatus sets the value of the lensStabilizationStatus property.
// Information about the use of lens stabilization during bracketed photo capture.

//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturephoto/lensstabilizationstatus
func (c_ CapturePhoto) SetLensStabilizationStatus(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setLensStabilizationStatus:"), value)
}

// A dictionary of metadata describing the captured image.
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturephoto/metadata
func (c_ CapturePhoto) Metadata() string {
	rv := objc.Send[string](c_.ID, objc.Sel("metadata"))
	return rv
}


// SetMetadata sets the value of the metadata property.
// A dictionary of metadata describing the captured image.

//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturephoto/metadata
func (c_ CapturePhoto) SetMetadata(value string) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setMetadata:"), objc.String(value))
}

// The 1-based index of this photo capture relative to other results from the same capture request.
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturephoto/photocount
func (c_ CapturePhoto) PhotoCount() int {
	rv := objc.Send[int](c_.ID, objc.Sel("photoCount"))
	return rv
}


// SetPhotoCount sets the value of the photoCount property.
// The 1-based index of this photo capture relative to other results from the same capture request.

//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturephoto/photocount
func (c_ CapturePhoto) SetPhotoCount(value int) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setPhotoCount:"), value)
}

// The portrait effects matte captured with the photo.
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturephoto/portraiteffectsmatte
func (c_ CapturePhoto) PortraitEffectsMatte() AVPortraitEffectsMatte {
	rv := objc.Send[AVPortraitEffectsMatte](c_.ID, objc.Sel("portraitEffectsMatte"))
	return rv
}


// SetPortraitEffectsMatte sets the value of the portraitEffectsMatte property.
// The portrait effects matte captured with the photo.

//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturephoto/portraiteffectsmatte
func (c_ CapturePhoto) SetPortraitEffectsMatte(value IAVPortraitEffectsMatte) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setPortraitEffectsMatte:"), value)
}

// The pixel data for a preview-sized version of the photo, if requested.
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturephoto/previewpixelbuffer
func (c_ CapturePhoto) PreviewPixelBuffer() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("previewPixelBuffer"))
	return rv
}


// SetPreviewPixelBuffer sets the value of the previewPixelBuffer property.
// The pixel data for a preview-sized version of the photo, if requested.

//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturephoto/previewpixelbuffer
func (c_ CapturePhoto) SetPreviewPixelBuffer(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setPreviewPixelBuffer:"), value)
}

// The settings object that was used to request this photo capture.
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturephoto/resolvedsettings
func (c_ CapturePhoto) ResolvedSettings() AVCaptureResolvedPhotoSettings {
	rv := objc.Send[AVCaptureResolvedPhotoSettings](c_.ID, objc.Sel("resolvedSettings"))
	return rv
}


// SetResolvedSettings sets the value of the resolvedSettings property.
// The settings object that was used to request this photo capture.

//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturephoto/resolvedsettings
func (c_ CapturePhoto) SetResolvedSettings(value IAVCaptureResolvedPhotoSettings) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setResolvedSettings:"), value)
}

// The 1-based index of this photo in a bracketed capture sequence.
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturephoto/sequencecount
func (c_ CapturePhoto) SequenceCount() int {
	rv := objc.Send[int](c_.ID, objc.Sel("sequenceCount"))
	return rv
}


// SetSequenceCount sets the value of the sequenceCount property.
// The 1-based index of this photo in a bracketed capture sequence.

//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturephoto/sequencecount
func (c_ CapturePhoto) SetSequenceCount(value int) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setSequenceCount:"), value)
}

// The type of device that captured the photo.
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturephoto/sourcedevicetype
func (c_ CapturePhoto) SourceDeviceType() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("sourceDeviceType"))
	return rv
}


// SetSourceDeviceType sets the value of the sourceDeviceType property.
// The type of device that captured the photo.

//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturephoto/sourcedevicetype
func (c_ CapturePhoto) SetSourceDeviceType(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setSourceDeviceType:"), value)
}

// The time at which the image was captured.
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturephoto/timestamp
func (c_ CapturePhoto) Timestamp() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("timestamp"))
	return rv
}


// SetTimestamp sets the value of the timestamp property.
// The time at which the image was captured.

//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturephoto/timestamp
func (c_ CapturePhoto) SetTimestamp(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setTimestamp:"), value)
}



