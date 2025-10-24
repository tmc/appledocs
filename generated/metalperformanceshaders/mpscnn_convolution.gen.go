// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MPSCNNConvolution */


/* debug [class_header]: Header for MPSCNNConvolution */
// The class instance for the [CNNConvolution] class.
var (
	CNNConvolutionClass     _CNNConvolutionClass
	CNNConvolutionClassOnce sync.Once
)

func getCNNConvolutionClass() _CNNConvolutionClass {
	CNNConvolutionClassOnce.Do(func() {
		CNNConvolutionClass = _CNNConvolutionClass{objc.GetClass("MPSCNNConvolution")}
	})
	return CNNConvolutionClass
}

type _CNNConvolutionClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for CNNConvolution */
// An interface definition for the [CNNConvolution] class.
type ICNNConvolution interface {
	ICNNKernel
	
/* debug [class_interface_properties]: Properties for CNNConvolution */
	// properties:
	InputFeatureChannels() objectivec.IObject
	SetInputFeatureChannels(value objectivec.IObject)
	Groups() objectivec.IObject
	SetGroups(value objectivec.IObject)
	OutputFeatureChannels() objectivec.IObject
	SetOutputFeatureChannels(value objectivec.IObject)
	Neuron() IMPSCNNNeuron
	SetNeuron(value IMPSCNNNeuron)
	SubPixelScaleFactor() objectivec.IObject
	SetSubPixelScaleFactor(value objectivec.IObject)
	NeuronType() CNNNeuronType get /* not a class type */
	SetNeuronType(value CNNNeuronType get /* not a class type */)
	NeuronParameterA() objectivec.IObject
	SetNeuronParameterA(value objectivec.IObject)
	NeuronParameterB() objectivec.IObject
	SetNeuronParameterB(value objectivec.IObject)
	ChannelMultiplier() objectivec.IObject
	SetChannelMultiplier(value objectivec.IObject)
	NeuronParameterC() objectivec.IObject
	SetNeuronParameterC(value objectivec.IObject)
	AccumulatorPrecisionOption() ConvolutionAccumulatorPrecisionOption get set /* not a class type */
	SetAccumulatorPrecisionOption(value ConvolutionAccumulatorPrecisionOption get set /* not a class type */)
	DataSource() CNNConvolutionDataSource get /* not a class type */
	SetDataSource(value CNNConvolutionDataSource get /* not a class type */)
	FusedNeuronDescriptor() IMPSNNNeuronDescriptor
	SetFusedNeuronDescriptor(value IMPSNNNeuronDescriptor)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for CNNConvolution */
	// methods:
	ResultStateBatch()
	ResultStateBatchForSourceImageSourceStatesDestinationImage(sourceImage ImageBatch /* not a class type */, sourceStates StateBatch /* not a class type */, destinationImage ImageBatch /* not a class type */) CNNConvolutionGradientStateBatch /* not a class type */
	ResultState()
	ResultStateForSourceImageSourceStatesDestinationImage(sourceImage IImage, sourceStates unsafe.Pointer, destinationImage IImage) ICNNConvolutionGradientState
	TemporaryResultState()
	TemporaryResultStateForCommandBufferSourceImageSourceStatesDestinationImage(commandBuffer unsafe.Pointer, sourceImage IImage, sourceStates unsafe.Pointer, destinationImage IImage) ICNNConvolutionGradientState
	TemporaryResultStateBatch()
	TemporaryResultStateBatchForCommandBufferSourceImageSourceStatesDestinationImage(commandBuffer unsafe.Pointer, sourceImage ImageBatch /* not a class type */, sourceStates StateBatch /* not a class type */, destinationImage ImageBatch /* not a class type */) CNNConvolutionGradientStateBatch /* not a class type */
	ExportWeightsAndBiases()
	ExportWeightsAndBiasesWithCommandBufferResultStateCanBeTemporary(commandBuffer unsafe.Pointer, resultStateCanBeTemporary bool) ICNNConvolutionWeightsAndBiasesState
	ReloadWeightsAndBiasesWithCommandBufferState(commandBuffer unsafe.Pointer, state ICNNConvolutionWeightsAndBiasesState)
	ReloadWeightsAndBiasesFromDataSource()
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for CNNConvolution */
// Alloc allocates a new instance without initialization.
func (cc _CNNConvolutionClass) Alloc() CNNConvolution {
	rv := objc.Send[CNNConvolution](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _CNNConvolutionClass) New() CNNConvolution {
	rv := objc.Send[CNNConvolution](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CNNConvolution) Init() CNNConvolution {
	rv := objc.Send[CNNConvolution](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CNNConvolution) Autorelease() CNNConvolution {
	rv := objc.Send[CNNConvolution](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCNNConvolution creates a new CNNConvolution instance.
func NewCNNConvolution() CNNConvolution {
	return getCNNConvolutionClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for CNNConvolution */
// A convolution kernel that convolves the input image with a set of filters, with each producing one feature map in the output image.
//
// The attributes of a convolution operation are described by an object.


// A convolution kernel that convolves the input image with a set of filters, with each producing one feature map in the output image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNConvolution
type CNNConvolution struct {
	CNNKernel
}

// CNNConvolutionFrom constructs a [CNNConvolution] from an unsafe.Pointer.
//
// A convolution kernel that convolves the input image with a set of filters, with each producing one feature map in the output image.
func CNNConvolutionFrom(ptr unsafe.Pointer) CNNConvolution {
	return CNNConvolution{
		CNNKernel: CNNKernelFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for CNNConvolution */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNConvolution/init(coder:device:)
func NewCNNConvolutionWithCoderDevice(aDecoder foundation.Coder, device unsafe.Pointer) CNNConvolution {
	instance := getCNNConvolutionClass().Alloc()
	rv := objc.Send[CNNConvolution](instance.ID, objc.Sel("initWithCoder:device:"), aDecoder, device)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewCNNConvolutionWithCoderDevice */


// Initializes a convolution kernel.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNConvolution/init(device:convolutionDescriptor:kernelWeights:biasTerms:flags:)
func NewCNNConvolutionWithDeviceConvolutionDescriptorKernelWeightsBiasTermsFlags(device unsafe.Pointer, convolutionDescriptor IMPSCNNConvolutionDescriptor, kernelWeights objectivec.IObject, biasTerms objectivec.IObject, flags CNNConvolutionFlags) CNNConvolution {
	instance := getCNNConvolutionClass().Alloc()
	rv := objc.Send[CNNConvolution](instance.ID, objc.Sel("initWithDevice:convolutionDescriptor:kernelWeights:biasTerms:flags:"), device, convolutionDescriptor, kernelWeights, biasTerms, flags)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewCNNConvolutionWithDeviceConvolutionDescriptorKernelWeightsBiasTermsFlags */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNConvolution/init(device:weights:)
func NewCNNConvolutionWithDeviceWeights(device unsafe.Pointer, weights unsafe.Pointer) CNNConvolution {
	instance := getCNNConvolutionClass().Alloc()
	rv := objc.Send[CNNConvolution](instance.ID, objc.Sel("initWithDevice:weights:"), device, weights)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewCNNConvolutionWithDeviceWeights */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for CNNConvolution */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for CNNConvolution */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for CNNConvolution */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnconvolution/2947881-resultstatebatch
func (c_ CNNConvolution) ResultStateBatch() {
	objc.Send[objc.ID](c_.ID, objc.Sel("resultStateBatch"))
}/* debug [instance_methods/method]: ResultStateBatch */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnconvolution/2947881-resultstatebatchforsourceimage
func (c_ CNNConvolution) ResultStateBatchForSourceImageSourceStatesDestinationImage(sourceImage ImageBatch /* not a class type */, sourceStates StateBatch /* not a class type */, destinationImage ImageBatch /* not a class type */) CNNConvolutionGradientStateBatch /* not a class type */ {
	rv := objc.Send[CNNConvolutionGradientStateBatch](c_.ID, objc.Sel("resultStateBatchForSourceImage:sourceStates:destinationImage:"), sourceImage, sourceStates, destinationImage)
	return rv
}/* debug [instance_methods/method]: ResultStateBatchForSourceImageSourceStatesDestinationImage */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnconvolution/2947883-resultstate
func (c_ CNNConvolution) ResultState() {
	objc.Send[objc.ID](c_.ID, objc.Sel("resultState"))
}/* debug [instance_methods/method]: ResultState */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnconvolution/2947883-resultstateforsourceimage
func (c_ CNNConvolution) ResultStateForSourceImageSourceStatesDestinationImage(sourceImage IImage, sourceStates unsafe.Pointer, destinationImage IImage) ICNNConvolutionGradientState {
	rv := objc.Send[CNNConvolutionGradientState](c_.ID, objc.Sel("resultStateForSourceImage:sourceStates:destinationImage:"), sourceImage, sourceStates, destinationImage)
	return rv
}/* debug [instance_methods/method]: ResultStateForSourceImageSourceStatesDestinationImage */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnconvolution/2947885-temporaryresultstate
func (c_ CNNConvolution) TemporaryResultState() {
	objc.Send[objc.ID](c_.ID, objc.Sel("temporaryResultState"))
}/* debug [instance_methods/method]: TemporaryResultState */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnconvolution/2947885-temporaryresultstateforcommandbu
func (c_ CNNConvolution) TemporaryResultStateForCommandBufferSourceImageSourceStatesDestinationImage(commandBuffer unsafe.Pointer, sourceImage IImage, sourceStates unsafe.Pointer, destinationImage IImage) ICNNConvolutionGradientState {
	rv := objc.Send[CNNConvolutionGradientState](c_.ID, objc.Sel("temporaryResultStateForCommandBuffer:sourceImage:sourceStates:destinationImage:"), commandBuffer, sourceImage, sourceStates, destinationImage)
	return rv
}/* debug [instance_methods/method]: TemporaryResultStateForCommandBufferSourceImageSourceStatesDestinationImage */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnconvolution/2947886-temporaryresultstatebatch
func (c_ CNNConvolution) TemporaryResultStateBatch() {
	objc.Send[objc.ID](c_.ID, objc.Sel("temporaryResultStateBatch"))
}/* debug [instance_methods/method]: TemporaryResultStateBatch */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnconvolution/2947886-temporaryresultstatebatchforcomm
func (c_ CNNConvolution) TemporaryResultStateBatchForCommandBufferSourceImageSourceStatesDestinationImage(commandBuffer unsafe.Pointer, sourceImage ImageBatch /* not a class type */, sourceStates StateBatch /* not a class type */, destinationImage ImageBatch /* not a class type */) CNNConvolutionGradientStateBatch /* not a class type */ {
	rv := objc.Send[CNNConvolutionGradientStateBatch](c_.ID, objc.Sel("temporaryResultStateBatchForCommandBuffer:sourceImage:sourceStates:destinationImage:"), commandBuffer, sourceImage, sourceStates, destinationImage)
	return rv
}/* debug [instance_methods/method]: TemporaryResultStateBatchForCommandBufferSourceImageSourceStatesDestinationImage */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnconvolution/2953001-exportweightsandbiases
func (c_ CNNConvolution) ExportWeightsAndBiases() {
	objc.Send[objc.ID](c_.ID, objc.Sel("exportWeightsAndBiases"))
}/* debug [instance_methods/method]: ExportWeightsAndBiases */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnconvolution/2953001-exportweightsandbiaseswithcomman
func (c_ CNNConvolution) ExportWeightsAndBiasesWithCommandBufferResultStateCanBeTemporary(commandBuffer unsafe.Pointer, resultStateCanBeTemporary bool) ICNNConvolutionWeightsAndBiasesState {
	rv := objc.Send[CNNConvolutionWeightsAndBiasesState](c_.ID, objc.Sel("exportWeightsAndBiasesWithCommandBuffer:resultStateCanBeTemporary:"), commandBuffer, resultStateCanBeTemporary)
	return rv
}/* debug [instance_methods/method]: ExportWeightsAndBiasesWithCommandBufferResultStateCanBeTemporary */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnconvolution/2953962-reloadweightsandbiaseswithcomman
func (c_ CNNConvolution) ReloadWeightsAndBiasesWithCommandBufferState(commandBuffer unsafe.Pointer, state ICNNConvolutionWeightsAndBiasesState) {
	objc.Send[objc.ID](c_.ID, objc.Sel("reloadWeightsAndBiasesWithCommandBuffer:state:"), commandBuffer, state)
}/* debug [instance_methods/method]: ReloadWeightsAndBiasesWithCommandBufferState */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnconvolution/2966657-reloadweightsandbiasesfromdataso
func (c_ CNNConvolution) ReloadWeightsAndBiasesFromDataSource() {
	objc.Send[objc.ID](c_.ID, objc.Sel("reloadWeightsAndBiasesFromDataSource"))
}/* debug [instance_methods/method]: ReloadWeightsAndBiasesFromDataSource */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for CNNConvolution */

// The number of feature channels per pixel in the input image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnconvolution/1845268-inputfeaturechannels
func (c_ CNNConvolution) InputFeatureChannels() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](c_.ID, objc.Sel("inputFeatureChannels"))
	return rv
}/* debug [instance_properties/getter]: inputFeatureChannels */


// The number of feature channels per pixel in the input image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnconvolution/1845268-inputfeaturechannels
func (c_ CNNConvolution) SetInputFeatureChannels(value objectivec.IObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setInputFeatureChannels:"), value)
}/* debug [instance_properties/setter]: inputFeatureChannels */


// The number of groups that the input and output channels are divided into.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnconvolution/1845269-groups
func (c_ CNNConvolution) Groups() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](c_.ID, objc.Sel("groups"))
	return rv
}/* debug [instance_properties/getter]: groups */


// The number of groups that the input and output channels are divided into.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnconvolution/1845269-groups
func (c_ CNNConvolution) SetGroups(value objectivec.IObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setGroups:"), value)
}/* debug [instance_properties/setter]: groups */


// The number of feature channels per pixel in the output image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnconvolution/1845271-outputfeaturechannels
func (c_ CNNConvolution) OutputFeatureChannels() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](c_.ID, objc.Sel("outputFeatureChannels"))
	return rv
}/* debug [instance_properties/getter]: outputFeatureChannels */


// The number of feature channels per pixel in the output image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnconvolution/1845271-outputfeaturechannels
func (c_ CNNConvolution) SetOutputFeatureChannels(value objectivec.IObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setOutputFeatureChannels:"), value)
}/* debug [instance_properties/setter]: outputFeatureChannels */


// The neuron filter to be applied as part of the convolution operation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnconvolution/1845274-neuron
func (c_ CNNConvolution) Neuron() IMPSCNNNeuron {
	rv := objc.Send[CNNNeuron](c_.ID, objc.Sel("neuron"))
	return rv
}/* debug [instance_properties/getter]: neuron */


// The neuron filter to be applied as part of the convolution operation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnconvolution/1845274-neuron
func (c_ CNNConvolution) SetNeuron(value IMPSCNNNeuron) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setNeuron:"), value)
}/* debug [instance_properties/setter]: neuron */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnconvolution/2873341-subpixelscalefactor
func (c_ CNNConvolution) SubPixelScaleFactor() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](c_.ID, objc.Sel("subPixelScaleFactor"))
	return rv
}/* debug [instance_properties/getter]: subPixelScaleFactor */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnconvolution/2873341-subpixelscalefactor
func (c_ CNNConvolution) SetSubPixelScaleFactor(value objectivec.IObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setSubPixelScaleFactor:"), value)
}/* debug [instance_properties/setter]: subPixelScaleFactor */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnconvolution/2875190-neurontype
func (c_ CNNConvolution) NeuronType() CNNNeuronType get /* not a class type */ {
	rv := objc.Send[objc.ID](c_.ID, objc.Sel("neuronType"))
	return rv
}/* debug [instance_properties/getter]: neuronType */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnconvolution/2875190-neurontype
func (c_ CNNConvolution) SetNeuronType(value CNNNeuronType get /* not a class type */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setNeuronType:"), value)
}/* debug [instance_properties/setter]: neuronType */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnconvolution/2875214-neuronparametera
func (c_ CNNConvolution) NeuronParameterA() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](c_.ID, objc.Sel("neuronParameterA"))
	return rv
}/* debug [instance_properties/getter]: neuronParameterA */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnconvolution/2875214-neuronparametera
func (c_ CNNConvolution) SetNeuronParameterA(value objectivec.IObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setNeuronParameterA:"), value)
}/* debug [instance_properties/setter]: neuronParameterA */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnconvolution/2875218-neuronparameterb
func (c_ CNNConvolution) NeuronParameterB() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](c_.ID, objc.Sel("neuronParameterB"))
	return rv
}/* debug [instance_properties/getter]: neuronParameterB */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnconvolution/2875218-neuronparameterb
func (c_ CNNConvolution) SetNeuronParameterB(value objectivec.IObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setNeuronParameterB:"), value)
}/* debug [instance_properties/setter]: neuronParameterB */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnconvolution/2919729-channelmultiplier
func (c_ CNNConvolution) ChannelMultiplier() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](c_.ID, objc.Sel("channelMultiplier"))
	return rv
}/* debug [instance_properties/getter]: channelMultiplier */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnconvolution/2919729-channelmultiplier
func (c_ CNNConvolution) SetChannelMultiplier(value objectivec.IObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setChannelMultiplier:"), value)
}/* debug [instance_properties/setter]: channelMultiplier */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnconvolution/2935626-neuronparameterc
func (c_ CNNConvolution) NeuronParameterC() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](c_.ID, objc.Sel("neuronParameterC"))
	return rv
}/* debug [instance_properties/getter]: neuronParameterC */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnconvolution/2935626-neuronparameterc
func (c_ CNNConvolution) SetNeuronParameterC(value objectivec.IObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setNeuronParameterC:"), value)
}/* debug [instance_properties/setter]: neuronParameterC */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnconvolution/2942410-accumulatorprecisionoption
func (c_ CNNConvolution) AccumulatorPrecisionOption() ConvolutionAccumulatorPrecisionOption get set /* not a class type */ {
	rv := objc.Send[objc.ID](c_.ID, objc.Sel("accumulatorPrecisionOption"))
	return rv
}/* debug [instance_properties/getter]: accumulatorPrecisionOption */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnconvolution/2942410-accumulatorprecisionoption
func (c_ CNNConvolution) SetAccumulatorPrecisionOption(value ConvolutionAccumulatorPrecisionOption get set /* not a class type */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setAccumulatorPrecisionOption:"), value)
}/* debug [instance_properties/setter]: accumulatorPrecisionOption */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnconvolution/2953961-datasource
func (c_ CNNConvolution) DataSource() CNNConvolutionDataSource get /* not a class type */ {
	rv := objc.Send[objc.ID](c_.ID, objc.Sel("dataSource"))
	return rv
}/* debug [instance_properties/getter]: dataSource */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnconvolution/2953961-datasource
func (c_ CNNConvolution) SetDataSource(value CNNConvolutionDataSource get /* not a class type */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setDataSource:"), value)
}/* debug [instance_properties/setter]: dataSource */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnconvolution/3013776-fusedneurondescriptor
func (c_ CNNConvolution) FusedNeuronDescriptor() IMPSNNNeuronDescriptor {
	rv := objc.Send[NeuronDescriptor](c_.ID, objc.Sel("fusedNeuronDescriptor"))
	return rv
}/* debug [instance_properties/getter]: fusedNeuronDescriptor */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnconvolution/3013776-fusedneurondescriptor
func (c_ CNNConvolution) SetFusedNeuronDescriptor(value IMPSNNNeuronDescriptor) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setFusedNeuronDescriptor:"), value)
}/* debug [instance_properties/setter]: fusedNeuronDescriptor */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MPSCNNConvolution */


