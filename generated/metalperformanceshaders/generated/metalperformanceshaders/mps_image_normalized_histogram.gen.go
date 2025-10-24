// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MPSImageNormalizedHistogram */


/* debug [class_header]: Header for MPSImageNormalizedHistogram */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for ImageNormalizedHistogram */
// An interface definition for the [ImageNormalizedHistogram] class.
type IImageNormalizedHistogram interface {
	IKernel
	
/* debug [class_interface_properties]: Properties for ImageNormalizedHistogram */
	// properties:
	ClipRectSource() Region get set /* not a class type */
	SetClipRectSource(value Region get set /* not a class type */)
	HistogramInfo() ImageHistogramInfo get /* not a class type */
	SetHistogramInfo(value ImageHistogramInfo get /* not a class type */)
	ZeroHistogram() objectivec.IObject
	SetZeroHistogram(value objectivec.IObject)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for ImageNormalizedHistogram */
	// methods:
	Encode()
	EncodeToCommandBufferSourceTextureMinmaxTextureHistogramHistogramOffset(commandBuffer unsafe.Pointer, source unsafe.Pointer, minmaxTexture unsafe.Pointer, histogram unsafe.Pointer, histogramOffset uint)
	HistogramSize()
	HistogramSizeForSourceFormat(sourceFormat PixelFormat /* not a class type */) uintptr /* not a class type */
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for ImageNormalizedHistogram */
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for ImageNormalizedHistogram */
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for ImageNormalizedHistogram */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagenormalizedhistogram/3019325-initwithcoder
func NewImageNormalizedHistogramWithCoderDevice(aDecoder Coder /* not a class type */, device unsafe.Pointer) ImageNormalizedHistogram {
	instance := getImageNormalizedHistogramClass().Alloc()
	rv := objc.Send[ImageNormalizedHistogram](instance.ID, objc.Sel("initWithCoder:device:"), aDecoder, device)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewImageNormalizedHistogramWithCoderDevice */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagenormalizedhistogram/3019326-initwithdevice
func NewImageNormalizedHistogramWithDeviceHistogramInfo(device unsafe.Pointer, histogramInfo objc.IObject /* cross-framework: MPSImageHistogramInfo */) ImageNormalizedHistogram {
	instance := getImageNormalizedHistogramClass().Alloc()
	rv := objc.Send[ImageNormalizedHistogram](instance.ID, objc.Sel("initWithDevice:histogramInfo:"), device, histogramInfo)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewImageNormalizedHistogramWithDeviceHistogramInfo */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for ImageNormalizedHistogram */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for ImageNormalizedHistogram */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for ImageNormalizedHistogram */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagenormalizedhistogram/3019322-encode
func (i_ ImageNormalizedHistogram) Encode() {
	objc.Send[objc.ID](i_.ID, objc.Sel("encode"))
}/* debug [instance_methods/method]: Encode */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagenormalizedhistogram/3019322-encodetocommandbuffer
func (i_ ImageNormalizedHistogram) EncodeToCommandBufferSourceTextureMinmaxTextureHistogramHistogramOffset(commandBuffer unsafe.Pointer, source unsafe.Pointer, minmaxTexture unsafe.Pointer, histogram unsafe.Pointer, histogramOffset uint) {
	objc.Send[objc.ID](i_.ID, objc.Sel("encodeToCommandBuffer:sourceTexture:minmaxTexture:histogram:histogramOffset:"), commandBuffer, source, minmaxTexture, histogram, histogramOffset)
}/* debug [instance_methods/method]: EncodeToCommandBufferSourceTextureMinmaxTextureHistogramHistogramOffset */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagenormalizedhistogram/3019324-histogramsize
func (i_ ImageNormalizedHistogram) HistogramSize() {
	objc.Send[objc.ID](i_.ID, objc.Sel("histogramSize"))
}/* debug [instance_methods/method]: HistogramSize */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagenormalizedhistogram/3019324-histogramsizeforsourceformat
func (i_ ImageNormalizedHistogram) HistogramSizeForSourceFormat(sourceFormat PixelFormat /* not a class type */) uintptr /* not a class type */ {
	rv := objc.Send[uintptr](i_.ID, objc.Sel("histogramSizeForSourceFormat:"), sourceFormat)
	return rv
}/* debug [instance_methods/method]: HistogramSizeForSourceFormat */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for ImageNormalizedHistogram */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagenormalizedhistogram/3019321-cliprectsource
func (i_ ImageNormalizedHistogram) ClipRectSource() Region get set /* not a class type */ {
	rv := objc.Send[objc.ID](i_.ID, objc.Sel("clipRectSource"))
	return rv
}/* debug [instance_properties/getter]: clipRectSource */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagenormalizedhistogram/3019321-cliprectsource
func (i_ ImageNormalizedHistogram) SetClipRectSource(value Region get set /* not a class type */) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setClipRectSource:"), value)
}/* debug [instance_properties/setter]: clipRectSource */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagenormalizedhistogram/3019323-histograminfo
func (i_ ImageNormalizedHistogram) HistogramInfo() ImageHistogramInfo get /* not a class type */ {
	rv := objc.Send[objc.ID](i_.ID, objc.Sel("histogramInfo"))
	return rv
}/* debug [instance_properties/getter]: histogramInfo */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagenormalizedhistogram/3019323-histograminfo
func (i_ ImageNormalizedHistogram) SetHistogramInfo(value ImageHistogramInfo get /* not a class type */) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setHistogramInfo:"), value)
}/* debug [instance_properties/setter]: histogramInfo */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagenormalizedhistogram/3019327-zerohistogram
func (i_ ImageNormalizedHistogram) ZeroHistogram() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](i_.ID, objc.Sel("zeroHistogram"))
	return rv
}/* debug [instance_properties/getter]: zeroHistogram */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagenormalizedhistogram/3019327-zerohistogram
func (i_ ImageNormalizedHistogram) SetZeroHistogram(value objectivec.IObject) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setZeroHistogram:"), value)
}/* debug [instance_properties/setter]: zeroHistogram */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MPSImageNormalizedHistogram */


