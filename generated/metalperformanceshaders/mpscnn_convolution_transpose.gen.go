// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)





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





// An interface definition for the [CNNConvolutionTranspose] class.
type ICNNConvolutionTranspose interface {
	ICNNKernel
	

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


}





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






// Initializes a transposed convolution kernel.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnconvolutiontranspose/2866995-initwithcoder
func NewCNNConvolutionTransposeWithCoderDevice(aDecoder foundation.Coder, device unsafe.Pointer) CNNConvolutionTranspose {
	instance := getCNNConvolutionTransposeClass().Alloc()
	rv := objc.Send[CNNConvolutionTranspose](instance.ID, objc.Sel("initWithCoder:device:"), aDecoder, device)
	rv.Autorelease()
	return rv
}


// Initializes a transposed convolution kernel.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnconvolutiontranspose/2867157-initwithdevice
func NewCNNConvolutionTransposeWithDeviceWeights(device unsafe.Pointer, weights unsafe.Pointer) CNNConvolutionTranspose {
	instance := getCNNConvolutionTransposeClass().Alloc()
	rv := objc.Send[CNNConvolutionTranspose](instance.ID, objc.Sel("initWithDevice:weights:"), device, weights)
	rv.Autorelease()
	return rv
}

















// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnconvolutiontranspose/2942406-encodebatch
func (c_ CNNConvolutionTranspose) EncodeBatch() {
	objc.Send[objc.ID](c_.ID, objc.Sel("encodeBatch"))
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnconvolutiontranspose/2942406-encodebatchtocommandbuffer
func (c_ CNNConvolutionTranspose) EncodeBatchToCommandBufferSourceImagesConvolutionGradientStates(commandBuffer unsafe.Pointer, sourceImage ImageBatch /* not a class type */, convolutionGradientState CNNConvolutionGradientStateBatch /* not a class type */) ImageBatch /* not a class type */ {
	rv := objc.Send[ImageBatch](c_.ID, objc.Sel("encodeBatchToCommandBuffer:sourceImages:convolutionGradientStates:"), commandBuffer, sourceImage, convolutionGradientState)
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnconvolutiontranspose/2942409-encode
func (c_ CNNConvolutionTranspose) Encode() {
	objc.Send[objc.ID](c_.ID, objc.Sel("encode"))
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnconvolutiontranspose/2942409-encodetocommandbuffer
func (c_ CNNConvolutionTranspose) EncodeToCommandBufferSourceImageConvolutionGradientState(commandBuffer unsafe.Pointer, sourceImage IImage, convolutionGradientState ICNNConvolutionGradientState) IImage {
	rv := objc.Send[Image](c_.ID, objc.Sel("encodeToCommandBuffer:sourceImage:convolutionGradientState:"), commandBuffer, sourceImage, convolutionGradientState)
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnconvolutiontranspose/2942411-encodebatchtocommandbuffer
func (c_ CNNConvolutionTranspose) EncodeBatchToCommandBufferSourceImagesConvolutionGradientStatesDestinationImages(commandBuffer unsafe.Pointer, sourceImage ImageBatch /* not a class type */, convolutionGradientState CNNConvolutionGradientStateBatch /* not a class type */, destinationImage ImageBatch /* not a class type */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("encodeBatchToCommandBuffer:sourceImages:convolutionGradientStates:destinationImages:"), commandBuffer, sourceImage, convolutionGradientState, destinationImage)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnconvolutiontranspose/2942429-encodetocommandbuffer
func (c_ CNNConvolutionTranspose) EncodeToCommandBufferSourceImageConvolutionGradientStateDestinationImage(commandBuffer unsafe.Pointer, sourceImage IImage, convolutionGradientState ICNNConvolutionGradientState, destinationImage IImage) {
	objc.Send[objc.ID](c_.ID, objc.Sel("encodeToCommandBuffer:sourceImage:convolutionGradientState:destinationImage:"), commandBuffer, sourceImage, convolutionGradientState, destinationImage)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnconvolutiontranspose/3131770-encodebatchtocommandbuffer
func (c_ CNNConvolutionTranspose) EncodeBatchToCommandBufferSourceImagesConvolutionGradientStatesDestinationStatesDestinationStateIsTemporary(commandBuffer unsafe.Pointer, sourceImages ImageBatch /* not a class type */, convolutionGradientStates CNNConvolutionGradientStateBatch /* not a class type */, outStates CNNConvolutionTransposeGradientStateBatch /* not a class type */, isTemporary bool) ImageBatch /* not a class type */ {
	rv := objc.Send[ImageBatch](c_.ID, objc.Sel("encodeBatchToCommandBuffer:sourceImages:convolutionGradientStates:destinationStates:destinationStateIsTemporary:"), commandBuffer, sourceImages, convolutionGradientStates, outStates, isTemporary)
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnconvolutiontranspose/3131771-encodetocommandbuffer
func (c_ CNNConvolutionTranspose) EncodeToCommandBufferSourceImageConvolutionGradientStateDestinationStateDestinationStateIsTemporary(commandBuffer unsafe.Pointer, sourceImage IImage, convolutionGradientState ICNNConvolutionGradientState, outState objectivec.IObject, isTemporary bool) IImage {
	rv := objc.Send[Image](c_.ID, objc.Sel("encodeToCommandBuffer:sourceImage:convolutionGradientState:destinationState:destinationStateIsTemporary:"), commandBuffer, sourceImage, convolutionGradientState, outState, isTemporary)
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnconvolutiontranspose/3131772-exportweightsandbiases
func (c_ CNNConvolutionTranspose) ExportWeightsAndBiases() {
	objc.Send[objc.ID](c_.ID, objc.Sel("exportWeightsAndBiases"))
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnconvolutiontranspose/3131772-exportweightsandbiaseswithcomman
func (c_ CNNConvolutionTranspose) ExportWeightsAndBiasesWithCommandBufferResultStateCanBeTemporary(commandBuffer unsafe.Pointer, resultStateCanBeTemporary bool) ICNNConvolutionWeightsAndBiasesState {
	rv := objc.Send[CNNConvolutionWeightsAndBiasesState](c_.ID, objc.Sel("exportWeightsAndBiasesWithCommandBuffer:resultStateCanBeTemporary:"), commandBuffer, resultStateCanBeTemporary)
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnconvolutiontranspose/3131773-reloadweightsandbiasesfromdataso
func (c_ CNNConvolutionTranspose) ReloadWeightsAndBiasesFromDataSource() {
	objc.Send[objc.ID](c_.ID, objc.Sel("reloadWeightsAndBiasesFromDataSource"))
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnconvolutiontranspose/3131774-reloadweightsandbiases
func (c_ CNNConvolutionTranspose) ReloadWeightsAndBiases() {
	objc.Send[objc.ID](c_.ID, objc.Sel("reloadWeightsAndBiases"))
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnconvolutiontranspose/3131774-reloadweightsandbiaseswithcomman
func (c_ CNNConvolutionTranspose) ReloadWeightsAndBiasesWithCommandBufferState(commandBuffer unsafe.Pointer, state ICNNConvolutionWeightsAndBiasesState) {
	objc.Send[objc.ID](c_.ID, objc.Sel("reloadWeightsAndBiasesWithCommandBuffer:state:"), commandBuffer, state)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnconvolutiontranspose/3131775-resultstatebatch
func (c_ CNNConvolutionTranspose) ResultStateBatch() {
	objc.Send[objc.ID](c_.ID, objc.Sel("resultStateBatch"))
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnconvolutiontranspose/3131775-resultstatebatchforsourceimage
func (c_ CNNConvolutionTranspose) ResultStateBatchForSourceImageSourceStatesDestinationImage(sourceImage ImageBatch /* not a class type */, sourceStates CNNConvolutionGradientStateBatch /* not a class type */, destinationImage ImageBatch /* not a class type */) CNNConvolutionTransposeGradientStateBatch /* not a class type */ {
	rv := objc.Send[CNNConvolutionTransposeGradientStateBatch](c_.ID, objc.Sel("resultStateBatchForSourceImage:sourceStates:destinationImage:"), sourceImage, sourceStates, destinationImage)
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnconvolutiontranspose/3131776-resultstate
func (c_ CNNConvolutionTranspose) ResultState() {
	objc.Send[objc.ID](c_.ID, objc.Sel("resultState"))
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnconvolutiontranspose/3131776-resultstateforsourceimage
func (c_ CNNConvolutionTranspose) ResultStateForSourceImageSourceStatesDestinationImage(sourceImage IImage, sourceStates unsafe.Pointer, destinationImage IImage) ICNNConvolutionTransposeGradientState {
	rv := objc.Send[CNNConvolutionTransposeGradientState](c_.ID, objc.Sel("resultStateForSourceImage:sourceStates:destinationImage:"), sourceImage, sourceStates, destinationImage)
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnconvolutiontranspose/3131777-temporaryresultstatebatch
func (c_ CNNConvolutionTranspose) TemporaryResultStateBatch() {
	objc.Send[objc.ID](c_.ID, objc.Sel("temporaryResultStateBatch"))
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnconvolutiontranspose/3131777-temporaryresultstatebatchforcomm
func (c_ CNNConvolutionTranspose) TemporaryResultStateBatchForCommandBufferSourceImageSourceStatesDestinationImage(commandBuffer unsafe.Pointer, sourceImage ImageBatch /* not a class type */, sourceStates CNNConvolutionGradientStateBatch /* not a class type */, destinationImage ImageBatch /* not a class type */) CNNConvolutionTransposeGradientStateBatch /* not a class type */ {
	rv := objc.Send[CNNConvolutionTransposeGradientStateBatch](c_.ID, objc.Sel("temporaryResultStateBatchForCommandBuffer:sourceImage:sourceStates:destinationImage:"), commandBuffer, sourceImage, sourceStates, destinationImage)
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnconvolutiontranspose/3131778-temporaryresultstate
func (c_ CNNConvolutionTranspose) TemporaryResultState() {
	objc.Send[objc.ID](c_.ID, objc.Sel("temporaryResultState"))
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnconvolutiontranspose/3131778-temporaryresultstateforcommandbu
func (c_ CNNConvolutionTranspose) TemporaryResultStateForCommandBufferSourceImageSourceStatesDestinationImage(commandBuffer unsafe.Pointer, sourceImage IImage, sourceStates unsafe.Pointer, destinationImage IImage) ICNNConvolutionTransposeGradientState {
	rv := objc.Send[CNNConvolutionTransposeGradientState](c_.ID, objc.Sel("temporaryResultStateForCommandBuffer:sourceImage:sourceStates:destinationImage:"), commandBuffer, sourceImage, sourceStates, destinationImage)
	return rv
}







// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnconvolutiontranspose/2867016-outputfeaturechannels
func (c_ CNNConvolutionTranspose) OutputFeatureChannels() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](c_.ID, objc.Sel("outputFeatureChannels"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnconvolutiontranspose/2867016-outputfeaturechannels
func (c_ CNNConvolutionTranspose) SetOutputFeatureChannels(value objectivec.IObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setOutputFeatureChannels:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnconvolutiontranspose/2867086-kerneloffsety
func (c_ CNNConvolutionTranspose) KernelOffsetY() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](c_.ID, objc.Sel("kernelOffsetY"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnconvolutiontranspose/2867086-kerneloffsety
func (c_ CNNConvolutionTranspose) SetKernelOffsetY(value objectivec.IObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setKernelOffsetY:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnconvolutiontranspose/2867099-groups
func (c_ CNNConvolutionTranspose) Groups() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](c_.ID, objc.Sel("groups"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnconvolutiontranspose/2867099-groups
func (c_ CNNConvolutionTranspose) SetGroups(value objectivec.IObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setGroups:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnconvolutiontranspose/2867174-inputfeaturechannels
func (c_ CNNConvolutionTranspose) InputFeatureChannels() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](c_.ID, objc.Sel("inputFeatureChannels"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnconvolutiontranspose/2867174-inputfeaturechannels
func (c_ CNNConvolutionTranspose) SetInputFeatureChannels(value objectivec.IObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setInputFeatureChannels:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnconvolutiontranspose/2867176-kerneloffsetx
func (c_ CNNConvolutionTranspose) KernelOffsetX() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](c_.ID, objc.Sel("kernelOffsetX"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnconvolutiontranspose/2867176-kerneloffsetx
func (c_ CNNConvolutionTranspose) SetKernelOffsetX(value objectivec.IObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setKernelOffsetX:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnconvolutiontranspose/2951924-accumulatorprecisionoption
func (c_ CNNConvolutionTranspose) AccumulatorPrecisionOption() ConvolutionAccumulatorPrecisionOption get set /* not a class type */ {
	rv := objc.Send[objc.ID](c_.ID, objc.Sel("accumulatorPrecisionOption"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnconvolutiontranspose/2951924-accumulatorprecisionoption
func (c_ CNNConvolutionTranspose) SetAccumulatorPrecisionOption(value ConvolutionAccumulatorPrecisionOption get set /* not a class type */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setAccumulatorPrecisionOption:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnconvolutiontranspose/3131769-datasource
func (c_ CNNConvolutionTranspose) DataSource() CNNConvolutionDataSource get /* not a class type */ {
	rv := objc.Send[objc.ID](c_.ID, objc.Sel("dataSource"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnconvolutiontranspose/3131769-datasource
func (c_ CNNConvolutionTranspose) SetDataSource(value CNNConvolutionDataSource get /* not a class type */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setDataSource:"), value)
}







