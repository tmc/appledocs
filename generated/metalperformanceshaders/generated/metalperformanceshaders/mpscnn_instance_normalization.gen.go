// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MPSCNNInstanceNormalization */


/* debug [class_header]: Header for MPSCNNInstanceNormalization */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for CNNInstanceNormalization */
// An interface definition for the [CNNInstanceNormalization] class.
type ICNNInstanceNormalization interface {
	ICNNKernel
	
/* debug [class_interface_properties]: Properties for CNNInstanceNormalization */
	// properties:
	Epsilon() objectivec.IObject
	SetEpsilon(value objectivec.IObject)
	DataSource() CNNInstanceNormalizationDataSource get /* not a class type */
	SetDataSource(value CNNInstanceNormalizationDataSource get /* not a class type */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for CNNInstanceNormalization */
	// methods:
	ResultState()
	ResultStateForSourceImageSourceStatesDestinationImage(sourceImage IImage, sourceStates unsafe.Pointer, destinationImage IImage) ICNNInstanceNormalizationGradientState
	ReloadGammaAndBeta()
	ReloadGammaAndBetaWithCommandBufferGammaAndBetaState(commandBuffer unsafe.Pointer, gammaAndBetaState ICNNNormalizationGammaAndBetaState)
	TemporaryResultState()
	TemporaryResultStateForCommandBufferSourceImageSourceStatesDestinationImage(commandBuffer unsafe.Pointer, sourceImage IImage, sourceStates unsafe.Pointer, destinationImage IImage) ICNNInstanceNormalizationGradientState
	ReloadGammaAndBetaFromDataSource()
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for CNNInstanceNormalization */
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for CNNInstanceNormalization */
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for CNNInstanceNormalization */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnninstancenormalization/2947948-initwithcoder
func NewCNNInstanceNormalizationWithCoderDevice(aDecoder Coder /* not a class type */, device unsafe.Pointer) CNNInstanceNormalization {
	instance := getCNNInstanceNormalizationClass().Alloc()
	rv := objc.Send[CNNInstanceNormalization](instance.ID, objc.Sel("initWithCoder:device:"), aDecoder, device)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewCNNInstanceNormalizationWithCoderDevice */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnninstancenormalization/2947947-initwithdevice
func NewCNNInstanceNormalizationWithDeviceDataSource(device unsafe.Pointer, dataSource unsafe.Pointer) CNNInstanceNormalization {
	instance := getCNNInstanceNormalizationClass().Alloc()
	rv := objc.Send[CNNInstanceNormalization](instance.ID, objc.Sel("initWithDevice:dataSource:"), device, dataSource)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewCNNInstanceNormalizationWithDeviceDataSource */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for CNNInstanceNormalization */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for CNNInstanceNormalization */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for CNNInstanceNormalization */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnninstancenormalization/2947959-resultstate
func (c_ CNNInstanceNormalization) ResultState() {
	objc.Send[objc.ID](c_.ID, objc.Sel("resultState"))
}/* debug [instance_methods/method]: ResultState */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnninstancenormalization/2947959-resultstateforsourceimage
func (c_ CNNInstanceNormalization) ResultStateForSourceImageSourceStatesDestinationImage(sourceImage IImage, sourceStates unsafe.Pointer, destinationImage IImage) ICNNInstanceNormalizationGradientState {
	rv := objc.Send[CNNInstanceNormalizationGradientState](c_.ID, objc.Sel("resultStateForSourceImage:sourceStates:destinationImage:"), sourceImage, sourceStates, destinationImage)
	return rv
}/* debug [instance_methods/method]: ResultStateForSourceImageSourceStatesDestinationImage */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnninstancenormalization/2953921-reloadgammaandbeta
func (c_ CNNInstanceNormalization) ReloadGammaAndBeta() {
	objc.Send[objc.ID](c_.ID, objc.Sel("reloadGammaAndBeta"))
}/* debug [instance_methods/method]: ReloadGammaAndBeta */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnninstancenormalization/2953921-reloadgammaandbetawithcommandbuf
func (c_ CNNInstanceNormalization) ReloadGammaAndBetaWithCommandBufferGammaAndBetaState(commandBuffer unsafe.Pointer, gammaAndBetaState ICNNNormalizationGammaAndBetaState) {
	objc.Send[objc.ID](c_.ID, objc.Sel("reloadGammaAndBetaWithCommandBuffer:gammaAndBetaState:"), commandBuffer, gammaAndBetaState)
}/* debug [instance_methods/method]: ReloadGammaAndBetaWithCommandBufferGammaAndBetaState */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnninstancenormalization/2956160-temporaryresultstate
func (c_ CNNInstanceNormalization) TemporaryResultState() {
	objc.Send[objc.ID](c_.ID, objc.Sel("temporaryResultState"))
}/* debug [instance_methods/method]: TemporaryResultState */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnninstancenormalization/2956160-temporaryresultstateforcommandbu
func (c_ CNNInstanceNormalization) TemporaryResultStateForCommandBufferSourceImageSourceStatesDestinationImage(commandBuffer unsafe.Pointer, sourceImage IImage, sourceStates unsafe.Pointer, destinationImage IImage) ICNNInstanceNormalizationGradientState {
	rv := objc.Send[CNNInstanceNormalizationGradientState](c_.ID, objc.Sel("temporaryResultStateForCommandBuffer:sourceImage:sourceStates:destinationImage:"), commandBuffer, sourceImage, sourceStates, destinationImage)
	return rv
}/* debug [instance_methods/method]: TemporaryResultStateForCommandBufferSourceImageSourceStatesDestinationImage */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnninstancenormalization/2976471-reloadgammaandbetafromdatasource
func (c_ CNNInstanceNormalization) ReloadGammaAndBetaFromDataSource() {
	objc.Send[objc.ID](c_.ID, objc.Sel("reloadGammaAndBetaFromDataSource"))
}/* debug [instance_methods/method]: ReloadGammaAndBetaFromDataSource */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for CNNInstanceNormalization */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnninstancenormalization/2947943-epsilon
func (c_ CNNInstanceNormalization) Epsilon() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](c_.ID, objc.Sel("epsilon"))
	return rv
}/* debug [instance_properties/getter]: epsilon */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnninstancenormalization/2947943-epsilon
func (c_ CNNInstanceNormalization) SetEpsilon(value objectivec.IObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setEpsilon:"), value)
}/* debug [instance_properties/setter]: epsilon */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnninstancenormalization/2953927-datasource
func (c_ CNNInstanceNormalization) DataSource() CNNInstanceNormalizationDataSource get /* not a class type */ {
	rv := objc.Send[objc.ID](c_.ID, objc.Sel("dataSource"))
	return rv
}/* debug [instance_properties/getter]: dataSource */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnninstancenormalization/2953927-datasource
func (c_ CNNInstanceNormalization) SetDataSource(value CNNInstanceNormalizationDataSource get /* not a class type */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setDataSource:"), value)
}/* debug [instance_properties/setter]: dataSource */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MPSCNNInstanceNormalization */


