// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)





// The class instance for the [ImageHistogramEqualization] class.
var (
	ImageHistogramEqualizationClass     _ImageHistogramEqualizationClass
	ImageHistogramEqualizationClassOnce sync.Once
)

func getImageHistogramEqualizationClass() _ImageHistogramEqualizationClass {
	ImageHistogramEqualizationClassOnce.Do(func() {
		ImageHistogramEqualizationClass = _ImageHistogramEqualizationClass{objc.GetClass("MPSImageHistogramEqualization")}
	})
	return ImageHistogramEqualizationClass
}

type _ImageHistogramEqualizationClass struct {
	class objc.Class
}





// An interface definition for the [ImageHistogramEqualization] class.
type IImageHistogramEqualization interface {
	IUnaryImageKernel
	

	// properties:
	HistogramInfo() ImageHistogramInfo get /* not a class type */
	SetHistogramInfo(value ImageHistogramInfo get /* not a class type */)


	

	// methods:
	EncodeTransform()
	EncodeTransformToCommandBufferSourceTextureHistogramHistogramOffset(commandBuffer unsafe.Pointer, source unsafe.Pointer, histogram unsafe.Pointer, histogramOffset uint)


}





// Alloc allocates a new instance without initialization.
func (ic _ImageHistogramEqualizationClass) Alloc() ImageHistogramEqualization {
	rv := objc.Send[ImageHistogramEqualization](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (ic _ImageHistogramEqualizationClass) New() ImageHistogramEqualization {
	rv := objc.Send[ImageHistogramEqualization](objc.ID(ic.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (i_ ImageHistogramEqualization) Init() ImageHistogramEqualization {
	rv := objc.Send[ImageHistogramEqualization](i_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i_ ImageHistogramEqualization) Autorelease() ImageHistogramEqualization {
	rv := objc.Send[ImageHistogramEqualization](i_.ID, objc.Sel("autorelease"))
	return rv
}

// NewImageHistogramEqualization creates a new ImageHistogramEqualization instance.
func NewImageHistogramEqualization() ImageHistogramEqualization {
	return getImageHistogramEqualizationClass().New()
}





// A filter that equalizes the histogram of an image.
//
// The process is divided into three steps: Call the method to create a object. Call the method. This creates a privately held image transform (i.e. a cumulative distribution function of the histogram) which will be used to equalize the distribution of the histogram of the source image. This process runs on a command buffer when it is committed to a command queue. It must complete before the next step can be run. It may be performed on the same command buffer. The argument specifies the histogram buffer which contains the histogram values for the source texture. The argument is used by the method to determine the number of channels and therefore which histogram data in the histogram buffer to use. The histogram for the source texture must have been computed either on the CPU or using the kernel. Call the method to read data from the source texture, apply the equalization transform to it, and write to the destination texture. This step is also done on the GPU on a command queue.


// A filter that equalizes the histogram of an image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImageHistogramEqualization
type ImageHistogramEqualization struct {
	UnaryImageKernel
}

// ImageHistogramEqualizationFrom constructs a [ImageHistogramEqualization] from an unsafe.Pointer.
//
// A filter that equalizes the histogram of an image.
func ImageHistogramEqualizationFrom(ptr unsafe.Pointer) ImageHistogramEqualization {
	return ImageHistogramEqualization{
		UnaryImageKernel: UnaryImageKernelFrom(ptr),
	}
}






// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagehistogramequalization/2866993-initwithcoder
func NewImageHistogramEqualizationWithCoderDevice(aDecoder foundation.Coder, device unsafe.Pointer) ImageHistogramEqualization {
	instance := getImageHistogramEqualizationClass().Alloc()
	rv := objc.Send[ImageHistogramEqualization](instance.ID, objc.Sel("initWithCoder:device:"), aDecoder, device)
	rv.Autorelease()
	return rv
}


// Initializes a histogram with specific information.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagehistogramequalization/1618856-initwithdevice
func NewImageHistogramEqualizationWithDeviceHistogramInfo(device unsafe.Pointer, histogramInfo objc.IObject /* cross-framework: MPSImageHistogramInfo */) ImageHistogramEqualization {
	instance := getImageHistogramEqualizationClass().Alloc()
	rv := objc.Send[ImageHistogramEqualization](instance.ID, objc.Sel("initWithDevice:histogramInfo:"), device, histogramInfo)
	rv.Autorelease()
	return rv
}

















// Encodes the transform function to a command buffer using a compute command encoder. The transform function computes the equalization lookup table.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagehistogramequalization/1618746-encodetransform
func (i_ ImageHistogramEqualization) EncodeTransform() {
	objc.Send[objc.ID](i_.ID, objc.Sel("encodeTransform"))
}


// Encodes the transform function to a command buffer using a compute command encoder. The transform function computes the equalization lookup table.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagehistogramequalization/1618746-encodetransformtocommandbuffer
func (i_ ImageHistogramEqualization) EncodeTransformToCommandBufferSourceTextureHistogramHistogramOffset(commandBuffer unsafe.Pointer, source unsafe.Pointer, histogram unsafe.Pointer, histogramOffset uint) {
	objc.Send[objc.ID](i_.ID, objc.Sel("encodeTransformToCommandBuffer:sourceTexture:histogram:histogramOffset:"), commandBuffer, source, histogram, histogramOffset)
}







// A structure describing the histogram content.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagehistogramequalization/1618775-histograminfo
func (i_ ImageHistogramEqualization) HistogramInfo() ImageHistogramInfo get /* not a class type */ {
	rv := objc.Send[objc.ID](i_.ID, objc.Sel("histogramInfo"))
	return rv
}


// A structure describing the histogram content.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagehistogramequalization/1618775-histograminfo
func (i_ ImageHistogramEqualization) SetHistogramInfo(value ImageHistogramInfo get /* not a class type */) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setHistogramInfo:"), value)
}







