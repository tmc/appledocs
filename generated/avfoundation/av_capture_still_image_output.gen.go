// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)





// The class instance for the [CaptureStillImageOutput] class.
var (
	CaptureStillImageOutputClass     _CaptureStillImageOutputClass
	CaptureStillImageOutputClassOnce sync.Once
)

func getCaptureStillImageOutputClass() _CaptureStillImageOutputClass {
	CaptureStillImageOutputClassOnce.Do(func() {
		CaptureStillImageOutputClass = _CaptureStillImageOutputClass{objc.GetClass("AVCaptureStillImageOutput")}
	})
	return CaptureStillImageOutputClass
}

type _CaptureStillImageOutputClass struct {
	class objc.Class
}





// An interface definition for the [CaptureStillImageOutput] class.
type ICaptureStillImageOutput interface {
	ICaptureOutput
	

	// properties:
	AvailableImageDataCodecTypes() []string
	AvailableImageDataCVPixelFormatTypes() []foundation.Number
	CapturingStillImage() bool
	HighResolutionStillImageOutputEnabled() bool
	SetHighResolutionStillImageOutputEnabled(value bool)
	OutputSettings() foundation.IDictionary
	SetOutputSettings(value foundation.IDictionary)
	IsCameraSensorOrientationCompensationEnabled() bool
	SetIsCameraSensorOrientationCompensationEnabled(value bool)
	IsCameraSensorOrientationCompensationSupported() bool
	SetIsCameraSensorOrientationCompensationSupported(value bool)
	IsCapturingStillImage() bool
	SetIsCapturingStillImage(value bool)
	IsHighResolutionStillImageOutputEnabled() bool
	SetIsHighResolutionStillImageOutputEnabled(value bool)
	IsLensStabilizationDuringBracketedCaptureEnabled() bool
	SetIsLensStabilizationDuringBracketedCaptureEnabled(value bool)
	IsLensStabilizationDuringBracketedCaptureSupported() bool
	SetIsLensStabilizationDuringBracketedCaptureSupported(value bool)
	IsStillImageStabilizationActive() bool
	SetIsStillImageStabilizationActive(value bool)
	IsStillImageStabilizationSupported() bool
	SetIsStillImageStabilizationSupported(value bool)


	

	// methods:


}





// Alloc allocates a new instance without initialization.
func (cc _CaptureStillImageOutputClass) Alloc() CaptureStillImageOutput {
	rv := objc.Send[CaptureStillImageOutput](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _CaptureStillImageOutputClass) New() CaptureStillImageOutput {
	rv := objc.Send[CaptureStillImageOutput](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CaptureStillImageOutput) Init() CaptureStillImageOutput {
	rv := objc.Send[CaptureStillImageOutput](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CaptureStillImageOutput) Autorelease() CaptureStillImageOutput {
	rv := objc.Send[CaptureStillImageOutput](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCaptureStillImageOutput creates a new CaptureStillImageOutput instance.
func NewCaptureStillImageOutput() CaptureStillImageOutput {
	return getCaptureStillImageOutputClass().New()
}





// A capture output for capturing still photos.


// A capture output for capturing still photos.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureStillImageOutput
type CaptureStillImageOutput struct {
	CaptureOutput
}

// CaptureStillImageOutputFrom constructs a [CaptureStillImageOutput] from an unsafe.Pointer.
//
// A capture output for capturing still photos.
func CaptureStillImageOutputFrom(ptr unsafe.Pointer) CaptureStillImageOutput {
	return CaptureStillImageOutput{
		CaptureOutput: CaptureOutputFrom(ptr),
	}
}











// Returns an representation of a still image data and metadata attachments in a JPEG sample buffer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureStillImageOutput/jpegStillImageNSDataRepresentation(_:)
func (cc _CaptureStillImageOutputClass) JpegStillImageNSDataRepresentation(jpegSampleBuffer SampleBufferRef /* not a class type */) foundation.Data {
	rv := objc.Send[foundation.Data](objc.ID(cc.class), objc.Sel("jpegStillImageNSDataRepresentation:"), jpegSampleBuffer)
	return rv
}

















// The supported image codec formats that can be specified as output settings.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureStillImageOutput/availableImageDataCodecTypes
func (c_ CaptureStillImageOutput) AvailableImageDataCodecTypes() []string {
	rv := objc.Send[[]string](c_.ID, objc.Sel("availableImageDataCodecTypes"))
	return rv
}


// The supported image pixel formats that can be specified as output settings.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureStillImageOutput/availableImageDataCVPixelFormatTypes
func (c_ CaptureStillImageOutput) AvailableImageDataCVPixelFormatTypes() []foundation.Number {
	rv := objc.Send[[]foundation.Number](c_.ID, objc.Sel("availableImageDataCVPixelFormatTypes"))
	return rv
}


// Indicates whether a still image is being captured.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureStillImageOutput/isCapturingStillImage
func (c_ CaptureStillImageOutput) CapturingStillImage() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("capturingStillImage"))
	return rv
}


// A Boolean value that indicates whether the receiver should emit still images at the highest resolution supported by its source objects property.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureStillImageOutput/isHighResolutionStillImageOutputEnabled
func (c_ CaptureStillImageOutput) HighResolutionStillImageOutputEnabled() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("highResolutionStillImageOutputEnabled"))
	return rv
}


// A Boolean value that indicates whether the receiver should emit still images at the highest resolution supported by its source objects property.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureStillImageOutput/isHighResolutionStillImageOutputEnabled
func (c_ CaptureStillImageOutput) SetHighResolutionStillImageOutputEnabled(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setHighResolutionStillImageOutputEnabled:"), value)
}


// The compression settings for the output.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureStillImageOutput/outputSettings
func (c_ CaptureStillImageOutput) OutputSettings() foundation.IDictionary {
	rv := objc.Send[foundation.IDictionary](c_.ID, objc.Sel("outputSettings"))
	return rv
}


// The compression settings for the output.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureStillImageOutput/outputSettings
func (c_ CaptureStillImageOutput) SetOutputSettings(value foundation.IDictionary) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setOutputSettings:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturestillimageoutput/iscamerasensororientationcompensationenabled
func (c_ CaptureStillImageOutput) IsCameraSensorOrientationCompensationEnabled() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("isCameraSensorOrientationCompensationEnabled"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturestillimageoutput/iscamerasensororientationcompensationenabled
func (c_ CaptureStillImageOutput) SetIsCameraSensorOrientationCompensationEnabled(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIsCameraSensorOrientationCompensationEnabled:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturestillimageoutput/iscamerasensororientationcompensationsupported
func (c_ CaptureStillImageOutput) IsCameraSensorOrientationCompensationSupported() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("isCameraSensorOrientationCompensationSupported"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturestillimageoutput/iscamerasensororientationcompensationsupported
func (c_ CaptureStillImageOutput) SetIsCameraSensorOrientationCompensationSupported(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIsCameraSensorOrientationCompensationSupported:"), value)
}


// Indicates whether a still image is being captured.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturestillimageoutput/iscapturingstillimage
func (c_ CaptureStillImageOutput) IsCapturingStillImage() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("isCapturingStillImage"))
	return rv
}


// Indicates whether a still image is being captured.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturestillimageoutput/iscapturingstillimage
func (c_ CaptureStillImageOutput) SetIsCapturingStillImage(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIsCapturingStillImage:"), value)
}


// A Boolean value that indicates whether the receiver should emit still images at the highest resolution supported by its source
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturestillimageoutput/ishighresolutionstillimageoutputenabled
func (c_ CaptureStillImageOutput) IsHighResolutionStillImageOutputEnabled() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("isHighResolutionStillImageOutputEnabled"))
	return rv
}


// A Boolean value that indicates whether the receiver should emit still images at the highest resolution supported by its source
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturestillimageoutput/ishighresolutionstillimageoutputenabled
func (c_ CaptureStillImageOutput) SetIsHighResolutionStillImageOutputEnabled(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIsHighResolutionStillImageOutputEnabled:"), value)
}


// A Boolean value that specifies whether to stabilize the lens across the duration of a bracketed capture.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturestillimageoutput/islensstabilizationduringbracketedcaptureenabled
func (c_ CaptureStillImageOutput) IsLensStabilizationDuringBracketedCaptureEnabled() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("isLensStabilizationDuringBracketedCaptureEnabled"))
	return rv
}


// A Boolean value that specifies whether to stabilize the lens across the duration of a bracketed capture.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturestillimageoutput/islensstabilizationduringbracketedcaptureenabled
func (c_ CaptureStillImageOutput) SetIsLensStabilizationDuringBracketedCaptureEnabled(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIsLensStabilizationDuringBracketedCaptureEnabled:"), value)
}


// A Boolean value that indicates whether the capture output supports lens stabilization across the duration of a bracketed capture.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturestillimageoutput/islensstabilizationduringbracketedcapturesupported
func (c_ CaptureStillImageOutput) IsLensStabilizationDuringBracketedCaptureSupported() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("isLensStabilizationDuringBracketedCaptureSupported"))
	return rv
}


// A Boolean value that indicates whether the capture output supports lens stabilization across the duration of a bracketed capture.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturestillimageoutput/islensstabilizationduringbracketedcapturesupported
func (c_ CaptureStillImageOutput) SetIsLensStabilizationDuringBracketedCaptureSupported(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIsLensStabilizationDuringBracketedCaptureSupported:"), value)
}


// Indicates whether still image stabilization is in use for the current capture.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturestillimageoutput/isstillimagestabilizationactive
func (c_ CaptureStillImageOutput) IsStillImageStabilizationActive() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("isStillImageStabilizationActive"))
	return rv
}


// Indicates whether still image stabilization is in use for the current capture.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturestillimageoutput/isstillimagestabilizationactive
func (c_ CaptureStillImageOutput) SetIsStillImageStabilizationActive(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIsStillImageStabilizationActive:"), value)
}


// A Boolean value that indicates whether the still image currently being captured supports still image stabilization.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturestillimageoutput/isstillimagestabilizationsupported
func (c_ CaptureStillImageOutput) IsStillImageStabilizationSupported() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("isStillImageStabilizationSupported"))
	return rv
}


// A Boolean value that indicates whether the still image currently being captured supports still image stabilization.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturestillimageoutput/isstillimagestabilizationsupported
func (c_ CaptureStillImageOutput) SetIsStillImageStabilizationSupported(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIsStillImageStabilizationSupported:"), value)
}







