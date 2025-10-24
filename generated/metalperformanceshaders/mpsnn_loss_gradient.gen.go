// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MPSNNLossGradient */


/* debug [class_header]: Header for MPSNNLossGradient */
// The class instance for the [LossGradient] class.
var (
	LossGradientClass     _LossGradientClass
	LossGradientClassOnce sync.Once
)

func getLossGradientClass() _LossGradientClass {
	LossGradientClassOnce.Do(func() {
		LossGradientClass = _LossGradientClass{objc.GetClass("MPSNNLossGradient")}
	})
	return LossGradientClass
}

type _LossGradientClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for LossGradient */
// An interface definition for the [LossGradient] class.
type ILossGradient interface {
	ICNNBinaryKernel
	
/* debug [class_interface_properties]: Properties for LossGradient */
	// properties:
	ComputeLabelGradients() objectivec.IObject
	SetComputeLabelGradients(value objectivec.IObject)
	Delta() objectivec.IObject
	SetDelta(value objectivec.IObject)
	Epsilon() objectivec.IObject
	SetEpsilon(value objectivec.IObject)
	LabelSmoothing() objectivec.IObject
	SetLabelSmoothing(value objectivec.IObject)
	LossType() CNNLossType get /* not a class type */
	SetLossType(value CNNLossType get /* not a class type */)
	NumberOfClasses() objectivec.IObject
	SetNumberOfClasses(value objectivec.IObject)
	ReductionType() CNNReductionType get /* not a class type */
	SetReductionType(value CNNReductionType get /* not a class type */)
	Weight() objectivec.IObject
	SetWeight(value objectivec.IObject)
	ReduceAcrossBatch() objectivec.IObject
	SetReduceAcrossBatch(value objectivec.IObject)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for LossGradient */
	// methods:
	EncodeBatch()
	EncodeBatchToCommandBufferSourceGradientsSourceImagesLabelsWeightsSourceStates(commandBuffer unsafe.Pointer, sourceGradients ImageBatch /* not a class type */, sourceImages ImageBatch /* not a class type */, labels ImageBatch /* not a class type */, weights ImageBatch /* not a class type */, sourceStates StateBatch /* not a class type */) ImageBatch /* not a class type */
	EncodeBatchToCommandBufferSourceGradientsSourceImagesLabelsWeightsSourceStatesDestinationGradients(commandBuffer unsafe.Pointer, sourceGradients ImageBatch /* not a class type */, sourceImages ImageBatch /* not a class type */, labels ImageBatch /* not a class type */, weights ImageBatch /* not a class type */, sourceStates StateBatch /* not a class type */, destinationGradients ImageBatch /* not a class type */)
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for LossGradient */
// Alloc allocates a new instance without initialization.
func (lc _LossGradientClass) Alloc() LossGradient {
	rv := objc.Send[LossGradient](objc.ID(lc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (lc _LossGradientClass) New() LossGradient {
	rv := objc.Send[LossGradient](objc.ID(lc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (l_ LossGradient) Init() LossGradient {
	rv := objc.Send[LossGradient](l_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (l_ LossGradient) Autorelease() LossGradient {
	rv := objc.Send[LossGradient](l_.ID, objc.Sel("autorelease"))
	return rv
}

// NewLossGradient creates a new LossGradient instance.
func NewLossGradient() LossGradient {
	return getLossGradientClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for LossGradient */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNLossGradient
type LossGradient struct {
	CNNBinaryKernel
}

// LossGradientFrom constructs a [LossGradient] from an unsafe.Pointer.
func LossGradientFrom(ptr unsafe.Pointer) LossGradient {
	return LossGradient{
		CNNBinaryKernel: CNNBinaryKernelFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for LossGradient */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnlossgradient/3131816-initwithcoder
func NewLossGradientWithCoderDevice(aDecoder foundation.Coder, device unsafe.Pointer) LossGradient {
	instance := getLossGradientClass().Alloc()
	rv := objc.Send[LossGradient](instance.ID, objc.Sel("initWithCoder:device:"), aDecoder, device)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewLossGradientWithCoderDevice */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnlossgradient/3131817-initwithdevice
func NewLossGradientWithDeviceLossDescriptor(device unsafe.Pointer, lossDescriptor ICNNLossDescriptor) LossGradient {
	instance := getLossGradientClass().Alloc()
	rv := objc.Send[LossGradient](instance.ID, objc.Sel("initWithDevice:lossDescriptor:"), device, lossDescriptor)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewLossGradientWithDeviceLossDescriptor */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for LossGradient */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for LossGradient */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for LossGradient */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnlossgradient/3131813-encodebatch
func (l_ LossGradient) EncodeBatch() {
	objc.Send[objc.ID](l_.ID, objc.Sel("encodeBatch"))
}/* debug [instance_methods/method]: EncodeBatch */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnlossgradient/3131813-encodebatchtocommandbuffer
func (l_ LossGradient) EncodeBatchToCommandBufferSourceGradientsSourceImagesLabelsWeightsSourceStates(commandBuffer unsafe.Pointer, sourceGradients ImageBatch /* not a class type */, sourceImages ImageBatch /* not a class type */, labels ImageBatch /* not a class type */, weights ImageBatch /* not a class type */, sourceStates StateBatch /* not a class type */) ImageBatch /* not a class type */ {
	rv := objc.Send[ImageBatch](l_.ID, objc.Sel("encodeBatchToCommandBuffer:sourceGradients:sourceImages:labels:weights:sourceStates:"), commandBuffer, sourceGradients, sourceImages, labels, weights, sourceStates)
	return rv
}/* debug [instance_methods/method]: EncodeBatchToCommandBufferSourceGradientsSourceImagesLabelsWeightsSourceStates */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnlossgradient/3131814-encodebatchtocommandbuffer
func (l_ LossGradient) EncodeBatchToCommandBufferSourceGradientsSourceImagesLabelsWeightsSourceStatesDestinationGradients(commandBuffer unsafe.Pointer, sourceGradients ImageBatch /* not a class type */, sourceImages ImageBatch /* not a class type */, labels ImageBatch /* not a class type */, weights ImageBatch /* not a class type */, sourceStates StateBatch /* not a class type */, destinationGradients ImageBatch /* not a class type */) {
	objc.Send[objc.ID](l_.ID, objc.Sel("encodeBatchToCommandBuffer:sourceGradients:sourceImages:labels:weights:sourceStates:destinationGradients:"), commandBuffer, sourceGradients, sourceImages, labels, weights, sourceStates, destinationGradients)
}/* debug [instance_methods/method]: EncodeBatchToCommandBufferSourceGradientsSourceImagesLabelsWeightsSourceStatesDestinationGradients */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for LossGradient */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnlossgradient/3131811-computelabelgradients
func (l_ LossGradient) ComputeLabelGradients() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](l_.ID, objc.Sel("computeLabelGradients"))
	return rv
}/* debug [instance_properties/getter]: computeLabelGradients */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnlossgradient/3131811-computelabelgradients
func (l_ LossGradient) SetComputeLabelGradients(value objectivec.IObject) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setComputeLabelGradients:"), value)
}/* debug [instance_properties/setter]: computeLabelGradients */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnlossgradient/3131812-delta
func (l_ LossGradient) Delta() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](l_.ID, objc.Sel("delta"))
	return rv
}/* debug [instance_properties/getter]: delta */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnlossgradient/3131812-delta
func (l_ LossGradient) SetDelta(value objectivec.IObject) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setDelta:"), value)
}/* debug [instance_properties/setter]: delta */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnlossgradient/3131815-epsilon
func (l_ LossGradient) Epsilon() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](l_.ID, objc.Sel("epsilon"))
	return rv
}/* debug [instance_properties/getter]: epsilon */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnlossgradient/3131815-epsilon
func (l_ LossGradient) SetEpsilon(value objectivec.IObject) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setEpsilon:"), value)
}/* debug [instance_properties/setter]: epsilon */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnlossgradient/3131818-labelsmoothing
func (l_ LossGradient) LabelSmoothing() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](l_.ID, objc.Sel("labelSmoothing"))
	return rv
}/* debug [instance_properties/getter]: labelSmoothing */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnlossgradient/3131818-labelsmoothing
func (l_ LossGradient) SetLabelSmoothing(value objectivec.IObject) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setLabelSmoothing:"), value)
}/* debug [instance_properties/setter]: labelSmoothing */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnlossgradient/3131819-losstype
func (l_ LossGradient) LossType() CNNLossType get /* not a class type */ {
	rv := objc.Send[objc.ID](l_.ID, objc.Sel("lossType"))
	return rv
}/* debug [instance_properties/getter]: lossType */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnlossgradient/3131819-losstype
func (l_ LossGradient) SetLossType(value CNNLossType get /* not a class type */) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setLossType:"), value)
}/* debug [instance_properties/setter]: lossType */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnlossgradient/3131820-numberofclasses
func (l_ LossGradient) NumberOfClasses() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](l_.ID, objc.Sel("numberOfClasses"))
	return rv
}/* debug [instance_properties/getter]: numberOfClasses */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnlossgradient/3131820-numberofclasses
func (l_ LossGradient) SetNumberOfClasses(value objectivec.IObject) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setNumberOfClasses:"), value)
}/* debug [instance_properties/setter]: numberOfClasses */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnlossgradient/3131821-reductiontype
func (l_ LossGradient) ReductionType() CNNReductionType get /* not a class type */ {
	rv := objc.Send[objc.ID](l_.ID, objc.Sel("reductionType"))
	return rv
}/* debug [instance_properties/getter]: reductionType */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnlossgradient/3131821-reductiontype
func (l_ LossGradient) SetReductionType(value CNNReductionType get /* not a class type */) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setReductionType:"), value)
}/* debug [instance_properties/setter]: reductionType */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnlossgradient/3131822-weight
func (l_ LossGradient) Weight() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](l_.ID, objc.Sel("weight"))
	return rv
}/* debug [instance_properties/getter]: weight */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnlossgradient/3131822-weight
func (l_ LossGradient) SetWeight(value objectivec.IObject) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setWeight:"), value)
}/* debug [instance_properties/setter]: weight */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnlossgradient/3547986-reduceacrossbatch
func (l_ LossGradient) ReduceAcrossBatch() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](l_.ID, objc.Sel("reduceAcrossBatch"))
	return rv
}/* debug [instance_properties/getter]: reduceAcrossBatch */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnlossgradient/3547986-reduceacrossbatch
func (l_ LossGradient) SetReduceAcrossBatch(value objectivec.IObject) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setReduceAcrossBatch:"), value)
}/* debug [instance_properties/setter]: reduceAcrossBatch */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MPSNNLossGradient */


