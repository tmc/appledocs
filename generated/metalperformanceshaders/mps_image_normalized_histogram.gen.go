// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)





// The class instance for the [ImageNormalizedHistogram] class.
var (
	ImageNormalizedHistogramClass     _ImageNormalizedHistogramClass
	ImageNormalizedHistogramClassOnce sync.Once
)

func getImageNormalizedHistogramClass() _ImageNormalizedHistogramClass {
	ImageNormalizedHistogramClassOnce.Do(func() {
		ImageNormalizedHistogramClass = _ImageNormalizedHistogramClass{objc.GetClass("MPSImageNormalizedHistogram")}
	})
	return ImageNormalizedHistogramClass
}

type _ImageNormalizedHistogramClass struct {
	class objc.Class
}





// An interface definition for the [ImageNormalizedHistogram] class.
type IImageNormalizedHistogram interface {
	IKernel
	

	// properties:
	ClipRectSource() Region get set /* not a class type */
	SetClipRectSource(value Region get set /* not a class type */)
	HistogramInfo() ImageHistogramInfo get /* not a class type */
	SetHistogramInfo(value ImageHistogramInfo get /* not a class type */)
	ZeroHistogram() objectivec.IObject
	SetZeroHistogram(value objectivec.IObject)


	

	// methods:
	Encode()
	EncodeToCommandBufferSourceTextureMinmaxTextureHistogramHistogramOffset(commandBuffer unsafe.Pointer, source unsafe.Pointer, minmaxTexture unsafe.Pointer, histogram unsafe.Pointer, histogramOffset uint)
	HistogramSize()
	HistogramSizeForSourceFormat(sourceFormat PixelFormat /* not a class type */) uintptr /* not a class type */


}





// Alloc allocates a new instance without initialization.
func (ic _ImageNormalizedHistogramClass) Alloc() ImageNormalizedHistogram {
	rv := objc.Send[ImageNormalizedHistogram](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (ic _ImageNormalizedHistogramClass) New() ImageNormalizedHistogram {
	rv := objc.Send[ImageNormalizedHistogram](objc.ID(ic.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (i_ ImageNormalizedHistogram) Init() ImageNormalizedHistogram {
	rv := objc.Send[ImageNormalizedHistogram](i_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i_ ImageNormalizedHistogram) Autorelease() ImageNormalizedHistogram {
	rv := objc.Send[ImageNormalizedHistogram](i_.ID, objc.Sel("autorelease"))
	return rv
}

// NewImageNormalizedHistogram creates a new ImageNormalizedHistogram instance.
func NewImageNormalizedHistogram() ImageNormalizedHistogram {
	return getImageNormalizedHistogramClass().New()
}





// A filter that computes the normalized histogram of an image.


// A filter that computes the normalized histogram of an image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImageNormalizedHistogram
type ImageNormalizedHistogram struct {
	Kernel
}

// ImageNormalizedHistogramFrom constructs a [ImageNormalizedHistogram] from an unsafe.Pointer.
//
// A filter that computes the normalized histogram of an image.
func ImageNormalizedHistogramFrom(ptr unsafe.Pointer) ImageNormalizedHistogram {
	return ImageNormalizedHistogram{
		Kernel: KernelFrom(ptr),
	}
}






// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagenormalizedhistogram/3019325-initwithcoder
func NewImageNormalizedHistogramWithCoderDevice(aDecoder foundation.Coder, device unsafe.Pointer) ImageNormalizedHistogram {
	instance := getImageNormalizedHistogramClass().Alloc()
	rv := objc.Send[ImageNormalizedHistogram](instance.ID, objc.Sel("initWithCoder:device:"), aDecoder, device)
	rv.Autorelease()
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagenormalizedhistogram/3019326-initwithdevice
func NewImageNormalizedHistogramWithDeviceHistogramInfo(device unsafe.Pointer, histogramInfo ImageHistogramInfo) ImageNormalizedHistogram {
	instance := getImageNormalizedHistogramClass().Alloc()
	rv := objc.Send[ImageNormalizedHistogram](instance.ID, objc.Sel("initWithDevice:histogramInfo:"), device, histogramInfo)
	rv.Autorelease()
	return rv
}

















// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagenormalizedhistogram/3019322-encode
func (i_ ImageNormalizedHistogram) Encode() {
	objc.Send[objc.ID](i_.ID, objc.Sel("encode"))
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagenormalizedhistogram/3019322-encodetocommandbuffer
func (i_ ImageNormalizedHistogram) EncodeToCommandBufferSourceTextureMinmaxTextureHistogramHistogramOffset(commandBuffer unsafe.Pointer, source unsafe.Pointer, minmaxTexture unsafe.Pointer, histogram unsafe.Pointer, histogramOffset uint) {
	objc.Send[objc.ID](i_.ID, objc.Sel("encodeToCommandBuffer:sourceTexture:minmaxTexture:histogram:histogramOffset:"), commandBuffer, source, minmaxTexture, histogram, histogramOffset)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagenormalizedhistogram/3019324-histogramsize
func (i_ ImageNormalizedHistogram) HistogramSize() {
	objc.Send[objc.ID](i_.ID, objc.Sel("histogramSize"))
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagenormalizedhistogram/3019324-histogramsizeforsourceformat
func (i_ ImageNormalizedHistogram) HistogramSizeForSourceFormat(sourceFormat PixelFormat /* not a class type */) uintptr /* not a class type */ {
	rv := objc.Send[uintptr](i_.ID, objc.Sel("histogramSizeForSourceFormat:"), sourceFormat)
	return rv
}







// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagenormalizedhistogram/3019321-cliprectsource
func (i_ ImageNormalizedHistogram) ClipRectSource() Region get set /* not a class type */ {
	rv := objc.Send[objc.ID](i_.ID, objc.Sel("clipRectSource"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagenormalizedhistogram/3019321-cliprectsource
func (i_ ImageNormalizedHistogram) SetClipRectSource(value Region get set /* not a class type */) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setClipRectSource:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagenormalizedhistogram/3019323-histograminfo
func (i_ ImageNormalizedHistogram) HistogramInfo() ImageHistogramInfo get /* not a class type */ {
	rv := objc.Send[objc.ID](i_.ID, objc.Sel("histogramInfo"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagenormalizedhistogram/3019323-histograminfo
func (i_ ImageNormalizedHistogram) SetHistogramInfo(value ImageHistogramInfo get /* not a class type */) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setHistogramInfo:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagenormalizedhistogram/3019327-zerohistogram
func (i_ ImageNormalizedHistogram) ZeroHistogram() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](i_.ID, objc.Sel("zeroHistogram"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagenormalizedhistogram/3019327-zerohistogram
func (i_ ImageNormalizedHistogram) SetZeroHistogram(value objectivec.IObject) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setZeroHistogram:"), value)
}







