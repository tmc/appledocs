// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MPSNNForwardLoss */


/* debug [class_header]: Header for MPSNNForwardLoss */
// The class instance for the [ForwardLoss] class.
var (
	ForwardLossClass     _ForwardLossClass
	ForwardLossClassOnce sync.Once
)

func getForwardLossClass() _ForwardLossClass {
	ForwardLossClassOnce.Do(func() {
		ForwardLossClass = _ForwardLossClass{objc.GetClass("MPSNNForwardLoss")}
	})
	return ForwardLossClass
}

type _ForwardLossClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for ForwardLoss */
// An interface definition for the [ForwardLoss] class.
type IForwardLoss interface {
	ICNNKernel
	
/* debug [class_interface_properties]: Properties for ForwardLoss */
	// properties:
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

	
/* debug [class_interface_methods]: Methods for ForwardLoss */
	// methods:
	EncodeBatch()
	EncodeBatchToCommandBufferSourceImagesLabelsWeightsDestinationStatesDestinationImages(commandBuffer unsafe.Pointer, sourceImages ImageBatch /* not a class type */, labels ImageBatch /* not a class type */, weights ImageBatch /* not a class type */, destinationStates StateBatch /* not a class type */, destinationImages ImageBatch /* not a class type */)
	EncodeBatchToCommandBufferSourceImagesLabelsWeightsDestinationStatesDestinationStateIsTemporary(commandBuffer unsafe.Pointer, sourceImages ImageBatch /* not a class type */, labels ImageBatch /* not a class type */, weights ImageBatch /* not a class type */, outStates StateBatch /* not a class type */, isTemporary bool) ImageBatch /* not a class type */
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for ForwardLoss */
// Alloc allocates a new instance without initialization.
func (fc _ForwardLossClass) Alloc() ForwardLoss {
	rv := objc.Send[ForwardLoss](objc.ID(fc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (fc _ForwardLossClass) New() ForwardLoss {
	rv := objc.Send[ForwardLoss](objc.ID(fc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (f_ ForwardLoss) Init() ForwardLoss {
	rv := objc.Send[ForwardLoss](f_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (f_ ForwardLoss) Autorelease() ForwardLoss {
	rv := objc.Send[ForwardLoss](f_.ID, objc.Sel("autorelease"))
	return rv
}

// NewForwardLoss creates a new ForwardLoss instance.
func NewForwardLoss() ForwardLoss {
	return getForwardLossClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for ForwardLoss */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNForwardLoss
type ForwardLoss struct {
	CNNKernel
}

// ForwardLossFrom constructs a [ForwardLoss] from an unsafe.Pointer.
func ForwardLossFrom(ptr unsafe.Pointer) ForwardLoss {
	return ForwardLoss{
		CNNKernel: CNNKernelFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for ForwardLoss */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnforwardloss/3131801-initwithcoder
func NewForwardLossWithCoderDevice(aDecoder Coder /* not a class type */, device unsafe.Pointer) ForwardLoss {
	instance := getForwardLossClass().Alloc()
	rv := objc.Send[ForwardLoss](instance.ID, objc.Sel("initWithCoder:device:"), aDecoder, device)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewForwardLossWithCoderDevice */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnforwardloss/3131802-initwithdevice
func NewForwardLossWithDeviceLossDescriptor(device unsafe.Pointer, lossDescriptor ICNNLossDescriptor) ForwardLoss {
	instance := getForwardLossClass().Alloc()
	rv := objc.Send[ForwardLoss](instance.ID, objc.Sel("initWithDevice:lossDescriptor:"), device, lossDescriptor)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewForwardLossWithDeviceLossDescriptor */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for ForwardLoss */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for ForwardLoss */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for ForwardLoss */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnforwardloss/3131798-encodebatch
func (f_ ForwardLoss) EncodeBatch() {
	objc.Send[objc.ID](f_.ID, objc.Sel("encodeBatch"))
}/* debug [instance_methods/method]: EncodeBatch */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnforwardloss/3131798-encodebatchtocommandbuffer
func (f_ ForwardLoss) EncodeBatchToCommandBufferSourceImagesLabelsWeightsDestinationStatesDestinationImages(commandBuffer unsafe.Pointer, sourceImages ImageBatch /* not a class type */, labels ImageBatch /* not a class type */, weights ImageBatch /* not a class type */, destinationStates StateBatch /* not a class type */, destinationImages ImageBatch /* not a class type */) {
	objc.Send[objc.ID](f_.ID, objc.Sel("encodeBatchToCommandBuffer:sourceImages:labels:weights:destinationStates:destinationImages:"), commandBuffer, sourceImages, labels, weights, destinationStates, destinationImages)
}/* debug [instance_methods/method]: EncodeBatchToCommandBufferSourceImagesLabelsWeightsDestinationStatesDestinationImages */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnforwardloss/3131799-encodebatchtocommandbuffer
func (f_ ForwardLoss) EncodeBatchToCommandBufferSourceImagesLabelsWeightsDestinationStatesDestinationStateIsTemporary(commandBuffer unsafe.Pointer, sourceImages ImageBatch /* not a class type */, labels ImageBatch /* not a class type */, weights ImageBatch /* not a class type */, outStates StateBatch /* not a class type */, isTemporary bool) ImageBatch /* not a class type */ {
	rv := objc.Send[ImageBatch](f_.ID, objc.Sel("encodeBatchToCommandBuffer:sourceImages:labels:weights:destinationStates:destinationStateIsTemporary:"), commandBuffer, sourceImages, labels, weights, outStates, isTemporary)
	return rv
}/* debug [instance_methods/method]: EncodeBatchToCommandBufferSourceImagesLabelsWeightsDestinationStatesDestinationStateIsTemporary */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for ForwardLoss */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnforwardloss/3131797-delta
func (f_ ForwardLoss) Delta() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](f_.ID, objc.Sel("delta"))
	return rv
}/* debug [instance_properties/getter]: delta */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnforwardloss/3131797-delta
func (f_ ForwardLoss) SetDelta(value objectivec.IObject) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setDelta:"), value)
}/* debug [instance_properties/setter]: delta */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnforwardloss/3131800-epsilon
func (f_ ForwardLoss) Epsilon() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](f_.ID, objc.Sel("epsilon"))
	return rv
}/* debug [instance_properties/getter]: epsilon */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnforwardloss/3131800-epsilon
func (f_ ForwardLoss) SetEpsilon(value objectivec.IObject) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setEpsilon:"), value)
}/* debug [instance_properties/setter]: epsilon */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnforwardloss/3131803-labelsmoothing
func (f_ ForwardLoss) LabelSmoothing() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](f_.ID, objc.Sel("labelSmoothing"))
	return rv
}/* debug [instance_properties/getter]: labelSmoothing */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnforwardloss/3131803-labelsmoothing
func (f_ ForwardLoss) SetLabelSmoothing(value objectivec.IObject) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setLabelSmoothing:"), value)
}/* debug [instance_properties/setter]: labelSmoothing */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnforwardloss/3131804-losstype
func (f_ ForwardLoss) LossType() CNNLossType get /* not a class type */ {
	rv := objc.Send[objc.ID](f_.ID, objc.Sel("lossType"))
	return rv
}/* debug [instance_properties/getter]: lossType */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnforwardloss/3131804-losstype
func (f_ ForwardLoss) SetLossType(value CNNLossType get /* not a class type */) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setLossType:"), value)
}/* debug [instance_properties/setter]: lossType */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnforwardloss/3131805-numberofclasses
func (f_ ForwardLoss) NumberOfClasses() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](f_.ID, objc.Sel("numberOfClasses"))
	return rv
}/* debug [instance_properties/getter]: numberOfClasses */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnforwardloss/3131805-numberofclasses
func (f_ ForwardLoss) SetNumberOfClasses(value objectivec.IObject) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setNumberOfClasses:"), value)
}/* debug [instance_properties/setter]: numberOfClasses */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnforwardloss/3131806-reductiontype
func (f_ ForwardLoss) ReductionType() CNNReductionType get /* not a class type */ {
	rv := objc.Send[objc.ID](f_.ID, objc.Sel("reductionType"))
	return rv
}/* debug [instance_properties/getter]: reductionType */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnforwardloss/3131806-reductiontype
func (f_ ForwardLoss) SetReductionType(value CNNReductionType get /* not a class type */) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setReductionType:"), value)
}/* debug [instance_properties/setter]: reductionType */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnforwardloss/3131807-weight
func (f_ ForwardLoss) Weight() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](f_.ID, objc.Sel("weight"))
	return rv
}/* debug [instance_properties/getter]: weight */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnforwardloss/3131807-weight
func (f_ ForwardLoss) SetWeight(value objectivec.IObject) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setWeight:"), value)
}/* debug [instance_properties/setter]: weight */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnforwardloss/3547985-reduceacrossbatch
func (f_ ForwardLoss) ReduceAcrossBatch() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](f_.ID, objc.Sel("reduceAcrossBatch"))
	return rv
}/* debug [instance_properties/getter]: reduceAcrossBatch */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnforwardloss/3547985-reduceacrossbatch
func (f_ ForwardLoss) SetReduceAcrossBatch(value objectivec.IObject) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setReduceAcrossBatch:"), value)
}/* debug [instance_properties/setter]: reduceAcrossBatch */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MPSNNForwardLoss */


