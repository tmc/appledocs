// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MPSCNNGroupNormalization */


/* debug [class_header]: Header for MPSCNNGroupNormalization */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for CNNGroupNormalization */
// An interface definition for the [CNNGroupNormalization] class.
type ICNNGroupNormalization interface {
	ICNNKernel
	
/* debug [class_interface_properties]: Properties for CNNGroupNormalization */
	// properties:
	DataSource() CNNGroupNormalizationDataSource get /* not a class type */
	SetDataSource(value CNNGroupNormalizationDataSource get /* not a class type */)
	Epsilon() objectivec.IObject
	SetEpsilon(value objectivec.IObject)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for CNNGroupNormalization */
	// methods:
	ReloadGammaAndBetaFromDataSource()
	ReloadGammaAndBeta()
	ReloadGammaAndBetaWithCommandBufferGammaAndBetaState(commandBuffer unsafe.Pointer, gammaAndBetaState ICNNNormalizationGammaAndBetaState)
	ResultState()
	ResultStateForSourceImageSourceStatesDestinationImage(sourceImage IImage, sourceStates unsafe.Pointer, destinationImage IImage) ICNNGroupNormalizationGradientState
	TemporaryResultState()
	TemporaryResultStateForCommandBufferSourceImageSourceStatesDestinationImage(commandBuffer unsafe.Pointer, sourceImage IImage, sourceStates unsafe.Pointer, destinationImage IImage) ICNNGroupNormalizationGradientState
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for CNNGroupNormalization */
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for CNNGroupNormalization */


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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for CNNGroupNormalization */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnngroupnormalization/3152536-initwithcoder
func NewCNNGroupNormalizationWithCoderDevice(aDecoder Coder /* not a class type */, device unsafe.Pointer) CNNGroupNormalization {
	instance := getCNNGroupNormalizationClass().Alloc()
	rv := objc.Send[CNNGroupNormalization](instance.ID, objc.Sel("initWithCoder:device:"), aDecoder, device)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewCNNGroupNormalizationWithCoderDevice */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnngroupnormalization/3152537-initwithdevice
func NewCNNGroupNormalizationWithDeviceDataSource(device unsafe.Pointer, dataSource unsafe.Pointer) CNNGroupNormalization {
	instance := getCNNGroupNormalizationClass().Alloc()
	rv := objc.Send[CNNGroupNormalization](instance.ID, objc.Sel("initWithDevice:dataSource:"), device, dataSource)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewCNNGroupNormalizationWithDeviceDataSource */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for CNNGroupNormalization */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for CNNGroupNormalization */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for CNNGroupNormalization */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnngroupnormalization/3152538-reloadgammaandbetafromdatasource
func (c_ CNNGroupNormalization) ReloadGammaAndBetaFromDataSource() {
	objc.Send[objc.ID](c_.ID, objc.Sel("reloadGammaAndBetaFromDataSource"))
}/* debug [instance_methods/method]: ReloadGammaAndBetaFromDataSource */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnngroupnormalization/3152539-reloadgammaandbeta
func (c_ CNNGroupNormalization) ReloadGammaAndBeta() {
	objc.Send[objc.ID](c_.ID, objc.Sel("reloadGammaAndBeta"))
}/* debug [instance_methods/method]: ReloadGammaAndBeta */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnngroupnormalization/3152539-reloadgammaandbetawithcommandbuf
func (c_ CNNGroupNormalization) ReloadGammaAndBetaWithCommandBufferGammaAndBetaState(commandBuffer unsafe.Pointer, gammaAndBetaState ICNNNormalizationGammaAndBetaState) {
	objc.Send[objc.ID](c_.ID, objc.Sel("reloadGammaAndBetaWithCommandBuffer:gammaAndBetaState:"), commandBuffer, gammaAndBetaState)
}/* debug [instance_methods/method]: ReloadGammaAndBetaWithCommandBufferGammaAndBetaState */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnngroupnormalization/3152540-resultstate
func (c_ CNNGroupNormalization) ResultState() {
	objc.Send[objc.ID](c_.ID, objc.Sel("resultState"))
}/* debug [instance_methods/method]: ResultState */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnngroupnormalization/3152540-resultstateforsourceimage
func (c_ CNNGroupNormalization) ResultStateForSourceImageSourceStatesDestinationImage(sourceImage IImage, sourceStates unsafe.Pointer, destinationImage IImage) ICNNGroupNormalizationGradientState {
	rv := objc.Send[CNNGroupNormalizationGradientState](c_.ID, objc.Sel("resultStateForSourceImage:sourceStates:destinationImage:"), sourceImage, sourceStates, destinationImage)
	return rv
}/* debug [instance_methods/method]: ResultStateForSourceImageSourceStatesDestinationImage */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnngroupnormalization/3152541-temporaryresultstate
func (c_ CNNGroupNormalization) TemporaryResultState() {
	objc.Send[objc.ID](c_.ID, objc.Sel("temporaryResultState"))
}/* debug [instance_methods/method]: TemporaryResultState */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnngroupnormalization/3152541-temporaryresultstateforcommandbu
func (c_ CNNGroupNormalization) TemporaryResultStateForCommandBufferSourceImageSourceStatesDestinationImage(commandBuffer unsafe.Pointer, sourceImage IImage, sourceStates unsafe.Pointer, destinationImage IImage) ICNNGroupNormalizationGradientState {
	rv := objc.Send[CNNGroupNormalizationGradientState](c_.ID, objc.Sel("temporaryResultStateForCommandBuffer:sourceImage:sourceStates:destinationImage:"), commandBuffer, sourceImage, sourceStates, destinationImage)
	return rv
}/* debug [instance_methods/method]: TemporaryResultStateForCommandBufferSourceImageSourceStatesDestinationImage */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for CNNGroupNormalization */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnngroupnormalization/3152534-datasource
func (c_ CNNGroupNormalization) DataSource() CNNGroupNormalizationDataSource get /* not a class type */ {
	rv := objc.Send[objc.ID](c_.ID, objc.Sel("dataSource"))
	return rv
}/* debug [instance_properties/getter]: dataSource */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnngroupnormalization/3152534-datasource
func (c_ CNNGroupNormalization) SetDataSource(value CNNGroupNormalizationDataSource get /* not a class type */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setDataSource:"), value)
}/* debug [instance_properties/setter]: dataSource */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnngroupnormalization/3152535-epsilon
func (c_ CNNGroupNormalization) Epsilon() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](c_.ID, objc.Sel("epsilon"))
	return rv
}/* debug [instance_properties/getter]: epsilon */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnngroupnormalization/3152535-epsilon
func (c_ CNNGroupNormalization) SetEpsilon(value objectivec.IObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setEpsilon:"), value)
}/* debug [instance_properties/setter]: epsilon */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MPSCNNGroupNormalization */


