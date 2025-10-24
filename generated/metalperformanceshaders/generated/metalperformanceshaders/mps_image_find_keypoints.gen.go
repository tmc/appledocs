// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MPSImageFindKeypoints */


/* debug [class_header]: Header for MPSImageFindKeypoints */
// The class instance for the [ImageFindKeypoints] class.
var (
	ImageFindKeypointsClass     _ImageFindKeypointsClass
	ImageFindKeypointsClassOnce sync.Once
)

func getImageFindKeypointsClass() _ImageFindKeypointsClass {
	ImageFindKeypointsClassOnce.Do(func() {
		ImageFindKeypointsClass = _ImageFindKeypointsClass{objc.GetClass("MPSImageFindKeypoints")}
	})
	return ImageFindKeypointsClass
}

type _ImageFindKeypointsClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for ImageFindKeypoints */
// An interface definition for the [ImageFindKeypoints] class.
type IImageFindKeypoints interface {
	IKernel
	
/* debug [class_interface_properties]: Properties for ImageFindKeypoints */
	// properties:
	KeypointRangeInfo() ImageKeypointRangeInfo get /* not a class type */
	SetKeypointRangeInfo(value ImageKeypointRangeInfo get /* not a class type */)
	MinimumThresholdValue() float32
	SetMinimumThresholdValue(value float32)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for ImageFindKeypoints */
	// methods:
	Encode()
	EncodeToCommandBufferSourceTextureRegionsNumberOfRegionsKeypointCountBufferKeypointCountBufferOffsetKeypointDataBufferKeypointDataBufferOffset(commandBuffer unsafe.Pointer, source unsafe.Pointer, regions Region /* not a class type */, numberOfRegions uint, keypointCountBuffer unsafe.Pointer, keypointCountBufferOffset uint, keypointDataBuffer unsafe.Pointer, keypointDataBufferOffset uint)
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for ImageFindKeypoints */
// Alloc allocates a new instance without initialization.
func (ic _ImageFindKeypointsClass) Alloc() ImageFindKeypoints {
	rv := objc.Send[ImageFindKeypoints](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (ic _ImageFindKeypointsClass) New() ImageFindKeypoints {
	rv := objc.Send[ImageFindKeypoints](objc.ID(ic.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (i_ ImageFindKeypoints) Init() ImageFindKeypoints {
	rv := objc.Send[ImageFindKeypoints](i_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i_ ImageFindKeypoints) Autorelease() ImageFindKeypoints {
	rv := objc.Send[ImageFindKeypoints](i_.ID, objc.Sel("autorelease"))
	return rv
}

// NewImageFindKeypoints creates a new ImageFindKeypoints instance.
func NewImageFindKeypoints() ImageFindKeypoints {
	return getImageFindKeypointsClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for ImageFindKeypoints */
// A kernel that is used to find a list of keypoints.
//
// This kernel is used to find a list of keypoints whose values are greater than the in . The keypoints are generated for a specified region in the image. The pixel format of the source image must be .


// A kernel that is used to find a list of keypoints.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImageFindKeypoints
type ImageFindKeypoints struct {
	Kernel
}

// ImageFindKeypointsFrom constructs a [ImageFindKeypoints] from an unsafe.Pointer.
//
// A kernel that is used to find a list of keypoints.
func ImageFindKeypointsFrom(ptr unsafe.Pointer) ImageFindKeypoints {
	return ImageFindKeypoints{
		Kernel: KernelFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for ImageFindKeypoints */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagefindkeypoints/2873306-initwithcoder
func NewImageFindKeypointsWithCoderDevice(aDecoder Coder /* not a class type */, device unsafe.Pointer) ImageFindKeypoints {
	instance := getImageFindKeypointsClass().Alloc()
	rv := objc.Send[ImageFindKeypoints](instance.ID, objc.Sel("initWithCoder:device:"), aDecoder, device)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewImageFindKeypointsWithCoderDevice */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagefindkeypoints/2873309-initwithdevice
func NewImageFindKeypointsWithDeviceInfo(device unsafe.Pointer, info objc.IObject /* cross-framework: MPSImageKeypointRangeInfo */) ImageFindKeypoints {
	instance := getImageFindKeypointsClass().Alloc()
	rv := objc.Send[ImageFindKeypoints](instance.ID, objc.Sel("initWithDevice:info:"), device, info)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewImageFindKeypointsWithDeviceInfo */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for ImageFindKeypoints */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for ImageFindKeypoints */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for ImageFindKeypoints */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagefindkeypoints/2873307-encode
func (i_ ImageFindKeypoints) Encode() {
	objc.Send[objc.ID](i_.ID, objc.Sel("encode"))
}/* debug [instance_methods/method]: Encode */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagefindkeypoints/2873307-encodetocommandbuffer
func (i_ ImageFindKeypoints) EncodeToCommandBufferSourceTextureRegionsNumberOfRegionsKeypointCountBufferKeypointCountBufferOffsetKeypointDataBufferKeypointDataBufferOffset(commandBuffer unsafe.Pointer, source unsafe.Pointer, regions Region /* not a class type */, numberOfRegions uint, keypointCountBuffer unsafe.Pointer, keypointCountBufferOffset uint, keypointDataBuffer unsafe.Pointer, keypointDataBufferOffset uint) {
	objc.Send[objc.ID](i_.ID, objc.Sel("encodeToCommandBuffer:sourceTexture:regions:numberOfRegions:keypointCountBuffer:keypointCountBufferOffset:keypointDataBuffer:keypointDataBufferOffset:"), commandBuffer, source, regions, numberOfRegions, keypointCountBuffer, keypointCountBufferOffset, keypointDataBuffer, keypointDataBufferOffset)
}/* debug [instance_methods/method]: EncodeToCommandBufferSourceTextureRegionsNumberOfRegionsKeypointCountBufferKeypointCountBufferOffsetKeypointDataBufferKeypointDataBufferOffset */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for ImageFindKeypoints */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagefindkeypoints/2873310-keypointrangeinfo
func (i_ ImageFindKeypoints) KeypointRangeInfo() ImageKeypointRangeInfo get /* not a class type */ {
	rv := objc.Send[objc.ID](i_.ID, objc.Sel("keypointRangeInfo"))
	return rv
}/* debug [instance_properties/getter]: keypointRangeInfo */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagefindkeypoints/2873310-keypointrangeinfo
func (i_ ImageFindKeypoints) SetKeypointRangeInfo(value ImageKeypointRangeInfo get /* not a class type */) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setKeypointRangeInfo:"), value)
}/* debug [instance_properties/setter]: keypointRangeInfo */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagekeypointrangeinfo/minimumthresholdvalue
func (i_ ImageFindKeypoints) MinimumThresholdValue() float32 {
	rv := objc.Send[float32](i_.ID, objc.Sel("minimumThresholdValue"))
	return rv
}/* debug [instance_properties/getter]: minimumThresholdValue */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagekeypointrangeinfo/minimumthresholdvalue
func (i_ ImageFindKeypoints) SetMinimumThresholdValue(value float32) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setMinimumThresholdValue:"), value)
}/* debug [instance_properties/setter]: minimumThresholdValue */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MPSImageFindKeypoints */


