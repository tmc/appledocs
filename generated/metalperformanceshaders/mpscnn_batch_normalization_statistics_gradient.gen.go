// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MPSCNNBatchNormalizationStatisticsGradient */


/* debug [class_header]: Header for MPSCNNBatchNormalizationStatisticsGradient */
// The class instance for the [CNNBatchNormalizationStatisticsGradient] class.
var (
	CNNBatchNormalizationStatisticsGradientClass     _CNNBatchNormalizationStatisticsGradientClass
	CNNBatchNormalizationStatisticsGradientClassOnce sync.Once
)

func getCNNBatchNormalizationStatisticsGradientClass() _CNNBatchNormalizationStatisticsGradientClass {
	CNNBatchNormalizationStatisticsGradientClassOnce.Do(func() {
		CNNBatchNormalizationStatisticsGradientClass = _CNNBatchNormalizationStatisticsGradientClass{objc.GetClass("MPSCNNBatchNormalizationStatisticsGradient")}
	})
	return CNNBatchNormalizationStatisticsGradientClass
}

type _CNNBatchNormalizationStatisticsGradientClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for CNNBatchNormalizationStatisticsGradient */
// An interface definition for the [CNNBatchNormalizationStatisticsGradient] class.
type ICNNBatchNormalizationStatisticsGradient interface {
	ICNNGradientKernel
	
/* debug [class_interface_properties]: Properties for CNNBatchNormalizationStatisticsGradient */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for CNNBatchNormalizationStatisticsGradient */
	// methods:
	EncodeBatch()
	EncodeBatchToCommandBufferSourceGradientsSourceImagesBatchNormalizationState(commandBuffer unsafe.Pointer, sourceGradients ImageBatch /* not a class type */, sourceImages ImageBatch /* not a class type */, batchNormalizationState ICNNBatchNormalizationState)
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for CNNBatchNormalizationStatisticsGradient */
// Alloc allocates a new instance without initialization.
func (cc _CNNBatchNormalizationStatisticsGradientClass) Alloc() CNNBatchNormalizationStatisticsGradient {
	rv := objc.Send[CNNBatchNormalizationStatisticsGradient](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _CNNBatchNormalizationStatisticsGradientClass) New() CNNBatchNormalizationStatisticsGradient {
	rv := objc.Send[CNNBatchNormalizationStatisticsGradient](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CNNBatchNormalizationStatisticsGradient) Init() CNNBatchNormalizationStatisticsGradient {
	rv := objc.Send[CNNBatchNormalizationStatisticsGradient](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CNNBatchNormalizationStatisticsGradient) Autorelease() CNNBatchNormalizationStatisticsGradient {
	rv := objc.Send[CNNBatchNormalizationStatisticsGradient](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCNNBatchNormalizationStatisticsGradient creates a new CNNBatchNormalizationStatisticsGradient instance.
func NewCNNBatchNormalizationStatisticsGradient() CNNBatchNormalizationStatisticsGradient {
	return getCNNBatchNormalizationStatisticsGradientClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for CNNBatchNormalizationStatisticsGradient */
// An object that stores the gradient of the loss function with respect to the batch statistics and batch normalization weights.


// An object that stores the gradient of the loss function with respect to the batch statistics and batch normalization weights.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNBatchNormalizationStatisticsGradient
type CNNBatchNormalizationStatisticsGradient struct {
	CNNGradientKernel
}

// CNNBatchNormalizationStatisticsGradientFrom constructs a [CNNBatchNormalizationStatisticsGradient] from an unsafe.Pointer.
//
// An object that stores the gradient of the loss function with respect to the batch statistics and batch normalization weights.
func CNNBatchNormalizationStatisticsGradientFrom(ptr unsafe.Pointer) CNNBatchNormalizationStatisticsGradient {
	return CNNBatchNormalizationStatisticsGradient{
		CNNGradientKernel: CNNGradientKernelFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for CNNBatchNormalizationStatisticsGradient */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnbatchnormalizationstatisticsgradient/3013774-initwithcoder
func NewCNNBatchNormalizationStatisticsGradientWithCoderDevice(aDecoder foundation.Coder, device unsafe.Pointer) CNNBatchNormalizationStatisticsGradient {
	instance := getCNNBatchNormalizationStatisticsGradientClass().Alloc()
	rv := objc.Send[CNNBatchNormalizationStatisticsGradient](instance.ID, objc.Sel("initWithCoder:device:"), aDecoder, device)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewCNNBatchNormalizationStatisticsGradientWithCoderDevice */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnbatchnormalizationstatisticsgradient/3013775-initwithdevice
func NewCNNBatchNormalizationStatisticsGradientWithDeviceFusedNeuronDescriptor(device unsafe.Pointer, fusedNeuronDescriptor INeuronDescriptor) CNNBatchNormalizationStatisticsGradient {
	instance := getCNNBatchNormalizationStatisticsGradientClass().Alloc()
	rv := objc.Send[CNNBatchNormalizationStatisticsGradient](instance.ID, objc.Sel("initWithDevice:fusedNeuronDescriptor:"), device, fusedNeuronDescriptor)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewCNNBatchNormalizationStatisticsGradientWithDeviceFusedNeuronDescriptor */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for CNNBatchNormalizationStatisticsGradient */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for CNNBatchNormalizationStatisticsGradient */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for CNNBatchNormalizationStatisticsGradient */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnbatchnormalizationstatisticsgradient/2953964-encodebatch
func (c_ CNNBatchNormalizationStatisticsGradient) EncodeBatch() {
	objc.Send[objc.ID](c_.ID, objc.Sel("encodeBatch"))
}/* debug [instance_methods/method]: EncodeBatch */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnbatchnormalizationstatisticsgradient/2953964-encodebatchtocommandbuffer
func (c_ CNNBatchNormalizationStatisticsGradient) EncodeBatchToCommandBufferSourceGradientsSourceImagesBatchNormalizationState(commandBuffer unsafe.Pointer, sourceGradients ImageBatch /* not a class type */, sourceImages ImageBatch /* not a class type */, batchNormalizationState ICNNBatchNormalizationState) {
	objc.Send[objc.ID](c_.ID, objc.Sel("encodeBatchToCommandBuffer:sourceGradients:sourceImages:batchNormalizationState:"), commandBuffer, sourceGradients, sourceImages, batchNormalizationState)
}/* debug [instance_methods/method]: EncodeBatchToCommandBufferSourceGradientsSourceImagesBatchNormalizationState */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for CNNBatchNormalizationStatisticsGradient */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MPSCNNBatchNormalizationStatisticsGradient */


