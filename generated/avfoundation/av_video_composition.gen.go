// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/coregraphics"
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
	IsValidForAssetTimeRangeValidationDelegate(asset unsafe.Pointer, timeRange unsafe.Pointer, validationDelegate objc.ID) bool
}

// An object that describes how to compose video frames at particular points in time.
//
// If you use the built-in video compositor, the instructions a video composition contain can specify a spatial transformation, an opacity value, and a cropping rectangle for each video source. This values can vary over time by applying linear ramping functions. You can create a custom video compositor by implementing the protocol. The system provides the custom video compositor with pixel buffers for each of its video sources during playback, and can perform arbitrary graphical operations on them to produce visual output.
//
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

// Alloc allocates a new instance without initialization.
func (vc _VideoCompositionClass) Alloc() VideoComposition {
	rv := objc.Send[VideoComposition](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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


// Indicates whether the time ranges of the composition’s instructions conform to validation requirements.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVVideoComposition/isValid(for:timeRange:validationDelegate:)
func (v_ VideoComposition) IsValidForAssetTimeRangeValidationDelegate(asset unsafe.Pointer, timeRange unsafe.Pointer, validationDelegate objc.ID) bool {
	rv := objc.Send[bool](v_.ID, objc.Sel("isValidForAsset:timeRange:validationDelegate:"), asset, timeRange, validationDelegate)
	return rv
}

// A custom compositor class to use.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVVideoComposition/customVideoCompositorClass
func (v_ VideoComposition) CustomVideoCompositorClass() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](v_.ID, objc.Sel("customVideoCompositorClass"))
	return rv
}

// A time interval for which the video composition should render composed video frames.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVVideoComposition/frameDuration
func (v_ VideoComposition) FrameDuration() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](v_.ID, objc.Sel("frameDuration"))
	return rv
}

// The video composition instructions.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVVideoComposition/instructions
func (v_ VideoComposition) Instructions() []objc.ID {
	rv := objc.Send[[]objc.ID](v_.ID, objc.Sel("instructions"))
	return rv
}

// The scale at which the video composition should render.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVVideoComposition/renderScale
func (v_ VideoComposition) RenderScale() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](v_.ID, objc.Sel("renderScale"))
	return rv
}

// A video composition tool to use with Core Animation in offline rendering.
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avvideocomposition/animationtool
func (v_ VideoComposition) AnimationTool() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](v_.ID, objc.Sel("animationTool"))
	return rv
}


// SetAnimationTool sets the value of the animationTool property.
// A video composition tool to use with Core Animation in offline rendering.

//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avvideocomposition/animationtool
func (v_ VideoComposition) SetAnimationTool(value unsafe.Pointer) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setAnimationTool:"), value)
}

// The color primaries used for video composition.
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avvideocomposition/colorprimaries
func (v_ VideoComposition) ColorPrimaries() string {
	rv := objc.Send[string](v_.ID, objc.Sel("colorPrimaries"))
	return rv
}


// SetColorPrimaries sets the value of the colorPrimaries property.
// The color primaries used for video composition.

//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avvideocomposition/colorprimaries
func (v_ VideoComposition) SetColorPrimaries(value string) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setColorPrimaries:"), objc.String(value))
}

// The transfer function used for video composition.
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avvideocomposition/colortransferfunction
func (v_ VideoComposition) ColorTransferFunction() string {
	rv := objc.Send[string](v_.ID, objc.Sel("colorTransferFunction"))
	return rv
}


// SetColorTransferFunction sets the value of the colorTransferFunction property.
// The transfer function used for video composition.

//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avvideocomposition/colortransferfunction
func (v_ VideoComposition) SetColorTransferFunction(value string) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setColorTransferFunction:"), objc.String(value))
}

// The YCbCr matrix used for video composition.
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avvideocomposition/colorycbcrmatrix
func (v_ VideoComposition) ColorYCbCrMatrix() string {
	rv := objc.Send[string](v_.ID, objc.Sel("colorYCbCrMatrix"))
	return rv
}


// SetColorYCbCrMatrix sets the value of the colorYCbCrMatrix property.
// The YCbCr matrix used for video composition.

//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avvideocomposition/colorycbcrmatrix
func (v_ VideoComposition) SetColorYCbCrMatrix(value string) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setColorYCbCrMatrix:"), objc.String(value))
}

// The output buffers of the video composition can be specified with the outputBufferDescription. The value is an array of an array of CMTag objects that describes the output buffers.
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avvideocomposition/outputbufferdescription-3ayt8
func (v_ VideoComposition) OutputBufferDescription() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](v_.ID, objc.Sel("outputBufferDescription"))
	return rv
}


// SetOutputBufferDescription sets the value of the outputBufferDescription property.
// The output buffers of the video composition can be specified with the outputBufferDescription. The value is an array of an array of CMTag objects that describes the output buffers.

//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avvideocomposition/outputbufferdescription-3ayt8
func (v_ VideoComposition) SetOutputBufferDescription(value unsafe.Pointer) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setOutputBufferDescription:"), value)
}

// The policy for display of HDR display metadata on the rendered frame.
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avvideocomposition/perframehdrdisplaymetadatapolicy-swift.property
func (v_ VideoComposition) PerFrameHDRDisplayMetadataPolicy() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](v_.ID, objc.Sel("perFrameHDRDisplayMetadataPolicy"))
	return rv
}


// SetPerFrameHDRDisplayMetadataPolicy sets the value of the perFrameHDRDisplayMetadataPolicy property.
// The policy for display of HDR display metadata on the rendered frame.

//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avvideocomposition/perframehdrdisplaymetadatapolicy-swift.property
func (v_ VideoComposition) SetPerFrameHDRDisplayMetadataPolicy(value unsafe.Pointer) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setPerFrameHDRDisplayMetadataPolicy:"), value)
}

// The size at which the video composition should render.
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avvideocomposition/rendersize
func (v_ VideoComposition) RenderSize() coregraphics.CGSize {
	rv := objc.Send[coregraphics.CGSize](v_.ID, objc.Sel("renderSize"))
	return rv
}


// SetRenderSize sets the value of the renderSize property.
// The size at which the video composition should render.

//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avvideocomposition/rendersize
func (v_ VideoComposition) SetRenderSize(value coregraphics.CGSize) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setRenderSize:"), value)
}

// The identifiers of source sample data tracks in the composition that the object requires to compose frames.
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avvideocomposition/sourcesampledatatrackids-2hgue
func (v_ VideoComposition) SourceSampleDataTrackIDs() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](v_.ID, objc.Sel("sourceSampleDataTrackIDs"))
	return rv
}


// SetSourceSampleDataTrackIDs sets the value of the sourceSampleDataTrackIDs property.
// The identifiers of source sample data tracks in the composition that the object requires to compose frames.

//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avvideocomposition/sourcesampledatatrackids-2hgue
func (v_ VideoComposition) SetSourceSampleDataTrackIDs(value unsafe.Pointer) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setSourceSampleDataTrackIDs:"), value)
}

// An identifier of the source track from which the video composition derives frame timing.
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avvideocomposition/sourcetrackidforframetiming
func (v_ VideoComposition) SourceTrackIDForFrameTiming() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](v_.ID, objc.Sel("sourceTrackIDForFrameTiming"))
	return rv
}


// SetSourceTrackIDForFrameTiming sets the value of the sourceTrackIDForFrameTiming property.
// An identifier of the source track from which the video composition derives frame timing.

//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avvideocomposition/sourcetrackidforframetiming
func (v_ VideoComposition) SetSourceTrackIDForFrameTiming(value unsafe.Pointer) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setSourceTrackIDForFrameTiming:"), value)
}

// Indicates the spatial configurations that are available to associate with the output of the video composition.
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avvideocomposition/spatialvideoconfigurations-80iab
func (v_ VideoComposition) SpatialVideoConfigurations() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](v_.ID, objc.Sel("spatialVideoConfigurations"))
	return rv
}


// SetSpatialVideoConfigurations sets the value of the spatialVideoConfigurations property.
// Indicates the spatial configurations that are available to associate with the output of the video composition.

//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avvideocomposition/spatialvideoconfigurations-80iab
func (v_ VideoComposition) SetSpatialVideoConfigurations(value unsafe.Pointer) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setSpatialVideoConfigurations:"), value)
}



