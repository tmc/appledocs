// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/corefoundation"
	"github.com/tmc/appledocs/generated/objectivec"
)





// The class instance for the [AssetImageGenerator] class.
var (
	AssetImageGeneratorClass     _AssetImageGeneratorClass
	AssetImageGeneratorClassOnce sync.Once
)

func getAssetImageGeneratorClass() _AssetImageGeneratorClass {
	AssetImageGeneratorClassOnce.Do(func() {
		AssetImageGeneratorClass = _AssetImageGeneratorClass{objc.GetClass("AVAssetImageGenerator")}
	})
	return AssetImageGeneratorClass
}

type _AssetImageGeneratorClass struct {
	class objc.Class
}





// An interface definition for the [AssetImageGenerator] class.
type IAssetImageGenerator interface {
	objectivec.IObject
	

	// properties:
	ApertureMode() AssetImageGeneratorApertureMode
	SetApertureMode(value AssetImageGeneratorApertureMode)
	AppliesPreferredTrackTransform() bool
	SetAppliesPreferredTrackTransform(value bool)
	Asset() IAVAsset
	CustomVideoCompositor() unsafe.Pointer
	DynamicRangePolicy() AssetImageGeneratorDynamicRangePolicy
	SetDynamicRangePolicy(value AssetImageGeneratorDynamicRangePolicy)
	MaximumSize() corefoundation.CGSize
	SetMaximumSize(value corefoundation.CGSize)
	RequestedTimeToleranceAfter() objectivec.IObject
	SetRequestedTimeToleranceAfter(value objectivec.IObject)
	RequestedTimeToleranceBefore() objectivec.IObject
	SetRequestedTimeToleranceBefore(value objectivec.IObject)
	VideoComposition() IAVVideoComposition
	SetVideoComposition(value IAVVideoComposition)


	

	// methods:
	CancelAllCGImageGeneration()
	GenerateCGImageAsynchronouslyForTimeCompletionHandler(requestedTime objectivec.IObject, handler unsafe.Pointer)
	GenerateCGImagesAsynchronouslyForTimesCompletionHandler(requestedTimes []foundation.Value, handler AssetImageGeneratorCompletionHandler /* not a class type */)


}





// Alloc allocates a new instance without initialization.
func (ac _AssetImageGeneratorClass) Alloc() AssetImageGenerator {
	rv := objc.Send[AssetImageGenerator](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (ac _AssetImageGeneratorClass) New() AssetImageGenerator {
	rv := objc.Send[AssetImageGenerator](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ AssetImageGenerator) Init() AssetImageGenerator {
	rv := objc.Send[AssetImageGenerator](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ AssetImageGenerator) Autorelease() AssetImageGenerator {
	rv := objc.Send[AssetImageGenerator](a_.ID, objc.Sel("autorelease"))
	return rv
}

// NewAssetImageGenerator creates a new AssetImageGenerator instance.
func NewAssetImageGenerator() AssetImageGenerator {
	return getAssetImageGeneratorClass().New()
}





// An object that generates images from a video asset.
//
// Use an image generator to extract images from a video asset at particular times within its timeline.


// An object that generates images from a video asset.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetImageGenerator
type AssetImageGenerator struct {
	objectivec.Object
}

// AssetImageGeneratorFrom constructs a [AssetImageGenerator] from an unsafe.Pointer.
//
// An object that generates images from a video asset.
func AssetImageGeneratorFrom(ptr unsafe.Pointer) AssetImageGenerator {
	return AssetImageGenerator{objectivec.Object{objc.ID(ptr)}}
}






// Creates an object that generates images for times within a video asset.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetImageGenerator/init(asset:)
func NewAssetImageGeneratorWithAsset(asset IAVAsset) AssetImageGenerator {
	instance := getAssetImageGeneratorClass().Alloc()
	rv := objc.Send[AssetImageGenerator](instance.ID, objc.Sel("initWithAsset:"), asset)
	rv.Autorelease()
	return rv
}







// Returns a new object that generates images for times within a video asset.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetImageGenerator/assetImageGeneratorWithAsset:
func (ac _AssetImageGeneratorClass) AssetImageGeneratorWithAsset(asset IAVAsset) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(ac.class), objc.Sel("assetImageGeneratorWithAsset:"), asset)
	return rv
}












// Cancels all pending image generation requests.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetImageGenerator/cancelAllCGImageGeneration()
func (a_ AssetImageGenerator) CancelAllCGImageGeneration() {
	objc.Send[objc.ID](a_.ID, objc.Sel("cancelAllCGImageGeneration"))
}


// Generates an image asynchronously for a requested time, and returns the result in a callback.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetImageGenerator/generateCGImageAsynchronously(for:completionHandler:)
func (a_ AssetImageGenerator) GenerateCGImageAsynchronouslyForTimeCompletionHandler(requestedTime objectivec.IObject, handler unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("generateCGImageAsynchronouslyForTime:completionHandler:"), requestedTime, handler)
}


// Generates images asynchronously for an array of requested times, and returns the results in a callback.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetImageGenerator/generateCGImagesAsynchronously(forTimes:completionHandler:)
func (a_ AssetImageGenerator) GenerateCGImagesAsynchronouslyForTimesCompletionHandler(requestedTimes []foundation.Value, handler AssetImageGeneratorCompletionHandler /* not a class type */) {
	objc.Send[objc.ID](a_.ID, objc.Sel("generateCGImagesAsynchronouslyForTimes:completionHandler:"), requestedTimes, handler)
}







// Specifies the aperture mode for the generated image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetImageGenerator/apertureMode-swift.property
func (a_ AssetImageGenerator) ApertureMode() AssetImageGeneratorApertureMode {
	rv := objc.Send[AssetImageGeneratorApertureMode](a_.ID, objc.Sel("apertureMode"))
	return rv
}


// Specifies the aperture mode for the generated image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetImageGenerator/apertureMode-swift.property
func (a_ AssetImageGenerator) SetApertureMode(value AssetImageGeneratorApertureMode) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setApertureMode:"), value)
}


// A Boolean value that specifies whether to apply the track matrix or matrices when generating an image from the asset.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetImageGenerator/appliesPreferredTrackTransform
func (a_ AssetImageGenerator) AppliesPreferredTrackTransform() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("appliesPreferredTrackTransform"))
	return rv
}


// A Boolean value that specifies whether to apply the track matrix or matrices when generating an image from the asset.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetImageGenerator/appliesPreferredTrackTransform
func (a_ AssetImageGenerator) SetAppliesPreferredTrackTransform(value bool) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setAppliesPreferredTrackTransform:"), value)
}


// The asset that initialized the image generator.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetImageGenerator/asset
func (a_ AssetImageGenerator) Asset() IAVAsset {
	rv := objc.Send[Asset](a_.ID, objc.Sel("asset"))
	return rv
}


// A custom video compositor to use when extracting images from assets with multiple video tracks.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetImageGenerator/customVideoCompositor
func (a_ AssetImageGenerator) CustomVideoCompositor() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("customVideoCompositor"))
	return rv
}


// The dynamic range policy to use when generating images.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetImageGenerator/dynamicRangePolicy-swift.property
func (a_ AssetImageGenerator) DynamicRangePolicy() AssetImageGeneratorDynamicRangePolicy {
	rv := objc.Send[AssetImageGeneratorDynamicRangePolicy](a_.ID, objc.Sel("dynamicRangePolicy"))
	return rv
}


// The dynamic range policy to use when generating images.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetImageGenerator/dynamicRangePolicy-swift.property
func (a_ AssetImageGenerator) SetDynamicRangePolicy(value AssetImageGeneratorDynamicRangePolicy) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setDynamicRangePolicy:"), value)
}


// The maximum size of images to generate.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetImageGenerator/maximumSize
func (a_ AssetImageGenerator) MaximumSize() corefoundation.CGSize {
	rv := objc.Send[corefoundation.CGSize](a_.ID, objc.Sel("maximumSize"))
	return rv
}


// The maximum size of images to generate.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetImageGenerator/maximumSize
func (a_ AssetImageGenerator) SetMaximumSize(value corefoundation.CGSize) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setMaximumSize:"), value)
}


// A maximum length of time after the requested time to allow image generation to occur.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetImageGenerator/requestedTimeToleranceAfter
func (a_ AssetImageGenerator) RequestedTimeToleranceAfter() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](a_.ID, objc.Sel("requestedTimeToleranceAfter"))
	return rv
}


// A maximum length of time after the requested time to allow image generation to occur.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetImageGenerator/requestedTimeToleranceAfter
func (a_ AssetImageGenerator) SetRequestedTimeToleranceAfter(value objectivec.IObject) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setRequestedTimeToleranceAfter:"), value)
}


// A maximum length of time before the requested time to allow image generation to occur.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetImageGenerator/requestedTimeToleranceBefore
func (a_ AssetImageGenerator) RequestedTimeToleranceBefore() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](a_.ID, objc.Sel("requestedTimeToleranceBefore"))
	return rv
}


// A maximum length of time before the requested time to allow image generation to occur.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetImageGenerator/requestedTimeToleranceBefore
func (a_ AssetImageGenerator) SetRequestedTimeToleranceBefore(value objectivec.IObject) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setRequestedTimeToleranceBefore:"), value)
}


// A video composition to use when extracting images from assets with multiple video tracks.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetImageGenerator/videoComposition
func (a_ AssetImageGenerator) VideoComposition() IAVVideoComposition {
	rv := objc.Send[VideoComposition](a_.ID, objc.Sel("videoComposition"))
	return rv
}


// A video composition to use when extracting images from assets with multiple video tracks.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetImageGenerator/videoComposition
func (a_ AssetImageGenerator) SetVideoComposition(value IAVVideoComposition) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setVideoComposition:"), value)
}







