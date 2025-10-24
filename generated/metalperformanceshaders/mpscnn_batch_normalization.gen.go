// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)





// The class instance for the [CNNBatchNormalization] class.
var (
	CNNBatchNormalizationClass     _CNNBatchNormalizationClass
	CNNBatchNormalizationClassOnce sync.Once
)

func getCNNBatchNormalizationClass() _CNNBatchNormalizationClass {
	CNNBatchNormalizationClassOnce.Do(func() {
		CNNBatchNormalizationClass = _CNNBatchNormalizationClass{objc.GetClass("MPSCNNBatchNormalization")}
	})
	return CNNBatchNormalizationClass
}

type _CNNBatchNormalizationClass struct {
	class objc.Class
}





// An interface definition for the [CNNBatchNormalization] class.
type ICNNBatchNormalization interface {
	ICNNKernel
	

	// properties:
	Epsilon() objectivec.IObject
	SetEpsilon(value objectivec.IObject)
	NumberOfFeatureChannels() objectivec.IObject
	SetNumberOfFeatureChannels(value objectivec.IObject)
	DataSource() CNNBatchNormalizationDataSource get /* not a class type */
	SetDataSource(value CNNBatchNormalizationDataSource get /* not a class type */)


	

	// methods:
	Encode()
	EncodeToCommandBufferSourceImageBatchNormalizationStateDestinationImage(commandBuffer unsafe.Pointer, sourceImage IImage, batchNormalizationState ICNNBatchNormalizationState, destinationImage IImage)
	EncodeBatch()
	EncodeBatchToCommandBufferSourceImagesBatchNormalizationStateDestinationImages(commandBuffer unsafe.Pointer, sourceImages ImageBatch /* not a class type */, batchNormalizationState ICNNBatchNormalizationState, destinationImages ImageBatch /* not a class type */)
	ReloadGammaAndBeta()
	ReloadGammaAndBetaWithCommandBufferGammaAndBetaState(commandBuffer unsafe.Pointer, gammaAndBetaState ICNNNormalizationGammaAndBetaState)
	ResultState()
	ResultStateForSourceImageSourceStatesDestinationImage(sourceImage IImage, sourceStates unsafe.Pointer, destinationImage IImage) ICNNBatchNormalizationState
	TemporaryResultState()
	TemporaryResultStateForCommandBufferSourceImageSourceStatesDestinationImage(commandBuffer unsafe.Pointer, sourceImage IImage, sourceStates unsafe.Pointer, destinationImage IImage) ICNNBatchNormalizationState
	ReloadGammaAndBetaFromDataSource()
	ReloadMeanAndVarianceFromDataSource()
	ReloadMeanAndVariance()
	ReloadMeanAndVarianceWithCommandBufferMeanAndVarianceState(commandBuffer unsafe.Pointer, meanAndVarianceState ICNNNormalizationMeanAndVarianceState)


}





// Alloc allocates a new instance without initialization.
func (cc _CNNBatchNormalizationClass) Alloc() CNNBatchNormalization {
	rv := objc.Send[CNNBatchNormalization](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _CNNBatchNormalizationClass) New() CNNBatchNormalization {
	rv := objc.Send[CNNBatchNormalization](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CNNBatchNormalization) Init() CNNBatchNormalization {
	rv := objc.Send[CNNBatchNormalization](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CNNBatchNormalization) Autorelease() CNNBatchNormalization {
	rv := objc.Send[CNNBatchNormalization](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCNNBatchNormalization creates a new CNNBatchNormalization instance.
func NewCNNBatchNormalization() CNNBatchNormalization {
	return getCNNBatchNormalizationClass().New()
}





// A batch normalization kernel.


// A batch normalization kernel.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNBatchNormalization
type CNNBatchNormalization struct {
	CNNKernel
}

// CNNBatchNormalizationFrom constructs a [CNNBatchNormalization] from an unsafe.Pointer.
//
// A batch normalization kernel.
func CNNBatchNormalizationFrom(ptr unsafe.Pointer) CNNBatchNormalization {
	return CNNBatchNormalization{
		CNNKernel: CNNKernelFrom(ptr),
	}
}






// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnbatchnormalization/2942602-initwithcoder
func NewCNNBatchNormalizationWithCoderDevice(aDecoder foundation.Coder, device unsafe.Pointer) CNNBatchNormalization {
	instance := getCNNBatchNormalizationClass().Alloc()
	rv := objc.Send[CNNBatchNormalization](instance.ID, objc.Sel("initWithCoder:device:"), aDecoder, device)
	rv.Autorelease()
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnbatchnormalization/2942600-initwithdevice
func NewCNNBatchNormalizationWithDeviceDataSource(device unsafe.Pointer, dataSource unsafe.Pointer) CNNBatchNormalization {
	instance := getCNNBatchNormalizationClass().Alloc()
	rv := objc.Send[CNNBatchNormalization](instance.ID, objc.Sel("initWithDevice:dataSource:"), device, dataSource)
	rv.Autorelease()
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnbatchnormalization/3013771-initwithdevice
func NewCNNBatchNormalizationWithDeviceDataSourceFusedNeuronDescriptor(device unsafe.Pointer, dataSource unsafe.Pointer, fusedNeuronDescriptor INeuronDescriptor) CNNBatchNormalization {
	instance := getCNNBatchNormalizationClass().Alloc()
	rv := objc.Send[CNNBatchNormalization](instance.ID, objc.Sel("initWithDevice:dataSource:fusedNeuronDescriptor:"), device, dataSource, fusedNeuronDescriptor)
	rv.Autorelease()
	return rv
}

















// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnbatchnormalization/2942591-encode
func (c_ CNNBatchNormalization) Encode() {
	objc.Send[objc.ID](c_.ID, objc.Sel("encode"))
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnbatchnormalization/2942591-encodetocommandbuffer
func (c_ CNNBatchNormalization) EncodeToCommandBufferSourceImageBatchNormalizationStateDestinationImage(commandBuffer unsafe.Pointer, sourceImage IImage, batchNormalizationState ICNNBatchNormalizationState, destinationImage IImage) {
	objc.Send[objc.ID](c_.ID, objc.Sel("encodeToCommandBuffer:sourceImage:batchNormalizationState:destinationImage:"), commandBuffer, sourceImage, batchNormalizationState, destinationImage)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnbatchnormalization/2942610-encodebatch
func (c_ CNNBatchNormalization) EncodeBatch() {
	objc.Send[objc.ID](c_.ID, objc.Sel("encodeBatch"))
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnbatchnormalization/2942610-encodebatchtocommandbuffer
func (c_ CNNBatchNormalization) EncodeBatchToCommandBufferSourceImagesBatchNormalizationStateDestinationImages(commandBuffer unsafe.Pointer, sourceImages ImageBatch /* not a class type */, batchNormalizationState ICNNBatchNormalizationState, destinationImages ImageBatch /* not a class type */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("encodeBatchToCommandBuffer:sourceImages:batchNormalizationState:destinationImages:"), commandBuffer, sourceImages, batchNormalizationState, destinationImages)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnbatchnormalization/2953965-reloadgammaandbeta
func (c_ CNNBatchNormalization) ReloadGammaAndBeta() {
	objc.Send[objc.ID](c_.ID, objc.Sel("reloadGammaAndBeta"))
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnbatchnormalization/2953965-reloadgammaandbetawithcommandbuf
func (c_ CNNBatchNormalization) ReloadGammaAndBetaWithCommandBufferGammaAndBetaState(commandBuffer unsafe.Pointer, gammaAndBetaState ICNNNormalizationGammaAndBetaState) {
	objc.Send[objc.ID](c_.ID, objc.Sel("reloadGammaAndBetaWithCommandBuffer:gammaAndBetaState:"), commandBuffer, gammaAndBetaState)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnbatchnormalization/2954874-resultstate
func (c_ CNNBatchNormalization) ResultState() {
	objc.Send[objc.ID](c_.ID, objc.Sel("resultState"))
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnbatchnormalization/2954874-resultstateforsourceimage
func (c_ CNNBatchNormalization) ResultStateForSourceImageSourceStatesDestinationImage(sourceImage IImage, sourceStates unsafe.Pointer, destinationImage IImage) ICNNBatchNormalizationState {
	rv := objc.Send[CNNBatchNormalizationState](c_.ID, objc.Sel("resultStateForSourceImage:sourceStates:destinationImage:"), sourceImage, sourceStates, destinationImage)
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnbatchnormalization/2954875-temporaryresultstate
func (c_ CNNBatchNormalization) TemporaryResultState() {
	objc.Send[objc.ID](c_.ID, objc.Sel("temporaryResultState"))
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnbatchnormalization/2954875-temporaryresultstateforcommandbu
func (c_ CNNBatchNormalization) TemporaryResultStateForCommandBufferSourceImageSourceStatesDestinationImage(commandBuffer unsafe.Pointer, sourceImage IImage, sourceStates unsafe.Pointer, destinationImage IImage) ICNNBatchNormalizationState {
	rv := objc.Send[CNNBatchNormalizationState](c_.ID, objc.Sel("temporaryResultStateForCommandBuffer:sourceImage:sourceStates:destinationImage:"), commandBuffer, sourceImage, sourceStates, destinationImage)
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnbatchnormalization/2976464-reloadgammaandbetafromdatasource
func (c_ CNNBatchNormalization) ReloadGammaAndBetaFromDataSource() {
	objc.Send[objc.ID](c_.ID, objc.Sel("reloadGammaAndBetaFromDataSource"))
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnbatchnormalization/3002358-reloadmeanandvariancefromdatasou
func (c_ CNNBatchNormalization) ReloadMeanAndVarianceFromDataSource() {
	objc.Send[objc.ID](c_.ID, objc.Sel("reloadMeanAndVarianceFromDataSource"))
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnbatchnormalization/3002359-reloadmeanandvariance
func (c_ CNNBatchNormalization) ReloadMeanAndVariance() {
	objc.Send[objc.ID](c_.ID, objc.Sel("reloadMeanAndVariance"))
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnbatchnormalization/3002359-reloadmeanandvariancewithcommand
func (c_ CNNBatchNormalization) ReloadMeanAndVarianceWithCommandBufferMeanAndVarianceState(commandBuffer unsafe.Pointer, meanAndVarianceState ICNNNormalizationMeanAndVarianceState) {
	objc.Send[objc.ID](c_.ID, objc.Sel("reloadMeanAndVarianceWithCommandBuffer:meanAndVarianceState:"), commandBuffer, meanAndVarianceState)
}







// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnbatchnormalization/2942599-epsilon
func (c_ CNNBatchNormalization) Epsilon() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](c_.ID, objc.Sel("epsilon"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnbatchnormalization/2942599-epsilon
func (c_ CNNBatchNormalization) SetEpsilon(value objectivec.IObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setEpsilon:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnbatchnormalization/2942604-numberoffeaturechannels
func (c_ CNNBatchNormalization) NumberOfFeatureChannels() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](c_.ID, objc.Sel("numberOfFeatureChannels"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnbatchnormalization/2942604-numberoffeaturechannels
func (c_ CNNBatchNormalization) SetNumberOfFeatureChannels(value objectivec.IObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setNumberOfFeatureChannels:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnbatchnormalization/2953967-datasource
func (c_ CNNBatchNormalization) DataSource() CNNBatchNormalizationDataSource get /* not a class type */ {
	rv := objc.Send[objc.ID](c_.ID, objc.Sel("dataSource"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnbatchnormalization/2953967-datasource
func (c_ CNNBatchNormalization) SetDataSource(value CNNBatchNormalizationDataSource get /* not a class type */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setDataSource:"), value)
}







