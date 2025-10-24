// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MPSImageHistogram */


/* debug [class_header]: Header for MPSImageHistogram */
// The class instance for the [ImageHistogram] class.
var (
	ImageHistogramClass     _ImageHistogramClass
	ImageHistogramClassOnce sync.Once
)

func getImageHistogramClass() _ImageHistogramClass {
	ImageHistogramClassOnce.Do(func() {
		ImageHistogramClass = _ImageHistogramClass{objc.GetClass("MPSImageHistogram")}
	})
	return ImageHistogramClass
}

type _ImageHistogramClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for ImageHistogram */
// An interface definition for the [ImageHistogram] class.
type IImageHistogram interface {
	IKernel
	
/* debug [class_interface_properties]: Properties for ImageHistogram */
	// properties:
	ClipRectSource() Region get set /* not a class type */
	SetClipRectSource(value Region get set /* not a class type */)
	HistogramInfo() ImageHistogramInfo get /* not a class type */
	SetHistogramInfo(value ImageHistogramInfo get /* not a class type */)
	ZeroHistogram() objectivec.IObject
	SetZeroHistogram(value objectivec.IObject)
	MinPixelThresholdValue() objectivec.IObject
	SetMinPixelThresholdValue(value objectivec.IObject)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for ImageHistogram */
	// methods:
	HistogramSize()
	HistogramSizeForSourceFormat(sourceFormat PixelFormat /* not a class type */) uintptr /* not a class type */
	Encode()
	EncodeToCommandBufferSourceTextureHistogramHistogramOffset(commandBuffer unsafe.Pointer, source unsafe.Pointer, histogram unsafe.Pointer, histogramOffset uint)
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for ImageHistogram */
// Alloc allocates a new instance without initialization.
func (ic _ImageHistogramClass) Alloc() ImageHistogram {
	rv := objc.Send[ImageHistogram](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (ic _ImageHistogramClass) New() ImageHistogram {
	rv := objc.Send[ImageHistogram](objc.ID(ic.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (i_ ImageHistogram) Init() ImageHistogram {
	rv := objc.Send[ImageHistogram](i_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i_ ImageHistogram) Autorelease() ImageHistogram {
	rv := objc.Send[ImageHistogram](i_.ID, objc.Sel("autorelease"))
	return rv
}

// NewImageHistogram creates a new ImageHistogram instance.
func NewImageHistogram() ImageHistogram {
	return getImageHistogramClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for ImageHistogram */
// A filter that computes the histogram of an image.
//
// Typically, you use an filter to calculate an image’s histogram that is passed to a subsequent filter such as or . The following listing shows how you can create a histogram filter to calculate the histogram of the , . The filter is passed an instance of that specifies information to compute the histogram for the channels of an image. After encoding, contains the histogram information and can be used for further operations such as equalization or specification. Listing 1. Creating a histogram filter


// A filter that computes the histogram of an image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImageHistogram
type ImageHistogram struct {
	Kernel
}

// ImageHistogramFrom constructs a [ImageHistogram] from an unsafe.Pointer.
//
// A filter that computes the histogram of an image.
func ImageHistogramFrom(ptr unsafe.Pointer) ImageHistogram {
	return ImageHistogram{
		Kernel: KernelFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for ImageHistogram */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagehistogram/2867090-initwithcoder
func NewImageHistogramWithCoderDevice(aDecoder Coder /* not a class type */, device unsafe.Pointer) ImageHistogram {
	instance := getImageHistogramClass().Alloc()
	rv := objc.Send[ImageHistogram](instance.ID, objc.Sel("initWithCoder:device:"), aDecoder, device)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewImageHistogramWithCoderDevice */


// Initializes a histogram with specific information.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagehistogram/1618910-initwithdevice
func NewImageHistogramWithDeviceHistogramInfo(device unsafe.Pointer, histogramInfo objc.IObject /* cross-framework: MPSImageHistogramInfo */) ImageHistogram {
	instance := getImageHistogramClass().Alloc()
	rv := objc.Send[ImageHistogram](instance.ID, objc.Sel("initWithDevice:histogramInfo:"), device, histogramInfo)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewImageHistogramWithDeviceHistogramInfo */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for ImageHistogram */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for ImageHistogram */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for ImageHistogram */

// The amount of space the histogram will take up in the output buffer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagehistogram/1618839-histogramsize
func (i_ ImageHistogram) HistogramSize() {
	objc.Send[objc.ID](i_.ID, objc.Sel("histogramSize"))
}/* debug [instance_methods/method]: HistogramSize */


// The amount of space the histogram will take up in the output buffer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagehistogram/1618839-histogramsizeforsourceformat
func (i_ ImageHistogram) HistogramSizeForSourceFormat(sourceFormat PixelFormat /* not a class type */) uintptr /* not a class type */ {
	rv := objc.Send[uintptr](i_.ID, objc.Sel("histogramSizeForSourceFormat:"), sourceFormat)
	return rv
}/* debug [instance_methods/method]: HistogramSizeForSourceFormat */


// Encodes the filter to a command buffer using a compute command encoder.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagehistogram/1618853-encode
func (i_ ImageHistogram) Encode() {
	objc.Send[objc.ID](i_.ID, objc.Sel("encode"))
}/* debug [instance_methods/method]: Encode */


// Encodes the filter to a command buffer using a compute command encoder.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagehistogram/1618853-encodetocommandbuffer
func (i_ ImageHistogram) EncodeToCommandBufferSourceTextureHistogramHistogramOffset(commandBuffer unsafe.Pointer, source unsafe.Pointer, histogram unsafe.Pointer, histogramOffset uint) {
	objc.Send[objc.ID](i_.ID, objc.Sel("encodeToCommandBuffer:sourceTexture:histogram:histogramOffset:"), commandBuffer, source, histogram, histogramOffset)
}/* debug [instance_methods/method]: EncodeToCommandBufferSourceTextureHistogramHistogramOffset */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for ImageHistogram */

// The source rectangle to use when reading data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagehistogram/1618765-cliprectsource
func (i_ ImageHistogram) ClipRectSource() Region get set /* not a class type */ {
	rv := objc.Send[objc.ID](i_.ID, objc.Sel("clipRectSource"))
	return rv
}/* debug [instance_properties/getter]: clipRectSource */


// The source rectangle to use when reading data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagehistogram/1618765-cliprectsource
func (i_ ImageHistogram) SetClipRectSource(value Region get set /* not a class type */) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setClipRectSource:"), value)
}/* debug [instance_properties/setter]: clipRectSource */


// A structure describing the histogram content.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagehistogram/1618844-histograminfo
func (i_ ImageHistogram) HistogramInfo() ImageHistogramInfo get /* not a class type */ {
	rv := objc.Send[objc.ID](i_.ID, objc.Sel("histogramInfo"))
	return rv
}/* debug [instance_properties/getter]: histogramInfo */


// A structure describing the histogram content.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagehistogram/1618844-histograminfo
func (i_ ImageHistogram) SetHistogramInfo(value ImageHistogramInfo get /* not a class type */) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setHistogramInfo:"), value)
}/* debug [instance_properties/setter]: histogramInfo */


// Determines whether to zero-initialize the histogram results.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagehistogram/1618891-zerohistogram
func (i_ ImageHistogram) ZeroHistogram() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](i_.ID, objc.Sel("zeroHistogram"))
	return rv
}/* debug [instance_properties/getter]: zeroHistogram */


// Determines whether to zero-initialize the histogram results.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagehistogram/1618891-zerohistogram
func (i_ ImageHistogram) SetZeroHistogram(value objectivec.IObject) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setZeroHistogram:"), value)
}/* debug [instance_properties/setter]: zeroHistogram */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagehistogram/2867008-minpixelthresholdvalue
func (i_ ImageHistogram) MinPixelThresholdValue() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](i_.ID, objc.Sel("minPixelThresholdValue"))
	return rv
}/* debug [instance_properties/getter]: minPixelThresholdValue */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagehistogram/2867008-minpixelthresholdvalue
func (i_ ImageHistogram) SetMinPixelThresholdValue(value objectivec.IObject) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setMinPixelThresholdValue:"), value)
}/* debug [instance_properties/setter]: minPixelThresholdValue */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MPSImageHistogram */


