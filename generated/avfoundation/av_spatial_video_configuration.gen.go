// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/corefoundation"
	"github.com/tmc/appledocs/generated/objectivec"
)





// The class instance for the [SpatialVideoConfiguration] class.
var (
	SpatialVideoConfigurationClass     _SpatialVideoConfigurationClass
	SpatialVideoConfigurationClassOnce sync.Once
)

func getSpatialVideoConfigurationClass() _SpatialVideoConfigurationClass {
	SpatialVideoConfigurationClassOnce.Do(func() {
		SpatialVideoConfigurationClass = _SpatialVideoConfigurationClass{objc.GetClass("AVSpatialVideoConfiguration")}
	})
	return SpatialVideoConfigurationClass
}

type _SpatialVideoConfigurationClass struct {
	class objc.Class
}





// An interface definition for the [SpatialVideoConfiguration] class.
type ISpatialVideoConfiguration interface {
	objectivec.IObject
	

	// properties:
	CameraCalibrationDataLensCollection() foundation.IDictionary
	SetCameraCalibrationDataLensCollection(value foundation.IDictionary)
	CameraSystemBaseline() foundation.foundation.INSNumber
	SetCameraSystemBaseline(value foundation.foundation.INSNumber)
	DisparityAdjustment() foundation.foundation.INSNumber
	SetDisparityAdjustment(value foundation.foundation.INSNumber)
	HorizontalFieldOfView() foundation.foundation.INSNumber
	SetHorizontalFieldOfView(value foundation.foundation.INSNumber)
	AnimationTool() IAVVideoCompositionCoreAnimationTool
	SetAnimationTool(value IAVVideoCompositionCoreAnimationTool)
	ColorPrimaries() foundation.foundation.INSString
	SetColorPrimaries(value foundation.foundation.INSString)
	ColorTransferFunction() foundation.foundation.INSString
	SetColorTransferFunction(value foundation.foundation.INSString)
	ColorYCbCrMatrix() foundation.foundation.INSString
	SetColorYCbCrMatrix(value foundation.foundation.INSString)
	CustomVideoCompositorClass() VideoCompositing /* not a class type */
	SetCustomVideoCompositorClass(value VideoCompositing /* not a class type */)
	FrameDuration() objectivec.IObject
	SetFrameDuration(value objectivec.IObject)
	RenderScale() float32
	SetRenderScale(value float32)
	RenderSize() corefoundation.CGSize
	SetRenderSize(value corefoundation.CGSize)


	

	// methods:


}





// Alloc allocates a new instance without initialization.
func (sc _SpatialVideoConfigurationClass) Alloc() SpatialVideoConfiguration {
	rv := objc.Send[SpatialVideoConfiguration](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (sc _SpatialVideoConfigurationClass) New() SpatialVideoConfiguration {
	rv := objc.Send[SpatialVideoConfiguration](objc.ID(sc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (s_ SpatialVideoConfiguration) Init() SpatialVideoConfiguration {
	rv := objc.Send[SpatialVideoConfiguration](s_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s_ SpatialVideoConfiguration) Autorelease() SpatialVideoConfiguration {
	rv := objc.Send[SpatialVideoConfiguration](s_.ID, objc.Sel("autorelease"))
	return rv
}

// NewSpatialVideoConfiguration creates a new SpatialVideoConfiguration instance.
func NewSpatialVideoConfiguration() SpatialVideoConfiguration {
	return getSpatialVideoConfigurationClass().New()
}





// An AVSpatialVideoConfiguration specifies spatial video properties.


// An AVSpatialVideoConfiguration specifies spatial video properties.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVSpatialVideoConfiguration-c.class
type SpatialVideoConfiguration struct {
	objectivec.Object
}

// SpatialVideoConfigurationFrom constructs a [SpatialVideoConfiguration] from an unsafe.Pointer.
//
// An AVSpatialVideoConfiguration specifies spatial video properties.
func SpatialVideoConfigurationFrom(ptr unsafe.Pointer) SpatialVideoConfiguration {
	return SpatialVideoConfiguration{objectivec.Object{objc.ID(ptr)}}
}






// Initializes an AVSpatialVideoConfiguration with a format description.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVSpatialVideoConfiguration-c.class/initWithFormatDescription:
func NewSpatialVideoConfigurationWithFormatDescription(formatDescription FormatDescriptionRef /* not a class type */) SpatialVideoConfiguration {
	instance := getSpatialVideoConfigurationClass().Alloc()
	rv := objc.Send[SpatialVideoConfiguration](instance.ID, objc.Sel("initWithFormatDescription:"), formatDescription)
	rv.Autorelease()
	return rv
}






















// Specifies intrinsic and extrinsic parameters for single or multiple lenses.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVSpatialVideoConfiguration-c.class/cameraCalibrationDataLensCollection
func (s_ SpatialVideoConfiguration) CameraCalibrationDataLensCollection() foundation.IDictionary {
	rv := objc.Send[foundation.IDictionary](s_.ID, objc.Sel("cameraCalibrationDataLensCollection"))
	return rv
}


// Specifies intrinsic and extrinsic parameters for single or multiple lenses.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVSpatialVideoConfiguration-c.class/cameraCalibrationDataLensCollection
func (s_ SpatialVideoConfiguration) SetCameraCalibrationDataLensCollection(value foundation.IDictionary) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setCameraCalibrationDataLensCollection:"), value)
}


// Specifies the distance between centers of the lenses of the camera system that created the video.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVSpatialVideoConfiguration-c.class/cameraSystemBaseline
func (s_ SpatialVideoConfiguration) CameraSystemBaseline() foundation.foundation.INSNumber {
	rv := objc.Send[foundation.NSNumber](s_.ID, objc.Sel("cameraSystemBaseline"))
	return rv
}


// Specifies the distance between centers of the lenses of the camera system that created the video.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVSpatialVideoConfiguration-c.class/cameraSystemBaseline
func (s_ SpatialVideoConfiguration) SetCameraSystemBaseline(value foundation.foundation.INSNumber) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setCameraSystemBaseline:"), value)
}


// Specifies a relative shift of the left and right images, which changes the zero parallax plane.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVSpatialVideoConfiguration-c.class/disparityAdjustment
func (s_ SpatialVideoConfiguration) DisparityAdjustment() foundation.foundation.INSNumber {
	rv := objc.Send[foundation.NSNumber](s_.ID, objc.Sel("disparityAdjustment"))
	return rv
}


// Specifies a relative shift of the left and right images, which changes the zero parallax plane.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVSpatialVideoConfiguration-c.class/disparityAdjustment
func (s_ SpatialVideoConfiguration) SetDisparityAdjustment(value foundation.foundation.INSNumber) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setDisparityAdjustment:"), value)
}


// Specifies horizontal field of view in thousandths of a degree. Can be nil if the value is unknown.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVSpatialVideoConfiguration-c.class/horizontalFieldOfView
func (s_ SpatialVideoConfiguration) HorizontalFieldOfView() foundation.foundation.INSNumber {
	rv := objc.Send[foundation.NSNumber](s_.ID, objc.Sel("horizontalFieldOfView"))
	return rv
}


// Specifies horizontal field of view in thousandths of a degree. Can be nil if the value is unknown.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVSpatialVideoConfiguration-c.class/horizontalFieldOfView
func (s_ SpatialVideoConfiguration) SetHorizontalFieldOfView(value foundation.foundation.INSNumber) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setHorizontalFieldOfView:"), value)
}


// A video composition tool to use with Core Animation in offline rendering.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avvideocomposition/animationtool
func (s_ SpatialVideoConfiguration) AnimationTool() IAVVideoCompositionCoreAnimationTool {
	rv := objc.Send[VideoCompositionCoreAnimationTool](s_.ID, objc.Sel("animationTool"))
	return rv
}


// A video composition tool to use with Core Animation in offline rendering.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avvideocomposition/animationtool
func (s_ SpatialVideoConfiguration) SetAnimationTool(value IAVVideoCompositionCoreAnimationTool) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setAnimationTool:"), value)
}


// The color primaries used for video composition.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avvideocomposition/colorprimaries
func (s_ SpatialVideoConfiguration) ColorPrimaries() foundation.foundation.INSString {
	rv := objc.Send[foundation.NSString](s_.ID, objc.Sel("colorPrimaries"))
	return rv
}


// The color primaries used for video composition.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avvideocomposition/colorprimaries
func (s_ SpatialVideoConfiguration) SetColorPrimaries(value foundation.foundation.INSString) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setColorPrimaries:"), value)
}


// The transfer function used for video composition.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avvideocomposition/colortransferfunction
func (s_ SpatialVideoConfiguration) ColorTransferFunction() foundation.foundation.INSString {
	rv := objc.Send[foundation.NSString](s_.ID, objc.Sel("colorTransferFunction"))
	return rv
}


// The transfer function used for video composition.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avvideocomposition/colortransferfunction
func (s_ SpatialVideoConfiguration) SetColorTransferFunction(value foundation.foundation.INSString) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setColorTransferFunction:"), value)
}


// The YCbCr matrix used for video composition.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avvideocomposition/colorycbcrmatrix
func (s_ SpatialVideoConfiguration) ColorYCbCrMatrix() foundation.foundation.INSString {
	rv := objc.Send[foundation.NSString](s_.ID, objc.Sel("colorYCbCrMatrix"))
	return rv
}


// The YCbCr matrix used for video composition.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avvideocomposition/colorycbcrmatrix
func (s_ SpatialVideoConfiguration) SetColorYCbCrMatrix(value foundation.foundation.INSString) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setColorYCbCrMatrix:"), value)
}


// A custom compositor class to use.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avvideocomposition/customvideocompositorclass
func (s_ SpatialVideoConfiguration) CustomVideoCompositorClass() VideoCompositing /* not a class type */ {
	rv := objc.Send[VideoCompositing](s_.ID, objc.Sel("customVideoCompositorClass"))
	return rv
}


// A custom compositor class to use.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avvideocomposition/customvideocompositorclass
func (s_ SpatialVideoConfiguration) SetCustomVideoCompositorClass(value VideoCompositing /* not a class type */) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setCustomVideoCompositorClass:"), value)
}


// A time interval for which the video composition should render composed video frames.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avvideocomposition/frameduration
func (s_ SpatialVideoConfiguration) FrameDuration() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](s_.ID, objc.Sel("frameDuration"))
	return rv
}


// A time interval for which the video composition should render composed video frames.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avvideocomposition/frameduration
func (s_ SpatialVideoConfiguration) SetFrameDuration(value objectivec.IObject) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setFrameDuration:"), value)
}


// The scale at which the video composition should render.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avvideocomposition/renderscale
func (s_ SpatialVideoConfiguration) RenderScale() float32 {
	rv := objc.Send[float32](s_.ID, objc.Sel("renderScale"))
	return rv
}


// The scale at which the video composition should render.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avvideocomposition/renderscale
func (s_ SpatialVideoConfiguration) SetRenderScale(value float32) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setRenderScale:"), value)
}


// The size at which the video composition should render.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avvideocomposition/rendersize
func (s_ SpatialVideoConfiguration) RenderSize() corefoundation.CGSize {
	rv := objc.Send[corefoundation.CGSize](s_.ID, objc.Sel("renderSize"))
	return rv
}


// The size at which the video composition should render.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avvideocomposition/rendersize
func (s_ SpatialVideoConfiguration) SetRenderSize(value corefoundation.CGSize) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setRenderSize:"), value)
}







