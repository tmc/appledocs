// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/corefoundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class AVVideoComposition */


/* debug [class_header]: Header for AVVideoComposition */
// The class instance for the [VideoComposition] class.
var (
	VideoCompositionClass     _VideoCompositionClass
	VideoCompositionClassOnce sync.Once
)

func getVideoCompositionClass() _VideoCompositionClass {
	VideoCompositionClassOnce.Do(func() {
		VideoCompositionClass = _VideoCompositionClass{objc.GetClass("AVVideoComposition")}
	})
	return VideoCompositionClass
}

type _VideoCompositionClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for VideoComposition */
// An interface definition for the [VideoComposition] class.
type IVideoComposition interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for VideoComposition */
	// properties:
	AnimationTool() IAVVideoCompositionCoreAnimationTool
	ColorPrimaries() objc.IObject /* cross-framework: NSString */
	ColorTransferFunction() objc.IObject /* cross-framework: NSString */
	ColorYCbCrMatrix() objc.IObject /* cross-framework: NSString */
	CustomVideoCompositorClass() unsafe.Pointer
	FrameDuration() objc.IObject /* cross-framework: Time */
	Instructions() []objc.ID
	OutputBufferDescription() objc.IObject /* cross-framework: NSArray */
	PerFrameHDRDisplayMetadataPolicy() VideoCompositionPerFrameHDRDisplayMetadataPolicy /* typedef */
	RenderScale() float32
	RenderSize() corefoundation.CGSize
	SourceSampleDataTrackIDs() []foundation.Number
	SourceTrackIDForFrameTiming() PersistentTrackID /* not a class type */
	SpatialVideoConfigurations() []SpatialVideoConfiguration
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for VideoComposition */
	// methods:
	IsValidForTracksAssetDurationTimeRangeValidationDelegate(tracks []AssetTrack, duration objc.IObject /* cross-framework: Time */, timeRange TimeRange /* not a class type */, validationDelegate unsafe.Pointer) bool
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for VideoComposition */
// Alloc allocates a new instance without initialization.
func (vc _VideoCompositionClass) Alloc() VideoComposition {
	rv := objc.Send[VideoComposition](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (vc _VideoCompositionClass) New() VideoComposition {
	rv := objc.Send[VideoComposition](objc.ID(vc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (v_ VideoComposition) Init() VideoComposition {
	rv := objc.Send[VideoComposition](v_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v_ VideoComposition) Autorelease() VideoComposition {
	rv := objc.Send[VideoComposition](v_.ID, objc.Sel("autorelease"))
	return rv
}

// NewVideoComposition creates a new VideoComposition instance.
func NewVideoComposition() VideoComposition {
	return getVideoCompositionClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for VideoComposition */
// An object that describes how to compose video frames at particular points in time.
//
// If you use the built-in video compositor, the instructions a video composition contain can specify a spatial transformation, an opacity value, and a cropping rectangle for each video source. This values can vary over time by applying linear ramping functions. You can create a custom video compositor by implementing the protocol. The system provides the custom video compositor with pixel buffers for each of its video sources during playback, and can perform arbitrary graphical operations on them to produce visual output.


// An object that describes how to compose video frames at particular points in time.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVVideoComposition
type VideoComposition struct {
	objectivec.Object
}

// VideoCompositionFrom constructs a [VideoComposition] from an unsafe.Pointer.
//
// An object that describes how to compose video frames at particular points in time.
func VideoCompositionFrom(ptr unsafe.Pointer) VideoComposition {
	return VideoComposition{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for VideoComposition */

// Creates a video composition configured to apply Core Image filters to each video frame of the specified asset.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVVideoComposition/init(asset:applyingCIFiltersWithHandler:)
func NewVideoCompositionWithAssetApplyingCIFiltersWithHandler(asset IAVAsset, applier unsafe.Pointer) VideoComposition {
	rv := objc.Send[VideoComposition](objc.ID(getVideoCompositionClass().class), objc.Sel("videoCompositionWithAsset:applyingCIFiltersWithHandler:"), asset, applier)
	return rv
}/* debug [class_init_methods/constructor]: NewVideoCompositionWithAssetApplyingCIFiltersWithHandler */


// Creates a video composition object configured to present the video tracks of the specified asset.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVVideoComposition/init(propertiesOf:)
func NewVideoCompositionWithPropertiesOfAsset(asset IAVAsset) VideoComposition {
	rv := objc.Send[VideoComposition](objc.ID(getVideoCompositionClass().class), objc.Sel("videoCompositionWithPropertiesOfAsset:"), asset)
	return rv
}/* debug [class_init_methods/constructor]: NewVideoCompositionWithPropertiesOfAsset */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for VideoComposition */

// Creates a video composition configured to apply Core Image filters to each video frame of the specified asset.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVVideoComposition/init(asset:applyingCIFiltersWithHandler:)
func (vc _VideoCompositionClass) VideoCompositionWithAssetApplyingCIFiltersWithHandler(asset IAVAsset, applier unsafe.Pointer) IVideoComposition {
	rv := objc.Send[VideoComposition](objc.ID(vc.class), objc.Sel("videoCompositionWithAsset:applyingCIFiltersWithHandler:"), asset, applier)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=VideoCompositionWithAssetApplyingCIFiltersWithHandler) */


// Creates a video composition object configured to present the video tracks of the specified asset.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVVideoComposition/init(propertiesOf:)
func (vc _VideoCompositionClass) VideoCompositionWithPropertiesOfAsset(asset IAVAsset) IVideoComposition {
	rv := objc.Send[VideoComposition](objc.ID(vc.class), objc.Sel("videoCompositionWithPropertiesOfAsset:"), asset)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=VideoCompositionWithPropertiesOfAsset) */


// Pass-through initializer, for internal use in AVFoundation only
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVVideoComposition/videoCompositionWithVideoComposition:
func (vc _VideoCompositionClass) VideoCompositionWithVideoComposition(videoComposition IAVVideoComposition) IVideoComposition {
	rv := objc.Send[VideoComposition](objc.ID(vc.class), objc.Sel("videoCompositionWithVideoComposition:"), videoComposition)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=VideoCompositionWithVideoComposition) */


// Returns a new video composition that’s configured to apply Core Image filters to each video frame of the specified asset.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVVideoComposition/videoComposition(with:applyingCIFiltersWithHandler:completionHandler:)
func (vc _VideoCompositionClass) VideoCompositionWithAssetApplyingCIFiltersWithHandlerCompletionHandler(asset IAVAsset, applier unsafe.Pointer, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(vc.class), objc.Sel("videoCompositionWithAsset:applyingCIFiltersWithHandler:completionHandler:"), asset, applier, completionHandler)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=VideoCompositionWithAssetApplyingCIFiltersWithHandlerCompletionHandler) */


// Returns a new video composition that’s configured to present the video tracks of the specified asset.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVVideoComposition/videoComposition(withPropertiesOf:completionHandler:)
func (vc _VideoCompositionClass) VideoCompositionWithPropertiesOfAssetCompletionHandler(asset IAVAsset, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(vc.class), objc.Sel("videoCompositionWithPropertiesOfAsset:completionHandler:"), asset, completionHandler)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=VideoCompositionWithPropertiesOfAssetCompletionHandler) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for VideoComposition */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for VideoComposition */

// Indicates whether the time ranges of the composition’s instructions conform to validation requirements.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVVideoComposition/isValid(for:assetDuration:timeRange:validationDelegate:)
func (v_ VideoComposition) IsValidForTracksAssetDurationTimeRangeValidationDelegate(tracks []AssetTrack, duration objc.IObject /* cross-framework: Time */, timeRange TimeRange /* not a class type */, validationDelegate unsafe.Pointer) bool {
	rv := objc.Send[bool](v_.ID, objc.Sel("isValidForTracks:assetDuration:timeRange:validationDelegate:"), tracks, duration, timeRange, validationDelegate)
	return rv
}/* debug [instance_methods/method]: IsValidForTracksAssetDurationTimeRangeValidationDelegate */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for VideoComposition */

// A video composition tool to use with Core Animation in offline rendering.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVVideoComposition/animationTool
func (v_ VideoComposition) AnimationTool() IAVVideoCompositionCoreAnimationTool {
	rv := objc.Send[VideoCompositionCoreAnimationTool](v_.ID, objc.Sel("animationTool"))
	return rv
}/* debug [instance_properties/getter]: animationTool */


// The color primaries used for video composition.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVVideoComposition/colorPrimaries
func (v_ VideoComposition) ColorPrimaries() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](v_.ID, objc.Sel("colorPrimaries"))
	return rv
}/* debug [instance_properties/getter]: colorPrimaries */


// The transfer function used for video composition.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVVideoComposition/colorTransferFunction
func (v_ VideoComposition) ColorTransferFunction() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](v_.ID, objc.Sel("colorTransferFunction"))
	return rv
}/* debug [instance_properties/getter]: colorTransferFunction */


// The YCbCr matrix used for video composition.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVVideoComposition/colorYCbCrMatrix
func (v_ VideoComposition) ColorYCbCrMatrix() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](v_.ID, objc.Sel("colorYCbCrMatrix"))
	return rv
}/* debug [instance_properties/getter]: colorYCbCrMatrix */


// A custom compositor class to use.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVVideoComposition/customVideoCompositorClass
func (v_ VideoComposition) CustomVideoCompositorClass() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](v_.ID, objc.Sel("customVideoCompositorClass"))
	return rv
}/* debug [instance_properties/getter]: customVideoCompositorClass */


// A time interval for which the video composition should render composed video frames.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVVideoComposition/frameDuration
func (v_ VideoComposition) FrameDuration() objc.IObject /* cross-framework: Time */ {
	rv := objc.Send[corevideo.Time](v_.ID, objc.Sel("frameDuration"))
	return rv
}/* debug [instance_properties/getter]: frameDuration */


// The video composition instructions.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVVideoComposition/instructions
func (v_ VideoComposition) Instructions() []objc.ID {
	rv := objc.Send[[]objc.ID](v_.ID, objc.Sel("instructions"))
	return rv
}/* debug [instance_properties/getter]: instructions */


// The output buffers of the video composition can be specified with the outputBufferDescription. The value is an array of CMTagCollectionRef objects that describes the output buffers.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVVideoComposition/outputBufferDescription-3wsar
func (v_ VideoComposition) OutputBufferDescription() objc.IObject /* cross-framework: NSArray */ {
	rv := objc.Send[foundation.NSArray](v_.ID, objc.Sel("outputBufferDescription"))
	return rv
}/* debug [instance_properties/getter]: outputBufferDescription */


// The policy for display of HDR display metadata on the rendered frame.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVVideoComposition/perFrameHDRDisplayMetadataPolicy-swift.property
func (v_ VideoComposition) PerFrameHDRDisplayMetadataPolicy() VideoCompositionPerFrameHDRDisplayMetadataPolicy /* typedef */ {
	rv := objc.Send[foundation.NSString](v_.ID, objc.Sel("perFrameHDRDisplayMetadataPolicy"))
	return rv
}/* debug [instance_properties/getter]: perFrameHDRDisplayMetadataPolicy */


// The scale at which the video composition should render.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVVideoComposition/renderScale
func (v_ VideoComposition) RenderScale() float32 {
	rv := objc.Send[float32](v_.ID, objc.Sel("renderScale"))
	return rv
}/* debug [instance_properties/getter]: renderScale */


// The size at which the video composition should render.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVVideoComposition/renderSize
func (v_ VideoComposition) RenderSize() corefoundation.CGSize {
	rv := objc.Send[corefoundation.CGSize](v_.ID, objc.Sel("renderSize"))
	return rv
}/* debug [instance_properties/getter]: renderSize */


// The identifiers of source sample data tracks in the composition that the compositor requires to compose frames.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVVideoComposition/sourceSampleDataTrackIDs-3nrgi
func (v_ VideoComposition) SourceSampleDataTrackIDs() []foundation.Number {
	rv := objc.Send[[]foundation.Number](v_.ID, objc.Sel("sourceSampleDataTrackIDs"))
	return rv
}/* debug [instance_properties/getter]: sourceSampleDataTrackIDs */


// An identifier of the source track from which the video composition derives frame timing.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVVideoComposition/sourceTrackIDForFrameTiming
func (v_ VideoComposition) SourceTrackIDForFrameTiming() PersistentTrackID /* not a class type */ {
	rv := objc.Send[PersistentTrackID](v_.ID, objc.Sel("sourceTrackIDForFrameTiming"))
	return rv
}/* debug [instance_properties/getter]: sourceTrackIDForFrameTiming */


// Indicates the spatial configurations that are available to associate with the output of the video composition.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVVideoComposition/spatialVideoConfigurations-2ipps
func (v_ VideoComposition) SpatialVideoConfigurations() []SpatialVideoConfiguration {
	rv := objc.Send[[]SpatialVideoConfiguration](v_.ID, objc.Sel("spatialVideoConfigurations"))
	return rv
}/* debug [instance_properties/getter]: spatialVideoConfigurations */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class AVVideoComposition */


