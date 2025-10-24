// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)





// The class instance for the [CNNGroupNormalization] class.
var (
	CNNGroupNormalizationClass     _CNNGroupNormalizationClass
	CNNGroupNormalizationClassOnce sync.Once
)

func getCNNGroupNormalizationClass() _CNNGroupNormalizationClass {
	CNNGroupNormalizationClassOnce.Do(func() {
		CNNGroupNormalizationClass = _CNNGroupNormalizationClass{objc.GetClass("MPSCNNGroupNormalization")}
	})
	return CNNGroupNormalizationClass
}

type _CNNGroupNormalizationClass struct {
	class objc.Class
}





// An interface definition for the [CNNGroupNormalization] class.
type ICNNGroupNormalization interface {
	ICNNKernel
	

	// properties:
	DataSource() CNNGroupNormalizationDataSource get /* not a class type */
	SetDataSource(value CNNGroupNormalizationDataSource get /* not a class type */)
	Epsilon() objectivec.IObject
	SetEpsilon(value objectivec.IObject)


	

	// methods:
	ReloadGammaAndBetaFromDataSource()
	ReloadGammaAndBeta()
	ReloadGammaAndBetaWithCommandBufferGammaAndBetaState(commandBuffer unsafe.Pointer, gammaAndBetaState ICNNNormalizationGammaAndBetaState)
	ResultState()
	ResultStateForSourceImageSourceStatesDestinationImage(sourceImage IImage, sourceStates unsafe.Pointer, destinationImage IImage) ICNNGroupNormalizationGradientState
	TemporaryResultState()
	TemporaryResultStateForCommandBufferSourceImageSourceStatesDestinationImage(commandBuffer unsafe.Pointer, sourceImage IImage, sourceStates unsafe.Pointer, destinationImage IImage) ICNNGroupNormalizationGradientState


}





// Alloc allocates a new instance without initialization.
func (cc _CNNGroupNormalizationClass) Alloc() CNNGroupNormalization {
	rv := objc.Send[CNNGroupNormalization](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _CNNGroupNormalizationClass) New() CNNGroupNormalization {
	rv := objc.Send[CNNGroupNormalization](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CNNGroupNormalization) Init() CNNGroupNormalization {
	rv := objc.Send[CNNGroupNormalization](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CNNGroupNormalization) Autorelease() CNNGroupNormalization {
	rv := objc.Send[CNNGroupNormalization](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCNNGroupNormalization creates a new CNNGroupNormalization instance.
func NewCNNGroupNormalization() CNNGroupNormalization {
	return getCNNGroupNormalizationClass().New()
}







// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNGroupNormalization
type CNNGroupNormalization struct {
	CNNKernel
}

// CNNGroupNormalizationFrom constructs a [CNNGroupNormalization] from an unsafe.Pointer.
func CNNGroupNormalizationFrom(ptr unsafe.Pointer) CNNGroupNormalization {
	return CNNGroupNormalization{
		CNNKernel: CNNKernelFrom(ptr),
	}
}






// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnngroupnormalization/3152536-initwithcoder
func NewCNNGroupNormalizationWithCoderDevice(aDecoder foundation.Coder, device unsafe.Pointer) CNNGroupNormalization {
	instance := getCNNGroupNormalizationClass().Alloc()
	rv := objc.Send[CNNGroupNormalization](instance.ID, objc.Sel("initWithCoder:device:"), aDecoder, device)
	rv.Autorelease()
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnngroupnormalization/3152537-initwithdevice
func NewCNNGroupNormalizationWithDeviceDataSource(device unsafe.Pointer, dataSource unsafe.Pointer) CNNGroupNormalization {
	instance := getCNNGroupNormalizationClass().Alloc()
	rv := objc.Send[CNNGroupNormalization](instance.ID, objc.Sel("initWithDevice:dataSource:"), device, dataSource)
	rv.Autorelease()
	return rv
}

















// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnngroupnormalization/3152538-reloadgammaandbetafromdatasource
func (c_ CNNGroupNormalization) ReloadGammaAndBetaFromDataSource() {
	objc.Send[objc.ID](c_.ID, objc.Sel("reloadGammaAndBetaFromDataSource"))
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnngroupnormalization/3152539-reloadgammaandbeta
func (c_ CNNGroupNormalization) ReloadGammaAndBeta() {
	objc.Send[objc.ID](c_.ID, objc.Sel("reloadGammaAndBeta"))
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnngroupnormalization/3152539-reloadgammaandbetawithcommandbuf
func (c_ CNNGroupNormalization) ReloadGammaAndBetaWithCommandBufferGammaAndBetaState(commandBuffer unsafe.Pointer, gammaAndBetaState ICNNNormalizationGammaAndBetaState) {
	objc.Send[objc.ID](c_.ID, objc.Sel("reloadGammaAndBetaWithCommandBuffer:gammaAndBetaState:"), commandBuffer, gammaAndBetaState)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnngroupnormalization/3152540-resultstate
func (c_ CNNGroupNormalization) ResultState() {
	objc.Send[objc.ID](c_.ID, objc.Sel("resultState"))
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnngroupnormalization/3152540-resultstateforsourceimage
func (c_ CNNGroupNormalization) ResultStateForSourceImageSourceStatesDestinationImage(sourceImage IImage, sourceStates unsafe.Pointer, destinationImage IImage) ICNNGroupNormalizationGradientState {
	rv := objc.Send[CNNGroupNormalizationGradientState](c_.ID, objc.Sel("resultStateForSourceImage:sourceStates:destinationImage:"), sourceImage, sourceStates, destinationImage)
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnngroupnormalization/3152541-temporaryresultstate
func (c_ CNNGroupNormalization) TemporaryResultState() {
	objc.Send[objc.ID](c_.ID, objc.Sel("temporaryResultState"))
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnngroupnormalization/3152541-temporaryresultstateforcommandbu
func (c_ CNNGroupNormalization) TemporaryResultStateForCommandBufferSourceImageSourceStatesDestinationImage(commandBuffer unsafe.Pointer, sourceImage IImage, sourceStates unsafe.Pointer, destinationImage IImage) ICNNGroupNormalizationGradientState {
	rv := objc.Send[CNNGroupNormalizationGradientState](c_.ID, objc.Sel("temporaryResultStateForCommandBuffer:sourceImage:sourceStates:destinationImage:"), commandBuffer, sourceImage, sourceStates, destinationImage)
	return rv
}







// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnngroupnormalization/3152534-datasource
func (c_ CNNGroupNormalization) DataSource() CNNGroupNormalizationDataSource get /* not a class type */ {
	rv := objc.Send[objc.ID](c_.ID, objc.Sel("dataSource"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnngroupnormalization/3152534-datasource
func (c_ CNNGroupNormalization) SetDataSource(value CNNGroupNormalizationDataSource get /* not a class type */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setDataSource:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnngroupnormalization/3152535-epsilon
func (c_ CNNGroupNormalization) Epsilon() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](c_.ID, objc.Sel("epsilon"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnngroupnormalization/3152535-epsilon
func (c_ CNNGroupNormalization) SetEpsilon(value objectivec.IObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setEpsilon:"), value)
}







