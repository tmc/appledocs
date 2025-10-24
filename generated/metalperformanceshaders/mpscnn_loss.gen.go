// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MPSCNNLoss */


/* debug [class_header]: Header for MPSCNNLoss */
// The class instance for the [CNNLoss] class.
var (
	CNNLossClass     _CNNLossClass
	CNNLossClassOnce sync.Once
)

func getCNNLossClass() _CNNLossClass {
	CNNLossClassOnce.Do(func() {
		CNNLossClass = _CNNLossClass{objc.GetClass("MPSCNNLoss")}
	})
	return CNNLossClass
}

type _CNNLossClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for CNNLoss */
// An interface definition for the [CNNLoss] class.
type ICNNLoss interface {
	ICNNKernel
	
/* debug [class_interface_properties]: Properties for CNNLoss */
	// properties:
	LabelSmoothing() objectivec.IObject
	SetLabelSmoothing(value objectivec.IObject)
	LossType() CNNLossType get /* not a class type */
	SetLossType(value CNNLossType get /* not a class type */)
	Delta() objectivec.IObject
	SetDelta(value objectivec.IObject)
	ReductionType() CNNReductionType get /* not a class type */
	SetReductionType(value CNNReductionType get /* not a class type */)
	Epsilon() objectivec.IObject
	SetEpsilon(value objectivec.IObject)
	Weight() objectivec.IObject
	SetWeight(value objectivec.IObject)
	NumberOfClasses() objectivec.IObject
	SetNumberOfClasses(value objectivec.IObject)
	ReduceAcrossBatch() objectivec.IObject
	SetReduceAcrossBatch(value objectivec.IObject)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for CNNLoss */
	// methods:
	Encode()
	EncodeToCommandBufferSourceImageLabels(commandBuffer unsafe.Pointer, sourceImage IImage, labels ICNNLossLabels) IImage
	EncodeBatchToCommandBufferSourceImagesLabels(commandBuffer unsafe.Pointer, sourceImage ImageBatch /* not a class type */, labels CNNLossLabelsBatch /* not a class type */) ImageBatch /* not a class type */
	EncodeToCommandBufferSourceImageLabelsDestinationImage(commandBuffer unsafe.Pointer, sourceImage IImage, labels ICNNLossLabels, destinationImage IImage)
	EncodeBatchToCommandBufferSourceImagesLabelsDestinationImages(commandBuffer unsafe.Pointer, sourceImage ImageBatch /* not a class type */, labels CNNLossLabelsBatch /* not a class type */, destinationImage ImageBatch /* not a class type */)
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for CNNLoss */
// Alloc allocates a new instance without initialization.
func (cc _CNNLossClass) Alloc() CNNLoss {
	rv := objc.Send[CNNLoss](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _CNNLossClass) New() CNNLoss {
	rv := objc.Send[CNNLoss](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CNNLoss) Init() CNNLoss {
	rv := objc.Send[CNNLoss](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CNNLoss) Autorelease() CNNLoss {
	rv := objc.Send[CNNLoss](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCNNLoss creates a new CNNLoss instance.
func NewCNNLoss() CNNLoss {
	return getCNNLossClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for CNNLoss */
// A kernel that computes the loss and loss gradient between specified predictions and labels.


// A kernel that computes the loss and loss gradient between specified predictions and labels.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNLoss
type CNNLoss struct {
	CNNKernel
}

// CNNLossFrom constructs a [CNNLoss] from an unsafe.Pointer.
//
// A kernel that computes the loss and loss gradient between specified predictions and labels.
func CNNLossFrom(ptr unsafe.Pointer) CNNLoss {
	return CNNLoss{
		CNNKernel: CNNKernelFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for CNNLoss */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnloss/2942379-initwithcoder
func NewCNNLossWithCoderDevice(aDecoder foundation.Coder, device unsafe.Pointer) CNNLoss {
	instance := getCNNLossClass().Alloc()
	rv := objc.Send[CNNLoss](instance.ID, objc.Sel("initWithCoder:device:"), aDecoder, device)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewCNNLossWithCoderDevice */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnloss/2942377-initwithdevice
func NewCNNLossWithDeviceLossDescriptor(device unsafe.Pointer, lossDescriptor ICNNLossDescriptor) CNNLoss {
	instance := getCNNLossClass().Alloc()
	rv := objc.Send[CNNLoss](instance.ID, objc.Sel("initWithDevice:lossDescriptor:"), device, lossDescriptor)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewCNNLossWithDeviceLossDescriptor */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for CNNLoss */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for CNNLoss */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for CNNLoss */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnloss/2951838-encode
func (c_ CNNLoss) Encode() {
	objc.Send[objc.ID](c_.ID, objc.Sel("encode"))
}/* debug [instance_methods/method]: Encode */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnloss/2951838-encodetocommandbuffer
func (c_ CNNLoss) EncodeToCommandBufferSourceImageLabels(commandBuffer unsafe.Pointer, sourceImage IImage, labels ICNNLossLabels) IImage {
	rv := objc.Send[Image](c_.ID, objc.Sel("encodeToCommandBuffer:sourceImage:labels:"), commandBuffer, sourceImage, labels)
	return rv
}/* debug [instance_methods/method]: EncodeToCommandBufferSourceImageLabels */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnloss/2951839-encodebatchtocommandbuffer
func (c_ CNNLoss) EncodeBatchToCommandBufferSourceImagesLabels(commandBuffer unsafe.Pointer, sourceImage ImageBatch /* not a class type */, labels CNNLossLabelsBatch /* not a class type */) ImageBatch /* not a class type */ {
	rv := objc.Send[ImageBatch](c_.ID, objc.Sel("encodeBatchToCommandBuffer:sourceImages:labels:"), commandBuffer, sourceImage, labels)
	return rv
}/* debug [instance_methods/method]: EncodeBatchToCommandBufferSourceImagesLabels */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnloss/2951843-encodetocommandbuffer
func (c_ CNNLoss) EncodeToCommandBufferSourceImageLabelsDestinationImage(commandBuffer unsafe.Pointer, sourceImage IImage, labels ICNNLossLabels, destinationImage IImage) {
	objc.Send[objc.ID](c_.ID, objc.Sel("encodeToCommandBuffer:sourceImage:labels:destinationImage:"), commandBuffer, sourceImage, labels, destinationImage)
}/* debug [instance_methods/method]: EncodeToCommandBufferSourceImageLabelsDestinationImage */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnloss/2951846-encodebatchtocommandbuffer
func (c_ CNNLoss) EncodeBatchToCommandBufferSourceImagesLabelsDestinationImages(commandBuffer unsafe.Pointer, sourceImage ImageBatch /* not a class type */, labels CNNLossLabelsBatch /* not a class type */, destinationImage ImageBatch /* not a class type */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("encodeBatchToCommandBuffer:sourceImages:labels:destinationImages:"), commandBuffer, sourceImage, labels, destinationImage)
}/* debug [instance_methods/method]: EncodeBatchToCommandBufferSourceImagesLabelsDestinationImages */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for CNNLoss */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnloss/2942358-labelsmoothing
func (c_ CNNLoss) LabelSmoothing() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](c_.ID, objc.Sel("labelSmoothing"))
	return rv
}/* debug [instance_properties/getter]: labelSmoothing */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnloss/2942358-labelsmoothing
func (c_ CNNLoss) SetLabelSmoothing(value objectivec.IObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setLabelSmoothing:"), value)
}/* debug [instance_properties/setter]: labelSmoothing */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnloss/2942359-losstype
func (c_ CNNLoss) LossType() CNNLossType get /* not a class type */ {
	rv := objc.Send[objc.ID](c_.ID, objc.Sel("lossType"))
	return rv
}/* debug [instance_properties/getter]: lossType */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnloss/2942359-losstype
func (c_ CNNLoss) SetLossType(value CNNLossType get /* not a class type */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setLossType:"), value)
}/* debug [instance_properties/setter]: lossType */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnloss/2942360-delta
func (c_ CNNLoss) Delta() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](c_.ID, objc.Sel("delta"))
	return rv
}/* debug [instance_properties/getter]: delta */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnloss/2942360-delta
func (c_ CNNLoss) SetDelta(value objectivec.IObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setDelta:"), value)
}/* debug [instance_properties/setter]: delta */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnloss/2942365-reductiontype
func (c_ CNNLoss) ReductionType() CNNReductionType get /* not a class type */ {
	rv := objc.Send[objc.ID](c_.ID, objc.Sel("reductionType"))
	return rv
}/* debug [instance_properties/getter]: reductionType */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnloss/2942365-reductiontype
func (c_ CNNLoss) SetReductionType(value CNNReductionType get /* not a class type */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setReductionType:"), value)
}/* debug [instance_properties/setter]: reductionType */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnloss/2942371-epsilon
func (c_ CNNLoss) Epsilon() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](c_.ID, objc.Sel("epsilon"))
	return rv
}/* debug [instance_properties/getter]: epsilon */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnloss/2942371-epsilon
func (c_ CNNLoss) SetEpsilon(value objectivec.IObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setEpsilon:"), value)
}/* debug [instance_properties/setter]: epsilon */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnloss/2942387-weight
func (c_ CNNLoss) Weight() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](c_.ID, objc.Sel("weight"))
	return rv
}/* debug [instance_properties/getter]: weight */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnloss/2942387-weight
func (c_ CNNLoss) SetWeight(value objectivec.IObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setWeight:"), value)
}/* debug [instance_properties/setter]: weight */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnloss/2942389-numberofclasses
func (c_ CNNLoss) NumberOfClasses() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](c_.ID, objc.Sel("numberOfClasses"))
	return rv
}/* debug [instance_properties/getter]: numberOfClasses */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnloss/2942389-numberofclasses
func (c_ CNNLoss) SetNumberOfClasses(value objectivec.IObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setNumberOfClasses:"), value)
}/* debug [instance_properties/setter]: numberOfClasses */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnloss/3547981-reduceacrossbatch
func (c_ CNNLoss) ReduceAcrossBatch() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](c_.ID, objc.Sel("reduceAcrossBatch"))
	return rv
}/* debug [instance_properties/getter]: reduceAcrossBatch */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnloss/3547981-reduceacrossbatch
func (c_ CNNLoss) SetReduceAcrossBatch(value objectivec.IObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setReduceAcrossBatch:"), value)
}/* debug [instance_properties/setter]: reduceAcrossBatch */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MPSCNNLoss */


