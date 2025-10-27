// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)





// The class instance for the [ImageHistogramSpecification] class.
var (
	ImageHistogramSpecificationClass     _ImageHistogramSpecificationClass
	ImageHistogramSpecificationClassOnce sync.Once
)

func getImageHistogramSpecificationClass() _ImageHistogramSpecificationClass {
	ImageHistogramSpecificationClassOnce.Do(func() {
		ImageHistogramSpecificationClass = _ImageHistogramSpecificationClass{objc.GetClass("MPSImageHistogramSpecification")}
	})
	return ImageHistogramSpecificationClass
}

type _ImageHistogramSpecificationClass struct {
	class objc.Class
}





// An interface definition for the [ImageHistogramSpecification] class.
type IImageHistogramSpecification interface {
	IUnaryImageKernel
	

	// properties:
	HistogramInfo() ImageHistogramInfo get /* not a class type */
	SetHistogramInfo(value ImageHistogramInfo get /* not a class type */)


	

	// methods:
	EncodeTransform()
	EncodeTransformToCommandBufferSourceTextureSourceHistogramSourceHistogramOffsetDesiredHistogramDesiredHistogramOffset(commandBuffer unsafe.Pointer, source unsafe.Pointer, sourceHistogram unsafe.Pointer, sourceHistogramOffset uint, desiredHistogram unsafe.Pointer, desiredHistogramOffset uint)


}





// Alloc allocates a new instance without initialization.
func (ic _ImageHistogramSpecificationClass) Alloc() ImageHistogramSpecification {
	rv := objc.Send[ImageHistogramSpecification](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (ic _ImageHistogramSpecificationClass) New() ImageHistogramSpecification {
	rv := objc.Send[ImageHistogramSpecification](objc.ID(ic.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (i_ ImageHistogramSpecification) Init() ImageHistogramSpecification {
	rv := objc.Send[ImageHistogramSpecification](i_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i_ ImageHistogramSpecification) Autorelease() ImageHistogramSpecification {
	rv := objc.Send[ImageHistogramSpecification](i_.ID, objc.Sel("autorelease"))
	return rv
}

// NewImageHistogramSpecification creates a new ImageHistogramSpecification instance.
func NewImageHistogramSpecification() ImageHistogramSpecification {
	return getImageHistogramSpecificationClass().New()
}





// A filter that performs a histogram specification operation on an image.
//
// is a generalized version of histogram equalization operation. The histogram specification filter converts the image so that its histogram matches the desired histogram. The process is divided into three steps: Call the method to create a object. Call the method. This creates a privately held image transform which will convert the distribution of the source histogram to the desired histogram. This process runs on a command buffer when it is committed to a command queue. It must complete before the next step can be run. It may be performed on the same command buffer. The argument is used by the method to determine the number of channels and therefore which histogram data in the source histogram buffer to use. The source histogram and desired histogram must have been computed either on the CPU or using the kernel. Call the method to read data from the source texture, apply the equalization transform to it, and write to the destination texture. This step is also done on the GPU on a command queue.


// A filter that performs a histogram specification operation on an image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImageHistogramSpecification
type ImageHistogramSpecification struct {
	UnaryImageKernel
}

// ImageHistogramSpecificationFrom constructs a [ImageHistogramSpecification] from an unsafe.Pointer.
//
// A filter that performs a histogram specification operation on an image.
func ImageHistogramSpecificationFrom(ptr unsafe.Pointer) ImageHistogramSpecification {
	return ImageHistogramSpecification{
		UnaryImageKernel: UnaryImageKernelFrom(ptr),
	}
}






// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagehistogramspecification/2867143-initwithcoder
func NewImageHistogramSpecificationWithCoderDevice(aDecoder foundation.Coder, device unsafe.Pointer) ImageHistogramSpecification {
	instance := getImageHistogramSpecificationClass().Alloc()
	rv := objc.Send[ImageHistogramSpecification](instance.ID, objc.Sel("initWithCoder:device:"), aDecoder, device)
	rv.Autorelease()
	return rv
}


// Initializes a histogram with specific information.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagehistogramspecification/1618907-initwithdevice
func NewImageHistogramSpecificationWithDeviceHistogramInfo(device unsafe.Pointer, histogramInfo ImageHistogramInfo) ImageHistogramSpecification {
	instance := getImageHistogramSpecificationClass().Alloc()
	rv := objc.Send[ImageHistogramSpecification](instance.ID, objc.Sel("initWithDevice:histogramInfo:"), device, histogramInfo)
	rv.Autorelease()
	return rv
}

















// Encodes the transform function to a command buffer using a compute command encoder. The transform function computes the equalization lookup table.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagehistogramspecification/1618854-encodetransform
func (i_ ImageHistogramSpecification) EncodeTransform() {
	objc.Send[objc.ID](i_.ID, objc.Sel("encodeTransform"))
}


// Encodes the transform function to a command buffer using a compute command encoder. The transform function computes the equalization lookup table.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagehistogramspecification/1618854-encodetransformtocommandbuffer
func (i_ ImageHistogramSpecification) EncodeTransformToCommandBufferSourceTextureSourceHistogramSourceHistogramOffsetDesiredHistogramDesiredHistogramOffset(commandBuffer unsafe.Pointer, source unsafe.Pointer, sourceHistogram unsafe.Pointer, sourceHistogramOffset uint, desiredHistogram unsafe.Pointer, desiredHistogramOffset uint) {
	objc.Send[objc.ID](i_.ID, objc.Sel("encodeTransformToCommandBuffer:sourceTexture:sourceHistogram:sourceHistogramOffset:desiredHistogram:desiredHistogramOffset:"), commandBuffer, source, sourceHistogram, sourceHistogramOffset, desiredHistogram, desiredHistogramOffset)
}







// A structure describing the histogram content.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagehistogramspecification/1618810-histograminfo
func (i_ ImageHistogramSpecification) HistogramInfo() ImageHistogramInfo get /* not a class type */ {
	rv := objc.Send[objc.ID](i_.ID, objc.Sel("histogramInfo"))
	return rv
}


// A structure describing the histogram content.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagehistogramspecification/1618810-histograminfo
func (i_ ImageHistogramSpecification) SetHistogramInfo(value ImageHistogramInfo get /* not a class type */) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setHistogramInfo:"), value)
}







