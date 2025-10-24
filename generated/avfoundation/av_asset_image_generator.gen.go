// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/corevideo"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class AVAssetImageGenerator */


/* debug [class_header]: Header for AVAssetImageGenerator */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for AssetImageGenerator */
// An interface definition for the [AssetImageGenerator] class.
type IAssetImageGenerator interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for AssetImageGenerator */
	// properties:
	ApertureMode() AssetImageGeneratorApertureMode /* typedef */
	SetApertureMode(value AssetImageGeneratorApertureMode /* typedef */)
	AppliesPreferredTrackTransform() bool
	SetAppliesPreferredTrackTransform(value bool)
	Asset() IAVAsset
	CustomVideoCompositor() unsafe.Pointer
	DynamicRangePolicy() AssetImageGeneratorDynamicRangePolicy /* typedef */
	SetDynamicRangePolicy(value AssetImageGeneratorDynamicRangePolicy /* typedef */)
	MaximumSize() corefoundation.CGSize
	SetMaximumSize(value corefoundation.CGSize)
	RequestedTimeToleranceAfter() objc.IObject /* cross-framework: Time */
	SetRequestedTimeToleranceAfter(value objc.IObject /* cross-framework: Time */)
	RequestedTimeToleranceBefore() objc.IObject /* cross-framework: Time */
	SetRequestedTimeToleranceBefore(value objc.IObject /* cross-framework: Time */)
	VideoComposition() IAVVideoComposition
	SetVideoComposition(value IAVVideoComposition)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for AssetImageGenerator */
	// methods:
	CancelAllCGImageGeneration()
	GenerateCGImageAsynchronouslyForTimeCompletionHandler(requestedTime objc.IObject /* cross-framework: Time */, handler unsafe.Pointer)
	GenerateCGImagesAsynchronouslyForTimesCompletionHandler(requestedTimes []foundation.Value, handler AssetImageGeneratorCompletionHandler /* not a class type */)
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for AssetImageGenerator */
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for AssetImageGenerator */
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for AssetImageGenerator */

// Creates an object that generates images for times within a video asset.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetImageGenerator/init(asset:)
func NewAssetImageGeneratorWithAsset(asset IAVAsset) AssetImageGenerator {
	instance := getAssetImageGeneratorClass().Alloc()
	rv := objc.Send[AssetImageGenerator](instance.ID, objc.Sel("initWithAsset:"), asset)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewAssetImageGeneratorWithAsset */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for AssetImageGenerator */

// Returns a new object that generates images for times within a video asset.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetImageGenerator/assetImageGeneratorWithAsset:
func (ac _AssetImageGeneratorClass) AssetImageGeneratorWithAsset(asset IAVAsset) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(ac.class), objc.Sel("assetImageGeneratorWithAsset:"), asset)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=AssetImageGeneratorWithAsset) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for AssetImageGenerator */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for AssetImageGenerator */

// Cancels all pending image generation requests.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetImageGenerator/cancelAllCGImageGeneration()
func (a_ AssetImageGenerator) CancelAllCGImageGeneration() {
	objc.Send[objc.ID](a_.ID, objc.Sel("cancelAllCGImageGeneration"))
}/* debug [instance_methods/method]: CancelAllCGImageGeneration */


// Generates an image asynchronously for a requested time, and returns the result in a callback.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetImageGenerator/generateCGImageAsynchronously(for:completionHandler:)
func (a_ AssetImageGenerator) GenerateCGImageAsynchronouslyForTimeCompletionHandler(requestedTime objc.IObject /* cross-framework: Time */, handler unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("generateCGImageAsynchronouslyForTime:completionHandler:"), requestedTime, handler)
}/* debug [instance_methods/method]: GenerateCGImageAsynchronouslyForTimeCompletionHandler */


// Generates images asynchronously for an array of requested times, and returns the results in a callback.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetImageGenerator/generateCGImagesAsynchronously(forTimes:completionHandler:)
func (a_ AssetImageGenerator) GenerateCGImagesAsynchronouslyForTimesCompletionHandler(requestedTimes []foundation.Value, handler AssetImageGeneratorCompletionHandler /* not a class type */) {
	objc.Send[objc.ID](a_.ID, objc.Sel("generateCGImagesAsynchronouslyForTimes:completionHandler:"), requestedTimes, handler)
}/* debug [instance_methods/method]: GenerateCGImagesAsynchronouslyForTimesCompletionHandler */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for AssetImageGenerator */

// Specifies the aperture mode for the generated image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetImageGenerator/apertureMode-swift.property
func (a_ AssetImageGenerator) ApertureMode() AssetImageGeneratorApertureMode /* typedef */ {
	rv := objc.Send[foundation.NSString](a_.ID, objc.Sel("apertureMode"))
	return rv
}/* debug [instance_properties/getter]: apertureMode */


// Specifies the aperture mode for the generated image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetImageGenerator/apertureMode-swift.property
func (a_ AssetImageGenerator) SetApertureMode(value AssetImageGeneratorApertureMode /* typedef */) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setApertureMode:"), value)
}/* debug [instance_properties/setter]: apertureMode */


// A Boolean value that specifies whether to apply the track matrix or matrices when generating an image from the asset.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetImageGenerator/appliesPreferredTrackTransform
func (a_ AssetImageGenerator) AppliesPreferredTrackTransform() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("appliesPreferredTrackTransform"))
	return rv
}/* debug [instance_properties/getter]: appliesPreferredTrackTransform */


// A Boolean value that specifies whether to apply the track matrix or matrices when generating an image from the asset.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetImageGenerator/appliesPreferredTrackTransform
func (a_ AssetImageGenerator) SetAppliesPreferredTrackTransform(value bool) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setAppliesPreferredTrackTransform:"), value)
}/* debug [instance_properties/setter]: appliesPreferredTrackTransform */


// The asset that initialized the image generator.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetImageGenerator/asset
func (a_ AssetImageGenerator) Asset() IAVAsset {
	rv := objc.Send[Asset](a_.ID, objc.Sel("asset"))
	return rv
}/* debug [instance_properties/getter]: asset */


// A custom video compositor to use when extracting images from assets with multiple video tracks.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetImageGenerator/customVideoCompositor
func (a_ AssetImageGenerator) CustomVideoCompositor() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("customVideoCompositor"))
	return rv
}/* debug [instance_properties/getter]: customVideoCompositor */


// The dynamic range policy to use when generating images.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetImageGenerator/dynamicRangePolicy-swift.property
func (a_ AssetImageGenerator) DynamicRangePolicy() AssetImageGeneratorDynamicRangePolicy /* typedef */ {
	rv := objc.Send[foundation.NSString](a_.ID, objc.Sel("dynamicRangePolicy"))
	return rv
}/* debug [instance_properties/getter]: dynamicRangePolicy */


// The dynamic range policy to use when generating images.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetImageGenerator/dynamicRangePolicy-swift.property
func (a_ AssetImageGenerator) SetDynamicRangePolicy(value AssetImageGeneratorDynamicRangePolicy /* typedef */) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setDynamicRangePolicy:"), value)
}/* debug [instance_properties/setter]: dynamicRangePolicy */


// The maximum size of images to generate.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetImageGenerator/maximumSize
func (a_ AssetImageGenerator) MaximumSize() corefoundation.CGSize {
	rv := objc.Send[corefoundation.CGSize](a_.ID, objc.Sel("maximumSize"))
	return rv
}/* debug [instance_properties/getter]: maximumSize */


// The maximum size of images to generate.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetImageGenerator/maximumSize
func (a_ AssetImageGenerator) SetMaximumSize(value corefoundation.CGSize) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setMaximumSize:"), value)
}/* debug [instance_properties/setter]: maximumSize */


// A maximum length of time after the requested time to allow image generation to occur.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetImageGenerator/requestedTimeToleranceAfter
func (a_ AssetImageGenerator) RequestedTimeToleranceAfter() objc.IObject /* cross-framework: Time */ {
	rv := objc.Send[corevideo.Time](a_.ID, objc.Sel("requestedTimeToleranceAfter"))
	return rv
}/* debug [instance_properties/getter]: requestedTimeToleranceAfter */


// A maximum length of time after the requested time to allow image generation to occur.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetImageGenerator/requestedTimeToleranceAfter
func (a_ AssetImageGenerator) SetRequestedTimeToleranceAfter(value objc.IObject /* cross-framework: Time */) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setRequestedTimeToleranceAfter:"), value)
}/* debug [instance_properties/setter]: requestedTimeToleranceAfter */


// A maximum length of time before the requested time to allow image generation to occur.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetImageGenerator/requestedTimeToleranceBefore
func (a_ AssetImageGenerator) RequestedTimeToleranceBefore() objc.IObject /* cross-framework: Time */ {
	rv := objc.Send[corevideo.Time](a_.ID, objc.Sel("requestedTimeToleranceBefore"))
	return rv
}/* debug [instance_properties/getter]: requestedTimeToleranceBefore */


// A maximum length of time before the requested time to allow image generation to occur.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetImageGenerator/requestedTimeToleranceBefore
func (a_ AssetImageGenerator) SetRequestedTimeToleranceBefore(value objc.IObject /* cross-framework: Time */) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setRequestedTimeToleranceBefore:"), value)
}/* debug [instance_properties/setter]: requestedTimeToleranceBefore */


// A video composition to use when extracting images from assets with multiple video tracks.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetImageGenerator/videoComposition
func (a_ AssetImageGenerator) VideoComposition() IAVVideoComposition {
	rv := objc.Send[VideoComposition](a_.ID, objc.Sel("videoComposition"))
	return rv
}/* debug [instance_properties/getter]: videoComposition */


// A video composition to use when extracting images from assets with multiple video tracks.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetImageGenerator/videoComposition
func (a_ AssetImageGenerator) SetVideoComposition(value IAVVideoComposition) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setVideoComposition:"), value)
}/* debug [instance_properties/setter]: videoComposition */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class AVAssetImageGenerator */


