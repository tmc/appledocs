// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)





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





// An interface definition for the [ImageFindKeypoints] class.
type IImageFindKeypoints interface {
	IKernel
	

	// properties:
	KeypointRangeInfo() ImageKeypointRangeInfo get /* not a class type */
	SetKeypointRangeInfo(value ImageKeypointRangeInfo get /* not a class type */)
	MinimumThresholdValue() float32
	SetMinimumThresholdValue(value float32)


	

	// methods:
	Encode()
	EncodeToCommandBufferSourceTextureRegionsNumberOfRegionsKeypointCountBufferKeypointCountBufferOffsetKeypointDataBufferKeypointDataBufferOffset(commandBuffer unsafe.Pointer, source unsafe.Pointer, regions metal.IMTLRegion, numberOfRegions uint, keypointCountBuffer unsafe.Pointer, keypointCountBufferOffset uint, keypointDataBuffer unsafe.Pointer, keypointDataBufferOffset uint)


}





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






// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagefindkeypoints/2873306-initwithcoder
func NewImageFindKeypointsWithCoderDevice(aDecoder foundation.Coder, device unsafe.Pointer) ImageFindKeypoints {
	instance := getImageFindKeypointsClass().Alloc()
	rv := objc.Send[ImageFindKeypoints](instance.ID, objc.Sel("initWithCoder:device:"), aDecoder, device)
	rv.Autorelease()
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagefindkeypoints/2873309-initwithdevice
func NewImageFindKeypointsWithDeviceInfo(device unsafe.Pointer, info ImageKeypointRangeInfo) ImageFindKeypoints {
	instance := getImageFindKeypointsClass().Alloc()
	rv := objc.Send[ImageFindKeypoints](instance.ID, objc.Sel("initWithDevice:info:"), device, info)
	rv.Autorelease()
	return rv
}

















// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagefindkeypoints/2873307-encode
func (i_ ImageFindKeypoints) Encode() {
	objc.Send[objc.ID](i_.ID, objc.Sel("encode"))
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagefindkeypoints/2873307-encodetocommandbuffer
func (i_ ImageFindKeypoints) EncodeToCommandBufferSourceTextureRegionsNumberOfRegionsKeypointCountBufferKeypointCountBufferOffsetKeypointDataBufferKeypointDataBufferOffset(commandBuffer unsafe.Pointer, source unsafe.Pointer, regions metal.IMTLRegion, numberOfRegions uint, keypointCountBuffer unsafe.Pointer, keypointCountBufferOffset uint, keypointDataBuffer unsafe.Pointer, keypointDataBufferOffset uint) {
	objc.Send[objc.ID](i_.ID, objc.Sel("encodeToCommandBuffer:sourceTexture:regions:numberOfRegions:keypointCountBuffer:keypointCountBufferOffset:keypointDataBuffer:keypointDataBufferOffset:"), commandBuffer, source, regions, numberOfRegions, keypointCountBuffer, keypointCountBufferOffset, keypointDataBuffer, keypointDataBufferOffset)
}







// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagefindkeypoints/2873310-keypointrangeinfo
func (i_ ImageFindKeypoints) KeypointRangeInfo() ImageKeypointRangeInfo get /* not a class type */ {
	rv := objc.Send[objc.ID](i_.ID, objc.Sel("keypointRangeInfo"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagefindkeypoints/2873310-keypointrangeinfo
func (i_ ImageFindKeypoints) SetKeypointRangeInfo(value ImageKeypointRangeInfo get /* not a class type */) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setKeypointRangeInfo:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagekeypointrangeinfo/minimumthresholdvalue
func (i_ ImageFindKeypoints) MinimumThresholdValue() float32 {
	rv := objc.Send[float32](i_.ID, objc.Sel("minimumThresholdValue"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagekeypointrangeinfo/minimumthresholdvalue
func (i_ ImageFindKeypoints) SetMinimumThresholdValue(value float32) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setMinimumThresholdValue:"), value)
}







