// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)





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





// An interface definition for the [CNNConvolutionDescriptor] class.
type ICNNConvolutionDescriptor interface {
	objectivec.IObject
	

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


	

	// methods:
	Encode()
	SetBatchNormalizationParametersForInferenceWithMean()
	SetBatchNormalizationParametersForInferenceWithMeanVarianceGammaBetaEpsilon(mean objectivec.IObject, variance objectivec.IObject, gamma objectivec.IObject, beta objectivec.IObject, epsilon float32)


}





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






// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNConvolutionDescriptor/init(kernelWidth:kernelHeight:inputFeatureChannels:outputFeatureChannels:)
func NewCNNConvolutionDescriptorCnnConvolutionDescriptorWithKernelWidthKernelHeightInputFeatureChannelsOutputFeatureChannels(kernelWidth uint, kernelHeight uint, inputFeatureChannels uint, outputFeatureChannels uint) CNNConvolutionDescriptor {
	rv := objc.Send[CNNConvolutionDescriptor](objc.ID(getCNNConvolutionDescriptorClass().class), objc.Sel("cnnConvolutionDescriptorWithKernelWidth:kernelHeight:inputFeatureChannels:outputFeatureChannels:"), kernelWidth, kernelHeight, inputFeatureChannels, outputFeatureChannels)
	return rv
}


// Creates a convolution descriptor with an optional neuron filter.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNConvolutionDescriptor/init(kernelWidth:kernelHeight:inputFeatureChannels:outputFeatureChannels:neuronFilter:)
func NewCNNConvolutionDescriptorCnnConvolutionDescriptorWithKernelWidthKernelHeightInputFeatureChannelsOutputFeatureChannelsNeuronFilter(kernelWidth uint, kernelHeight uint, inputFeatureChannels uint, outputFeatureChannels uint, neuronFilter IMPSCNNNeuron) CNNConvolutionDescriptor {
	rv := objc.Send[CNNConvolutionDescriptor](objc.ID(getCNNConvolutionDescriptorClass().class), objc.Sel("cnnConvolutionDescriptorWithKernelWidth:kernelHeight:inputFeatureChannels:outputFeatureChannels:neuronFilter:"), kernelWidth, kernelHeight, inputFeatureChannels, outputFeatureChannels, neuronFilter)
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNConvolutionDescriptor/init(coder:)
func NewCNNConvolutionDescriptorWithCoder(aDecoder foundation.Coder) CNNConvolutionDescriptor {
	instance := getCNNConvolutionDescriptorClass().Alloc()
	rv := objc.Send[CNNConvolutionDescriptor](instance.ID, objc.Sel("initWithCoder:"), aDecoder)
	rv.Autorelease()
	return rv
}







// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNConvolutionDescriptor/init(kernelWidth:kernelHeight:inputFeatureChannels:outputFeatureChannels:)
func (cc _CNNConvolutionDescriptorClass) CnnConvolutionDescriptorWithKernelWidthKernelHeightInputFeatureChannelsOutputFeatureChannels(kernelWidth uint, kernelHeight uint, inputFeatureChannels uint, outputFeatureChannels uint) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(cc.class), objc.Sel("cnnConvolutionDescriptorWithKernelWidth:kernelHeight:inputFeatureChannels:outputFeatureChannels:"), kernelWidth, kernelHeight, inputFeatureChannels, outputFeatureChannels)
	return rv
}


// Creates a convolution descriptor with an optional neuron filter.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNConvolutionDescriptor/init(kernelWidth:kernelHeight:inputFeatureChannels:outputFeatureChannels:neuronFilter:)
func (cc _CNNConvolutionDescriptorClass) CnnConvolutionDescriptorWithKernelWidthKernelHeightInputFeatureChannelsOutputFeatureChannelsNeuronFilter(kernelWidth uint, kernelHeight uint, inputFeatureChannels uint, outputFeatureChannels uint, neuronFilter IMPSCNNNeuron) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(cc.class), objc.Sel("cnnConvolutionDescriptorWithKernelWidth:kernelHeight:inputFeatureChannels:outputFeatureChannels:neuronFilter:"), kernelWidth, kernelHeight, inputFeatureChannels, outputFeatureChannels, neuronFilter)
	return rv
}







// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnconvolutiondescriptor/2867154-supportssecurecoding
func (cc _CNNConvolutionDescriptorClass) SupportsSecureCoding() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(cc.class), objc.Sel("supportsSecureCoding"))
	return rv
}






// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnconvolutiondescriptor/2866972-encode
func (c_ CNNConvolutionDescriptor) Encode() {
	objc.Send[objc.ID](c_.ID, objc.Sel("encode"))
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnconvolutiondescriptor/2867057-setbatchnormalizationparametersf
func (c_ CNNConvolutionDescriptor) SetBatchNormalizationParametersForInferenceWithMean() {
	objc.Send[objc.ID](c_.ID, objc.Sel("setBatchNormalizationParametersForInferenceWithMean"))
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNConvolutionDescriptor/encode(with:)
func (c_ CNNConvolutionDescriptor) EncodeWithCoder(aCoder foundation.Coder) {
	objc.Send[objc.ID](c_.ID, objc.Sel("encodeWithCoder:"), aCoder)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNConvolutionDescriptor/setBatchNormalizationParametersForInferenceWithMean(_:variance:gamma:beta:epsilon:)
func (c_ CNNConvolutionDescriptor) SetBatchNormalizationParametersForInferenceWithMeanVarianceGammaBetaEpsilon(mean objectivec.IObject, variance objectivec.IObject, gamma objectivec.IObject, beta objectivec.IObject, epsilon float32) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setBatchNormalizationParametersForInferenceWithMean:variance:gamma:beta:epsilon:"), mean, variance, gamma, beta, epsilon)
}







// The output stride (downsampling factor) in the y dimension.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnconvolutiondescriptor/1648847-strideinpixelsy
func (c_ CNNConvolutionDescriptor) StrideInPixelsY() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](c_.ID, objc.Sel("strideInPixelsY"))
	return rv
}


// The output stride (downsampling factor) in the y dimension.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnconvolutiondescriptor/1648847-strideinpixelsy
func (c_ CNNConvolutionDescriptor) SetStrideInPixelsY(value objectivec.IObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setStrideInPixelsY:"), value)
}


// The number of groups that the input and output channels are divided into.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnconvolutiondescriptor/1648849-groups
func (c_ CNNConvolutionDescriptor) Groups() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](c_.ID, objc.Sel("groups"))
	return rv
}


// The number of groups that the input and output channels are divided into.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnconvolutiondescriptor/1648849-groups
func (c_ CNNConvolutionDescriptor) SetGroups(value objectivec.IObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setGroups:"), value)
}


// The number of feature channels per pixel in the output image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnconvolutiondescriptor/1648852-outputfeaturechannels
func (c_ CNNConvolutionDescriptor) OutputFeatureChannels() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](c_.ID, objc.Sel("outputFeatureChannels"))
	return rv
}


// The number of feature channels per pixel in the output image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnconvolutiondescriptor/1648852-outputfeaturechannels
func (c_ CNNConvolutionDescriptor) SetOutputFeatureChannels(value objectivec.IObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setOutputFeatureChannels:"), value)
}


// The height of the kernel window.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnconvolutiondescriptor/1648904-kernelheight
func (c_ CNNConvolutionDescriptor) KernelHeight() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](c_.ID, objc.Sel("kernelHeight"))
	return rv
}


// The height of the kernel window.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnconvolutiondescriptor/1648904-kernelheight
func (c_ CNNConvolutionDescriptor) SetKernelHeight(value objectivec.IObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setKernelHeight:"), value)
}


// The output stride (downsampling factor) in the x dimension.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnconvolutiondescriptor/1648908-strideinpixelsx
func (c_ CNNConvolutionDescriptor) StrideInPixelsX() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](c_.ID, objc.Sel("strideInPixelsX"))
	return rv
}


// The output stride (downsampling factor) in the x dimension.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnconvolutiondescriptor/1648908-strideinpixelsx
func (c_ CNNConvolutionDescriptor) SetStrideInPixelsX(value objectivec.IObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setStrideInPixelsX:"), value)
}


// The number of feature channels per pixel in the input image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnconvolutiondescriptor/1648934-inputfeaturechannels
func (c_ CNNConvolutionDescriptor) InputFeatureChannels() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](c_.ID, objc.Sel("inputFeatureChannels"))
	return rv
}


// The number of feature channels per pixel in the input image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnconvolutiondescriptor/1648934-inputfeaturechannels
func (c_ CNNConvolutionDescriptor) SetInputFeatureChannels(value objectivec.IObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setInputFeatureChannels:"), value)
}


// The width of the kernel window.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnconvolutiondescriptor/1648959-kernelwidth
func (c_ CNNConvolutionDescriptor) KernelWidth() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](c_.ID, objc.Sel("kernelWidth"))
	return rv
}


// The width of the kernel window.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnconvolutiondescriptor/1648959-kernelwidth
func (c_ CNNConvolutionDescriptor) SetKernelWidth(value objectivec.IObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setKernelWidth:"), value)
}


// The neuron filter to be applied as part of the convolution operation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnconvolutiondescriptor/1829442-neuron
func (c_ CNNConvolutionDescriptor) Neuron() IMPSCNNNeuron {
	rv := objc.Send[CNNNeuron](c_.ID, objc.Sel("neuron"))
	return rv
}


// The neuron filter to be applied as part of the convolution operation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnconvolutiondescriptor/1829442-neuron
func (c_ CNNConvolutionDescriptor) SetNeuron(value IMPSCNNNeuron) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setNeuron:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnconvolutiondescriptor/2867154-supportssecurecoding
func (c_ CNNConvolutionDescriptor) SupportsSecureCoding() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](c_.ID, objc.Sel("supportsSecureCoding"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnconvolutiondescriptor/2867154-supportssecurecoding
func (c_ CNNConvolutionDescriptor) SetSupportsSecureCoding(value objectivec.IObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setSupportsSecureCoding:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnconvolutiondescriptor/2881195-dilationratex
func (c_ CNNConvolutionDescriptor) DilationRateX() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](c_.ID, objc.Sel("dilationRateX"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnconvolutiondescriptor/2881195-dilationratex
func (c_ CNNConvolutionDescriptor) SetDilationRateX(value objectivec.IObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setDilationRateX:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnconvolutiondescriptor/2881196-dilationratey
func (c_ CNNConvolutionDescriptor) DilationRateY() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](c_.ID, objc.Sel("dilationRateY"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnconvolutiondescriptor/2881196-dilationratey
func (c_ CNNConvolutionDescriptor) SetDilationRateY(value objectivec.IObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setDilationRateY:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnconvolutiondescriptor/2953957-fusedneurondescriptor
func (c_ CNNConvolutionDescriptor) FusedNeuronDescriptor() IMPSNNNeuronDescriptor {
	rv := objc.Send[NeuronDescriptor](c_.ID, objc.Sel("fusedNeuronDescriptor"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnconvolutiondescriptor/2953957-fusedneurondescriptor
func (c_ CNNConvolutionDescriptor) SetFusedNeuronDescriptor(value IMPSNNNeuronDescriptor) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setFusedNeuronDescriptor:"), value)
}







