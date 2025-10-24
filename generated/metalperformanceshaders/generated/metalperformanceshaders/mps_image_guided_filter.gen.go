// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MPSImageGuidedFilter */


/* debug [class_header]: Header for MPSImageGuidedFilter */
// The class instance for the [ImageGuidedFilter] class.
var (
	ImageGuidedFilterClass     _ImageGuidedFilterClass
	ImageGuidedFilterClassOnce sync.Once
)

func getImageGuidedFilterClass() _ImageGuidedFilterClass {
	ImageGuidedFilterClassOnce.Do(func() {
		ImageGuidedFilterClass = _ImageGuidedFilterClass{objc.GetClass("MPSImageGuidedFilter")}
	})
	return ImageGuidedFilterClass
}

type _ImageGuidedFilterClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for ImageGuidedFilter */
// An interface definition for the [ImageGuidedFilter] class.
type IImageGuidedFilter interface {
	IKernel
	
/* debug [class_interface_properties]: Properties for ImageGuidedFilter */
	// properties:
	Epsilon() objectivec.IObject
	SetEpsilon(value objectivec.IObject)
	KernelDiameter() objectivec.IObject
	SetKernelDiameter(value objectivec.IObject)
	ReconstructScale() objectivec.IObject
	SetReconstructScale(value objectivec.IObject)
	ReconstructOffset() objectivec.IObject
	SetReconstructOffset(value objectivec.IObject)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for ImageGuidedFilter */
	// methods:
	EncodeReconstruction()
	EncodeReconstructionToCommandBufferGuidanceTextureCoefficientsTextureDestinationTexture(commandBuffer unsafe.Pointer, guidanceTexture unsafe.Pointer, coefficientsTexture unsafe.Pointer, destinationTexture unsafe.Pointer)
	EncodeRegression()
	EncodeRegressionToCommandBufferSourceTextureGuidanceTextureWeightsTextureDestinationCoefficientsTexture(commandBuffer unsafe.Pointer, sourceTexture unsafe.Pointer, guidanceTexture unsafe.Pointer, weightsTexture unsafe.Pointer, destinationCoefficientsTexture unsafe.Pointer)
	EncodeReconstructionToCommandBufferGuidanceTextureCoefficientsTextureACoefficientsTextureBDestinationTexture(commandBuffer unsafe.Pointer, guidanceTexture unsafe.Pointer, coefficientsTextureA unsafe.Pointer, coefficientsTextureB unsafe.Pointer, destinationTexture unsafe.Pointer)
	EncodeRegressionToCommandBufferSourceTextureGuidanceTextureWeightsTextureDestinationCoefficientsTextureADestinationCoefficientsTextureB(commandBuffer unsafe.Pointer, sourceTexture unsafe.Pointer, guidanceTexture unsafe.Pointer, weightsTexture unsafe.Pointer, destinationCoefficientsTextureA unsafe.Pointer, destinationCoefficientsTextureB unsafe.Pointer)
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for ImageGuidedFilter */
// Alloc allocates a new instance without initialization.
func (ic _ImageGuidedFilterClass) Alloc() ImageGuidedFilter {
	rv := objc.Send[ImageGuidedFilter](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (ic _ImageGuidedFilterClass) New() ImageGuidedFilter {
	rv := objc.Send[ImageGuidedFilter](objc.ID(ic.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (i_ ImageGuidedFilter) Init() ImageGuidedFilter {
	rv := objc.Send[ImageGuidedFilter](i_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i_ ImageGuidedFilter) Autorelease() ImageGuidedFilter {
	rv := objc.Send[ImageGuidedFilter](i_.ID, objc.Sel("autorelease"))
	return rv
}

// NewImageGuidedFilter creates a new ImageGuidedFilter instance.
func NewImageGuidedFilter() ImageGuidedFilter {
	return getImageGuidedFilterClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for ImageGuidedFilter */
// A filter that performs edge-aware filtering on an image.


// A filter that performs edge-aware filtering on an image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImageGuidedFilter
type ImageGuidedFilter struct {
	Kernel
}

// ImageGuidedFilterFrom constructs a [ImageGuidedFilter] from an unsafe.Pointer.
//
// A filter that performs edge-aware filtering on an image.
func ImageGuidedFilterFrom(ptr unsafe.Pointer) ImageGuidedFilter {
	return ImageGuidedFilter{
		Kernel: KernelFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for ImageGuidedFilter */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimageguidedfilter/2951903-initwithcoder
func NewImageGuidedFilterWithCoderDevice(aDecoder Coder /* not a class type */, device unsafe.Pointer) ImageGuidedFilter {
	instance := getImageGuidedFilterClass().Alloc()
	rv := objc.Send[ImageGuidedFilter](instance.ID, objc.Sel("initWithCoder:device:"), aDecoder, device)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewImageGuidedFilterWithCoderDevice */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimageguidedfilter/2951910-initwithdevice
func NewImageGuidedFilterWithDeviceKernelDiameter(device unsafe.Pointer, kernelDiameter uint) ImageGuidedFilter {
	instance := getImageGuidedFilterClass().Alloc()
	rv := objc.Send[ImageGuidedFilter](instance.ID, objc.Sel("initWithDevice:kernelDiameter:"), device, kernelDiameter)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewImageGuidedFilterWithDeviceKernelDiameter */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for ImageGuidedFilter */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for ImageGuidedFilter */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for ImageGuidedFilter */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimageguidedfilter/2951906-encodereconstruction
func (i_ ImageGuidedFilter) EncodeReconstruction() {
	objc.Send[objc.ID](i_.ID, objc.Sel("encodeReconstruction"))
}/* debug [instance_methods/method]: EncodeReconstruction */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimageguidedfilter/2951906-encodereconstructiontocommandbuf
func (i_ ImageGuidedFilter) EncodeReconstructionToCommandBufferGuidanceTextureCoefficientsTextureDestinationTexture(commandBuffer unsafe.Pointer, guidanceTexture unsafe.Pointer, coefficientsTexture unsafe.Pointer, destinationTexture unsafe.Pointer) {
	objc.Send[objc.ID](i_.ID, objc.Sel("encodeReconstructionToCommandBuffer:guidanceTexture:coefficientsTexture:destinationTexture:"), commandBuffer, guidanceTexture, coefficientsTexture, destinationTexture)
}/* debug [instance_methods/method]: EncodeReconstructionToCommandBufferGuidanceTextureCoefficientsTextureDestinationTexture */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimageguidedfilter/2951907-encoderegression
func (i_ ImageGuidedFilter) EncodeRegression() {
	objc.Send[objc.ID](i_.ID, objc.Sel("encodeRegression"))
}/* debug [instance_methods/method]: EncodeRegression */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimageguidedfilter/2951907-encoderegressiontocommandbuffer
func (i_ ImageGuidedFilter) EncodeRegressionToCommandBufferSourceTextureGuidanceTextureWeightsTextureDestinationCoefficientsTexture(commandBuffer unsafe.Pointer, sourceTexture unsafe.Pointer, guidanceTexture unsafe.Pointer, weightsTexture unsafe.Pointer, destinationCoefficientsTexture unsafe.Pointer) {
	objc.Send[objc.ID](i_.ID, objc.Sel("encodeRegressionToCommandBuffer:sourceTexture:guidanceTexture:weightsTexture:destinationCoefficientsTexture:"), commandBuffer, sourceTexture, guidanceTexture, weightsTexture, destinationCoefficientsTexture)
}/* debug [instance_methods/method]: EncodeRegressionToCommandBufferSourceTextureGuidanceTextureWeightsTextureDestinationCoefficientsTexture */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimageguidedfilter/3516398-encodereconstructiontocommandbuf
func (i_ ImageGuidedFilter) EncodeReconstructionToCommandBufferGuidanceTextureCoefficientsTextureACoefficientsTextureBDestinationTexture(commandBuffer unsafe.Pointer, guidanceTexture unsafe.Pointer, coefficientsTextureA unsafe.Pointer, coefficientsTextureB unsafe.Pointer, destinationTexture unsafe.Pointer) {
	objc.Send[objc.ID](i_.ID, objc.Sel("encodeReconstructionToCommandBuffer:guidanceTexture:coefficientsTextureA:coefficientsTextureB:destinationTexture:"), commandBuffer, guidanceTexture, coefficientsTextureA, coefficientsTextureB, destinationTexture)
}/* debug [instance_methods/method]: EncodeReconstructionToCommandBufferGuidanceTextureCoefficientsTextureACoefficientsTextureBDestinationTexture */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimageguidedfilter/3516399-encoderegressiontocommandbuffer
func (i_ ImageGuidedFilter) EncodeRegressionToCommandBufferSourceTextureGuidanceTextureWeightsTextureDestinationCoefficientsTextureADestinationCoefficientsTextureB(commandBuffer unsafe.Pointer, sourceTexture unsafe.Pointer, guidanceTexture unsafe.Pointer, weightsTexture unsafe.Pointer, destinationCoefficientsTextureA unsafe.Pointer, destinationCoefficientsTextureB unsafe.Pointer) {
	objc.Send[objc.ID](i_.ID, objc.Sel("encodeRegressionToCommandBuffer:sourceTexture:guidanceTexture:weightsTexture:destinationCoefficientsTextureA:destinationCoefficientsTextureB:"), commandBuffer, sourceTexture, guidanceTexture, weightsTexture, destinationCoefficientsTextureA, destinationCoefficientsTextureB)
}/* debug [instance_methods/method]: EncodeRegressionToCommandBufferSourceTextureGuidanceTextureWeightsTextureDestinationCoefficientsTextureADestinationCoefficientsTextureB */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for ImageGuidedFilter */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimageguidedfilter/2951908-epsilon
func (i_ ImageGuidedFilter) Epsilon() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](i_.ID, objc.Sel("epsilon"))
	return rv
}/* debug [instance_properties/getter]: epsilon */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimageguidedfilter/2951908-epsilon
func (i_ ImageGuidedFilter) SetEpsilon(value objectivec.IObject) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setEpsilon:"), value)
}/* debug [instance_properties/setter]: epsilon */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimageguidedfilter/2951909-kerneldiameter
func (i_ ImageGuidedFilter) KernelDiameter() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](i_.ID, objc.Sel("kernelDiameter"))
	return rv
}/* debug [instance_properties/getter]: kernelDiameter */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimageguidedfilter/2951909-kerneldiameter
func (i_ ImageGuidedFilter) SetKernelDiameter(value objectivec.IObject) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setKernelDiameter:"), value)
}/* debug [instance_properties/setter]: kernelDiameter */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimageguidedfilter/2953078-reconstructscale
func (i_ ImageGuidedFilter) ReconstructScale() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](i_.ID, objc.Sel("reconstructScale"))
	return rv
}/* debug [instance_properties/getter]: reconstructScale */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimageguidedfilter/2953078-reconstructscale
func (i_ ImageGuidedFilter) SetReconstructScale(value objectivec.IObject) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setReconstructScale:"), value)
}/* debug [instance_properties/setter]: reconstructScale */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimageguidedfilter/2953079-reconstructoffset
func (i_ ImageGuidedFilter) ReconstructOffset() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](i_.ID, objc.Sel("reconstructOffset"))
	return rv
}/* debug [instance_properties/getter]: reconstructOffset */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimageguidedfilter/2953079-reconstructoffset
func (i_ ImageGuidedFilter) SetReconstructOffset(value objectivec.IObject) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setReconstructOffset:"), value)
}/* debug [instance_properties/setter]: reconstructOffset */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MPSImageGuidedFilter */


