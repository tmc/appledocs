// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MPSCNNConvolutionTranspose */


/* debug [class_header]: Header for MPSCNNConvolutionTranspose */
// The class instance for the [CNNConvolutionTranspose] class.
var (
	CNNConvolutionTransposeClass     _CNNConvolutionTransposeClass
	CNNConvolutionTransposeClassOnce sync.Once
)

func getCNNConvolutionTransposeClass() _CNNConvolutionTransposeClass {
	CNNConvolutionTransposeClassOnce.Do(func() {
		CNNConvolutionTransposeClass = _CNNConvolutionTransposeClass{objc.GetClass("MPSCNNConvolutionTranspose")}
	})
	return CNNConvolutionTransposeClass
}

type _CNNConvolutionTransposeClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for CNNConvolutionTranspose */
// An interface definition for the [CNNConvolutionTranspose] class.
type ICNNConvolutionTranspose interface {
	ICNNKernel
	
/* debug [class_interface_properties]: Properties for CNNConvolutionTranspose */
	// properties:
	OutputFeatureChannels() objectivec.IObject
	SetOutputFeatureChannels(value objectivec.IObject)
	KernelOffsetY() objectivec.IObject
	SetKernelOffsetY(value objectivec.IObject)
	Groups() objectivec.IObject
	SetGroups(value objectivec.IObject)
	InputFeatureChannels() objectivec.IObject
	SetInputFeatureChannels(value objectivec.IObject)
	KernelOffsetX() objectivec.IObject
	SetKernelOffsetX(value objectivec.IObject)
	AccumulatorPrecisionOption() ConvolutionAccumulatorPrecisionOption get set /* not a class type */
	SetAccumulatorPrecisionOption(value ConvolutionAccumulatorPrecisionOption get set /* not a class type */)
	DataSource() CNNConvolutionDataSource get /* not a class type */
	SetDataSource(value CNNConvolutionDataSource get /* not a class type */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for CNNConvolutionTranspose */
	// methods:
	EncodeBatch()
	EncodeBatchToCommandBufferSourceImagesConvolutionGradientStates(commandBuffer unsafe.Pointer, sourceImage ImageBatch /* not a class type */, convolutionGradientState CNNConvolutionGradientStateBatch /* not a class type */) ImageBatch /* not a class type */
	Encode()
	EncodeToCommandBufferSourceImageConvolutionGradientState(commandBuffer unsafe.Pointer, sourceImage IImage, convolutionGradientState ICNNConvolutionGradientState) IImage
	EncodeBatchToCommandBufferSourceImagesConvolutionGradientStatesDestinationImages(commandBuffer unsafe.Pointer, sourceImage ImageBatch /* not a class type */, convolutionGradientState CNNConvolutionGradientStateBatch /* not a class type */, destinationImage ImageBatch /* not a class type */)
	EncodeToCommandBufferSourceImageConvolutionGradientStateDestinationImage(commandBuffer unsafe.Pointer, sourceImage IImage, convolutionGradientState ICNNConvolutionGradientState, destinationImage IImage)
	EncodeBatchToCommandBufferSourceImagesConvolutionGradientStatesDestinationStatesDestinationStateIsTemporary(commandBuffer unsafe.Pointer, sourceImages ImageBatch /* not a class type */, convolutionGradientStates CNNConvolutionGradientStateBatch /* not a class type */, outStates CNNConvolutionTransposeGradientStateBatch /* not a class type */, isTemporary bool) ImageBatch /* not a class type */
	EncodeToCommandBufferSourceImageConvolutionGradientStateDestinationStateDestinationStateIsTemporary(commandBuffer unsafe.Pointer, sourceImage IImage, convolutionGradientState ICNNConvolutionGradientState, outState objectivec.IObject, isTemporary bool) IImage
	ExportWeightsAndBiases()
	ExportWeightsAndBiasesWithCommandBufferResultStateCanBeTemporary(commandBuffer unsafe.Pointer, resultStateCanBeTemporary bool) ICNNConvolutionWeightsAndBiasesState
	ReloadWeightsAndBiasesFromDataSource()
	ReloadWeightsAndBiases()
	ReloadWeightsAndBiasesWithCommandBufferState(commandBuffer unsafe.Pointer, state ICNNConvolutionWeightsAndBiasesState)
	ResultStateBatch()
	ResultStateBatchForSourceImageSourceStatesDestinationImage(sourceImage ImageBatch /* not a class type */, sourceStates CNNConvolutionGradientStateBatch /* not a class type */, destinationImage ImageBatch /* not a class type */) CNNConvolutionTransposeGradientStateBatch /* not a class type */
	ResultState()
	ResultStateForSourceImageSourceStatesDestinationImage(sourceImage IImage, sourceStates unsafe.Pointer, destinationImage IImage) ICNNConvolutionTransposeGradientState
	TemporaryResultStateBatch()
	TemporaryResultStateBatchForCommandBufferSourceImageSourceStatesDestinationImage(commandBuffer unsafe.Pointer, sourceImage ImageBatch /* not a class type */, sourceStates CNNConvolutionGradientStateBatch /* not a class type */, destinationImage ImageBatch /* not a class type */) CNNConvolutionTransposeGradientStateBatch /* not a class type */
	TemporaryResultState()
	TemporaryResultStateForCommandBufferSourceImageSourceStatesDestinationImage(commandBuffer unsafe.Pointer, sourceImage IImage, sourceStates unsafe.Pointer, destinationImage IImage) ICNNConvolutionTransposeGradientState
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for CNNConvolutionTranspose */
// Alloc allocates a new instance without initialization.
func (cc _CNNConvolutionTransposeClass) Alloc() CNNConvolutionTranspose {
	rv := objc.Send[CNNConvolutionTranspose](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _CNNConvolutionTransposeClass) New() CNNConvolutionTranspose {
	rv := objc.Send[CNNConvolutionTranspose](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CNNConvolutionTranspose) Init() CNNConvolutionTranspose {
	rv := objc.Send[CNNConvolutionTranspose](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CNNConvolutionTranspose) Autorelease() CNNConvolutionTranspose {
	rv := objc.Send[CNNConvolutionTranspose](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCNNConvolutionTranspose creates a new CNNConvolutionTranspose instance.
func NewCNNConvolutionTranspose() CNNConvolutionTranspose {
	return getCNNConvolutionTransposeClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for CNNConvolutionTranspose */
// A transposed convolution kernel.


// A transposed convolution kernel.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNConvolutionTranspose
type CNNConvolutionTranspose struct {
	CNNKernel
}

// CNNConvolutionTransposeFrom constructs a [CNNConvolutionTranspose] from an unsafe.Pointer.
//
// A transposed convolution kernel.
func CNNConvolutionTransposeFrom(ptr unsafe.Pointer) CNNConvolutionTranspose {
	return CNNConvolutionTranspose{
		CNNKernel: CNNKernelFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for CNNConvolutionTranspose */

// Initializes a transposed convolution kernel.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnconvolutiontranspose/2866995-initwithcoder
func NewCNNConvolutionTransposeWithCoderDevice(aDecoder Coder /* not a class type */, device unsafe.Pointer) CNNConvolutionTranspose {
	instance := getCNNConvolutionTransposeClass().Alloc()
	rv := objc.Send[CNNConvolutionTranspose](instance.ID, objc.Sel("initWithCoder:device:"), aDecoder, device)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewCNNConvolutionTransposeWithCoderDevice */


// Initializes a transposed convolution kernel.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnconvolutiontranspose/2867157-initwithdevice
func NewCNNConvolutionTransposeWithDeviceWeights(device unsafe.Pointer, weights unsafe.Pointer) CNNConvolutionTranspose {
	instance := getCNNConvolutionTransposeClass().Alloc()
	rv := objc.Send[CNNConvolutionTranspose](instance.ID, objc.Sel("initWithDevice:weights:"), device, weights)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewCNNConvolutionTransposeWithDeviceWeights */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for CNNConvolutionTranspose */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for CNNConvolutionTranspose */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for CNNConvolutionTranspose */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnconvolutiontranspose/2942406-encodebatch
func (c_ CNNConvolutionTranspose) EncodeBatch() {
	objc.Send[objc.ID](c_.ID, objc.Sel("encodeBatch"))
}/* debug [instance_methods/method]: EncodeBatch */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnconvolutiontranspose/2942406-encodebatchtocommandbuffer
func (c_ CNNConvolutionTranspose) EncodeBatchToCommandBufferSourceImagesConvolutionGradientStates(commandBuffer unsafe.Pointer, sourceImage ImageBatch /* not a class type */, convolutionGradientState CNNConvolutionGradientStateBatch /* not a class type */) ImageBatch /* not a class type */ {
	rv := objc.Send[ImageBatch](c_.ID, objc.Sel("encodeBatchToCommandBuffer:sourceImages:convolutionGradientStates:"), commandBuffer, sourceImage, convolutionGradientState)
	return rv
}/* debug [instance_methods/method]: EncodeBatchToCommandBufferSourceImagesConvolutionGradientStates */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnconvolutiontranspose/2942409-encode
func (c_ CNNConvolutionTranspose) Encode() {
	objc.Send[objc.ID](c_.ID, objc.Sel("encode"))
}/* debug [instance_methods/method]: Encode */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnconvolutiontranspose/2942409-encodetocommandbuffer
func (c_ CNNConvolutionTranspose) EncodeToCommandBufferSourceImageConvolutionGradientState(commandBuffer unsafe.Pointer, sourceImage IImage, convolutionGradientState ICNNConvolutionGradientState) IImage {
	rv := objc.Send[Image](c_.ID, objc.Sel("encodeToCommandBuffer:sourceImage:convolutionGradientState:"), commandBuffer, sourceImage, convolutionGradientState)
	return rv
}/* debug [instance_methods/method]: EncodeToCommandBufferSourceImageConvolutionGradientState */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnconvolutiontranspose/2942411-encodebatchtocommandbuffer
func (c_ CNNConvolutionTranspose) EncodeBatchToCommandBufferSourceImagesConvolutionGradientStatesDestinationImages(commandBuffer unsafe.Pointer, sourceImage ImageBatch /* not a class type */, convolutionGradientState CNNConvolutionGradientStateBatch /* not a class type */, destinationImage ImageBatch /* not a class type */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("encodeBatchToCommandBuffer:sourceImages:convolutionGradientStates:destinationImages:"), commandBuffer, sourceImage, convolutionGradientState, destinationImage)
}/* debug [instance_methods/method]: EncodeBatchToCommandBufferSourceImagesConvolutionGradientStatesDestinationImages */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnconvolutiontranspose/2942429-encodetocommandbuffer
func (c_ CNNConvolutionTranspose) EncodeToCommandBufferSourceImageConvolutionGradientStateDestinationImage(commandBuffer unsafe.Pointer, sourceImage IImage, convolutionGradientState ICNNConvolutionGradientState, destinationImage IImage) {
	objc.Send[objc.ID](c_.ID, objc.Sel("encodeToCommandBuffer:sourceImage:convolutionGradientState:destinationImage:"), commandBuffer, sourceImage, convolutionGradientState, destinationImage)
}/* debug [instance_methods/method]: EncodeToCommandBufferSourceImageConvolutionGradientStateDestinationImage */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnconvolutiontranspose/3131770-encodebatchtocommandbuffer
func (c_ CNNConvolutionTranspose) EncodeBatchToCommandBufferSourceImagesConvolutionGradientStatesDestinationStatesDestinationStateIsTemporary(commandBuffer unsafe.Pointer, sourceImages ImageBatch /* not a class type */, convolutionGradientStates CNNConvolutionGradientStateBatch /* not a class type */, outStates CNNConvolutionTransposeGradientStateBatch /* not a class type */, isTemporary bool) ImageBatch /* not a class type */ {
	rv := objc.Send[ImageBatch](c_.ID, objc.Sel("encodeBatchToCommandBuffer:sourceImages:convolutionGradientStates:destinationStates:destinationStateIsTemporary:"), commandBuffer, sourceImages, convolutionGradientStates, outStates, isTemporary)
	return rv
}/* debug [instance_methods/method]: EncodeBatchToCommandBufferSourceImagesConvolutionGradientStatesDestinationStatesDestinationStateIsTemporary */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnconvolutiontranspose/3131771-encodetocommandbuffer
func (c_ CNNConvolutionTranspose) EncodeToCommandBufferSourceImageConvolutionGradientStateDestinationStateDestinationStateIsTemporary(commandBuffer unsafe.Pointer, sourceImage IImage, convolutionGradientState ICNNConvolutionGradientState, outState objectivec.IObject, isTemporary bool) IImage {
	rv := objc.Send[Image](c_.ID, objc.Sel("encodeToCommandBuffer:sourceImage:convolutionGradientState:destinationState:destinationStateIsTemporary:"), commandBuffer, sourceImage, convolutionGradientState, outState, isTemporary)
	return rv
}/* debug [instance_methods/method]: EncodeToCommandBufferSourceImageConvolutionGradientStateDestinationStateDestinationStateIsTemporary */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnconvolutiontranspose/3131772-exportweightsandbiases
func (c_ CNNConvolutionTranspose) ExportWeightsAndBiases() {
	objc.Send[objc.ID](c_.ID, objc.Sel("exportWeightsAndBiases"))
}/* debug [instance_methods/method]: ExportWeightsAndBiases */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnconvolutiontranspose/3131772-exportweightsandbiaseswithcomman
func (c_ CNNConvolutionTranspose) ExportWeightsAndBiasesWithCommandBufferResultStateCanBeTemporary(commandBuffer unsafe.Pointer, resultStateCanBeTemporary bool) ICNNConvolutionWeightsAndBiasesState {
	rv := objc.Send[CNNConvolutionWeightsAndBiasesState](c_.ID, objc.Sel("exportWeightsAndBiasesWithCommandBuffer:resultStateCanBeTemporary:"), commandBuffer, resultStateCanBeTemporary)
	return rv
}/* debug [instance_methods/method]: ExportWeightsAndBiasesWithCommandBufferResultStateCanBeTemporary */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnconvolutiontranspose/3131773-reloadweightsandbiasesfromdataso
func (c_ CNNConvolutionTranspose) ReloadWeightsAndBiasesFromDataSource() {
	objc.Send[objc.ID](c_.ID, objc.Sel("reloadWeightsAndBiasesFromDataSource"))
}/* debug [instance_methods/method]: ReloadWeightsAndBiasesFromDataSource */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnconvolutiontranspose/3131774-reloadweightsandbiases
func (c_ CNNConvolutionTranspose) ReloadWeightsAndBiases() {
	objc.Send[objc.ID](c_.ID, objc.Sel("reloadWeightsAndBiases"))
}/* debug [instance_methods/method]: ReloadWeightsAndBiases */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnconvolutiontranspose/3131774-reloadweightsandbiaseswithcomman
func (c_ CNNConvolutionTranspose) ReloadWeightsAndBiasesWithCommandBufferState(commandBuffer unsafe.Pointer, state ICNNConvolutionWeightsAndBiasesState) {
	objc.Send[objc.ID](c_.ID, objc.Sel("reloadWeightsAndBiasesWithCommandBuffer:state:"), commandBuffer, state)
}/* debug [instance_methods/method]: ReloadWeightsAndBiasesWithCommandBufferState */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnconvolutiontranspose/3131775-resultstatebatch
func (c_ CNNConvolutionTranspose) ResultStateBatch() {
	objc.Send[objc.ID](c_.ID, objc.Sel("resultStateBatch"))
}/* debug [instance_methods/method]: ResultStateBatch */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnconvolutiontranspose/3131775-resultstatebatchforsourceimage
func (c_ CNNConvolutionTranspose) ResultStateBatchForSourceImageSourceStatesDestinationImage(sourceImage ImageBatch /* not a class type */, sourceStates CNNConvolutionGradientStateBatch /* not a class type */, destinationImage ImageBatch /* not a class type */) CNNConvolutionTransposeGradientStateBatch /* not a class type */ {
	rv := objc.Send[CNNConvolutionTransposeGradientStateBatch](c_.ID, objc.Sel("resultStateBatchForSourceImage:sourceStates:destinationImage:"), sourceImage, sourceStates, destinationImage)
	return rv
}/* debug [instance_methods/method]: ResultStateBatchForSourceImageSourceStatesDestinationImage */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnconvolutiontranspose/3131776-resultstate
func (c_ CNNConvolutionTranspose) ResultState() {
	objc.Send[objc.ID](c_.ID, objc.Sel("resultState"))
}/* debug [instance_methods/method]: ResultState */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnconvolutiontranspose/3131776-resultstateforsourceimage
func (c_ CNNConvolutionTranspose) ResultStateForSourceImageSourceStatesDestinationImage(sourceImage IImage, sourceStates unsafe.Pointer, destinationImage IImage) ICNNConvolutionTransposeGradientState {
	rv := objc.Send[CNNConvolutionTransposeGradientState](c_.ID, objc.Sel("resultStateForSourceImage:sourceStates:destinationImage:"), sourceImage, sourceStates, destinationImage)
	return rv
}/* debug [instance_methods/method]: ResultStateForSourceImageSourceStatesDestinationImage */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnconvolutiontranspose/3131777-temporaryresultstatebatch
func (c_ CNNConvolutionTranspose) TemporaryResultStateBatch() {
	objc.Send[objc.ID](c_.ID, objc.Sel("temporaryResultStateBatch"))
}/* debug [instance_methods/method]: TemporaryResultStateBatch */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnconvolutiontranspose/3131777-temporaryresultstatebatchforcomm
func (c_ CNNConvolutionTranspose) TemporaryResultStateBatchForCommandBufferSourceImageSourceStatesDestinationImage(commandBuffer unsafe.Pointer, sourceImage ImageBatch /* not a class type */, sourceStates CNNConvolutionGradientStateBatch /* not a class type */, destinationImage ImageBatch /* not a class type */) CNNConvolutionTransposeGradientStateBatch /* not a class type */ {
	rv := objc.Send[CNNConvolutionTransposeGradientStateBatch](c_.ID, objc.Sel("temporaryResultStateBatchForCommandBuffer:sourceImage:sourceStates:destinationImage:"), commandBuffer, sourceImage, sourceStates, destinationImage)
	return rv
}/* debug [instance_methods/method]: TemporaryResultStateBatchForCommandBufferSourceImageSourceStatesDestinationImage */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnconvolutiontranspose/3131778-temporaryresultstate
func (c_ CNNConvolutionTranspose) TemporaryResultState() {
	objc.Send[objc.ID](c_.ID, objc.Sel("temporaryResultState"))
}/* debug [instance_methods/method]: TemporaryResultState */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnconvolutiontranspose/3131778-temporaryresultstateforcommandbu
func (c_ CNNConvolutionTranspose) TemporaryResultStateForCommandBufferSourceImageSourceStatesDestinationImage(commandBuffer unsafe.Pointer, sourceImage IImage, sourceStates unsafe.Pointer, destinationImage IImage) ICNNConvolutionTransposeGradientState {
	rv := objc.Send[CNNConvolutionTransposeGradientState](c_.ID, objc.Sel("temporaryResultStateForCommandBuffer:sourceImage:sourceStates:destinationImage:"), commandBuffer, sourceImage, sourceStates, destinationImage)
	return rv
}/* debug [instance_methods/method]: TemporaryResultStateForCommandBufferSourceImageSourceStatesDestinationImage */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for CNNConvolutionTranspose */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnconvolutiontranspose/2867016-outputfeaturechannels
func (c_ CNNConvolutionTranspose) OutputFeatureChannels() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](c_.ID, objc.Sel("outputFeatureChannels"))
	return rv
}/* debug [instance_properties/getter]: outputFeatureChannels */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnconvolutiontranspose/2867016-outputfeaturechannels
func (c_ CNNConvolutionTranspose) SetOutputFeatureChannels(value objectivec.IObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setOutputFeatureChannels:"), value)
}/* debug [instance_properties/setter]: outputFeatureChannels */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnconvolutiontranspose/2867086-kerneloffsety
func (c_ CNNConvolutionTranspose) KernelOffsetY() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](c_.ID, objc.Sel("kernelOffsetY"))
	return rv
}/* debug [instance_properties/getter]: kernelOffsetY */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnconvolutiontranspose/2867086-kerneloffsety
func (c_ CNNConvolutionTranspose) SetKernelOffsetY(value objectivec.IObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setKernelOffsetY:"), value)
}/* debug [instance_properties/setter]: kernelOffsetY */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnconvolutiontranspose/2867099-groups
func (c_ CNNConvolutionTranspose) Groups() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](c_.ID, objc.Sel("groups"))
	return rv
}/* debug [instance_properties/getter]: groups */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnconvolutiontranspose/2867099-groups
func (c_ CNNConvolutionTranspose) SetGroups(value objectivec.IObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setGroups:"), value)
}/* debug [instance_properties/setter]: groups */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnconvolutiontranspose/2867174-inputfeaturechannels
func (c_ CNNConvolutionTranspose) InputFeatureChannels() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](c_.ID, objc.Sel("inputFeatureChannels"))
	return rv
}/* debug [instance_properties/getter]: inputFeatureChannels */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnconvolutiontranspose/2867174-inputfeaturechannels
func (c_ CNNConvolutionTranspose) SetInputFeatureChannels(value objectivec.IObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setInputFeatureChannels:"), value)
}/* debug [instance_properties/setter]: inputFeatureChannels */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnconvolutiontranspose/2867176-kerneloffsetx
func (c_ CNNConvolutionTranspose) KernelOffsetX() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](c_.ID, objc.Sel("kernelOffsetX"))
	return rv
}/* debug [instance_properties/getter]: kernelOffsetX */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnconvolutiontranspose/2867176-kerneloffsetx
func (c_ CNNConvolutionTranspose) SetKernelOffsetX(value objectivec.IObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setKernelOffsetX:"), value)
}/* debug [instance_properties/setter]: kernelOffsetX */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnconvolutiontranspose/2951924-accumulatorprecisionoption
func (c_ CNNConvolutionTranspose) AccumulatorPrecisionOption() ConvolutionAccumulatorPrecisionOption get set /* not a class type */ {
	rv := objc.Send[objc.ID](c_.ID, objc.Sel("accumulatorPrecisionOption"))
	return rv
}/* debug [instance_properties/getter]: accumulatorPrecisionOption */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnconvolutiontranspose/2951924-accumulatorprecisionoption
func (c_ CNNConvolutionTranspose) SetAccumulatorPrecisionOption(value ConvolutionAccumulatorPrecisionOption get set /* not a class type */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setAccumulatorPrecisionOption:"), value)
}/* debug [instance_properties/setter]: accumulatorPrecisionOption */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnconvolutiontranspose/3131769-datasource
func (c_ CNNConvolutionTranspose) DataSource() CNNConvolutionDataSource get /* not a class type */ {
	rv := objc.Send[objc.ID](c_.ID, objc.Sel("dataSource"))
	return rv
}/* debug [instance_properties/getter]: dataSource */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnconvolutiontranspose/3131769-datasource
func (c_ CNNConvolutionTranspose) SetDataSource(value CNNConvolutionDataSource get /* not a class type */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setDataSource:"), value)
}/* debug [instance_properties/setter]: dataSource */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MPSCNNConvolutionTranspose */


