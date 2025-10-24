// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MPSImageEDLines */


/* debug [class_header]: Header for MPSImageEDLines */
// The class instance for the [ImageEDLines] class.
var (
	ImageEDLinesClass     _ImageEDLinesClass
	ImageEDLinesClassOnce sync.Once
)

func getImageEDLinesClass() _ImageEDLinesClass {
	ImageEDLinesClassOnce.Do(func() {
		ImageEDLinesClass = _ImageEDLinesClass{objc.GetClass("MPSImageEDLines")}
	})
	return ImageEDLinesClass
}

type _ImageEDLinesClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for ImageEDLines */
// An interface definition for the [ImageEDLines] class.
type IImageEDLines interface {
	IKernel
	
/* debug [class_interface_properties]: Properties for ImageEDLines */
	// properties:
	ClipRectSource() Region get set /* not a class type */
	SetClipRectSource(value Region get set /* not a class type */)
	DetailRatio() objectivec.IObject
	SetDetailRatio(value objectivec.IObject)
	GaussianSigma() objectivec.IObject
	SetGaussianSigma(value objectivec.IObject)
	GradientThreshold() objectivec.IObject
	SetGradientThreshold(value objectivec.IObject)
	LineErrorThreshold() objectivec.IObject
	SetLineErrorThreshold(value objectivec.IObject)
	MaxLines() objectivec.IObject
	SetMaxLines(value objectivec.IObject)
	MergeLocalityThreshold() objectivec.IObject
	SetMergeLocalityThreshold(value objectivec.IObject)
	MinLineLength() objectivec.IObject
	SetMinLineLength(value objectivec.IObject)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for ImageEDLines */
	// methods:
	Encode()
	EncodeToCommandBufferSourceTextureDestinationTextureEndpointBufferEndpointOffset(commandBuffer unsafe.Pointer, source unsafe.Pointer, dest unsafe.Pointer, endpointBuffer unsafe.Pointer, endpointOffset uint)
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for ImageEDLines */
// Alloc allocates a new instance without initialization.
func (ic _ImageEDLinesClass) Alloc() ImageEDLines {
	rv := objc.Send[ImageEDLines](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (ic _ImageEDLinesClass) New() ImageEDLines {
	rv := objc.Send[ImageEDLines](objc.ID(ic.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (i_ ImageEDLines) Init() ImageEDLines {
	rv := objc.Send[ImageEDLines](i_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i_ ImageEDLines) Autorelease() ImageEDLines {
	rv := objc.Send[ImageEDLines](i_.ID, objc.Sel("autorelease"))
	return rv
}

// NewImageEDLines creates a new ImageEDLines instance.
func NewImageEDLines() ImageEDLines {
	return getImageEDLinesClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for ImageEDLines */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImageEDLines
type ImageEDLines struct {
	Kernel
}

// ImageEDLinesFrom constructs a [ImageEDLines] from an unsafe.Pointer.
func ImageEDLinesFrom(ptr unsafe.Pointer) ImageEDLines {
	return ImageEDLines{
		Kernel: KernelFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for ImageEDLines */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimageedlines/3618920-initwithcoder
func NewImageEDLinesWithCoderDevice(aDecoder foundation.Coder, device unsafe.Pointer) ImageEDLines {
	instance := getImageEDLinesClass().Alloc()
	rv := objc.Send[ImageEDLines](instance.ID, objc.Sel("initWithCoder:device:"), aDecoder, device)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewImageEDLinesWithCoderDevice */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimageedlines/3618921-initwithdevice
func NewImageEDLinesWithDeviceGaussianSigmaMinLineLengthMaxLinesDetailRatioGradientThresholdLineErrorThresholdMergeLocalityThreshold(device unsafe.Pointer, gaussianSigma float32, minLineLength objectivec.IObject, maxLines uint, detailRatio objectivec.IObject, gradientThreshold float32, lineErrorThreshold float32, mergeLocalityThreshold float32) ImageEDLines {
	instance := getImageEDLinesClass().Alloc()
	rv := objc.Send[ImageEDLines](instance.ID, objc.Sel("initWithDevice:gaussianSigma:minLineLength:maxLines:detailRatio:gradientThreshold:lineErrorThreshold:mergeLocalityThreshold:"), device, gaussianSigma, minLineLength, maxLines, detailRatio, gradientThreshold, lineErrorThreshold, mergeLocalityThreshold)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewImageEDLinesWithDeviceGaussianSigmaMinLineLengthMaxLinesDetailRatioGradientThresholdLineErrorThresholdMergeLocalityThreshold */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for ImageEDLines */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for ImageEDLines */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for ImageEDLines */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimageedlines/3618917-encode
func (i_ ImageEDLines) Encode() {
	objc.Send[objc.ID](i_.ID, objc.Sel("encode"))
}/* debug [instance_methods/method]: Encode */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimageedlines/3618917-encodetocommandbuffer
func (i_ ImageEDLines) EncodeToCommandBufferSourceTextureDestinationTextureEndpointBufferEndpointOffset(commandBuffer unsafe.Pointer, source unsafe.Pointer, dest unsafe.Pointer, endpointBuffer unsafe.Pointer, endpointOffset uint) {
	objc.Send[objc.ID](i_.ID, objc.Sel("encodeToCommandBuffer:sourceTexture:destinationTexture:endpointBuffer:endpointOffset:"), commandBuffer, source, dest, endpointBuffer, endpointOffset)
}/* debug [instance_methods/method]: EncodeToCommandBufferSourceTextureDestinationTextureEndpointBufferEndpointOffset */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for ImageEDLines */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimageedlines/3618915-cliprectsource
func (i_ ImageEDLines) ClipRectSource() Region get set /* not a class type */ {
	rv := objc.Send[objc.ID](i_.ID, objc.Sel("clipRectSource"))
	return rv
}/* debug [instance_properties/getter]: clipRectSource */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimageedlines/3618915-cliprectsource
func (i_ ImageEDLines) SetClipRectSource(value Region get set /* not a class type */) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setClipRectSource:"), value)
}/* debug [instance_properties/setter]: clipRectSource */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimageedlines/3618916-detailratio
func (i_ ImageEDLines) DetailRatio() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](i_.ID, objc.Sel("detailRatio"))
	return rv
}/* debug [instance_properties/getter]: detailRatio */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimageedlines/3618916-detailratio
func (i_ ImageEDLines) SetDetailRatio(value objectivec.IObject) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setDetailRatio:"), value)
}/* debug [instance_properties/setter]: detailRatio */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimageedlines/3618918-gaussiansigma
func (i_ ImageEDLines) GaussianSigma() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](i_.ID, objc.Sel("gaussianSigma"))
	return rv
}/* debug [instance_properties/getter]: gaussianSigma */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimageedlines/3618918-gaussiansigma
func (i_ ImageEDLines) SetGaussianSigma(value objectivec.IObject) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setGaussianSigma:"), value)
}/* debug [instance_properties/setter]: gaussianSigma */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimageedlines/3618919-gradientthreshold
func (i_ ImageEDLines) GradientThreshold() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](i_.ID, objc.Sel("gradientThreshold"))
	return rv
}/* debug [instance_properties/getter]: gradientThreshold */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimageedlines/3618919-gradientthreshold
func (i_ ImageEDLines) SetGradientThreshold(value objectivec.IObject) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setGradientThreshold:"), value)
}/* debug [instance_properties/setter]: gradientThreshold */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimageedlines/3618922-lineerrorthreshold
func (i_ ImageEDLines) LineErrorThreshold() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](i_.ID, objc.Sel("lineErrorThreshold"))
	return rv
}/* debug [instance_properties/getter]: lineErrorThreshold */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimageedlines/3618922-lineerrorthreshold
func (i_ ImageEDLines) SetLineErrorThreshold(value objectivec.IObject) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setLineErrorThreshold:"), value)
}/* debug [instance_properties/setter]: lineErrorThreshold */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimageedlines/3618923-maxlines
func (i_ ImageEDLines) MaxLines() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](i_.ID, objc.Sel("maxLines"))
	return rv
}/* debug [instance_properties/getter]: maxLines */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimageedlines/3618923-maxlines
func (i_ ImageEDLines) SetMaxLines(value objectivec.IObject) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setMaxLines:"), value)
}/* debug [instance_properties/setter]: maxLines */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimageedlines/3618924-mergelocalitythreshold
func (i_ ImageEDLines) MergeLocalityThreshold() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](i_.ID, objc.Sel("mergeLocalityThreshold"))
	return rv
}/* debug [instance_properties/getter]: mergeLocalityThreshold */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimageedlines/3618924-mergelocalitythreshold
func (i_ ImageEDLines) SetMergeLocalityThreshold(value objectivec.IObject) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setMergeLocalityThreshold:"), value)
}/* debug [instance_properties/setter]: mergeLocalityThreshold */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimageedlines/3618925-minlinelength
func (i_ ImageEDLines) MinLineLength() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](i_.ID, objc.Sel("minLineLength"))
	return rv
}/* debug [instance_properties/getter]: minLineLength */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimageedlines/3618925-minlinelength
func (i_ ImageEDLines) SetMinLineLength(value objectivec.IObject) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setMinLineLength:"), value)
}/* debug [instance_properties/setter]: minLineLength */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MPSImageEDLines */


