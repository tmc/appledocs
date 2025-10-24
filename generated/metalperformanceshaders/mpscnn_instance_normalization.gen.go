// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)





// The class instance for the [CNNInstanceNormalization] class.
var (
	CNNInstanceNormalizationClass     _CNNInstanceNormalizationClass
	CNNInstanceNormalizationClassOnce sync.Once
)

func getCNNInstanceNormalizationClass() _CNNInstanceNormalizationClass {
	CNNInstanceNormalizationClassOnce.Do(func() {
		CNNInstanceNormalizationClass = _CNNInstanceNormalizationClass{objc.GetClass("MPSCNNInstanceNormalization")}
	})
	return CNNInstanceNormalizationClass
}

type _CNNInstanceNormalizationClass struct {
	class objc.Class
}





// An interface definition for the [CNNInstanceNormalization] class.
type ICNNInstanceNormalization interface {
	ICNNKernel
	

	// properties:
	Epsilon() objectivec.IObject
	SetEpsilon(value objectivec.IObject)
	DataSource() CNNInstanceNormalizationDataSource get /* not a class type */
	SetDataSource(value CNNInstanceNormalizationDataSource get /* not a class type */)


	

	// methods:
	ResultState()
	ResultStateForSourceImageSourceStatesDestinationImage(sourceImage IImage, sourceStates unsafe.Pointer, destinationImage IImage) ICNNInstanceNormalizationGradientState
	ReloadGammaAndBeta()
	ReloadGammaAndBetaWithCommandBufferGammaAndBetaState(commandBuffer unsafe.Pointer, gammaAndBetaState ICNNNormalizationGammaAndBetaState)
	TemporaryResultState()
	TemporaryResultStateForCommandBufferSourceImageSourceStatesDestinationImage(commandBuffer unsafe.Pointer, sourceImage IImage, sourceStates unsafe.Pointer, destinationImage IImage) ICNNInstanceNormalizationGradientState
	ReloadGammaAndBetaFromDataSource()


}





// Alloc allocates a new instance without initialization.
func (cc _CNNInstanceNormalizationClass) Alloc() CNNInstanceNormalization {
	rv := objc.Send[CNNInstanceNormalization](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _CNNInstanceNormalizationClass) New() CNNInstanceNormalization {
	rv := objc.Send[CNNInstanceNormalization](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CNNInstanceNormalization) Init() CNNInstanceNormalization {
	rv := objc.Send[CNNInstanceNormalization](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CNNInstanceNormalization) Autorelease() CNNInstanceNormalization {
	rv := objc.Send[CNNInstanceNormalization](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCNNInstanceNormalization creates a new CNNInstanceNormalization instance.
func NewCNNInstanceNormalization() CNNInstanceNormalization {
	return getCNNInstanceNormalizationClass().New()
}





// An instance normalization kernel.


// An instance normalization kernel.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNInstanceNormalization
type CNNInstanceNormalization struct {
	CNNKernel
}

// CNNInstanceNormalizationFrom constructs a [CNNInstanceNormalization] from an unsafe.Pointer.
//
// An instance normalization kernel.
func CNNInstanceNormalizationFrom(ptr unsafe.Pointer) CNNInstanceNormalization {
	return CNNInstanceNormalization{
		CNNKernel: CNNKernelFrom(ptr),
	}
}






// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnninstancenormalization/2947948-initwithcoder
func NewCNNInstanceNormalizationWithCoderDevice(aDecoder foundation.Coder, device unsafe.Pointer) CNNInstanceNormalization {
	instance := getCNNInstanceNormalizationClass().Alloc()
	rv := objc.Send[CNNInstanceNormalization](instance.ID, objc.Sel("initWithCoder:device:"), aDecoder, device)
	rv.Autorelease()
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnninstancenormalization/2947947-initwithdevice
func NewCNNInstanceNormalizationWithDeviceDataSource(device unsafe.Pointer, dataSource unsafe.Pointer) CNNInstanceNormalization {
	instance := getCNNInstanceNormalizationClass().Alloc()
	rv := objc.Send[CNNInstanceNormalization](instance.ID, objc.Sel("initWithDevice:dataSource:"), device, dataSource)
	rv.Autorelease()
	return rv
}

















// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnninstancenormalization/2947959-resultstate
func (c_ CNNInstanceNormalization) ResultState() {
	objc.Send[objc.ID](c_.ID, objc.Sel("resultState"))
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnninstancenormalization/2947959-resultstateforsourceimage
func (c_ CNNInstanceNormalization) ResultStateForSourceImageSourceStatesDestinationImage(sourceImage IImage, sourceStates unsafe.Pointer, destinationImage IImage) ICNNInstanceNormalizationGradientState {
	rv := objc.Send[CNNInstanceNormalizationGradientState](c_.ID, objc.Sel("resultStateForSourceImage:sourceStates:destinationImage:"), sourceImage, sourceStates, destinationImage)
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnninstancenormalization/2953921-reloadgammaandbeta
func (c_ CNNInstanceNormalization) ReloadGammaAndBeta() {
	objc.Send[objc.ID](c_.ID, objc.Sel("reloadGammaAndBeta"))
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnninstancenormalization/2953921-reloadgammaandbetawithcommandbuf
func (c_ CNNInstanceNormalization) ReloadGammaAndBetaWithCommandBufferGammaAndBetaState(commandBuffer unsafe.Pointer, gammaAndBetaState ICNNNormalizationGammaAndBetaState) {
	objc.Send[objc.ID](c_.ID, objc.Sel("reloadGammaAndBetaWithCommandBuffer:gammaAndBetaState:"), commandBuffer, gammaAndBetaState)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnninstancenormalization/2956160-temporaryresultstate
func (c_ CNNInstanceNormalization) TemporaryResultState() {
	objc.Send[objc.ID](c_.ID, objc.Sel("temporaryResultState"))
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnninstancenormalization/2956160-temporaryresultstateforcommandbu
func (c_ CNNInstanceNormalization) TemporaryResultStateForCommandBufferSourceImageSourceStatesDestinationImage(commandBuffer unsafe.Pointer, sourceImage IImage, sourceStates unsafe.Pointer, destinationImage IImage) ICNNInstanceNormalizationGradientState {
	rv := objc.Send[CNNInstanceNormalizationGradientState](c_.ID, objc.Sel("temporaryResultStateForCommandBuffer:sourceImage:sourceStates:destinationImage:"), commandBuffer, sourceImage, sourceStates, destinationImage)
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnninstancenormalization/2976471-reloadgammaandbetafromdatasource
func (c_ CNNInstanceNormalization) ReloadGammaAndBetaFromDataSource() {
	objc.Send[objc.ID](c_.ID, objc.Sel("reloadGammaAndBetaFromDataSource"))
}







// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnninstancenormalization/2947943-epsilon
func (c_ CNNInstanceNormalization) Epsilon() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](c_.ID, objc.Sel("epsilon"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnninstancenormalization/2947943-epsilon
func (c_ CNNInstanceNormalization) SetEpsilon(value objectivec.IObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setEpsilon:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnninstancenormalization/2953927-datasource
func (c_ CNNInstanceNormalization) DataSource() CNNInstanceNormalizationDataSource get /* not a class type */ {
	rv := objc.Send[objc.ID](c_.ID, objc.Sel("dataSource"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnninstancenormalization/2953927-datasource
func (c_ CNNInstanceNormalization) SetDataSource(value CNNInstanceNormalizationDataSource get /* not a class type */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setDataSource:"), value)
}







