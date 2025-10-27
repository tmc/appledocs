// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/corefoundation"
	"github.com/tmc/appledocs/generated/objectivec"
)





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





// An interface definition for the [VideoComposition] class.
type IVideoComposition interface {
	objectivec.IObject
	

	// properties:
	AnimationTool() IAVVideoCompositionCoreAnimationTool
	ColorPrimaries() foundation.foundation.INSString
	ColorTransferFunction() foundation.foundation.INSString
	ColorYCbCrMatrix() foundation.foundation.INSString
	CustomVideoCompositorClass() unsafe.Pointer
	FrameDuration() objectivec.IObject
	Instructions() []objc.ID
	OutputBufferDescription() foundation.foundation.INSArray
	PerFrameHDRDisplayMetadataPolicy() VideoCompositionPerFrameHDRDisplayMetadataPolicy
	RenderScale() float32
	RenderSize() corefoundation.CGSize
	SourceSampleDataTrackIDs() []foundation.Number
	SourceTrackIDForFrameTiming() PersistentTrackID /* not a class type */
	SpatialVideoConfigurations() []SpatialVideoConfiguration


	

	// methods:
	IsValidForTracksAssetDurationTimeRangeValidationDelegate(tracks []AssetTrack, duration objectivec.IObject, timeRange objectivec.IObject, validationDelegate unsafe.Pointer) bool


}





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






// Creates a video composition configured to apply Core Image filters to each video frame of the specified asset.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVVideoComposition/init(asset:applyingCIFiltersWithHandler:)
func NewVideoCompositionWithAssetApplyingCIFiltersWithHandler(asset IAVAsset, applier unsafe.Pointer) VideoComposition {
	rv := objc.Send[VideoComposition](objc.ID(getVideoCompositionClass().class), objc.Sel("videoCompositionWithAsset:applyingCIFiltersWithHandler:"), asset, applier)
	return rv
}


// Creates a video composition object configured to present the video tracks of the specified asset.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVVideoComposition/init(propertiesOf:)
func NewVideoCompositionWithPropertiesOfAsset(asset IAVAsset) VideoComposition {
	rv := objc.Send[VideoComposition](objc.ID(getVideoCompositionClass().class), objc.Sel("videoCompositionWithPropertiesOfAsset:"), asset)
	return rv
}







// Creates a video composition configured to apply Core Image filters to each video frame of the specified asset.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVVideoComposition/init(asset:applyingCIFiltersWithHandler:)
func (vc _VideoCompositionClass) VideoCompositionWithAssetApplyingCIFiltersWithHandler(asset IAVAsset, applier unsafe.Pointer) IVideoComposition {
	rv := objc.Send[VideoComposition](objc.ID(vc.class), objc.Sel("videoCompositionWithAsset:applyingCIFiltersWithHandler:"), asset, applier)
	return rv
}


// Creates a video composition object configured to present the video tracks of the specified asset.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVVideoComposition/init(propertiesOf:)
func (vc _VideoCompositionClass) VideoCompositionWithPropertiesOfAsset(asset IAVAsset) IVideoComposition {
	rv := objc.Send[VideoComposition](objc.ID(vc.class), objc.Sel("videoCompositionWithPropertiesOfAsset:"), asset)
	return rv
}


// Pass-through initializer, for internal use in AVFoundation only
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVVideoComposition/videoCompositionWithVideoComposition:
func (vc _VideoCompositionClass) VideoCompositionWithVideoComposition(videoComposition IAVVideoComposition) IVideoComposition {
	rv := objc.Send[VideoComposition](objc.ID(vc.class), objc.Sel("videoCompositionWithVideoComposition:"), videoComposition)
	return rv
}


// Returns a new video composition that’s configured to apply Core Image filters to each video frame of the specified asset.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVVideoComposition/videoComposition(with:applyingCIFiltersWithHandler:completionHandler:)
func (vc _VideoCompositionClass) VideoCompositionWithAssetApplyingCIFiltersWithHandlerCompletionHandler(asset IAVAsset, applier unsafe.Pointer, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(vc.class), objc.Sel("videoCompositionWithAsset:applyingCIFiltersWithHandler:completionHandler:"), asset, applier, completionHandler)
}


// Returns a new video composition that’s configured to present the video tracks of the specified asset.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVVideoComposition/videoComposition(withPropertiesOf:completionHandler:)
func (vc _VideoCompositionClass) VideoCompositionWithPropertiesOfAssetCompletionHandler(asset IAVAsset, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(vc.class), objc.Sel("videoCompositionWithPropertiesOfAsset:completionHandler:"), asset, completionHandler)
}












// Indicates whether the time ranges of the composition’s instructions conform to validation requirements.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVVideoComposition/isValid(for:assetDuration:timeRange:validationDelegate:)
func (v_ VideoComposition) IsValidForTracksAssetDurationTimeRangeValidationDelegate(tracks []AssetTrack, duration objectivec.IObject, timeRange objectivec.IObject, validationDelegate unsafe.Pointer) bool {
	rv := objc.Send[bool](v_.ID, objc.Sel("isValidForTracks:assetDuration:timeRange:validationDelegate:"), tracks, duration, timeRange, validationDelegate)
	return rv
}







// A video composition tool to use with Core Animation in offline rendering.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVVideoComposition/animationTool
func (v_ VideoComposition) AnimationTool() IAVVideoCompositionCoreAnimationTool {
	rv := objc.Send[VideoCompositionCoreAnimationTool](v_.ID, objc.Sel("animationTool"))
	return rv
}


// The color primaries used for video composition.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVVideoComposition/colorPrimaries
func (v_ VideoComposition) ColorPrimaries() foundation.foundation.INSString {
	rv := objc.Send[foundation.NSString](v_.ID, objc.Sel("colorPrimaries"))
	return rv
}


// The transfer function used for video composition.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVVideoComposition/colorTransferFunction
func (v_ VideoComposition) ColorTransferFunction() foundation.foundation.INSString {
	rv := objc.Send[foundation.NSString](v_.ID, objc.Sel("colorTransferFunction"))
	return rv
}


// The YCbCr matrix used for video composition.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVVideoComposition/colorYCbCrMatrix
func (v_ VideoComposition) ColorYCbCrMatrix() foundation.foundation.INSString {
	rv := objc.Send[foundation.NSString](v_.ID, objc.Sel("colorYCbCrMatrix"))
	return rv
}


// A custom compositor class to use.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVVideoComposition/customVideoCompositorClass
func (v_ VideoComposition) CustomVideoCompositorClass() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](v_.ID, objc.Sel("customVideoCompositorClass"))
	return rv
}


// A time interval for which the video composition should render composed video frames.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVVideoComposition/frameDuration
func (v_ VideoComposition) FrameDuration() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](v_.ID, objc.Sel("frameDuration"))
	return rv
}


// The video composition instructions.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVVideoComposition/instructions
func (v_ VideoComposition) Instructions() []objc.ID {
	rv := objc.Send[[]objc.ID](v_.ID, objc.Sel("instructions"))
	return rv
}


// The output buffers of the video composition can be specified with the outputBufferDescription. The value is an array of CMTagCollectionRef objects that describes the output buffers.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVVideoComposition/outputBufferDescription-3wsar
func (v_ VideoComposition) OutputBufferDescription() foundation.foundation.INSArray {
	rv := objc.Send[foundation.NSArray](v_.ID, objc.Sel("outputBufferDescription"))
	return rv
}


// The policy for display of HDR display metadata on the rendered frame.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVVideoComposition/perFrameHDRDisplayMetadataPolicy-swift.property
func (v_ VideoComposition) PerFrameHDRDisplayMetadataPolicy() VideoCompositionPerFrameHDRDisplayMetadataPolicy {
	rv := objc.Send[VideoCompositionPerFrameHDRDisplayMetadataPolicy](v_.ID, objc.Sel("perFrameHDRDisplayMetadataPolicy"))
	return rv
}


// The scale at which the video composition should render.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVVideoComposition/renderScale
func (v_ VideoComposition) RenderScale() float32 {
	rv := objc.Send[float32](v_.ID, objc.Sel("renderScale"))
	return rv
}


// The size at which the video composition should render.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVVideoComposition/renderSize
func (v_ VideoComposition) RenderSize() corefoundation.CGSize {
	rv := objc.Send[corefoundation.CGSize](v_.ID, objc.Sel("renderSize"))
	return rv
}


// The identifiers of source sample data tracks in the composition that the compositor requires to compose frames.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVVideoComposition/sourceSampleDataTrackIDs-3nrgi
func (v_ VideoComposition) SourceSampleDataTrackIDs() []foundation.Number {
	rv := objc.Send[[]foundation.Number](v_.ID, objc.Sel("sourceSampleDataTrackIDs"))
	return rv
}


// An identifier of the source track from which the video composition derives frame timing.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVVideoComposition/sourceTrackIDForFrameTiming
func (v_ VideoComposition) SourceTrackIDForFrameTiming() PersistentTrackID /* not a class type */ {
	rv := objc.Send[PersistentTrackID](v_.ID, objc.Sel("sourceTrackIDForFrameTiming"))
	return rv
}


// Indicates the spatial configurations that are available to associate with the output of the video composition.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVVideoComposition/spatialVideoConfigurations-2ipps
func (v_ VideoComposition) SpatialVideoConfigurations() []SpatialVideoConfiguration {
	rv := objc.Send[[]SpatialVideoConfiguration](v_.ID, objc.Sel("spatialVideoConfigurations"))
	return rv
}







