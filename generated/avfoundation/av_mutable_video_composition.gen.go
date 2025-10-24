// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class AVMutableVideoComposition */


/* debug [class_header]: Header for AVMutableVideoComposition */
// The class instance for the [MutableVideoComposition] class.
var (
	MutableVideoCompositionClass     _MutableVideoCompositionClass
	MutableVideoCompositionClassOnce sync.Once
)

func getMutableVideoCompositionClass() _MutableVideoCompositionClass {
	MutableVideoCompositionClassOnce.Do(func() {
		MutableVideoCompositionClass = _MutableVideoCompositionClass{objc.GetClass("AVMutableVideoComposition")}
	})
	return MutableVideoCompositionClass
}

type _MutableVideoCompositionClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MutableVideoComposition */
// An interface definition for the [MutableVideoComposition] class.
type IMutableVideoComposition interface {
	IVideoComposition
	
/* debug [class_interface_properties]: Properties for MutableVideoComposition */
	// properties:
	AnimationTool() IAVVideoCompositionCoreAnimationTool
	SetAnimationTool(value IAVVideoCompositionCoreAnimationTool)
	ColorPrimaries() objc.IObject /* cross-framework: NSString */
	SetColorPrimaries(value objc.IObject /* cross-framework: NSString */)
	ColorTransferFunction() objc.IObject /* cross-framework: NSString */
	SetColorTransferFunction(value objc.IObject /* cross-framework: NSString */)
	ColorYCbCrMatrix() objc.IObject /* cross-framework: NSString */
	SetColorYCbCrMatrix(value objc.IObject /* cross-framework: NSString */)
	CustomVideoCompositorClass() unsafe.Pointer
	SetCustomVideoCompositorClass(value unsafe.Pointer)
	FrameDuration() objc.IObject /* cross-framework: Time */
	SetFrameDuration(value objc.IObject /* cross-framework: Time */)
	Instructions() []objc.ID
	SetInstructions(value []objc.ID)
	OutputBufferDescription() objc.IObject /* cross-framework: NSArray */
	SetOutputBufferDescription(value objc.IObject /* cross-framework: NSArray */)
	PerFrameHDRDisplayMetadataPolicy() VideoCompositionPerFrameHDRDisplayMetadataPolicy /* typedef */
	SetPerFrameHDRDisplayMetadataPolicy(value VideoCompositionPerFrameHDRDisplayMetadataPolicy /* typedef */)
	RenderScale() float32
	SetRenderScale(value float32)
	RenderSize() corefoundation.CGSize
	SetRenderSize(value corefoundation.CGSize)
	SourceSampleDataTrackIDs() []foundation.Number
	SetSourceSampleDataTrackIDs(value []foundation.Number)
	SourceTrackIDForFrameTiming() PersistentTrackID /* not a class type */
	SetSourceTrackIDForFrameTiming(value PersistentTrackID /* not a class type */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MutableVideoComposition */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MutableVideoComposition */
// Alloc allocates a new instance without initialization.
func (mc _MutableVideoCompositionClass) Alloc() MutableVideoComposition {
	rv := objc.Send[MutableVideoComposition](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MutableVideoCompositionClass) New() MutableVideoComposition {
	rv := objc.Send[MutableVideoComposition](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MutableVideoComposition) Init() MutableVideoComposition {
	rv := objc.Send[MutableVideoComposition](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MutableVideoComposition) Autorelease() MutableVideoComposition {
	rv := objc.Send[MutableVideoComposition](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMutableVideoComposition creates a new MutableVideoComposition instance.
func NewMutableVideoComposition() MutableVideoComposition {
	return getMutableVideoCompositionClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MutableVideoComposition */
// A mutable video composition subclass.
//
// If you use the built-in video compositor, the instructions a video composition contain can specify a spatial transformation, an opacity value, and a cropping rectangle for each video source. This values can vary over time by applying linear ramping functions. You can create a custom video compositor by implementing the protocol. The system provides the custom video compositor with pixel buffers for each of its video sources during playback, and can perform arbitrary graphical operations on them to produce visual output.


// A mutable video composition subclass.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMutableVideoComposition
type MutableVideoComposition struct {
	VideoComposition
}

// MutableVideoCompositionFrom constructs a [MutableVideoComposition] from an unsafe.Pointer.
//
// A mutable video composition subclass.
func MutableVideoCompositionFrom(ptr unsafe.Pointer) MutableVideoComposition {
	return MutableVideoComposition{
		VideoComposition: VideoCompositionFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MutableVideoComposition */

// Creates a mutable video composition configured to apply Core Image filters to each video frame of the specified asset.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMutableVideoComposition/init(asset:applyingCIFiltersWithHandler:)
func NewMutableVideoCompositionWithAssetApplyingCIFiltersWithHandler(asset IAVAsset, applier unsafe.Pointer) MutableVideoComposition {
	rv := objc.Send[MutableVideoComposition](objc.ID(getMutableVideoCompositionClass().class), objc.Sel("videoCompositionWithAsset:applyingCIFiltersWithHandler:"), asset, applier)
	return rv
}/* debug [class_init_methods/constructor]: NewMutableVideoCompositionWithAssetApplyingCIFiltersWithHandler */


// Creates a mutable video composition with the specified asset properties.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMutableVideoComposition/init(propertiesOf:)
func NewMutableVideoCompositionWithPropertiesOfAsset(asset IAVAsset) MutableVideoComposition {
	rv := objc.Send[MutableVideoComposition](objc.ID(getMutableVideoCompositionClass().class), objc.Sel("videoCompositionWithPropertiesOfAsset:"), asset)
	return rv
}/* debug [class_init_methods/constructor]: NewMutableVideoCompositionWithPropertiesOfAsset */


// Creates a mutable video composition with the specified asset properties and a prototype video composition instruction.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMutableVideoComposition/init(propertiesOf:prototypeInstruction:)
func NewMutableVideoCompositionWithPropertiesOfAssetPrototypeInstruction(asset IAVAsset, prototypeInstruction IAVVideoCompositionInstruction) MutableVideoComposition {
	rv := objc.Send[MutableVideoComposition](objc.ID(getMutableVideoCompositionClass().class), objc.Sel("videoCompositionWithPropertiesOfAsset:prototypeInstruction:"), asset, prototypeInstruction)
	return rv
}/* debug [class_init_methods/constructor]: NewMutableVideoCompositionWithPropertiesOfAssetPrototypeInstruction */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MutableVideoComposition */

// Creates a mutable video composition configured to apply Core Image filters to each video frame of the specified asset.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMutableVideoComposition/init(asset:applyingCIFiltersWithHandler:)
func (mc _MutableVideoCompositionClass) VideoCompositionWithAssetApplyingCIFiltersWithHandler(asset IAVAsset, applier unsafe.Pointer) IMutableVideoComposition {
	rv := objc.Send[MutableVideoComposition](objc.ID(mc.class), objc.Sel("videoCompositionWithAsset:applyingCIFiltersWithHandler:"), asset, applier)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=VideoCompositionWithAssetApplyingCIFiltersWithHandler) */


// Creates a mutable video composition with the specified asset properties.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMutableVideoComposition/init(propertiesOf:)
func (mc _MutableVideoCompositionClass) VideoCompositionWithPropertiesOfAsset(asset IAVAsset) IMutableVideoComposition {
	rv := objc.Send[MutableVideoComposition](objc.ID(mc.class), objc.Sel("videoCompositionWithPropertiesOfAsset:"), asset)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=VideoCompositionWithPropertiesOfAsset) */


// Creates a mutable video composition with the specified asset properties and a prototype video composition instruction.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMutableVideoComposition/init(propertiesOf:prototypeInstruction:)
func (mc _MutableVideoCompositionClass) VideoCompositionWithPropertiesOfAssetPrototypeInstruction(asset IAVAsset, prototypeInstruction IAVVideoCompositionInstruction) IMutableVideoComposition {
	rv := objc.Send[MutableVideoComposition](objc.ID(mc.class), objc.Sel("videoCompositionWithPropertiesOfAsset:prototypeInstruction:"), asset, prototypeInstruction)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=VideoCompositionWithPropertiesOfAssetPrototypeInstruction) */


// Creates a new mutable video composition.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMutableVideoComposition/videoComposition
func (mc _MutableVideoCompositionClass) VideoComposition() IMutableVideoComposition {
	rv := objc.Send[MutableVideoComposition](objc.ID(mc.class), objc.Sel("videoComposition"))
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=VideoComposition) */


// Returns a new video composition that’s configured to apply Core Image filters to each video frame of the specified asset.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMutableVideoComposition/videoComposition(with:applyingCIFiltersWithHandler:completionHandler:)
func (mc _MutableVideoCompositionClass) VideoCompositionWithAssetApplyingCIFiltersWithHandlerCompletionHandler(asset IAVAsset, applier unsafe.Pointer, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("videoCompositionWithAsset:applyingCIFiltersWithHandler:completionHandler:"), asset, applier, completionHandler)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=VideoCompositionWithAssetApplyingCIFiltersWithHandlerCompletionHandler) */


// Returns a new video composition that’s configured to present the video tracks of the specified asset.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMutableVideoComposition/videoComposition(withPropertiesOf:completionHandler:)
func (mc _MutableVideoCompositionClass) VideoCompositionWithPropertiesOfAssetCompletionHandler(asset IAVAsset, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("videoCompositionWithPropertiesOfAsset:completionHandler:"), asset, completionHandler)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=VideoCompositionWithPropertiesOfAssetCompletionHandler) */


// Returns a new mutable video composition with the specified asset properties and a prototype video composition instruction.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMutableVideoComposition/videoComposition(withPropertiesOf:prototypeInstruction:completionHandler:)
func (mc _MutableVideoCompositionClass) VideoCompositionWithPropertiesOfAssetPrototypeInstructionCompletionHandler(asset IAVAsset, prototypeInstruction IAVVideoCompositionInstruction, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("videoCompositionWithPropertiesOfAsset:prototypeInstruction:completionHandler:"), asset, prototypeInstruction, completionHandler)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=VideoCompositionWithPropertiesOfAssetPrototypeInstructionCompletionHandler) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MutableVideoComposition */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MutableVideoComposition */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MutableVideoComposition */

// A video composition tool to use with Core Animation in offline rendering.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMutableVideoComposition/animationTool
func (m_ MutableVideoComposition) AnimationTool() IAVVideoCompositionCoreAnimationTool {
	rv := objc.Send[VideoCompositionCoreAnimationTool](m_.ID, objc.Sel("animationTool"))
	return rv
}/* debug [instance_properties/getter]: animationTool */


// A video composition tool to use with Core Animation in offline rendering.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMutableVideoComposition/animationTool
func (m_ MutableVideoComposition) SetAnimationTool(value IAVVideoCompositionCoreAnimationTool) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setAnimationTool:"), value)
}/* debug [instance_properties/setter]: animationTool */


// The color primaries used for video composition.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMutableVideoComposition/colorPrimaries
func (m_ MutableVideoComposition) ColorPrimaries() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("colorPrimaries"))
	return rv
}/* debug [instance_properties/getter]: colorPrimaries */


// The color primaries used for video composition.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMutableVideoComposition/colorPrimaries
func (m_ MutableVideoComposition) SetColorPrimaries(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setColorPrimaries:"), value)
}/* debug [instance_properties/setter]: colorPrimaries */


// The transfer function used for video composition.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMutableVideoComposition/colorTransferFunction
func (m_ MutableVideoComposition) ColorTransferFunction() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("colorTransferFunction"))
	return rv
}/* debug [instance_properties/getter]: colorTransferFunction */


// The transfer function used for video composition.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMutableVideoComposition/colorTransferFunction
func (m_ MutableVideoComposition) SetColorTransferFunction(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setColorTransferFunction:"), value)
}/* debug [instance_properties/setter]: colorTransferFunction */


// The YCbCr matrix used for video composition.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMutableVideoComposition/colorYCbCrMatrix
func (m_ MutableVideoComposition) ColorYCbCrMatrix() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("colorYCbCrMatrix"))
	return rv
}/* debug [instance_properties/getter]: colorYCbCrMatrix */


// The YCbCr matrix used for video composition.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMutableVideoComposition/colorYCbCrMatrix
func (m_ MutableVideoComposition) SetColorYCbCrMatrix(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setColorYCbCrMatrix:"), value)
}/* debug [instance_properties/setter]: colorYCbCrMatrix */


// The custom compositor class to use.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMutableVideoComposition/customVideoCompositorClass
func (m_ MutableVideoComposition) CustomVideoCompositorClass() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("customVideoCompositorClass"))
	return rv
}/* debug [instance_properties/getter]: customVideoCompositorClass */


// The custom compositor class to use.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMutableVideoComposition/customVideoCompositorClass
func (m_ MutableVideoComposition) SetCustomVideoCompositorClass(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setCustomVideoCompositorClass:"), value)
}/* debug [instance_properties/setter]: customVideoCompositorClass */


// A time interval for which the video composition should render composed video frames.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMutableVideoComposition/frameDuration
func (m_ MutableVideoComposition) FrameDuration() objc.IObject /* cross-framework: Time */ {
	rv := objc.Send[corevideo.Time](m_.ID, objc.Sel("frameDuration"))
	return rv
}/* debug [instance_properties/getter]: frameDuration */


// A time interval for which the video composition should render composed video frames.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMutableVideoComposition/frameDuration
func (m_ MutableVideoComposition) SetFrameDuration(value objc.IObject /* cross-framework: Time */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setFrameDuration:"), value)
}/* debug [instance_properties/setter]: frameDuration */


// The video composition instructions.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMutableVideoComposition/instructions
func (m_ MutableVideoComposition) Instructions() []objc.ID {
	rv := objc.Send[[]objc.ID](m_.ID, objc.Sel("instructions"))
	return rv
}/* debug [instance_properties/getter]: instructions */


// The video composition instructions.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMutableVideoComposition/instructions
func (m_ MutableVideoComposition) SetInstructions(value []objc.ID) {
	var nsArray objc.ID
	if len(value) > 0 {
		nsArray = objc.ID(objc.GetClass("NSMutableArray")).Send(objc.Sel("arrayWithCapacity:"), len(value))
		for _, item := range value {
			nsArray.Send(objc.Sel("addObject:"), item)
		}
	} else {
		nsArray = objc.ID(objc.GetClass("NSArray")).Send(objc.Sel("array"))
	}
	objc.Send[objc.ID](m_.ID, objc.Sel("setInstructions:"), nsArray)
}/* debug [instance_properties/setter]: instructions */


// The output buffers of the video composition can be specified with the outputBufferDescription. The value is an array of CMTagCollectionRef objects that describes the output buffers.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMutableVideoComposition/outputBufferDescription
func (m_ MutableVideoComposition) OutputBufferDescription() objc.IObject /* cross-framework: NSArray */ {
	rv := objc.Send[foundation.NSArray](m_.ID, objc.Sel("outputBufferDescription"))
	return rv
}/* debug [instance_properties/getter]: outputBufferDescription */


// The output buffers of the video composition can be specified with the outputBufferDescription. The value is an array of CMTagCollectionRef objects that describes the output buffers.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMutableVideoComposition/outputBufferDescription
func (m_ MutableVideoComposition) SetOutputBufferDescription(value objc.IObject /* cross-framework: NSArray */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setOutputBufferDescription:"), value)
}/* debug [instance_properties/setter]: outputBufferDescription */


// Configures the policy for display of HDR display metadata on the rendered frame.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMutableVideoComposition/perFrameHDRDisplayMetadataPolicy
func (m_ MutableVideoComposition) PerFrameHDRDisplayMetadataPolicy() VideoCompositionPerFrameHDRDisplayMetadataPolicy /* typedef */ {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("perFrameHDRDisplayMetadataPolicy"))
	return rv
}/* debug [instance_properties/getter]: perFrameHDRDisplayMetadataPolicy */


// Configures the policy for display of HDR display metadata on the rendered frame.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMutableVideoComposition/perFrameHDRDisplayMetadataPolicy
func (m_ MutableVideoComposition) SetPerFrameHDRDisplayMetadataPolicy(value VideoCompositionPerFrameHDRDisplayMetadataPolicy /* typedef */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setPerFrameHDRDisplayMetadataPolicy:"), value)
}/* debug [instance_properties/setter]: perFrameHDRDisplayMetadataPolicy */


// The scale at which the video composition should render.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMutableVideoComposition/renderScale
func (m_ MutableVideoComposition) RenderScale() float32 {
	rv := objc.Send[float32](m_.ID, objc.Sel("renderScale"))
	return rv
}/* debug [instance_properties/getter]: renderScale */


// The scale at which the video composition should render.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMutableVideoComposition/renderScale
func (m_ MutableVideoComposition) SetRenderScale(value float32) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setRenderScale:"), value)
}/* debug [instance_properties/setter]: renderScale */


// The size at which the video composition should render.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMutableVideoComposition/renderSize
func (m_ MutableVideoComposition) RenderSize() corefoundation.CGSize {
	rv := objc.Send[corefoundation.CGSize](m_.ID, objc.Sel("renderSize"))
	return rv
}/* debug [instance_properties/getter]: renderSize */


// The size at which the video composition should render.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMutableVideoComposition/renderSize
func (m_ MutableVideoComposition) SetRenderSize(value corefoundation.CGSize) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setRenderSize:"), value)
}/* debug [instance_properties/setter]: renderSize */


// The identifiers of source sample data tracks in the composition that the compositor requires to compose frames.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMutableVideoComposition/sourceSampleDataTrackIDs-21o6b
func (m_ MutableVideoComposition) SourceSampleDataTrackIDs() []foundation.Number {
	rv := objc.Send[[]foundation.Number](m_.ID, objc.Sel("sourceSampleDataTrackIDs"))
	return rv
}/* debug [instance_properties/getter]: sourceSampleDataTrackIDs */


// The identifiers of source sample data tracks in the composition that the compositor requires to compose frames.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMutableVideoComposition/sourceSampleDataTrackIDs-21o6b
func (m_ MutableVideoComposition) SetSourceSampleDataTrackIDs(value []foundation.Number) {
	var nsArray objc.ID
	if len(value) > 0 {
		nsArray = objc.ID(objc.GetClass("NSMutableArray")).Send(objc.Sel("arrayWithCapacity:"), len(value))
		for _, item := range value {
			nsArray.Send(objc.Sel("addObject:"), item)
		}
	} else {
		nsArray = objc.ID(objc.GetClass("NSArray")).Send(objc.Sel("array"))
	}
	objc.Send[objc.ID](m_.ID, objc.Sel("setSourceSampleDataTrackIDs:"), nsArray)
}/* debug [instance_properties/setter]: sourceSampleDataTrackIDs */


// An identifier of the source track from which the video composition derives frame timing.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMutableVideoComposition/sourceTrackIDForFrameTiming
func (m_ MutableVideoComposition) SourceTrackIDForFrameTiming() PersistentTrackID /* not a class type */ {
	rv := objc.Send[PersistentTrackID](m_.ID, objc.Sel("sourceTrackIDForFrameTiming"))
	return rv
}/* debug [instance_properties/getter]: sourceTrackIDForFrameTiming */


// An identifier of the source track from which the video composition derives frame timing.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMutableVideoComposition/sourceTrackIDForFrameTiming
func (m_ MutableVideoComposition) SetSourceTrackIDForFrameTiming(value PersistentTrackID /* not a class type */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setSourceTrackIDForFrameTiming:"), value)
}/* debug [instance_properties/setter]: sourceTrackIDForFrameTiming */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class AVMutableVideoComposition */


