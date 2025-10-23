// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
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
	AutomaticallyEnablesStillImageStabilizationWhenAvailable() bool /* primitive/slice/pointer. */
	SetAutomaticallyEnablesStillImageStabilizationWhenAvailable(value bool /* primitive/slice/pointer. */)
	AvailableImageDataCVPixelFormatTypes() objc.IObject /* cross-framework: NSNumber */
	SetAvailableImageDataCVPixelFormatTypes(value objc.IObject /* cross-framework: NSNumber */)
	AvailableImageDataCodecTypes() VideoCodecType /* not a class type */
	SetAvailableImageDataCodecTypes(value VideoCodecType /* not a class type */)
	IsCameraSensorOrientationCompensationEnabled() bool /* primitive/slice/pointer. */
	SetIsCameraSensorOrientationCompensationEnabled(value bool /* primitive/slice/pointer. */)
	IsCameraSensorOrientationCompensationSupported() bool /* primitive/slice/pointer. */
	SetIsCameraSensorOrientationCompensationSupported(value bool /* primitive/slice/pointer. */)
	IsCapturingStillImage() bool /* primitive/slice/pointer. */
	SetIsCapturingStillImage(value bool /* primitive/slice/pointer. */)
	IsHighResolutionStillImageOutputEnabled() bool /* primitive/slice/pointer. */
	SetIsHighResolutionStillImageOutputEnabled(value bool /* primitive/slice/pointer. */)
	IsLensStabilizationDuringBracketedCaptureEnabled() bool /* primitive/slice/pointer. */
	SetIsLensStabilizationDuringBracketedCaptureEnabled(value bool /* primitive/slice/pointer. */)
	IsLensStabilizationDuringBracketedCaptureSupported() bool /* primitive/slice/pointer. */
	SetIsLensStabilizationDuringBracketedCaptureSupported(value bool /* primitive/slice/pointer. */)
	IsStillImageStabilizationActive() bool /* primitive/slice/pointer. */
	SetIsStillImageStabilizationActive(value bool /* primitive/slice/pointer. */)
	IsStillImageStabilizationSupported() bool /* primitive/slice/pointer. */
	SetIsStillImageStabilizationSupported(value bool /* primitive/slice/pointer. */)
	MaxBracketedCaptureStillImageCount() int /* primitive/slice/pointer. */
	SetMaxBracketedCaptureStillImageCount(value int /* primitive/slice/pointer. */)
	OutputSettings() objc.IObject /* cross-framework: NSString */
	SetOutputSettings(value objc.IObject /* cross-framework: NSString */)
	// methods:
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

// Alloc allocates a new instance without initialization.
func (cc _CaptureStillImageOutputClass) Alloc() CaptureStillImageOutput {
	rv := objc.Send[CaptureStillImageOutput](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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



// A Boolean value that indicates whether still image stabilization should be automatically enabled.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturestillimageoutput/automaticallyenablesstillimagestabilizationwhenavailable
func (c_ CaptureStillImageOutput) AutomaticallyEnablesStillImageStabilizationWhenAvailable() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](c_.ID, objc.Sel("automaticallyEnablesStillImageStabilizationWhenAvailable"))
	return rv
}


// A Boolean value that indicates whether still image stabilization should be automatically enabled.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturestillimageoutput/automaticallyenablesstillimagestabilizationwhenavailable
func (c_ CaptureStillImageOutput) SetAutomaticallyEnablesStillImageStabilizationWhenAvailable(value bool /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setAutomaticallyEnablesStillImageStabilizationWhenAvailable:"), value)
}


// The supported image pixel formats that can be specified as output settings.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturestillimageoutput/availableimagedatacvpixelformattypes
func (c_ CaptureStillImageOutput) AvailableImageDataCVPixelFormatTypes() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](c_.ID, objc.Sel("availableImageDataCVPixelFormatTypes"))
	return rv
}


// The supported image pixel formats that can be specified as output settings.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturestillimageoutput/availableimagedatacvpixelformattypes
func (c_ CaptureStillImageOutput) SetAvailableImageDataCVPixelFormatTypes(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setAvailableImageDataCVPixelFormatTypes:"), value)
}


// The supported image codec formats that can be specified as output settings.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturestillimageoutput/availableimagedatacodectypes
func (c_ CaptureStillImageOutput) AvailableImageDataCodecTypes() VideoCodecType /* not a class type */ {
	rv := objc.Send[VideoCodecType](c_.ID, objc.Sel("availableImageDataCodecTypes"))
	return rv
}


// The supported image codec formats that can be specified as output settings.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturestillimageoutput/availableimagedatacodectypes
func (c_ CaptureStillImageOutput) SetAvailableImageDataCodecTypes(value VideoCodecType /* not a class type */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setAvailableImageDataCodecTypes:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturestillimageoutput/iscamerasensororientationcompensationenabled
func (c_ CaptureStillImageOutput) IsCameraSensorOrientationCompensationEnabled() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](c_.ID, objc.Sel("isCameraSensorOrientationCompensationEnabled"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturestillimageoutput/iscamerasensororientationcompensationenabled
func (c_ CaptureStillImageOutput) SetIsCameraSensorOrientationCompensationEnabled(value bool /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIsCameraSensorOrientationCompensationEnabled:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturestillimageoutput/iscamerasensororientationcompensationsupported
func (c_ CaptureStillImageOutput) IsCameraSensorOrientationCompensationSupported() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](c_.ID, objc.Sel("isCameraSensorOrientationCompensationSupported"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturestillimageoutput/iscamerasensororientationcompensationsupported
func (c_ CaptureStillImageOutput) SetIsCameraSensorOrientationCompensationSupported(value bool /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIsCameraSensorOrientationCompensationSupported:"), value)
}


// Indicates whether a still image is being captured.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturestillimageoutput/iscapturingstillimage
func (c_ CaptureStillImageOutput) IsCapturingStillImage() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](c_.ID, objc.Sel("isCapturingStillImage"))
	return rv
}


// Indicates whether a still image is being captured.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturestillimageoutput/iscapturingstillimage
func (c_ CaptureStillImageOutput) SetIsCapturingStillImage(value bool /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIsCapturingStillImage:"), value)
}


// A Boolean value that indicates whether the receiver should emit still images at the highest resolution supported by its source
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturestillimageoutput/ishighresolutionstillimageoutputenabled
func (c_ CaptureStillImageOutput) IsHighResolutionStillImageOutputEnabled() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](c_.ID, objc.Sel("isHighResolutionStillImageOutputEnabled"))
	return rv
}


// A Boolean value that indicates whether the receiver should emit still images at the highest resolution supported by its source
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturestillimageoutput/ishighresolutionstillimageoutputenabled
func (c_ CaptureStillImageOutput) SetIsHighResolutionStillImageOutputEnabled(value bool /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIsHighResolutionStillImageOutputEnabled:"), value)
}


// A Boolean value that specifies whether to stabilize the lens across the duration of a bracketed capture.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturestillimageoutput/islensstabilizationduringbracketedcaptureenabled
func (c_ CaptureStillImageOutput) IsLensStabilizationDuringBracketedCaptureEnabled() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](c_.ID, objc.Sel("isLensStabilizationDuringBracketedCaptureEnabled"))
	return rv
}


// A Boolean value that specifies whether to stabilize the lens across the duration of a bracketed capture.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturestillimageoutput/islensstabilizationduringbracketedcaptureenabled
func (c_ CaptureStillImageOutput) SetIsLensStabilizationDuringBracketedCaptureEnabled(value bool /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIsLensStabilizationDuringBracketedCaptureEnabled:"), value)
}


// A Boolean value that indicates whether the capture output supports lens stabilization across the duration of a bracketed capture.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturestillimageoutput/islensstabilizationduringbracketedcapturesupported
func (c_ CaptureStillImageOutput) IsLensStabilizationDuringBracketedCaptureSupported() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](c_.ID, objc.Sel("isLensStabilizationDuringBracketedCaptureSupported"))
	return rv
}


// A Boolean value that indicates whether the capture output supports lens stabilization across the duration of a bracketed capture.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturestillimageoutput/islensstabilizationduringbracketedcapturesupported
func (c_ CaptureStillImageOutput) SetIsLensStabilizationDuringBracketedCaptureSupported(value bool /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIsLensStabilizationDuringBracketedCaptureSupported:"), value)
}


// Indicates whether still image stabilization is in use for the current capture.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturestillimageoutput/isstillimagestabilizationactive
func (c_ CaptureStillImageOutput) IsStillImageStabilizationActive() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](c_.ID, objc.Sel("isStillImageStabilizationActive"))
	return rv
}


// Indicates whether still image stabilization is in use for the current capture.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturestillimageoutput/isstillimagestabilizationactive
func (c_ CaptureStillImageOutput) SetIsStillImageStabilizationActive(value bool /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIsStillImageStabilizationActive:"), value)
}


// A Boolean value that indicates whether the still image currently being captured supports still image stabilization.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturestillimageoutput/isstillimagestabilizationsupported
func (c_ CaptureStillImageOutput) IsStillImageStabilizationSupported() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](c_.ID, objc.Sel("isStillImageStabilizationSupported"))
	return rv
}


// A Boolean value that indicates whether the still image currently being captured supports still image stabilization.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturestillimageoutput/isstillimagestabilizationsupported
func (c_ CaptureStillImageOutput) SetIsStillImageStabilizationSupported(value bool /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIsStillImageStabilizationSupported:"), value)
}


// Specifies the maximum number of still images that may be taken in a single bracket.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturestillimageoutput/maxbracketedcapturestillimagecount
func (c_ CaptureStillImageOutput) MaxBracketedCaptureStillImageCount() int /* primitive/slice/pointer. */ {
	rv := objc.Send[int](c_.ID, objc.Sel("maxBracketedCaptureStillImageCount"))
	return rv
}


// Specifies the maximum number of still images that may be taken in a single bracket.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturestillimageoutput/maxbracketedcapturestillimagecount
func (c_ CaptureStillImageOutput) SetMaxBracketedCaptureStillImageCount(value int /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setMaxBracketedCaptureStillImageCount:"), value)
}


// The compression settings for the output.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturestillimageoutput/outputsettings
func (c_ CaptureStillImageOutput) OutputSettings() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("outputSettings"))
	return rv
}


// The compression settings for the output.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturestillimageoutput/outputsettings
func (c_ CaptureStillImageOutput) SetOutputSettings(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setOutputSettings:"), value)
}



