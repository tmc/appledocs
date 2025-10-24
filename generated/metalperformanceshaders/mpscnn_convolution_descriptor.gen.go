// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MPSCNNConvolutionDescriptor */


/* debug [class_header]: Header for MPSCNNConvolutionDescriptor */
// The class instance for the [CNNConvolutionDescriptor] class.
var (
	CNNConvolutionDescriptorClass     _CNNConvolutionDescriptorClass
	CNNConvolutionDescriptorClassOnce sync.Once
)

func getCNNConvolutionDescriptorClass() _CNNConvolutionDescriptorClass {
	CNNConvolutionDescriptorClassOnce.Do(func() {
		CNNConvolutionDescriptorClass = _CNNConvolutionDescriptorClass{objc.GetClass("MPSCNNConvolutionDescriptor")}
	})
	return CNNConvolutionDescriptorClass
}

type _CNNConvolutionDescriptorClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for CNNConvolutionDescriptor */
// An interface definition for the [CNNConvolutionDescriptor] class.
type ICNNConvolutionDescriptor interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for CNNConvolutionDescriptor */
	// properties:
	StrideInPixelsY() objectivec.IObject
	SetStrideInPixelsY(value objectivec.IObject)
	Groups() objectivec.IObject
	SetGroups(value objectivec.IObject)
	OutputFeatureChannels() objectivec.IObject
	SetOutputFeatureChannels(value objectivec.IObject)
	KernelHeight() objectivec.IObject
	SetKernelHeight(value objectivec.IObject)
	StrideInPixelsX() objectivec.IObject
	SetStrideInPixelsX(value objectivec.IObject)
	InputFeatureChannels() objectivec.IObject
	SetInputFeatureChannels(value objectivec.IObject)
	KernelWidth() objectivec.IObject
	SetKernelWidth(value objectivec.IObject)
	Neuron() IMPSCNNNeuron
	SetNeuron(value IMPSCNNNeuron)
	DilationRateX() objectivec.IObject
	SetDilationRateX(value objectivec.IObject)
	DilationRateY() objectivec.IObject
	SetDilationRateY(value objectivec.IObject)
	FusedNeuronDescriptor() IMPSNNNeuronDescriptor
	SetFusedNeuronDescriptor(value IMPSNNNeuronDescriptor)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for CNNConvolutionDescriptor */
	// methods:
	Encode()
	SetBatchNormalizationParametersForInferenceWithMean()
	SetBatchNormalizationParametersForInferenceWithMeanVarianceGammaBetaEpsilon(mean objectivec.IObject, variance objectivec.IObject, gamma objectivec.IObject, beta objectivec.IObject, epsilon float32)
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for CNNConvolutionDescriptor */
// Alloc allocates a new instance without initialization.
func (cc _CNNConvolutionDescriptorClass) Alloc() CNNConvolutionDescriptor {
	rv := objc.Send[CNNConvolutionDescriptor](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _CNNConvolutionDescriptorClass) New() CNNConvolutionDescriptor {
	rv := objc.Send[CNNConvolutionDescriptor](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CNNConvolutionDescriptor) Init() CNNConvolutionDescriptor {
	rv := objc.Send[CNNConvolutionDescriptor](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CNNConvolutionDescriptor) Autorelease() CNNConvolutionDescriptor {
	rv := objc.Send[CNNConvolutionDescriptor](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCNNConvolutionDescriptor creates a new CNNConvolutionDescriptor instance.
func NewCNNConvolutionDescriptor() CNNConvolutionDescriptor {
	return getCNNConvolutionDescriptorClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for CNNConvolutionDescriptor */
// A description of the attributes of a convolution kernel.
//
// You use an object to describe the properties of an kernel such as its size, pixel format and CPU cache mode.


// A description of the attributes of a convolution kernel.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNConvolutionDescriptor
type CNNConvolutionDescriptor struct {
	objectivec.Object
}

// CNNConvolutionDescriptorFrom constructs a [CNNConvolutionDescriptor] from an unsafe.Pointer.
//
// A description of the attributes of a convolution kernel.
func CNNConvolutionDescriptorFrom(ptr unsafe.Pointer) CNNConvolutionDescriptor {
	return CNNConvolutionDescriptor{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for CNNConvolutionDescriptor */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNConvolutionDescriptor/init(kernelWidth:kernelHeight:inputFeatureChannels:outputFeatureChannels:)
func NewCNNConvolutionDescriptorCnnConvolutionDescriptorWithKernelWidthKernelHeightInputFeatureChannelsOutputFeatureChannels(kernelWidth uint, kernelHeight uint, inputFeatureChannels uint, outputFeatureChannels uint) CNNConvolutionDescriptor {
	rv := objc.Send[CNNConvolutionDescriptor](objc.ID(getCNNConvolutionDescriptorClass().class), objc.Sel("cnnConvolutionDescriptorWithKernelWidth:kernelHeight:inputFeatureChannels:outputFeatureChannels:"), kernelWidth, kernelHeight, inputFeatureChannels, outputFeatureChannels)
	return rv
}/* debug [class_init_methods/constructor]: NewCNNConvolutionDescriptorCnnConvolutionDescriptorWithKernelWidthKernelHeightInputFeatureChannelsOutputFeatureChannels */


// Creates a convolution descriptor with an optional neuron filter.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNConvolutionDescriptor/init(kernelWidth:kernelHeight:inputFeatureChannels:outputFeatureChannels:neuronFilter:)
func NewCNNConvolutionDescriptorCnnConvolutionDescriptorWithKernelWidthKernelHeightInputFeatureChannelsOutputFeatureChannelsNeuronFilter(kernelWidth uint, kernelHeight uint, inputFeatureChannels uint, outputFeatureChannels uint, neuronFilter IMPSCNNNeuron) CNNConvolutionDescriptor {
	rv := objc.Send[CNNConvolutionDescriptor](objc.ID(getCNNConvolutionDescriptorClass().class), objc.Sel("cnnConvolutionDescriptorWithKernelWidth:kernelHeight:inputFeatureChannels:outputFeatureChannels:neuronFilter:"), kernelWidth, kernelHeight, inputFeatureChannels, outputFeatureChannels, neuronFilter)
	return rv
}/* debug [class_init_methods/constructor]: NewCNNConvolutionDescriptorCnnConvolutionDescriptorWithKernelWidthKernelHeightInputFeatureChannelsOutputFeatureChannelsNeuronFilter */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNConvolutionDescriptor/init(coder:)
func NewCNNConvolutionDescriptorWithCoder(aDecoder foundation.Coder) CNNConvolutionDescriptor {
	instance := getCNNConvolutionDescriptorClass().Alloc()
	rv := objc.Send[CNNConvolutionDescriptor](instance.ID, objc.Sel("initWithCoder:"), aDecoder)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewCNNConvolutionDescriptorWithCoder */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for CNNConvolutionDescriptor */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNConvolutionDescriptor/init(kernelWidth:kernelHeight:inputFeatureChannels:outputFeatureChannels:)
func (cc _CNNConvolutionDescriptorClass) CnnConvolutionDescriptorWithKernelWidthKernelHeightInputFeatureChannelsOutputFeatureChannels(kernelWidth uint, kernelHeight uint, inputFeatureChannels uint, outputFeatureChannels uint) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(cc.class), objc.Sel("cnnConvolutionDescriptorWithKernelWidth:kernelHeight:inputFeatureChannels:outputFeatureChannels:"), kernelWidth, kernelHeight, inputFeatureChannels, outputFeatureChannels)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=CnnConvolutionDescriptorWithKernelWidthKernelHeightInputFeatureChannelsOutputFeatureChannels) */


// Creates a convolution descriptor with an optional neuron filter.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNConvolutionDescriptor/init(kernelWidth:kernelHeight:inputFeatureChannels:outputFeatureChannels:neuronFilter:)
func (cc _CNNConvolutionDescriptorClass) CnnConvolutionDescriptorWithKernelWidthKernelHeightInputFeatureChannelsOutputFeatureChannelsNeuronFilter(kernelWidth uint, kernelHeight uint, inputFeatureChannels uint, outputFeatureChannels uint, neuronFilter IMPSCNNNeuron) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(cc.class), objc.Sel("cnnConvolutionDescriptorWithKernelWidth:kernelHeight:inputFeatureChannels:outputFeatureChannels:neuronFilter:"), kernelWidth, kernelHeight, inputFeatureChannels, outputFeatureChannels, neuronFilter)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=CnnConvolutionDescriptorWithKernelWidthKernelHeightInputFeatureChannelsOutputFeatureChannelsNeuronFilter) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for CNNConvolutionDescriptor */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnconvolutiondescriptor/2867154-supportssecurecoding
func (cc _CNNConvolutionDescriptorClass) SupportsSecureCoding() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(cc.class), objc.Sel("supportsSecureCoding"))
	return rv
}/* debug [class_properties_class/property]: supportsSecureCoding */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for CNNConvolutionDescriptor */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnconvolutiondescriptor/2866972-encode
func (c_ CNNConvolutionDescriptor) Encode() {
	objc.Send[objc.ID](c_.ID, objc.Sel("encode"))
}/* debug [instance_methods/method]: Encode */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnconvolutiondescriptor/2867057-setbatchnormalizationparametersf
func (c_ CNNConvolutionDescriptor) SetBatchNormalizationParametersForInferenceWithMean() {
	objc.Send[objc.ID](c_.ID, objc.Sel("setBatchNormalizationParametersForInferenceWithMean"))
}/* debug [instance_methods/method]: SetBatchNormalizationParametersForInferenceWithMean */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNConvolutionDescriptor/encode(with:)
func (c_ CNNConvolutionDescriptor) EncodeWithCoder(aCoder foundation.Coder) {
	objc.Send[objc.ID](c_.ID, objc.Sel("encodeWithCoder:"), aCoder)
}/* debug [instance_methods/method]: EncodeWithCoder */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNConvolutionDescriptor/setBatchNormalizationParametersForInferenceWithMean(_:variance:gamma:beta:epsilon:)
func (c_ CNNConvolutionDescriptor) SetBatchNormalizationParametersForInferenceWithMeanVarianceGammaBetaEpsilon(mean objectivec.IObject, variance objectivec.IObject, gamma objectivec.IObject, beta objectivec.IObject, epsilon float32) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setBatchNormalizationParametersForInferenceWithMean:variance:gamma:beta:epsilon:"), mean, variance, gamma, beta, epsilon)
}/* debug [instance_methods/method]: SetBatchNormalizationParametersForInferenceWithMeanVarianceGammaBetaEpsilon */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for CNNConvolutionDescriptor */

// The output stride (downsampling factor) in the y dimension.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnconvolutiondescriptor/1648847-strideinpixelsy
func (c_ CNNConvolutionDescriptor) StrideInPixelsY() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](c_.ID, objc.Sel("strideInPixelsY"))
	return rv
}/* debug [instance_properties/getter]: strideInPixelsY */


// The output stride (downsampling factor) in the y dimension.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnconvolutiondescriptor/1648847-strideinpixelsy
func (c_ CNNConvolutionDescriptor) SetStrideInPixelsY(value objectivec.IObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setStrideInPixelsY:"), value)
}/* debug [instance_properties/setter]: strideInPixelsY */


// The number of groups that the input and output channels are divided into.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnconvolutiondescriptor/1648849-groups
func (c_ CNNConvolutionDescriptor) Groups() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](c_.ID, objc.Sel("groups"))
	return rv
}/* debug [instance_properties/getter]: groups */


// The number of groups that the input and output channels are divided into.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnconvolutiondescriptor/1648849-groups
func (c_ CNNConvolutionDescriptor) SetGroups(value objectivec.IObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setGroups:"), value)
}/* debug [instance_properties/setter]: groups */


// The number of feature channels per pixel in the output image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnconvolutiondescriptor/1648852-outputfeaturechannels
func (c_ CNNConvolutionDescriptor) OutputFeatureChannels() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](c_.ID, objc.Sel("outputFeatureChannels"))
	return rv
}/* debug [instance_properties/getter]: outputFeatureChannels */


// The number of feature channels per pixel in the output image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnconvolutiondescriptor/1648852-outputfeaturechannels
func (c_ CNNConvolutionDescriptor) SetOutputFeatureChannels(value objectivec.IObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setOutputFeatureChannels:"), value)
}/* debug [instance_properties/setter]: outputFeatureChannels */


// The height of the kernel window.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnconvolutiondescriptor/1648904-kernelheight
func (c_ CNNConvolutionDescriptor) KernelHeight() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](c_.ID, objc.Sel("kernelHeight"))
	return rv
}/* debug [instance_properties/getter]: kernelHeight */


// The height of the kernel window.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnconvolutiondescriptor/1648904-kernelheight
func (c_ CNNConvolutionDescriptor) SetKernelHeight(value objectivec.IObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setKernelHeight:"), value)
}/* debug [instance_properties/setter]: kernelHeight */


// The output stride (downsampling factor) in the x dimension.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnconvolutiondescriptor/1648908-strideinpixelsx
func (c_ CNNConvolutionDescriptor) StrideInPixelsX() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](c_.ID, objc.Sel("strideInPixelsX"))
	return rv
}/* debug [instance_properties/getter]: strideInPixelsX */


// The output stride (downsampling factor) in the x dimension.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnconvolutiondescriptor/1648908-strideinpixelsx
func (c_ CNNConvolutionDescriptor) SetStrideInPixelsX(value objectivec.IObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setStrideInPixelsX:"), value)
}/* debug [instance_properties/setter]: strideInPixelsX */


// The number of feature channels per pixel in the input image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnconvolutiondescriptor/1648934-inputfeaturechannels
func (c_ CNNConvolutionDescriptor) InputFeatureChannels() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](c_.ID, objc.Sel("inputFeatureChannels"))
	return rv
}/* debug [instance_properties/getter]: inputFeatureChannels */


// The number of feature channels per pixel in the input image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnconvolutiondescriptor/1648934-inputfeaturechannels
func (c_ CNNConvolutionDescriptor) SetInputFeatureChannels(value objectivec.IObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setInputFeatureChannels:"), value)
}/* debug [instance_properties/setter]: inputFeatureChannels */


// The width of the kernel window.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnconvolutiondescriptor/1648959-kernelwidth
func (c_ CNNConvolutionDescriptor) KernelWidth() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](c_.ID, objc.Sel("kernelWidth"))
	return rv
}/* debug [instance_properties/getter]: kernelWidth */


// The width of the kernel window.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnconvolutiondescriptor/1648959-kernelwidth
func (c_ CNNConvolutionDescriptor) SetKernelWidth(value objectivec.IObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setKernelWidth:"), value)
}/* debug [instance_properties/setter]: kernelWidth */


// The neuron filter to be applied as part of the convolution operation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnconvolutiondescriptor/1829442-neuron
func (c_ CNNConvolutionDescriptor) Neuron() IMPSCNNNeuron {
	rv := objc.Send[CNNNeuron](c_.ID, objc.Sel("neuron"))
	return rv
}/* debug [instance_properties/getter]: neuron */


// The neuron filter to be applied as part of the convolution operation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnconvolutiondescriptor/1829442-neuron
func (c_ CNNConvolutionDescriptor) SetNeuron(value IMPSCNNNeuron) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setNeuron:"), value)
}/* debug [instance_properties/setter]: neuron */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnconvolutiondescriptor/2867154-supportssecurecoding
func (c_ CNNConvolutionDescriptor) SupportsSecureCoding() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](c_.ID, objc.Sel("supportsSecureCoding"))
	return rv
}/* debug [instance_properties/getter]: supportsSecureCoding */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnconvolutiondescriptor/2867154-supportssecurecoding
func (c_ CNNConvolutionDescriptor) SetSupportsSecureCoding(value objectivec.IObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setSupportsSecureCoding:"), value)
}/* debug [instance_properties/setter]: supportsSecureCoding */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnconvolutiondescriptor/2881195-dilationratex
func (c_ CNNConvolutionDescriptor) DilationRateX() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](c_.ID, objc.Sel("dilationRateX"))
	return rv
}/* debug [instance_properties/getter]: dilationRateX */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnconvolutiondescriptor/2881195-dilationratex
func (c_ CNNConvolutionDescriptor) SetDilationRateX(value objectivec.IObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setDilationRateX:"), value)
}/* debug [instance_properties/setter]: dilationRateX */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnconvolutiondescriptor/2881196-dilationratey
func (c_ CNNConvolutionDescriptor) DilationRateY() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](c_.ID, objc.Sel("dilationRateY"))
	return rv
}/* debug [instance_properties/getter]: dilationRateY */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnconvolutiondescriptor/2881196-dilationratey
func (c_ CNNConvolutionDescriptor) SetDilationRateY(value objectivec.IObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setDilationRateY:"), value)
}/* debug [instance_properties/setter]: dilationRateY */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnconvolutiondescriptor/2953957-fusedneurondescriptor
func (c_ CNNConvolutionDescriptor) FusedNeuronDescriptor() IMPSNNNeuronDescriptor {
	rv := objc.Send[NeuronDescriptor](c_.ID, objc.Sel("fusedNeuronDescriptor"))
	return rv
}/* debug [instance_properties/getter]: fusedNeuronDescriptor */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnconvolutiondescriptor/2953957-fusedneurondescriptor
func (c_ CNNConvolutionDescriptor) SetFusedNeuronDescriptor(value IMPSNNNeuronDescriptor) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setFusedNeuronDescriptor:"), value)
}/* debug [instance_properties/setter]: fusedNeuronDescriptor */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MPSCNNConvolutionDescriptor */


