// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class AVSpatialVideoConfiguration */


/* debug [class_header]: Header for AVSpatialVideoConfiguration */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for SpatialVideoConfiguration */
// An interface definition for the [SpatialVideoConfiguration] class.
type ISpatialVideoConfiguration interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for SpatialVideoConfiguration */
	// properties:
	CameraCalibrationDataLensCollection() foundation.IDictionary
	SetCameraCalibrationDataLensCollection(value foundation.IDictionary)
	CameraSystemBaseline() objc.IObject /* cross-framework: NSNumber */
	SetCameraSystemBaseline(value objc.IObject /* cross-framework: NSNumber */)
	DisparityAdjustment() objc.IObject /* cross-framework: NSNumber */
	SetDisparityAdjustment(value objc.IObject /* cross-framework: NSNumber */)
	HorizontalFieldOfView() objc.IObject /* cross-framework: NSNumber */
	SetHorizontalFieldOfView(value objc.IObject /* cross-framework: NSNumber */)
	AnimationTool() IAVVideoCompositionCoreAnimationTool
	SetAnimationTool(value IAVVideoCompositionCoreAnimationTool)
	ColorPrimaries() objc.IObject /* cross-framework: NSString */
	SetColorPrimaries(value objc.IObject /* cross-framework: NSString */)
	ColorTransferFunction() objc.IObject /* cross-framework: NSString */
	SetColorTransferFunction(value objc.IObject /* cross-framework: NSString */)
	ColorYCbCrMatrix() objc.IObject /* cross-framework: NSString */
	SetColorYCbCrMatrix(value objc.IObject /* cross-framework: NSString */)
	CustomVideoCompositorClass() VideoCompositing /* not a class type */
	SetCustomVideoCompositorClass(value VideoCompositing /* not a class type */)
	FrameDuration() objc.IObject /* cross-framework: Time */
	SetFrameDuration(value objc.IObject /* cross-framework: Time */)
	RenderScale() float32
	SetRenderScale(value float32)
	RenderSize() corefoundation.CGSize
	SetRenderSize(value corefoundation.CGSize)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for SpatialVideoConfiguration */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for SpatialVideoConfiguration */
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for SpatialVideoConfiguration */
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for SpatialVideoConfiguration */

// Initializes an AVSpatialVideoConfiguration with a format description.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVSpatialVideoConfiguration-c.class/initWithFormatDescription:
func NewSpatialVideoConfigurationWithFormatDescription(formatDescription FormatDescriptionRef /* not a class type */) SpatialVideoConfiguration {
	instance := getSpatialVideoConfigurationClass().Alloc()
	rv := objc.Send[SpatialVideoConfiguration](instance.ID, objc.Sel("initWithFormatDescription:"), formatDescription)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewSpatialVideoConfigurationWithFormatDescription */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for SpatialVideoConfiguration */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for SpatialVideoConfiguration */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for SpatialVideoConfiguration */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for SpatialVideoConfiguration */

// Specifies intrinsic and extrinsic parameters for single or multiple lenses.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVSpatialVideoConfiguration-c.class/cameraCalibrationDataLensCollection
func (s_ SpatialVideoConfiguration) CameraCalibrationDataLensCollection() foundation.IDictionary {
	rv := objc.Send[foundation.IDictionary](s_.ID, objc.Sel("cameraCalibrationDataLensCollection"))
	return rv
}/* debug [instance_properties/getter]: cameraCalibrationDataLensCollection */


// Specifies intrinsic and extrinsic parameters for single or multiple lenses.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVSpatialVideoConfiguration-c.class/cameraCalibrationDataLensCollection
func (s_ SpatialVideoConfiguration) SetCameraCalibrationDataLensCollection(value foundation.IDictionary) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setCameraCalibrationDataLensCollection:"), value)
}/* debug [instance_properties/setter]: cameraCalibrationDataLensCollection */


// Specifies the distance between centers of the lenses of the camera system that created the video.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVSpatialVideoConfiguration-c.class/cameraSystemBaseline
func (s_ SpatialVideoConfiguration) CameraSystemBaseline() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](s_.ID, objc.Sel("cameraSystemBaseline"))
	return rv
}/* debug [instance_properties/getter]: cameraSystemBaseline */


// Specifies the distance between centers of the lenses of the camera system that created the video.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVSpatialVideoConfiguration-c.class/cameraSystemBaseline
func (s_ SpatialVideoConfiguration) SetCameraSystemBaseline(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setCameraSystemBaseline:"), value)
}/* debug [instance_properties/setter]: cameraSystemBaseline */


// Specifies a relative shift of the left and right images, which changes the zero parallax plane.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVSpatialVideoConfiguration-c.class/disparityAdjustment
func (s_ SpatialVideoConfiguration) DisparityAdjustment() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](s_.ID, objc.Sel("disparityAdjustment"))
	return rv
}/* debug [instance_properties/getter]: disparityAdjustment */


// Specifies a relative shift of the left and right images, which changes the zero parallax plane.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVSpatialVideoConfiguration-c.class/disparityAdjustment
func (s_ SpatialVideoConfiguration) SetDisparityAdjustment(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setDisparityAdjustment:"), value)
}/* debug [instance_properties/setter]: disparityAdjustment */


// Specifies horizontal field of view in thousandths of a degree. Can be nil if the value is unknown.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVSpatialVideoConfiguration-c.class/horizontalFieldOfView
func (s_ SpatialVideoConfiguration) HorizontalFieldOfView() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](s_.ID, objc.Sel("horizontalFieldOfView"))
	return rv
}/* debug [instance_properties/getter]: horizontalFieldOfView */


// Specifies horizontal field of view in thousandths of a degree. Can be nil if the value is unknown.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVSpatialVideoConfiguration-c.class/horizontalFieldOfView
func (s_ SpatialVideoConfiguration) SetHorizontalFieldOfView(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setHorizontalFieldOfView:"), value)
}/* debug [instance_properties/setter]: horizontalFieldOfView */


// A video composition tool to use with Core Animation in offline rendering.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avvideocomposition/animationtool
func (s_ SpatialVideoConfiguration) AnimationTool() IAVVideoCompositionCoreAnimationTool {
	rv := objc.Send[VideoCompositionCoreAnimationTool](s_.ID, objc.Sel("animationTool"))
	return rv
}/* debug [instance_properties/getter]: animationTool */


// A video composition tool to use with Core Animation in offline rendering.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avvideocomposition/animationtool
func (s_ SpatialVideoConfiguration) SetAnimationTool(value IAVVideoCompositionCoreAnimationTool) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setAnimationTool:"), value)
}/* debug [instance_properties/setter]: animationTool */


// The color primaries used for video composition.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avvideocomposition/colorprimaries
func (s_ SpatialVideoConfiguration) ColorPrimaries() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](s_.ID, objc.Sel("colorPrimaries"))
	return rv
}/* debug [instance_properties/getter]: colorPrimaries */


// The color primaries used for video composition.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avvideocomposition/colorprimaries
func (s_ SpatialVideoConfiguration) SetColorPrimaries(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setColorPrimaries:"), value)
}/* debug [instance_properties/setter]: colorPrimaries */


// The transfer function used for video composition.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avvideocomposition/colortransferfunction
func (s_ SpatialVideoConfiguration) ColorTransferFunction() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](s_.ID, objc.Sel("colorTransferFunction"))
	return rv
}/* debug [instance_properties/getter]: colorTransferFunction */


// The transfer function used for video composition.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avvideocomposition/colortransferfunction
func (s_ SpatialVideoConfiguration) SetColorTransferFunction(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setColorTransferFunction:"), value)
}/* debug [instance_properties/setter]: colorTransferFunction */


// The YCbCr matrix used for video composition.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avvideocomposition/colorycbcrmatrix
func (s_ SpatialVideoConfiguration) ColorYCbCrMatrix() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](s_.ID, objc.Sel("colorYCbCrMatrix"))
	return rv
}/* debug [instance_properties/getter]: colorYCbCrMatrix */


// The YCbCr matrix used for video composition.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avvideocomposition/colorycbcrmatrix
func (s_ SpatialVideoConfiguration) SetColorYCbCrMatrix(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setColorYCbCrMatrix:"), value)
}/* debug [instance_properties/setter]: colorYCbCrMatrix */


// A custom compositor class to use.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avvideocomposition/customvideocompositorclass
func (s_ SpatialVideoConfiguration) CustomVideoCompositorClass() VideoCompositing /* not a class type */ {
	rv := objc.Send[VideoCompositing](s_.ID, objc.Sel("customVideoCompositorClass"))
	return rv
}/* debug [instance_properties/getter]: customVideoCompositorClass */


// A custom compositor class to use.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avvideocomposition/customvideocompositorclass
func (s_ SpatialVideoConfiguration) SetCustomVideoCompositorClass(value VideoCompositing /* not a class type */) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setCustomVideoCompositorClass:"), value)
}/* debug [instance_properties/setter]: customVideoCompositorClass */


// A time interval for which the video composition should render composed video frames.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avvideocomposition/frameduration
func (s_ SpatialVideoConfiguration) FrameDuration() objc.IObject /* cross-framework: Time */ {
	rv := objc.Send[corevideo.Time](s_.ID, objc.Sel("frameDuration"))
	return rv
}/* debug [instance_properties/getter]: frameDuration */


// A time interval for which the video composition should render composed video frames.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avvideocomposition/frameduration
func (s_ SpatialVideoConfiguration) SetFrameDuration(value objc.IObject /* cross-framework: Time */) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setFrameDuration:"), value)
}/* debug [instance_properties/setter]: frameDuration */


// The scale at which the video composition should render.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avvideocomposition/renderscale
func (s_ SpatialVideoConfiguration) RenderScale() float32 {
	rv := objc.Send[float32](s_.ID, objc.Sel("renderScale"))
	return rv
}/* debug [instance_properties/getter]: renderScale */


// The scale at which the video composition should render.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avvideocomposition/renderscale
func (s_ SpatialVideoConfiguration) SetRenderScale(value float32) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setRenderScale:"), value)
}/* debug [instance_properties/setter]: renderScale */


// The size at which the video composition should render.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avvideocomposition/rendersize
func (s_ SpatialVideoConfiguration) RenderSize() corefoundation.CGSize {
	rv := objc.Send[corefoundation.CGSize](s_.ID, objc.Sel("renderSize"))
	return rv
}/* debug [instance_properties/getter]: renderSize */


// The size at which the video composition should render.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avvideocomposition/rendersize
func (s_ SpatialVideoConfiguration) SetRenderSize(value corefoundation.CGSize) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setRenderSize:"), value)
}/* debug [instance_properties/setter]: renderSize */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class AVSpatialVideoConfiguration */


